package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int64 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 string }

func main() {
    var v1 uint = 255
    v2 := make(map[string]bool)
    var v3 uint16 = 255
    v4 := []int{127, -58}
    var v5 uint = 18446744073709551615
    v1 = v5
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: 1000, F1: -1, F2: 9223372036854775807}
        _ = fi6
        _ = v7
    }
    v4 = append(v4, -9223372036854775808)
    v8 := GenS2{F0: 2147483647, F1: 689}
    if ("a" <= "hello") {
        var v9 string = ("" + ("hello" + ""))
        _ = v9
    }
    switch 4294967295 {
    case 4294967295:
        var v10 uint32 = 800
        _ = v10
    case -232:
        v3 = (v3 | 828)
    }
    var v11 int32 = 1
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            fmt.Println(fi13, v1, v5)
            _ = len(v4)
            var v14 uint32 = 4294967295
            v15 := GenS1{F0: 1000, F1: 70, F2: -484}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        fmt.Println(v3, v1)
        v17 := GenS3{F0: ""}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi16
        _ = v17
    }
    var v18 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
