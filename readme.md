# eve-ng-restapi-go-client

[![Go Report Card](https://goreportcard.com/badge/github.com/inexio/eve-ng-restapi-go-client)](https://goreportcard.com/report/github.com/inexio/eve-ng-restapi-go-client)
[![GitHub license](https://img.shields.io/badge/license-BSD-blue.svg)](https://github.com/inexio/check_eve_ng/blob/master/LICENSE)
[![GitHub code style](https://img.shields.io/badge/code%20style-uber--go-brightgreen)](https://github.com/uber-go/guide/blob/master/style.md)
[![GoDoc doc](https://img.shields.io/badge/godoc-reference-blue)](https://godoc.org/github.com/inexio/eve-ng-restapi-go-client)

## Description

Golang package - client library for the [Eve-NG  REST API](https://www.eve-ng.net/index.php/documentation/howtos/how-to-eve-ng-api/).

This client is an open-source library for communicating with Eve-NG Server applications that are written in php.

## Code Style

This project was written according to the **[uber-go](https://github.com/uber-go/guide/blob/master/style.md)** coding style.

## Features

### Full Eve-NG Support

This client allows you to create:

- Folders

- Labs

- Nodes

- Networks

- Users

and also enables you to:

- Move folders and labs

- Edit existing labs, nodes, networks and users

- Connect nodes to networks

- Start / stop nodes (single or bulk operations)

- Wipe / export node starting configurations

- Check the system status

## Requirements

Requires a running instance of Eve-NG.

Further information about Eve-NG can be found [here](https://www.eve-ng.net)

To check if your setup works, follow the steps provided in the **'Tests'** section of this document. 

## Installation

```
go get github.com/inexio/eve-ng-restapi-go-client
```

or 

```
git clone https://github.com/inexio/eve-ng-restapi-go-client.git
```

## Setup

After installing the library you have to either declare a path to your config file or set certain environment variables for the client to work.

These can be set as follows:

#### Config File

In the **client.go** file, in the **init()** function you can see the following lines of code:

```
viper.AddConfigPath("config/")
viper.SetConfigType("yaml")
viper.SetConfigName("eve-ng-api")
```

ConfigPath is relativ to the package location.
ConfigType and name can also be changed to match your desired type of config.

#### Environment Variables

Also in the **client.go** file, in the **init()** function you will find:

```
//Set env var prefix to only match certain vars
viper.SetEnvPrefix("eve_ng_api")
```

SetEnvPrefix can be changed to whatever prefix you prefer to have in your environment vars.

The needed environment vars can then be added as follows:

```
export EVE_NG_API_BASEURL="<your_base_url>"
export EVE_NG_API_USERNAME="<your_username>"
export EVE_NG_API_PASSWORD="<your_password>"
```

## Usage

The following section will show you how to create a lab and do various operations in it.

```go
baseUrl := "https://<your eve-ng server>"
username := <username>
password := <password>
eveNgClient, err := NewEveNgClient(baseUrl)
_ := eveNgClient.SetUsernameAndPassword(username, password)

_ = eveNgClient.Login()
defer func() {
  _ = eveNgClient.Logout()
}()

_ = eveNgClient.AddFolder("/", "TestFolder")

labFolder := "/TestFolder" //path to the desired folder
labName := "TestLaboratory" //name of the laboratory

_ = eveNgClient.AddLab(labFolder, labName, "1", "admin", "A test laboratory", "Test laboratory for unit and integration tests")

networkId, _ = eveNgClient.AddNetwork("/TestFolder/TestLaboratory.unl", "nat0", "TestNetwork", "69", "420", 1, 0)

nodeId, _ := eveNgClient.AddNode("/TestFolder/TestLaboratory.unl", "qemu", "veos", "0", 0, "AristaSW.png", "veos-4.16.14M", "vEOS", "420", "69", "512", "telnet", 1, "undefined", 4, "", "", "", "", 1)

_ = eveNgClient.ConnectNodeInterfaceToNetwork("/TestFolder/TestLaboratory.unl", nodeId, 1, networkId)

_ = eveNgClient.StartNode("/TestFolder/TestLaboratory.unl", nodeId)

labTopology, _ := eveNgClient.GetTopology("/TestFolder/TestLaboratory.unl")

_ = eveNgClient.StartNodes("/TestFolder/TestLaboratory.unl")
```

After running the code above, the lab you just created should look like this when viewed from the web-interface

![](https://user-images.githubusercontent.com/55132811/74844336-99f7a980-532d-11ea-966f-1611f4705102.png)

## Tests

The library comes with a few unit and integrations tests. To use these tests you have to either use a config file giving the client the correct base-url, username and password or set certain environment variables.

In order to run these test, run the follwing command inside root directory of this repository:

```
go test
```

If you want to check if your setup works, run:

```
go test -run TestEveNgClient_Nodes
```

## Getting Help

If there are any problems or something does not work as intended, open an issue on GitHub.

## Contribution

Contributions to the project are welcome.

We are looking forward to your bug reports, suggestions and fixes.

If you want to make any contributions make sure your go report does match up with our projects score of **A+**.

When you contribute make sure your code is conform to the **uber-go** coding style.

Happy Coding!
