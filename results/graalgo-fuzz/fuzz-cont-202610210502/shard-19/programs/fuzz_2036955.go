package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint }
func genFn0() uint64 {
    if true {
        var v7 bool = (-941 == 42)
        _ = v7
    }
    var v8 int16 = 42
    v8 -= v8
    _ = v8
    return 906
}
func genFn1(p0 uint16, p1 string) int32 {
    var v9 uint32 = 117
    p0 = 65535
    v9 += 4294967295
    if (716 < 639) {
        p0 -= p0
        fmt.Println(p1, p1, v9)
    } else {
        var v10 float32 = 1.5
        _ = v10
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 uint8 = 0
        v12 = 255
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 0
    var v3 int16 = 1000
    v4 := make(map[string]uint32)
    var v5 float32 = 1.5
    var v6 string = "a"
    var v13 bool = false
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 float32 = 0.0
        var v16 int64 = 9223372036854775807
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: v6, F1: v13, F2: 0}
    var v18 uint8 = 255
    fmt.Println(v6, v5, v13)
    fmt.Println(v13, v2, v13)
    var v19 bool = (193 == 2)
    v6 = v6
    v20 := v1[1]
    var v21 float64 = (1.5 * (1e10 / (0.0 / ((((1.5 / 1.5) + 0.1) / (3.14159 + 1.5)) * ((((((((1.5 / 1.0) / 3.14159) + 0.1) / 3.14159) + 1e10) * (3.14159 + ((0.0 - (3.14159 * 1.5)) / float64(v5)))) / float64(((v5 * -2.25) + v5))) + 3.14159)))))
    var v22 uint8 = (v18 % ((127) | 1))
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 uint64 = 893
        var v25 int16 = (^(v3))
        var v26 uint16 = 32767
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    if false {
        v27 := GenS1{F0: (v6 + (v6 + (((("\n\t" + v6) + ((v20 + ("hello" + (v6 + v20))) + v6)) + "hello") + (v6 + v6)))), F1: true, F2: 4294967295}
        _ = v27
    } else {
        for fi28 := 0; fi28 < 8; fi28++ {
            if (65535 != 18446744073709551615) {
                v3 -= (v3 ^ 17)
                v29 := GenS1{F0: v20, F1: ((^(v2)) <= -573), F2: 2}
                switch (fi28 & (fi28 / (((fi28 | (fi28 % (((fi28 * -2147483648)) | 1)))) | 1))) {
                case 4294967295:
                    v30 := GenS2{F0: 386}
                    _ = v30
                case 647:
                    if (fi28 >= 2147483647) {
                        var v31 uint = 4294967295
                        _ = v31
                    }
                default:
                    v19 = true
                }
                _ = v29
            }
            _ = fi28
        }
        v5 = v5
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
