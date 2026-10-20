package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 string }
func genFn0(p0 int64) int32 {
    v5 := []int{32768, -1, 0}
    var v6 uint32 = 2147483648
    var v7 float32 = float32((float64(1.5) * ((0.0 + (1.5 - 1e10)) - 0.1)))
    var v8 bool = false
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 uint16 = 62
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int32, p1 uint8) uint {
    v11 := GenS1{F0: -9223372036854775808, F1: 92}
    var v12 uint8 = 94
    v13 := GenS2{F0: 0}
    p1 += v12
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    var v3 uint8 = 1
    var v4 int16 = 1000
    var v14 int = 128
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 float64 = (1e10 + 3.14159)
            var v18 uint16 = 27
            v19 := []int{fi16, (v14 ^ fi16), fi15, 549}
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi15
    }
    v4 -= 32767
    var v20 int = v14
    delete(v2, "\n\t")
    switch "a" {
    case "foo bar":
        fmt.Println(v20, v20)
    case "":
        var v21 uint32 = 549
        _ = v21
    case "a":
        v3 -= (v3 ^ v3)
    default:
        fmt.Println(v14, v20)
    }
    v22 := GenS2{F0: 255}
    for fi23 := 0; fi23 < 2; fi23++ {
        v24 := GenS3{F0: (true && true), F1: 475}
        var v25 bool = (-1 > -2147483648)
        v26 := GenS2{F0: (v3 << ((4294967295) & 63))}
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        if (true && true) {
            v28 := []int{}
            v29 := GenS2{F0: 200}
            _ = v28
            _ = v29
        } else {
            for fi30 := 0; fi30 < 7; fi30++ {
                _ = len(v2)
                _ = fi30
            }
            fmt.Println(v14)
            v20 = fi27
        }
        var v31 uint64 = 18446744073709551615
        var v32 int = (fi27 / (((fi27 * -2147483648)) | 1))
        _ = fi27
        _ = v31
        _ = v32
    }
    v33 := GenS4{F0: ("\n\t" + "foo bar")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v33
    _ = v4
}
