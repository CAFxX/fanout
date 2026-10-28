package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint64; F2 string }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int8; F1 uint8; F2 int64 }
type GenS7 struct { F0 uint32; F1 bool; F2 int8 }
type GenS8 struct { F0 int32; F1 int32 }
type GenS9 struct { F0 uint8; F1 string }
type GenS10 struct { F0 int; F1 int16; F2 int32 }
type GenS11 struct { F0 int32 }
func genFn0(p0 bool) bool {
    fmt.Println(p0)
    var v4 uint32 = 0
    v5 := []int{725, 1}
    var v6 uint16 = 501
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return ((-1 > 0) && ("" > "\n\t"))
}

func main() {
    var v1 int32 = 42
    var v2 uint16 = 32767
    v3 := make(map[uint32]int)
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 481}
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := GenS2{F0: "a", F1: (fi7 > (fi9 - 1))}
            v11 := GenS2{F0: "\n\t", F1: true}
            v1 = 1
            _ = fi9
            _ = v10
            _ = v11
        }
        v2 = v2
        _ = fi7
        _ = v8
    }
    v12 := v3[780]
    fmt.Println(v12, v2)
    v13 := GenS2{F0: "a", F1: (4294967295 <= 2)}
    for fi14 := 0; fi14 < 17; fi14++ {
        if true {
            v2 -= 65535
            v12 = fi14
        } else {
            if false {
                v3[2] = 9223372036854775807
            }
            v15 := GenS3{F0: 0}
            for fi16 := 0; fi16 < 5; fi16++ {
                v17 := GenS4{F0: 132, F1: 1, F2: "foo bar"}
                if (3.14159 >= (0.1 - 3.14159)) {
                    var v18 uint = 254
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
            _ = v15
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            fmt.Println(v2, v1, v1)
            v20 := []int{v12, 9223372036854775807, fi14, 127}
            var v21 int = fi14
            v1 += v1
            var v22 int16 = int16(v12)
            v21 += fi14
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        if (42 >= 844) {
            var v23 int32 = -2147483648
            switch 128 {
            case 580:
                if (32767 > (v12 - 4294967295)) {
                    v24 := v3[0]
                    var v25 int = 42
                    _ = v24
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        switch "\n\t" {
                        case "hello":
                            v12 -= v12
                        case "":
                            v23 += (v23 - 1)
                        case "a":
                            v23 = (v1 & 0)
                        default:
                            v2 = v2
                        }
                        _ = fi26
                    }
                }
            case 32768:
                for fi27 := 0; fi27 < 18; fi27++ {
                    fmt.Println(v23, fi27)
                    var v28 uint16 = v2
                    _ = fi27
                    _ = v28
                }
            case 128:
                v29 := GenS5{F0: 518, F1: 255}
                _ = v29
            }
            _ = v23
        }
        _ = fi14
    }
    v30 := GenS6{F0: 42, F1: 255, F2: 42}
    var v31 uint = 255
    v2 += 622
    var v32 float64 = 1e10
    switch -173 {
    case 0:
        var v33 int = v12
        _ = v33
    default:
        var v34 uint32 = 770
        _ = v34
    }
    v35 := v3[285]
    if (395 == 0) {
        if (v2 == v2) {
            v36 := []int{127, (v35 ^ v35)}
            _ = v36
        } else {
            if false {
                v37 := GenS7{F0: 1, F1: false, F2: -118}
                switch "\n\t" {
                case "":
                    delete(v3, 42)
                case "foo bar":
                    v1 = v1
                }
                _ = v37
            } else {
                v38 := v3[4294967295]
                switch 124 {
                case 32767:
                    if true {
                        switch v35 {
                        case 812:
                            if false {
                                var v39 uint32 = 450
                                _ = v39
                            }
                        default:
                            if false {
                                switch (v35 >> ((v31) & 63)) {
                                case -9223372036854775808:
                                    v32 = 0.0
                                default:
                                    var v40 uint64 = 0
                                    _ = v40
                                }
                                for fi41 := 0; fi41 < 4; fi41++ {
                                    v42 := GenS5{F0: 2147483648, F1: v31}
                                    v43 := []int{v12, (v35 >> ((36) & 63)), v35}
                                    var v44 uint64 = 281
                                    _ = fi41
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                }
                                v45 := GenS8{F0: v1, F1: v1}
                                _ = v45
                            } else {
                                v46 := GenS9{F0: 23, F1: "a"}
                                v1 = (v1 ^ (v1 * 808))
                                var v47 uint8 = 0
                                _ = v46
                                _ = v47
                            }
                        }
                    }
                }
                _ = v38
            }
            var v48 int8 = 59
            _ = v48
        }
        for fi49 := 0; fi49 < 16; fi49++ {
            if (!(65535 > (v2 % ((v2) | 1)))) {
                var v50 int8 = 4
                if (v32 <= float64((0.0 + 1e10))) {
                    for fi51 := 0; fi51 < 17; fi51++ {
                        v31 = v31
                        if ((v1 + (v1 - (v1 / ((v1) | 1)))) > (v1 * v1)) {
                            v3[636] = (-(v12))
                            fmt.Println(v2, v31)
                        } else {
                            v52 := GenS9{F0: 255, F1: ("hello" + "\n\t")}
                            _ = v52
                        }
                        v31 = v31
                        _ = fi51
                    }
                } else {
                    v53 := GenS7{F0: 2147483647, F1: false, F2: 127}
                    var v54 int64 = 0
                    v50 += int8(v54)
                    _ = v53
                    _ = v54
                }
                switch "hello" {
                case "a":
                    var v55 int = (-(v12))
                    _ = v55
                case "hello":
                    v3[2147483647] = (fi49 + v35)
                }
                _ = v50
            }
            if (true || (((v32 * -2.25) + (((v32 * float64(1.5)) - v32) - float64(3.14159))) > -2.25)) {
                var v56 string = (("foo bar" + "a") + "hello")
                var v57 int32 = (v1 * (v1 % ((v1) | 1)))
                for fi58 := 0; fi58 < 10; fi58++ {
                    v59 := GenS10{F0: v12, F1: 32767, F2: v1}
                    v57 = v57
                    switch v56 {
                    case "a":
                        v1 = -2147483648
                    case "hello":
                        if ((fi58 | v35) <= fi58) {
                            var v60 int32 = v57
                            fmt.Println(fi58)
                            fmt.Println(fi58, fi58, v32)
                            _ = v60
                        } else {
                            v61 := GenS4{F0: v31, F1: 0, F2: (v56 + "hello")}
                            _ = v61
                        }
                    default:
                        var v62 int64 = 42
                        _ = v62
                    }
                    _ = fi58
                    _ = v59
                }
                _ = v56
                _ = v57
            } else {
                switch (("foo bar" + "hello") + "foo bar") {
                case "\n\t":
                    v31 = v31
                case "":
                    fmt.Println(v35, v1)
                case "hello":
                    switch v35 {
                    case 4294967295:
                        v2 = 0
                    case 32768:
                        v1 -= v1
                    }
                default:
                    v63 := GenS4{F0: (v31 % ((v31) | 1)), F1: 4294967295, F2: ("foo bar" + "hello")}
                    _ = v63
                }
                v64 := []int{0, fi49, fi49}
                fmt.Println(v12)
                _ = v64
            }
            _ = fi49
        }
        delete(v3, 75)
    } else {
        var v65 float32 = float32(v32)
        _ = v65
    }
    fmt.Println(v31, v12, v35)
    v32 = (v32 * 1e10)
    switch "" {
    case "hello":
        v66 := GenS5{F0: 2147483647, F1: (v31 * v31)}
        _ = v66
    case "\n\t":
        if (!true) {
            v12 -= (v35 & (v12 >> ((v31) & 63)))
            v32 -= ((((v32 - (v32 - v32)) * ((1.5 / (v32 + v32)) * v32)) / 0.1) - v32)
            if (v2 < (v2 % ((v2) | 1))) {
                v12 = v35
                var v67 int64 = 9223372036854775807
                for fi68 := 0; fi68 < 19; fi68++ {
                    v69 := GenS7{F0: 65535, F1: (false || (42 != 42)), F2: 101}
                    if false {
                        for fi70 := 0; fi70 < 6; fi70++ {
                            v71 := GenS6{F0: -128, F1: 7, F2: 9223372036854775807}
                            if true {
                                switch 4294967295 {
                                case 857:
                                    if (!(!(!true))) {
                                        v2 -= 255
                                        fmt.Println(v67, v12)
                                        v72 := GenS5{F0: 1, F1: 1}
                                        switch "hello" {
                                        case "a":
                                            delete(v3, 470)
                                        default:
                                            var v73 int32 = v1
                                            _ = v73
                                        }
                                        _ = v72
                                    }
                                case 9223372036854775807:
                                    var v74 uint8 = 42
                                    _ = v74
                                case 2147483647:
                                    fmt.Println(v67, v31)
                                }
                                var v75 int8 = -1
                                var v76 uint8 = 134
                                _ = v75
                                _ = v76
                            }
                            _ = fi70
                            _ = v71
                        }
                        var v77 int8 = 119
                        v78 := v3[1]
                        _ = v77
                        _ = v78
                    } else {
                        switch (^(v35)) {
                        case 127:
                            var v79 int32 = (v1 >> (((v31 & v31)) & 63))
                            _ = v79
                        }
                        fmt.Println(fi68, v35, v32)
                    }
                    switch "" {
                    case "a":
                        for fi80 := 0; fi80 < 2; fi80++ {
                            switch (^(fi68)) {
                            case -1:
                                if (!false) {
                                    v3[991] = -19
                                }
                                if (!(0 != (^(v12)))) {
                                    if (!(!false)) {
                                        v1 -= v1
                                    } else {
                                        _ = len(v3)
                                    }
                                    if true {
                                        v81 := []int{v12, (fi80 % (((v35 - fi68)) | 1))}
                                        v81 = append(v81, 127)
                                        _ = v81
                                    }
                                    v82 := GenS11{F0: 761}
                                    _ = v82
                                } else {
                                    v83 := GenS11{F0: v1}
                                    _ = v83
                                }
                                v84 := []int{}
                                _ = v84
                            }
                            v32 += v32
                            fmt.Println(v1)
                            _ = fi80
                        }
                    default:
                        v1 += (v1 - v1)
                    }
                    _ = fi68
                    _ = v69
                }
                _ = v67
            } else {
                for fi85 := 0; fi85 < 15; fi85++ {
                    for fi86 := 0; fi86 < 12; fi86++ {
                        var v87 int32 = (v1 >> ((v31) & 63))
                        _ = fi86
                        _ = v87
                    }
                    _ = fi85
                }
                var v88 int = 32768
                for fi89 := 0; fi89 < 15; fi89++ {
                    v90 := []int{fi89}
                    if (true || false) {
                        var v91 float32 = float32(v32)
                        var v92 int8 = 100
                        var v93 float64 = (float64(v91) + 0.0)
                        _ = v91
                        _ = v92
                        _ = v93
                    }
                    _ = fi89
                    _ = v90
                }
                _ = v88
            }
        }
    default:
        fmt.Println(v2, v35)
    }
    for fi94 := 0; fi94 < 6; fi94++ {
        var v95 uint64 = 4294967295
        switch "" {
        case "a":
            v12 += int(v1)
        case "":
            if (3.14159 <= 1e10) {
                fmt.Println(v95, v31, v35)
            } else {
                v1 -= (-(v1))
                for fi96 := 0; fi96 < 7; fi96++ {
                    _ = len(v3)
                    v35 += v12
                    v31 = (v31 % ((v31) | 1))
                    _ = fi96
                }
            }
        case "hello":
            var v97 int8 = 43
            _ = v97
        }
        fmt.Println(v2, v2, v2)
        _ = fi94
        _ = v95
    }
    var v98 float64 = (v32 * (v32 / v32))
    v31 -= 375
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v98
}
