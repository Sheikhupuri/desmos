package profile

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/simulation"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
)

const (
	OpWeightMsgSaveProfile   = simulation.OpWeightMsgSaveProfile
	OpWeightMsgDeleteProfile = simulation.OpWeightMsgDeleteProfile
	DefaultGasValue          = simulation.DefaultGasValue
	EventTypeProfileSaved    = types.EventTypeProfileSaved
	EventTypeProfileDeleted  = types.EventTypeProfileDeleted
	AttributeProfileMoniker  = types.AttributeProfileMoniker
	AttributeProfileCreator  = types.AttributeProfileCreator
	ModuleName               = types.ModuleName
	RouterKey                = types.RouterKey
	StoreKey                 = types.StoreKey
	MinNameSurnameLength     = types.MinNameSurnameLength
	MaxNameSurnameLength     = types.MaxNameSurnameLength
	MinMonikerLength         = types.MinMonikerLength
	MaxMonikerLength         = types.MaxMonikerLength
	MaxBioLength             = types.MaxBioLength
	ActionSaveProfile        = types.ActionSaveProfile
	ActionDeleteProfile      = types.ActionDeleteProfile
	QuerierRoute             = types.QuerierRoute
	QueryProfile             = types.QueryProfile
	QueryProfiles            = types.QueryProfiles
)

var (
	// functions aliases
	NewKeeper                = keeper.NewKeeper
	NewQuerier               = keeper.NewQuerier
	RegisterInvariants       = keeper.RegisterInvariants
	AllInvariants            = keeper.AllInvariants
	ValidProfileInvariant    = keeper.ValidProfileInvariant
	NewHandler               = keeper.NewHandler
	DecodeStore              = simulation.DecodeStore
	SimulateMsgSaveProfile   = simulation.SimulateMsgSaveProfile
	SimulateMsgDeleteProfile = simulation.SimulateMsgDeleteProfile
	RandomProfileData        = simulation.RandomProfileData
	RandomProfile            = simulation.RandomProfile
	RandomMoniker            = simulation.RandomMoniker
	RandomName               = simulation.RandomName
	RandomSurname            = simulation.RandomSurname
	RandomBio                = simulation.RandomBio
	RandomProfilePic         = simulation.RandomProfilePic
	RandomProfileCover       = simulation.RandomProfileCover
	GetSimAccount            = simulation.GetSimAccount
	WeightedOperations       = simulation.WeightedOperations
	RandomizedGenState       = simulation.RandomizedGenState
	NewProfile               = types.NewProfile
	NewPictures              = types.NewPictures
	NewGenesisState          = types.NewGenesisState
	DefaultGenesisState      = types.DefaultGenesisState
	ValidateGenesis          = types.ValidateGenesis
	RegisterCodec            = types.RegisterCodec
	NewMsgSaveProfile        = types.NewMsgSaveProfile
	NewMsgDeleteProfile      = types.NewMsgDeleteProfile
	ProfileStoreKey          = types.ProfileStoreKey
	MonikerStoreKey          = types.MonikerStoreKey

	// variable aliases
	ModuleCdc          = types.ModuleCdc
	TxHashRegEx        = types.TxHashRegEx
	URIRegEx           = types.URIRegEx
	ProfileStorePrefix = types.ProfileStorePrefix
	MonikerStorePrefix = types.MonikerStorePrefix
)

type (
	Profile          = types.Profile
	Profiles         = types.Profiles
	Pictures         = types.Pictures
	GenesisState     = types.GenesisState
	MsgSaveProfile   = types.MsgSaveProfile
	MsgDeleteProfile = types.MsgDeleteProfile
	Keeper           = keeper.Keeper
	ProfileData      = simulation.ProfileData
)
