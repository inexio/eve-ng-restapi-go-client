package evengclient

import (
	"github.com/go-resty/resty"
	"github.com/pkg/errors"
	"strconv"
)

//---------- Client operations ----------//

/*
EveNgClient is an implementation of the client specified for eve-ng
*/
type EveNgClient struct {
	client
}

/*
NewEveNgClient generates a new eve-ng api-client object which can be used to communicate with the eve-ng REST API
*/
func NewEveNgClient(baseUrl string) (*EveNgClient, error) {
	if baseUrl == "" {
		return nil, errors.New("invalid base url")
	}

	//if baseUrl does not end with an "/" it has to be added to the string
	if lastChar := baseUrl[len(baseUrl)-1:]; lastChar != "/" {
		baseUrl += "/"
	}
	clientData := clientData{baseUrl: baseUrl, resty: resty.New(), useAuth: false}
	newClient := client{&clientData}
	return &EveNgClient{newClient}, nil
}

/*
Login performs a login via an eve-ng api-client
*/
func (c *EveNgClient) Login() error {
	if !c.isValid() {
		return &NotValidError{}
	}
	escapedUsername, err := jsonEscape(c.username)
	if err != nil {
		return errors.Wrap(err, "error during json escaping username")
	}

	escapedPassword, err := jsonEscape(c.password)
	if err != nil {
		return errors.Wrap(err, "error during json escaping password")
	}
	_, err = c.request("POST", endpointPath+"auth/login", `{"username":"`+escapedUsername+`","password":"`+escapedPassword+`"}`, nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http login request")
	}
	return nil
}

/*
Logout performs a logout via an eve-ng api-client
*/
func (c *EveNgClient) Logout() error {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("GET", endpointPath+"auth/logout", "", nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http loqout request")
	}
	return nil
}



//---------- System health operations ----------//

/*
GetSystemStatus returns the system status of eve-ng
*/
func (c *EveNgClient) GetSystemStatus() (SystemStatus, error) {
	if !c.isValid() {
		return SystemStatus{}, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"status", "", nil, nil)
	if err != nil {
		return SystemStatus{}, errors.Wrap(err, "error during http get system status request")
	}
	var systemStatusResponse SystemStatus
	err = c.unmarshalDataIntoStruct(response.Body(), &systemStatusResponse)
	if err != nil {
		return SystemStatus{}, errors.Wrap(err, "error during unmarshal")
	}
	return systemStatusResponse, nil
}



//---------- Lab operations ----------//

/*
AddLab adds a lab to
*/
func (c *EveNgClient) AddLab(path string, name string, version string, author string, description string, body string) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("POST", endpointPath+"labs", `{"path":"`+path+`","name":"`+name+`","version":"`+version+`","author":"`+author+`","description":"`+description+`","body":"`+body+`"}`, nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}

	return nil
}

/*
RemoveLab remove an existing lab
*/
func (c *EveNgClient) RemoveLab(labPath string) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("DELETE", endpointPath+"labs/"+labPath, "", nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}

	return nil
}

/*
MoveLab moves a lab to an existing folder
*/
func (c *EveNgClient) MoveLab(labPath string, newPath string) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	response, err := c.request("PUT", endpointPath+"labs/"+labPath+"/move", `{"path":"`+newPath+`"}`, nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}
	var lab Lab
	err = c.unmarshalDataIntoStruct(response.Body(), &lab)
	if err != nil {
		return err
	}
	return nil
}

/*
EditLab edit an existing lab
*/
func (c *EveNgClient) EditLab(labPath string, name string, version string, author string, description string) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	response, err := c.request("PUT", endpointPath+"labs/"+labPath+"", `{"name":"`+name+`","version":"`+version+`","author":"`+author+`","description":"`+description+`"}`, nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}
	var lab Lab
	err = c.unmarshalDataIntoStruct(response.Body(), &lab)
	if err != nil {
		return err
	}
	return nil
}

