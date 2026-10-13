package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 uint64 }
func genFn0() int8 {
    if true {
        var v5 int32 = 42
        if (!true) {
            var v6 uint64 = 2
            _ = v6
        } else {
            fmt.Println(v5)
            fmt.Println(v5)
            v5 = -366
        }
        _ = v5
    } else {
        var v7 float64 = (0.1 - (0.1 * float64(3.14159)))
        if (265 <= 32767) {
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 uint8 = 255
                for fi10 := 0; fi10 < 15; fi10++ {
                    switch -2147483648 {
                    case 42:
                        fmt.Println(v7, v9)
                        v11 := GenS1{F0: 32767}
                        _ = v11
                    case -365:
                        v9 += (^(v9))
                    case -1:
                        var v12 uint64 = 1
                        _ = v12
                    default:
                        var v13 float64 = v7
                        _ = v13
                    }
                    if (2147483647 <= 2) {
                        v7 = (v7 + float64(1.5))
                        var v14 int64 = 9223372036854775807
                        v14 += v14
                        _ = v14
                    } else {
                        v15 := GenS1{F0: -9}
                        var v16 uint64 = 42
                        v17 := GenS1{F0: int16(v9)}
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    switch ((("hello" + ("foo bar" + (("hello" + ("\n\t" + "\n\t")) + ("hello" + (("\n\t" + "") + ""))))) + ((("a" + ((("" + "a") + "\n\t") + (("" + "a") + "\n\t"))) + "") + "a")) + "hello") {
                    case "a":
                        var v18 int = (fi10 & (fi10 - fi8))
                        _ = v18
                    case "foo bar":
                        var v19 int32 = 42
                        _ = v19
                    }
                    _ = fi10
                }
                _ = fi8
                _ = v9
            }
            v7 = ((v7 / 3.14159) - v7)
            v7 -= (0.1 / 0.1)
        }
        var v20 float32 = (1e10 * ((((-2.25 - ((float32(v7) + 0.0) - ((3.14159 / (1.5 + (1e10 * -2.25))) - (-2.25 * 1e10)))) * 0.0) / -2.25) + 1.5))
        _ = v20
        _ = v7
    }
    var v21 uint64 = 18446744073709551615
    _ = v21
    return -1
}
func genFn1(p0 int32, p1 int8, p2 int64) int8 {
    v22 := GenS1{F0: 1000}
    for fi23 := 0; fi23 < 3; fi23++ {
        p0 = (p0 & 1)
        _ = fi23
    }
    fmt.Println(p2, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 1
}

func main() {
    var v1 int = 1
    v2 := []string{"\n\t", "hello"}
    var v3 string = ("a" + "\n\t")
    v4 := []int{(v1 | v1), v1, (v1 + 127)}
    v2 = append(v2, "\n\t")
    v1 = (v1 % (((v1 / (((v1 | (^(v1)))) | 1))) | 1))
    if (243 >= 2147483647) {
        var v24 uint32 = 956
        if false {
            if (false && (false && true)) {
                fmt.Println(v24)
            }
            for fi25 := 0; fi25 < 20; fi25++ {
                fmt.Println(v3)
                _ = fi25
            }
        }
        v24 -= v24
        _ = v24
    } else {
        if false {
            var v26 string = ("" + "hello")
            _ = v26
        } else {
            if true {
                for fi27 := 0; fi27 < 9; fi27++ {
                    if (false && (!(197 > 42))) {
                        switch (v1 - (-(v1))) {
                        case 2147483647:
                            var v28 int64 = 42
                            _ = v28
                        default:
                            v3 = v3
                        }
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                        v29 := GenS1{F0: 42}
                        _ = v29
                    } else {
                        v30 := GenS2{F0: 0, F1: (!(true && (true && false))), F2: 18446744073709551615}
                        if (true && (42 >= 65535)) {
                            var v31 string = "foo bar"
                            _ = v31
                        }
                        _ = v30
                    }
                    if (32767 < 255) {
                        fmt.Println(v1, v3)
                        switch "a" {
                        case "hello":
                            for fi32 := 0; fi32 < 3; fi32++ {
                                v1 = 127
                                _ = fi32
                            }
                        case "foo bar":
                            for fi33 := 0; fi33 < 15; fi33++ {
                                switch "hello" {
                                case "a":
                                    var v34 int32 = 42
                                    _ = v34
                                case "foo bar":
                                    v35 := GenS2{F0: 255, F1: (878 == 65535), F2: 18446744073709551615}
                                    _ = v35
                                default:
                                    fmt.Println(v1, fi27)
                                }
                                _ = fi33
                            }
                        case "":
                            for fi36 := 0; fi36 < 16; fi36++ {
                                var v37 int32 = -1
                                var v38 float32 = float32(3.14159)
                                v39 := GenS1{F0: 0}
                                if len(v2) > 0 {
                                    v2[0] = v3
                                }
                                v2 = append(v2, "\n\t")
                                _ = fi36
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                        }
                    }
                    _ = fi27
                }
            }
        }
        v40 := GenS2{F0: 170, F1: true, F2: 0}
        var v41 int64 = -1
        _ = v40
        _ = v41
    }
    if (-9223372036854775808 >= 0) {
        for fi42 := 0; fi42 < 3; fi42++ {
            var v43 int8 = 42
            var v44 int16 = 42
            _ = fi42
            _ = v43
            _ = v44
        }
    }
    var v45 uint32 = 2147483647
    for fi46 := 0; fi46 < 7; fi46++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v47 int32 = -232
        for fi48 := 0; fi48 < 12; fi48++ {
            v3 = (v3 + v3)
            if false {
                fmt.Println(fi48, fi48)
                _ = len(v4)
                v49 := GenS1{F0: 0}
                _ = v49
            }
            _ = fi48
        }
        _ = fi46
        _ = v47
    }
    v50 := GenS3{F0: 736}
    var v51 float64 = 0.1
    switch v3 {
    case "foo bar":
        for fi52 := 0; fi52 < 19; fi52++ {
            if (!true) {
                if (-289 >= 32767) {
                    switch (v1 | (v1 ^ (fi52 & (fi52 | 4294967295)))) {
                    case 561:
                        if len(v4) > 0 {
                            v4[0] = 127
                        }
                    case 32767:
                        var v53 uint32 = v45
                        _ = v53
                    }
                }
                var v54 uint = 65535
                v1 += fi52
                _ = v54
            }
            var v55 int8 = -29
            _ = fi52
            _ = v55
        }
    default:
        var v56 int64 = 42
        _ = v56
    }
    if (true && (1 <= 0)) {
        for fi57 := 0; fi57 < 12; fi57++ {
            var v58 int = (v1 * v1)
            for fi59 := 0; fi59 < 9; fi59++ {
                var v60 string = ("\n\t" + v3)
                v45 = v45
                if len(v4) > 0 {
                    v4[0] = (v58 ^ fi57)
                }
                _ = fi59
                _ = v60
            }
            var v61 float64 = (0.0 - 1.5)
            _ = fi57
            _ = v58
            _ = v61
        }
        switch v3 {
        case "hello":
            v51 = float64(((float32(0.0) + 0.0) / (0.1 * (float32((v51 - v51)) / 1.0))))
        default:
            if (-564 != 0) {
                fmt.Println(v3, v51)
                v51 += float64(0.0)
                switch ("foo bar" + ("a" + v3)) {
                case "":
                    for fi62 := 0; fi62 < 4; fi62++ {
                        var v63 uint64 = 4294967295
                        var v64 int64 = -9223372036854775808
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                case "\n\t":
                    if (-37 >= 127) {
                        for fi65 := 0; fi65 < 17; fi65++ {
                            if true {
                                for fi66 := 0; fi66 < 17; fi66++ {
                                    var v67 int32 = -208
                                    _ = fi66
                                    _ = v67
                                }
                            } else {
                                if ((true || true) && true) {
                                    var v68 uint64 = 18446744073709551615
                                    v68 += 18446744073709551615
                                    _ = v68
                                } else {
                                    var v69 uint64 = 42
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v69
                                }
                                _ = len(v4)
                                var v70 int = fi65
                                _ = v70
                            }
                            v1 -= v1
                            v71 := GenS2{F0: 0, F1: false, F2: 557}
                            _ = fi65
                            _ = v71
                        }
                        switch ("a" + v3) {
                        case "foo bar":
                            if false {
                                if (true || (!true)) {
                                    for fi72 := 0; fi72 < 19; fi72++ {
                                        var v73 int16 = 32767
                                        _ = fi72
                                        _ = v73
                                    }
                                    if false {
                                        v74 := GenS4{F0: -32768, F1: 347}
                                        switch 32768 {
                                        case 42:
                                            var v75 uint8 = 42
                                            _ = v75
                                        case 0:
                                            v76 := GenS5{F0: 1, F1: (100 <= 100)}
                                            _ = v76
                                        case 9223372036854775807:
                                            for fi77 := 0; fi77 < 17; fi77++ {
                                                v78 := GenS5{F0: 4294967295, F1: (!(v51 != ((v51 * (v51 - 3.14159)) / v51)))}
                                                _ = fi77
                                                _ = v78
                                            }
                                        }
                                        for fi79 := 0; fi79 < 20; fi79++ {
                                            var v80 int8 = -29
                                            _ = fi79
                                            _ = v80
                                        }
                                        _ = v74
                                    }
                                }
                                if (4294967295 != 916) {
                                    v45 = v45
                                    v45 = 1
                                    if (42 == 234) {
                                        for fi81 := 0; fi81 < 14; fi81++ {
                                            v82 := GenS6{F0: 0, F1: v45}
                                            _ = fi81
                                            _ = v82
                                        }
                                        var v83 int = (v1 * v1)
                                        v45 -= (v45 % ((v45) | 1))
                                        _ = v83
                                    } else {
                                        fmt.Println(v3, v51, v3)
                                        _ = len(v2)
                                    }
                                }
                            } else {
                                if (float64(1e10) == 3.14159) {
                                    for fi84 := 0; fi84 < 10; fi84++ {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v85 int32 = 1
                                        var v86 int8 = 1
                                        _ = fi84
                                        _ = v85
                                        _ = v86
                                    }
                                }
                                var v87 uint16 = 0
                                v88 := GenS7{F0: 9223372036854775807}
                                _ = v87
                                _ = v88
                            }
                        case "hello":
                            var v89 uint64 = 0
                            _ = v89
                        default:
                            var v90 uint8 = 2
                            _ = v90
                        }
                    } else {
                        v2 = append(v2, v3)
                        switch (v3 + ((("hello" + v3) + v3) + (v3 + v3))) {
                        case "\n\t":
                            v91 := GenS8{F0: false, F1: 18446744073709551615}
                            _ = v91
                        case "hello":
                            v92 := GenS3{F0: 420}
                            _ = v92
                        case "":
                            if ((878 == 255) && false) {
                                v1 -= v1
                                fmt.Println(v1, v3, v51)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                _ = len(v4)
                                var v93 uint16 = 255
                                _ = v93
                            }
                        default:
                            switch v1 {
                            case -2147483648:
                                v94 := GenS5{F0: 255, F1: (255 == 230)}
                                _ = v94
                            case 615:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case -109:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                    }
                case "hello":
                    var v95 int16 = 42
                    _ = v95
                }
            } else {
                var v96 int32 = 1
                v3 = ((v3 + ("hello" + (("" + "foo bar") + ("foo bar" + "\n\t")))) + "a")
                v3 = v3
                _ = v96
            }
        }
    }
    v97 := GenS5{F0: 65535, F1: (!(127 <= (v1 | v1)))}
    v45 = v45
    var v98 float32 = (-2.25 * 0.0)
    for fi99 := 0; fi99 < 15; fi99++ {
        v100 := GenS1{F0: -1}
        for fi101 := 0; fi101 < 15; fi101++ {
            fmt.Println(v3, fi101)
            switch v3 {
            case "":
                _ = len(v4)
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi101
        }
        var v102 int64 = 1
        _ = fi99
        _ = v100
        _ = v102
    }
    switch v3 {
    case "\n\t":
        fmt.Println(v45, v45, v3)
    case "foo bar":
        var v103 string = v3
        _ = v103
    default:
        v104 := GenS7{F0: v1}
        _ = v104
    }
    if len(v4) > 0 {
        v4[0] = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v51:", v51)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v50
    _ = v51
    _ = v97
    _ = v98
}
