package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int; F1 int; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 int, p1 string) int {
    var v6 int8 = -12
    v6 += v6
    _ = p0
    _ = p1
    _ = v6
    return p0
}
func genFn1(p0 int8) bool {
    for fi7 := 0; fi7 < 1; fi7++ {
        switch (fi7 | (fi7 >> ((686) & 63))) {
        case 127:
            v8 := []int{-2147483648, -1}
            _ = v8
        default:
            v9 := GenS1{F0: 1, F1: (^(p0))}
            _ = v9
        }
        if false {
            p0 -= p0
            v10 := []int{fi7, 1}
            _ = v10
        }
        var v11 int8 = (^(p0))
        _ = fi7
        _ = v11
    }
    var v12 int16 = -529
    _ = p0
    _ = v12
    return (!true)
}

func main() {
    var v1 int8 = 106
    var v2 string = ("" + ("hello" + "\n\t"))
    var v3 bool = true
    v4 := make(map[string]uint32)
    var v5 int64 = 1
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint = 42
        var v15 string = v2
        switch ("a" + v15) {
        case "foo bar":
            if v3 {
                v16 := []int{(fi13 | 2147483647), (fi13 ^ (fi13 % ((int(v14)) | 1))), (-(fi13)), -557}
                if v3 {
                    v1 += (v1 * v1)
                    for fi17 := 0; fi17 < 2; fi17++ {
                        v14 = (^(v14))
                        v18 := GenS2{F0: (^(fi13)), F1: genFn0(-593, ("a" + "a")), F2: 255}
                        v19 := v4[v2]
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    v14 += (v14 & v14)
                }
                v16 = append(v16, fi13)
                _ = v16
            }
        case "a":
            v3 = (!v3)
        case "":
            v3 = v3
        default:
            v20 := GenS1{F0: 576, F1: v1}
            _ = v20
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v21 := []int{-35, -886}
    v1 -= v1
    var v22 string = (("a" + v2) + ("\n\t" + (v2 + "\n\t")))
    v1 += 127
    var v23 int32 = 397
    var v24 int8 = (v1 + 88)
    v25 := GenS3{F0: 0, F1: (-(v1))}
    _ = len(v21)
    var v26 float32 = -2.25
    v21 = append(v21, 9223372036854775807)
    var v27 int64 = 0
    v28 := GenS4{F0: (122 != -9223372036854775808), F1: 42}
    var v29 uint = 4294967295
    var v30 float32 = ((v26 - (((-2.25 * (v26 / (v26 * (v26 / v26)))) * 1e10) - v26)) * 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
