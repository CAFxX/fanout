package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS8 struct { F0 int16; F1 int64; F2 int }
type GenS9 struct { F0 uint32; F1 int16 }

func main() {
    var v1 int16 = 0
    var v2 int64 = 42
    var v3 bool = true
    var v4 int32 = -1
    var v5 string = "a"
    v6 := []int{623, -9223372036854775808, -213, 891}
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 0}
        _ = fi7
        _ = v8
    }
    if len(v6) > 0 {
        v6[0] = 42
    }
    if (999 < 2) {
        v2 += (v2 * 0)
        for fi9 := 0; fi9 < 20; fi9++ {
            if (((!(v4 == v4)) || ((!v3) && true)) || (v1 == int16(v4))) {
                fmt.Println(v1)
                var v10 int = (^(fi9))
                var v11 bool = (!false)
                _ = v10
                _ = v11
            } else {
                if len(v6) > 0 {
                    v6[0] = -921
                }
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 uint64 = 1
                    v14 := GenS2{F0: (-(v2)), F1: fi12}
                    if len(v6) > 0 {
                        v6[0] = fi12
                    }
                    if (!v3) {
                        v4 -= v4
                        var v15 float64 = 0.0
                        _ = v15
                    } else {
                        var v16 int16 = (^(v1))
                        var v17 int16 = v1
                        v18 := GenS3{F0: 1000}
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    v2 += v2
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                v5 = "hello"
            }
            _ = fi9
        }
    }
    v19 := GenS3{F0: v1}
    if (0 != 255) {
        if (false && true) {
            if len(v6) > 0 {
                v6[0] = -902
            }
            var v20 int64 = (v2 / ((v2) | 1))
            var v21 string = "hello"
            _ = v20
            _ = v21
        } else {
            var v22 uint8 = 195
            _ = v22
        }
    } else {
        v23 := GenS4{F0: (v4 >> ((0) & 63)), F1: 32767}
        _ = v23
    }
    v2 -= (v2 - v2)
    if v3 {
        for fi24 := 0; fi24 < 17; fi24++ {
            if true {
                var v25 float64 = (1.5 - (3.14159 - 0.0))
                for fi26 := 0; fi26 < 5; fi26++ {
                    v27 := GenS3{F0: (v1 & (v1 | (v1 % ((v1) | 1))))}
                    _ = fi26
                    _ = v27
                }
                var v28 int32 = (v4 << ((4294967295) & 63))
                _ = v25
                _ = v28
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                switch -2147483648 {
                case 2147483647:
                    if v3 {
                        fmt.Println(fi24)
                    }
                default:
                    switch (fi24 * fi24) {
                    case -9223372036854775808:
                        if ((3.14159 - 1.5) == 0.0) {
                            fmt.Println(v4, fi24)
                            var v29 int8 = 0
                            var v30 int8 = 1
                            _ = v29
                            _ = v30
                        } else {
                            for fi31 := 0; fi31 < 3; fi31++ {
                                _ = len(v6)
                                fmt.Println(v3, v2, v4)
                                _ = fi31
                            }
                        }
                    case -477:
                        var v32 int64 = v2
                        _ = v32
                    case 976:
                        if v3 {
                            for fi33 := 0; fi33 < 15; fi33++ {
                                v34 := GenS5{F0: 42}
                                fmt.Println(v3, fi24)
                                switch (^(fi24)) {
                                case -9223372036854775808:
                                    if false {
                                        v3 = (2147483647 == v4)
                                        if (0 >= 2147483648) {
                                            switch fi24 {
                                            case 570:
                                                var v35 float32 = 3.14159
                                                _ = v35
                                            }
                                            var v36 int64 = v2
                                            v1 = -773
                                            _ = v36
                                        } else {
                                            var v37 int64 = v2
                                            if len(v6) > 0 {
                                                v6[0] = (fi24 & fi24)
                                            }
                                            v6 = append(v6, -747)
                                            _ = v37
                                        }
                                    } else {
                                        var v38 uint = 2
                                        _ = v38
                                    }
                                }
                                _ = fi33
                                _ = v34
                            }
                            if true {
                                v4 = v4
                                for fi39 := 0; fi39 < 7; fi39++ {
                                    v2 += v2
                                    _ = fi39
                                }
                                for fi40 := 0; fi40 < 5; fi40++ {
                                    if (2 > 1) {
                                        var v41 int32 = 42
                                        _ = v41
                                    }
                                    _ = fi40
                                }
                            } else {
                                v42 := GenS3{F0: v1}
                                if ((128 > 127) && v3) {
                                    v5 = ""
                                    var v43 int = fi24
                                    fmt.Println(v4, v1)
                                    _ = v43
                                } else {
                                    if v3 {
                                        if v3 {
                                            switch v5 {
                                            case "a":
                                                v6 = append(v6, -812)
                                            }
                                            v44 := GenS6{F0: (v1 >> ((2) & 63)), F1: 2}
                                            if v3 {
                                                v4 -= v4
                                                if (!true) {
                                                    for fi45 := 0; fi45 < 4; fi45++ {
                                                        v6 = append(v6, fi45)
                                                        if v3 {
                                                            if (42 != 1) {
                                                                v46 := GenS7{F0: 4294967295, F1: 2, F2: 418}
                                                                _ = v46
                                                            } else {
                                                                if v3 {
                                                                    fmt.Println(v3, v2, v3)
                                                                    if v3 {
                                                                        if (!((^(v1)) > v1)) {
                                                                            switch v5 {
                                                                            case "\n\t":
                                                                                if false {
                                                                                    var v47 string = ("" + "")
                                                                                    var v48 int16 = 42
                                                                                    _ = v47
                                                                                    _ = v48
                                                                                }
                                                                            }
                                                                            v5 = ("\n\t" + (("\n\t" + v5) + (v5 + "foo bar")))
                                                                        }
                                                                        if false {
                                                                            v6 = append(v6, (fi24 * (fi45 - fi24)))
                                                                            v49 := GenS4{F0: (v4 / ((-2147483648) | 1)), F1: (fi45 & fi45)}
                                                                            fmt.Println(fi45, fi24)
                                                                            if (921 >= 594) {
                                                                                if len(v6) > 0 {
                                                                                    v6[0] = (fi24 | fi45)
                                                                                }
                                                                                v2 += -1
                                                                            } else {
                                                                                switch (fi24 | -326) {
                                                                                case 9223372036854775807:
                                                                                    v1 -= v1
                                                                                }
                                                                                if len(v6) > 0 {
                                                                                    v6[0] = 2147483647
                                                                                }
                                                                                if len(v6) >= 2 {
                                                                                    v6 = v6[1:len(v6)]
                                                                                }
                                                                            }
                                                                            _ = v49
                                                                        }
                                                                    }
                                                                    v50 := GenS8{F0: v1, F1: v2, F2: (fi24 >> ((16) & 63))}
                                                                    _ = v50
                                                                }
                                                            }
                                                        } else {
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            var v51 int16 = (v1 / (((v1 | (-(v1)))) | 1))
                                                            var v52 uint64 = 0
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                        v5 = ("a" + ("hello" + (v5 + "a")))
                                                        _ = fi45
                                                    }
                                                    v53 := GenS3{F0: -1}
                                                    fmt.Println(v5)
                                                    _ = v53
                                                } else {
                                                    if v3 {
                                                        v54 := GenS6{F0: v1, F1: 221}
                                                        if (!true) {
                                                            v1 += v1
                                                        }
                                                        v4 = v4
                                                        _ = v54
                                                    } else {
                                                        v1 -= v1
                                                    }
                                                    v5 = "a"
                                                    v4 = (v4 | v4)
                                                }
                                                switch 104 {
                                                case -480:
                                                    _ = len(v6)
                                                case -9223372036854775808:
                                                    v2 -= -728
                                                case 32768:
                                                    v55 := GenS9{F0: 386, F1: v1}
                                                    _ = v55
                                                default:
                                                    v5 = v5
                                                }
                                            }
                                            _ = v44
                                        }
                                        for fi56 := 0; fi56 < 14; fi56++ {
                                            var v57 uint = 65535
                                            _ = fi56
                                            _ = v57
                                        }
                                    }
                                    var v58 int32 = v4
                                    _ = v58
                                }
                                _ = v42
                            }
                            _ = len(v6)
                        } else {
                            var v59 string = (v5 + (("foo bar" + (((v5 + "") + (v5 + (v5 + (("a" + ("" + v5)) + "hello")))) + ("foo bar" + ""))) + v5))
                            var v60 uint = 18446744073709551615
                            fmt.Println(v4)
                            _ = v59
                            _ = v60
                        }
                    default:
                        var v61 bool = true
                        _ = v61
                    }
                }
            }
            _ = fi24
        }
        v2 = -9223372036854775808
        var v62 float64 = ((-2.25 / 1e10) - (0.0 + 1e10))
        _ = v62
    }
    fmt.Println(v3, v3)
    v6 = append(v6, -137)
    _ = len(v6)
    var v63 uint = 148
    if ((v2 - 538) >= -1) {
        fmt.Println(v3, v63, v3)
    }
    for fi64 := 0; fi64 < 12; fi64++ {
        v4 += v4
        _ = fi64
    }
    v4 += (^(v4))
    for fi65 := 0; fi65 < 19; fi65++ {
        var v66 float64 = (1e10 - 0.1)
        _ = fi65
        _ = v66
    }
    var v67 string = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v63:", v63)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v67
}