/*
GetLab retrieves data for the given lab
*/
func (c *EveNgClient) GetLab(labPath string) (Lab, error) {
	if !c.isValid() {
		return Lab{}, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"labs/"+labPath+"", "", nil, nil)
	if err != nil {
		return Lab{}, errors.Wrap(err, "error during http get request")
	}
	var lab Lab
	err = c.unmarshalDataIntoStruct(response.Body(), &lab)
	if err != nil {
		return Lab{}, err
	}
	return lab, nil
}

/*
GetLabTopology retrieves topology for given lab
*/
func (c *EveNgClient) GetLabTopology(labPath string) (TopologyPoints, error) {
	if !c.isValid() {
		return nil, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"labs/"+labPath+"/topology", "", nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error during http get request")
	}
	var topologyPoints TopologyPoints
	err = c.unmarshalDataIntoStruct(response.Body(), &topologyPoints)
	if err != nil {
		return nil, err
	}
	return topologyPoints, nil
}



//---------- Node operations ----------//

/*
AddLabNode add a new node to a lab
*/
func (c *EveNgClient) AddLabNode(labPath string, nodeType string, template string, config string, delay int, icon string, image string, name string, left string, top string, ram string, console string, cpu int, cpuLimit string, ethernet int, firstMac string, rdpUser string, rdpPassword string, uuid string, count int) (int, error) {
	if !c.isValid() {
		return 0, &NotValidError{}
	}
	response, err := c.request("POST", endpointPath+"labs/"+labPath+"/nodes", `{"path":"`+labPath+`","type":"`+nodeType+`","template":"`+template+`","config":"`+config+`","delay":"`+strconv.Itoa(delay)+`","icon":"`+icon+`","image":"`+image+`","name":"`+name+`","left":"`+left+`","top":"`+top+`","ram":"`+ram+`","console":"`+console+`","cpu":"`+strconv.Itoa(cpu)+`","cpulimit":"`+cpuLimit+`","fistmac":"`+firstMac+`","ethernet":"`+strconv.Itoa(ethernet)+`","rdp_user":"`+rdpUser+`","rdp_password":"`+rdpPassword+`","uuid":"`+uuid+`","count":"`+strconv.Itoa(count)+`"}`, nil, nil)
	if err != nil {
		return 0, errors.Wrap(err, "error during http get request")
	}

	var createResponse CreateResponse
	err = c.unmarshalDataIntoStruct(response.Body(), &createResponse)
	if err != nil {
		return 0, err
	}

	return createResponse.Id, nil
}

/*
RemoveLabNode removes a node from a lab
*/
func (c *EveNgClient) RemoveLabNode(labPath string, nodeId int) error {
	if !c.isValid() {
		return &NotValidError{}
	}

	_, err := c.request("DELETE", endpointPath+"labs/"+labPath+"/nodes/"+strconv.Itoa(nodeId), "", nil, nil)
	if err != nil {
		return errors.Wrap(err, "Error during http request")
	}

	return nil
}

/*
GetLabNodes returns all nodes in a lab
*/
func (c *EveNgClient) GetLabNodes(labPath string) (Nodes, error) {
	if !c.isValid() {
		return nil, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"labs/"+labPath+"/nodes", "", nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error during http get request")
	}
	var nodes Nodes
	err = c.unmarshalDataIntoStruct(response.Body(), &nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

/*
GetLabNode - Returns data for a specific lab node
*/
func (c *EveNgClient) GetLabNode(labPath string, nodeId int) (Node, error) {
	if !c.isValid() {
		return Node{}, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"labs/"+labPath+"/nodes/"+strconv.Itoa(nodeId), "", nil, nil)
	if err != nil {
		return Node{}, errors.Wrap(err, "error during http get request")
	}
	var node Node
	err = c.unmarshalDataIntoStruct(response.Body(), &node)
	if err != nil {
		return Node{}, err
	}
	return node, err
}

/*
StartLabNodes starts all nodes in a lab
*/
func (c *EveNgClient) StartLabNodes(labPath string) (error) {
	nodes, err := c.GetLabNodes(labPath)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}

	for _, node := range nodes {
		err = c.StartLabNode(labPath, node.Id)
		if err != nil {
			return errors.Wrap(err, "error during http get request")
		}
	}

	return nil
}

/*
StartLabNode starts a specific node in a lab
*/
func (c *EveNgClient) StartLabNode(labPath string, nodeId int) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("GET", endpointPath+"labs/"+labPath+"/nodes/"+strconv.Itoa(nodeId)+"/start", "", nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}
	return nil
}

