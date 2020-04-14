package evengclient

/*
BasicResponse contains the data returned by the api in case of a get
*/
type BasicResponse struct {
	Code    interface{} `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

/*
CreateResponse contains the data returned by the api in case of a create
*/
type CreateResponse struct {
	ID int `json:"id"`
}

type systemStatusResponse struct {
	Data SystemStatus
}

/*
SystemStatus contains information regarding the system status
*/
type SystemStatus struct {
	Cached      int    `json:"cached"`
	CPU         int    `json:"cpu"`
	Disk        int    `json:"disk"`
	Dynamips    int    `json:"dynamips"`
	Iol         int    `json:"iol"`
	Mem         int    `json:"mem"`
	Qemu        int    `json:"qemu"`
	Qemuversion string `json:"qemu_version"`
	Swap        int    `json:"swap"`
	Version     string `json:"version"`
}

/*
FolderContents is a list of folders and labs inside a folder
*/
type FolderContents struct {
	Folders  Folders  `json:"folders"`
	LabFiles LabFiles `json:"labs"`
}

/*
Folders is a list of folders
*/
type Folders []Folder

/*
Folder contains information regarding a folder
*/
type Folder struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

/*
LabFiles is a list of labfiles
*/
type LabFiles []LabFile

/*
LabFile contains information regarding a lab file
*/
type LabFile struct {
	File string `json:"file"`
	Path string `json:"path"`
}

/*
Lab contains information about a lab
*/
type Lab struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	Body        string `json:"body"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
}

