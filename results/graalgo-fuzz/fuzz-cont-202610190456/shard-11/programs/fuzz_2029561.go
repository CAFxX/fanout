package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint8; F2 int16 }
func genFn0(p0 bool, p1 uint, p2 string) int {
    v4 := []int{1}
    if (528 > 618) {
        p2 = (p2 + "foo bar")
        if (p2 == (("" + "a") + "\n\t")) {
            p1 -= p1
        } else {
            if (32767 > 42) {
                fmt.Println(p2, p2)
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -249
}
func genFn1(p0 int16, p1 int8, p2 int8) uint64 {
    p0 = (p0 + 451)
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            if (-888 < 56) {
                v7 := []int{}
                fmt.Println(p0, p1, p1)
                _ = v7
            }
            var v8 int8 = p1
            p0 = -1
            _ = fi6
            _ = v8
        }
        _ = fi5
    }
    v9 := GenS1{F0: -394, F1: 0, F2: 922}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []int{0, -654, 0}
    v3 := make(map[int]string)
    var v10 float32 = -2.25
    if true {
        switch "" {
        case "a":
            v11 := GenS1{F0: 9223372036854775807, F1: (^(v1)), F2: 1}
            _ = v11
        case "foo bar":
            v1 += 0
        default:
            for fi12 := 0; fi12 < 18; fi12++ {
                _ = len(v2)
                v13 := GenS2{F0: -1}
                _ = fi12
                _ = v13
            }
        }
    }
    var v14 uint32 = 2
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 uint32 = (v14 | (v14 % (((v14 << ((629) & 63))) | 1)))
        v16 = (v16 * (^(v14)))
        _ = fi15
        _ = v16
    }
    if (v1 <= v1) {
        v17 := GenS3{F0: -1, F1: uint8(v1), F2: 42}
        if false {
            if len(v2) > 0 {
                v2[0] = -674
            }
            _ = len(v3)
            var v18 uint16 = 214
            _ = v18
        }
        if (false || (0.0 < 0.0)) {
            v19 := GenS1{F0: 1, F1: (v1 / ((genFn1(-579, -1, 42)) | 1)), F2: -1}
            _ = v19
        } else {
            var v20 int64 = 9223372036854775807
            v2 = append(v2, -1)
            v21 := GenS1{F0: (v20 << ((4294967295) & 63)), F1: v1, F2: (v20 >> ((872) & 63))}
            _ = v20
            _ = v21
        }
        _ = v17
    }
    var v22 uint8 = 128
    var v23 int64 = 9223372036854775807
    var v24 bool = (1 <= -36)
    _ = len(v2)
    v14 = 899
    if ((v1 % ((genFn1(42, -1, -1)) | 1)) <= 4294967295) {
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 bool = v24
            _ = fi25
            _ = v26
        }
        _ = len(v3)
    }
    var v27 float32 = (float32(1e10) - float32((-2.25 - float64(3.14159))))
    for fi28 := 0; fi28 < 3; fi28++ {
        if (((((!(v1 < 0)) || (v24 || (!(((v22 % (((^(v22))) | 1)) < 2) && v24)))) && (v23 >= -268)) || false) && false) {
            var v29 uint = 552
            var v30 int16 = -443
            _ = v29
            _ = v30
        } else {
            if (v24 && v24) {
                v23 -= v23
            } else {
                fmt.Println(v22, v24)
            }
            fmt.Println(v22, v1, v1)
        }
        _ = fi28
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
}
