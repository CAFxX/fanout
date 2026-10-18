package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint }
func genFn0(p0 int32, p1 int, p2 int64) bool {
    if ((1e10 - 1e10) > -2.25) {
        v5 := []int{(p1 | 248), -1}
        p2 = (p2 ^ p2)
        _ = v5
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            if ((1.5 < (1.5 + ((1e10 * (1.5 * (-2.25 / float32(0.1)))) / (1e10 + (float32((1e10 + 1e10)) - 3.14159))))) || (14 < 1)) {
                var v8 float32 = 0.1
                p2 = (p2 % (((p2 - 1)) | 1))
                v9 := GenS1{F0: p2, F1: -1}
                _ = v8
                _ = v9
            }
            fmt.Println(p2, p0, p1)
            var v10 int64 = p2
            _ = fi7
            _ = v10
        }
        _ = fi6
    }
    fmt.Println(p0, p1, p1)
    p0 = -1
    _ = p0
    _ = p1
    _ = p2
    return false
}
func genFn1() bool {
    var v11 float32 = float32((-2.25 / 1.0))
    v12 := GenS1{F0: 0, F1: -128}
    _ = v11
    _ = v12
    return false
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]uint32)
    var v3 bool = false
    var v4 uint64 = 1
    v3 = (float64(1.5) > (float64((1e10 - (float32(v1) - 1.5))) - (0.1 - -2.25)))
    v1 = v1
    v4 += (^(v4))
    if v3 {
        if true {
            if (true || v3) {
                var v13 uint8 = 9
                var v14 float64 = v1
                var v15 int32 = -2147483648
                _ = v13
                _ = v14
                _ = v15
            } else {
                v16 := GenS2{F0: 127, F1: 0, F2: 775}
                v17 := GenS3{F0: 0}
                if false {
                    if (v4 >= 330) {
                        v18 := []int{661, 32768, -9223372036854775808}
                        _ = v18
                    } else {
                        var v19 uint8 = 177
                        for fi20 := 0; fi20 < 9; fi20++ {
                            v21 := GenS1{F0: -214, F1: -128}
                            _ = fi20
                            _ = v21
                        }
                        _ = v19
                    }
                } else {
                    switch "" {
                    case "a":
                        var v22 uint64 = v4
                        _ = v22
                    case "":
                        v23 := v2[-50]
                        _ = v23
                    default:
                        v24 := []int{32768}
                        _ = v24
                    }
                    var v25 float64 = v1
                    for fi26 := 0; fi26 < 9; fi26++ {
                        var v27 int64 = 1
                        if v3 {
                            var v28 uint8 = 160
                            _ = v28
                        } else {
                            var v29 int64 = v27
                            var v30 int32 = 0
                            _ = v29
                            _ = v30
                        }
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                }
                _ = v16
                _ = v17
            }
            switch int(v4) {
            case 32768:
                for fi31 := 0; fi31 < 1; fi31++ {
                    var v32 bool = v3
                    var v33 int32 = 1
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            case 32767:
                v1 = v1
            case -9223372036854775808:
                v34 := GenS4{F0: (340 < 0)}
                _ = v34
            default:
                fmt.Println(v3, v4, v3)
            }
        }
    } else {
        v1 -= (((v1 * float64(1.5)) - v1) * (((v1 / (v1 / (v1 + v1))) - 0.1) - v1))
        var v35 uint8 = 20
        v36 := GenS5{F0: (^(v35))}
        _ = v35
        _ = v36
    }
    v37 := []int{-1, 792, 2147483647, -928}
    fmt.Println(v4, v3)
    fmt.Println(v1)
    for fi38 := 0; fi38 < 8; fi38++ {
        for fi39 := 0; fi39 < 7; fi39++ {
            v40 := GenS6{F0: 989}
            fmt.Println(v4, v3, fi39)
            v41 := GenS1{F0: -49, F1: -1}
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = fi38
    }
    var v42 uint32 = 4294967295
    if (true && (!false)) {
        var v43 bool = ((-1 >= 1000) || (!true))
        _ = v43
    }
    if v3 {
        delete(v2, int(v4))
        var v44 uint64 = (v4 ^ 2)
        _ = v44
    }
    delete(v2, 506)
    var v45 uint = 42
    switch 875 {
    case 2147483647:
        var v46 uint64 = v4
        _ = v46
    default:
        fmt.Println(v4, v45, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v42
    _ = v45
}
