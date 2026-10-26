package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint }
type GenS2 struct { F0 int64; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint64; F1 int; F2 uint8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int64, p1 uint16, p2 int32) uint8 {
    p0 = p0
    p2 += -2147483648
    _ = p0
    _ = p1
    _ = p2
    return 67
}

func main() {
    v1 := make(map[int]int)
    var v2 string = ("foo bar" + "hello")
    v3 := make(map[string]int)
    var v4 uint = 1
    v5 := []int{1, 97, 2147483647, -537}
    var v6 string = v2
    v6 = "foo bar"
    if false {
        fmt.Println(v6)
    } else {
        var v7 float32 = (float32(float64(3.14159)) * -2.25)
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 11; fi9++ {
                fmt.Println(fi9)
                fmt.Println(v7, fi8)
                if (!true) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v4 -= 1
                    var v10 uint16 = 255
                    _ = v10
                }
                v11 := GenS1{F0: true, F1: (fi9 + fi9), F2: 42}
                switch fi8 {
                case 2147483647:
                    v12 := v1[fi8]
                    _ = v12
                case 288:
                    v6 = v2
                case 0:
                    var v13 bool = false
                    _ = v13
                }
                switch (^(fi8)) {
                case 4294967295:
                    var v14 uint32 = 305
                    _ = v14
                case 32768:
                    if true {
                        if (-2.25 < 1e10) {
                            v2 = v2
                        } else {
                            fmt.Println(v7)
                            v15 := GenS2{F0: 9223372036854775807, F1: 534, F2: -539}
                            _ = v15
                        }
                        var v16 uint = (v4 / (((v4 ^ 65535)) | 1))
                        _ = v16
                    }
                default:
                    v6 = (v2 + v2)
                }
                _ = fi9
                _ = v11
            }
            var v17 uint = 4294967295
            _ = fi8
            _ = v17
        }
        _ = v7
    }
    v18 := GenS3{F0: 42, F1: 127, F2: 255}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        v19 := GenS2{F0: 0, F1: 1, F2: -2147483648}
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 float32 = (-2.25 / (3.14159 / 1.0))
            v6 = v6
            _ = fi20
            _ = v21
        }
        _ = v19
    } else {
        fmt.Println(v4, v4, v2)
        v22 := GenS4{F0: -566}
        _ = v22
    }
    v23 := GenS2{F0: -327, F1: 255, F2: 2147483647}
    _ = len(v3)
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v24 uint8 = 1
    v2 = v6
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch 2147483647 {
    case -478:
        v25 := GenS2{F0: 9223372036854775807, F1: 780, F2: 2147483647}
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
