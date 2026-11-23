package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 int64; F1 uint; F2 uint8 }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 int32; F1 uint64; F2 string }
type GenS11 struct { F0 uint64; F1 uint; F2 int8 }
type GenS12 struct { F0 uint16 }
func genFn0() uint16 {
    var v7 int = -2147483648
    var v8 uint = 255
    var v9 uint64 = 846
    v10 := GenS1{F0: 0}
    for fi11 := 0; fi11 < 3; fi11++ {
        if (-1 != 42) {
            if true {
                v12 := GenS1{F0: 128}
                var v13 string = (("a" + "a") + "hello")
                var v14 int16 = -985
                _ = v12
                _ = v13
                _ = v14
            } else {
                var v15 int32 = -559
                switch -9223372036854775808 {
                case 621:
                    v16 := GenS2{F0: 471, F1: "\n\t"}
                    _ = v16
                case 583:
                    for fi17 := 0; fi17 < 19; fi17++ {
                        switch v7 {
                        case 1:
                            fmt.Println(fi17)
                        case 0:
                            var v18 int8 = -1
                            _ = v18
                        case 4294967295:
                            v19 := GenS3{F0: (v8 + uint(v9))}
                            _ = v19
                        }
                        var v20 float64 = 0.0
                        _ = fi17
                        _ = v20
                    }
                case -9223372036854775808:
                    var v21 float64 = ((3.14159 * -2.25) - (3.14159 - 0.0))
                    _ = v21
                default:
                    v22 := GenS1{F0: 42}
                    _ = v22
                }
                _ = v15
            }
            v23 := GenS3{F0: 265}
            _ = v23
        }
        if false {
            if true {
                if true {
                    var v24 string = "\n\t"
                    switch 2147483647 {
                    case -2147483648:
                        v24 = ""
                    case 128:
                        switch v24 {
                        case "hello":
                            for fi25 := 0; fi25 < 6; fi25++ {
                                var v26 float32 = (float32(float64(1e10)) - (1.5 + (0.1 / 1e10)))
                                var v27 int64 = 822
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        case "":
                            switch (-(fi11)) {
                            case -1:
                                fmt.Println(fi11)
                            case 389:
                                v9 = (v9 & v9)
                            case 824:
                                switch v24 {
                                case "foo bar":
                                    for fi28 := 0; fi28 < 12; fi28++ {
                                        v8 -= (v8 + 369)
                                        v29 := GenS1{F0: 41}
                                        _ = fi28
                                        _ = v29
                                    }
                                case "\n\t":
                                    v9 = (^(v9))
                                case "a":
                                    var v30 bool = (132 == int32(v8))
                                    _ = v30
                                default:
                                    v8 += v8
                                }
                            default:
                                var v31 string = ("foo bar" + v24)
                                _ = v31
                            }
                        case "foo bar":
                            switch v24 {
                            case "":
                                if true {
                                    if (true && ("" > "\n\t")) {
                                        var v32 int8 = 127
                                        fmt.Println(v24, v8)
                                        for fi33 := 0; fi33 < 3; fi33++ {
                                            fmt.Println(v24, v32, fi11)
                                            v32 = 15
                                            var v34 int8 = -128
                                            _ = fi33
                                            _ = v34
                                        }
                                        _ = v32
                                    } else {
                                        v8 += v8
                                        var v35 uint8 = 255
                                        var v36 int8 = -128
                                        _ = v35
                                        _ = v36
                                    }
                                } else {
                                    if true {
                                        fmt.Println(v9)
                                        var v37 uint8 = 91
                                        _ = v37
                                    }
                                    v38 := GenS1{F0: 220}
                                    if true {
                                        var v39 int64 = -197
                                        fmt.Println(v7, v24)
                                        var v40 int8 = 0
                                        _ = v39
                                        _ = v40
                                    }
                                    _ = v38
                                }
                            case "a":
                                for fi41 := 0; fi41 < 13; fi41++ {
                                    v24 = "a"
                                    var v42 uint64 = v9
                                    _ = fi41
                                    _ = v42
                                }
                            default:
                                v43 := []int{(fi11 ^ fi11), (^(v7))}
                                _ = v43
                            }
                        default:
                            if true {
                                for fi44 := 0; fi44 < 5; fi44++ {
                                    v8 = (v8 / (((v8 % (((v8 | 1)) | 1))) | 1))
                                    v45 := []int{int(v8), fi44, fi11}
                                    _ = fi44
                                    _ = v45
                                }
                                var v46 uint = v8
                                var v47 float64 = (-2.25 + (0.1 / 3.14159))
                                _ = v46
                                _ = v47
                            } else {
                                v48 := GenS4{F0: 127, F1: (v8 << ((42) & 63)), F2: 1}
                                fmt.Println(v9, v7)
                                switch (fi11 % ((9223372036854775807) | 1)) {
                                case 32768:
                                    if true {
                                        v49 := GenS5{F0: true}
                                        switch (^(fi11)) {
                                        case 127:
                                            v50 := []int{(v7 - (fi11 ^ (fi11 << (((v8 & v8)) & 63))))}
                                            _ = v50
                                        }
                                        var v51 int64 = 209
                                        _ = v49
                                        _ = v51
                                    }
                                case -465:
                                    var v52 uint16 = uint16(v9)
                                    _ = v52
                                }
                                _ = v48
                            }
                        }
                    }
                    _ = v24
                } else {
                    switch fi11 {
                    case 32767:
                        for fi53 := 0; fi53 < 17; fi53++ {
                            fmt.Println(fi11, v8, v7)
                            v8 -= (v8 + v8)
                            _ = fi53
                        }
                    case 128:
                        var v54 bool = true
                        _ = v54
                    }
                    var v55 string = "\n\t"
                    switch v55 {
                    case "foo bar":
                        v56 := GenS5{F0: true}
                        _ = v56
                    case "a":
                        switch 0 {
                        case 42:
                            var v57 float32 = -2.25
                            _ = v57
                        default:
                            switch ((v55 + (((("\n\t" + v55) + v55) + v55) + ("\n\t" + v55))) + ((v55 + (v55 + (v55 + "hello"))) + v55)) {
                            case "":
                                v58 := GenS6{F0: 18446744073709551615, F1: v55}
                                _ = v58
                            case "hello":
                                v7 = (-(v7))
                            default:
                                for fi59 := 0; fi59 < 7; fi59++ {
                                    v55 = ""
                                    v8 += v8
                                    _ = fi59
                                }
                            }
                        }
                    }
                    _ = v55
                }
                for fi60 := 0; fi60 < 9; fi60++ {
                    v61 := []int{(v7 % (((^(v7))) | 1)), -9223372036854775808, fi11, (v7 - v7)}
                    _ = fi60
                    _ = v61
                }
            } else {
                v7 = (v7 + (fi11 & fi11))
                switch (fi11 >> ((v8) & 63)) {
                case 42:
                    v62 := []int{}
                    _ = v62
                case 9223372036854775807:
                    v63 := []int{32767, int(v8), (fi11 + int(v8))}
                    _ = v63
                }
                var v64 uint16 = 2
                _ = v64
            }
            fmt.Println(fi11)
        }
        _ = fi11
    }
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 308
}

func main() {
    var v1 float64 = 1e10
    var v2 int8 = -1
    v3 := []string{"\n\t"}
    v4 := []uint32{2}
    v5 := make(map[string]uint32)
    var v6 uint = 65535
    v2 = (v2 >> ((v6) & 63))
    if true {
        v65 := GenS7{F0: (-(v2)), F1: 0}
        fmt.Println(v1, v6)
        _ = v65
    } else {
        v2 += (v2 | (v2 | v2))
        switch -308 {
        case -473:
            var v66 int8 = (v2 % ((v2) | 1))
            _ = v66
        case 1:
            switch (("" + "foo bar") + "hello") {
            case "a":
                var v67 int64 = 684
                _ = v67
            case "":
                v2 = 0
            case "\n\t":
                var v68 uint64 = 859
                _ = v68
            }
        case -2147483648:
            if true {
                _ = len(v5)
                for fi69 := 0; fi69 < 20; fi69++ {
                    var v70 int = fi69
                    switch (v70 + v70) {
                    case 9223372036854775807:
                        var v71 int = v70
                        _ = v71
                    case 42:
                        switch "a" {
                        case "a":
                            v6 += (v6 >> ((741) & 63))
                        case "foo bar":
                            for fi72 := 0; fi72 < 1; fi72++ {
                                v73 := GenS8{F0: -1, F1: v6, F2: 255}
                                var v74 int64 = 0
                                var v75 int32 = -1
                                _ = fi72
                                _ = v73
                                _ = v74
                                _ = v75
                            }
                        case "hello":
                            var v76 float32 = 3.14159
                            _ = v76
                        default:
                            var v77 uint = v6
                            _ = v77
                        }
                    case -9223372036854775808:
                        v78 := GenS3{F0: (v6 >> ((v6) & 63))}
                        _ = v78
                    default:
                        if (2 != 18446744073709551615) {
                            var v79 bool = (32767 >= 2)
                            for fi80 := 0; fi80 < 20; fi80++ {
                                v81 := GenS9{F0: 0, F1: -2147483648}
                                _ = len(v5)
                                _ = fi80
                                _ = v81
                            }
                            _ = v79
                        } else {
                            v1 -= v1
                        }
                    }
                    _ = fi69
                    _ = v70
                }
            } else {
                switch ("hello" + "foo bar") {
                case "":
                    v82 := GenS10{F0: 552, F1: 784, F2: "a"}
                    _ = v82
                case "a":
                    v83 := GenS9{F0: 4294967295, F1: -1}
                    _ = v83
                case "\n\t":
                    v84 := GenS11{F0: 700, F1: v6, F2: v2}
                    _ = v84
                default:
                    if ((3.14159 >= 3.14159) && (32767 == -167)) {
                        v5["a"] = 65535
                        for fi85 := 0; fi85 < 19; fi85++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if false {
                                v1 -= v1
                            } else {
                                var v86 int = (fi85 << (((^(v6))) & 63))
                                v3 = append(v3, ("" + "foo bar"))
                                _ = v86
                            }
                            _ = fi85
                        }
                        v5[("" + ("a" + "a"))] = 2147483647
                    }
                }
                var v87 uint32 = 844
                _ = v87
            }
        default:
            if true {
                if (42 < 822) {
                    v5[""] = 42
                    v5["\n\t"] = 2147483648
                }
            }
        }
    }
    if true {
        var v88 int64 = 928
        v1 = float64((((3.14159 + (float32(v1) - 3.14159)) * ((0.1 * (0.1 + -2.25)) * -2.25)) * (1.5 + 1.5)))
        _ = v88
    } else {
        var v89 int16 = -512
        _ = v89
    }
    var v90 int32 = 2147483647
    v1 = (0.0 / 3.14159)
    v91 := GenS6{F0: v6, F1: "\n\t"}
    fmt.Println(v90, v2)
    var v92 uint = (v6 * v6)
    v93 := GenS1{F0: 1}
    for fi94 := 0; fi94 < 8; fi94++ {
        switch -860 {
        case -2147483648:
            v95 := GenS9{F0: 42, F1: 2147483647}
            _ = v95
        case 0:
            var v96 int64 = -1
            _ = v96
        default:
            var v97 int64 = 42
            _ = v97
        }
        _ = fi94
    }
    if ("a" >= "hello") {
        var v98 uint16 = genFn0()
        _ = v98
    }
    for fi99 := 0; fi99 < 5; fi99++ {
        if true {
            if ((1e10 - 0.0) >= 1.5) {
                v92 = (v92 * 65535)
                if false {
                    v2 += (v2 & v2)
                    var v100 uint32 = 65535
                    var v101 int = fi99
                    _ = v100
                    _ = v101
                }
            } else {
                v6 -= v6
                var v102 int8 = (v2 & -1)
                _ = v102
            }
        }
        var v103 bool = false
        _ = fi99
        _ = v103
    }
    fmt.Println(v2)
    for fi104 := 0; fi104 < 11; fi104++ {
        v1 = 0.0
        _ = fi104
    }
    switch -579 {
    case -2147483648:
        for fi105 := 0; fi105 < 4; fi105++ {
            v106 := GenS4{F0: 127, F1: (^(v6)), F2: 128}
            _ = fi105
            _ = v106
        }
    case 9223372036854775807:
        switch 392 {
        case -451:
            _ = len(v5)
        case -153:
            for fi107 := 0; fi107 < 5; fi107++ {
                v108 := GenS12{F0: 0}
                v1 -= float64(0.1)
                _ = fi107
                _ = v108
            }
        default:
            var v109 uint16 = 65535
            _ = v109
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v90:", v90)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v90
    _ = v91
    _ = v92
    _ = v93
}
