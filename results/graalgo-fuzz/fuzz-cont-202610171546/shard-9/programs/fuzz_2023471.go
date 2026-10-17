package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int64 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32; F1 string; F2 uint32 }
type GenS8 struct { F0 uint; F1 int32; F2 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int64; F1 int }
type GenS11 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint64) int32 {
    v5 := []int{0, 1, -1, -932}
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 float32 = -2.25
        var v8 int16 = -1
        if ((true || false) || ((!false) && (-580 <= 2147483647))) {
            v5 = append(v5, (fi6 * -1))
            if false {
                p1 = (^(p1))
            } else {
                v8 -= v8
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                switch fi6 {
                case 127:
                    var v9 float32 = v7
                    _ = v9
                case 1:
                    var v10 uint32 = 2147483648
                    _ = v10
                default:
                    v11 := GenS1{F0: 0, F1: -1, F2: int64(fi6)}
                    _ = v11
                }
            }
        } else {
            if false {
                var v12 int64 = -812
                var v13 int16 = v8
                _ = v12
                _ = v13
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            p0 -= p0
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if true {
        for fi14 := 0; fi14 < 11; fi14++ {
            for fi15 := 0; fi15 < 12; fi15++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi15
            }
            _ = fi14
        }
        v16 := GenS1{F0: 979, F1: int32(p1), F2: -207}
        v17 := GenS2{F0: 244, F1: -1}
        _ = v16
        _ = v17
    } else {
        p0 = p0
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 int16 = 0
            _ = fi18
            _ = v19
        }
        p0 -= (p0 >> ((65535) & 63))
    }
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    _ = v5
    return 2147483647
}

