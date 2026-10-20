package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int32; F1 int32; F2 int }
func genFn0(p0 uint64, p1 int) int {
    var v4 uint32 = 2147483648
    v5 := GenS1{F0: 127, F1: 0, F2: 0}
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 79, F1: 0, F2: 392}
        v8 := []int{(p1 + (p1 * 1)), -823}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 130, F1: -9223372036854775808, F2: 0}
    if false {
        v10 := []int{(p1 - p1), p1, 3, (p1 % ((p1) | 1))}
        _ = v10
    } else {
        if (true && (21 >= 89)) {
            switch (p1 / ((p1) | 1)) {
            case -44:
                var v11 int16 = -1
                _ = v11
            case -1:
                if ((v4 | v4) <= (^(v4))) {
                    if false {
                        var v12 uint32 = v4
                        p0 = 4294967295
                        switch (p1 / (((-(p1))) | 1)) {
                        case -897:
                            if (-1 >= 561) {
                                v13 := []int{p1, 4294967295, (p1 % (((p1 ^ (p1 % (((p1 >> ((2) & 63))) | 1)))) | 1)), p1}
                                p0 = 18446744073709551615
                                v14 := GenS1{F0: 1, F1: -9223372036854775808, F2: 2}
                                _ = v13
                                _ = v14
                            } else {
                                for fi15 := 0; fi15 < 16; fi15++ {
                                    v16 := GenS2{F0: "a", F1: 127}
                                    var v17 uint64 = 0
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                }
                                var v18 int8 = 1
                                switch "hello" {
                                case "hello":
                                    switch (^(p1)) {
                                    case 827:
                                        v19 := []int{}
                                        _ = v19
                                    case 9223372036854775807:
                                        var v20 uint64 = (p0 * p0)
                                        _ = v20
                                    }
                                case "a":
                                    v21 := []int{}
                                    _ = v21
                                case "foo bar":
                                    var v22 int = -1
                                    _ = v22
                                }
                                _ = v18
                            }
                        case 2147483647:
                            for fi23 := 0; fi23 < 18; fi23++ {
                                if (0 > 100) {
                                    v24 := []int{(fi23 ^ fi23)}
                                    var v25 uint64 = (^(p0))
                                    _ = v24
                                    _ = v25
                                } else {
                                    p1 = 676
                                }
                                if (!false) {
                                    var v26 uint8 = uint8(p0)
                                    _ = v26
                                }
                                if ((1 >= 65535) || true) {
                                    v27 := GenS1{F0: 42, F1: 1, F2: 0}
                                    _ = v27
                                }
                                _ = fi23
                            }
                        default:
                            for fi28 := 0; fi28 < 7; fi28++ {
                                for fi29 := 0; fi29 < 2; fi29++ {
                                    v30 := []int{42, (-(p1)), fi28}
                                    _ = fi29
                                    _ = v30
                                }
                                if false {
                                    v31 := GenS3{F0: "hello", F1: 9223372036854775807}
                                    switch "a" {
                                    case "":
                                        p0 -= (p0 << ((uint(v4)) & 63))
                                    default:
                                        fmt.Println(fi28)
                                    }
                                    _ = v31
                                }
                                _ = fi28
                            }
                        }
                        _ = v12
                    } else {
                        v32 := GenS2{F0: "hello", F1: -128}
                        if true {
                            v33 := GenS4{F0: (p0 | p0), F1: 156, F2: 255}
                            fmt.Println(p0, v4)
                            p1 = 32768
                            _ = v33
                        }
                        _ = v32
                    }
                } else {
                    v34 := GenS5{F0: 0, F1: ""}
                    fmt.Println(v4)
                    var v35 uint16 = uint16(v4)
                    _ = v34
                    _ = v35
                }
            case -132:
                v4 -= 42
            default:
                if ((299 < 2) || true) {
                    p0 -= p0
                    switch p1 {
                    case 542:
                        v36 := GenS3{F0: ("" + "\n\t"), F1: 42}
                        _ = v36
                    case -1:
                        fmt.Println(p0, p0, p0)
                    default:
                        v4 = (v4 ^ 0)
                    }
                    fmt.Println(p1)
                }
            }
            var v37 uint32 = v4
            _ = v37
        } else {
            v38 := []int{(-(p1))}
            for fi39 := 0; fi39 < 15; fi39++ {
                if false {
                    switch "foo bar" {
                    case "a":
                        for fi40 := 0; fi40 < 4; fi40++ {
                            var v41 int16 = 1000
                            _ = fi40
                            _ = v41
                        }
                    }
                    _ = len(v38)
                }
                _ = fi39
            }
            _ = v38
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v9
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    v3 := []int{}
    for fi42 := 0; fi42 < 20; fi42++ {
        var v43 int = fi42
        v43 = v43
        if true {
            switch genFn0(676, (v43 >> ((589) & 63))) {
            case -1:
                var v44 uint = 255
                _ = v44
            }
            v43 = v43
            if len(v3) > 0 {
                v3[0] = -1
            }
        }
        _ = fi42
        _ = v43
    }
    var v45 float64 = ((0.1 / -2.25) + 3.14159)
    v2[("a" + "hello")] = (true && true)
    v45 = (v45 + (1.5 * -2.25))
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v45)
    case "a":
        for fi46 := 0; fi46 < 20; fi46++ {
            var v47 uint32 = 42
            v48 := v2[("" + "")]
            for fi49 := 0; fi49 < 15; fi49++ {
                switch "\n\t" {
                case "hello":
                    if false {
                        var v50 float32 = 0.1
                        var v51 int64 = 9223372036854775807
                        _ = v50
                        _ = v51
                    } else {
                        var v52 uint16 = 42
                        v53 := GenS6{F0: -901, F1: 2147483647, F2: fi46}
                        _ = v52
                        _ = v53
                    }
                case "\n\t":
                    switch ((("foo bar" + "\n\t") + "") + ("\n\t" + "")) {
                    case "":
                        var v54 int64 = 1
                        _ = v54
                    case "\n\t":
                        var v55 uint16 = 641
                        _ = v55
                    default:
                        v56 := GenS1{F0: 231, F1: -221, F2: 42}
                        _ = v56
                    }
                    v48 = true
                    if v48 {
                        v57 := GenS3{F0: "hello", F1: 42}
                        if (!true) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (18446744073709551615 < 392) {
                                v48 = (!(1 < 42))
                            } else {
                                var v58 uint32 = v47
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v58
                            }
                            v45 = 1.5
                        } else {
                            switch "" {
                            case "":
                                v48 = (1 == 9223372036854775807)
                            }
                            v47 += (v47 >> ((65535) & 63))
                            v47 = (v47 / (((v47 | (v47 + (^(v47))))) | 1))
                        }
                        _ = v57
                    } else {
                        _ = len(v3)
                        _ = len(v3)
                        switch fi46 {
                        case 32767:
                            delete(v2, "")
                        case 1:
                            _ = len(v3)
                        case -534:
                            switch (fi46 ^ fi46) {
                            case 128:
                                v3 = append(v3, 128)
                            case 32768:
                                var v59 uint = 324
                                _ = v59
                            case 127:
                                fmt.Println(v45, fi49, v48)
                            default:
                                var v60 uint64 = 0
                                _ = v60
                            }
                        default:
                            v47 = (^(v47))
                        }
                    }
                    if false {
                        fmt.Println(fi46)
                        v61 := GenS1{F0: 42, F1: 701, F2: 634}
                        v62 := GenS1{F0: 128, F1: 1, F2: 828}
                        _ = v61
                        _ = v62
                    }
                case "":
                    if v48 {
                        v47 += (v47 << ((0) & 63))
                        v63 := GenS4{F0: 289, F1: 1, F2: 65535}
                        _ = v63
                    }
                }
                var v64 uint32 = 2147483648
                var v65 uint = 4294967295
                _ = fi49
                _ = v64
                _ = v65
            }
            _ = fi46
            _ = v47
            _ = v48
        }
    }
    if len(v3) > 0 {
        v3[0] = -796
    }
    v3 = append(v3, -317)
    v66 := GenS3{F0: "\n\t", F1: -9223372036854775808}
    fmt.Println(v45)
    if (0 >= 43) {
        if (true && ("foo bar" >= ("\n\t" + (("" + "\n\t") + "hello")))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v67 uint64 = 18446744073709551615
            _ = v67
        }
    } else {
        v45 += (((v45 - v45) * v45) * 1.5)
        _ = len(v2)
    }
    v45 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v66
}
