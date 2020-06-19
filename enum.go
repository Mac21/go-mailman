package gomailman

type Action int

const (
	//
	Hold Action = iota
	Reject
	Discard
	Accept
	Defer
)

type FilterAction int

const (
	HoldFilter FilterAction = iota
	RejectFilter
	DiscardFilter
	AcceptFilter
	DeferFilter
	ForwardFilter
	PreserveFilter
)

type MemberRole int

const (
	Member MemberRole = iota + 1
	Owner
	Moderator
	NonMember
)

type DeliveryMode int

const (
	RegularDigests DeliveryMode = iota + 1
	PlainTextDigests
	MimeDigests
	SummaryDigsts
)

type DeliveryStatus int

const (
	EnabledStatus DeliveryStatus = iota + 1
	ByUserStatus
	ByBouncesStatus
	ByModeratorStatus
	UnknownStatus
)

type SubscriptionMode int

const (
	AsUserMode SubscriptionMode = iota + 1
	AsAddressMode
)
