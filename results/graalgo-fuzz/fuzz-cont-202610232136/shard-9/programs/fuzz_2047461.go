package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 int64; F1 int16; F2 int64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int8; F1 int32 }
func genFn0() bool {
    var v6 int32 = -2147483648
    if (4294967295 < 32768) {
        v6 -= (-(v6))
        switch 127 {
        case -1:
            if false {
                if (-1 == -97) {
                    fmt.Println(v6)
                }
            } else {
                var v7 uint8 = 42
                var v8 uint16 = 255
                for fi9 := 0; fi9 < 8; fi9++ {
                    v7 += v7
                    _ = fi9
                }
                _ = v7
                _ = v8
            }
        default:
            v6 += v6
        }
    } else {
        fmt.Println(v6)
    }
    _ = v6
    return (false && (549 != 255))
}

func main() {
    v1 := []string{"a"}
    var v2 int64 = -1
    var v3 int16 = 0
    var v4 uint64 = 157
    v5 := []string{"a", ("\n\t" + "a"), "a", (("\n\t" + "\n\t") + (("" + "") + ((("" + (("foo bar" + ("foo bar" + ("\n\t" + ""))) + "\n\t")) + "hello") + "foo bar")))}
    var v10 int32 = 849
    if ((v2 * 0) == v2) {
        var v11 float32 = 1.5
        var v12 string = ("hello" + ("\n\t" + "\n\t"))
        _ = v11
        _ = v12
    } else {
        if (true && false) {
            v4 -= v4
            v13 := GenS1{F0: 62, F1: 255, F2: 128}
            _ = v13
        }
    }
    if ((((v2 ^ (v2 | v2)) != -1) || (!true)) || (true || true)) {
        var v14 float32 = (float32(0.1) / 1.0)
        fmt.Println(v3, v4)
        v15 := GenS2{F0: 42, F1: v4}
        _ = v14
        _ = v15
    } else {
        var v16 float32 = -2.25
        _ = v16
    }
    v3 = v3
    var v17 uint64 = v4
    v4 -= 869
    if false {
        v1 = append(v1, "a")
    } else {
        if (0.1 != 1e10) {
            v17 = v4
            v17 = (v4 ^ v4)
            var v18 uint32 = 1
            _ = v18
        }
    }
    if true {
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
    } else {
        for fi19 := 0; fi19 < 20; fi19++ {
            fmt.Println(v3)
            _ = fi19
        }
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v2 += (v2 * (^(v2)))
            var v20 int64 = (v2 | v2)
            _ = v20
        }
        v2 = (^(v2))
    }
    var v21 uint16 = 694
    switch "hello" {
    case "a":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "":
        v22 := GenS2{F0: 572, F1: 4294967295}
        _ = v22
    case "hello":
        var v23 int64 = (v2 / (((-(v2))) | 1))
        _ = v23
    }
    var v24 int16 = v3
    if ((!true) && true) {
        v4 = v17
        switch ("foo bar" + "") {
        case "hello":
            var v25 uint32 = 65535
            _ = v25
        case "a":
            for fi26 := 0; fi26 < 1; fi26++ {
                fmt.Println(v17, v21, v24)
                fmt.Println(v2, v17)
                v1 = append(v1, ("foo bar" + "hello"))
                _ = fi26
            }
        }
        var v27 int64 = -1
        _ = v27
    }
    var v28 uint32 = 4294967295
    if false {
        v29 := GenS1{F0: v28, F1: (v21 / (((v21 ^ 42)) | 1)), F2: 128}
        fmt.Println(v2, v3)
        var v30 uint = 255
        _ = v29
        _ = v30
    } else {
        fmt.Println(v2, v21)
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v4 = 381
            _ = len(v1)
        } else {
            fmt.Println(v10, v17, v10)
        }
        if (uint(v28) >= 255) {
            if ((v24 != v24) && false) {
                for fi31 := 0; fi31 < 5; fi31++ {
                    var v32 int = fi31
                    _ = fi31
                    _ = v32
                }
                if len(v5) > 0 {
                    v5[0] = ("hello" + "")
                }
            }
            v33 := GenS3{F0: 18446744073709551615}
            _ = v33
        } else {
            var v34 uint = 1
            _ = v34
        }
    }
    if true {
        var v35 int8 = 127
        if (true && ((v2 << ((185) & 63)) > v2)) {
            for fi36 := 0; fi36 < 10; fi36++ {
                if true {
                    v37 := GenS2{F0: (v21 + 255), F1: (v4 * (v17 ^ v17))}
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        v38 := GenS2{F0: 255, F1: 1}
                        _ = v38
                    default:
                        if true {
                            switch "\n\t" {
                            case "":
                                if (("" != "foo bar") && false) {
                                    var v39 int16 = v24
                                    var v40 bool = false
                                    _ = v39
                                    _ = v40
                                }
                            default:
                                var v41 bool = ("" <= "hello")
                                _ = v41
                            }
                            if (v17 == v17) {
                                for fi42 := 0; fi42 < 2; fi42++ {
                                    if false {
                                        fmt.Println(v28, v24, v17)
                                        var v43 uint64 = 2
                                        _ = v43
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v10 -= 1
                                    }
                                    _ = fi42
                                }
                                v3 -= (v3 << ((uint(v21)) & 63))
                            } else {
                                if false {
                                    if false {
                                        var v44 int16 = (v24 / ((v24) | 1))
                                        switch 42 {
                                        case 1:
                                            fmt.Println(fi36, v4)
                                        case -2147483648:
                                            v35 -= (v35 + (v35 / ((32) | 1)))
                                        case 976:
                                            var v45 uint64 = (v17 << ((42) & 63))
                                            _ = v45
                                        default:
                                            var v46 int8 = (v35 ^ v35)
                                            _ = v46
                                        }
                                        _ = v44
                                    }
                                } else {
                                    var v47 float64 = 0.0
                                    _ = v47
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if (uint64(v28) != v17) {
                                    if len(v1) > 0 {
                                        v1[0] = ("foo bar" + ("\n\t" + "foo bar"))
                                    }
                                } else {
                                    v48 := GenS4{F0: 0, F1: "foo bar"}
                                    var v49 float32 = 1.5
                                    _ = v48
                                    _ = v49
                                }
                            }
                        }
                    }
                    switch 127 {
                    case -1:
                        fmt.Println(v28)
                    case -2147483648:
                        fmt.Println(v24, v28)
                    }
                    _ = v37
                } else {
                    v10 = 42
                    for fi50 := 0; fi50 < 18; fi50++ {
                        fmt.Println(v35, v4, v4)
                        fmt.Println(v28, v2, fi36)
                        v28 = 598
                        _ = fi50
                    }
                    var v51 int16 = 1000
                    _ = v51
                }
                _ = fi36
            }
            v2 = int64(v21)
        }
        for fi52 := 0; fi52 < 13; fi52++ {
            v53 := GenS2{F0: 1, F1: v17}
            if (387 > v2) {
                for fi54 := 0; fi54 < 5; fi54++ {
                    if false {
                        var v55 uint32 = 65535
                        switch (fi54 << ((42) & 63)) {
                        case -9223372036854775808:
                            v56 := GenS5{F0: false, F1: 414}
                            _ = v56
                        case -3:
                            fmt.Println(v28, v28, v24)
                        default:
                            var v57 int64 = (v2 * (v2 - v2))
                            _ = v57
                        }
                        _ = v55
                    }
                    if (true || false) {
                        fmt.Println(v3)
                        if true {
                            switch "a" {
                            case "a":
                                v35 = (v35 / ((v35) | 1))
                                var v58 float32 = ((float32(0.1) - 1.5) / 1.5)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v58
                            case "\n\t":
                                if false {
                                    var v59 float64 = 0.0
                                    v2 -= v2
                                    _ = v59
                                } else {
                                    if true {
                                        var v60 int8 = 127
                                        var v61 int = fi52
                                        _ = v60
                                        _ = v61
                                    }
                                }
                            default:
                                var v62 int = fi52
                                _ = v62
                            }
                            v24 = (v3 * 32767)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            v1 = append(v1, ("" + ""))
                            if (1e10 <= 0.1) {
                                v63 := GenS6{F0: (fi52 * fi54), F1: (!true)}
                                _ = v63
                            } else {
                                switch fi52 {
                                case 127:
                                    if len(v5) > 0 {
                                        v5[0] = ("foo bar" + ("hello" + ""))
                                    }
                                case -1:
                                    var v64 string = ("foo bar" + "")
                                    v65 := GenS1{F0: (v28 * 2147483647), F1: (^(v21)), F2: 255}
                                    _ = v64
                                    _ = v65
                                case 733:
                                    if len(v5) > 0 {
                                        v5[0] = "a"
                                    }
                                default:
                                    var v66 uint16 = v21
                                    _ = v66
                                }
                            }
                        }
                    }
                    _ = fi54
                }
                if (!true) {
                    _ = len(v5)
                }
            } else {
                if ((-(v35)) == (v35 - (v35 & v35))) {
                    switch ((("\n\t" + "a") + "hello") + "hello") {
                    case "foo bar":
                        v21 -= (v21 - 2)
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    v2 = int64(v17)
                    v67 := GenS7{F0: (v2 + v2), F1: (v3 & int16(v2)), F2: v2}
                    _ = v67
                }
                v5 = append(v5, "")
                for fi68 := 0; fi68 < 16; fi68++ {
                    v69 := GenS8{F0: int16(v4)}
                    v17 = 42
                    v17 = uint64(v28)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v70 := GenS9{F0: v35, F1: (v10 % (((v10 * v10)) | 1))}
                    _ = fi68
                    _ = v69
                    _ = v70
                }
            }
            v10 = v10
            _ = fi52
            _ = v53
        }
        _ = v35
    }
    v71 := GenS3{F0: 722}
    if (!((^(v28)) <= (v28 >> ((911) & 63)))) {
        var v72 float32 = (3.14159 - (float32(0.1) - 1e10))
        v17 = v4
        _ = v72
    }
    var v73 bool = ((false && (!true)) && true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v21
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v71
    _ = v73
}
