package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint64 }
func genFn0() uint8 {
    v5 := GenS1{F0: 127, F1: 2, F2: 2}
    var v6 string = "hello"
    if true {
        fmt.Println(v6)
        var v7 int = -9223372036854775808
        var v8 float64 = 1e10
        _ = v7
        _ = v8
    }
    _ = v5
    _ = v6
    return 128
}
func genFn1() uint32 {
    var v9 int32 = -600
    if (0.1 != ((float64(0.0) + ((1.5 * 1e10) - (1e10 / 0.1))) / 3.14159)) {
        var v10 int = -2147483648
        if (42 > 37) {
            var v11 float64 = ((0.0 + (3.14159 - -2.25)) - ((-2.25 + -2.25) / (1.5 / 0.1)))
            var v12 int64 = 895
            _ = v11
            _ = v12
        }
        if true {
            v13 := []int{v10, 42}
            var v14 int8 = -39
            fmt.Println(v14, v9)
            _ = v13
            _ = v14
        } else {
            var v15 uint16 = 65535
            v10 -= (v10 / ((v10) | 1))
            _ = v15
        }
        _ = v10
    }
    if (1e10 <= -2.25) {
        var v16 int16 = 655
        var v17 int8 = 127
        v18 := GenS2{F0: 1}
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 int = 128
    v20 := []int{(v19 * v19), (v19 & (-(v19))), (v19 >> ((0) & 63)), (v19 | 4294967295)}
    _ = v19
    _ = v20
    _ = v9
    return 0
}

func main() {
    var v1 uint32 = 217
    var v2 float64 = (1e10 / float64(((((-2.25 - 3.14159) - (0.0 - 1e10)) - 3.14159) / 1.0)))
    var v3 int = 9223372036854775807
    var v4 int64 = -9223372036854775808
    var v21 bool = false
    var v22 int64 = v4
    var v23 int64 = (v4 % ((v4) | 1))
    v24 := []int{-65, v3, 684, (v3 & v3)}
    _ = len(v24)
    var v25 int8 = 100
    for fi26 := 0; fi26 < 6; fi26++ {
        switch ("a" + "") {
        case "\n\t":
            var v27 int32 = 0
            _ = v27
        case "":
            var v28 float32 = 1.5
            _ = v28
        case "hello":
            _ = len(v24)
        }
        v22 += (v23 - v23)
        _ = fi26
    }
    switch v3 {
    case 32768:
        for fi29 := 0; fi29 < 19; fi29++ {
            for fi30 := 0; fi30 < 7; fi30++ {
                var v31 int64 = v22
                var v32 int64 = int64(v25)
                v33 := GenS1{F0: -21, F1: v1, F2: v1}
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
            _ = fi29
        }
    case 32767:
        v24 = append(v24, -1)
    }
    switch "a" {
    case "\n\t":
        v21 = v21
    case "":
        v3 += 159
    default:
        var v34 float64 = v2
        _ = v34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 string = ("hello" + "hello")
        v3 = (fi35 + fi35)
        _ = fi35
        _ = v36
    }
    v24 = append(v24, (v3 & v3))
    v3 -= v3
    v22 = v23
    v37 := GenS3{F0: -1, F1: 226}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v4
}
