package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16; F1 string; F2 int32 }
type GenS8 struct { F0 uint; F1 int8; F2 int8 }
type GenS9 struct { F0 uint64; F1 int64; F2 uint }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint8; F1 int8 }
type GenS12 struct { F0 int; F1 int32; F2 int32 }
type GenS13 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint64) uint {
    v6 := []int{711, 2147483647, 4294967295}
    v7 := GenS1{F0: 0}
    var v8 int64 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{42, 0}
    var v2 string = ""
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    v5 := []string{v2, "foo bar", ("" + v2)}
    var v9 string = v2
    v10 := GenS1{F0: 165}
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            v9 = v9
            var v13 uint64 = 18446744073709551615
            v13 -= (v13 - v13)
            var v14 int = 9223372036854775807
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := GenS2{F0: -464}
        _ = fi11
        _ = v15
    }
    delete(v4, 2147483647)
    var v16 uint = 18446744073709551615
    v9 = ""
    for fi17 := 0; fi17 < 10; fi17++ {
        if (false && (true && (v2 != v9))) {
            v18 := GenS2{F0: -9223372036854775808}
            for fi19 := 0; fi19 < 12; fi19++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi19
            }
            var v20 uint = 65535
            _ = v18
            _ = v20
        } else {
            if (!false) {
                v2 = (v9 + v2)
                v16 -= (^(v16))
            } else {
                fmt.Println(v2)
                if ((fi17 + (^(fi17))) >= fi17) {
                    if false {
                        var v21 float32 = 1.5
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 4; fi22++ {
                        var v23 int32 = -2147483648
                        var v24 int16 = -1
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    var v25 uint = genFn0(1, 4294967295)
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        switch v9 {
                        case "foo bar":
                            v4[4294967295] = (fi17 + (fi26 << ((v16) & 63)))
                        case "":
                            if false {
                                if (!(fi26 >= (fi26 + -419))) {
                                    var v27 int8 = 42
                                    v28 := GenS2{F0: 1}
                                    _ = v27
                                    _ = v28
                                }
                                v16 -= (v16 + (v16 >> ((v16) & 63)))
                            }
                        default:
                            if true {
                                fmt.Println(fi26)
                            } else {
                                switch fi17 {
                                case 42:
                                    if true {
                                        var v29 int64 = -561
                                        _ = v29
                                    }
                                case -9223372036854775808:
                                    if (!true) {
                                        var v30 int8 = 0
                                        var v31 uint64 = 0
                                        v32 := GenS3{F0: fi17}
                                        var v33 int8 = v30
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = v9
                                        }
                                    }
                                }
                                if false {
                                    if ((820 != v16) || true) {
                                        v34 := GenS1{F0: 32767}
                                        v2 = (v2 + "foo bar")
                                        _ = v34
                                    } else {
                                        switch fi26 {
                                        case 32768:
                                            switch fi17 {
                                            case -9223372036854775808:
                                                switch fi17 {
                                                case 4294967295:
                                                    var v35 float64 = 1.5
                                                    _ = v35
                                                case -768:
                                                    fmt.Println(v16)
                                                case 329:
                                                    if ((v9 + (v9 + v9)) != v9) {
                                                        if false {
                                                            if (!false) {
                                                                var v36 string = ((v2 + ((("foo bar" + v2) + "hello") + (("" + v9) + ((v2 + v9) + v2)))) + ((v9 + v2) + v9))
                                                                if (0 >= 1) {
                                                                    var v37 float64 = 1e10
                                                                    v38 := GenS4{F0: false, F1: 2}
                                                                    switch fi17 {
                                                                    case 0:
                                                                        if (((1 < 128) || true) && false) {
                                                                            v39 := GenS5{F0: (!(-218 <= -1))}
                                                                            var v40 int64 = 42
                                                                            var v41 float32 = float32(v37)
                                                                            _ = v39
                                                                            _ = v40
                                                                            _ = v41
                                                                        }
                                                                    case 128:
                                                                        fmt.Println(v37)
                                                                    default:
                                                                        var v42 uint32 = 42
                                                                        _ = v42
                                                                    }
                                                                    v16 -= genFn0(42, 1)
                                                                    _ = v37
                                                                    _ = v38
                                                                } else {
                                                                    if (false || (0 == 1)) {
                                                                        if len(v5) > 0 {
                                                                            v5[0] = "\n\t"
                                                                        }
                                                                        switch fi26 {
                                                                        case 32768:
                                                                            switch v36 {
                                                                            case "":
                                                                                if (!(!(false || true))) {
                                                                                    v2 = v2
                                                                                }
                                                                                switch ((v36 + (v36 + ((((v2 + "\n\t") + v9) + v9) + "hello"))) + v2) {
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        fmt.Println(v16, v36)
                                                                                    }
                                                                                case "a":
                                                                                    v2 = "\n\t"
                                                                                }
                                                                            case "foo bar":
                                                                                v9 = ((v2 + "hello") + "a")
                                                                                v43 := GenS6{F0: 937}
                                                                                var v44 uint16 = 337
                                                                                _ = v43
                                                                                _ = v44
                                                                            case "hello":
                                                                                if true {
                                                                                    v45 := GenS7{F0: 1, F1: v9, F2: 42}
                                                                                    switch 692 {
                                                                                    case 42:
                                                                                        v36 = "a"
                                                                                    default:
                                                                                        v16 += v16
                                                                                    }
                                                                                    var v46 float32 = 0.0
                                                                                    _ = v45
                                                                                    _ = v46
                                                                                } else {
                                                                                    v16 += v16
                                                                                    v9 = v9
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v47 int8 = 42
                                                                            _ = v47
                                                                        }
                                                                        v9 = "hello"
                                                                    } else {
                                                                        var v48 int64 = 9223372036854775807
                                                                        _ = v48
                                                                    }
                                                                    fmt.Println(v16, v36, v2)
                                                                    v2 = v36
                                                                }
                                                                _ = v36
                                                            }
                                                            v16 = genFn0(198, 18446744073709551615)
                                                            v49 := GenS5{F0: (18446744073709551615 == 915)}
                                                            _ = v49
                                                        } else {
                                                            fmt.Println(fi17)
                                                            if ((1.5 / (1e10 * ((1.5 + 0.0) / 1.0))) == (0.1 * 0.0)) {
                                                                switch (fi26 + 42) {
                                                                case -1:
                                                                    fmt.Println(fi26, v9, fi26)
                                                                }
                                                                if (631 != fi26) {
                                                                    v9 = "hello"
                                                                }
                                                            }
                                                        }
                                                        v9 = v9
                                                        _ = len(v3)
                                                    }
                                                }
                                            case 42:
                                                var v50 float32 = ((1e10 + (-2.25 * -2.25)) / (1.5 * float32(-2.25)))
                                                _ = v50
                                            case -429:
                                                _ = len(v1)
                                                if false {
                                                    switch v9 {
                                                    case "":
                                                        v51 := GenS1{F0: 806}
                                                        var v52 uint = genFn0(938, 42)
                                                        v53 := GenS8{F0: (^(v16)), F1: -1, F2: 1}
                                                        v2 = v9
                                                        _ = v51
                                                        _ = v52
                                                        _ = v53
                                                    case "hello":
                                                        v16 -= v16
                                                        v16 -= v16
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case "foo bar":
                                                        v1 = append(v1, 1)
                                                    default:
                                                        if false {
                                                            if ((-871 > 122) && true) {
                                                                var v54 float64 = 1.5
                                                                if (861 >= 42) {
                                                                    v54 = v54
                                                                    if (127 == 1) {
                                                                        v4[65535] = (^(fi26))
                                                                    }
                                                                    v54 += ((0.0 + (0.0 - 3.14159)) * v54)
                                                                    switch (^(fi26)) {
                                                                    case -86:
                                                                        if (!(1e10 != v54)) {
                                                                            _ = len(v5)
                                                                        } else {
                                                                            v16 += (^(v16))
                                                                        }
                                                                    case -9223372036854775808:
                                                                        if (-1 == 32767) {
                                                                            if true {
                                                                                if true {
                                                                                    fmt.Println(v16, fi26, v16)
                                                                                    var v55 bool = true
                                                                                    _ = v55
                                                                                } else {
                                                                                    var v56 uint64 = 382
                                                                                    _ = v56
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v57 := GenS9{F0: 2, F1: 1, F2: (^(v16))}
                                                                            var v58 bool = false
                                                                            var v59 int16 = 0
                                                                            _ = v57
                                                                            _ = v58
                                                                            _ = v59
                                                                        }
                                                                    case 127:
                                                                        switch (fi26 & 777) {
                                                                        case -9223372036854775808:
                                                                            var v60 int64 = -1
                                                                            _ = v60
                                                                        default:
                                                                            v16 = v16
                                                                        }
                                                                    default:
                                                                        if (false || true) {
                                                                            switch (-(fi26)) {
                                                                            case -2147483648:
                                                                                v61 := GenS10{F0: 0}
                                                                                _ = v61
                                                                            case -358:
                                                                                var v62 uint32 = 2147483648
                                                                                _ = v62
                                                                            case -9223372036854775808:
                                                                                var v63 uint64 = 0
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = ((v9 + v2) + "hello")
                                                                                }
                                                                                _ = v63
                                                                            default:
                                                                                v54 = v54
                                                                            }
                                                                            var v64 uint32 = uint32(v16)
                                                                            _ = v64
                                                                        }
                                                                    }
                                                                    if len(v5) > 0 {
                                                                        v5[0] = "a"
                                                                    }
                                                                } else {
                                                                    v54 += v54
                                                                    fmt.Println(fi17)
                                                                    fmt.Println(v9, v9, v54)
                                                                }
                                                                if true {
                                                                    var v65 uint8 = 1
                                                                    var v66 uint8 = v65
                                                                    _ = len(v4)
                                                                    _ = v65
                                                                    _ = v66
                                                                } else {
                                                                    v54 = v54
                                                                    v67 := GenS3{F0: fi17}
                                                                    _ = v67
                                                                }
                                                                _ = v54
                                                            } else {
                                                                v68 := GenS7{F0: 255, F1: v9, F2: 42}
                                                                switch (-(fi26)) {
                                                                case 9223372036854775807:
                                                                    if true {
                                                                        var v69 uint16 = 325
                                                                        var v70 int = (^(fi26))
                                                                        v71 := GenS11{F0: 111, F1: 0}
                                                                        _ = v69
                                                                        _ = v70
                                                                        _ = v71
                                                                    } else {
                                                                        v72 := v3[(v2 + v2)]
                                                                        fmt.Println(v16)
                                                                        if false {
                                                                            var v73 int32 = -353
                                                                            v74 := GenS8{F0: (v16 - v16), F1: 50, F2: -40}
                                                                            _ = v73
                                                                            _ = v74
                                                                        } else {
                                                                            var v75 float32 = -2.25
                                                                            var v76 int16 = 0
                                                                            var v77 string = v2
                                                                            _ = v75
                                                                            _ = v76
                                                                            _ = v77
                                                                        }
                                                                        fmt.Println(v9, v2, fi17)
                                                                        _ = v72
                                                                    }
                                                                case 32767:
                                                                    v3["a"] = (-(fi17))
                                                                case 448:
                                                                    var v78 int8 = 127
                                                                    _ = v78
                                                                }
                                                                var v79 uint16 = 255
                                                                _ = v68
                                                                _ = v79
                                                            }
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "":
                                                        var v80 float32 = 1.5
                                                        _ = v80
                                                    case "\n\t":
                                                        var v81 int = fi17
                                                        switch 558 {
                                                        case 32768:
                                                            switch v9 {
                                                            case "":
                                                                v82 := GenS6{F0: 4294967295}
                                                                _ = v82
                                                            }
                                                        case 4294967295:
                                                            var v83 uint = genFn0(154, 4294967295)
                                                            _ = v83
                                                        case 127:
                                                            switch 942 {
                                                            case 0:
                                                                v3["hello"] = (v81 - -1)
                                                            case 2147483647:
                                                                v2 = ((v2 + v2) + v9)
                                                            }
                                                        default:
                                                            v84 := GenS12{F0: fi17, F1: 42, F2: 1}
                                                            _ = v84
                                                        }
                                                        _ = v81
                                                    }
                                                }
                                            default:
                                                var v85 int = (fi17 << ((v16) & 63))
                                                _ = v85
                                            }
                                        case -2147483648:
                                            v16 = (^(v16))
                                        default:
                                            var v86 int32 = 1
                                            _ = v86
                                        }
                                        switch "a" {
                                        case "\n\t":
                                            var v87 int32 = 193
                                            _ = v87
                                        case "":
                                            v88 := GenS13{F0: 65535}
                                            _ = v88
                                        case "foo bar":
                                            if (float64((0.1 * float32(1.5))) > 0.0) {
                                                fmt.Println(v2)
                                                v9 = ((v2 + (v2 + "\n\t")) + v9)
                                            } else {
                                                var v89 uint16 = 255
                                                v9 = "a"
                                                _ = v89
                                            }
                                        default:
                                            var v90 int16 = 1
                                            _ = v90
                                        }
                                    }
                                }
                                var v91 int = (fi26 % (((fi17 & fi17)) | 1))
                                _ = v91
                            }
                        }
                        v92 := GenS3{F0: 1}
                        _ = fi26
                        _ = v92
                    }
                    var v93 int = (fi17 * fi17)
                    var v94 uint16 = 0
                    _ = v93
                    _ = v94
                }
                fmt.Println(v2, v9, v16)
            }
            v3[(v9 + (v9 + "hello"))] = fi17
            for fi95 := 0; fi95 < 14; fi95++ {
                var v96 uint16 = 523
                v96 = v96
                _ = fi95
                _ = v96
            }
        }
        delete(v4, 65535)
        if len(v1) > 0 {
            v1[0] = 2
        }
        _ = fi17
    }
    v16 -= (^(v16))
    var v97 int16 = 1000
    if false {
        switch 791 {
        case 0:
            if (!(-2147483648 == 0)) {
                delete(v4, 487)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        case -2147483648:
            _ = len(v4)
        }
    }
    var v98 string = (v2 + "hello")
    v2 = ""
    var v99 int8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
    _ = v97
    _ = v98
    _ = v99
}
