package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int16 }
type GenS2 struct { F0 bool; F1 int8; F2 uint8 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int16 }

func main() {
    var v1 uint8 = 0
    var v2 bool = false
    v3 := make(map[uint32]bool)
    fmt.Println(v1, v1)
    var v4 float64 = 1e10
    if (("a" + "\n\t") > (("" + ("" + ("a" + ("" + "\n\t")))) + "")) {
        v1 -= v1
        switch "hello" {
        case "hello":
            for fi5 := 0; fi5 < 10; fi5++ {
                var v6 uint16 = 78
                v7 := []int{fi5, -2147483648, fi5, 32767}
                _ = fi5
                _ = v6
                _ = v7
            }
        case "":
            for fi8 := 0; fi8 < 15; fi8++ {
                fmt.Println(v1, v1, v2)
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 uint8 = (v1 % ((2) | 1))
                    _ = fi9
                    _ = v10
                }
                var v11 uint32 = 42
                _ = fi8
                _ = v11
            }
        case "foo bar":
            v12 := []int{2147483647}
            _ = v12
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            for fi14 := 0; fi14 < 10; fi14++ {
                if true {
                    v2 = v2
                    var v15 int8 = 42
                    _ = v15
                }
                if (!((!(!v2)) && true)) {
                    v4 = v4
                    var v16 int64 = 9223372036854775807
                    v17 := GenS1{F0: 42, F1: "\n\t", F2: 1000}
                    var v18 string = "\n\t"
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    var v19 uint = 2
                    v2 = (!v2)
                    var v20 int32 = -2147483648
                    _ = v19
                    _ = v20
                }
                _ = fi14
            }
            _ = fi13
        }
    } else {
        fmt.Println(v2, v2)
        if (!(("a" + "a") == "foo bar")) {
            v4 += (v4 / (0.1 + 1e10))
        }
        if v2 {
            var v21 int64 = 0
            var v22 uint16 = 372
            _ = v21
            _ = v22
        } else {
            fmt.Println(v4, v4, v1)
            v23 := GenS2{F0: (42 >= (^(v1))), F1: -128, F2: (^(v1))}
            if (!v2) {
                for fi24 := 0; fi24 < 9; fi24++ {
                    if false {
                        v2 = (v2 || (127 >= -621))
                        fmt.Println(fi24, v1)
                    } else {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            switch "\n\t" {
                            case "a":
                                v26 := []int{fi24, 9223372036854775807, fi24, fi25}
                                _ = v26
                            case "hello":
                                fmt.Println(v4, v1, v1)
                            case "":
                                switch "hello" {
                                case "":
                                    v2 = true
                                    var v27 uint16 = 2
                                    fmt.Println(v1)
                                    _ = v27
                                }
                            }
                            v28 := GenS1{F0: 1, F1: "a", F2: 32767}
                            switch fi24 {
                            case 128:
                                fmt.Println(v4)
                            case -1:
                                if v2 {
                                    var v29 float32 = 0.0
                                    if (42 != -181) {
                                        v30 := GenS1{F0: 1, F1: "foo bar", F2: 0}
                                        switch "hello" {
                                        case "\n\t":
                                            fmt.Println(fi24)
                                        case "hello":
                                            _ = len(v3)
                                        default:
                                            v31 := []int{fi24, 32768, 9223372036854775807, fi24}
                                            _ = v31
                                        }
                                        _ = v30
                                    }
                                    _ = v29
                                } else {
                                    var v32 int = (^(fi24))
                                    v33 := GenS1{F0: int16(fi24), F1: "a", F2: -1}
                                    _ = v32
                                    _ = v33
                                }
                            case 42:
                                v2 = v2
                            }
                            if v2 {
                                if false {
                                    var v34 uint = 554
                                    switch -2147483648 {
                                    case 32768:
                                        switch ("\n\t" + (((("" + ("foo bar" + ("" + "hello"))) + "\n\t") + "foo bar") + "\n\t")) {
                                        case "foo bar":
                                            var v35 uint = v34
                                            if v2 {
                                                v36 := GenS2{F0: (9223372036854775807 == 42), F1: -108, F2: v1}
                                                _ = v36
                                            }
                                            v35 = v34
                                            _ = v35
                                        }
                                    case 1:
                                        var v37 float64 = v4
                                        _ = v37
                                    default:
                                        var v38 float64 = 1e10
                                        _ = v38
                                    }
                                    switch fi25 {
                                    case -997:
                                        if ((!v2) || (false && (4294967295 != 65535))) {
                                            var v39 int64 = 42
                                            _ = v39
                                        } else {
                                            v2 = ((65535 != 449) || (false || true))
                                        }
                                    case 508:
                                        v40 := []int{fi24, fi24}
                                        _ = v40
                                    case -9223372036854775808:
                                        v41 := GenS3{F0: 1, F1: ((v34 < v34) && v2)}
                                        _ = v41
                                    default:
                                        v4 = float64(1e10)
                                    }
                                    _ = v34
                                }
                                fmt.Println(fi25, v1, fi25)
                                v42 := GenS3{F0: -1, F1: v2}
                                _ = v42
                            }
                            if (fi25 < 234) {
                                if v2 {
                                    var v43 int16 = 32767
                                    _ = v43
                                }
                            } else {
                                fmt.Println(v1, fi24, v4)
                                _ = len(v3)
                                if (798 == 70) {
                                    var v44 uint64 = 17
                                    var v45 uint16 = 65535
                                    var v46 uint = 0
                                    var v47 uint16 = 2
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                } else {
                                    v48 := []int{}
                                    v49 := GenS4{F0: -9223372036854775808, F1: 2, F2: 42}
                                    _ = v48
                                    _ = v49
                                }
                                switch ("foo bar" + "") {
                                case "hello":
                                    var v50 uint32 = 2147483647
                                    _ = v50
                                case "a":
                                    if true {
                                        v51 := []int{fi25, 4294967295}
                                        v2 = (255 < 469)
                                        v52 := GenS5{F0: (!(!(127 <= (fi25 + (fi25 % (((fi24 + (fi25 % (((fi25 / (((fi24 << ((255) & 63))) | 1))) | 1)))) | 1))))))}
                                        _ = v51
                                        _ = v52
                                    }
                                case "":
                                    v4 -= v4
                                default:
                                    v53 := GenS2{F0: (1 <= -168), F1: -128, F2: v1}
                                    _ = v53
                                }
                            }
                            switch (fi25 >> ((18446744073709551615) & 63)) {
                            case -90:
                                var v54 bool = v2
                                _ = v54
                            case 0:
                                if (v1 == v1) {
                                    if (!(0.0 < (0.0 - 1.5))) {
                                        var v55 float32 = (((float32(((-2.25 + (((v4 + v4) - 1.5) - 1e10)) - float64((1.5 * -2.25)))) * (-2.25 * 1.5)) + (-2.25 - 1e10)) * -2.25)
                                        var v56 float32 = v55
                                        _ = v55
                                        _ = v56
                                    }
                                    v57 := []int{32768, (fi24 * 32767), 128, fi24}
                                    _ = v57
                                } else {
                                    v58 := GenS6{F0: 69}
                                    switch (^(fi25)) {
                                    case 931:
                                        delete(v3, 199)
                                    case 127:
                                        var v59 float32 = 1e10
                                        _ = v59
                                    case 0:
                                        if v2 {
                                            var v60 float32 = float32((0.0 - v4))
                                            switch 2147483647 {
                                            case 128:
                                                if (((((float32(0.0) / v60) * 3.14159) > 0.1) && v2) && (!v2)) {
                                                    var v61 uint32 = 2147483647
                                                    if (!(false && (-32768 > 0))) {
                                                        var v62 uint16 = 65535
                                                        v63 := []int{4294967295}
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        v64 := []int{}
                                                        _ = v64
                                                    }
                                                    fmt.Println(v60, v61)
                                                    _ = v61
                                                } else {
                                                    v1 = 128
                                                }
                                            }
                                            v1 -= v1
                                            _ = v60
                                        }
                                    }
                                    _ = v58
                                }
                            }
                            if false {
                                var v65 uint = 1
                                var v66 int16 = int16(v1)
                                var v67 uint16 = 1
                                _ = v65
                                _ = v66
                                _ = v67
                            }
                            _ = fi25
                            _ = v28
                        }
                        if ("hello" <= "\n\t") {
                            var v68 uint16 = 918
                            _ = v68
                        }
                    }
                    var v69 int16 = 32767
                    var v70 uint8 = 1
                    _ = fi24
                    _ = v69
                    _ = v70
                }
                fmt.Println(v2, v2, v1)
            }
            _ = v23
        }
    }
    for fi71 := 0; fi71 < 7; fi71++ {
        v3[2147483648] = (713 > 1)
        if v2 {
            for fi72 := 0; fi72 < 11; fi72++ {
                v1 = v1
                _ = fi72
            }
            if (993 > 4294967295) {
                var v73 int = 127
                _ = v73
            } else {
                var v74 bool = true
                var v75 int16 = 0
                var v76 int = fi71
                _ = v74
                _ = v75
                _ = v76
            }
            v77 := GenS7{F0: -564}
            _ = v77
        }
        _ = fi71
    }
    v78 := GenS5{F0: v2}
    v2 = (95 >= -4)
    if false {
        v4 = 1.5
    } else {
        fmt.Println(v1, v1, v4)
        v1 = v1
        for fi79 := 0; fi79 < 12; fi79++ {
            var v80 uint = 821
            v81 := GenS8{F0: 1}
            _ = fi79
            _ = v80
            _ = v81
        }
    }
    fmt.Println(v4, v2)
    v1 -= v1
    var v82 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v78
    _ = v82
}
