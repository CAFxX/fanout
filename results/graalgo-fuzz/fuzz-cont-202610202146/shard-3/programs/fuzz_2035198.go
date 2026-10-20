package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []string{"a", (("foo bar" + ("a" + "a")) + ("a" + "")), ""}
    var v2 int32 = 2147483647
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    var v5 bool = (0 >= -1)
    var v6 uint8 = 127
    var v7 uint16 = 19
    v7 = v7
    _ = len(v4)
    v5 = v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "hello" {
    case "foo bar":
        if v5 {
            var v8 uint32 = 2
            var v9 uint8 = 127
            for fi10 := 0; fi10 < 5; fi10++ {
                switch ("foo bar" + ("foo bar" + "foo bar")) {
                case "\n\t":
                    var v11 uint32 = v8
                    _ = v11
                case "":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "hello":
                    var v12 int64 = -1
                    _ = v12
                }
                v13 := GenS1{F0: (v8 >> ((uint(v7)) & 63)), F1: "a"}
                _ = fi10
                _ = v13
            }
            _ = v8
            _ = v9
        } else {
            switch 9223372036854775807 {
            case 32767:
                switch -87 {
                case 1:
                    v14 := GenS1{F0: 4294967295, F1: ("" + "foo bar")}
                    _ = v14
                default:
                    v15 := GenS1{F0: 2147483647, F1: "hello"}
                    _ = v15
                }
            case -2147483648:
                v16 := GenS1{F0: 1, F1: (("" + ("hello" + ((("hello" + "hello") + "foo bar") + "hello"))) + "hello")}
                _ = v16
            case 410:
                v5 = (0 < 65535)
            default:
                delete(v3, 604)
            }
        }
    case "hello":
        var v17 float64 = 0.1
        _ = v17
    case "":
        fmt.Println(v5, v5)
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        v19 := GenS1{F0: 1, F1: ("hello" + "foo bar")}
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := GenS1{F0: 2147483647, F1: ("a" + "a")}
            _ = fi21
            _ = v22
        }
        v23 := GenS2{F0: 32767}
        fmt.Println(v7, v5)
        _ = fi20
        _ = v23
    }
    var v24 int16 = 0
    switch 32768 {
    case 128:
        v25 := GenS1{F0: 42, F1: "hello"}
        _ = v25
    case 32767:
        delete(v4, 550)
    }
    v24 -= v24
    for fi26 := 0; fi26 < 15; fi26++ {
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 string = ("" + ("" + "foo bar"))
            var v29 uint16 = v7
            _ = fi27
            _ = v28
            _ = v29
        }
        var v30 int32 = v2
        _ = fi26
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
