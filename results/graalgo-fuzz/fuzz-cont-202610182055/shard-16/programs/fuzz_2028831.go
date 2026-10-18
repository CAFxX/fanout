package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS4 struct { F0 bool; F1 int16; F2 int32 }
type GenS5 struct { F0 int32; F1 int64; F2 int16 }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 int16; F1 int8 }
func genFn0(p0 int64, p1 uint16, p2 uint32) uint32 {
    v4 := GenS1{F0: 2, F1: p0, F2: -9223372036854775808}
    fmt.Println(p2, p1)
    v5 := GenS2{F0: p1}
    for fi6 := 0; fi6 < 20; fi6++ {
        if false {
            p2 = (p2 % (((p2 | 613)) | 1))
        }
        for fi7 := 0; fi7 < 17; fi7++ {
            if false {
                var v8 bool = false
                var v9 int16 = 1
                fmt.Println(p2, fi6, fi6)
                _ = v8
                _ = v9
            }
            fmt.Println(p1, p1)
            var v10 string = ("\n\t" + "foo bar")
            _ = fi7
            _ = v10
        }
        _ = fi6
    }
    var v11 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v4
    _ = v5
    return (p2 * p2)
}
func genFn1(p0 int8, p1 uint) int32 {
    if true {
        v12 := GenS3{F0: 255, F1: -752, F2: 255}
        v13 := GenS3{F0: 2, F1: 1, F2: 1}
        _ = v12
        _ = v13
    }
    var v14 float32 = (float32(1e10) - 1e10)
    _ = p0
    _ = p1
    _ = v14
    return 2147483647
}

