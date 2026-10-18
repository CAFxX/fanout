package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int16; F1 uint32; F2 int }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 string; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{120, 2147483648}
    var v3 int64 = -1
    if (!false) {
        switch ("foo bar" + "foo bar") {
        case "hello":
            if (false || false) {
                v4 := v1[1]
                _ = v4
            } else {
                for fi5 := 0; fi5 < 3; fi5++ {
                    for fi6 := 0; fi6 < 18; fi6++ {
                        v3 -= (v3 - v3)
                        v7 := GenS1{F0: 9223372036854775807, F1: 2147483647}
                        var v8 uint32 = 1
                        _ = fi6
                        _ = v7
                        _ = v8
                    }
                    if true {
                        v3 += (v3 ^ (v3 - v3))
                    }
                    _ = fi5
                }
                var v9 int16 = -32768
                _ = v9
            }
        case "\n\t":
            v3 = v3
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int = (-(fi10))
            switch ("hello" + "") {
            case "hello":
                v12 := GenS2{F0: -32768, F1: 42, F2: fi10}
                _ = v12
            case "\n\t":
                switch "" {
                case "":
                    if (3.14159 > (((3.14159 * (0.1 * 1e10)) - ((((3.14159 + ((1.5 + 3.14159) / (1.5 - -2.25))) * 3.14159) * 3.14159) + 3.14159)) / 0.1)) {
                        for fi13 := 0; fi13 < 10; fi13++ {
                            var v14 uint = 4294967295
                            if (-1 <= (fi10 / ((fi13) | 1))) {
                                var v15 int = fi13
                                v16 := GenS2{F0: 1000, F1: 65535, F2: fi10}
                                var v17 uint32 = 42
                                _ = v15
                                _ = v16
                                _ = v17
                            } else {
                                v14 = (^(v14))
                            }
                            _ = fi13
                            _ = v14
                        }
                        v3 = (-(v3))
                    } else {
                        v18 := GenS2{F0: 1000, F1: 42, F2: 128}
                        v19 := GenS2{F0: 0, F1: 866, F2: (-(fi10))}
                        _ = v18
                        _ = v19
                    }
                case "foo bar":
                    _ = len(v2)
                case "a":
                    if true {
                        v11 -= (-(v11))
                        v3 = 0
                        var v20 bool = true
                        _ = v20
                    } else {
                        fmt.Println(v3, fi10)
                        for fi21 := 0; fi21 < 8; fi21++ {
                            v22 := GenS1{F0: int64(fi21), F1: 1}
                            v3 -= (v3 ^ -1)
                            v23 := GenS3{F0: 100, F1: (v3 % ((v3) | 1))}
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                    }
                }
            }
            fmt.Println(v11, v11, v3)
            _ = fi10
            _ = v11
        }
        if false {
            switch 634 {
            case 662:
                v3 += (v3 % ((0) | 1))
            case 127:
                v3 = v3
            default:
                for fi24 := 0; fi24 < 19; fi24++ {
                    for fi25 := 0; fi25 < 13; fi25++ {
                        switch fi24 {
                        case 506:
                            var v26 uint64 = 67
                            _ = v26
                        case -9223372036854775808:
                            v27 := GenS4{F0: "a", F1: 2147483648, F2: 209}
                            delete(v1, 393)
                            _ = v27
                        default:
                            v3 += v3
                        }
                        _ = fi25
                    }
                    _ = fi24
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    } else {
        var v28 int = 127
        _ = v28
    }
    if (!(!true)) {
        for fi29 := 0; fi29 < 20; fi29++ {
            v3 -= v3
            for fi30 := 0; fi30 < 14; fi30++ {
                _ = len(v1)
                _ = fi30
            }
            var v31 uint16 = 42
            _ = fi29
            _ = v31
        }
        if (-201 > 1) {
            var v32 bool = (true || (false && ((!((true && ((v3 % ((42) | 1)) < v3)) && false)) && true)))
            _ = v32
        } else {
            fmt.Println(v3)
            v33 := GenS1{F0: v3, F1: -2147483648}
            if false {
                v3 -= (v3 | (v3 % ((v3) | 1)))
                var v34 int16 = 263
                fmt.Println(v3, v3)
                _ = v34
            } else {
                v35 := GenS3{F0: -104, F1: 0}
                var v36 bool = false
                _ = v35
                _ = v36
            }
            _ = v33
        }
        for fi37 := 0; fi37 < 2; fi37++ {
            v38 := v1[119]
            switch "a" {
            case "":
                for fi39 := 0; fi39 < 4; fi39++ {
                    var v40 int64 = (v3 << ((42) & 63))
                    _ = fi39
                    _ = v40
                }
            case "hello":
                fmt.Println(fi37, v38)
            case "\n\t":
                for fi41 := 0; fi41 < 5; fi41++ {
                    v42 := GenS5{F0: 18446744073709551615, F1: 896}
                    v3 = v3
                    v43 := GenS1{F0: v3, F1: 2147483647}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            default:
                var v44 uint16 = 392
                _ = v44
            }
            switch (fi37 + -679) {
            case 128:
                v3 -= (v3 * v3)
            case 4294967295:
                if false {
                    switch "\n\t" {
                    case "hello":
                        for fi45 := 0; fi45 < 2; fi45++ {
                            switch fi37 {
                            case -523:
                                v3 = v3
                            case 535:
                                v3 += v3
                            }
                            if ((("a" + ("\n\t" + "hello")) + "foo bar") <= "\n\t") {
                                v3 = v3
                                _ = len(v1)
                            } else {
                                if (517 >= 32767) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                } else {
                                    v38 = v38
                                }
                            }
                            var v46 float64 = 1.5
                            switch fi37 {
                            case 617:
                                var v47 uint16 = 32767
                                _ = v47
                            }
                            v48 := GenS4{F0: "", F1: 2147483647, F2: 1}
                            var v49 int64 = v3
                            _ = fi45
                            _ = v46
                            _ = v48
                            _ = v49
                        }
                    }
                    if v38 {
                        var v50 int = fi37
                        if len(v2) > 0 {
                            v2[0] = 205
                        }
                        fmt.Println(fi37)
                        _ = v50
                    } else {
                        for fi51 := 0; fi51 < 16; fi51++ {
                            v52 := GenS6{F0: (("hello" + "\n\t") + "\n\t")}
                            _ = fi51
                            _ = v52
                        }
                    }
                    v3 = 9223372036854775807
                } else {
                    var v53 uint64 = 1
                    if v38 {
                        var v54 float64 = (float64(0.1) + 0.1)
                        _ = v54
                    } else {
                        if false {
                            v55 := GenS7{F0: (((("" + "") + "foo bar") + "hello") + ("foo bar" + ""))}
                            v3 = v3
                            _ = v55
                        } else {
                            var v56 uint8 = 236
                            var v57 uint = 140
                            var v58 uint8 = 163
                            _ = v56
                            _ = v57
                            _ = v58
                        }
                    }
                    for fi59 := 0; fi59 < 17; fi59++ {
                        switch (fi37 ^ (fi59 & fi37)) {
                        case 0:
                            switch (("a" + "hello") + ("" + "foo bar")) {
                            case "":
                                fmt.Println(fi59, fi37)
                            default:
                                v60 := v1[65535]
                                _ = v60
                            }
                        case 239:
                            var v61 int64 = v3
                            _ = v61
                        case 605:
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                        }
                        fmt.Println(v3)
                        v53 += (v53 / ((766) | 1))
                        var v62 float32 = 1.5
                        _ = len(v2)
                        var v63 int16 = 42
                        _ = fi59
                        _ = v62
                        _ = v63
                    }
                    _ = v53
                }
            }
            _ = fi37
            _ = v38
        }
    } else {
        _ = len(v2)
        var v64 int32 = 1
        _ = v64
    }
    v3 += (v3 + (^(v3)))
    if (!(!((!((false || true) || true)) && true))) {
        v3 = v3
        v3 = (^(v3))
    } else {
        v2 = append(v2, 1)
        v65 := GenS7{F0: ""}
        _ = v65
    }
    var v66 int64 = v3
    v3 = (v3 ^ v3)
    fmt.Println(v3, v3)
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v66
}
