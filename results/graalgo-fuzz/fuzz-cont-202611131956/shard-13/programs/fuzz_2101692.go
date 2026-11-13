package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int16; F2 bool }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 int16; F1 string; F2 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 uint64; F1 uint32; F2 int8 }

func main() {
    var v1 int16 = 1000
    v2 := []string{}
    v3 := make(map[int]string)
    var v4 string = ("\n\t" + "foo bar")
    var v5 int64 = -858
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            fmt.Println(v1)
            fmt.Println(v4, v1, v5)
            var v8 float64 = (0.0 - 0.0)
            _ = fi7
            _ = v8
        }
        v2 = append(v2, "foo bar")
        var v9 uint64 = 2
        _ = fi6
        _ = v9
    }
    v10 := GenS1{F0: 42}
    v1 = v1
    var v11 string = v4
    for fi12 := 0; fi12 < 12; fi12++ {
        if (0 > 18446744073709551615) {
            v13 := GenS2{F0: 0, F1: (v1 / ((v1) | 1)), F2: false}
            switch -744 {
            case -9223372036854775808:
                switch "hello" {
                case "hello":
                    var v14 uint64 = 1
                    _ = v14
                case "foo bar":
                    switch "\n\t" {
                    case "\n\t":
                        var v15 string = ("a" + v11)
                        _ = v15
                    case "":
                        v16 := GenS2{F0: 1, F1: (v1 | v1), F2: false}
                        _ = v16
                    case "a":
                        v17 := GenS1{F0: -56}
                        _ = v17
                    default:
                        for fi18 := 0; fi18 < 2; fi18++ {
                            v1 -= v1
                            v1 -= 1
                            _ = fi18
                        }
                    }
                case "":
                    v1 = v1
                default:
                    v5 = (-(v5))
                }
            default:
                switch fi12 {
                case 1:
                    if ((42 > 97) || true) {
                        if false {
                            v19 := GenS3{F0: (v5 * 0), F1: 255}
                            if len(v2) > 0 {
                                v2[0] = ("foo bar" + (v4 + v11))
                            }
                            _ = v19
                        }
                        v20 := GenS2{F0: 1, F1: int16(v5), F2: false}
                        _ = v20
                    } else {
                        if true {
                            var v21 int8 = -128
                            var v22 uint64 = 0
                            switch v11 {
                            case "":
                                var v23 uint64 = (v22 - v22)
                                _ = v23
                            case "a":
                                v22 = v22
                            default:
                                var v24 string = v4
                                _ = v24
                            }
                            _ = v21
                            _ = v22
                        }
                    }
                }
            }
            if (v5 >= 0) {
                v11 = v11
                for fi25 := 0; fi25 < 2; fi25++ {
                    if (v11 > (("hello" + "foo bar") + (v11 + (v11 + v11)))) {
                        if true {
                            _ = len(v3)
                        }
                        var v26 uint8 = 42
                        _ = v26
                    }
                    _ = len(v2)
                    if len(v2) > 0 {
                        v2[0] = v4
                    }
                    _ = fi25
                }
                if (false && true) {
                    fmt.Println(v5, v5)
                    switch 2147483647 {
                    case 0:
                        v27 := GenS4{F0: true}
                        _ = v27
                    case -1:
                        v1 += (v1 - 294)
                    case 32768:
                        for fi28 := 0; fi28 < 3; fi28++ {
                            v4 = ((v4 + "a") + v4)
                            delete(v3, fi12)
                            var v29 int32 = 42
                            if ((true && (897 > 933)) || false) {
                                v30 := GenS4{F0: (true && true)}
                                _ = v30
                            }
                            _ = fi28
                            _ = v29
                        }
                    default:
                        v4 = v4
                    }
                } else {
                    delete(v3, 2147483647)
                    var v31 int8 = -128
                    v1 -= v1
                    _ = v31
                }
            }
            _ = v13
        }
        v1 = (v1 / ((v1) | 1))
        switch 868 {
        case 407:
            for fi32 := 0; fi32 < 13; fi32++ {
                v1 = 1000
                fmt.Println(v11, v1)
                _ = fi32
            }
        }
        _ = fi12
    }
    fmt.Println(v4, v4)
    var v33 int16 = -1
    v11 = v11
    for fi34 := 0; fi34 < 19; fi34++ {
        var v35 string = v11
        var v36 uint8 = 255
        if (641 > 2) {
            v37 := GenS5{F0: 42, F1: 0}
            if false {
                if (v36 <= 64) {
                    v11 = v11
                } else {
                    var v38 uint64 = 18446744073709551615
                    var v39 bool = true
                    switch (("\n\t" + (v4 + ("hello" + v35))) + "\n\t") {
                    case "foo bar":
                        for fi40 := 0; fi40 < 4; fi40++ {
                            switch v35 {
                            case "\n\t":
                                if (!true) {
                                    var v41 uint16 = 2
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v33, v33)
                                    _ = v41
                                }
                            }
                            _ = fi40
                        }
                    case "hello":
                        if (!false) {
                            var v42 int = fi34
                            _ = v42
                        }
                    default:
                        if v39 {
                            if v39 {
                                v35 = (v4 + ((("\n\t" + v4) + v11) + "hello"))
                            } else {
                                switch v11 {
                                case "foo bar":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            }
                            v43 := GenS4{F0: false}
                            if v39 {
                                if (uint32(v38) != 4294967295) {
                                    if len(v2) > 0 {
                                        v2[0] = (v4 + v35)
                                    }
                                    switch (("a" + ("foo bar" + "")) + v35) {
                                    case "hello":
                                        delete(v3, fi34)
                                    default:
                                        var v44 float64 = (float64(0.0) * 1e10)
                                        _ = v44
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = v35
                                    }
                                }
                                for fi45 := 0; fi45 < 20; fi45++ {
                                    switch (fi45 + (fi45 >> ((0) & 63))) {
                                    case 42:
                                        v46 := GenS1{F0: -128}
                                        _ = v46
                                    case 4294967295:
                                        var v47 int64 = (^(v5))
                                        _ = v47
                                    case -9223372036854775808:
                                        v4 = "\n\t"
                                    }
                                    if v39 {
                                        switch v35 {
                                        case "\n\t":
                                            if true {
                                                switch int(v1) {
                                                case -670:
                                                    if true {
                                                        fmt.Println(v38, fi45)
                                                        v2 = append(v2, "foo bar")
                                                    } else {
                                                        var v48 uint16 = 255
                                                        delete(v3, fi45)
                                                        _ = v48
                                                    }
                                                default:
                                                    v4 = "hello"
                                                    fmt.Println(fi45)
                                                    var v49 int8 = 42
                                                    _ = v49
                                                }
                                                v50 := v3[fi45]
                                                v51 := v3[(fi34 / ((32767) | 1))]
                                                _ = v50
                                                _ = v51
                                            } else {
                                                fmt.Println(fi45, v4)
                                            }
                                        case "hello":
                                            v52 := GenS4{F0: ((v38 + 2) >= 2)}
                                            _ = v52
                                        case "a":
                                            if (0 != 846) {
                                                v36 -= 127
                                                if len(v2) > 0 {
                                                    v2[0] = v11
                                                }
                                                if v39 {
                                                    fmt.Println(v38)
                                                    var v53 float64 = -2.25
                                                    switch -1 {
                                                    case 2147483647:
                                                        if v39 {
                                                            switch fi34 {
                                                            case 2147483647:
                                                                fmt.Println(v53, v11)
                                                            case 199:
                                                                if ((v38 / ((v38) | 1)) > 1) {
                                                                    if v39 {
                                                                        _ = len(v3)
                                                                        v1 -= v33
                                                                        var v54 int32 = -1
                                                                        _ = v54
                                                                    } else {
                                                                        fmt.Println(v4, v36, v4)
                                                                        var v55 int32 = 384
                                                                        _ = v55
                                                                    }
                                                                }
                                                            }
                                                            if (!v39) {
                                                                if (!(57 <= 0)) {
                                                                    var v56 float32 = 0.1
                                                                    switch "" {
                                                                    case "":
                                                                        v57 := GenS1{F0: 100}
                                                                        _ = v57
                                                                    default:
                                                                        fmt.Println(v5, v38, v39)
                                                                    }
                                                                    if (false && v39) {
                                                                        v5 += v5
                                                                    } else {
                                                                        var v58 uint32 = 551
                                                                        _ = v58
                                                                    }
                                                                    _ = v56
                                                                }
                                                                var v59 int = (^(fi34))
                                                                var v60 float64 = 1.5
                                                                _ = v59
                                                                _ = v60
                                                            } else {
                                                                switch (fi45 * fi45) {
                                                                case -2147483648:
                                                                    v61 := GenS5{F0: 0, F1: 0}
                                                                    _ = v61
                                                                case 405:
                                                                    if v39 {
                                                                        switch -810 {
                                                                        case -516:
                                                                            v11 = v11
                                                                        case 32767:
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                v35 = (v35 + v4)
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if (!v39) {
                                                                            v53 -= float64((0.0 / float32(v53)))
                                                                            v62 := GenS6{F0: ((v4 + "hello") + v11), F1: 255}
                                                                            _ = v62
                                                                        }
                                                                        var v63 uint = 2
                                                                        v33 = (v1 | -1)
                                                                        v5 = v5
                                                                        _ = v63
                                                                    }
                                                                case 0:
                                                                    v11 = v35
                                                                default:
                                                                    var v64 uint64 = (v38 & (v38 * v38))
                                                                    _ = v64
                                                                }
                                                                v65 := GenS6{F0: "hello", F1: 324}
                                                                switch fi45 {
                                                                case -1:
                                                                    v38 = v38
                                                                case -185:
                                                                    if v39 {
                                                                        switch v11 {
                                                                        case "":
                                                                            v66 := v3[(fi45 | (fi34 / ((fi34) | 1)))]
                                                                            _ = v66
                                                                        case "a":
                                                                            if ("foo bar" != "hello") {
                                                                                v11 = v4
                                                                                fmt.Println(v39, v36)
                                                                                v67 := GenS7{F0: v33, F1: v4, F2: 4294967295}
                                                                                _ = v67
                                                                            }
                                                                        case "hello":
                                                                            var v68 int16 = (-(v1))
                                                                            _ = v68
                                                                        }
                                                                    } else {
                                                                        var v69 float64 = v53
                                                                        _ = v69
                                                                    }
                                                                case 0:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v35
                                                                    }
                                                                default:
                                                                    var v70 int = (fi45 | int(v38))
                                                                    _ = v70
                                                                }
                                                                _ = v65
                                                            }
                                                            var v71 uint16 = 0
                                                            _ = v71
                                                        }
                                                    default:
                                                        v53 = float64(0.1)
                                                    }
                                                    fmt.Println(fi34)
                                                    switch -500 {
                                                    case 0:
                                                        var v72 int32 = 0
                                                        _ = v72
                                                    }
                                                    _ = v53
                                                }
                                            }
                                        default:
                                            delete(v3, (fi34 + fi34))
                                        }
                                        fmt.Println(v1, v39)
                                    }
                                    _ = fi45
                                }
                                fmt.Println(v38)
                            }
                            _ = v43
                        }
                    }
                    _ = v38
                    _ = v39
                }
                var v73 uint8 = (^(v36))
                var v74 uint32 = 65535
                _ = v73
                _ = v74
            }
            v75 := GenS8{F0: (fi34 * (-(fi34)))}
            _ = v37
            _ = v75
        } else {
            v76 := GenS4{F0: true}
            switch v11 {
            case "foo bar":
                var v77 string = v11
                _ = v77
            case "hello":
                switch ("foo bar" + v11) {
                case "":
                    v5 -= v5
                }
            }
            var v78 string = "foo bar"
            _ = v76
            _ = v78
        }
        _ = fi34
        _ = v35
        _ = v36
    }
    switch v4 {
    case "hello":
        v4 = (v4 + (v11 + "hello"))
    case "\n\t":
        for fi79 := 0; fi79 < 9; fi79++ {
            v80 := GenS9{F0: 65535, F1: 9223372036854775807}
            fmt.Println(v1)
            for fi81 := 0; fi81 < 19; fi81++ {
                var v82 uint16 = 308
                var v83 uint = 4294967295
                v2 = append(v2, v11)
                _ = fi81
                _ = v82
                _ = v83
            }
            _ = fi79
            _ = v80
        }
    case "a":
        v84 := GenS10{F0: 2, F1: 2147483647, F2: 100}
        _ = v84
    default:
        v4 = v4
    }
    v85 := v3[-788]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v85
}
