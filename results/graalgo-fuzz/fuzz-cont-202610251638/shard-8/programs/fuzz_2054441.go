package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 int }
func genFn0(p0 int64) bool {
    if true {
        v7 := GenS1{F0: 629}
        var v8 int64 = p0
        _ = v7
        _ = v8
    }
    v9 := []int{128, -73, -293, -602}
    v10 := GenS1{F0: 128}
    _ = p0
    _ = v10
    _ = v9
    return false
}
func genFn1(p0 uint8) uint8 {
    p0 += p0
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint64 = 42
        v12 += 894
        _ = fi11
        _ = v12
    }
    p0 += 255
    _ = p0
    return p0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    v3 := []int{}
    var v4 float64 = (float64(3.14159) / 1.5)
    var v5 bool = (1 >= 42)
    var v6 int = 1
    v13 := GenS1{F0: v6}
    if v5 {
        var v14 int8 = -128
        switch ("hello" + "a") {
        case "a":
            switch int(v14) {
            case -2147483648:
                if ((v6 & 497) < (v6 % ((v6) | 1))) {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        for fi16 := 0; fi16 < 4; fi16++ {
                            switch ("a" + "hello") {
                            case "hello":
                                v17 := GenS2{F0: 0, F1: (v4 == 1e10)}
                                _ = v17
                            default:
                                _ = len(v1)
                            }
                            v4 = ((((-2.25 + (-2.25 * v4)) - 0.1) * (0.1 * float64(3.14159))) - (v4 - 0.1))
                            v2[0] = ""
                            v6 -= (fi16 >> ((107) & 63))
                            _ = fi16
                        }
                        _ = fi15
                    }
                } else {
                    switch 42 {
                    case -1:
                        v4 += float64(float32(((v4 * -2.25) - v4)))
                    case 32768:
                        fmt.Println(v6)
                    }
                }
            case 0:
                v4 = (v4 / (v4 / ((v4 * v4) + ((0.0 + ((0.0 / float64(-2.25)) * float64(float32(((((0.0 - v4) + float64(0.0)) + ((-2.25 + float64(1.5)) - (-2.25 + (1.5 / float64(-2.25))))) / 1e10))))) * (1.5 / (v4 / 1.0))))))
            case -9223372036854775808:
                var v18 uint = 4294967295
                _ = v18
            }
        default:
            var v19 int32 = 589
            _ = v19
        }
        var v20 float32 = 1e10
        _ = v14
        _ = v20
    } else {
        for fi21 := 0; fi21 < 16; fi21++ {
            v5 = v5
            _ = fi21
        }
        if (42 >= 2147483647) {
            v4 = 0.1
        }
        if v5 {
            var v22 float64 = (3.14159 * float64(-2.25))
            var v23 uint32 = 4294967295
            var v24 float32 = (((0.0 + 0.1) / 0.1) / -2.25)
            _ = v22
            _ = v23
            _ = v24
        } else {
            switch "" {
            case "":
                if (v6 != (v6 << ((42) & 63))) {
                    fmt.Println(v6, v4)
                    fmt.Println(v4, v5, v4)
                }
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "a":
                _ = len(v2)
            default:
                if v5 {
                    if (42 >= 42) {
                        var v25 uint16 = 2
                        var v26 uint = 907
                        _ = v25
                        _ = v26
                    }
                    if v5 {
                        v5 = v5
                        v27 := GenS2{F0: genFn1(216), F1: false}
                        v28 := GenS2{F0: 1, F1: (v5 || v5)}
                        _ = v27
                        _ = v28
                    }
                    v29 := GenS3{F0: (v6 - v6), F1: 749}
                    _ = v29
                } else {
                    v5 = false
                    v5 = v5
                    for fi30 := 0; fi30 < 5; fi30++ {
                        v5 = (int64(fi30) == -9223372036854775808)
                        _ = fi30
                    }
                }
            }
        }
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        if len(v3) > 0 {
            v3[0] = fi31
        }
        v32 := GenS4{F0: 65535}
        _ = fi31
        _ = v32
    }
    v5 = (v5 && v5)
    fmt.Println(v5)
    v33 := GenS3{F0: (^(v6)), F1: 974}
    var v34 uint16 = 255
    fmt.Println(v4, v4, v4)
    if (((v4 / v4) * v4) != v4) {
        var v35 int8 = -128
        v35 = v35
        _ = v35
    }
    v36 := GenS5{F0: 879, F1: (v6 ^ v6)}
    v5 = (901 < 2147483648)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
