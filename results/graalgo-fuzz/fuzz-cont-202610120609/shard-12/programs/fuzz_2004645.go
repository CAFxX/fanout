package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint64 }
func genFn0() uint32 {
    if (!false) {
        v5 := GenS1{F0: 42, F1: (-161 <= 2147483647), F2: 65}
        _ = v5
    } else {
        var v6 float64 = (0.0 + (1e10 + (-2.25 + 1e10)))
        _ = v6
    }
    v7 := []int{-2147483648, 232, 1}
    _ = v7
    return 2147483648
}
func genFn1(p0 int16, p1 bool) uint32 {
    if p1 {
        v8 := []int{-352, -281, 42, -1}
        _ = v8
    }
    v9 := GenS1{F0: (p0 % ((1000) | 1)), F1: true, F2: 1}
    var v10 uint8 = 0
    var v11 uint32 = 42
    for fi12 := 0; fi12 < 6; fi12++ {
        v10 -= v10
        v13 := GenS1{F0: 1, F1: false, F2: (v10 & (v10 - v10))}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return (v11 * 42)
}

func main() {
    var v1 uint16 = 899
    v2 := []int{32768, 0, 610, 42}
    v3 := make(map[int]uint32)
    var v4 int32 = -2147483648
    if true {
        for fi14 := 0; fi14 < 18; fi14++ {
            switch (fi14 - -975) {
            case -9223372036854775808:
                if len(v2) > 0 {
                    v2[0] = (^(fi14))
                }
            }
            _ = fi14
        }
    }
    var v15 int16 = 32767
    if ((v4 + (v4 >> ((18446744073709551615) & 63))) >= (v4 + (^(v4)))) {
        var v16 uint64 = 656
        fmt.Println(v16, v1, v16)
        _ = v16
    }
    v17 := GenS2{F0: 42}
    switch ("\n\t" + "\n\t") {
    case "a":
        var v18 int8 = 42
        _ = v18
    }
    v4 = (v4 | v4)
    fmt.Println(v15, v15)
    v19 := GenS2{F0: 255}
    if (true && false) {
        v15 += 0
    } else {
        switch 4294967295 {
        case 273:
            v4 -= -1
        case 127:
            v15 -= int16(v4)
        default:
            var v20 float64 = 0.0
            _ = v20
        }
        if false {
            if false {
                v15 = 1
                v15 = (v15 % ((v15) | 1))
                delete(v3, -9223372036854775808)
            }
            fmt.Println(v15)
        } else {
            if true {
                fmt.Println(v15, v1, v1)
            } else {
                switch 42 {
                case 128:
                    var v21 string = "hello"
                    _ = v21
                case 1:
                    v22 := GenS2{F0: 1}
                    _ = v22
                case 638:
                    v1 -= 65535
                }
                v23 := v3[2147483647]
                _ = v23
            }
            v24 := GenS3{F0: false, F1: 4294967295}
            switch "" {
            case "hello":
                v4 += (v4 / (((v4 * v4)) | 1))
            }
            _ = v24
        }
        v15 += (v15 / ((1) | 1))
    }
    v25 := GenS1{F0: v15, F1: (-12 != 42), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
