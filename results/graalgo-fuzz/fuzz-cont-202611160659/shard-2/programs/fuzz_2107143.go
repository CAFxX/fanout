package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 int64; F1 uint64 }
type GenS9 struct { F0 int32; F1 uint16 }
type GenS10 struct { F0 string; F1 uint32 }
type GenS11 struct { F0 bool; F1 int }
type GenS12 struct { F0 uint; F1 uint; F2 uint8 }
type GenS13 struct { F0 uint64; F1 uint64 }
type GenS14 struct { F0 int64; F1 uint16; F2 uint16 }
func genFn0() int64 {
    v4 := []int{-9223372036854775808}
    v5 := GenS1{F0: 2}
    v6 := GenS1{F0: 1}
    var v7 uint = 42
    var v8 string = "a"
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 792
    v2 := []string{""}
    var v3 uint16 = (^(v1))
    if (1 == 1) {
        var v9 uint32 = 545
        _ = v9
    } else {
        var v10 uint32 = 2147483647
        if len(v2) > 0 {
            v2[0] = ("\n\t" + (("\n\t" + (("a" + "\n\t") + ("hello" + "a"))) + ("foo bar" + "")))
        }
        v11 := GenS1{F0: uint32(v3)}
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 2147483648}
    v13 := GenS2{F0: 1}
    var v14 uint8 = 101
    switch 4294967295 {
    case 278:
        var v15 uint64 = 998
        _ = v15
    default:
        switch ("\n\t" + "") {
        case "hello":
            v1 = v3
        case "a":
            v14 -= v14
        default:
            if len(v2) > 0 {
                v2[0] = "hello"
            }
        }
    }
    v16 := GenS3{F0: 2}
    if true {
        if true {
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 int64 = 1
                _ = fi17
                _ = v18
            }
        } else {
            var v19 int64 = -9223372036854775808
            fmt.Println(v1)
            _ = v19
        }
    }
    switch "hello" {
    case "hello":
        if (false && true) {
            for fi20 := 0; fi20 < 7; fi20++ {
                for fi21 := 0; fi21 < 9; fi21++ {
                    _ = len(v2)
                    fmt.Println(fi20)
                    if (!false) {
                        if ((!(!true)) || (0.0 >= -2.25)) {
                            switch ("hello" + (("foo bar" + "hello") + ("hello" + ""))) {
                            case "a":
                                v22 := GenS4{F0: 1, F1: v14}
                                var v23 int64 = 9223372036854775807
                                var v24 int64 = 9223372036854775807
                                _ = v22
                                _ = v23
                                _ = v24
                            case "":
                                v25 := GenS3{F0: 393}
                                _ = v25
                            case "hello":
                                v14 = (^(v14))
                            }
                        } else {
                            v1 += 0
                            v3 = v3
                        }
                        var v26 int64 = 478
                        _ = v26
                    } else {
                        switch ("\n\t" + ("foo bar" + (("a" + "a") + ("" + (("hello" + ("" + "\n\t")) + "\n\t"))))) {
                        case "a":
                            var v27 uint8 = v14
                            _ = v27
                        case "":
                            if ((464 <= 427) || true) {
                                switch "hello" {
                                case "hello":
                                    if true {
                                        v3 += 0
                                        v28 := GenS5{F0: 42}
                                        _ = v28
                                    } else {
                                        if (911 >= (^(v1))) {
                                            fmt.Println(v3, fi21, v3)
                                            v29 := GenS6{F0: 262, F1: v1, F2: 2147483647}
                                            v30 := GenS4{F0: 18446744073709551615, F1: 120}
                                            _ = v29
                                            _ = v30
                                        }
                                        v31 := GenS7{F0: (!((0 == 531) && true)), F1: 70}
                                        _ = v31
                                    }
                                case "\n\t":
                                    fmt.Println(v1, fi21)
                                default:
                                    if true {
                                        var v32 uint64 = 2
                                        fmt.Println(fi21)
                                        _ = v32
                                    }
                                }
                                v3 += (v3 >> ((42) & 63))
                                v33 := GenS2{F0: 1000}
                                _ = v33
                            } else {
                                var v34 uint8 = 255
                                var v35 uint32 = 42
                                _ = v34
                                _ = v35
                            }
                            v36 := GenS6{F0: 15, F1: 2, F2: 477}
                            _ = v36
                        }
                        if true {
                            var v37 string = ("" + (("hello" + "\n\t") + ""))
                            switch fi21 {
                            case 8:
                                fmt.Println(v1, v3, fi21)
                            case 128:
                                v37 = (v37 + v37)
                            case 32768:
                                v1 = 32767
                            default:
                                var v38 uint32 = 65535
                                _ = v38
                            }
                            var v39 float64 = 0.1
                            _ = v37
                            _ = v39
                        }
                        var v40 uint64 = 415
                        _ = v40
                    }
                    _ = fi21
                }
                var v41 int16 = 32767
                v42 := GenS8{F0: 9223372036854775807, F1: 18446744073709551615}
                _ = fi20
                _ = v41
                _ = v42
            }
            for fi43 := 0; fi43 < 20; fi43++ {
                var v44 uint64 = 4294967295
                _ = fi43
                _ = v44
            }
            var v45 uint64 = 2
            _ = v45
        } else {
            for fi46 := 0; fi46 < 11; fi46++ {
                v14 -= v14
                v47 := GenS9{F0: 1, F1: 42}
                _ = fi46
                _ = v47
            }
            v3 -= uint16(v14)
        }
    case "":
        if true {
            v3 = v3
        } else {
            v1 -= v3
            fmt.Println(v14, v3, v1)
            v48 := GenS3{F0: 1}
            _ = v48
        }
    case "\n\t":
        if false {
            var v49 int32 = -2147483648
            var v50 uint = 4294967295
            _ = v49
            _ = v50
        } else {
            if false {
                switch -1 {
                case -1:
                    var v51 uint = 42
                    _ = v51
                case 128:
                    v14 += v14
                }
                v52 := GenS9{F0: 137, F1: uint16(v14)}
                for fi53 := 0; fi53 < 11; fi53++ {
                    if ((42 > 1) && (-274 <= 0)) {
                        var v54 uint64 = 0
                        var v55 uint64 = (v54 / ((42) | 1))
                        _ = v54
                        _ = v55
                    }
                    _ = fi53
                }
                _ = v52
            } else {
                _ = len(v2)
                fmt.Println(v14, v14)
                if false {
                    fmt.Println(v3, v1)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (!((((1e10 - 0.1) - 1e10) / (-2.25 / (3.14159 / 0.1))) <= 1e10)) {
                        var v56 uint64 = 0
                        var v57 int32 = 1
                        if (-149 < 1000) {
                            v58 := GenS5{F0: -1}
                            v2 = append(v2, "\n\t")
                            if ((2147483648 < 2147483648) || false) {
                                var v59 float32 = -2.25
                                if false {
                                    var v60 int32 = (v57 >> ((148) & 63))
                                    _ = v60
                                }
                                _ = v59
                            } else {
                                v61 := GenS10{F0: "foo bar", F1: 42}
                                _ = v61
                            }
                            _ = v58
                        } else {
                            if true {
                                if (0.0 != -2.25) {
                                    v56 += v56
                                    _ = len(v2)
                                    v3 = (v3 / ((v3) | 1))
                                }
                                for fi62 := 0; fi62 < 15; fi62++ {
                                    for fi63 := 0; fi63 < 11; fi63++ {
                                        v57 -= v57
                                        _ = fi63
                                    }
                                    var v64 uint32 = 796
                                    v65 := GenS11{F0: true, F1: fi62}
                                    _ = fi62
                                    _ = v64
                                    _ = v65
                                }
                            } else {
                                var v66 bool = false
                                v67 := GenS4{F0: v56, F1: (v14 % ((255) | 1))}
                                var v68 uint = 2
                                _ = v66
                                _ = v67
                                _ = v68
                            }
                            v57 = (v57 + (-(v57)))
                        }
                        _ = v56
                        _ = v57
                    } else {
                        for fi69 := 0; fi69 < 2; fi69++ {
                            v70 := GenS12{F0: uint(v3), F1: 400, F2: v14}
                            if len(v2) > 0 {
                                v2[0] = (("a" + "a") + "")
                            }
                            v14 += v14
                            _ = fi69
                            _ = v70
                        }
                        var v71 bool = false
                        _ = v71
                    }
                }
            }
        }
    }
    v3 = (v1 >> ((65535) & 63))
    switch "hello" {
    case "\n\t":
        v72 := GenS11{F0: true, F1: -9223372036854775808}
        _ = v72
    case "":
        if len(v2) > 0 {
            v2[0] = "a"
        }
    case "a":
        switch 32767 {
        case 9223372036854775807:
            var v73 int8 = 0
            _ = v73
        case 127:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case -299:
            if true {
                v74 := GenS10{F0: "hello", F1: 65535}
                for fi75 := 0; fi75 < 12; fi75++ {
                    for fi76 := 0; fi76 < 19; fi76++ {
                        var v77 string = ("" + "")
                        var v78 int32 = 442
                        _ = fi76
                        _ = v77
                        _ = v78
                    }
                    _ = fi75
                }
                var v79 uint = 434
                _ = v74
                _ = v79
            } else {
                var v80 float64 = float64(1e10)
                if false {
                    var v81 int = 127
                    _ = v81
                }
                _ = v80
            }
        default:
            for fi82 := 0; fi82 < 15; fi82++ {
                if (1e10 == 3.14159) {
                    v83 := GenS13{F0: 42, F1: 18446744073709551615}
                    var v84 float64 = -2.25
                    v85 := GenS14{F0: 42, F1: v1, F2: 32767}
                    _ = v83
                    _ = v84
                    _ = v85
                }
                _ = fi82
            }
        }
    }
    v3 += v3
    v2 = append(v2, "")
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
}
