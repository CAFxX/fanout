package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 bool }
func genFn0(p0 uint16, p1 int32, p2 int16) uint64 {
    var v7 uint32 = 114
    var v8 int8 = 6
    v9 := GenS1{F0: int64(v7)}
    v10 := GenS2{F0: p0}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    v3 := []uint32{42}
    var v4 uint32 = 42
    var v5 int64 = 687
    var v6 uint8 = 127
    v4 = v4
    v11 := GenS2{F0: 2}
    if false {
        var v12 float64 = (1e10 * 1e10)
        _ = v12
    }
    var v13 int16 = 1000
    var v14 uint32 = 101
    fmt.Println(v5, v5, v5)
    v13 += (v13 ^ 0)
    v13 = (v13 + v13)
    for fi15 := 0; fi15 < 9; fi15++ {
        fmt.Println(v6, v5)
        _ = fi15
    }
    switch 968 {
    case 128:
        v14 += (v14 - 331)
    case 42:
        var v16 int32 = 0
        _ = v16
    case -880:
        _ = len(v1)
    default:
        v17 := GenS1{F0: v5}
        _ = v17
    }
    v4 -= (v4 ^ 13)
    for fi18 := 0; fi18 < 19; fi18++ {
        v14 += (v4 & v4)
        v6 += (v6 ^ (v6 & (v6 / ((78) | 1))))
        _ = fi18
    }
    v6 = v6
    v1[2147483647] = (v4 - v14)
    for fi19 := 0; fi19 < 19; fi19++ {
        switch "foo bar" {
        case "\n\t":
            var v20 uint32 = v14
            _ = v20
        default:
            v21 := GenS3{F0: 42, F1: false}
            _ = v21
        }
        fmt.Println(v6, fi19, fi19)
        _ = fi19
    }
    v6 -= v6
    switch -1 {
    case 9223372036854775807:
        var v22 float64 = -2.25
        _ = v22
    default:
        switch 2147483647 {
        case 291:
            var v23 uint32 = 710
            _ = v23
        case 9223372036854775807:
            if (2 < 148) {
                v6 += 128
            } else {
                var v24 int8 = -113
                v25 := GenS2{F0: 42}
                var v26 string = "hello"
                _ = v24
                _ = v25
                _ = v26
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
