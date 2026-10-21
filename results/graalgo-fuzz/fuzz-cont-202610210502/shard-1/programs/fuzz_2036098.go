package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int; F1 int64; F2 int16 }
type GenS5 struct { F0 uint }
func genFn0(p0 int32, p1 int8) uint8 {
    for fi7 := 0; fi7 < 5; fi7++ {
        p0 = p0
        fmt.Println(fi7, p0, p0)
        _ = fi7
    }
    v8 := GenS1{F0: 2}
    for fi9 := 0; fi9 < 18; fi9++ {
        if false {
            if false {
                var v10 float32 = (((((1e10 - 3.14159) * 3.14159) - 3.14159) - 1e10) - (1e10 + (((-2.25 / 0.1) + (-2.25 * (1e10 + 0.1))) - 0.1)))
                var v11 uint32 = 0
                switch 884 {
                case -9223372036854775808:
                    p1 = (p1 >> ((1) & 63))
                case -807:
                    if true {
                        var v12 uint8 = 127
                        var v13 uint16 = 255
                        var v14 int32 = (p0 / ((p0) | 1))
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                case 128:
                    v11 -= v11
                }
                _ = v10
                _ = v11
            }
            var v15 int32 = (p0 | int32(fi9))
            _ = v15
        } else {
            for fi16 := 0; fi16 < 7; fi16++ {
                if true {
                    switch 1 {
                    case 156:
                        if (!((-2.25 + (1e10 - 1e10)) >= 0.0)) {
                            if true {
                                var v17 uint = 2
                                v18 := []int{fi16, (fi16 | fi16)}
                                _ = v17
                                _ = v18
                            } else {
                                var v19 uint16 = 1
                                switch ("hello" + "") {
                                case "\n\t":
                                    v20 := GenS2{F0: "foo bar", F1: (v19 >> ((4294967295) & 63))}
                                    var v21 bool = false
                                    v22 := GenS3{F0: (fi9 ^ fi9), F1: fi9}
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                case "hello":
                                    if true {
                                        fmt.Println(fi16, fi16, p0)
                                        var v23 uint = 669
                                        _ = v23
                                    }
                                default:
                                    p1 = (-(p1))
                                }
                                var v24 uint = 255
                                _ = v19
                                _ = v24
                            }
                            v25 := []int{fi16, (fi9 ^ (fi16 ^ fi16)), 42}
                            p0 = -1
                            _ = v25
                        }
                    case -1:
                        if (p1 != (p1 * (-(p1)))) {
                            p0 = 1
                            if ((!true) || (-2.25 >= 1e10)) {
                                v26 := []int{fi9, -2147483648, fi16}
                                switch "a" {
                                case "":
                                    v27 := GenS1{F0: 2}
                                    _ = v27
                                case "hello":
                                    switch "a" {
                                    case "":
                                        var v28 int32 = 42
                                        _ = v28
                                    case "\n\t":
                                        fmt.Println(fi16, fi9, fi9)
                                    }
                                }
                                _ = v26
                            }
                        } else {
                            p0 = p0
                            p1 = 100
                        }
                    default:
                        fmt.Println(p1)
                    }
                    p1 -= 127
                    p0 = (p0 / ((p0) | 1))
                }
                _ = fi16
            }
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v8
    return 42
}

func main() {
    var v1 int8 = 42
    var v2 int32 = -1
    var v3 int8 = (-(v1))
    v4 := []string{"hello", ("a" + "\n\t"), "foo bar"}
    var v5 int32 = (v2 & (v2 ^ (v2 << ((18446744073709551615) & 63))))
    var v6 bool = true
    v5 = (-(v2))
    if true {
        if (0 >= 65535) {
            v5 -= (v2 - (v2 * 233))
        } else {
            for fi29 := 0; fi29 < 8; fi29++ {
                switch fi29 {
                case -1:
                    fmt.Println(v2, v3)
                case 4294967295:
                    fmt.Println(v2)
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi29
            }
            v30 := GenS4{F0: 1, F1: 9223372036854775807, F2: -418}
            var v31 int = 32767
            _ = v30
            _ = v31
        }
        var v32 uint8 = 0
        fmt.Println(v3)
        _ = v32
    } else {
        for fi33 := 0; fi33 < 9; fi33++ {
            if v6 {
                if (!v6) {
                    if v6 {
                        v3 -= 100
                        v4 = append(v4, "a")
                        v34 := GenS4{F0: 4294967295, F1: -736, F2: 42}
                        _ = v34
                    }
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                }
            }
            var v35 int32 = (v5 | (-(v5)))
            v5 -= v5
            _ = fi33
            _ = v35
        }
        var v36 string = ("a" + ((("foo bar" + ("" + "a")) + (("\n\t" + "a") + "a")) + ("hello" + "foo bar")))
        _ = v36
    }
    v3 = (-(v1))
    for fi37 := 0; fi37 < 10; fi37++ {
        if ((-697 <= (fi37 >> ((847) & 63))) && (268 <= 320)) {
            v38 := GenS2{F0: "", F1: 972}
            _ = v38
        }
        v1 += int8(fi37)
        _ = fi37
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        v40 := GenS4{F0: (fi39 << ((896) & 63)), F1: 0, F2: 0}
        v6 = (1 > 0)
        _ = fi39
        _ = v40
    }
    var v41 int8 = (v1 << ((4294967295) & 63))
    v6 = v6
    v42 := GenS2{F0: ("" + "foo bar"), F1: 853}
    switch 128 {
    case 2147483647:
        if (float32((float64((1e10 + 1.5)) / 3.14159)) >= -2.25) {
            switch 127 {
            case -2147483648:
                v43 := GenS5{F0: 42}
                _ = v43
            default:
                v4 = append(v4, "foo bar")
            }
            v2 = (v2 * (v5 * v2))
            var v44 int = 4294967295
            _ = v44
        } else {
            v6 = true
        }
    case 32768:
        if v6 {
            for fi45 := 0; fi45 < 6; fi45++ {
                v6 = v6
                var v46 int64 = -1
                _ = fi45
                _ = v46
            }
            v47 := GenS1{F0: 65535}
            _ = v47
        }
    default:
        v48 := GenS3{F0: 1, F1: -610}
        _ = v48
    }
    switch 128 {
    case -526:
        for fi49 := 0; fi49 < 19; fi49++ {
            switch ("hello" + "\n\t") {
            case "\n\t":
                for fi50 := 0; fi50 < 6; fi50++ {
                    switch "a" {
                    case "":
                        v1 = 127
                    default:
                        var v51 float32 = 3.14159
                        _ = v51
                    }
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                    v3 += v41
                    _ = fi50
                }
            case "":
                if v6 {
                    for fi52 := 0; fi52 < 13; fi52++ {
                        var v53 int8 = (v41 | 100)
                        _ = fi52
                        _ = v53
                    }
                    fmt.Println(v5)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            default:
                var v54 bool = (v6 || (!false))
                _ = v54
            }
            _ = fi49
        }
    case 728:
        if v6 {
            for fi55 := 0; fi55 < 16; fi55++ {
                v5 -= v5
                _ = fi55
            }
        } else {
            v4 = append(v4, "")
            var v56 bool = (765 == 0)
            for fi57 := 0; fi57 < 18; fi57++ {
                for fi58 := 0; fi58 < 16; fi58++ {
                    v3 = 0
                    v56 = v6
                    v59 := GenS4{F0: (fi58 / ((4294967295) | 1)), F1: 0, F2: 32767}
                    _ = fi58
                    _ = v59
                }
                v56 = v6
                var v60 uint16 = 854
                _ = fi57
                _ = v60
            }
            _ = v56
        }
    case -209:
        v41 = (v41 - v41)
    default:
        var v61 uint8 = 255
        _ = v61
    }
    for fi62 := 0; fi62 < 9; fi62++ {
        var v63 uint8 = 2
        _ = fi62
        _ = v63
    }
    if len(v4) > 0 {
        v4[0] = ("foo bar" + ("\n\t" + "\n\t"))
    }
    v4 = append(v4, (("\n\t" + "\n\t") + "a"))
    if v6 {
        fmt.Println(v3, v41)
    } else {
        switch "a" {
        case "\n\t":
            var v64 int8 = (v3 >> ((81) & 63))
            _ = v64
        default:
            fmt.Println(v6)
        }
        v3 = (^(v3))
    }
    fmt.Println(v5, v5, v1)
    for fi65 := 0; fi65 < 15; fi65++ {
        var v66 float64 = 0.0
        v4 = append(v4, ("foo bar" + "foo bar"))
        for fi67 := 0; fi67 < 11; fi67++ {
            var v68 string = "hello"
            _ = fi67
            _ = v68
        }
        _ = fi65
        _ = v66
    }
    if ("" < "hello") {
        if (!false) {
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
        }
        fmt.Println(v6, v6, v1)
        v3 = (v3 * v3)
    } else {
        var v69 int16 = 32767
        if (true || false) {
            fmt.Println(v6, v6, v6)
        }
        if v6 {
            var v70 int32 = -1
            var v71 int16 = v69
            for fi72 := 0; fi72 < 5; fi72++ {
                for fi73 := 0; fi73 < 19; fi73++ {
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                    fmt.Println(v3)
                    var v74 int = fi73
                    _ = fi73
                    _ = v74
                }
                fmt.Println(v69, v6)
                _ = fi72
            }
            _ = v70
            _ = v71
        } else {
            fmt.Println(v1, v41, v41)
        }
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
