package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {

		case types.QueryPost:
			return queryPost(ctx, path[1:], req, keeper)

		case types.QueryPosts:
			return queryPosts(ctx, req, keeper)

		case types.QueryPollAnswers:
			return queryPollAnswers(ctx, path[1:], req, keeper)

		default:
			return nil, fmt.Errorf("unknown post query endpoint")
		}
	}
}

// getPostResponse allows to get a PostQueryResponse from the given post retrieving the other information
// using the given Context and Keeper.
func getPostResponse(ctx sdk.Context, keeper Keeper, post types.Post) types.PostQueryResponse {
	// Get the likes
	postLikes := keeper.GetPostReactions(ctx, post.PostID)
	if postLikes == nil {
		postLikes = types.Reactions{}
	}

	// Get the children
	childrenIDs := keeper.GetPostChildrenIDs(ctx, post.PostID)
	if childrenIDs == nil {
		childrenIDs = types.PostIDs{}
	}

	//Get the poll answers if poll exist
	var answers []types.AnswersDetails
	if post.PollData != nil {
		answers = keeper.GetPollAnswers(ctx, post.PostID)
	}

	// Crete the response object
	return types.NewPostResponse(post, answers, postLikes, childrenIDs)
}

// queryPost handles the request to get a post having a specific id
func queryPost(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	id, err := types.ParsePostID(path[0])
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Invalid post id: %s", path[0]))
	}

	post, found := keeper.GetPost(ctx, id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Post with id %s not found", id))
	}

	postResponse := getPostResponse(ctx, keeper, post)
	bz, err2 := codec.MarshalJSONIndent(keeper.Cdc, &postResponse)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// queryPosts handles the request of listing all the posts that satisfy a specific filter
func queryPosts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var params types.QueryPostsParams

	err := keeper.Cdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	posts := keeper.GetPostsFiltered(ctx, params)

	postResponses := make([]types.PostQueryResponse, len(posts))
	for index, post := range posts {
		postResponses[index] = getPostResponse(ctx, keeper, post)
	}

	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &postResponses)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

//queryPollAnswers handles the request to get poll answers related to a post with given id
func queryPollAnswers(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	id, err := types.ParsePostID(path[0])

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Invalid post id: %s", path[0]))
	}

	post, found := keeper.GetPost(ctx, id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Post with id %s not found", id))
	}

	if post.PollData == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Post with id %s has no poll associated", id))
	}

	pollAnswers := keeper.GetPollAnswers(ctx, id)

	pollAnswersResponse := types.PollAnswersQueryResponse{
		PostID:         id,
		AnswersDetails: pollAnswers,
	}
	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &pollAnswersResponse)

	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}