package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int32; F2 bool }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int32; F1 uint64 }
func genFn0() string {
    var v4 uint64 = 18446744073709551615
    fmt.Println(v4)
    for fi5 := 0; fi5 < 16; fi5++ {
        if (false || false) {
            for fi6 := 0; fi6 < 16; fi6++ {
                var v7 uint = 162
                fmt.Println(fi6, v7, fi6)
                var v8 int32 = -56
                _ = fi6
                _ = v7
                _ = v8
            }
        } else {
            v9 := GenS1{F0: fi5}
            _ = v9
        }
        v4 -= v4
        _ = fi5
    }
    _ = v4
    return "foo bar"
}
func genFn1(p0 int16) uint32 {
    p0 -= (p0 | p0)
    p0 = -32768
    v10 := []int{376, 2147483647, 399}
    var v11 int8 = -114
    switch "" {
    case "\n\t":
        p0 = p0
    default:
        for fi12 := 0; fi12 < 3; fi12++ {
            if (1 == 65535) {
                for fi13 := 0; fi13 < 5; fi13++ {
                    v11 = (v11 & (v11 & v11))
                    _ = fi13
                }
                for fi14 := 0; fi14 < 7; fi14++ {
                    v11 = v11
                    if len(v10) > 0 {
                        v10[0] = 127
                    }
                    var v15 uint64 = 1000
                    _ = fi14
                    _ = v15
                }
                var v16 bool = true
                _ = v16
            } else {
                if false {
                    v17 := GenS2{F0: v11, F1: 0, F2: (2147483648 >= 696)}
                    v11 = (v11 >> ((4294967295) & 63))
                    _ = v17
                } else {
                    v18 := GenS2{F0: (v11 | v11), F1: -2147483648, F2: false}
                    _ = v18
                }
                for fi19 := 0; fi19 < 16; fi19++ {
                    if (fi19 < (fi19 >> ((198) & 63))) {
                        var v20 uint8 = 127
                        var v21 uint32 = 2147483648
                        _ = v20
                        _ = v21
                    }
                    switch 32768 {
                    case -2147483648:
                        fmt.Println(fi19)
                    case -736:
                        p0 = (p0 >> ((18446744073709551615) & 63))
                    }
                    p0 -= p0
                    _ = fi19
                }
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
            for fi22 := 0; fi22 < 11; fi22++ {
                v11 = int8(fi22)
                p0 += (p0 - (p0 + (p0 ^ (p0 % ((p0) | 1)))))
                var v23 int8 = (v11 / (((^(v11))) | 1))
                _ = fi22
                _ = v23
            }
            _ = fi12
        }
    }
    _ = p0
    _ = v10
    _ = v11
    return 233
}

func main() {
    v1 := []string{"a", "a"}
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    switch -9223372036854775808 {
    case 2147483647:
        v24 := GenS1{F0: 889}
        _ = v24
    case 42:
        _ = len(v1)
    case 4294967295:
        var v25 uint16 = 1
        _ = v25
    default:
        var v26 int16 = -1
        _ = v26
    }
    if ((!(301 != -90)) || ((762 <= 4294967295) && true)) {
        v27 := v3[""]
        var v28 float32 = 1.5
        switch "a" {
        case "\n\t":
            switch -2147483648 {
            case 1:
                for fi29 := 0; fi29 < 7; fi29++ {
                    v27 = (false && ((v27 || (("a" <= "") && (613 == 80))) && false))
                    v30 := GenS3{F0: (fi29 * 127), F1: v27}
                    v31 := GenS2{F0: 42, F1: -1, F2: v27}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            case 2147483647:
                v28 = v28
            }
        case "a":
            v27 = true
        case "":
            var v32 uint = 0
            _ = v32
        }
        _ = v27
        _ = v28
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        for fi34 := 0; fi34 < 3; fi34++ {
            switch (fi33 - fi33) {
            case -1:
                var v35 int16 = -634
                _ = v35
            case -365:
                var v36 int64 = -196
                _ = v36
            case -194:
                delete(v2, -800)
            }
            v37 := GenS4{F0: -122, F1: 65535}
            _ = fi34
            _ = v37
        }
        var v38 uint64 = 1
        v2[fi33] = 844
        _ = fi33
        _ = v38
    }
    v39 := GenS1{F0: 32767}
    if true {
        v40 := GenS1{F0: -768}
        _ = v40
    }
    for fi41 := 0; fi41 < 19; fi41++ {
        if true {
            switch 425 {
            case -781:
                switch "\n\t" {
                case "foo bar":
                    var v42 uint16 = 2
                    _ = v42
                case "a":
                    fmt.Println(fi41)
                case "":
                    var v43 int8 = -1
                    _ = v43
                default:
                    if (!false) {
                        v44 := GenS4{F0: -127, F1: 0}
                        _ = v44
                    }
                }
            case 128:
                var v45 int64 = -670
                _ = v45
            default:
                var v46 float32 = (-2.25 + 1e10)
                _ = v46
            }
        }
        var v47 uint32 = 0
        if false {
            for fi48 := 0; fi48 < 12; fi48++ {
                var v49 int64 = int64(fi48)
                if (343 == 18446744073709551615) {
                    var v50 uint = 52
                    v49 += 42
                    var v51 float32 = (1e10 / 1.5)
                    _ = v50
                    _ = v51
                }
                _ = fi48
                _ = v49
            }
            switch "\n\t" {
            case "":
                v52 := v3["a"]
                _ = v52
            case "hello":
                v53 := v2[fi41]
                _ = v53
            }
            switch (fi41 / (((fi41 % (((fi41 / ((fi41) | 1))) | 1))) | 1)) {
            case 4294967295:
                switch fi41 {
                case -651:
                    v47 = genFn1(-250)
                case 0:
                    v47 = genFn1(0)
                case 32768:
                    v3[("\n\t" + "hello")] = (!false)
                default:
                    switch fi41 {
                    case 128:
                        for fi54 := 0; fi54 < 4; fi54++ {
                            v55 := GenS5{F0: false, F1: 1}
                            v47 += genFn1(32767)
                            v56 := GenS1{F0: fi41}
                            _ = fi54
                            _ = v55
                            _ = v56
                        }
                    case 1:
                        var v57 uint32 = (v47 ^ (v47 - (v47 ^ 535)))
                        _ = v57
                    }
                }
            case 2147483647:
                for fi58 := 0; fi58 < 13; fi58++ {
                    var v59 float64 = 0.0
                    _ = fi58
                    _ = v59
                }
            case -492:
                var v60 uint64 = 18446744073709551615
                _ = v60
            default:
                v61 := GenS5{F0: ((-(fi41)) >= fi41), F1: 4294967295}
                _ = v61
            }
        }
        _ = fi41
        _ = v47
    }
    var v62 int32 = 0
    v62 += (v62 * v62)
    fmt.Println(v62)
    v3["foo bar"] = true
    switch 127 {
    case 2147483647:
        for fi63 := 0; fi63 < 10; fi63++ {
            v62 = int32(fi63)
            _ = fi63
        }
    case -9223372036854775808:
        var v64 int32 = 297
        _ = v64
    default:
        v65 := GenS5{F0: true, F1: 1}
        _ = v65
    }
    v66 := GenS6{F0: (v62 - (v62 % (((v62 << ((255) & 63))) | 1))), F1: 0}
    v62 = v62
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v62
    _ = v66
}
