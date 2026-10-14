package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint) uint8 {
    p0 += (p0 >> ((p0) & 63))
    if true {
        v7 := GenS1{F0: 127, F1: -100}
        v8 := []int{-2147483648}
        _ = v7
        _ = v8
    } else {
        var v9 bool = true
        p0 = p0
        p0 += 4294967295
        _ = v9
    }
    _ = p0
    return 243
}
func genFn1(p0 int32) int16 {
    var v10 float32 = (((((0.1 * (0.0 * 3.14159)) + -2.25) + (((0.0 + 3.14159) * 0.1) / (-2.25 * float32(1.5)))) - -2.25) * 1.5)
    v11 := GenS1{F0: 2, F1: -128}
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int = (fi12 ^ (fi12 / (((fi12 ^ fi12)) | 1)))
        _ = fi12
        _ = v13
    }
    v14 := []int{128, -9223372036854775808, 240}
    _ = p0
    _ = v10
    _ = v11
    _ = v14
    return -1
}

func main() {
    var v1 bool = false
    v2 := make(map[int]string)
    var v3 uint16 = 42
    var v4 uint32 = 1
    var v5 uint32 = v4
    v6 := make(map[uint32]int)
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := GenS1{F0: 1, F1: 1}
            v18 := GenS2{F0: v4, F1: 177, F2: (v5 ^ 0)}
            v5 -= (v5 / (((v5 << ((199) & 63))) | 1))
            _ = fi16
            _ = v17
            _ = v18
        }
        v19 := GenS1{F0: 211, F1: -93}
        var v20 int = fi15
        _ = fi15
        _ = v19
        _ = v20
    }
    v6[v5] = 127
    if v1 {
        v21 := GenS2{F0: (v4 * v5), F1: genFn0(255), F2: 2147483647}
        if true {
            if (-9223372036854775808 != 42) {
                var v22 float32 = 3.14159
                _ = v22
            } else {
                switch -436 {
                case 32768:
                    fmt.Println(v1, v5)
                case -396:
                    if v1 {
                        v23 := GenS2{F0: v4, F1: 0, F2: v4}
                        fmt.Println(v4)
                        _ = v23
                    } else {
                        v5 = 2147483648
                        var v24 float32 = -2.25
                        _ = v24
                    }
                default:
                    if (!true) {
                        v25 := GenS3{F0: -9223372036854775808}
                        _ = v25
                    }
                }
                for fi26 := 0; fi26 < 1; fi26++ {
                    v27 := []int{fi26, (fi26 - fi26), fi26}
                    fmt.Println(v5, v5)
                    _ = fi26
                    _ = v27
                }
                var v28 bool = (v1 || (!(1 >= -825)))
                _ = v28
            }
            v1 = v1
        } else {
            var v29 int16 = genFn1(2147483647)
            for fi30 := 0; fi30 < 15; fi30++ {
                var v31 uint64 = 4294967295
                v29 = (v29 | (v29 - v29))
                delete(v6, v4)
                _ = fi30
                _ = v31
            }
            for fi32 := 0; fi32 < 18; fi32++ {
                var v33 uint64 = 943
                v33 = 2
                fmt.Println(v3, v33)
                _ = fi32
                _ = v33
            }
            _ = v29
        }
        _ = v21
    }
    if (-489 < 4294967295) {
        if v1 {
            switch 1 {
            case 42:
                if v1 {
                    v5 = v5
                } else {
                    v34 := []int{9223372036854775807, int(v3)}
                    _ = v34
                }
            case 127:
                for fi35 := 0; fi35 < 16; fi35++ {
                    v3 = v3
                    var v36 float64 = 1e10
                    var v37 int = (^(fi35))
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            default:
                fmt.Println(v3, v5, v3)
            }
        }
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        var v39 uint64 = 0
        v1 = v1
        if (!true) {
            var v40 float32 = 1.5
            v2[fi38] = "foo bar"
            for fi41 := 0; fi41 < 3; fi41++ {
                v42 := GenS1{F0: 255, F1: -128}
                _ = fi41
                _ = v42
            }
            _ = v40
        } else {
            var v43 int16 = 1000
            for fi44 := 0; fi44 < 1; fi44++ {
                fmt.Println(v4)
                var v45 float32 = (-2.25 / 1.5)
                _ = fi44
                _ = v45
            }
            fmt.Println(fi38, v4, v39)
            _ = v43
        }
        _ = fi38
        _ = v39
    }
    var v46 int16 = genFn1(-1)
    v2[42] = (("a" + ("" + "hello")) + "a")
    for fi47 := 0; fi47 < 5; fi47++ {
        var v48 uint64 = 1
        _ = fi47
        _ = v48
    }
    for fi49 := 0; fi49 < 16; fi49++ {
        var v50 int8 = -1
        for fi51 := 0; fi51 < 15; fi51++ {
            var v52 uint32 = v4
            var v53 int16 = (v46 >> ((4294967295) & 63))
            v54 := v2[fi49]
            _ = fi51
            _ = v52
            _ = v53
            _ = v54
        }
        var v55 float64 = ((-2.25 * 3.14159) + -2.25)
        _ = fi49
        _ = v50
        _ = v55
    }
    v56 := []int{}
    v57 := GenS1{F0: 2, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v56
    _ = v57
    _ = v6
}
