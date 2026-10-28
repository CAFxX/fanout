package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint32; F1 int8; F2 bool }
type GenS4 struct { F0 uint32; F1 uint8 }
func genFn0(p0 string, p1 int32) string {
    var v4 uint16 = 2
    v5 := GenS1{F0: 0, F1: 1000}
    fmt.Println(p0, p1)
    if (989 < -958) {
        for fi6 := 0; fi6 < 3; fi6++ {
            fmt.Println(p1, p0)
            if (65535 >= v4) {
                var v7 int64 = 864
                if (-2.25 >= 1e10) {
                    var v8 float64 = 0.1
                    for fi9 := 0; fi9 < 14; fi9++ {
                        var v10 uint32 = 2147483648
                        v10 = (v10 ^ v10)
                        switch -2147483648 {
                        case 484:
                            switch fi9 {
                            case 1:
                                fmt.Println(v8)
                            case 128:
                                fmt.Println(v8, v10, v8)
                            case 127:
                                v11 := []int{fi6, fi9}
                                _ = v11
                            }
                            v12 := GenS1{F0: 4294967295, F1: 380}
                            _ = v12
                        case -1:
                            v8 -= 1.5
                        case 2147483647:
                            var v13 uint8 = 137
                            _ = v13
                        default:
                            v14 := GenS2{F0: (^(v10)), F1: -128}
                            _ = v14
                        }
                        p0 = ""
                        _ = fi9
                        _ = v10
                    }
                    v15 := []int{-2147483648, (fi6 ^ (fi6 >> ((2) & 63))), 271}
                    _ = v15
                    _ = v8
                }
                _ = v7
            } else {
                v16 := GenS1{F0: 1, F1: 380}
                p0 = ("foo bar" + ((p0 + (p0 + p0)) + p0))
                v4 = (^(v4))
                _ = v16
            }
            _ = fi6
        }
    }
    var v17 uint = 65535
    _ = p0
    _ = p1
    _ = v17
    _ = v4
    _ = v5
    return p0
}
func genFn1() bool {
    v18 := GenS1{F0: 0, F1: 32767}
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 int32 = -2147483648
        _ = fi19
        _ = v20
    }
    v21 := GenS3{F0: 2147483647, F1: 1, F2: false}
    _ = v18
    _ = v21
    return (32767 < 2)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 0
    var v3 float32 = float32(1e10)
    for fi22 := 0; fi22 < 6; fi22++ {
        v23 := GenS4{F0: 2, F1: 214}
        var v24 float32 = (v3 - v3)
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int64 = -591
    var v26 float64 = ((0.0 - -2.25) - 0.0)
    var v27 bool = true
    switch 42 {
    case -250:
        v28 := GenS4{F0: 2147483648, F1: 127}
        _ = v28
    case 42:
        v25 -= -1
    case 128:
        v3 += (((v3 * 0.1) + 1e10) + (float32(0.0) - ((((1e10 - (float32((float64(v3) * 3.14159)) + v3)) - float32(float64(v3))) - ((0.0 / float32(((0.1 / v26) - -2.25))) / (-2.25 + v3))) - (v3 * ((-2.25 - (v3 * v3)) + ((v3 + v3) * 3.14159))))))
    }
    var v29 uint16 = 65535
    for fi30 := 0; fi30 < 15; fi30++ {
        for fi31 := 0; fi31 < 3; fi31++ {
            v32 := []int{-2147483648}
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    fmt.Println(v2, v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
}
