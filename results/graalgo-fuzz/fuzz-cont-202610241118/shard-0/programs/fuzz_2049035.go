package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int8; F2 uint8 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int; F2 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 string; F1 uint16; F2 int }
type GenS9 struct { F0 int; F1 string; F2 uint16 }
type GenS10 struct { F0 int16; F1 int16; F2 uint16 }
type GenS11 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS12 struct { F0 int32; F1 int32; F2 uint64 }
type GenS13 struct { F0 int16; F1 int; F2 string }
func genFn0(p0 uint8, p1 bool) int16 {
    if p1 {
        p0 = (^(p0))
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 bool = (((0.1 >= 3.14159) || p1) || (1 > 42))
            v10 := GenS1{F0: 2147483647, F1: -128}
            var v11 uint16 = 65535
            var v12 float32 = 1.5
            v13 := GenS1{F0: 1, F1: 42}
            _ = fi8
            _ = v10
            _ = v11
            _ = v12
            _ = v13
            _ = v9
        }
        v14 := GenS1{F0: fi7, F1: 1}
        v15 := GenS1{F0: (fi7 + fi7), F1: -1}
        _ = fi7
        _ = v14
        _ = v15
    }
    switch ("a" + "hello") {
    case "\n\t":
        p1 = p1
    }
    _ = p0
    _ = p1
    return 372
}

