package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int64; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 int16; F2 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool; F1 int64; F2 int32 }
func genFn0(p0 bool, p1 int16, p2 int64) int16 {
    v4 := []int{}
    if (p0 && (1 == 119)) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if len(v4) > 0 {
            v4[0] = 593
        }
    }
    var v5 uint16 = 304
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p1
}
func genFn1(p0 uint64, p1 int16) int {
    if ((0.1 + (0.0 * 3.14159)) != 1.5) {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 string = "\n\t"
            _ = fi6
            _ = v7
        }
        p1 -= -1
    } else {
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS1{F0: ""}
                var v11 int8 = -1
                fmt.Println(fi9)
                v12 := []int{int(p1), int(v11), fi9}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 float64 = (0.0 * float64(float32(1e10)))
                fmt.Println(fi13)
                fmt.Println(p0, p0, v14)
                _ = fi13
                _ = v14
            }
            if (!(true && (127 != 126))) {
                if true {
                    switch fi8 {
                    case -9223372036854775808:
                        var v15 int32 = 2147483647
                        _ = v15
                    case -1:
                        var v16 int = -2147483648
                        _ = v16
                    }
                    fmt.Println(fi8, p0)
                    switch (fi8 + fi8) {
                    case 32767:
                        v17 := []int{653, (fi8 % ((fi8) | 1))}
                        _ = v17
                    case 484:
                        v18 := GenS2{F0: 42, F1: int64(p1), F2: 1}
                        _ = v18
                    case 858:
                        for fi19 := 0; fi19 < 16; fi19++ {
                            v20 := GenS3{F0: (p0 & p0), F1: 32767, F2: 2147483648}
                            p1 = genFn0(true, p1, -1)
                            if false {
                                if true {
                                    var v21 int = 32767
                                    v21 -= fi19
                                    _ = v21
                                }
                                var v22 int32 = 0
                                v22 += v22
                                _ = v22
                            }
                            _ = fi19
                            _ = v20
                        }
                    default:
                        if (1e10 <= float32(0.1)) {
                            if (true && (true || (fi8 == (-(fi8))))) {
                                p1 += (p1 | (-(p1)))
                            } else {
                                for fi23 := 0; fi23 < 13; fi23++ {
                                    v24 := GenS2{F0: 127, F1: -818, F2: int64(fi8)}
                                    _ = fi23
                                    _ = v24
                                }
                                var v25 string = "foo bar"
                                _ = v25
                            }
                            if (false || (true && false)) {
                                switch "hello" {
                                case "a":
                                    p0 = 18446744073709551615
                                case "\n\t":
                                    var v26 bool = true
                                    _ = v26
                                default:
                                    var v27 uint8 = 0
                                    _ = v27
                                }
                                switch (fi8 ^ fi8) {
                                case 0:
                                    p0 = p0
                                default:
                                    p0 = p0
                                }
                                v28 := GenS3{F0: p0, F1: 341, F2: 0}
                                _ = v28
                            }
                            v29 := GenS3{F0: p0, F1: 182, F2: 916}
                            _ = v29
                        }
                    }
                } else {
                    switch fi8 {
                    case 590:
                        if true {
                            var v30 int32 = 1
                            _ = v30
                        } else {
                            fmt.Println(fi8)
                        }
                    case -708:
                        p0 = p0
                    }
                }
                p0 = (^(p0))
            }
            _ = fi8
        }
        v31 := GenS2{F0: 100, F1: 9223372036854775807, F2: -1}
        for fi32 := 0; fi32 < 10; fi32++ {
            var v33 int32 = 0
            switch ("foo bar" + "hello") {
            case "hello":
                v34 := GenS2{F0: -87, F1: 1, F2: -9223372036854775808}
                _ = v34
            case "foo bar":
                v35 := GenS1{F0: ("foo bar" + "")}
                _ = v35
            default:
                switch ("\n\t" + "hello") {
                case "a":
                    v36 := GenS4{F0: 188, F1: 255, F2: 127}
                    _ = v36
                case "":
                    var v37 uint16 = 858
                    _ = v37
                case "foo bar":
                    var v38 float64 = 0.1
                    _ = v38
                }
            }
            v39 := []int{(^(fi32)), fi32}
            _ = fi32
            _ = v33
            _ = v39
        }
        _ = v31
    }
    v40 := GenS2{F0: 73, F1: 42, F2: 9223372036854775807}
    p1 -= p1
    _ = p0
    _ = p1
    _ = v40
    return 32768
}

