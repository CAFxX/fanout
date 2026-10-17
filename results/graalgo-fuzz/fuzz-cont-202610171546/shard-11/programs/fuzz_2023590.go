package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint; F1 int8; F2 uint8 }
type GenS3 struct { F0 int; F1 int16; F2 int }
func genFn0(p0 int64, p1 int64) int64 {
    var v6 string = "a"
    var v7 uint = 930
    var v8 uint = (^(v7))
    v9 := []int{9223372036854775807, -9223372036854775808, 42, -406}
    var v10 uint8 = 2
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int, p1 string) uint16 {
    switch p1 {
    case "foo bar":
        v11 := GenS1{F0: 576, F1: 1, F2: 255}
        _ = v11
    default:
        v12 := []int{-793, (p0 * p0)}
        _ = v12
    }
    switch (("\n\t" + p1) + (p1 + (("hello" + p1) + p1))) {
    case "a":
        v13 := GenS1{F0: 65535, F1: -128, F2: 42}
        _ = v13
    default:
        var v14 uint = 0
        _ = v14
    }
    v15 := GenS2{F0: 647, F1: 20, F2: 127}
    _ = p0
    _ = p1
    _ = v15
    return 2
}

func main() {
    var v1 string = "foo bar"
    var v2 int8 = 126
    var v3 int = 32768
    var v4 int16 = -32768
    var v5 float64 = 0.1
    var v16 float64 = (v5 + (float64(((0.0 * 0.0) / float32(float64(1.5)))) - (3.14159 - float64((3.14159 - ((0.1 / (0.0 - 3.14159)) + 3.14159))))))
    v17 := []int{(v3 + v3), (v3 >> ((864) & 63)), v3, (v3 * v3)}
    fmt.Println(v1)
    for fi18 := 0; fi18 < 10; fi18++ {
        v4 = v4
        if true {
            fmt.Println(v4, fi18, v2)
            switch (-(v3)) {
            case -368:
                var v19 int16 = -350
                _ = v19
            case 2147483647:
                if (!true) {
                    var v20 uint16 = 255
                    _ = v20
                }
            default:
                v21 := GenS1{F0: 409, F1: (v2 * v2), F2: 42}
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            v16 -= (float64(0.1) - v5)
            if len(v17) > 0 {
                v17[0] = fi22
            }
            if (33 > 127) {
                var v23 float64 = 1.5
                _ = v23
            }
            _ = fi22
        }
        _ = fi18
    }
    v24 := GenS3{F0: (v3 / ((v3) | 1)), F1: (v4 / (((v4 >> ((730) & 63))) | 1)), F2: v3}
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    _ = len(v17)
    var v25 bool = (v4 > -605)
    var v26 uint64 = 1
    v25 = v25
    for fi27 := 0; fi27 < 14; fi27++ {
        if (32767 < (v3 + fi27)) {
            v5 = (1e10 + v16)
            v2 = (v2 - 1)
        }
        if v25 {
            v28 := GenS3{F0: fi27, F1: (v4 | int16(v26)), F2: -2147483648}
            switch fi27 {
            case -338:
                _ = len(v17)
            case 32767:
                v1 = "\n\t"
            }
            fmt.Println(v5, v16)
            _ = v28
        }
        _ = fi27
    }
    v26 = (v26 + (v26 + v26))
    var v29 int8 = 0
    v1 = v1
    var v30 uint16 = 662
    var v31 uint32 = 2147483648
    _ = len(v17)
    if (false && v25) {
        v5 -= float64(float32(0.1))
    } else {
        var v32 float64 = (((v5 * v16) / (float64(0.0) / (v5 - v5))) + (((1e10 - v16) + v5) * v5))
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
