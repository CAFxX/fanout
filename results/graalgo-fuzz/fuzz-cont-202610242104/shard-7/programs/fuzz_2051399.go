package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 uint32; F1 int; F2 uint64 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int }
func genFn0(p0 uint32, p1 uint64) uint16 {
    fmt.Println(p0)
    switch ("foo bar" + "foo bar") {
    case "":
        var v5 int64 = 9223372036854775807
        _ = v5
    case "\n\t":
        fmt.Println(p0, p1)
    case "hello":
        var v6 int16 = 32767
        _ = v6
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 int64) uint32 {
    var v7 string = "\n\t"
    p0 += p0
    v8 := []int{}
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(fi9, v7)
        var v10 string = v7
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (-783 < 0)
    v3 := make(map[int]int)
    v4 := []string{(("\n\t" + "foo bar") + ""), ("a" + "hello"), "", "a"}
    v2 = false
    if (v2 || v2) {
        for fi11 := 0; fi11 < 4; fi11++ {
            if ((3.14159 * 0.0) == -2.25) {
                var v12 int32 = 0
                var v13 string = ("a" + "\n\t")
                var v14 float64 = (3.14159 - 0.0)
                _ = v12
                _ = v13
                _ = v14
            } else {
                v2 = (42 > 457)
            }
            if v2 {
                switch (fi11 >> ((756) & 63)) {
                case -9223372036854775808:
                    switch "\n\t" {
                    case "":
                        switch "hello" {
                        case "hello":
                            var v15 int64 = -1
                            _ = v15
                        case "a":
                            var v16 int64 = 37
                            _ = v16
                        case "":
                            switch "\n\t" {
                            case "\n\t":
                                v2 = (v2 && (494 >= genFn1(-49)))
                            }
                        default:
                            for fi17 := 0; fi17 < 9; fi17++ {
                                var v18 uint32 = 2147483648
                                v18 = (v18 / ((genFn1(0)) | 1))
                                _ = fi17
                                _ = v18
                            }
                        }
                    case "hello":
                        for fi19 := 0; fi19 < 15; fi19++ {
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                            v2 = ((true || (618 == 42)) || v2)
                            var v20 uint32 = genFn1(9223372036854775807)
                            _ = fi19
                            _ = v20
                        }
                    }
                case 128:
                    switch fi11 {
                    case 2147483647:
                        switch "\n\t" {
                        case "a":
                            for fi21 := 0; fi21 < 7; fi21++ {
                                v22 := GenS1{F0: 2, F1: 498, F2: 42}
                                var v23 uint = 65535
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        }
                    case 766:
                        v24 := GenS2{F0: 4294967295, F1: (fi11 ^ fi11), F2: 976}
                        _ = v24
                    case -1:
                        var v25 uint64 = 1
                        _ = v25
                    }
                default:
                    v2 = ((!v2) || (2147483647 == 0))
                }
            } else {
                for fi26 := 0; fi26 < 7; fi26++ {
                    v4 = append(v4, (("\n\t" + "") + "\n\t"))
                    if (!v2) {
                        v2 = (!false)
                        v2 = v2
                        var v27 int32 = 1
                        _ = v27
                    } else {
                        var v28 float64 = (float64(float32(1.5)) / 1e10)
                        if (-495 == -1) {
                            var v29 int16 = -32768
                            _ = len(v3)
                            _ = v29
                        }
                        _ = v28
                    }
                    _ = fi26
                }
            }
            for fi30 := 0; fi30 < 17; fi30++ {
                v2 = (fi30 > fi11)
                if (true || false) {
                    fmt.Println(v2, fi30, fi30)
                    fmt.Println(fi11, v2, fi30)
                } else {
                    var v31 uint = 1
                    var v32 int16 = 42
                    switch "\n\t" {
                    case "a":
                        if v2 {
                            v32 += v32
                            v33 := v3[-6]
                            _ = v33
                        } else {
                            if v2 {
                                _ = len(v4)
                                switch "\n\t" {
                                case "":
                                    v1["hello"] = 2
                                    var v34 uint8 = 1
                                    _ = v34
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                v35 := GenS2{F0: 2147483647, F1: fi11, F2: 18446744073709551615}
                                _ = v35
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    default:
                        switch "" {
                        case "":
                            fmt.Println(fi11, v32)
                        default:
                            v31 = (v31 ^ v31)
                        }
                    }
                    _ = v31
                    _ = v32
                }
                _ = fi30
            }
            _ = fi11
        }
        v2 = v2
    } else {
        var v36 uint32 = 1
        var v37 int8 = 42
        v2 = (2 >= 255)
        _ = v36
        _ = v37
    }
    var v38 int32 = -1
    var v39 uint8 = 21
    for fi40 := 0; fi40 < 15; fi40++ {
        var v41 int8 = -128
        _ = fi40
        _ = v41
    }
    if v2 {
        for fi42 := 0; fi42 < 16; fi42++ {
            fmt.Println(v38, v39, fi42)
            _ = fi42
        }
        var v43 float64 = (1e10 + float64((float32((-2.25 + (0.0 + -2.25))) + 1e10)))
        _ = v43
    }
    v44 := GenS1{F0: 4294967295, F1: v38, F2: 42}
    for fi45 := 0; fi45 < 2; fi45++ {
        switch (fi45 << ((18446744073709551615) & 63)) {
        case -1:
            for fi46 := 0; fi46 < 16; fi46++ {
                var v47 int = 127
                fmt.Println(v2, v47)
                if (false || true) {
                    if v2 {
                        fmt.Println(fi46, fi46, v2)
                        v48 := GenS1{F0: 695, F1: (v38 & -1), F2: int16(fi46)}
                        if v2 {
                            switch fi45 {
                            case 1:
                                if ((1e10 * 1.5) >= 0.1) {
                                    v38 -= (v38 - v38)
                                    v39 = (v39 % (((v39 & v39)) | 1))
                                }
                            case 42:
                                var v49 int8 = 42
                                _ = v49
                            case -589:
                                switch "hello" {
                                case "hello":
                                    v39 -= 194
                                case "a":
                                    var v50 uint8 = v39
                                    _ = v50
                                default:
                                    v51 := GenS1{F0: 698, F1: v38, F2: -846}
                                    _ = v51
                                }
                            }
                            v38 -= v38
                            var v52 string = "foo bar"
                            _ = v52
                        }
                        v39 += (^(v39))
                        if v2 {
                            if v2 {
                                var v53 int32 = v38
                                v2 = v2
                                _ = v53
                            } else {
                                var v54 bool = (!(float32((1e10 / 1e10)) != 0.1))
                                _ = v54
                            }
                            switch (v47 / (((v47 / ((1) | 1))) | 1)) {
                            case 42:
                                var v55 uint16 = 42
                                _ = v55
                            }
                            v47 = (fi45 * (fi46 / ((-1) | 1)))
                        } else {
                            v47 += fi45
                            if (!(!(381 == 0))) {
                                v47 += (fi45 % ((fi46) | 1))
                                v56 := GenS1{F0: 42, F1: -328, F2: -446}
                                var v57 int16 = 1
                                _ = v56
                                _ = v57
                            } else {
                                switch "foo bar" {
                                case "\n\t":
                                    var v58 int = v47
                                    _ = v58
                                case "foo bar":
                                    var v59 uint8 = v39
                                    _ = v59
                                case "hello":
                                    v2 = v2
                                }
                                var v60 uint32 = 2147483648
                                var v61 int8 = 27
                                var v62 float32 = float32(float64((0.0 * (float32(0.1) / 3.14159))))
                                switch 513 {
                                case 713:
                                    if true {
                                        v3[fi46] = (v47 >> ((255) & 63))
                                        _ = len(v3)
                                    } else {
                                        fmt.Println(v61, v61)
                                        if true {
                                            v2 = (true || v2)
                                        }
                                    }
                                case -1:
                                    if (!v2) {
                                        if false {
                                            v62 += ((3.14159 / float32(1e10)) / (v62 + (float32(3.14159) + ((((v62 + v62) - (-2.25 + ((1.5 / (0.0 - (v62 * (0.0 / ((v62 / v62) + (v62 / (v62 + (v62 - 1.5)))))))) + float32((0.1 + 0.0))))) + (0.0 + v62)) + 3.14159))))
                                            switch "\n\t" {
                                            case "foo bar":
                                                _ = len(v4)
                                            default:
                                                var v63 float64 = 0.1
                                                _ = v63
                                            }
                                            var v64 uint = 2
                                            _ = v64
                                        } else {
                                            v38 -= v38
                                            v65 := GenS3{F0: v61, F1: 1}
                                            if v2 {
                                                if (250 != 65535) {
                                                    v2 = true
                                                    var v66 int8 = (v61 & v61)
                                                    _ = v66
                                                }
                                                _ = len(v4)
                                                switch "" {
                                                case "\n\t":
                                                    fmt.Println(v2)
                                                case "":
                                                    v67 := GenS3{F0: v61, F1: 182}
                                                    _ = v67
                                                case "hello":
                                                    v47 -= fi45
                                                default:
                                                    switch fi46 {
                                                    case 713:
                                                        if v2 {
                                                            var v68 uint8 = (v39 - 42)
                                                            _ = v68
                                                        } else {
                                                            var v69 int = (fi45 / (((^(fi45))) | 1))
                                                            v60 += v60
                                                            if v2 {
                                                                var v70 int16 = 32767
                                                                fmt.Println(v47)
                                                                delete(v3, 127)
                                                                v71 := GenS3{F0: 100, F1: v39}
                                                                fmt.Println(v60, v62)
                                                                _ = v70
                                                                _ = v71
                                                            } else {
                                                                var v72 int64 = 1
                                                                fmt.Println(fi46, fi46, v69)
                                                                var v73 int16 = 32767
                                                                var v74 uint8 = v39
                                                                var v75 bool = v2
                                                                var v76 int32 = v38
                                                                _ = v72
                                                                _ = v73
                                                                _ = v74
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                            _ = len(v1)
                                                            fmt.Println(v69, v39, fi46)
                                                            _ = v69
                                                        }
                                                    case 9223372036854775807:
                                                        var v77 int16 = 1
                                                        _ = v77
                                                    }
                                                }
                                            } else {
                                                if (v2 || ((v61 << ((18446744073709551615) & 63)) >= v61)) {
                                                    fmt.Println(v62)
                                                    var v78 string = ""
                                                    if len(v4) > 0 {
                                                        v4[0] = "hello"
                                                    }
                                                    _ = v78
                                                }
                                            }
                                            _ = v65
                                        }
                                        if (v61 < v61) {
                                            var v79 bool = v2
                                            var v80 int32 = v38
                                            if (((("hello" + "") + "") + "\n\t") > ("\n\t" + "\n\t")) {
                                                v81 := v3[2147483647]
                                                v82 := GenS4{F0: ("a" + "foo bar"), F1: 4294967295}
                                                var v83 int32 = (v80 % (((v80 ^ v38)) | 1))
                                                _ = v81
                                                _ = v82
                                                _ = v83
                                            } else {
                                                v47 -= fi46
                                            }
                                            var v84 string = ("a" + ("\n\t" + ("hello" + "")))
                                            var v85 string = "\n\t"
                                            _ = v79
                                            _ = v80
                                            _ = v84
                                            _ = v85
                                        }
                                    }
                                }
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                        }
                        _ = v48
                    }
                    fmt.Println(fi46, fi45, fi45)
                    if (fi45 < fi45) {
                        var v86 string = "a"
                        var v87 int = v47
                        _ = v86
                        _ = v87
                    } else {
                        v2 = (1 <= v38)
                        fmt.Println(v39, v47, v38)
                    }
                    if (!v2) {
                        v47 += fi46
                        if v2 {
                            switch (-(fi46)) {
                            case -1:
                                v88 := GenS5{F0: (fi45 * v47)}
                                _ = v88
                            case 4294967295:
                                v38 = v38
                            case 984:
                                var v89 string = "hello"
                                _ = v89
                            default:
                                v90 := v1["foo bar"]
                                _ = v90
                            }
                        }
                        v47 += (-(fi45))
                    }
                    if (v2 || v2) {
                        v39 += v39
                        v91 := GenS4{F0: (("hello" + "foo bar") + ""), F1: 730}
                        if true {
                            v92 := v1["hello"]
                            if (1000 < 1) {
                                _ = len(v1)
                                if (!(!true)) {
                                    v47 += 4294967295
                                    delete(v3, -975)
                                    switch (-(fi46)) {
                                    case -437:
                                        if true {
                                            var v93 string = "foo bar"
                                            var v94 float64 = (0.0 / -2.25)
                                            _ = v93
                                            _ = v94
                                        } else {
                                            v92 = (v92 | v92)
                                            delete(v3, (v47 / ((32767) | 1)))
                                        }
                                    default:
                                        if v2 {
                                            fmt.Println(v2)
                                            v38 = (-(v38))
                                        }
                                    }
                                }
                                fmt.Println(v2, v92)
                            }
                            v39 -= (v39 << ((255) & 63))
                            _ = v92
                        }
                        _ = v91
                    }
                    v3[(fi46 ^ fi46)] = int(v39)
                    _ = len(v1)
                }
                _ = fi46
                _ = v47
            }
        default:
            v95 := v1["foo bar"]
            _ = v95
        }
        _ = fi45
    }
    v96 := GenS3{F0: -1, F1: (v39 | v39)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v44
    _ = v96
}
