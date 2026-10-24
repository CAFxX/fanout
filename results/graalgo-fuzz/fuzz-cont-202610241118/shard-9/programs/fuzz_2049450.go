package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 uint16; F1 string; F2 uint64 }
type GenS7 struct { F0 uint32; F1 int; F2 uint32 }
func genFn0(p0 uint64, p1 int) uint16 {
    var v4 int = 127
    v5 := []int{p1}
    p1 = v4
    v5 = append(v5, (p1 << ((255) & 63)))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 bool, p1 uint8) bool {
    v6 := GenS1{F0: 2, F1: 1}
    p1 -= p1
    var v7 int64 = 1
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int16 = 0
        v10 := GenS1{F0: 1, F1: 0}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return true
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 0
    var v3 uint64 = 872
    if false {
        var v11 int64 = -355
        _ = v11
    } else {
        v2 += v2
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        for fi13 := 0; fi13 < 7; fi13++ {
            v2 = 4294967295
            v1[438] = fi13
            v14 := GenS2{F0: (uint8(v3) >= 42), F1: 42}
            var v15 string = "foo bar"
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 string = ""
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := GenS2{F0: true, F1: (v3 << ((671) & 63))}
            v2 += (v2 * v2)
            var v19 int32 = 1
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi12
        _ = v16
    }
    var v20 int = 127
    if (1 <= 100) {
        if false {
            switch 127 {
            case -2147483648:
                fmt.Println(v20)
            case 926:
                for fi21 := 0; fi21 < 1; fi21++ {
                    if (((v2 / (((v2 * 1)) | 1)) >= (v2 & v2)) || (0 < 897)) {
                        v3 = 2
                        var v22 uint32 = 581
                        v3 += (^(v3))
                        _ = v22
                    }
                    v3 = v3
                    _ = fi21
                }
            default:
                switch ("\n\t" + (("\n\t" + "") + "")) {
                case "":
                    if (("foo bar" != "hello") && true) {
                        v23 := GenS1{F0: uint32(v2), F1: (v2 * v2)}
                        fmt.Println(v3, v3)
                        _ = v23
                    } else {
                        v2 += (v2 >> ((1) & 63))
                        v3 = (v3 & 4294967295)
                        v24 := GenS3{F0: -1}
                        _ = v24
                    }
                case "\n\t":
                    v25 := v1[4294967295]
                    _ = v25
                default:
                    fmt.Println(v2, v3)
                }
            }
            switch "\n\t" {
            case "\n\t":
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 float32 = -2.25
                    _ = fi26
                    _ = v27
                }
            default:
                for fi28 := 0; fi28 < 6; fi28++ {
                    if true {
                        var v29 int = 32767
                        _ = v29
                    } else {
                        switch "foo bar" {
                        case "a":
                            _ = len(v1)
                        case "":
                            v30 := GenS3{F0: 462}
                            _ = v30
                        case "foo bar":
                            v31 := []int{(fi28 | v20)}
                            _ = v31
                        default:
                            if true {
                                if (-2.25 >= (-2.25 - -2.25)) {
                                    for fi32 := 0; fi32 < 13; fi32++ {
                                        var v33 bool = true
                                        switch "a" {
                                        case "":
                                            _ = len(v1)
                                        default:
                                            if (42 < 42) {
                                                if (255 >= 307) {
                                                    if (((^(v20)) < fi32) || (!true)) {
                                                        v1[2147483648] = fi32
                                                        switch "" {
                                                        case "":
                                                            v34 := GenS3{F0: 672}
                                                            _ = v34
                                                        case "hello":
                                                            v20 = (v20 * (fi28 * (^(fi32))))
                                                        case "\n\t":
                                                            if ((1 == 2147483647) || ((((("foo bar" + "") + ("a" + ("hello" + "a"))) + "foo bar") + ("" + "foo bar")) <= "")) {
                                                                var v35 int = (^(v20))
                                                                _ = v35
                                                            }
                                                        }
                                                        var v36 uint64 = 18446744073709551615
                                                        var v37 uint64 = (v3 % ((4294967295) | 1))
                                                        if (630 > 1000) {
                                                            v38 := GenS4{F0: (v36 >> ((v2) & 63)), F1: (("foo bar" + (("foo bar" + ("a" + "foo bar")) + ((("hello" + "\n\t") + "a") + ("a" + "")))) + ("\n\t" + "a"))}
                                                            v39 := GenS4{F0: v3, F1: "\n\t"}
                                                            var v40 int = fi32
                                                            _ = v38
                                                            _ = v39
                                                            _ = v40
                                                        } else {
                                                            v37 = v3
                                                            v33 = false
                                                            fmt.Println(v2)
                                                        }
                                                        _ = v36
                                                        _ = v37
                                                    } else {
                                                        v33 = v33
                                                        if (false && true) {
                                                            switch (v20 & -1) {
                                                            case 42:
                                                                v41 := []int{v20, 275, (v20 << (((v2 - v2)) & 63)), 127}
                                                                var v42 uint32 = 428
                                                                if len(v41) >= 2 {
                                                                    v41 = v41[1:len(v41)]
                                                                }
                                                                _ = v41
                                                                _ = v42
                                                            case 438:
                                                                v43 := GenS4{F0: (v3 ^ v3), F1: "foo bar"}
                                                                _ = v43
                                                            default:
                                                                v44 := GenS5{F0: v33, F1: 0}
                                                                _ = v44
                                                            }
                                                            switch (fi32 % ((v20) | 1)) {
                                                            case 4294967295:
                                                                if false {
                                                                    if v33 {
                                                                        var v45 float64 = float64((float32((0.1 - (-2.25 * -2.25))) * 1e10))
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            var v46 int32 = 42
                                                                            _ = v46
                                                                        case "\n\t":
                                                                            var v47 bool = ((true && (v33 && (0 < 1))) && (v33 || true))
                                                                            _ = v47
                                                                        default:
                                                                            v48 := []int{}
                                                                            _ = v48
                                                                        }
                                                                        v49 := GenS6{F0: 0, F1: (("" + ("foo bar" + "a")) + ""), F2: v3}
                                                                        v3 = (v3 ^ (^(v3)))
                                                                        fmt.Println(v20, v3, v20)
                                                                        _ = v45
                                                                        _ = v49
                                                                    } else {
                                                                        var v50 uint8 = 230
                                                                        if (false && false) {
                                                                            v51 := v1[491]
                                                                            var v52 string = ("\n\t" + ("a" + "a"))
                                                                            _ = v51
                                                                            _ = v52
                                                                        } else {
                                                                            v20 = 128
                                                                            switch -661 {
                                                                            case 188:
                                                                                _ = len(v1)
                                                                            case 32767:
                                                                                v20 -= fi32
                                                                            default:
                                                                                v53 := []int{(fi32 | v20), (v20 + (fi28 >> ((v2) & 63))), v20, (fi28 % (((fi32 % ((9223372036854775807) | 1))) | 1))}
                                                                                _ = v53
                                                                            }
                                                                        }
                                                                        v3 = (v3 - (v3 % (((v3 / (((v3 + v3)) | 1))) | 1)))
                                                                        _ = v50
                                                                    }
                                                                    var v54 uint8 = 2
                                                                    _ = v54
                                                                }
                                                            case -1:
                                                                v55 := GenS5{F0: v33, F1: -1}
                                                                _ = v55
                                                            case 32768:
                                                                var v56 int16 = -424
                                                                _ = v56
                                                            default:
                                                                switch ("" + "") {
                                                                case "":
                                                                    fmt.Println(fi28)
                                                                case "hello":
                                                                    if ((!v33) || v33) {
                                                                        if (!false) {
                                                                            v20 += 156
                                                                            v57 := []int{4294967295}
                                                                            _ = v57
                                                                        } else {
                                                                            v58 := []int{fi32, 42, fi28, fi32}
                                                                            switch (-(fi28)) {
                                                                            case 32768:
                                                                                var v59 string = "a"
                                                                                _ = v59
                                                                            default:
                                                                                v33 = (239 < 255)
                                                                            }
                                                                            v33 = true
                                                                            var v60 float32 = (1.5 + (1e10 / (0.1 + 0.1)))
                                                                            _ = v58
                                                                            _ = v60
                                                                        }
                                                                        fmt.Println(v20)
                                                                        var v61 int64 = 42
                                                                        _ = v61
                                                                    }
                                                                    v1[2147483648] = (fi28 & fi28)
                                                                default:
                                                                    v20 = -2147483648
                                                                }
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    var v62 uint32 = 42
                                                    var v63 uint32 = (v62 | (v62 / ((v62) | 1)))
                                                    _ = v62
                                                    _ = v63
                                                }
                                            } else {
                                                if true {
                                                    v2 = v2
                                                    v64 := []int{fi32, fi28, v20, -28}
                                                    _ = v64
                                                }
                                                v65 := v1[uint32(v2)]
                                                _ = v65
                                            }
                                        }
                                        _ = fi32
                                        _ = v33
                                    }
                                    v66 := GenS5{F0: true, F1: 9223372036854775807}
                                    _ = v66
                                } else {
                                    v1[2147483647] = fi28
                                    fmt.Println(fi28, v20)
                                }
                            } else {
                                v67 := v1[903]
                                var v68 int = v67
                                v69 := []int{v67, (v20 ^ (v67 / ((-2147483648) | 1))), int(v3), (fi28 / ((917) | 1))}
                                _ = v67
                                _ = v68
                                _ = v69
                            }
                        }
                        if ((v3 >= v3) || true) {
                            if true {
                                v70 := []int{(v20 << (((^(v2))) & 63)), (v20 & v20)}
                                _ = v70
                            }
                            v2 = v2
                        }
                        v2 -= (^(v2))
                    }
                    _ = fi28
                }
            }
        } else {
            v20 += 200
        }
        switch v20 {
        case 42:
            var v71 uint16 = 2
            _ = v71
        }
    } else {
        v20 = (v20 & (v20 / ((v20) | 1)))
        for fi72 := 0; fi72 < 4; fi72++ {
            v3 = (v3 - v3)
            v2 = v2
            _ = fi72
        }
    }
    v73 := GenS7{F0: 988, F1: 42, F2: 42}
    var v74 int8 = int8(v3)
    v3 = 18446744073709551615
    v74 += -128
    for fi75 := 0; fi75 < 20; fi75++ {
        for fi76 := 0; fi76 < 7; fi76++ {
            v2 += v2
            _ = fi76
        }
        _ = fi75
    }
    v77 := []int{v20, -9223372036854775808, -660}
    if false {
        var v78 uint64 = (^(v3))
        var v79 bool = true
        v78 = v78
        _ = v78
        _ = v79
    }
    var v80 bool = (42 != genFn0((^(v3)), (v20 * v20)))
    v81 := GenS3{F0: 0}
    if true {
        v74 = v74
    } else {
        for fi82 := 0; fi82 < 3; fi82++ {
            switch "a" {
            case "\n\t":
                for fi83 := 0; fi83 < 17; fi83++ {
                    v20 = (fi82 | 2147483647)
                    _ = fi83
                }
            case "":
                v84 := GenS7{F0: 2, F1: fi82, F2: 30}
                _ = v84
            case "foo bar":
                var v85 bool = v80
                _ = v85
            }
            _ = fi82
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v74:", v74)
    fmt.Println("slicelen v77:", len(v77), v77)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v73
    _ = v74
    _ = v77
    _ = v80
    _ = v81
}
