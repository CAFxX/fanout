package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int16; F1 int8; F2 uint64 }
type GenS4 struct { F0 int32; F1 int16; F2 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint64, p1 uint8) int {
    if true {
        v7 := GenS1{F0: 1, F1: -23}
        p1 += p1
        _ = v7
    } else {
        p1 -= p1
        p0 -= p0
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := []int{fi8, (fi8 >> ((624) & 63)), fi8, fi8}
            _ = fi8
            _ = v9
        }
    }
    p1 = (p1 + p1)
    var v10 float32 = 1.5
    v11 := []int{}
    var v12 uint8 = (p1 << ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return -1
}
func genFn1(p0 uint16, p1 int, p2 uint16) uint {
    if (42 > 122) {
        v13 := GenS1{F0: 65535, F1: -66}
        _ = v13
    } else {
        p1 = p1
    }
    fmt.Println(p0, p1, p0)
    switch "hello" {
    case "a":
        if true {
            var v14 float64 = ((((0.1 - (((0.0 / 1e10) * (((-2.25 * 1e10) / -2.25) * 3.14159)) + float64(3.14159))) + (-2.25 / 1.0)) / 3.14159) + (3.14159 / float64(-2.25)))
            _ = v14
        }
    default:
        v15 := GenS1{F0: 2147483648, F1: 533}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    return 387
}

func main() {
    v1 := []uint32{}
    var v2 float64 = float64((1.5 * 0.0))
    var v3 int = 9223372036854775807
    v4 := make(map[int]bool)
    var v5 uint = 220
    var v6 int64 = 441
    var v16 int16 = 1
    if ("" > "foo bar") {
        var v17 uint8 = 128
        _ = v17
    } else {
        v18 := GenS1{F0: 1, F1: 0}
        _ = v18
    }
    if false {
        _ = len(v4)
        var v19 int32 = 2147483647
        var v20 uint = (v5 + 1)
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: -47, F1: (v5 & 255)}
    fmt.Println(v3, v3)
    if (!(2 >= 4294967295)) {
        var v22 int = v3
        _ = v22
    }
    var v23 int8 = 42
    if (!((463 == 2) || (2147483648 >= 65535))) {
        v24 := GenS1{F0: 1, F1: 42}
        _ = v24
    } else {
        if false {
            if (-51 <= -1) {
                v25 := GenS3{F0: -1, F1: (v23 | 42), F2: 1}
                var v26 float32 = float32((v2 / (1e10 * (v2 - v2))))
                var v27 int8 = (v23 + -57)
                _ = v25
                _ = v26
                _ = v27
            } else {
                fmt.Println(v2, v16, v6)
                v28 := GenS4{F0: -2147483648, F1: (v16 - (v16 % (((v16 - int16(v3))) | 1))), F2: -9223372036854775808}
                var v29 uint = (v5 - v5)
                _ = v28
                _ = v29
            }
            var v30 bool = (168 >= 2)
            _ = v30
        } else {
            v6 = (v6 * 42)
        }
        v31 := GenS1{F0: 4294967295, F1: 878}
        var v32 uint8 = 255
        _ = v31
        _ = v32
    }
    v2 = 0.0
    v33 := GenS5{F0: 128}
    v6 -= -430
    switch "foo bar" {
    case "\n\t":
        v4[(-(v3))] = (2147483647 < 727)
    case "":
        fmt.Println(v16)
    case "foo bar":
        switch genFn0(1, 128) {
        case 32768:
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        }
    default:
        v34 := GenS6{F0: 1, F1: (v16 << ((v5) & 63))}
        _ = v34
    }
    for fi35 := 0; fi35 < 15; fi35++ {
        if true {
            if (401 <= 843) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = len(v1)
                var v36 int64 = v6
                _ = v36
            } else {
                v16 = v16
                var v37 bool = false
                if v37 {
                    v38 := GenS2{F0: v23, F1: (v5 ^ v5)}
                    _ = v38
                } else {
                    if v37 {
                        var v39 uint64 = 42
                        v40 := v4[genFn0(v39, 42)]
                        _ = v39
                        _ = v40
                    } else {
                        if v37 {
                            _ = len(v1)
                        }
                        v5 += genFn1(32767, (^(fi35)), 42)
                        switch fi35 {
                        case 128:
                            var v41 uint8 = 128
                            _ = v41
                        case -156:
                            if (!((-1 >= 2147483647) || true)) {
                                fmt.Println(v5, fi35)
                            }
                        case 42:
                            for fi42 := 0; fi42 < 6; fi42++ {
                                v3 = genFn0(982, 0)
                                v23 = (v23 + v23)
                                _ = fi42
                            }
                        default:
                            v16 += 42
                        }
                    }
                    for fi43 := 0; fi43 < 4; fi43++ {
                        var v44 int8 = v23
                        _ = len(v4)
                        _ = fi43
                        _ = v44
                    }
                    if v37 {
                        if v37 {
                            _ = len(v1)
                        }
                        var v45 uint32 = 2147483647
                        for fi46 := 0; fi46 < 12; fi46++ {
                            v47 := GenS5{F0: 104}
                            var v48 int64 = (v6 * v6)
                            var v49 int8 = 0
                            _ = fi46
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                        _ = v45
                    }
                }
                _ = v37
            }
            v50 := v4[(v3 >> ((v5) & 63))]
            for fi51 := 0; fi51 < 10; fi51++ {
                var v52 uint8 = 128
                _ = fi51
                _ = v52
            }
            _ = v50
        }
        v23 = int8(v16)
        _ = fi35
    }
    fmt.Println(v2)
    for fi53 := 0; fi53 < 10; fi53++ {
        v6 = v6
        v3 += fi53
        _ = fi53
    }
    v2 = (v2 - (1e10 / -2.25))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
