package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 int64; F2 int16 }
type GenS7 struct { F0 int; F1 uint32; F2 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int; F1 int }
type GenS10 struct { F0 string; F1 uint32; F2 int32 }
func genFn0() uint {
    v7 := []int{4294967295}
    for fi8 := 0; fi8 < 3; fi8++ {
        fmt.Println(fi8)
        for fi9 := 0; fi9 < 2; fi9++ {
            fmt.Println(fi9)
            var v10 string = "foo bar"
            var v11 int = (fi8 ^ (fi8 / ((fi9) | 1)))
            _ = fi9
            _ = v10
            _ = v11
        }
        fmt.Println(fi8)
        _ = fi8
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        switch ("a" + "") {
        case "hello":
            var v13 float32 = 0.0
            _ = v13
        case "foo bar":
            var v14 uint = 18446744073709551615
            _ = v14
        default:
            var v15 int32 = -1
            _ = v15
        }
        _ = len(v7)
        v16 := GenS1{F0: -527, F1: (fi12 + (fi12 * 4294967295)), F2: 153}
        _ = fi12
        _ = v16
    }
    v17 := GenS1{F0: 344, F1: 32767, F2: 23}
    if false {
        if true {
            var v18 uint64 = 42
            v18 += v18
            _ = v18
        }
    } else {
        v19 := GenS2{F0: 9223372036854775807}
        v7 = append(v7, -9223372036854775808)
        _ = v19
    }
    _ = v17
    _ = v7
    return 65535
}
func genFn1(p0 bool) uint16 {
    v20 := []int{72}
    var v21 uint16 = 32767
    var v22 float64 = float64(0.1)
    if p0 {
        fmt.Println(v21, p0)
    } else {
        v20 = append(v20, 2147483647)
        if p0 {
            for fi23 := 0; fi23 < 7; fi23++ {
                if (-1 < 1) {
                    v20 = append(v20, fi23)
                }
                _ = fi23
            }
            for fi24 := 0; fi24 < 9; fi24++ {
                v25 := GenS2{F0: -925}
                v26 := GenS1{F0: 9223372036854775807, F1: (fi24 & -18), F2: 2}
                for fi27 := 0; fi27 < 6; fi27++ {
                    fmt.Println(fi24, v21, v21)
                    _ = fi27
                }
                _ = fi24
                _ = v25
                _ = v26
            }
            switch "hello" {
            case "":
                v21 = v21
            case "hello":
                p0 = (-953 != -967)
            case "foo bar":
                v22 += v22
            }
        } else {
            var v28 float32 = (0.0 / 3.14159)
            _ = v28
        }
    }
    var v29 string = "a"
    _ = p0
    _ = v20
    _ = v21
    _ = v22
    _ = v29
    return v21
}

