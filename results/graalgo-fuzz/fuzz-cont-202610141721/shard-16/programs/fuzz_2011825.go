package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 bool }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int8, p1 string, p2 uint) string {
    p1 = "hello"
    v6 := []int{128, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p1 + "foo bar")
}
func genFn1(p0 uint, p1 bool) uint {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: -428, F1: 9223372036854775807, F2: p1}
        fmt.Println(p0, p0)
        _ = fi7
        _ = v8
    }
    p1 = ("\n\t" <= ("" + "hello"))
    v9 := []int{-2147483648, 2147483647, 9223372036854775807, -9223372036854775808}
    if p1 {
        v10 := GenS1{F0: 2147483647, F1: 387, F2: (2 > 109)}
        var v11 bool = p1
        if ((!p1) || (!false)) {
            switch -1 {
            case 359:
                var v12 float32 = 0.0
                _ = v12
            case -2147483648:
                p1 = v11
            default:
                if false {
                    v11 = true
                    for fi13 := 0; fi13 < 12; fi13++ {
                        var v14 uint8 = 255
                        p0 = (p0 >> ((p0) & 63))
                        _ = fi13
                        _ = v14
                    }
                    if true {
                        var v15 int8 = 0
                        var v16 float64 = (0.1 - float64(((((1.5 * 1.5) + 1.5) * (float32(1e10) - ((1.5 / (float32(1.5) * (1e10 - (float32(3.14159) + (0.0 * float32(((float64(3.14159) / (1e10 / 1e10)) * -2.25))))))) * 0.0))) / 1e10)))
                        fmt.Println(v15, v15, v15)
                        _ = v15
                        _ = v16
                    } else {
                        var v17 bool = (3.14159 >= 0.0)
                        _ = v17
                    }
                }
            }
        } else {
            if true {
                for fi18 := 0; fi18 < 6; fi18++ {
                    if v11 {
                        if false {
                            fmt.Println(v11, fi18, p1)
                            fmt.Println(p1, p1)
                            if (!(42 != 39)) {
                                v11 = true
                                for fi19 := 0; fi19 < 16; fi19++ {
                                    if ((0.0 - (1.5 - 1.5)) < 3.14159) {
                                        if ((753 <= 65535) || false) {
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                            p0 = (p0 - (^(p0)))
                                            if (((p1 && p1) && v11) && v11) {
                                                var v20 int8 = 127
                                                var v21 int64 = 42
                                                var v22 uint16 = 2
                                                _ = v20
                                                _ = v21
                                                _ = v22
                                            }
                                        } else {
                                            p0 = 18446744073709551615
                                            v9 = append(v9, fi18)
                                        }
                                        var v23 int = (fi19 * fi19)
                                        var v24 uint8 = 105
                                        _ = v23
                                        _ = v24
                                    } else {
                                        v11 = p1
                                    }
                                    fmt.Println(p1, fi19, p0)
                                    _ = fi19
                                }
                                p0 += p0
                            }
                        }
                        var v25 string = (("" + "") + "foo bar")
                        _ = v25
                    }
                    _ = fi18
                }
                var v26 int64 = 423
                _ = v26
            }
            p1 = p1
            var v27 bool = false
            _ = v27
        }
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return (^(p0))
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{-1, 127, -99, -2147483648}
    var v3 int64 = -9223372036854775808
    var v4 float64 = 0.0
    var v5 uint16 = 1
    v28 := GenS1{F0: 72, F1: 0, F2: false}
    v29 := GenS1{F0: -301, F1: v3, F2: true}
    if (127 != 42) {
        var v30 uint32 = 562
        _ = v30
    }
    v5 += v5
    for fi31 := 0; fi31 < 12; fi31++ {
        for fi32 := 0; fi32 < 14; fi32++ {
            v5 = v5
            switch -381 {
            case 4294967295:
                var v33 uint64 = 42
                _ = v33
            case -289:
                var v34 int8 = 127
                _ = v34
            default:
                fmt.Println(v3)
            }
            fmt.Println(fi31)
            v1[1] = fi31
            _ = fi32
        }
        _ = fi31
    }
    var v35 uint32 = 475
    v5 += v5
    var v36 int = 32768
    for fi37 := 0; fi37 < 18; fi37++ {
        v35 -= v35
        _ = fi37
    }
    for fi38 := 0; fi38 < 6; fi38++ {
        var v39 uint8 = 128
        if false {
            v40 := GenS2{F0: (v5 * v5), F1: (v35 >> ((genFn1(255, (v5 <= v5))) & 63))}
            var v41 uint = 4294967295
            _ = v40
            _ = v41
        }
        _ = fi38
        _ = v39
    }
    for fi42 := 0; fi42 < 19; fi42++ {
        v5 += (v5 << ((2) & 63))
        fmt.Println(v5)
        _ = fi42
    }
    var v43 float32 = 1.5
    for fi44 := 0; fi44 < 13; fi44++ {
        fmt.Println(v5, fi44, v4)
        v3 -= v3
        _ = fi44
    }
    v43 += 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v43
    _ = v5
}
