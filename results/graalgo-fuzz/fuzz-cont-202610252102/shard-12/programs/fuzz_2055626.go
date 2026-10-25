package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int16; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint; F1 bool; F2 string }
type GenS4 struct { F0 int; F1 uint8; F2 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 int16; F2 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32; F1 string }

func main() {
    v1 := []string{("\n\t" + ("foo bar" + "hello")), "foo bar", "\n\t", "foo bar"}
    var v2 uint16 = 32767
    var v3 float32 = (0.1 / (-2.25 - float32((0.1 + (1e10 - -2.25)))))
    var v4 float32 = v3
    var v5 float64 = 1e10
    fmt.Println(v4, v5, v2)
    fmt.Println(v3, v3)
    if (uint(v2) != 297) {
        switch ("foo bar" + "hello") {
        case "a":
            for fi6 := 0; fi6 < 5; fi6++ {
                for fi7 := 0; fi7 < 20; fi7++ {
                    if true {
                        if false {
                            v8 := GenS1{F0: 0, F1: 4294967295}
                            var v9 float32 = v4
                            if false {
                                v9 -= float32(v5)
                                if (42 > -28) {
                                    v3 = v9
                                    if ((v4 / float32((v5 / (v5 - v5)))) > v4) {
                                        v9 = (float32((v5 / (v5 * 0.1))) / v9)
                                    }
                                    if false {
                                        switch "a" {
                                        case "\n\t":
                                            var v10 uint8 = 42
                                            _ = v10
                                        case "hello":
                                            var v11 uint = 42
                                            _ = v11
                                        }
                                    } else {
                                        var v12 int8 = 1
                                        var v13 int64 = 42
                                        if false {
                                            v13 += v13
                                        } else {
                                            v14 := GenS2{F0: -32768, F1: 32767, F2: 4294967295}
                                            v15 := GenS1{F0: 70, F1: 830}
                                            _ = v14
                                            _ = v15
                                        }
                                        _ = v12
                                        _ = v13
                                    }
                                } else {
                                    if (!(-254 > fi6)) {
                                        var v16 int64 = -806
                                        v17 := GenS3{F0: 1, F1: (1e10 < float32(v5)), F2: (("\n\t" + "\n\t") + "a")}
                                        _ = v16
                                        _ = v17
                                    } else {
                                        var v18 uint = 21
                                        fmt.Println(v9)
                                        _ = v18
                                    }
                                    _ = len(v1)
                                    var v19 uint32 = 65535
                                    _ = v19
                                }
                                if len(v1) > 0 {
                                    v1[0] = ("" + "hello")
                                }
                            } else {
                                var v20 int16 = 42
                                v20 = 0
                                _ = v20
                            }
                            _ = len(v1)
                            _ = v8
                            _ = v9
                        }
                        if (722 < 776) {
                            v2 -= (v2 | 2)
                        } else {
                            v21 := GenS3{F0: 18446744073709551615, F1: ((!(!false)) || (!false)), F2: ("a" + "\n\t")}
                            switch fi6 {
                            case -213:
                                v5 += (0.0 * (0.0 / float64(1.5)))
                            case -1:
                                v22 := GenS4{F0: fi6, F1: 164, F2: 68}
                                _ = v22
                            }
                            var v23 int32 = -1
                            _ = v21
                            _ = v23
                        }
                        if (v3 == 0.0) {
                            if true {
                                v5 = -2.25
                            } else {
                                _ = len(v1)
                            }
                            switch ("" + "a") {
                            case "foo bar":
                                var v24 int8 = 42
                                _ = v24
                            case "a":
                                var v25 uint8 = 153
                                _ = v25
                            default:
                                if true {
                                    switch (fi7 / ((1) | 1)) {
                                    case 32768:
                                        if ((2147483647 == fi7) && false) {
                                            var v26 float64 = v5
                                            if false {
                                                v5 = (v26 * v26)
                                            } else {
                                                if len(v1) > 0 {
                                                    v1[0] = "a"
                                                }
                                                v27 := GenS2{F0: 1000, F1: int16(v2), F2: 18446744073709551615}
                                                v26 -= (1.5 - 0.0)
                                                _ = v27
                                            }
                                            var v28 string = "foo bar"
                                            _ = v26
                                            _ = v28
                                        }
                                    case 74:
                                        v1 = append(v1, "foo bar")
                                    default:
                                        var v29 int64 = 9223372036854775807
                                        _ = v29
                                    }
                                    v30 := GenS3{F0: 688, F1: ((fi6 - (^(fi7))) <= fi7), F2: "foo bar"}
                                    _ = v30
                                } else {
                                    if (v4 <= float32(v5)) {
                                        fmt.Println(fi7, v4, fi7)
                                        var v31 uint16 = 42
                                        v4 = v3
                                        var v32 uint = 1
                                        v3 = v4
                                        _ = v31
                                        _ = v32
                                    } else {
                                        switch (fi7 << ((18446744073709551615) & 63)) {
                                        case 128:
                                            v2 -= v2
                                        case -902:
                                            var v33 int64 = 1
                                            _ = v33
                                        case -9223372036854775808:
                                            if len(v1) > 0 {
                                                v1[0] = ""
                                            }
                                        default:
                                            if (!(true && true)) {
                                                v34 := GenS5{F0: "a"}
                                                v2 = (^(v2))
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v34
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v35 uint = 34
                                                var v36 int8 = 1
                                                _ = v35
                                                _ = v36
                                            }
                                        }
                                    }
                                }
                            }
                            var v37 int = 904
                            _ = v37
                        } else {
                            v38 := GenS1{F0: 255, F1: 4294967295}
                            _ = v38
                        }
                    }
                    v5 += v5
                    fmt.Println(v3, fi6)
                    _ = fi7
                }
                _ = fi6
            }
        case "\n\t":
            v4 = (1e10 * (v3 * (1e10 * 0.1)))
        case "":
            var v39 int16 = 0
            _ = v39
        }
        if false {
            v3 += ((v4 * v4) * float32(v5))
            _ = len(v1)
            var v40 uint = 244
            _ = v40
        }
        for fi41 := 0; fi41 < 4; fi41++ {
            v1 = append(v1, "foo bar")
            _ = len(v1)
            for fi42 := 0; fi42 < 12; fi42++ {
                _ = len(v1)
                var v43 int = (-(fi41))
                var v44 string = "a"
                var v45 int = (-(fi41))
                if (!true) {
                    switch (v44 + ((v44 + ("hello" + v44)) + (v44 + v44))) {
                    case "a":
                        v44 = ((v44 + v44) + "hello")
                    case "hello":
                        var v46 int32 = -2147483648
                        _ = v46
                    }
                    v47 := GenS6{F0: 2147483647, F1: -1, F2: 946}
                    _ = v47
                }
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
            _ = fi41
        }
    }
    v48 := GenS4{F0: 9223372036854775807, F1: 167, F2: 42}
    v5 += -2.25
    for fi49 := 0; fi49 < 16; fi49++ {
        v50 := GenS5{F0: "foo bar"}
        switch "foo bar" {
        case "":
            v51 := GenS1{F0: 0, F1: 2}
            _ = v51
        case "\n\t":
            for fi52 := 0; fi52 < 11; fi52++ {
                switch ((("foo bar" + (("" + "foo bar") + ("hello" + (("\n\t" + "hello") + ("" + ""))))) + ("a" + "\n\t")) + (((("\n\t" + "hello") + "") + ((((("foo bar" + "") + "") + (("a" + "a") + "\n\t")) + "\n\t") + (("a" + "a") + "foo bar"))) + "")) {
                case "foo bar":
                    var v53 float64 = (v5 / (1e10 * v5))
                    _ = v53
                case "":
                    v54 := GenS3{F0: 299, F1: (42 < -824), F2: ("" + ("hello" + "foo bar"))}
                    _ = v54
                default:
                    var v55 uint8 = 147
                    _ = v55
                }
                switch 1 {
                case -363:
                    if (2 <= 1) {
                        switch "\n\t" {
                        case "":
                            var v56 int = (fi52 << ((42) & 63))
                            _ = v56
                        case "foo bar":
                            var v57 uint8 = 42
                            _ = v57
                        case "\n\t":
                            if true {
                                var v58 bool = (false || true)
                                v58 = v58
                                var v59 uint8 = 1
                                _ = v58
                                _ = v59
                            } else {
                                if false {
                                    switch 2147483647 {
                                    case 32768:
                                        if false {
                                            v3 += (v4 - 1.5)
                                            v1 = append(v1, "hello")
                                            var v60 bool = false
                                            if len(v1) > 0 {
                                                v1[0] = ("" + "hello")
                                            }
                                            _ = v60
                                        }
                                        var v61 string = ""
                                        var v62 int = fi49
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(v2, fi49, fi49)
                                        _ = v61
                                        _ = v62
                                    case 451:
                                        fmt.Println(fi49, fi49, fi49)
                                    }
                                } else {
                                    if false {
                                        if true {
                                            if (v2 != v2) {
                                                fmt.Println(fi49, fi49)
                                            } else {
                                                v4 = float32(0.0)
                                                var v63 uint32 = 2147483647
                                                if (0 <= 648) {
                                                    v1 = append(v1, "a")
                                                    v64 := GenS7{F0: (fi52 - 32767)}
                                                    v1 = append(v1, "")
                                                    _ = v64
                                                } else {
                                                    fmt.Println(v4, fi49, v4)
                                                    v5 = -2.25
                                                }
                                                _ = v63
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            switch "hello" {
                                            case "a":
                                                var v65 int = fi49
                                                _ = v65
                                            default:
                                                var v66 uint16 = v2
                                                _ = v66
                                            }
                                        }
                                    }
                                    _ = len(v1)
                                }
                                if false {
                                    v67 := GenS8{F0: 0, F1: (("a" + "\n\t") + ("" + "\n\t"))}
                                    v1 = append(v1, ("" + "\n\t"))
                                    _ = v67
                                } else {
                                    v2 += (^(v2))
                                }
                            }
                        default:
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                        }
                        switch fi52 {
                        case 9223372036854775807:
                            var v68 bool = (!true)
                            _ = v68
                        }
                    } else {
                        v69 := GenS8{F0: 1, F1: "a"}
                        _ = len(v1)
                        fmt.Println(v4, v5)
                        _ = v69
                    }
                case -322:
                    fmt.Println(fi52, v2, v4)
                }
                _ = fi52
            }
        }
        _ = fi49
        _ = v50
    }
    switch 9223372036854775807 {
    case 127:
        v2 = 42
    }
    for fi70 := 0; fi70 < 17; fi70++ {
        for fi71 := 0; fi71 < 16; fi71++ {
            fmt.Println(v4)
            v2 -= (^(v2))
            _ = fi71
        }
        _ = fi70
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
}
