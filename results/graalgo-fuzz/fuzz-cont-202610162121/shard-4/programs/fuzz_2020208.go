package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint8; F2 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 uint64; F1 uint16; F2 string }
type GenS10 struct { F0 int64; F1 uint64 }
func genFn0() int8 {
    var v7 float32 = float32(((1.5 + 1e10) + -2.25))
    v8 := []int{}
    v9 := GenS1{F0: 32768, F1: (42 == 760)}
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 uint) string {
    var v10 int64 = -9223372036854775808
    p0 = p0
    _ = p0
    _ = v10
    return ("" + ("a" + "\n\t"))
}

func main() {
    v1 := []string{"\n\t", "", (("foo bar" + "") + "a"), "foo bar"}
    var v2 int8 = 101
    var v3 uint = 117
    var v4 int = 459
    v5 := make(map[string]uint32)
    v6 := make(map[int]uint32)
    var v11 uint8 = 2
    var v12 float64 = ((1e10 * (1e10 / 1e10)) - 0.0)
    v13 := GenS1{F0: 127, F1: (9223372036854775807 >= 942)}
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS2{F0: v2, F1: true}
        v2 += -1
        v12 = v12
        _ = fi14
        _ = v15
    }
    fmt.Println(v12, v11)
    if false {
        switch "" {
        case "hello":
            switch "\n\t" {
            case "hello":
                if true {
                    var v16 float64 = (v12 * -2.25)
                    switch "hello" {
                    case "\n\t":
                        var v17 int8 = (v2 << ((v3) & 63))
                        _ = v17
                    case "foo bar":
                        var v18 uint64 = 2
                        _ = v18
                    case "a":
                        v11 = v11
                    default:
                        if true {
                            for fi19 := 0; fi19 < 20; fi19++ {
                                var v20 int32 = 0
                                var v21 uint8 = (v11 & (v11 << ((4294967295) & 63)))
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            if false {
                                switch v4 {
                                case 4294967295:
                                    var v22 int64 = -1
                                    _ = v22
                                case 42:
                                    var v23 int16 = 1
                                    _ = v23
                                }
                                switch v4 {
                                case 32767:
                                    v24 := GenS3{F0: ("foo bar" + "foo bar")}
                                    _ = v24
                                }
                            }
                        } else {
                            switch "hello" {
                            case "foo bar":
                                v11 -= (^(v11))
                            case "hello":
                                if true {
                                    var v25 uint = v3
                                    fmt.Println(v12)
                                    _ = v25
                                } else {
                                    var v26 int8 = v2
                                    var v27 float64 = ((v16 + v16) * -2.25)
                                    v3 = 255
                                    _ = v26
                                    _ = v27
                                }
                            default:
                                fmt.Println(v11)
                            }
                            v11 = v11
                            for fi28 := 0; fi28 < 3; fi28++ {
                                for fi29 := 0; fi29 < 17; fi29++ {
                                    var v30 int32 = 73
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v11 = v11
                                    v2 += v2
                                    var v31 bool = (1000 == 1)
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                                _ = fi28
                            }
                        }
                    }
                    var v32 float32 = -2.25
                    _ = v16
                    _ = v32
                }
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "a":
                delete(v5, (("foo bar" + (("" + "foo bar") + "a")) + "\n\t"))
            default:
                if (!(((1e10 - 0.1) > 3.14159) && false)) {
                    if (32767 == 1000) {
                        v33 := GenS2{F0: (v2 / ((42) | 1)), F1: ((0.0 + -2.25) <= 0.0)}
                        var v34 string = ("a" + ("hello" + "a"))
                        _ = v33
                        _ = v34
                    }
                    v35 := GenS3{F0: ("hello" + "hello")}
                    switch 127 {
                    case 4294967295:
                        for fi36 := 0; fi36 < 20; fi36++ {
                            v37 := GenS2{F0: 42, F1: (18446744073709551615 <= 18446744073709551615)}
                            v5["foo bar"] = 2
                            _ = fi36
                            _ = v37
                        }
                    default:
                        switch ("a" + "hello") {
                        case "":
                            var v38 string = ("\n\t" + "hello")
                            _ = v38
                        case "a":
                            for fi39 := 0; fi39 < 4; fi39++ {
                                for fi40 := 0; fi40 < 18; fi40++ {
                                    if (734 == 1) {
                                        var v41 int32 = int32(v2)
                                        _ = v41
                                    }
                                    _ = fi40
                                }
                                v11 += (v11 & v11)
                                v42 := GenS4{F0: v11, F1: v11, F2: ((false || (!((v12 == -2.25) && (!(-766 >= 9223372036854775807))))) || true)}
                                _ = fi39
                                _ = v42
                            }
                        default:
                            fmt.Println(v12)
                        }
                    }
                    _ = v35
                } else {
                    for fi43 := 0; fi43 < 15; fi43++ {
                        var v44 uint8 = 2
                        v45 := GenS5{F0: 235}
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                }
            }
        }
        v3 = 4294967295
    }
    var v46 int32 = int32(v11)
    if (0.1 <= 3.14159) {
        for fi47 := 0; fi47 < 17; fi47++ {
            _ = len(v5)
            for fi48 := 0; fi48 < 9; fi48++ {
                if true {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v49 bool = true
                    fmt.Println(v3, v11)
                    _ = v49
                }
                v50 := GenS4{F0: v11, F1: (v11 | v11), F2: true}
                _ = fi48
                _ = v50
            }
            var v51 uint16 = 32767
            _ = fi47
            _ = v51
        }
        v52 := GenS6{F0: 4294967295, F1: genFn0()}
        _ = v52
    }
    v5["foo bar"] = 788
    v11 = v11
    switch ("hello" + ("foo bar" + "a")) {
    case "":
        _ = len(v6)
    case "foo bar":
        var v53 float32 = (float32(v12) - -2.25)
        _ = v53
    case "\n\t":
        var v54 int64 = 1
        _ = v54
    default:
        for fi55 := 0; fi55 < 13; fi55++ {
            v11 = 213
            var v56 float32 = 0.1
            if (true || true) {
                switch (("\n\t" + "\n\t") + "hello") {
                case "":
                    v57 := GenS3{F0: "\n\t"}
                    _ = v57
                case "foo bar":
                    v58 := GenS7{F0: 787, F1: true}
                    _ = v58
                default:
                    for fi59 := 0; fi59 < 11; fi59++ {
                        var v60 uint32 = 795
                        if ((((("" + "") + ("" + "a")) + "") + ("\n\t" + "foo bar")) <= "") {
                            v5[""] = (v60 >> ((18446744073709551615) & 63))
                            var v61 uint64 = uint64(v11)
                            if true {
                                v11 += v11
                                switch "" {
                                case "\n\t":
                                    v3 = 65535
                                case "a":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = len(v6)
                                if true {
                                    v4 = (fi59 ^ 1)
                                    v62 := v5[(("" + "hello") + "\n\t")]
                                    _ = v62
                                } else {
                                    var v63 uint64 = 1
                                    var v64 int64 = -625
                                    _ = v63
                                    _ = v64
                                }
                            }
                            _ = v61
                        } else {
                            if (!(1000 >= 32767)) {
                                var v65 int8 = -128
                                v11 = v11
                                fmt.Println(v56)
                                if true {
                                    v60 = (v60 ^ (^(v60)))
                                } else {
                                    var v66 uint = v3
                                    v3 += (v3 ^ (v66 ^ v3))
                                    _ = len(v6)
                                    _ = v66
                                }
                                if true {
                                    v67 := GenS8{F0: 1, F1: 2}
                                    _ = len(v1)
                                    var v68 string = "a"
                                    var v69 int8 = int8(v3)
                                    if ((9223372036854775807 != -9223372036854775808) || (v11 != 255)) {
                                        var v70 uint8 = (^(v11))
                                        var v71 uint8 = (v70 / (((^(v70))) | 1))
                                        fmt.Println(v65, v68)
                                        fmt.Println(v60)
                                        var v72 int = int(v11)
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    }
                                    _ = v67
                                    _ = v68
                                    _ = v69
                                }
                                _ = v65
                            } else {
                                var v73 uint = 65535
                                v73 = (v3 >> ((v3) & 63))
                                v74 := GenS2{F0: (v2 % ((int8(v11)) | 1)), F1: (!(v12 <= float64(v56)))}
                                _ = v73
                                _ = v74
                            }
                        }
                        v75 := GenS7{F0: v60, F1: (v46 != v46)}
                        _ = fi59
                        _ = v60
                        _ = v75
                    }
                }
                for fi76 := 0; fi76 < 20; fi76++ {
                    var v77 uint = (^(v3))
                    switch (fi55 & fi76) {
                    case 127:
                        fmt.Println(fi76, v56, v11)
                    case 286:
                        switch "\n\t" {
                        case "a":
                            if false {
                                v3 = 233
                                v1 = append(v1, "\n\t")
                            }
                        case "foo bar":
                            _ = len(v5)
                        case "":
                            if false {
                                switch int(v46) {
                                case 42:
                                    v5["foo bar"] = 2147483648
                                    var v78 uint64 = 2
                                    _ = v78
                                }
                                v56 -= -2.25
                            } else {
                                switch fi76 {
                                case -1:
                                    v46 = -1
                                case -784:
                                    v79 := GenS8{F0: 42, F1: 32767}
                                    _ = v79
                                }
                            }
                        default:
                            var v80 int8 = (v2 % ((v2) | 1))
                            _ = v80
                        }
                    }
                    _ = fi76
                    _ = v77
                }
            } else {
                v46 = int32(v11)
                v1 = append(v1, "foo bar")
                v81 := GenS1{F0: int(v46), F1: ("hello" != "\n\t")}
                _ = v81
            }
            _ = fi55
            _ = v56
        }
    }
    v82 := GenS8{F0: 42, F1: 1}
    v83 := GenS9{F0: 58, F1: 1, F2: "a"}
    v3 -= (v3 | v3)
    v84 := GenS5{F0: (v11 << ((v3) & 63))}
    for fi85 := 0; fi85 < 15; fi85++ {
        var v86 float64 = (v12 * float64((3.14159 / -2.25)))
        _ = fi85
        _ = v86
    }
    if (!false) {
        fmt.Println(v46, v46, v3)
    } else {
        for fi87 := 0; fi87 < 12; fi87++ {
            fmt.Println(v4, v46, v2)
            v2 = (^(v2))
            _ = fi87
        }
        v88 := GenS5{F0: (v11 ^ v11)}
        if true {
            if (255 <= 65535) {
                v3 -= (^(v3))
                v3 = v3
            } else {
                fmt.Println(v12, v46, v4)
                var v89 float32 = 1.5
                _ = len(v1)
                _ = v89
            }
            v90 := GenS2{F0: 101, F1: true}
            v91 := v5["a"]
            _ = v90
            _ = v91
        } else {
            for fi92 := 0; fi92 < 7; fi92++ {
                switch -203 {
                case -285:
                    var v93 int32 = v46
                    _ = v93
                }
                v46 += 2147483647
                _ = fi92
            }
            v94 := GenS10{F0: -644, F1: 53}
            _ = v94
        }
        _ = v88
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v82
    _ = v83
    _ = v84
}
