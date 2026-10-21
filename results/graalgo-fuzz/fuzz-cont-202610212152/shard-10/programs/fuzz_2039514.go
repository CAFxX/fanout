package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int64; F1 string }

func main() {
    var v1 string = "hello"
    var v2 int8 = -128
    v3 := []string{v1, v1}
    if true {
        v1 = v1
    }
    _ = len(v3)
    for fi4 := 0; fi4 < 14; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 bool = (32767 < 42)
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v1 = v1
    fmt.Println(v2)
    v2 = (v2 % ((v2) | 1))
    v7 := GenS1{F0: 127, F1: v2}
    if true {
        var v8 float64 = 1e10
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 uint32 = 716
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 0, F1: v2}
        _ = v11
        _ = v8
    }
    v2 = (v2 >> ((65535) & 63))
    fmt.Println(v2, v1)
    var v12 uint = 1
    v13 := GenS2{F0: 42, F1: ((v1 + v1) + "foo bar")}
    v12 = v12
    v2 = (v2 ^ (v2 % (((v2 * v2)) | 1)))
    fmt.Println(v12, v12)
    v12 += (v12 * 255)
    _ = len(v3)
    var v14 uint16 = 328
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
}
