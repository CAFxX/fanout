package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 string; F1 int64 }
func genFn0(p0 uint8, p1 int, p2 uint16) uint {
    var v4 uint = 42
    v5 := GenS1{F0: true, F1: -128}
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := []int{(-(fi6)), 2147483647}
        var v8 int = (fi6 * p1)
        p2 = (p2 >> (((v4 + (^(v4)))) & 63))
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2
}
func genFn1() int16 {
    var v9 uint16 = 42
    fmt.Println(v9)
    _ = v9
    return 1000
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 1e10
    v3 := []uint32{289, 42}
    v2 += ((v2 + ((0.0 / (float64((1e10 * 3.14159)) - v2)) * 1.5)) * (v2 * 3.14159))
    v1[-9223372036854775808] = true
    if true {
        v10 := GenS2{F0: 82, F1: 464}
        for fi11 := 0; fi11 < 14; fi11++ {
            if false {
                if (0 <= -109) {
                    var v12 uint32 = 0
                    var v13 int = fi11
                    v14 := v1[(v13 ^ (v13 << ((0) & 63)))]
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    v2 = v2
                }
                if false {
                    v2 += 0.0
                    var v15 int8 = 112
                    _ = v15
                }
                var v16 float64 = 0.0
                _ = v16
            } else {
                v3 = append(v3, 2147483647)
                v2 -= v2
            }
            var v17 int = fi11
            _ = fi11
            _ = v17
        }
        _ = v10
    } else {
        var v18 int8 = 4
        _ = v18
    }
    v19 := v1[127]
    if len(v3) > 0 {
        v3[0] = 849
    }
    var v20 uint16 = 65535
    v21 := GenS1{F0: (409 <= 279), F1: -1}
    _ = len(v1)
    var v22 string = "\n\t"
    v23 := GenS3{F0: v22, F1: -221}
    var v24 uint8 = 255
    v22 = "foo bar"
    var v25 uint16 = (v20 & v20)
    v20 += (^(v20))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
