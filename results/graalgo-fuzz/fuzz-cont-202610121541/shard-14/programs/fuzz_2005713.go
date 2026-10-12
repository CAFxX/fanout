package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int64; F2 int16 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 string, p1 uint16, p2 int64) uint32 {
    v5 := GenS1{F0: 65535, F1: 17, F2: p2}
    v6 := GenS2{F0: -9223372036854775808}
    v7 := []int{9223372036854775807, 555, 490}
    if ((^(p1)) > 0) {
        v8 := GenS3{F0: 42, F1: (-(p2)), F2: 42}
        _ = v8
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        var v9 uint8 = 192
        for fi10 := 0; fi10 < 1; fi10++ {
            if true {
                if ((p1 - (^(p1))) != 2) {
                    if false {
                        fmt.Println(p2, p0, v9)
                        var v11 int = 32768
                        _ = v11
                    }
                    v12 := GenS4{F0: false, F1: -32768}
                    v9 = v9
                    _ = v12
                }
            } else {
                v13 := GenS5{F0: 1}
                _ = v13
            }
            _ = fi10
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []string{}
    v2 := []int{-1, 32767, -9223372036854775808}
    var v3 float64 = float64((3.14159 + (((3.14159 + 3.14159) + (float32((((((0.0 + 3.14159) / 1.0) + float64((0.0 / 1.0))) + (3.14159 * float64(1e10))) / 1.5)) - -2.25)) + (1.5 * float32(0.1)))))
    v4 := make(map[string]uint32)
    var v14 int = 32768
    var v15 string = "hello"
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := GenS4{F0: true, F1: 1000}
        var v18 string = ("" + v15)
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS4{F0: false, F1: 1}
    for fi20 := 0; fi20 < 16; fi20++ {
        v3 = 3.14159
        v14 = (-(v14))
        _ = fi20
    }
    var v21 uint = 18446744073709551615
    v22 := GenS3{F0: 18446744073709551615, F1: 0, F2: -160}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v2)
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := v4["foo bar"]
        if len(v2) > 0 {
            v2[0] = 244
        }
        _ = fi23
        _ = v24
    }
    var v25 uint8 = 135
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    var v26 uint16 = 284
    var v27 uint16 = v26
    var v28 int64 = int64(v21)
    var v29 uint = v21
    v28 += 787
    v1 = append(v1, v15)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
