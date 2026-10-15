package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS7 struct { F0 int64; F1 string; F2 uint }
func genFn0(p0 int8, p1 int8) uint64 {
    v6 := GenS1{F0: false}
    fmt.Println(p0, p0)
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: true}
        v9 := GenS1{F0: (false || ((2 >= 183) && true))}
        var v10 uint8 = 95
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 float32 = (3.14159 * ((((((0.0 + (0.0 + 0.1)) * (0.1 * -2.25)) + -2.25) * (0.1 + (1e10 / -2.25))) / 1.0) * float32(3.14159)))
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            p0 -= (p1 >> ((255) & 63))
            v11 = float32(3.14159)
            _ = fi13
        }
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 string = ("hello" + "\n\t")
            v16 := []int{509, (fi14 % ((fi12) | 1)), (fi14 & fi12)}
            p1 = 1
            switch v15 {
            case "a":
                switch fi12 {
                case 128:
                    var v17 int32 = 0
                    _ = v17
                default:
                    p0 = 0
                }
            }
            var v18 int = fi12
            _ = fi14
            _ = v15
            _ = v16
            _ = v18
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return 42
}
func genFn1(p0 bool, p1 uint64, p2 int16) uint16 {
    if (((3.14159 + 1.5) < 3.14159) && ((p1 + (p1 >> ((2) & 63))) >= p1)) {
        fmt.Println(p0, p0)
    } else {
        v19 := GenS2{F0: 985}
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v21 := GenS1{F0: p0}
        v22 := GenS3{F0: 2}
        v23 := []int{0, (^(fi20)), int(p1), fi20}
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint8 = 2
    v3 := make(map[int]string)
    var v4 bool = false
    v5 := make(map[int]uint32)
    v24 := GenS1{F0: (!true)}
    var v25 uint8 = v2
    for fi26 := 0; fi26 < 17; fi26++ {
        fmt.Println(v1, v25, v1)
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 int32 = 42
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    v29 := []int{(v1 & v1), 32768, (v1 + v1), (v1 - v1)}
    for fi30 := 0; fi30 < 19; fi30++ {
        if v4 {
            var v31 bool = true
            _ = len(v5)
            _ = v31
        } else {
            v32 := GenS1{F0: (((1000 <= 255) || v4) || v4)}
            if v4 {
                if v4 {
                    switch ("foo bar" + (("hello" + "") + ("\n\t" + "a"))) {
                    case "":
                        var v33 int8 = -1
                        _ = v33
                    case "hello":
                        if v4 {
                            for fi34 := 0; fi34 < 12; fi34++ {
                                if v4 {
                                    var v35 int16 = -643
                                    v36 := GenS1{F0: v4}
                                    v37 := GenS4{F0: -9223372036854775808}
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                } else {
                                    var v38 float64 = (3.14159 * 3.14159)
                                    var v39 float64 = v38
                                    _ = v38
                                    _ = v39
                                }
                                var v40 int64 = 1
                                _ = fi34
                                _ = v40
                            }
                            for fi41 := 0; fi41 < 4; fi41++ {
                                fmt.Println(v25)
                                var v42 int8 = -128
                                _ = fi41
                                _ = v42
                            }
                        }
                    default:
                        v2 = 128
                    }
                    v43 := GenS5{F0: 42}
                    if v4 {
                        v1 = 127
                        for fi44 := 0; fi44 < 14; fi44++ {
                            var v45 uint8 = 42
                            if len(v29) >= 2 {
                                v29 = v29[1:len(v29)]
                            }
                            _ = fi44
                            _ = v45
                        }
                    } else {
                        var v46 float64 = 0.0
                        if (1 == 100) {
                            var v47 int32 = 2147483647
                            _ = v47
                        }
                        _ = v46
                    }
                    _ = v43
                }
                if v4 {
                    var v48 int64 = 0
                    if v4 {
                        for fi49 := 0; fi49 < 18; fi49++ {
                            switch "\n\t" {
                            case "\n\t":
                                v48 = (v48 | -9223372036854775808)
                            default:
                                var v50 int16 = 42
                                _ = v50
                            }
                            fmt.Println(v1)
                            _ = fi49
                        }
                    }
                    _ = v48
                } else {
                    if (!v4) {
                        if (65535 == 18446744073709551615) {
                            v25 -= v2
                            var v51 int = 128
                            _ = v51
                        }
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(v2, v1)
                        case "hello":
                            v1 -= -2147483648
                        case "a":
                            v52 := GenS6{F0: 360, F1: (!(!true)), F2: 693}
                            _ = v52
                        }
                        for fi53 := 0; fi53 < 4; fi53++ {
                            var v54 int32 = 2147483647
                            _ = fi53
                            _ = v54
                        }
                    } else {
                        _ = len(v3)
                    }
                    v55 := GenS7{F0: 609, F1: "hello", F2: 555}
                    if (88 >= 2) {
                        fmt.Println(fi30, v2, fi30)
                    }
                    _ = v55
                }
                v1 = 32767
            }
            _ = v32
        }
        if (0.1 == 1.5) {
            v2 = v2
        } else {
            v56 := v5[21]
            v56 = 2
            _ = len(v3)
            _ = v56
        }
        _ = fi30
    }
    v25 = 92
    for fi57 := 0; fi57 < 13; fi57++ {
        var v58 int32 = 2147483647
        if v4 {
            v4 = false
            var v59 int64 = 9223372036854775807
            v60 := GenS6{F0: 32767, F1: ((("\n\t" + "a") + "") >= "foo bar"), F2: 0}
            _ = v59
            _ = v60
        }
        for fi61 := 0; fi61 < 3; fi61++ {
            var v62 int64 = -9223372036854775808
            fmt.Println(fi61, v58, v1)
            _ = fi61
            _ = v62
        }
        _ = fi57
        _ = v58
    }
    if true {
        delete(v5, v1)
    } else {
        v1 = 42
        var v63 uint8 = v25
        for fi64 := 0; fi64 < 14; fi64++ {
            switch v1 {
            case 551:
                v2 = (v2 - v25)
            case -157:
                fmt.Println(v1)
            }
            for fi65 := 0; fi65 < 11; fi65++ {
                _ = len(v29)
                var v66 uint8 = 1
                _ = len(v29)
                _ = fi65
                _ = v66
            }
            v67 := GenS2{F0: 583}
            _ = fi64
            _ = v67
        }
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
