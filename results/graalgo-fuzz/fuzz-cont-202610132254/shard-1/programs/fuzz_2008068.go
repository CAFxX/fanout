package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 string }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 32767
    var v3 bool = true
    v4 := make(map[int]bool)
    var v5 uint16 = 0
    v6 := GenS1{F0: (v5 | v5), F1: (v2 + v2), F2: ("hello" + "hello")}
    switch (v2 >> ((42) & 63)) {
    case -16:
        fmt.Println(v3)
    case 334:
        v7 := v4[v2]
        _ = v7
    case 431:
        for fi8 := 0; fi8 < 7; fi8++ {
            switch v2 {
            case 42:
                if false {
                    delete(v4, v2)
                }
            case 127:
                var v9 int32 = -61
                _ = v9
            case -162:
                v10 := []int{(fi8 << ((2) & 63)), fi8, fi8}
                _ = v10
            }
            _ = fi8
        }
    default:
        v11 := GenS2{F0: 128, F1: -1}
        _ = v11
    }
    v12 := GenS2{F0: 234, F1: -128}
    for fi13 := 0; fi13 < 13; fi13++ {
        switch fi13 {
        case 127:
            for fi14 := 0; fi14 < 19; fi14++ {
                v2 += (fi13 - (fi14 % (((fi13 - fi14)) | 1)))
                fmt.Println(v2, v3, v2)
                if v3 {
                    switch "a" {
                    case "\n\t":
                        if (true || false) {
                            fmt.Println(v5)
                            v15 := []int{-593, fi13, fi14, v2}
                            v3 = true
                            _ = v15
                        } else {
                            v16 := GenS3{F0: -20}
                            v17 := GenS4{F0: -1, F1: 2147483648}
                            switch "" {
                            case "a":
                                var v18 uint64 = 18446744073709551615
                                _ = v18
                            case "foo bar":
                                if (!(false || (42 < 127))) {
                                    switch "a" {
                                    case "\n\t":
                                        v19 := GenS5{F0: 1, F1: 25}
                                        _ = v19
                                    case "foo bar":
                                        var v20 string = ("a" + "hello")
                                        _ = v20
                                    case "hello":
                                        switch (-(v2)) {
                                        case -9223372036854775808:
                                            var v21 uint8 = 127
                                            _ = v21
                                        case -2147483648:
                                            var v22 int32 = 0
                                            _ = v22
                                        default:
                                            fmt.Println(fi13, v2)
                                        }
                                    default:
                                        var v23 int8 = 100
                                        _ = v23
                                    }
                                }
                            case "\n\t":
                                fmt.Println(fi14, fi14, v5)
                            default:
                                v24 := GenS1{F0: 255, F1: 127, F2: ("" + "")}
                                _ = v24
                            }
                            _ = v16
                            _ = v17
                        }
                        var v25 int32 = 42
                        _ = v25
                    default:
                        v5 = v5
                    }
                } else {
                    v3 = (128 >= 127)
                    v3 = (!v3)
                    if (!(255 != 2)) {
                        var v26 uint32 = 4294967295
                        var v27 float32 = -2.25
                        v3 = false
                        _ = v26
                        _ = v27
                    }
                    switch (fi13 / (((v2 ^ fi13)) | 1)) {
                    case -956:
                        var v28 int16 = -32768
                        _ = v28
                    case -900:
                        switch fi13 {
                        case 9223372036854775807:
                            var v29 int8 = -1
                            _ = v29
                        case 128:
                            v5 = (v5 * v5)
                        case 1:
                            var v30 float32 = (((float32(0.1) * float32(float64(((3.14159 / (0.1 + 0.0)) * (1e10 - 1.5))))) - 3.14159) / -2.25)
                            _ = v30
                        default:
                            var v31 uint32 = 476
                            _ = v31
                        }
                        v32 := GenS5{F0: 348, F1: 18446744073709551615}
                        var v33 uint32 = 42
                        v3 = (v33 > (v33 ^ (v33 & v33)))
                        _ = v32
                        _ = v33
                    default:
                        switch fi14 {
                        case 127:
                            v34 := GenS6{F0: 42}
                            _ = v34
                        case 2147483647:
                            var v35 float64 = -2.25
                            _ = v35
                        case -1:
                            var v36 int32 = -739
                            _ = v36
                        default:
                            fmt.Println(v2, v2)
                        }
                    }
                    if v3 {
                        fmt.Println(v3)
                        var v37 float64 = (-2.25 - 0.0)
                        if (1 == -351) {
                            fmt.Println(fi13, fi14)
                            v38 := []int{1, 9223372036854775807, fi14}
                            if len(v38) > 0 {
                                v38[0] = int(v5)
                            }
                            _ = v38
                        }
                        _ = v37
                    }
                    var v39 string = "\n\t"
                    _ = v39
                }
                _ = fi14
            }
        case 458:
            for fi40 := 0; fi40 < 4; fi40++ {
                switch (fi13 >> ((706) & 63)) {
                case 128:
                    var v41 float32 = (3.14159 / 3.14159)
                    _ = v41
                }
                _ = fi40
            }
        case 128:
            var v42 int32 = int32(fi13)
            _ = v42
        }
        v43 := GenS7{F0: 2}
        switch "a" {
        case "a":
            var v44 int8 = 33
            _ = v44
        case "":
            v3 = (!true)
        default:
            for fi45 := 0; fi45 < 6; fi45++ {
                fmt.Println(fi13)
                v46 := GenS1{F0: (v5 % ((v5) | 1)), F1: v2, F2: ("\n\t" + ((("a" + ((("\n\t" + "") + "foo bar") + "")) + "\n\t") + "hello"))}
                _ = fi45
                _ = v46
            }
        }
        _ = fi13
        _ = v43
    }
    var v47 int8 = -78
    switch 127 {
    case -9223372036854775808:
        v47 = v47
    case 9223372036854775807:
        for fi48 := 0; fi48 < 3; fi48++ {
            v2 = 32767
            if true {
                var v49 float64 = 0.1
                _ = v49
            }
            _ = fi48
        }
    }
    v3 = (0 != 65535)
    v2 = (v2 | (v2 * v2))
    for fi50 := 0; fi50 < 10; fi50++ {
        switch fi50 {
        case -176:
            if false {
                v47 = (-(v47))
                v2 -= fi50
            }
        case 1:
            var v51 uint8 = 0
            _ = v51
        default:
            v52 := GenS3{F0: 32767}
            _ = v52
        }
        _ = fi50
    }
    v47 += 58
    if (!v3) {
        if (255 <= 0) {
            if (3.14159 != float32(1e10)) {
                v3 = (18446744073709551615 < 2)
                for fi53 := 0; fi53 < 1; fi53++ {
                    if true {
                        var v54 uint16 = (^(v5))
                        _ = v54
                    }
                    var v55 int = fi53
                    _ = fi53
                    _ = v55
                }
            } else {
                v56 := v1[2147483647]
                var v57 int16 = -1
                _ = v56
                _ = v57
            }
            for fi58 := 0; fi58 < 17; fi58++ {
                v59 := GenS7{F0: 1}
                v3 = true
                switch "" {
                case "a":
                    if v3 {
                        v60 := GenS1{F0: 0, F1: (fi58 | v2), F2: "foo bar"}
                        _ = v60
                    } else {
                        var v61 uint16 = (v5 / ((v5) | 1))
                        for fi62 := 0; fi62 < 3; fi62++ {
                            v63 := GenS6{F0: 9223372036854775807}
                            if ((!(682 > 255)) || v3) {
                                fmt.Println(v5, fi62, fi62)
                                var v64 uint64 = 294
                                var v65 int = fi62
                                _ = v64
                                _ = v65
                            } else {
                                v66 := []int{}
                                var v67 uint32 = 1
                                v2 += v2
                                _ = v66
                                _ = v67
                            }
                            v5 = (v61 * (^(v61)))
                            v68 := []int{}
                            _ = fi62
                            _ = v63
                            _ = v68
                        }
                        _ = v61
                    }
                case "":
                    v3 = v3
                case "\n\t":
                    v3 = v3
                }
                _ = fi58
                _ = v59
            }
        }
        fmt.Println(v2, v47, v2)
    } else {
        for fi69 := 0; fi69 < 16; fi69++ {
            if ((174 < 207) && v3) {
                v70 := []int{(fi69 << ((255) & 63))}
                if ((236 > 127) || v3) {
                    fmt.Println(fi69, fi69)
                    v71 := GenS8{F0: v5}
                    _ = v71
                } else {
                    var v72 uint16 = v5
                    v5 = v72
                    v72 = v5
                    _ = v72
                }
                _ = v70
            }
            for fi73 := 0; fi73 < 20; fi73++ {
                var v74 string = "\n\t"
                fmt.Println(v47)
                var v75 string = (v74 + v74)
                _ = fi73
                _ = v74
                _ = v75
            }
            _ = fi69
        }
    }
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
