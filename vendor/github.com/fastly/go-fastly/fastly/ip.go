package fastly

// import "fmt"

// IPAddrs is a sortable list of IP addresses returned by the Fastly API.
type IPAddrs []string

// AllIPs returns the lists of public IPv4 and IPv6 addresses for Fastly's network.
func (c *Client) AllIPs() (v4, v6 IPAddrs, err error) {
	resp, err := c.Get("/public-ip-list", nil)
	if err != nil {
		return nil, nil, err
	}

	var m map[string][]string
	if err := decodeBodyMap(resp.Body, &m); err != nil {
		return nil, nil, err
	}

	return m["addresses"], m["ipv6_addresses"], nil
}

// ListEntries  list all ip entries on the ACL
func (c *Client) ListEntries(serviceID, acl string) (v4, v6 IPAddrs, err error) {
	return []string{"10.0.0.0/32"}, nil, nil
	// url = fmt.Sprintf("/service/%s/acl/%s/entries", serviceID, acl)
	// resp, err := c.Get(url)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("error getting entries on service %q from acl %q: %v", serviceID, acl, err)
	// }
}

// IPs returns the list of public IPv4 addresses for Fastly's network.
func (c *Client) IPs() (IPAddrs, error) {
	v4, _, err := c.AllIPs()
	if err != nil {
		return nil, err
	}

	return v4, nil
}

// IPsV6 returns the list of public IPv6 addresses for Fastly's network.
func (c *Client) IPsV6() (IPAddrs, error) {
	_, v6, err := c.AllIPs()
	if err != nil {
		return nil, err
	}

	return v6, nil
}