func main() {
    var v1 string = ("\n\t" + "a")
    var v2 uint64 = 743
    v3 := make(map[string]int)
    var v4 uint64 = v2
    v20 := []int{}
    if len(v20) > 0 {
        v20[0] = 1
    }
    if (((!false) || false) && (127 != 1)) {
        v21 := GenS2{F0: genFn0(2, 42), F1: 9223372036854775807}
        v2 = 2
        _ = v21
    }
    if true {
        delete(v3, v1)
        fmt.Println(v4, v4, v2)
        if (426 != 42) {
            v22 := GenS3{F0: -9223372036854775808}
            for fi23 := 0; fi23 < 12; fi23++ {
                v1 = ((v1 + "hello") + "hello")
                v1 = (v1 + ("hello" + (v1 + v1)))
                for fi24 := 0; fi24 < 18; fi24++ {
                    var v25 uint64 = 4294967295
                    switch v1 {
                    case "hello":
                        v4 -= v25
                    case "foo bar":
                        if (!(false || false)) {
                            v26 := GenS1{F0: 255, F1: -2147483648, F2: 0}
                            v4 = v4
                            _ = v26
                        } else {
                            var v27 bool = (0 != 85)
                            if (!v27) {
                                switch v1 {
                                case "":
                                    fmt.Println(v27, v1, v4)
                                    switch "\n\t" {
                                    case "a":
                                        fmt.Println(v4, v27)
                                    default:
                                        v25 += v25
                                    }
                                    if true {
                                        v1 = v1
                                    } else {
                                        switch (fi24 << ((2) & 63)) {
                                        case 0:
                                            v28 := GenS3{F0: 208}
                                            _ = v28
                                        case 616:
                                            fmt.Println(v25)
                                        }
                                        if ((0.1 + ((1e10 + (3.14159 + (1.5 * (0.1 / 3.14159)))) - float64((float32((0.0 - (1e10 / 1.0))) + -2.25)))) == (0.0 / 3.14159)) {
                                            v29 := GenS2{F0: -2147483648, F1: 1}
                                            _ = v29
                                        }
                                        v3[v1] = fi24
                                    }
                                    v27 = true
                                    v30 := GenS4{F0: 2}
                                    _ = v30
                                case "\n\t":
                                    v31 := GenS5{F0: -68, F1: uint16(v4)}
                                    _ = v31
                                case "hello":
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                default:
                                    fmt.Println(v4, v2, v4)
                                }
                            } else {
                                if v27 {
                                    var v32 uint64 = 2
                                    var v33 float32 = 3.14159
                                    if true {
                                        var v34 float64 = float64(((1e10 + v33) * 3.14159))
                                        var v35 int32 = 2147483647
                                        _ = v34
                                        _ = v35
                                    } else {
                                        if v27 {
                                            v36 := GenS6{F0: v2}
                                            switch v1 {
                                            case "a":
                                                _ = len(v3)
                                            default:
                                                switch (fi23 * fi24) {
                                                case 42:
                                                    fmt.Println(fi24, fi23)
                                                case -140:
                                                    v2 = (v4 * (v4 & v32))
                                                default:
                                                    if v27 {
                                                        switch v1 {
                                                        case "hello":
                                                            v37 := GenS3{F0: -1}
                                                            _ = v37
                                                        }
                                                        v4 += 0
                                                        var v38 int32 = -203
                                                        _ = v38
                                                    }
                                                }
                                            }
                                            var v39 uint16 = 255
                                            v40 := GenS3{F0: 42}
                                            switch (("" + v1) + ("a" + "")) {
                                            case "hello":
                                                if len(v20) > 0 {
                                                    v20[0] = (-(fi24))
                                                }
                                            case "a":
                                                v25 -= (v4 & v32)
                                            case "\n\t":
                                                v41 := v3[(v1 + "foo bar")]
                                                _ = v41
                                            }
                                            v42 := GenS6{F0: v25}
                                            _ = v36
                                            _ = v39
                                            _ = v40
                                            _ = v42
                                        } else {
                                            if len(v20) > 0 {
                                                v20[0] = fi23
                                            }
                                            if v27 {
                                                v33 -= float32(3.14159)
                                                fmt.Println(v1)
                                                var v43 uint16 = 1
                                                _ = v43
                                            }
                                        }
                                        v44 := GenS7{F0: 42, F1: "", F2: 363}
                                        fmt.Println(v25, v1)
                                        _ = v44
                                    }
                                    _ = v32
                                    _ = v33
                                }
                                v20 = append(v20, fi23)
                                switch v1 {
                                case "a":
                                    var v45 int16 = 1000
                                    _ = v45
                                case "foo bar":
                                    var v46 bool = false
                                    _ = v46
                                }
                            }
                            v25 -= 14
                            _ = v27
                        }
                    }
                    _ = fi24
                    _ = v25
                }
                _ = fi23
            }
            var v47 int32 = 1
            _ = v22
            _ = v47
        } else {
            var v48 int8 = 127
            for fi49 := 0; fi49 < 7; fi49++ {
                var v50 uint32 = 0
                v51 := GenS2{F0: -963, F1: 1}
                _ = fi49
                _ = v50
                _ = v51
            }
            _ = v48
        }
    }
    for fi52 := 0; fi52 < 7; fi52++ {
        switch ((v1 + v1) + "") {
        case "\n\t":
            if false {
                v1 = "\n\t"
                var v53 bool = true
                v54 := GenS4{F0: 42}
                _ = v53
                _ = v54
            }
        case "foo bar":
            v2 = v2
        default:
            v55 := GenS7{F0: 1, F1: ((v1 + (v1 + v1)) + v1), F2: 374}
            _ = v55
        }
        _ = fi52
    }
    var v56 int64 = -9223372036854775808
    if (-2.25 != 0.1) {
        v57 := GenS8{F0: 101, F1: -710, F2: 42}
        if (((false || (42 >= 0)) && ((!(float32(-2.25) >= 3.14159)) || true)) || (true && true)) {
            v58 := GenS9{F0: -9223372036854775808}
            _ = v58
        }
        var v59 int = -527
        _ = v57
        _ = v59
    }
    v3[v1] = -500
    for fi60 := 0; fi60 < 19; fi60++ {
        var v61 float32 = 0.0
        _ = fi60
        _ = v61
    }
    for fi62 := 0; fi62 < 12; fi62++ {
        switch -768 {
        case -9223372036854775808:
            for fi63 := 0; fi63 < 8; fi63++ {
                _ = len(v20)
                v64 := GenS10{F0: v56, F1: (fi62 ^ (-(fi62)))}
                v65 := GenS11{F0: 0}
                _ = fi63
                _ = v64
                _ = v65
            }
        case 42:
            v66 := GenS4{F0: 12}
            _ = v66
        case 514:
            var v67 uint64 = v2
            _ = v67
        default:
            v4 = (v4 % (((^(v2))) | 1))
        }
        switch ((v1 + v1) + "") {
        case "":
            for fi68 := 0; fi68 < 8; fi68++ {
                var v69 uint64 = v4
                _ = fi68
                _ = v69
            }
        case "a":
            v1 = v1
        case "foo bar":
            var v70 uint32 = 4294967295
            _ = v70
        default:
            var v71 float64 = -2.25
            _ = v71
        }
        if ((("hello" + v1) < (v1 + v1)) && true) {
            v20 = append(v20, fi62)
        }
        _ = fi62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v56
}
