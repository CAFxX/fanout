package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint8; F1 int; F2 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 int32; F1 bool; F2 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -1
    v3 := make(map[string]int)
    v4 := []string{"", "", ("" + "")}
    var v5 uint32 = 2147483647
    if false {
        fmt.Println(v5, v2)
        switch "hello" {
        case "\n\t":
            for fi6 := 0; fi6 < 7; fi6++ {
                if false {
                    v2 += 42
                    fmt.Println(v2)
                } else {
                    var v7 uint64 = 826
                    for fi8 := 0; fi8 < 16; fi8++ {
                        if true {
                            v9 := GenS1{F0: 42, F1: "foo bar", F2: 4294967295}
                            _ = len(v4)
                            var v10 bool = (((v7 == 42) && false) || (255 != 42))
                            v11 := GenS1{F0: -128, F1: "hello", F2: v5}
                            var v12 int16 = 0
                            _ = v10
                            _ = v11
                            _ = v12
                            _ = v9
                        } else {
                            switch (("" + "foo bar") + "foo bar") {
                            case "":
                                fmt.Println(v7)
                            default:
                                if false {
                                    v13 := GenS1{F0: 0, F1: ((("hello" + "foo bar") + "foo bar") + ("hello" + "foo bar")), F2: (^(v5))}
                                    _ = v13
                                }
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        fmt.Println(fi8, v2)
                        _ = fi8
                    }
                    _ = v7
                }
                _ = len(v1)
                _ = fi6
            }
        case "a":
            v14 := GenS2{F0: 389}
            _ = v14
        default:
            for fi15 := 0; fi15 < 13; fi15++ {
                for fi16 := 0; fi16 < 8; fi16++ {
                    v5 = (v5 >> ((929) & 63))
                    if (!true) {
                        if (65535 > 65535) {
                            if true {
                                fmt.Println(fi16, v2, fi15)
                                v2 += v2
                                var v17 int16 = 42
                                _ = v17
                            } else {
                                v5 = (v5 >> ((826) & 63))
                                var v18 int = -1
                                var v19 int32 = v2
                                _ = v18
                                _ = v19
                            }
                            if (!true) {
                                v2 -= v2
                            } else {
                                v3["\n\t"] = (fi15 ^ fi16)
                                var v20 bool = true
                                v20 = v20
                                _ = v20
                            }
                            v21 := GenS2{F0: -1}
                            switch fi15 {
                            case 127:
                                fmt.Println(fi16, v2)
                            default:
                                v22 := GenS3{F0: -1, F1: 4294967295}
                                _ = v22
                            }
                            v2 += (-(v2))
                            _ = v21
                        }
                    } else {
                        fmt.Println(v5, v2)
                    }
                    _ = fi16
                }
                switch fi15 {
                case 9223372036854775807:
                    if (-111 >= 0) {
                        var v23 int32 = -402
                        _ = v23
                    } else {
                        if (!(!((!(255 > 42)) && true))) {
                            var v24 int = -275
                            var v25 uint32 = (^(v5))
                            var v26 uint8 = 255
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    }
                case 42:
                    for fi27 := 0; fi27 < 13; fi27++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(fi15, fi15)
                        _ = fi27
                    }
                case -608:
                    if false {
                        v28 := GenS2{F0: -32768}
                        fmt.Println(v5, v2, v5)
                        _ = v28
                    } else {
                        fmt.Println(fi15, v5, fi15)
                        v2 = (v2 - 42)
                        for fi29 := 0; fi29 < 9; fi29++ {
                            if len(v4) > 0 {
                                v4[0] = ("a" + "hello")
                            }
                            v30 := GenS4{F0: -647, F1: 127}
                            _ = fi29
                            _ = v30
                        }
                    }
                default:
                    if true {
                        var v31 int = (fi15 / (((fi15 | (fi15 << ((185) & 63)))) | 1))
                        var v32 int32 = (v2 * (^(v2)))
                        _ = v31
                        _ = v32
                    }
                }
                _ = fi15
            }
        }
        switch ("\n\t" + "\n\t") {
        case "hello":
            if ((float32(0.0) != 3.14159) || true) {
                switch -1 {
                case 181:
                    if (((true || false) || (true && false)) && (128 >= 1)) {
                        v5 += v5
                        for fi33 := 0; fi33 < 13; fi33++ {
                            if (-326 > 0) {
                                v2 -= (v2 | v2)
                                v2 = v2
                            } else {
                                for fi34 := 0; fi34 < 7; fi34++ {
                                    fmt.Println(fi33, v5)
                                    switch (fi33 << ((42) & 63)) {
                                    case -55:
                                        v5 -= v5
                                    }
                                    fmt.Println(v5, fi34, v2)
                                    _ = fi34
                                }
                                switch "foo bar" {
                                case "hello":
                                    v35 := GenS5{F0: 146, F1: -26, F2: (v5 % (((^(v5))) | 1))}
                                    _ = v35
                                default:
                                    v2 -= -1
                                }
                            }
                            _ = fi33
                        }
                        var v36 string = "foo bar"
                        _ = v36
                    } else {
                        for fi37 := 0; fi37 < 5; fi37++ {
                            v38 := GenS3{F0: 42, F1: 807}
                            var v39 uint8 = 54
                            v2 -= (-(v2))
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                        v4 = append(v4, "a")
                        fmt.Println(v5, v2)
                    }
                case -501:
                    if true {
                        _ = len(v4)
                        v40 := GenS6{F0: -1}
                        for fi41 := 0; fi41 < 16; fi41++ {
                            v42 := GenS7{F0: 0, F1: 42}
                            if false {
                                var v43 int16 = -857
                                if (1.5 < 0.1) {
                                    for fi44 := 0; fi44 < 18; fi44++ {
                                        var v45 string = ("hello" + "a")
                                        v1[(v5 ^ (v5 + 666))] = (fi41 * (fi44 >> ((408) & 63)))
                                        _ = fi44
                                        _ = v45
                                    }
                                }
                                _ = v43
                            } else {
                                v5 = 65535
                            }
                            if false {
                                var v46 int16 = 312
                                v46 += v46
                                _ = v46
                            }
                            _ = fi41
                            _ = v42
                        }
                        _ = v40
                    } else {
                        _ = len(v4)
                        for fi47 := 0; fi47 < 9; fi47++ {
                            var v48 uint8 = 88
                            switch fi47 {
                            case 2147483647:
                                _ = len(v4)
                            case 4294967295:
                                if (42 < 18446744073709551615) {
                                    v48 = v48
                                    v3[("foo bar" + "\n\t")] = fi47
                                }
                            case -9223372036854775808:
                                fmt.Println(v5)
                            }
                            v4 = append(v4, "hello")
                            _ = fi47
                            _ = v48
                        }
                        fmt.Println(v5, v5)
                    }
                }
                v5 = v5
                var v49 uint = 674
                _ = v49
            }
        default:
            v50 := v1[(^(v5))]
            _ = v50
        }
    } else {
        var v51 float32 = ((1.5 - (((3.14159 * 1.5) * -2.25) / -2.25)) - ((0.1 - 3.14159) * 1e10))
        var v52 int8 = 6
        _ = len(v4)
        _ = v51
        _ = v52
    }
    switch 452 {
    case 9223372036854775807:
        var v53 uint64 = 4294967295
        _ = v53
    default:
        var v54 float64 = 0.1
        _ = v54
    }
    if false {
        for fi55 := 0; fi55 < 8; fi55++ {
            _ = len(v4)
            v2 += v2
            var v56 int8 = 127
            _ = fi55
            _ = v56
        }
    } else {
        v5 = (^(v5))
    }
    v2 = v2
    _ = len(v1)
    v57 := v1[572]
    if true {
        switch "foo bar" {
        case "":
            v58 := GenS8{F0: v2, F1: true, F2: -314}
            _ = v58
        }
        fmt.Println(v5, v2, v2)
        var v59 int32 = v2
        _ = v59
    } else {
        for fi60 := 0; fi60 < 10; fi60++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v61 int16 = -32768
            _ = fi60
            _ = v61
        }
        var v62 uint = 18446744073709551615
        v57 = (-(v57))
        _ = v62
    }
    v2 += int32(v5)
    if (0.0 != 0.0) {
        for fi63 := 0; fi63 < 8; fi63++ {
            var v64 uint8 = 255
            _ = fi63
            _ = v64
        }
        var v65 bool = ((124 > 1) && (((v2 + -997) <= -310) && true))
        var v66 string = ""
        _ = v65
        _ = v66
    } else {
        switch "foo bar" {
        case "\n\t":
            var v67 uint = 1
            _ = v67
        case "foo bar":
            var v68 int8 = 127
            _ = v68
        default:
            _ = len(v3)
        }
        for fi69 := 0; fi69 < 13; fi69++ {
            v57 -= fi69
            _ = fi69
        }
    }
    var v70 int = v57
    v57 = (v70 - v70)
    v71 := GenS2{F0: 8}
    var v72 uint64 = 0
    var v73 float64 = 3.14159
    var v74 int = (v70 + (v57 << ((32) & 63)))
    v75 := GenS1{F0: -1, F1: "hello", F2: (v5 | v5)}
    var v76 int = v74
    var v77 uint8 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v70:", v70)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v70
    _ = v71
    _ = v72
    _ = v73
    _ = v74
    _ = v75
    _ = v76
    _ = v77
}
