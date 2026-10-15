package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int64; F1 int64; F2 uint8 }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 uint64; F1 uint; F2 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint32; F1 int64 }
type GenS11 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS12 struct { F0 int8; F1 int64; F2 uint16 }
type GenS13 struct { F0 uint32; F1 int; F2 string }
func genFn0(p0 int32) int {
    var v6 uint8 = 0
    var v7 uint32 = 1
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: (fi8 & fi8), F1: (p0 % ((p0) | 1))}
        v10 := []int{}
        if (true || false) {
            v11 := GenS1{F0: fi8, F1: (^(p0))}
            var v12 string = "hello"
            v10 = append(v10, fi8)
            _ = v11
            _ = v12
        } else {
            p0 += (p0 / (((p0 - (p0 & p0))) | 1))
            fmt.Println(v7, v7)
            if (false || false) {
                for fi13 := 0; fi13 < 14; fi13++ {
                    p0 = p0
                    _ = fi13
                }
                if false {
                    if true {
                        v14 := GenS2{F0: (true && true)}
                        v15 := GenS2{F0: (v7 != (v7 * v7))}
                        var v16 float64 = 0.0
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        var v17 uint8 = 1
                        if ((p0 % ((p0) | 1)) > 1) {
                            if len(v10) > 0 {
                                v10[0] = fi8
                            }
                            v18 := GenS1{F0: -707, F1: -658}
                            if (!true) {
                                v19 := GenS3{F0: -643}
                                var v20 uint8 = v17
                                v7 -= (^(v7))
                                _ = v19
                                _ = v20
                            } else {
                                switch -2147483648 {
                                case 874:
                                    fmt.Println(p0)
                                default:
                                    for fi21 := 0; fi21 < 15; fi21++ {
                                        v22 := GenS4{F0: 4294967295, F1: 32767}
                                        var v23 int64 = -1
                                        _ = fi21
                                        _ = v22
                                        _ = v23
                                    }
                                }
                                var v24 uint = 942
                                _ = v24
                            }
                            _ = v18
                        } else {
                            for fi25 := 0; fi25 < 15; fi25++ {
                                if true {
                                    v26 := GenS5{F0: 1, F1: -804, F2: (^(v17))}
                                    switch fi25 {
                                    case -103:
                                        if (-1 <= 100) {
                                            switch (fi8 * fi8) {
                                            case -421:
                                                if (1 != 0) {
                                                    if false {
                                                        var v27 int = (fi8 ^ fi25)
                                                        var v28 string = "foo bar"
                                                        _ = v27
                                                        _ = v28
                                                    } else {
                                                        if ("hello" >= ("foo bar" + "foo bar")) {
                                                            v29 := GenS6{F0: p0, F1: 619}
                                                            fmt.Println(v7, fi8, v6)
                                                            v17 -= (v17 << ((0) & 63))
                                                            _ = v29
                                                        } else {
                                                            p0 += 0
                                                            v30 := GenS7{F0: (v17 / ((v6) | 1)), F1: fi25}
                                                            var v31 uint16 = 42
                                                            _ = v30
                                                            _ = v31
                                                        }
                                                        var v32 int = (-(fi25))
                                                        switch 9223372036854775807 {
                                                        case 2147483647:
                                                            if (!false) {
                                                                v33 := GenS7{F0: 24, F1: fi25}
                                                                v34 := GenS5{F0: 9223372036854775807, F1: 485, F2: v17}
                                                                v6 -= v17
                                                                _ = v33
                                                                _ = v34
                                                            } else {
                                                                var v35 int = (fi8 | 32767)
                                                                v32 = fi25
                                                                v36 := GenS8{F0: 889, F1: 72, F2: 42}
                                                                _ = v35
                                                                _ = v36
                                                            }
                                                        }
                                                        _ = v32
                                                    }
                                                    switch (-(fi25)) {
                                                    case 9223372036854775807:
                                                        v37 := GenS9{F0: 255}
                                                        _ = v37
                                                    default:
                                                        var v38 uint16 = 255
                                                        _ = v38
                                                    }
                                                    p0 += (p0 * p0)
                                                } else {
                                                    v39 := GenS7{F0: (v17 + (v17 >> ((2) & 63))), F1: 32768}
                                                    _ = v39
                                                }
                                            }
                                        } else {
                                            if true {
                                                var v40 int64 = -9223372036854775808
                                                v41 := GenS9{F0: 4294967295}
                                                _ = v40
                                                _ = v41
                                            }
                                            var v42 int8 = 42
                                            var v43 int64 = 0
                                            v10 = append(v10, (fi8 >> ((65535) & 63)))
                                            fmt.Println(fi25, v43)
                                            var v44 bool = (!true)
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        }
                                    case 0:
                                        var v45 bool = (((0.1 / 1.0) * -2.25) <= (0.0 / -2.25))
                                        _ = v45
                                    default:
                                        if (("a" + ("hello" + "foo bar")) < "\n\t") {
                                            if ((0.1 > float32(1e10)) && (!false)) {
                                                v7 += (v7 | 2)
                                                if len(v10) > 0 {
                                                    v10[0] = fi25
                                                }
                                            }
                                            v46 := GenS6{F0: p0, F1: uint64(v17)}
                                            var v47 int16 = 909
                                            _ = v46
                                            _ = v47
                                        }
                                        if true {
                                            if ((4294967295 == (v7 + 2147483647)) || ("" > "foo bar")) {
                                                var v48 uint64 = 1
                                                if ((^(v7)) == 65535) {
                                                    v49 := GenS5{F0: -1, F1: 1, F2: v17}
                                                    _ = v49
                                                }
                                                _ = v48
                                            } else {
                                                v10 = append(v10, fi25)
                                                p0 -= 1
                                                var v50 uint8 = (v6 ^ (v6 ^ (^(v17))))
                                                if true {
                                                    if false {
                                                        fmt.Println(v50)
                                                        var v51 int8 = -1
                                                        v17 = (v6 / ((v50) | 1))
                                                        _ = v51
                                                    }
                                                    v52 := GenS6{F0: p0, F1: 4294967295}
                                                    _ = v52
                                                }
                                                v53 := GenS10{F0: 369, F1: -9223372036854775808}
                                                fmt.Println(v50)
                                                _ = v50
                                                _ = v53
                                            }
                                        } else {
                                            switch "hello" {
                                            case "":
                                                v54 := GenS11{F0: v17, F1: 18446744073709551615, F2: 63}
                                                _ = v54
                                            case "foo bar":
                                                fmt.Println(v7, fi8)
                                            case "hello":
                                                v10 = append(v10, 127)
                                            }
                                            v55 := GenS7{F0: v17, F1: (fi25 | (fi8 | (fi8 << ((2) & 63))))}
                                            _ = v55
                                        }
                                    }
                                    _ = v26
                                }
                                var v56 int = fi8
                                v57 := GenS4{F0: (fi25 + (v56 * (fi25 ^ (fi8 >> ((42) & 63))))), F1: -1}
                                _ = fi25
                                _ = v56
                                _ = v57
                            }
                            fmt.Println(v6)
                            _ = len(v10)
                        }
                        _ = v17
                    }
                    if (true || false) {
                        p0 = 2147483647
                        for fi58 := 0; fi58 < 2; fi58++ {
                            v59 := GenS1{F0: fi58, F1: (^(p0))}
                            _ = fi58
                            _ = v59
                        }
                        p0 = p0
                    } else {
                        v60 := GenS6{F0: (p0 / ((int32(fi8)) | 1)), F1: 18446744073709551615}
                        for fi61 := 0; fi61 < 20; fi61++ {
                            v6 = v6
                            var v62 bool = (!false)
                            var v63 int = fi8
                            v10 = append(v10, fi8)
                            _ = len(v10)
                            v64 := GenS12{F0: 42, F1: 0, F2: 255}
                            _ = fi61
                            _ = v62
                            _ = v63
                            _ = v64
                        }
                        _ = v60
                    }
                }
            } else {
                var v65 bool = (735 == -1)
                _ = len(v10)
                _ = v65
            }
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 int32, p1 int32) uint8 {
    var v66 int64 = 42
    var v67 int16 = -1
    var v68 uint32 = 65535
    v69 := []int{32768, 42}
    var v70 uint64 = 841
    _ = p0
    _ = p1
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v70
    return 42
}

