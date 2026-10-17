package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 int }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 int8; F1 int; F2 int16 }
type GenS11 struct { F0 int8; F1 int8 }
type GenS12 struct { F0 uint; F1 uint; F2 int64 }
type GenS13 struct { F0 bool; F1 uint32; F2 int64 }
func genFn0(p0 uint8, p1 uint) string {
    v6 := []int{-971, 32768, -2147483648, -1}
    v7 := GenS1{F0: 134}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (("hello" + "\n\t") + "foo bar")
}
func genFn1(p0 uint32, p1 uint) uint32 {
    v8 := []int{-338, 160, -53}
    p0 = 65535
    var v9 int32 = 305
    v10 := GenS1{F0: 42}
    var v11 uint32 = (p0 & (p0 - (p0 >> ((p1) & 63))))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return (^(v11))
}

func main() {
    var v1 int8 = -55
    var v2 float32 = -2.25
    v3 := []uint32{}
    v4 := []string{("" + "\n\t"), "a", "foo bar", "foo bar"}
    var v5 int8 = v1
    v1 = (v5 - 2)
    v1 -= -106
    v2 = 1e10
    v12 := GenS2{F0: 32767, F1: 2}
    v13 := GenS3{F0: ("\n\t" + (("\n\t" + "a") + ""))}
    var v14 string = ((("" + "\n\t") + "\n\t") + "\n\t")
    v3 = append(v3, genFn1(2147483647, 0))
    if false {
        var v15 bool = true
        v1 = v5
        v16 := GenS4{F0: 1, F1: 1}
        _ = v15
        _ = v16
    } else {
        if false {
            if (!(2 == 938)) {
                var v17 int = 831
                v18 := GenS1{F0: 1}
                if false {
                    switch v17 {
                    case 32767:
                        fmt.Println(v17)
                    case 1:
                        switch "" {
                        case "a":
                            var v19 uint = 65535
                            _ = v19
                        case "foo bar":
                            v2 -= ((-2.25 - (3.14159 + v2)) * ((1e10 * v2) / 1.5))
                        case "":
                            v20 := GenS2{F0: -1, F1: 72}
                            _ = v20
                        }
                    default:
                        v21 := GenS1{F0: -1}
                        _ = v21
                    }
                }
                _ = v17
                _ = v18
            } else {
                for fi22 := 0; fi22 < 6; fi22++ {
                    if (266 <= 0) {
                        if ((fi22 / (((fi22 & 127)) | 1)) <= fi22) {
                            var v23 uint = 0
                            _ = v23
                        } else {
                            v14 = v14
                            for fi24 := 0; fi24 < 4; fi24++ {
                                v25 := GenS5{F0: (fi24 | (fi24 << ((2) & 63))), F1: 0, F2: 936}
                                _ = fi24
                                _ = v25
                            }
                            v3 = append(v3, 769)
                        }
                        v4 = append(v4, v14)
                    }
                    v26 := GenS3{F0: "a"}
                    _ = fi22
                    _ = v26
                }
            }
            v14 = v14
        } else {
            var v27 int8 = (^(v5))
            v14 = "\n\t"
            _ = v27
        }
        if (1 >= 392) {
            if (2147483647 != 339) {
                fmt.Println(v1)
                if true {
                    v14 = (v14 + (v14 + "hello"))
                    v28 := GenS2{F0: 437, F1: 2}
                    _ = v28
                }
                if true {
                    if (!((float64((v2 / (float32(1e10) / v2))) < float64(v2)) && false)) {
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                    } else {
                        var v29 float64 = 0.0
                        if (41 < 1) {
                            for fi30 := 0; fi30 < 18; fi30++ {
                                v31 := GenS4{F0: 9223372036854775807, F1: 2}
                                _ = fi30
                                _ = v31
                            }
                        } else {
                            fmt.Println(v1, v2)
                            v1 = 42
                            var v32 uint8 = 127
                            _ = v32
                        }
                        if true {
                            switch (v14 + v14) {
                            case "a":
                                v33 := GenS5{F0: 4294967295, F1: 2, F2: 4294967295}
                                _ = v33
                            case "hello":
                                if false {
                                    var v34 int16 = 770
                                    _ = v34
                                }
                            case "foo bar":
                                v3 = append(v3, 4294967295)
                            }
                            for fi35 := 0; fi35 < 4; fi35++ {
                                for fi36 := 0; fi36 < 3; fi36++ {
                                    if false {
                                        if true {
                                            v37 := GenS5{F0: fi36, F1: 470, F2: 657}
                                            _ = v37
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v4 = append(v4, v14)
                                        }
                                        v38 := GenS1{F0: -1}
                                        if (9223372036854775807 > 1) {
                                            fmt.Println(v1, v2, v29)
                                            v5 = (v1 * v5)
                                        } else {
                                            if (!((127 > v1) && (!(true || ("a" != (v14 + (("foo bar" + "a") + ""))))))) {
                                                v39 := GenS1{F0: -9223372036854775808}
                                                _ = v39
                                            }
                                            if true {
                                                var v40 uint = 4294967295
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                switch fi36 {
                                                case 4294967295:
                                                    var v41 float32 = v2
                                                    _ = v41
                                                case 32767:
                                                    switch (fi35 * (-(fi35))) {
                                                    case 32768:
                                                        var v42 uint64 = 1
                                                        _ = v42
                                                    case 4294967295:
                                                        v3 = append(v3, 0)
                                                    }
                                                default:
                                                    if true {
                                                        if ((false && false) || (true || true)) {
                                                            var v43 uint = v40
                                                            v44 := GenS5{F0: fi36, F1: 42, F2: 4294967295}
                                                            switch v14 {
                                                            case "foo bar":
                                                                v45 := GenS6{F0: 30, F1: 262, F2: 42}
                                                                _ = v45
                                                            default:
                                                                if (!false) {
                                                                    if (!true) {
                                                                        v2 -= float32(v29)
                                                                        if (413 < 1) {
                                                                            switch v14 {
                                                                            case "foo bar":
                                                                                v3 = append(v3, 65535)
                                                                            case "\n\t":
                                                                                v29 = float64(((1e10 + float32(3.14159)) - (v2 + (v2 - v2))))
                                                                            }
                                                                        }
                                                                    }
                                                                    switch (v14 + v14) {
                                                                    case "\n\t":
                                                                        if (fi35 < (-(fi36))) {
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            var v46 string = (v14 + v14)
                                                                            v29 = 0.1
                                                                            v47 := GenS3{F0: v46}
                                                                            _ = v46
                                                                            _ = v47
                                                                        } else {
                                                                            v2 += (v2 * 1e10)
                                                                            v48 := GenS7{F0: 9223372036854775807}
                                                                            if true {
                                                                                v1 = v5
                                                                                fmt.Println(fi35)
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = 693
                                                                                }
                                                                            } else {
                                                                                v49 := GenS1{F0: int64(fi35)}
                                                                                v14 = v14
                                                                                v40 = 1
                                                                                _ = v49
                                                                            }
                                                                            _ = v48
                                                                        }
                                                                    default:
                                                                        v50 := GenS3{F0: v14}
                                                                        _ = v50
                                                                    }
                                                                }
                                                            }
                                                            _ = v43
                                                            _ = v44
                                                        } else {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            switch "\n\t" {
                                                            case "":
                                                                switch v14 {
                                                                case "a":
                                                                    if (-2.25 >= v29) {
                                                                        var v51 float64 = v29
                                                                        v5 = -1
                                                                        _ = len(v3)
                                                                        _ = len(v3)
                                                                        v40 -= 4294967295
                                                                        switch int(v40) {
                                                                        case 0:
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                        }
                                                                        _ = v51
                                                                    } else {
                                                                        v52 := GenS1{F0: -1}
                                                                        v53 := GenS8{F0: v40, F1: (fi35 * (fi36 & fi35))}
                                                                        if (3.14159 > 3.14159) {
                                                                            var v54 int16 = 32767
                                                                            v29 += v29
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                var v55 float32 = v2
                                                                                _ = v55
                                                                            }
                                                                            _ = v54
                                                                        }
                                                                        _ = v52
                                                                        _ = v53
                                                                    }
                                                                default:
                                                                    if (v14 <= "foo bar") {
                                                                        if (346 > 18446744073709551615) {
                                                                            switch v14 {
                                                                            case "\n\t":
                                                                                v56 := GenS4{F0: 1, F1: 426}
                                                                                _ = v56
                                                                            case "hello":
                                                                                fmt.Println(fi35, v2)
                                                                            case "a":
                                                                                _ = len(v4)
                                                                                fmt.Println(v5)
                                                                            }
                                                                            switch (((v14 + v14) + v14) + "") {
                                                                            case "foo bar":
                                                                                v57 := GenS3{F0: (v14 + "")}
                                                                                _ = v57
                                                                            case "a":
                                                                                var v58 int = 2147483647
                                                                                _ = v58
                                                                            case "hello":
                                                                                fmt.Println(v5)
                                                                            }
                                                                        } else {
                                                                            var v59 int32 = -2147483648
                                                                            v60 := GenS2{F0: 32767, F1: 2}
                                                                            fmt.Println(v29, v59)
                                                                            _ = v59
                                                                            _ = v60
                                                                        }
                                                                        v61 := GenS8{F0: (v40 << ((v40) & 63)), F1: (fi36 % (((^(fi35))) | 1))}
                                                                        switch v14 {
                                                                        case "hello":
                                                                            if (v2 <= v2) {
                                                                                v2 += 0.1
                                                                            } else {
                                                                                v62 := GenS2{F0: 32767, F1: 352}
                                                                                fmt.Println(fi36, v1, fi36)
                                                                                v63 := GenS4{F0: 0, F1: 4294967295}
                                                                                _ = v62
                                                                                _ = v63
                                                                            }
                                                                            fmt.Println(v1, fi35, v5)
                                                                            fmt.Println(fi36)
                                                                        default:
                                                                            fmt.Println(v1, v40)
                                                                        }
                                                                        if (706 != 42) {
                                                                            if len(v3) > 0 {
                                                                                v3[0] = 1
                                                                            }
                                                                            if true {
                                                                                v64 := GenS1{F0: -691}
                                                                                _ = v64
                                                                            }
                                                                            var v65 uint64 = 18446744073709551615
                                                                            var v66 int32 = -170
                                                                            _ = v65
                                                                            _ = v66
                                                                        } else {
                                                                            fmt.Println(v14, v40, v40)
                                                                            v67 := GenS1{F0: 0}
                                                                            _ = v67
                                                                        }
                                                                        v3 = append(v3, genFn1(2, (v40 << (((^(v40))) & 63))))
                                                                        fmt.Println(v1, v5)
                                                                        _ = v61
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                v68 := GenS3{F0: ((v14 + (v14 + v14)) + "")}
                                                                fmt.Println(fi35, v40, fi35)
                                                                if true {
                                                                    switch ("foo bar" + "\n\t") {
                                                                    case "hello":
                                                                        v29 += ((((v29 * v29) + 0.1) - (v29 * v29)) - v29)
                                                                    case "a":
                                                                        v2 += ((0.0 * (v2 + (float32(-2.25) * float32((v29 / float64((v2 - v2))))))) + v2)
                                                                    default:
                                                                        var v69 int32 = 1
                                                                        _ = v69
                                                                    }
                                                                    _ = len(v4)
                                                                    var v70 float32 = float32(v29)
                                                                    _ = v70
                                                                } else {
                                                                    if true {
                                                                        v3 = append(v3, 65535)
                                                                        v1 = (-(v1))
                                                                        if true {
                                                                            var v71 string = "foo bar"
                                                                            _ = v71
                                                                        }
                                                                        v29 += v29
                                                                    } else {
                                                                        v14 = v14
                                                                    }
                                                                    v1 -= int8(v40)
                                                                    switch v14 {
                                                                    case "foo bar":
                                                                        v2 = (v2 - (-2.25 - v2))
                                                                    default:
                                                                        v40 += v40
                                                                    }
                                                                }
                                                                _ = v68
                                                            case "\n\t":
                                                                v3 = append(v3, 42)
                                                            default:
                                                                v72 := GenS9{F0: 487, F1: -2147483648}
                                                                if (v40 < v40) {
                                                                    fmt.Println(fi35, v40)
                                                                } else {
                                                                    fmt.Println(fi36, v5, fi35)
                                                                }
                                                                fmt.Println(v2)
                                                                _ = v72
                                                            }
                                                        }
                                                        var v73 int64 = -691
                                                        _ = v73
                                                    }
                                                }
                                                _ = v40
                                            } else {
                                                if ((fi36 % ((9223372036854775807) | 1)) <= fi35) {
                                                    v74 := GenS3{F0: v14}
                                                    fmt.Println(fi36, v29, fi36)
                                                    if len(v3) > 0 {
                                                        v3[0] = 42
                                                    }
                                                    _ = v74
                                                }
                                            }
                                        }
                                        v29 = v29
                                        _ = v38
                                    } else {
                                        fmt.Println(v14)
                                        var v75 uint = 490
                                        _ = v75
                                    }
                                    v5 += v1
                                    _ = fi36
                                }
                                v76 := GenS10{F0: (v1 - (^(v1))), F1: (fi35 >> ((65535) & 63)), F2: 0}
                                _ = fi35
                                _ = v76
                            }
                            var v77 uint64 = 534
                            _ = v77
                        } else {
                            if len(v4) > 0 {
                                v4[0] = v14
                            }
                            v14 = v14
                            switch 9223372036854775807 {
                            case 0:
                                v78 := GenS6{F0: -1, F1: 2, F2: 811}
                                _ = v78
                            default:
                                v79 := GenS7{F0: -1}
                                _ = v79
                            }
                        }
                        _ = v29
                    }
                }
            } else {
                fmt.Println(v1, v2, v14)
            }
            v80 := GenS11{F0: (v5 % (((v1 & -1)) | 1)), F1: (^(v1))}
            if false {
                v14 = v14
                switch 122 {
                case 887:
                    var v81 uint8 = 42
                    _ = v81
                case 876:
                    var v82 int8 = (v1 % (((v1 - (v1 & -1))) | 1))
                    _ = v82
                }
            }
            _ = v80
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v83 := GenS12{F0: 1, F1: 1, F2: 42}
    var v84 uint64 = 18446744073709551615
    if true {
        switch 208 {
        case 9223372036854775807:
            for fi85 := 0; fi85 < 20; fi85++ {
                var v86 int16 = 32767
                if (-832 == 1) {
                    v1 = v5
                }
                fmt.Println(v86, v1, v5)
                _ = fi85
                _ = v86
            }
        case 1:
            switch ((v14 + v14) + "\n\t") {
            case "foo bar":
                v84 = v84
            case "":
                v2 = v2
            }
        case 128:
            v14 = (v14 + "foo bar")
        }
        var v87 int = 127
        v87 = v87
        _ = v87
    }
    if true {
        v88 := GenS2{F0: 42, F1: (v84 - 18446744073709551615)}
        if true {
            var v89 int64 = 42
            _ = v89
        } else {
            for fi90 := 0; fi90 < 6; fi90++ {
                v91 := GenS13{F0: false, F1: 2147483647, F2: 1}
                fmt.Println(v14, v1, v1)
                if (-53 == -566) {
                    var v92 int16 = -32768
                    v14 = "hello"
                    _ = len(v3)
                    _ = v92
                }
                _ = fi90
                _ = v91
            }
        }
        _ = v88
    }
    v93 := GenS8{F0: 65535, F1: 0}
    var v94 bool = (!(false || true))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v84:", v84)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v84
    _ = v93
    _ = v94
}
