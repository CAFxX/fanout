package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 65535
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 5; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    v5 := GenS1{F0: 578}
    switch -573 {
    case 9223372036854775807:
        var v6 int16 = 32767
        _ = v6
    default:
        if true {
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 uint64 = 171
                var v9 int64 = -360
                _ = fi7
                _ = v8
                _ = v9
            }
        }
    }
    v10 := GenS1{F0: 32768}
    v11 := GenS1{F0: -131}
    if (2147483648 <= 0) {
        v12 := []int{97, -390}
        _ = v12
    } else {
        fmt.Println(v2)
        fmt.Println(v2)
        var v13 int = 42
        _ = v13
    }
    var v14 float64 = float64((((1.5 * 0.1) + (0.1 + (float32(0.1) + 0.1))) + (0.1 / ((0.0 + 1.5) + (0.0 + 3.14159)))))
    var v15 float32 = (1e10 / 0.1)
    v14 = 0.0
    if ((v14 + 0.1) < v14) {
        if (1 > 1000) {
            v2 -= v2
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            if (float64((0.1 / ((v15 + -2.25) - -2.25))) != v14) {
                var v17 float32 = v15
                v2 += 255
                v15 = v17
                _ = v17
            } else {
                if (0 > 286) {
                    v14 = ((v14 * v14) / v14)
                    v14 -= v14
                    v2 = (v2 & (v2 | (v2 >> ((1) & 63))))
                } else {
                    v18 := []int{(fi16 ^ fi16), fi16, (^(fi16))}
                    _ = v18
                }
            }
            _ = fi16
        }
    }
    v19 := []int{}
    if (true && true) {
        switch 671 {
        case -9223372036854775808:
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 int64 = 0
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                _ = fi20
                _ = v21
            }
        default:
            v15 = -2.25
        }
    }
    v15 += (v15 + 1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v5
}
