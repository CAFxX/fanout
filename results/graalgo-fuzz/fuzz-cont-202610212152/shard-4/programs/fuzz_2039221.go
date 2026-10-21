package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint }
type GenS4 struct { F0 uint; F1 bool; F2 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8; F1 int16; F2 int32 }
type GenS10 struct { F0 uint64; F1 uint16 }
type GenS11 struct { F0 bool; F1 uint16; F2 int64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint; F1 int16 }
type GenS14 struct { F0 int; F1 int; F2 int64 }
func genFn0() uint8 {
    var v6 uint8 = 127
    v6 -= v6
    if (((1.5 + 0.0) + float64(1.5)) <= float64((1.5 / -2.25))) {
        for fi7 := 0; fi7 < 7; fi7++ {
            v6 += (v6 ^ 0)
            v6 -= (v6 & 128)
            _ = fi7
        }
    }
    _ = v6
    return (v6 & 101)
}

func main() {
    var v1 int8 = -128
    var v2 uint8 = 123
    v3 := []int{128}
    var v4 int32 = 2147483647
    v5 := make(map[int]bool)
    if ((0.1 + 1.5) < ((-2.25 / 1.0) - (1e10 / (1.5 / 1e10)))) {
        var v8 uint = uint(v2)
        if (true || false) {
            v2 = 198
            v2 += v2
        }
        switch 127 {
        case 853:
            switch (("a" + "") + "a") {
            case "a":
                var v9 int8 = v1
                _ = v9
            case "":
                if (!false) {
                    var v10 uint = v8
                    switch "foo bar" {
                    case "":
                        var v11 int8 = 127
                        _ = v11
                    }
                    _ = v10
                } else {
                    if (867 != 433) {
                        var v12 uint8 = (^(v2))
                        var v13 uint = v8
                        _ = v12
                        _ = v13
                    } else {
                        v3 = append(v3, 0)
                        for fi14 := 0; fi14 < 8; fi14++ {
                            v15 := GenS1{F0: v4, F1: 2, F2: 4294967295}
                            delete(v5, -636)
                            var v16 uint = (v8 * 262)
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        var v17 int = -96
                        _ = v17
                    }
                }
            default:
                v5[127] = (128 <= 40)
            }
        case -894:
            var v18 float64 = ((1e10 * 0.1) / 0.1)
            _ = v18
        case 9223372036854775807:
            if (!false) {
                for fi19 := 0; fi19 < 12; fi19++ {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        v3 = append(v3, 128)
                        v8 += 153
                        _ = fi20
                    }
                    _ = fi19
                }
                delete(v5, 42)
                v4 = 1
            } else {
                fmt.Println(v4)
                v21 := GenS2{F0: (v4 << ((v8) & 63)), F1: (v4 * 328)}
                _ = v21
            }
        default:
            v8 += (v8 << ((v8) & 63))
        }
        _ = v8
    }
    switch "" {
    case "\n\t":
        if (561 == 65535) {
            switch "" {
            case "foo bar":
                var v22 int16 = 618
                _ = v22
            default:
                v23 := GenS3{F0: (v4 / ((0) | 1)), F1: 112, F2: 4294967295}
                _ = v23
            }
            var v24 float64 = -2.25
            v1 -= v1
            _ = v24
        }
    case "":
        v2 = (v2 | v2)
    case "hello":
        _ = len(v5)
    }
    if len(v3) > 0 {
        v3[0] = -1
    }
    var v25 float32 = 1e10
    var v26 int8 = v1
    for fi27 := 0; fi27 < 1; fi27++ {
        if true {
            v25 = v25
            v1 += v1
        }
        for fi28 := 0; fi28 < 19; fi28++ {
            v4 = (v4 + (v4 / ((v4) | 1)))
            v29 := GenS1{F0: 1, F1: 18446744073709551615, F2: 0}
            _ = len(v3)
            if ((987 >= 1) || false) {
                switch "" {
                case "a":
                    v30 := GenS4{F0: 1, F1: (!(!(9223372036854775807 < -1))), F2: 214}
                    _ = v30
                case "\n\t":
                    v4 = (v4 | (v4 + v4))
                case "":
                    if len(v3) > 0 {
                        v3[0] = fi27
                    }
                }
                var v31 int = (fi28 - (fi27 ^ (fi28 ^ 128)))
                v31 += (^(v31))
                v1 -= (v1 * v1)
                _ = v31
            } else {
                v26 = v26
                if (!(true && true)) {
                    v2 = (v2 & (v2 / ((v2) | 1)))
                    delete(v5, (fi27 & -2147483648))
                    if (42 >= 879) {
                        var v32 uint64 = 1
                        v2 = 128
                        var v33 int = fi28
                        _ = v32
                        _ = v33
                    }
                    v34 := GenS1{F0: 2147483647, F1: 42, F2: 1}
                    var v35 uint8 = (v2 & v2)
                    _ = v34
                    _ = v35
                }
            }
            var v36 uint = 2
            _ = fi28
            _ = v29
            _ = v36
        }
        _ = fi27
    }
    for fi37 := 0; fi37 < 17; fi37++ {
        for fi38 := 0; fi38 < 4; fi38++ {
            if len(v3) > 0 {
                v3[0] = -822
            }
            v26 -= (v1 * 1)
            var v39 string = (("foo bar" + ((("hello" + "a") + "\n\t") + ("a" + ((("" + "foo bar") + ("" + "\n\t")) + ("hello" + ""))))) + "a")
            v1 += 127
            fmt.Println(v26, v25)
            _ = fi38
            _ = v39
        }
        if true {
            if (125 >= -9223372036854775808) {
                for fi40 := 0; fi40 < 10; fi40++ {
                    v2 = 212
                    if false {
                        v25 += v25
                        v41 := GenS5{F0: 32767}
                        switch "hello" {
                        case "a":
                            if false {
                                if len(v3) > 0 {
                                    v3[0] = -2147483648
                                }
                                v42 := GenS6{F0: fi37}
                                _ = v42
                            } else {
                                if (-2.25 != -2.25) {
                                    v43 := GenS7{F0: 0}
                                    var v44 int8 = 0
                                    switch (^(fi37)) {
                                    case 4294967295:
                                        v1 = 0
                                    default:
                                        v26 -= v1
                                    }
                                    _ = len(v3)
                                    if ((true && true) && true) {
                                        var v45 float64 = -2.25
                                        _ = v45
                                    }
                                    v46 := GenS8{F0: ("" + ("" + "a"))}
                                    _ = v43
                                    _ = v44
                                    _ = v46
                                } else {
                                    v25 += (v25 / (1e10 - 0.0))
                                    if false {
                                        var v47 int64 = 1
                                        v48 := GenS9{F0: v1, F1: 530, F2: 773}
                                        if (false && (!true)) {
                                            var v49 uint64 = 2
                                            _ = v49
                                        } else {
                                            var v50 uint64 = 653
                                            v47 = -870
                                            _ = v50
                                        }
                                        _ = v47
                                        _ = v48
                                    } else {
                                        v1 += (v26 - (v26 & (^(v1))))
                                    }
                                    if false {
                                        v25 = (v25 - v25)
                                        var v51 bool = false
                                        var v52 uint32 = 398
                                        _ = v51
                                        _ = v52
                                    } else {
                                        var v53 bool = true
                                        _ = v53
                                    }
                                }
                                v54 := GenS10{F0: 341, F1: 32767}
                                var v55 uint = 65535
                                _ = v54
                                _ = v55
                            }
                        }
                        _ = v41
                    } else {
                        var v56 uint16 = 255
                        if true {
                            v25 = 0.0
                        } else {
                            _ = len(v5)
                            v57 := GenS9{F0: 42, F1: 376, F2: (v4 ^ (v4 / (((v4 * (v4 - v4))) | 1)))}
                            _ = v57
                        }
                        v58 := v5[fi40]
                        _ = len(v5)
                        switch "" {
                        case "a":
                            var v59 int64 = -234
                            _ = v59
                        }
                        if ((v56 & v56) < 32767) {
                            if len(v3) > 0 {
                                v3[0] = (fi37 - int(v56))
                            }
                            var v60 uint32 = 528
                            var v61 int16 = 0
                            _ = v60
                            _ = v61
                        }
                        if (42 < 18446744073709551615) {
                            v1 += int8(v2)
                            v5[253] = (!(v58 || ((-2147483648 != v4) || v58)))
                        } else {
                            var v62 bool = (((!(1 >= 42)) || (!(((v56 - v56) > (v56 + v56)) && ("a" == "\n\t")))) && (false && (-2.25 < v25)))
                            fmt.Println(v25)
                            var v63 int32 = v4
                            _ = v62
                            _ = v63
                        }
                        _ = v56
                        _ = v58
                    }
                    _ = fi40
                }
                v26 = (v1 | 100)
                v64 := GenS10{F0: 4294967295, F1: 255}
                _ = v64
            }
            var v65 float64 = float64((float32(float64(float32(1.5))) / (v25 / 1.5)))
            _ = v65
        } else {
            if len(v3) > 0 {
                v3[0] = fi37
            }
            for fi66 := 0; fi66 < 4; fi66++ {
                switch "a" {
                case "\n\t":
                    var v67 uint32 = 4294967295
                    _ = v67
                case "foo bar":
                    var v68 int32 = (-(v4))
                    _ = v68
                }
                fmt.Println(v26, v2, v4)
                fmt.Println(v4)
                _ = fi66
            }
        }
        _ = fi37
    }
    if true {
        v1 = v26
    } else {
        v2 = v2
        v69 := GenS9{F0: (^(v26)), F1: 32767, F2: (v4 % (((v4 & v4)) | 1))}
        _ = v69
    }
    for fi70 := 0; fi70 < 1; fi70++ {
        v71 := v5[fi70]
        _ = len(v3)
        if len(v3) > 0 {
            v3[0] = fi70
        }
        _ = fi70
        _ = v71
    }
    if false {
        delete(v5, 514)
    } else {
        v2 -= v2
        var v72 int32 = v4
        if len(v3) > 0 {
            v3[0] = 42
        }
        _ = v72
    }
    _ = len(v5)
    if true {
        for fi73 := 0; fi73 < 20; fi73++ {
            var v74 bool = false
            _ = fi73
            _ = v74
        }
        var v75 int32 = (v4 / ((v4) | 1))
        var v76 uint = 2
        _ = v75
        _ = v76
    } else {
        v77 := GenS11{F0: false, F1: 1, F2: -1}
        v78 := GenS11{F0: (877 < 992), F1: 1, F2: -9223372036854775808}
        var v79 int32 = (v4 & 0)
        _ = v77
        _ = v78
        _ = v79
    }
    var v80 int = 346
    for fi81 := 0; fi81 < 18; fi81++ {
        switch (fi81 << ((205) & 63)) {
        case 2147483647:
            var v82 int64 = 9223372036854775807
            _ = v82
        case 260:
            if (!false) {
                if (!true) {
                    fmt.Println(v4, v26, fi81)
                } else {
                    var v83 int = v80
                    _ = v83
                }
            } else {
                var v84 uint = 42
                v85 := GenS12{F0: 871}
                _ = v84
                _ = v85
            }
        }
        for fi86 := 0; fi86 < 3; fi86++ {
            v1 += v26
            if false {
                fmt.Println(v25)
                var v87 uint = 260
                v3 = append(v3, int(v87))
                v88 := GenS13{F0: (v87 >> ((2) & 63)), F1: 32767}
                v4 = -1
                v89 := GenS9{F0: v1, F1: -32768, F2: 466}
                switch v80 {
                case 32767:
                    v90 := GenS14{F0: 32768, F1: (v80 >> (((v87 * (v87 + (v87 + v87)))) & 63)), F2: 9223372036854775807}
                    _ = v90
                case -808:
                    if false {
                        v87 += (v87 & (v87 << ((v87) & 63)))
                        v25 = 3.14159
                        if true {
                            v25 += v25
                        }
                        var v91 uint = (^(v87))
                        _ = v91
                    } else {
                        fmt.Println(v2, v1)
                    }
                case -2147483648:
                    var v92 uint8 = (v2 / ((0) | 1))
                    _ = v92
                }
                _ = v87
                _ = v88
                _ = v89
            } else {
                delete(v5, 161)
                var v93 uint16 = 255
                _ = v93
            }
            fmt.Println(v2, fi81, fi86)
            _ = fi86
        }
        v2 -= (^(v2))
        _ = fi81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v80
}
