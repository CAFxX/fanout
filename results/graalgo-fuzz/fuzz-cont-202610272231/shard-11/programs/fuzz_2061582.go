package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int64; F1 int16; F2 int16 }
type GenS4 struct { F0 uint16 }
func genFn0() int32 {
    var v7 uint32 = 4294967295
    switch 800 {
    case 2147483647:
        v8 := []int{}
        _ = v8
    case -612:
        switch "foo bar" {
        case "foo bar":
            v7 = v7
        }
    case -384:
        if true {
            switch 32767 {
            case 9223372036854775807:
                v9 := []int{128, 127, 32767, -376}
                _ = v9
            default:
                if true {
                    v10 := []int{128, 532}
                    if false {
                        fmt.Println(v7)
                        var v11 int16 = 0
                        _ = v11
                    }
                    v7 = 858
                    _ = v10
                }
            }
        } else {
            v12 := []int{-608, 9223372036854775807, 128}
            _ = v12
        }
    default:
        var v13 int8 = 42
        _ = v13
    }
    v7 = (v7 ^ v7)
    v7 -= (v7 - 2147483648)
    var v14 uint64 = 458
    _ = v14
    _ = v7
    return -863
}
func genFn1(p0 int, p1 uint32) int {
    if true {
        for fi15 := 0; fi15 < 7; fi15++ {
            for fi16 := 0; fi16 < 2; fi16++ {
                p0 = (p0 & p0)
                p0 = (fi16 * fi16)
                _ = fi16
            }
            var v17 float64 = -2.25
            _ = fi15
            _ = v17
        }
        switch p0 {
        case -9223372036854775808:
            for fi18 := 0; fi18 < 10; fi18++ {
                if ((1.5 + (1e10 - float64(-2.25))) >= (3.14159 / 0.1)) {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v20 := []int{}
                        _ = len(v20)
                        p1 += 452
                        _ = fi19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 16; fi21++ {
                        v22 := GenS1{F0: 42, F1: 7, F2: 65535}
                        fmt.Println(fi18, p1)
                        v23 := GenS2{F0: 9223372036854775807, F1: true}
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    v24 := []int{p0, (-(fi18)), -238, (fi18 & (fi18 + fi18))}
                    _ = v24
                } else {
                    switch p0 {
                    case -644:
                        var v25 int32 = -337
                        _ = v25
                    default:
                        fmt.Println(fi18, fi18)
                    }
                }
                _ = fi18
            }
        case 843:
            var v26 float32 = -2.25
            _ = v26
        case 42:
            fmt.Println(p1, p0)
        }
        v27 := GenS1{F0: 0, F1: 42, F2: 324}
        _ = v27
    } else {
        if true {
            v28 := GenS3{F0: -9223372036854775808, F1: int16(p1), F2: 32767}
            p1 += p1
            _ = v28
        }
        v29 := GenS4{F0: 2}
        _ = v29
    }
    v30 := GenS1{F0: 466, F1: 606, F2: 971}
    var v31 int64 = -591
    _ = p0
    _ = p1
    _ = v30
    _ = v31
    return (p0 & p0)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{983}
    var v3 float64 = (((1e10 / 1.0) + float64(0.1)) / 3.14159)
    var v4 int64 = 163
    v5 := make(map[string]int)
    var v6 int8 = 72
    if ((!false) && ((v4 | (^(v4))) >= (v4 | 1))) {
        fmt.Println(v3, v3, v3)
    } else {
        for fi32 := 0; fi32 < 6; fi32++ {
            v4 = (-(v4))
            for fi33 := 0; fi33 < 10; fi33++ {
                v4 -= 9223372036854775807
                v4 = v4
                _ = fi33
            }
            switch "hello" {
            case "a":
                switch ("hello" + "") {
                case "hello":
                    var v34 string = "\n\t"
                    _ = v34
                }
            case "hello":
                v3 += ((v3 - v3) * ((v3 * 1.5) - 1.5))
            case "\n\t":
                switch (fi32 * fi32) {
                case 2147483647:
                    for fi35 := 0; fi35 < 7; fi35++ {
                        fmt.Println(fi32, fi35)
                        _ = fi35
                    }
                }
            }
            _ = fi32
        }
    }
    switch "" {
    case "a":
        v5["a"] = genFn1(128, 413)
    case "hello":
        switch 42 {
        case -340:
            v5[("a" + "a")] = 42
        case 32767:
            fmt.Println(v4)
        case 197:
            v36 := GenS4{F0: 39}
            _ = v36
        }
    default:
        v6 = -128
    }
    v2 = append(v2, 42)
    v37 := GenS2{F0: (v4 | (v4 - (-(v4)))), F1: (true && (!(656 < 796)))}
    v4 = (v4 ^ (-(v4)))
    v38 := GenS3{F0: v4, F1: 42, F2: -1}
    for fi39 := 0; fi39 < 19; fi39++ {
        v3 -= v3
        v40 := GenS3{F0: (v4 * 42), F1: -1, F2: 42}
        for fi41 := 0; fi41 < 6; fi41++ {
            fmt.Println(v4, v4)
            var v42 string = ("foo bar" + ("\n\t" + ""))
            _ = fi41
            _ = v42
        }
        _ = fi39
        _ = v40
    }
    v3 = ((v3 + (v3 + 1e10)) / v3)
    v43 := GenS2{F0: 9223372036854775807, F1: (v3 < (v3 + v3))}
    switch 207 {
    case -382:
        v4 -= v4
    case 746:
        if true {
            v44 := GenS4{F0: 0}
            if (true && false) {
                v1[2147483647] = "\n\t"
                v45 := v5["\n\t"]
                if len(v2) > 0 {
                    v2[0] = v45
                }
                _ = v45
            }
            if (true || (!true)) {
                var v46 float64 = float64(0.0)
                _ = v46
            } else {
                var v47 uint16 = 752
                fmt.Println(v3, v6)
                for fi48 := 0; fi48 < 19; fi48++ {
                    for fi49 := 0; fi49 < 14; fi49++ {
                        var v50 string = "\n\t"
                        if false {
                            v4 = (-(v4))
                        } else {
                            if ((v47 << ((296) & 63)) <= (v47 & (v47 & 0))) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v51 int64 = v4
                                _ = v51
                            } else {
                                var v52 uint = 1
                                _ = v52
                            }
                            v53 := GenS2{F0: (v4 % ((v4) | 1)), F1: false}
                            var v54 uint32 = 885
                            _ = v53
                            _ = v54
                        }
                        _ = fi49
                        _ = v50
                    }
                    var v55 bool = (!true)
                    _ = fi48
                    _ = v55
                }
                _ = v47
            }
            _ = v44
        }
    case -424:
        _ = len(v2)
    }
    var v56 uint16 = 65535
    if false {
        if false {
            for fi57 := 0; fi57 < 15; fi57++ {
                v3 = -2.25
                var v58 int = (fi57 | (-(fi57)))
                var v59 int32 = -1
                _ = fi57
                _ = v58
                _ = v59
            }
        }
    } else {
        _ = len(v5)
    }
    for fi60 := 0; fi60 < 14; fi60++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = len(v2)
        v56 -= (v56 | (v56 << ((18446744073709551615) & 63)))
        _ = fi60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v5
    _ = v56
    _ = v6
}
