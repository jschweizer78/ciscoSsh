package configXML

// JabberConfigXMLFile to model jabber-config.xml
type JabberConfigXMLFile struct {
	Config ConfigXML `xml:"config"`
}

// ConfigXML to model jabber-config.xml
type ConfigXML struct {
	AttrVersion string            `xml:"version,attr"`
	Client      ClientXML         `xml:"Client"`
	UpdateURL   string            `xml:"UpdateUrl"`
	Policies    map[string]string // TODO change to struct
	Directory   map[string]string // TODO change to struct
}

// ClientXML to model jabber-config.xml
type ClientXML struct {
	JabberPluginConfig PluginConfigXML    `xml:"jabber-plugin-config"`
	Options            []ClientOptionsXML `xml:"Options"`
}

// ClientOptionsXML to model jabber-config.xml
type ClientOptionsXML struct {
	AllowUserCustomTabs bool `xml:"AllowUserCustomTabs"`
}

// PluginConfigXML to model jabber-config.xml
type PluginConfigXML struct {
	BrowserPlugin []BrowserPlugInXML `xml:"browser-plugin"`
}

// BrowserPlugInXML to model jabber-config.xml
type BrowserPlugInXML struct {
	Page BrowserPlugInPageXML `xmml:"page"`
}

/*
You can specify the ${UserID} token as part of the value for the url parameter.
When users sign in, the client replaces the ${UserID} token with the username of the logged in user.
*/

// BrowserPlugInPageXML to model jabber-config.xml
type BrowserPlugInPageXML struct {
	AttrRefresh bool   `xml:"refresh,attr"`
	AttrPreload bool   `xml:"preload,attr"`
	Tooltip     string `xml:"tooltip"`
	Icon        string `xml:"icon"`
	URL         string `xml:"url"`
}

/*
<?xml version="1.0" encoding="utf-8"?>
<config version="1.0">
 <Client>
  <PrtLogServerUrl>http://server_name:port/path/prt_script.php</PrtLogServerUrl>
  <jabber-plugin-config>
   <browser-plugin>
    <page refresh="true" preload="true">
     <tooltip>Cisco</tooltip>
     <icon>http://www.cisco.com/web/fw/i/logo.gif</icon>
     <url>www.cisco.com</url>
    </page>
   </browser-plugin>
  </jabber-plugin-config>
  </Client>
  <Options>
    <Set_Status_Inactive_Timeout>20</Set_Status_Inactive_Timeout>
    <StartCallWithVideo>false</StartCallWithVideo>
  </Options>
  <Policies>
    <Disallowed_File_Transfer_Types>.exe;.msi</Disallowed_File_Transfer_Types>
  </Policies>
  <Directory>
   <BDIPresenceDomain>example.com</BDIPresenceDomain>
    <BDIPrimaryServerName>dir.example.com</BDIPrimaryServerName>
    <BDISearchBase1>ou=staff,dc=example,dc=com</BDISearchBase1>
    <BDIConnectionUsername>ad_jabber_access@example.com</BDIConnectionUsername>
    <BDIConnectionPassword>jabber</BDIConnectionPassword>
    <BDIPhotoUriSubstitutionEnabled>True</BDIPhotoUriSubstitutionEnabled>
    <BDIPhotoUriSubstitutionToken>sAMAccountName</BDIPhotoUriSubstitutionToken>
    <BDIPhotoUriWithToken>http://example.com/photo/sAMAccountName.jpg
   </BDIPhotoUriWithToken>
  </Directory>
</config>
*/
