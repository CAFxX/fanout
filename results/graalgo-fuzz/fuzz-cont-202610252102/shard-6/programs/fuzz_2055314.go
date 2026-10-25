package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int; F1 string; F2 uint }
type GenS4 struct { F0 int32; F1 uint16; F2 string }
type GenS5 struct { F0 int32; F1 bool; F2 bool }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 int32; F1 bool }
type GenS8 struct { F0 uint8 }
func genFn0() uint8 {
    var v7 string = ((("a" + "foo bar") + ((("a" + (("foo bar" + "foo bar") + "\n\t")) + "\n\t") + "")) + "foo bar")
    v8 := GenS1{F0: 100, F1: 2}
    _ = v7
    _ = v8
    return 190
}
func genFn1(p0 string, p1 uint16) int32 {
    for fi9 := 0; fi9 < 13; fi9++ {
        switch (fi9 - (fi9 - (fi9 + -135))) {
        case 9223372036854775807:
            var v10 uint32 = 65535
            _ = v10
        case 2147483647:
            p0 = p0
        }
        switch p0 {
        case "hello":
            var v11 int = fi9
            _ = v11
        case "a":
            switch 127 {
            case 128:
                var v12 int64 = -183
                _ = v12
            case 127:
                switch fi9 {
                case -9223372036854775808:
                    v13 := []int{(-(fi9)), fi9, 268, (^(fi9))}
                    _ = v13
                case 2147483647:
                    fmt.Println(fi9)
                default:
                    var v14 uint = 409
                    _ = v14
                }
            default:
                switch p0 {
                case "hello":
                    var v15 float64 = 3.14159
                    _ = v15
                default:
                    var v16 int = (fi9 << ((255) & 63))
                    _ = v16
                }
            }
        case "\n\t":
            p0 = (p0 + p0)
        default:
            p0 = p0
        }
        _ = fi9
    }
    var v17 int8 = 96
    p1 = (p1 / ((p1) | 1))
    var v18 uint32 = 65535
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 558
}

