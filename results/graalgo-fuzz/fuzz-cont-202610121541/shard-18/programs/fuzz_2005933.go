package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 int8; F2 int16 }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 int; F1 int64 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]bool)
    var v3 uint64 = 318
    v4 := []string{"a"}
    v5 := []uint32{1, uint32(v3)}
    v6 := make(map[int]bool)
    v7 := GenS1{F0: (v1 & 869)}
    fmt.Println(v3, v3)
    if true {
        switch -887 {
        case 2147483647:
            var v8 int16 = 1000
            _ = v8
        case 9223372036854775807:
            fmt.Println(v1, v3)
        default:
            v4 = append(v4, "foo bar")
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            fmt.Println(v1, v1)
            var v10 string = "\n\t"
            _ = fi9
            _ = v10
        }
    } else {
        var v11 float64 = (0.0 / float64(-2.25))
        v12 := GenS2{F0: 65535, F1: (v3 - v3)}
        _ = v11
        _ = v12
    }
    _ = len(v2)
    v13 := GenS2{F0: 1, F1: v3}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch -1 {
    case 1:
        var v14 uint64 = 1
        _ = v14
    default:
        switch -712 {
        case 897:
            var v15 uint = (v1 ^ v1)
            _ = v15
        case 4294967295:
            switch 2147483647 {
            case 551:
                if true {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        v17 := GenS3{F0: 65535}
                        _ = fi16
                        _ = v17
                    }
                    v18 := GenS4{F0: "foo bar", F1: 1, F2: 1}
                    v3 -= (v3 - v3)
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        var v20 int32 = 42
                        var v21 int16 = 1
                        var v22 uint32 = 42
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    fmt.Println(v1, v3)
                    for fi23 := 0; fi23 < 7; fi23++ {
                        for fi24 := 0; fi24 < 14; fi24++ {
                            switch "" {
                            case "foo bar":
                                v25 := GenS4{F0: ("a" + "a"), F1: -1, F2: 1}
                                _ = v25
                            case "\n\t":
                                fmt.Println(v1)
                            case "":
                                v1 -= (v1 | v1)
                            default:
                                v26 := GenS4{F0: "a", F1: int8(v3), F2: 42}
                                _ = v26
                            }
                            var v27 uint64 = (v3 >> ((181) & 63))
                            _ = fi24
                            _ = v27
                        }
                        v1 -= 121
                        _ = fi23
                    }
                }
            case 992:
                switch "" {
                case "foo bar":
                    v1 = v1
                case "a":
                    if (!((v1 % ((319) | 1)) < (^(v1)))) {
                        v1 = v1
                        v3 = 42
                        for fi28 := 0; fi28 < 5; fi28++ {
                            var v29 string = "a"
                            v30 := GenS5{F0: 685, F1: 128}
                            for fi31 := 0; fi31 < 20; fi31++ {
                                var v32 uint16 = 24
                                switch "a" {
                                case "a":
                                    v3 = (v3 * (v3 / ((1) | 1)))
                                case "foo bar":
                                    v33 := GenS6{F0: fi31, F1: 98}
                                    _ = v33
                                case "\n\t":
                                    var v34 int8 = -1
                                    _ = v34
                                }
                                var v35 string = (v29 + ("\n\t" + "\n\t"))
                                _ = fi31
                                _ = v32
                                _ = v35
                            }
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    }
                default:
                    fmt.Println(v3, v1)
                }
            default:
                var v36 string = ""
                _ = v36
            }
        default:
            fmt.Println(v1)
        }
    }
    switch 127 {
    case 9223372036854775807:
        var v37 float64 = (0.0 + (((1e10 * (float64((0.1 / 0.1)) * 0.0)) + (1e10 / (0.0 - ((-2.25 / 0.1) - (0.1 + 1.5))))) * ((1.5 - 0.0) / (1.5 - 0.0))))
        _ = v37
    default:
        v1 -= v1
    }
    if (839 < 32767) {
        var v38 uint16 = 2
        for fi39 := 0; fi39 < 7; fi39++ {
            if (!((1 <= -47) || (v3 > v3))) {
                v40 := GenS2{F0: 1, F1: v3}
                _ = v40
            }
            var v41 uint32 = 0
            _ = fi39
            _ = v41
        }
        _ = v38
    }
    var v42 uint16 = 255
    var v43 int16 = 259
    var v44 uint64 = (v3 >> ((v1) & 63))
    var v45 uint8 = 131
    var v46 float64 = 1.5
    v3 += 986
    for fi47 := 0; fi47 < 17; fi47++ {
        for fi48 := 0; fi48 < 15; fi48++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch ("\n\t" + "hello") {
            case "\n\t":
                switch "foo bar" {
                case "a":
                    var v49 int = fi48
                    _ = v49
                case "foo bar":
                    v50 := GenS6{F0: (fi47 >> ((65535) & 63)), F1: v43}
                    if (false || false) {
                        v51 := GenS6{F0: (fi47 % ((fi48) | 1)), F1: int16(v45)}
                        v43 = v43
                        _ = v51
                    } else {
                        fmt.Println(v44)
                        if (((("\n\t" + (("" + (((("a" + ("a" + "foo bar")) + ("\n\t" + (("hello" + "hello") + "foo bar"))) + ("\n\t" + "")) + "")) + "foo bar")) + (("" + "\n\t") + "a")) + "hello") != "a") {
                            if false {
                                v52 := v2[1]
                                _ = v52
                            } else {
                                v4 = append(v4, "foo bar")
                                if (42 <= 993) {
                                    if true {
                                        v53 := GenS7{F0: (^(v1))}
                                        switch (fi47 % ((fi47) | 1)) {
                                        case 128:
                                            switch -593 {
                                            case 2147483647:
                                                v42 += v42
                                                v54 := GenS8{F0: 127}
                                                _ = v54
                                            case 127:
                                                var v55 uint8 = (v45 / ((v45) | 1))
                                                if true {
                                                    var v56 int32 = int32(v42)
                                                    v43 = v43
                                                    v57 := GenS9{F0: 0, F1: true}
                                                    v58 := GenS4{F0: "hello", F1: 1, F2: v43}
                                                    var v59 string = "hello"
                                                    v60 := GenS10{F0: int(v1), F1: 0}
                                                    v61 := v6[128]
                                                    _ = v56
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                    _ = v60
                                                    _ = v61
                                                }
                                                var v62 int64 = -9223372036854775808
                                                _ = v55
                                                _ = v62
                                            case 32767:
                                                if (false || ((!(false && false)) || false)) {
                                                    if len(v5) > 0 {
                                                        v5[0] = 162
                                                    }
                                                    if (fi47 == (fi47 % ((fi47) | 1))) {
                                                        v1 = (v1 - 255)
                                                        if len(v5) > 0 {
                                                            v5[0] = 42
                                                        }
                                                        var v63 int64 = -1
                                                        _ = v63
                                                    }
                                                    var v64 uint8 = v45
                                                    _ = v64
                                                } else {
                                                    fmt.Println(v1, v1, fi48)
                                                    var v65 int16 = (^(v43))
                                                    v42 = (v42 * v42)
                                                    _ = v65
                                                }
                                            }
                                        case 127:
                                            var v66 float32 = 1e10
                                            _ = v66
                                        }
                                        _ = len(v4)
                                        _ = v53
                                    } else {
                                        v4 = append(v4, "foo bar")
                                        var v67 uint16 = (^(v42))
                                        _ = v67
                                    }
                                    v68 := GenS3{F0: v42}
                                    _ = v68
                                } else {
                                    fmt.Println(v42, v46, v43)
                                }
                                var v69 string = "a"
                                _ = v69
                            }
                        }
                        v3 = (v44 / ((945) | 1))
                    }
                    if false {
                        v70 := GenS6{F0: 1, F1: (v43 & (^(v43)))}
                        _ = v70
                    }
                    _ = v50
                default:
                    var v71 bool = false
                    _ = v71
                }
            case "hello":
                var v72 float64 = float64(-2.25)
                _ = v72
            default:
                if (!(int(v43) < fi48)) {
                    var v73 string = (("hello" + "a") + ("foo bar" + ("" + "foo bar")))
                    if false {
                        v46 += 1e10
                    }
                    v43 += 42
                    _ = v73
                }
            }
            var v74 int8 = 127
            _ = fi48
            _ = v74
        }
        fmt.Println(v45, fi47, v3)
        v43 -= v43
        _ = fi47
    }
    var v75 int = -881
    for fi76 := 0; fi76 < 16; fi76++ {
        v42 = v42
        switch 1 {
        case 32768:
            var v77 uint32 = 580
            _ = v77
        case -487:
            for fi78 := 0; fi78 < 8; fi78++ {
                delete(v6, (fi78 % ((fi76) | 1)))
                switch "\n\t" {
                case "hello":
                    v79 := GenS7{F0: v1}
                    _ = v79
                }
                _ = fi78
            }
        case -2147483648:
            if (4294967295 < 200) {
                for fi80 := 0; fi80 < 1; fi80++ {
                    v3 = 246
                    _ = fi80
                }
                for fi81 := 0; fi81 < 15; fi81++ {
                    v46 += v46
                    var v82 float64 = v46
                    v5 = append(v5, 2147483647)
                    _ = fi81
                    _ = v82
                }
                var v83 uint16 = (v42 % ((v42) | 1))
                _ = v83
            } else {
                var v84 uint = 2
                fmt.Println(v45)
                _ = v84
            }
        }
        _ = fi76
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v75:", v75)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v7
    _ = v75
}
