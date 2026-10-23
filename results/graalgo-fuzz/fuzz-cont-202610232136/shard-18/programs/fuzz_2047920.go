package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 bool; F2 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint8; F1 int32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int32 }
func genFn0(p0 int32, p1 int8) uint {
    for fi4 := 0; fi4 < 9; fi4++ {
        p1 = (^(p1))
        _ = fi4
    }
    p0 = p0
    v5 := GenS1{F0: 2147483647, F1: -62, F2: 2}
    _ = p0
    _ = p1
    _ = v5
    return 255
}

func main() {
    v1 := []uint32{881, 4294967295, 2}
    v2 := []string{}
    v3 := []string{"\n\t", "foo bar", "hello", ("hello" + "a")}
    switch -1 {
    case -870:
        if (((true || (!false)) || true) || false) {
            _ = len(v2)
            for fi6 := 0; fi6 < 18; fi6++ {
                for fi7 := 0; fi7 < 1; fi7++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v8 bool = true
                    _ = fi7
                    _ = v8
                }
                var v9 int = fi6
                _ = fi6
                _ = v9
            }
            if (42 < 2147483647) {
                var v10 int = 4294967295
                var v11 int16 = 1000
                if (!true) {
                    v10 = v10
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                    _ = len(v3)
                }
                _ = v10
                _ = v11
            } else {
                var v12 uint8 = 255
                v13 := GenS2{F0: 32767, F1: 2, F2: 2}
                for fi14 := 0; fi14 < 15; fi14++ {
                    v12 -= v12
                    _ = fi14
                }
                _ = v12
                _ = v13
            }
        } else {
            var v15 int = 32767
            _ = v15
        }
    case -970:
        var v16 uint16 = 0
        _ = v16
    default:
        var v17 uint32 = 2147483647
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            _ = len(v2)
            if true {
                var v20 uint32 = 4294967295
                v21 := GenS3{F0: genFn0(-2147483648, -1)}
                _ = v20
                _ = v21
            }
            _ = fi19
        }
        var v22 bool = (4294967295 != 42)
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 int16 = -32768
            if (!v22) {
                if v22 {
                    if true {
                        v24 -= 1
                    } else {
                        var v25 int = 2147483647
                        _ = v25
                    }
                } else {
                    if ((("foo bar" + ("foo bar" + "foo bar")) == ("hello" + ("hello" + ""))) && v22) {
                        fmt.Println(fi18, v22, fi18)
                        v26 := GenS4{F0: 1}
                        v27 := GenS1{F0: 2, F1: 55, F2: 94}
                        var v28 uint16 = 0
                        _ = v26
                        _ = v27
                        _ = v28
                    } else {
                        v3 = append(v3, "a")
                        switch fi23 {
                        case 71:
                            _ = len(v3)
                        default:
                            switch "foo bar" {
                            case "\n\t":
                                var v29 string = ("foo bar" + ((("\n\t" + ("hello" + "hello")) + "a") + "foo bar"))
                                _ = v29
                            case "foo bar":
                                v30 := GenS5{F0: 65535, F1: v22, F2: 4294967295}
                                _ = v30
                            }
                        }
                    }
                }
                fmt.Println(v22, v22)
                var v31 int = fi18
                _ = v31
            }
            _ = fi23
            _ = v24
        }
        _ = fi18
        _ = v22
    }
    if true {
        for fi32 := 0; fi32 < 12; fi32++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi32
        }
        var v33 bool = true
        v33 = (!false)
        _ = v33
    } else {
        v34 := GenS1{F0: 784, F1: 77, F2: 2147483648}
        if (4294967295 >= 360) {
            v35 := GenS1{F0: 42, F1: 0, F2: 585}
            var v36 string = "a"
            _ = v35
            _ = v36
        }
        _ = v34
    }
    v37 := GenS6{F0: false}
    var v38 string = "foo bar"
    v38 = ((v38 + "foo bar") + (("foo bar" + v38) + v38))
    switch 13 {
    case 128:
        for fi39 := 0; fi39 < 2; fi39++ {
            switch "\n\t" {
            case "foo bar":
                v38 = "\n\t"
            case "a":
                if false {
                    if (!false) {
                        for fi40 := 0; fi40 < 11; fi40++ {
                            v38 = v38
                            var v41 float32 = float32((float64((float32(1e10) * 0.1)) / -2.25))
                            _ = fi40
                            _ = v41
                        }
                    } else {
                        var v42 uint32 = 2147483648
                        _ = len(v3)
                        _ = v42
                    }
                    fmt.Println(fi39, v38)
                    var v43 uint = genFn0(2147483647, 1)
                    _ = v43
                }
            default:
                switch (fi39 % ((fi39) | 1)) {
                case -2147483648:
                    _ = len(v3)
                case 416:
                    for fi44 := 0; fi44 < 19; fi44++ {
                        fmt.Println(v38, fi39)
                        _ = fi44
                    }
                }
            }
            var v45 int8 = 127
            _ = fi39
            _ = v45
        }
    case 0:
        v38 = v38
    }
    v46 := GenS3{F0: genFn0(2147483647, -1)}
    v47 := GenS7{F0: 0, F1: 99}
    for fi48 := 0; fi48 < 1; fi48++ {
        var v49 int32 = 1
        fmt.Println(v49)
        fmt.Println(v38)
        _ = fi48
        _ = v49
    }
    var v50 int16 = -1
    switch -786 {
    case 2147483647:
        if true {
            fmt.Println(v38, v38)
            if (false && ((false || (true && false)) && (!(v38 > v38)))) {
                switch v38 {
                case "\n\t":
                    fmt.Println(v38, v50)
                case "hello":
                    var v51 int8 = -128
                    _ = v51
                }
                fmt.Println(v50, v50)
                switch v38 {
                case "\n\t":
                    fmt.Println(v38)
                case "hello":
                    var v52 int64 = 9223372036854775807
                    _ = v52
                default:
                    if ((4294967295 > 0) || (100 == 100)) {
                        var v53 uint16 = 65535
                        switch v38 {
                        case "":
                            if (true || ((-14 > 42) && false)) {
                                var v54 uint8 = 42
                                for fi55 := 0; fi55 < 12; fi55++ {
                                    v53 = 218
                                    for fi56 := 0; fi56 < 4; fi56++ {
                                        switch (fi55 + 1) {
                                        case 2147483647:
                                            if (-634 != 1) {
                                                var v57 int16 = v50
                                                v58 := GenS8{F0: (fi56 + (fi56 ^ (fi55 * int(v57)))), F1: -1}
                                                var v59 float32 = -2.25
                                                _ = v57
                                                _ = v58
                                                _ = v59
                                            }
                                        default:
                                            if false {
                                                if ((9223372036854775807 != 4294967295) && (2147483647 > uint32(v53))) {
                                                    var v60 int64 = -1
                                                    switch v38 {
                                                    case "hello":
                                                        var v61 int8 = 100
                                                        _ = v61
                                                    default:
                                                        var v62 bool = ((v38 + v38) != (v38 + ("\n\t" + "\n\t")))
                                                        _ = v62
                                                    }
                                                    _ = v60
                                                } else {
                                                    if ((42 <= 904) && (!true)) {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        var v63 float64 = 1.5
                                                        var v64 int64 = -573
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                    var v65 int8 = 42
                                                    _ = v65
                                                }
                                            } else {
                                                var v66 int = fi55
                                                fmt.Println(fi56)
                                                v50 = v50
                                                _ = v66
                                            }
                                        }
                                        _ = fi56
                                    }
                                    _ = fi55
                                }
                                _ = v54
                            } else {
                                var v67 int32 = -512
                                v67 = v67
                                var v68 int64 = 0
                                _ = v67
                                _ = v68
                            }
                        default:
                            v38 = (v38 + ("a" + "\n\t"))
                        }
                        _ = v53
                    } else {
                        if false {
                            for fi69 := 0; fi69 < 2; fi69++ {
                                var v70 int32 = 1
                                fmt.Println(fi69, fi69)
                                _ = fi69
                                _ = v70
                            }
                            if (42 != 628) {
                                switch "" {
                                case "":
                                    v71 := GenS9{F0: 4294967295}
                                    _ = v71
                                case "a":
                                    fmt.Println(v38, v50)
                                default:
                                    v38 = v38
                                }
                                fmt.Println(v50, v50)
                                for fi72 := 0; fi72 < 19; fi72++ {
                                    v1 = append(v1, 4294967295)
                                    switch v38 {
                                    case "\n\t":
                                        if true {
                                            switch 0 {
                                            case -2147483648:
                                                if false {
                                                    var v73 uint8 = 255
                                                    switch (fi72 ^ (fi72 * fi72)) {
                                                    case 127:
                                                        switch fi72 {
                                                        case 2147483647:
                                                            var v74 uint8 = 1
                                                            _ = v74
                                                        case 42:
                                                            if ((v73 ^ (^(v73))) >= v73) {
                                                                if true {
                                                                    switch v38 {
                                                                    case "a":
                                                                        v75 := GenS4{F0: 0}
                                                                        _ = v75
                                                                    case "foo bar":
                                                                        v38 = (v38 + v38)
                                                                    }
                                                                }
                                                                switch (v38 + v38) {
                                                                case "":
                                                                    for fi76 := 0; fi76 < 6; fi76++ {
                                                                        fmt.Println(fi76)
                                                                        v77 := GenS8{F0: (fi72 - fi72), F1: v50}
                                                                        _ = fi76
                                                                        _ = v77
                                                                    }
                                                                case "\n\t":
                                                                    if true {
                                                                        v38 = v38
                                                                    } else {
                                                                        v73 += (v73 * v73)
                                                                        v38 = v38
                                                                        switch ("\n\t" + v38) {
                                                                        case "hello":
                                                                            v50 += 32767
                                                                        case "a":
                                                                            var v78 uint32 = 79
                                                                            _ = v78
                                                                        case "foo bar":
                                                                            if (false && (-2147483648 <= 0)) {
                                                                                v50 = v50
                                                                                var v79 string = v38
                                                                                v73 -= v73
                                                                                _ = v79
                                                                            }
                                                                        default:
                                                                            var v80 uint16 = 1
                                                                            _ = v80
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                if true {
                                                                    for fi81 := 0; fi81 < 2; fi81++ {
                                                                        v38 = (v38 + "hello")
                                                                        _ = fi81
                                                                    }
                                                                    fmt.Println(fi72, v38)
                                                                }
                                                                var v82 bool = true
                                                                fmt.Println(fi72, v73, v82)
                                                                _ = v82
                                                            }
                                                        case 0:
                                                            if false {
                                                                v50 += v50
                                                                v38 = v38
                                                                var v83 uint = genFn0(42, 42)
                                                                _ = v83
                                                            }
                                                        }
                                                    }
                                                    if false {
                                                        for fi84 := 0; fi84 < 3; fi84++ {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v85 := GenS10{F0: (v73 % ((203) | 1)), F1: 42}
                                                            v50 = -94
                                                            _ = fi84
                                                            _ = v85
                                                        }
                                                        for fi86 := 0; fi86 < 13; fi86++ {
                                                            var v87 string = v38
                                                            v50 -= (^(v50))
                                                            _ = fi86
                                                            _ = v87
                                                        }
                                                    }
                                                    _ = v73
                                                }
                                            default:
                                                for fi88 := 0; fi88 < 6; fi88++ {
                                                    v3 = append(v3, v38)
                                                    if ((!true) || false) {
                                                        if true {
                                                            v89 := GenS4{F0: -2147483648}
                                                            _ = v89
                                                        } else {
                                                            var v90 uint16 = 32767
                                                            _ = v90
                                                        }
                                                        v91 := GenS11{F0: -13}
                                                        var v92 uint = 904
                                                        fmt.Println(v38, fi88)
                                                        var v93 int64 = 1
                                                        if len(v1) > 0 {
                                                            v1[0] = 882
                                                        }
                                                        _ = v91
                                                        _ = v92
                                                        _ = v93
                                                    } else {
                                                        var v94 string = "\n\t"
                                                        var v95 float64 = (1.5 + 0.0)
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                    _ = fi88
                                                }
                                            }
                                            for fi96 := 0; fi96 < 5; fi96++ {
                                                v38 = v38
                                                _ = fi96
                                            }
                                        }
                                    case "":
                                        fmt.Println(v50)
                                    default:
                                        var v97 bool = true
                                        _ = v97
                                    }
                                    var v98 float32 = 0.0
                                    _ = fi72
                                    _ = v98
                                }
                            }
                            v99 := GenS3{F0: 65535}
                            _ = v99
                        }
                    }
                }
            } else {
                var v100 uint8 = 255
                if (!(!false)) {
                    v100 = (^(v100))
                    var v101 uint16 = 0
                    for fi102 := 0; fi102 < 4; fi102++ {
                        var v103 uint8 = (v100 / ((v100) | 1))
                        _ = fi102
                        _ = v103
                    }
                    _ = v101
                }
                fmt.Println(v50, v38, v50)
                _ = v100
            }
        }
    case 751:
        if true {
            v104 := GenS12{F0: -1}
            _ = len(v2)
            _ = v104
        } else {
            v50 += v50
        }
    case -764:
        var v105 uint16 = 42
        _ = v105
    }
    v38 = "hello"
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v46
    _ = v47
    _ = v50
}
