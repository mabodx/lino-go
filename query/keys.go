package query

import (
	"bytes"
	"encoding/hex"
	"strconv"

	crypto "github.com/tendermint/go-crypto"
)

const (
	MainKVStoreKey      = "main"
	AccountKVStoreKey   = "account"
	PostKVStoreKey      = "post"
	ValidatorKVStoreKey = "validator"
	GlobalKVStoreKey    = "global"
	VoteKVStoreKey      = "vote"
	InfraKVStoreKey     = "infra"
	DeveloperKVStoreKey = "developer"
	ParamKVStoreKey     = "param"
	ProposalKVStoreKey  = "proposal"
)

var (
	KeySeparator = "/"

	accountInfoSubstore              = []byte{0x00}
	accountBankSubstore              = []byte{0x01}
	accountMetaSubstore              = []byte{0x02}
	accountFollowerSubstore          = []byte{0x03}
	accountFollowingSubstore         = []byte{0x04}
	accountRewardSubstore            = []byte{0x05}
	accountPendingStakeQueueSubstore = []byte{0x06}
	accountRelationshipSubstore      = []byte{0x07}
	accountBalanceHistorySubstore    = []byte{0x08}
	accountGrantPubKeySubstore       = []byte{0x09}

	postInfoSubStore           = []byte{0x00} // SubStore for all post info
	postMetaSubStore           = []byte{0x01} // SubStore for all post mata info
	postLikeSubStore           = []byte{0x02} // SubStore for all like to post
	postReportOrUpvoteSubStore = []byte{0x03} // SubStore for all like to post
	postCommentSubStore        = []byte{0x04} // SubStore for all comments
	postViewsSubStore          = []byte{0x05} // SubStore for all views
	postDonationsSubStore      = []byte{0x06} // SubStore for all donations

	validatorSubstore     = []byte{0x00}
	validatorListSubstore = []byte{0x01}

	delegationSubstore    = []byte{0x00}
	voterSubstore         = []byte{0x01}
	voteSubstore          = []byte{0x02}
	referenceListSubStore = []byte{0x03}
	delegateeSubStore     = []byte{0x04}

	developerSubstore     = []byte{0x00}
	developerListSubstore = []byte{0x01}

	infraProviderSubstore     = []byte{0x00}
	infraProviderListSubstore = []byte{0x01}

	proposalSubstore     = []byte{0x00}
	proposalListSubStore = []byte{0x01}

	allocationParamSubStore              = []byte{0x00} // SubStore for allocation
	infraInternalAllocationParamSubStore = []byte{0x01} // SubStore for infrat internal allocation
	evaluateOfContentValueParamSubStore  = []byte{0x02} // Substore for evaluate of content value
	developerParamSubStore               = []byte{0x03} // Substore for developer param
	voteParamSubStore                    = []byte{0x04} // Substore for vote param
	proposalParamSubStore                = []byte{0x05} // Substore for proposal param
	validatorParamSubStore               = []byte{0x06} // Substore for validator param
	coinDayParamSubStore                 = []byte{0x07} // Substore for coin day param
	bandwidthParamSubStore               = []byte{0x08} // Substore for bandwidth param
	accountParamSubstore                 = []byte{0x09} // Substore for account param
	postParamSubStore                    = []byte{0x10} // Substore for evaluate of content value
)

func getHexSubstringAfterKeySeparator(key []byte) string {
	return hex.EncodeToString(key[bytes.Index(key, []byte(KeySeparator)):])
}

func getSubstringAfterKeySeparator(key []byte) string {
	return string(key[bytes.Index(key, []byte(KeySeparator)):])
}

func getSubstringAfterSubstore(key []byte) string {
	return string(key[1:])
}

//
// account related
//
func getAccountInfoKey(accKey string) []byte {
	return append(accountInfoSubstore, accKey...)
}

func getAccountBankKey(accKey string) []byte {
	return append(accountBankSubstore, accKey...)
}

func getAccountMetaKey(accKey string) []byte {
	return append(accountMetaSubstore, accKey...)
}

func getFollowerKey(me string, myFollower string) []byte {
	return append(getFollowerPrefix(me), myFollower...)
}

func getFollowerPrefix(me string) []byte {
	return append(append(accountFollowerSubstore, me...), KeySeparator...)
}

func getFollowingKey(me string, myFollowing string) []byte {
	return append(getFollowingPrefix(me), myFollowing...)
}

func getFollowingPrefix(me string) []byte {
	return append(append(accountFollowingSubstore, me...), KeySeparator...)
}

func getRewardKey(accKey string) []byte {
	return append(accountRewardSubstore, accKey...)
}

func getRelationshipKey(me string, other string) []byte {
	return append(getRelationshipPrefix(me), other...)
}

func getRelationshipPrefix(me string) []byte {
	return append(append(accountRelationshipSubstore, me...), KeySeparator...)
}

func getPendingStakeQueueKey(accKey string) []byte {
	return append(accountPendingStakeQueueSubstore, accKey...)
}

