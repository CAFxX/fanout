package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int32; F1 uint; F2 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int8; F1 uint8 }
type GenS9 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS10 struct { F0 uint; F1 int64 }
type GenS11 struct { F0 string; F1 int; F2 int8 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 bool; F1 int8 }
type GenS14 struct { F0 string; F1 uint }

func main() {
    v1 := make(map[string]string)
    var v2 int = 9223372036854775807
    var v3 uint = 65535
    v4 := GenS1{F0: ("foo bar" + "\n\t"), F1: 359, F2: "hello"}
    v5 := GenS2{F0: 897}
    var v6 bool = (false || false)
    for fi7 := 0; fi7 < 17; fi7++ {
        v3 += v3
        v8 := []int{(v2 * -1), 0}
        v9 := GenS2{F0: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v6 = v6
        _ = fi10
    }
    var v11 uint64 = 42
    _ = len(v1)
    if v6 {
        v12 := []int{v2, 4294967295, v2, -1}
        v1[""] = ""
        _ = v12
    } else {
        switch ("" + "\n\t") {
        case "":
            for fi13 := 0; fi13 < 17; fi13++ {
                if false {
                    v11 += v11
                    for fi14 := 0; fi14 < 9; fi14++ {
                        v6 = (!false)
                        _ = fi14
                    }
                    var v15 int = 724
                    _ = v15
                }
                _ = fi13
            }
        case "foo bar":
            var v16 float32 = (0.0 - 3.14159)
            _ = v16
        case "\n\t":
            v2 = (v2 + -1)
        default:
            if (9223372036854775807 < 9223372036854775807) {
                fmt.Println(v3, v6)
            } else {
                v17 := GenS3{F0: 6, F1: false}
                var v18 uint8 = 42
                fmt.Println(v6, v11, v11)
                _ = v17
                _ = v18
            }
        }
        v19 := []int{}
        _ = v19
    }
    v20 := GenS1{F0: ((("" + "foo bar") + "") + ("\n\t" + "a")), F1: uint32(v3), F2: ("a" + "foo bar")}
    v11 = v11
    for fi21 := 0; fi21 < 2; fi21++ {
        v6 = false
        _ = fi21
    }
    v3 += v3
    if v6 {
        if v6 {
            fmt.Println(v3, v11, v2)
            fmt.Println(v2, v6, v11)
            var v22 int = (v2 << ((42) & 63))
            _ = v22
        } else {
            switch 927 {
            case -9223372036854775808:
                if v6 {
                    if ((!("a" > "foo bar")) && (!(!false))) {
                        v23 := GenS4{F0: (v2 & v2), F1: -44}
                        fmt.Println(v2)
                        v24 := GenS1{F0: (("" + "\n\t") + "\n\t"), F1: 0, F2: ""}
                        _ = v23
                        _ = v24
                    } else {
                        switch 1 {
                        case -2147483648:
                            if (float32(float64(1e10)) != 3.14159) {
                                for fi25 := 0; fi25 < 5; fi25++ {
                                    v26 := GenS5{F0: 2147483647, F1: (v3 & (v3 << (((^(v3))) & 63))), F2: 255}
                                    _ = fi25
                                    _ = v26
                                }
                                fmt.Println(v2, v6, v6)
                                v27 := GenS6{F0: 479}
                                _ = v27
                            }
                        case 321:
                            switch (v2 / (((v2 >> ((4294967295) & 63))) | 1)) {
                            case -733:
                                for fi28 := 0; fi28 < 5; fi28++ {
                                    v29 := []int{(v2 >> ((v3) & 63)), (v2 / ((fi28) | 1)), v2}
                                    v30 := GenS2{F0: -510}
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                            case -9223372036854775808:
                                switch 1 {
                                case 32767:
                                    for fi31 := 0; fi31 < 8; fi31++ {
                                        var v32 uint = uint(v11)
                                        switch fi31 {
                                        case 4294967295:
                                            var v33 float64 = 0.1
                                            _ = v33
                                        default:
                                            switch "a" {
                                            case "a":
                                                var v34 bool = (2147483648 >= 2147483647)
                                                _ = v34
                                            }
                                        }
                                        if (v6 || (255 != 42)) {
                                            switch v2 {
                                            case 734:
                                                for fi35 := 0; fi35 < 13; fi35++ {
                                                    v36 := []int{}
                                                    v1[""] = ("foo bar" + "foo bar")
                                                    _ = fi35
                                                    _ = v36
                                                }
                                            }
                                            if (9223372036854775807 != 42) {
                                                fmt.Println(v3, v2)
                                            }
                                        } else {
                                            v37 := []int{fi31, (fi31 | 128), v2, (v2 + fi31)}
                                            v11 += 0
                                            v11 = v11
                                            _ = v37
                                        }
                                        _ = fi31
                                        _ = v32
                                    }
                                case 32768:
                                    v38 := v1[""]
                                    _ = v38
                                default:
                                    var v39 uint8 = 20
                                    _ = v39
                                }
                            }
                        default:
                            v11 += (^(v11))
                        }
                    }
                } else {
                    var v40 uint8 = 255
                    v41 := []int{v2, v2, (v2 % (((v2 % ((v2) | 1))) | 1))}
                    switch v2 {
                    case 128:
                        for fi42 := 0; fi42 < 17; fi42++ {
                            v3 += (^(v3))
                            v11 = (v11 - (v11 & (v11 % (((v11 ^ (v11 + 18446744073709551615))) | 1))))
                            var v43 uint32 = 113
                            _ = fi42
                            _ = v43
                        }
                    case 344:
                        if (!v6) {
                            switch "foo bar" {
                            case "\n\t":
                                v44 := v1[("\n\t" + "a")]
                                _ = v44
                            case "hello":
                                for fi45 := 0; fi45 < 10; fi45++ {
                                    switch (^(v2)) {
                                    case 42:
                                        for fi46 := 0; fi46 < 11; fi46++ {
                                            if true {
                                                v6 = ((v40 >> (((v3 % (((v3 & v3)) | 1))) & 63)) != 127)
                                                v11 = (v11 / ((4294967295) | 1))
                                            } else {
                                                switch fi45 {
                                                case -119:
                                                    var v47 uint32 = uint32(v11)
                                                    _ = v47
                                                default:
                                                    switch ("\n\t" + "hello") {
                                                    case "a":
                                                        var v48 int = v2
                                                        _ = v48
                                                    case "hello":
                                                        if (!((fi45 / ((-2147483648) | 1)) < (v2 ^ 418))) {
                                                            _ = len(v41)
                                                            var v49 int16 = 0
                                                            _ = v49
                                                        } else {
                                                            if (-9223372036854775808 > -488) {
                                                                if (!v6) {
                                                                    v40 = v40
                                                                    if (v6 && (v11 >= (v11 | (^(v11))))) {
                                                                        if v6 {
                                                                            if (18446744073709551615 >= uint(v40)) {
                                                                                if v6 {
                                                                                    var v50 uint = 176
                                                                                    _ = v50
                                                                                }
                                                                                if (9223372036854775807 >= -1) {
                                                                                    if v6 {
                                                                                        v2 -= 933
                                                                                        fmt.Println(v2, v40)
                                                                                        v2 -= 127
                                                                                    } else {
                                                                                        v2 = (^(v2))
                                                                                    }
                                                                                    v6 = (!v6)
                                                                                    v2 = (-(v2))
                                                                                }
                                                                                v6 = (42 >= 255)
                                                                            }
                                                                            v40 += v40
                                                                            _ = len(v1)
                                                                        } else {
                                                                            v51 := GenS6{F0: 2147483647}
                                                                            if (42 > 58) {
                                                                                var v52 int8 = 100
                                                                                v52 = v52
                                                                                v53 := GenS2{F0: -1}
                                                                                _ = v52
                                                                                _ = v53
                                                                            }
                                                                            if (v40 < 1) {
                                                                                if v6 {
                                                                                    v11 = (^(v11))
                                                                                    if len(v41) >= 2 {
                                                                                        v41 = v41[1:len(v41)]
                                                                                    }
                                                                                    v54 := GenS7{F0: 2147483648, F1: fi45}
                                                                                    fmt.Println(fi46, v11)
                                                                                    if true {
                                                                                        v2 -= (fi46 - 127)
                                                                                        v3 -= v3
                                                                                        var v55 uint = (v3 / ((uint(v11)) | 1))
                                                                                        v56 := GenS6{F0: 845}
                                                                                        _ = v55
                                                                                        _ = v56
                                                                                    } else {
                                                                                        v57 := GenS4{F0: fi46, F1: -128}
                                                                                        v3 += 88
                                                                                        if len(v41) >= 2 {
                                                                                            v41 = v41[1:len(v41)]
                                                                                        }
                                                                                        _ = v57
                                                                                    }
                                                                                    _ = v54
                                                                                }
                                                                                v41 = append(v41, 1)
                                                                            } else {
                                                                                v2 = v2
                                                                                v40 = 42
                                                                                v58 := v1["a"]
                                                                                _ = v58
                                                                            }
                                                                            _ = v51
                                                                        }
                                                                        v41 = append(v41, (fi45 % ((fi45) | 1)))
                                                                        v11 -= (v11 ^ (v11 ^ (^(v11))))
                                                                    } else {
                                                                        v11 = v11
                                                                    }
                                                                }
                                                            }
                                                            v59 := v1["a"]
                                                            v40 += (^(v40))
                                                            v3 -= v3
                                                            _ = v59
                                                        }
                                                    case "":
                                                        v40 -= (v40 * v40)
                                                    }
                                                }
                                                if (v11 > (v11 * (^(v11)))) {
                                                    if ((-117 == -128) && (!(42 >= 1))) {
                                                        fmt.Println(v40, v40, v11)
                                                        v60 := GenS8{F0: -1, F1: (v40 ^ 128)}
                                                        v11 -= v11
                                                        delete(v1, "\n\t")
                                                        v41 = append(v41, fi45)
                                                        _ = v60
                                                    }
                                                    var v61 uint8 = 117
                                                    var v62 uint16 = 0
                                                    _ = v61
                                                    _ = v62
                                                }
                                            }
                                            switch "\n\t" {
                                            case "hello":
                                                v63 := GenS4{F0: fi46, F1: -1}
                                                var v64 uint8 = (v40 >> ((v3) & 63))
                                                _ = v63
                                                _ = v64
                                            }
                                            _ = fi46
                                        }
                                    case -617:
                                        var v65 uint = (v3 | 65535)
                                        _ = v65
                                    case 0:
                                        v66 := GenS2{F0: 996}
                                        _ = v66
                                    default:
                                        v67 := GenS5{F0: 42, F1: 18446744073709551615, F2: 127}
                                        _ = v67
                                    }
                                    for fi68 := 0; fi68 < 7; fi68++ {
                                        var v69 uint32 = 2147483648
                                        switch "hello" {
                                        case "foo bar":
                                            v11 -= (v11 - (^(v11)))
                                        case "a":
                                            _ = len(v1)
                                        default:
                                            var v70 uint = v3
                                            _ = v70
                                        }
                                        v6 = v6
                                        v71 := GenS9{F0: -1, F1: v11, F2: v69}
                                        _ = fi68
                                        _ = v69
                                        _ = v71
                                    }
                                    var v72 int16 = -582
                                    _ = fi45
                                    _ = v72
                                }
                            case "foo bar":
                                var v73 float64 = 0.0
                                _ = v73
                            }
                            v11 = (v11 | 18446744073709551615)
                        }
                    }
                    _ = v40
                    _ = v41
                }
            case 2147483647:
                if (194 <= 1) {
                    switch "hello" {
                    case "a":
                        v74 := GenS10{F0: (v3 + (v3 ^ (v3 ^ v3))), F1: 1}
                        _ = v74
                    case "foo bar":
                        if false {
                            fmt.Println(v11, v6)
                            fmt.Println(v11, v11, v6)
                        } else {
                            var v75 uint64 = (^(v11))
                            _ = v75
                        }
                    }
                } else {
                    if (0.1 == 1e10) {
                        var v76 uint = 65535
                        v2 = (v2 + (v2 / ((918) | 1)))
                        _ = v76
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            v77 := GenS11{F0: "", F1: int(v11), F2: 127}
                            _ = v77
                        case "\n\t":
                            for fi78 := 0; fi78 < 7; fi78++ {
                                v79 := GenS7{F0: 42, F1: (^(v2))}
                                v80 := GenS12{F0: 1}
                                _ = fi78
                                _ = v79
                                _ = v80
                            }
                        case "a":
                            v6 = (v3 <= (v3 % ((v3) | 1)))
                        default:
                            fmt.Println(v3, v6)
                        }
                        var v81 int64 = 9223372036854775807
                        _ = v81
                    }
                }
            }
            v82 := []int{v2}
            _ = v82
        }
        var v83 uint64 = (v11 % ((191) | 1))
        _ = v83
    }
    if (v3 < (v3 & v3)) {
        for fi84 := 0; fi84 < 8; fi84++ {
            var v85 uint64 = 0
            switch ("a" + "hello") {
            case "a":
                var v86 float64 = 0.0
                _ = v86
            case "":
                delete(v1, "")
            }
            _ = fi84
            _ = v85
        }
        var v87 float64 = float64(float32(3.14159))
        _ = v87
    } else {
        v88 := GenS7{F0: 4294967295, F1: v2}
        var v89 int32 = -2147483648
        v90 := GenS13{F0: (!v6), F1: 110}
        _ = v88
        _ = v89
        _ = v90
    }
    v91 := GenS14{F0: "", F1: v3}
    v92 := []int{}
    if v6 {
        v93 := v1["\n\t"]
        _ = v93
    }
    for fi94 := 0; fi94 < 19; fi94++ {
        for fi95 := 0; fi95 < 15; fi95++ {
            fmt.Println(fi95)
            _ = fi95
        }
        _ = fi94
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v92:", len(v92), v92)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v91
    _ = v92
}
