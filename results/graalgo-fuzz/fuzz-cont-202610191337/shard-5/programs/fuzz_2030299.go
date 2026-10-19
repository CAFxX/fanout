package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int32 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 bool; F1 uint32; F2 uint }
func genFn0(p0 int16) string {
    var v6 float64 = 1e10
    switch ("foo bar" + "") {
    case "\n\t":
        fmt.Println(v6)
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := []int{fi8, fi8}
            _ = len(v9)
            _ = fi8
            _ = v9
        }
        v6 -= 1e10
        fmt.Println(v6, p0, v6)
        _ = fi7
    }
    _ = p0
    _ = v6
    return ""
}

func main() {
    var v1 int32 = -644
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    var v4 string = "a"
    v5 := []int{-747, -1}
    v10 := v2[((v4 + (v4 + ("\n\t" + v4))) + "")]
    v4 = ""
    v10 -= v10
    v1 = (v1 >> ((42) & 63))
    v10 = (v10 / ((v10) | 1))
    var v11 uint16 = 42
    v10 += (v10 % ((v10) | 1))
    _ = len(v2)
    v12 := GenS1{F0: (v1 >> ((4294967295) & 63)), F1: -1, F2: (v1 >> ((42) & 63))}
    v13 := GenS2{F0: v11, F1: v1, F2: 233}
    switch v10 {
    case 9223372036854775807:
        var v14 uint32 = 2147483647
        _ = v14
    default:
        v15 := GenS3{F0: 1000, F1: 55}
        _ = v15
    }
    if (!false) {
        v16 := GenS1{F0: 431, F1: -276, F2: v1}
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 int = fi17
            v5 = append(v5, (fi17 << ((641) & 63)))
            var v20 uint = 237
            _ = fi18
            _ = v19
            _ = v20
        }
        v11 = (v11 ^ (v11 / (((v11 * v11)) | 1)))
        if false {
            v21 := GenS2{F0: v11, F1: 0, F2: 42}
            v22 := GenS4{F0: (!false), F1: 44, F2: 18446744073709551615}
            _ = v21
            _ = v22
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