func main() {
    v1 := []string{}
    var v2 int64 = 9223372036854775807
    var v3 bool = true
    v4 := make(map[int]uint32)
    v5 := make(map[string]uint32)
    var v6 int32 = -1
    _ = len(v1)
    var v16 uint64 = 18446744073709551615
    v6 -= v6
    v1 = append(v1, "foo bar")
    v6 = v6
    v2 = -1
    switch ("a" + (("foo bar" + "a") + "foo bar")) {
    case "a":
        fmt.Println(v16)
    }
    if (v3 || (0 > 1)) {
        var v17 float64 = -2.25
        _ = v17
    } else {
        _ = len(v4)
    }
    var v18 int32 = (v6 + (v6 * (v6 + 2147483647)))
    switch -793 {
    case 2147483647:
        var v19 uint32 = 1
        _ = v19
    case -2147483648:
        var v20 uint32 = 2
        _ = v20
    case 948:
        v21 := GenS1{F0: 1, F1: -41}
        _ = v21
    }
    switch 9223372036854775807 {
    case 4294967295:
        var v22 string = "hello"
        _ = v22
    case 32767:
        var v23 float64 = -2.25
        _ = v23
    case 505:
        v3 = (!v3)
    default:
        for fi24 := 0; fi24 < 4; fi24++ {
            v2 -= (^(v2))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v25 int16 = genFn0(1, (!true))
            _ = fi24
            _ = v25
        }
    }
    v18 = v6
    v18 = (v18 << ((2) & 63))
    if ("hello" <= ((("" + "a") + "foo bar") + "\n\t")) {
        delete(v5, ("" + (("" + "") + ("hello" + "a"))))
        var v26 int16 = genFn0(2, v3)
        _ = v26
    } else {
        var v27 float32 = float32(0.0)
        if (!false) {
            var v28 uint32 = uint32(v16)
            _ = v28
        }
        v29 := GenS1{F0: 42, F1: 1}
        _ = v27
        _ = v29
    }
    if v3 {
        var v30 bool = (!v3)
        v6 = (-(v6))
        for fi31 := 0; fi31 < 5; fi31++ {
            if true {
                fmt.Println(v30, fi31)
                if v3 {
                    _ = len(v1)
                    var v32 uint16 = 971
                    for fi33 := 0; fi33 < 13; fi33++ {
                        fmt.Println(v16, fi31)
                        fmt.Println(v2)
                        v34 := GenS2{F0: (v2 + (v2 * 0))}
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                }
            }
            _ = fi31
        }
        _ = v30
    }
    v18 += v6
    for fi35 := 0; fi35 < 20; fi35++ {
        if v3 {
            var v36 uint = 255
            _ = v36
        } else {
            switch 127 {
            case 9223372036854775807:
                v37 := GenS3{F0: -1, F1: 42, F2: 127}
                _ = v37
            default:
                if (!(42 > -16)) {
                    var v38 float64 = 3.14159
                    _ = v38
                } else {
                    v39 := GenS4{F0: v3, F1: 711}
                    var v40 uint16 = 2
                    var v41 int8 = 100
                    _ = v39
                    _ = v40
                    _ = v41
                }
            }
            v16 -= 0
        }
        switch "a" {
        case "":
            for fi42 := 0; fi42 < 19; fi42++ {
                var v43 float32 = 0.1
                v44 := GenS2{F0: (-(v2))}
                var v45 int32 = v18
                fmt.Println(v3, v43)
                switch 32768 {
                case 32768:
                    v46 := GenS5{F0: 0, F1: fi42, F2: 919}
                    _ = v46
                case 998:
                    var v47 uint64 = v16
                    _ = v47
                case 2147483647:
                    switch 4294967295 {
                    case 128:
                        if (71 <= -1) {
                            _ = len(v5)
                        }
                    case -417:
                        fmt.Println(fi35, v43)
                    case -2147483648:
                        fmt.Println(v6)
                    default:
                        v18 = (v45 * -528)
                    }
                }
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
        default:
            if (!(!v3)) {
                fmt.Println(v2)
            }
        }
        _ = fi35
    }
    if v3 {
        var v48 uint64 = v16
        switch "" {
        case "hello":
            v1 = append(v1, "")
        case "foo bar":
            for fi49 := 0; fi49 < 4; fi49++ {
                if v3 {
                    switch (^(fi49)) {
                    case 984:
                        var v50 int32 = (v18 + v6)
                        _ = v50
                    }
                    _ = len(v4)
                } else {
                    for fi51 := 0; fi51 < 6; fi51++ {
                        if true {
                            v2 -= (v2 & v2)
                        } else {
                            if ((v3 || v3) && (((!(false || (v3 || v3))) || v3) && (v3 || v3))) {
                                if (!("hello" == "hello")) {
                                    fmt.Println(v18, v6)
                                }
                                v52 := GenS4{F0: v3, F1: 42}
                                _ = v52
                            } else {
                                var v53 int = fi51
                                _ = v53
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v16 = v16
                            if (1 > 32767) {
                                v16 = v16
                                var v54 uint = 596
                                _ = v54
                            }
                            var v55 string = ("" + ((("hello" + "hello") + "foo bar") + "a"))
                            _ = v55
                        }
                        if (2147483648 > 65535) {
                            if v3 {
                                var v56 uint32 = 313
                                if (255 < 137) {
                                    switch fi49 {
                                    case 32768:
                                        if true {
                                            v4[fi49] = v56
                                            if (565 > 242) {
                                                if true {
                                                    switch int(v16) {
                                                    case 4294967295:
                                                        if (v3 || v3) {
                                                            var v57 float64 = 1.5
                                                            _ = v57
                                                        } else {
                                                            fmt.Println(fi49, v16, v6)
                                                            v58 := GenS6{F0: 42}
                                                            var v59 int16 = -32768
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                    default:
                                                        v18 = v18
                                                    }
                                                    v2 = v2
                                                    var v60 uint64 = 633
                                                    _ = v60
                                                } else {
                                                    v4[32767] = v56
                                                }
                                                var v61 string = "\n\t"
                                                _ = v61
                                            } else {
                                                var v62 uint64 = v48
                                                var v63 bool = (!(!(65535 != 1)))
                                                _ = v62
                                                _ = v63
                                            }
                                            var v64 float64 = 3.14159
                                            if v3 {
                                                _ = len(v1)
                                                var v65 uint = 18446744073709551615
                                                _ = v65
                                            } else {
                                                var v66 string = "\n\t"
                                                var v67 string = "hello"
                                                v48 = (v48 % (((v48 % ((1) | 1))) | 1))
                                                _ = v66
                                                _ = v67
                                            }
                                            _ = v64
                                        } else {
                                            if v3 {
                                                fmt.Println(v56, v3, v2)
                                                fmt.Println(v18, v48)
                                            } else {
                                                var v68 float32 = 3.14159
                                                _ = v68
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    case 2147483647:
                                        if v3 {
                                            v69 := GenS7{F0: 42, F1: "\n\t"}
                                            switch ("\n\t" + "") {
                                            case "":
                                                v70 := GenS4{F0: false, F1: v56}
                                                _ = v70
                                            case "foo bar":
                                                v71 := GenS8{F0: ("\n\t" + "a"), F1: 65535, F2: (fi49 + 347)}
                                                _ = v71
                                            case "hello":
                                                v72 := GenS9{F0: fi49, F1: ("" + "foo bar"), F2: 536}
                                                _ = v72
                                            }
                                            _ = v69
                                        }
                                    }
                                } else {
                                    var v73 uint16 = 85
                                    fmt.Println(v48)
                                    _ = v73
                                }
                                _ = v56
                            } else {
                                switch (^(fi51)) {
                                case 3:
                                    var v74 uint32 = uint32(v16)
                                    _ = v74
                                case 2147483647:
                                    v3 = v3
                                case -1:
                                    if v3 {
                                        var v75 bool = v3
                                        _ = v75
                                    } else {
                                        var v76 string = ((("" + ("a" + "\n\t")) + "a") + "foo bar")
                                        v6 -= v6
                                        v6 += (v18 ^ v6)
                                        v1 = append(v1, ("" + ("foo bar" + v76)))
                                        if false {
                                            v48 = v48
                                        } else {
                                            var v77 uint = 633
                                            _ = v77
                                        }
                                        _ = v76
                                    }
                                default:
                                    switch "a" {
                                    case "foo bar":
                                        var v78 int = -195
                                        _ = v78
                                    case "\n\t":
                                        var v79 uint64 = 0
                                        _ = v79
                                    case "hello":
                                        v80 := GenS10{F0: 42, F1: -761, F2: 32767}
                                        v81 := GenS11{F0: 123, F1: 4294967295, F2: 65535}
                                        fmt.Println(v16)
                                        var v82 int8 = -3
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    default:
                                        var v83 uint32 = 42
                                        _ = v83
                                    }
                                }
                                v84 := GenS12{F0: v6, F1: v6, F2: (v48 - 791)}
                                _ = v84
                            }
                            v16 = (^(v48))
                        }
                        _ = fi51
                    }
                    for fi85 := 0; fi85 < 6; fi85++ {
                        v16 += 1
                        var v86 int8 = -8
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v2)
                        case "\n\t":
                            v48 += (v16 + (^(v48)))
                        case "foo bar":
                            v16 = (v48 << ((255) & 63))
                            if (255 != 445) {
                                fmt.Println(v86, v6, v16)
                                var v87 int8 = (v86 / ((v86) | 1))
                                _ = v87
                            }
                            v48 = (^(v16))
                        }
                        _ = len(v1)
                        v48 += v48
                        _ = fi85
                        _ = v86
                    }
                    for fi88 := 0; fi88 < 18; fi88++ {
                        v2 = v2
                        v89 := GenS2{F0: (v2 % ((v2) | 1))}
                        if v3 {
                            _ = len(v1)
                            var v90 uint64 = (v48 ^ (v48 >> ((42) & 63)))
                            v2 = v2
                            v91 := GenS13{F0: 42, F1: (fi49 + (fi88 * (fi88 * fi88))), F2: "a"}
                            var v92 bool = true
                            var v93 uint64 = (v90 + v16)
                            var v94 uint64 = v90
                            fmt.Println(fi49)
                            var v95 int16 = 42
                            fmt.Println(v3, fi49)
                            _ = v90
                            _ = v91
                            _ = v92
                            _ = v93
                            _ = v94
                            _ = v95
                        }
                        var v96 int8 = -1
                        _ = fi88
                        _ = v89
                        _ = v96
                    }
                }
                v48 += 163
                _ = fi49
            }
        default:
            for fi97 := 0; fi97 < 17; fi97++ {
                for fi98 := 0; fi98 < 18; fi98++ {
                    v2 = v2
                    fmt.Println(v48, v16, v16)
                    v2 -= 0
                    _ = fi98
                }
                v2 += (v2 ^ (v2 * v2))
                _ = fi97
            }
        }
        for fi99 := 0; fi99 < 14; fi99++ {
            switch ("foo bar" + "\n\t") {
            case "hello":
                var v100 uint = 255
                _ = v100
            case "\n\t":
                var v101 uint = 1
                _ = v101
            }
            _ = fi99
        }
        _ = v48
    } else {
        var v102 float32 = 1.5
        _ = v102
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
