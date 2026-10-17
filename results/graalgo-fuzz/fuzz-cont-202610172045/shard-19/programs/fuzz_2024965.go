package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int32; F1 int8 }
func genFn0(p0 uint32) int16 {
    p0 = p0
    p0 -= p0
    p0 += (p0 / (((p0 ^ p0)) | 1))
    fmt.Println(p0)
    _ = p0
    return 1000
}
func genFn1(p0 string) uint16 {
    p0 = p0
    for fi6 := 0; fi6 < 13; fi6++ {
        if (-302 == -802) {
            p0 = "a"
        }
        var v7 uint16 = 2
        _ = fi6
        _ = v7
    }
    var v8 int = -2147483648
    _ = p0
    _ = v8
    return 255
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 0.0
    v3 := make(map[string]uint32)
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    switch "a" {
    case "foo bar":
        v9 := GenS1{F0: -1}
        _ = v9
    case "hello":
        fmt.Println(v2)
    case "":
        v2 = (v2 * (v2 / (3.14159 - v2)))
    }
    v10 := GenS2{F0: 833}
    var v11 float32 = float32((-2.25 * -2.25))
    fmt.Println(v2)
    if (1.5 >= v11) {
        v11 += (float32(float64((v2 / v11))) + (3.14159 / (v11 * float32(3.14159))))
        for fi12 := 0; fi12 < 16; fi12++ {
            if (true && (!(821 < 1))) {
                var v13 uint16 = 2
                v14 := []int{fi12, fi12}
                _ = v13
                _ = v14
            }
            fmt.Println(fi12)
            for fi15 := 0; fi15 < 14; fi15++ {
                v2 += ((-2.25 * float32(0.1)) * -2.25)
                _ = fi15
            }
            _ = fi12
        }
    } else {
        v11 = (float32(0.1) / (-2.25 * float32(0.1)))
        v11 += (float32(float64(0.1)) + v11)
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        v4[fi16] = "hello"
        if (!false) {
            v11 = (v11 + float32(((1e10 + 1e10) * -2.25)))
            fmt.Println(v11)
        } else {
            v11 += (v2 * (v2 + (v11 / v11)))
            for fi17 := 0; fi17 < 1; fi17++ {
                v18 := GenS3{F0: 32767, F1: 42}
                _ = fi17
                _ = v18
            }
            if false {
                for fi19 := 0; fi19 < 11; fi19++ {
                    var v20 uint8 = 12
                    switch (fi19 + (fi19 * 2147483647)) {
                    case 4294967295:
                        _ = len(v1)
                    case -2147483648:
                        switch "hello" {
                        case "hello":
                            v4[(fi19 * fi16)] = "hello"
                        }
                    default:
                        v2 = (v2 * (1.5 * (1.5 + v2)))
                    }
                    var v21 int32 = 42
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                fmt.Println(fi16, v11)
            } else {
                v22 := GenS2{F0: 0}
                _ = v22
            }
        }
        _ = fi16
    }
    var v23 float64 = (-2.25 / 0.1)
    var v24 int = 32767
    v25 := GenS4{F0: 2147483647, F1: -93}
    var v26 uint8 = 255
    v26 -= 95
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