/*
StopLabNodes stops all nodes in a lab
*/
func (c *EveNgClient) StopLabNodes(labPath string) (error) {
	nodes, err := c.GetLabNodes(labPath)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}

	for _, node := range nodes {
		err = c.StopLabNode(labPath, node.Id)
		if err != nil {
			return errors.Wrap(err, "error during http get request")
		}
	}

	return nil
}

/*
StopLabNode stops a specific node in a lab
*/
func (c *EveNgClient) StopLabNode(labPath string, nodeId int) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("GET", endpointPath+"labs/"+labPath+"/nodes/"+strconv.Itoa(nodeId)+"/stop/stopmode=3", "", nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}
	return nil
}

/*
WipeLabNodes wipes all nodes in a lab
*/
func (c *EveNgClient) WipeLabNodes(labPath string) (error) {
	nodes, err := c.GetLabNodes(labPath)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}

	for _, node := range nodes {
		err = c.WipeLabNode(labPath, node.Id)
		if err != nil {
			return errors.Wrap(err, "error during WipeLabNode")
		}
	}

	return nil
}

/*
WipeLabNode wipes a specific node in a lab
*/
func (c *EveNgClient) WipeLabNode(labPath string, nodeId int) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("GET", endpointPath+"labs/"+labPath+"/nodes/"+strconv.Itoa(nodeId)+"/wipe", "", nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}
	return err
}

/*
ExportLabNodes exports all nodes in a lab
*/
func (c *EveNgClient) ExportLabNodes(labPath string) (error) {
	nodes, err := c.GetLabNodes(labPath)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}

	for _, node := range nodes {
		err = c.ExportLabNode(labPath, node.Id)
		if err != nil {
			return errors.Wrap(err, "error during WipeLabNode")
		}
	}

	return nil
}

/*
ExportLabNode exports a specific node in a lab
*/
func (c *EveNgClient) ExportLabNode(labPath string, nodeId int) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("GET", endpointPath+"labs/"+labPath+"/nodes/"+strconv.Itoa(nodeId)+"/wipe", "", nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}

	return err
}



//---------- Node Interface operations ----------//

/*
ConnectLabNodeInterfaceToNetwork connects the given node interface to a network
*/
func (c *EveNgClient) ConnectLabNodeInterfaceToNetwork(labPath string, nodeId int, interfaceId int, networkId int) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("PUT", endpointPath+"labs/"+labPath+"/nodes/"+strconv.Itoa(nodeId)+"/interfaces", `{"`+strconv.Itoa(interfaceId)+`":"`+strconv.Itoa(networkId)+`"}`, nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}

	return nil
}

/*
GetLabNodeInterfaces returns all interfaces for a specific lab node
*/
func (c *EveNgClient) GetLabNodeInterfaces(labPath string, nodeId int) (Interfaces, error) {
	if !c.isValid() {
		return Interfaces{}, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"labs/"+labPath+"/nodes/"+strconv.Itoa(nodeId)+"/interfaces", "", nil, nil)
	if err != nil {
		return Interfaces{}, errors.Wrap(err, "error during http get request")
	}
	var interfaces Interfaces
	err = c.unmarshalDataIntoStruct(response.Body(), &interfaces)
	if err != nil {
		return Interfaces{}, err
	}
	return interfaces, nil
}



//---------- Node Template operations ----------//

/*
GetNodeTemplates returns all node templates
*/
func (c *EveNgClient) GetNodeTemplates() (Templates, error) {
	if !c.isValid() {
		return nil, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"list/templates/", "", nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error during http get request")
	}
	var templates Templates
	err = c.unmarshalDataIntoStruct(response.Body(), &templates)
	if err != nil {
		return nil, err
	}
	return templates, nil
}

