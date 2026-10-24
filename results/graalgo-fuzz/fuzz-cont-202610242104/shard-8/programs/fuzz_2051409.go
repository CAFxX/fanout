package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 bool; F1 uint32; F2 int }
type GenS5 struct { F0 uint; F1 uint8; F2 uint }
type GenS6 struct { F0 uint; F1 bool; F2 uint }
func genFn0(p0 string, p1 int32) int16 {
    switch "" {
    case "":
        v6 := []int{}
        _ = v6
    case "\n\t":
        p0 = p0
    case "foo bar":
        fmt.Println(p1)
    }
    v7 := []int{390, 2147483647}
    _ = p0
    _ = p1
    _ = v7
    return 32767
}
func genFn1(p0 bool, p1 int32) int32 {
    if p0 {
        var v8 uint16 = 2
        v9 := []int{128, 1, 2147483647}
        _ = v8
        _ = v9
    }
    p0 = (!p0)
    var v10 uint8 = 255
    var v11 float32 = -2.25
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -1
}

func main() {
    v1 := []string{"hello", ("" + (((("foo bar" + "") + "") + "\n\t") + "hello")), "a", "foo bar"}
    v2 := make(map[string]int)
    var v3 int32 = -2147483648
    var v4 bool = true
    var v5 string = "\n\t"
    var v12 int32 = -831
    if (3.14159 == 1.5) {
        fmt.Println(v5)
        for fi13 := 0; fi13 < 16; fi13++ {
            switch 4294967295 {
            case 32767:
                v14 := GenS1{F0: (!(42 != -77))}
                _ = v14
            }
            v4 = v4
            _ = fi13
        }
        fmt.Println(v12, v12)
    } else {
        var v15 int64 = 9223372036854775807
        var v16 string = ("hello" + "\n\t")
        _ = v15
        _ = v16
    }
    v5 = ""
    if true {
        if (2147483647 == 1) {
            _ = len(v2)
        } else {
            fmt.Println(v3, v3)
        }
        v2[v5] = 42
        var v17 bool = (1e10 > float32((3.14159 - 3.14159)))
        _ = v17
    } else {
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := GenS1{F0: true}
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 int32 = (v12 ^ 0)
                var v22 int8 = 100
                _ = fi20
                _ = v21
                _ = v22
            }
            if len(v1) > 0 {
                v1[0] = v5
            }
            _ = fi18
            _ = v19
        }
        for fi23 := 0; fi23 < 3; fi23++ {
            if false {
                switch (-(fi23)) {
                case -1:
                    for fi24 := 0; fi24 < 9; fi24++ {
                        var v25 int = fi24
                        switch v5 {
                        case "hello":
                            var v26 int8 = -128
                            _ = v26
                        case "\n\t":
                            v27 := GenS1{F0: (!(float64(3.14159) > 3.14159))}
                            _ = v27
                        }
                        _ = fi24
                        _ = v25
                    }
                case 2147483647:
                    var v28 uint32 = 2147483647
                    _ = v28
                }
            }
            if v4 {
                var v29 int32 = genFn1(v4, (v3 ^ 1))
                v30 := GenS1{F0: (1.5 >= ((0.1 + (0.0 / -2.25)) + float32(1.5)))}
                _ = v29
                _ = v30
            } else {
                switch fi23 {
                case 32767:
                    if v4 {
                        v31 := GenS1{F0: v4}
                        v12 -= (^(v3))
                        var v32 uint = 2
                        _ = v31
                        _ = v32
                    }
                case 32768:
                    v12 -= 42
                case 69:
                    v12 = v3
                }
                for fi33 := 0; fi33 < 10; fi33++ {
                    v5 = v5
                    fmt.Println(v4, fi33)
                    _ = fi33
                }
            }
            _ = fi23
        }
    }
    var v34 uint8 = 228
    for fi35 := 0; fi35 < 19; fi35++ {
        v3 += (v3 >> ((65535) & 63))
        _ = fi35
    }
    switch -983 {
    case -1:
        for fi36 := 0; fi36 < 7; fi36++ {
            switch (fi36 | fi36) {
            case 2147483647:
                var v37 uint = 229
                _ = v37
            }
            v12 += genFn1(v4, 0)
            v12 -= 740
            _ = fi36
        }
    }
    var v38 string = v5
    for fi39 := 0; fi39 < 19; fi39++ {
        if ((-2.25 + -2.25) != (0.0 - 3.14159)) {
            switch "" {
            case "hello":
                var v40 string = v38
                _ = v40
            default:
                for fi41 := 0; fi41 < 10; fi41++ {
                    v42 := GenS2{F0: 1, F1: (-(fi41))}
                    v34 = 128
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    _ = fi41
                    _ = v42
                }
            }
            for fi43 := 0; fi43 < 15; fi43++ {
                fmt.Println(v34)
                v44 := GenS2{F0: v34, F1: 127}
                v5 = "foo bar"
                _ = fi43
                _ = v44
            }
            for fi45 := 0; fi45 < 2; fi45++ {
                var v46 string = ""
                _ = fi45
                _ = v46
            }
        }
        for fi47 := 0; fi47 < 9; fi47++ {
            switch int(v3) {
            case -9223372036854775808:
                v48 := GenS3{F0: 448, F1: v4}
                _ = v48
            case 846:
                switch v38 {
                case "a":
                    switch -9223372036854775808 {
                    case 4294967295:
                        var v49 int16 = 0
                        _ = v49
                    case 2147483647:
                        if (v34 <= 1) {
                            v12 -= 0
                            if true {
                                v50 := GenS4{F0: v4, F1: 0, F2: (fi47 - (fi47 % ((186) | 1)))}
                                _ = v50
                            } else {
                                v5 = ("foo bar" + ("\n\t" + v38))
                            }
                            v51 := GenS5{F0: 4294967295, F1: v34, F2: 4294967295}
                            _ = v51
                        } else {
                            var v52 int = fi39
                            var v53 float64 = -2.25
                            var v54 int16 = -323
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                    default:
                        switch fi39 {
                        case -574:
                            switch v5 {
                            case "":
                                switch fi39 {
                                case 634:
                                    if v4 {
                                        if v4 {
                                            var v55 int64 = -1
                                            var v56 uint16 = 293
                                            fmt.Println(fi47, v3, v4)
                                            v5 = v38
                                            v12 += (v12 << ((65535) & 63))
                                            _ = v55
                                            _ = v56
                                        }
                                        var v57 float32 = (3.14159 + (1.5 / 1.0))
                                        var v58 uint16 = 271
                                        _ = v57
                                        _ = v58
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = v5
                                        }
                                    }
                                case 2147483647:
                                    if false {
                                        fmt.Println(v5, v12, v12)
                                        var v59 uint16 = 32767
                                        _ = v59
                                    }
                                    if (!(1 < 0)) {
                                        fmt.Println(v5, v4)
                                        if true {
                                            v4 = (873 >= 2)
                                        } else {
                                            v3 -= v12
                                            v3 -= genFn1(true, v12)
                                        }
                                        v34 -= (v34 + (v34 * 42))
                                    } else {
                                        var v60 bool = (-2.25 < 0.1)
                                        _ = v60
                                    }
                                    fmt.Println(v38)
                                default:
                                    var v61 uint8 = 2
                                    _ = v61
                                }
                            case "\n\t":
                                var v62 int8 = -128
                                _ = v62
                            case "hello":
                                switch (("" + ("\n\t" + ((v38 + v5) + v38))) + "foo bar") {
                                case "":
                                    v2[""] = fi39
                                }
                            default:
                                var v63 string = "\n\t"
                                _ = v63
                            }
                        case -1:
                            switch -578 {
                            case -589:
                                v12 = (v3 | (v3 / ((genFn1(false, (v12 & (-(v3))))) | 1)))
                            case -609:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        case 0:
                            fmt.Println(fi39, fi47, v3)
                        default:
                            if (v4 && v4) {
                                v64 := GenS4{F0: (!((v4 && (0 < 173)) || v4)), F1: 4294967295, F2: (fi39 << ((369) & 63))}
                                v38 = ("a" + v38)
                                _ = v64
                            }
                        }
                    }
                case "foo bar":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "\n\t":
                    switch fi47 {
                    case 2147483647:
                        fmt.Println(v34, v3)
                    default:
                        v2[v38] = fi47
                    }
                default:
                    v12 += 2147483647
                }
            default:
                _ = len(v2)
            }
            v65 := GenS6{F0: 598, F1: v4, F2: 507}
            var v66 int32 = int32(v34)
            v12 = (-(v12))
            _ = fi47
            _ = v65
            _ = v66
        }
        v1 = append(v1, v38)
        _ = fi39
    }
    v12 += -2147483648
    if (true && (v4 && v4)) {
        for fi67 := 0; fi67 < 5; fi67++ {
            v12 = (v3 + (v12 & (v12 + v12)))
            _ = len(v2)
            _ = fi67
        }
    } else {
        v34 = (^(v34))
        fmt.Println(v4)
        for fi68 := 0; fi68 < 6; fi68++ {
            for fi69 := 0; fi69 < 16; fi69++ {
                v70 := GenS1{F0: (true || (true && v4))}
                var v71 int32 = (v3 | v3)
                _ = fi69
                _ = v70
                _ = v71
            }
            v72 := GenS1{F0: (1e10 >= 1.5)}
            delete(v2, v5)
            _ = fi68
            _ = v72
        }
    }
    v3 = int32(v34)
    var v73 bool = (true && true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v5
    _ = v73
}
