package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 bool; F1 string; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 bool; F1 int32 }

func main() {
    var v1 int64 = -1
    v2 := make(map[string]bool)
    var v3 int = 345
    for fi4 := 0; fi4 < 12; fi4++ {
        v3 = (fi4 - (v3 >> ((4294967295) & 63)))
        switch fi4 {
        case -1:
            v3 -= 4294967295
        default:
            v3 = 323
        }
        for fi5 := 0; fi5 < 14; fi5++ {
            switch ("foo bar" + "foo bar") {
            case "foo bar":
                var v6 uint16 = 1
                _ = v6
            case "hello":
                v1 -= (v1 ^ -332)
            case "\n\t":
                if true {
                    v7 := GenS1{F0: 2, F1: ("hello" + ("hello" + ""))}
                    _ = v7
                } else {
                    fmt.Println(v3, fi5, fi5)
                    var v8 int64 = v1
                    var v9 uint64 = 4294967295
                    var v10 float32 = 0.1
                    if true {
                        v11 := []int{}
                        v2["foo bar"] = ((v8 >> ((372) & 63)) <= (v1 << ((230) & 63)))
                        _ = v11
                    }
                    _ = v10
                    _ = v8
                    _ = v9
                }
                var v12 uint64 = 2
                fmt.Println(v1, v3, v3)
                var v13 int64 = v1
                v14 := GenS2{F0: (249 != 1), F1: (v12 & 0)}
                v1 = (v1 | (v13 - (v13 - 9223372036854775807)))
                _ = v12
                _ = v13
                _ = v14
            default:
                if false {
                    v15 := GenS3{F0: true, F1: ("hello" + ("\n\t" + "")), F2: ("foo bar" + "foo bar")}
                    v1 = (v1 - 399)
                    v2["foo bar"] = (18446744073709551615 >= 42)
                    _ = v15
                } else {
                    var v16 uint16 = 2
                    fmt.Println(v3, v1)
                    _ = v16
                }
            }
            v3 += (v3 >> ((42) & 63))
            if ((2 == 678) && true) {
                delete(v2, "foo bar")
            }
            var v17 uint64 = 18446744073709551615
            _ = fi5
            _ = v17
        }
        _ = fi4
    }
    v1 -= (v1 - (v1 + (v1 - v1)))
    for fi18 := 0; fi18 < 18; fi18++ {
        v19 := GenS3{F0: (!((!((127 >= 42) && (!(("hello" == "\n\t") && true)))) && (-32768 > 1000))), F1: ("a" + (("foo bar" + ((("a" + ("\n\t" + ("" + "a"))) + "hello") + "")) + "hello")), F2: (("a" + "hello") + ("a" + ("foo bar" + "\n\t")))}
        var v20 int = fi18
        _ = fi18
        _ = v19
        _ = v20
    }
    switch (v3 | v3) {
    case 2147483647:
        for fi21 := 0; fi21 < 1; fi21++ {
            if true {
                var v22 bool = ((true && false) && false)
                v23 := GenS4{F0: 0}
                var v24 int8 = 1
                _ = v22
                _ = v23
                _ = v24
            } else {
                v25 := GenS2{F0: true, F1: 4294967295}
                delete(v2, ("hello" + "a"))
                v26 := []int{}
                _ = v25
                _ = v26
            }
            _ = fi21
        }
    case 42:
        var v27 string = ("" + ((("hello" + "\n\t") + "a") + "hello"))
        _ = v27
    default:
        for fi28 := 0; fi28 < 4; fi28++ {
            switch 1 {
            case 128:
                if (!((0 == 127) && (2 != 0))) {
                    delete(v2, "hello")
                    for fi29 := 0; fi29 < 14; fi29++ {
                        if true {
                            v30 := []int{fi28, 351}
                            if (155 == 2147483648) {
                                var v31 uint64 = 2
                                delete(v2, "")
                                switch (v3 & fi29) {
                                case 32768:
                                    v32 := GenS2{F0: true, F1: v31}
                                    _ = v32
                                case 128:
                                    if (!true) {
                                        if true {
                                            v1 = v1
                                            v1 = v1
                                            v33 := GenS2{F0: false, F1: v31}
                                            _ = v33
                                        } else {
                                            switch (fi28 + -2147483648) {
                                            case -709:
                                                _ = len(v30)
                                            }
                                            if len(v30) >= 2 {
                                                v30 = v30[1:len(v30)]
                                            }
                                            if false {
                                                var v34 float32 = 0.1
                                                switch ("" + ("\n\t" + "")) {
                                                case "foo bar":
                                                    v31 -= (v31 ^ 912)
                                                case "hello":
                                                    fmt.Println(v3, v34, v34)
                                                case "a":
                                                    fmt.Println(fi29)
                                                }
                                                _ = v34
                                            }
                                        }
                                        if (!false) {
                                            var v35 int64 = v1
                                            if (("foo bar" + "") != ("" + "a")) {
                                                v35 = (v35 & (v35 | (v1 << ((428) & 63))))
                                                v36 := GenS5{F0: 1, F1: 757, F2: 4294967295}
                                                switch 0 {
                                                case 32767:
                                                    var v37 int32 = 0
                                                    _ = v37
                                                case 2147483647:
                                                    if len(v30) >= 2 {
                                                        v30 = v30[1:len(v30)]
                                                    }
                                                }
                                                _ = v36
                                            }
                                            _ = v35
                                        }
                                    } else {
                                        v3 = fi29
                                        fmt.Println(fi28, v31)
                                    }
                                }
                                _ = v31
                            }
                            var v38 uint16 = 820
                            _ = v30
                            _ = v38
                        }
                        _ = fi29
                    }
                    if (-1 <= v3) {
                        var v39 bool = false
                        var v40 float64 = 0.1
                        _ = v39
                        _ = v40
                    } else {
                        fmt.Println(v3, fi28)
                        fmt.Println(fi28, fi28, v1)
                    }
                } else {
                    var v41 uint8 = 56
                    var v42 string = ("foo bar" + ("" + ("foo bar" + "a")))
                    v3 += (fi28 ^ int(v1))
                    _ = v41
                    _ = v42
                }
            case 4294967295:
                v43 := GenS1{F0: 797, F1: "a"}
                _ = v43
            case -22:
                switch ("hello" + "\n\t") {
                case "\n\t":
                    v44 := []int{v3, 128, (fi28 % ((int(v1)) | 1))}
                    _ = v44
                case "foo bar":
                    var v45 int64 = 0
                    _ = v45
                case "a":
                    var v46 float32 = (1.5 + -2.25)
                    _ = v46
                }
            default:
                v47 := GenS4{F0: 0}
                _ = v47
            }
            v2[(("\n\t" + ("foo bar" + "")) + ("foo bar" + ""))] = (32 < 42)
            v1 += v1
            _ = fi28
        }
    }
    v3 += -931
    v1 -= 323
    for fi48 := 0; fi48 < 1; fi48++ {
        fmt.Println(v1, v3)
        _ = fi48
    }
    fmt.Println(v3, v1)
    v1 = 0
    if (-773 != -2147483648) {
        switch v3 {
        case 42:
            for fi49 := 0; fi49 < 16; fi49++ {
                var v50 int32 = -1
                if (-32768 > 88) {
                    switch ("foo bar" + "a") {
                    case "hello":
                        var v51 uint8 = 1
                        _ = v51
                    case "":
                        var v52 uint64 = 863
                        _ = v52
                    }
                } else {
                    if (142 <= 231) {
                        if true {
                            switch (v3 >> ((677) & 63)) {
                            case 951:
                                v53 := []int{v3}
                                _ = v53
                            }
                            var v54 uint16 = 996
                            if (4294967295 == 1) {
                                var v55 int = (v3 | (^(v3)))
                                _ = v55
                            } else {
                                switch 2147483647 {
                                case 446:
                                    var v56 int16 = 42
                                    _ = v56
                                case 32767:
                                    v54 -= v54
                                case -257:
                                    v57 := GenS6{F0: 1, F1: 0}
                                    _ = v57
                                }
                                var v58 uint8 = 127
                                _ = v58
                            }
                            _ = v54
                        }
                        switch ("" + "\n\t") {
                        case "a":
                            for fi59 := 0; fi59 < 18; fi59++ {
                                v60 := v2["\n\t"]
                                _ = fi59
                                _ = v60
                            }
                        default:
                            if false {
                                if false {
                                    for fi61 := 0; fi61 < 7; fi61++ {
                                        if false {
                                            v62 := []int{fi61, fi49}
                                            _ = v62
                                        }
                                        _ = fi61
                                    }
                                }
                                var v63 uint = 0
                                var v64 float64 = 0.1
                                _ = v63
                                _ = v64
                            } else {
                                fmt.Println(v50)
                                var v65 string = "foo bar"
                                _ = v65
                            }
                        }
                        var v66 uint32 = 778
                        _ = v66
                    } else {
                        v67 := GenS4{F0: 32767}
                        _ = v67
                    }
                    v68 := GenS1{F0: 0, F1: "\n\t"}
                    var v69 int16 = -1
                    _ = v68
                    _ = v69
                }
                switch (fi49 / ((32768) | 1)) {
                case 9223372036854775807:
                    v3 = v3
                default:
                    fmt.Println(fi49)
                }
                _ = fi49
                _ = v50
            }
        case 1:
            var v70 uint = 278
            _ = v70
        case 128:
            for fi71 := 0; fi71 < 3; fi71++ {
                switch v3 {
                case -9223372036854775808:
                    var v72 int = (fi71 % (((fi71 - (v3 - 915))) | 1))
                    _ = v72
                case -369:
                    v73 := GenS7{F0: (false || false), F1: 754}
                    _ = v73
                default:
                    v3 += (v3 + v3)
                }
                var v74 float64 = (0.0 - (((0.1 - -2.25) * 1e10) / 1.0))
                _ = fi71
                _ = v74
            }
        default:
            fmt.Println(v3)
        }
        if (!false) {
            v3 = v3
        }
    } else {
        fmt.Println(v1)
        v75 := []int{v3, v3, 268, (v3 * 32767)}
        _ = v75
    }
    delete(v2, "")
    if (true || true) {
        var v76 float32 = float32(1.5)
        _ = v76
    } else {
        v77 := GenS1{F0: 622, F1: "foo bar"}
        fmt.Println(v3)
        _ = v77
    }
    var v78 int32 = -303
    fmt.Println(v1, v3)
    v78 = v78
    var v79 int64 = (v1 | v1)
    for fi80 := 0; fi80 < 16; fi80++ {
        fmt.Println(v3, v1)
        _ = fi80
    }
    fmt.Println(v79, v79, v79)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v78
    _ = v79
}
