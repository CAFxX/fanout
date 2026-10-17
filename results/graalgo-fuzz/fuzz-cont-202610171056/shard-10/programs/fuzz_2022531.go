package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int64; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 int16; F1 int32; F2 bool }
type GenS8 struct { F0 uint64; F1 int; F2 int }
func genFn0(p0 string) bool {
    fmt.Println(p0)
    var v4 string = "a"
    var v5 uint16 = 332
    v6 := GenS1{F0: 8}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (42 < -1)
}

func main() {
    v1 := []string{"hello", ("\n\t" + "foo bar")}
    var v2 uint = 2
    var v3 int = 668
    for fi7 := 0; fi7 < 15; fi7++ {
        if (((((("hello" + "foo bar") + "foo bar") + "\n\t") + "hello") != "\n\t") && true) {
            v2 = 255
            var v8 uint8 = 122
            v8 -= (v8 << ((v2) & 63))
            _ = v8
        }
        v1 = append(v1, "\n\t")
        var v9 int32 = 42
        _ = fi7
        _ = v9
    }
    if (879 > 0) {
        if true {
            v10 := GenS2{F0: (!((!true) && true))}
            for fi11 := 0; fi11 < 18; fi11++ {
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := GenS3{F0: (v3 * (-(v3))), F1: int64(v2), F2: 0}
                    _ = fi12
                    _ = v13
                }
                _ = fi11
            }
            v14 := GenS4{F0: (v3 + 127)}
            _ = v10
            _ = v14
        } else {
            var v15 uint64 = 4294967295
            _ = v15
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 -= 42
        switch (("\n\t" + "foo bar") + "\n\t") {
        case "\n\t":
            switch "a" {
            case "\n\t":
                _ = len(v1)
            case "":
                var v16 int64 = -1
                _ = v16
            default:
                v17 := GenS5{F0: "hello", F1: 922}
                _ = v17
            }
        default:
            fmt.Println(v2)
        }
    }
    if true {
        v18 := GenS2{F0: (-1 <= 588)}
        v19 := GenS2{F0: false}
        var v20 uint64 = 2
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 int16 = 0
    fmt.Println(v21, v2)
    switch "hello" {
    case "hello":
        v21 += (^(v21))
    default:
        switch ("" + "\n\t") {
        case "":
            var v22 int = -200
            _ = v22
        case "foo bar":
            switch (v3 ^ (^(v3))) {
            case 4294967295:
                if ((1.5 * 1.5) >= 1.5) {
                    _ = len(v1)
                    if true {
                        var v23 uint = v2
                        v2 = v23
                        fmt.Println(v23, v2)
                        _ = v23
                    } else {
                        switch -2147483648 {
                        case -9223372036854775808:
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                        case 128:
                            if (1 != -1) {
                                v21 = v21
                            } else {
                                for fi24 := 0; fi24 < 5; fi24++ {
                                    v25 := GenS1{F0: 1}
                                    _ = fi24
                                    _ = v25
                                }
                                for fi26 := 0; fi26 < 18; fi26++ {
                                    var v27 uint32 = 2147483647
                                    v3 = (fi26 % ((4294967295) | 1))
                                    _ = fi26
                                    _ = v27
                                }
                                v3 = (v3 % ((v3) | 1))
                            }
                        case 4294967295:
                            for fi28 := 0; fi28 < 4; fi28++ {
                                v21 += 1000
                                v29 := GenS3{F0: fi28, F1: -1, F2: 597}
                                switch v3 {
                                case -176:
                                    if (!false) {
                                        fmt.Println(fi28, v3, v3)
                                    } else {
                                        for fi30 := 0; fi30 < 2; fi30++ {
                                            switch (fi28 - -574) {
                                            case 32768:
                                                switch v3 {
                                                case 228:
                                                    v31 := GenS1{F0: 155}
                                                    _ = v31
                                                default:
                                                    _ = len(v1)
                                                }
                                            case 0:
                                                var v32 uint16 = 32767
                                                _ = v32
                                            case 9223372036854775807:
                                                var v33 int16 = 42
                                                _ = v33
                                            }
                                            _ = fi30
                                        }
                                        v21 -= v21
                                        v3 -= v3
                                    }
                                }
                                _ = fi28
                                _ = v29
                            }
                        }
                    }
                }
            default:
                if false {
                    fmt.Println(v2)
                    if (-2.25 > float64(-2.25)) {
                        var v34 uint8 = 2
                        fmt.Println(v34, v2)
                        _ = v34
                    }
                    v35 := GenS4{F0: (-(v3))}
                    _ = v35
                }
            }
        case "hello":
            v21 += v21
        default:
            var v36 bool = true
            _ = v36
        }
    }
    fmt.Println(v2)
    if (v21 != -1) {
        switch (("a" + ("foo bar" + "a")) + "hello") {
        case "foo bar":
            for fi37 := 0; fi37 < 8; fi37++ {
                v3 = fi37
                v3 = fi37
                _ = fi37
            }
        default:
            var v38 uint = 65535
            _ = v38
        }
        if len(v1) > 0 {
            v1[0] = (((("\n\t" + "hello") + "") + "") + "")
        }
    } else {
        var v39 uint32 = 216
        var v40 int16 = v21
        var v41 uint8 = 128
        _ = v39
        _ = v40
        _ = v41
    }
    var v42 uint64 = 1
    if (1e10 == 1e10) {
        switch v3 {
        case 4294967295:
            v42 -= (v42 * (v42 % ((v42) | 1)))
        case 105:
            for fi43 := 0; fi43 < 14; fi43++ {
                if false {
                    fmt.Println(v3)
                } else {
                    switch fi43 {
                    case 1:
                        v2 -= (^(v2))
                    case 928:
                        v42 = (^(v42))
                    case -644:
                        if (true && false) {
                            for fi44 := 0; fi44 < 20; fi44++ {
                                var v45 float64 = 0.1
                                v21 = 1000
                                if (1.5 < 1.5) {
                                    var v46 string = (("" + "a") + "\n\t")
                                    _ = v46
                                } else {
                                    var v47 uint8 = 255
                                    _ = v47
                                }
                                _ = fi44
                                _ = v45
                            }
                        } else {
                            v2 -= v2
                        }
                    }
                    v21 += (^(v21))
                }
                _ = fi43
            }
        }
        v3 -= (v3 % (((v3 ^ (v3 - (v3 + v3)))) | 1))
    }
    v48 := GenS4{F0: int(v42)}
    var v49 string = "foo bar"
    if (2 >= (v42 & 558)) {
        var v50 float32 = (0.0 * (-2.25 * 3.14159))
        var v51 int32 = 2147483647
        _ = v50
        _ = v51
    }
    for fi52 := 0; fi52 < 12; fi52++ {
        switch 32767 {
        case 2147483647:
            switch 871 {
            case 32768:
                v53 := GenS6{F0: 1, F1: 18446744073709551615}
                _ = v53
            }
        default:
            switch "hello" {
            case "foo bar":
                var v54 uint8 = 0
                _ = v54
            default:
                v55 := GenS4{F0: (fi52 % (((fi52 | (v3 | (v3 * fi52)))) | 1))}
                _ = v55
            }
        }
        _ = fi52
    }
    for fi56 := 0; fi56 < 14; fi56++ {
        _ = len(v1)
        _ = len(v1)
        _ = fi56
    }
    if (false && true) {
        if (!true) {
            if false {
                v57 := GenS7{F0: (v21 - (v21 << ((65535) & 63))), F1: 42, F2: false}
                var v58 uint8 = 110
                switch -787 {
                case 2147483647:
                    for fi59 := 0; fi59 < 18; fi59++ {
                        if (!(1 < (^(v42)))) {
                            if (!(1 <= v58)) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            v3 -= fi59
                            v60 := GenS1{F0: (v58 % (((v58 >> ((uint(v42)) & 63))) | 1))}
                            _ = v60
                        } else {
                            if false {
                                v49 = "a"
                                var v61 float64 = 0.1
                                v42 = v42
                                _ = v61
                            }
                        }
                        var v62 uint32 = 0
                        _ = fi59
                        _ = v62
                    }
                case 32767:
                    v63 := GenS3{F0: (v3 | v3), F1: 0, F2: (v21 << ((4294967295) & 63))}
                    _ = v63
                }
                _ = v57
                _ = v58
            } else {
                for fi64 := 0; fi64 < 10; fi64++ {
                    var v65 uint = 1
                    _ = fi64
                    _ = v65
                }
                if (2147483647 <= 881) {
                    var v66 int = v3
                    _ = len(v1)
                    fmt.Println(v21)
                    _ = v66
                }
            }
        }
        for fi67 := 0; fi67 < 7; fi67++ {
            for fi68 := 0; fi68 < 20; fi68++ {
                v42 -= (v42 / ((v42) | 1))
                _ = fi68
            }
            v69 := GenS8{F0: (v42 >> ((735) & 63)), F1: fi67, F2: (fi67 & (fi67 * (fi67 ^ (v3 % ((int(v2)) | 1)))))}
            _ = fi67
            _ = v69
        }
    } else {
        v70 := GenS1{F0: 128}
        var v71 int8 = 41
        v72 := GenS2{F0: ((v21 << ((v2) & 63)) >= (-(v21)))}
        _ = v70
        _ = v71
        _ = v72
    }
    switch -256 {
    case -463:
        for fi73 := 0; fi73 < 5; fi73++ {
            if (42 == 36) {
                fmt.Println(v49)
                if false {
                    v21 += v21
                    var v74 int32 = 392
                    switch -2147483648 {
                    case 32767:
                        if false {
                            v75 := GenS1{F0: 42}
                            _ = v75
                        }
                    default:
                        var v76 uint8 = 1
                        _ = v76
                    }
                    _ = v74
                } else {
                    if len(v1) > 0 {
                        v1[0] = v49
                    }
                    fmt.Println(v42, v3, fi73)
                }
                v1 = append(v1, (v49 + v49))
            }
            var v77 int8 = 127
            var v78 string = v49
            _ = fi73
            _ = v77
            _ = v78
        }
    }
    if len(v1) > 0 {
        v1[0] = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v42
    _ = v48
    _ = v49
}
