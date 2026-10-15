package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 bool; F1 uint32; F2 bool }
type GenS4 struct { F0 bool; F1 int64; F2 int32 }
func genFn0(p0 int16, p1 int16, p2 int64) int16 {
    var v7 uint64 = 42
    var v8 uint32 = 4294967295
    if true {
        var v9 uint8 = 2
        switch (((("" + ("\n\t" + ("\n\t" + "foo bar"))) + "\n\t") + "") + ("" + ("a" + (("\n\t" + "hello") + ("a" + (("\n\t" + ("a" + "\n\t")) + ("hello" + "\n\t"))))))) {
        case "":
            var v10 uint16 = 42
            _ = v10
        case "\n\t":
            var v11 string = "foo bar"
            _ = v11
        default:
            var v12 uint32 = 42
            _ = v12
        }
        _ = v9
    } else {
        var v13 string = ""
        p1 -= p0
        _ = v13
    }
    switch 1 {
    case -2147483648:
        v14 := GenS1{F0: v7, F1: 42, F2: (p0 % (((p0 >> ((1) & 63))) | 1))}
        _ = v14
    case -1:
        fmt.Println(v8)
    case 9223372036854775807:
        if false {
            for fi15 := 0; fi15 < 15; fi15++ {
                for fi16 := 0; fi16 < 5; fi16++ {
                    v17 := GenS2{F0: 441, F1: (("foo bar" == (("foo bar" + ("foo bar" + "\n\t")) + ("hello" + ("" + (("\n\t" + "\n\t") + "a"))))) || (fi15 >= fi16))}
                    _ = fi16
                    _ = v17
                }
                v18 := []int{-359, -365, (fi15 * (fi15 | 888)), fi15}
                _ = fi15
                _ = v18
            }
            var v19 uint64 = (v7 & v7)
            for fi20 := 0; fi20 < 19; fi20++ {
                p0 = (p0 ^ (p0 & -307))
                v19 = (v7 * (v19 << ((1) & 63)))
                v21 := []int{int(v19)}
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 int8) int {
    switch ("hello" + "") {
    case "foo bar":
        p0 -= p0
    default:
        v22 := GenS2{F0: 65535, F1: (0.0 >= (float32((3.14159 + 1e10)) * (-2.25 * 1e10)))}
        _ = v22
    }
    p0 = p0
    _ = p0
    return 33
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 1e10
    v3 := []uint32{}
    v4 := make(map[string]int)
    var v5 int64 = 225
    var v6 uint64 = 59
    v5 -= -1
    var v23 uint32 = 426
    if ((255 <= 171) || true) {
        var v24 int64 = 1
        var v25 int16 = -1
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 5; fi26++ {
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 float32 = 1.5
                if true {
                    if (true || (true && false)) {
                        fmt.Println(v23, v23)
                        if (((genFn0(573, 42, (v5 | (^(v5)))) <= 42) || (1 < 53)) || (1e10 < (1e10 - 0.1))) {
                            var v29 int16 = 32767
                            _ = v29
                        } else {
                            if (1 >= -1) {
                                if (3.14159 >= v2) {
                                    v1["foo bar"] = 254
                                    v2 += 0.1
                                } else {
                                    var v30 uint64 = v6
                                    _ = v30
                                }
                                if (true && false) {
                                    v6 = v6
                                    fmt.Println(fi26, v5, v5)
                                    fmt.Println(fi27)
                                }
                                v3 = append(v3, 65535)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            v4[""] = genFn1(-6)
                            v31 := GenS1{F0: v6, F1: (fi26 / (((fi26 + fi26)) | 1)), F2: 32767}
                            _ = v31
                        }
                    } else {
                        var v32 float64 = 3.14159
                        switch (^(fi27)) {
                        case 703:
                            v33 := GenS1{F0: 277, F1: genFn1(0), F2: -32768}
                            _ = v33
                        case 128:
                            v23 += v23
                        default:
                            if (("foo bar" + ("\n\t" + (("hello" + "hello") + (("" + "") + (("" + "\n\t") + ("" + "\n\t")))))) >= "hello") {
                                var v34 int32 = -2147483648
                                fmt.Println(v6, fi26)
                                _ = v34
                            }
                        }
                        fmt.Println(fi27, v28)
                        if len(v3) > 0 {
                            v3[0] = 676
                        }
                        var v35 int = (fi26 << ((0) & 63))
                        var v36 bool = true
                        var v37 uint = 103
                        if len(v3) > 0 {
                            v3[0] = v23
                        }
                        v38 := GenS2{F0: v37, F1: (int8(v23) != 82)}
                        switch (fi26 + v35) {
                        case 1:
                            v6 -= v6
                        case 32767:
                            if (-32768 <= -32768) {
                                switch (v35 + v35) {
                                case 960:
                                    v6 = v6
                                case 204:
                                    var v39 uint16 = 0
                                    _ = v39
                                }
                                v40 := GenS2{F0: (^(v37)), F1: (v37 == (v37 - 2))}
                                var v41 int8 = 0
                                _ = v40
                                _ = v41
                            }
                        default:
                            if (v36 || (fi26 < (fi27 / ((823) | 1)))) {
                                if (-282 <= 627) {
                                    switch "foo bar" {
                                    case "":
                                        v42 := GenS2{F0: (^(v37)), F1: v36}
                                        _ = v42
                                    default:
                                        if true {
                                            var v43 float32 = (v28 + float32((v32 / -2.25)))
                                            if v36 {
                                                v44 := v4[(("\n\t" + ("" + ("a" + "foo bar"))) + "\n\t")]
                                                var v45 string = ("a" + (("\n\t" + "") + "a"))
                                                _ = v44
                                                _ = v45
                                            }
                                            var v46 int8 = -12
                                            fmt.Println(v2, v43, v28)
                                            _ = v43
                                            _ = v46
                                        }
                                    }
                                    v4["a"] = fi26
                                    if v36 {
                                        v37 = (v37 << ((uint(v6)) & 63))
                                        v37 = (v37 ^ (v37 | v37))
                                        v4["\n\t"] = 236
                                    }
                                } else {
                                    if ((false && (!(2 > uint16(v23)))) || (!(671 <= 305))) {
                                        v47 := GenS1{F0: (v6 + (v6 >> (((v37 / ((1) | 1))) & 63))), F1: (fi27 & -1), F2: -1}
                                        var v48 uint8 = 202
                                        v23 = uint32(v48)
                                        _ = v47
                                        _ = v48
                                    } else {
                                        v49 := v1["a"]
                                        v50 := GenS3{F0: v36, F1: v23, F2: v36}
                                        var v51 uint = v37
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    }
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                switch ("a" + ("a" + "")) {
                                case "hello":
                                    v52 := GenS2{F0: v37, F1: ((true && v36) && (false || (true && (!(0 != 365)))))}
                                    _ = v52
                                case "":
                                    if v36 {
                                        fmt.Println(v5)
                                        v28 = v28
                                    } else {
                                        v35 = (fi27 >> ((v37) & 63))
                                        var v53 float32 = 1e10
                                        var v54 float64 = v2
                                        var v55 uint32 = (v23 % ((1) | 1))
                                        fmt.Println(fi27)
                                        fmt.Println(v2)
                                        switch (("hello" + "a") + "\n\t") {
                                        case "foo bar":
                                            v6 = (^(v6))
                                        case "":
                                            delete(v4, "\n\t")
                                            v55 -= 2
                                        case "\n\t":
                                            v55 -= (v23 + 2147483648)
                                        }
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    }
                                case "foo bar":
                                    v56 := GenS1{F0: v6, F1: int(v6), F2: genFn0(0, 1, (v5 / ((v5) | 1)))}
                                    _ = v56
                                default:
                                    switch "\n\t" {
                                    case "\n\t":
                                        v2 = ((((v32 + (-2.25 * v2)) + v32) - (v2 - 0.0)) * v32)
                                    case "hello":
                                        v28 += ((1e10 - float32((((0.0 * (v2 - v32)) + v32) * v32))) + (v28 / 3.14159))
                                    case "a":
                                        v35 -= 128
                                    }
                                }
                            }
                        }
                        var v57 string = ("a" + "\n\t")
                        if len(v3) > 0 {
                            v3[0] = 2147483648
                        }
                        var v58 int16 = 222
                        fmt.Println(v5, v6)
                        var v59 uint = v37
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v57, v57, v32)
                        }
                        var v60 int16 = 401
                        _ = len(v3)
                        _ = v32
                        _ = v35
                        _ = v36
                        _ = v37
                        _ = v38
                        _ = v57
                        _ = v58
                        _ = v59
                        _ = v60
                    }
                } else {
                    var v61 int8 = 100
                    if true {
                        v62 := GenS4{F0: ("a" > ("foo bar" + "")), F1: 1, F2: -751}
                        v61 -= v61
                        v5 += (v5 % ((v5) | 1))
                        _ = v62
                    } else {
                        var v63 int8 = 100
                        var v64 uint16 = 210
                        _ = v63
                        _ = v64
                    }
                    v61 += int8(fi26)
                    _ = v61
                }
                v23 -= (v23 << ((42) & 63))
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    }
    fmt.Println(v6, v5, v6)
    var v65 string = "\n\t"
    v2 = -2.25
    if (v23 != (v23 % (((^(v23))) | 1))) {
        v5 -= (-(v5))
    }
    for fi66 := 0; fi66 < 18; fi66++ {
        v23 -= v23
        fmt.Println(fi66)
        _ = fi66
    }
    switch v65 {
    case "":
        for fi67 := 0; fi67 < 10; fi67++ {
            for fi68 := 0; fi68 < 19; fi68++ {
                var v69 float64 = v2
                var v70 float32 = -2.25
                v6 += v6
                _ = fi68
                _ = v69
                _ = v70
            }
            _ = fi67
        }
    case "\n\t":
        v2 = (3.14159 + (v2 + v2))
    }
    fmt.Println(v65, v5)
    var v71 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v71
}