/*
Network contains information about a network
*/
type Network struct {
	Count      int    `json:"count"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Top        int    `json:"top"`
	Left       int    `json:"left"`
	Style      string `json:"style"`
	Linkstyle  string `json:"linkstyle"`
	Color      string `json:"color"`
	Label      string `json:"label"`
	Visibility int    `json:"visibility"`
}

/*
NetworkWithID contains information about a network including its id
*/
type NetworkWithID struct {
	ID int `json:"id"`
	Network
}

/*
Networks an array of Networks
*/
type Networks map[string]NetworkWithID

/*
NetworkTypes is an array of network types
*/
type NetworkTypes map[string]string

/*
Endpoints contains information about ethernet and serial endpoints
*/
type Endpoints struct {
	Ethernet EthernetEndpoints `json:"ethernet"`
	Serial   SerialEndpoints   `json:"serial"`
}

/*
EthernetEndpoints contains information about EthernetEndpoints
*/
type EthernetEndpoints map[string]string

/*
SerialEndpoints contains information about SerialEndpoints
*/
type SerialEndpoints map[string]string

/*
Node contains information about a node
*/
type Node struct {
	UUID       string      `json:"uuid"`
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Status     int         `json:"status"`
	Template   string      `json:"template"`
	CPU        int         `json:"cpu"`
	RAM        int         `json:"ram"`
	Image      string      `json:"image"`
	Console    string      `json:"console"`
	Ethernet   int         `json:"ethernet"`
	Delay      int         `json:"delay"`
	Icon       string      `json:"icon"`
	URL        string      `json:"url"`
	Top        int         `json:"top"`
	Left       int         `json:"left"`
	Config     string      `json:"config"`
	Firstmac   string      `json:"firstmac"`
	Configlist interface{} `json:"configlist"`
}

/*
NodeWithID contains information about a node including its id
*/
type NodeWithID struct {
	ID int `json:"id"`
	Node
}

/*
Nodes an array of Nodes
*/
type Nodes map[string]NodeWithID

/*
Interfaces contains information about ethernet and serial interfaces
*/
type Interfaces struct {
	Ethernet EthernetInterfaces `json:"ethernet"`
	Serial   SerialInterfaces   `json:"serial"`
}

/*
EthernetInterfaces an array of EthernetInterfaces
*/
type EthernetInterfaces []Interface

/*
SerialInterfaces an array of SerialInterfaces
*/
type SerialInterfaces []Interface

/*
Interface basic interface structure
*/
type Interface struct {
	Name      string `json:"name"`
	NetworkID *int   `json:"network_id"`
}

/*
TopologyPoints an array of network topology points
*/
type TopologyPoints []Topology

/*
Topology contains information about a network topology point
*/
type Topology struct {
	Destination            string `json:"destination"`
	DestinationLabel       string `json:"destination_label"`
	DestinationType        string `json:"destination_type"`
	DestinationInterfaceID string `json:"destinationinterfaceid"`
	DestinationNodename    string `json:"destinationnodename"`
	DestinationSuspend     int    `json:"destinationsuspend"`
	DestinationDelay       int    `json:"destinationdelay"`
	DestinationLoss        int    `json:"destinationloss"`
	DestinationBandwidth   int    `json:"destinationbandwidth"`
	DestinationJitter      int    `json:"destinationjitter"`
	Source                 string `json:"source"`
	SourceLabel            string `json:"source_label"`
	SourceType             string `json:"source_type"`
	SourceNodename         string `json:"sourcenodename"`
	SourceInterfaceID      int    `json:"sourceinterface"`
	SourceSuspend          int    `json:"sourcesuspend"`
	SourceDelay            int    `json:"sourcedelay"`
	SourceLoss             int    `json:"sourceloss"`
	SourceBandwidth        int    `json:"sourcebandwidth"`
	SourceJitter           int    `json:"sourcejitter"`
	Type                   string `json:"type"`
	NetworkID              int    `json:"networkid"`
	Style                  string `json:"style"`
	Linkstyle              string `json:"linkstyle"`
	Label                  string `json:"label"`
	Color                  string `json:"color"`
}

/*
Pictures an array containing pictures
*/
type Pictures []Picture

/*
Picture contains information about a specific picture
*/
type Picture struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Map    string `json:"map"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

/*
Templates contains information about all templates
*/
type Templates map[string]string

/*
Template contains information about a specific template
*/
type Template struct {
	Description string  `json:"description"`
	Options     Options `json:"options"`
	Type        string  `json:"type"`
	Qemu        Qemu    `json:"qemu"`
}

/*
Options contains information about the templates options
*/
type Options struct {
	Config      Config      `json:"config"`
	Delay       Delay       `json:"delay"`
	Ethernet    Ethernet    `json:"ethernet"`
	Icon        Icon        `json:"icon"`
	Image       Image       `json:"image"`
	Name        Name        `json:"name"`
	Nvram       Nvram       `json:"nvram"`
	RAM         RAM         `json:"ram"`
	Serial      Serial      `json:"serial"`
	UUID        UUID        `json:"uuid"`
	Cpulimit    Cpulimit    `json:"cpulimit"`
	CPU         CPU         `json:"cpu"`
	Firstmac    Firstmac    `json:"firstmac"`
	Qemuversion Qemuversion `json:"qemuversion"`
	Qemuarch    Qemuarch    `json:"qemuarch"`
	Qemunic     Qemunic     `json:"qemunic"`
	Qemuoptions Qemuoptions `json:"qemuoptions"`
	Console     Console     `json:"console"`
	Rdpuser     Rdpuser     `json:"rdpuser"`
	Rdppassword Rdppassword `json:"rdppassword"`
}

/*
Config contains information about a template config
*/
type Config struct {
	List StringArray `json:"list"`
	StringValTemplateOption
}

/*
List resembles the field List of most stucts returned by the api
*/
type List map[string]string

/*
StringArray resembles the field List of some stucts returned by the api because they are different than the List struct
*/
type StringArray []string

/*
Delay contains information about the templates delay option
*/
type Delay struct {
	IntValTemplateOption
}

/*
Ethernet contains information about the templates ethernet option
*/
type Ethernet struct {
	IntValTemplateOption
}

/*
Icon contains information about the templates icon option
*/
type Icon struct {
	List List `json:"list"`
	StringValTemplateOption
}

/*
Image contains information about the templates image option
*/
type Image struct {
	List interface{} `json:"list"`
	StringValTemplateOption
}

/*
Name contains information about the templates name option
*/
type Name struct {
	StringValTemplateOption
}

/*
Nvram contains information about the templates nvram option
*/
type Nvram struct {
	IntValTemplateOption
}

/*
RAM contains information about the templates ram option
*/
type RAM struct {
	IntValTemplateOption
}

/*
Serial contains information about the templates Serial option
*/
type Serial struct {
	IntValTemplateOption
}

/*
UUID contains information about the templates UUID option
*/
type UUID struct {
	StringValTemplateOption
}

/*
Cpulimit contains information about the templates Cpulimit option
*/
type Cpulimit struct {
	IntValTemplateOption
}

/*
CPU contains information about the templates cpu option
*/
type CPU struct {
	IntValTemplateOption
}

/*
Firstmac contains information about the templates firstmac option
*/
type Firstmac struct {
	StringValTemplateOption
}

/*
Qemuversion contains information about the templates qemuversion option
*/
type Qemuversion struct {
	List List `json:"list"`
	StringValTemplateOption
}

/*
Qemuarch contains information about the templates qemuarch option
*/
type Qemuarch struct {
	List List `json:"list"`
	StringValTemplateOption
}

/*
Qemunic contains information about the templates qemunic option
*/
type Qemunic struct {
	List List `json:"list"`
	StringValTemplateOption
}

/*
Qemuoptions contains information about the templates qemuoptions option
*/
type Qemuoptions struct {
	StringValTemplateOption
}

/*
Console contains information about the templates qonsole option
*/
type Console struct {
	List List `json:"list"`
	StringValTemplateOption
}

/*
Rdpuser contains information about the templates rdpuser option
*/
type Rdpuser struct {
	StringValTemplateOption
}

/*
Rdppassword contains information about the templates rdppassword option
*/
type Rdppassword struct {
	StringValTemplateOption
}

/*
StringValTemplateOption contains the standard fields of a templateOption
*/
type StringValTemplateOption struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

/*
IntValTemplateOption contains the standard fields of a templateOption
*/
type IntValTemplateOption struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value int    `json:"value"`
}

/*
Qemu contains information about the templates qemu option
*/
type Qemu struct {
	Arch    string `json:"arch"`
	Nic     string `json:"nic"`
	Options string `json:"options"`
}

/*
Users a list of all users in a lab
*/
type Users map[string]User

/*
User contains information about a specific user
*/
type User struct {
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	ExtAuth     string  `json:"ext_auth"`
	RAM         string  `json:"ram"`
	CPU         string  `json:"cpu"`
	Expiration  string  `json:"expiration"`
	Name        string  `json:"name"`
	Session     string  `json:"session"`
	Role        string  `json:"role"`
	Online      int     `json:"online"`
	IP          string  `json:"ip"`
	Folder      string  `json:"folder"`
	Lab         string  `json:"lab"`
	Pod         string  `json:"pod"`
	Pexpiration string  `json:"pexpiration"`
	DateStart   string  `json:"datestart"`
	DiskUsage   float64 `json:"diskusage"`
}

/*
UserRoles - contains information about user roles
*/
type UserRoles map[string]string
