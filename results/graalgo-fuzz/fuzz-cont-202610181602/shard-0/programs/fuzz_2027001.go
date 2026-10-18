package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint64; F1 uint; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS6 struct { F0 string; F1 uint; F2 int }
type GenS7 struct { F0 int16; F1 int; F2 uint32 }
func genFn0() uint16 {
    v6 := GenS1{F0: 1, F1: 18446744073709551615}
    var v7 uint8 = 0
    _ = v6
    _ = v7
    return 636
}
func genFn1(p0 uint64) uint32 {
    p0 = (p0 | p0)
    v8 := GenS2{F0: (^(p0)), F1: 18446744073709551615, F2: 297}
    var v9 int = -1
    _ = p0
    _ = v8
    _ = v9
    return 213
}

func main() {
    var v1 int32 = 941
    var v2 int64 = -1
    var v3 string = "foo bar"
    v4 := []uint32{}
    var v5 int32 = 2147483647
    if true {
        for fi10 := 0; fi10 < 5; fi10++ {
            v5 = 217
            v5 -= (v1 << ((65535) & 63))
            _ = fi10
        }
        var v11 int32 = v1
        _ = v11
    }
    var v12 uint16 = 742
    fmt.Println(v2, v2, v5)
    var v13 uint64 = 2
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 bool = false
            fmt.Println(v13)
            if v16 {
                var v17 float32 = float32(1e10)
                v3 = ("" + "foo bar")
                if v16 {
                    v18 := GenS3{F0: 42}
                    var v19 int64 = v2
                    _ = v18
                    _ = v19
                } else {
                    var v20 uint8 = 2
                    _ = v20
                }
                v4 = append(v4, 2147483648)
                if true {
                    v2 = -1
                }
                fmt.Println(v3, v2)
                _ = v17
            }
            v3 = v3
            v21 := GenS1{F0: (fi15 >> ((40) & 63)), F1: 772}
            v3 = ("hello" + (v3 + v3))
            v1 = 0
            _ = fi15
            _ = v16
            _ = v21
        }
        if false {
            v22 := GenS4{F0: v5, F1: 65535}
            fmt.Println(v12, v3, v2)
            var v23 int32 = v1
            _ = v22
            _ = v23
        }
        _ = fi14
    }
    fmt.Println(v3, v2, v5)
    for fi24 := 0; fi24 < 12; fi24++ {
        v1 -= v1
        for fi25 := 0; fi25 < 5; fi25++ {
            v2 -= v2
            switch fi25 {
            case 1:
                if true {
                    var v26 int16 = 0
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v27 string = (v3 + (((v3 + ("foo bar" + v3)) + v3) + ("foo bar" + "a")))
                    _ = v26
                    _ = v27
                }
            case 710:
                v5 -= (v1 + 0)
            case 933:
                v2 = v2
            default:
                fmt.Println(v13, v3)
            }
            if false {
                var v28 int16 = -1
                if ((v12 % ((genFn0()) | 1)) > v12) {
                    _ = len(v4)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                fmt.Println(fi25)
                _ = v28
            } else {
                v4 = append(v4, 2)
                var v29 int32 = -2147483648
                if true {
                    var v30 uint16 = v12
                    _ = v30
                }
                _ = v29
            }
            _ = fi25
        }
        var v31 string = (v3 + v3)
        _ = fi24
        _ = v31
    }
    switch v3 {
    case "":
        if (42 >= 1000) {
            if (3.14159 <= 1.5) {
                v5 = (v5 - v1)
                v5 += (v1 << ((18446744073709551615) & 63))
                var v32 uint64 = (v13 & v13)
                _ = v32
            }
            v2 += -945
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case "\n\t":
        v12 -= (v12 / (((v12 - (v12 + 963))) | 1))
    case "a":
        fmt.Println(v1)
    default:
        var v33 float64 = ((1e10 / (0.1 / 0.1)) * 0.1)
        _ = v33
    }
    if (!(v12 != v12)) {
        if true {
            switch 32768 {
            case 32767:
                var v34 uint = 795
                _ = v34
            case -9223372036854775808:
                var v35 int16 = 32767
                _ = v35
            case 9223372036854775807:
                v5 = (v5 >> ((931) & 63))
            }
            var v36 int8 = 77
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v36
        } else {
            if ((v12 | (v12 << ((65535) & 63))) <= 0) {
                var v37 int32 = (v1 >> ((595) & 63))
                _ = v37
            } else {
                var v38 uint16 = (v12 << ((0) & 63))
                _ = len(v4)
                _ = v38
            }
            var v39 uint32 = uint32(v12)
            switch 479 {
            case 127:
                var v40 bool = false
                _ = v40
            case 4294967295:
                for fi41 := 0; fi41 < 17; fi41++ {
                    for fi42 := 0; fi42 < 16; fi42++ {
                        if ((false && (4294967295 > uint(v13))) || true) {
                            if false {
                                v5 += (v1 / ((-4) | 1))
                            } else {
                                var v43 uint8 = 169
                                switch v3 {
                                case "hello":
                                    if false {
                                        _ = len(v4)
                                    }
                                default:
                                    fmt.Println(v1)
                                }
                                _ = v43
                            }
                            var v44 int8 = int8(v1)
                            _ = v44
                        } else {
                            var v45 int64 = v2
                            _ = v45
                        }
                        v46 := GenS5{F0: v39, F1: (v2 / ((v2) | 1)), F2: 128}
                        _ = fi42
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 10; fi47++ {
                        var v48 float64 = (3.14159 / 1e10)
                        _ = len(v4)
                        v49 := GenS5{F0: (v39 | (v39 & 42)), F1: (v2 >> ((2) & 63)), F2: 128}
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                    fmt.Println(v5, v2)
                    _ = fi41
                }
            case 234:
                var v50 int64 = v2
                _ = v50
            default:
                switch -1 {
                case -104:
                    if len(v4) > 0 {
                        v4[0] = 708
                    }
                case 4294967295:
                    if len(v4) > 0 {
                        v4[0] = v39
                    }
                }
            }
            _ = v39
        }
    } else {
        _ = len(v4)
    }
    fmt.Println(v3)
    var v51 string = v3
    switch 127 {
    case -92:
        fmt.Println(v3, v51)
    case -230:
        if (-9223372036854775808 <= 476) {
            v1 -= -1
            var v52 bool = (-1 >= 42)
            _ = len(v4)
            _ = v52
        }
    }
    v2 += v2
    fmt.Println(v2, v5)
    if true {
        var v53 uint8 = 1
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi54 := 0; fi54 < 20; fi54++ {
            if false {
                if ((!(!true)) && false) {
                    var v55 int32 = int32(v12)
                    switch fi54 {
                    case 464:
                        switch (("hello" + v51) + v3) {
                        case "foo bar":
                            v2 = -802
                        default:
                            fmt.Println(v2, v3, v13)
                        }
                    case 128:
                        var v56 bool = ("hello" == v3)
                        _ = v56
                    case -197:
                        var v57 uint = uint(v12)
                        _ = v57
                    default:
                        v55 = int32(fi54)
                    }
                    for fi58 := 0; fi58 < 10; fi58++ {
                        var v59 bool = (v12 != 0)
                        var v60 int32 = 25
                        fmt.Println(fi58, v5)
                        v51 = v51
                        var v61 float64 = (3.14159 * 3.14159)
                        _ = fi58
                        _ = v59
                        _ = v60
                        _ = v61
                    }
                    _ = v55
                } else {
                    if (!true) {
                        _ = len(v4)
                        if true {
                            v53 = uint8(v12)
                            if true {
                                v62 := GenS2{F0: 1, F1: 1, F2: -32768}
                                _ = v62
                            }
                            var v63 uint = 42
                            _ = v63
                        }
                    } else {
                        var v64 string = "foo bar"
                        switch 2147483647 {
                        case 4294967295:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case 32768:
                            fmt.Println(v2, v12)
                        default:
                            v53 = (v53 + v53)
                        }
                        v2 += 9223372036854775807
                        _ = v64
                    }
                }
                for fi65 := 0; fi65 < 20; fi65++ {
                    switch (((v51 + v51) + "\n\t") + "\n\t") {
                    case "a":
                        v51 = v3
                        var v66 uint16 = 1
                        _ = v66
                    case "foo bar":
                        v5 = 1
                    case "":
                        var v67 uint8 = (v53 & 128)
                        var v68 uint = 2
                        var v69 float32 = (1e10 * (((1.5 - float32(1.5)) - 0.0) / ((float32(float64(3.14159)) / (0.1 - 1.5)) * 1e10)))
                        _ = v67
                        _ = v68
                        _ = v69
                    }
                    _ = fi65
                }
            }
            var v70 uint = 718
            v12 = (v12 / (((v12 % ((v12) | 1))) | 1))
            _ = fi54
            _ = v70
        }
        _ = v53
    } else {
        _ = len(v4)
        for fi71 := 0; fi71 < 6; fi71++ {
            v72 := GenS2{F0: 42, F1: 56, F2: -1}
            var v73 string = "foo bar"
            if ((42 >= 4294967295) && (-823 < 32768)) {
                v74 := GenS6{F0: v3, F1: 0, F2: 32768}
                switch v3 {
                case "\n\t":
                    switch (v73 + (v51 + "hello")) {
                    case "foo bar":
                        var v75 int64 = v2
                        _ = v75
                    default:
                        for fi76 := 0; fi76 < 11; fi76++ {
                            v2 = (^(v2))
                            v73 = (("\n\t" + (v51 + (v51 + v73))) + v73)
                            v77 := GenS7{F0: 1000, F1: fi71, F2: 1}
                            v73 = (((v73 + "a") + v51) + "")
                            _ = fi76
                            _ = v77
                        }
                    }
                }
                for fi78 := 0; fi78 < 18; fi78++ {
                    v4 = append(v4, 649)
                    v4 = append(v4, genFn1((^(v13))))
                    _ = fi78
                }
                _ = v74
            } else {
                v13 = (v13 / (((v13 % ((2) | 1))) | 1))
            }
            _ = fi71
            _ = v72
            _ = v73
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
