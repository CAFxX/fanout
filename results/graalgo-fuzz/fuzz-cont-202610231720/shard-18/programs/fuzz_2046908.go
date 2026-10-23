package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 string }
type GenS2 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint }
func genFn0(p0 uint, p1 bool, p2 bool) uint32 {
    fmt.Println(p1, p1, p2)
    v7 := []int{4294967295}
    v8 := GenS1{F0: 1, F1: 2, F2: ("foo bar" + "foo bar")}
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 int8 = 81
            v12 := GenS1{F0: 139, F1: 964, F2: ("a" + ("" + ""))}
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 int = 1
        _ = fi9
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := []uint32{781, 2147483648, 895, 1}
    var v2 float32 = (0.1 - 1e10)
    var v3 int64 = -189
    v4 := make(map[int]uint32)
    v5 := make(map[int]int)
    var v6 int32 = 700
    if (127 > 128) {
        v14 := GenS2{F0: 1, F1: 601, F2: 2}
        _ = v14
    }
    if (1e10 != 0.1) {
        switch ("hello" + ("\n\t" + "\n\t")) {
        case "":
            var v15 uint32 = 355
            _ = v15
        case "foo bar":
            for fi16 := 0; fi16 < 11; fi16++ {
                if (32768 != fi16) {
                    v1 = append(v1, 149)
                    if (float64(((0.0 * v2) / ((3.14159 - (float32((((-2.25 + ((float64(v2) - -2.25) + 0.1)) - (0.0 / 1e10)) / 1e10)) + v2)) + (v2 / v2)))) > (float64(0.0) - ((1.5 + -2.25) - -2.25))) {
                        var v17 int = -589
                        _ = v17
                    } else {
                        for fi18 := 0; fi18 < 8; fi18++ {
                            fmt.Println(v3, v2)
                            _ = fi18
                        }
                        if ((v3 % (((-(v3))) | 1)) >= v3) {
                            if (!true) {
                                v6 += (-(v6))
                            } else {
                                var v19 float64 = (1e10 + float64(((-2.25 - (0.1 * (((float32((((0.0 * 0.0) * 0.1) + (((float64(0.0) + ((0.0 * 1.5) + (1e10 * 1e10))) * 0.0) / 1e10))) * v2) - (float32(0.1) * float32(0.1))) / ((v2 + (v2 / -2.25)) + v2)))) + float32((1.5 + ((0.0 + 1e10) / -2.25))))))
                                for fi20 := 0; fi20 < 17; fi20++ {
                                    _ = len(v1)
                                    v5[(fi16 << ((1) & 63))] = fi16
                                    if (!true) {
                                        _ = len(v1)
                                    } else {
                                        var v21 uint64 = 4294967295
                                        _ = v21
                                    }
                                    _ = fi20
                                }
                                var v22 uint32 = 448
                                _ = v19
                                _ = v22
                            }
                            var v23 string = "a"
                            var v24 int16 = -171
                            _ = v23
                            _ = v24
                        }
                    }
                }
                _ = fi16
            }
        }
    } else {
        if (!true) {
            switch "" {
            case "foo bar":
                var v25 uint16 = 1
                _ = v25
            case "hello":
                for fi26 := 0; fi26 < 3; fi26++ {
                    v3 = (v3 + (v3 % ((v3) | 1)))
                    var v27 string = ("a" + ("a" + ""))
                    _ = fi26
                    _ = v27
                }
            default:
                var v28 int64 = v3
                _ = v28
            }
            if false {
                v2 = (v2 / (((-2.25 - (3.14159 - v2)) - v2) * (v2 / (v2 * 0.1))))
                var v29 uint64 = 42
                for fi30 := 0; fi30 < 6; fi30++ {
                    var v31 uint = 42
                    switch "foo bar" {
                    case "hello":
                        v2 += v2
                    case "a":
                        if (true || (true && (603 == 1))) {
                            fmt.Println(v3, v3)
                            fmt.Println(v3)
                            switch "hello" {
                            case "\n\t":
                                v32 := GenS1{F0: v29, F1: uint16(v29), F2: ((("foo bar" + "hello") + "\n\t") + (("" + "") + "hello"))}
                                _ = v32
                            case "foo bar":
                                var v33 uint8 = 128
                                _ = v33
                            }
                        }
                    case "foo bar":
                        switch ((("hello" + "hello") + ((("a" + "foo bar") + "foo bar") + "a")) + "") {
                        case "\n\t":
                            var v34 int = (^(fi30))
                            _ = v34
                        case "hello":
                            delete(v4, fi30)
                        }
                    }
                    for fi35 := 0; fi35 < 3; fi35++ {
                        v2 += v2
                        var v36 int = (-(fi30))
                        _ = fi35
                        _ = v36
                    }
                    _ = fi30
                    _ = v31
                }
                _ = v29
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v37 uint = 4294967295
                fmt.Println(v37)
                _ = v37
            }
        }
        switch 32767 {
        case -615:
            v6 += v6
        case -9223372036854775808:
            v38 := GenS1{F0: 2, F1: 255, F2: ""}
            _ = v38
        case 0:
            if (42 == 42) {
                v2 -= v2
                var v39 int32 = (v6 ^ (v6 | (v6 / ((v6) | 1))))
                _ = v39
            } else {
                if len(v1) > 0 {
                    v1[0] = 42
                }
                if (127 == 1) {
                    if true {
                        v3 += 42
                        switch "" {
                        case "a":
                            for fi40 := 0; fi40 < 16; fi40++ {
                                v41 := v4[42]
                                v42 := GenS3{F0: int8(v41)}
                                var v43 int16 = 42
                                _ = fi40
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                        case "":
                            var v44 uint64 = 18446744073709551615
                            _ = v44
                        }
                    } else {
                        fmt.Println(v6, v3)
                    }
                    fmt.Println(v6, v2)
                } else {
                    var v45 uint = 42
                    _ = len(v4)
                    v3 = v3
                    _ = v45
                }
            }
        default:
            v2 = 1e10
        }
        for fi46 := 0; fi46 < 6; fi46++ {
            if (((!false) || true) || false) {
                v47 := GenS4{F0: (-128 >= 42), F1: 255}
                _ = v47
            } else {
                var v48 int16 = -32768
                var v49 uint8 = 127
                _ = v48
                _ = v49
            }
            _ = fi46
        }
    }
    var v50 uint = 18446744073709551615
    v50 += 65535
    v2 = float32((0.1 + 3.14159))
    var v51 uint8 = 127
    v6 -= v6
    if true {
        if false {
            v52 := GenS4{F0: false, F1: (v50 - v50)}
            var v53 uint32 = 146
            var v54 bool = (18446744073709551615 > 0)
            _ = v52
            _ = v53
            _ = v54
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
