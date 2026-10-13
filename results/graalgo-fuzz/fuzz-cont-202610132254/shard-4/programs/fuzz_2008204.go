package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 uint64; F1 uint; F2 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint64 }

func main() {
    var v1 int8 = 100
    v2 := make(map[int]uint32)
    var v3 uint16 = 0
    v4 := []int{-2147483648, int(v3)}
    var v5 int32 = 2147483647
    var v6 uint8 = 127
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: 1}
        _ = fi7
        _ = v8
    }
    if (1.5 == 3.14159) {
        v5 = 1
        v9 := GenS1{F0: 65535}
        if false {
            switch "foo bar" {
            case "foo bar":
                var v10 float64 = (1e10 / float64((0.1 * (1e10 * 0.1))))
                _ = v10
            case "\n\t":
                if len(v4) > 0 {
                    v4[0] = int(v5)
                }
            case "":
                switch "a" {
                case "\n\t":
                    var v11 float64 = float64(((((3.14159 / ((-2.25 / 1.0) / -2.25)) * -2.25) - 0.0) - ((1.5 - ((1e10 / 1.5) * float32(float64(0.1)))) - 0.1)))
                    _ = v11
                case "hello":
                    switch -2147483648 {
                    case 127:
                        if true {
                            for fi12 := 0; fi12 < 4; fi12++ {
                                v13 := GenS1{F0: 2}
                                _ = fi12
                                _ = v13
                            }
                        } else {
                            fmt.Println(v6, v1)
                            v14 := GenS1{F0: 42}
                            var v15 int64 = -9223372036854775808
                            _ = v14
                            _ = v15
                        }
                    case -1:
                        for fi16 := 0; fi16 < 6; fi16++ {
                            fmt.Println(v3)
                            var v17 uint32 = 697
                            _ = fi16
                            _ = v17
                        }
                    case 32767:
                        for fi18 := 0; fi18 < 19; fi18++ {
                            v19 := GenS1{F0: 445}
                            _ = fi18
                            _ = v19
                        }
                    default:
                        for fi20 := 0; fi20 < 10; fi20++ {
                            var v21 uint16 = v3
                            _ = fi20
                            _ = v21
                        }
                    }
                case "a":
                    if ((((uint32(v3) <= 2) && true) || true) && false) {
                        v6 += v6
                        var v22 int8 = v1
                        _ = v22
                    } else {
                        fmt.Println(v3, v3)
                    }
                default:
                    v1 += v1
                }
            default:
                v23 := GenS2{F0: 4294967295, F1: 0}
                _ = v23
            }
            for fi24 := 0; fi24 < 8; fi24++ {
                v25 := GenS1{F0: 189}
                v1 += v1
                _ = fi24
                _ = v25
            }
            fmt.Println(v6)
        } else {
            v5 += -985
            _ = len(v2)
            switch 833 {
            case 9223372036854775807:
                switch 1 {
                case 4294967295:
                    v26 := GenS2{F0: 0, F1: 0}
                    _ = v26
                case 32768:
                    v27 := GenS2{F0: 42, F1: 113}
                    _ = v27
                default:
                    fmt.Println(v1)
                }
            case 495:
                for fi28 := 0; fi28 < 15; fi28++ {
                    var v29 bool = ("" < "hello")
                    v5 -= 42
                    if (754 < 1000) {
                        if (v29 && ((v29 || false) || (v6 != 42))) {
                            v30 := GenS1{F0: uint(v6)}
                            _ = v30
                        }
                        fmt.Println(fi28, v1)
                        var v31 float32 = 1e10
                        _ = v31
                    }
                    _ = fi28
                    _ = v29
                }
            default:
                switch "foo bar" {
                case "hello":
                    v32 := GenS1{F0: 353}
                    _ = v32
                case "a":
                    v33 := GenS3{F0: v6, F1: (v1 & (v1 ^ -32))}
                    _ = v33
                case "foo bar":
                    if false {
                        v34 := GenS4{F0: v1}
                        for fi35 := 0; fi35 < 1; fi35++ {
                            switch (fi35 + (fi35 / (((fi35 * fi35)) | 1))) {
                            case -9223372036854775808:
                                if len(v4) > 0 {
                                    v4[0] = (fi35 / ((fi35) | 1))
                                }
                            case 128:
                                v36 := v2[fi35]
                                _ = v36
                            case 32768:
                                if (1 == 18446744073709551615) {
                                    var v37 int16 = 0
                                    v38 := GenS3{F0: 0, F1: (-(v1))}
                                    _ = v37
                                    _ = v38
                                } else {
                                    for fi39 := 0; fi39 < 3; fi39++ {
                                        fmt.Println(v3)
                                        switch "foo bar" {
                                        case "foo bar":
                                            v3 = v3
                                        case "a":
                                            switch ("foo bar" + "a") {
                                            case "\n\t":
                                                v40 := GenS5{F0: (v6 % ((174) | 1)), F1: (fi35 % ((fi39) | 1))}
                                                _ = v40
                                            default:
                                                _ = len(v4)
                                            }
                                        case "hello":
                                            var v41 int16 = 32767
                                            v2[int(v1)] = 2147483647
                                            _ = v41
                                        default:
                                            var v42 int8 = v1
                                            _ = v42
                                        }
                                        v43 := GenS2{F0: 528, F1: 42}
                                        _ = fi39
                                        _ = v43
                                    }
                                }
                            default:
                                v3 -= v3
                            }
                            v44 := GenS1{F0: 598}
                            _ = fi35
                            _ = v44
                        }
                        if (1.5 == 0.1) {
                            v6 += v6
                            for fi45 := 0; fi45 < 18; fi45++ {
                                _ = len(v2)
                                for fi46 := 0; fi46 < 15; fi46++ {
                                    switch "\n\t" {
                                    case "":
                                        v1 = int8(v5)
                                    default:
                                        v1 = 92
                                    }
                                    fmt.Println(v1, fi45, v6)
                                    _ = fi46
                                }
                                if true {
                                    v47 := GenS6{F0: 576, F1: 248, F2: "hello"}
                                    var v48 string = "\n\t"
                                    _ = v47
                                    _ = v48
                                } else {
                                    if (!false) {
                                        v6 += v6
                                        var v49 int64 = 834
                                        _ = v49
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    switch (fi45 | (^(fi45))) {
                                    case 0:
                                        v1 -= (v1 % ((v1) | 1))
                                    default:
                                        _ = len(v2)
                                    }
                                }
                                _ = fi45
                            }
                        }
                        _ = v34
                    }
                }
            }
        }
        _ = v9
    }
    fmt.Println(v5, v6, v5)
    fmt.Println(v1, v6)
    v2[-2147483648] = 588
    v50 := GenS7{F0: int64(v1)}
    v51 := GenS5{F0: 1, F1: -1}
    if true {
        fmt.Println(v1)
        for fi52 := 0; fi52 < 13; fi52++ {
            if true {
                fmt.Println(v1, v3)
            }
            if false {
                v53 := GenS8{F0: 0}
                _ = v53
            }
            _ = fi52
        }
    } else {
        v54 := GenS3{F0: v6, F1: v1}
        fmt.Println(v3, v6, v1)
        _ = v54
    }
    v55 := GenS8{F0: 2}
    if (!(("\n\t" + ("hello" + "hello")) != "foo bar")) {
        if true {
            var v56 uint = 256
            v57 := GenS7{F0: 293}
            v4 = append(v4, 10)
            _ = v56
            _ = v57
        } else {
            if true {
                var v58 int32 = 42
                _ = v58
            } else {
                fmt.Println(v1, v5)
            }
            switch ("\n\t" + "") {
            case "":
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v1)
                case "hello":
                    for fi59 := 0; fi59 < 19; fi59++ {
                        for fi60 := 0; fi60 < 6; fi60++ {
                            var v61 int16 = 32767
                            _ = fi60
                            _ = v61
                        }
                        if true {
                            v3 = v3
                            var v62 float32 = ((0.0 - -2.25) * float32(float64(0.1)))
                            var v63 int16 = 1
                            _ = v62
                            _ = v63
                        } else {
                            switch (fi59 | 1) {
                            case 42:
                                v3 -= (v3 << ((65535) & 63))
                            }
                        }
                        if (1 <= 0) {
                            var v64 float64 = (0.0 * ((1.5 - 0.1) + (0.0 / float64(((0.1 + 1.5) * -2.25)))))
                            _ = v64
                        } else {
                            v3 = v3
                            fmt.Println(v6)
                            v1 = 124
                        }
                        _ = fi59
                    }
                case "foo bar":
                    fmt.Println(v6, v1, v5)
                }
            }
        }
        fmt.Println(v5, v6, v6)
    } else {
        fmt.Println(v5, v6, v5)
    }
    v65 := GenS7{F0: int64(v5)}
    switch "" {
    case "\n\t":
        var v66 int32 = int32(v1)
        _ = v66
    default:
        v1 = -106
    }
    fmt.Println(v3, v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v55
    _ = v6
    _ = v65
}
