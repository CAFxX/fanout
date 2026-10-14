package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int; F1 int32; F2 int16 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint64; F1 int64 }
func genFn0(p0 int64) uint {
    var v5 string = ("hello" + ("\n\t" + "a"))
    v5 = v5
    v6 := GenS1{F0: 0}
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 string = "foo bar"
    var v2 int8 = 127
    v3 := []int{-794}
    var v4 uint64 = 1
    v1 = v1
    v7 := GenS2{F0: 0, F1: -200}
    var v8 int16 = -1
    if (0 <= 42) {
        _ = len(v3)
    } else {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 int8 = v2
            for fi11 := 0; fi11 < 17; fi11++ {
                if len(v3) > 0 {
                    v3[0] = int(v4)
                }
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
        var v12 uint64 = v4
        switch "foo bar" {
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = 128
            }
        case "hello":
            v13 := GenS3{F0: 703, F1: -296, F2: v8}
            _ = v13
        case "":
            v8 = (v8 >> ((18446744073709551615) & 63))
        default:
            var v14 uint64 = (v4 << ((genFn0(-1)) & 63))
            _ = v14
        }
        _ = v12
    }
    fmt.Println(v2, v1, v2)
    for fi15 := 0; fi15 < 1; fi15++ {
        if (65535 > uint16(v4)) {
            v16 := GenS4{F0: true, F1: v8}
            var v17 int32 = 44
            v17 -= 2147483647
            _ = v16
            _ = v17
        }
        if len(v3) > 0 {
            v3[0] = fi15
        }
        var v18 int16 = (v8 - v8)
        _ = fi15
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        fmt.Println(fi19)
        _ = fi19
    }
    var v20 float32 = 0.1
    if (!(1 >= 42)) {
        v4 = 18446744073709551615
        if (1 != 240) {
            for fi21 := 0; fi21 < 3; fi21++ {
                v2 += v2
                _ = fi21
            }
        }
        var v22 int = 4294967295
        _ = v22
    } else {
        v20 = ((((v20 - v20) * v20) / v20) - v20)
    }
    v23 := GenS5{F0: 893, F1: 9223372036854775807}
    fmt.Println(v20)
    v24 := GenS5{F0: (v4 & v4), F1: -1}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
