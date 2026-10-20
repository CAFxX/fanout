package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 int8; F2 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64; F1 int64 }
type GenS9 struct { F0 string; F1 uint64 }
type GenS10 struct { F0 uint64; F1 bool; F2 int32 }
func genFn0(p0 int, p1 uint16) int16 {
    p1 += (p1 >> ((18446744073709551615) & 63))
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 uint32 = 2147483647
        switch (p0 / (((fi5 + fi5)) | 1)) {
        case 719:
            for fi7 := 0; fi7 < 12; fi7++ {
                v8 := GenS1{F0: true}
                _ = fi7
                _ = v8
            }
        }
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS1{F0: false}
            v6 = v6
            _ = fi9
            _ = v10
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 bool = ((2 != 2) && false)
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    var v4 int32 = 42
    _ = len(v3)
    var v11 float32 = 1.5
    v11 = v11
    for fi12 := 0; fi12 < 2; fi12++ {
        switch fi12 {
        case 32767:
            switch "" {
            case "":
                if true {
                    if v1 {
                        v11 -= float32((3.14159 * (0.1 / 1e10)))
                        v4 += v4
                        v13 := []int{}
                        _ = v13
                    }
                    if v1 {
                        var v14 bool = (1e10 != (1e10 - 3.14159))
                        fmt.Println(v11, fi12, v14)
                        for fi15 := 0; fi15 < 4; fi15++ {
                            switch 992 {
                            case -9223372036854775808:
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    v14 = ((!false) || v1)
                                case "a":
                                    var v16 uint = 294
                                    _ = v16
                                case "\n\t":
                                    v14 = ((false || (872 >= 258)) || true)
                                default:
                                    switch fi12 {
                                    case -2147483648:
                                        var v17 int16 = 42
                                        _ = v17
                                    case 32767:
                                        if true {
                                            var v18 uint8 = 200
                                            _ = v18
                                        }
                                    case 0:
                                        v4 -= 42
                                    default:
                                        switch 32767 {
                                        case 9223372036854775807:
                                            delete(v3, "")
                                        default:
                                            if ((^(v4)) >= v4) {
                                                v19 := []int{}
                                                fmt.Println(v1, v1)
                                                v11 = v11
                                                _ = v19
                                            }
                                        }
                                    }
                                    var v20 uint64 = 1
                                    if ((!v14) && (false || (91 > 128))) {
                                        fmt.Println(v14, v1, fi15)
                                        v1 = v14
                                        if true {
                                            v1 = (v11 >= (1.5 + ((0.0 * v11) + float32(3.14159))))
                                            v11 -= 1.5
                                        }
                                    }
                                    _ = v20
                                }
                            case 32767:
                                v21 := GenS2{F0: 154}
                                _ = v21
                            case 42:
                                v22 := []int{fi12, 4294967295, int(v4), (fi12 * fi12)}
                                _ = v22
                            default:
                                fmt.Println(fi15, fi12, fi12)
                            }
                            var v23 int8 = -1
                            _ = len(v3)
                            v24 := []int{(fi15 | fi12), fi15}
                            _ = fi15
                            _ = v23
                            _ = v24
                        }
                        _ = v14
                    } else {
                        var v25 float32 = v11
                        _ = v25
                    }
                }
            default:
                v4 = 417
            }
        case 32768:
            v26 := GenS3{F0: 32767}
            _ = v26
        case 2147483647:
            fmt.Println(v1)
        default:
            v27 := GenS4{F0: 2147483648, F1: 1, F2: false}
            _ = v27
        }
        _ = fi12
    }
    var v28 uint64 = 4294967295
    for fi29 := 0; fi29 < 2; fi29++ {
        v30 := GenS5{F0: 411}
        v31 := GenS6{F0: v28, F1: genFn0((fi29 / ((fi29) | 1)), uint16(v28))}
        _ = fi29
        _ = v30
        _ = v31
    }
    v11 -= (v11 * (((v11 * 3.14159) + 3.14159) / 1e10))
    var v32 int16 = 1000
    v28 -= (v28 | v28)
    v33 := GenS7{F0: v4}
    switch "\n\t" {
    case "":
        switch 32767 {
        case 432:
            if v1 {
                fmt.Println(v32, v11, v1)
                if ((float64((v11 * 1.5)) * (0.1 / 0.1)) >= 1.5) {
                    var v34 uint16 = 617
                    v35 := GenS3{F0: 1000}
                    if v1 {
                        if (0.1 > float64(v11)) {
                            for fi36 := 0; fi36 < 10; fi36++ {
                                v37 := v3["hello"]
                                v38 := GenS5{F0: fi36}
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                            _ = len(v3)
                        } else {
                            if (v1 || v1) {
                                if (!(!(v11 >= v11))) {
                                    var v39 uint16 = v34
                                    delete(v2, "")
                                    if true {
                                        fmt.Println(v34, v34, v32)
                                    } else {
                                        v32 += v32
                                        v39 -= (v39 << ((786) & 63))
                                        var v40 float32 = -2.25
                                        _ = v40
                                    }
                                    _ = v39
                                }
                                if ((!v1) && (-9223372036854775808 != -191)) {
                                    switch 1 {
                                    case 32767:
                                        for fi41 := 0; fi41 < 4; fi41++ {
                                            v42 := GenS8{F0: 9223372036854775807, F1: 9223372036854775807}
                                            _ = fi41
                                            _ = v42
                                        }
                                    default:
                                        var v43 int16 = (v32 << ((419) & 63))
                                        _ = v43
                                    }
                                    v44 := GenS9{F0: ("" + "hello"), F1: v28}
                                    _ = v44
                                } else {
                                    for fi45 := 0; fi45 < 19; fi45++ {
                                        if (v34 >= uint16(v28)) {
                                            switch ("hello" + ("hello" + "\n\t")) {
                                            case "a":
                                                var v46 uint64 = v28
                                                _ = v46
                                            case "hello":
                                                var v47 uint64 = 947
                                                _ = v47
                                            default:
                                                for fi48 := 0; fi48 < 14; fi48++ {
                                                    fmt.Println(v28, v32, v28)
                                                    v32 -= v32
                                                    fmt.Println(v32, fi45, v28)
                                                    var v49 int8 = 42
                                                    _ = fi48
                                                    _ = v49
                                                }
                                            }
                                            for fi50 := 0; fi50 < 10; fi50++ {
                                                if true {
                                                    v51 := []int{(fi45 % (((-(fi45))) | 1)), fi50}
                                                    _ = v51
                                                } else {
                                                    var v52 int64 = -898
                                                    v28 += (^(v28))
                                                    _ = v52
                                                }
                                                if (false || v1) {
                                                    v53 := []int{-2147483648, (^(fi45)), (-(fi50)), fi45}
                                                    _ = v53
                                                } else {
                                                    v34 = 123
                                                    v54 := []int{(fi45 & 1)}
                                                    _ = v54
                                                }
                                                _ = fi50
                                            }
                                            if ((3.14159 != v11) && ("a" > "")) {
                                                v55 := []int{(fi45 & (fi45 - (fi45 / ((fi45) | 1))))}
                                                if (fi45 >= 1) {
                                                    for fi56 := 0; fi56 < 8; fi56++ {
                                                        v4 = (v4 << ((927) & 63))
                                                        _ = fi56
                                                    }
                                                }
                                                switch "" {
                                                case "":
                                                    for fi57 := 0; fi57 < 16; fi57++ {
                                                        v32 -= v32
                                                        _ = fi57
                                                    }
                                                }
                                                _ = v55
                                            }
                                        } else {
                                            v58 := []int{fi45, (fi45 * (-(fi45))), (^(fi45))}
                                            _ = v58
                                        }
                                        _ = fi45
                                    }
                                    var v59 float64 = 1e10
                                    _ = v59
                                }
                                var v60 int64 = 42
                                _ = v60
                            } else {
                                fmt.Println(v32, v1, v11)
                            }
                        }
                        switch 32767 {
                        case -1:
                            var v61 float64 = 0.0
                            _ = v61
                        case 32768:
                            var v62 int = 707
                            _ = v62
                        }
                    }
                    _ = v34
                    _ = v35
                }
                var v63 int32 = v4
                _ = v63
            } else {
                for fi64 := 0; fi64 < 1; fi64++ {
                    var v65 uint = 1
                    var v66 uint32 = 42
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                for fi67 := 0; fi67 < 19; fi67++ {
                    var v68 uint16 = 2
                    v4 = (v4 ^ v4)
                    v69 := GenS6{F0: 312, F1: genFn0(fi67, (^(v68)))}
                    _ = fi67
                    _ = v68
                    _ = v69
                }
                v70 := v3[("a" + "foo bar")]
                _ = v70
            }
        }
    case "a":
        var v71 uint = 2
        _ = v71
    case "foo bar":
        v3["hello"] = "hello"
    default:
        var v72 int64 = 0
        _ = v72
    }
    v11 = v11
    switch "" {
    case "":
        if false {
            for fi73 := 0; fi73 < 15; fi73++ {
                var v74 bool = v1
                _ = fi73
                _ = v74
            }
            var v75 float64 = (-2.25 - (0.0 - (0.1 + (1e10 + ((1e10 - 0.0) + -2.25)))))
            _ = v75
        }
    case "\n\t":
        delete(v3, "")
    case "foo bar":
        var v76 uint32 = 2
        _ = v76
    default:
        v77 := v3[(("a" + ("a" + "hello")) + "")]
        _ = v77
    }
    for fi78 := 0; fi78 < 16; fi78++ {
        _ = len(v3)
        _ = fi78
    }
    var v79 uint8 = 1
    for fi80 := 0; fi80 < 12; fi80++ {
        v81 := GenS8{F0: 610, F1: 42}
        switch "\n\t" {
        case "hello":
            v82 := GenS10{F0: 4294967295, F1: true, F2: v4}
            _ = v82
        default:
            if (false || v1) {
                v83 := []int{fi80}
                v84 := GenS4{F0: 42, F1: 100, F2: v1}
                _ = v83
                _ = v84
            } else {
                var v85 uint8 = v79
                _ = v85
            }
        }
        _ = fi80
        _ = v81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v79
}