/*
GetNodeTemplate returns data of a specific template
*/
func (c *EveNgClient) GetNodeTemplate(templateName string) (Template, error) {
	if !c.isValid() {
		return Template{}, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"list/templates/"+templateName, "", nil, nil)
	if err != nil {
		return Template{}, errors.Wrap(err, "error during http get request")
	}
	var template Template
	err = c.unmarshalDataIntoStruct(response.Body(), &template)
	if err != nil {
		return Template{}, err
	}
	return template, nil
}



//---------- Network operations ----------//

/*
AddLabNetwork add a new network to a lab
*/
func (c *EveNgClient) AddLabNetwork(labPath string, networkType string, networkName string, left string, top string, visibility int, postfix int) (int, error) {
	if !c.isValid() {
		return 0, &NotValidError{}
	}
	response, err := c.request("POST", endpointPath+"labs/"+labPath+"/networks", `{"type":"`+networkType+`","name":"`+networkName+`","left":"`+left+`","top":"`+top+`","visibility":"`+strconv.Itoa(visibility)+`","postfix":"`+strconv.Itoa(postfix)+`"}`, nil, nil)
	if err != nil {
		return 0, errors.Wrap(err, "error during http get request")
	}

	var createResponse CreateResponse
	err = c.unmarshalDataIntoStruct(response.Body(), &createResponse)
	if err != nil {
		return 0, err
	}

	return createResponse.Id, err
}

/*
RemoveLabNetwork removes a given network
*/
func (c *EveNgClient) RemoveLabNetwork(labPath string, networkId int) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}

	_, err := c.request("DELETE", endpointPath+"labs/"+labPath+"/networks/"+strconv.Itoa(networkId), "", nil, nil)
	if err != nil {
		return errors.Wrap(err, "Error during http delete request")
	}

	return nil
}

/*
GetLabNetworks returns a list of all networks configured in a lab
*/
func (c *EveNgClient) GetLabNetworks(labPath string) (Networks, error) {
	if !c.isValid() {
		return nil, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"labs/"+labPath+"/networks", "", nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error during http get request")
	}
	var networks Networks
	err = c.unmarshalDataIntoStruct(response.Body(), &networks)
	if err != nil {
		return nil, err
	}
	return networks, nil
}

/*
GetLabNetwork returns data for given network id for lab
*/
func (c *EveNgClient) GetLabNetwork(labPath string, networkId int) (Network, error) {
	if !c.isValid() {
		return Network{}, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"labs/"+labPath+"/networks/"+strconv.Itoa(networkId), "", nil, nil)
	if err != nil {
		return Network{}, errors.Wrap(err, "error during http get request")
	}
	var network Network
	err = c.unmarshalDataIntoStruct(response.Body(), &network)
	if err != nil {
		return Network{}, err
	}
	return network, nil
}

/*
GetNetworkTypes returns all available network types
*/
func (c *EveNgClient) GetNetworkTypes() (NetworkTypes, error) {
	if !c.isValid() {
		return nil, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"list/networks", "", nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error during http get request")
	}
	var networkTypes NetworkTypes
	err = c.unmarshalDataIntoStruct(response.Body(), &networkTypes)
	if err != nil {
		return nil, err
	}
	return networkTypes, nil
}


//---------- User operations ----------//

/*
AddUser adds a new user
*/
func (c *EveNgClient) AddUser(username string, name string, email string, password string, role string, expiration string, dateStart string, extAuth string, pod int, pexpiration string, cpu int, ram int) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("POST", endpointPath+"users", `{"username":"`+username+`","name":"`+name+`","email":"`+email+`","password":"`+password+`","role":"`+role+`","expiration":"`+expiration+`","datestart":"`+dateStart+`","extauth":"`+extAuth+`","pod":`+strconv.Itoa(pod)+`,"pexpiration":"`+pexpiration+`","cpu":`+strconv.Itoa(cpu)+`,"ram":`+strconv.Itoa(ram)+`}`, nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}
	return nil
}

