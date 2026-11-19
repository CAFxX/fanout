package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int16; F1 uint16; F2 int16 }
type GenS3 struct { F0 int8; F1 uint16; F2 int32 }
func genFn0() int {
    v4 := GenS1{F0: 1, F1: 1}
    v5 := GenS1{F0: -1, F1: 0}
    var v6 uint64 = 0
    v6 = (v6 & 4294967295)
    _ = v4
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = ("a" + ("" + "foo bar"))
    v3 := []string{}
    v2 = "a"
    switch v2 {
    case "":
        fmt.Println(v2)
    default:
        var v7 string = ""
        _ = v7
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v8 := GenS1{F0: -9223372036854775808, F1: -2147483648}
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS2{F0: -32768, F1: 42, F2: -35}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(v2)
        _ = fi11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 int16 = 924
        if false {
            if (4294967295 < 2) {
                v14 := GenS3{F0: -12, F1: 161, F2: 0}
                _ = len(v1)
                _ = v14
            }
        }
        _ = fi12
        _ = v13
    }
    v2 = v2
    if false {
        v1["foo bar"] = (true || false)
        var v15 uint32 = 42
        _ = v15
    }
    var v16 int64 = -36
    v2 = "hello"
    fmt.Println(v16)
    v16 = (^(v16))
    _ = len(v3)
    v17 := GenS3{F0: 1, F1: 2, F2: 443}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
