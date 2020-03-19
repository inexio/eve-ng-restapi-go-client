package evengclient

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"strconv"
	"testing"
)

/*
TestEveNgClient_LoginLogout covers:
	- Login
	- Logout
*/
func TestEveNgClient_LoginLogout(t *testing.T) {
	eveNgClient, err := NewEveNgClient(viper.GetString("BaseUrl"))
	if !assert.NoError(t, err, "Error while creating API client") {
		return
	}

	err = eveNgClient.SetUsernameAndPassword(viper.GetString("Username"), viper.GetString("Password"))
	if !assert.NoError(t, err, "Error while setting username and password") {
		return
	}

	err = eveNgClient.Login()
	if !assert.NoError(t, err, "Error during login") {
		return
	}
	defer func() {
		err = eveNgClient.Logout()
		if !assert.NoError(t, err, "Error during logout") {
			return
		}
	}()
}

/*
TestEveNgClient_GetSystemStatus covers:
	- GetSystemStatus
*/
func TestEveNgClient_GetSystemStatus(t *testing.T) {
	eveNgClient, err := NewEveNgClient(viper.GetString("BaseUrl"))
	if !assert.NoError(t, err, "Error while creating API client") {
		return
	}

	err = eveNgClient.SetUsernameAndPassword(viper.GetString("Username"), viper.GetString("Password"))
	if !assert.NoError(t, err, "Error while setting username and password") {
		return
	}

	err = eveNgClient.Login()
	if !assert.NoError(t, err, "Error during login") {
		return
	}
	defer func() {
		err = eveNgClient.Logout()
		if !assert.NoError(t, err, "Error during logout") {
			return
		}
	}()

	systemStatus, err := eveNgClient.GetSystemStatus()
	if assert.NoError(t, err, "Error during GetSystemStatus operation") {
		assert.NotNil(t, systemStatus.Cached, "Dynamips is nil")
		assert.NotNil(t, systemStatus.Cpu, "Vpcs is nil")
		assert.NotNil(t, systemStatus.Disk, "Docker is nil")
		assert.NotNil(t, systemStatus.Dynamips, "Qemu is nil")
		assert.NotNil(t, systemStatus.Iol, "Iol is nil")
		assert.NotNil(t, systemStatus.Mem, "Mem is nil")
		assert.NotNil(t, systemStatus.Qemu, "Qemu is nil")
		assert.NotEmpty(t, systemStatus.Qemuversion, "QemuVersion is empty")
		assert.NotNil(t, systemStatus.Swap, "Swap is nil")
		assert.NotEmpty(t, systemStatus.Version, "Version is empty")
	}
}

