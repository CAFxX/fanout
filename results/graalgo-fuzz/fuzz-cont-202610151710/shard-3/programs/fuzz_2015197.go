package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint32; F1 int32 }

func main() {
    var v1 uint32 = 42
    var v2 int8 = 100
    v3 := []int{488, 127, 128}
    v4 := make(map[int]uint32)
    v5 := make(map[int]bool)
    v6 := make(map[uint32]uint32)
    var v7 int8 = 127
    var v8 uint8 = 1
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            v8 += v8
            delete(v6, v1)
            _ = fi10
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            switch "hello" {
            case "foo bar":
                v1 -= v1
            case "hello":
                delete(v5, fi11)
            case "":
                v1 = 347
            }
            var v12 int8 = v7
            _ = fi11
            _ = v12
        }
        _ = fi9
    }
    v2 += v7
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v13 uint = 18446744073709551615
        v14 := GenS1{F0: 2}
        v7 = v7
        _ = v13
        _ = v14
    }
    if (v1 >= v1) {
        v3 = append(v3, -996)
        for fi15 := 0; fi15 < 15; fi15++ {
            fmt.Println(v7, fi15)
            _ = fi15
        }
    }
    switch 32767 {
    case -215:
        v1 += v1
    default:
        v3 = append(v3, -574)
    }
    v4[-439] = v1
    var v16 int32 = 2
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 int = (-(fi17))
        v19 := GenS2{F0: 728, F1: 586}
        if false {
            v20 := GenS2{F0: 103, F1: 921}
            v21 := GenS3{F0: (v1 + (v1 % ((v1) | 1))), F1: 2147483647}
            v16 = v16
            _ = v20
            _ = v21
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
