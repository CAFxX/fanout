package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int8; F1 int8; F2 int32 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint }

func main() {
    var v1 uint32 = 42
    v2 := []uint32{v1}
    v3 := make(map[int]uint32)
    for fi4 := 0; fi4 < 19; fi4++ {
        if true {
            v1 = 245
            _ = len(v2)
            switch "foo bar" {
            case "\n\t":
                delete(v3, (fi4 >> ((250) & 63)))
            default:
                for fi5 := 0; fi5 < 19; fi5++ {
                    var v6 int32 = -1
                    v6 -= -1
                    var v7 int32 = (-(v6))
                    _ = fi5
                    _ = v6
                    _ = v7
                }
            }
        }
        fmt.Println(fi4)
        _ = fi4
    }
    var v8 uint8 = 255
    v9 := GenS1{F0: -420, F1: 4294967295}
    v10 := GenS1{F0: -67, F1: 42}
    var v11 string = "hello"
    if (("hello" != "foo bar") || ((^(v1)) >= (v1 ^ 370))) {
        for fi12 := 0; fi12 < 6; fi12++ {
            if ((!(42 <= 42)) || true) {
                v13 := GenS2{F0: -1, F1: 127, F2: 42}
                v2 = append(v2, v1)
                _ = v13
            }
            fmt.Println(v1)
            v14 := GenS3{F0: 2, F1: true, F2: 527}
            _ = fi12
            _ = v14
        }
    } else {
        var v15 bool = (false || (1 <= 1))
        _ = v15
    }
    if len(v2) > 0 {
        v2[0] = v1
    }
    v16 := GenS3{F0: v1, F1: true, F2: 1}
    switch -1 {
    case 2147483647:
        v8 += 2
    case 0:
        v8 -= (v8 << ((18446744073709551615) & 63))
    }
    _ = len(v2)
    var v17 uint8 = (v8 - v8)
    v18 := v3[4294967295]
    delete(v3, -996)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
