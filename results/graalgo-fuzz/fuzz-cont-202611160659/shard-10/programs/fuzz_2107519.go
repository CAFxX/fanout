package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 string }
type GenS2 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 uint64; F1 int32 }
func genFn0(p0 bool, p1 uint64) bool {
    var v4 int64 = -1
    v5 := GenS1{F0: -625, F1: -128, F2: "hello"}
    switch -9223372036854775808 {
    case -728:
        if p0 {
            var v6 uint16 = 65535
            _ = v6
        }
    default:
        switch "a" {
        case "a":
            v7 := GenS1{F0: -1, F1: -128, F2: "a"}
            _ = v7
        default:
            v8 := []int{2147483647}
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        p0 = (-56 == -128)
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := GenS2{F0: p1, F1: -518, F2: 42}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 float32 = (0.0 + 1e10)
    _ = p0
    _ = p1
    _ = v12
    _ = v4
    _ = v5
    return false
}
func genFn1(p0 uint64, p1 int) int32 {
    v13 := []int{4294967295, p1, p1, (^(p1))}
    if (125 > 1) {
        var v14 int16 = 42
        _ = v14
    } else {
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := GenS3{F0: 0}
            fmt.Println(p1, fi15, fi15)
            var v17 int64 = -896
            _ = fi15
            _ = v16
            _ = v17
        }
        if len(v13) > 0 {
            v13[0] = p1
        }
    }
    v18 := GenS1{F0: 1, F1: 42, F2: ""}
    _ = p0
    _ = p1
    _ = v13
    _ = v18
    return 772
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]int)
    var v3 string = ("" + ("hello" + ("hello" + "\n\t")))
    v3 = v3
    v2[2] = 2147483647
    fmt.Println(v1, v3)
    v19 := []int{-1, 32768}
    switch v3 {
    case "\n\t":
        fmt.Println(v3, v3)
    case "foo bar":
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
    case "":
        fmt.Println(v3, v1)
    default:
        v1 = false
    }
    v20 := GenS2{F0: 0, F1: genFn1(1, 2147483647), F2: 127}
    var v21 int16 = 1000
    _ = len(v2)
    v22 := GenS4{F0: 2147483648, F1: 216}
    if true {
        v1 = (42 == 128)
        v21 += (v21 % ((32767) | 1))
        v23 := GenS3{F0: -1}
        _ = v23
    }
    if ((0.0 > 1.5) && (v1 || v1)) {
        var v24 uint16 = 78
        _ = v24
    } else {
        v3 = ("\n\t" + v3)
        for fi25 := 0; fi25 < 16; fi25++ {
            var v26 float64 = 0.1
            fmt.Println(v26, v21, v3)
            v1 = (18446744073709551615 <= 1)
            _ = fi25
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        if v1 {
            v1 = (!(v1 || ((!(v1 && ((v1 || (2147483648 <= 110)) && v1))) || v1)))
            delete(v2, 2147483648)
            switch ("" + ((v3 + (v3 + "")) + v3)) {
            case "\n\t":
                for fi28 := 0; fi28 < 14; fi28++ {
                    v21 += (v21 | (^(v21)))
                    _ = fi28
                }
            }
        } else {
            var v29 uint16 = 255
            if (4294967295 >= 65535) {
                var v30 string = v3
                for fi31 := 0; fi31 < 1; fi31++ {
                    v32 := v2[4294967295]
                    if (1 != (v21 & v21)) {
                        switch ("foo bar" + (v3 + v30)) {
                        case "hello":
                            if v1 {
                                fmt.Println(v21)
                                switch v32 {
                                case 42:
                                    v1 = true
                                }
                            } else {
                                var v33 int32 = genFn1(2, (fi27 / ((fi27) | 1)))
                                switch "" {
                                case "foo bar":
                                    v32 = -1
                                case "a":
                                    var v34 uint16 = v29
                                    _ = v34
                                case "":
                                    v35 := GenS3{F0: genFn1(42, (v32 | (fi27 / (((fi27 & fi31)) | 1))))}
                                    _ = v35
                                default:
                                    if false {
                                        var v36 uint64 = 702
                                        fmt.Println(v36)
                                        _ = v36
                                    } else {
                                        v1 = true
                                        if len(v19) >= 2 {
                                            v19 = v19[1:len(v19)]
                                        }
                                        switch (^(fi31)) {
                                        case -319:
                                            v29 += v29
                                        case -962:
                                            v37 := GenS4{F0: 1, F1: 2}
                                            _ = v37
                                        default:
                                            v32 += (fi27 * -9223372036854775808)
                                        }
                                    }
                                }
                                v32 = 1
                                _ = v33
                            }
                            v3 = v3
                        case "\n\t":
                            fmt.Println(v29, v21)
                        }
                        v1 = ((((v1 && (v21 >= -1)) || (1 >= 42)) || v1) || ((genFn1(42, (v32 | (fi27 << ((4294967295) & 63)))) != 931) || v1))
                        v19 = append(v19, (fi31 << ((0) & 63)))
                        var v38 int64 = -9223372036854775808
                        if (!(v1 && ((v21 <= int16(v38)) || false))) {
                            v1 = (!v1)
                        } else {
                            if v1 {
                                v21 += 371
                                var v39 int64 = (^(v38))
                                if false {
                                    var v40 int8 = 127
                                    var v41 float32 = 1e10
                                    _ = v40
                                    _ = v41
                                }
                                _ = v39
                            } else {
                                if (61 >= 65535) {
                                    switch v30 {
                                    case "hello":
                                        var v42 bool = ((1e10 * -2.25) <= ((3.14159 * 0.0) - 3.14159))
                                        _ = v42
                                    case "foo bar":
                                        if v1 {
                                            var v43 int = (fi27 - fi31)
                                            _ = v43
                                        } else {
                                            delete(v2, 1)
                                            switch (fi31 + (v32 | -318)) {
                                            case 2147483647:
                                                switch v30 {
                                                case "\n\t":
                                                    var v44 int32 = 42
                                                    _ = v44
                                                case "foo bar":
                                                    var v45 string = (v30 + v3)
                                                    _ = v45
                                                case "a":
                                                    v38 = (v38 * v38)
                                                }
                                            case 0:
                                                v3 = v30
                                            default:
                                                delete(v2, 5)
                                            }
                                            _ = len(v2)
                                        }
                                    }
                                } else {
                                    _ = len(v19)
                                    if (!(v1 && (!v1))) {
                                        v29 -= 504
                                        switch v3 {
                                        case "\n\t":
                                            v1 = (!v1)
                                            var v46 int64 = v38
                                            _ = v46
                                        case "hello":
                                            switch "" {
                                            case "hello":
                                                if (v1 && (!v1)) {
                                                    v29 += 2
                                                    v1 = (!true)
                                                    fmt.Println(v3)
                                                } else {
                                                    v3 = ""
                                                }
                                            case "foo bar":
                                                fmt.Println(v1)
                                            }
                                        case "":
                                            delete(v2, 1)
                                        }
                                        var v47 int16 = v21
                                        _ = v47
                                    } else {
                                        if (v1 || (1 < 225)) {
                                            if v1 {
                                                if false {
                                                    v2[2147483647] = v32
                                                    v48 := GenS2{F0: 4294967295, F1: 348, F2: 129}
                                                    var v49 uint32 = 2147483647
                                                    _ = v48
                                                    _ = v49
                                                }
                                            } else {
                                                fmt.Println(v3, v38)
                                            }
                                            v19 = append(v19, (fi31 - fi31))
                                        }
                                        if len(v19) > 0 {
                                            v19[0] = (^(fi31))
                                        }
                                    }
                                    var v50 uint8 = 255
                                    _ = v50
                                }
                                if v1 {
                                    v51 := GenS5{F0: true, F1: 542}
                                    v52 := GenS2{F0: 352, F1: 395, F2: 128}
                                    _ = v51
                                    _ = v52
                                } else {
                                    var v53 uint = 922
                                    fmt.Println(fi31)
                                    var v54 float64 = 1.5
                                    _ = v53
                                    _ = v54
                                }
                                if (42 <= 42) {
                                    var v55 int = (fi31 & fi31)
                                    _ = v55
                                }
                                var v56 string = v3
                                v1 = (!true)
                                _ = v56
                            }
                        }
                        switch (fi31 << ((1) & 63)) {
                        case 0:
                            var v57 uint32 = 2147483647
                            _ = v57
                        case 128:
                            if false {
                                v58 := GenS2{F0: 573, F1: -2147483648, F2: uint8(v29)}
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                                _ = v58
                            }
                        }
                        _ = v38
                    } else {
                        var v59 uint8 = 10
                        v60 := GenS4{F0: 2147483648, F1: v29}
                        var v61 float64 = (0.0 / 1.0)
                        _ = v59
                        _ = v60
                        _ = v61
                    }
                    v21 = v21
                    _ = fi31
                    _ = v32
                }
                var v62 int = (^(fi27))
                _ = v30
                _ = v62
            }
            for fi63 := 0; fi63 < 16; fi63++ {
                var v64 int16 = v21
                var v65 uint64 = 0
                if (v1 && (int32(fi27) > genFn1((v65 >> ((1) & 63)), 0))) {
                    v21 += (v64 | (v64 & (v21 + (v21 & (-(v21))))))
                    v3 = "\n\t"
                    v1 = false
                    v21 = (^(v21))
                } else {
                    v66 := GenS5{F0: false, F1: 4294967295}
                    v67 := v2[840]
                    fmt.Println(fi63)
                    v65 += (^(v65))
                    var v68 int32 = -1
                    if ((^(v68)) == v68) {
                        if (v68 < (^(v68))) {
                            if (v1 && (v1 && (!(9 != 0)))) {
                                if ((v21 % (((v64 + -1)) | 1)) <= int16(v68)) {
                                    v69 := GenS1{F0: genFn1(v65, fi27), F1: 127, F2: v3}
                                    v70 := GenS3{F0: v68}
                                    var v71 uint16 = v29
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                } else {
                                    var v72 uint8 = 104
                                    v19 = append(v19, 1)
                                    v73 := GenS4{F0: 65535, F1: v29}
                                    if (-701 > (v68 - v68)) {
                                        v74 := GenS3{F0: (v68 * -2147483648)}
                                        var v75 int32 = (v68 / (((-(v68))) | 1))
                                        _ = v74
                                        _ = v75
                                    } else {
                                        var v76 uint8 = v72
                                        _ = v76
                                    }
                                    var v77 string = ""
                                    _ = v72
                                    _ = v73
                                    _ = v77
                                }
                            } else {
                                var v78 int = (fi27 & (fi63 ^ fi63))
                                v68 -= genFn1(0, -2147483648)
                                _ = v78
                            }
                            _ = len(v2)
                        } else {
                            fmt.Println(v3)
                            switch v3 {
                            case "\n\t":
                                v79 := GenS6{F0: v65, F1: v1}
                                _ = v79
                            case "a":
                                switch 42 {
                                case 758:
                                    switch v67 {
                                    case 2147483647:
                                        v1 = v1
                                    case 127:
                                        var v80 uint = 1
                                        _ = v80
                                    }
                                case -778:
                                    var v81 int8 = 1
                                    _ = v81
                                case 815:
                                    fmt.Println(v21, v29, fi27)
                                default:
                                    if len(v19) > 0 {
                                        v19[0] = int(v68)
                                    }
                                }
                            case "hello":
                                v19 = append(v19, 2147483647)
                            }
                            var v82 int16 = (v21 % ((1000) | 1))
                            _ = v82
                        }
                        fmt.Println(v21)
                    }
                    var v83 uint64 = (v65 + (^(v65)))
                    _ = v66
                    _ = v67
                    _ = v68
                    _ = v83
                }
                _ = fi63
                _ = v64
                _ = v65
            }
            _ = v29
        }
        _ = fi27
    }
    for fi84 := 0; fi84 < 5; fi84++ {
        switch -537 {
        case -764:
            for fi85 := 0; fi85 < 17; fi85++ {
                var v86 uint16 = 42
                _ = fi85
                _ = v86
            }
        case 9223372036854775807:
            if (1 == 206) {
                var v87 int32 = 67
                if (!v1) {
                    v88 := GenS7{F0: (v21 <= v21), F1: (42 <= -9223372036854775808)}
                    _ = v88
                } else {
                    v89 := GenS8{F0: 106, F1: (-(v87))}
                    for fi90 := 0; fi90 < 3; fi90++ {
                        v91 := v2[0]
                        _ = fi90
                        _ = v91
                    }
                    v3 = v3
                    _ = v89
                }
                _ = v87
            }
        }
        v92 := GenS3{F0: 2147483647}
        v21 = v21
        _ = fi84
        _ = v92
    }
    var v93 int64 = -9223372036854775808
    for fi94 := 0; fi94 < 3; fi94++ {
        switch v3 {
        case "foo bar":
            var v95 uint64 = 4294967295
            _ = v95
        default:
            v96 := GenS4{F0: 2, F1: 1}
            _ = v96
        }
        for fi97 := 0; fi97 < 9; fi97++ {
            fmt.Println(fi97, v3, fi97)
            _ = fi97
        }
        _ = fi94
    }
    v98 := GenS6{F0: 1, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v93
    _ = v98
}
