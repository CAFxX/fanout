package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 uint32; F1 int8; F2 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32; F1 bool }
type GenS9 struct { F0 uint8; F1 bool }
func genFn0(p0 bool) string {
    for fi4 := 0; fi4 < 19; fi4++ {
        fmt.Println(p0, fi4)
        _ = fi4
    }
    v5 := GenS1{F0: 2}
    var v6 uint8 = 255
    var v7 uint32 = 880
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int16 = 1000
        if p0 {
            for fi10 := 0; fi10 < 1; fi10++ {
                fmt.Println(v7)
                v11 := []int{}
                if false {
                    if len(v11) > 0 {
                        v11[0] = -1
                    }
                } else {
                    fmt.Println(v9)
                    v9 += 991
                    fmt.Println(p0, p0)
                    fmt.Println(v6, v7)
                }
                _ = len(v11)
                _ = fi10
                _ = v11
            }
            if false {
                var v12 string = (("\n\t" + ("" + "a")) + "foo bar")
                switch "" {
                case "foo bar":
                    if (127 < (-(fi8))) {
                        fmt.Println(v12, fi8)
                    } else {
                        v6 = v6
                        v13 := GenS1{F0: 32767}
                        switch ("hello" + ("foo bar" + (((v12 + v12) + (v12 + "foo bar")) + v12))) {
                        case "foo bar":
                            var v14 uint32 = (v7 ^ (^(v7)))
                            _ = v14
                        }
                        _ = v13
                    }
                }
                if (253 != 464) {
                    if (!p0) {
                        v7 -= v7
                        var v15 float64 = ((1.5 / 1e10) / 0.1)
                        var v16 float64 = v15
                        _ = v15
                        _ = v16
                    }
                }
                _ = v12
            } else {
                if p0 {
                    v7 -= (v7 & v7)
                    v17 := []int{fi8}
                    _ = v17
                } else {
                    v7 = 2
                }
            }
        } else {
            for fi18 := 0; fi18 < 17; fi18++ {
                if (-596 < 911) {
                    v7 = v7
                } else {
                    switch 127 {
                    case 602:
                        var v19 uint8 = v6
                        _ = v19
                    case 952:
                        v20 := GenS1{F0: 2}
                        _ = v20
                    }
                    fmt.Println(v7, v7)
                    if (22 < (v6 >> ((65535) & 63))) {
                        fmt.Println(p0, v9)
                        v9 -= v9
                        fmt.Println(p0)
                        var v21 uint8 = v6
                        if p0 {
                            v22 := GenS1{F0: 65535}
                            if p0 {
                                p0 = false
                            } else {
                                var v23 int = -502
                                _ = v23
                            }
                            _ = v22
                        }
                        _ = v21
                    } else {
                        if p0 {
                            v24 := GenS2{F0: ("\n\t" + "\n\t")}
                            switch "\n\t" {
                            case "hello":
                                var v25 int64 = -9223372036854775808
                                _ = v25
                            case "\n\t":
                                v6 = (v6 | v6)
                            default:
                                fmt.Println(v9)
                            }
                            _ = v24
                        } else {
                            var v26 string = ""
                            _ = v26
                        }
                        var v27 string = "\n\t"
                        if (102 < 127) {
                            var v28 int16 = (v9 * (v9 - 42))
                            var v29 uint32 = v7
                            _ = v28
                            _ = v29
                        } else {
                            var v30 int64 = -636
                            var v31 string = (v27 + v27)
                            var v32 int = fi18
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        _ = v27
                    }
                }
                fmt.Println(v7)
                _ = fi18
            }
            var v33 float32 = -2.25
            if p0 {
                if (v6 < (v6 / ((v6) | 1))) {
                    switch "hello" {
                    case "\n\t":
                        switch -338 {
                        case 651:
                            p0 = true
                        case -247:
                            v34 := []int{(fi8 >> ((uint(v6)) & 63)), (fi8 / ((32768) | 1)), 42}
                            _ = v34
                        case 1:
                            var v35 uint16 = 772
                            _ = v35
                        default:
                            if ((v6 / ((255) | 1)) < v6) {
                                v36 := []int{(fi8 % ((-204) | 1)), fi8, (-(fi8)), -9223372036854775808}
                                switch fi8 {
                                case 671:
                                    v37 := GenS3{F0: ("" + "hello"), F1: fi8}
                                    _ = v37
                                default:
                                    _ = len(v36)
                                }
                                switch -1 {
                                case 42:
                                    var v38 int = 1
                                    _ = v38
                                case -55:
                                    fmt.Println(v9)
                                case 4294967295:
                                    v39 := GenS1{F0: 2}
                                    _ = v39
                                default:
                                    p0 = (0.1 == (0.0 - float32(((((-2.25 * 0.1) - -2.25) / 1.5) - 0.0))))
                                }
                                _ = v36
                            } else {
                                switch ("\n\t" + "\n\t") {
                                case "":
                                    p0 = true
                                case "hello":
                                    if p0 {
                                        v40 := []int{2147483647, fi8, -2147483648}
                                        _ = v40
                                    }
                                case "\n\t":
                                    v41 := GenS3{F0: "", F1: fi8}
                                    _ = v41
                                default:
                                    switch fi8 {
                                    case 9223372036854775807:
                                        for fi42 := 0; fi42 < 15; fi42++ {
                                            var v43 int64 = -9223372036854775808
                                            v44 := GenS3{F0: ("a" + "hello"), F1: fi42}
                                            v7 = v7
                                            _ = fi42
                                            _ = v43
                                            _ = v44
                                        }
                                    case -1:
                                        var v45 int8 = 127
                                        _ = v45
                                    }
                                }
                                var v46 float32 = float32(3.14159)
                                v9 = v9
                                _ = v46
                            }
                        }
                    case "a":
                        var v47 string = "hello"
                        _ = v47
                    case "foo bar":
                        var v48 int32 = -156
                        _ = v48
                    default:
                        var v49 uint64 = 4294967295
                        _ = v49
                    }
                    v50 := []int{}
                    if ((!p0) && p0) {
                        p0 = p0
                        v7 -= v7
                        var v51 uint32 = (v7 ^ v7)
                        _ = v51
                    } else {
                        if (65535 <= 255) {
                            v52 := GenS3{F0: ("" + "a"), F1: fi8}
                            v33 = (1.5 + v33)
                            _ = v52
                        }
                        var v53 int32 = 153
                        if p0 {
                            var v54 uint64 = 2
                            if p0 {
                                var v55 uint16 = 65535
                                _ = v55
                            } else {
                                var v56 float32 = v33
                                if (float64((((v33 + 1e10) * 0.1) + 1.5)) >= 0.0) {
                                    v6 -= (v6 % ((v6) | 1))
                                    v57 := GenS2{F0: "hello"}
                                    for fi58 := 0; fi58 < 15; fi58++ {
                                        if false {
                                            if (1 != 9223372036854775807) {
                                                var v59 float32 = ((v56 + v56) / v33)
                                                v56 -= v59
                                                v60 := GenS2{F0: "a"}
                                                _ = v59
                                                _ = v60
                                            } else {
                                                if p0 {
                                                    if ("\n\t" >= "") {
                                                        if len(v50) > 0 {
                                                            v50[0] = fi8
                                                        }
                                                        _ = len(v50)
                                                        v6 = (v6 / (((v6 % (((v6 * 255)) | 1))) | 1))
                                                    } else {
                                                        var v61 uint32 = v7
                                                        fmt.Println(v56, fi58, v7)
                                                        var v62 int8 = -1
                                                        if (!(v56 > (v56 * (v33 / 0.1)))) {
                                                            v9 -= int16(v6)
                                                        }
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                } else {
                                                    switch "a" {
                                                    case "\n\t":
                                                        v50 = append(v50, fi8)
                                                        switch (fi8 | fi58) {
                                                        case 919:
                                                            v63 := GenS4{F0: 1000}
                                                            _ = v63
                                                        case 0:
                                                            switch fi8 {
                                                            case 32767:
                                                                p0 = false
                                                            case 777:
                                                                fmt.Println(v56)
                                                            default:
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v6 -= v6
                                                                }
                                                            }
                                                        }
                                                    case "hello":
                                                        var v64 int16 = (^(v9))
                                                        _ = v64
                                                    case "":
                                                        fmt.Println(v56)
                                                    default:
                                                        v65 := GenS4{F0: (v9 + v9)}
                                                        _ = v65
                                                    }
                                                    switch fi8 {
                                                    case -923:
                                                        p0 = p0
                                                    case 1:
                                                        var v66 string = "a"
                                                        _ = v66
                                                    default:
                                                        var v67 uint32 = 396
                                                        _ = v67
                                                    }
                                                    v7 = v7
                                                }
                                                v50 = append(v50, (^(fi58)))
                                                v50 = append(v50, fi58)
                                            }
                                        } else {
                                            var v68 uint8 = (^(v6))
                                            _ = v68
                                        }
                                        v6 += (^(v6))
                                        v33 -= v33
                                        fmt.Println(v53, v9, v54)
                                        _ = fi58
                                    }
                                    _ = v57
                                }
                                fmt.Println(v54)
                                _ = v56
                            }
                            v7 = v7
                            _ = v54
                        } else {
                            var v69 uint = 0
                            switch fi8 {
                            case 128:
                                fmt.Println(v9, p0)
                            default:
                                v70 := GenS4{F0: (v9 / ((v9) | 1))}
                                _ = v70
                            }
                            v6 += (v6 + 128)
                            _ = v69
                        }
                        _ = v53
                    }
                    _ = v50
                } else {
                    v7 = (v7 & (v7 * v7))
                    v71 := []int{fi8}
                    var v72 int32 = 2147483647
                    _ = v71
                    _ = v72
                }
                v33 = float32(3.14159)
            } else {
                v73 := []int{(fi8 & (fi8 >> ((238) & 63)))}
                switch "a" {
                case "a":
                    v7 = 2147483648
                }
                var v74 int32 = 0
                _ = v73
                _ = v74
            }
            _ = v33
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return ("hello" + (("" + (("" + "\n\t") + "")) + ((("foo bar" + "hello") + ("\n\t" + "hello")) + "")))
}
func genFn1(p0 int32, p1 string) uint32 {
    p0 = 2147483647
    var v75 float64 = float64(1.5)
    p1 = p1
    fmt.Println(p1, p1, v75)
    _ = p0
    _ = p1
    _ = v75
    return 905
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 117
    var v3 uint32 = 2
    switch "hello" {
    case "a":
        var v76 int8 = 42
        _ = v76
    case "":
        var v77 int64 = 9223372036854775807
        _ = v77
    }
    v3 = v3
    var v78 uint16 = (^(v2))
    var v79 uint32 = v3
    for fi80 := 0; fi80 < 3; fi80++ {
        v81 := GenS4{F0: -32768}
        v3 = v79
        _ = fi80
        _ = v81
    }
    if false {
        v78 += 32767
        v79 += 2
        v79 -= 4294967295
    } else {
        if false {
            switch "\n\t" {
            case "\n\t":
                var v82 uint = 4294967295
                _ = v82
            }
            fmt.Println(v79, v3)
        }
    }
    var v83 uint = 354
    var v84 float32 = (0.1 / -2.25)
    v79 += (v3 & 42)
    for fi85 := 0; fi85 < 3; fi85++ {
        if true {
            v86 := GenS5{F0: -1, F1: int64(v2)}
            for fi87 := 0; fi87 < 16; fi87++ {
                switch int(v83) {
                case -9223372036854775808:
                    v88 := GenS4{F0: 1000}
                    _ = v88
                default:
                    var v89 uint64 = 18446744073709551615
                    _ = v89
                }
                v90 := []int{(fi85 & fi87), 9223372036854775807, fi87, 0}
                _ = fi87
                _ = v90
            }
            var v91 float64 = -2.25
            _ = v86
            _ = v91
        } else {
            var v92 bool = (!(v2 >= 449))
            _ = v92
        }
        if (128 == 2) {
            v93 := []int{}
            v94 := GenS6{F0: v79, F1: -9, F2: (v84 < float32(1e10))}
            v95 := v1["a"]
            _ = v93
            _ = v94
            _ = v95
        } else {
            for fi96 := 0; fi96 < 12; fi96++ {
                delete(v1, "")
                _ = fi96
            }
        }
        if true {
            var v97 string = ""
            v98 := GenS7{F0: v83}
            _ = v97
            _ = v98
        } else {
            v84 += (3.14159 * ((v84 - 0.0) * (1.5 / float32(-2.25))))
            for fi99 := 0; fi99 < 15; fi99++ {
                var v100 int8 = -109
                _ = fi99
                _ = v100
            }
        }
        _ = fi85
    }
    if (9223372036854775807 >= -9223372036854775808) {
        _ = len(v1)
    } else {
        v1["a"] = (v3 * (v79 ^ v3))
        if false {
            if (255 != 127) {
                if (((("a" + "") + "hello") + "foo bar") >= ("foo bar" + "a")) {
                    switch "a" {
                    case "a":
                        v101 := GenS8{F0: v3, F1: (9223372036854775807 > -252)}
                        _ = v101
                    }
                }
            } else {
                switch "hello" {
                case "foo bar":
                    if false {
                        var v102 string = ""
                        _ = v102
                    } else {
                        v3 += (v3 % ((65535) | 1))
                    }
                case "a":
                    v3 = genFn1(42, ("" + (("foo bar" + "foo bar") + "")))
                default:
                    var v103 string = ""
                    _ = v103
                }
                v104 := GenS4{F0: -102}
                var v105 uint16 = (v78 >> ((687) & 63))
                _ = v104
                _ = v105
            }
            v106 := GenS9{F0: 0, F1: true}
            _ = v106
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v78
    _ = v79
    _ = v83
    _ = v84
}
