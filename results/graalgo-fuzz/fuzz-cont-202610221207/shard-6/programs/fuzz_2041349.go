package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 int32; F1 string; F2 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int }
func genFn0() int64 {
    if (!true) {
        v7 := []int{}
        _ = v7
    } else {
        var v8 int16 = -1
        _ = v8
    }
    v9 := GenS1{F0: 626}
    var v10 uint8 = 0
    var v11 uint8 = 219
    v10 += (v11 + (^(v11)))
    _ = v10
    _ = v11
    _ = v9
    return 444
}

func main() {
    v1 := []string{"a", "foo bar", "foo bar", ("a" + "\n\t")}
    var v2 int16 = -1
    var v3 uint64 = 2
    var v4 uint32 = 0
    var v5 float32 = 0.1
    var v6 uint16 = 32767
    var v12 int = 2147483647
    for fi13 := 0; fi13 < 18; fi13++ {
        switch (fi13 ^ -1) {
        case -444:
            if true {
                var v14 int32 = 2147483647
                v15 := GenS1{F0: v2}
                _ = v14
                _ = v15
            }
        default:
            if (!(false && ((^(v3)) <= 1))) {
                v5 += (0.0 + v5)
            } else {
                v6 -= (^(v6))
                for fi16 := 0; fi16 < 8; fi16++ {
                    v3 = 0
                    var v17 int16 = (v2 % (((^(v2))) | 1))
                    var v18 bool = false
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
        }
        if true {
            switch "" {
            case "a":
                v4 = 65535
            case "foo bar":
                if false {
                    if (!false) {
                        fmt.Println(v2, fi13)
                        v19 := GenS1{F0: 0}
                        _ = v19
                    } else {
                        fmt.Println(fi13)
                        var v20 uint64 = v3
                        var v21 bool = (true && (!true))
                        _ = v20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v12 = 9223372036854775807
                        var v23 int64 = int64(v3)
                        _ = fi22
                        _ = v23
                    }
                } else {
                    fmt.Println(v3)
                }
            case "\n\t":
                v1 = append(v1, "hello")
            }
            var v24 int64 = -1
            if false {
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                switch v12 {
                case 2147483647:
                    v6 -= v6
                case 162:
                    for fi25 := 0; fi25 < 12; fi25++ {
                        fmt.Println(v24, v2, v4)
                        switch "" {
                        case "\n\t":
                            fmt.Println(v3, fi13, v4)
                        }
                        _ = fi25
                    }
                case -982:
                    var v26 uint = 830
                    _ = v26
                }
                v24 = (v24 & int64(v6))
            } else {
                var v27 int = (v12 << ((1) & 63))
                v27 += (v12 ^ (fi13 | v27))
                var v28 int64 = (v24 / (((v24 | (v24 - v24))) | 1))
                _ = v27
                _ = v28
            }
            _ = v24
        }
        var v29 bool = true
        _ = fi13
        _ = v29
    }
    if (0 <= (-(v2))) {
        fmt.Println(v2, v12, v3)
    }
    fmt.Println(v4, v4)
    switch (("foo bar" + "a") + "hello") {
    case "foo bar":
        var v30 int64 = 455
        _ = v30
    default:
        switch "a" {
        case "":
            v6 -= v6
        case "\n\t":
            switch v12 {
            case 127:
                var v31 int = 127
                _ = v31
            case 2147483647:
                var v32 int16 = 1000
                _ = v32
            case -2147483648:
                v5 += v5
            default:
                switch (("foo bar" + "foo bar") + "") {
                case "foo bar":
                    if true {
                        v4 = v4
                        v4 += 2147483647
                    }
                case "hello":
                    if (true && true) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch "" {
                        case "a":
                            var v33 int8 = 103
                            _ = v33
                        case "":
                            v34 := GenS1{F0: (v2 * v2)}
                            _ = v34
                        case "\n\t":
                            fmt.Println(v5)
                        default:
                            switch (("hello" + ((("" + "\n\t") + ("\n\t" + (("\n\t" + ("hello" + "\n\t")) + "\n\t"))) + ("" + "foo bar"))) + "a") {
                            case "hello":
                                for fi35 := 0; fi35 < 2; fi35++ {
                                    var v36 int32 = -659
                                    v12 = int(v6)
                                    switch "hello" {
                                    case "foo bar":
                                        for fi37 := 0; fi37 < 7; fi37++ {
                                            v2 = v2
                                            if len(v1) > 0 {
                                                v1[0] = "foo bar"
                                            }
                                            if false {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v36 += (v36 ^ v36)
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v2 += (-(v2))
                                            v38 := GenS2{F0: 1}
                                            fmt.Println(v2, v2, fi35)
                                            _ = fi37
                                            _ = v38
                                        }
                                    default:
                                        v39 := GenS1{F0: 1000}
                                        _ = v39
                                    }
                                    _ = fi35
                                    _ = v36
                                }
                            case "a":
                                v6 -= 705
                            case "":
                                v40 := GenS1{F0: 400}
                                _ = v40
                            default:
                                if ("\n\t" == "a") {
                                    v12 = (v12 >> ((18446744073709551615) & 63))
                                }
                            }
                        }
                    }
                }
            }
        default:
            switch "foo bar" {
            case "":
                fmt.Println(v5, v5)
            }
        }
    }
    if (2147483648 <= v4) {
        if (false && (v3 >= v3)) {
            if len(v1) > 0 {
                v1[0] = ("\n\t" + "\n\t")
            }
            if true {
                if true {
                    var v41 bool = false
                    if v41 {
                        fmt.Println(v3, v5)
                    }
                    v1 = append(v1, ("foo bar" + ("\n\t" + "")))
                    _ = v41
                }
                if (3.14159 <= 1e10) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        if (!(-128 >= 42)) {
                            var v42 uint8 = 255
                            var v43 uint = 18446744073709551615
                            v44 := GenS3{F0: false}
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                    case "":
                        if len(v1) > 0 {
                            v1[0] = ("" + ("\n\t" + "\n\t"))
                        }
                    case "foo bar":
                        if len(v1) > 0 {
                            v1[0] = ("" + ("a" + "hello"))
                        }
                    }
                    fmt.Println(v3)
                }
                var v45 uint = 0
                _ = v45
            } else {
                switch v12 {
                case 1:
                    switch "hello" {
                    case "foo bar":
                        v3 = 691
                    default:
                        v5 = (((((v5 * v5) + (v5 + v5)) - ((v5 * (((1e10 + ((v5 + float32(1.5)) * ((v5 - v5) + (((0.0 * v5) / v5) - v5)))) / (1e10 - (float32(0.1) + v5))) * (v5 * v5))) + v5)) - (v5 * v5)) * (3.14159 * 0.1))
                    }
                }
                if false {
                    fmt.Println(v4, v2, v6)
                    v46 := GenS4{F0: (v4 % ((v4) | 1)), F1: -128}
                    for fi47 := 0; fi47 < 6; fi47++ {
                        v12 -= (v12 - (-(v12)))
                        for fi48 := 0; fi48 < 7; fi48++ {
                            v49 := GenS1{F0: 0}
                            switch 4294967295 {
                            case 9223372036854775807:
                                v4 += (v4 * (^(v4)))
                                v50 := GenS1{F0: (-(v2))}
                                _ = v50
                            case 33:
                                var v51 uint8 = 46
                                _ = v51
                            }
                            _ = fi48
                            _ = v49
                        }
                        for fi52 := 0; fi52 < 10; fi52++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi52
                        }
                        _ = fi47
                    }
                    _ = v46
                }
            }
            for fi53 := 0; fi53 < 18; fi53++ {
                if (!((!(false && false)) && (false || ((-9223372036854775808 >= -391) || true)))) {
                    var v54 int = v12
                    var v55 uint16 = v6
                    _ = v54
                    _ = v55
                } else {
                    v3 = 4294967295
                }
                for fi56 := 0; fi56 < 10; fi56++ {
                    v12 += fi53
                    v4 = v4
                    v1 = append(v1, ((("" + (("hello" + ("a" + ("\n\t" + ""))) + "foo bar")) + "\n\t") + ""))
                    var v57 uint32 = (^(v4))
                    if true {
                        var v58 int8 = 1
                        v2 -= 111
                        var v59 int32 = -1
                        _ = v58
                        _ = v59
                    }
                    _ = fi56
                    _ = v57
                }
                _ = fi53
            }
        } else {
            switch "a" {
            case "\n\t":
                v1 = append(v1, "\n\t")
            case "foo bar":
                var v60 uint32 = 364
                _ = v60
            }
            if false {
                for fi61 := 0; fi61 < 2; fi61++ {
                    var v62 uint32 = v4
                    v62 = (v4 % (((^(v62))) | 1))
                    _ = fi61
                    _ = v62
                }
                if (!true) {
                    var v63 uint16 = (v6 * (^(v6)))
                    v3 = 778
                    if true {
                        v1 = append(v1, "foo bar")
                        v6 = v6
                    } else {
                        v3 = v3
                        if false {
                            for fi64 := 0; fi64 < 10; fi64++ {
                                for fi65 := 0; fi65 < 1; fi65++ {
                                    v4 += v4
                                    var v66 uint64 = (^(v3))
                                    var v67 int8 = 42
                                    var v68 uint8 = 214
                                    _ = fi65
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                }
                                v69 := GenS5{F0: 0, F1: "foo bar", F2: genFn0()}
                                _ = fi64
                                _ = v69
                            }
                        } else {
                            v6 -= (^(v63))
                        }
                        for fi70 := 0; fi70 < 14; fi70++ {
                            for fi71 := 0; fi71 < 12; fi71++ {
                                switch v12 {
                                case 127:
                                    if true {
                                        v72 := GenS6{F0: 42}
                                        if false {
                                            v2 -= v2
                                        }
                                        _ = v72
                                    } else {
                                        fmt.Println(v4, fi71)
                                    }
                                case -2147483648:
                                    fmt.Println(v3, v3, v4)
                                case -9223372036854775808:
                                    if (!(v2 <= 543)) {
                                        fmt.Println(v6, v12, v4)
                                    } else {
                                        var v73 string = ("a" + "foo bar")
                                        _ = v73
                                    }
                                }
                                _ = fi71
                            }
                            for fi74 := 0; fi74 < 5; fi74++ {
                                fmt.Println(v5, fi74)
                                _ = fi74
                            }
                            _ = fi70
                        }
                    }
                    _ = v63
                } else {
                    var v75 uint64 = (v3 * (v3 ^ (v3 | 680)))
                    var v76 int8 = -1
                    _ = v75
                    _ = v76
                }
                for fi77 := 0; fi77 < 15; fi77++ {
                    fmt.Println(v6, fi77)
                    var v78 int64 = 349
                    fmt.Println(v3, v4, v4)
                    _ = fi77
                    _ = v78
                }
            } else {
                switch "" {
                case "hello":
                    fmt.Println(v6, v2)
                default:
                    if true {
                        _ = len(v1)
                        v79 := GenS7{F0: 0}
                        _ = v79
                    } else {
                        for fi80 := 0; fi80 < 5; fi80++ {
                            v3 -= v3
                            if (!true) {
                                fmt.Println(v5)
                                if len(v1) > 0 {
                                    v1[0] = "a"
                                }
                                var v81 bool = (-755 != 971)
                                _ = v81
                            }
                            v1 = append(v1, "a")
                            _ = fi80
                        }
                        var v82 string = "\n\t"
                        _ = v82
                    }
                }
                _ = len(v1)
            }
        }
        if true {
            var v83 int8 = 127
            v12 = -652
            var v84 int16 = 1000
            _ = v83
            _ = v84
        } else {
            if false {
                fmt.Println(v6, v6, v6)
                var v85 uint = 42
                _ = v85
            } else {
                if (0.0 >= 1.5) {
                    v86 := GenS8{F0: (v12 - -494)}
                    _ = v86
                }
            }
            var v87 uint64 = v3
            var v88 uint64 = v3
            _ = v87
            _ = v88
        }
    } else {
        if (float64(1e10) <= 3.14159) {
            v12 = 682
        } else {
            fmt.Println(v4, v12, v2)
        }
    }
    for fi89 := 0; fi89 < 14; fi89++ {
        v90 := GenS2{F0: -1}
        if (!false) {
            var v91 uint = 543
            var v92 int8 = int8(v6)
            _ = v91
            _ = v92
        }
        v4 -= v4
        _ = fi89
        _ = v90
    }
    for fi93 := 0; fi93 < 5; fi93++ {
        if len(v1) > 0 {
            v1[0] = ""
        }
        _ = fi93
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
