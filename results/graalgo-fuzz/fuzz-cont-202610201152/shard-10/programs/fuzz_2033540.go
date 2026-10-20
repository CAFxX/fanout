package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 int16; F2 int }
func genFn0(p0 uint64, p1 uint64) int {
    v6 := GenS1{F0: 2, F1: p0, F2: 602}
    v7 := GenS2{F0: "a"}
    v8 := GenS3{F0: 1, F1: -32768, F2: -758}
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := []int{-9223372036854775808, (^(fi9)), (fi10 + -505)}
            p0 = p1
            var v12 float64 = (float64(3.14159) - float64(((1e10 - (-2.25 * -2.25)) - 0.1)))
            _ = fi10
            _ = v11
            _ = v12
        }
        p1 -= 18446744073709551615
        for fi13 := 0; fi13 < 19; fi13++ {
            if ((!((float32(1e10) * float32(1e10)) == -2.25)) && (float32(1.5) >= 0.0)) {
                v14 := GenS3{F0: -991, F1: 840, F2: fi9}
                _ = v14
            }
            _ = fi13
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 128
}

func main() {
    var v1 int = -9223372036854775808
    v2 := []int{(v1 / ((v1) | 1)), v1, (v1 / (((v1 % (((v1 & v1)) | 1))) | 1))}
    v3 := []int{v1, (v1 - v1), v1}
    var v4 uint64 = 0
    var v5 uint16 = 0
    var v15 uint64 = uint64(v5)
    var v16 uint32 = uint32(v15)
    fmt.Println(v16, v5, v4)
    v17 := GenS3{F0: 2147483647, F1: 781, F2: (v1 ^ (v1 / (((v1 * genFn0(v15, v4))) | 1)))}
    if true {
        var v18 int = genFn0((^(v15)), (v4 * (v4 + (^(v4)))))
        _ = v18
    } else {
        fmt.Println(v1)
        var v19 int64 = -24
        _ = v19
    }
    v5 -= 255
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v20 int8 = 100
    if len(v2) > 0 {
        v2[0] = (^(v1))
    }
    switch v1 {
    case -2147483648:
        v3 = append(v3, 9223372036854775807)
    case -1:
        var v21 int32 = 1
        _ = v21
    }
    var v22 uint64 = (v4 % (((^(v15))) | 1))
    var v23 uint = 0
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 float32 = 1e10
        var v26 string = "\n\t"
        v5 = (v5 << ((v23) & 63))
        _ = fi24
        _ = v25
        _ = v26
    }
    if ((false && false) || false) {
        var v27 float64 = 1.5
        _ = v27
    }
    var v28 int8 = v20
    v5 = (v5 & uint16(v16))
    var v29 int64 = 108
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
