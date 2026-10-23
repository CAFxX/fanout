package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 string; F1 string }
type GenS9 struct { F0 int; F1 bool }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint16, p1 uint32, p2 uint) int64 {
    fmt.Println(p0, p0)
    var v7 int64 = -79
    var v8 float32 = ((float32((0.1 / 1.0)) * -2.25) + 1.5)
    v9 := []int{2147483647, int(p0), 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return v7
}
func genFn1(p0 int8) int8 {
    v10 := GenS1{F0: 255, F1: 463}
    var v11 int = -9223372036854775808
    for fi12 := 0; fi12 < 4; fi12++ {
        fmt.Println(p0, fi12)
        v13 := GenS1{F0: 1, F1: 2147483647}
        _ = fi12
        _ = v13
    }
    p0 -= -128
    var v14 int16 = -952
    _ = p0
    _ = v10
    _ = v11
    _ = v14
    return -1
}

func main() {
    var v1 uint16 = 42
    var v2 bool = true
    var v3 int16 = -1
    v4 := []string{"\n\t", "hello"}
    v5 := []uint32{}
    v6 := []int{-528, 387, 1, 9223372036854775807}
    fmt.Println(v1)
    v15 := GenS2{F0: (v3 | v3)}
    v16 := GenS2{F0: 1000}
    if (4 < 0) {
        switch 9223372036854775807 {
        case 32767:
            v17 := GenS3{F0: 1}
            _ = v17
        default:
            for fi18 := 0; fi18 < 2; fi18++ {
                fmt.Println(fi18, v1, v1)
                if ((v3 - v3) <= (-(v3))) {
                    if len(v5) > 0 {
                        v5[0] = 0
                    }
                    if ((v2 && v2) || v2) {
                        v19 := GenS4{F0: 2147483647, F1: (v2 && true)}
                        if (false && v2) {
                            if (440 < -2147483648) {
                                v6 = append(v6, 1)
                            } else {
                                v20 := GenS5{F0: 1}
                                v2 = false
                                _ = v20
                            }
                            v1 = 2
                        } else {
                            if (!(-301 != 734)) {
                                v2 = (false || false)
                                for fi21 := 0; fi21 < 5; fi21++ {
                                    v22 := GenS5{F0: -1}
                                    _ = fi21
                                    _ = v22
                                }
                            }
                        }
                        var v23 int32 = 956
                        _ = v19
                        _ = v23
                    } else {
                        v24 := GenS4{F0: 2147483648, F1: (!(!v2))}
                        if (!(0 >= 127)) {
                            var v25 string = ("foo bar" + "\n\t")
                            fmt.Println(fi18, v3, fi18)
                            var v26 uint16 = 32767
                            _ = v25
                            _ = v26
                        } else {
                            v1 = v1
                            v2 = (!(!v2))
                            v3 = (v3 | v3)
                        }
                        _ = v24
                    }
                    v27 := GenS6{F0: 4294967295}
                    _ = v27
                } else {
                    if (18446744073709551615 == 42) {
                        for fi28 := 0; fi28 < 13; fi28++ {
                            switch int(v1) {
                            case 32767:
                                v3 = 0
                                var v29 int = -1
                                switch (v29 & fi18) {
                                case 32767:
                                    v3 -= 75
                                case 0:
                                    switch (-(fi28)) {
                                    case -693:
                                        if len(v6) > 0 {
                                            v6[0] = (^(fi18))
                                        }
                                    case -455:
                                        if v2 {
                                            fmt.Println(v29)
                                            v30 := GenS7{F0: 18446744073709551615, F1: ""}
                                            _ = v30
                                        }
                                    default:
                                        fmt.Println(v1, fi28, fi28)
                                    }
                                case -9223372036854775808:
                                    v31 := GenS8{F0: "a", F1: ("foo bar" + ("\n\t" + "\n\t"))}
                                    _ = v31
                                default:
                                    var v32 int16 = (v3 | v3)
                                    _ = v32
                                }
                                _ = v29
                            case 698:
                                v2 = (0.1 == -2.25)
                                var v33 uint64 = 703
                                _ = v33
                            case 232:
                                if (v3 > -1) {
                                    if v2 {
                                        var v34 uint32 = 2
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        fmt.Println(v34, v34)
                                        _ = v34
                                    }
                                } else {
                                    v35 := GenS9{F0: 1, F1: (0 == 365)}
                                    fmt.Println(v3, v2, v1)
                                    _ = v35
                                }
                            default:
                                v36 := GenS4{F0: 2, F1: false}
                                _ = v36
                            }
                            v3 = (v3 - (v3 / ((1000) | 1)))
                            _ = fi28
                        }
                    }
                    switch (("" + "") + "a") {
                    case "a":
                        var v37 string = "\n\t"
                        _ = v37
                    default:
                        v38 := GenS2{F0: (v3 ^ 1000)}
                        _ = v38
                    }
                    var v39 string = "a"
                    _ = v39
                }
                for fi40 := 0; fi40 < 11; fi40++ {
                    if false {
                        switch int(v3) {
                        case 9223372036854775807:
                            if v2 {
                                v2 = (18446744073709551615 >= 395)
                                var v41 float32 = 0.1
                                _ = v41
                            }
                        case 32767:
                            if v2 {
                                v42 := GenS3{F0: 42}
                                _ = v42
                            }
                        case 0:
                            if (42 == -128) {
                                if ("" == ("hello" + "hello")) {
                                    switch "foo bar" {
                                    case "hello":
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    case "\n\t":
                                        var v43 uint64 = 0
                                        _ = v43
                                    case "a":
                                        v1 = (v1 + v1)
                                        v6 = append(v6, (fi18 % ((42) | 1)))
                                        var v44 float32 = float32(float64((1.5 - (1.5 + -2.25))))
                                        _ = v44
                                    default:
                                        v1 -= 642
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = 2147483647
                                    }
                                    var v45 bool = (0 >= 137)
                                    v6 = append(v6, fi18)
                                    _ = v45
                                } else {
                                    var v46 int = fi18
                                    _ = v46
                                }
                                v4 = append(v4, "a")
                                v47 := GenS1{F0: (^(v1)), F1: 2147483647}
                                _ = v47
                            } else {
                                switch -456 {
                                case 128:
                                    var v48 uint16 = v1
                                    _ = v48
                                default:
                                    var v49 float32 = 3.14159
                                    _ = v49
                                }
                                v50 := GenS9{F0: fi40, F1: v2}
                                _ = v50
                            }
                        default:
                            _ = len(v5)
                        }
                        var v51 int8 = 72
                        _ = v51
                    }
                    switch fi18 {
                    case 42:
                        v52 := GenS10{F0: 0}
                        _ = v52
                    case 2147483647:
                        v1 += (v1 + (^(v1)))
                    }
                    v2 = (1 < 1)
                    _ = fi40
                }
                _ = fi18
            }
        }
        fmt.Println(v1, v2)
        v2 = v2
    } else {
        var v53 uint16 = v1
        var v54 uint8 = 0
        _ = v53
        _ = v54
    }
    var v55 uint8 = 85
    if v2 {
        fmt.Println(v2, v2)
        if len(v5) > 0 {
            v5[0] = 2147483647
        }
    } else {
        for fi56 := 0; fi56 < 13; fi56++ {
            v55 = (^(v55))
            if len(v5) > 0 {
                v5[0] = 65535
            }
            _ = fi56
        }
        var v57 bool = (!(42 < int64(v3)))
        v4 = append(v4, "")
        _ = v57
    }
    if len(v6) > 0 {
        v6[0] = 9223372036854775807
    }
    if (127 != 32) {
        switch 9223372036854775807 {
        case -1:
            var v58 uint = 255
            _ = v58
        case -383:
            fmt.Println(v55, v3, v55)
        }
        if (25 == 65535) {
            _ = len(v5)
            if v2 {
                fmt.Println(v2, v1, v1)
                v2 = (0 >= -505)
                if len(v5) > 0 {
                    v5[0] = 760
                }
            } else {
                var v59 string = ""
                v2 = false
                _ = v59
            }
        } else {
            var v60 int16 = (v3 / ((v3) | 1))
            v5 = append(v5, 0)
            _ = v60
        }
    } else {
        switch ("hello" + "") {
        case "":
            var v61 uint64 = 0
            _ = v61
        case "foo bar":
            fmt.Println(v3, v2)
        case "\n\t":
            for fi62 := 0; fi62 < 1; fi62++ {
                if (false && true) {
                    for fi63 := 0; fi63 < 3; fi63++ {
                        var v64 uint64 = 0
                        fmt.Println(v1, v64, v55)
                        _ = fi63
                        _ = v64
                    }
                    if true {
                        v5 = append(v5, 2147483648)
                        if (18446744073709551615 >= 42) {
                            var v65 string = "hello"
                            _ = v65
                        } else {
                            if (v2 && (-9223372036854775808 <= 859)) {
                                _ = len(v5)
                            }
                            if ("foo bar" <= "foo bar") {
                                v66 := GenS6{F0: 731}
                                _ = v66
                            } else {
                                switch (fi62 * 42) {
                                case 2147483647:
                                    for fi67 := 0; fi67 < 20; fi67++ {
                                        fmt.Println(v1, fi67)
                                        var v68 uint64 = 42
                                        _ = fi67
                                        _ = v68
                                    }
                                default:
                                    for fi69 := 0; fi69 < 1; fi69++ {
                                        v70 := GenS11{F0: -9223372036854775808}
                                        _ = fi69
                                        _ = v70
                                    }
                                }
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        var v71 int64 = -751
                        _ = v71
                    }
                }
                var v72 uint8 = (^(v55))
                _ = fi62
                _ = v72
            }
        }
        v3 = v3
    }
    if v2 {
        var v73 int8 = -1
        if v2 {
            fmt.Println(v73, v3, v1)
        }
        fmt.Println(v73)
        _ = v73
    } else {
        if v2 {
            var v74 int64 = 0
            for fi75 := 0; fi75 < 19; fi75++ {
                var v76 uint = 65535
                for fi77 := 0; fi77 < 15; fi77++ {
                    var v78 int16 = (v3 % ((v3) | 1))
                    _ = fi77
                    _ = v78
                }
                fmt.Println(v74, v74, fi75)
                _ = fi75
                _ = v76
            }
            v6 = append(v6, 42)
            _ = v74
        } else {
            v79 := GenS12{F0: 1, F1: 582}
            _ = len(v5)
            _ = v79
        }
        _ = len(v6)
        v80 := GenS1{F0: (v1 >> ((767) & 63)), F1: 0}
        _ = v80
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
