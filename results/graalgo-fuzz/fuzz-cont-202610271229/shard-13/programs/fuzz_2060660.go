package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint8) uint32 {
    fmt.Println(p0)
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return 771
}
func genFn1() uint32 {
    var v5 int16 = 512
    fmt.Println(v5)
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: 701, F1: ("\n\t" == "\n\t")}
        _ = fi6
        _ = v7
    }
    fmt.Println(v5)
    v8 := GenS1{F0: 0, F1: false}
    _ = v5
    _ = v8
    return 299
}

func main() {
    v1 := []int{9223372036854775807, 32767}
    var v2 int = -652
    var v3 float32 = float32(1.5)
    var v4 uint8 = 255
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 int32 = 2147483647
        _ = fi9
        _ = v10
    }
    var v11 uint64 = 2
    fmt.Println(v2, v2, v11)
    v4 += v4
    var v12 float64 = -2.25
    v13 := GenS2{F0: (^(v4)), F1: 1, F2: genFn1()}
    fmt.Println(v3, v11)
    v12 -= 1.5
    var v14 int16 = 32767
    if (((!false) || (-9223372036854775808 == 677)) && true) {
        if false {
            _ = len(v1)
            if len(v1) > 0 {
                v1[0] = v2
            }
            var v15 int32 = 1
            _ = v15
        }
        var v16 bool = (false || ((-(v2)) >= v2))
        v17 := GenS3{F0: 42}
        _ = v16
        _ = v17
    }
    v11 = (v11 >> ((2) & 63))
    v2 = (v2 / ((v2) | 1))
    _ = len(v1)
    if false {
        if (!((-1 != int64(v2)) || false)) {
            if len(v1) > 0 {
                v1[0] = (^(v2))
            }
            v2 -= -238
            var v18 float32 = v3
            _ = v18
        } else {
            if (!(!true)) {
                var v19 int16 = 32767
                switch "\n\t" {
                case "":
                    v20 := GenS4{F0: (v11 << ((0) & 63))}
                    _ = v20
                }
                _ = v19
            }
        }
        var v21 int32 = 2147483647
        var v22 uint64 = (v11 << ((42) & 63))
        _ = v21
        _ = v22
    } else {
        fmt.Println(v11, v12)
        v12 += (float64(v3) + float64(v3))
        v12 -= 1.5
    }
    if (!false) {
        for fi23 := 0; fi23 < 18; fi23++ {
            for fi24 := 0; fi24 < 13; fi24++ {
                v4 += v4
                _ = fi24
            }
            _ = fi23
        }
        switch "a" {
        case "\n\t":
            var v25 int64 = -1
            _ = v25
        default:
            v11 = 1
        }
        var v26 int32 = 135
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
