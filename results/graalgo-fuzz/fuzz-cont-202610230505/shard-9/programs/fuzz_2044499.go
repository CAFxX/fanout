package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint16; F2 string }
type GenS4 struct { F0 string; F1 string; F2 int }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 int16; F1 int64 }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 int; F1 bool }

func main() {
    var v1 uint8 = 73
    var v2 int = 678
    var v3 bool = false
    var v4 int16 = 32767
    v5 := []int{1, v2, (v2 / ((v2) | 1))}
    v6 := GenS1{F0: 1, F1: v2, F2: -307}
    switch v2 {
    case -775:
        for fi7 := 0; fi7 < 15; fi7++ {
            if (v2 >= -686) {
                var v8 uint64 = 42
                _ = v8
            }
            for fi9 := 0; fi9 < 4; fi9++ {
                switch v2 {
                case 4294967295:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case 215:
                    if true {
                        var v10 uint64 = 2
                        v11 := GenS1{F0: 2, F1: 0, F2: (-(v4))}
                        _ = v10
                        _ = v11
                    }
                case -1:
                    _ = len(v5)
                default:
                    var v12 float64 = (1.5 + -2.25)
                    _ = v12
                }
                var v13 uint8 = 76
                v5 = append(v5, v2)
                _ = fi9
                _ = v13
            }
            _ = fi7
        }
    case 42:
        if (("\n\t" < "foo bar") && ("hello" < "foo bar")) {
            fmt.Println(v4, v2, v4)
            v14 := GenS2{F0: 2}
            _ = v14
        }
    default:
        var v15 string = "foo bar"
        _ = v15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        v3 = v3
        if (!false) {
            for fi17 := 0; fi17 < 4; fi17++ {
                var v18 float64 = (1.5 * 1.5)
                var v19 uint8 = 0
                v19 += v1
                _ = fi17
                _ = v18
                _ = v19
            }
            v4 = 215
            if (!true) {
                v3 = v3
            }
        } else {
            v1 += v1
        }
        fmt.Println(v3)
        _ = fi16
    }
    v2 -= v2
    _ = len(v5)
    if true {
        var v20 uint16 = 377
        var v21 int32 = 54
        var v22 uint16 = (^(v20))
        _ = v20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        switch 128 {
        case -2147483648:
            v24 := GenS3{F0: v2, F1: 32767, F2: "\n\t"}
            _ = v24
        case 127:
            v1 -= (v1 << ((18446744073709551615) & 63))
        case 32768:
            v3 = (!(("foo bar" <= "hello") || v3))
        default:
            for fi25 := 0; fi25 < 5; fi25++ {
                if v3 {
                    fmt.Println(v1, fi23)
                    fmt.Println(v4, fi25)
                    if (false || v3) {
                        fmt.Println(fi25, fi23)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if (1e10 == -2.25) {
                            if len(v5) > 0 {
                                v5[0] = -2147483648
                            }
                            var v26 uint32 = 513
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if v3 {
                                if true {
                                    v27 := GenS3{F0: (-(fi25)), F1: 42, F2: "hello"}
                                    if len(v5) > 0 {
                                        v5[0] = (v2 * (fi23 - fi25))
                                    }
                                    var v28 uint32 = (v26 - 213)
                                    _ = v27
                                    _ = v28
                                }
                                var v29 uint64 = 33
                                _ = v29
                            } else {
                                var v30 int32 = 107
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v30
                            }
                            var v31 float64 = 1e10
                            _ = v26
                            _ = v31
                        } else {
                            if (float32(0.1) >= -2.25) {
                                switch ("" + "a") {
                                case "foo bar":
                                    if (360 <= 65535) {
                                        if (0 > v4) {
                                            v4 -= (v4 & v4)
                                            var v32 int8 = 127
                                            if (1 >= 2) {
                                                if len(v5) > 0 {
                                                    v5[0] = -486
                                                }
                                                var v33 int8 = -16
                                                v34 := GenS2{F0: 2147483648}
                                                _ = v33
                                                _ = v34
                                            } else {
                                                fmt.Println(v1, v1, v2)
                                                var v35 uint8 = 2
                                                _ = v35
                                            }
                                            _ = v32
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = (^(fi23))
                                        }
                                        fmt.Println(v4, v2, fi25)
                                    } else {
                                        var v36 int32 = 1
                                        if false {
                                            var v37 int8 = -1
                                            _ = v37
                                        }
                                        if true {
                                            var v38 bool = (uint64(v1) <= 114)
                                            _ = v38
                                        } else {
                                            if true {
                                                var v39 int = fi25
                                                v1 -= v1
                                                var v40 float32 = 1.5
                                                _ = v39
                                                _ = v40
                                            } else {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                switch (fi23 & (fi25 & 0)) {
                                                case 42:
                                                    var v41 int64 = -9223372036854775808
                                                    _ = v41
                                                default:
                                                    fmt.Println(v1, fi23)
                                                }
                                                switch ("" + ("\n\t" + ("" + "hello"))) {
                                                case "a":
                                                    v42 := GenS1{F0: 255, F1: (-(fi23)), F2: (-(v4))}
                                                    _ = v42
                                                default:
                                                    v36 += -2147483648
                                                }
                                                fmt.Println(fi23, v3, v1)
                                            }
                                            if (255 > 255) {
                                                if (true && true) {
                                                    v4 = (v4 | (v4 * 902))
                                                    if v3 {
                                                        fmt.Println(v2, fi23)
                                                        fmt.Println(v1)
                                                        fmt.Println(v4, v2, v3)
                                                    } else {
                                                        var v43 float64 = float64((0.1 * 3.14159))
                                                        v44 := GenS4{F0: ("hello" + ("foo bar" + ("" + "a"))), F1: ("hello" + ""), F2: (fi25 % ((fi23) | 1))}
                                                        v36 = v36
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                    if ((v2 / (((fi23 + (fi25 + v2))) | 1)) == (fi23 << ((65535) & 63))) {
                                                        v45 := GenS1{F0: 279, F1: v2, F2: 946}
                                                        var v46 bool = (0 <= 943)
                                                        _ = v45
                                                        _ = v46
                                                    } else {
                                                        if false {
                                                            v36 -= 2147483647
                                                            v5 = append(v5, (v2 - (fi25 - v2)))
                                                            if v3 {
                                                                fmt.Println(v4, fi23)
                                                            } else {
                                                                var v47 uint8 = 0
                                                                v3 = false
                                                                if len(v5) > 0 {
                                                                    v5[0] = v2
                                                                }
                                                                _ = v47
                                                            }
                                                        }
                                                    }
                                                    v48 := GenS4{F0: "\n\t", F1: "a", F2: (fi25 ^ v2)}
                                                    if v3 {
                                                        var v49 float64 = -2.25
                                                        v49 = ((v49 + 1.5) * 1.5)
                                                        v3 = (201 == 572)
                                                        _ = v49
                                                    }
                                                    _ = v48
                                                }
                                            }
                                        }
                                        if v3 {
                                            var v50 int32 = v36
                                            _ = v50
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v36
                                    }
                                case "\n\t":
                                    _ = len(v5)
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if ("foo bar" > ("\n\t" + "\n\t")) {
                                    var v51 int = 32768
                                    _ = v51
                                } else {
                                    v52 := GenS5{F0: 922, F1: 951}
                                    _ = v52
                                }
                            }
                        }
                    } else {
                        v1 -= (^(v1))
                        switch "\n\t" {
                        case "foo bar":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        var v53 int8 = 1
                        _ = v53
                    }
                } else {
                    var v54 float32 = -2.25
                    if false {
                        v2 -= (^(fi23))
                    }
                    v55 := GenS5{F0: 285, F1: 1}
                    _ = v54
                    _ = v55
                }
                if v3 {
                    v4 -= (v4 % ((v4) | 1))
                    var v56 int16 = -403
                    _ = v56
                } else {
                    v1 = v1
                }
                if v3 {
                    fmt.Println(v4, fi23)
                    v1 = (v1 * v1)
                    v57 := GenS6{F0: 0, F1: ("foo bar" + "a")}
                    _ = v57
                } else {
                    if ((9223372036854775807 <= 9223372036854775807) && (true || (("" + "foo bar") >= (((("foo bar" + "") + "foo bar") + ("\n\t" + "foo bar")) + "hello")))) {
                        var v58 uint64 = 2
                        switch fi25 {
                        case 1:
                            v1 = (v1 % (((v1 << ((438) & 63))) | 1))
                        }
                        _ = v58
                    } else {
                        v59 := GenS4{F0: "hello", F1: "foo bar", F2: fi23}
                        v2 = v2
                        switch "a" {
                        case "a":
                            v2 = 9223372036854775807
                        }
                        _ = v59
                    }
                    fmt.Println(v1, fi25, v4)
                    switch 4294967295 {
                    case 128:
                        v4 = 1000
                    case 32767:
                        var v60 uint32 = 0
                        _ = v60
                    default:
                        v61 := GenS7{F0: (v4 ^ v4), F1: 42}
                        _ = v61
                    }
                    if ((^(v1)) > v1) {
                        v3 = (true || false)
                        v1 = (v1 ^ 60)
                        if (0 < 2147483648) {
                            v1 += (^(v1))
                            switch "" {
                            case "\n\t":
                                v4 -= v4
                                v1 = 219
                                var v62 string = ""
                                _ = v62
                            case "a":
                                var v63 int8 = 1
                                _ = v63
                            case "hello":
                                v1 = (v1 | (v1 & v1))
                            }
                            switch "hello" {
                            case "":
                                var v64 int64 = 0
                                _ = v64
                            case "a":
                                v2 = (-(fi25))
                            }
                        }
                    } else {
                        v65 := GenS8{F0: 1, F1: (v4 << ((1) & 63))}
                        _ = v65
                    }
                }
                _ = fi25
            }
        }
        if (1e10 < 0.1) {
            var v66 float32 = 1e10
            v1 -= (v1 + v1)
            _ = v66
        }
        _ = fi23
    }
    if true {
        var v67 uint32 = 2147483647
        _ = v67
    }
    for fi68 := 0; fi68 < 16; fi68++ {
        switch "" {
        case "a":
            for fi69 := 0; fi69 < 10; fi69++ {
                v1 += v1
                _ = fi69
            }
        }
        _ = fi68
    }
    v2 = v2
    if len(v5) > 0 {
        v5[0] = v2
    }
    for fi70 := 0; fi70 < 6; fi70++ {
        for fi71 := 0; fi71 < 17; fi71++ {
            var v72 uint64 = 862
            v73 := GenS9{F0: 18446744073709551615, F1: fi71}
            _ = len(v5)
            _ = fi71
            _ = v72
            _ = v73
        }
        if (!v3) {
            fmt.Println(v3)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v74 := GenS10{F0: (^(v2)), F1: v3}
            _ = v74
        }
        var v75 uint = 0
        _ = fi70
        _ = v75
    }
    var v76 uint16 = 320
    fmt.Println(v2)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
}
