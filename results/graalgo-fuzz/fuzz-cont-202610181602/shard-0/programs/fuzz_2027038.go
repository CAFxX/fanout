package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8; F1 int16; F2 int8 }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint16, p1 uint) uint8 {
    var v5 uint32 = 2
    v6 := []int{-9223372036854775808, 127}
    if len(v6) > 0 {
        v6[0] = -896
    }
    var v7 float32 = 0.1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 131
}
func genFn1(p0 uint, p1 int8) uint {
    switch 42 {
    case 0:
        if (true && false) {
            for fi8 := 0; fi8 < 9; fi8++ {
                for fi9 := 0; fi9 < 4; fi9++ {
                    if true {
                        var v10 uint = (^(p0))
                        v11 := GenS1{F0: -587, F1: 117, F2: 1}
                        var v12 uint8 = 127
                        _ = v10
                        _ = v11
                        _ = v12
                    } else {
                        p1 = p1
                    }
                    if false {
                        p1 -= (p1 & (-(p1)))
                        var v13 uint = p0
                        _ = v13
                    } else {
                        fmt.Println(fi8, fi8, fi9)
                    }
                    _ = fi9
                }
                _ = fi8
            }
            fmt.Println(p0)
        }
    }
    if true {
        v14 := GenS1{F0: -878, F1: 776, F2: 216}
        _ = v14
    }
    _ = p0
    _ = p1
    return (p0 * p0)
}

func main() {
    var v1 int16 = -54
    var v2 uint32 = 2147483647
    var v3 float64 = 1.5
    v4 := []uint32{(^(v2))}
    fmt.Println(v2, v1)
    var v15 uint = 245
    v1 += -32768
    _ = len(v4)
    var v16 uint = v15
    var v17 float64 = 0.1
    for fi18 := 0; fi18 < 11; fi18++ {
        switch 32768 {
        case -148:
            var v19 uint32 = (v2 & v2)
            _ = v19
        case -9223372036854775808:
            for fi20 := 0; fi20 < 11; fi20++ {
                v15 -= v15
                v2 = (v2 ^ (v2 - v2))
                if (((v16 & (v15 << ((uint(v2)) & 63))) > (v16 / ((v16) | 1))) && false) {
                    switch 4294967295 {
                    case 8:
                        var v21 uint64 = 773
                        _ = v21
                    default:
                        v1 -= v1
                    }
                } else {
                    if len(v4) > 0 {
                        v4[0] = v2
                    }
                    fmt.Println(fi20, v15)
                    var v22 uint8 = 0
                    _ = v22
                }
                _ = fi20
            }
        case 0:
            v1 -= (-(v1))
        default:
            for fi23 := 0; fi23 < 8; fi23++ {
                if false {
                    _ = len(v4)
                    var v24 int16 = (v1 << (((v16 ^ genFn1(genFn1(0, 0), -1))) & 63))
                    switch (fi18 ^ (fi23 ^ (fi23 & (fi18 + 2147483647)))) {
                    case -7:
                        if (334 >= 248) {
                            if (!false) {
                                v25 := GenS2{F0: fi23}
                                if (fi18 > -572) {
                                    v26 := GenS3{F0: false}
                                    switch "" {
                                    case "a":
                                        switch (("" + ("a" + "")) + (((("foo bar" + "a") + "a") + "\n\t") + "\n\t")) {
                                        case "":
                                            var v27 uint16 = 0
                                            _ = v27
                                        }
                                    default:
                                        v1 -= (v1 & v1)
                                    }
                                    _ = v26
                                }
                                v2 = v2
                                _ = v25
                            } else {
                                var v28 float64 = (float64(float32(float64((1.5 / 1.0)))) - (((0.0 * (3.14159 * float64(-2.25))) / 3.14159) + 3.14159))
                                fmt.Println(v28)
                                fmt.Println(fi18, v3, v15)
                                _ = v28
                            }
                            v29 := GenS4{F0: true}
                            _ = v29
                        }
                    default:
                        var v30 float32 = 0.1
                        v31 := GenS2{F0: (-(fi18))}
                        v32 := GenS5{F0: 18, F1: v1, F2: -127}
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    v33 := GenS4{F0: (-73 == 0)}
                    _ = v24
                    _ = v33
                }
                var v34 string = "a"
                _ = fi23
                _ = v34
            }
        }
        v16 = v15
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi18
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        switch "" {
        case "":
            if len(v4) > 0 {
                v4[0] = v2
            }
        default:
            v16 = v15
        }
        if len(v4) > 0 {
            v4[0] = 2
        }
        if (false && ((0.0 * float32(1e10)) >= ((1.5 / (float32(float64(1e10)) - (float32((((0.0 * (1e10 / float64((0.0 * ((float32(v3) - float32(v3)) / 1.5))))) - 0.1) - ((3.14159 - (v3 * v17)) * v3))) / -2.25))) / -2.25))) {
            var v36 uint32 = v2
            if len(v4) > 0 {
                v4[0] = 0
            }
            _ = v36
        }
        _ = fi35
    }
    var v37 string = "a"
    v38 := GenS2{F0: 42}
    for fi39 := 0; fi39 < 6; fi39++ {
        for fi40 := 0; fi40 < 20; fi40++ {
            v3 -= v3
            _ = fi40
        }
        v37 = ((v37 + v37) + v37)
        var v41 float32 = 3.14159
        _ = fi39
        _ = v41
    }
    fmt.Println(v2, v17)
    fmt.Println(v17, v17, v15)
    if false {
        v37 = ""
    }
    v15 -= v15
    v42 := GenS6{F0: -422}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v42
}