func main() {
    v1 := []string{"a"}
    v2 := []uint32{65535, 65535, 1}
    var v3 uint = 701
    var v4 int = 128
    v5 := make(map[uint32]string)
    v6 := make(map[string]int)
    var v19 int64 = -9223372036854775808
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 int64 = (v19 / ((912) | 1))
        fmt.Println(v21, v19, v4)
        _ = fi20
        _ = v21
    }
    fmt.Println(v3, v4, v3)
    v3 += v3
    switch v4 {
    case 4294967295:
        v19 += v19
    case 400:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v22 int = (^(v4))
    if true {
        v23 := GenS2{F0: 2, F1: 2}
        switch "a" {
        case "":
            var v24 int64 = (^(v19))
            _ = v24
        case "hello":
            v6["foo bar"] = (^(v22))
        }
        _ = v23
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        v19 = v19
        fmt.Println(v4, fi25)
        _ = fi25
    }
    fmt.Println(v4, v3, v19)
    fmt.Println(v19, v3, v22)
    var v26 uint16 = 0
    switch "foo bar" {
    case "\n\t":
        var v27 int32 = 1
        _ = v27
    case "a":
        switch "a" {
        case "a":
            v28 := GenS2{F0: 4294967295, F1: v26}
            _ = v28
        case "":
            if true {
                v19 = (v19 & 1)
                v4 -= v22
                var v29 uint = 1
                _ = v29
            } else {
                v5[2147483647] = ("\n\t" + "a")
            }
        case "\n\t":
            v30 := GenS3{F0: (v4 << ((v3) & 63)), F1: ("hello" + "foo bar"), F2: 42}
            _ = v30
        default:
            switch 2147483647 {
            case 42:
                var v31 uint8 = 166
                _ = v31
            case 0:
                var v32 int32 = 2147483647
                _ = v32
            }
        }
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        v4 = 32767
        for fi34 := 0; fi34 < 5; fi34++ {
            var v35 uint32 = 276
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    if true {
        var v36 float32 = (1.5 / 1.0)
        _ = v36
    } else {
        if (650 == 2) {
            if (65535 < 42) {
                v26 -= v26
                var v37 float64 = (-2.25 / -2.25)
                if false {
                    v22 = 101
                    if (v3 != v3) {
                        if (!((v19 <= (^(v19))) || (1 != 0))) {
                            var v38 int64 = 822
                            if true {
                                v39 := GenS3{F0: (v4 & v4), F1: ("a" + "a"), F2: (v3 >> ((2) & 63))}
                                if ((v26 << ((42) & 63)) >= v26) {
                                    if false {
                                        if (0 < 1000) {
                                            var v40 uint32 = 2147483647
                                            for fi41 := 0; fi41 < 8; fi41++ {
                                                v4 = (v4 + (-(v22)))
                                                for fi42 := 0; fi42 < 19; fi42++ {
                                                    var v43 int8 = 42
                                                    if false {
                                                        v44 := GenS4{F0: genFn1("a", v26), F1: 665, F2: ((("\n\t" + "foo bar") + "") + "a")}
                                                        v45 := GenS3{F0: int(v43), F1: "hello", F2: (v3 & 65535)}
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                    v46 := GenS5{F0: 1, F1: (1 >= genFn1(("a" + (("" + "\n\t") + "hello")), v26)), F2: (!false)}
                                                    v47 := GenS1{F0: v43, F1: 69}
                                                    v1 = append(v1, ("hello" + "\n\t"))
                                                    _ = fi42
                                                    _ = v43
                                                    _ = v46
                                                    _ = v47
                                                }
                                                _ = fi41
                                            }
                                            _ = v40
                                        } else {
                                            v48 := GenS5{F0: -1, F1: false, F2: true}
                                            if true {
                                                var v49 uint16 = 65535
                                                if ((^(v49)) < (v26 % (((v26 - (v49 / (((v49 - v49)) | 1)))) | 1))) {
                                                    var v50 int32 = -1
                                                    _ = v50
                                                } else {
                                                    for fi51 := 0; fi51 < 11; fi51++ {
                                                        if ((((v37 + (v37 - ((v37 * v37) / v37))) + (v37 / (v37 * float64(((1e10 * 0.0) / ((1.5 * 3.14159) * -2.25)))))) / (v37 * 1e10)) == ((1e10 + ((0.0 + (((v37 / v37) - 1e10) - (v37 - v37))) * float64(-2.25))) / -2.25)) {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                fmt.Println(v49, v37)
                                                            case "hello":
                                                                _ = len(v5)
                                                            case "":
                                                                fmt.Println(v38, v19)
                                                            }
                                                            var v52 int = v22
                                                            v1 = append(v1, "hello")
                                                            _ = v52
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = 789
                                                        }
                                                        _ = fi51
                                                    }
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v26 = 2
                                                    }
                                                }
                                                _ = v49
                                            } else {
                                                v37 -= float64((((0.1 / float32((-2.25 / v37))) / (1e10 * 1e10)) / float32(3.14159)))
                                                v53 := GenS6{F0: 18446744073709551615, F1: true}
                                                var v54 int = 9223372036854775807
                                                _ = v53
                                                _ = v54
                                            }
                                            for fi55 := 0; fi55 < 1; fi55++ {
                                                fmt.Println(v38, v37)
                                                if false {
                                                    if false {
                                                        fmt.Println(v37, v38, v26)
                                                        v3 += v3
                                                    }
                                                } else {
                                                    var v56 int32 = 0
                                                    v57 := GenS7{F0: 42, F1: (!(true || (!(-1 == 0))))}
                                                    _ = v56
                                                    _ = v57
                                                }
                                                _ = fi55
                                            }
                                            _ = v48
                                        }
                                        switch (-(v22)) {
                                        case -369:
                                            fmt.Println(v3, v22)
                                        case -9223372036854775808:
                                            fmt.Println(v4)
                                        case 42:
                                            var v58 float64 = (((((float64(((-2.25 * float32(1.5)) - 3.14159)) / float64(1e10)) + v37) / 0.1) * (v37 / -2.25)) / float64(float32((float64((float32((v37 / v37)) - 1.5)) / (v37 * float64(float32((float64(1.5) * 1.5))))))))
                                            _ = v58
                                        default:
                                            v22 = 896
                                        }
                                    }
                                    var v59 uint32 = 2
                                    _ = v59
                                }
                                v4 = (v4 % ((v22) | 1))
                                _ = v39
                            } else {
                                switch "hello" {
                                case "foo bar":
                                    v26 += 42
                                }
                                v60 := GenS7{F0: 42, F1: (1e10 < (0.0 * 0.0))}
                                _ = v60
                            }
                            var v61 float32 = 0.1
                            _ = v38
                            _ = v61
                        } else {
                            fmt.Println(v22)
                            v37 += (v37 - 3.14159)
                        }
                        var v62 int16 = 1000
                        _ = v62
                    } else {
                        v1 = append(v1, "foo bar")
                        v63 := GenS1{F0: -128, F1: 1}
                        for fi64 := 0; fi64 < 1; fi64++ {
                            if (float32((v37 + (float64(3.14159) / 1.0))) == -2.25) {
                                v26 += 288
                            } else {
                                v65 := GenS7{F0: 0, F1: ((((v26 >> ((1) & 63)) < 0) || ((^(v3)) > (v3 - (v3 * (v3 / ((388) | 1)))))) || true)}
                                var v66 uint16 = v26
                                _ = v65
                                _ = v66
                            }
                            var v67 string = "foo bar"
                            fmt.Println(v22)
                            _ = fi64
                            _ = v67
                        }
                        _ = v63
                    }
                    var v68 float64 = ((0.0 / -2.25) - ((v37 + (1.5 + (v37 * ((v37 - v37) - v37)))) / 0.1))
                    _ = v68
                }
                _ = v37
            } else {
                for fi69 := 0; fi69 < 6; fi69++ {
                    var v70 int8 = -128
                    _ = fi69
                    _ = v70
                }
                v3 -= (v3 << ((647) & 63))
            }
        } else {
            v22 -= v22
            if false {
                for fi71 := 0; fi71 < 2; fi71++ {
                    fmt.Println(v19, v26, v19)
                    v19 += v19
                    _ = fi71
                }
                var v72 uint32 = 2147483647
                if true {
                    v3 += (v3 * 4294967295)
                }
                _ = v72
            } else {
                v19 = v19
            }
        }
        var v73 string = "\n\t"
        var v74 uint16 = (^(v26))
        _ = v73
        _ = v74
    }
    for fi75 := 0; fi75 < 14; fi75++ {
        var v76 int16 = -729
        for fi77 := 0; fi77 < 18; fi77++ {
            switch 2147483647 {
            case 4294967295:
                fmt.Println(fi75, v4, v3)
            case 9223372036854775807:
                v19 -= (v19 << ((v3) & 63))
            }
            v4 += v22
            v78 := GenS8{F0: 127}
            fmt.Println(fi77, v76)
            var v79 float32 = 3.14159
            _ = fi77
            _ = v78
            _ = v79
        }
        var v80 float32 = -2.25
        _ = fi75
        _ = v76
        _ = v80
    }
    var v81 int64 = (v19 / ((v19) | 1))
    if (true || true) {
        for fi82 := 0; fi82 < 18; fi82++ {
            v19 += (v81 * -1)
            _ = fi82
        }
    } else {
        if (2147483648 < 1) {
            var v83 uint = (v3 / ((0) | 1))
            _ = v83
        }
        v84 := v5[1]
        for fi85 := 0; fi85 < 15; fi85++ {
            if (0.0 >= 3.14159) {
                fmt.Println(v19)
                for fi86 := 0; fi86 < 2; fi86++ {
                    fmt.Println(v81, fi85, v4)
                    _ = fi86
                }
            }
            switch (fi85 << ((v3) & 63)) {
            case -840:
                var v87 int32 = -2147483648
                _ = v87
            default:
                if (true && (42 != 0)) {
                    var v88 int32 = 0
                    v19 += (v81 & (v19 & (v81 * (v19 / ((9223372036854775807) | 1)))))
                    switch v84 {
                    case "\n\t":
                        var v89 int16 = 0
                        _ = v89
                    case "a":
                        fmt.Println(v81, v3, v3)
                    case "foo bar":
                        v81 = 382
                    }
                    _ = v88
                }
            }
            _ = fi85
        }
        _ = v84
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v81:", v81)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
}
