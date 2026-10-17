package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{751}
    var v3 int32 = 1
    var v4 uint64 = 0
    var v5 uint = 0
    v6 := make(map[string]uint32)
    var v7 bool = true
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: v4, F1: (v7 || (127 < fi8))}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        switch ("" + "") {
        case "":
            v11 := v6["a"]
            _ = v11
        default:
            switch fi10 {
            case -9223372036854775808:
                v12 := GenS2{F0: 1}
                _ = v12
            case -2147483648:
                v7 = v7
            }
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := GenS3{F0: 534, F1: (v5 | v5)}
            v7 = v7
            _ = fi13
            _ = v14
        }
        v7 = (v3 <= (v3 / (((v3 - (v3 << ((262) & 63)))) | 1)))
        _ = fi10
    }
    fmt.Println(v7, v7, v3)
    if ((2147483648 >= 1) || v7) {
        var v15 int16 = 0
        if (((3.14159 / 1.0) + -2.25) >= 0.1) {
            if false {
                var v16 bool = v7
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 14; fi17++ {
                v5 = 1
                v18 := v6[(("foo bar" + "") + "hello")]
                _ = fi17
                _ = v18
            }
            if true {
                v15 -= 378
                _ = len(v2)
            } else {
                v19 := GenS4{F0: 2, F1: 42}
                v7 = false
                _ = v19
            }
        }
        _ = v15
    } else {
        var v20 float32 = 1e10
        fmt.Println(v20, v4, v3)
        _ = v20
    }
    v21 := GenS4{F0: 499, F1: uint16(v5)}
    if v7 {
        v22 := v6["hello"]
        var v23 int32 = 1
        _ = v22
        _ = v23
    }
    if true {
        var v24 int64 = -54
        v2 = append(v2, -917)
        v2 = append(v2, 9223372036854775807)
        _ = v24
    }
    v7 = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
