package gomailman

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ArchivePolicy int

const (
	NeverArchive ArchivePolicy = iota
	PrivateArchive
	PublicArchive
)

type ArchiveRenderingMode int

const (
	RenderText ArchiveRenderingMode = iota
	RenderMarkdown
)

type List struct {
	ID           string `json:"list_id"`
	Name         string `json:"list_name"`
	MailHost     string `json:"mail_host"`
	DisplayName  string `json:"display_name"`
	FQDNListName string `json:"fqdn_listname"`
	Description  string `json:"description"`
	Advertised   bool   `json:"advertised"`
	Volume       int    `json:"volume"`
	MemberCount  int    `json:"member_count"`
}

type ListConfig struct {
	AcceptableAliases        []string             `json:"acceptable_aliases"`
	AcceptTheseNonMembers    []string             `json:"accept_these_non_members"`
	AdminImmediateNotify     bool                 `json:"admin_immed_notify"`
	AdminNotifyMemberChanges bool                 `json:"admin_notify_mchanges"`
	Administrivia            bool                 `json:"administrivia"`
	Advertised               bool                 `json:"advertised"`
	AllowListPosts           bool                 `json:"allow_list_posts"`
	AnonymousList            bool                 `json:"anonymous_list"`
	ArchivePolicy            ArchivePolicy        `json:"archive_policy"`
	ArchiveRenderingMode     ArchiveRenderingMode `json:"archive_rendering_mode"`
	// TODO:
	// autorespond_owner=GetterSetter(enum_validator(ResponseAction)),
	// autorespond_postings=GetterSetter(enum_validator(ResponseAction)),
	// autorespond_requests=GetterSetter(enum_validator(ResponseAction)),
	// autoresponse_grace_period=GetterSetter(as_timedelta),
	// autoresponse_owner_text=GetterSetter(str),
	// autoresponse_postings_text=GetterSetter(str),
	// autoresponse_request_text=GetterSetter(str),
	// bounces_address=GetterSetter(None),
	// bounce_info_stale_after=GetterSetter(as_timedelta),
	// bounce_notify_owner_on_disable=GetterSetter(as_boolean),
	// bounce_notify_owner_on_removal=GetterSetter(as_boolean),
	// bounce_score_threshold=GetterSetter(integer_ge_zero_validator),
	// bounce_you_are_disabled_warnings=GetterSetter(integer_ge_zero_validator),
	// bounce_you_are_disabled_warnings_interval=GetterSetter(
	//     as_timedelta),
	// collapse_alternatives=GetterSetter(as_boolean),
	// convert_html_to_plaintext=GetterSetter(as_boolean),
	// created_at=GetterSetter(None),
	// default_member_action=GetterSetter(enum_validator(Action)),
	// default_nonmember_action=GetterSetter(enum_validator(Action)),
	// description=GetterSetter(no_newlines_validator),
	// digest_last_sent_at=GetterSetter(None),
	// digest_send_periodic=GetterSetter(as_boolean),
	// digest_size_threshold=GetterSetter(float),
	// digest_volume_frequency=GetterSetter(enum_validator(DigestFrequency)),
	// digests_enabled=GetterSetter(as_boolean),
	// display_name=GetterSetter(str),
	// discard_these_nonmembers=GetterSetter(list_of_strings_validator),
	// dmarc_mitigate_action=GetterSetter(enum_validator(DMARCMitigateAction)),
	// dmarc_mitigate_unconditionally=GetterSetter(as_boolean),
	// dmarc_moderation_notice=GetterSetter(str),
	// dmarc_wrapped_message_text=GetterSetter(str),
	// emergency=GetterSetter(as_boolean),
	// filter_action=GetterSetter(enum_validator(FilterAction)),
	// filter_content=GetterSetter(as_boolean),
	// filter_extensions=GetterSetter(list_of_strings_validator),
	// filter_types=GetterSetter(list_of_strings_validator),
	// first_strip_reply_to=GetterSetter(as_boolean),
	// forward_unrecognized_bounces_to=GetterSetter(
	//     enum_validator(UnrecognizedBounceDisposition)),
	// fqdn_listname=GetterSetter(None),
	// gateway_to_mail=GetterSetter(as_boolean),
	// gateway_to_news=GetterSetter(as_boolean),
	// hold_these_nonmembers=GetterSetter(list_of_strings_validator),
	// include_rfc2369_headers=GetterSetter(as_boolean),
	// info=GetterSetter(str),
	// join_address=GetterSetter(None),
	// last_post_at=GetterSetter(None),
	// leave_address=GetterSetter(None),
	// linked_newsgroup=GetterSetter(str),
	// list_name=GetterSetter(None),
	// mail_host=GetterSetter(None),
	// max_message_size=GetterSetter(integer_ge_zero_validator),
	// max_num_recipients=GetterSetter(integer_ge_zero_validator),
	// max_days_to_hold=GetterSetter(integer_ge_zero_validator),
	// member_roster_visibility=GetterSetter(enum_validator(RosterVisibility)),
	// moderator_password=GetterSetter(password_bytes_validator),
	// newsgroup_moderation=GetterSetter(enum_validator(NewsgroupModeration)),
	// next_digest_number=GetterSetter(None),
	// nntp_prefix_subject_too=GetterSetter(as_boolean),
	// no_reply_address=GetterSetter(None),
	// owner_address=GetterSetter(None),
	// pass_types=GetterSetter(list_of_strings_validator),
	// pass_extensions=GetterSetter(list_of_strings_validator),
	// personalize=GetterSetter(enum_validator(Personalization)),
	// post_id=GetterSetter(None),
	// posting_address=GetterSetter(None),
	// posting_pipeline=GetterSetter(pipeline_validator),
	// preferred_language=LanguageGetterSetter(language_validator),
	// process_bounces=GetterSetter(as_boolean),
	// reject_these_nonmembers=GetterSetter(list_of_strings_validator),
	// reply_goes_to_list=GetterSetter(enum_validator(ReplyToMunging)),
	// reply_to_address=GetterSetter(str),
	// request_address=GetterSetter(None),
	// require_explicit_destination=GetterSetter(as_boolean),
	// respond_to_post_requests=GetterSetter(as_boolean),
	// subject_prefix=GetterSetter(str),
	// subscription_policy=GetterSetter(enum_validator(SubscriptionPolicy)),
	// unsubscription_policy=GetterSetter(enum_validator(SubscriptionPolicy)),
	// usenet_watermark=GetterSetter(None),
	// volume=GetterSetter(None)
	SendGoodbyeMessage bool `json:"send_goodbye_message"`
	SendWelcomeMessage bool `json:"send_welcome_message"`
}

