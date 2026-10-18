package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 bool; F2 int }
func genFn0(p0 int64, p1 int8) uint {
    var v7 uint64 = 1
    var v8 bool = (!(2 <= 121))
    v9 := GenS1{F0: ("" + (("\n\t" + "foo bar") + ("foo bar" + "")))}
    var v10 uint8 = 128
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1() int64 {
    if false {
        var v11 int32 = 0
        for fi12 := 0; fi12 < 2; fi12++ {
            switch 4294967295 {
            case -17:
                switch ("" + ("a" + "")) {
                case "foo bar":
                    v13 := GenS1{F0: "foo bar"}
                    _ = v13
                default:
                    var v14 uint16 = 42
                    _ = v14
                }
            }
            fmt.Println(fi12, fi12)
            v15 := []int{-1, fi12, fi12}
            _ = fi12
            _ = v15
        }
        var v16 string = "hello"
        _ = v11
        _ = v16
    }
    var v17 int16 = 42
    _ = v17
    return 1
}

func main() {
    v1 := []int{2147483647, -2147483648}
    var v2 int32 = -872
    v3 := make(map[string]bool)
    var v4 uint64 = 951
    var v5 int32 = v2
    v6 := make(map[uint32]bool)
    if (true && (256 == 1000)) {
        for fi18 := 0; fi18 < 3; fi18++ {
            v2 = (v2 % ((v2) | 1))
            var v19 uint8 = 127
            v19 += (v19 / (((v19 % ((v19) | 1))) | 1))
            _ = fi18
            _ = v19
        }
    } else {
        var v20 bool = false
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS1{F0: "\n\t"}
        v23 := GenS1{F0: "foo bar"}
        fmt.Println(v2)
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v4, v5)
    v4 -= v4
    switch 127 {
    case 128:
        v24 := GenS1{F0: "a"}
        _ = v24
    case 42:
        if true {
            v5 += (v5 << ((255) & 63))
            v5 -= 0
            fmt.Println(v4, v5, v2)
        }
    case 32768:
        v25 := GenS1{F0: ""}
        _ = v25
    default:
        v26 := GenS1{F0: "\n\t"}
        _ = v26
    }
    switch -614 {
    case -572:
        v4 += v4
    case -9223372036854775808:
        fmt.Println(v2, v2)
    default:
        v4 = v4
    }
    v27 := GenS2{F0: 128, F1: 337}
    if false {
        v28 := GenS1{F0: "foo bar"}
        v5 += (-(v2))
        _ = v28
    } else {
        var v29 int64 = -1
        v30 := GenS1{F0: "hello"}
        _ = v29
        _ = v30
    }
    v4 = (v4 - v4)
    if (false && false) {
        for fi31 := 0; fi31 < 5; fi31++ {
            if ((^(v4)) > 2) {
                var v32 float32 = ((float32(1.5) * -2.25) + 1.5)
                _ = v32
            } else {
                switch (("\n\t" + "a") + ("foo bar" + "hello")) {
                case "hello":
                    if false {
                        _ = len(v1)
                        for fi33 := 0; fi33 < 13; fi33++ {
                            _ = len(v1)
                            _ = fi33
                        }
                    }
                }
                for fi34 := 0; fi34 < 1; fi34++ {
                    var v35 int16 = 427
                    if true {
                        v2 = (v5 / ((v5) | 1))
                    }
                    _ = fi34
                    _ = v35
                }
            }
            v4 = 106
            fmt.Println(fi31, v5, v4)
            _ = fi31
        }
    } else {
        for fi36 := 0; fi36 < 4; fi36++ {
            for fi37 := 0; fi37 < 17; fi37++ {
                if false {
                    fmt.Println(v2)
                    fmt.Println(fi37, v5)
                } else {
                    v5 -= (v5 % (((v2 + (v5 ^ v5))) | 1))
                    switch "\n\t" {
                    case "foo bar":
                        var v38 float64 = 0.0
                        _ = v38
                    case "hello":
                        v6[uint32(v4)] = true
                    }
                    switch fi37 {
                    case 32768:
                        var v39 float64 = 1.5
                        _ = v39
                    case -708:
                        v40 := GenS1{F0: "\n\t"}
                        _ = v40
                    case -2147483648:
                        switch -2147483648 {
                        case 32768:
                            fmt.Println(v2)
                        case -909:
                            if false {
                                v5 = v2
                            } else {
                                var v41 bool = (fi37 == 837)
                                var v42 int16 = 1000
                                fmt.Println(fi36, v2, v42)
                                _ = v41
                                _ = v42
                            }
                            var v43 int = fi36
                            _ = v43
                        default:
                            fmt.Println(fi36)
                        }
                    }
                    var v44 int8 = -128
                    _ = v44
                }
                _ = fi37
            }
            _ = fi36
        }
        for fi45 := 0; fi45 < 15; fi45++ {
            var v46 uint = 0
            v47 := GenS2{F0: 2, F1: 4294967295}
            for fi48 := 0; fi48 < 18; fi48++ {
                var v49 uint16 = 39
                _ = fi48
                _ = v49
            }
            _ = fi45
            _ = v46
            _ = v47
        }
    }
    if (true && true) {
        if (!((false || true) || (genFn0(0, -117) > 18446744073709551615))) {
            if true {
                v5 = v2
                fmt.Println(v2)
            } else {
                v2 -= 1
            }
            fmt.Println(v4, v2, v2)
        }
        switch 4294967295 {
        case 156:
            for fi50 := 0; fi50 < 11; fi50++ {
                v4 += (v4 >> ((genFn0(42, -1)) & 63))
                v51 := GenS3{F0: true, F1: 0}
                _ = fi50
                _ = v51
            }
        case 1:
            _ = len(v1)
        }
    }
    switch -2147483648 {
    case 1:
        delete(v3, "\n\t")
    }
    v52 := GenS2{F0: 255, F1: 772}
    for fi53 := 0; fi53 < 8; fi53++ {
        v54 := GenS4{F0: (fi53 / ((fi53) | 1))}
        var v55 int = 32767
        _ = fi53
        _ = v54
        _ = v55
    }
    if ((int16(v2) == 1) && true) {
        v4 -= v4
        if (!true) {
            v4 = (v4 & v4)
        } else {
            if true {
                v5 += 42
            } else {
                var v56 int16 = -586
                for fi57 := 0; fi57 < 17; fi57++ {
                    var v58 int16 = (v56 << ((genFn0(226, -128)) & 63))
                    v5 = -2147483648
                    var v59 int32 = v2
                    _ = fi57
                    _ = v58
                    _ = v59
                }
                var v60 float64 = float64((float32(3.14159) - (1e10 + -2.25)))
                _ = v56
                _ = v60
            }
            if len(v1) > 0 {
                v1[0] = -656
            }
        }
        var v61 float64 = (((1e10 + ((float64(float32((((float64(1e10) + 3.14159) * (float64(((1e10 + float32(1.5)) / -2.25)) * -2.25)) * (3.14159 - (3.14159 - 1.5))))) * (((-2.25 + 0.1) + float64(1.5)) + float64(1.5))) / ((0.1 * ((((3.14159 - ((float64((0.0 + 0.1)) * 3.14159) + ((3.14159 - (1e10 + (1.5 * 0.1))) * -2.25))) - 1e10) / 1e10) / (((1e10 * (float64(1e10) - 1.5)) - (3.14159 / (3.14159 + (float64(((0.1 / 1e10) + float32(-2.25))) * (0.1 * (0.0 * (1.5 + 0.1))))))) * (-2.25 - 1e10)))) / 1.0))) + 3.14159) / 0.1)
        _ = v61
    } else {
        v2 -= 151
        fmt.Println(v4, v5)
    }
    v62 := GenS5{F0: -492, F1: true, F2: -853}
    for fi63 := 0; fi63 < 5; fi63++ {
        v4 += v4
        _ = fi63
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v62
}
