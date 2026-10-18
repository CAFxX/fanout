package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int8; F1 uint8; F2 int }
type GenS5 struct { F0 uint32; F1 bool; F2 uint }
type GenS6 struct { F0 int; F1 int64; F2 int }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS9 struct { F0 string; F1 int; F2 int8 }
type GenS10 struct { F0 int16; F1 int8 }
type GenS11 struct { F0 int64; F1 uint; F2 uint32 }
type GenS12 struct { F0 int; F1 string }

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    fmt.Println(v1)
    v1 += -32768
    if (1 >= 4294967295) {
        for fi4 := 0; fi4 < 17; fi4++ {
            var v5 uint8 = 2
            _ = fi4
            _ = v5
        }
        var v6 float64 = ((((float64(float32(1e10)) * (1.5 + 0.1)) - 1.5) + 3.14159) + float64((0.0 * -2.25)))
        switch 9223372036854775807 {
        case -2147483648:
            var v7 string = ("a" + "a")
            _ = v7
        case -859:
            switch (("" + "foo bar") + "") {
            case "a":
                var v8 bool = (true || false)
                _ = v8
            default:
                fmt.Println(v6, v1)
            }
        default:
            if false {
                for fi9 := 0; fi9 < 17; fi9++ {
                    for fi10 := 0; fi10 < 9; fi10++ {
                        v11 := GenS1{F0: 2}
                        fmt.Println(fi10, v6)
                        var v12 int16 = v1
                        var v13 float32 = float32(((1e10 * v6) - (float64(3.14159) - float64(-2.25))))
                        if true {
                            if true {
                                switch "" {
                                case "foo bar":
                                    v12 += 1
                                case "":
                                    switch (fi9 >> ((4294967295) & 63)) {
                                    case -420:
                                        v13 += v13
                                    case 9223372036854775807:
                                        var v14 uint = 18446744073709551615
                                        _ = v14
                                    default:
                                        var v15 string = "a"
                                        _ = v15
                                    }
                                case "hello":
                                    switch "a" {
                                    case "foo bar":
                                        var v16 bool = (true || (2147483647 <= -1))
                                        _ = v16
                                    case "":
                                        fmt.Println(fi9, v13, v6)
                                    case "\n\t":
                                        v17 := GenS1{F0: 255}
                                        _ = v17
                                    }
                                }
                                var v18 int = (fi10 / ((fi10) | 1))
                                if true {
                                    switch "" {
                                    case "\n\t":
                                        var v19 string = ("" + ("a" + "a"))
                                        _ = v19
                                    case "":
                                        var v20 int16 = (v12 - (v12 << ((42) & 63)))
                                        _ = v20
                                    }
                                    v3[42] = (v18 + fi9)
                                    v6 -= v6
                                } else {
                                    v18 = (fi10 << ((42) & 63))
                                    v21 := GenS1{F0: 128}
                                    _ = v21
                                }
                                _ = v18
                            } else {
                                v22 := GenS2{F0: -128, F1: 576, F2: 42}
                                var v23 string = "foo bar"
                                v24 := []int{}
                                if (((v6 * v6) / v6) <= 1e10) {
                                    var v25 int16 = (v12 | (^(v1)))
                                    _ = v25
                                }
                                if true {
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                    v26 := GenS3{F0: v1, F1: 42}
                                    _ = v26
                                } else {
                                    var v27 int = (^(fi9))
                                    var v28 uint64 = 929
                                    _ = v27
                                    _ = v28
                                }
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                        } else {
                            fmt.Println(fi10, fi10)
                            v13 = v13
                        }
                        v29 := GenS3{F0: v12, F1: -942}
                        var v30 int = fi10
                        var v31 int16 = (v1 / (((v12 - 1000)) | 1))
                        _ = fi10
                        _ = v11
                        _ = v12
                        _ = v13
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                    _ = fi9
                }
                var v32 int64 = 1
                _ = v32
            }
        }
        _ = v6
    } else {
        v33 := []int{0, 471}
        var v34 uint = 79
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        _ = v33
        _ = v34
    }
    _ = len(v2)
    switch 84 {
    case -515:
        var v35 uint32 = 696
        _ = v35
    case 32767:
        if true {
            fmt.Println(v1)
            for fi36 := 0; fi36 < 3; fi36++ {
                var v37 uint32 = 2147483648
                switch (-(fi36)) {
                case -441:
                    var v38 int16 = v1
                    _ = v38
                case -1:
                    var v39 uint = 147
                    _ = v39
                }
                for fi40 := 0; fi40 < 17; fi40++ {
                    v41 := GenS3{F0: (v1 >> ((712) & 63)), F1: 1}
                    var v42 int32 = 2147483647
                    switch "a" {
                    case "hello":
                        var v43 int8 = -1
                        _ = v43
                    case "a":
                        if ((-207 == 1) && true) {
                            fmt.Println(v1, v42, fi36)
                            v44 := []int{fi40, fi40}
                            _ = v44
                        }
                        v45 := GenS3{F0: (-(v1)), F1: (v42 & -150)}
                        _ = v45
                    case "":
                        switch 32767 {
                        case -1:
                            if (32767 > 42) {
                                fmt.Println(fi40, v37, fi36)
                                v46 := []int{fi36, fi40, fi36}
                                _ = v46
                            } else {
                                v37 += v37
                                v47 := GenS1{F0: 0}
                                _ = v47
                            }
                        case 726:
                            if false {
                                if (!false) {
                                    if ((fi40 & (fi36 - fi36)) < fi36) {
                                        if ("foo bar" != "hello") {
                                            v42 = v42
                                            v42 += (v42 & (v42 * 42))
                                            if false {
                                                switch fi40 {
                                                case -9223372036854775808:
                                                    var v48 uint8 = 0
                                                    _ = v48
                                                default:
                                                    fmt.Println(v42, fi40)
                                                }
                                                var v49 int16 = v1
                                                fmt.Println(fi40)
                                                _ = v49
                                            } else {
                                                var v50 float64 = 1e10
                                                v42 -= -843
                                                _ = v50
                                            }
                                        } else {
                                            var v51 int32 = 809
                                            _ = v51
                                        }
                                    }
                                }
                                switch (fi40 << ((1) & 63)) {
                                case 9223372036854775807:
                                    v37 += v37
                                case 42:
                                    if false {
                                        var v52 bool = (!(false && true))
                                        if v52 {
                                            if (973 == (v42 << ((42) & 63))) {
                                                v42 -= (-(v42))
                                                v42 = (v42 << ((1) & 63))
                                                _ = len(v3)
                                            } else {
                                                switch 32767 {
                                                case 42:
                                                    v37 += (v37 | (v37 >> ((42) & 63)))
                                                case 127:
                                                    v37 = (v37 | (^(v37)))
                                                case 32768:
                                                    switch 128 {
                                                    case 4294967295:
                                                        var v53 string = ("foo bar" + "hello")
                                                        _ = v53
                                                    case 42:
                                                        switch "" {
                                                        case "a":
                                                            v54 := GenS1{F0: 255}
                                                            _ = v54
                                                        default:
                                                            fmt.Println(v42)
                                                        }
                                                    default:
                                                        v55 := GenS4{F0: 1, F1: 127, F2: 583}
                                                        _ = v55
                                                    }
                                                }
                                            }
                                            v42 = (v42 + (v42 << ((656) & 63)))
                                            v56 := GenS5{F0: (v37 - v37), F1: (!((("\n\t" + ("a" + "")) + "a") > "")), F2: 18446744073709551615}
                                            _ = v56
                                        } else {
                                            if v52 {
                                                v57 := GenS5{F0: (v37 * (v37 - (v37 << ((2) & 63)))), F1: (true || false), F2: 988}
                                                v58 := GenS1{F0: 127}
                                                _ = v57
                                                _ = v58
                                            } else {
                                                fmt.Println(v52, fi40, v1)
                                                var v59 uint64 = 1
                                                _ = v59
                                            }
                                        }
                                        var v60 int8 = 100
                                        _ = v52
                                        _ = v60
                                    }
                                    if true {
                                        v61 := []int{fi36, fi36, fi40, -9223372036854775808}
                                        switch (fi36 & (^(fi40))) {
                                        case -419:
                                            var v62 bool = (1 > 255)
                                            _ = v62
                                        case 547:
                                            if false {
                                                v61 = append(v61, 128)
                                            }
                                        default:
                                            var v63 uint32 = 2147483648
                                            _ = v63
                                        }
                                        v2[(fi40 & (fi36 >> ((437) & 63)))] = int(v1)
                                        var v64 bool = false
                                        _ = v61
                                        _ = v64
                                    }
                                    if (!(!false)) {
                                        var v65 bool = ((128 >= 2) && true)
                                        v66 := []int{(fi36 << ((1) & 63)), (fi36 % ((fi40) | 1)), (fi36 + (fi36 * (fi36 & fi36))), -359}
                                        switch fi36 {
                                        case -398:
                                            fmt.Println(fi36)
                                        }
                                        _ = v65
                                        _ = v66
                                    } else {
                                        var v67 int = (fi36 | (fi36 / ((42) | 1)))
                                        fmt.Println(v37, fi40, v42)
                                        v68 := GenS3{F0: v1, F1: 2147483647}
                                        _ = v67
                                        _ = v68
                                    }
                                case -2147483648:
                                    var v69 uint = 65535
                                    _ = v69
                                default:
                                    var v70 uint16 = 32767
                                    _ = v70
                                }
                                var v71 uint = 4294967295
                                _ = v71
                            }
                        case -2147483648:
                            v72 := GenS6{F0: 2147483647, F1: 1, F2: fi36}
                            _ = v72
                        }
                    default:
                        var v73 uint16 = 1
                        _ = v73
                    }
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                _ = fi36
                _ = v37
            }
        }
    case -2147483648:
        var v74 int = 956
        _ = v74
    default:
        var v75 float64 = ((-2.25 - ((-2.25 - float64(1.5)) * (1.5 - ((1e10 * ((0.0 + 0.0) + 1.5)) + float64(-2.25))))) - ((0.0 * (float64(0.0) * 3.14159)) + 1.5))
        _ = v75
    }
    var v76 int = 32767
    if false {
        var v77 uint64 = 0
        v76 = (v76 | v76)
        var v78 int8 = -1
        _ = v77
        _ = v78
    }
    v2[(v76 >> ((718) & 63))] = (v76 | v76)
    switch "" {
    case "a":
        for fi79 := 0; fi79 < 10; fi79++ {
            for fi80 := 0; fi80 < 20; fi80++ {
                v1 = v1
                switch "foo bar" {
                case "hello":
                    switch fi80 {
                    case 32768:
                        v1 += v1
                    case 0:
                        v81 := GenS3{F0: v1, F1: 0}
                        v76 = fi80
                        v1 = v1
                        _ = v81
                    case 1:
                        _ = len(v3)
                    }
                case "a":
                    v82 := GenS7{F0: (2 != 128), F1: ((v1 == (v1 >> ((18446744073709551615) & 63))) || (true && false))}
                    _ = v82
                case "foo bar":
                    v83 := GenS5{F0: 65535, F1: false, F2: 851}
                    _ = v83
                default:
                    fmt.Println(fi79)
                }
                _ = fi80
            }
            _ = fi79
        }
    case "foo bar":
        for fi84 := 0; fi84 < 8; fi84++ {
            switch ("foo bar" + "foo bar") {
            case "foo bar":
                if true {
                    var v85 int64 = 9223372036854775807
                    if false {
                        for fi86 := 0; fi86 < 19; fi86++ {
                            var v87 int32 = 42
                            var v88 uint16 = 32767
                            _ = fi86
                            _ = v87
                            _ = v88
                        }
                    } else {
                        if ((3.14159 - float32(0.0)) <= float32(float64(0.0))) {
                            v89 := GenS8{F0: 1, F1: (v1 ^ -1), F2: 128}
                            if true {
                                var v90 float64 = ((float64((float32(1.5) - (1.5 + 0.1))) + 1.5) + 0.1)
                                var v91 bool = false
                                _ = v90
                                _ = v91
                            }
                            _ = v89
                        }
                        switch (fi84 + (^(v76))) {
                        case 9223372036854775807:
                            v85 += v85
                        default:
                            switch "\n\t" {
                            case "":
                                v92 := []int{}
                                _ = v92
                            case "foo bar":
                                fmt.Println(v85)
                            case "a":
                                var v93 uint8 = 0
                                _ = v93
                            }
                        }
                    }
                    _ = v85
                }
            case "a":
                v94 := []int{}
                _ = v94
            }
            fmt.Println(v1, v1, fi84)
            fmt.Println(v1)
            _ = fi84
        }
    case "":
        if false {
            v95 := []int{0, -9223372036854775808}
            _ = v95
        }
    default:
        var v96 uint64 = 18446744073709551615
        _ = v96
    }
    if (32767 <= (v1 & v1)) {
        fmt.Println(v76)
        v76 = (-(v76))
        for fi97 := 0; fi97 < 20; fi97++ {
            v98 := []int{}
            var v99 uint32 = 403
            v1 = (v1 / ((-1) | 1))
            _ = fi97
            _ = v98
            _ = v99
        }
    }
    v76 = v76
    if true {
        for fi100 := 0; fi100 < 8; fi100++ {
            var v101 uint16 = 32767
            _ = fi100
            _ = v101
        }
        v102 := GenS9{F0: "a", F1: (v76 | -9223372036854775808), F2: 1}
        _ = v102
    } else {
        fmt.Println(v1)
        if ("" < "") {
            v103 := GenS10{F0: v1, F1: 0}
            _ = v103
        } else {
            fmt.Println(v1, v76)
        }
        fmt.Println(v1, v1)
    }
    v104 := GenS4{F0: -93, F1: 127, F2: 32767}
    if true {
        if true {
            v105 := GenS11{F0: -9223372036854775808, F1: 255, F2: 2147483647}
            for fi106 := 0; fi106 < 1; fi106++ {
                fmt.Println(v1)
                var v107 int16 = (v1 / ((v1) | 1))
                _ = fi106
                _ = v107
            }
            _ = v105
        } else {
            v76 -= v76
        }
        v76 += (^(v76))
    } else {
        v108 := v2[(v76 % (((v76 ^ (v76 / (((^(v76))) | 1)))) | 1))]
        v109 := GenS4{F0: 0, F1: 127, F2: 2147483647}
        v110 := GenS12{F0: (v76 >> ((506) & 63)), F1: "hello"}
        _ = v108
        _ = v109
        _ = v110
    }
    var v111 uint32 = 0
    var v112 uint16 = 382
    v1 = -842
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v76:", v76)
    _ = v1
    _ = v104
    _ = v111
    _ = v112
    _ = v2
    _ = v3
    _ = v76
}
