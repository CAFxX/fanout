package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int16; F1 string; F2 uint }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]int)
    var v3 int16 = -1
    v4 := []int{95, -669}
    var v5 float32 = -2.25
    v6 := GenS1{F0: (^(v3)), F1: 97}
    v7 := GenS2{F0: -402, F1: 42}
    if true {
        var v8 uint = 18446744073709551615
        var v9 float32 = float32(0.0)
        v9 = v5
        _ = v8
        _ = v9
    } else {
        var v10 int16 = (v3 + 1000)
        v11 := GenS2{F0: -1, F1: 0}
        if ((1 < 1) || ((-(v10)) >= (v3 * v3))) {
            v1["\n\t"] = 4294967295
            fmt.Println(v5, v10, v10)
        }
        _ = v10
        _ = v11
    }
    fmt.Println(v5, v5)
    if (0 > 100) {
        var v12 uint = 255
        _ = v12
    }
    var v13 uint64 = 18446744073709551615
    var v14 string = "hello"
    v13 = (v13 % ((v13) | 1))
    fmt.Println(v3, v13, v3)
    if (v14 == "\n\t") {
        if (true && (16 < 1)) {
            var v15 int32 = 406
            _ = v15
        } else {
            if (1e10 < -2.25) {
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 uint16 = 65535
                    fmt.Println(v3, v3, v5)
                    v18 := GenS3{F0: 4294967295, F1: v17}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            } else {
                var v19 uint16 = 0
                var v20 int = int(v3)
                _ = v19
                _ = v20
            }
        }
        v21 := GenS4{F0: 0, F1: 929}
        _ = v21
    } else {
        v13 = (v13 * 4294967295)
        if (!(3.14159 != ((v5 - 1e10) - float32(3.14159)))) {
            var v22 uint = 793
            v23 := GenS1{F0: v3, F1: 128}
            _ = v22
            _ = v23
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = len(v2)
        }
        v24 := GenS1{F0: v3, F1: uint8(v13)}
        _ = v24
    }
    if true {
        if true {
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 uint = 792
                if true {
                    switch ("hello" + v14) {
                    case "a":
                        v4 = append(v4, fi25)
                    case "\n\t":
                        v14 = "a"
                    default:
                        v27 := GenS5{F0: -1, F1: (("" + (v14 + (v14 + ((v14 + ((((("\n\t" + (("\n\t" + "a") + v14)) + v14) + v14) + "\n\t") + v14)) + v14)))) >= v14)}
                        _ = v27
                    }
                    fmt.Println(fi25, v5)
                    var v28 uint = 18446744073709551615
                    _ = v28
                } else {
                    if true {
                        switch v14 {
                        case "foo bar":
                            v5 = 0.0
                        default:
                            var v29 int = fi25
                            _ = v29
                        }
                        var v30 bool = true
                        _ = v30
                    } else {
                        for fi31 := 0; fi31 < 13; fi31++ {
                            v13 += (v13 << (((^(v26))) & 63))
                            switch fi31 {
                            case 32768:
                                fmt.Println(fi31, v5)
                            default:
                                if len(v4) > 0 {
                                    v4[0] = (fi25 << ((v26) & 63))
                                }
                            }
                            var v32 int32 = -890
                            _ = fi31
                            _ = v32
                        }
                    }
                }
                _ = fi25
                _ = v26
            }
            var v33 int16 = (v3 - 819)
            v34 := GenS4{F0: 127, F1: 32767}
            _ = v33
            _ = v34
        }
        switch "\n\t" {
        case "\n\t":
            if false {
                v4 = append(v4, 4294967295)
                fmt.Println(v14)
                var v35 uint16 = 469
                _ = v35
            }
        case "foo bar":
            _ = len(v1)
        }
        switch v14 {
        case "foo bar":
            for fi36 := 0; fi36 < 2; fi36++ {
                v14 = "a"
                _ = fi36
            }
        default:
            if true {
                for fi37 := 0; fi37 < 5; fi37++ {
                    if (false && (32767 < 42)) {
                        var v38 float64 = (float64((v5 + ((-2.25 / (v5 * v5)) / (v5 + (v5 * (1.5 + float32(float64(float32((0.1 * (float64(v5) * 0.0))))))))))) - 3.14159)
                        _ = v38
                    }
                    _ = fi37
                }
                v39 := GenS6{F0: v3, F1: v14, F2: 42}
                _ = v39
            } else {
                var v40 uint16 = 2
                _ = v40
            }
        }
    } else {
        v41 := GenS4{F0: 42, F1: 2}
        if (!true) {
            switch 4294967295 {
            case 223:
                v14 = v14
            case -1:
                v5 = 0.1
            case -319:
                var v42 int = 82
                _ = v42
            }
            fmt.Println(v3, v3, v13)
        }
        switch 793 {
        case -2147483648:
            switch v14 {
            case "\n\t":
                var v43 int64 = -435
                _ = v43
            }
        case -1:
            if (v3 <= (v3 - v3)) {
                var v44 float64 = 1e10
                v14 = (v14 + v14)
                v14 = "\n\t"
                _ = v44
            }
        default:
            if (648 >= 1) {
                if false {
                    for fi45 := 0; fi45 < 20; fi45++ {
                        if (!(18446744073709551615 <= 164)) {
                            var v46 int16 = 42
                            var v47 uint = 1
                            _ = v46
                            _ = v47
                        }
                        _ = fi45
                    }
                }
                for fi48 := 0; fi48 < 3; fi48++ {
                    if true {
                        v49 := GenS7{F0: 2, F1: 262}
                        if true {
                            v3 -= (^(v3))
                        } else {
                            var v50 uint16 = 656
                            v14 = ("a" + (v14 + v14))
                            v51 := GenS8{F0: 2}
                            _ = v50
                            _ = v51
                        }
                        _ = v49
                    } else {
                        if false {
                            var v52 int = (fi48 % (((fi48 % ((fi48) | 1))) | 1))
                            _ = v52
                        } else {
                            var v53 uint16 = 2
                            _ = v53
                        }
                        var v54 int8 = 1
                        _ = v54
                    }
                    _ = fi48
                }
                switch 392 {
                case 0:
                    v3 += v3
                default:
                    var v55 int16 = (v3 * v3)
                    _ = v55
                }
            }
        }
        _ = v41
    }
    if ("a" == (v14 + (v14 + (v14 + v14)))) {
        v56 := GenS7{F0: 42, F1: -356}
        if true {
            v3 -= (v3 * 822)
        }
        if len(v4) > 0 {
            v4[0] = 692
        }
        _ = v56
    }
    var v57 int64 = 791
    v57 = v57
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v7
}
