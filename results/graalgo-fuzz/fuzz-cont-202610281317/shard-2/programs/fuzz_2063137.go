package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint64; F1 int8; F2 int8 }
type GenS3 struct { F0 int; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 string; F1 string; F2 int32 }
type GenS7 struct { F0 int }
func genFn0(p0 bool) uint8 {
    if p0 {
        if (255 == 65535) {
            v4 := []int{-1, 881, 4294967295}
            p0 = p0
            _ = v4
        }
        fmt.Println(p0)
    } else {
        v5 := []int{1}
        var v6 int64 = 0
        _ = v5
        _ = v6
    }
    p0 = p0
    p0 = true
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 float32 = 3.14159
        var v9 int16 = -318
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    return 237
}

func main() {
    var v1 float64 = float64((1.5 * float32((3.14159 * (0.0 / 1e10)))))
    v2 := make(map[int]bool)
    v3 := []uint32{2147483647, 65535, 1}
    var v10 float64 = v1
    if false {
        if false {
            v11 := v2[104]
            v11 = false
            _ = v11
        }
    }
    _ = len(v3)
    v12 := GenS1{F0: 42, F1: true}
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 float64 = v10
        v15 := GenS2{F0: 18446744073709551615, F1: -55, F2: -22}
        switch (("a" + "a") + "a") {
        case "":
            for fi16 := 0; fi16 < 14; fi16++ {
                v14 = v14
                v17 := GenS3{F0: (fi16 ^ fi16), F1: -9223372036854775808, F2: 0}
                _ = fi16
                _ = v17
            }
        case "hello":
            var v18 int32 = -1
            _ = v18
        case "\n\t":
            switch fi13 {
            case 32767:
                var v19 string = ("a" + ("\n\t" + "a"))
                _ = v19
            case -9223372036854775808:
                fmt.Println(v14)
            default:
                v2[fi13] = ((fi13 & (fi13 & fi13)) < (fi13 & fi13))
            }
        default:
            for fi20 := 0; fi20 < 14; fi20++ {
                v21 := GenS4{F0: 137, F1: true, F2: 255}
                v22 := GenS1{F0: 658, F1: ((-32768 <= 1000) && false)}
                v23 := GenS5{F0: -85, F1: -1}
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    var v24 float64 = (v1 / v10)
    if (700 > 2) {
        v24 = v24
        v1 = v10
    } else {
        fmt.Println(v1)
    }
    if (((!false) || (255 < 2)) && (true || false)) {
        v25 := GenS6{F0: ("" + (((("\n\t" + ("hello" + "")) + "\n\t") + ((("" + "foo bar") + "hello") + "\n\t")) + (("" + ("foo bar" + "foo bar")) + "\n\t"))), F1: "a", F2: 2147483647}
        switch "hello" {
        case "":
            for fi26 := 0; fi26 < 15; fi26++ {
                v1 = v1
                var v27 string = "\n\t"
                fmt.Println(v27, v1, v27)
                _ = fi26
                _ = v27
            }
        case "a":
            v28 := v2[774]
            _ = v28
        case "\n\t":
            v29 := GenS7{F0: -225}
            _ = v29
        default:
            v10 = (-2.25 / ((0.0 + 1e10) - float64((float32(v24) * 1.5))))
        }
        _ = v25
    }
    if true {
        if (-604 != 2147483647) {
            v1 = (float64(-2.25) * v10)
            var v30 int32 = 0
            _ = v30
        } else {
            fmt.Println(v24, v24, v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v24
    _ = v3
}
