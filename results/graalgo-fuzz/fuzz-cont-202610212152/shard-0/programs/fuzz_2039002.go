package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int8; F1 int16; F2 int64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{42, 4294967295, 662, 4294967295}
    v3 := make(map[int]int)
    v4 := []int{791, 128}
    v5 := []uint32{42, 4294967295, 1}
    var v6 uint64 = 814
    v7 := GenS1{F0: "foo bar", F1: 173, F2: 87}
    v6 = 18446744073709551615
    for fi8 := 0; fi8 < 9; fi8++ {
        v6 -= (v6 + 42)
        v6 -= 389
        _ = fi8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(v6, fi9)
        var v10 int8 = -1
        fmt.Println(v10, fi9, v10)
        _ = fi9
        _ = v10
    }
    var v11 bool = false
    switch ("foo bar" + ("hello" + ("" + (("a" + ("" + "")) + (("a" + "a") + ""))))) {
    case "\n\t":
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 bool = v11
            for fi14 := 0; fi14 < 9; fi14++ {
                fmt.Println(v11, v11)
                v15 := GenS1{F0: ("hello" + ""), F1: 42, F2: 1}
                v11 = (0 != 1)
                fmt.Println(fi12, fi14, v6)
                if (((255 == 2) && v13) && true) {
                    v16 := GenS2{F0: -53, F1: -88, F2: 42}
                    if ((3.14159 == 0.1) || v13) {
                        if (!(!(4294967295 <= 42))) {
                            v4 = append(v4, fi12)
                            if (v6 >= v6) {
                                fmt.Println(v6)
                            } else {
                                v17 := GenS1{F0: "a", F1: 11, F2: 634}
                                _ = v17
                            }
                            fmt.Println(fi14, v6)
                            v18 := GenS2{F0: 1, F1: 32767, F2: int64(fi14)}
                            _ = v18
                        } else {
                            v6 += v6
                            var v19 uint8 = 128
                            v19 = (v19 * v19)
                            _ = v19
                        }
                        if len(v4) > 0 {
                            v4[0] = (fi14 ^ -165)
                        }
                        switch (fi14 << ((18446744073709551615) & 63)) {
                        case -578:
                            v6 -= 2
                        case -9223372036854775808:
                            v2 = append(v2, 4294967295)
                        case 427:
                            var v20 int = fi12
                            _ = v20
                        default:
                            v5 = append(v5, 105)
                        }
                    }
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
            var v21 int16 = -1
            _ = fi12
            _ = v13
            _ = v21
        }
    case "hello":
        v6 = (^(v6))
    case "a":
        switch "hello" {
        case "\n\t":
            fmt.Println(v11, v11)
        default:
            delete(v1, 42)
        }
    default:
        fmt.Println(v6, v6)
    }
    var v22 int8 = 0
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 float64 = 3.14159
        _ = fi23
        _ = v24
    }
    switch "hello" {
    case "\n\t":
        var v25 float64 = 0.0
        _ = v25
    case "a":
        var v26 float32 = -2.25
        _ = v26
    default:
        if (v11 || (v11 && (!(!true)))) {
            if (v11 || v11) {
                switch 32767 {
                case -2147483648:
                    if v11 {
                        v1[2147483648] = "hello"
                    }
                case -9223372036854775808:
                    v27 := v1[904]
                    _ = v27
                case 128:
                    var v28 int16 = 388
                    _ = v28
                default:
                    var v29 int = 2147483647
                    _ = v29
                }
                var v30 float64 = (3.14159 - float64(0.1))
                v31 := GenS1{F0: ("a" + "a"), F1: 128, F2: 65535}
                _ = v30
                _ = v31
            } else {
                var v32 uint64 = (^(v6))
                _ = v32
            }
            for fi33 := 0; fi33 < 10; fi33++ {
                v6 -= (^(v6))
                _ = fi33
            }
        } else {
            v11 = (0 > 127)
            fmt.Println(v6)
            v34 := GenS2{F0: v22, F1: -435, F2: 42}
            _ = v34
        }
    }
    var v35 int16 = 663
    var v36 uint = 42
    switch -847 {
    case 1:
        if v11 {
            var v37 uint32 = 230
            _ = len(v2)
            _ = v37
        }
    case 127:
        if false {
            v36 -= (v36 << ((237) & 63))
            var v38 string = "a"
            switch 557 {
            case 115:
                _ = len(v5)
            case 988:
                for fi39 := 0; fi39 < 18; fi39++ {
                    if false {
                        switch "hello" {
                        case "":
                            fmt.Println(v11, fi39, v11)
                        case "\n\t":
                            for fi40 := 0; fi40 < 4; fi40++ {
                                if false {
                                    v38 = v38
                                    v6 = 4294967295
                                    fmt.Println(v11, v36)
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = fi40
                                    }
                                    fmt.Println(v36, fi40, v6)
                                    delete(v3, 4294967295)
                                }
                                _ = fi40
                            }
                        }
                        var v41 uint32 = 2147483647
                        if (!v11) {
                            var v42 int16 = (v35 ^ -599)
                            if (v36 >= 361) {
                                var v43 bool = (!v11)
                                for fi44 := 0; fi44 < 12; fi44++ {
                                    var v45 float64 = (float64(1.5) * (1.5 * 0.0))
                                    v35 = 875
                                    var v46 string = "a"
                                    var v47 int8 = -81
                                    v48 := GenS1{F0: ((v38 + "\n\t") + ((v46 + v38) + "\n\t")), F1: 128, F2: 255}
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                }
                                var v49 uint8 = 0
                                _ = v43
                                _ = v49
                            } else {
                                v50 := GenS2{F0: v22, F1: (^(v42)), F2: 42}
                                fmt.Println(v42, v35, v6)
                                var v51 uint16 = 830
                                _ = v50
                                _ = v51
                            }
                            _ = v42
                        } else {
                            v52 := v1[v41]
                            v35 = -32768
                            _ = v52
                        }
                        _ = v41
                    } else {
                        if ((true || false) && v11) {
                            v2 = append(v2, uint32(v36))
                            var v53 float32 = (-2.25 - (float32(1e10) / 1e10))
                            _ = v53
                        } else {
                            for fi54 := 0; fi54 < 12; fi54++ {
                                var v55 int = (fi39 | int(v6))
                                v56 := GenS2{F0: int8(v36), F1: (^(v35)), F2: 0}
                                var v57 uint64 = v6
                                _ = fi54
                                _ = v55
                                _ = v56
                                _ = v57
                            }
                            _ = len(v3)
                        }
                    }
                    var v58 int64 = -800
                    v35 += v35
                    _ = fi39
                    _ = v58
                }
            case 9223372036854775807:
                v38 = ((v38 + v38) + v38)
            default:
                v38 = v38
            }
            _ = v38
        } else {
            var v59 uint32 = 0
            var v60 uint32 = 42
            _ = v59
            _ = v60
        }
    }
    if true {
        fmt.Println(v36, v6)
        _ = len(v3)
        if v11 {
            switch 4294967295 {
            case 0:
                if ((v35 % ((1) | 1)) > v35) {
                    if true {
                        v1[353] = "a"
                        switch "a" {
                        case "a":
                            fmt.Println(v35, v35, v6)
                        case "hello":
                            var v61 string = ("a" + "")
                            _ = v61
                        }
                        var v62 int8 = -59
                        _ = v62
                    } else {
                        var v63 int16 = 32767
                        for fi64 := 0; fi64 < 3; fi64++ {
                            switch "\n\t" {
                            case "hello":
                                v36 = (v36 % ((685) | 1))
                            case "\n\t":
                                for fi65 := 0; fi65 < 2; fi65++ {
                                    var v66 int16 = (^(v63))
                                    _ = fi65
                                    _ = v66
                                }
                            case "foo bar":
                                for fi67 := 0; fi67 < 15; fi67++ {
                                    v1[42] = "\n\t"
                                    _ = fi67
                                }
                            default:
                                v35 -= v35
                            }
                            v22 -= (v22 % ((v22) | 1))
                            v68 := v3[fi64]
                            _ = fi64
                            _ = v68
                        }
                        _ = v63
                    }
                    switch "foo bar" {
                    case "\n\t":
                        switch "foo bar" {
                        case "":
                            for fi69 := 0; fi69 < 2; fi69++ {
                                fmt.Println(v11)
                                _ = fi69
                            }
                        case "foo bar":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            v70 := GenS1{F0: "", F1: 91, F2: 1}
                            _ = v70
                        }
                    }
                }
            case 32767:
                switch (("\n\t" + "hello") + "hello") {
                case "hello":
                    var v71 uint16 = 447
                    _ = v71
                }
            }
            var v72 uint = 0
            _ = v72
        } else {
            if v11 {
                _ = len(v1)
                if len(v4) > 0 {
                    v4[0] = -2147483648
                }
                var v73 int8 = -1
                _ = v73
            } else {
                var v74 uint = v36
                fmt.Println(v11, v6, v22)
                _ = len(v2)
                _ = v74
            }
            for fi75 := 0; fi75 < 15; fi75++ {
                var v76 uint = (v36 ^ 676)
                _ = fi75
                _ = v76
            }
        }
    }
    v11 = v11
    v35 = (v35 >> ((v36) & 63))
    v77 := GenS2{F0: v22, F1: v35, F2: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v77
}
