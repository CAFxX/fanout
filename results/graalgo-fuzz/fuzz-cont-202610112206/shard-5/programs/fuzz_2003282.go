package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 string; F1 uint32; F2 bool }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS5 struct { F0 bool; F1 string; F2 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int64; F1 string }
func genFn0(p0 int32, p1 uint64, p2 string) int {
    fmt.Println(p0)
    v5 := []int{-9223372036854775808}
    switch p2 {
    case "hello":
        fmt.Println(p0, p1, p2)
    case "\n\t":
        var v6 int16 = -779
        _ = v6
    case "foo bar":
        var v7 uint16 = 32767
        _ = v7
    default:
        for fi8 := 0; fi8 < 20; fi8++ {
            switch fi8 {
            case -940:
                if (true && true) {
                    var v9 int32 = 42
                    switch p2 {
                    case "a":
                        switch ("a" + "a") {
                        case "":
                            p1 = p1
                        default:
                            for fi10 := 0; fi10 < 5; fi10++ {
                                v5 = append(v5, (^(fi10)))
                                _ = fi10
                            }
                        }
                    case "hello":
                        for fi11 := 0; fi11 < 2; fi11++ {
                            var v12 int32 = p0
                            switch "\n\t" {
                            case "\n\t":
                                fmt.Println(fi8)
                            case "hello":
                                var v13 float64 = float64(3.14159)
                                var v14 int32 = (^(v12))
                                if ((1000 > 1) && true) {
                                    if ("\n\t" > p2) {
                                        fmt.Println(v13, fi11)
                                        v9 += 1
                                    } else {
                                        fmt.Println(p2, v14)
                                        if true {
                                            v5 = append(v5, (fi11 >> ((144) & 63)))
                                        }
                                        var v15 uint16 = 1
                                        _ = v15
                                    }
                                    var v16 float32 = float32((0.1 / v13))
                                    var v17 float64 = v13
                                    _ = v16
                                    _ = v17
                                }
                                var v18 int = 32767
                                if false {
                                    v19 := GenS1{F0: 42, F1: 65535}
                                    v13 += 0.1
                                    v20 := GenS2{F0: (p2 + ("" + ((p2 + ((p2 + "") + "\n\t")) + ""))), F1: 2, F2: true}
                                    _ = v19
                                    _ = v20
                                }
                                _ = v13
                                _ = v14
                                _ = v18
                            default:
                                if len(v5) > 0 {
                                    v5[0] = fi8
                                }
                            }
                            _ = fi11
                            _ = v12
                        }
                    case "":
                        var v21 uint64 = p1
                        _ = v21
                    default:
                        if len(v5) > 0 {
                            v5[0] = fi8
                        }
                    }
                    switch ((("a" + ("hello" + "hello")) + (("a" + p2) + "\n\t")) + p2) {
                    case "a":
                        if false {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            fmt.Println(v9, p2, p2)
                        }
                    case "":
                        var v22 int32 = 0
                        _ = v22
                    }
                    _ = v9
                } else {
                    if (true && (-2.25 == float32(float64(float32((-2.25 / 1.0)))))) {
                        v23 := GenS3{F0: -50, F1: 66}
                        p0 -= (p0 | -1)
                        var v24 uint64 = (p1 / ((p1) | 1))
                        _ = v23
                        _ = v24
                    }
                    v25 := GenS4{F0: 293, F1: 291, F2: -1}
                    v26 := GenS3{F0: 100, F1: 42}
                    _ = v25
                    _ = v26
                }
            case 128:
                for fi27 := 0; fi27 < 4; fi27++ {
                    var v28 uint8 = 226
                    if false {
                        var v29 uint = 551
                        _ = v29
                    } else {
                        if (65535 < 32767) {
                            var v30 float32 = (0.1 - 0.1)
                            _ = v30
                        }
                        var v31 float64 = ((1.5 + 3.14159) + 1.5)
                        _ = v31
                    }
                    _ = fi27
                    _ = v28
                }
            }
            if (255 < 0) {
                v32 := GenS5{F0: true, F1: p2, F2: 1000}
                _ = v32
            } else {
                for fi33 := 0; fi33 < 3; fi33++ {
                    switch fi8 {
                    case -2147483648:
                        switch (p2 + ("foo bar" + p2)) {
                        case "\n\t":
                            if false {
                                var v34 int = fi8
                                if (130 < 2) {
                                    var v35 uint32 = 2147483647
                                    if true {
                                        v36 := GenS6{F0: -23}
                                        var v37 bool = false
                                        if v37 {
                                            switch ("a" + p2) {
                                            case "hello":
                                                var v38 float32 = (1e10 - 0.0)
                                                _ = v38
                                            case "foo bar":
                                                switch (v34 ^ fi33) {
                                                case -482:
                                                    var v39 int32 = p0
                                                    _ = v39
                                                case 32:
                                                    p0 -= (p0 >> ((2) & 63))
                                                case 452:
                                                    var v40 uint = 563
                                                    _ = v40
                                                }
                                            }
                                            switch p2 {
                                            case "hello":
                                                if v37 {
                                                    fmt.Println(v35, fi33, v34)
                                                    v5 = append(v5, v34)
                                                    v34 -= fi33
                                                }
                                            case "":
                                                fmt.Println(p0, v34, p0)
                                            }
                                            if v37 {
                                                var v41 int16 = 0
                                                v5 = append(v5, fi8)
                                                var v42 int64 = -1
                                                _ = v41
                                                _ = v42
                                            } else {
                                                p1 -= uint64(v35)
                                            }
                                        } else {
                                            var v43 int8 = 42
                                            fmt.Println(p0, v35)
                                            _ = v43
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v35 = (v35 >> ((213) & 63))
                                        _ = v36
                                        _ = v37
                                    }
                                    _ = v35
                                } else {
                                    switch -195 {
                                    case 127:
                                        if true {
                                            var v44 float64 = 0.1
                                            v45 := GenS4{F0: 124, F1: uint16(p1), F2: 42}
                                            v34 = 291
                                            _ = v44
                                            _ = v45
                                        }
                                    case 4294967295:
                                        var v46 uint64 = (p1 >> ((65535) & 63))
                                        _ = v46
                                    case -284:
                                        var v47 uint16 = 65535
                                        _ = v47
                                    default:
                                        p2 = "a"
                                    }
                                    var v48 uint32 = 2147483647
                                    _ = v48
                                }
                                _ = v34
                            }
                        }
                    }
                    v49 := GenS7{F0: 1}
                    p2 = ("" + p2)
                    p2 = "\n\t"
                    _ = fi33
                    _ = v49
                }
            }
            if true {
                var v50 string = "hello"
                fmt.Println(v50, p0)
                _ = v50
            } else {
                var v51 int32 = 284
                switch fi8 {
                case 32768:
                    if true {
                        if (!(!false)) {
                            p1 = 451
                            for fi52 := 0; fi52 < 16; fi52++ {
                                switch -383 {
                                case 32767:
                                    var v53 float32 = 1e10
                                    _ = v53
                                }
                                _ = fi52
                            }
                        } else {
                            v54 := GenS7{F0: 100}
                            for fi55 := 0; fi55 < 19; fi55++ {
                                _ = len(v5)
                                _ = fi55
                            }
                            fmt.Println(p2, p1, p0)
                            _ = v54
                        }
                        if (false && true) {
                            v5 = append(v5, (fi8 * (fi8 + fi8)))
                            v56 := GenS5{F0: false, F1: p2, F2: 1}
                            var v57 uint8 = 40
                            _ = v56
                            _ = v57
                        } else {
                            p2 = ""
                        }
                    } else {
                        if false {
                            var v58 float64 = -2.25
                            if (fi8 != fi8) {
                                _ = len(v5)
                                v51 -= (p0 * v51)
                                switch (fi8 - 32767) {
                                case 42:
                                    p2 = p2
                                default:
                                    if false {
                                        for fi59 := 0; fi59 < 9; fi59++ {
                                            p1 = p1
                                            fmt.Println(fi8, v51)
                                            v58 = float64((1.5 + 0.0))
                                            _ = fi59
                                        }
                                        p0 -= (p0 & v51)
                                    } else {
                                        if ((v58 == -2.25) && (true || true)) {
                                            v60 := GenS1{F0: -1, F1: 42}
                                            _ = v60
                                        } else {
                                            p2 = p2
                                            fmt.Println(p0)
                                        }
                                    }
                                }
                            } else {
                                v61 := GenS1{F0: -128, F1: 211}
                                switch p2 {
                                case "foo bar":
                                    fmt.Println(v51, v58)
                                case "a":
                                    var v62 int64 = 9223372036854775807
                                    _ = v62
                                default:
                                    if (!false) {
                                        var v63 int32 = v51
                                        var v64 int32 = (v51 ^ p0)
                                        _ = v63
                                        _ = v64
                                    }
                                }
                                v65 := GenS5{F0: (false && true), F1: "", F2: 563}
                                _ = v61
                                _ = v65
                            }
                            _ = v58
                        }
                        var v66 uint8 = 2
                        var v67 uint64 = 569
                        _ = v66
                        _ = v67
                    }
                case 128:
                    var v68 int32 = -1
                    _ = v68
                default:
                    var v69 uint64 = (p1 << ((2) & 63))
                    _ = v69
                }
                if len(v5) > 0 {
                    v5[0] = 472
                }
                _ = v51
            }
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}
func genFn1(p0 int) uint16 {
    var v70 int16 = 1000
    if ((v70 & int16(p0)) < (v70 / ((v70) | 1))) {
        var v71 int16 = v70
        v70 -= (v71 / ((61) | 1))
        _ = v71
    }
    _ = p0
    _ = v70
    return 375
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]int)
    var v3 int32 = -825
    var v4 uint32 = 2
    v2[2] = 128
    var v72 int32 = v3
    var v73 uint = 4294967295
    var v74 uint = v73
    for fi75 := 0; fi75 < 11; fi75++ {
        v76 := GenS8{F0: (fi75 ^ -9223372036854775808)}
        for fi77 := 0; fi77 < 14; fi77++ {
            switch "" {
            case "\n\t":
                if false {
                    v78 := []int{(-(fi75))}
                    v79 := GenS4{F0: 0, F1: 32767, F2: -562}
                    if len(v78) >= 2 {
                        v78 = v78[1:len(v78)]
                    }
                    _ = v78
                    _ = v79
                } else {
                    fmt.Println(v4, v1)
                }
            case "a":
                if true {
                    fmt.Println(v3)
                    v80 := GenS3{F0: -1, F1: 255}
                    switch "" {
                    case "hello":
                        if true {
                            v81 := GenS9{F0: v73}
                            _ = v81
                        }
                    case "":
                        v82 := v2[2147483647]
                        _ = v82
                    }
                    _ = v80
                }
            }
            v83 := GenS9{F0: v73}
            _ = fi77
            _ = v83
        }
        _ = fi75
        _ = v76
    }
    v84 := GenS6{F0: 100}
    fmt.Println(v72, v72)
    fmt.Println(v4, v74)
    v85 := GenS10{F0: -198, F1: "a"}
    for fi86 := 0; fi86 < 5; fi86++ {
        var v87 int = fi86
        var v88 uint8 = 127
        v72 -= v72
        _ = fi86
        _ = v87
        _ = v88
    }
    if false {
        var v89 uint16 = 177
        var v90 uint32 = (v4 / ((v4) | 1))
        _ = len(v2)
        _ = v89
        _ = v90
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v72
    _ = v73
    _ = v74
    _ = v84
    _ = v85
}