/*
TestEveNgClient_GetNodeTemplates covers:
	- GetNodeTemplates
	- GetNodeTemplate
*/
func TestEveNgClient_NodeTemplates(t *testing.T) {
	eveNgClient, err := NewEveNgClient(viper.GetString("BaseUrl"))
	if !assert.NoError(t, err, "Error while creating API client") {
		return
	}

	err = eveNgClient.SetUsernameAndPassword(viper.GetString("Username"), viper.GetString("Password"))
	if !assert.NoError(t, err, "Error while setting username and password") {
		return
	}

	err = eveNgClient.Login()
	if !assert.NoError(t, err, "Error during login") {
		return
	}
	defer func() {
		err = eveNgClient.Logout()
		if !assert.NoError(t, err, "Error during logout") {
			return
		}
	}()

	nodeTemplates, err := eveNgClient.GetNodeTemplates()
	if assert.NoError(t, err, "Error during GetNodeTemplates operation") {
		assert.True(t, len(nodeTemplates) > 0, "No node templates found during GetNodeTemplates operation")
	}

	for nodeTemplateName := range nodeTemplates {
		nodeTemplate, err := eveNgClient.GetNodeTemplate(nodeTemplateName)
		if assert.NoError(t, err, "Error during GetNodeTemplate") {
			assert.True(t, nodeTemplate.Description != "", "No description for node template found during GetNodeTemplate operation")
			assert.NotEmpty(t, nodeTemplate.Type, "Node template type is empty")
			//Options
			//Options.Config
			assert.IsType(t, StringArray{}, nodeTemplate.Options.Config.List, "Node template options config list is not an actual list")
			assert.NotEmpty(t, nodeTemplate.Options.Config.Name, "Node template options config name is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Config.Type, "Node template options config type is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Config.Value, "Node template options config value is empty")
			//Options.Delay
			assert.NotEmpty(t, nodeTemplate.Options.Delay.Name, "Node template options delay name is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Delay.Type, "Node template options delay type is empty")
			assert.NotNil(t, nodeTemplate.Options.Delay.Value, "Node template options delay value is nil")
			//Options.Ethernet
			assert.NotEmpty(t, nodeTemplate.Options.Ethernet.Name, "Node template options ethernet name is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Ethernet.Type, "Node template options ethernet type is empty")
			assert.NotNil(t, nodeTemplate.Options.Ethernet.Value, "Node template options ethernet value is nil")
			//Options.Icon
			assert.IsType(t, List{}, nodeTemplate.Options.Icon.List, "Node template options icon list is not an actual list")
			assert.NotEmpty(t, nodeTemplate.Options.Icon.Name, "Node template options icon name is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Icon.Type, "Node template options icon type is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Icon.Value, "Node template options icon value is empty")
			//Options.Image
			var interfaceType []interface{}
			assert.IsType(t, interfaceType, nodeTemplate.Options.Image.List, "Node template options image list is not an actual []interface{}")
			assert.NotEmpty(t, nodeTemplate.Options.Image.Name, "Node template options image name is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Image.Type, "Node template options image type is empty")
			assert.NotNil(t, nodeTemplate.Options.Image.Value, "Node template options image value is nil")
			//Options.Name
			assert.NotEmpty(t, nodeTemplate.Options.Name.Name, "Node template options name name is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Name.Type, "Node template options name type is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Name.Value, "Node template options name value is empty")
			//Options.Nvram
			assert.NotNil(t, nodeTemplate.Options.Nvram.Name, "Node template options nvram name is nil")
			assert.NotNil(t, nodeTemplate.Options.Nvram.Type, "Node template options nvram type is nil")
			assert.NotNil(t, nodeTemplate.Options.Nvram.Value, "Node template options nvram value is nil")
			//Options.Ram
			assert.NotEmpty(t, nodeTemplate.Options.Ram.Name, "Node template options ram name is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Ram.Type, "Node template options ram type is empty")
			assert.NotNil(t, nodeTemplate.Options.Ram.Value, "Node template options ram value is nil")
			//Options.Serial
			assert.NotNil(t, nodeTemplate.Options.Serial.Name, "Node template options serial name is nil")
			assert.NotNil(t, nodeTemplate.Options.Serial.Type, "Node template options serial type is nil")
			assert.NotNil(t, nodeTemplate.Options.Serial.Value, "Node template options serial value is nil")
			//Options.Uuid
			assert.NotEmpty(t, nodeTemplate.Options.Uuid.Name, "Node template options uuid name is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Uuid.Type, "Node template options uuid type is empty")
			assert.NotNil(t, nodeTemplate.Options.Uuid.Value, "Node template options uuid value is nil")
			//Options.CpuLimit
			assert.NotEmpty(t, nodeTemplate.Options.Cpulimit.Name, "Node template options cpulimit name is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Cpulimit.Type, "Node template options cpulimit type is empty")
			assert.NotNil(t, nodeTemplate.Options.Cpulimit.Value, "Node template options cpulimit value is nil")
			//Options.Cpu
			assert.NotEmpty(t, nodeTemplate.Options.Cpu.Name, "Node template options cpu name is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Cpu.Type, "Node template options cpu type is empty")
			assert.NotNil(t, nodeTemplate.Options.Cpu.Value, "Node template options cpu value is nil")
			//Options.Firstmac
			assert.NotEmpty(t, nodeTemplate.Options.Firstmac.Name, "Node template options firstmac name is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Firstmac.Type, "Node template options firstmac type is empty")
			assert.NotNil(t, nodeTemplate.Options.Firstmac.Value, "Node template options firstmac value is nil")
			//Options.Qemuversion
			assert.IsType(t, List{}, nodeTemplate.Options.Qemuversion.List, "Node template options qemuversions list is not an actual list")
			assert.NotNil(t, nodeTemplate.Options.Qemuversion.Name, "Node template options qemuversion name is nil")
			assert.NotNil(t, nodeTemplate.Options.Qemuversion.Type, "Node template options qemuversion type is nil")
			assert.NotNil(t, nodeTemplate.Options.Qemuversion.Value, "Node template options qemuversion value is nil")
			//Options.Qemuarch
			assert.IsType(t, List{}, nodeTemplate.Options.Qemuarch.List, "Node template options qemuarch list is not an actual list")
			assert.NotNil(t, nodeTemplate.Options.Qemuarch.Name, "Node template options qemuarch name is nil")
			assert.NotNil(t, nodeTemplate.Options.Qemuarch.Type, "Node template options qemuarch type is nil")
			assert.NotNil(t, nodeTemplate.Options.Qemuarch.Value, "Node template options qemuarch value is nil")
			//Options.Qemunic
			assert.IsType(t, List{}, nodeTemplate.Options.Qemunic.List, "Node template options qemunic list is not an actual list")
			assert.NotNil(t, nodeTemplate.Options.Qemunic.Name, "Node template options qemunic name is nil")
			assert.NotNil(t, nodeTemplate.Options.Qemunic.Type, "Node template options qemunic type is nil")
			assert.NotNil(t, nodeTemplate.Options.Qemunic.Value, "Node template options qemunic value is nil")
			//Options.Qemuoptions
			assert.NotNil(t, nodeTemplate.Options.Qemuoptions.Name, "Node template options qemuoptions name is nil")
			assert.NotNil(t, nodeTemplate.Options.Qemuoptions.Type, "Node template options qemuoptions type is nil")
			assert.NotNil(t, nodeTemplate.Options.Qemuoptions.Value, "Node template options qemuoptions value is nil")
			//Options.Console
			assert.IsType(t, List{}, nodeTemplate.Options.Console.List, "Node template options console list is not an actual list")
			assert.NotEmpty(t, nodeTemplate.Options.Console.Name, "Node template options console name is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Console.Type, "Node template options console type is empty")
			assert.NotEmpty(t, nodeTemplate.Options.Console.Value, "Node template options console value is empty")
			//Options.Rdpuser
			assert.NotNil(t, nodeTemplate.Options.Rdpuser.Name, "Node template options rdpuser name is nil")
			assert.NotNil(t, nodeTemplate.Options.Rdpuser.Type, "Node template options rdpuser type is nil")
			assert.NotNil(t, nodeTemplate.Options.Rdpuser.Value, "Node template options rdpuser value is nil")
			//Options.Rdppassword
			assert.NotNil(t, nodeTemplate.Options.Rdppassword.Name, "Node template options rdppassword name is nil")
			assert.NotNil(t, nodeTemplate.Options.Rdppassword.Type, "Node template options rdppassword type is nil")
			assert.NotNil(t, nodeTemplate.Options.Rdppassword.Value, "Node template options rdppassword value is nil")
			//Qemu
			assert.NotEmpty(t, nodeTemplate.Qemu.Arch, "Node template qemu arch is empty")
			assert.NotNil(t, nodeTemplate.Qemu.Nic, "Node template qemu nic is nil")
			assert.NotEmpty(t, nodeTemplate.Qemu.Options, "Node template qemu options is empty")
		}
		break
	}
}

/*
TestEveNgClient_getFolderContents covers:
	- getFolderContents
*/
func TestEveNgClient_getFolderContents(t *testing.T) {
	eveNgClient, err := NewEveNgClient(viper.GetString("BaseUrl"))
	if !assert.NoError(t, err, "Error while creating API client") {
		return
	}

	err = eveNgClient.SetUsernameAndPassword(viper.GetString("Username"), viper.GetString("Password"))
	if !assert.NoError(t, err, "Error while setting username and password") {
		return
	}

	err = eveNgClient.Login()
	if !assert.NoError(t, err, "Error during login") {
		return
	}
	defer func() {
		err = eveNgClient.Logout()
		if !assert.NoError(t, err, "Error during logout") {
			return
		}
	}()

	folderContents, err := eveNgClient.getFolderContents("")
	if assert.NoError(t, err, "Error during GetFolderContents operation") {
		assert.True(t, len(folderContents.Folders) > 0, "No folders found during GetFolderContents operation")
		assert.True(t, len(folderContents.LabFiles) > 0, "No lab folders found during GetFolderContents operation")
	}
}

/*
TestEveNgClient_GetLabFiles covers:
	- GetLabFiles
*/
func TestEveNgClient_GetLabFiles(t *testing.T) {
	eveNgClient, err := NewEveNgClient(viper.GetString("BaseUrl"))
	if !assert.NoError(t, err, "Error while creating API client") {
		return
	}

	err = eveNgClient.SetUsernameAndPassword(viper.GetString("Username"), viper.GetString("Password"))
	if !assert.NoError(t, err, "Error while setting username and password") {
		return
	}

	err = eveNgClient.Login()
	if !assert.NoError(t, err, "Error during login") {
		return
	}
	defer func() {
		err = eveNgClient.Logout()
		if !assert.NoError(t, err, "Error during logout") {
			return
		}
	}()

	labFiles, err := eveNgClient.GetLabFiles("")
	if assert.NoError(t, err, "Error during GetFolderContents operation") {
		assert.True(t, len(labFiles) > 0, "No folders found during GetFolderContents operation")
	}
}

/*
TestEveNgClient_GetFolders covers:
	- GetFolders
*/
func TestEveNgClient_GetFolders(t *testing.T) {
	eveNgClient, err := NewEveNgClient(viper.GetString("BaseUrl"))
	if !assert.NoError(t, err, "Error while creating API client") {
		return
	}

	err = eveNgClient.SetUsernameAndPassword(viper.GetString("Username"), viper.GetString("Password"))
	if !assert.NoError(t, err, "Error while setting username and password") {
		return
	}

	err = eveNgClient.Login()
	if !assert.NoError(t, err, "Error during login") {
		return
	}
	defer func() {
		err = eveNgClient.Logout()
		if !assert.NoError(t, err, "Error during logout") {
			return
		}
	}()

	folders, err := eveNgClient.GetFolders("")
	if assert.NoError(t, err, "Error during GetFolders operation") {
		assert.True(t, len(folders) > 0, "No folders found during GetFolders operation")
	}
}

/*
TestEveNgClient_GetUserRoles covers:
	- GetUserRoles
*/
func TestEveNgClient_GetUserRoles(t *testing.T) {
	eveNgClient, err := NewEveNgClient(viper.GetString("BaseUrl"))
	if !assert.NoError(t, err, "Error while creating API client") {
		return
	}

	err = eveNgClient.SetUsernameAndPassword(viper.GetString("Username"), viper.GetString("Password"))
	if !assert.NoError(t, err, "Error while setting username and password") {
		return
	}

	err = eveNgClient.Login()
	if !assert.NoError(t, err, "Error during login") {
		return
	}
	defer func() {
		err = eveNgClient.Logout()
		if !assert.NoError(t, err, "Error during logout") {
			return
		}
	}()

	userRoles, err := eveNgClient.GetUserRoles()
	if assert.NoError(t, err, "Error during GetUserRoles operation") {
		if assert.True(t, len(userRoles) > 0, "No user roles found during GetUserRoles") {
			for roleName, roleDescription := range userRoles {
				assert.NotEmpty(t, roleName, "RoleName is empty")
				assert.NotEmpty(t, roleDescription, "RoleDescription is empty")
				break
			}
		}
	}
}

/*
TestEveNgClient_Users covers:
	- AddUser
	- EditUser
	- GetUsers
	- GetUser
	- RemoveUser
*/
func TestEveNgClient_Users(t *testing.T) {
	eveNgClient, err := NewEveNgClient(viper.GetString("BaseUrl"))
	if !assert.NoError(t, err, "Error while creating API client") {
		return
	}

	err = eveNgClient.SetUsernameAndPassword(viper.GetString("Username"), viper.GetString("Password"))
	if !assert.NoError(t, err, "Error while setting username and password") {
		return
	}

	err = eveNgClient.Login()
	if !assert.NoError(t, err, "Error during login") {
		return
	}
	defer func() {
		err = eveNgClient.Logout()
		if !assert.NoError(t, err, "Error during logout") {
			return
		}
	}()

	usersBeforeAdd, err := eveNgClient.GetUsers()

	testusername := "testuser"
	foundUsers := 0
	err = eveNgClient.AddUser(testusername, "Test User", "test@test.test", "testpassword", "admin", "-1", "-1", "internal", 127, "-1", -1, -1)
	if assert.NoError(t, err, "Error during AddUser operation") {
		users, err := eveNgClient.GetUsers()
		if assert.NoError(t, err, "Error during GetUsers") {
			if assert.Greater(t, len(users), len(usersBeforeAdd), "No users found during GetUsers operation") {
				foundUsers = len(users)
			}

			user, err := eveNgClient.GetUser(testusername)
			if assert.NoError(t, err, "Error during GetUsers") {
				assert.Equal(t, "Test User", user.Name, "Username does not match the expected")
				assert.Equal(t, "-1", user.Expiration, "Expiration does not match the expected")
				assert.Equal(t, "-1", user.Pexpiration, "Pexpiration does not match the expected")
				assert.Equal(t, "admin", user.Role, "Role does not match the expected")
				assert.Equal(t, "127", user.Pod, "Pod does not match the expected")
				assert.NotNil(t, user.Session, "Session does not match the expected")
				assert.Equal(t, testusername, user.Username, "Username does not match the expected")
			}
		}
	}
	defer func() {
		err = eveNgClient.RemoveUser(testusername)
		if assert.NoError(t, err, "Error during RemoveUser operation") {
			users, err := eveNgClient.GetUsers()
			if assert.NoError(t, err, "Error during GetUsers") {
				assert.Less(t, len(users), foundUsers, "Users found during GetUsers; none expected")
			}
		}
	}()

	//Editing the user
	err = eveNgClient.EditUser(testusername, "User Test", "changes@changed.user", "changedpassword", "admin", "-1", 125, "202")
	if assert.NoError(t, err, "Error during EditUser operation") {
		user, err := eveNgClient.GetUser(testusername)
		if assert.NoError(t, err, "Error during GetUsers") {
			assert.Equal(t, "User Test", user.Name, "Username does not match the expected")
			assert.Equal(t, "-1", user.Expiration, "Expiration does not match the expected")
			assert.Equal(t, "202", user.Pexpiration, "Pexpiration does not match the expected")
			assert.Equal(t, "admin", user.Role, "Role does not match the expected")
			assert.Equal(t, "125", user.Pod, "Pod does not match the expected")
			assert.NotNil(t, user.Session, "Session does not match the expected")
			assert.Equal(t, testusername, user.Username, "Username does not match the expected")
		}
	}
}

/*
TestEveNgClient_GetNetworkTypes covers:
	- GetNetworkTypes
*/
func TestEveNgClient_GetNetworkTypes(t *testing.T) {
	eveNgClient, err := NewEveNgClient(viper.GetString("BaseUrl"))
	if !assert.NoError(t, err, "Error while creating API client") {
		return
	}

	err = eveNgClient.SetUsernameAndPassword(viper.GetString("Username"), viper.GetString("Password"))
	if !assert.NoError(t, err, "Error while setting username and password") {
		return
	}

	err = eveNgClient.Login()
	if !assert.NoError(t, err, "Error during login") {
		return
	}
	defer func() {
		err = eveNgClient.Logout()
		if !assert.NoError(t, err, "Error during logout") {
			return
		}
	}()

	networkTypes, err := eveNgClient.GetNetworkTypes()
	if assert.NoError(t, err, "Error during GetNetworkTypes operation") {
		if assert.True(t, len(networkTypes) > 0, "No network types found during GetNetworkTypes") {
			for networkType, typeDescription := range networkTypes {
				assert.NotEmpty(t, networkType, "NetworkType is empty")
				assert.NotEmpty(t, typeDescription, "TypeDescription is empty")
				break
			}
		}
	}
}

/*
TestEveNgClient_Folders covers:
	- AddFolder
	- GetFolders
	- RemoveFolder
	- MoveFolder
*/
func TestEveNgClient_Folders(t *testing.T) {
	eveNgClient, err := NewEveNgClient(viper.GetString("BaseUrl"))
	if !assert.NoError(t, err, "Error while creating API client") {
		return
	}

	err = eveNgClient.SetUsernameAndPassword(viper.GetString("Username"), viper.GetString("Password"))
	if !assert.NoError(t, err, "Error while setting username and password") {
		return
	}

	err = eveNgClient.Login()
	if !assert.NoError(t, err, "Error during login") {
		return
	}
	defer func() {
		err = eveNgClient.Logout()
		if !assert.NoError(t, err, "Error during logout") {
			return
		}
	}()

	folderPath := ""
	folderName := "FolderTesting"
	foundFolders := 0

	//Add a new folder
	err = eveNgClient.AddFolder(folderPath, folderName)
	if assert.NoError(t, err, "Error during AddFolderOperation") {
		foldersAfterAdd, err := eveNgClient.GetFolders("")
		if assert.NoError(t, err, "Error during GetFolders operation") {
			assert.True(t, len(foldersAfterAdd) > 0, "No folders found during GetFolders operation")
			foundFolders = len(foldersAfterAdd)
		}

		testFolders, err := eveNgClient.GetFolders("/" + folderName)
		if assert.NoError(t, err, "Error during GetFolders operation") {
			assert.True(t, len(testFolders) > 0, "No folders found insider of TestFolder during GetFolders operation")
		}
	}
	defer func() {
		err = eveNgClient.RemoveFolder(folderName)
		if assert.NoError(t, err, "Error during RemoveFolder operation") {
			foldersAfterRemove, err := eveNgClient.GetFolders("")
			if assert.NoError(t, err, "Error during GetFolders operation") {
				assert.True(t, len(foldersAfterRemove) > 0, "No folders found during GetFolders operation")
				assert.Less(t, len(foldersAfterRemove), foundFolders, "Folder hasn't been removed correctly")
			}
		}
	}()

	newFolderPath := ""
	newFolderName := "FolderTestingMove"
	newFoundFolders := 0

	//Add a second new folder
	err = eveNgClient.AddFolder(newFolderPath, newFolderName)
	if assert.NoError(t, err, "Error during AddFolderOperation") {
		foldersAfterAdd, err := eveNgClient.GetFolders("")
		if assert.NoError(t, err, "Error during GetFolders operation") {
			assert.True(t, len(foldersAfterAdd) > 0, "No folders found during GetFolders operation")
			newFoundFolders = len(foldersAfterAdd)
		}

		testFolders, err := eveNgClient.GetFolders(folderName)
		if assert.NoError(t, err, "Error during GetFolders operation") {
			assert.True(t, len(testFolders) > 0, "No folders found insider of TestFolder during GetFolders operation")
		}
	}
	defer func() {
		err = eveNgClient.RemoveFolder(newFolderName)
		if assert.NoError(t, err, "Error during RemoveFolder operation") {
			foldersAfterRemove, err := eveNgClient.GetFolders("")
			if assert.NoError(t, err, "Error during GetFolders operation") {
				assert.True(t, len(foldersAfterRemove) > 0, "No folders found during GetFolders operation")
				assert.Less(t, len(foldersAfterRemove), newFoundFolders, "Folder hasn't been removed correctly")
			}
		}
	}()

	//Move folder and check if it worked correctly
	foldersInOldFolderBeforeMove, err := eveNgClient.GetFolders("")
	assert.NoError(t, err, "Error during GetFolders operation")
	foldersInNewFolderBeforeMove, err := eveNgClient.GetFolders(newFolderName)
	assert.NoError(t, err, "Error during GetFolders operation")
	err = eveNgClient.MoveFolder(folderName, "/"+newFolderName+"/"+folderName)
	if assert.NoError(t, err, "Error during MoveFolder operation") {
		foldersInOldFolderAfterMove, err := eveNgClient.GetFolders("")
		if assert.NoError(t, err, "Error during FetFolders operation") {
			assert.Less(t, len(foldersInOldFolderAfterMove), len(foldersInOldFolderBeforeMove), "Folder hasn't been moved correctly")
		}
		foldersInNewFolderAfterMove, err := eveNgClient.GetFolders(newFolderName)
		if assert.NoError(t, err, "Error during GetFolders operation") {
			assert.Greater(t, len(foldersInNewFolderAfterMove), len(foldersInNewFolderBeforeMove), "Folder hasn't been moved correctly")
		}
	}
	defer func() {
		foldersInOldFolderBefore2ndMove, err := eveNgClient.GetFolders("")
		assert.NoError(t, err, "Error during GetFolders operation")
		foldersInNewFolderBefore2ndMove, err := eveNgClient.GetFolders(newFolderName)
		assert.NoError(t, err, "Error during GetFolders operation")
		err = eveNgClient.MoveFolder(newFolderName+"/"+folderName, "/"+folderName)
		if assert.NoError(t, err, "Error during MoveFolder operation") {
			foldersInOldFolderAfter2ndMove, err := eveNgClient.GetFolders("")
			if assert.NoError(t, err, "Error during FetFolders operation") {
				assert.Greater(t, len(foldersInOldFolderAfter2ndMove), len(foldersInOldFolderBefore2ndMove), "Folder hasn't been moved correctly")
			}
			foldersInNewFolderAfter2ndMove, err := eveNgClient.GetFolders(newFolderName)
			if assert.NoError(t, err, "Error during GetFolders operation") {
				assert.Less(t, len(foldersInNewFolderAfter2ndMove), len(foldersInNewFolderBefore2ndMove), "Folder hasn't been moved correctly")
			}
		}
	}()
}

/*
TestEveNgClient_Labs covers:
	- AddLab
	- GetLabFiles
	- GetLab
	- EditLab
	- RemoveLab
	- MoveLab
	- AddLabNetwork
	- GetLabNetworks
	- GetLabNetwork
	- RemoveLabNetwork
*/
func TestEveNgClient_Labs(t *testing.T) {
	eveNgClient, err := NewEveNgClient(viper.GetString("BaseUrl"))
	if !assert.NoError(t, err, "Error while creating API client") {
		return
	}

	err = eveNgClient.SetUsernameAndPassword(viper.GetString("Username"), viper.GetString("Password"))
	if !assert.NoError(t, err, "Error while setting username and password") {
		return
	}

	err = eveNgClient.Login()
	if !assert.NoError(t, err, "Error during login") {
		return
	}
	defer func() {
		err = eveNgClient.Logout()
		if !assert.NoError(t, err, "Error during logout") {
			return
		}
	}()

	//Add, Get, Remove Lab test functions
	labFolder := ""
	labName := "LabTesting2"
	labPath := labName + ".unl"
	foundLabs := 0

	err = eveNgClient.AddLab(labFolder, labName, "1", "admin", "A test laboratory", "Test laboratory for unit and integration tests")
	if assert.NoError(t, err, "Error during AddLab operation") {
		labFilesAfterAdd, err := eveNgClient.GetLabFiles(labFolder)
		if assert.NoError(t, err, "Error during GetLabFiles operation") {
			assert.True(t, len(labFilesAfterAdd) > 0, "No lab folders found during GetLabFiles operation")
			foundLabs = len(labFilesAfterAdd)
		}

		testLab, err := eveNgClient.GetLab(labPath)
		if assert.NoError(t, err, "Error during GetLab operation") {
			assert.NotEmpty(t, testLab.Id, "Lab ID is empty")
			assert.Equal(t, "admin", testLab.Author, "TestLab author does not match expected value")
			assert.Equal(t, "Test laboratory for unit and integration tests", testLab.Body, "TestLab body does not match expected value")
			assert.Equal(t, "A test laboratory", testLab.Description, "TestLab description does not match expected value")
			assert.Equal(t, labName+".unl", testLab.Filename, "TestLab filename does not match expected value")
			assert.Equal(t, labName, testLab.Name, "TestLab name does not match expected value")
			assert.Equal(t, "1", testLab.Version, "TestLab version does not match expected value")
		}
	}
	defer func() {
		err = eveNgClient.RemoveLab(labPath)
		if assert.NoError(t, err, "Error during RemoveLab operation") {
			labFilesAfterRemove, err := eveNgClient.GetLabFiles(labFolder)
			if assert.NoError(t, err, "Error during GetLabFiles operation") {
				assert.Less(t, len(labFilesAfterRemove), foundLabs, "Lab has not been removed correctly")
			}
		}
	}()

	renamedTestLabName := "RenmaedLabTesting"

	err = eveNgClient.EditLab(labPath, renamedTestLabName, "2", "testauthor", "A changed test laboratory")
	if assert.NoError(t, err, "Error during EditLab operation") {
		editedTestLab, err := eveNgClient.GetLab(renamedTestLabName + ".unl")
		if assert.NoError(t, err, "Error during GetLab operation") {
			assert.NotEmpty(t, editedTestLab.Id, "Lab ID is empty")
			assert.Equal(t, "testauthor", editedTestLab.Author, "TestLab author does not match expected value")
			assert.Equal(t, "A changed test laboratory", editedTestLab.Description, "TestLab description does not match expected value")
			assert.Equal(t, renamedTestLabName+".unl", editedTestLab.Filename, "TestLab filename does not match expected value")
			assert.Equal(t, renamedTestLabName, editedTestLab.Name, "TestLab name does not match expected value")
			assert.Equal(t, "2", editedTestLab.Version, "TestLab version does not match expected value")
		}
	}
	err = eveNgClient.EditLab(renamedTestLabName+".unl", labName, "2", "testauthor", "A changed test laboratory")
	if assert.NoError(t, err, "Error during EditLab operation") {
		testLab, err := eveNgClient.GetLab(labPath)
		if assert.NoError(t, err, "Error during GetLab operation") {
			assert.Equal(t, labName, testLab.Name, "TestLab name does not match expected value")
		}
	}

	//Add, Get, Remove network test functions
	foundNetworks := 0

	networkId, err := eveNgClient.AddNetwork(labPath, "bridge", "TestNetwork", 69, 420, 1, 0)
	if assert.NoError(t, err, "Error during AddLabNetwork operation") {
		labNetworks, err := eveNgClient.GetNetworks(labPath)
		if assert.NoError(t, err, "Error during GetLabNetworks operation") {
			if assert.True(t, len(labNetworks) > 0, "No lab networks found during GetLabNetworks operation") {
				foundNetworks = len(labNetworks)
				for _, labNetwork := range labNetworks {
					labNetworkDetails, err := eveNgClient.GetNetwork(labPath, labNetwork.Id)
					if !assert.NoError(t, err, "Error during GetLabNetwork operation") {
						assert.NotNil(t, labNetworkDetails.Count, "Network count is nil")
						assert.Equal(t, "TestNetwork", labNetworkDetails.Name, "Network name is empty")
						assert.Equal(t, "bridge", labNetworkDetails.Type, "Network type is empty")
						assert.Equal(t, "420", labNetworkDetails.Top, "Network top is nil")
						assert.Equal(t, "69", labNetworkDetails.Left, "Network left is nil")
						assert.NotEmpty(t, labNetworkDetails.Style, "Network style is empty")
						assert.NotEmpty(t, labNetworkDetails.Linkstyle, "Network linkstyle is empty")
						assert.NotNil(t, labNetworkDetails.Color, "Network color is nil")
						assert.NotNil(t, labNetworkDetails.Label, "Network label is nil")
						assert.Equal(t, 1, labNetworkDetails.Visibility, "Network visibility is nil")
					}
					break
				}
			}
		}
	}
	defer func() {
		err = eveNgClient.RemoveNetwork(labPath, networkId)
		if assert.NoError(t, err, "Error during RemoveLabNetwork operation") {
			labNetworks, err := eveNgClient.GetNetworks(labPath)
			if assert.NoError(t, err, "Error during GetLabNetworks operation") {
				assert.Less(t, len(labNetworks), foundNetworks, "Network has not been removed correctly")
			}
		}
	}()

	//Add a new folder and remove it afterwards
	folderName := "FolderTesting"
	folderPath := "/" + folderName
	err = eveNgClient.AddFolder("", folderName)
	defer func() {
		err = eveNgClient.RemoveFolder("/" + folderName)
	}()

	//MoveLab
	newLabPath := folderPath + "/" + labName + ".unl"

	labFilesInOldFolderBeforeMove, err := eveNgClient.GetLabFiles("")
	labFilesInNewFolderBeforeMove, err := eveNgClient.GetLabFiles(folderPath)
	err = eveNgClient.MoveLab(labPath, folderPath)
	if assert.NoError(t, err, "Error during MoveLab operation") {
		labFilesInOldFolderAfterMove, err := eveNgClient.GetLabFiles("")
		if assert.NoError(t, err, "Error during GetLabFiles operation") {
			assert.Less(t, len(labFilesInOldFolderAfterMove), len(labFilesInOldFolderBeforeMove), "Lab hasn't been moved out the old folder correctly")
		}
		labFilesInNewFolderAfterMove, err := eveNgClient.GetLabFiles(folderPath)
		if assert.NoError(t, err, "Error during GetLabFiles operation in old folder") {
			assert.Greater(t, len(labFilesInNewFolderAfterMove), len(labFilesInNewFolderBeforeMove), "Lab hasn't been moved into the new folder correctly")
		}
	}

	defer func() {
		labFilesInOldFolderBeforeMove, err := eveNgClient.GetLabFiles("")
		assert.NoError(t, err, "Error during GetLabFiles operation")
		labFilesInNewFolderBeforeMove, err := eveNgClient.GetLabFiles(folderPath)
		assert.NoError(t, err, "Error during GetLabFiles operation")
		err = eveNgClient.MoveLab(newLabPath, "")
		if assert.NoError(t, err, "Error during MoveLab operation") {
			labFilesInOldFolderAfterMove, err := eveNgClient.GetLabFiles("")
			if assert.NoError(t, err, "Error during GetLabFiles operation") {
				assert.Greater(t, len(labFilesInOldFolderAfterMove), len(labFilesInOldFolderBeforeMove), "Lab hasn't been moved out the old folder correctly")
			}
			labFilesInNewFolderAfterMove, err := eveNgClient.GetLabFiles(folderPath)
			if assert.NoError(t, err, "Error during GetLabFiles operation in old folder") {
				assert.Less(t, len(labFilesInNewFolderAfterMove), len(labFilesInNewFolderBeforeMove), "Lab hasn't been moved into the new folder correctly")
			}
		}
	}()
}

/*
TestEveNgClient_Nodes covers:
	- AddLabNode
	- GetLabNodes
	- GetLabNode
	- RemoveLabNode
	- StartLabNodes
	- StartLabNode
	- StopLabNodes
	- StopLabNode
	- AddLabNetwork
	- ConnectLabNodeInterfaceToNetwork
	- GetLabNodeInterfaces
	- GetLabTopology
*/
func TestEveNgClient_Nodes(t *testing.T) {
	eveNgClient, err := NewEveNgClient(viper.GetString("BaseUrl"))
	if !assert.NoError(t, err, "Error while creating API client") {
		return
	}

	err = eveNgClient.SetUsernameAndPassword(viper.GetString("Username"), viper.GetString("Password"))
	if !assert.NoError(t, err, "Error while setting username and password") {
		return
	}

	err = eveNgClient.Login()
	if !assert.NoError(t, err, "Error during login") {
		return
	}
	defer func() {
		err = eveNgClient.Logout()
		if !assert.NoError(t, err, "Error during logout") {
			return
		}
	}()

	//Add a new lab
	labFolder := ""
	labName := "NodeTesting"
	labPath := labName + ".unl"
	err = eveNgClient.AddLab(labFolder, labName, "1", "admin", "A test laboratory", "Test laboratory for unit and integration tests")
	defer func() {
		err = eveNgClient.RemoveLab(labPath)
	}()

	//Add a network to the lab
	networkId, err := eveNgClient.AddNetwork(labPath, "nat0", "TestNetwork", 69, 420, 1, 0)
	defer func() {
		labNetworks, _ := eveNgClient.GetNetworks(labPath)
		for _, labNetwork := range labNetworks {
			err = eveNgClient.RemoveNetwork(labPath, labNetwork.Id)
		}
	}()

	//Add nodes to the lab
	foundNodes := 0

	nodeId, err := eveNgClient.AddNode(labPath, "qemu", "asav", "0", 0, "ASA.png", "asav-952-204", "ASAv", 404, 227, 2048, "telnet", 1, "undefined", 8, "", "", "", "", 1)
	if assert.NoError(t, err, "Error during AddLabNode operation") {
		labNodesBeforeRemove, err := eveNgClient.GetNodes(labPath)
		foundNodes = len(labNodesBeforeRemove)
		if assert.NoError(t, err, "Error during GetLabNodes operation") {
			if assert.True(t, len(labNodesBeforeRemove) > 0, "No lab nodes found during GetLabNodes operation") {
				for _, labNode := range labNodesBeforeRemove {
					labNodeDetails, err := eveNgClient.GetNode(labPath, labNode.Id)
					if assert.NoError(t, err, "Error during GetLabNode operation") {
						assert.NotEmpty(t, labNodeDetails.Uuid, "Node uuid does is empty")
						assert.Equal(t, "ASAv", labNodeDetails.Name, "Node name does not match expected value")
						assert.Equal(t, "qemu", labNodeDetails.Type, "Node type does not match expected value")
						assert.Equal(t, 0, labNodeDetails.Status, "Node status does not match expected value")
						assert.Equal(t, "asav", labNodeDetails.Template, "Node template does not match expected value")
						assert.Equal(t, 1, labNodeDetails.Cpu, "Node cpu does not match expected value")
						assert.Equal(t, 2048, labNodeDetails.Ram, "Node ram does not match expected value")
						assert.Equal(t, "asav-952-204", labNodeDetails.Image, "Node image does not match expected value")
						assert.Equal(t, "telnet", labNodeDetails.Console, "Node console does not match expected value")
						assert.Equal(t, 8, labNodeDetails.Ethernet, "Node ethernet does not match expected value")
						assert.Equal(t, 0, labNodeDetails.Delay, "Node Delay does not match expected value")
						assert.Equal(t, "ASA.png", labNodeDetails.Icon, "Node icon does not match expected value")
						assert.NotEmpty(t, labNodeDetails.Url, "Node url is empty")
						assert.Equal(t, 227, labNodeDetails.Top, "Node top does not match expected value")
						assert.Equal(t, 404, labNodeDetails.Left, "Node left does not match expected value")
						assert.Equal(t, "0", labNodeDetails.Config, "Node Config does not match expected value")
						assert.NotEmpty(t, labNodeDetails.Firstmac, "Node firstmac is empty")
					}
					break
				}
			}
		}
	}
	defer func() {
		err = eveNgClient.RemoveNode(labPath, nodeId)
		if assert.NoError(t, err, "Error during RemoveLabNode operation") {
			labNodesAfterRemove, err := eveNgClient.GetNodes(labPath)
			if assert.NoError(t, err, "Error during GetLabFiles operation") {
				assert.Less(t, len(labNodesAfterRemove), foundNodes, "Node has not been removed correctly")
			}
		}
	}()

	//Start single node
	err = eveNgClient.StartNode(labPath, nodeId)
	if assert.NoError(t, err, "Error during StartLabNode operation") {
		labNode, err := eveNgClient.GetNode(labPath, nodeId)
		if assert.NoError(t, err, "Error during GetLabNode operation") {
			assert.Equal(t, 2, labNode.Status, "Starting LabNode didn't work")
		}
	}

	//Stop single node
	err = eveNgClient.StopNode(labPath, nodeId)
	if assert.NoError(t, err, "Error during StopLabNode operation") {
		labNode, err := eveNgClient.GetNode(labPath, nodeId)
		if assert.NoError(t, err, "Error during GetLabNode operation") {
			assert.Equal(t, 0, labNode.Status, "Stopping LabNode didn't work")
		}
	}

	//Start all nodes
	err = eveNgClient.StartNodes(labPath)
	if assert.NoError(t, err, "Error during StartLabNodes operation") {
		labNodes, err := eveNgClient.GetNodes(labPath)
		if assert.NoError(t, err, "Error during GetLabNodes operation") {
			for _, labNode := range labNodes {
				assert.Equal(t, 2, labNode.Status, "Node "+strconv.Itoa(labNode.Id)+" wasn't started correctly")
			}
		}
	}
	defer func() {
		//Stop all nodes after execution
		err = eveNgClient.StopNodes(labPath)
		if assert.NoError(t, err, "Error during StopLabNodes operation") {
			labNodes, err := eveNgClient.GetNodes(labPath)
			if assert.NoError(t, err, "Error during GetLabNodes operation") {
				for _, labNode := range labNodes {
					assert.Equal(t, 0, labNode.Status, "Node "+strconv.Itoa(labNode.Id)+" wasn't stopped correctly")
				}
			}
		}
	}()

	//Connect node interface to network
	err = eveNgClient.ConnectNodeInterfaceToNetwork(labPath, nodeId, 1, networkId)
	if assert.NoError(t, err, "Error during ConnectLabNodeInterfaceToNetwork operation") {
		nodeInterfaces, err := eveNgClient.GetNodeInterfaces(labPath, nodeId)
		if assert.NoError(t, err, "Error during GetLabNodeInterfaces operation") {
			for _, ethernetInterface := range nodeInterfaces.Ethernet {
				if ethernetInterface.Name == "Eth1" {
					assert.Equal(t, networkId, ethernetInterface.NetworkId, "Network was not correctly added to NodeInterface")
				}
			}
		}
	}

	//Get lab topology
	labTopology, err := eveNgClient.GetTopology(labPath)
	if assert.NoError(t, err, "Error during GetLabTopology operation") {
		if assert.True(t, len(labTopology) > 0, "No topologies found during GetLabTopology") {
			for _, topologyPoint := range labTopology {
				assert.NotEmpty(t, topologyPoint.Destination, "Destination is empty")
				assert.NotNil(t, topologyPoint.DestinationLabel, "DestinationLabel is nil")
				assert.NotEmpty(t, topologyPoint.DestinationType, "DestinationType is empty")
				assert.NotNil(t, topologyPoint.DestinationInterfaceId, "DestinationInterfaceId is nil")
				assert.NotNil(t, topologyPoint.DestinationNodename, "DestinationNodename is nil")
				assert.NotNil(t, topologyPoint.DestinationSuspend, "DestinationSuspend is nil")
				assert.NotNil(t, topologyPoint.DestinationDelay, "DestinationDelay is nil")
				assert.NotNil(t, topologyPoint.DestinationLoss, "DestinationLoss is nil")
				assert.NotNil(t, topologyPoint.DestinationBandwidth, "DestinationBandwidth is nil")
				assert.NotNil(t, topologyPoint.DestinationJitter, "DestinationJitter is nil")
				assert.NotEmpty(t, topologyPoint.Source, "Source is empty")
				assert.NotNil(t, topologyPoint.SourceLabel, "SourceLabel is nil")
				assert.NotEmpty(t, topologyPoint.SourceType, "SourceType is empty")
				assert.NotNil(t, topologyPoint.SourceNodename, "SourceNodenam is empty")
				assert.NotNil(t, topologyPoint.SourceInterfaceId, "SourceInterfaceId is nil")
				assert.NotNil(t, topologyPoint.SourceSuspend, "SourceSuspend is nil")
				assert.NotNil(t, topologyPoint.SourceDelay, "SourceDelay is nil")
				assert.NotNil(t, topologyPoint.SourceLoss, "SourceLoss is nil")
				assert.NotNil(t, topologyPoint.SourceBandwidth, "SourceBandwidth is nil")
				assert.NotNil(t, topologyPoint.SourceJitter, "SourceJitter is nil")
				assert.NotEmpty(t, topologyPoint.Type, "Type is empty")
				assert.NotNil(t, topologyPoint.NetworkId, "NetworkId is nil")
				assert.NotNil(t, topologyPoint.Style, "Style is nil")
				assert.NotNil(t, topologyPoint.Linkstyle, "Linkstyle is nil")
				assert.NotNil(t, topologyPoint.Label, "Label is nil")
				assert.NotNil(t, topologyPoint.Color, "Color is nil")
				break
			}
		}
	}
}

/*
TestEveNgClient_ExportWipeNodes covers:
	- ExportLabNode
	- ExportLabNodes
	- WipeLabNode
	- WipeLabNodes
*/
func TestEveNgClient_ExportWipeNodes(t *testing.T) {
	eveNgClient, err := NewEveNgClient(viper.GetString("BaseUrl"))
	if !assert.NoError(t, err, "Error while creating API client") {
		return
	}

	err = eveNgClient.SetUsernameAndPassword(viper.GetString("Username"), viper.GetString("Password"))
	if !assert.NoError(t, err, "Error while setting username and password") {
		return
	}

	err = eveNgClient.Login()
	if !assert.NoError(t, err, "Error during login") {
		return
	}
	defer func() {
		err = eveNgClient.Logout()
		if !assert.NoError(t, err, "Error during logout") {
			return
		}
	}()

	//Add a new lab
	labFolder := ""
	labName := "NodeWipeAndExportTesting"
	labPath := labName + ".unl"
	err = eveNgClient.AddLab(labFolder, labName, "1", "admin", "A test laboratory", "Test laboratory for unit and integration tests")
	defer func() {
		err = eveNgClient.RemoveLab(labPath)
	}()

	nodeId, _ := eveNgClient.AddNode(labPath, "qemu", "asav", "0", 0, "ASA.png", "asav-952-204", "ASAv", 404, 227, 2048, "telnet", 1, "undefined", 8, "", "", "", "", 1)
	defer func() {
		_ = eveNgClient.RemoveNode(labPath, nodeId)
	}()

	//Export the nodes start configs
	err = eveNgClient.ExportNode(labPath, nodeId)
	assert.NoError(t, err, "Error during ExportLabNode operation")

	//Export all nodes start configs
	err = eveNgClient.ExportNodes(labPath)
	assert.NoError(t, err, "Error during ExportLabNodes operation")

	//Wipe the node
	err = eveNgClient.WipeNode(labPath, nodeId)
	assert.NoError(t, err, "Error during WipeNode operation")

	//Wipe all nodes
	err = eveNgClient.WipeNodes(labPath)
	assert.NoError(t, err, "Error during WipeNodes operation")
}
