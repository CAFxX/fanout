package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 bool; F1 bool; F2 int16 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int32; F1 string; F2 int8 }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint }
func genFn0(p0 bool) uint8 {
    fmt.Println(p0)
    switch 128 {
    case 0:
        if p0 {
            var v6 bool = (-844 != 128)
            _ = v6
        }
    case 303:
        if p0 {
            v7 := GenS1{F0: 1, F1: -128}
            fmt.Println(p0)
            _ = v7
        }
    case 177:
        var v8 bool = (!((0.0 / 0.1) > ((0.0 / 1.0) + 0.0)))
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{(fi10 & (-(fi10)))}
        v11 = append(v11, (fi10 & (fi10 * -2147483648)))
        _ = fi10
        _ = v11
    }
    var v12 int16 = -1
    _ = p0
    _ = v12
    return 127
}

func main() {
    var v1 uint64 = 2
    var v2 int16 = 1
    v3 := []string{}
    v4 := []int{0}
    var v5 int = -1
    v13 := GenS1{F0: uint8(v1), F1: -69}
    var v14 float32 = (0.0 - (0.0 * 0.1))
    var v15 bool = true
    if (v15 && true) {
        if (((^(v1)) != v1) && ((1.5 + 1e10) <= 3.14159)) {
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := GenS1{F0: 169, F1: -128}
                var v18 string = (("foo bar" + "\n\t") + (("a" + "foo bar") + "a"))
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 uint64 = (v1 % ((v1) | 1))
            v5 = (^(v5))
            _ = v19
        }
        if v15 {
            for fi20 := 0; fi20 < 4; fi20++ {
                switch ("hello" + ("" + "a")) {
                case "a":
                    var v21 int64 = 9223372036854775807
                    _ = v21
                case "foo bar":
                    var v22 int = fi20
                    _ = v22
                case "\n\t":
                    for fi23 := 0; fi23 < 12; fi23++ {
                        if len(v4) > 0 {
                            v4[0] = 4294967295
                        }
                        _ = fi23
                    }
                default:
                    var v24 uint32 = 813
                    _ = v24
                }
                v2 = (v2 & v2)
                v2 -= v2
                _ = fi20
            }
            _ = len(v3)
        }
        if true {
            if len(v4) > 0 {
                v4[0] = 42
            }
            v2 = v2
        } else {
            if (float32(1.5) != (0.0 / (0.0 / (-2.25 - v14)))) {
                switch "" {
                case "":
                    if false {
                        v25 := GenS2{F0: 42, F1: 2}
                        _ = v25
                    }
                case "foo bar":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    v5 += (v5 << ((164) & 63))
                }
                v15 = v15
            }
        }
    } else {
        var v26 int32 = 0
        v27 := GenS2{F0: 2, F1: 42}
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 uint64 = (v1 & v1)
            fmt.Println(fi28, v1, v5)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi28
            _ = v29
        }
        _ = v26
        _ = v27
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        if (v15 || true) {
            v15 = (42 == 1)
            v31 := GenS1{F0: 255, F1: -128}
            _ = v31
        }
        for fi32 := 0; fi32 < 11; fi32++ {
            if true {
                if v15 {
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    if (204 > -2147483648) {
                        var v33 float64 = (-2.25 - 1e10)
                        fmt.Println(v33, fi32, fi32)
                        _ = v33
                    }
                } else {
                    var v34 string = "hello"
                    if len(v4) > 0 {
                        v4[0] = (^(fi30))
                    }
                    v35 := GenS1{F0: 222, F1: -12}
                    switch fi32 {
                    case -824:
                        switch int(v2) {
                        case -665:
                            v36 := GenS3{F0: (65535 <= 1), F1: (!false), F2: 32767}
                            _ = v36
                        case 42:
                            var v37 uint = 18446744073709551615
                            _ = v37
                        case 128:
                            v38 := GenS3{F0: v15, F1: false, F2: v2}
                            _ = v38
                        default:
                            v15 = v15
                        }
                        var v39 uint32 = 236
                        if v15 {
                            if len(v4) > 0 {
                                v4[0] = -1
                            }
                            fmt.Println(v1)
                        }
                        _ = v39
                    case 796:
                        var v40 int16 = v2
                        _ = v40
                    default:
                        fmt.Println(fi32)
                    }
                    switch v34 {
                    case "a":
                        v41 := GenS4{F0: 2, F1: v2}
                        _ = v41
                    case "hello":
                        if (42 >= 4294967295) {
                            var v42 string = v34
                            v43 := GenS3{F0: (0.0 >= (1e10 / (1e10 / 1.0))), F1: (-1 > 42), F2: v2}
                            fmt.Println(v14)
                            _ = v42
                            _ = v43
                        } else {
                            v5 = (v5 << ((1) & 63))
                        }
                    }
                    _ = v34
                    _ = v35
                }
                var v44 int16 = (-(v2))
                _ = v44
            } else {
                var v45 string = "\n\t"
                var v46 int8 = 127
                var v47 int32 = 0
                _ = v45
                _ = v46
                _ = v47
            }
            var v48 float32 = 1e10
            _ = fi32
            _ = v48
        }
        if len(v4) > 0 {
            v4[0] = v5
        }
        _ = fi30
    }
    v2 -= 854
    fmt.Println(v1)
    v2 += v2
    var v49 uint16 = 118
    _ = len(v3)
    v49 -= (v49 | 1)
    if v15 {
        v2 = (v2 * (-(v2)))
    }
    for fi50 := 0; fi50 < 15; fi50++ {
        if v15 {
            fmt.Println(fi50, v2)
            var v51 uint64 = (v1 - (^(v1)))
            _ = v51
        }
        if (2147483647 < 0) {
            for fi52 := 0; fi52 < 20; fi52++ {
                v53 := GenS5{F0: -1, F1: ("" + ""), F2: -128}
                _ = fi52
                _ = v53
            }
        } else {
            v49 += (v49 << ((186) & 63))
            var v54 int64 = -9223372036854775808
            _ = v54
        }
        _ = fi50
    }
    if true {
        var v55 uint32 = 169
        var v56 uint = 65535
        _ = v55
        _ = v56
    } else {
        v57 := GenS6{F0: v49, F1: 2, F2: 1}
        fmt.Println(v49, v5, v15)
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
}
