package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int }
type GenS3 struct { F0 int; F1 int; F2 bool }
func genFn0(p0 uint64, p1 uint64, p2 int8) uint32 {
    var v4 string = ("a" + "foo bar")
    fmt.Println(p0)
    var v5 int8 = p2
    v6 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int16, p1 uint8) uint8 {
    v7 := GenS1{F0: 1, F1: 0}
    p0 -= p0
    p1 += p1
    var v8 int8 = -117
    v9 := []int{32767, 32768, 169, 323}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 uint16 = 73
    var v2 float32 = 1.5
    var v3 uint16 = (^(v1))
    v10 := GenS2{F0: 303, F1: 1, F2: 787}
    v11 := []int{-2147483648, 127, 4294967295}
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            if false {
                var v14 float32 = 1.5
                var v15 int64 = 1
                var v16 uint = 442
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                var v17 int32 = 340
                _ = v14
                _ = v15
                _ = v16
                _ = v17
            } else {
                var v18 float32 = 1e10
                _ = v18
            }
            _ = fi13
        }
        _ = fi12
    }
    v19 := GenS2{F0: 42, F1: 42, F2: 2147483647}
    var v20 uint = 255
    v1 = (v1 / (((v3 + 783)) | 1))
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 uint16 = (v1 & v1)
        var v23 uint = 1
        _ = fi21
        _ = v22
        _ = v23
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        _ = len(v11)
        var v25 bool = (!(v20 > (^(v20))))
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        var v27 string = "\n\t"
        var v28 float32 = v2
        v29 := GenS1{F0: 42, F1: 42}
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    var v30 float64 = (0.1 / float64(((1.5 / 1.0) - float32((-2.25 + -2.25)))))
    fmt.Println(v1, v1)
    var v31 int16 = -297
    for fi32 := 0; fi32 < 15; fi32++ {
        v33 := GenS3{F0: -87, F1: (fi32 << (((v20 - (^(v20)))) & 63)), F2: false}
        if false {
            v34 := GenS2{F0: genFn0(18446744073709551615, 1, -1), F1: 192, F2: (fi32 >> (((v20 ^ uint(v1))) & 63))}
            _ = v34
        } else {
            v35 := GenS2{F0: 65535, F1: genFn1(v31, 127), F2: (fi32 * fi32)}
            var v36 uint8 = 255
            _ = v35
            _ = v36
        }
        _ = fi32
        _ = v33
    }
    v30 += (1e10 + (v30 + v30))
    fmt.Println(v3, v1)
    v20 += (v20 | 255)
    var v37 float32 = (v2 * v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v37
}
