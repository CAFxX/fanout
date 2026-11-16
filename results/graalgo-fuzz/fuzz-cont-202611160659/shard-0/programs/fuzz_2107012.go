package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int; F2 int16 }
func genFn0(p0 int16, p1 uint) int {
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 int32 = 0
        fmt.Println(p1)
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: p1}
    v7 := []int{128, 4294967295}
    for fi8 := 0; fi8 < 10; fi8++ {
        switch fi8 {
        case 175:
            v9 := GenS1{F0: p1}
            _ = v9
        case -760:
            v7 = append(v7, fi8)
        case -2147483648:
            fmt.Println(fi8)
        }
        var v10 int16 = (^(p0))
        if false {
            v11 := GenS2{F0: -128}
            if false {
                var v12 int16 = 32767
                var v13 int16 = v10
                _ = v12
                _ = v13
            }
            v10 += (p0 * (p0 % (((p0 + v10)) | 1)))
            _ = v11
        } else {
            p0 = 13
            if true {
                switch fi8 {
                case 459:
                    v10 = v10
                default:
                    var v14 float64 = 3.14159
                    _ = v14
                }
            }
            var v15 int = (fi8 & fi8)
            _ = v15
        }
        _ = fi8
        _ = v10
    }
    var v16 uint32 = 0
    _ = p0
    _ = p1
    _ = v16
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 int8) int {
    p0 = p0
    p0 += p0
    _ = p0
    return 32767
}

func main() {
    v1 := []string{"", "hello", "\n\t", "\n\t"}
    var v2 uint8 = 33
    var v3 int = 127
    v2 += v2
    v17 := GenS2{F0: 127}
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 uint32 = 423
        v20 := GenS3{F0: 1, F1: fi18, F2: 56}
        var v21 int = (fi18 * v3)
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    v2 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (((float32((1.5 - (1.5 + 0.0))) + 0.1) + 0.0) != 0.0) {
        v3 += (^(v3))
        fmt.Println(v3, v2)
    }
    v3 -= v3
    switch "a" {
    case "foo bar":
        if (true && (int32(v3) != 2147483647)) {
            _ = len(v1)
            if (true && true) {
                v3 += (v3 << ((255) & 63))
            }
            var v22 uint32 = 453
            _ = v22
        } else {
            var v23 int16 = -575
            var v24 string = "a"
            var v25 int = 4294967295
            _ = v23
            _ = v24
            _ = v25
        }
    case "":
        if (v3 >= (v3 % ((260) | 1))) {
            switch (v3 % (((v3 << ((2) & 63))) | 1)) {
            case 2147483647:
                var v26 int16 = -137
                _ = v26
            case 0:
                fmt.Println(v2)
            default:
                v27 := GenS3{F0: 831, F1: v3, F2: 691}
                _ = v27
            }
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v2 -= (v2 - 196)
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
}
