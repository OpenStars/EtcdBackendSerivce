// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"github.com/apache/thrift/lib/go/thrift"
	"OpenStars/Platform/MarketPlace"
)

var _ = MarketPlace.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TDataResult getData(i64 key)")
  fmt.Fprintln(os.Stderr, "  TErrorCode putData(i64 key, TMarketPlaceItem data)")
  fmt.Fprintln(os.Stderr, "  TErrorCode removeData(i64 key)")
  fmt.Fprintln(os.Stderr, "  TListDataResult getListData( key)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := MarketPlace.NewTDataServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "getData":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetData requires 1 args")
      flag.Usage()
    }
    argvalue0, err25 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err25 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetData(context.Background(), value0))
    fmt.Print("\n")
    break
  case "putData":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "PutData requires 2 args")
      flag.Usage()
    }
    argvalue0, err26 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err26 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    arg27 := flag.Arg(2)
    mbTrans28 := thrift.NewTMemoryBufferLen(len(arg27))
    defer mbTrans28.Close()
    _, err29 := mbTrans28.WriteString(arg27)
    if err29 != nil {
      Usage()
      return
    }
    factory30 := thrift.NewTJSONProtocolFactory()
    jsProt31 := factory30.GetProtocol(mbTrans28)
    argvalue1 := MarketPlace.NewTMarketPlaceItem()
    err32 := argvalue1.Read(jsProt31)
    if err32 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.PutData(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "removeData":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RemoveData requires 1 args")
      flag.Usage()
    }
    argvalue0, err33 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err33 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.RemoveData(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getListData":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetListData requires 1 args")
      flag.Usage()
    }
    arg34 := flag.Arg(1)
    mbTrans35 := thrift.NewTMemoryBufferLen(len(arg34))
    defer mbTrans35.Close()
    _, err36 := mbTrans35.WriteString(arg34)
    if err36 != nil { 
      Usage()
      return
    }
    factory37 := thrift.NewTJSONProtocolFactory()
    jsProt38 := factory37.GetProtocol(mbTrans35)
    containerStruct0 := MarketPlace.NewTDataServiceGetListDataArgs()
    err39 := containerStruct0.ReadField1(jsProt38)
    if err39 != nil {
      Usage()
      return
    }
    argvalue0 := containerStruct0.Key
    value0 := argvalue0
    fmt.Print(client.GetListData(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
