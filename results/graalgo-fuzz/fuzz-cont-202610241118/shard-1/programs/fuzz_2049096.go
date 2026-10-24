package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int8; F1 uint; F2 uint }
type GenS4 struct { F0 int32; F1 string; F2 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint }

func main() {
    var v1 int64 = -1
    var v2 uint32 = 2
    v3 := make(map[int]string)
    var v4 int = 127
    v5 := make(map[uint32]bool)
    var v6 int32 = 376
    v4 -= 1
    switch "hello" {
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            v7 := GenS1{F0: 1}
            _ = v7
        case "\n\t":
            for fi8 := 0; fi8 < 16; fi8++ {
                switch "foo bar" {
                case "foo bar":
                    v6 = v6
                case "":
                    switch (^(v4)) {
                    case 127:
                        var v9 int64 = -9223372036854775808
                        _ = v9
                    }
                }
                for fi10 := 0; fi10 < 11; fi10++ {
                    switch fi8 {
                    case -1:
                        v1 -= v1
                    }
                    _ = fi10
                }
                v11 := []int{(v4 / ((fi8) | 1)), -146, 849, v4}
                _ = fi8
                _ = v11
            }
        case "":
            v12 := GenS1{F0: 0}
            _ = v12
        }
    case "hello":
        for fi13 := 0; fi13 < 11; fi13++ {
            switch ("\n\t" + ("\n\t" + ("foo bar" + "a"))) {
            case "foo bar":
                switch v4 {
                case -1:
                    var v14 uint32 = (v2 - 367)
                    _ = v14
                case 2147483647:
                    switch (-(v4)) {
                    case 0:
                        v2 = v2
                    }
                }
            }
            var v15 int16 = 0
            fmt.Println(v2, v1)
            _ = fi13
            _ = v15
        }
    case "\n\t":
        v6 += v6
    }
    var v16 int64 = 9223372036854775807
    v17 := GenS1{F0: 1}
    var v18 float64 = 3.14159
    var v19 bool = (149 <= 2)
    v2 = v2
    fmt.Println(v16, v1, v2)
    fmt.Println(v18)
    for fi20 := 0; fi20 < 2; fi20++ {
        switch (fi20 ^ v4) {
        case 9223372036854775807:
            v1 = 0
        case 40:
            if false {
                if v19 {
                    switch "\n\t" {
                    case "a":
                        for fi21 := 0; fi21 < 8; fi21++ {
                            if (1 != 4294967295) {
                                v22 := GenS2{F0: 255, F1: 139}
                                v6 += (v6 / ((v6) | 1))
                                fmt.Println(v6, v2, v6)
                                fmt.Println(v6)
                                var v23 uint64 = 18446744073709551615
                                _ = v22
                                _ = v23
                            } else {
                                fmt.Println(v1, v4, v6)
                                v2 = 4294967295
                                v24 := GenS2{F0: 255, F1: 0}
                                _ = v24
                            }
                            _ = fi21
                        }
                    }
                    for fi25 := 0; fi25 < 16; fi25++ {
                        switch ((("hello" + "a") + (("foo bar" + "a") + "foo bar")) + ("hello" + ("" + ""))) {
                        case "\n\t":
                            if ((v6 + v6) <= int32(v1)) {
                                v26 := GenS1{F0: 65535}
                                if false {
                                    v16 = v16
                                    var v27 uint32 = v2
                                    _ = v27
                                }
                                v28 := GenS1{F0: 204}
                                _ = v26
                                _ = v28
                            }
                        }
                        if v19 {
                            v29 := GenS2{F0: 118, F1: 0}
                            if v19 {
                                var v30 int8 = 127
                                if (v18 <= v18) {
                                    v31 := GenS1{F0: 570}
                                    var v32 float32 = (1.5 * 3.14159)
                                    v33 := GenS1{F0: 0}
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                } else {
                                    v19 = v19
                                    v19 = v19
                                    var v34 int16 = 373
                                    _ = v34
                                }
                                var v35 uint64 = 2
                                _ = v30
                                _ = v35
                            }
                            _ = v29
                        } else {
                            if (!true) {
                                v36 := []int{(fi20 >> ((42) & 63))}
                                v5[v2] = (4294967295 != 18446744073709551615)
                                _ = v36
                            } else {
                                if ((v1 + (v1 >> ((1) & 63))) < (v16 ^ (v1 + 976))) {
                                    var v37 uint = uint(v2)
                                    v16 = v16
                                    v37 = v37
                                    _ = v37
                                }
                            }
                            v4 = (fi25 << ((1) & 63))
                            if (v6 <= (v6 >> ((2) & 63))) {
                                var v38 int64 = 1
                                if v19 {
                                    fmt.Println(v2, v16)
                                }
                                v16 -= (v16 >> ((uint(v2)) & 63))
                                var v39 uint = 65535
                                switch "" {
                                case "hello":
                                    if v19 {
                                        if (737 == 42) {
                                            v40 := GenS3{F0: -128, F1: v39, F2: (^(v39))}
                                            fmt.Println(v4)
                                            var v41 float64 = v18
                                            _ = v40
                                            _ = v41
                                        } else {
                                            _ = len(v3)
                                            fmt.Println(v39, v2)
                                            v42 := []int{fi20, -9223372036854775808, (fi25 | v4)}
                                            _ = v42
                                        }
                                        var v43 int16 = -505
                                        var v44 int16 = 1
                                        _ = v43
                                        _ = v44
                                    }
                                case "foo bar":
                                    var v45 int32 = v6
                                    _ = v45
                                }
                                fmt.Println(v2, v18, fi25)
                                _ = v38
                                _ = v39
                            } else {
                                v46 := []int{v4, fi20, (fi20 / ((fi25) | 1)), fi20}
                                _ = v46
                            }
                        }
                        var v47 uint16 = 32767
                        _ = fi25
                        _ = v47
                    }
                    var v48 uint32 = (v2 + (v2 * v2))
                    _ = v48
                } else {
                    var v49 float64 = (0.1 + (v18 * (v18 + (v18 * 0.1))))
                    _ = v49
                }
            }
        case 4294967295:
            var v50 uint8 = 128
            _ = v50
        default:
            fmt.Println(v4, v6, v18)
        }
        var v51 uint64 = 1
        _ = fi20
        _ = v51
    }
    if (65535 != 1) {
        v52 := GenS4{F0: (^(v6)), F1: "", F2: (^(v2))}
        _ = v52
    } else {
        var v53 bool = true
        _ = v53
    }
    fmt.Println(v18, v16, v1)
    if false {
        for fi54 := 0; fi54 < 20; fi54++ {
            var v55 int16 = 42
            switch 32768 {
            case 42:
                if (true || v19) {
                    v56 := v3[fi54]
                    if (("hello" > v56) || v19) {
                        v55 = (v55 | 1)
                    }
                    for fi57 := 0; fi57 < 12; fi57++ {
                        switch (fi57 & (v4 << ((uint(v2)) & 63))) {
                        case -1:
                            if (float32((-2.25 * v18)) < 0.1) {
                                var v58 int64 = (v1 % ((v1) | 1))
                                switch v56 {
                                case "a":
                                    var v59 uint8 = 128
                                    _ = v59
                                }
                                _ = v58
                            }
                        case 9223372036854775807:
                            delete(v3, (fi54 | (^(v4))))
                        case 1:
                            v2 -= (v2 ^ (v2 + v2))
                        default:
                            var v60 int32 = (v6 - v6)
                            _ = v60
                        }
                        var v61 uint64 = 439
                        fmt.Println(v2, v1)
                        _ = fi57
                        _ = v61
                    }
                    _ = v56
                } else {
                    v62 := GenS5{F0: false}
                    _ = v62
                }
            case 1:
                fmt.Println(v4, v16)
            default:
                for fi63 := 0; fi63 < 1; fi63++ {
                    v18 = v18
                    switch "a" {
                    case "\n\t":
                        v6 = 2147483647
                        var v64 float32 = ((0.1 * float32((((-2.25 / float64(3.14159)) / 1e10) + (float64(((0.1 - 0.0) - -2.25)) - v18)))) + (3.14159 / (1e10 + 0.0)))
                        v65 := GenS6{F0: 42}
                        _ = v64
                        _ = v65
                    case "":
                        v66 := []int{}
                        _ = v66
                    default:
                        v18 = 1e10
                    }
                    _ = fi63
                }
            }
            _ = fi54
            _ = v55
        }
        switch 9223372036854775807 {
        case 48:
            if (18446744073709551615 == 255) {
                fmt.Println(v2, v16, v6)
                _ = len(v5)
            }
        default:
            var v67 float32 = 0.0
            _ = v67
        }
        switch (v4 | v4) {
        case 0:
            for fi68 := 0; fi68 < 3; fi68++ {
                var v69 uint = 42
                _ = fi68
                _ = v69
            }
        }
    } else {
        v70 := v3[(^(v4))]
        fmt.Println(v70, v70)
        _ = v70
    }
    for fi71 := 0; fi71 < 17; fi71++ {
        if (((v4 <= (v4 / ((v4) | 1))) || v19) && (32767 != 1000)) {
            v1 -= (v1 << ((0) & 63))
            v3[(v4 >> ((255) & 63))] = "\n\t"
        } else {
            _ = len(v3)
        }
        _ = fi71
    }
    for fi72 := 0; fi72 < 9; fi72++ {
        var v73 uint32 = (v2 % ((42) | 1))
        for fi74 := 0; fi74 < 19; fi74++ {
            switch (-(fi72)) {
            case -941:
                var v75 string = ("a" + "")
                _ = v75
            case 32767:
                v16 = v16
            }
            v16 += -9223372036854775808
            v16 = -375
            _ = fi74
        }
        _ = fi72
        _ = v73
    }
    for fi76 := 0; fi76 < 3; fi76++ {
        var v77 uint16 = 590
        v78 := []int{v4, (v4 * (fi76 >> ((18446744073709551615) & 63)))}
        _ = fi76
        _ = v77
        _ = v78
    }
    v79 := GenS2{F0: 106, F1: 42}
    if true {
        var v80 int = v4
        var v81 float64 = (1.5 + v18)
        var v82 int = -532
        _ = v80
        _ = v81
        _ = v82
    } else {
        fmt.Println(v19, v18)
        var v83 int64 = (v1 ^ (v16 << ((18446744073709551615) & 63)))
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v79
}
