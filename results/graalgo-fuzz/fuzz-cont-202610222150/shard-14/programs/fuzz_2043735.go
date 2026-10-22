package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16 }
func genFn0(p0 bool, p1 uint64) uint8 {
    v4 := GenS1{F0: 2, F1: ""}
    p1 = (p1 + p1)
    _ = p0
    _ = p1
    _ = v4
    return 199
}
func genFn1(p0 bool, p1 int32) int16 {
    p1 += 355
    fmt.Println(p0, p0)
    var v5 float32 = 0.1
    _ = p0
    _ = p1
    _ = v5
    return 32767
}

func main() {
    var v1 int16 = 1
    var v2 uint64 = 177
    v3 := make(map[string]uint32)
    v2 -= (v2 % ((v2) | 1))
    for fi6 := 0; fi6 < 10; fi6++ {
        v2 = 416
        v1 += v1
        _ = fi6
    }
    var v7 int = 621
    v1 = (v1 + (-(v1)))
    v1 = int16(v7)
    var v8 int32 = 0
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(fi9, v7)
        _ = fi9
    }
    v10 := GenS2{F0: 42, F1: 32767}
    var v11 uint64 = (^(v2))
    var v12 uint16 = 255
    if false {
        v13 := []int{}
        if ((-1 == 9223372036854775807) && false) {
            v14 := GenS3{F0: 255}
            if false {
                var v15 uint64 = v11
                for fi16 := 0; fi16 < 16; fi16++ {
                    for fi17 := 0; fi17 < 16; fi17++ {
                        var v18 string = ""
                        v18 = "\n\t"
                        switch "\n\t" {
                        case "foo bar":
                            switch fi17 {
                            case 0:
                                if ((v12 / (((v12 >> ((1) & 63))) | 1)) < v12) {
                                    var v19 uint = 1
                                    var v20 uint = v19
                                    var v21 int8 = 100
                                    switch fi17 {
                                    case -9223372036854775808:
                                        if false {
                                            if true {
                                                v15 = v11
                                                if true {
                                                    fmt.Println(fi17)
                                                }
                                                var v22 float64 = -2.25
                                                _ = v22
                                            }
                                        }
                                    }
                                    v2 += 206
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                            case 9223372036854775807:
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                            case 32768:
                                var v23 float64 = (0.0 + 1.5)
                                _ = v23
                            }
                        default:
                            v12 -= v12
                        }
                        v1 = (v1 * v1)
                        _ = len(v3)
                        _ = fi17
                        _ = v18
                    }
                    _ = fi16
                }
                if (-9223372036854775808 < 32767) {
                    v7 = (v7 % ((v7) | 1))
                }
                _ = v15
            } else {
                for fi24 := 0; fi24 < 5; fi24++ {
                    fmt.Println(v7, v1, v12)
                    _ = fi24
                }
            }
            _ = v14
        }
        var v25 uint64 = v11
        _ = v13
        _ = v25
    } else {
        delete(v3, "hello")
        var v26 uint16 = (v12 | v12)
        _ = v26
    }
    v27 := GenS4{F0: -1}
    v11 += v11
    var v28 float64 = float64(float32(0.0))
    if false {
        v29 := []int{(v7 * (v7 & 0))}
        _ = v29
    }
    var v30 uint64 = v11
    v12 = (v12 % ((v12) | 1))
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v7
    _ = v8
}
