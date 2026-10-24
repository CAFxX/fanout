package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 uint }
func genFn0(p0 int, p1 uint, p2 int16) uint16 {
    fmt.Println(p0, p0, p0)
    switch "\n\t" {
    case "hello":
        for fi5 := 0; fi5 < 5; fi5++ {
            if ((false || (973 < -9223372036854775808)) && false) {
                v6 := []int{(p0 / (((-(fi5))) | 1)), fi5, (fi5 >> ((p1) & 63))}
                _ = v6
            } else {
                var v7 uint8 = 42
                _ = v7
            }
            _ = fi5
        }
    case "foo bar":
        v8 := GenS1{F0: 461}
        _ = v8
    case "":
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := []int{-62, fi9}
            var v11 uint = p1
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        switch 1 {
        case 127:
            switch -659 {
            case 128:
                var v12 uint64 = 0
                _ = v12
            case -831:
                v13 := GenS2{F0: uint8(p1), F1: p1, F2: p0}
                _ = v13
            case 0:
                v14 := []int{}
                _ = v14
            default:
                v15 := GenS2{F0: 2, F1: p1, F2: (-(p0))}
                _ = v15
            }
        case -1:
            p1 = p1
        case 42:
            var v16 string = ""
            _ = v16
        default:
            for fi17 := 0; fi17 < 10; fi17++ {
                if (!(248 == 42)) {
                    v18 := GenS3{F0: 255}
                    p0 -= 127
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 13; fi19++ {
                        v20 := []int{(^(p0)), (-(fi17)), (^(fi17)), -1}
                        var v21 uint16 = 32767
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 15; fi22++ {
                        p0 = 2147483647
                        _ = fi22
                    }
                    switch "" {
                    case "hello":
                        for fi23 := 0; fi23 < 13; fi23++ {
                            p1 = p1
                            v24 := []int{fi17}
                            _ = fi23
                            _ = v24
                        }
                    case "":
                        p1 = 65535
                    case "\n\t":
                        var v25 bool = true
                        _ = v25
                    default:
                        if true {
                            var v26 int64 = -1
                            p1 += (p1 / (((p1 % (((p1 / (((p1 + (p1 | 0))) | 1))) | 1))) | 1))
                            p1 -= (p1 >> ((932) & 63))
                            _ = v26
                        } else {
                            v27 := GenS4{F0: 2}
                            var v28 int8 = -128
                            _ = v27
                            _ = v28
                        }
                    }
                }
                var v29 int32 = int32(p1)
                _ = fi17
                _ = v29
            }
        }
    }
    if (true || (2 <= 255)) {
        v30 := GenS3{F0: 0}
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 int8 = 100
    var v2 uint8 = 127
    var v3 int32 = 42
    var v4 float64 = float64(float32((1.5 / ((-2.25 + 3.14159) * 0.1))))
    v31 := []int{0, 4294967295}
    v32 := GenS5{F0: (^(v3)), F1: 2}
    v33 := GenS3{F0: (v2 * v2)}
    v3 = 2147483647
    v1 = (v1 * (v1 + (v1 >> ((0) & 63))))
    var v34 uint = 65535
    v3 += v3
    if false {
        if len(v31) > 0 {
            v31[0] = 4294967295
        }
    } else {
        var v35 int8 = (-(v1))
        v1 = (v1 - (v35 ^ 127))
        _ = v35
    }
    v36 := GenS6{F0: (v34 | (^(v34)))}
    v31 = append(v31, 42)
    switch -1 {
    case 32768:
        if true {
            if ((v1 * v1) < v1) {
                fmt.Println(v1)
            }
            v37 := GenS1{F0: 400}
            _ = v37
        }
    case -617:
        var v38 uint64 = 1
        _ = v38
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        for fi40 := 0; fi40 < 3; fi40++ {
            v4 -= 3.14159
            fmt.Println(fi40, v2)
            var v41 int16 = -1
            var v42 float64 = (v4 / v4)
            var v43 int8 = int8(v41)
            _ = fi40
            _ = v41
            _ = v42
            _ = v43
        }
        _ = fi39
    }
    _ = len(v31)
    var v44 int = 798
    var v45 float32 = float32(v4)
    v46 := GenS2{F0: (v2 - v2), F1: v34, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v44
    _ = v45
    _ = v46
}
