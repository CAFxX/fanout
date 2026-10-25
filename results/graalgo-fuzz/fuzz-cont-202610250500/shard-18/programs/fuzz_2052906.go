package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 bool }
type GenS2 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint }
type GenS4 struct { F0 int64; F1 int64 }
func genFn0(p0 string, p1 uint8) uint32 {
    if (false && false) {
        v5 := []int{-107, 1, 32768, -2147483648}
        _ = v5
    }
    v6 := GenS1{F0: -419, F1: (p1 ^ 0), F2: (((1e10 / 0.1) / -2.25) >= ((3.14159 / -2.25) * 0.1))}
    v7 := GenS2{F0: true, F1: 42, F2: (p1 & (p1 >> ((606) & 63)))}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int32 = 1
        v10 := GenS2{F0: (2 == 4294967295), F1: 42, F2: p1}
        v11 := GenS2{F0: true, F1: 2, F2: (p1 & p1)}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2147483648
}
func genFn1() uint16 {
    var v12 float64 = float64(1e10)
    v12 = 1e10
    v13 := []int{4294967295, 42, 128}
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 uint = 2
    var v2 uint64 = 1
    var v3 int32 = -933
    var v4 uint16 = 42
    v14 := []int{1}
    var v15 uint16 = 102
    v16 := GenS3{F0: -128, F1: v2, F2: 1}
    var v17 uint32 = 2
    fmt.Println(v4, v2, v4)
    v15 += uint16(v1)
    v18 := GenS3{F0: 42, F1: (v2 / (((^(v2))) | 1)), F2: v1}
    var v19 int8 = 100
    for fi20 := 0; fi20 < 1; fi20++ {
        if true {
            fmt.Println(v15, fi20)
            if (!false) {
                var v21 string = "hello"
                _ = v21
            }
            v15 += v4
        } else {
            if true {
                v4 = (v15 - uint16(v2))
                if ((v15 / ((v15) | 1)) != (v15 & genFn1())) {
                    var v22 uint8 = 0
                    _ = v22
                }
            } else {
                v1 = 65535
            }
            if (-1 != 9223372036854775807) {
                for fi23 := 0; fi23 < 19; fi23++ {
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    if false {
                        v4 += (v4 + 42)
                        v24 := GenS3{F0: 42, F1: (^(v2)), F2: (v1 + v1)}
                        var v25 float64 = (1.5 / (1e10 + 0.0))
                        var v26 uint8 = uint8(v1)
                        var v27 int8 = -128
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = v24
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        v4 += (v15 + v4)
                        var v28 bool = (v3 >= v3)
                        v19 = 1
                        _ = v28
                    }
                    _ = fi23
                }
                for fi29 := 0; fi29 < 9; fi29++ {
                    _ = len(v14)
                    _ = fi29
                }
                if ((v17 ^ 0) >= genFn0("hello", 255)) {
                    if len(v14) > 0 {
                        v14[0] = 595
                    }
                }
            }
        }
        if (669 > 969) {
            switch (fi20 / ((fi20) | 1)) {
            case 9223372036854775807:
                fmt.Println(fi20)
            default:
                for fi30 := 0; fi30 < 5; fi30++ {
                    var v31 uint = (^(v1))
                    var v32 bool = false
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
            if (fi20 < int(v17)) {
                var v33 uint8 = 255
                if ((v33 - v33) <= (^(v33))) {
                    v17 = v17
                } else {
                    v14 = append(v14, 0)
                }
                _ = v33
            }
        } else {
            fmt.Println(v4, v4)
            v3 += 187
            if (!false) {
                v19 -= (v19 * v19)
            } else {
                if true {
                    var v34 uint64 = (v2 | (v2 & 2))
                    _ = len(v14)
                    if false {
                        v15 = genFn1()
                        v35 := GenS4{F0: -9223372036854775808, F1: -1}
                        _ = v35
                    }
                    _ = v34
                }
                var v36 bool = (uint32(v2) == (v17 + (v17 & v17)))
                for fi37 := 0; fi37 < 18; fi37++ {
                    if v36 {
                        v14 = append(v14, (fi37 << ((v1) & 63)))
                        var v38 int16 = 0
                        _ = v38
                    }
                    v36 = v36
                    _ = fi37
                }
                _ = v36
            }
        }
        v3 += v3
        _ = fi20
    }
    var v39 bool = true
    _ = len(v14)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v4
}
