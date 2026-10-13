package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int64; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0(p0 bool) uint32 {
    p0 = p0
    if p0 {
        for fi7 := 0; fi7 < 10; fi7++ {
            fmt.Println(fi7, p0)
            p0 = (true || p0)
            _ = fi7
        }
        if (!(false || p0)) {
            for fi8 := 0; fi8 < 8; fi8++ {
                p0 = true
                _ = fi8
            }
            var v9 uint16 = 857
            var v10 int = 4294967295
            _ = v10
            _ = v9
        } else {
            switch 32767 {
            case -9223372036854775808:
                if p0 {
                    p0 = p0
                }
            case 4294967295:
                var v11 int32 = 1
                _ = v11
            default:
                var v12 int64 = 42
                _ = v12
            }
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := []int{(fi13 + (fi13 << ((396) & 63)))}
                v15 := GenS1{F0: fi13, F1: 699}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        p0 = p0
    }
    p0 = true
    v16 := GenS1{F0: 0, F1: 281}
    fmt.Println(p0)
    _ = p0
    _ = v16
    return 495
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    var v3 int32 = 1
    v4 := make(map[uint32]string)
    var v5 uint8 = 136
    var v6 int32 = (v3 & v3)
    v6 = v6
    v17 := GenS2{F0: -9223372036854775808, F1: 42, F2: 42}
    var v18 uint = 42
    v19 := GenS1{F0: int(v18), F1: v18}
    var v20 float64 = (1.5 / float64(float32(0.1)))
    switch 552 {
    case 42:
        var v21 int64 = -833
        _ = v21
    case 2147483647:
        switch "\n\t" {
        case "a":
            var v22 string = "hello"
            _ = v22
        default:
            var v23 int = -514
            _ = v23
        }
    case -9223372036854775808:
        var v24 string = ("" + "")
        _ = v24
    default:
        var v25 float32 = 1e10
        _ = v25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        v20 = float64(0.1)
        v5 += (v5 % ((v5) | 1))
        _ = fi26
    }
    if true {
        v1["hello"] = true
        for fi27 := 0; fi27 < 12; fi27++ {
            v28 := GenS2{F0: 147, F1: (v5 ^ (v5 & 1)), F2: 0}
            fmt.Println(v18)
            fmt.Println(v18, v20)
            _ = fi27
            _ = v28
        }
    }
    var v29 uint16 = uint16(v5)
    fmt.Println(v29, v29, v29)
    v1["foo bar"] = ((v29 << ((v18) & 63)) > (^(v29)))
    v30 := GenS3{F0: v29, F1: (^(v5))}
    var v31 int32 = (v6 / ((v3) | 1))
    var v32 int64 = -9223372036854775808
    var v33 int32 = v31
    fmt.Println(v32)
    v34 := []int{284, -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
