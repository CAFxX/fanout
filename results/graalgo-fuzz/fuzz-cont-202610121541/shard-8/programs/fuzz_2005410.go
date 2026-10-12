package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 int32; F1 int16; F2 uint32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64; F1 int8 }

func main() {
    var v1 uint16 = 231
    var v2 int64 = -9223372036854775808
    v3 := make(map[uint32]uint32)
    v1 -= 65535
    v2 -= 1
    switch "\n\t" {
    case "hello":
        _ = len(v3)
    case "foo bar":
        var v4 int8 = 0
        _ = v4
    default:
        v5 := v3[216]
        _ = v5
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(v1)
        v7 := GenS1{F0: (fi6 * fi6), F1: -2147483648, F2: -69}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 2147483647, F1: int32(v2), F2: 42}
    if (false && (1 != 2147483647)) {
        switch "a" {
        case "":
            v1 -= (v1 - (v1 % ((v1) | 1)))
        case "foo bar":
            v3[789] = 2147483648
        default:
            fmt.Println(v1)
        }
        if true {
            for fi9 := 0; fi9 < 10; fi9++ {
                for fi10 := 0; fi10 < 7; fi10++ {
                    v2 += 1
                    var v11 uint16 = v1
                    var v12 uint8 = 19
                    v13 := GenS2{F0: v12}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                _ = fi9
            }
            for fi14 := 0; fi14 < 3; fi14++ {
                if (32767 > 1) {
                    var v15 int64 = v2
                    _ = v15
                }
                _ = fi14
            }
        }
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 uint16 = v1
            var v18 uint8 = 0
            if true {
                for fi19 := 0; fi19 < 13; fi19++ {
                    if false {
                        switch fi19 {
                        case 9223372036854775807:
                            v20 := GenS3{F0: false, F1: v18}
                            _ = v20
                        case -2147483648:
                            if true {
                                v21 := GenS1{F0: fi19, F1: 155, F2: 100}
                                if (!false) {
                                    v22 := GenS4{F0: 702}
                                    v1 -= (v17 << ((2) & 63))
                                    var v23 int = fi19
                                    _ = v22
                                    _ = v23
                                } else {
                                    var v24 bool = (0 == 1)
                                    switch (^(fi19)) {
                                    case 942:
                                        v25 := GenS3{F0: (0 >= (v1 * (v17 - v1))), F1: v18}
                                        _ = v25
                                    case 127:
                                        if false {
                                            var v26 uint16 = 32767
                                            _ = v26
                                        }
                                    case 4294967295:
                                        var v27 float32 = (1e10 - (3.14159 + 1.5))
                                        _ = v27
                                    }
                                    _ = v24
                                }
                                if true {
                                    v28 := v3[1]
                                    fmt.Println(v1, v18, v1)
                                    fmt.Println(fi16, fi16, v18)
                                    var v29 uint8 = (v18 + (v18 ^ 246))
                                    _ = v28
                                    _ = v29
                                }
                                v30 := GenS5{F0: 2, F1: 0}
                                _ = len(v3)
                                _ = v21
                                _ = v30
                            }
                        default:
                            switch fi19 {
                            case 542:
                                var v31 int32 = -1
                                _ = v31
                            case 42:
                                var v32 int8 = 1
                                _ = v32
                            case 1:
                                v33 := GenS1{F0: (fi19 >> ((65535) & 63)), F1: -2147483648, F2: 42}
                                _ = v33
                            }
                        }
                    } else {
                        if true {
                            fmt.Println(v1)
                            if false {
                                fmt.Println(v17, v18, v2)
                            }
                        }
                    }
                    _ = fi19
                }
                fmt.Println(fi16)
            } else {
                var v34 uint8 = (v18 << ((255) & 63))
                var v35 int64 = (v2 >> ((42) & 63))
                _ = v34
                _ = v35
            }
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    if (0 > 255) {
        var v36 int16 = 32767
        var v37 uint8 = 255
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        v1 += 2
        if true {
            var v39 int32 = -2147483648
            for fi40 := 0; fi40 < 18; fi40++ {
                v2 += v2
                switch "hello" {
                case "\n\t":
                    var v41 int64 = 1
                    var v42 int32 = 0
                    switch -9223372036854775808 {
                    case -198:
                        v39 = 1
                        v3[0] = 695
                    case -629:
                        var v43 uint16 = (v1 + 0)
                        _ = v43
                    case 127:
                        var v44 int = 32768
                        _ = v44
                    default:
                        var v45 uint64 = 142
                        _ = v45
                    }
                    _ = v41
                    _ = v42
                case "foo bar":
                    var v46 string = "a"
                    _ = v46
                }
                _ = fi40
            }
            _ = v39
        }
        v47 := GenS6{F0: "", F1: 32767}
        _ = fi38
        _ = v47
    }
    if (!((1e10 / 0.1) >= 0.0)) {
        var v48 float64 = 1e10
        var v49 int16 = 1000
        _ = v48
        _ = v49
    } else {
        v50 := GenS6{F0: "foo bar", F1: 65535}
        if (42 >= 18446744073709551615) {
            switch "hello" {
            case "hello":
                switch "foo bar" {
                case "foo bar":
                    var v51 uint64 = 18446744073709551615
                    _ = v51
                }
            default:
                v2 -= (v2 & v2)
            }
            v52 := GenS7{F0: -432, F1: -164, F2: 2147483648}
            for fi53 := 0; fi53 < 8; fi53++ {
                v54 := v3[42]
                v2 = (v2 * (v2 * -9223372036854775808))
                v54 -= (v54 * 125)
                _ = fi53
                _ = v54
            }
            _ = v52
        }
        v55 := []int{}
        _ = v50
        _ = v55
    }
    v56 := GenS2{F0: 255}
    delete(v3, 480)
    v2 = (v2 % ((v2) | 1))
    v2 -= -42
    if false {
        var v57 int = -1000
        v58 := GenS8{F0: 127}
        _ = v57
        _ = v58
    } else {
        v1 += (^(v1))
        v1 -= (v1 / ((v1) | 1))
        switch 599 {
        case 32768:
            for fi59 := 0; fi59 < 3; fi59++ {
                v60 := []int{(fi59 % ((fi59) | 1)), 1}
                for fi61 := 0; fi61 < 6; fi61++ {
                    v62 := GenS2{F0: 113}
                    if len(v60) >= 2 {
                        v60 = v60[1:len(v60)]
                    }
                    switch fi59 {
                    case 32767:
                        var v63 int16 = -1
                        _ = v63
                    case 128:
                        v64 := GenS6{F0: ((("\n\t" + "") + "") + ("\n\t" + "\n\t")), F1: 2}
                        _ = v64
                    case -2147483648:
                        fmt.Println(v1, fi61)
                    default:
                        v65 := GenS9{F0: v2, F1: 100}
                        var v66 uint8 = 187
                        fmt.Println(v2, v1)
                        if (-888 > int16(v66)) {
                            var v67 uint16 = 0
                            _ = v67
                        } else {
                            v1 += v1
                            v60 = append(v60, 127)
                        }
                        _ = v65
                        _ = v66
                    }
                    _ = fi61
                    _ = v62
                }
                _ = fi59
                _ = v60
            }
        case 0:
            fmt.Println(v1)
        case 1:
            v1 = v1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v56
    _ = v8
}
