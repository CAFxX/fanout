package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int32 }
func genFn0(p0 int64, p1 int, p2 int32) string {
    p0 = 616
    if (100 > 127) {
        var v4 uint16 = 32767
        var v5 uint = 4294967295
        if true {
            switch "foo bar" {
            case "foo bar":
                v5 += 65535
            case "":
                v6 := []int{p1, 911}
                _ = v6
            case "\n\t":
                switch -1 {
                case 148:
                    fmt.Println(p2)
                case 42:
                    v7 := []int{(p1 * p1), p1}
                    _ = v7
                }
            }
            var v8 int16 = 535
            v9 := GenS1{F0: 0}
            _ = v8
            _ = v9
        }
        _ = v4
        _ = v5
    } else {
        var v10 int = 127
        switch v10 {
        case -2147483648:
            v11 := []int{v10, -402, -9223372036854775808}
            _ = v11
        case -9223372036854775808:
            v12 := []int{-430, (p1 + v10)}
            _ = v12
        }
        fmt.Println(v10, p0, p0)
        _ = v10
    }
    v13 := GenS2{F0: p2}
    for fi14 := 0; fi14 < 18; fi14++ {
        switch ((("hello" + ("\n\t" + "hello")) + "foo bar") + "a") {
        case "":
            v15 := []int{32768, 928, (-(p1))}
            _ = v15
        default:
            var v16 float64 = (-2.25 * 0.0)
            _ = v16
        }
        var v17 int8 = -68
        _ = fi14
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return "a"
}
func genFn1(p0 uint64) uint32 {
    var v18 int8 = 42
    var v19 int8 = -125
    _ = p0
    _ = v18
    _ = v19
    return 2
}

func main() {
    var v1 uint8 = 1
    var v2 float32 = 1.5
    var v3 int32 = 1
    fmt.Println(v3, v1, v3)
    for fi20 := 0; fi20 < 18; fi20++ {
        if (genFn1(2) < 1) {
            fmt.Println(v1)
            v21 := GenS3{F0: 310, F1: 113}
            v1 -= 82
            _ = v21
        } else {
            var v22 float32 = v2
            _ = v22
        }
        _ = fi20
    }
    v23 := []int{42, -9223372036854775808, -555}
    v3 = (v3 | 537)
    if true {
        v2 = (v2 * v2)
        if true {
            switch "hello" {
            case "hello":
                switch 127 {
                case 9223372036854775807:
                    switch ("hello" + "a") {
                    case "":
                        v1 = v1
                    case "foo bar":
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                    case "\n\t":
                        v3 = -2147483648
                    }
                case 864:
                    for fi24 := 0; fi24 < 15; fi24++ {
                        if false {
                            for fi25 := 0; fi25 < 4; fi25++ {
                                switch (("" + "") + "a") {
                                case "":
                                    switch 420 {
                                    case 4294967295:
                                        v1 += (^(v1))
                                    case -996:
                                        switch 128 {
                                        case 394:
                                            fmt.Println(v2, v2)
                                        case 592:
                                            if len(v23) >= 2 {
                                                v23 = v23[1:len(v23)]
                                            }
                                        case -1:
                                            if len(v23) > 0 {
                                                v23[0] = fi25
                                            }
                                        }
                                    }
                                }
                                var v26 uint64 = 1
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                                _ = fi25
                                _ = v26
                            }
                        }
                        for fi27 := 0; fi27 < 12; fi27++ {
                            v1 = 42
                            v28 := GenS1{F0: 114}
                            switch fi24 {
                            case 32767:
                                _ = len(v23)
                            case 4294967295:
                                var v29 bool = true
                                _ = v29
                            }
                            _ = fi27
                            _ = v28
                        }
                        if (0.1 == 1e10) {
                            v30 := GenS4{F0: 0, F1: 127, F2: 805}
                            _ = v30
                        }
                        _ = fi24
                    }
                }
            case "":
                switch 765 {
                case 0:
                    var v31 int32 = v3
                    _ = v31
                default:
                    if (float64((((v2 / (v2 * float32(-2.25))) - 0.1) - (1e10 * v2))) < -2.25) {
                        v32 := GenS3{F0: uint64(v1), F1: 1}
                        _ = v32
                    } else {
                        v1 = v1
                    }
                }
            case "foo bar":
                switch "" {
                case "hello":
                    v2 = v2
                case "":
                    v33 := GenS4{F0: -1, F1: 98, F2: 536}
                    _ = v33
                default:
                    v1 = (v1 ^ (v1 ^ (v1 & v1)))
                }
            default:
                var v34 float64 = ((1.5 - -2.25) / (((0.1 + 0.0) * 0.0) + 0.1))
                _ = v34
            }
        } else {
            for fi35 := 0; fi35 < 17; fi35++ {
                v36 := GenS4{F0: 32767, F1: v1, F2: 42}
                var v37 uint64 = 1
                for fi38 := 0; fi38 < 19; fi38++ {
                    if len(v23) > 0 {
                        v23[0] = 32767
                    }
                    switch -9223372036854775808 {
                    case 32768:
                        if len(v23) > 0 {
                            v23[0] = fi38
                        }
                    }
                    var v39 uint32 = 0
                    if (!false) {
                        if (65535 >= 255) {
                            var v40 int = (fi35 << ((1) & 63))
                            if false {
                                v2 = float32(0.0)
                                fmt.Println(v40, v40)
                                _ = len(v23)
                            }
                            _ = v40
                        }
                        var v41 int = fi35
                        _ = v41
                    }
                    _ = fi38
                    _ = v39
                }
                _ = fi35
                _ = v36
                _ = v37
            }
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
        }
    } else {
        if true {
            v2 -= v2
            if false {
                switch ("\n\t" + ("foo bar" + ("hello" + "\n\t"))) {
                case "foo bar":
                    v3 = -324
                }
            }
        } else {
            var v42 string = "\n\t"
            if false {
                v43 := GenS5{F0: 4294967295, F1: genFn1(1)}
                if (!(42 > -954)) {
                    for fi44 := 0; fi44 < 11; fi44++ {
                        var v45 int64 = -709
                        if (127 > 100) {
                            v42 = (((v42 + (("a" + v42) + "foo bar")) + ("a" + "foo bar")) + v42)
                            if (false || true) {
                                if (42 > genFn1(326)) {
                                    var v46 string = v42
                                    v47 := GenS6{F0: (v3 >> ((65535) & 63))}
                                    _ = v46
                                    _ = v47
                                } else {
                                    var v48 int8 = 0
                                    var v49 uint16 = 255
                                    _ = v48
                                    _ = v49
                                }
                            } else {
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                                switch "" {
                                case "hello":
                                    v45 += -185
                                default:
                                    v50 := GenS3{F0: 1, F1: 699}
                                    _ = v50
                                }
                            }
                        }
                        _ = fi44
                        _ = v45
                    }
                    _ = len(v23)
                    fmt.Println(v42, v42)
                }
                var v51 float32 = (3.14159 * v2)
                _ = v43
                _ = v51
            }
            _ = v42
        }
    }
    for fi52 := 0; fi52 < 19; fi52++ {
        fmt.Println(v2, v3, v1)
        var v53 int32 = 2147483647
        switch fi52 {
        case 391:
            v1 += 161
        }
        _ = fi52
        _ = v53
    }
    fmt.Println(v1)
    v54 := GenS5{F0: -9223372036854775808, F1: 2147483648}
    v1 += (v1 * (v1 % ((v1) | 1)))
    for fi55 := 0; fi55 < 8; fi55++ {
        if false {
            fmt.Println(v3, fi55, v1)
            v2 -= 3.14159
            var v56 float32 = v2
            _ = v56
        }
        _ = fi55
    }
    var v57 int = -1
    fmt.Println(v57, v1)
    v58 := GenS4{F0: 42, F1: 2, F2: uint(v1)}
    v23 = append(v23, (v57 & v57))
    fmt.Println(v57, v2, v2)
    if false {
        v59 := GenS4{F0: 284, F1: 127, F2: 42}
        v3 = v3
        if (("\n\t" + ("" + "a")) > "hello") {
            var v60 float32 = -2.25
            v60 = (0.0 + (v60 - (float32(1.5) - -2.25)))
            if true {
                _ = len(v23)
            } else {
                var v61 bool = (!false)
                v62 := GenS6{F0: 485}
                _ = v61
                _ = v62
            }
            _ = v60
        } else {
            v63 := GenS4{F0: 0, F1: (v1 * (v1 & (v1 / ((v1) | 1)))), F2: 1}
            _ = v63
        }
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v54
    _ = v57
    _ = v58
}
