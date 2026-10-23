package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint64; F1 uint16; F2 bool }
type GenS6 struct { F0 int8; F1 int; F2 int64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    var v5 float64 = 0.0
    if (!false) {
        v5 = v5
    }
    v5 += 3.14159
    for fi6 := 0; fi6 < 20; fi6++ {
        switch "\n\t" {
        case "hello":
            v5 = (1.5 - (float64(((0.0 + 3.14159) - float32(((((3.14159 * -2.25) + 1.5) - 0.1) * (v5 + ((0.0 * (1.5 - v5)) - -2.25)))))) - 0.1))
        default:
            for fi7 := 0; fi7 < 6; fi7++ {
                v5 = 1.5
                v8 := GenS1{F0: 42}
                v5 = (float64((((((0.0 / 1e10) * -2.25) + -2.25) * 0.1) - 3.14159)) * v5)
                _ = fi7
                _ = v8
            }
        }
        v5 = v5
        if false {
            v5 -= (((v5 / 1e10) + (float64(0.0) / (float64(1e10) - (-2.25 * v5)))) + 1.5)
        }
        _ = fi6
    }
    var v9 string = ("hello" + "hello")
    v10 := []int{-1, 4294967295}
    var v11 uint16 = 770
    fmt.Println(v5)
    v12 := GenS2{F0: 0, F1: 91}
    v5 = (v5 + (0.1 / 1e10))
    for fi13 := 0; fi13 < 14; fi13++ {
        v5 -= (v5 + v5)
        v14 := GenS3{F0: int8(v11), F1: -32768}
        _ = len(v10)
        _ = fi13
        _ = v14
    }
    switch 1 {
    case 974:
        v15 := GenS2{F0: 0, F1: 0}
        _ = v15
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 6; fi17++ {
            if (100 <= -17) {
                var v18 uint = 135
                var v19 uint64 = 42
                v9 = v9
                _ = v18
                _ = v19
            }
            var v20 uint32 = 2
            _ = fi17
            _ = v20
        }
        v5 -= float64(-2.25)
        _ = fi16
    }
    fmt.Println(v11, v11, v9)
    v2[32768] = true
    if (((v5 / v5) + float64((1e10 * 1e10))) > v5) {
        v11 -= (v11 * 255)
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    } else {
        if true {
            if (true || false) {
                v5 = float64((((1.5 - -2.25) / (1e10 / 1.0)) * 0.0))
                if (9223372036854775807 <= 0) {
                    fmt.Println(v5, v5, v5)
                    switch ("hello" + "") {
                    case "":
                        if len(v10) > 0 {
                            v10[0] = 0
                        }
                    case "a":
                        v11 -= v11
                    default:
                        fmt.Println(v5, v11, v5)
                    }
                } else {
                    var v21 float32 = 1e10
                    v11 = 255
                    _ = v21
                }
                v22 := GenS4{F0: 69, F1: 42}
                _ = v22
            }
            v23 := GenS5{F0: 18446744073709551615, F1: (v11 >> ((255) & 63)), F2: false}
            if false {
                if (float32((v5 + v5)) < ((1e10 + 1.5) * 3.14159)) {
                    v10 = append(v10, 9223372036854775807)
                }
                var v24 bool = (("hello" + "a") <= (v9 + "hello"))
                v24 = v24
                _ = v24
            } else {
                v25 := GenS6{F0: 127, F1: 42, F2: 0}
                v26 := GenS6{F0: 0, F1: -2147483648, F2: 54}
                _ = v25
                _ = v26
            }
            _ = v23
        }
    }
    fmt.Println(v11, v11, v9)
    v27 := v3[v9]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
