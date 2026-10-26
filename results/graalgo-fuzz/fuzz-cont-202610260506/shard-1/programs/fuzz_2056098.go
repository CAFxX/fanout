package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }
type GenS2 struct { F0 int64 }
func genFn0(p0 string, p1 int64, p2 uint16) int32 {
    fmt.Println(p2, p0, p0)
    if (!(0 == 42)) {
        if (32768 > 501) {
            var v6 uint64 = 4294967295
            if false {
                var v7 int32 = int32(p1)
                _ = v7
            }
            _ = v6
        }
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        if ((p1 ^ (p1 >> ((65535) & 63))) < 9223372036854775807) {
            p1 = p1
            p2 -= 292
        } else {
            switch p0 {
            case "\n\t":
                var v9 uint8 = 171
                _ = v9
            default:
                v10 := []int{fi8, 32768, (^(fi8)), fi8}
                _ = v10
            }
            switch fi8 {
            case 32767:
                var v11 int64 = (p1 - p1)
                _ = v11
            case 32768:
                fmt.Println(fi8, p1)
            case -2147483648:
                var v12 int32 = -1
                _ = v12
            default:
                p1 += (p1 + p1)
            }
            var v13 uint32 = 1
            _ = v13
        }
        var v14 bool = false
        _ = fi8
        _ = v14
    }
    if (4294967295 != 128) {
        if (100 <= 42) {
            if ((-718 >= 0) && false) {
                var v15 uint8 = 42
                var v16 float64 = 3.14159
                _ = v15
                _ = v16
            }
            p0 = p0
            v17 := GenS1{F0: (p2 % (((^(p2))) | 1)), F1: p0, F2: (p1 >> ((0) & 63))}
            _ = v17
        }
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int16) uint32 {
    fmt.Println(p0)
    fmt.Println(p0)
    if ((p0 < p0) || (!false)) {
        fmt.Println(p0)
        if true {
            var v18 uint32 = 828
            _ = v18
        }
        v19 := []int{42}
        _ = v19
    }
    var v20 float64 = 3.14159
    var v21 uint8 = 93
    _ = p0
    _ = v20
    _ = v21
    return 0
}

func main() {
    var v1 float64 = -2.25
    var v2 int = 128
    var v3 int64 = 67
    var v4 float32 = (3.14159 / float32(v1))
    var v5 float32 = float32(1e10)
    v4 = 0.1
    v22 := GenS1{F0: 22, F1: "a", F2: (v3 + (v3 % ((v3) | 1)))}
    fmt.Println(v5, v3, v4)
    v23 := []int{}
    v5 -= v5
    v3 += (v3 >> ((18446744073709551615) & 63))
    var v24 int8 = -1
    v24 = v24
    _ = len(v23)
    var v25 bool = (v2 >= (v2 >> ((274) & 63)))
    var v26 uint8 = 255
    fmt.Println(v3, v2, v3)
    switch "hello" {
    case "foo bar":
        var v27 uint64 = 61
        _ = v27
    case "hello":
        var v28 uint64 = 18446744073709551615
        _ = v28
    case "a":
        v29 := GenS1{F0: 255, F1: "\n\t", F2: v3}
        _ = v29
    default:
        v23 = append(v23, v2)
    }
    v30 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
