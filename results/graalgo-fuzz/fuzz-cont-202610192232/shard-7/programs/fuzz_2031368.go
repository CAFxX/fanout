package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int64, p1 uint8, p2 bool) int16 {
    p1 += (p1 ^ (p1 << ((18446744073709551615) & 63)))
    p0 += (p0 / ((-826) | 1))
    var v7 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1(p0 uint32, p1 bool, p2 bool) uint32 {
    var v8 float64 = -2.25
    if p1 {
        v9 := []int{-218}
        switch int(p0) {
        case 468:
            if p2 {
                if (!true) {
                    v10 := GenS1{F0: 997, F1: 1}
                    var v11 uint16 = 203
                    _ = v10
                    _ = v11
                } else {
                    var v12 uint64 = 0
                    var v13 int16 = 1
                    var v14 int8 = 35
                    _ = v12
                    _ = v13
                    _ = v14
                }
                p2 = ((334 <= 287) && p1)
            } else {
                v15 := GenS2{F0: 824}
                _ = v15
            }
        case 362:
            p1 = ((42 != -2147483648) && p1)
        case 32768:
            p0 = 65535
        default:
            var v16 int = 991
            _ = v16
        }
        if true {
            if (2147483647 >= -610) {
                _ = len(v9)
                v17 := GenS2{F0: 42}
                v18 := GenS2{F0: 967}
                _ = v17
                _ = v18
            }
        } else {
            if (!p1) {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                p2 = (53 >= uint8(p0))
                p0 = (p0 | p0)
            } else {
                var v19 float64 = (((v8 + v8) - (v8 * ((float64(1e10) - 1.5) / v8))) / 0.1)
                if (32768 == 288) {
                    for fi20 := 0; fi20 < 2; fi20++ {
                        var v21 uint16 = 2
                        p1 = p2
                        for fi22 := 0; fi22 < 10; fi22++ {
                            v23 := GenS3{F0: v21, F1: (^(fi22))}
                            _ = fi22
                            _ = v23
                        }
                        _ = fi20
                        _ = v21
                    }
                }
                p1 = false
                _ = v19
            }
            v8 = 1.5
        }
        _ = v9
    } else {
        v24 := []int{1}
        if p2 {
            fmt.Println(p2, p1, p0)
        } else {
            p1 = (!(2147483647 < -2147483648))
            if false {
                if p1 {
                    switch 576 {
                    case 9223372036854775807:
                        var v25 float32 = float32(v8)
                        _ = v25
                    case 221:
                        for fi26 := 0; fi26 < 16; fi26++ {
                            for fi27 := 0; fi27 < 3; fi27++ {
                                v28 := GenS4{F0: 631, F1: p0}
                                _ = fi27
                                _ = v28
                            }
                            _ = fi26
                        }
                    default:
                        if len(v24) > 0 {
                            v24[0] = -1
                        }
                    }
                    var v29 uint32 = (p0 | 2)
                    _ = v29
                }
                for fi30 := 0; fi30 < 6; fi30++ {
                    p2 = (1 <= 42)
                    _ = fi30
                }
            } else {
                p0 -= (p0 % ((p0) | 1))
                p2 = true
            }
            var v31 int = 2147483647
            _ = v31
        }
        _ = v24
    }
    v32 := GenS3{F0: 0, F1: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    _ = v8
    return 71
}

func main() {
    var v1 int32 = -243
    v2 := []uint32{173, 2147483647}
    var v3 uint8 = 127
    var v4 int = 32768
    v5 := make(map[uint32]int)
    var v6 int32 = v1
    if (true || false) {
        v33 := GenS1{F0: 1, F1: 864}
        for fi34 := 0; fi34 < 6; fi34++ {
            fmt.Println(v6, v1)
            _ = fi34
        }
        fmt.Println(v4)
        _ = v33
    }
    _ = len(v5)
    for fi35 := 0; fi35 < 3; fi35++ {
        var v36 float64 = (1e10 - ((float64(3.14159) / (0.0 - float64(float32(-2.25)))) + (-2.25 * ((0.0 / (((1e10 * 0.1) + (3.14159 * 0.0)) / 0.1)) - 3.14159))))
        for fi37 := 0; fi37 < 19; fi37++ {
            if (-9223372036854775808 <= 0) {
                var v38 uint64 = 0
                v39 := GenS5{F0: false}
                v40 := GenS4{F0: 2, F1: 109}
                switch (fi37 << ((875) & 63)) {
                case 0:
                    var v41 int64 = 994
                    _ = v41
                case -712:
                    v4 = fi37
                case 127:
                    v42 := GenS6{F0: 2147483647, F1: true}
                    _ = v42
                default:
                    var v43 uint16 = 622
                    _ = v43
                }
                v36 += v36
                _ = v38
                _ = v39
                _ = v40
            }
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    var v44 string = "foo bar"
    v45 := GenS7{F0: 491, F1: 0}
    fmt.Println(v6, v6)
    if false {
        var v46 int = v4
        fmt.Println(v6, v4)
        _ = v46
    } else {
        var v47 float64 = float64(-2.25)
        _ = v47
    }
    if (32767 <= 671) {
        var v48 int32 = (v1 ^ (v1 / (((v6 ^ int32(v4))) | 1)))
        v49 := GenS2{F0: 0}
        _ = v48
        _ = v49
    } else {
        var v50 bool = true
        _ = v50
    }
    delete(v5, genFn1(863, false, (95 <= 1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
