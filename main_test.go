package main

import (
	"context"
	"testing"
)

func TestXML(t *testing.T) {
	a := `<config>
    <subdomains>
        <subdomain>http://secureline.tools.avast.com</subdomain>
        <subdomain>http://gf.tools.avast.com</subdomain>
        <subdomain>http://files.avast.com</subdomain>
		<submdomains>
		<cookies>
        <!-- avast -->
        <cookie name="dlp-avast" host="amazon">mmm_amz_dlp_777_ppc_m</cookie>
        <cookie name="dlp-avast" host="baixaki">mmm_bxk_dlp_777_ppc_m</cookie>
		<cookies>
	</config>`
	config := ReadXML(context.TODO(), []byte(a))
	if config.Subdomains.Subdomains[0] != "http://secureline.tools.avast.com" {
		t.Fatal(config.Subdomains.Subdomains[0])
	}

}
