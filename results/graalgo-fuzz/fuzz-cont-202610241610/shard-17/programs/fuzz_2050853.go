package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 string; F2 uint64 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 bool; F1 bool; F2 int8 }
type GenS7 struct { F0 int64; F1 uint8 }
type GenS8 struct { F0 uint16; F1 string }
type GenS9 struct { F0 uint8; F1 uint64 }
type GenS10 struct { F0 int16; F1 bool; F2 uint16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int32; F1 uint8 }
type GenS13 struct { F0 string; F1 int32 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{42, 1, 32767}
    v3 := []uint32{1}
    var v4 bool = (42 < 2147483648)
    var v5 float32 = 1e10
    v6 := []uint32{1, 2, 777, 65535}
    v5 += (v5 / v5)
    if ((false && true) && (0 < 375)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        if (657 <= 729) {
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 float64 = 1e10
                _ = fi7
                _ = v8
            }
            if v4 {
                if len(v3) > 0 {
                    v3[0] = 65535
                }
            }
            if v4 {
                v1 += v1
            } else {
                fmt.Println(v5, v4, v5)
            }
        }
        if v4 {
            v1 -= v1
            fmt.Println(v1)
            v9 := GenS1{F0: v1, F1: 42}
            _ = v9
        }
        v1 = v1
    }
    v10 := GenS1{F0: (-(v1)), F1: 100}
    switch -2147483648 {
    case -843:
        if false {
            v1 -= (v1 << ((255) & 63))
            fmt.Println(v5)
            v1 -= (v1 & (-(v1)))
        } else {
            v11 := GenS2{F0: 493, F1: 18446744073709551615}
            for fi12 := 0; fi12 < 14; fi12++ {
                v13 := GenS3{F0: (-(v1))}
                v1 = v1
                v14 := GenS4{F0: ("hello" + ""), F1: ("\n\t" + ("a" + "hello")), F2: 4294967295}
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        }
    case 908:
        if (0 <= 2) {
            fmt.Println(v5)
            switch 0 {
            case -9223372036854775808:
                v5 -= 1.5
            case 42:
                v15 := GenS2{F0: 1, F1: 42}
                _ = v15
            case 9223372036854775807:
                var v16 string = ""
                _ = v16
            }
            v17 := GenS5{F0: 4294967295, F1: 42}
            _ = v17
        } else {
            v3 = append(v3, 344)
            for fi18 := 0; fi18 < 20; fi18++ {
                v4 = (-902 != v1)
                _ = fi18
            }
        }
    }
    if (4294967295 > 65535) {
        var v19 bool = false
        var v20 int64 = v1
        var v21 uint16 = 32767
        _ = v19
        _ = v20
        _ = v21
    } else {
        if len(v6) > 0 {
            v6[0] = 825
        }
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        if (0 <= 127) {
            _ = len(v6)
        }
        switch (fi22 << ((65535) & 63)) {
        case 0:
            if false {
                _ = len(v6)
                v23 := GenS6{F0: false, F1: (v4 || v4), F2: 127}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v23
            }
        case 127:
            v5 -= 3.14159
        default:
            if len(v6) > 0 {
                v6[0] = 2147483647
            }
        }
        _ = fi22
    }
    var v24 uint = 863
    v25 := GenS2{F0: (v24 - (v24 % ((v24) | 1))), F1: v24}
    v26 := GenS7{F0: (v1 << ((42) & 63)), F1: 127}
    for fi27 := 0; fi27 < 3; fi27++ {
        if (84 != -128) {
            switch "a" {
            case "hello":
                for fi28 := 0; fi28 < 13; fi28++ {
                    switch "\n\t" {
                    case "":
                        v2 = append(v2, (fi27 % ((-1) | 1)))
                    default:
                        var v29 uint32 = uint32(v24)
                        _ = v29
                    }
                    _ = fi28
                }
            default:
                if true {
                    v5 -= v5
                    var v30 uint8 = 49
                    for fi31 := 0; fi31 < 13; fi31++ {
                        if (!v4) {
                            v24 = 65535
                            switch (-(fi27)) {
                            case 42:
                                v32 := GenS3{F0: (-(v1))}
                                _ = v32
                            case -362:
                                fmt.Println(v5, v24, fi27)
                            case -9223372036854775808:
                                v1 = (v1 / ((v1) | 1))
                                var v33 int = fi31
                                _ = v33
                            }
                            if v4 {
                                if ((42 >= (v30 | (v30 | 239))) && v4) {
                                    v34 := GenS8{F0: 479, F1: "foo bar"}
                                    v35 := GenS9{F0: v30, F1: 42}
                                    _ = v34
                                    _ = v35
                                }
                                switch (-(fi27)) {
                                case 0:
                                    _ = len(v2)
                                case 4294967295:
                                    v36 := GenS7{F0: 701, F1: (v30 / ((13) | 1))}
                                    _ = v36
                                default:
                                    v37 := GenS8{F0: 176, F1: ("hello" + "foo bar")}
                                    _ = v37
                                }
                                v30 -= 0
                                fmt.Println(v1)
                            } else {
                                v30 = (v30 ^ 127)
                                v38 := GenS10{F0: 418, F1: false, F2: 65535}
                                if (true || true) {
                                    var v39 uint8 = (v30 * 0)
                                    var v40 float32 = (0.1 - float32((0.1 + 3.14159)))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v39
                                    _ = v40
                                } else {
                                    switch fi27 {
                                    case 88:
                                        switch "" {
                                        case "hello":
                                            v5 = 0.1
                                            if (!false) {
                                                fmt.Println(v4, v4)
                                                if (32767 > 1000) {
                                                    fmt.Println(v24)
                                                    v41 := GenS8{F0: 42, F1: "a"}
                                                    _ = len(v3)
                                                    _ = v41
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                v6 = append(v6, 0)
                                            }
                                            v42 := GenS5{F0: 2147483647, F1: -78}
                                            _ = v42
                                        }
                                    case 0:
                                        var v43 float32 = v5
                                        _ = v43
                                    case -53:
                                        if (v30 <= v30) {
                                            v24 += (v24 | uint(v30))
                                            v44 := GenS9{F0: (v30 | 128), F1: 42}
                                            _ = v44
                                        } else {
                                            if false {
                                                fmt.Println(v30, v30, v30)
                                                if (!v4) {
                                                    v24 = v24
                                                    _ = len(v2)
                                                }
                                            } else {
                                                _ = len(v6)
                                            }
                                            switch -1 {
                                            case 9223372036854775807:
                                                v4 = v4
                                                v45 := GenS11{F0: 0}
                                                _ = v45
                                            case -589:
                                                v30 = 6
                                            default:
                                                if (0 != (^(v24))) {
                                                    var v46 int32 = 0
                                                    _ = v46
                                                }
                                            }
                                            v47 := GenS12{F0: 1, F1: (v30 >> ((492) & 63))}
                                            _ = v47
                                        }
                                    }
                                }
                                _ = v38
                            }
                            var v48 int = (fi31 ^ (^(fi27)))
                            var v49 float32 = ((v5 * float32(-2.25)) / (1.5 + (float32(1e10) * (v5 * (v5 / -2.25)))))
                            v6 = append(v6, 1)
                            _ = v48
                            _ = v49
                        } else {
                            var v50 int = fi31
                            _ = v50
                        }
                        if v4 {
                            v30 = (v30 + v30)
                            v3 = append(v3, 527)
                            switch 2147483647 {
                            case -2147483648:
                                var v51 int = fi27
                                fmt.Println(v5)
                                _ = v51
                            }
                        } else {
                            fmt.Println(v30, v30, fi27)
                            var v52 uint64 = 0
                            _ = v52
                        }
                        _ = fi31
                    }
                    _ = v30
                } else {
                    switch (("a" + "hello") + ("\n\t" + (("a" + ("hello" + (("foo bar" + "a") + "\n\t"))) + "\n\t"))) {
                    case "hello":
                        var v53 uint64 = 1
                        _ = v53
                    default:
                        v3 = append(v3, 65535)
                    }
                    var v54 float32 = (3.14159 - v5)
                    if (!false) {
                        for fi55 := 0; fi55 < 2; fi55++ {
                            if (26 > 100) {
                                _ = len(v3)
                                v54 -= v5
                            } else {
                                var v56 int16 = -264
                                v1 += v1
                                _ = v56
                            }
                            _ = fi55
                        }
                        switch -1 {
                        case -1:
                            for fi57 := 0; fi57 < 9; fi57++ {
                                v5 -= float32((0.1 + float64(v5)))
                                v4 = (1 < 1)
                                _ = fi57
                            }
                        case -203:
                            for fi58 := 0; fi58 < 10; fi58++ {
                                v59 := GenS4{F0: "hello", F1: ("foo bar" + "a"), F2: 1}
                                var v60 uint = v24
                                v61 := GenS13{F0: "a", F1: 2147483647}
                                _ = fi58
                                _ = v59
                                _ = v60
                                _ = v61
                            }
                        case 128:
                            v24 -= (v24 | v24)
                        }
                    }
                    _ = v54
                }
            }
            v62 := GenS2{F0: 204, F1: 1}
            _ = v62
        }
        switch -9223372036854775808 {
        case 699:
            if len(v3) > 0 {
                v3[0] = 2
            }
        case 128:
            for fi63 := 0; fi63 < 5; fi63++ {
                v4 = true
                _ = fi63
            }
        default:
            v1 = -1
        }
        var v64 uint8 = uint8(v24)
        _ = fi27
        _ = v64
    }
    if v4 {
        var v65 string = "a"
        if len(v2) > 0 {
            v2[0] = 932
        }
        _ = v65
    } else {
        fmt.Println(v5)
    }
    for fi66 := 0; fi66 < 14; fi66++ {
        var v67 bool = v4
        switch (fi66 << ((v24) & 63)) {
        case -1:
            for fi68 := 0; fi68 < 1; fi68++ {
                _ = len(v2)
                _ = fi68
            }
        default:
            for fi69 := 0; fi69 < 17; fi69++ {
                var v70 string = "\n\t"
                _ = fi69
                _ = v70
            }
        }
        for fi71 := 0; fi71 < 4; fi71++ {
            var v72 int8 = -1
            _ = fi71
            _ = v72
        }
        _ = fi66
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
