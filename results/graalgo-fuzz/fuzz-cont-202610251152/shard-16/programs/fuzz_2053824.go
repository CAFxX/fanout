package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int; F1 int64 }
func genFn0(p0 int8, p1 string) bool {
    var v6 uint8 = 2
    var v7 uint8 = (v6 % ((v6) | 1))
    if (0 < 258) {
        v8 := []int{-2147483648, -9223372036854775808, 42, 127}
        _ = v8
    }
    var v9 float64 = ((-2.25 + (float64(float32(1.5)) / 1.0)) + (1.5 + (float64(3.14159) / (0.1 * (1.5 + -2.25)))))
    if (false || (0 < 32767)) {
        v10 := []int{32767, 32768, 32767, 9223372036854775807}
        var v11 int16 = 1000
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v9
    return false
}
func genFn1(p0 uint16) uint64 {
    p0 = p0
    var v12 int32 = -2147483648
    v13 := GenS1{F0: 65}
    fmt.Println(v12, v12)
    _ = p0
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := []uint32{1}
    var v2 string = ""
    var v3 uint32 = 430
    var v4 uint = 18446744073709551615
    v5 := make(map[int]uint32)
    v14 := GenS1{F0: uint8(v4)}
    switch "hello" {
    case "hello":
        var v15 int64 = -1
        _ = v15
    case "a":
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 uint16 = 893
            _ = fi16
            _ = v17
        }
    default:
        var v18 int32 = 0
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 bool = (true || true)
        var v21 uint32 = (v3 - (^(v3)))
        delete(v5, (fi19 - (fi19 % ((2147483647) | 1))))
        _ = fi19
        _ = v20
        _ = v21
    }
    v2 = "\n\t"
    fmt.Println(v4, v4)
    var v22 uint32 = v3
    var v23 int32 = 2147483647
    v3 = 416
    v24 := GenS2{F0: (v2 + v2), F1: (v2 + v2)}
    var v25 int = 1
    v3 += (v3 ^ (v3 | (v22 % (((v3 << (((^(v4))) & 63))) | 1))))
    v26 := GenS3{F0: v25, F1: 458}
    if len(v1) > 0 {
        v1[0] = v3
    }
    var v27 string = "a"
    v27 = (("foo bar" + v2) + "foo bar")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