/*
RemoveUser removes an existing user
*/
func (c *EveNgClient) RemoveUser(username string) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("DELETE", endpointPath+"users/"+username, "", nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}
	return nil
}

/*
EditUser edits an existing user
*/
func (c *EveNgClient) EditUser(username string, name string, email string, password string, role string, expiration string, pod int, pexpiration string) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("PUT", endpointPath+"users/"+username, `{"name":"`+name+`","email":"`+email+`","password":"`+password+`","role":"`+role+`","expiration":"`+expiration+`","pod":`+strconv.Itoa(pod)+`,"pexpiration":"`+pexpiration+`"}`, nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}
	return nil
}

/*
GetUsers retreives a list of all users
*/
func (c *EveNgClient) GetUsers() (Users, error) {
	if !c.isValid() {
		return nil, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"users/", "", nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error during http get request")
	}
	var users Users
	err = c.unmarshalDataIntoStruct(response.Body(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

/*
GetUser retreives data for given user
*/
func (c *EveNgClient) GetUser(username string) (User, error) {
	if !c.isValid() {
		return User{}, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"users/"+username, "", nil, nil)
	if err != nil {
		return User{}, errors.Wrap(err, "error during http get request")
	}
	var user User
	err = c.unmarshalDataIntoStruct(response.Body(), &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

/*
GetUserRoles returns all available user roles
*/
func (c *EveNgClient) GetUserRoles() (UserRoles, error) {
	if !c.isValid() {
		return nil, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"list/roles", "", nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error during http get request")
	}
	var userRoles UserRoles
	err = c.unmarshalDataIntoStruct(response.Body(), &userRoles)
	if err != nil {
		return nil, err
	}
	return userRoles, nil
}



//---------- Folder operations ----------//

/*
AddFolder adds a new folder to the given directory
*/
func (c *EveNgClient) AddFolder(path string, folderName string) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("POST", endpointPath+"folders", `{"path":"`+path+`","name":"`+folderName+`"}`, nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}
	return nil
}

/*
MoveFolder moves/renames an existing folder
*/
func (c *EveNgClient) MoveFolder(oldPath string, newPath string) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("PUT", endpointPath+"folders/"+oldPath, `{"path":"`+newPath+`"}`, nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}
	return nil
}

/*
RemoveFolder deletes an existing folder
*/
func (c *EveNgClient) RemoveFolder(path string) (error) {
	if !c.isValid() {
		return &NotValidError{}
	}
	_, err := c.request("DELETE", endpointPath+"folders/"+path, "", nil, nil)
	if err != nil {
		return errors.Wrap(err, "error during http get request")
	}
	return nil
}

/*
GetFolderContents returns contents of a given folder
*/
func (c *EveNgClient) getFolderContents(folder string) (FolderContents, error) {
	if !c.isValid() {
		return FolderContents{}, &NotValidError{}
	}
	response, err := c.request("GET", endpointPath+"folders/"+folder, "", nil, nil)
	if err != nil {
		return FolderContents{}, errors.Wrap(err, "error during http get request")
	}
	var folderContents FolderContents
	err = c.unmarshalDataIntoStruct(response.Body(), &folderContents)
	if err != nil {
		return FolderContents{}, err
	}
	return folderContents, nil
}

/*
GetLabFiles returns all lab files in a given path
*/
func (c * EveNgClient) GetLabFiles(path string) (LabFiles, error) {
	folderContents, err := c.getFolderContents(path)
	if err != nil {
		return LabFiles{}, errors.Wrap(err, "error while retrieving lab files for given path")
	}

	return folderContents.LabFiles, err
}

/*
GetFolders returns all folders in a given path
*/
func (c * EveNgClient) GetFolders(path string) (Folders, error) {
	folderContents, err := c.getFolderContents(path)
	if err != nil {
		return Folders{}, errors.Wrap(err, "error while retrieving folder for given path")
	}

	return folderContents.Folders, err
}