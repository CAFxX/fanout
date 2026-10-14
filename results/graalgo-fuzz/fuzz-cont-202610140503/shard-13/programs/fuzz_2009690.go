package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 string; F2 uint64 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int64; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int; F1 bool }
type GenS8 struct { F0 bool; F1 uint16; F2 int32 }
type GenS9 struct { F0 int; F1 uint32 }
type GenS10 struct { F0 uint64; F1 int32; F2 int64 }
type GenS11 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = -797
    v3 := []uint32{65535, 42, 65535}
    var v4 uint32 = 42
    var v5 int = int(v4)
    var v6 uint32 = (v4 << ((42) & 63))
    v7 := GenS1{F0: false}
    var v8 uint64 = 42
    var v9 string = "foo bar"
    v10 := GenS1{F0: true}
    switch (v9 + "\n\t") {
    case "hello":
        if (v6 != (v6 + v4)) {
            _ = len(v1)
            var v11 uint16 = 58
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 string = v9
                _ = fi12
                _ = v13
            }
            _ = v11
        }
    case "foo bar":
        switch -275 {
        case 128:
            var v14 bool = true
            _ = v14
        case 273:
            var v15 int8 = 0
            _ = v15
        case 32767:
            if false {
                for fi16 := 0; fi16 < 12; fi16++ {
                    switch fi16 {
                    case 0:
                        var v17 uint = 584
                        _ = v17
                    case 9223372036854775807:
                        for fi18 := 0; fi18 < 10; fi18++ {
                            if (1 <= -699) {
                                v5 = -1
                            }
                            if false {
                                if true {
                                    switch (v2 / ((v5) | 1)) {
                                    case 2147483647:
                                        v5 = fi18
                                    case -933:
                                        switch v9 {
                                        case "hello":
                                            if false {
                                                if true {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if true {
                                                            v19 := GenS1{F0: false}
                                                            _ = v19
                                                        } else {
                                                            v6 = uint32(v8)
                                                        }
                                                    case "":
                                                        switch ((v9 + v9) + v9) {
                                                        case "":
                                                            switch (fi18 | (fi16 ^ 1)) {
                                                            case 82:
                                                                var v20 int16 = -32768
                                                                _ = v20
                                                            case 476:
                                                                var v21 bool = (4294967295 < 196)
                                                                v22 := GenS2{F0: 2, F1: "hello", F2: (v8 % ((2) | 1))}
                                                                if false {
                                                                    v5 += (fi16 & v5)
                                                                    switch v2 {
                                                                    case 327:
                                                                        v4 = v4
                                                                    case -94:
                                                                        switch (((v9 + "\n\t") + ("hello" + v9)) + (v9 + v9)) {
                                                                        case "":
                                                                            var v23 uint16 = 65535
                                                                            _ = v23
                                                                        default:
                                                                            var v24 int8 = 100
                                                                            _ = v24
                                                                        }
                                                                    default:
                                                                        if (!(!true)) {
                                                                            var v25 uint64 = (v8 ^ (v8 + 2))
                                                                            v26 := GenS1{F0: (0 >= v25)}
                                                                            _ = v25
                                                                            _ = v26
                                                                        } else {
                                                                            var v27 int64 = 1
                                                                            var v28 bool = (-567 < -338)
                                                                            var v29 float64 = 0.1
                                                                            _ = v27
                                                                            _ = v28
                                                                            _ = v29
                                                                        }
                                                                    }
                                                                    var v30 uint8 = 127
                                                                    _ = v30
                                                                }
                                                                _ = v21
                                                                _ = v22
                                                            }
                                                        default:
                                                            fmt.Println(v6, v2, v5)
                                                        }
                                                    case "foo bar":
                                                        v3 = append(v3, (v6 / (((v6 / (((v6 >> ((2) & 63))) | 1))) | 1)))
                                                    default:
                                                        if (!((0 > 0) && (!false))) {
                                                            v5 += fi18
                                                            v31 := GenS1{F0: (229 <= 2)}
                                                            if true {
                                                                v1[v4] = (v9 + "foo bar")
                                                                var v32 int8 = 42
                                                                _ = v32
                                                            } else {
                                                                if len(v3) > 0 {
                                                                    v3[0] = 0
                                                                }
                                                                v5 += v2
                                                                var v33 uint32 = (v4 | v6)
                                                                _ = v33
                                                            }
                                                            switch (fi18 * (fi18 + (v5 + (fi18 | (fi18 / (((fi16 - v5)) | 1)))))) {
                                                            case 4294967295:
                                                                v34 := GenS1{F0: true}
                                                                _ = v34
                                                            case 9223372036854775807:
                                                                var v35 string = ((v9 + "hello") + "hello")
                                                                _ = v35
                                                            case -2147483648:
                                                                var v36 float64 = 1e10
                                                                v8 -= v8
                                                                _ = v36
                                                            default:
                                                                var v37 int64 = 1
                                                                _ = v37
                                                            }
                                                            fmt.Println(v5, v8)
                                                            _ = v31
                                                        }
                                                    }
                                                    var v38 int = -2147483648
                                                    _ = v38
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                if (false || true) {
                                                    if false {
                                                        var v39 float32 = 0.0
                                                        if (false && false) {
                                                            if (v4 <= v6) {
                                                                if (!(!true)) {
                                                                    if true {
                                                                        var v40 uint8 = 49
                                                                        _ = v40
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        v41 := GenS3{F0: (!((65535 > 2) || true)), F1: ((251 <= 1) && (true && true))}
                                                                        v2 -= 569
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        _ = v41
                                                                    } else {
                                                                        var v42 uint64 = (v8 * (v8 ^ (v8 + v8)))
                                                                        v1[v4] = v9
                                                                        _ = v42
                                                                    }
                                                                    fmt.Println(fi18)
                                                                    if true {
                                                                        var v43 uint64 = v8
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        _ = len(v3)
                                                                        _ = v43
                                                                    } else {
                                                                        fmt.Println(v8, v4, v5)
                                                                        _ = len(v3)
                                                                    }
                                                                }
                                                                v2 += fi16
                                                            } else {
                                                                v1[666] = "a"
                                                                if false {
                                                                    var v44 int16 = int16(v6)
                                                                    _ = v44
                                                                } else {
                                                                    var v45 float64 = 3.14159
                                                                    var v46 int = v5
                                                                    fmt.Println(v5)
                                                                    _ = v45
                                                                    _ = v46
                                                                }
                                                            }
                                                        }
                                                        v2 += (v2 - -837)
                                                        _ = v39
                                                    }
                                                    v47 := GenS4{F0: 238, F1: 276, F2: (v6 + (v6 + 0))}
                                                    _ = v47
                                                } else {
                                                    v48 := GenS2{F0: 29, F1: ("a" + "foo bar"), F2: (v8 / ((v8) | 1))}
                                                    switch (fi18 ^ (-(fi18))) {
                                                    case -590:
                                                        fmt.Println(fi16, fi16, v5)
                                                    case 42:
                                                        fmt.Println(fi18)
                                                    case -9223372036854775808:
                                                        v49 := GenS5{F0: 904}
                                                        _ = v49
                                                    }
                                                    fmt.Println(v6, v4, fi16)
                                                    _ = v48
                                                }
                                                fmt.Println(v2)
                                            } else {
                                                if false {
                                                    v50 := GenS1{F0: false}
                                                    var v51 int16 = 1
                                                    fmt.Println(fi16, fi18, fi18)
                                                    var v52 float32 = (float32(((0.1 + float64(-2.25)) + (-2.25 / 1e10))) + ((1.5 + 1e10) - 0.0))
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    _ = len(v3)
                                                    v53 := GenS5{F0: 376}
                                                    _ = v53
                                                }
                                                v54 := v1[v4]
                                                v5 = fi18
                                                v55 := GenS1{F0: true}
                                                _ = len(v1)
                                                var v56 uint = 65535
                                                switch 648 {
                                                case 1:
                                                    var v57 bool = true
                                                    _ = v57
                                                case -9223372036854775808:
                                                    v58 := v1[(v6 ^ 1)]
                                                    _ = v58
                                                case 128:
                                                    if (1 >= 2) {
                                                        delete(v1, v6)
                                                        v59 := GenS3{F0: false, F1: (false || true)}
                                                        _ = v59
                                                    }
                                                }
                                                v3 = append(v3, 224)
                                                _ = v54
                                                _ = v55
                                                _ = v56
                                            }
                                        case "":
                                            if (0.1 >= (-2.25 * 1e10)) {
                                                if true {
                                                    v6 += 247
                                                }
                                                v60 := GenS6{F0: 42}
                                                v2 = fi18
                                                _ = v60
                                            } else {
                                                switch (fi18 + v2) {
                                                case 751:
                                                    var v61 uint16 = 2
                                                    _ = v61
                                                case 9223372036854775807:
                                                    delete(v1, (v4 + (^(v4))))
                                                default:
                                                    if false {
                                                        if false {
                                                            _ = len(v1)
                                                            v2 += v2
                                                            switch -13 {
                                                            case -9223372036854775808:
                                                                v9 = "a"
                                                            case 32768:
                                                                fmt.Println(v5, v2)
                                                            case 32767:
                                                                v5 = v2
                                                            }
                                                        } else {
                                                            v62 := GenS3{F0: (uint16(v6) > 2), F1: false}
                                                            if false {
                                                                v8 = (v8 * v8)
                                                                var v63 int8 = -92
                                                                var v64 uint32 = 65535
                                                                _ = v63
                                                                _ = v64
                                                            } else {
                                                                v65 := v1[v4]
                                                                _ = v65
                                                            }
                                                            switch v9 {
                                                            case "foo bar":
                                                                var v66 uint8 = 196
                                                                _ = v66
                                                            case "hello":
                                                                fmt.Println(v4)
                                                            case "a":
                                                                v6 -= v6
                                                                switch v9 {
                                                                case "a":
                                                                    if false {
                                                                        v2 += (v2 - (-(v2)))
                                                                        v5 -= (v5 & (fi16 << ((4294967295) & 63)))
                                                                    }
                                                                    var v67 uint8 = 22
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    _ = v67
                                                                case "\n\t":
                                                                    var v68 string = v9
                                                                    _ = v68
                                                                case "hello":
                                                                    var v69 uint16 = 319
                                                                    _ = v69
                                                                }
                                                                v2 = (^(v2))
                                                            default:
                                                                v70 := GenS7{F0: (v5 * (^(v2))), F1: true}
                                                                _ = v70
                                                            }
                                                            _ = v62
                                                        }
                                                        v6 = (v4 / (((v6 ^ v4)) | 1))
                                                    }
                                                }
                                                _ = len(v3)
                                            }
                                        default:
                                            v71 := GenS1{F0: true}
                                            _ = v71
                                        }
                                    }
                                    if (v8 <= 18446744073709551615) {
                                        var v72 int64 = 9223372036854775807
                                        v8 = (^(v8))
                                        fmt.Println(v72)
                                        _ = v72
                                    }
                                } else {
                                    v1[(v4 + (^(v4)))] = (v9 + v9)
                                    var v73 float64 = (0.1 / 1.5)
                                    var v74 int64 = 1
                                    v73 = v73
                                    v75 := GenS8{F0: (!(!((v9 == ("a" + (v9 + ((v9 + v9) + "hello")))) || true))), F1: 0, F2: 42}
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                }
                                var v76 uint = uint(v6)
                                v77 := GenS7{F0: fi18, F1: false}
                                _ = v76
                                _ = v77
                            } else {
                                switch "foo bar" {
                                case "foo bar":
                                    var v78 uint8 = 42
                                    _ = v78
                                default:
                                    var v79 int32 = 42
                                    _ = v79
                                }
                            }
                            var v80 bool = true
                            _ = fi18
                            _ = v80
                        }
                    case 32767:
                        switch "a" {
                        case "foo bar":
                            var v81 uint32 = 1
                            _ = v81
                        case "":
                            fmt.Println(v8)
                        case "hello":
                            var v82 float32 = ((3.14159 / -2.25) * 0.1)
                            _ = v82
                        default:
                            for fi83 := 0; fi83 < 3; fi83++ {
                                switch -968 {
                                case -1:
                                    var v84 int16 = -133
                                    _ = v84
                                }
                                _ = fi83
                            }
                        }
                    default:
                        for fi85 := 0; fi85 < 3; fi85++ {
                            v5 = (^(v2))
                            _ = len(v3)
                            v86 := GenS9{F0: (v5 / ((fi16) | 1)), F1: v4}
                            var v87 uint32 = v6
                            if len(v3) > 0 {
                                v3[0] = 910
                            }
                            _ = fi85
                            _ = v86
                            _ = v87
                        }
                    }
                    v88 := GenS4{F0: -962, F1: 18446744073709551615, F2: v6}
                    v1[v4] = v9
                    _ = fi16
                    _ = v88
                }
                v89 := GenS8{F0: false, F1: 709, F2: -1}
                _ = v89
            } else {
                v90 := v1[(v4 >> ((4294967295) & 63))]
                _ = v90
            }
        }
    }
    v5 += (v2 * v5)
    for fi91 := 0; fi91 < 15; fi91++ {
        for fi92 := 0; fi92 < 18; fi92++ {
            var v93 string = v9
            var v94 int = fi91
            _ = fi92
            _ = v93
            _ = v94
        }
        _ = fi91
    }
    v95 := GenS10{F0: (v8 % (((v8 - 1)) | 1)), F1: 42, F2: -512}
    if (!false) {
        v9 = "hello"
        var v96 int16 = -1
        var v97 int = (-(v2))
        _ = v96
        _ = v97
    } else {
        v5 = 4294967295
        if false {
            fmt.Println(v6, v4, v8)
        } else {
            if len(v3) > 0 {
                v3[0] = v4
            }
            switch v5 {
            case 192:
                var v98 int64 = 398
                _ = v98
            case -762:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                switch ((("foo bar" + ("foo bar" + v9)) + v9) + v9) {
                case "foo bar":
                    var v99 uint32 = 42
                    _ = v99
                case "hello":
                    if (float64(0.0) > 0.1) {
                        v6 = 2147483647
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (2 == 255) {
                            var v100 uint = 845
                            v9 = (v9 + ("" + v9))
                            _ = v100
                        }
                    } else {
                        v5 = -1
                        v6 += (^(v4))
                    }
                default:
                    fmt.Println(v6, v9)
                }
            }
        }
    }
    fmt.Println(v2, v5)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v6, v4, v6)
    v101 := GenS11{F0: 1}
    v3 = append(v3, 65535)
    v6 -= v4
    if (9223372036854775807 <= -1) {
        if ((((v9 + (("hello" + v9) + v9)) + (v9 + "hello")) + (("hello" + v9) + "hello")) != (v9 + v9)) {
            fmt.Println(v4)
        }
        v1[v4] = v9
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    _ = v95
}
