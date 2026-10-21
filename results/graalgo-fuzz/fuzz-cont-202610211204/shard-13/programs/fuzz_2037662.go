package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 string }
func genFn0(p0 uint64) int8 {
    var v7 bool = true
    p0 += 4294967295
    _ = p0
    _ = v7
    return -128
}
func genFn1(p0 uint64, p1 int8, p2 uint32) uint32 {
    p0 = (^(p0))
    p0 = uint64(p2)
    _ = p0
    _ = p1
    _ = p2
    return (^(p2))
}

func main() {
    var v1 int64 = 232
    var v2 float64 = 0.1
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    var v5 uint64 = 1
    var v6 uint16 = 0
    if (false && false) {
        var v8 uint8 = 255
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            if true {
                var v9 uint = 42
                fmt.Println(v5)
                v10 := GenS1{F0: 2, F1: (v8 & v8), F2: -1}
                _ = v10
                _ = v9
            }
        }
        _ = v8
    }
    switch 124 {
    case 81:
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint32 = 2147483647
            fmt.Println(v5)
            v13 := GenS2{F0: -460}
            _ = fi11
            _ = v12
            _ = v13
        }
    default:
        fmt.Println(v2)
    }
    if (100 >= 127) {
        delete(v4, 128)
    } else {
        fmt.Println(v5, v5)
        if (false && false) {
            switch 245 {
            case 32767:
                var v14 int8 = 115
                _ = v14
            case -9223372036854775808:
                for fi15 := 0; fi15 < 19; fi15++ {
                    if ((^(v5)) != (v5 / ((950) | 1))) {
                        v5 = v5
                        fmt.Println(fi15)
                    }
                    v1 += v1
                    _ = fi15
                }
            case 2147483647:
                if (!true) {
                    var v16 float64 = (0.0 * v2)
                    if false {
                        var v17 uint64 = 340
                        v2 = v2
                        fmt.Println(v2, v17, v2)
                        _ = v17
                    }
                    if (!(1 > -990)) {
                        v6 = v6
                        if (4294967295 < -100) {
                            _ = len(v3)
                        }
                    }
                    _ = v16
                } else {
                    var v18 uint = 42
                    _ = v18
                }
            default:
                var v19 int = -2147483648
                _ = v19
            }
            fmt.Println(v5, v6)
        }
        v1 = (v1 + (^(v1)))
    }
    if (true || (v6 <= (^(v6)))) {
        var v20 uint16 = v6
        var v21 float32 = ((1.5 / 1.0) - 0.1)
        var v22 int8 = 127
        _ = v20
        _ = v21
        _ = v22
    } else {
        var v23 uint32 = uint32(v5)
        _ = v23
    }
    switch -431 {
    case 1:
        fmt.Println(v2, v1, v2)
    default:
        var v24 string = "hello"
        _ = v24
    }
    switch "\n\t" {
    case "hello":
        if true {
            for fi25 := 0; fi25 < 9; fi25++ {
                for fi26 := 0; fi26 < 11; fi26++ {
                    v27 := GenS2{F0: 32767}
                    v1 = -9223372036854775808
                    if true {
                        if (!false) {
                            v1 = 9223372036854775807
                            v28 := GenS3{F0: 32767, F1: (v1 & (v1 >> ((1) & 63)))}
                            fmt.Println(v2)
                            _ = v28
                        }
                    }
                    _ = fi26
                    _ = v27
                }
                var v29 string = ("hello" + "a")
                _ = fi25
                _ = v29
            }
            if (true || ((v6 ^ 2) < (^(v6)))) {
                var v30 uint64 = (v5 >> ((18446744073709551615) & 63))
                var v31 int8 = 2
                _ = v30
                _ = v31
            } else {
                v1 -= v1
                for fi32 := 0; fi32 < 6; fi32++ {
                    var v33 int32 = 0
                    _ = fi32
                    _ = v33
                }
            }
            v5 = v5
        }
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        v2 -= v2
        switch (fi34 / ((-9223372036854775808) | 1)) {
        case -9223372036854775808:
            var v35 uint32 = genFn1(v5, genFn0(v5), 308)
            _ = v35
        case 128:
            if (int32(v5) > 42) {
                var v36 int16 = 590
                if true {
                    v37 := v4[9223372036854775807]
                    if (false && (1 <= 4294967295)) {
                        var v38 uint32 = 727
                        if ((v37 && false) && v37) {
                            v1 = v1
                            v36 = -1
                        }
                        _ = v38
                    } else {
                        v6 = (^(v6))
                        var v39 int = (fi34 % ((fi34) | 1))
                        fmt.Println(fi34, v2, v6)
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 6; fi40++ {
                        if (115 != 681) {
                            v41 := GenS3{F0: v36, F1: v1}
                            var v42 float64 = v2
                            if ((v1 + (v1 - v1)) != v1) {
                                v36 += (^(v36))
                            }
                            _ = v41
                            _ = v42
                        } else {
                            var v43 int8 = 1
                            _ = v43
                        }
                        v44 := GenS1{F0: 851, F1: 128, F2: -1}
                        v45 := GenS4{F0: (fi34 + (fi40 + (fi40 >> ((263) & 63)))), F1: (-(v1))}
                        _ = fi40
                        _ = v44
                        _ = v45
                    }
                    _ = v37
                }
                _ = v36
            }
        }
        for fi46 := 0; fi46 < 15; fi46++ {
            v5 = (v5 % (((v5 >> ((306) & 63))) | 1))
            switch int(v5) {
            case -227:
                if false {
                    v47 := GenS2{F0: 0}
                    _ = v47
                } else {
                    v48 := GenS3{F0: 32767, F1: v1}
                    var v49 uint64 = 965
                    v6 = 32767
                    _ = v48
                    _ = v49
                }
            case 4294967295:
                v6 += v6
            case 32768:
                v1 += (^(v1))
            default:
                v50 := GenS5{F0: -32768, F1: ""}
                fmt.Println(fi34)
                switch "foo bar" {
                case "\n\t":
                    var v51 uint64 = (v5 | (v5 ^ v5))
                    _ = v51
                case "foo bar":
                    v52 := []int{(fi46 / (((fi34 ^ (fi46 * fi34))) | 1)), fi46, 0}
                    _ = v52
                }
                _ = v50
            }
            _ = fi46
        }
        _ = fi34
    }
    if ((v1 <= 858) && (1.5 != -2.25)) {
        if false {
            for fi53 := 0; fi53 < 15; fi53++ {
                v54 := GenS2{F0: -715}
                _ = fi53
                _ = v54
            }
        } else {
            if true {
                for fi55 := 0; fi55 < 10; fi55++ {
                    var v56 int8 = 127
                    _ = fi55
                    _ = v56
                }
                v5 = (^(v5))
            }
        }
        var v57 uint64 = 2
        if ((("foo bar" + "\n\t") + "") >= (("" + "\n\t") + "hello")) {
            var v58 int8 = 70
            switch 32767 {
            case 127:
                if true {
                    var v59 int = 0
                    var v60 uint = 1
                    _ = v59
                    _ = v60
                } else {
                    v61 := GenS6{F0: "foo bar"}
                    var v62 uint32 = 42
                    v62 -= (v62 >> ((461) & 63))
                    _ = v61
                    _ = v62
                }
            case -967:
                var v63 uint = 2
                _ = v63
            case 0:
                v6 += (v6 / ((42) | 1))
            }
            var v64 uint16 = (v6 | v6)
            _ = v58
            _ = v64
        }
        _ = v57
    } else {
        delete(v4, 426)
    }
    v65 := GenS5{F0: 32767, F1: "a"}
    if true {
        v66 := []int{32767, 32767, -2147483648, int(v1)}
        _ = v66
    }
    v6 = (^(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
}
