package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int32 }
func genFn0(p0 uint64) bool {
    var v6 float64 = 0.0
    v7 := []int{32767, 755, -1, 32767}
    v8 := GenS1{F0: 42, F1: 138}
    var v9 uint16 = 0
    v6 += ((float64(1e10) * (v6 - v6)) + (v6 / v6))
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1(p0 uint16, p1 string, p2 uint64) string {
    v10 := []int{-137, 9223372036854775807, 32767, 929}
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int8 = 127
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p1
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 9223372036854775807
    var v3 int32 = 42
    var v4 bool = true
    var v5 float32 = -2.25
    switch v2 {
    case 25:
        v13 := []int{472, 405, (^(v2))}
        _ = v13
    case 4294967295:
        if v4 {
            for fi14 := 0; fi14 < 8; fi14++ {
                if v4 {
                    switch (fi14 + (v2 - 128)) {
                    case -1:
                        v15 := GenS2{F0: 1}
                        _ = v15
                    }
                }
                v16 := []int{fi14, 352, (fi14 * (fi14 + (^(fi14))))}
                var v17 int = (^(fi14))
                _ = fi14
                _ = v16
                _ = v17
            }
            v18 := []int{v2, -957, 32767}
            _ = v18
        } else {
            fmt.Println(v4, v1)
        }
    }
    fmt.Println(v1, v5, v4)
    var v19 int8 = -1
    var v20 uint32 = 245
    v21 := []int{(v2 / (((v2 + (^(v2)))) | 1)), -1}
    var v22 uint8 = 179
    var v23 bool = (uint32(v22) != v20)
    var v24 uint64 = 0
    var v25 int8 = (v19 * 42)
    v22 = (v22 / (((v22 ^ 255)) | 1))
    v26 := GenS3{F0: (^(v2)), F1: (-(v3))}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
