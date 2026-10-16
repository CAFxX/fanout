package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 int8; F2 int16 }
type GenS6 struct { F0 bool; F1 uint32; F2 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8; F1 bool }
type GenS9 struct { F0 uint64; F1 int8; F2 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 42
    v3 := make(map[uint32]bool)
    var v4 int16 = v2
    var v5 uint16 = 1
    v5 = (v5 << ((2) & 63))
    v4 = v4
    v2 = v2
    if (v2 <= v4) {
        v4 += (v4 + v4)
        switch 32768 {
        case 2147483647:
            v6 := GenS1{F0: "\n\t"}
            _ = v6
        default:
            v5 -= 932
        }
        for fi7 := 0; fi7 < 2; fi7++ {
            if true {
                v8 := []int{(^(fi7)), (^(fi7)), 0, (^(fi7))}
                var v9 float64 = (-2.25 - 0.1)
                var v10 int64 = 1
                _ = v10
                _ = v8
                _ = v9
            }
            v11 := []int{fi7}
            for fi12 := 0; fi12 < 8; fi12++ {
                fmt.Println(v5)
                if (true && true) {
                    var v13 int32 = 1
                    var v14 float32 = (1.5 / 1e10)
                    v14 -= (float32(3.14159) / v14)
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS1{F0: ("" + "")}
                    _ = v15
                }
                switch -9223372036854775808 {
                case 128:
                    v2 -= v4
                case 9223372036854775807:
                    fmt.Println(v5)
                case 42:
                    if len(v11) > 0 {
                        v11[0] = 0
                    }
                }
                if false {
                    v4 = v2
                    var v16 uint16 = 135
                    v5 = v16
                    v17 := GenS2{F0: true}
                    var v18 int16 = 0
                    v19 := GenS3{F0: 169, F1: ""}
                    v18 -= -328
                    v16 -= 32767
                    v16 += (v5 % ((42) | 1))
                    v16 = (^(v16))
                    var v20 uint16 = (v5 - (v5 << ((1) & 63)))
                    if false {
                        var v21 uint64 = 795
                        var v22 uint64 = 4294967295
                        switch (^(fi12)) {
                        case 652:
                            var v23 int16 = v2
                            _ = v23
                        case 218:
                            fmt.Println(v2, v4)
                        default:
                            var v24 uint = 362
                            _ = v24
                        }
                        _ = v21
                        _ = v22
                    } else {
                        switch "a" {
                        case "":
                            if false {
                                var v25 uint8 = 115
                                _ = v25
                            } else {
                                v5 += v16
                                v5 = v16
                                var v26 uint32 = 1
                                _ = v26
                            }
                        case "foo bar":
                            fmt.Println(v2, v16, fi12)
                            switch (("\n\t" + "a") + "foo bar") {
                            case "hello":
                                v5 = (v16 ^ 692)
                            case "\n\t":
                                switch 32767 {
                                case 9223372036854775807:
                                    switch "" {
                                    case "\n\t":
                                        if ((("" + "hello") + "") <= "foo bar") {
                                            v27 := GenS2{F0: false}
                                            _ = v27
                                        } else {
                                            var v28 bool = false
                                            _ = v28
                                        }
                                    case "foo bar":
                                        v5 += (v5 / (((v20 ^ 115)) | 1))
                                    case "a":
                                        var v29 float64 = float64((((0.1 * float32((float64(0.0) * 0.1))) + 3.14159) * 0.1))
                                        _ = v29
                                    }
                                    v2 += v4
                                default:
                                    v30 := GenS4{F0: 873}
                                    _ = v30
                                }
                            }
                            v4 -= v2
                        case "hello":
                            switch ("" + "\n\t") {
                            case "\n\t":
                                v31 := GenS2{F0: true}
                                _ = v31
                            case "hello":
                                if (666 <= 65535) {
                                    if len(v11) > 0 {
                                        v11[0] = 127
                                    }
                                } else {
                                    v32 := GenS5{F0: false, F1: -31, F2: v4}
                                    v18 -= (v2 >> ((42) & 63))
                                    v20 += v20
                                    _ = v32
                                }
                            }
                        default:
                            if false {
                                var v33 int64 = -93
                                switch "" {
                                case "a":
                                    if true {
                                        v4 = 42
                                    } else {
                                        switch (fi12 + fi7) {
                                        case -9223372036854775808:
                                            if ((100 >= 1) || false) {
                                                var v34 int32 = 105
                                                _ = v34
                                            }
                                        default:
                                            if (v33 < (-(v33))) {
                                                var v35 uint16 = 904
                                                _ = v35
                                            } else {
                                                if true {
                                                    switch "" {
                                                    case "hello":
                                                        v36 := GenS6{F0: (true || (0 <= 613)), F1: 685, F2: 42}
                                                        _ = v36
                                                    case "\n\t":
                                                        if true {
                                                            v37 := GenS2{F0: false}
                                                            v5 += v20
                                                            _ = v37
                                                        }
                                                    case "a":
                                                        switch (fi12 | fi12) {
                                                        case 4294967295:
                                                            var v38 float32 = 3.14159
                                                            _ = v38
                                                        case 128:
                                                            v39 := GenS2{F0: true}
                                                            _ = v39
                                                        case 32768:
                                                            v18 -= (v2 << ((701) & 63))
                                                        }
                                                    }
                                                    v16 += (v16 ^ v16)
                                                    v40 := GenS2{F0: (42 < 42)}
                                                    v41 := GenS2{F0: (!false)}
                                                    v5 = (v16 - 173)
                                                    _ = v40
                                                    _ = v41
                                                }
                                                v2 += (v2 + 32767)
                                            }
                                        }
                                        if true {
                                            var v42 float32 = (3.14159 * 0.1)
                                            _ = v42
                                        }
                                        var v43 uint = 980
                                        _ = v43
                                    }
                                default:
                                    v33 = -9223372036854775808
                                }
                                if len(v11) > 0 {
                                    v11[0] = 593
                                }
                                _ = v33
                            }
                        }
                    }
                    _ = v16
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    _ = len(v3)
                    v3[65535] = (1.5 <= 1e10)
                    v3[2] = false
                }
                v11 = append(v11, (fi7 >> ((255) & 63)))
                _ = fi12
            }
            _ = fi7
            _ = v11
        }
    }
    v44 := []int{32768, 42, -469}
    v1[2147483648] = true
    if len(v44) > 0 {
        v44[0] = 4294967295
    }
    v5 += (v5 << ((335) & 63))
    if false {
        var v45 uint8 = 39
        _ = v45
    } else {
        for fi46 := 0; fi46 < 3; fi46++ {
            if (!(false || true)) {
                switch ("" + "foo bar") {
                case "":
                    switch fi46 {
                    case -985:
                        fmt.Println(v2)
                    case 128:
                        var v47 uint8 = 42
                        _ = v47
                    case 42:
                        fmt.Println(v5, v5, fi46)
                    default:
                        v2 = (v2 << ((805) & 63))
                    }
                case "foo bar":
                    v5 = (v5 / ((544) | 1))
                case "\n\t":
                    v48 := GenS7{F0: 127}
                    _ = v48
                }
                fmt.Println(v5)
                for fi49 := 0; fi49 < 6; fi49++ {
                    v44 = append(v44, 750)
                    v2 += (v2 - (^(v4)))
                    v50 := GenS4{F0: 451}
                    _ = fi49
                    _ = v50
                }
            }
            v2 = (v4 >> ((0) & 63))
            _ = fi46
        }
        var v51 int64 = -400
        if ((!false) && true) {
            for fi52 := 0; fi52 < 7; fi52++ {
                fmt.Println(v5)
                if false {
                    v53 := GenS2{F0: false}
                    fmt.Println(fi52, v5)
                    if len(v44) > 0 {
                        v44[0] = 490
                    }
                    _ = v53
                } else {
                    v54 := GenS1{F0: "foo bar"}
                    _ = v54
                }
                for fi55 := 0; fi55 < 5; fi55++ {
                    if (-1 < 0) {
                        var v56 int64 = 42
                        if true {
                            _ = len(v44)
                        }
                        _ = v56
                    } else {
                        fmt.Println(fi55, v51, v5)
                        if true {
                            if true {
                                switch (fi55 + 5) {
                                case 128:
                                    if (2 != 42) {
                                        var v57 bool = true
                                        if v57 {
                                            v3[42] = (23 != 128)
                                            switch ("\n\t" + "") {
                                            case "hello":
                                                var v58 int8 = -128
                                                _ = v58
                                            case "":
                                                if true {
                                                    v59 := v1[2147483648]
                                                    switch -573 {
                                                    case 548:
                                                        switch int(v51) {
                                                        case -2147483648:
                                                            v57 = (!v57)
                                                        case 9223372036854775807:
                                                            _ = len(v1)
                                                        }
                                                    case -535:
                                                        v57 = false
                                                    default:
                                                        if false {
                                                            if false {
                                                                var v60 int16 = (v2 << ((65535) & 63))
                                                                if v59 {
                                                                    v61 := GenS2{F0: v57}
                                                                    v62 := GenS6{F0: false, F1: 2, F2: 41}
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                                _ = v60
                                                            }
                                                        }
                                                    }
                                                    _ = v59
                                                }
                                            case "a":
                                                var v63 float32 = -2.25
                                                _ = v63
                                            }
                                        } else {
                                            switch (fi55 | (fi55 | (fi55 / ((128) | 1)))) {
                                            case 127:
                                                var v64 string = "\n\t"
                                                _ = v64
                                            case 393:
                                                v4 += (-(v2))
                                            case 42:
                                                var v65 uint = 18446744073709551615
                                                _ = v65
                                            default:
                                                switch fi52 {
                                                case -1:
                                                    if true {
                                                        var v66 string = "hello"
                                                        var v67 uint32 = 2
                                                        v5 += 2
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                case 32767:
                                                    v68 := GenS3{F0: 56, F1: "hello"}
                                                    _ = v68
                                                case 9223372036854775807:
                                                    var v69 int32 = -371
                                                    _ = v69
                                                default:
                                                    v51 += (v51 * v51)
                                                }
                                            }
                                            if (false && (65535 == 2147483648)) {
                                                if ((false || true) && (v57 && (!((1e10 - -2.25) >= -2.25)))) {
                                                    if v57 {
                                                        v57 = (1 == 2)
                                                    } else {
                                                        v57 = (("foo bar" < "a") || v57)
                                                        switch 127 {
                                                        case 0:
                                                            v70 := GenS8{F0: 0, F1: (-743 >= -405)}
                                                            _ = v70
                                                        case 326:
                                                            v3[492] = ("foo bar" == "a")
                                                            _ = len(v44)
                                                            v5 = 42
                                                        default:
                                                            var v71 float32 = (float32(-2.25) - 0.0)
                                                            if len(v44) >= 2 {
                                                                v44 = v44[1:len(v44)]
                                                            }
                                                            _ = v71
                                                        }
                                                        v72 := GenS1{F0: "foo bar"}
                                                        _ = v72
                                                    }
                                                    if false {
                                                        v57 = false
                                                        v73 := GenS8{F0: -128, F1: v57}
                                                        _ = v73
                                                    }
                                                    _ = len(v44)
                                                    var v74 uint8 = 114
                                                    fmt.Println(fi55, v2)
                                                    _ = v74
                                                } else {
                                                    var v75 uint8 = 0
                                                    _ = v75
                                                }
                                                fmt.Println(fi55)
                                            } else {
                                                fmt.Println(v57)
                                            }
                                            switch fi55 {
                                            case 42:
                                                var v76 int64 = v51
                                                _ = v76
                                            default:
                                                var v77 int32 = -2
                                                _ = v77
                                            }
                                        }
                                        fmt.Println(v5)
                                        _ = v57
                                    }
                                case 0:
                                    var v78 uint32 = 0
                                    _ = v78
                                }
                                v79 := GenS8{F0: -128, F1: true}
                                _ = v79
                            } else {
                                var v80 int = (fi52 - (fi55 - 1))
                                v81 := GenS2{F0: true}
                                _ = v80
                                _ = v81
                            }
                            v51 = (v51 << ((18446744073709551615) & 63))
                        }
                    }
                    v2 -= 0
                    if (42 >= 116) {
                        var v82 string = (("a" + "hello") + (("" + ((("" + "hello") + "\n\t") + "")) + "foo bar"))
                        v83 := GenS9{F0: 1, F1: -128, F2: v2}
                        _ = v82
                        _ = v83
                    } else {
                        var v84 int64 = v51
                        _ = v84
                    }
                    _ = fi55
                }
                _ = fi52
            }
            var v85 uint = 341
            v3[95] = false
            _ = v85
        }
        _ = v51
    }
    v86 := GenS9{F0: 2, F1: 0, F2: v2}
    var v87 int64 = 42
    var v88 int8 = 100
    var v89 float64 = 1.5
    fmt.Println(v4)
    var v90 int8 = (-(v88))
    v5 -= v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v5:", v5)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v86
    _ = v87
    _ = v88
    _ = v89
    _ = v90
}
