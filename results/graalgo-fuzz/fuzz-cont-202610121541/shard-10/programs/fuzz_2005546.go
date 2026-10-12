package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 bool }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 uint8; F1 int64; F2 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 bool; F1 int64 }
type GenS10 struct { F0 uint16; F1 int64; F2 uint8 }
func genFn0(p0 int64, p1 uint64) string {
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 uint = 4294967295
        _ = fi4
        _ = v5
    }
    p0 = 0
    switch 128 {
    case 0:
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int16 = 42
            _ = fi6
            _ = v7
        }
    case 9223372036854775807:
        if false {
            var v8 uint8 = 1
            var v9 int16 = 1000
            var v10 uint32 = 0
            _ = v10
            _ = v8
            _ = v9
        } else {
            v11 := []int{566, 229, 1}
            _ = v11
        }
    case -342:
        v12 := []int{}
        _ = v12
    }
    _ = p0
    _ = p1
    return "a"
}
func genFn1() uint8 {
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint64 = 172
        _ = fi13
        _ = v14
    }
    var v15 float64 = 1.5
    _ = v15
    return 42
}

func main() {
    v1 := []uint32{2147483647, 159, 4294967295}
    v2 := []string{"", "", (("" + "\n\t") + ("" + ("a" + "")))}
    var v3 int8 = 100
    v3 = (^(v3))
    if true {
        for fi16 := 0; fi16 < 5; fi16++ {
            if ((1 >= 1) && false) {
                switch "a" {
                case "a":
                    v3 = (v3 ^ 1)
                case "hello":
                    if false {
                        fmt.Println(v3)
                        for fi17 := 0; fi17 < 9; fi17++ {
                            v3 = (^(v3))
                            var v18 int16 = 42
                            var v19 int = 32768
                            if (true || true) {
                                v18 += 32767
                            } else {
                                var v20 float64 = (float64(-2.25) / (3.14159 / 1.0))
                                _ = v20
                            }
                            v19 = (v19 + (v19 / ((fi17) | 1)))
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        var v21 uint16 = 921
                        _ = v21
                    }
                default:
                    for fi22 := 0; fi22 < 2; fi22++ {
                        v23 := GenS1{F0: 1, F1: 255, F2: ((v3 + v3) == (v3 << ((248) & 63)))}
                        v24 := GenS1{F0: 255, F1: 222, F2: false}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
                var v25 int8 = v3
                _ = v25
            } else {
                var v26 int64 = -894
                _ = v26
            }
            _ = fi16
        }
    }
    switch 127 {
    case 128:
        var v27 uint16 = 0
        _ = v27
    case -644:
        for fi28 := 0; fi28 < 6; fi28++ {
            v3 = (v3 >> ((18446744073709551615) & 63))
            fmt.Println(v3, v3)
            if (!false) {
                v29 := GenS1{F0: 930, F1: 188, F2: false}
                if (907 > 1) {
                    for fi30 := 0; fi30 < 5; fi30++ {
                        v3 = 44
                        var v31 uint32 = 2147483647
                        switch int(v3) {
                        case 127:
                            if false {
                                if (37 >= 255) {
                                    var v32 int64 = 9223372036854775807
                                    var v33 uint64 = 1
                                    _ = v32
                                    _ = v33
                                } else {
                                    if false {
                                        var v34 int32 = -1
                                        var v35 bool = true
                                        _ = v34
                                        _ = v35
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = ("foo bar" + ((((("a" + "") + ("" + ("hello" + "foo bar"))) + "") + ("a" + "\n\t")) + "hello"))
                                        }
                                        var v36 uint64 = 667
                                        if len(v2) > 0 {
                                            v2[0] = "\n\t"
                                        }
                                        _ = v36
                                    }
                                }
                                var v37 float32 = 0.1
                                if false {
                                    fmt.Println(fi28, fi28, v37)
                                    fmt.Println(v37)
                                }
                                _ = v37
                            }
                        }
                        _ = fi30
                        _ = v31
                    }
                    if false {
                        v3 -= v3
                    }
                    if (!true) {
                        fmt.Println(fi28)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        v1 = append(v1, 2147483647)
                        var v38 uint32 = 2
                        _ = v38
                    }
                } else {
                    v3 += 0
                    v3 = v3
                }
                var v39 bool = false
                _ = v29
                _ = v39
            } else {
                v40 := GenS2{F0: 878, F1: 204, F2: 42}
                for fi41 := 0; fi41 < 19; fi41++ {
                    v42 := GenS3{F0: 0, F1: 42}
                    _ = fi41
                    _ = v42
                }
                _ = v40
            }
            _ = fi28
        }
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    v43 := GenS4{F0: -9223372036854775808, F1: ("\n\t" + "a")}
    for fi44 := 0; fi44 < 18; fi44++ {
        switch "a" {
        case "hello":
            fmt.Println(fi44, v3)
        }
        for fi45 := 0; fi45 < 16; fi45++ {
            if (542 < 951) {
                v3 = 100
                var v46 uint32 = 2147483647
                _ = v46
            }
            _ = fi45
        }
        var v47 uint32 = 65535
        _ = fi44
        _ = v47
    }
    switch -691 {
    case 4294967295:
        v48 := GenS2{F0: 0, F1: 128, F2: 398}
        _ = v48
    case 32768:
        for fi49 := 0; fi49 < 11; fi49++ {
            v50 := GenS4{F0: 0, F1: "\n\t"}
            var v51 int8 = v3
            _ = fi49
            _ = v50
            _ = v51
        }
    default:
        for fi52 := 0; fi52 < 4; fi52++ {
            v53 := GenS1{F0: 65535, F1: 2, F2: (!true)}
            if false {
                v54 := GenS1{F0: 650, F1: 215, F2: (32767 <= 822)}
                if (genFn1() > 0) {
                    if true {
                        var v55 int = -9223372036854775808
                        _ = v55
                    }
                    if (true && false) {
                        fmt.Println(fi52)
                    }
                }
                v3 = (v3 + v3)
                _ = v54
            }
            _ = fi52
            _ = v53
        }
    }
    v3 -= (v3 / ((v3) | 1))
    var v56 float64 = 0.0
    var v57 float64 = (v56 / ((((v56 * float64(-2.25)) + (0.1 + (v56 * (v56 * (1.5 - v56))))) + 1.5) + v56))
    v58 := GenS5{F0: 128, F1: 9223372036854775807, F2: 500}
    v59 := GenS4{F0: 42, F1: ("" + "foo bar")}
    for fi60 := 0; fi60 < 7; fi60++ {
        v57 = v57
        fmt.Println(v3, fi60, v3)
        _ = fi60
    }
    switch "hello" {
    case "hello":
        v3 -= -82
    default:
        for fi61 := 0; fi61 < 5; fi61++ {
            switch ("\n\t" + ("" + "")) {
            case "a":
                v62 := GenS1{F0: 4294967295, F1: 1, F2: (42 > 2)}
                _ = v62
            case "hello":
                v63 := GenS4{F0: -9223372036854775808, F1: "\n\t"}
                _ = v63
            default:
                if len(v2) > 0 {
                    v2[0] = ("hello" + "\n\t")
                }
            }
            v1 = append(v1, 61)
            _ = fi61
        }
    }
    for fi64 := 0; fi64 < 9; fi64++ {
        v65 := GenS6{F0: 128}
        for fi66 := 0; fi66 < 14; fi66++ {
            v57 = v56
            _ = fi66
        }
        _ = fi64
        _ = v65
    }
    switch int(v3) {
    case 222:
        if true {
            v67 := GenS1{F0: 1, F1: 42, F2: false}
            _ = len(v1)
            _ = v67
        } else {
            v56 -= float64(3.14159)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if ((((98 < 2) && false) && true) && false) {
                var v68 int64 = 42
                if (!false) {
                    v57 += v57
                } else {
                    var v69 uint8 = 42
                    v70 := GenS6{F0: (v69 >> ((45) & 63))}
                    _ = v69
                    _ = v70
                }
                if (!true) {
                    switch ("" + ("foo bar" + "")) {
                    case "foo bar":
                        fmt.Println(v3, v57)
                    case "":
                        v71 := GenS7{F0: "a", F1: "foo bar"}
                        _ = v71
                    default:
                        if (("a" + "hello") >= "\n\t") {
                            for fi72 := 0; fi72 < 9; fi72++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v3 = v3
                                if (242 != 42) {
                                    for fi73 := 0; fi73 < 7; fi73++ {
                                        if (!(255 < genFn1())) {
                                            v57 = (float64(0.0) - ((((v56 / v56) - v57) / 0.1) / (0.1 * float64(3.14159))))
                                            var v74 string = ("" + "hello")
                                            v74 = v74
                                            v74 = v74
                                            if len(v1) > 0 {
                                                v1[0] = 1
                                            }
                                            _ = v74
                                        } else {
                                            v3 = v3
                                            v56 -= (v56 * v56)
                                            var v75 int8 = -110
                                            _ = v75
                                        }
                                        v56 = ((1.5 * 0.1) / (v56 * 1e10))
                                        v76 := GenS2{F0: 448, F1: 2, F2: 4294967295}
                                        _ = len(v2)
                                        _ = fi73
                                        _ = v76
                                    }
                                } else {
                                    v2 = append(v2, (("" + "a") + "a"))
                                    var v77 uint32 = 4294967295
                                    _ = v77
                                }
                                _ = fi72
                            }
                            if false {
                                v78 := GenS6{F0: 66}
                                var v79 float64 = v57
                                v57 = (1e10 * float64((3.14159 * 1.5)))
                                _ = v78
                                _ = v79
                            } else {
                                switch "a" {
                                case "\n\t":
                                    if false {
                                        v80 := GenS8{F0: 42}
                                        fmt.Println(v3, v3)
                                        var v81 string = ""
                                        _ = v80
                                        _ = v81
                                    }
                                case "hello":
                                    var v82 uint8 = 0
                                    _ = v82
                                case "a":
                                    if true {
                                        for fi83 := 0; fi83 < 17; fi83++ {
                                            v57 = v56
                                            _ = fi83
                                        }
                                        _ = len(v2)
                                    }
                                }
                                switch 127 {
                                case 32768:
                                    var v84 uint32 = 954
                                    _ = v84
                                default:
                                    for fi85 := 0; fi85 < 14; fi85++ {
                                        var v86 bool = (-393 <= 1)
                                        if (v86 && v86) {
                                            switch (fi85 << ((133) & 63)) {
                                            case 329:
                                                fmt.Println(v56, v56, v3)
                                            case 42:
                                                if ((((("foo bar" + "foo bar") + ("" + "foo bar")) + ("foo bar" + "\n\t")) + "") < ("a" + "foo bar")) {
                                                    var v87 uint16 = 0
                                                    _ = v87
                                                } else {
                                                    if v86 {
                                                        if (848 > 65535) {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            fmt.Println(v57)
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    } else {
                                                        for fi88 := 0; fi88 < 5; fi88++ {
                                                            _ = len(v2)
                                                            if (810 != -1) {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                if (!((v3 << ((42) & 63)) != (v3 * (v3 - (v3 | (v3 + v3)))))) {
                                                                    var v89 uint = 65535
                                                                    v57 = (v56 - float64((3.14159 + 3.14159)))
                                                                    fmt.Println(v68)
                                                                    _ = v89
                                                                }
                                                            } else {
                                                                fmt.Println(v86, fi88, v68)
                                                                v3 = v3
                                                            }
                                                            _ = fi88
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = 65535
                                                        }
                                                        if v86 {
                                                            if v86 {
                                                                switch 42 {
                                                                case 32768:
                                                                    v90 := GenS7{F0: "hello", F1: "foo bar"}
                                                                    _ = v90
                                                                case 128:
                                                                    if ((v3 >> ((255) & 63)) == v3) {
                                                                        v56 -= v57
                                                                        v91 := GenS9{F0: false, F1: 1}
                                                                        _ = v91
                                                                    }
                                                                }
                                                                var v92 uint64 = 66
                                                                v1 = append(v1, 42)
                                                                _ = v92
                                                            }
                                                            fmt.Println(v3, v3, v68)
                                                        }
                                                    }
                                                    if true {
                                                        var v93 string = "\n\t"
                                                        _ = v93
                                                    } else {
                                                        _ = len(v2)
                                                    }
                                                    for fi94 := 0; fi94 < 6; fi94++ {
                                                        var v95 string = (("foo bar" + "hello") + "")
                                                        v96 := GenS10{F0: 820, F1: 681, F2: 0}
                                                        _ = fi94
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                }
                                            }
                                        }
                                        var v97 bool = (!(956 < 0))
                                        _ = fi85
                                        _ = v86
                                        _ = v97
                                    }
                                }
                                for fi98 := 0; fi98 < 9; fi98++ {
                                    var v99 int64 = v68
                                    v57 = 0.0
                                    fmt.Println(fi98)
                                    _ = fi98
                                    _ = v99
                                }
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            v3 = (v3 % ((v3) | 1))
                            v68 = v68
                        }
                    }
                    var v100 int = int(v3)
                    for fi101 := 0; fi101 < 11; fi101++ {
                        v3 = int8(v100)
                        _ = fi101
                    }
                    _ = v100
                }
                _ = v68
            } else {
                switch ("\n\t" + "hello") {
                case "\n\t":
                    v102 := GenS10{F0: 255, F1: -9223372036854775808, F2: 89}
                    _ = v102
                case "hello":
                    var v103 uint16 = 2
                    _ = v103
                default:
                    switch "hello" {
                    case "hello":
                        v57 -= (v57 * v56)
                    case "\n\t":
                        v2 = append(v2, "a")
                    case "foo bar":
                        var v104 int = 128
                        _ = v104
                    }
                }
                var v105 uint16 = 42
                _ = v105
            }
        }
    default:
        if len(v2) > 0 {
            v2[0] = ("hello" + "")
        }
    }
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    fmt.Println(v57, v56)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
