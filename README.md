# WooCommerce API - Golang Client
A Golang wrapper for the WooCommerce REST API

[![build status](https://travis-ci.org/hpcslag/wc-api-golang.svg)](https://travis-ci.org/hpcslag/wc-api-golang)

# Descrption
Only support `http` protocol (not `https`), and API Version is `v2`, follow OAuth 1.0a rules and condition to make request to `WooCommerce`.

## Installation

```bash
$ go get github.com/hpcslag/wc-api-golang/woocommerce
```

## Getting started

Generate API credentials (Consumer Key & Consumer Secret) following this instructions <http://docs.woocommerce.com/document/woocommerce-rest-api/>
.

Check out the WooCommerce API endpoints and data that can be manipulated in <https://woocommerce.github.io/woocommerce-rest-api-docs/>.

## Setup

Setup for the new WP REST API integration (WooCommerce 2.6 or later):

```golang
import (
  wc "github.com/hpcslag/wc-api-golang/woocommerce"
)

client := wc.NewClient(
    "http://path_to_wordpress.com", 
    "ck_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", 
    "cs_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
)
```

### Options

|       Option      |   Type   | Required |                Description                 |
| ----------------- | -------- | -------- | ------------------------------------------ |
| `domain`             | `string` | yes      | Your Store URL, example: http://woo.dev/   |
| `consumer_key`    | `string` | yes      | Your API consumer key                      |
| `consumer_secret` | `string` | yes      | Your API consumer secret                   |


## Methods

|    Params    |      Type      |                         Description                          |
| ------------ | -------------- | ------------------------------------------------------------ |
| `path`       | `string`       | WooCommerce API endpoint, example: `customers` or `order/12` |
| `param`      | `*url.Values`  | Only for GET and DELETE, request query string                |
| `data`       | `interface{}`  | Only for POST and PUT, data that will be converted to JSON   |

### GET

```golang
responseString := client.Get( path , param)

//example:
responseString := client.Get("products",nil)
```

### POST

```golang
responseString := client.Post( path , data)

//example:
var data map[string]interface{} = map[string]interface{}{
    "code":"10off",
    "discount_type": "percent",
    "amount": "10",
    "individual_use": true,
    "exclude_sale_items": true,
    "minimum_amount": "100.00",
}

responseString := client.Post("coupons",data)
```

### PUT

```golang
responseString := client.Put( path , data)
```

### DELETE

```golang
responseString := client.Delete( path , param)
```

### OPTIONS

```golang
responseString := client.Options( path , param)
```

#### Response

Methods will response `io.ReadCloser` type.

You can read by buffer decode:
```golang

body := client.Get("products",nil)

buf := new(bytes.Buffer)
buf.ReadFrom(body)
resp := buf.String()

fmt.Printf(resp)

defer body.Close()
```

also can use `map` type to receive json(unmarshal):
```golang
jsonDecoder := json.NewDecoder(body)

var data map[string]interface{}

//sometime response json is array:
//var data []map[string]interface{}

if err := jsonDecoder.Decode(&data); err != nil && err != io.EOF {
    t.Fatal(err)
}

fmt.Println(body)

defer body.Close()
```
