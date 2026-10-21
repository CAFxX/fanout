package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 uint32; F2 int32 }
type GenS6 struct { F0 int32; F1 int8; F2 uint }
type GenS7 struct { F0 bool; F1 uint16 }

func main() {
    var v1 bool = (4294967295 > 737)
    var v2 float32 = 0.1
    var v3 uint32 = 42
    var v4 bool = v1
    v5 := []string{"", "\n\t", "\n\t", "foo bar"}
    v6 := []string{}
    v7 := GenS1{F0: 0, F1: 2}
    switch "" {
    case "":
        for fi8 := 0; fi8 < 11; fi8++ {
            v3 -= (v3 ^ (^(v3)))
            _ = len(v5)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi8
        }
    case "\n\t":
        v2 = 1.5
    case "hello":
        var v9 bool = false
        _ = v9
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        for fi10 := 0; fi10 < 9; fi10++ {
            for fi11 := 0; fi11 < 9; fi11++ {
                if (v2 > 3.14159) {
                    var v12 int8 = 82
                    _ = v12
                } else {
                    if v4 {
                        switch (^(fi10)) {
                        case -1:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v5 = append(v5, ("\n\t" + ("a" + "hello")))
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        var v13 int64 = -1
                        var v14 uint8 = 0
                        if v4 {
                            v4 = true
                        } else {
                            var v15 uint = 4294967295
                            var v16 uint32 = v3
                            v2 = (v2 * 0.1)
                            _ = v15
                            _ = v16
                        }
                        var v17 uint = 539
                        _ = v13
                        _ = v14
                        _ = v17
                    } else {
                        fmt.Println(v2, v3)
                    }
                    var v18 bool = v4
                    _ = v18
                }
                fmt.Println(v3, v1)
                _ = fi11
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                v5 = append(v5, (("hello" + ("\n\t" + "a")) + ""))
                if (("" + ("\n\t" + ("foo bar" + "hello"))) >= (((("\n\t" + (("foo bar" + ("hello" + ("a" + ("foo bar" + "foo bar")))) + ("hello" + "foo bar"))) + "hello") + ("\n\t" + ("hello" + "hello"))) + "")) {
                    var v20 uint32 = 2147483648
                    v1 = v1
                    var v21 uint = 255
                    _ = v20
                    _ = v21
                } else {
                    fmt.Println(v3, v1, fi19)
                    if (2 >= 4294967295) {
                        v22 := GenS2{F0: (("a" + "") + "a")}
                        _ = v22
                    } else {
                        var v23 uint32 = 4294967295
                        _ = v23
                    }
                    var v24 uint64 = 0
                    _ = len(v6)
                    _ = v24
                }
                var v25 int32 = 0
                _ = fi19
                _ = v25
            }
            for fi26 := 0; fi26 < 13; fi26++ {
                var v27 int = 501
                if len(v6) > 0 {
                    v6[0] = "\n\t"
                }
                fmt.Println(fi26, v2, v27)
                v27 = v27
                _ = fi26
                _ = v27
            }
            _ = fi10
        }
        var v28 string = "hello"
        _ = v28
    } else {
        for fi29 := 0; fi29 < 20; fi29++ {
            var v30 uint32 = (v3 << ((65535) & 63))
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v1 = (42 > 923)
            _ = fi29
            _ = v30
        }
    }
    v31 := GenS3{F0: 44}
    if v4 {
        v32 := GenS1{F0: 606, F1: 1}
        fmt.Println(v1, v4)
        v33 := GenS1{F0: 0, F1: 2}
        _ = v32
        _ = v33
    }
    fmt.Println(v1)
    v34 := GenS3{F0: (v3 | (v3 * v3))}
    var v35 int16 = -311
    fmt.Println(v2, v3)
    for fi36 := 0; fi36 < 2; fi36++ {
        var v37 bool = (v1 && ((-1 < -1) || v4))
        _ = fi36
        _ = v37
    }
    v6 = append(v6, "\n\t")
    for fi38 := 0; fi38 < 16; fi38++ {
        for fi39 := 0; fi39 < 5; fi39++ {
            v2 -= v2
            switch (fi38 << ((732) & 63)) {
            case 882:
                var v40 uint16 = 255
                _ = v40
            case 412:
                switch (-(fi39)) {
                case 9223372036854775807:
                    var v41 uint16 = 360
                    _ = v41
                case 0:
                    v3 += v3
                default:
                    var v42 int8 = 100
                    _ = v42
                }
            default:
                if v4 {
                    v35 -= v35
                    v5 = append(v5, (("foo bar" + "a") + "hello"))
                } else {
                    v43 := GenS3{F0: v3}
                    switch "foo bar" {
                    case "a":
                        v5 = append(v5, "a")
                    default:
                        switch (("" + "\n\t") + "hello") {
                        case "hello":
                            fmt.Println(fi38, v35)
                        default:
                            v44 := GenS2{F0: ""}
                            _ = v44
                        }
                    }
                    _ = v43
                }
            }
            var v45 int64 = 0
            switch ("a" + "\n\t") {
            case "hello":
                v46 := GenS1{F0: -870, F1: 218}
                switch "foo bar" {
                case "hello":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case "":
                    if false {
                        var v47 float32 = v2
                        fmt.Println(v35, v35)
                        _ = v47
                    } else {
                        v45 += 9223372036854775807
                    }
                case "\n\t":
                    if (0 <= 0) {
                        var v48 int8 = 23
                        _ = v48
                    }
                default:
                    v49 := GenS4{F0: 128}
                    _ = v49
                }
                v3 -= 2147483647
                v50 := GenS3{F0: 42}
                _ = v46
                _ = v50
            case "foo bar":
                if (2 >= v3) {
                    switch (-(fi38)) {
                    case 1:
                        var v51 uint32 = 0
                        _ = v51
                    default:
                        if v1 {
                            v52 := GenS2{F0: "\n\t"}
                            _ = v52
                        } else {
                            fmt.Println(v45)
                        }
                    }
                    var v53 uint64 = 2
                    var v54 int64 = (v45 | 639)
                    switch "\n\t" {
                    case "a":
                        var v55 float32 = float32(((1e10 - 1.5) - float64(-2.25)))
                        _ = v55
                    case "foo bar":
                        if v1 {
                            fmt.Println(v54, v45)
                            v56 := GenS1{F0: -1, F1: 1}
                            var v57 int16 = v35
                            _ = v56
                            _ = v57
                        }
                    case "\n\t":
                        v2 += ((float32((-2.25 / 1.0)) - 0.0) * (v2 * v2))
                        v35 = v35
                        var v58 uint = 517
                        v59 := GenS5{F0: v58, F1: (v3 >> ((uint(v53)) & 63)), F2: 0}
                        _ = v58
                        _ = v59
                    default:
                        v35 = v35
                    }
                    v2 -= 0.0
                    _ = v53
                    _ = v54
                } else {
                    switch 1 {
                    case 389:
                        v2 += float32(3.14159)
                    case 9223372036854775807:
                        if (v4 && (-112 <= 54)) {
                            v2 = (v2 + -2.25)
                            if (2147483647 <= v3) {
                                var v60 int32 = 1
                                var v61 int64 = 1
                                v3 = (v3 >> ((65535) & 63))
                                _ = v60
                                _ = v61
                            }
                            if len(v5) > 0 {
                                v5[0] = "a"
                            }
                            v2 -= (v2 - 0.0)
                        }
                    case 32767:
                        v45 = (v45 | (v45 | v45))
                    default:
                        if len(v6) > 0 {
                            v6[0] = "\n\t"
                        }
                        var v62 int32 = -20
                        _ = v62
                    }
                    switch (fi39 + (fi38 - fi39)) {
                    case -2147483648:
                        v63 := GenS5{F0: uint(v3), F1: v3, F2: -2147483648}
                        _ = v63
                    case 0:
                        if v1 {
                            if ((!((v2 - 3.14159) > v2)) && (v4 && (!v4))) {
                                var v64 int32 = 0
                                v3 = (^(v3))
                                _ = v64
                            }
                            switch "a" {
                            case "\n\t":
                                switch "a" {
                                case "hello":
                                    v4 = false
                                case "\n\t":
                                    var v65 float64 = (1.5 * -2.25)
                                    _ = v65
                                }
                            default:
                                fmt.Println(fi39)
                            }
                            fmt.Println(fi39, fi38)
                        } else {
                            var v66 int8 = 42
                            var v67 float32 = ((3.14159 * v2) + (1.5 / (0.1 + ((v2 + (v2 * v2)) * v2))))
                            fmt.Println(fi38, v45)
                            var v68 uint16 = 255
                            var v69 int16 = -1
                            v68 -= v68
                            v35 += (v69 - v35)
                            _ = v66
                            _ = v67
                            _ = v68
                            _ = v69
                        }
                    case 128:
                        var v70 float32 = (0.1 / (1e10 - (((v2 * v2) + v2) + v2)))
                        _ = v70
                    }
                }
                if (!(1 >= 183)) {
                    switch fi39 {
                    case -1:
                        var v71 int32 = 0
                        _ = v71
                    case -2147483648:
                        if (0 <= 30) {
                            fmt.Println(fi39)
                            var v72 bool = v1
                            _ = v72
                        }
                    case 799:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    default:
                        var v73 int32 = 42
                        _ = v73
                    }
                } else {
                    if true {
                        var v74 float32 = (v2 - v2)
                        v75 := GenS2{F0: ""}
                        _ = v74
                        _ = v75
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v76 int8 = -128
                        _ = v76
                    }
                    v77 := GenS6{F0: -848, F1: -1, F2: 1}
                    v78 := GenS7{F0: (false && (int16(v3) >= (^(v35)))), F1: 42}
                    v2 = v2
                    _ = v77
                    _ = v78
                }
                v6 = append(v6, ("hello" + "a"))
                if true {
                    switch (fi38 | (fi39 - 341)) {
                    case 127:
                        fmt.Println(fi38, v45, v3)
                    case 0:
                        v3 = (v3 - (^(v3)))
                    case 781:
                        var v79 uint8 = 0
                        _ = v79
                    }
                }
                v3 = v3
            default:
                var v80 string = "a"
                _ = v80
            }
            _ = fi39
            _ = v45
        }
        v81 := GenS1{F0: 2147483647, F1: 164}
        _ = fi38
        _ = v81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
