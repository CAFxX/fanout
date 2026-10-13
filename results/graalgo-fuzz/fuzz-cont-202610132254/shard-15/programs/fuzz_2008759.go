package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint }

func main() {
    var v1 int16 = 32767
    v2 := []int{128, 566, 0, 128}
    var v3 int = 1
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]bool)
    v6 := v5[4294967295]
    var v7 int16 = v1
    switch 32767 {
    case 292:
        if v6 {
            delete(v4, 2147483648)
        } else {
            var v8 uint8 = 1
            _ = v8
        }
    case -9223372036854775808:
        v9 := GenS1{F0: 1, F1: -9223372036854775808}
        _ = v9
    }
    fmt.Println(v1, v3, v6)
    v10 := v4[1]
    if ((1 != 552) && (767 != 2)) {
        var v11 int32 = 2147483647
        _ = v11
    }
    if (2147483648 <= 379) {
        v7 += v7
    } else {
        if v6 {
            if v6 {
                var v12 float64 = 0.1
                switch (^(v3)) {
                case -685:
                    v12 = (v12 + 0.0)
                case 101:
                    if len(v2) > 0 {
                        v2[0] = -1
                    }
                case 1:
                    fmt.Println(v3)
                }
                _ = v12
            } else {
                if (!(42 <= 117)) {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v14 := GenS1{F0: 0, F1: 42}
                        _ = fi13
                        _ = v14
                    }
                    var v15 uint16 = 970
                    var v16 uint16 = v15
                    _ = v15
                    _ = v16
                }
                fmt.Println(v6, v1)
                fmt.Println(v7)
            }
            if true {
                v6 = true
                var v17 int64 = 849
                if (-1 == -128) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi18 := 0; fi18 < 13; fi18++ {
                        if true {
                            v19 := GenS2{F0: 128, F1: 65535}
                            var v20 bool = v6
                            _ = v19
                            _ = v20
                        } else {
                            for fi21 := 0; fi21 < 15; fi21++ {
                                var v22 int64 = v17
                                v23 := GenS2{F0: 2, F1: 42}
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                            var v24 uint64 = 4294967295
                            v24 -= (^(v24))
                            _ = v24
                        }
                        _ = fi18
                    }
                    switch v3 {
                    case 9223372036854775807:
                        if (635 != 984) {
                            delete(v5, 42)
                            v25 := GenS2{F0: 70, F1: 42}
                            if v6 {
                                fmt.Println(v10, v3, v7)
                                v5[2] = v6
                            } else {
                                fmt.Println(v6, v17)
                                v26 := GenS2{F0: 42, F1: 0}
                                var v27 int64 = 9223372036854775807
                                _ = v26
                                _ = v27
                            }
                            _ = v25
                        } else {
                            v3 -= (v3 | (v3 * v3))
                        }
                    default:
                        switch 42 {
                        case 32767:
                            _ = len(v4)
                        case 128:
                            for fi28 := 0; fi28 < 11; fi28++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch 1 {
                                case 9223372036854775807:
                                    if true {
                                        v10 = v10
                                        var v29 float32 = ((1.5 - ((3.14159 + (float32(0.0) * (3.14159 + 3.14159))) * 1.5)) - 0.1)
                                        if (0 != 2147483647) {
                                            var v30 uint16 = 125
                                            v29 -= v29
                                            _ = v30
                                        } else {
                                            if v6 {
                                                v6 = (2 <= 2147483648)
                                                fmt.Println(fi28, v6, v6)
                                            }
                                        }
                                        _ = v29
                                    }
                                }
                                _ = fi28
                            }
                        case -1:
                            var v31 int = v3
                            _ = v31
                        }
                    }
                }
                _ = v17
            }
            var v32 float32 = 3.14159
            _ = v32
        } else {
            var v33 float64 = 1.5
            v3 = (-(v3))
            _ = v33
        }
        _ = len(v5)
        if (2147483648 <= 68) {
            var v34 string = "foo bar"
            if false {
                var v35 uint32 = 2
                for fi36 := 0; fi36 < 11; fi36++ {
                    if ((v35 / ((v35) | 1)) > v35) {
                        var v37 uint = 18446744073709551615
                        if (v35 != (^(v35))) {
                            v38 := GenS3{F0: (v37 << (((v37 * v37)) & 63))}
                            var v39 uint8 = 240
                            _ = v38
                            _ = v39
                        } else {
                            if (!v6) {
                                for fi40 := 0; fi40 < 5; fi40++ {
                                    _ = len(v5)
                                    var v41 int8 = 127
                                    v42 := GenS3{F0: (v37 | v37)}
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                }
                                v43 := GenS4{F0: v35}
                                v5[(v35 >> (((v37 & (v37 | 2))) & 63))] = v6
                                _ = v43
                            }
                        }
                        for fi44 := 0; fi44 < 2; fi44++ {
                            v37 = v37
                            _ = fi44
                        }
                        _ = v37
                    } else {
                        v4[v35] = "a"
                        v2 = append(v2, (v3 + -785))
                    }
                    for fi45 := 0; fi45 < 10; fi45++ {
                        var v46 int32 = 1
                        v7 -= 0
                        _ = fi45
                        _ = v46
                    }
                    fmt.Println(v3, v7)
                    _ = fi36
                }
                var v47 int64 = 9223372036854775807
                _ = v35
                _ = v47
            } else {
                switch (v3 & -2147483648) {
                case -9223372036854775808:
                    var v48 uint = 12
                    _ = v48
                case 42:
                    var v49 uint64 = 158
                    _ = v49
                case 732:
                    if v6 {
                        v7 -= (v1 << ((255) & 63))
                        if false {
                            delete(v4, 861)
                            if len(v2) > 0 {
                                v2[0] = (v3 ^ v3)
                            }
                        } else {
                            if false {
                                _ = len(v2)
                            } else {
                                var v50 uint64 = 635
                                _ = v50
                            }
                            v51 := GenS3{F0: 42}
                            _ = v51
                        }
                        var v52 uint64 = 547
                        _ = v52
                    }
                }
                if v6 {
                    switch v34 {
                    case "a":
                        var v53 int8 = -1
                        _ = v53
                    case "":
                        if (191 != 255) {
                            v54 := GenS5{F0: v34}
                            _ = v54
                        }
                    case "\n\t":
                        v55 := GenS6{F0: 6}
                        _ = v55
                    default:
                        var v56 uint8 = 93
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 14; fi57++ {
                        var v58 float64 = 1.5
                        v10 = ("\n\t" + ((v10 + "a") + v10))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi57
                        _ = v58
                    }
                    v7 = (-(v1))
                } else {
                    fmt.Println(v10)
                    for fi59 := 0; fi59 < 7; fi59++ {
                        _ = len(v5)
                        v60 := v5[920]
                        v61 := GenS4{F0: 42}
                        _ = fi59
                        _ = v60
                        _ = v61
                    }
                    if v6 {
                        var v62 float64 = (3.14159 + float64(0.1))
                        _ = v62
                    }
                }
                fmt.Println(v1)
            }
            switch (((v10 + (v34 + ("a" + ""))) + "") + v34) {
            case "\n\t":
                for fi63 := 0; fi63 < 12; fi63++ {
                    var v64 int8 = 119
                    _ = len(v2)
                    _ = fi63
                    _ = v64
                }
            case "":
                v1 += (v1 / ((-1) | 1))
            case "foo bar":
                var v65 uint8 = 1
                _ = v65
            default:
                v10 = "a"
            }
            _ = v34
        }
    }
    var v66 int = (v3 & 42)
    for fi67 := 0; fi67 < 14; fi67++ {
        var v68 string = v10
        _ = fi67
        _ = v68
    }
    for fi69 := 0; fi69 < 5; fi69++ {
        if v6 {
            for fi70 := 0; fi70 < 15; fi70++ {
                v10 = v10
                if v6 {
                    fmt.Println(fi69, v7, v10)
                    var v71 int8 = -46
                    fmt.Println(v66, v71)
                    _ = v71
                }
                var v72 int = fi70
                _ = fi70
                _ = v72
            }
            var v73 int32 = 560
            _ = v73
        }
        _ = fi69
    }
    var v74 string = (v10 + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v7:", v7)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v7
    _ = v74
}
