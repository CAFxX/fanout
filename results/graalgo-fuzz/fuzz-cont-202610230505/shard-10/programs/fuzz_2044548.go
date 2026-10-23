package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int32; F1 int8 }
func genFn0(p0 uint) bool {
    var v5 float64 = 0.1
    p0 = (p0 + p0)
    _ = p0
    _ = v5
    return (false && true)
}

func main() {
    var v1 uint = 2
    v2 := make(map[string]uint32)
    var v3 bool = (!(!false))
    v4 := []string{(("\n\t" + "hello") + ("" + "foo bar")), ("" + "foo bar")}
    v6 := v2[(("\n\t" + "foo bar") + ("" + "a"))]
    if v3 {
        switch "foo bar" {
        case "\n\t":
            var v7 int16 = -282
            _ = v7
        case "":
            v3 = true
        case "a":
            fmt.Println(v3, v3, v3)
        }
        v8 := GenS1{F0: -788, F1: 32767}
        v6 -= 4294967295
        _ = v8
    } else {
        var v9 int = 9223372036854775807
        var v10 float32 = -2.25
        v11 := GenS1{F0: (v9 / ((v9) | 1)), F1: v9}
        _ = v10
        _ = v11
        _ = v9
    }
    v3 = v3
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    if true {
        fmt.Println(v3, v3)
        var v12 string = ""
        var v13 uint32 = v6
        _ = v12
        _ = v13
    } else {
        v1 = (v1 * (v1 - v1))
        v1 = v1
    }
    if true {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 float32 = -2.25
            v1 = v1
            _ = fi14
            _ = v15
        }
        fmt.Println(v1, v1)
        v4 = append(v4, ("hello" + ("\n\t" + (((("hello" + ("a" + "a")) + "a") + "foo bar") + "a"))))
    } else {
        switch 127 {
        case 481:
            for fi16 := 0; fi16 < 4; fi16++ {
                for fi17 := 0; fi17 < 7; fi17++ {
                    v1 -= v1
                    var v18 uint64 = 2
                    v18 = (v18 / (((v18 ^ (v18 & v18))) | 1))
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        case -9223372036854775808:
            v4 = append(v4, "a")
        case -18:
            v19 := GenS2{F0: 2, F1: "hello"}
            _ = v19
        }
    }
    var v20 uint = 65535
    v20 -= 2
    switch "a" {
    case "foo bar":
        v3 = ((!true) && v3)
    case "":
        var v21 uint32 = 4294967295
        _ = v21
    }
    delete(v2, ((("\n\t" + "a") + "") + "hello"))
    if (0 != 127) {
        fmt.Println(v20, v6)
        v22 := v2["hello"]
        switch "" {
        case "":
            switch "\n\t" {
            case "hello":
                v23 := GenS1{F0: -348, F1: 4294967295}
                _ = v23
            case "foo bar":
                v24 := GenS2{F0: 713, F1: ""}
                _ = v24
            default:
                v6 -= v6
            }
        default:
            if (!(968 < 127)) {
                v1 = v20
                v25 := GenS3{F0: 223, F1: 1}
                _ = v25
            }
        }
        _ = v22
    } else {
        fmt.Println(v6, v3)
    }
    v1 = (v1 | v20)
    fmt.Println(v3, v6, v20)
    v26 := GenS2{F0: 2, F1: (((("foo bar" + "a") + "") + "") + "hello")}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v6
}