func main() {
    var v1 int32 = 2147483647
    var v2 float32 = (1.5 + -2.25)
    var v3 uint32 = 564
    switch -2147483648 {
    case 127:
        if false {
            var v41 int = 211
            v42 := GenS5{F0: 1}
            var v43 uint = 0
            _ = v41
            _ = v42
            _ = v43
        } else {
            if false {
                var v44 uint32 = v3
                var v45 uint32 = (v3 % ((v44) | 1))
                switch ("" + "hello") {
                case "":
                    switch ("a" + "\n\t") {
                    case "foo bar":
                        v46 := GenS2{F0: 127, F1: 1, F2: 0}
                        _ = v46
                    default:
                        v1 += 42
                    }
                case "a":
                    if true {
                        v47 := GenS3{F0: 739, F1: 0, F2: v45}
                        for fi48 := 0; fi48 < 1; fi48++ {
                            var v49 string = "a"
                            _ = fi48
                            _ = v49
                        }
                        _ = v47
                    }
                default:
                    for fi50 := 0; fi50 < 6; fi50++ {
                        switch genFn1(4294967295, 42) {
                        case 9223372036854775807:
                            var v51 bool = false
                            _ = v51
                        }
                        for fi52 := 0; fi52 < 5; fi52++ {
                            v1 = (v1 % (((v1 / ((670) | 1))) | 1))
                            if true {
                                if ("\n\t" != "") {
                                    v45 -= v3
                                    switch ("a" + "a") {
                                    case "a":
                                        if true {
                                            v53 := []int{}
                                            _ = v53
                                        } else {
                                            v45 -= v44
                                            v2 += 3.14159
                                            v54 := []int{fi50}
                                            _ = v54
                                        }
                                    default:
                                        if true {
                                            if false {
                                                switch 128 {
                                                case 2147483647:
                                                    v2 = (float32(3.14159) + (((((float32((1.5 - 0.1)) * 1.5) + -2.25) + v2) - 1e10) / float32(1e10)))
                                                case -9223372036854775808:
                                                    v55 := []int{fi50}
                                                    _ = v55
                                                }
                                            } else {
                                                v2 -= (((v2 - ((float32((3.14159 * (-2.25 * 1.5))) / 1e10) / v2)) * 1.5) - (v2 + v2))
                                            }
                                        } else {
                                            var v56 string = ("a" + "hello")
                                            fmt.Println(v3)
                                            var v57 int32 = 2147483647
                                            _ = v56
                                            _ = v57
                                        }
                                    }
                                    v58 := []int{}
                                    v59 := GenS6{F0: 0, F1: genFn0(false, 369, -9223372036854775808), F2: genFn1(511, genFn0((2 != (v45 * 2147483648)), int16(fi52), 668))}
                                    var v60 string = ""
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                            }
                            v1 -= (v1 | v1)
                            _ = fi52
                        }
                        _ = fi50
                    }
                }
                _ = v44
                _ = v45
            }
        }
    case 1:
        var v61 uint64 = 289
        _ = v61
    default:
        v62 := []int{genFn1(2, 32767), -31, 0}
        _ = v62
    }
    for fi63 := 0; fi63 < 16; fi63++ {
        if (v1 <= (v1 - 2147483647)) {
            v64 := GenS7{F0: 1}
            v65 := GenS2{F0: -99, F1: -9223372036854775808, F2: 1}
            fmt.Println(fi63, v3, fi63)
            _ = v64
            _ = v65
        }
        v66 := GenS1{F0: "hello"}
        if (!false) {
            for fi67 := 0; fi67 < 19; fi67++ {
                v68 := []int{-558, genFn1(412, -32768)}
                fmt.Println(v2)
                var v69 uint8 = 42
                _ = fi67
                _ = v68
                _ = v69
            }
            for fi70 := 0; fi70 < 4; fi70++ {
                v71 := []int{}
                _ = fi70
                _ = v71
            }
            v72 := []int{(^(fi63)), fi63}
            _ = v72
        }
        _ = fi63
        _ = v66
    }
    if false {
        v73 := GenS8{F0: 2}
        var v74 uint64 = 1
        _ = v73
        _ = v74
    } else {
        v2 = (1e10 / 1.0)
        fmt.Println(v3, v2)
        v75 := GenS9{F0: (1.5 > 0.0), F1: 9223372036854775807, F2: (v1 * int32(v3))}
        _ = v75
    }
    fmt.Println(v1, v3)
    var v76 uint16 = 323
    v77 := GenS5{F0: (v3 / ((v3) | 1))}
    fmt.Println(v2, v1)
    for fi78 := 0; fi78 < 7; fi78++ {
        v79 := GenS3{F0: 815, F1: v76, F2: 1}
        _ = fi78
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v76
    _ = v77
}
