package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 string; F1 int32 }
func genFn0() int {
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    var v8 uint64 = 611
    fmt.Println(v8)
    if false {
        v8 = (v8 ^ (^(v8)))
    } else {
        v8 = 18446744073709551615
        v9 := GenS1{F0: 25}
        _ = v9
    }
    _ = v8
    return 127
}
func genFn1(p0 string, p1 int, p2 int64) uint32 {
    p2 -= -9223372036854775808
    v10 := []int{p1, p1, (p1 / (((p1 & 2147483647)) | 1)), -300}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 2147483648
}

func main() {
    v1 := []int{-56, 916, -320}
    var v2 int8 = -21
    v3 := make(map[string]int)
    var v4 int64 = 42
    v5 := make(map[string]uint32)
    v6 := make(map[string]int)
    v2 += (v2 - (v2 | 22))
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    if (4294967295 != 738) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v11 := GenS1{F0: 0}
        var v12 uint = 42
        _ = v11
        _ = v12
    }
    if (856 <= 4294967295) {
        v13 := GenS2{F0: 1, F1: v2}
        var v14 float32 = 1e10
        var v15 string = "foo bar"
        _ = v13
        _ = v14
        _ = v15
    } else {
        v4 += 401
    }
    v16 := GenS2{F0: 42, F1: (v2 ^ (v2 + v2))}
    var v17 uint32 = 0
    if (!false) {
        v18 := GenS3{F0: (("a" + "\n\t") + "\n\t"), F1: 852}
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        if (!(255 >= 255)) {
            switch "" {
            case "foo bar":
                v20 := GenS2{F0: 65535, F1: (v2 / ((int8(v17)) | 1))}
                _ = v20
            case "\n\t":
                var v21 int16 = -32768
                _ = v21
            default:
                v22 := GenS2{F0: 65535, F1: 47}
                _ = v22
            }
        } else {
            switch "a" {
            case "":
                var v23 string = ("a" + "hello")
                _ = v23
            case "foo bar":
                v2 -= (v2 % ((-1) | 1))
            }
            v17 = (v17 - (v17 >> ((2) & 63)))
            _ = len(v5)
        }
        v24 := v3["a"]
        var v25 bool = false
        _ = fi19
        _ = v24
        _ = v25
    }
    var v26 uint32 = 2147483647
    for fi27 := 0; fi27 < 6; fi27++ {
        if len(v1) > 0 {
            v1[0] = fi27
        }
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
