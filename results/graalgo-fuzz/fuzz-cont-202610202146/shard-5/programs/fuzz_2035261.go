package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint8; F2 bool }
type GenS4 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint64; F1 uint64 }
type GenS9 struct { F0 bool; F1 uint8; F2 int32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint; F1 uint64 }
type GenS14 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int, p1 int8) uint {
    var v6 int32 = 1
    v6 = v6
    _ = p0
    _ = p1
    _ = v6
    return 4294967295
}
func genFn1(p0 string, p1 uint8, p2 string) uint64 {
    var v7 uint8 = p1
    for fi8 := 0; fi8 < 14; fi8++ {
        p0 = p2
        switch 127 {
        case -745:
            var v9 int16 = 725
            _ = v9
        case -1:
            var v10 int32 = 1
            _ = v10
        case 9223372036854775807:
            p2 = ("foo bar" + ("foo bar" + p2))
        }
        if true {
            if false {
                var v11 int32 = -2147483648
                var v12 bool = ((v11 << ((236) & 63)) != (v11 << ((18446744073709551615) & 63)))
                for fi13 := 0; fi13 < 17; fi13++ {
                    v14 := GenS1{F0: 0}
                    if v12 {
                        if (false && (!v12)) {
                            if v12 {
                                var v15 int32 = (v11 % (((-(v11))) | 1))
                                var v16 uint32 = 380
                                fmt.Println(fi13, v16, p2)
                                _ = v15
                                _ = v16
                            }
                        }
                    }
                    v17 := []int{32767, (fi8 / ((-175) | 1))}
                    _ = fi13
                    _ = v14
                    _ = v17
                }
                _ = v11
                _ = v12
            } else {
                fmt.Println(fi8, fi8, p1)
                fmt.Println(fi8, fi8)
            }
            if (false && (111 > 759)) {
                var v18 string = p2
                v19 := []int{fi8}
                _ = v18
                _ = v19
            }
        } else {
            if true {
                fmt.Println(fi8, p2)
                var v20 int = fi8
                fmt.Println(p1, v20)
                _ = v20
            }
            v21 := []int{fi8, -2147483648}
            _ = v21
        }
        _ = fi8
    }
    v22 := []int{1}
    var v23 int8 = 0
    v22 = append(v22, -771)
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    _ = v7
    return 2
}

