package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int16; F1 bool; F2 uint64 }
type GenS4 struct { F0 string; F1 uint8 }
func genFn0(p0 int8, p1 bool, p2 uint64) uint8 {
    v7 := []int{32767, 32768, 0, -9223372036854775808}
    for fi8 := 0; fi8 < 17; fi8++ {
        if false {
            p1 = p1
            p1 = p1
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        switch fi9 {
        case -9223372036854775808:
            fmt.Println(p0, p1, fi9)
        default:
            p2 = 2
        }
        fmt.Println(p1)
        var v10 int64 = -1
        _ = fi9
        _ = v10
    }
    p0 = p0
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            if p1 {
                fmt.Println(p2, fi12)
            }
            var v13 int = (-(fi12))
            var v14 float64 = 1e10
            var v15 int8 = (p0 >> ((0) & 63))
            _ = len(v7)
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        p0 -= (p0 % ((p0) | 1))
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 float32 = 1e10
    var v2 uint8 = 35
    v3 := []string{"foo bar", "hello", "foo bar", "a"}
    v4 := []int{4294967295, 320, 9223372036854775807, 128}
    var v5 bool = true
    var v6 uint32 = 2147483647
    _ = len(v3)
    v2 += 42
    if v5 {
        if (!((!false) || (0 <= -120))) {
            for fi16 := 0; fi16 < 15; fi16++ {
                if (false && (31 > 42)) {
                    if (1 < 42) {
                        for fi17 := 0; fi17 < 3; fi17++ {
                            _ = len(v4)
                            var v18 uint = 2
                            v19 := GenS1{F0: 0}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        v1 -= ((((v1 + v1) / v1) / ((v1 + v1) / 3.14159)) + v1)
                    }
                } else {
                    if false {
                        var v20 uint8 = 42
                        v2 -= (v20 >> ((65535) & 63))
                        _ = v20
                    }
                }
                for fi21 := 0; fi21 < 3; fi21++ {
                    v6 = 4294967295
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    if (v1 <= 1e10) {
                        fmt.Println(v1, v1, v6)
                        var v22 float64 = 0.1
                        var v23 uint16 = 608
                        _ = v22
                        _ = v23
                    }
                    _ = fi21
                }
                _ = fi16
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch 127 {
            case 1:
                for fi24 := 0; fi24 < 1; fi24++ {
                    if (v5 || (!((v2 ^ uint8(v6)) <= genFn0(63, false, 4294967295)))) {
                        if true {
                            if (v6 < v6) {
                                var v25 int64 = 42
                                _ = v25
                            } else {
                                var v26 int8 = -18
                                _ = v26
                            }
                        }
                    }
                    v6 += v6
                    _ = fi24
                }
            }
            var v27 uint32 = v6
            _ = v27
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v28 int64 = 811
    var v29 uint64 = 656
    if ((1.5 == 1.5) && (151 != v2)) {
        var v30 int32 = 601
        var v31 uint16 = 32767
        _ = len(v4)
        _ = v30
        _ = v31
    } else {
        v28 += v28
        var v32 int32 = 1
        _ = v32
    }
    v3 = append(v3, "hello")
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 int64 = v28
        v35 := GenS1{F0: 65535}
        switch 32767 {
        case 145:
            _ = len(v3)
        case -2147483648:
            var v36 int8 = -128
            _ = v36
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println(v5, v5, v29)
    fmt.Println(v6)
    switch 32767 {
    case 1:
        v4 = append(v4, -176)
    case 32767:
        var v37 uint8 = v2
        _ = v37
    case 2147483647:
        for fi38 := 0; fi38 < 3; fi38++ {
            fmt.Println(v29, fi38)
            v29 = (v29 + v29)
            _ = fi38
        }
    }
    v39 := GenS2{F0: 65535, F1: 42}
    v40 := GenS1{F0: 0}
    for fi41 := 0; fi41 < 10; fi41++ {
        switch (fi41 | fi41) {
        case 116:
            v4 = append(v4, -9223372036854775808)
        case 9223372036854775807:
            var v42 int16 = -874
            _ = v42
        default:
            v2 = 47
        }
        for fi43 := 0; fi43 < 13; fi43++ {
            v6 = 0
            v44 := GenS1{F0: uint16(v6)}
            _ = fi43
            _ = v44
        }
        _ = fi41
    }
    for fi45 := 0; fi45 < 15; fi45++ {
        v46 := GenS1{F0: 32767}
        if (v5 && v5) {
            for fi47 := 0; fi47 < 19; fi47++ {
                switch 32767 {
                case 32767:
                    v6 -= (v6 % ((609) | 1))
                case 32768:
                    if ((0.1 / v1) <= (float32(-2.25) * (3.14159 + (v1 * (1.5 / v1))))) {
                        v48 := GenS2{F0: v6, F1: 1}
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                        _ = v48
                    } else {
                        fmt.Println(v1, v6)
                        _ = len(v4)
                    }
                case -200:
                    switch (fi47 | fi45) {
                    case 60:
                        var v49 int8 = 1
                        _ = v49
                    case -722:
                        _ = len(v4)
                    }
                default:
                    v50 := GenS1{F0: 133}
                    _ = v50
                }
                _ = fi47
            }
            for fi51 := 0; fi51 < 16; fi51++ {
                fmt.Println(v6, v1)
                _ = fi51
            }
        } else {
            if v5 {
                v5 = (!(v5 && true))
            } else {
                fmt.Println(fi45, v29, v6)
                for fi52 := 0; fi52 < 12; fi52++ {
                    v2 -= (v2 - (v2 - (v2 - 2)))
                    _ = fi52
                }
                _ = len(v3)
            }
            switch fi45 {
            case -2147483648:
                var v53 int = 1
                _ = v53
            default:
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if v5 {
                        fmt.Println(v28, v2)
                        if true {
                            v1 = 1e10
                            v54 := GenS2{F0: (v6 | (^(v6))), F1: -128}
                            _ = v54
                        } else {
                            switch "foo bar" {
                            case "hello":
                                var v55 int64 = (v28 >> ((18446744073709551615) & 63))
                                _ = v55
                            }
                        }
                    } else {
                        v56 := GenS3{F0: 1, F1: v5, F2: (v29 * (v29 % (((v29 >> ((1) & 63))) | 1)))}
                        v57 := GenS4{F0: ((("a" + "\n\t") + "foo bar") + ("" + "foo bar")), F1: v2}
                        fmt.Println(v6)
                        _ = v56
                        _ = v57
                    }
                } else {
                    var v58 int64 = (v28 * v28)
                    _ = v58
                }
            }
        }
        v59 := GenS1{F0: 42}
        _ = fi45
        _ = v46
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
