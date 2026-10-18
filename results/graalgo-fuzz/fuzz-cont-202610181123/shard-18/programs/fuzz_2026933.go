package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int; F1 int32 }
func genFn0(p0 uint, p1 int8, p2 bool) int8 {
    v7 := GenS1{F0: (p0 * p0)}
    if false {
        if p2 {
            p1 -= 0
        }
    } else {
        var v8 int64 = int64(p1)
        _ = v8
    }
    v9 := GenS2{F0: 0}
    p2 = (-1 < -1)
    p1 = (-(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return (p1 % (((p1 >> ((936) & 63))) | 1))
}
func genFn1() uint32 {
    var v10 float64 = (0.1 * 1e10)
    if (!true) {
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS3{F0: genFn0(255, 42, (!false))}
            _ = fi11
            _ = v12
        }
    } else {
        v13 := GenS2{F0: 18446744073709551615}
        _ = v13
    }
    if (false || true) {
        for fi14 := 0; fi14 < 19; fi14++ {
            if (!((-1 < 1) && (!false))) {
                switch fi14 {
                case -1:
                    switch "" {
                    case "foo bar":
                        if (true || (true && true)) {
                            var v15 float32 = 1.5
                            v16 := []int{486}
                            _ = v15
                            _ = v16
                        }
                    }
                default:
                    v10 = 0.1
                }
                v17 := GenS3{F0: genFn0(2, -128, false)}
                var v18 int8 = genFn0(2, 42, ((fi14 * -78) <= 1))
                _ = v17
                _ = v18
            } else {
                for fi19 := 0; fi19 < 9; fi19++ {
                    fmt.Println(fi19, fi14)
                    fmt.Println(v10)
                    v20 := GenS2{F0: 4294967295}
                    _ = fi19
                    _ = v20
                }
                fmt.Println(v10, fi14)
            }
            var v21 bool = (375 < 1)
            var v22 uint8 = 91
            _ = fi14
            _ = v21
            _ = v22
        }
        if ((0 <= 488) && (v10 != 1.5)) {
            v23 := []int{32768}
            _ = v23
        } else {
            v24 := []int{-798, 4294967295, 127}
            v25 := GenS4{F0: 74, F1: -1}
            _ = v24
            _ = v25
        }
    } else {
        v26 := []int{-9223372036854775808}
        fmt.Println(v10)
        for fi27 := 0; fi27 < 20; fi27++ {
            switch -2147483648 {
            case 128:
                var v28 uint = 4294967295
                _ = v28
            case 1:
                v10 -= (v10 * 1.5)
            default:
                v29 := GenS2{F0: 42}
                _ = v29
            }
            for fi30 := 0; fi30 < 2; fi30++ {
                if len(v26) > 0 {
                    v26[0] = (fi30 * fi27)
                }
                if true {
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                    v10 = (float64(float32(v10)) - v10)
                }
                _ = fi30
            }
            v10 = (float64(float32(3.14159)) + 0.1)
            _ = fi27
        }
        _ = v26
    }
    var v31 int16 = 1000
    _ = v10
    _ = v31
    return 2
}

func main() {
    var v1 bool = false
    var v2 int = 645
    var v3 uint16 = 417
    v4 := make(map[string]string)
    v5 := make(map[uint32]int)
    v6 := make(map[string]int)
    v3 = v3
    switch "hello" {
    case "":
        v1 = v1
    case "\n\t":
        switch v2 {
        case -9223372036854775808:
            for fi32 := 0; fi32 < 7; fi32++ {
                fmt.Println(v3)
                _ = fi32
            }
        case 491:
            var v33 int32 = int32(v3)
            _ = v33
        case 296:
            var v34 uint8 = 112
            _ = v34
        default:
            for fi35 := 0; fi35 < 9; fi35++ {
                v5[1] = (fi35 << ((255) & 63))
                if (-9223372036854775808 <= 920) {
                    delete(v5, 675)
                    delete(v6, "foo bar")
                }
                fmt.Println(v3)
                _ = fi35
            }
        }
    case "a":
        v36 := GenS1{F0: 65535}
        _ = v36
    }
    v37 := GenS1{F0: 18446744073709551615}
    var v38 int16 = -1
    for fi39 := 0; fi39 < 6; fi39++ {
        v38 = 1000
        _ = fi39
    }
    var v40 uint64 = 482
    fmt.Println(v38)
    switch -805 {
    case 544:
        for fi41 := 0; fi41 < 5; fi41++ {
            switch -2147483648 {
            case 32768:
                var v42 int = 32768
                _ = v42
            case -1:
                var v43 int32 = 643
                _ = v43
            case 42:
                for fi44 := 0; fi44 < 8; fi44++ {
                    v45 := GenS3{F0: 51}
                    fmt.Println(v38, fi41, v3)
                    _ = fi44
                    _ = v45
                }
            }
            _ = fi41
        }
    case 714:
        delete(v4, "a")
    }
    if ((-790 > 2147483647) || v1) {
        v46 := []int{v2, (v2 ^ (v2 & -13)), v2}
        if (v1 || v1) {
            if v1 {
                var v47 int = (v2 | (v2 >> ((2) & 63)))
                _ = v47
            } else {
                if len(v46) > 0 {
                    v46[0] = (-(v2))
                }
                v6["foo bar"] = 919
                switch "\n\t" {
                case "":
                    if len(v46) >= 2 {
                        v46 = v46[1:len(v46)]
                    }
                case "\n\t":
                    for fi48 := 0; fi48 < 11; fi48++ {
                        var v49 uint = 255
                        if (v40 != v40) {
                            for fi50 := 0; fi50 < 9; fi50++ {
                                fmt.Println(v38, v49)
                                _ = fi50
                            }
                        } else {
                            if v1 {
                                var v51 uint = v49
                                _ = v51
                            }
                        }
                        _ = fi48
                        _ = v49
                    }
                case "hello":
                    v5[859] = int(v3)
                default:
                    var v52 int32 = -1
                    _ = v52
                }
            }
            switch 127 {
            case 4294967295:
                v3 -= (v3 % (((^(v3))) | 1))
            case 536:
                fmt.Println(v40, v2)
            case -2147483648:
                switch ("" + "\n\t") {
                case "\n\t":
                    if v1 {
                        for fi53 := 0; fi53 < 6; fi53++ {
                            _ = len(v46)
                            _ = fi53
                        }
                        delete(v5, 42)
                    } else {
                        fmt.Println(v38, v2, v38)
                    }
                case "hello":
                    v3 = (v3 << ((352) & 63))
                }
            }
        }
        _ = v46
    } else {
        v1 = v1
        if v1 {
            v38 -= (v38 >> ((4294967295) & 63))
            _ = len(v6)
        } else {
            var v54 int8 = 100
            _ = v54
        }
    }
    var v55 uint64 = v40
    fmt.Println(v2)
    if (true && (!false)) {
        v55 -= (v40 >> ((359) & 63))
        v1 = (42 >= -128)
    }
    switch "" {
    case "a":
        var v56 float64 = 0.1
        _ = v56
    case "foo bar":
        var v57 float32 = 0.1
        _ = v57
    default:
        v58 := []int{(v2 | (^(v2))), v2}
        _ = v58
    }
    for fi59 := 0; fi59 < 5; fi59++ {
        switch ("" + "hello") {
        case "foo bar":
            var v60 int64 = -1
            _ = v60
        case "hello":
            switch fi59 {
            case 4294967295:
                if (2147483647 <= 2147483647) {
                    v55 += 2
                } else {
                    if ((("" + "\n\t") <= "\n\t") && v1) {
                        var v61 uint16 = (v3 - (v3 ^ 0))
                        _ = v61
                    }
                    v38 += v38
                    for fi62 := 0; fi62 < 16; fi62++ {
                        switch ("a" + "foo bar") {
                        case "a":
                            var v63 int8 = -1
                            _ = v63
                        }
                        v64 := GenS4{F0: (v3 + v3), F1: (fi59 + fi62)}
                        _ = fi62
                        _ = v64
                    }
                }
            case 127:
                var v65 uint8 = 255
                _ = v65
            }
        }
        for fi66 := 0; fi66 < 3; fi66++ {
            v67 := GenS5{F0: (fi66 << ((42) & 63)), F1: 990}
            _ = len(v4)
            v68 := []int{(fi59 & (fi59 | 42)), 578, (fi59 ^ fi66), (fi59 - fi59)}
            fmt.Println(v55)
            _ = fi66
            _ = v67
            _ = v68
        }
        v38 += (v38 | v38)
        _ = fi59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v5
    _ = v55
    _ = v6
}
