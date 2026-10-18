package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int64; F2 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int32; F2 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 int64; F2 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    v3 := []int{0, 4294967295, 32768, 1}
    var v4 int = 32767
    var v5 float32 = 1.5
    _ = len(v1)
    if false {
        v4 = v4
        _ = len(v3)
    } else {
        if (42 != 1) {
            _ = len(v2)
        } else {
            switch "" {
            case "\n\t":
                fmt.Println(v5, v4)
            case "foo bar":
                var v6 int8 = 1
                _ = v6
            case "":
                v7 := GenS1{F0: 1, F1: 2147483647}
                _ = v7
            }
        }
        v8 := GenS1{F0: 2, F1: 4294967295}
        var v9 uint32 = 2
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if false {
            v11 := GenS2{F0: 42}
            _ = v11
        } else {
            var v12 uint64 = 791
            v12 = (v12 >> ((527) & 63))
            _ = v12
        }
        _ = fi10
    }
    if ((false && false) && ((true && (2147483648 != 1)) || (!(162 != 73)))) {
        v13 := GenS1{F0: 500, F1: 795}
        _ = v13
    }
    switch v4 {
    case 128:
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 float64 = float64((v5 + 0.0))
            _ = fi14
            _ = v15
        }
    case 2147483647:
        switch "\n\t" {
        case "":
            var v16 uint = 0
            _ = v16
        case "a":
            switch (v4 ^ (v4 | v4)) {
            case -2147483648:
                v17 := GenS3{F0: 890, F1: 0, F2: v4}
                _ = v17
            }
        default:
            fmt.Println(v4)
        }
    case 0:
        if false {
            v4 += v4
            for fi18 := 0; fi18 < 12; fi18++ {
                if false {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        if len(v3) > 0 {
                            v3[0] = (^(fi18))
                        }
                        v20 := GenS4{F0: 0}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v5 = 0.0
                        if (!true) {
                            v4 = (-(v4))
                            fmt.Println(v4, fi19, v5)
                            var v21 int64 = 1
                            if ((!true) && (42 == 1000)) {
                                var v22 uint64 = 212
                                if ((!false) && true) {
                                    var v23 int8 = -128
                                    v24 := GenS1{F0: 834, F1: 1}
                                    var v25 int8 = 100
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                } else {
                                    var v26 uint32 = 237
                                    fmt.Println(fi19)
                                    _ = v26
                                }
                                _ = v22
                            } else {
                                var v27 int16 = 709
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = -875
                                    }
                                }
                                _ = v27
                            }
                            v5 += 1e10
                            _ = v21
                        } else {
                            var v28 int = (v4 * (-(fi18)))
                            _ = v28
                        }
                        var v29 int32 = -879
                        v4 -= (fi18 / ((-321) | 1))
                        if (1 < 18446744073709551615) {
                            v4 -= 901
                        }
                        var v30 int64 = 956
                        var v31 int = int(v30)
                        v2[("" + "foo bar")] = ("a" + "foo bar")
                        switch fi19 {
                        case 2147483647:
                            v31 = int(v29)
                        case -449:
                            fmt.Println(v29, v4, v5)
                        case 208:
                            v31 = v31
                        }
                        _ = fi19
                        _ = v20
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                    v4 += fi18
                }
                _ = fi18
            }
            for fi32 := 0; fi32 < 7; fi32++ {
                var v33 uint16 = 800
                for fi34 := 0; fi34 < 15; fi34++ {
                    v4 = v4
                    _ = fi34
                }
                _ = len(v3)
                _ = fi32
                _ = v33
            }
        }
    default:
        for fi35 := 0; fi35 < 1; fi35++ {
            switch (v4 + v4) {
            case 1:
                var v36 uint16 = 65535
                _ = v36
            }
            _ = fi35
        }
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        v5 = v5
        for fi38 := 0; fi38 < 11; fi38++ {
            v39 := GenS1{F0: 42, F1: 42}
            var v40 int = (v4 - (fi38 * fi38))
            switch (("foo bar" + "\n\t") + "a") {
            case "foo bar":
                var v41 uint32 = 2
                _ = v41
            }
            _ = fi38
            _ = v39
            _ = v40
        }
        if true {
            var v42 int64 = 42
            v5 -= -2.25
            _ = v42
        } else {
            v4 -= 2147483647
            switch (fi37 + -9223372036854775808) {
            case 128:
                v3 = append(v3, (v4 >> ((4294967295) & 63)))
            case 32768:
                for fi43 := 0; fi43 < 2; fi43++ {
                    v4 = 618
                    var v44 uint64 = 0
                    if ((!true) && false) {
                        var v45 int16 = -1
                        _ = v45
                    } else {
                        v4 += fi37
                        v46 := GenS2{F0: -1}
                        if len(v3) > 0 {
                            v3[0] = (fi37 ^ 1)
                        }
                        _ = v46
                    }
                    _ = fi43
                    _ = v44
                }
            default:
                _ = len(v2)
            }
        }
        _ = fi37
    }
    v3 = append(v3, (v4 | -133))
    if (!true) {
        var v47 float32 = v5
        if true {
            var v48 int16 = 1000
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v49 := GenS4{F0: -896}
            _ = v48
            _ = v49
        } else {
            v47 = v47
            fmt.Println(v5, v5)
        }
        _ = v47
    } else {
        var v50 int16 = 42
        switch v4 {
        case 576:
            v51 := GenS4{F0: -1}
            _ = v51
        case -2147483648:
            fmt.Println(v4)
        case 4294967295:
            for fi52 := 0; fi52 < 9; fi52++ {
                var v53 uint32 = 0
                v4 = (v4 | 127)
                _ = fi52
                _ = v53
            }
        default:
            delete(v2, ("\n\t" + ("a" + ("\n\t" + ("\n\t" + ("a" + (("foo bar" + "") + ("" + ""))))))))
        }
        var v54 bool = (true && true)
        _ = v50
        _ = v54
    }
    switch (v4 % ((v4) | 1)) {
    case 42:
        v55 := GenS1{F0: 0, F1: 42}
        _ = v55
    default:
        if false {
            if true {
                switch v4 {
                case 9223372036854775807:
                    switch "foo bar" {
                    case "hello":
                        if (!true) {
                            v56 := GenS4{F0: 672}
                            _ = v56
                        } else {
                            var v57 int = (v4 + v4)
                            _ = v57
                        }
                    case "a":
                        v58 := GenS4{F0: 9223372036854775807}
                        _ = v58
                    case "\n\t":
                        v5 -= (-2.25 - v5)
                    default:
                        for fi59 := 0; fi59 < 16; fi59++ {
                            v60 := GenS1{F0: 42, F1: 4294967295}
                            for fi61 := 0; fi61 < 13; fi61++ {
                                var v62 uint64 = 1
                                _ = fi61
                                _ = v62
                            }
                            v63 := GenS5{F0: 110, F1: -2147483648, F2: "a"}
                            _ = fi59
                            _ = v60
                            _ = v63
                        }
                    }
                case -291:
                    for fi64 := 0; fi64 < 5; fi64++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi64
                    }
                case -1:
                    v65 := GenS5{F0: 10, F1: 2147483647, F2: "hello"}
                    _ = v65
                default:
                    for fi66 := 0; fi66 < 19; fi66++ {
                        var v67 int8 = 100
                        var v68 float64 = 0.1
                        _ = fi66
                        _ = v67
                        _ = v68
                    }
                }
            }
            var v69 float32 = ((1.5 - 0.1) - v5)
            _ = v69
        } else {
            switch "\n\t" {
            case "a":
                v1[2147483647] = 1
            case "foo bar":
                delete(v2, ("foo bar" + (("\n\t" + "foo bar") + "\n\t")))
            case "":
                for fi70 := 0; fi70 < 20; fi70++ {
                    var v71 uint8 = 127
                    v72 := GenS3{F0: 609, F1: 1, F2: v4}
                    var v73 uint16 = 0
                    _ = fi70
                    _ = v71
                    _ = v72
                    _ = v73
                }
            }
            if true {
                switch -9223372036854775808 {
                case 128:
                    var v74 string = ""
                    _ = v74
                case -1:
                    var v75 int32 = 42
                    _ = v75
                }
                if false {
                    var v76 uint64 = 804
                    _ = v76
                } else {
                    v5 = (v5 + 1.5)
                    var v77 string = "foo bar"
                    _ = v77
                }
                switch (-(v4)) {
                case -680:
                    var v78 bool = false
                    _ = v78
                case -558:
                    var v79 uint16 = 65535
                    _ = v79
                default:
                    v5 = 0.1
                }
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v5 = (float32((0.0 / 1.5)) * (v5 - (v5 - 0.0)))
                for fi80 := 0; fi80 < 2; fi80++ {
                    for fi81 := 0; fi81 < 4; fi81++ {
                        var v82 uint8 = 166
                        switch (fi80 - (fi80 ^ v4)) {
                        case 1:
                            v82 = (v82 - 46)
                        case 127:
                            if false {
                                if (false || ((((0.1 + (0.0 + (1.5 * (float64(float32(((-2.25 + float64(1.5)) / 1e10))) + 3.14159)))) * float64(((0.1 * v5) + v5))) / float64((v5 + (v5 - 0.1)))) > 0.0)) {
                                    if true {
                                        v83 := GenS2{F0: -310}
                                        v82 = v82
                                        _ = v83
                                    }
                                    fmt.Println(fi80, fi81)
                                    var v84 int32 = 578
                                    _ = v84
                                } else {
                                    v5 += v5
                                }
                            }
                        case 32768:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = fi81
                        _ = v82
                    }
                    if (((v5 * (((float32(1e10) + v5) * 0.1) - v5)) * ((v5 - (v5 * v5)) / 1.0)) >= v5) {
                        delete(v2, (("" + "") + "hello"))
                        if ((true || (0 != 1)) && true) {
                            v5 += v5
                            for fi85 := 0; fi85 < 17; fi85++ {
                                if false {
                                    var v86 bool = true
                                    _ = v86
                                } else {
                                    var v87 int = -641
                                    v4 = fi80
                                    if false {
                                        _ = len(v3)
                                        _ = len(v1)
                                    }
                                    _ = v87
                                }
                                _ = fi85
                            }
                        }
                        switch ("\n\t" + "hello") {
                        case "hello":
                            var v88 uint64 = 0
                            _ = v88
                        }
                    } else {
                        v4 -= fi80
                    }
                    _ = fi80
                }
            }
        }
    }
    v89 := GenS1{F0: 18446744073709551615, F1: 742}
    if true {
        v90 := GenS5{F0: 1, F1: -2147483648, F2: ("foo bar" + ("a" + "hello"))}
        var v91 uint16 = 255
        var v92 int = (v4 / (((v4 & v4)) | 1))
        _ = v90
        _ = v91
        _ = v92
    }
    if true {
        v93 := GenS6{F0: 2}
        var v94 int8 = 100
        var v95 int16 = 148
        _ = v93
        _ = v94
        _ = v95
    }
    v5 -= 1e10
    v96 := GenS7{F0: 763, F1: 42, F2: 0}
    v97 := GenS6{F0: 18446744073709551615}
    switch v4 {
    case -1:
        var v98 uint16 = 32767
        _ = v98
    case -970:
        var v99 int64 = -1
        _ = v99
    }
    v100 := v2[("hello" + ("" + ""))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v89
    _ = v96
    _ = v97
}
