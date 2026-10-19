package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 string; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 int; F2 int16 }
type GenS6 struct { F0 int16; F1 uint32; F2 bool }
type GenS7 struct { F0 int; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -122
    v3 := make(map[string]bool)
    var v4 float32 = 1e10
    var v5 string = ((("" + "foo bar") + "") + "")
    var v6 uint64 = 1
    if true {
        var v7 uint16 = 206
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int = fi8
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        if false {
            if ((0.1 > (-2.25 - 0.1)) && false) {
                if (!(18446744073709551615 == 31)) {
                    var v11 uint16 = 2
                    var v12 int = 871
                    _ = v11
                    _ = v12
                } else {
                    v2 = (v2 >> ((2) & 63))
                }
                if ((^(v6)) > (v6 % (((v6 % ((42) | 1))) | 1))) {
                    v2 += (v2 / ((v2) | 1))
                }
                fmt.Println(v5, fi10)
            } else {
                v13 := v3[("hello" + (v5 + v5))]
                v14 := []int{fi10}
                _ = v13
                _ = v14
            }
            var v15 uint = 2
            _ = v15
        } else {
            _ = len(v3)
            var v16 uint = 4294967295
            _ = v16
        }
        _ = fi10
    }
    v17 := []int{}
    v4 += v4
    v2 -= v2
    for fi18 := 0; fi18 < 2; fi18++ {
        if true {
            var v19 int16 = 1
            v20 := GenS1{F0: 42}
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            switch v5 {
            case "foo bar":
                v2 -= (v2 * v2)
            case "hello":
                _ = len(v3)
            default:
                if false {
                    v1[fi21] = 4294967295
                    v6 -= 811
                }
            }
            _ = fi21
        }
        for fi22 := 0; fi22 < 18; fi22++ {
            delete(v1, fi22)
            if (32767 == 2) {
                fmt.Println(fi22, v6, v2)
                var v23 uint = 746
                _ = v23
            }
            _ = fi22
        }
        _ = fi18
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        switch -386 {
        case -2147483648:
            fmt.Println(v5, v4)
        case 32768:
            v25 := GenS1{F0: 9223372036854775807}
            _ = v25
        case 82:
            switch fi24 {
            case -345:
                if true {
                    if (!(129 == 128)) {
                        var v26 uint32 = uint32(v6)
                        fmt.Println(v5)
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 9; fi27++ {
                        v28 := GenS2{F0: 663, F1: "\n\t", F2: -474}
                        _ = fi27
                        _ = v28
                    }
                }
            }
        }
        _ = fi24
    }
    var v29 int = -562
    if (false || false) {
        if true {
            var v30 int8 = v2
            _ = v30
        } else {
            switch v5 {
            case "a":
                v1[(v29 / ((v29) | 1))] = 0
            case "foo bar":
                var v31 uint32 = 42
                _ = v31
            }
            var v32 float32 = v4
            _ = v32
        }
        var v33 int8 = (v2 >> ((18446744073709551615) & 63))
        _ = v33
    } else {
        switch v5 {
        case "":
            for fi34 := 0; fi34 < 19; fi34++ {
                var v35 int16 = -639
                _ = fi34
                _ = v35
            }
        }
        v36 := GenS3{F0: -146}
        fmt.Println(v6, v5, v4)
        _ = v36
    }
    var v37 int = (v29 % ((v29) | 1))
    v38 := v1[(v37 << ((2) & 63))]
    var v39 int = (-(v37))
    fmt.Println(v5, v38, v4)
    if (v5 == "a") {
        v40 := GenS3{F0: 42}
        if (v29 > v37) {
            v41 := GenS4{F0: -1}
            _ = v41
        } else {
            for fi42 := 0; fi42 < 2; fi42++ {
                if (2 > 152) {
                    v6 = (v6 / ((1) | 1))
                }
                fmt.Println(v4, v4)
                switch v5 {
                case "hello":
                    fmt.Println(v38)
                case "":
                    var v43 uint = 845
                    _ = v43
                case "a":
                    var v44 int = v39
                    _ = v44
                }
                _ = fi42
            }
            for fi45 := 0; fi45 < 12; fi45++ {
                var v46 int = (v37 + v29)
                var v47 int64 = 0
                fmt.Println(v37)
                _ = fi45
                _ = v46
                _ = v47
            }
        }
        for fi48 := 0; fi48 < 17; fi48++ {
            for fi49 := 0; fi49 < 15; fi49++ {
                v38 = (v38 + v38)
                _ = fi49
            }
            if true {
                v1[-1] = 2147483648
                v38 -= (v38 >> ((47) & 63))
                v50 := GenS5{F0: -2147483648, F1: v37, F2: -1}
                _ = v50
            } else {
                if true {
                    var v51 uint = 185
                    v29 += (v29 << (((^(v51))) & 63))
                    v52 := GenS3{F0: -1}
                    _ = v51
                    _ = v52
                }
                fmt.Println(v4)
            }
            _ = fi48
        }
        _ = v40
    } else {
        var v53 int8 = (v2 ^ (v2 >> ((42) & 63)))
        switch ("\n\t" + "") {
        case "a":
            var v54 int = v37
            _ = v54
        }
        switch (("a" + (v5 + "hello")) + "foo bar") {
        case "a":
            switch 2147483647 {
            case 990:
                if len(v17) > 0 {
                    v17[0] = -203
                }
            case 128:
                if (((v6 / (((v6 | (v6 - (v6 / ((18446744073709551615) | 1))))) | 1)) >= 0) && (673 > 2147483647)) {
                    v55 := GenS6{F0: 42, F1: 218, F2: false}
                    switch (^(v29)) {
                    case 1:
                        v56 := v1[349]
                        _ = v56
                    case -2147483648:
                        var v57 uint = 0
                        _ = v57
                    }
                    _ = v55
                }
            case -2147483648:
                v29 -= (^(v37))
            }
        case "hello":
            switch v5 {
            case "foo bar":
                switch (v39 * (-(v37))) {
                case 127:
                    if (21 > 1) {
                        fmt.Println(v37, v5)
                    }
                case 2147483647:
                    var v58 int32 = 2147483647
                    _ = v58
                default:
                    if true {
                        switch "foo bar" {
                        case "":
                            v59 := GenS4{F0: 1}
                            _ = v59
                        case "\n\t":
                            for fi60 := 0; fi60 < 20; fi60++ {
                                v53 -= 110
                                _ = fi60
                            }
                        case "foo bar":
                            if false {
                                for fi61 := 0; fi61 < 10; fi61++ {
                                    var v62 string = "a"
                                    if false {
                                        if true {
                                            v63 := GenS3{F0: 2147483647}
                                            _ = v63
                                        }
                                        v2 += (v2 - (^(v53)))
                                    } else {
                                        v6 += v6
                                        var v64 string = (v5 + ("foo bar" + "hello"))
                                        for fi65 := 0; fi65 < 15; fi65++ {
                                            var v66 int8 = (v2 | v2)
                                            v67 := GenS7{F0: int(v66), F1: v38}
                                            _ = fi65
                                            _ = v66
                                            _ = v67
                                        }
                                        _ = v64
                                    }
                                    _ = fi61
                                    _ = v62
                                }
                                v68 := GenS2{F0: 1, F1: "hello", F2: 9223372036854775807}
                                v4 = 3.14159
                                _ = v68
                            }
                        default:
                            switch (v39 % (((v39 << ((1) & 63))) | 1)) {
                            case 128:
                                var v69 uint = 616
                                _ = v69
                            case 42:
                                var v70 int8 = (^(v2))
                                _ = v70
                            case -2147483648:
                                if ((1.5 - 1e10) == ((((0.1 - -2.25) + 0.1) / 1.5) + 0.1)) {
                                    fmt.Println(v2, v2, v5)
                                } else {
                                    if false {
                                        var v71 float32 = v4
                                        var v72 uint8 = 0
                                        _ = v71
                                        _ = v72
                                    }
                                }
                            }
                        }
                        if false {
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                        }
                        v53 = v2
                    } else {
                        fmt.Println(v29, v38)
                        v73 := GenS6{F0: -258, F1: (v38 >> ((65535) & 63)), F2: false}
                        _ = v73
                    }
                }
            case "hello":
                if (v38 == 42) {
                    var v74 int32 = -2147483648
                    var v75 int32 = v74
                    _ = v74
                    _ = v75
                }
            case "":
                v76 := v1[v39]
                _ = v76
            default:
                if false {
                    v29 = (v29 << ((42) & 63))
                }
            }
        }
        _ = v53
    }
    for fi77 := 0; fi77 < 8; fi77++ {
        for fi78 := 0; fi78 < 2; fi78++ {
            var v79 float64 = 3.14159
            switch 4294967295 {
            case 9223372036854775807:
                v5 = ""
            case 32767:
                v1[fi77] = 1
                v4 += v4
                v80 := GenS5{F0: 0, F1: fi77, F2: int16(fi78)}
                _ = v80
            }
            _ = fi78
            _ = v79
        }
        _ = fi77
    }
    var v81 float64 = 3.14159
    var v82 string = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
}
