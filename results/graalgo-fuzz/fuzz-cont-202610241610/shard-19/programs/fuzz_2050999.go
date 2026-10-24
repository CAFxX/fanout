package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int; F1 uint16; F2 int16 }
func genFn0() int64 {
    var v4 int8 = -128
    var v5 float32 = 1.5
    v6 := GenS1{F0: (-9223372036854775808 < 42), F1: -1, F2: -916}
    _ = v4
    _ = v5
    _ = v6
    return -9223372036854775808
}

func main() {
    v1 := []int{696, 2147483647, -2147483648, -1}
    var v2 int8 = -1
    v3 := make(map[int]string)
    delete(v3, 42)
    var v7 int16 = -979
    fmt.Println(v2, v2)
    for fi8 := 0; fi8 < 15; fi8++ {
        switch -9223372036854775808 {
        case -9223372036854775808:
            fmt.Println(v2)
        }
        _ = fi8
    }
    v7 = v7
    for fi9 := 0; fi9 < 7; fi9++ {
        if ("\n\t" <= "") {
            if (!true) {
                if false {
                    v2 = 0
                    var v10 float64 = 3.14159
                    var v11 float32 = 1e10
                    _ = v10
                    _ = v11
                } else {
                    var v12 uint8 = 42
                    v7 = v7
                    fmt.Println(v2, v12, v12)
                    _ = v12
                }
                _ = len(v1)
                var v13 uint16 = 675
                _ = v13
            } else {
                v2 += v2
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi14 := 0; fi14 < 18; fi14++ {
                    v7 = (^(v7))
                    var v15 int64 = 1
                    v16 := GenS2{F0: 219, F1: 943}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            var v17 int32 = 42
            _ = v17
        }
        v18 := GenS3{F0: fi9, F1: 0, F2: -21}
        _ = fi9
        _ = v18
    }
    var v19 uint16 = 786
    fmt.Println(v19)
    var v20 int8 = -63
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := v3[(fi21 | (fi21 * fi21))]
        var v23 int32 = 1
        fmt.Println(fi21, v22, v23)
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 int64 = 9223372036854775807
    var v25 int16 = (v7 | -1)
    v2 = 0
    switch 4294967295 {
    case 127:
        var v26 float32 = 1.5
        _ = v26
    case -883:
        v27 := GenS2{F0: 18446744073709551615, F1: 2147483647}
        _ = v27
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v7
}
