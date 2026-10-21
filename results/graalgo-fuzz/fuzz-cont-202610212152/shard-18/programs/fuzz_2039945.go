package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int32; F1 uint; F2 uint32 }
type GenS6 struct { F0 bool; F1 uint32 }
func genFn0(p0 bool, p1 bool, p2 int) uint16 {
    if (!(p2 != (p2 | (p2 >> ((65535) & 63))))) {
        v5 := GenS1{F0: "foo bar", F1: -9223372036854775808}
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := []int{(fi6 & p2), (fi6 << ((42) & 63)), p2, (fi6 - (p2 ^ -510))}
            for fi8 := 0; fi8 < 1; fi8++ {
                if p0 {
                    if (792 != -1) {
                        if (!true) {
                            var v9 uint16 = 721
                            v9 = v9
                            switch "" {
                            case "":
                                v10 := GenS1{F0: "a", F1: -635}
                                _ = v10
                            case "foo bar":
                                v11 := GenS1{F0: ((("foo bar" + ("\n\t" + (("\n\t" + "") + "\n\t"))) + "\n\t") + ((("hello" + "a") + "a") + ("\n\t" + ("" + "foo bar")))), F1: 794}
                                switch (("hello" + "a") + "") {
                                case "foo bar":
                                    var v12 uint = 42
                                    _ = v12
                                case "\n\t":
                                    switch (-(p2)) {
                                    case 32767:
                                        var v13 float32 = 3.14159
                                        _ = v13
                                    case -2147483648:
                                        p2 = (fi8 / (((-(fi6))) | 1))
                                    default:
                                        if (-228 >= 0) {
                                            p2 = 32767
                                            switch "foo bar" {
                                            case "foo bar":
                                                if (42 > 4294967295) {
                                                    var v14 float64 = 1.5
                                                    _ = v14
                                                } else {
                                                    if true {
                                                        var v15 int8 = -128
                                                        _ = v15
                                                    }
                                                    switch "\n\t" {
                                                    case "hello":
                                                        if len(v7) > 0 {
                                                            v7[0] = 508
                                                        }
                                                    case "foo bar":
                                                        p0 = p0
                                                    case "\n\t":
                                                        var v16 bool = false
                                                        _ = v16
                                                    default:
                                                        p1 = (3.14159 > 1.5)
                                                    }
                                                    if p0 {
                                                        if p1 {
                                                            var v17 uint32 = 2147483648
                                                            var v18 int16 = -932
                                                            v19 := GenS2{F0: (v17 & (v17 >> ((62) & 63)))}
                                                            _ = v17
                                                            _ = v18
                                                            _ = v19
                                                        }
                                                        p1 = p0
                                                    }
                                                }
                                            default:
                                                var v20 int32 = -593
                                                _ = v20
                                            }
                                            var v21 int64 = -764
                                            _ = v21
                                        } else {
                                            v9 = 32767
                                            switch fi8 {
                                            case -2147483648:
                                                fmt.Println(v9)
                                            case 42:
                                                switch ((("" + "hello") + "foo bar") + "\n\t") {
                                                case "a":
                                                    switch (p2 % (((fi8 << ((941) & 63))) | 1)) {
                                                    case 136:
                                                        p2 += 42
                                                    default:
                                                        var v22 uint32 = 942
                                                        _ = v22
                                                    }
                                                case "\n\t":
                                                    var v23 int8 = 42
                                                    _ = v23
                                                default:
                                                    switch ("a" + "a") {
                                                    case "\n\t":
                                                        v7 = append(v7, 127)
                                                    default:
                                                        v24 := GenS2{F0: 0}
                                                        _ = v24
                                                    }
                                                }
                                            case -375:
                                                if true {
                                                    fmt.Println(p2, v9)
                                                }
                                            }
                                            v25 := GenS1{F0: ("a" + (((("foo bar" + "hello") + ("foo bar" + (("" + ("" + "\n\t")) + (("a" + "") + "\n\t")))) + "hello") + "hello")), F1: int64(v9)}
                                            if (p0 || p0) {
                                                var v26 uint64 = 42
                                                _ = v26
                                            } else {
                                                v27 := GenS1{F0: "a", F1: 1}
                                                switch "hello" {
                                                case "hello":
                                                    p2 -= (p2 | (-(fi8)))
                                                    var v28 float64 = (0.0 * (1.5 * (3.14159 * -2.25)))
                                                    v9 = 1
                                                    _ = v28
                                                case "foo bar":
                                                    p0 = (!(127 > 0))
                                                case "\n\t":
                                                    v29 := GenS1{F0: "a", F1: int64(fi6)}
                                                    _ = v29
                                                }
                                                v30 := GenS3{F0: 0, F1: 0}
                                                _ = v27
                                                _ = v30
                                            }
                                            var v31 int64 = -1
                                            var v32 string = "foo bar"
                                            p0 = (p0 || p1)
                                            switch (fi6 - (p2 / ((fi6) | 1))) {
                                            case 220:
                                                fmt.Println(v32, p1)
                                            default:
                                                if (0 >= 189) {
                                                    v31 = v31
                                                } else {
                                                    fmt.Println(p2)
                                                }
                                            }
                                            v9 = v9
                                            _ = v25
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                case "":
                                    fmt.Println(fi8, fi8, fi8)
                                default:
                                    var v33 uint32 = 4294967295
                                    _ = v33
                                }
                                switch p2 {
                                case -9223372036854775808:
                                    v34 := GenS2{F0: 2147483648}
                                    _ = v34
                                case 1:
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                case 127:
                                    if true {
                                        switch "" {
                                        case "a":
                                            v35 := GenS3{F0: 0, F1: 42}
                                            _ = v35
                                        case "\n\t":
                                            if (((((1 <= 2147483647) || false) || p1) && (42 == 49)) && ("a" <= "hello")) {
                                                if (p1 && false) {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        var v36 int = p2
                                                        _ = v36
                                                    }
                                                    if len(v7) > 0 {
                                                        v7[0] = fi6
                                                    }
                                                    var v37 int = fi8
                                                    if p0 {
                                                        fmt.Println(v9)
                                                    } else {
                                                        switch -2147483648 {
                                                        case -989:
                                                            v9 += v9
                                                        }
                                                    }
                                                    v38 := GenS2{F0: 4294967295}
                                                    _ = v37
                                                    _ = v38
                                                }
                                                fmt.Println(p1, fi8, fi8)
                                            } else {
                                                var v39 float32 = 0.1
                                                if ((613 >= 65535) && false) {
                                                    v7 = append(v7, fi6)
                                                    if (!true) {
                                                        var v40 uint64 = 42
                                                        switch "" {
                                                        case "\n\t":
                                                            var v41 int = (fi6 << ((1) & 63))
                                                            _ = v41
                                                        case "hello":
                                                            v7 = append(v7, (p2 - 0))
                                                        case "a":
                                                            fmt.Println(v9)
                                                        }
                                                        _ = v40
                                                    } else {
                                                        var v42 uint32 = uint32(v9)
                                                        if (-900 > 329) {
                                                            var v43 uint64 = 18446744073709551615
                                                            _ = v43
                                                        }
                                                        switch (fi8 & fi6) {
                                                        case 0:
                                                            var v44 int = 128
                                                            _ = v44
                                                        default:
                                                            v45 := GenS2{F0: (v42 * v42)}
                                                            _ = v45
                                                        }
                                                        var v46 float32 = v39
                                                        switch fi6 {
                                                        case -1:
                                                            switch p2 {
                                                            case 32767:
                                                                if len(v7) > 0 {
                                                                    v7[0] = (fi6 + p2)
                                                                }
                                                            default:
                                                                _ = len(v7)
                                                            }
                                                        default:
                                                            p0 = (!(!(!(18446744073709551615 != 873))))
                                                        }
                                                        _ = v42
                                                        _ = v46
                                                    }
                                                }
                                                var v47 float32 = float32(0.1)
                                                _ = v39
                                                _ = v47
                                            }
                                        case "foo bar":
                                            p2 -= 127
                                        default:
                                            var v48 int16 = 1000
                                            _ = v48
                                        }
                                    } else {
                                        if len(v7) > 0 {
                                            v7[0] = fi8
                                        }
                                        p2 = 32768
                                        var v49 int = (fi8 + fi8)
                                        _ = v49
                                    }
                                default:
                                    v50 := GenS3{F0: 0, F1: 199}
                                    _ = v50
                                }
                                _ = v11
                            case "a":
                                var v51 int32 = -1
                                _ = v51
                            }
                            _ = v9
                        }
                    } else {
                        var v52 float32 = 1e10
                        _ = v52
                    }
                    var v53 uint8 = 42
                    if len(v7) > 0 {
                        v7[0] = -507
                    }
                    _ = v53
                } else {
                    v54 := GenS2{F0: 707}
                    _ = v54
                }
                var v55 int64 = -930
                _ = fi8
                _ = v55
            }
            _ = fi6
            _ = v7
        }
        var v56 int16 = 290
        _ = v5
        _ = v56
    }
    v57 := GenS2{F0: 98}
    var v58 uint8 = 128
    v59 := []int{p2}
    switch "hello" {
    case "\n\t":
        if len(v59) >= 2 {
            v59 = v59[1:len(v59)]
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v57
    _ = v58
    _ = v59
    return 352
}
func genFn1(p0 int8, p1 uint32, p2 uint8) int32 {
    p0 += 81
    for fi60 := 0; fi60 < 3; fi60++ {
        switch (fi60 * (fi60 & fi60)) {
        case -1:
            switch "hello" {
            case "":
                v61 := []int{(fi60 + 823), (fi60 ^ 87), (fi60 + fi60)}
                _ = v61
            default:
                var v62 float32 = (3.14159 * (1.5 / 1e10))
                _ = v62
            }
        case -408:
            switch "foo bar" {
            case "\n\t":
                p1 = (p1 << ((4294967295) & 63))
            }
        }
        _ = fi60
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 uint32 = 807
    var v2 int = 2147483647
    v3 := []uint32{}
    v4 := make(map[uint32]bool)
    v63 := GenS4{F0: 643, F1: 1000}
    if true {
        for fi64 := 0; fi64 < 13; fi64++ {
            var v65 float64 = -2.25
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi64
            _ = v65
        }
        for fi66 := 0; fi66 < 1; fi66++ {
            fmt.Println(v1, v1)
            for fi67 := 0; fi67 < 9; fi67++ {
                if false {
                    v4[v1] = true
                    var v68 uint = 65535
                    _ = v68
                } else {
                    fmt.Println(fi67, v1, v1)
                    var v69 int16 = 42
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                    v2 -= fi66
                    v70 := GenS5{F0: -406, F1: 879, F2: v1}
                    v69 = v69
                    _ = len(v3)
                    var v71 uint = 255
                    _ = v69
                    _ = v70
                    _ = v71
                }
                switch fi67 {
                case -9223372036854775808:
                    v1 -= (^(v1))
                case -1:
                    v2 = 32767
                }
                _ = fi67
            }
            for fi72 := 0; fi72 < 19; fi72++ {
                if (true || true) {
                    if true {
                        var v73 int32 = 749
                        v1 -= v1
                        if (65535 <= 1) {
                            v1 = 42
                            v2 = 121
                            var v74 float64 = 1e10
                            var v75 uint = 551
                            var v76 string = ""
                            _ = v74
                            _ = v75
                            _ = v76
                        } else {
                            if len(v3) > 0 {
                                v3[0] = (v1 - v1)
                            }
                            _ = len(v3)
                            fmt.Println(v73, v2, v1)
                        }
                        _ = v73
                    }
                    v1 -= (v1 * 95)
                }
                v2 += (v2 ^ -2147483648)
                v2 -= fi66
                _ = fi72
            }
            _ = fi66
        }
    }
    var v77 uint16 = 2
    v77 = (v77 + (v77 * v77))
    var v78 int = (v2 << ((4294967295) & 63))
    v79 := GenS6{F0: (232 != 178), F1: v1}
    v77 += (^(v77))
    _ = len(v3)
    var v80 uint8 = 127
    var v81 bool = (128 > (v80 * 128))
    fmt.Println(v1)
    for fi82 := 0; fi82 < 15; fi82++ {
        for fi83 := 0; fi83 < 16; fi83++ {
            v81 = v81
            _ = fi83
        }
        _ = fi82
    }
    v84 := GenS4{F0: 420, F1: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v63
    _ = v77
    _ = v78
    _ = v79
    _ = v80
    _ = v81
    _ = v84
}
