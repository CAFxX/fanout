package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int64; F1 int8; F2 uint64 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 int16, p1 bool) uint8 {
    var v5 string = (((("a" + (("hello" + "") + "hello")) + "") + "a") + ("a" + "\n\t"))
    var v6 float64 = (1e10 + 0.1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 uint8, p1 string) uint {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int64 = 515
        _ = fi7
        _ = v8
    }
    v9 := []int{}
    p0 = p0
    for fi10 := 0; fi10 < 3; fi10++ {
        p0 = 127
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    v11 := GenS1{F0: -996}
    var v12 uint8 = genFn0(0, false)
    if false {
        if ((42 == v1) || (-9223372036854775808 > 205)) {
            var v13 int64 = -756
            v12 += v12
            var v14 uint32 = 141
            _ = v13
            _ = v14
        } else {
            var v15 string = ("hello" + "")
            v12 = v12
            fmt.Println(v15, v12, v12)
            _ = v15
        }
        v12 -= v12
        fmt.Println(v12, v1)
    } else {
        var v16 uint64 = 1
        _ = v16
    }
    switch "a" {
    case "\n\t":
        delete(v4, "foo bar")
    case "foo bar":
        v1 = v1
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        v1 = -750
        for fi18 := 0; fi18 < 20; fi18++ {
            v1 = (v1 & -32768)
            v19 := GenS1{F0: 750}
            if false {
                v1 -= int16(v12)
                v20 := []int{(-(fi18))}
                _ = v20
            }
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    v21 := []int{4294967295, 9223372036854775807}
    if true {
        switch "\n\t" {
        case "":
            for fi22 := 0; fi22 < 4; fi22++ {
                _ = len(v21)
                _ = fi22
            }
        case "foo bar":
            var v23 uint = 255
            _ = v23
        }
        fmt.Println(v1, v1)
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 int32 = 0
            fmt.Println(fi24)
            _ = fi24
            _ = v25
        }
    }
    if false {
        var v26 float64 = -2.25
        _ = v26
    } else {
        var v27 uint = 42
        var v28 float64 = (float64(-2.25) * (0.0 / float64(0.1)))
        switch "\n\t" {
        case "":
            var v29 uint16 = 32767
            _ = v29
        case "hello":
            v28 = 1e10
        case "foo bar":
            var v30 int16 = v1
            _ = v30
        }
        _ = v27
        _ = v28
    }
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        switch "hello" {
        case "foo bar":
            var v32 int32 = -248
            _ = v32
        case "hello":
            switch fi31 {
            case -1:
                for fi33 := 0; fi33 < 3; fi33++ {
                    var v34 int64 = 1
                    if len(v21) > 0 {
                        v21[0] = (fi31 & (-(fi33)))
                    }
                    var v35 uint16 = 195
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            case -817:
                if (float64(-2.25) <= 0.1) {
                    for fi36 := 0; fi36 < 10; fi36++ {
                        var v37 uint8 = (v12 % (((v12 ^ v12)) | 1))
                        _ = fi36
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 6; fi38++ {
                        v21 = append(v21, 4294967295)
                        switch "\n\t" {
                        case "hello":
                            v39 := GenS1{F0: -1}
                            _ = v39
                        }
                        v12 = (v12 & (v12 | v12))
                        if false {
                            fmt.Println(fi31, fi38, v12)
                            var v40 string = "a"
                            var v41 uint16 = 609
                            v42 := GenS1{F0: (v1 / (((v1 ^ v1)) | 1))}
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                        switch (("hello" + "foo bar") + ("a" + "hello")) {
                        case "foo bar":
                            v12 += (v12 % ((128) | 1))
                        case "a":
                            if ((true || (2 < 65535)) || true) {
                                if false {
                                    v43 := GenS2{F0: v12}
                                    _ = v43
                                } else {
                                    v44 := GenS1{F0: v1}
                                    v12 = (^(v12))
                                    _ = v44
                                }
                                fmt.Println(fi38, v1, fi31)
                                fmt.Println(fi38)
                            } else {
                                v12 += (v12 * genFn0(v1, true))
                                switch 32768 {
                                case 9223372036854775807:
                                    v1 += (^(v1))
                                case 127:
                                    if false {
                                        fmt.Println(v12, v12)
                                        var v45 uint = 531
                                        var v46 uint16 = 1
                                        var v47 string = ("foo bar" + ("a" + "foo bar"))
                                        delete(v4, v47)
                                        fmt.Println(v1)
                                        if true {
                                            v1 = v1
                                            v1 += 1000
                                            switch (-(fi38)) {
                                            case 1:
                                                fmt.Println(v47)
                                            case 0:
                                                fmt.Println(v12, fi31, fi31)
                                            default:
                                                switch (("hello" + (v47 + v47)) + v47) {
                                                case "":
                                                    v46 += (^(v46))
                                                case "\n\t":
                                                    v47 = ""
                                                case "foo bar":
                                                    v1 -= v1
                                                default:
                                                    if ((((v47 + v47) + "a") + "a") <= v47) {
                                                        var v48 int = fi38
                                                        v48 -= (fi31 ^ (v48 + (v48 | (fi38 - -688))))
                                                        _ = v48
                                                    }
                                                }
                                            }
                                            var v49 int64 = 1
                                            v21 = append(v21, (fi38 >> ((genFn1(v12, ((v47 + v47) + "\n\t"))) & 63)))
                                            _ = v49
                                        }
                                        var v50 uint16 = (v46 | 824)
                                        switch (v47 + (((("" + ((v47 + ((v47 + "foo bar") + "\n\t")) + "hello")) + v47) + v47) + v47)) {
                                        case "\n\t":
                                            v51 := GenS1{F0: (v1 - 1)}
                                            _ = v51
                                        case "a":
                                            var v52 string = "\n\t"
                                            _ = v52
                                        default:
                                            v46 = v50
                                        }
                                        v53 := GenS3{F0: 128, F1: v12}
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                        _ = v50
                                        _ = v53
                                    } else {
                                        v12 += v12
                                        var v54 float64 = float64(float32(float64(((1.5 * 1e10) - (0.0 / (3.14159 / (0.0 - -2.25)))))))
                                        if (18446744073709551615 >= 18446744073709551615) {
                                            if false {
                                                fmt.Println(fi38)
                                            }
                                        } else {
                                            if ((v1 % ((v1) | 1)) >= 32767) {
                                                if false {
                                                    if (false || false) {
                                                        var v55 uint32 = 42
                                                        _ = len(v21)
                                                        var v56 uint8 = (^(v12))
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    if false {
                                                        fmt.Println(v54, v12)
                                                        if ((v1 << ((18446744073709551615) & 63)) >= (v1 | v1)) {
                                                            var v57 int64 = -247
                                                            fmt.Println(v12, fi38, v12)
                                                            v58 := GenS3{F0: fi31, F1: 2}
                                                            fmt.Println(v54, v54)
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                        v12 = (v12 / (((v12 & v12)) | 1))
                                                    }
                                                    v59 := GenS1{F0: v1}
                                                    _ = v59
                                                } else {
                                                    var v60 bool = (42 >= -2147483648)
                                                    if (2147483647 >= 790) {
                                                        var v61 float32 = 0.1
                                                        _ = len(v21)
                                                        _ = v61
                                                    } else {
                                                        switch 42 {
                                                        case 4294967295:
                                                            v1 = (v1 | v1)
                                                        case 128:
                                                            var v62 bool = (-740 < -55)
                                                            _ = v62
                                                        }
                                                        v63 := GenS4{F0: 42, F1: 0, F2: 285}
                                                        var v64 int8 = -128
                                                        var v65 bool = v60
                                                        _ = v63
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    if false {
                                                        var v66 bool = v60
                                                        var v67 int64 = -1
                                                        var v68 uint64 = uint64(v12)
                                                        _ = v66
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                    _ = v60
                                                }
                                                var v69 int32 = -928
                                                fmt.Println(v12)
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v70 string = "foo bar"
                                                    _ = v70
                                                default:
                                                    v21 = append(v21, fi31)
                                                }
                                                v54 = v54
                                                var v71 int = (fi31 >> ((606) & 63))
                                                _ = v69
                                                _ = v71
                                            } else {
                                                if (false || ("\n\t" != (("a" + "foo bar") + "\n\t"))) {
                                                    var v72 int64 = -858
                                                    _ = v72
                                                }
                                            }
                                            var v73 int64 = 0
                                            _ = v73
                                        }
                                        _ = v54
                                    }
                                case 2147483647:
                                    v74 := GenS2{F0: (v12 * genFn0((v1 * 42), false))}
                                    _ = v74
                                default:
                                    switch fi31 {
                                    case 42:
                                        v12 = (v12 % ((genFn0(-431, true)) | 1))
                                    default:
                                        if false {
                                            var v75 int16 = (v1 % ((v1) | 1))
                                            _ = v75
                                        } else {
                                            var v76 uint32 = 2147483648
                                            _ = v76
                                        }
                                    }
                                }
                                v12 = v12
                            }
                        case "":
                            v12 -= (v12 + v12)
                        }
                        _ = fi38
                    }
                } else {
                    var v77 int64 = 228
                    v78 := GenS5{F0: -322, F1: -9223372036854775808}
                    if (!false) {
                        var v79 int64 = (v77 >> ((42) & 63))
                        for fi80 := 0; fi80 < 10; fi80++ {
                            v81 := GenS2{F0: (^(v12))}
                            _ = fi80
                            _ = v81
                        }
                        for fi82 := 0; fi82 < 2; fi82++ {
                            if ((!true) || false) {
                                v1 -= (v1 ^ v1)
                                switch (fi82 % ((fi31) | 1)) {
                                case -2147483648:
                                    v77 -= (^(v79))
                                case 42:
                                    switch (fi82 >> ((4294967295) & 63)) {
                                    case 980:
                                        if (!(-92 <= 42)) {
                                            switch fi82 {
                                            case 248:
                                                fmt.Println(fi31, fi31)
                                            case 888:
                                                v83 := v2[-1]
                                                _ = v83
                                            case -943:
                                                v79 = (v77 & v79)
                                            }
                                            var v84 bool = false
                                            _ = v84
                                        }
                                    case 32767:
                                        var v85 float64 = 3.14159
                                        _ = v85
                                    default:
                                        var v86 uint64 = 283
                                        _ = v86
                                    }
                                case 2147483647:
                                    var v87 int16 = (v1 % ((v1) | 1))
                                    _ = v87
                                default:
                                    v88 := GenS4{F0: (^(v79)), F1: 0, F2: 4294967295}
                                    _ = v88
                                }
                                if false {
                                    v1 += 0
                                } else {
                                    var v89 int16 = v1
                                    _ = v89
                                }
                            } else {
                                v1 += 32767
                                fmt.Println(v79, v1, v12)
                                switch (fi31 + fi82) {
                                case 42:
                                    var v90 uint8 = genFn0(v1, true)
                                    _ = v90
                                default:
                                    v91 := GenS4{F0: v77, F1: -1, F2: 0}
                                    _ = v91
                                }
                                if true {
                                    v79 = (v79 / ((v79) | 1))
                                    v92 := GenS2{F0: (v12 + v12)}
                                    _ = v92
                                }
                            }
                            v1 = v1
                            var v93 float64 = float64(1.5)
                            _ = fi82
                            _ = v93
                        }
                        _ = v79
                    }
                    _ = v77
                    _ = v78
                }
            }
        }
        var v94 int32 = int32(fi31)
        v95 := GenS6{F0: 1}
        _ = fi31
        _ = v94
        _ = v95
    }
    v12 = (^(v12))
    v12 = v12
    v12 = v12
    if (180 < genFn0(v1, false)) {
        v1 -= int16(v12)
    }
    var v96 int = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v96:", v96)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v96
}