func getBalanceHistoryPrefix(me string) []byte {
	return append(append(accountBalanceHistorySubstore, me...), KeySeparator...)
}
func getBalanceHistoryKey(me string, bucketSlot int64) []byte {
	return strconv.AppendInt(getBalanceHistoryPrefix(me), bucketSlot, 10)
}

func getGrantPubKeyPrefix(me string) []byte {
	return append(append(accountGrantPubKeySubstore, me...), KeySeparator...)
}

func getGrantPubKeyKey(me string, pubKey crypto.PubKey) []byte {
	return append(getGrantPubKeyPrefix(me), pubKey.Bytes()...)
}

//
// post related
//
func getPermlink(author string, postID string) string {
	return string(author + "#" + postID)
}

func getUserPostInfoPrefix(me string) []byte {
	return append(postInfoSubStore, me...)
}

func getPostInfoKey(permlink string) []byte {
	return append(postInfoSubStore, permlink...)
}

func getUserPostMetaPrefix(me string) []byte {
	return append(postMetaSubStore, me...)
}

func getPostMetaKey(permlink string) []byte {
	return append(postMetaSubStore, permlink...)
}

func getPostLikePrefix(permlink string) []byte {
	return append(append(postLikeSubStore, permlink...), KeySeparator...)
}

func getPostLikeKey(permlink string, likeUser string) []byte {
	return append(getPostLikePrefix(permlink), likeUser...)
}

func getPostReportOrUpvotePrefix(permlink string) []byte {
	return append(append(postReportOrUpvoteSubStore, permlink...), KeySeparator...)
}

func getPostReportOrUpvoteKey(permlink string, user string) []byte {
	return append(getPostReportOrUpvotePrefix(permlink), user...)
}

func getPostViewPrefix(permlink string) []byte {
	return append(append(postViewsSubStore, permlink...), KeySeparator...)
}

func getPostViewKey(permlink string, viewUser string) []byte {
	return append(getPostViewPrefix(permlink), viewUser...)
}

func getPostCommentPrefix(permlink string) []byte {
	return append(append(postCommentSubStore, permlink...), KeySeparator...)
}

func getPostCommentKey(permlink string, commentPermlink string) []byte {
	return append(getPostCommentPrefix(permlink), commentPermlink...)
}

func getPostDonationsPrefix(permlink string) []byte {
	return append(append(postDonationsSubStore, permlink...), KeySeparator...)
}

func getPostDonationsKey(permlink string, donateUser string) []byte {
	return append(getPostDonationsPrefix(permlink), donateUser...)
}

//
//  validator related
//
func getValidatorKey(accKey string) []byte {
	return append(validatorSubstore, accKey...)
}

func getValidatorListKey() []byte {
	return validatorListSubstore
}

//
// vote related
//
func getDelegationPrefix(me string) []byte {
	return append(append(delegationSubstore, me...), KeySeparator...)
}

func getDelegationKey(me string, myDelegator string) []byte {
	return append(getDelegationPrefix(me), myDelegator...)
}

func getVotePrefix(id string) []byte {
	return append(append(voteSubstore, id...), KeySeparator...)
}

func getVoteKey(proposalID string, voter string) []byte {
	return append(getVotePrefix(proposalID), voter...)
}

func getVoterKey(me string) []byte {
	return append(voterSubstore, me...)
}

func GetReferenceListKey() []byte {
	return referenceListSubStore
}

func GetDelegateePrefix(me string) []byte {
	return append(append(delegateeSubStore, me...), KeySeparator...)
}

func GetDelegateeKey(me, delegatee string) []byte {
	return append(GetDelegateePrefix(me), delegatee...)
}

//
// developer related
//
func getDeveloperKey(accKey string) []byte {
	return append(developerSubstore, accKey...)
}

func getDeveloperListKey() []byte {
	return developerListSubstore
}

//
// infra related
//
func getInfraProviderKey(accKey string) []byte {
	return append(infraProviderSubstore, accKey...)
}

func getInfraProviderListKey() []byte {
	return infraProviderListSubstore
}

//
// proposal related
//
func getProposalKey(proposalID string) []byte {
	return append(proposalSubstore, proposalID...)
}

func getProposalListKey() []byte {
	return proposalListSubStore
}

//
// param related
//
func getEvaluateOfContentValueParamKey() []byte {
	return evaluateOfContentValueParamSubStore
}

func getGlobalAllocationParamKey() []byte {
	return allocationParamSubStore
}

func getInfraInternalAllocationParamKey() []byte {
	return infraInternalAllocationParamSubStore
}

func getDeveloperParamKey() []byte {
	return developerParamSubStore
}

func getVoteParamKey() []byte {
	return voteParamSubStore
}

func getValidatorParamKey() []byte {
	return validatorParamSubStore
}

func getProposalParamKey() []byte {
	return proposalParamSubStore
}

func getCoinDayParamKey() []byte {
	return coinDayParamSubStore
}

func getBandwidthParamKey() []byte {
	return bandwidthParamSubStore
}

func getAccountParamKey() []byte {
	return accountParamSubstore
}

func GetPostParamKey() []byte {
	return postParamSubStore
}