func main() {
    var v1 float64 = 1.5
    var v2 int64 = 0
    v3 := []string{(("" + "") + ""), (("\n\t" + ("a" + "")) + ""), "", "hello"}
    v4 := make(map[string]string)
    v5 := []int{-9223372036854775808, -1, 7, 117}
    v71 := GenS13{F0: 65535, F1: 32767, F2: (("\n\t" + "") + "\n\t")}
    _ = len(v4)
    for fi72 := 0; fi72 < 8; fi72++ {
        fmt.Println(fi72, v2, v2)
        if (!(true && (!(318 == 42)))) {
            var v73 bool = (((("hello" + "hello") + ("foo bar" + ("" + "\n\t"))) + "hello") <= "a")
            var v74 string = ""
            _ = v73
            _ = v74
        } else {
            var v75 int32 = -1
            v1 += (((v1 / (1.5 * ((float64(0.0) - 3.14159) * v1))) - float64(0.1)) - float64((1e10 - 1.5)))
            _ = len(v4)
            _ = v75
        }
        var v76 float64 = (0.0 - ((v1 - v1) * 0.1))
        _ = fi72
        _ = v76
    }
    for fi77 := 0; fi77 < 4; fi77++ {
        v1 += (v1 + (1.5 - -2.25))
        _ = fi77
    }
    if true {
        var v78 int16 = 1000
        for fi79 := 0; fi79 < 6; fi79++ {
            if ((!(143 != 32767)) || (!(1 <= 2))) {
                v1 += v1
                v80 := GenS1{F0: 1, F1: 42}
                _ = v80
            }
            switch (-(fi79)) {
            case 128:
                if false {
                    v1 += -2.25
                    v81 := GenS13{F0: 2, F1: genFn0(42), F2: "hello"}
                    _ = v81
                } else {
                    if (127 > -128) {
                        var v82 uint32 = 1
                        _ = v82
                    }
                }
            }
            for fi83 := 0; fi83 < 18; fi83++ {
                if (0 != -1) {
                    fmt.Println(v1, fi83)
                    v2 += -1
                }
                var v84 int8 = 1
                _ = fi83
                _ = v84
            }
            _ = fi79
        }
        _ = v78
    } else {
        fmt.Println(v2, v2)
        delete(v4, "")
    }
    for fi85 := 0; fi85 < 14; fi85++ {
        for fi86 := 0; fi86 < 1; fi86++ {
            switch fi85 {
            case 4294967295:
                fmt.Println(v2, fi86, v2)
            case 266:
                var v87 int16 = 42
                if ((281 != 65535) && false) {
                    v2 += (v2 * v2)
                    v2 += (v2 & v2)
                    delete(v4, "a")
                }
                _ = v87
            case 1:
                var v88 uint = 42
                _ = v88
            default:
                fmt.Println(v1)
            }
            v2 -= v2
            fmt.Println(fi85, fi85, v1)
            v89 := GenS5{F0: 284, F1: 42, F2: genFn1(0, 0)}
            switch (fi86 % ((fi85) | 1)) {
            case 9223372036854775807:
                switch ("hello" + "") {
                case "hello":
                    if true {
                        var v90 string = ("" + "a")
                        _ = v90
                    } else {
                        fmt.Println(fi86, fi85)
                        v2 = 211
                        fmt.Println(v2, fi85)
                    }
                default:
                    if len(v5) > 0 {
                        v5[0] = fi85
                    }
                }
            }
            _ = fi86
            _ = v89
        }
        var v91 uint = 1
        _ = fi85
        _ = v91
    }
    for fi92 := 0; fi92 < 20; fi92++ {
        v93 := GenS4{F0: fi92, F1: -255}
        var v94 int16 = 32767
        var v95 int8 = 42
        _ = fi92
        _ = v93
        _ = v94
        _ = v95
    }
    for fi96 := 0; fi96 < 9; fi96++ {
        fmt.Println(fi96)
        v2 -= (v2 >> ((2) & 63))
        _ = fi96
    }
    switch -1 {
    case 32768:
        _ = len(v4)
    case 42:
        if len(v3) > 0 {
            v3[0] = ("" + (("" + "foo bar") + "hello"))
        }
    case 127:
        var v97 int = 32767
        _ = v97
    }
    var v98 float64 = -2.25
    if true {
        var v99 int32 = 935
        var v100 float64 = v98
        _ = v100
        _ = v99
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v71
    _ = v98
}
