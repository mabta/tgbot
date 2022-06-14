package option

import (
	"fmt"
	"strings"
)

type Option struct {
	Token      string
	UpdatePath string
	Domain     string
}

func NewWithPath(token, updatePath, domain string) *Option {
	domain = strings.TrimSuffix(domain, "/")
	if !strings.HasPrefix(updatePath, "/") {
		updatePath = "/" + updatePath
	}
	return &Option{
		Token:      token,
		UpdatePath: updatePath,
		Domain:     domain,
	}
}
func New(token, domain string) *Option {
	updatePath := fmt.Sprintf("/tgbot-%s", token)
	updatePath = strings.ReplaceAll(updatePath, ":", "-")
	return &Option{
		Token:      token,
		UpdatePath: updatePath,
		Domain:     domain,
	}
}
func (o *Option) WebhookUrl() string {
	return fmt.Sprintf("https://%s%s", o.Domain, o.UpdatePath)
}
