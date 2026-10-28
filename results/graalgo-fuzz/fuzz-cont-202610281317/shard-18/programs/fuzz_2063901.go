package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint32; F2 bool }
type GenS5 struct { F0 int16 }
func genFn0() uint32 {
    if (42 >= 52) {
        if ("hello" > "a") {
            if true {
                var v4 float64 = (0.0 - (0.1 / 1.5))
                _ = v4
            }
            var v5 uint32 = 2
            _ = v5
        } else {
            v6 := GenS1{F0: 1, F1: 263}
            var v7 int32 = 1
            for fi8 := 0; fi8 < 7; fi8++ {
                v7 += -2147483648
                _ = fi8
            }
            _ = v6
            _ = v7
        }
        var v9 int16 = 1000
        v10 := GenS2{F0: 252, F1: (v9 * (-(v9)))}
        _ = v10
        _ = v9
    } else {
        var v11 uint8 = 1
        v11 = (v11 + (v11 / (((v11 % ((v11) | 1))) | 1)))
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 int64 = -257
            var v15 uint32 = 697
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 uint = 981
        _ = fi12
        _ = v16
    }
    var v17 uint = 957
    _ = v17
    return 2
}
func genFn1(p0 uint8) uint64 {
    var v18 int16 = 32767
    v18 += v18
    if (42 < 1) {
        switch -1 {
        case 42:
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 uint32 = 745
                _ = fi19
                _ = v20
            }
        }
        if (false && (!(false || (false || ((p0 | p0) == p0))))) {
            v21 := []int{127}
            v22 := GenS3{F0: 1}
            _ = v21
            _ = v22
        }
        var v23 int32 = 1
        _ = v23
    } else {
        fmt.Println(p0, v18)
        var v24 uint8 = p0
        switch "hello" {
        case "":
            var v25 uint64 = 660
            _ = v25
        }
        _ = v24
    }
    var v26 int64 = 9223372036854775807
    if true {
        for fi27 := 0; fi27 < 2; fi27++ {
            fmt.Println(fi27, v26, fi27)
            var v28 float32 = (0.0 - 1.5)
            var v29 uint64 = 42
            _ = fi27
            _ = v28
            _ = v29
        }
        v26 = (v26 << ((65535) & 63))
        var v30 int8 = 1
        _ = v30
    } else {
        v26 += v26
        p0 -= 71
    }
    _ = p0
    _ = v18
    _ = v26
    return 836
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    v3 := []string{("a" + "\n\t")}
    if (false && true) {
        for fi31 := 0; fi31 < 4; fi31++ {
            fmt.Println(fi31)
            var v32 uint32 = 2147483647
            _ = fi31
            _ = v32
        }
        var v33 int16 = 32767
        _ = v33
    }
    var v34 float64 = 0.1
    if (-2147483648 <= -2147483648) {
        v35 := GenS1{F0: -1, F1: 852}
        switch "hello" {
        case "a":
            var v36 uint8 = 237
            _ = v36
        default:
            fmt.Println(v34)
        }
        v34 = v34
        _ = v35
    }
    v37 := v2[42]
    v38 := v1[-9223372036854775808]
    v34 += (3.14159 - (v34 / 1.5))
    var v39 uint64 = 1
    var v40 uint16 = 1
    if true {
        _ = len(v1)
    } else {
        switch "a" {
        case "a":
            v41 := v1[9223372036854775807]
            _ = v41
        case "\n\t":
            v34 = 1e10
        case "":
            fmt.Println(v39)
        }
        fmt.Println(v34, v40)
        delete(v2, uint32(v40))
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        switch v37 {
        case "\n\t":
            var v43 uint32 = 891
            _ = v43
        case "foo bar":
            switch v37 {
            case "":
                delete(v1, 2147483647)
            case "hello":
                for fi44 := 0; fi44 < 13; fi44++ {
                    v45 := GenS3{F0: 2147483648}
                    var v46 uint64 = genFn1(194)
                    delete(v2, 827)
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            case "\n\t":
                switch "" {
                case "hello":
                    var v47 float32 = ((0.1 + 3.14159) * 1.5)
                    _ = v47
                case "foo bar":
                    v2[2147483648] = ""
                }
            default:
                var v48 uint = 365
                _ = v48
            }
        case "a":
            fmt.Println(v38, v37)
        }
        var v49 bool = (1 == 2147483647)
        for fi50 := 0; fi50 < 6; fi50++ {
            var v51 int32 = 381
            _ = fi50
            _ = v51
        }
        _ = fi42
        _ = v49
    }
    _ = len(v1)
    for fi52 := 0; fi52 < 4; fi52++ {
        if true {
            var v53 bool = (!false)
            if (42 > -2147483648) {
                fmt.Println(v40)
                if v53 {
                    v37 = "a"
                    v40 += (v40 + (^(v40)))
                } else {
                    for fi54 := 0; fi54 < 16; fi54++ {
                        v3 = append(v3, v37)
                        v53 = (!((v40 - (v40 << ((2) & 63))) <= v40))
                        switch 1 {
                        case 433:
                            v53 = true
                        case 0:
                            v55 := GenS4{F0: 255, F1: 903, F2: v53}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v56 int8 = -101
                            _ = v55
                            _ = v56
                        case 4294967295:
                            _ = len(v1)
                        }
                        _ = fi54
                    }
                    for fi57 := 0; fi57 < 2; fi57++ {
                        v1[int(v39)] = v37
                        _ = fi57
                    }
                    v53 = true
                }
            } else {
                fmt.Println(v39, v40)
                v38 = (v38 + "a")
            }
            delete(v2, 2147483648)
            _ = v53
        }
        var v58 uint8 = 216
        switch "foo bar" {
        case "":
            if false {
                v3 = append(v3, v38)
                if true {
                    if true {
                        fmt.Println(v37)
                        if (42 <= (v58 - (v58 * v58))) {
                            if len(v3) > 0 {
                                v3[0] = v38
                            }
                        }
                    }
                } else {
                    v59 := GenS5{F0: 32767}
                    for fi60 := 0; fi60 < 5; fi60++ {
                        var v61 uint = 232
                        if ((!true) || false) {
                            var v62 int32 = -1
                            var v63 uint64 = genFn1(v58)
                            _ = v62
                            _ = v63
                        }
                        fmt.Println(fi60)
                        _ = fi60
                        _ = v61
                    }
                    _ = v59
                }
            }
        case "foo bar":
            var v64 float32 = float32(0.0)
            _ = v64
        case "hello":
            var v65 uint = 18446744073709551615
            _ = v65
        }
        _ = fi52
        _ = v58
    }
    switch 9223372036854775807 {
    case 32768:
        for fi66 := 0; fi66 < 8; fi66++ {
            switch ("a" + v37) {
            case "":
                v67 := GenS3{F0: 42}
                _ = v67
            default:
                var v68 float64 = 0.1
                _ = v68
            }
            var v69 int = int(v39)
            v70 := GenS4{F0: (v40 >> ((65535) & 63)), F1: 751, F2: (!false)}
            _ = fi66
            _ = v69
            _ = v70
        }
    }
    var v71 uint32 = 65535
    v38 = "a"
    v39 -= (v39 - genFn1(1))
    var v72 int64 = -503
    for fi73 := 0; fi73 < 5; fi73++ {
        switch v38 {
        case "":
            var v74 uint16 = v40
            _ = v74
        }
        _ = fi73
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v71
    _ = v72
}
