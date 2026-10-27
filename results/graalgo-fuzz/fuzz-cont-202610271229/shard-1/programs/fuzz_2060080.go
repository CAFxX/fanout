package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 string; F2 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32; F1 int32; F2 int32 }

func main() {
    v1 := []string{("" + "\n\t"), ("hello" + "\n\t"), (("\n\t" + "\n\t") + "hello")}
    v2 := make(map[int]int)
    var v3 float32 = (1e10 - (1.5 * float32(1e10)))
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    var v6 float64 = (float64(float32(1e10)) / -2.25)
    switch "" {
    case "":
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    case "a":
        for fi7 := 0; fi7 < 5; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                if (1 <= -1) {
                    v9 := GenS1{F0: 0, F1: 18446744073709551615}
                    _ = v9
                } else {
                    v3 += ((v3 - -2.25) + float32(3.14159))
                    v10 := GenS2{F0: -672, F1: 1, F2: 262}
                    var v11 uint16 = 255
                    _ = v10
                    _ = v11
                }
                fmt.Println(fi8, v6, fi8)
                _ = fi8
            }
            _ = fi7
        }
    }
    v6 = float64(float32((3.14159 * (v6 + 3.14159))))
    v3 = (3.14159 + (v3 - 0.0))
    switch "foo bar" {
    case "":
        switch "a" {
        case "":
            var v12 uint64 = 189
            _ = v12
        }
    default:
        var v13 int32 = 104
        _ = v13
    }
    if (!false) {
        v6 = v6
    }
    var v14 uint32 = 674
    if false {
        v15 := GenS2{F0: 1000, F1: (^(v14)), F2: 498}
        var v16 uint8 = 42
        _ = v15
        _ = v16
    } else {
        var v17 string = "hello"
        _ = v17
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        v3 = (v3 + (0.0 - 3.14159))
        for fi19 := 0; fi19 < 12; fi19++ {
            if true {
                if true {
                    v20 := GenS3{F0: 4294967295}
                    var v21 int64 = 1
                    _ = v20
                    _ = v21
                } else {
                    if (!(18446744073709551615 > 407)) {
                        fmt.Println(v6, v6, v6)
                    }
                    var v22 int32 = 42
                    if true {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v22, fi18)
                    }
                    _ = v22
                }
                var v23 int = fi19
                _ = v23
            } else {
                var v24 uint32 = 2147483648
                if true {
                    fmt.Println(v3, fi19)
                } else {
                    switch (fi18 & -2147483648) {
                    case 32767:
                        switch "" {
                        case "foo bar":
                            var v25 float32 = v3
                            _ = v25
                        default:
                            v24 = (v14 + (^(v14)))
                        }
                    case 1:
                        v24 -= v24
                    default:
                        v26 := GenS3{F0: 0}
                        _ = v26
                    }
                    if (int16(v24) > -566) {
                        var v27 string = "hello"
                        if true {
                            v28 := GenS2{F0: 32767, F1: v14, F2: 0}
                            var v29 uint64 = 975
                            _ = v28
                            _ = v29
                        }
                        _ = v27
                    }
                    switch (((("\n\t" + "") + "") + "\n\t") + "\n\t") {
                    case "\n\t":
                        v3 -= (v3 / 3.14159)
                    case "foo bar":
                        if true {
                            switch 53 {
                            case 128:
                                if (!true) {
                                    v30 := GenS2{F0: -32768, F1: (v24 >> ((2) & 63)), F2: 544}
                                    switch "hello" {
                                    case "hello":
                                        v14 = 65535
                                        v14 += (^(v24))
                                        v6 -= (v6 - v6)
                                        var v31 uint16 = 2
                                        if false {
                                            if false {
                                                if true {
                                                    if false {
                                                        fmt.Println(v3)
                                                        switch (fi19 & (fi18 + fi18)) {
                                                        case 2147483647:
                                                            v6 = v6
                                                        }
                                                    }
                                                    v14 -= v24
                                                    fmt.Println(v31, fi19)
                                                } else {
                                                    if (true || false) {
                                                        fmt.Println(v31, v24)
                                                        var v32 int8 = 26
                                                        _ = v32
                                                    }
                                                    v33 := GenS4{F0: 42, F1: "hello", F2: 970}
                                                    var v34 int32 = 556
                                                    _ = v33
                                                    _ = v34
                                                }
                                                v31 += (v31 ^ v31)
                                            } else {
                                                var v35 float64 = (((((v6 + v6) * v6) + v6) * v6) + ((0.1 * (3.14159 * 0.0)) + (v6 + (v6 - 1.5))))
                                                fmt.Println(v3, v35, fi19)
                                                var v36 int8 = 100
                                                _ = v35
                                                _ = v36
                                            }
                                        }
                                        v31 = (^(v31))
                                        v6 = v6
                                        v14 += v14
                                        v37 := GenS5{F0: (fi19 ^ fi18)}
                                        switch "\n\t" {
                                        case "a":
                                            v6 = 1.5
                                        default:
                                            fmt.Println(fi18, v14)
                                        }
                                        _ = v31
                                        _ = v37
                                    case "":
                                        v24 += (^(v14))
                                    default:
                                        switch "\n\t" {
                                        case "foo bar":
                                            fmt.Println(fi18, v24)
                                        default:
                                            v6 = 1e10
                                        }
                                    }
                                    _ = v30
                                } else {
                                    v3 -= 1e10
                                }
                            case 42:
                                fmt.Println(fi19, v14)
                            case 32767:
                                v14 += (^(v14))
                                v14 = (^(v24))
                            default:
                                v6 += v6
                            }
                            var v38 uint64 = 42
                            if (381 < 1) {
                                v39 := GenS4{F0: 0, F1: "hello", F2: 876}
                                v24 = (v14 - v24)
                                switch fi19 {
                                case -324:
                                    delete(v4, v14)
                                case -346:
                                    v24 -= v14
                                case 32767:
                                    v38 += v38
                                default:
                                    if false {
                                        fmt.Println(fi18, v6, v3)
                                    } else {
                                        v14 = (v24 + v14)
                                    }
                                }
                                var v40 uint = 1
                                v40 = v40
                                var v41 uint32 = 42
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                            _ = v38
                        }
                    }
                }
                _ = v24
            }
            _ = fi19
        }
        switch "\n\t" {
        case "foo bar":
            var v42 int32 = 1
            _ = v42
        case "hello":
            if (!(("" + ("" + "hello")) <= ("a" + ("" + (("foo bar" + "foo bar") + "\n\t"))))) {
                if false {
                    v43 := GenS5{F0: fi18}
                    var v44 int8 = 1
                    switch int(v44) {
                    case -1:
                        fmt.Println(fi18, v14, v44)
                    case 9223372036854775807:
                        var v45 uint16 = 386
                        _ = v45
                    }
                    _ = v43
                    _ = v44
                }
                var v46 uint = 397
                if (((1 != 10) || (153 <= -237)) && (("\n\t" <= "hello") || true)) {
                    v47 := GenS2{F0: 122, F1: 909, F2: 152}
                    v14 += v14
                    _ = v47
                } else {
                    if (472 != 868) {
                        var v48 float64 = (0.1 * v6)
                        v49 := GenS3{F0: 2}
                        _ = v48
                        _ = v49
                    }
                    var v50 int16 = 0
                    _ = v50
                }
                _ = v46
            }
        case "\n\t":
            v14 = (v14 ^ v14)
        default:
            if len(v1) > 0 {
                v1[0] = ""
            }
        }
        _ = fi18
    }
    v51 := GenS6{F0: 127}
    var v52 float32 = -2.25
    if false {
        var v53 int32 = -1
        v54 := GenS2{F0: -258, F1: 65535, F2: 0}
        _ = v53
        _ = v54
    } else {
        v55 := GenS3{F0: 1}
        for fi56 := 0; fi56 < 2; fi56++ {
            var v57 uint64 = 1
            var v58 uint16 = 32767
            _ = fi56
            _ = v57
            _ = v58
        }
        switch "" {
        case "":
            for fi59 := 0; fi59 < 4; fi59++ {
                for fi60 := 0; fi60 < 6; fi60++ {
                    fmt.Println(fi60)
                    var v61 bool = (false || ((fi59 >> ((217) & 63)) <= (fi59 | 679)))
                    _ = fi60
                    _ = v61
                }
                _ = fi59
            }
        case "a":
            if (!(2 < 2)) {
                var v62 int64 = 833
                _ = v62
            }
        case "\n\t":
            v63 := GenS7{F0: -2147483648, F1: 0, F2: -1}
            _ = v63
        }
        _ = v55
    }
    v3 += v52
    for fi64 := 0; fi64 < 10; fi64++ {
        var v65 int8 = 127
        _ = fi64
        _ = v65
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
