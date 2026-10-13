package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint; F1 int; F2 int32 }
type GenS4 struct { F0 string }
func genFn0(p0 int16) uint8 {
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint16 = 0
        v7 := []int{2147483647, -2147483648, (fi5 & fi5)}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: (p0 << ((42) & 63)), F1: (65535 != 18446744073709551615), F2: 1}
    var v9 string = ("hello" + "")
    v9 = (v9 + (v9 + "a"))
    _ = p0
    _ = v8
    _ = v9
    return 128
}
func genFn1(p0 int16, p1 bool) int32 {
    if (1e10 != float64((3.14159 - 0.1))) {
        v10 := GenS2{F0: 41, F1: 140, F2: 746}
        var v11 int = 4294967295
        _ = v10
        _ = v11
    }
    v12 := []int{-549}
    if (-2.25 != 1.5) {
        var v13 uint64 = 4294967295
        var v14 string = ""
        _ = v13
        _ = v14
    }
    if len(v12) > 0 {
        v12[0] = -612
    }
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    v1 := []int{32768, 308}
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 int8 = -69
    v15 := GenS3{F0: 42, F1: 804, F2: -1}
    var v16 uint64 = 770
    v16 = 18446744073709551615
    v17 := v3[-2147483648]
    var v18 uint16 = 0
    var v19 float64 = float64((-2.25 * 0.0))
    v4 -= v4
    for fi20 := 0; fi20 < 4; fi20++ {
        v16 += v16
        _ = fi20
    }
    fmt.Println(v19, v17, v4)
    var v21 uint64 = (v16 - 42)
    if ((-2.25 * ((0.1 - (0.1 - 0.1)) + (3.14159 / 3.14159))) != ((float32(v19) - (3.14159 / 1.0)) - (3.14159 * 1.5))) {
        var v22 uint32 = 927
        _ = v22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        switch "hello" {
        case "hello":
            if ("" != "\n\t") {
                v19 = ((v19 * float64(1e10)) * v19)
                var v24 uint64 = 13
                _ = v24
            } else {
                v3[(^(v17))] = (fi23 | 128)
                v1 = append(v1, v17)
                v16 += 509
            }
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            v25 := GenS4{F0: "hello"}
            _ = v25
        }
        var v26 bool = false
        _ = fi23
        _ = v26
    }
    var v27 uint = 4294967295
    v28 := GenS3{F0: 1, F1: 9223372036854775807, F2: -2147483648}
    v18 = 34
    var v29 uint64 = v21
    if ((v27 != (v27 >> ((v27) & 63))) || (false || (255 == 634))) {
        var v30 int64 = -9223372036854775808
        v31 := GenS2{F0: -21, F1: 2, F2: v30}
        _ = v30
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
