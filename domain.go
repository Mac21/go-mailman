package gomailman

type Domain struct {
	AliasDomain string `json:"alias_domain,omitempty"`
	Description string `json:"description,omitempty"`
	MailHost    string `json:"mail_host"`
}
