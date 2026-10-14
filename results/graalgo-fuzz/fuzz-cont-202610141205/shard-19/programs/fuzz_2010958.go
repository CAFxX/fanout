package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int8; F2 int }
type GenS4 struct { F0 string }
func genFn0(p0 uint32, p1 int16, p2 int8) int32 {
    if (p2 > 0) {
        v6 := GenS1{F0: 0, F1: (p0 | p0), F2: 547}
        if true {
            v7 := []int{1, 890}
            _ = v7
        } else {
            switch "hello" {
            case "":
                var v8 int64 = 9223372036854775807
                _ = v8
            }
            var v9 int = -679
            _ = v9
        }
        _ = v6
    } else {
        v10 := GenS2{F0: true}
        _ = v10
    }
    p1 = (p1 & p1)
    v11 := GenS3{F0: p2, F1: (p2 ^ p2), F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -2147483648
}

func main() {
    var v1 int8 = -1
    v2 := []uint32{}
    v3 := make(map[string]bool)
    var v4 int8 = v1
    v5 := make(map[string]uint32)
    v12 := GenS4{F0: "hello"}
    v4 = (v1 - (v1 >> ((41) & 63)))
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 uint = 2
        _ = fi13
        _ = v14
    }
    v4 += (^(v4))
    for fi15 := 0; fi15 < 5; fi15++ {
        v4 += v1
        _ = fi15
    }
    var v16 int = -697
    fmt.Println(v1, v16, v16)
    fmt.Println(v4, v4)
    for fi17 := 0; fi17 < 20; fi17++ {
        v4 = (v1 - v4)
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
