package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 bool; F2 string }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint; F1 int64; F2 uint32 }
type GenS8 struct { F0 int64; F1 int32 }
func genFn0(p0 string, p1 int32, p2 uint32) uint16 {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint16 = 2
        p1 -= (p1 * p1)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        if (((true && false) || (197 != 2)) || true) {
            p1 += (p1 * (-(p1)))
        }
        var v10 float64 = ((1e10 - 0.1) + 1.5)
        fmt.Println(p1, fi9, fi9)
        _ = fi9
        _ = v10
    }
    v11 := []int{9223372036854775807, 0, 42, 262}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 1
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]int)
    var v3 string = ("" + "a")
    var v4 int64 = -1
    v5 := []int{-9223372036854775808, -83, 247}
    v6 := make(map[uint32]uint32)
    v12 := GenS1{F0: true}
    if (-1 != 42) {
        if false {
            fmt.Println(v3, v3, v4)
            var v13 uint16 = genFn0((v3 + v3), -1, 42)
            var v14 int16 = 1000
            _ = v13
            _ = v14
        } else {
            var v15 float32 = 3.14159
            _ = v15
        }
        v16 := GenS1{F0: true}
        _ = v16
    } else {
        v4 += 966
    }
    var v17 uint32 = 4294967295
    fmt.Println(v1)
    if false {
        var v18 float64 = ((0.1 * (((float64(((0.0 * (float32(float64(v1)) / (float32(float64((((v1 + (3.14159 + v1)) / (0.1 * (1.5 + 3.14159))) / v1))) * v1))) * float32(-2.25))) * -2.25) / 1.0) * (0.0 / 1e10))) * 0.1)
        _ = v18
    }
    var v19 int8 = 42
    if (299 == 799) {
        var v20 int64 = 1
        switch 676 {
        case 0:
            fmt.Println(v17)
        case -133:
            v6[v17] = v17
        case -2147483648:
            switch v3 {
            case "foo bar":
                if ((18446744073709551615 > 1) || false) {
                    if ((v3 >= (v3 + v3)) && (1e10 <= 0.0)) {
                        switch (v3 + "foo bar") {
                        case "hello":
                            _ = len(v5)
                        default:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        v1 -= v1
                        v1 -= (v1 - ((float32(-2.25) - 3.14159) / (float32(-2.25) + v1)))
                    }
                    for fi21 := 0; fi21 < 7; fi21++ {
                        switch "a" {
                        case "":
                            v1 = v1
                        case "a":
                            v3 = "hello"
                        case "\n\t":
                            v22 := GenS1{F0: (float64(1e10) == -2.25)}
                            _ = v22
                        default:
                            var v23 float32 = ((v1 + v1) - float32(1.5))
                            _ = v23
                        }
                        v24 := GenS1{F0: (191 == 255)}
                        _ = fi21
                        _ = v24
                    }
                } else {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        for fi26 := 0; fi26 < 17; fi26++ {
                            if false {
                                if false {
                                    var v27 float32 = v1
                                    _ = v27
                                }
                                var v28 int8 = (v19 >> ((987) & 63))
                                v29 := GenS2{F0: 0, F1: ((v17 << ((2) & 63)) >= (v17 >> ((65535) & 63))), F2: "hello"}
                                switch "hello" {
                                case "foo bar":
                                    v30 := GenS3{F0: true, F1: genFn0(((((("hello" + v3) + v3) + ("hello" + "")) + v3) + (("" + v3) + (v3 + v3))), 2147483647, 65535)}
                                    _ = v30
                                case "\n\t":
                                    if ((0.0 * 0.1) != 0.1) {
                                        v5 = append(v5, fi25)
                                    } else {
                                        var v31 uint64 = 18446744073709551615
                                        _ = v31
                                    }
                                }
                                _ = v28
                                _ = v29
                            }
                            fmt.Println(fi25)
                            _ = fi26
                        }
                        v4 = v20
                        _ = fi25
                    }
                    for fi32 := 0; fi32 < 13; fi32++ {
                        fmt.Println(v17, v17, v4)
                        var v33 string = (v3 + "")
                        _ = fi32
                        _ = v33
                    }
                }
            case "\n\t":
                var v34 int = 4294967295
                _ = v34
            default:
                v3 = (v3 + v3)
            }
        }
        if (v17 < 65535) {
            if false {
                if true {
                    v3 = ""
                    var v35 bool = true
                    _ = v35
                } else {
                    _ = len(v6)
                    v1 += (v1 * (1e10 / 0.1))
                }
                v36 := GenS4{F0: v3}
                _ = v36
            } else {
                v3 = v3
                if (2 != 2) {
                    v37 := GenS5{F0: 0, F1: -10}
                    _ = v37
                }
                v38 := GenS2{F0: 753, F1: true, F2: "foo bar"}
                _ = v38
            }
            var v39 uint = 18446744073709551615
            _ = v39
        } else {
            for fi40 := 0; fi40 < 11; fi40++ {
                v1 -= ((1e10 * (float32(3.14159) * v1)) / (v1 * 1.5))
                _ = fi40
            }
            v17 = 1
        }
        _ = v20
    } else {
        v17 = v17
    }
    v4 += (v4 ^ (v4 | v4))
    fmt.Println(v4)
    if ((v19 - v19) != v19) {
        if (!false) {
            v1 = v1
            for fi41 := 0; fi41 < 15; fi41++ {
                if false {
                    if true {
                        fmt.Println(fi41)
                    }
                    v42 := GenS6{F0: -128}
                    _ = v42
                } else {
                    fmt.Println(v19, fi41)
                    var v43 int16 = 1
                    _ = v43
                }
                v4 -= (v4 >> ((0) & 63))
                _ = fi41
            }
            v17 = (v17 >> ((1) & 63))
        } else {
            fmt.Println(v1, v4)
        }
    } else {
        v19 = v19
        if ((false && true) || (!true)) {
            if false {
                var v44 uint64 = 2
                v45 := GenS7{F0: 33, F1: 0, F2: uint32(v44)}
                if false {
                    fmt.Println(v19, v3)
                    var v46 uint32 = v17
                    _ = v46
                } else {
                    v47 := GenS8{F0: (v4 >> ((65535) & 63)), F1: -199}
                    _ = v47
                }
                _ = v44
                _ = v45
            } else {
                switch -2147483648 {
                case -1:
                    if true {
                        var v48 uint16 = genFn0(v3, -1, v17)
                        _ = v48
                    } else {
                        var v49 int16 = 1
                        for fi50 := 0; fi50 < 12; fi50++ {
                            for fi51 := 0; fi51 < 5; fi51++ {
                                v17 = (v17 >> ((255) & 63))
                                v4 += (v4 >> ((879) & 63))
                                _ = fi51
                            }
                            if (true || (!true)) {
                                v52 := v6[2147483648]
                                for fi53 := 0; fi53 < 14; fi53++ {
                                    fmt.Println(v1)
                                    if false {
                                        v17 += (v52 + 2147483648)
                                        v1 = -2.25
                                    }
                                    _ = fi53
                                }
                                v4 -= 42
                                _ = v52
                            } else {
                                fmt.Println(v17, v17, v1)
                                var v54 uint64 = 670
                                _ = v54
                            }
                            v49 -= (v49 * 32767)
                            _ = fi50
                        }
                        _ = v49
                    }
                case 42:
                    fmt.Println(v17, v17, v3)
                }
            }
            v19 += (v19 >> ((624) & 63))
            for fi55 := 0; fi55 < 12; fi55++ {
                var v56 int32 = 489
                for fi57 := 0; fi57 < 19; fi57++ {
                    if (v1 <= v1) {
                        var v58 uint32 = (v17 << ((4294967295) & 63))
                        _ = v58
                    } else {
                        var v59 int16 = 32767
                        fmt.Println(v3, v56)
                        _ = v59
                    }
                    var v60 float32 = 0.0
                    v61 := GenS7{F0: 65535, F1: (v4 | int64(v17)), F2: (v17 % (((v17 << ((260) & 63))) | 1))}
                    _ = fi57
                    _ = v60
                    _ = v61
                }
                var v62 int16 = -1
                _ = fi55
                _ = v56
                _ = v62
            }
        }
        switch v3 {
        case "a":
            var v63 int8 = v19
            _ = v63
        case "\n\t":
            if (!false) {
                var v64 uint = uint(v17)
                _ = v64
            } else {
                var v65 int64 = (v4 % ((v4) | 1))
                _ = v65
            }
        case "":
            var v66 uint32 = v17
            _ = v66
        default:
            var v67 int8 = (^(v19))
            _ = v67
        }
    }
    var v68 uint32 = (v17 + v17)
    fmt.Println(v19)
    if false {
        delete(v6, (v17 | 0))
        var v69 uint64 = 18446744073709551615
        _ = v69
    } else {
        switch 808 {
        case -2147483648:
            var v70 int = -9223372036854775808
            _ = v70
        default:
            var v71 int8 = v19
            _ = v71
        }
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v68:", v68)
    _ = v1
    _ = v12
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
}
