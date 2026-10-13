package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint32; F1 string; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 bool; F2 bool }
type GenS5 struct { F0 int64; F1 int; F2 int16 }
func genFn0() uint {
    var v5 int16 = 32767
    var v6 uint = 978
    if false {
        if true {
            v7 := GenS1{F0: -2147483648, F1: 212}
            var v8 int32 = -937
            var v9 float32 = float32(1.5)
            _ = v7
            _ = v8
            _ = v9
        }
        v10 := GenS2{F0: 2, F1: "\n\t", F2: 1}
        for fi11 := 0; fi11 < 19; fi11++ {
            v12 := []int{fi11, (-(fi11))}
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        for fi13 := 0; fi13 < 16; fi13++ {
            for fi14 := 0; fi14 < 5; fi14++ {
                if true {
                    v15 := []int{fi14, (fi14 | 2147483647), 408, -529}
                    var v16 float64 = (3.14159 + 1e10)
                    if (false || (!(v6 <= v6))) {
                        v17 := GenS1{F0: 1, F1: 1}
                        var v18 uint16 = 706
                        var v19 uint = v6
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = v15
                    _ = v16
                } else {
                    if ((-9223372036854775808 < 42) && (25 > 69)) {
                        switch (fi13 ^ fi14) {
                        case 128:
                            var v20 float64 = 1.5
                            _ = v20
                        }
                        fmt.Println(v5, fi13)
                        if (!false) {
                            v21 := GenS2{F0: 4294967295, F1: ("a" + ""), F2: 1}
                            v22 := []int{(^(fi13)), (fi13 % ((2147483647) | 1)), (fi13 * 32768), 2147483647}
                            _ = v21
                            _ = v22
                        }
                    } else {
                        fmt.Println(v5, v5)
                    }
                    switch fi13 {
                    case 1:
                        v6 += 2
                    case -1:
                        v6 = 18446744073709551615
                    case 42:
                        var v23 uint32 = 42
                        _ = v23
                    default:
                        v24 := GenS3{F0: 0}
                        _ = v24
                    }
                    var v25 bool = false
                    _ = v25
                }
                fmt.Println(fi13, fi13, fi13)
                if true {
                    v26 := []int{(-(fi14))}
                    _ = v26
                } else {
                    v5 += (v5 + 32767)
                    v27 := GenS2{F0: 98, F1: "foo bar", F2: 0}
                    _ = v27
                }
                _ = fi14
            }
            var v28 int = fi13
            _ = fi13
            _ = v28
        }
        var v29 string = ""
        switch 127 {
        case -633:
            var v30 int32 = 1
            _ = v30
        case 32767:
            var v31 float32 = 0.1
            _ = v31
        }
        _ = v29
    }
    v6 -= v6
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 int64, p1 int, p2 int8) uint64 {
    v32 := []int{(p1 & (p1 * 919)), p1, p1, p1}
    var v33 bool = ((p0 & (^(p0))) < p0)
    if v33 {
        for fi34 := 0; fi34 < 2; fi34++ {
            var v35 int16 = 32767
            _ = fi34
            _ = v35
        }
        for fi36 := 0; fi36 < 5; fi36++ {
            switch "hello" {
            case "hello":
                p2 = p2
            case "a":
                switch fi36 {
                case -2147483648:
                    if v33 {
                        var v37 float32 = (float32(0.1) - float32(0.1))
                        fmt.Println(v33, v37)
                        _ = v37
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            p1 = fi36
                        default:
                            v38 := GenS1{F0: 1, F1: 1}
                            _ = v38
                        }
                        fmt.Println(fi36, p0)
                        var v39 bool = (805 < 18446744073709551615)
                        _ = v39
                    }
                }
            default:
                for fi40 := 0; fi40 < 6; fi40++ {
                    switch (^(fi36)) {
                    case 513:
                        var v41 uint16 = 0
                        _ = v41
                    case 983:
                        fmt.Println(p1, v33, p0)
                    case -1:
                        fmt.Println(p1, p1, p0)
                    }
                    _ = fi40
                }
            }
            _ = fi36
        }
    } else {
        for fi42 := 0; fi42 < 8; fi42++ {
            v43 := GenS3{F0: p0}
            _ = fi42
            _ = v43
        }
        for fi44 := 0; fi44 < 20; fi44++ {
            v45 := GenS4{F0: p0, F1: v33, F2: v33}
            switch "\n\t" {
            case "\n\t":
                v46 := GenS1{F0: 1, F1: 128}
                _ = v46
            }
            if len(v32) > 0 {
                v32[0] = p1
            }
            _ = fi44
            _ = v45
        }
        var v47 uint64 = 342
        _ = v47
    }
    for fi48 := 0; fi48 < 11; fi48++ {
        if len(v32) >= 2 {
            v32 = v32[1:len(v32)]
        }
        _ = fi48
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    _ = v33
    return 2
}

func main() {
    var v1 uint64 = 535
    var v2 bool = true
    v3 := []int{9223372036854775807, -999, 1, 42}
    v4 := make(map[string]int)
    switch "\n\t" {
    case "\n\t":
        if false {
            v1 += v1
            switch "a" {
            case "hello":
                fmt.Println(v1, v2)
            case "foo bar":
                if (v2 || false) {
                    v3 = append(v3, 835)
                    for fi49 := 0; fi49 < 14; fi49++ {
                        v50 := GenS1{F0: 42, F1: 60}
                        v51 := GenS3{F0: -1}
                        v52 := GenS1{F0: 0, F1: 202}
                        _ = fi49
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                    v2 = false
                }
            default:
                delete(v4, "")
            }
            v1 += 2
        } else {
            if (!(!(3.14159 >= 3.14159))) {
                var v53 bool = (!(v2 || (((0.0 / float32(float64((1.5 + float32(0.1))))) * 1e10) > float32((((((((-2.25 + float64((1e10 / 3.14159))) + 3.14159) / (-2.25 - (0.1 * ((1.5 * (1.5 * (((0.1 * 1.5) * 1.5) / 1.0))) + ((1e10 / (0.1 - 1.5)) + -2.25))))) * ((((-2.25 - 0.1) / float64((1e10 - float32(float64(0.0))))) / 3.14159) / (1e10 + 0.0))) + (1e10 / 1.0)) + ((((3.14159 - (3.14159 / 1.0)) + 0.0) + -2.25) * float64(3.14159))) * -2.25)))))
                _ = v53
            }
            var v54 int16 = -1
            _ = v54
        }
    case "foo bar":
        v1 -= v1
    default:
        switch -9223372036854775808 {
        case -103:
            if (!(!true)) {
                v55 := GenS1{F0: -641, F1: 103}
                var v56 uint = 691
                _ = v55
                _ = v56
            }
        case 703:
            v2 = (true && v2)
        case -432:
            switch "" {
            case "foo bar":
                v57 := GenS5{F0: -9223372036854775808, F1: 127, F2: -69}
                _ = v57
            case "hello":
                var v58 uint16 = 0
                _ = v58
            }
        }
    }
    fmt.Println(v1)
    for fi59 := 0; fi59 < 17; fi59++ {
        fmt.Println(v1)
        v1 -= 18446744073709551615
        _ = fi59
    }
    v1 += v1
    var v60 bool = (-14 < -43)
    switch "hello" {
    case "":
        var v61 uint64 = v1
        _ = v61
    }
    var v62 bool = ((v1 << ((4294967295) & 63)) < genFn1(9223372036854775807, 127, -35))
    fmt.Println(v1)
    v63 := GenS1{F0: 1, F1: 0}
    if (!(42 < 2147483647)) {
        v64 := GenS2{F0: 2147483648, F1: "a", F2: (v1 | (v1 * 503))}
        _ = v64
    }
    delete(v4, "foo bar")
    switch "hello" {
    case "":
        if v60 {
            v4[""] = 797
        }
    case "hello":
        fmt.Println(v60, v62)
    }
    _ = len(v3)
    if true {
        for fi65 := 0; fi65 < 16; fi65++ {
            for fi66 := 0; fi66 < 3; fi66++ {
                var v67 int16 = int16(fi66)
                _ = fi66
                _ = v67
            }
            v2 = v62
            var v68 uint64 = v1
            _ = fi65
            _ = v68
        }
    }
    var v69 float64 = 1e10
    v62 = ("a" == "a")
    if false {
        switch 368 {
        case 9223372036854775807:
            var v70 int = -9223372036854775808
            _ = v70
        case -2147483648:
            for fi71 := 0; fi71 < 13; fi71++ {
                for fi72 := 0; fi72 < 18; fi72++ {
                    v73 := GenS2{F0: 900, F1: "a", F2: (v1 >> ((4294967295) & 63))}
                    _ = fi72
                    _ = v73
                }
                v1 -= (v1 % (((v1 | (v1 ^ v1))) | 1))
                _ = fi71
            }
        default:
            fmt.Println(v69, v1)
        }
        v69 += v69
        var v74 uint16 = 32767
        _ = v74
    }
    switch 1 {
    case 896:
        if v60 {
            var v75 int32 = 867
            v75 -= v75
            for fi76 := 0; fi76 < 3; fi76++ {
                for fi77 := 0; fi77 < 4; fi77++ {
                    v1 += 0
                    _ = fi77
                }
                _ = fi76
            }
            _ = v75
        } else {
            switch 897 {
            case -952:
                var v78 int8 = 127
                _ = v78
            case 4294967295:
                for fi79 := 0; fi79 < 3; fi79++ {
                    _ = len(v4)
                    if v60 {
                        var v80 bool = (v60 || (143 <= 4294967295))
                        if v60 {
                            var v81 int64 = 9223372036854775807
                            delete(v4, ("hello" + "\n\t"))
                            _ = v81
                        } else {
                            v82 := GenS3{F0: -757}
                            delete(v4, "")
                            _ = v82
                        }
                        v69 -= v69
                        _ = v80
                    } else {
                        v83 := GenS5{F0: 1, F1: 2147483647, F2: -32768}
                        _ = v83
                    }
                    _ = fi79
                }
            default:
                delete(v4, ("a" + "\n\t"))
            }
        }
    case 0:
        v69 -= 0.0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v60:", v60)
    fmt.Println("v62:", v62)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v62
    _ = v63
    _ = v69
}
