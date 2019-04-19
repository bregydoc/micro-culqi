![logo](./uculqilogo.png)

# MicroCulqi

Microculqi uCulqi is a microservice to mount a simple grpc service to make charges and save and send invoices to your customers. The first principle of this microservice is make very simple install your [Culqi](https://culqi.com/) backend. uCulqi uses [Mailjet](https://www.mailjet.com/) as SMTP service, in the future Im going to implement another backends for this task.

## Features
- Charges with Culqi
- Autogenerated Invoice from your order
- gRPC of the box
- Containable
- Auto-send invoices email

## Launching service

Microculqi (uCulqi) are compiled as a docker image and you can pull from docker hub with the following command:

```shell
docker pull bregymr/uculqi
```

### Basic Usage

You need two files to run your uculqi instance: your uculqi.config.yml file and a simple template.html file:

#### uculqi.config.yml

A minimal configuration example below, note that in this example doesn't have a mailjet configuration then uculqi can't to auto send emails, for a full example please read the Usage section. 

```yaml
company:
  name: "uCulqi Corp"
  email_support: "support@example.com"
culqi:
  merchant_code: "<YOUR MERCHANT CODE>"
  api_key: "<YOUR PRIVATE API KEY>"
```

[Here](https://github.com/bregydoc/micro-culqi/blob/master/examples/server-mount/uculqi.config.yml) you can see a complete config file.

#### template.html

The template html is an text-template based on [go template package](https://golang.org/pkg/text/template/), you can build your own invoice template based on your invoice information. You can see your allowed template values [here](https://github.com/bregydoc/micro-culqi/blob/master/invoice_template.go).

A simple and basic example of the template is the following html:

```html
<html>
  <body>
    <h1> Invoice {{.ID}} </h1>
    Total: {{.TotalCost}}
  </body>
</html>
```



With this requirements satisfied you can to run uCulqi with:

```shell
docker run -p 18000:18000 -v $(pwd)/uculqi.config.yml:/etc/uculqi/uculqi.config.yml -v $(pwd)/template.html:/etc/uculqi/template.html bregymr/uculqi
```

Well, we can a uculqi instance running but this doesn't have a complete features, for that we need to add some configurations, continue reading to know more.

### Full Usage

**TL:DR**: Full example [here](https://github.com/bregydoc/micro-culqi/tree/master/examples/server-mount).

The full usage of uCulqi is creating an STL keys to makes a secure grpc connection. I use openssl to create these keys, below you can see an example of this command. Note the CN (Common Name) refers to your domain, in a dev enviroment you can use `localhost` like in the example.

```shell
openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj '/CN=localhost'
```

Next, I recomend use a docker-compose file to launch your service, here you can see and example of the docker-compose.yml.

```yml
version: "3.7"

services:
  uculqi:
    image: bregymr/uculqi
    ports:
      - 18000:18000
    secrets:
      - cert
      - key
    volumes:
      - ./uculqi.config.yml:/etc/uculqi/uculqi.config.yml
      - ./template.html:/etc/uculqi/template.html
secrets:
  cert:
    file: ./server.crt
  key:
    file: ./server.key
```

Note that routes of uculqi.config.yml and template.html are necessary and you can configure it.

## Connecting to your service

To connect to your service you need to install the go micro-culqi package.

Install with:

```shell
go get -u github.com/bregydoc/micro-culqi
```

*Note: currently Iam only compiled a golang client with the .proto files, but you can compiled it to another languages, [here](https://grpc.io/) you can see how. The uCulqi proto file is [here](https://github.com/bregydoc/micro-culqi/blob/master/proto/system.proto).*

### Usage

Basic example:

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bregydoc/micro-culqi/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	c, _ := credentials.NewClientTLSFromFile(
		"/path/to/your/credentials/server.crt",
		"",
	)
	
	conn, err := grpc.Dial("localhost:18000", grpc.WithTransportCredentials(c))
	// error handling omitted
	
	client := pculqi.NewUCulqiClient(conn)

	invoice, err := client.ExecuteCharge(context.TODO(), &pculqi.MinimalInvoice{
		Currency: pculqi.AvailableCurrency_PEN,
		Token:    "<YOUR CULQI GENERATED TOKEN>",
		Products: []*pculqi.Product{
			{Name: "Dry Herb Vaporizer", Currency: pculqi.AvailableCurrency_PEN, Price: 20.0},
		},
		Email: "customer@example.com",
	})
  
	if err != nil {
		panic(err)
	}
	
	d, _ := json.Marshal(invoice)
	
	fmt.Println(string(d))
}
```



Another process than you can execute with uCulqi:

```go
func (s *Service) ExecuteCharge(c context.Context, params *pculqi.MinimalInvoice) (*pculqi.Invoice, error){}

func (s *Service) GetInvoiceByID(c context.Context, params *pculqi.InvoiceID) (*pculqi.Invoice, error){}

func (s *Service) UpdateEmailTemplate(c context.Context, params *pculqi.TemplateData) (*pculqi.IsOk, error){}

func (s *Service) ExecuteChargeWithOrder(c context.Context, order *pculqi.Order) (*pculqi.Invoice, error){}
```





## Todo

- [ ] Improve documentation
- [ ] Create a REST API to open access to your invoices
- [ ] Write more test files



## Contribution

You can contribute to this project creating a pull request or writing an issue.



