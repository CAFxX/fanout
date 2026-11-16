package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 int8; F1 uint32; F2 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint8) int8 {
    switch 4294967295 {
    case -360:
        if (152 == p1) {
            var v4 uint32 = 2147483648
            v5 := []int{32767, 2147483647, -1}
            v6 := GenS1{F0: 127}
            _ = v4
            _ = v5
            _ = v6
        } else {
            fmt.Println(p0, p1)
            v7 := []int{-2147483648, 399, 432, 228}
            _ = v7
        }
    default:
        var v8 int64 = -679
        _ = v8
    }
    v9 := GenS1{F0: 127}
    var v10 string = "foo bar"
    v11 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return -44
}
func genFn1(p0 uint) bool {
    v12 := GenS2{F0: 2147483648, F1: 579}
    for fi13 := 0; fi13 < 15; fi13++ {
        p0 = (^(p0))
        var v14 bool = (95 <= 127)
        var v15 int8 = 127
        _ = fi13
        _ = v14
        _ = v15
    }
    switch 5 {
    case 128:
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 uint64 = 0
            v18 := GenS2{F0: 873, F1: 0}
            _ = fi16
            _ = v17
            _ = v18
        }
    case -2147483648:
        var v19 int = 127
        _ = v19
    }
    if false {
        v20 := GenS3{F0: "\n\t"}
        v21 := GenS1{F0: 0}
        _ = v20
        _ = v21
    }
    var v22 uint8 = 37
    _ = p0
    _ = v12
    _ = v22
    return (1 < uint32(p0))
}

func main() {
    var v1 float32 = 0.1
    var v2 bool = (!(255 != 932))
    v3 := make(map[int]int)
    v2 = v2
    v2 = (!true)
    var v23 uint = 2
    v24 := []int{0}
    switch 0 {
    case 9223372036854775807:
        v25 := GenS4{F0: 42}
        _ = v25
    case 42:
        var v26 uint8 = 2
        _ = v26
    case 127:
        v23 -= (v23 << (((v23 ^ 18446744073709551615)) & 63))
    }
    v2 = true
    v27 := GenS5{F0: 9223372036854775807, F1: (1 > -128)}
    v1 = v1
    delete(v3, 639)
    v28 := GenS6{F0: 816}
    v1 -= -2.25
    v23 += (v23 >> (((^(v23))) & 63))
    var v29 int8 = 0
    if (v2 || false) {
        for fi30 := 0; fi30 < 20; fi30++ {
            if v2 {
                v31 := v3[fi30]
                _ = v31
            } else {
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
            }
            _ = fi30
        }
        var v32 float32 = (((0.1 - 0.1) + (3.14159 - 1e10)) - v1)
        _ = v32
    } else {
        v2 = (true && (365 <= -1))
        fmt.Println(v23, v29)
        var v33 uint16 = 32767
        _ = v33
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        for fi35 := 0; fi35 < 1; fi35++ {
            if len(v24) > 0 {
                v24[0] = fi34
            }
            if true {
                if v2 {
                    v23 = v23
                    v36 := v3[128]
                    v29 -= (v29 - (v29 % (((v29 & (-(v29)))) | 1)))
                    v24 = append(v24, fi35)
                    if v2 {
                        var v37 int16 = -1
                        fmt.Println(v37, v1)
                        fmt.Println(v2, v29)
                        _ = v37
                    } else {
                        if ((!((v36 % ((v36) | 1)) <= 4294967295)) || (4294967295 > 567)) {
                            if (("a" < "foo bar") || (!v2)) {
                                v2 = (-1 != (fi35 / ((fi35) | 1)))
                                fmt.Println(v2, v2, v29)
                                fmt.Println(fi35)
                            } else {
                                _ = len(v24)
                            }
                        }
                    }
                    fmt.Println(v2, fi35)
                    _ = v36
                } else {
                    v38 := GenS7{F0: 144, F1: -1}
                    _ = v38
                }
                switch (^(fi35)) {
                case 128:
                    var v39 string = (((("" + ("\n\t" + ("foo bar" + ""))) + "") + ("a" + "")) + "")
                    _ = v39
                case -2147483648:
                    if (32767 > 0) {
                        v29 = (v29 << ((v23) & 63))
                        var v40 float32 = (float32(float64((v1 + 3.14159))) * v1)
                        _ = v40
                    } else {
                        var v41 uint16 = 42
                        switch (fi35 >> (((^(v23))) & 63)) {
                        case -1:
                            v42 := GenS8{F0: -97, F1: 2147483647, F2: (v23 & (v23 ^ 0))}
                            _ = v42
                        case -388:
                            if len(v24) > 0 {
                                v24[0] = (fi35 >> ((128) & 63))
                            }
                            v43 := GenS4{F0: -32768}
                            _ = v43
                        default:
                            v44 := GenS6{F0: 2147483648}
                            _ = v44
                        }
                        v45 := v3[fi35]
                        v1 = 0.1
                        _ = v41
                        _ = v45
                    }
                case 477:
                    switch "foo bar" {
                    case "":
                        v24 = append(v24, fi35)
                    case "hello":
                        v29 -= (v29 >> ((v23) & 63))
                    case "\n\t":
                        v23 += (v23 + 18446744073709551615)
                    default:
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                    }
                    v46 := GenS8{F0: 0, F1: uint32(v23), F2: (v23 & (v23 + 1))}
                    _ = v46
                default:
                    var v47 uint16 = 32767
                    _ = v47
                }
                v48 := GenS9{F0: 478}
                v23 = v23
                fmt.Println(v29, fi34, v29)
                if (0.1 > 1.5) {
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                    delete(v3, 4294967295)
                    v49 := GenS5{F0: 489, F1: (1.5 > float64(v1))}
                    _ = v49
                } else {
                    v50 := GenS10{F0: 658}
                    _ = v50
                }
                _ = v48
            } else {
                var v51 int32 = -752
                _ = v51
            }
            v52 := GenS11{F0: 1}
            var v53 uint64 = 18446744073709551615
            _ = fi35
            _ = v52
            _ = v53
        }
        if (v2 || (false || ((-215 == 32767) && (217 <= 51)))) {
            var v54 int8 = (v29 - genFn0(868, 0))
            for fi55 := 0; fi55 < 1; fi55++ {
                switch (("foo bar" + "hello") + ("foo bar" + "hello")) {
                case "foo bar":
                    if (v2 && (v1 < v1)) {
                        var v56 int = 128
                        fmt.Println(v54, v23)
                        var v57 int = (fi34 / (((fi55 | (fi34 << ((v23) & 63)))) | 1))
                        _ = v56
                        _ = v57
                    } else {
                        v1 -= (v1 * (((v1 - 0.0) + (-2.25 / v1)) - v1))
                        v29 += genFn0(4294967295, 87)
                    }
                case "\n\t":
                    if (!v2) {
                        v58 := GenS11{F0: genFn0(938, 128)}
                        _ = v58
                    }
                }
                v59 := GenS11{F0: (v29 | v29)}
                var v60 string = ("a" + "\n\t")
                _ = fi55
                _ = v59
                _ = v60
            }
            var v61 int64 = 1
            _ = v54
            _ = v61
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
