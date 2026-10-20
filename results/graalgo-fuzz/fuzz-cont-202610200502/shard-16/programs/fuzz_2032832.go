package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0(p0 uint, p1 uint16, p2 uint) uint8 {
    p1 += p1
    var v6 int64 = 1
    if false {
        v7 := GenS1{F0: 858}
        _ = v7
    }
    p2 = (p0 ^ 4294967295)
    v6 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 115
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 575
    var v4 uint32 = (^(v3))
    var v5 string = "hello"
    v8 := []int{int(v3), -2147483648, 235, 845}
    v5 = v5
    if true {
        v9 := GenS2{F0: 1, F1: 65535}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 int8 = 0
            v4 -= (^(v3))
            _ = fi10
            _ = v11
        }
        fmt.Println(v4, v4, v5)
    }
    if true {
        switch 1 {
        case -598:
            var v12 int8 = -128
            _ = v12
        case 2147483647:
            var v13 uint64 = 2
            _ = v13
        case 32768:
            for fi14 := 0; fi14 < 7; fi14++ {
                _ = len(v2)
                delete(v2, 922)
                var v15 int32 = -2147483648
                _ = fi14
                _ = v15
            }
        default:
            if true {
                v4 = 0
            }
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    switch "a" {
    case "foo bar":
        v3 += (v4 / ((v3) | 1))
    case "":
        var v16 uint64 = 2
        _ = v16
    default:
        var v17 int32 = 42
        _ = v17
    }
    var v18 int16 = -789
    _ = len(v8)
    delete(v1, v4)
    var v19 int32 = 0
    var v20 uint64 = 18446744073709551615
    var v21 uint32 = 1
    var v22 float32 = 0.0
    if (685 >= -2147483648) {
        fmt.Println(v22, v4, v4)
        var v23 int = 42
        _ = v23
    } else {
        for fi24 := 0; fi24 < 14; fi24++ {
            _ = len(v1)
            v25 := GenS2{F0: v20, F1: 0}
            _ = fi24
            _ = v25
        }
        v19 = (v19 + -2147483648)
        v3 += (v3 & (v3 % (((v3 ^ v4)) | 1)))
    }
    var v26 float32 = 0.1
    var v27 uint64 = v20
    v28 := GenS2{F0: (^(v27)), F1: uint(v20)}
    v8 = append(v8, 685)
    v20 -= (v20 >> ((672) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