// GetList takes a list_id and returns a single list.
func (c *Client) GetList(listID string) (*List, error) {
	res, err := c.conn.do(http.MethodGet, buildURL("lists", listID), http.NoBody)
	if err != nil {
		return nil, err
	}

	if err := parseResponseError(res); err != nil {
		return nil, err
	}

	list := new(List)
	if err = json.NewDecoder(res.Body).Decode(list); err != nil {
		return nil, err
	}

	return list, res.Body.Close()
}

func (c *Client) AddList(listID string) error {
	fakeList := map[string]string{
		"fqdn_listname": listID,
	}

	b, err := json.Marshal(fakeList)
	if err != nil {
		return err
	}

	res, err := c.conn.do(http.MethodPost, buildURL("lists"), bytes.NewReader(b))
	if err != nil {
		return err
	}

	if err := parseResponseError(res); err != nil {
		return err
	}

	return res.Body.Close()
}

func (c *Client) UpdateListConfig(listID string, lc *ListConfig) error {
	b, err := json.Marshal(lc)
	if err != nil {
		return err
	}

	res, err := c.conn.do(http.MethodPost, buildURL("lists", listID, "config"), bytes.NewReader(b))
	if err != nil {
		return err
	}

	if err := parseResponseError(res); err != nil {
		return err
	}

	return res.Body.Close()
}

func (c *Client) DeleteList(listID string) error {
	res, err := c.conn.do(http.MethodDelete, buildURL("lists", listID), http.NoBody)
	if err != nil {
		return err
	}

	if err := parseResponseError(res); err != nil {
		return err
	}

	return res.Body.Close()
}
