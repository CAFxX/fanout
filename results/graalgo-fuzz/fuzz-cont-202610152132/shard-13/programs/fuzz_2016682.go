package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 string; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 uint64 }
func genFn0() int {
    var v5 int = -9223372036854775808
    var v6 int = v5
    v7 := []int{4294967295, (v5 | 9223372036854775807), 721, (v5 >> ((30) & 63))}
    if false {
        fmt.Println(v5, v6)
    }
    _ = v5
    _ = v6
    _ = v7
    return (v5 * v5)
}

func main() {
    v1 := []uint32{2147483648, 65535, 904}
    v2 := []uint32{}
    var v3 float32 = ((1.5 - 0.0) * 1e10)
    var v4 uint16 = 95
    if false {
        fmt.Println(v3)
    } else {
        var v8 uint32 = 42
        fmt.Println(v8, v4)
        _ = v8
    }
    if false {
        var v9 int8 = 127
        _ = v9
    } else {
        fmt.Println(v3, v4)
        for fi10 := 0; fi10 < 4; fi10++ {
            v3 = v3
            var v11 int16 = -798
            var v12 int64 = -1
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v3 += -2.25
    for fi13 := 0; fi13 < 3; fi13++ {
        if (-32768 < 1000) {
            v14 := GenS1{F0: false, F1: 4294967295, F2: -9223372036854775808}
            _ = v14
        } else {
            v4 = 115
            var v15 float32 = 1e10
            switch "\n\t" {
            case "foo bar":
                if false {
                    fmt.Println(fi13, fi13)
                    for fi16 := 0; fi16 < 7; fi16++ {
                        var v17 int16 = -1
                        _ = fi16
                        _ = v17
                    }
                } else {
                    for fi18 := 0; fi18 < 15; fi18++ {
                        v3 = ((((v15 * 0.1) * float32(float64((v3 * (v3 + v15))))) / v3) + float32(((0.1 * (1.5 + 1e10)) - -2.25)))
                        v1 = append(v1, 833)
                        switch ("\n\t" + (("foo bar" + "a") + ("hello" + "a"))) {
                        case "foo bar":
                            v19 := GenS1{F0: true, F1: 598, F2: -9223372036854775808}
                            _ = v19
                        }
                        _ = fi18
                    }
                    v4 = v4
                    v3 = 0.1
                }
            case "":
                var v20 int32 = 2147483647
                _ = v20
            case "hello":
                for fi21 := 0; fi21 < 14; fi21++ {
                    v4 = v4
                    var v22 uint64 = 2
                    _ = fi21
                    _ = v22
                }
            default:
                if (42 == 37) {
                    fmt.Println(v15, v15, v15)
                } else {
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    for fi23 := 0; fi23 < 6; fi23++ {
                        v4 = v4
                        var v24 int32 = 1
                        switch "foo bar" {
                        case "":
                            switch ((((((("\n\t" + ("foo bar" + "a")) + "foo bar") + "") + "a") + ("hello" + ((("foo bar" + "hello") + "hello") + "foo bar"))) + "hello") + "\n\t") {
                            case "a":
                                v4 += v4
                            case "":
                                v2 = append(v2, 75)
                            default:
                                var v25 uint64 = 4294967295
                                _ = v25
                            }
                        case "\n\t":
                            v3 += v3
                        case "a":
                            var v26 bool = ("" > "a")
                            _ = v26
                        }
                        _ = fi23
                        _ = v24
                    }
                }
            }
            _ = v15
        }
        v4 += v4
        if (true || true) {
            for fi27 := 0; fi27 < 12; fi27++ {
                fmt.Println(v4, fi27)
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    if true {
                        if (499 == -1) {
                            switch "foo bar" {
                            case "":
                                v4 += v4
                            case "hello":
                                fmt.Println(fi27, fi13)
                            }
                        } else {
                            switch ((("foo bar" + "") + "hello") + "hello") {
                            case "a":
                                fmt.Println(fi13, v4, v4)
                            case "\n\t":
                                var v28 uint = 42
                                _ = v28
                            case "foo bar":
                                fmt.Println(v4, fi27)
                            default:
                                var v29 uint8 = 2
                                _ = v29
                            }
                            if false {
                                v4 -= v4
                                var v30 uint64 = 0
                                v30 = v30
                                _ = v30
                            }
                            v2 = append(v2, 592)
                        }
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                    } else {
                        v3 = -2.25
                    }
                    if (127 < genFn0()) {
                        var v31 uint16 = 1
                        _ = v31
                    }
                }
                v32 := GenS1{F0: false, F1: 238, F2: 0}
                var v33 uint32 = 786
                _ = fi27
                _ = v32
                _ = v33
            }
        }
        _ = fi13
    }
    switch "hello" {
    case "a":
        v4 = (^(v4))
    case "hello":
        v34 := GenS2{F0: 35, F1: 0}
        _ = v34
    default:
        v4 = 42
    }
    v3 = 1.5
    for fi35 := 0; fi35 < 14; fi35++ {
        v36 := GenS2{F0: 585, F1: 5}
        var v37 uint32 = 0
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 int64 = 740
    v3 += 1.5
    for fi39 := 0; fi39 < 14; fi39++ {
        v40 := GenS2{F0: 311, F1: 42}
        if true {
            var v41 uint16 = v4
            var v42 uint64 = 728
            _ = v41
            _ = v42
        }
        if (2147483647 <= 0) {
            var v43 string = ("" + "hello")
            v38 -= int64(fi39)
            v44 := GenS3{F0: (((v43 + (v43 + "hello")) + "hello") + ("" + v43)), F1: 81, F2: 0}
            _ = v43
            _ = v44
        }
        _ = fi39
        _ = v40
    }
    v45 := GenS3{F0: ((("hello" + "") + "a") + (("\n\t" + "\n\t") + "hello")), F1: 65535, F2: 498}
    if (((!false) || (true || true)) || (((2 > 4294967295) || false) && true)) {
        v3 += v3
        v46 := GenS4{F0: v4}
        _ = v46
    } else {
        if (true || (!false)) {
            for fi47 := 0; fi47 < 6; fi47++ {
                v48 := GenS4{F0: v4}
                _ = fi47
                _ = v48
            }
            v49 := GenS5{F0: (v4 >> ((2) & 63)), F1: v4, F2: v38}
            fmt.Println(v38, v4)
            _ = v49
        } else {
            var v50 int8 = int8(v4)
            _ = v50
        }
        if (!true) {
            var v51 int8 = 21
            var v52 string = "\n\t"
            _ = v51
            _ = v52
        }
        var v53 float64 = 0.0
        _ = v53
    }
    var v54 uint8 = 107
    if true {
        if false {
            if len(v1) > 0 {
                v1[0] = 681
            }
            if (true || (!false)) {
                v3 -= ((v3 - (0.0 - v3)) * 1.5)
            }
            v55 := GenS5{F0: (^(v4)), F1: v4, F2: -738}
            _ = v55
        }
        for fi56 := 0; fi56 < 13; fi56++ {
            fmt.Println(v3, v38, v3)
            _ = fi56
        }
        switch "a" {
        case "foo bar":
            if false {
                v1 = append(v1, 999)
                v57 := GenS4{F0: (v4 / ((v4) | 1))}
                v38 -= (-(v38))
                _ = v57
            } else {
                v58 := GenS6{F0: 65535, F1: ("a" + (("\n\t" + "foo bar") + ("hello" + ("" + "\n\t"))))}
                _ = v58
            }
        case "":
            if false {
                v1 = append(v1, 2147483648)
            } else {
                var v59 uint16 = v4
                v60 := GenS7{F0: 0}
                var v61 int16 = 1
                _ = v59
                _ = v60
                _ = v61
            }
        case "hello":
            var v62 int8 = 127
            _ = v62
        default:
            v63 := GenS3{F0: (((("\n\t" + ("hello" + "\n\t")) + ("\n\t" + "hello")) + (("a" + "foo bar") + ("" + ("hello" + (("\n\t" + "a") + "hello"))))) + ("" + "a")), F1: 4294967295, F2: 2}
            _ = v63
        }
    }
    v54 = (v54 / (((v54 / ((127) | 1))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v45
    _ = v54
}