func main() {
    var v1 int16 = -32768
    var v2 uint32 = 2147483647
    var v3 uint32 = 261
    switch (("" + ("a" + "hello")) + "\n\t") {
    case "a":
        switch ("\n\t" + ((("foo bar" + "hello") + (("\n\t" + ("\n\t" + "\n\t")) + "hello")) + "")) {
        case "foo bar":
            var v15 uint16 = 1
            _ = v15
        default:
            fmt.Println(v3)
        }
    default:
        v16 := GenS1{F0: 213, F1: -1, F2: 42}
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 int16 = (v1 & (v1 & 0))
        _ = fi17
        _ = v18
    }
    switch 42 {
    case 0:
        var v19 uint = 2
        _ = v19
    case 4294967295:
        for fi20 := 0; fi20 < 17; fi20++ {
            switch fi20 {
            case -2147483648:
                v21 := []int{(-(fi20))}
                _ = v21
            case 1:
                if ((v1 >= 1000) || (-1 >= 0)) {
                    if (true && (!(!(true || (0 != 127))))) {
                        if (true || (!false)) {
                            v22 := GenS1{F0: 65535, F1: 9223372036854775807, F2: fi20}
                            for fi23 := 0; fi23 < 5; fi23++ {
                                switch "" {
                                case "\n\t":
                                    fmt.Println(v3)
                                    v2 -= genFn0(42, 65535, v2)
                                default:
                                    switch (fi23 % (((fi23 >> ((255) & 63))) | 1)) {
                                    case 32767:
                                        v1 += v1
                                    case 127:
                                        var v24 uint16 = 0
                                        _ = v24
                                    case 0:
                                        v25 := []int{(fi20 >> ((18446744073709551615) & 63))}
                                        _ = v25
                                    }
                                }
                                _ = fi23
                            }
                            _ = v22
                        }
                        v1 += v1
                        v26 := []int{(fi20 - (-(fi20)))}
                        _ = v26
                    } else {
                        fmt.Println(v3)
                        v2 -= (v3 | (v3 << ((951) & 63)))
                        v2 = genFn0(9223372036854775807, 32767, (^(v3)))
                    }
                } else {
                    fmt.Println(fi20, fi20, fi20)
                    v27 := []int{}
                    _ = v27
                }
            case 127:
                var v28 uint8 = 0
                _ = v28
            }
            if (1e10 == 0.1) {
                var v29 uint8 = 127
                var v30 uint16 = uint16(v29)
                var v31 float32 = 3.14159
                _ = v29
                _ = v30
                _ = v31
            }
            v2 += genFn0(290, 65535, (v2 >> ((1) & 63)))
            _ = fi20
        }
    case 9223372036854775807:
        v32 := []int{-536}
        _ = v32
    default:
        v2 -= (^(v2))
    }
    if (!(true || false)) {
        v33 := GenS2{F0: 55}
        v34 := GenS4{F0: true, F1: (v1 ^ -32768), F2: 42}
        v2 = v3
        _ = v33
        _ = v34
    } else {
        var v35 uint8 = 221
        var v36 uint16 = 42
        v2 = v2
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        var v38 string = ""
        fmt.Println(v3)
        _ = fi37
        _ = v38
    }
    fmt.Println(v3, v2, v2)
    if false {
        var v39 string = "\n\t"
        var v40 uint = uint(v3)
        var v41 float32 = 3.14159
        _ = v39
        _ = v40
        _ = v41
    } else {
        if ((1.5 - 1e10) > 0.0) {
            var v42 float64 = 3.14159
            v43 := GenS3{F0: 255, F1: 2147483647, F2: 0}
            if (812 < 18446744073709551615) {
                for fi44 := 0; fi44 < 11; fi44++ {
                    v3 -= genFn0(42, 320, v3)
                    _ = fi44
                }
                v45 := []int{404, 32767, 395}
                _ = v45
            }
            _ = v42
            _ = v43
        }
    }
    for fi46 := 0; fi46 < 18; fi46++ {
        v47 := []int{fi46, (fi46 | int(v1)), (^(fi46)), (fi46 & (fi46 | fi46))}
        var v48 float32 = 0.1
        _ = fi46
        _ = v47
        _ = v48
    }
    if true {
        v49 := []int{0, 9223372036854775807, 1, 4294967295}
        if (float32((1.5 / float64((3.14159 * 1.5)))) > (float32(-2.25) / 1e10)) {
            for fi50 := 0; fi50 < 7; fi50++ {
                switch ("foo bar" + ("\n\t" + "hello")) {
                case "a":
                    v3 -= 2147483648
                case "\n\t":
                    if len(v49) > 0 {
                        v49[0] = 0
                    }
                default:
                    v51 := GenS5{F0: -111, F1: -99, F2: v1}
                    _ = v51
                }
                switch ("foo bar" + "hello") {
                case "hello":
                    for fi52 := 0; fi52 < 9; fi52++ {
                        if (((genFn1(-120, 361) == -980) || (true && false)) && (!false)) {
                            if false {
                                if (false && true) {
                                    var v53 uint8 = 2
                                    _ = v53
                                } else {
                                    var v54 uint64 = 4294967295
                                    if ("" >= "foo bar") {
                                        var v55 uint64 = v54
                                        if (false && (true && false)) {
                                            fmt.Println(fi50, fi52, fi50)
                                            var v56 string = ""
                                            v1 += v1
                                            fmt.Println(v55, v56, v56)
                                            if true {
                                                var v57 uint32 = (v2 / ((v2) | 1))
                                                v58 := GenS6{F0: 528, F1: v2, F2: 0}
                                                _ = v57
                                                _ = v58
                                            }
                                            _ = v56
                                        } else {
                                            v59 := GenS7{F0: -128, F1: "a"}
                                            var v60 uint64 = v55
                                            var v61 uint16 = 65535
                                            _ = v59
                                            _ = v60
                                            _ = v61
                                        }
                                        _ = v55
                                    }
                                    _ = v54
                                }
                            }
                        } else {
                            switch ("hello" + "hello") {
                            case "":
                                var v62 int = fi52
                                _ = v62
                            }
                        }
                        _ = fi52
                    }
                default:
                    if ("\n\t" < (((("a" + ((("foo bar" + "") + "foo bar") + "foo bar")) + "") + "a") + "hello")) {
                        var v63 bool = (1 <= 4294967295)
                        v3 += genFn0(-584, 1, 65535)
                        _ = v63
                    }
                }
                _ = fi50
            }
            fmt.Println(v1, v1)
        }
        switch 2147483647 {
        case 32768:
            fmt.Println(v3, v2, v2)
        case 1:
            if (!("a" == "a")) {
                v64 := GenS6{F0: 32767, F1: genFn0(66, 2, (v2 ^ v2)), F2: 2}
                _ = v64
            }
        default:
            for fi65 := 0; fi65 < 18; fi65++ {
                if (!(-229 == 2147483647)) {
                    v3 += (v2 - v3)
                    var v66 int16 = -1
                    v3 += 114
                    _ = v66
                }
                switch "\n\t" {
                case "a":
                    if false {
                        switch (fi65 - 32767) {
                        case -9223372036854775808:
                            v2 -= v2
                        case 32768:
                            v49 = append(v49, (-(fi65)))
                        default:
                            fmt.Println(v3, v3)
                        }
                        v1 += (v1 << ((0) & 63))
                        var v67 int64 = 539
                        _ = v67
                    }
                case "hello":
                    v2 = (v2 | (v2 >> ((65535) & 63)))
                case "foo bar":
                    v68 := GenS8{F0: v1, F1: -117}
                    _ = v68
                default:
                    v69 := GenS8{F0: -1, F1: -77}
                    _ = v69
                }
                v3 = 593
                _ = fi65
            }
        }
        _ = v49
    }
    v1 += (v1 - (v1 / ((v1) | 1)))
    var v70 uint64 = 18446744073709551615
    v3 += v3
    v71 := []int{4294967295, 2147483647}
    var v72 float64 = 0.0
    switch -1 {
    case -171:
        var v73 uint64 = (v70 | v70)
        _ = v73
    case 1:
        var v74 int = 0
        _ = v74
    default:
        var v75 int = 32767
        _ = v75
    }
    if len(v71) > 0 {
        v71[0] = -9223372036854775808
    }
    v70 -= (^(v70))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v70:", v70)
    fmt.Println("slicelen v71:", len(v71), v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v70
    _ = v71
    _ = v72
}
