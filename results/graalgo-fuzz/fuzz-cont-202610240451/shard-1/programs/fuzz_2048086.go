package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint; F1 uint16 }

func main() {
    var v1 int32 = 2147483647
    var v2 uint64 = 334
    var v3 int32 = (v1 - 0)
    var v4 uint16 = 520
    v5 := []string{"hello", "foo bar", ("" + (("hello" + "") + "foo bar")), "hello"}
    if len(v5) > 0 {
        v5[0] = ("a" + "")
    }
    v3 = (v1 ^ -126)
    fmt.Println(v2)
    v3 += v1
    v3 = v3
    var v6 uint8 = 2
    v4 = v4
    var v7 bool = false
    for fi8 := 0; fi8 < 8; fi8++ {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v4 = v4
        _ = fi8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        if false {
            if len(v5) > 0 {
                v5[0] = "hello"
            }
            var v10 int16 = 0
            _ = v10
        } else {
            var v11 bool = v7
            _ = v11
        }
        var v12 int8 = -78
        switch "\n\t" {
        case "foo bar":
            v6 += v6
        case "a":
            v13 := GenS1{F0: "hello", F1: v7}
            _ = v13
        default:
            v14 := GenS2{F0: 989, F1: (v4 % ((v4) | 1))}
            _ = v14
        }
        _ = fi9
        _ = v12
    }
    var v15 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
