package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int; F2 string }
type GenS4 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint32) int16 {
    var v4 int64 = -9223372036854775808
    var v5 string = "foo bar"
    v6 := []int{32768, 675}
    var v7 float32 = float32((1e10 * (((3.14159 + 1.5) + (0.1 / 1e10)) * float64((1e10 + (0.1 * float32(0.1)))))))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{2147483647, 0, 2147483647, 9223372036854775807}
    var v3 int8 = -10
    for fi8 := 0; fi8 < 14; fi8++ {
        v3 = v3
        _ = len(v1)
        _ = fi8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v3 -= v3
        _ = fi9
    }
    v3 -= v3
    v10 := GenS1{F0: 1, F1: "\n\t"}
    v2 = append(v2, -1)
    v3 -= (^(v3))
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            v3 = 0
            if (!false) {
                v3 = (-(v3))
                if (v3 == v3) {
                    v13 := GenS2{F0: 1}
                    if false {
                        v14 := v1["hello"]
                        var v15 int8 = 121
                        _ = v14
                        _ = v15
                    }
                    v3 = (v3 - v3)
                    _ = v13
                } else {
                    var v16 uint8 = 19
                    var v17 int = fi11
                    _ = v16
                    _ = v17
                }
            }
            v3 += v3
            _ = fi12
        }
        v18 := GenS1{F0: -710, F1: "\n\t"}
        _ = fi11
        _ = v18
    }
    var v19 int32 = -1
    var v20 uint = 0
    if len(v2) > 0 {
        v2[0] = 32767
    }
    v3 += v3
    for fi21 := 0; fi21 < 7; fi21++ {
        v22 := GenS3{F0: 0, F1: fi21, F2: ""}
        _ = fi21
        _ = v22
    }
    v3 = (v3 / ((-128) | 1))
    if (!(!true)) {
        for fi23 := 0; fi23 < 12; fi23++ {
            v24 := GenS4{F0: 63, F1: 2, F2: (v19 + -2147483648)}
            var v25 bool = (!(!(1 != 334)))
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        if true {
            var v27 float64 = 1.5
            _ = v27
        } else {
            v28 := GenS5{F0: -32768}
            _ = v28
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