func main() {
    var v1 uint8 = 0
    var v2 int64 = 1
    var v3 int32 = 386
    var v4 int16 = -300
    var v5 float32 = (1e10 / (((0.1 + 1e10) + 1.5) * -2.25))
    var v6 float64 = ((3.14159 - float64(v5)) * (3.14159 + 1e10))
    v30 := GenS2{F0: 1}
    v31 := []int{32768}
    var v32 bool = true
    fmt.Println(v32, v5)
    for fi33 := 0; fi33 < 8; fi33++ {
        v5 = (v5 * 1.5)
        _ = fi33
    }
    v32 = false
    for fi34 := 0; fi34 < 19; fi34++ {
        v35 := GenS3{F0: 100}
        v32 = (!false)
        _ = fi34
        _ = v35
    }
    v2 += v2
    for fi36 := 0; fi36 < 12; fi36++ {
        if len(v31) > 0 {
            v31[0] = (fi36 + fi36)
        }
        var v37 int = fi36
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        switch fi38 {
        case 32767:
            for fi39 := 0; fi39 < 7; fi39++ {
                fmt.Println(v3, fi39)
                _ = fi39
            }
        case 4294967295:
            for fi40 := 0; fi40 < 2; fi40++ {
                v41 := GenS4{F0: 42}
                v42 := GenS1{F0: v2, F1: 571, F2: v1}
                _ = fi40
                _ = v41
                _ = v42
            }
        case 9223372036854775807:
            if len(v31) > 0 {
                v31[0] = fi38
            }
        default:
            for fi43 := 0; fi43 < 12; fi43++ {
                if v32 {
                    if v32 {
                        if ((v5 + v5) <= 1e10) {
                            if (v3 == (v3 % ((v3) | 1))) {
                                var v44 int16 = v4
                                _ = v44
                            }
                        } else {
                            v45 := GenS5{F0: 304}
                            fmt.Println(fi38)
                            if ("\n\t" <= "\n\t") {
                                if len(v31) > 0 {
                                    v31[0] = fi43
                                }
                                v1 += 2
                                switch "a" {
                                case "":
                                    fmt.Println(v4, fi43)
                                case "hello":
                                    v5 = 0.0
                                }
                                v46 := GenS3{F0: 127}
                                switch "a" {
                                case "":
                                    switch 2147483647 {
                                    case -591:
                                        switch ("\n\t" + "") {
                                        case "":
                                            switch "\n\t" {
                                            case "\n\t":
                                                v47 := GenS6{F0: (fi43 ^ int(v2)), F1: v2, F2: v4}
                                                _ = v47
                                            default:
                                                v48 := GenS4{F0: 72}
                                                _ = v48
                                            }
                                        default:
                                            v1 -= v1
                                        }
                                    case -2147483648:
                                        if v32 {
                                            fmt.Println(v32, fi38, v2)
                                            fmt.Println(v5)
                                            var v49 int = 32767
                                            switch "" {
                                            case "\n\t":
                                                if len(v31) >= 2 {
                                                    v31 = v31[1:len(v31)]
                                                }
                                            case "hello":
                                                switch fi43 {
                                                case -10:
                                                    fmt.Println(v5, v32, v1)
                                                }
                                            case "foo bar":
                                                var v50 string = ("hello" + ("\n\t" + (("" + "\n\t") + "hello")))
                                                _ = v50
                                            default:
                                                if true {
                                                    v51 := GenS2{F0: (v2 / (((-(v2))) | 1))}
                                                    var v52 float32 = v5
                                                    _ = v51
                                                    _ = v52
                                                }
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                v53 := GenS7{F0: 1, F1: 628, F2: 4294967295}
                                                _ = v53
                                            case "foo bar":
                                                v49 = v49
                                            case "hello":
                                                v54 := GenS6{F0: (^(fi43)), F1: (v2 | (v2 | 42)), F2: 1000}
                                                _ = v54
                                            default:
                                                v31 = append(v31, (fi38 + -2147483648))
                                            }
                                            _ = v49
                                        }
                                    default:
                                        var v55 uint = 255
                                        _ = v55
                                    }
                                case "\n\t":
                                    v56 := GenS4{F0: 268}
                                    _ = v56
                                case "hello":
                                    v57 := GenS1{F0: v2, F1: (^(fi43)), F2: (v1 >> ((263) & 63))}
                                    _ = v57
                                }
                                _ = v46
                            } else {
                                if len(v31) > 0 {
                                    v31[0] = int(v4)
                                }
                                if false {
                                    v1 = (v1 + v1)
                                }
                                v58 := GenS8{F0: (fi38 ^ (^(fi38)))}
                                _ = v58
                            }
                            _ = v45
                        }
                        var v59 uint32 = 911
                        var v60 uint32 = 243
                        _ = v59
                        _ = v60
                    } else {
                        if ((v4 >> ((42) & 63)) <= int16(v3)) {
                            fmt.Println(fi43)
                            v2 = 419
                        }
                        var v61 int = fi43
                        switch ("hello" + "") {
                        case "a":
                            switch (((("\n\t" + "hello") + ("\n\t" + "foo bar")) + "a") + ("\n\t" + ("\n\t" + ("a" + ("hello" + ""))))) {
                            case "hello":
                                var v62 uint16 = 32767
                                _ = v62
                            case "a":
                                v4 -= (v4 ^ 42)
                            default:
                                if len(v31) >= 2 {
                                    v31 = v31[1:len(v31)]
                                }
                            }
                        }
                        _ = v61
                    }
                    v6 = v6
                    v63 := GenS9{F0: fi38, F1: 4294967295}
                    _ = v63
                }
                var v64 int8 = -128
                v2 = (v2 | (-(v2)))
                v6 -= v6
                _ = fi43
                _ = v64
            }
        }
        for fi65 := 0; fi65 < 16; fi65++ {
            var v66 uint32 = 0
            _ = fi65
            _ = v66
        }
        if (!true) {
            v1 = v1
            if ("foo bar" < "foo bar") {
                for fi67 := 0; fi67 < 9; fi67++ {
                    switch "a" {
                    case "foo bar":
                        _ = len(v31)
                    case "":
                        _ = len(v31)
                    case "a":
                        var v68 string = ""
                        _ = v68
                    }
                    fmt.Println(fi38, fi67)
                    if (v5 > v5) {
                        v3 = (v3 / ((2147483647) | 1))
                        v69 := GenS10{F0: "a", F1: 2, F2: v3}
                        v3 = (v3 + (^(v3)))
                        _ = v69
                    }
                    _ = fi67
                }
            }
        } else {
            v4 += v4
            v6 -= v6
            if len(v31) >= 2 {
                v31 = v31[1:len(v31)]
            }
        }
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