func main() {
    v1 := []string{"foo bar"}
    var v2 uint = 1
    v3 := make(map[uint32]uint32)
    v4 := []string{"a", "a"}
    var v5 float32 = 1e10
    v24 := GenS1{F0: 1}
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    v2 += (v2 - genFn0(32767, -128))
    var v25 int8 = 42
    fmt.Println(v2)
    v2 = v2
    v2 = v2
    var v26 int16 = 32767
    var v27 int = 698
    v5 = 3.14159
    for fi28 := 0; fi28 < 12; fi28++ {
        v29 := GenS2{F0: (fi28 | 127)}
        _ = fi28
        _ = v29
    }
    v30 := GenS3{F0: genFn0(v27, v25), F1: 0, F2: false}
    var v31 float32 = v5
    var v32 uint = (v2 * 42)
    if false {
        v33 := GenS1{F0: 0}
        v34 := GenS2{F0: v27}
        v35 := GenS4{F0: 2147483647, F1: true, F2: 0}
        _ = v33
        _ = v34
        _ = v35
    } else {
        var v36 int32 = -1
        if false {
            switch (("a" + "\n\t") + "a") {
            case "":
                if false {
                    switch (("hello" + "\n\t") + "") {
                    case "foo bar":
                        switch "\n\t" {
                        case "foo bar":
                            for fi37 := 0; fi37 < 3; fi37++ {
                                fmt.Println(v27)
                                if false {
                                    fmt.Println(v5, v36, v32)
                                    switch 1 {
                                    case 0:
                                        v38 := GenS2{F0: 712}
                                        _ = v38
                                    default:
                                        switch "a" {
                                        case "hello":
                                            if ((true && false) || false) {
                                                v39 := GenS5{F0: 2, F1: 32767}
                                                _ = v39
                                            }
                                        default:
                                            for fi40 := 0; fi40 < 2; fi40++ {
                                                v2 = v2
                                                switch -2147483648 {
                                                case 282:
                                                    var v41 int8 = (v25 + v25)
                                                    _ = v41
                                                }
                                                var v42 uint64 = genFn1("a", 116, ("\n\t" + ("\n\t" + "\n\t")))
                                                v5 = 1.5
                                                if false {
                                                    v31 = v31
                                                }
                                                if ((false && ((116 == v25) && true)) || ((^(v32)) < (v2 - 0))) {
                                                    v32 = genFn0((-(v27)), 42)
                                                    var v43 int8 = (v25 >> ((genFn0((v27 ^ fi40), v25)) & 63))
                                                    _ = v43
                                                } else {
                                                    switch 127 {
                                                    case 128:
                                                        if true {
                                                            v44 := GenS3{F0: 408, F1: 255, F2: true}
                                                            _ = v44
                                                        }
                                                    case -2147483648:
                                                        v5 += (v5 * 0.0)
                                                    case 181:
                                                        if true {
                                                            v31 = 0.0
                                                            v42 = (v42 ^ (v42 & (v42 << (((v2 - v32)) & 63))))
                                                        } else {
                                                            v45 := GenS6{F0: 932}
                                                            var v46 float64 = 0.0
                                                            v47 := GenS7{F0: -466}
                                                            _ = v45
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v48 := GenS2{F0: fi40}
                                                        _ = v48
                                                    case "":
                                                        v32 = genFn0((^(fi37)), (v25 + int8(v26)))
                                                    case "hello":
                                                        var v49 int64 = 9223372036854775807
                                                        _ = v49
                                                    default:
                                                        if true {
                                                            v50 := GenS8{F0: (v42 / ((genFn1(((("hello" + ("hello" + ("\n\t" + "a"))) + ("\n\t" + ("hello" + "a"))) + ("" + (("" + ("a" + ("" + ""))) + ("hello" + ("a" + "foo bar"))))), 127, "foo bar")) | 1)), F1: 755}
                                                            if ((9223372036854775807 > 42) && false) {
                                                                fmt.Println(fi37)
                                                            } else {
                                                                if (false || false) {
                                                                    v25 += 0
                                                                }
                                                                fmt.Println(v2, v31, v36)
                                                                var v51 float32 = 1e10
                                                                _ = v51
                                                            }
                                                            _ = v50
                                                        }
                                                    }
                                                    delete(v3, 2147483648)
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v52 := GenS8{F0: 18446744073709551615, F1: 672}
                                                var v53 int = (fi37 + (fi37 + (fi37 ^ (fi40 / (((v27 | (fi40 | (fi40 | 128)))) | 1)))))
                                                _ = fi40
                                                _ = v42
                                                _ = v52
                                                _ = v53
                                            }
                                        }
                                    }
                                }
                                _ = fi37
                            }
                        case "\n\t":
                            v5 += v5
                        case "hello":
                            var v54 int64 = -9223372036854775808
                            _ = v54
                        default:
                            if (v31 < v31) {
                                for fi55 := 0; fi55 < 4; fi55++ {
                                    for fi56 := 0; fi56 < 6; fi56++ {
                                        v57 := GenS9{F0: false, F1: 42, F2: v36}
                                        var v58 int8 = (v25 % (((v25 | (^(v25)))) | 1))
                                        _ = fi56
                                        _ = v57
                                        _ = v58
                                    }
                                    if (834 > 2) {
                                        if false {
                                            v59 := GenS3{F0: v2, F1: 42, F2: ((!(299 <= 42)) && ((2147483647 > uint32(v32)) && ("\n\t" <= "")))}
                                            _ = v59
                                        } else {
                                            var v60 uint = v32
                                            var v61 uint8 = 1
                                            _ = v60
                                            _ = v61
                                        }
                                        var v62 string = (("foo bar" + "a") + ("foo bar" + "a"))
                                        _ = v62
                                    } else {
                                        var v63 uint16 = 2
                                        var v64 int32 = 1
                                        if len(v4) > 0 {
                                            v4[0] = "hello"
                                        }
                                        _ = v63
                                        _ = v64
                                    }
                                    _ = fi55
                                }
                            } else {
                                var v65 uint = genFn0(v27, 0)
                                v27 = int(v25)
                                switch v27 {
                                case -977:
                                    fmt.Println(v25, v25, v31)
                                case -2147483648:
                                    for fi66 := 0; fi66 < 15; fi66++ {
                                        if (503 < 971) {
                                            v67 := GenS10{F0: 1}
                                            _ = v67
                                        }
                                        v32 = 794
                                        v4 = append(v4, "foo bar")
                                        _ = fi66
                                    }
                                case 42:
                                    v68 := GenS8{F0: genFn1("a", 42, "foo bar"), F1: 0}
                                    _ = v68
                                }
                                _ = v65
                            }
                        }
                    }
                    var v69 float64 = -2.25
                    switch ("" + ("hello" + (("hello" + (("\n\t" + ("" + "")) + "hello")) + ("hello" + (("hello" + "a") + ("foo bar" + "foo bar")))))) {
                    case "a":
                        var v70 int16 = (v26 >> ((65535) & 63))
                        _ = v70
                    }
                    _ = v69
                } else {
                    for fi71 := 0; fi71 < 10; fi71++ {
                        v72 := GenS11{F0: 0, F1: 208, F2: v36}
                        v26 -= -32768
                        var v73 int32 = (v36 & 1)
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                    for fi74 := 0; fi74 < 1; fi74++ {
                        for fi75 := 0; fi75 < 10; fi75++ {
                            if (!true) {
                                fmt.Println(v27, fi75)
                                if false {
                                    var v76 int64 = 42
                                    var v77 uint16 = 155
                                    v25 -= v25
                                    _ = v76
                                    _ = v77
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        v78 := GenS7{F0: 42}
                                        _ = v78
                                    case "hello":
                                        fmt.Println(v27, v32, v2)
                                    case "foo bar":
                                        if (true && ((true && true) || false)) {
                                            var v79 uint64 = 579
                                            if false {
                                                if false {
                                                    v3[uint32(v2)] = 65535
                                                } else {
                                                    var v80 int32 = v36
                                                    fmt.Println(v25)
                                                    v81 := GenS3{F0: v2, F1: 30, F2: (42 < 246)}
                                                    v82 := GenS12{F0: 42}
                                                    delete(v3, 524)
                                                    v5 += v5
                                                    var v83 uint64 = 18446744073709551615
                                                    _ = v80
                                                    _ = v81
                                                    _ = v82
                                                    _ = v83
                                                }
                                                v3[2147483648] = 2147483648
                                                fmt.Println(v26, v31)
                                            } else {
                                                v26 += 1
                                                v84 := v3[209]
                                                _ = v84
                                            }
                                            v85 := GenS13{F0: 255, F1: (v79 >> (((^(v2))) & 63))}
                                            _ = v79
                                            _ = v85
                                        }
                                    default:
                                        v26 = -511
                                    }
                                    var v86 float32 = (((v5 - (v5 * v31)) - (v31 / ((0.0 + (3.14159 / 1.5)) - 3.14159))) / 3.14159)
                                    _ = v86
                                }
                                if false {
                                    fmt.Println(fi74, fi74, v31)
                                }
                            } else {
                                v25 = (v25 >> ((v32) & 63))
                            }
                            _ = fi75
                        }
                        for fi87 := 0; fi87 < 6; fi87++ {
                            _ = len(v3)
                            _ = fi87
                        }
                        _ = fi74
                    }
                    var v88 uint = v32
                    _ = v88
                }
            case "hello":
                fmt.Println(v25)
            case "foo bar":
                v31 -= (v31 * (float32(0.1) / ((-2.25 * (0.1 - (1e10 - (float32(((float64(float32((float64((1.5 - v5)) - 1.5))) / float64(v31)) - 0.0)) * float32(float64(v5)))))) + v5)))
            default:
                if false {
                    v89 := GenS14{F0: 1, F1: 1}
                    _ = v89
                }
            }
            for fi90 := 0; fi90 < 7; fi90++ {
                if true {
                    if false {
                        var v91 float64 = 0.0
                        var v92 uint64 = genFn1("hello", 140, "")
                        _ = v91
                        _ = v92
                    }
                    var v93 uint16 = 0
                    _ = v93
                }
                _ = fi90
            }
            v31 = 0.0
        }
        _ = v36
    }
    for fi94 := 0; fi94 < 3; fi94++ {
        var v95 int64 = 1
        _ = fi94
        _ = v95
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
