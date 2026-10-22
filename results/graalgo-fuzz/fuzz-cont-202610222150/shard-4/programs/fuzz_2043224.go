package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint32; F1 uint64; F2 int32 }
func genFn0(p0 bool, p1 int32, p2 uint64) bool {
    var v5 float32 = -2.25
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int32 = 1
        var v8 int16 = int16(p2)
        fmt.Println(p0)
        _ = fi6
        _ = v7
        _ = v8
    }
    p2 = p2
    if p0 {
        v9 := []int{}
        var v10 uint8 = 254
        v10 = v10
        _ = v10
        _ = v9
    }
    v5 += float32(-2.25)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (128 < 255)
}

func main() {
    var v1 int32 = 0
    var v2 float64 = (1e10 * (float64(3.14159) * float64(1e10)))
    var v3 bool = (18446744073709551615 > 165)
    var v4 uint64 = 933
    var v11 uint64 = (v4 - v4)
    var v12 int16 = 1000
    fmt.Println(v11, v12, v3)
    v13 := GenS1{F0: 353}
    fmt.Println(v11, v2)
    v2 -= v2
    v14 := []int{-808, 368}
    var v15 uint64 = 2
    v16 := GenS1{F0: 255}
    v17 := GenS1{F0: 1}
    v15 = v15
    var v18 uint32 = 65535
    var v19 uint16 = 42
    var v20 uint8 = uint8(v19)
    fmt.Println(v12, v18)
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 float64 = v2
        switch "foo bar" {
        case "":
            switch fi21 {
            case -1:
                for fi23 := 0; fi23 < 8; fi23++ {
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    v2 = v2
                    var v24 int16 = -888
                    _ = fi23
                    _ = v24
                }
            case 390:
                if (!v3) {
                    if len(v14) > 0 {
                        v14[0] = fi21
                    }
                } else {
                    _ = len(v14)
                    switch "a" {
                    case "hello":
                        v1 = v1
                    case "":
                        v25 := GenS2{F0: (v12 & v12), F1: (v20 == v20)}
                        _ = v25
                    case "a":
                        v15 -= (v11 % (((v11 & (v15 - v15))) | 1))
                    }
                }
            case -218:
                var v26 uint = 0
                _ = v26
            }
        }
        if (v20 > v20) {
            switch "a" {
            case "hello":
                for fi27 := 0; fi27 < 2; fi27++ {
                    fmt.Println(v4)
                    fmt.Println(v19)
                    _ = fi27
                }
            case "\n\t":
                var v28 uint64 = v4
                _ = v28
            case "a":
                if (v3 && ((^(v1)) > (^(v1)))) {
                    v29 := GenS3{F0: uint32(v4), F1: uint64(v20), F2: v1}
                    v18 = (v18 | 42)
                    var v30 bool = false
                    _ = v29
                    _ = v30
                } else {
                    var v31 int8 = 42
                    _ = v31
                }
            default:
                fmt.Println(fi21, v22)
            }
            var v32 int8 = 0
            _ = v32
        }
        _ = fi21
        _ = v22
    }
    var v33 int16 = (-(v12))
    v19 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v4
}
