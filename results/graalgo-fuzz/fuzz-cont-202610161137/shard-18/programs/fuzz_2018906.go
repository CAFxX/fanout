package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 string; F1 uint32 }
func genFn0(p0 int16) uint32 {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: (p0 << ((42) & 63))}
        v6 := []int{32768, 2147483647, (fi4 - fi4)}
        fmt.Println(fi4)
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(p0)
    for fi7 := 0; fi7 < 11; fi7++ {
        if true {
            v8 := GenS1{F0: p0}
            fmt.Println(fi7, fi7)
            v9 := []int{2147483647, fi7, (fi7 >> ((772) & 63)), -2147483648}
            _ = v8
            _ = v9
        }
        var v10 int8 = 100
        _ = fi7
        _ = v10
    }
    _ = p0
    return 2147483647
}
func genFn1(p0 int16, p1 int, p2 uint) uint64 {
    fmt.Println(p2, p2)
    var v11 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 3.14159
    v3 := []uint32{60, 0}
    fmt.Println(v2)
    if false {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 bool = (0 <= 1)
            fmt.Println(v13)
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 string = (("a" + "a") + "")
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS1{F0: 32767}
    v17 := GenS2{F0: 0, F1: 42}
    if (42 > 2) {
        var v18 uint64 = 0
        _ = v18
    }
    if (-781 < 359) {
        v19 := GenS2{F0: 141, F1: 0}
        switch ("foo bar" + "") {
        case "a":
            switch -1 {
            case 1:
                v20 := GenS3{F0: "", F1: 0}
                _ = v20
            case 32768:
                v21 := GenS2{F0: 1, F1: 1}
                _ = v21
            default:
                for fi22 := 0; fi22 < 18; fi22++ {
                    var v23 int = (fi22 << ((65535) & 63))
                    v23 = -9223372036854775808
                    _ = fi22
                    _ = v23
                }
            }
        case "foo bar":
            if false {
                v24 := GenS2{F0: 2147483647, F1: -11}
                fmt.Println(v2)
                switch 2147483647 {
                case 0:
                    v25 := GenS4{F0: ("\n\t" + "\n\t"), F1: 42}
                    _ = v25
                default:
                    switch "hello" {
                    case "a":
                        fmt.Println(v2)
                    case "foo bar":
                        var v26 float32 = 3.14159
                        _ = v26
                    case "hello":
                        for fi27 := 0; fi27 < 13; fi27++ {
                            if (42 == 4294967295) {
                                if true {
                                    if ((2147483647 >= 123) && true) {
                                        for fi28 := 0; fi28 < 7; fi28++ {
                                            fmt.Println(fi28, v2)
                                            _ = fi28
                                        }
                                        v2 -= 3.14159
                                    } else {
                                        v2 = v2
                                    }
                                    switch (fi27 >> ((42) & 63)) {
                                    case 32767:
                                        switch "a" {
                                        case "hello":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        default:
                                            var v29 int8 = 100
                                            _ = v29
                                        }
                                    case -496:
                                        v2 += -2.25
                                    }
                                }
                                var v30 int16 = 0
                                v2 += -2.25
                                _ = v30
                            } else {
                                switch -580 {
                                case 0:
                                    switch "\n\t" {
                                    case "":
                                        v2 += 1.5
                                    case "hello":
                                        var v31 uint64 = genFn1(1, (^(fi27)), 255)
                                        _ = v31
                                    default:
                                        var v32 float64 = (float64((0.0 - 1.5)) + 3.14159)
                                        _ = v32
                                    }
                                case 303:
                                    v1[2147483648] = "\n\t"
                                }
                            }
                            _ = fi27
                        }
                    }
                }
                _ = v24
            }
        case "":
            var v33 int16 = -203
            _ = v33
        default:
            var v34 uint64 = 2
            _ = v34
        }
        _ = v19
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        switch "hello" {
        case "foo bar":
            switch "" {
            case "\n\t":
                if len(v3) > 0 {
                    v3[0] = 669
                }
            case "a":
                var v36 int16 = -453
                _ = v36
            case "hello":
                if true {
                    for fi37 := 0; fi37 < 17; fi37++ {
                        if (fi37 >= fi37) {
                            var v38 int16 = 66
                            v2 = float64((0.1 - -2.25))
                            v2 = 0.1
                            _ = v38
                        }
                        _ = fi37
                    }
                    var v39 uint32 = genFn0(32767)
                    switch (fi35 * (^(fi35))) {
                    case -158:
                        if true {
                            if (!false) {
                                if true {
                                    v40 := GenS3{F0: "\n\t", F1: v39}
                                    _ = v40
                                }
                                var v41 int32 = -1
                                for fi42 := 0; fi42 < 13; fi42++ {
                                    v39 += (v39 + v39)
                                    v3 = append(v3, genFn0(-32768))
                                    v43 := GenS2{F0: (v41 - (v41 * v41)), F1: 42}
                                    _ = fi42
                                    _ = v43
                                }
                                _ = v41
                            }
                            v44 := GenS1{F0: -908}
                            _ = v44
                        }
                    default:
                        var v45 float32 = 3.14159
                        _ = v45
                    }
                    _ = v39
                } else {
                    switch "" {
                    case "":
                        for fi46 := 0; fi46 < 2; fi46++ {
                            fmt.Println(v2, fi35, fi46)
                            switch (fi35 & (-(fi46))) {
                            case 127:
                                v47 := GenS2{F0: 0, F1: 42}
                                _ = v47
                            case 2147483647:
                                v1[42] = ("a" + "foo bar")
                            case -9223372036854775808:
                                var v48 int16 = 0
                                _ = v48
                            }
                            _ = fi46
                        }
                    case "hello":
                        if ("a" != (("" + "a") + "hello")) {
                            for fi49 := 0; fi49 < 14; fi49++ {
                                if (2 > 371) {
                                    switch ("hello" + (("\n\t" + "") + "foo bar")) {
                                    case "":
                                        _ = len(v1)
                                    case "\n\t":
                                        v2 = (1.5 * v2)
                                        switch "a" {
                                        case "foo bar":
                                            v2 = v2
                                        default:
                                            v2 = -2.25
                                        }
                                        var v50 int16 = 1000
                                        _ = v50
                                    case "foo bar":
                                        v2 = (v2 - 1e10)
                                    default:
                                        fmt.Println(v2)
                                    }
                                }
                                _ = len(v1)
                                _ = fi49
                            }
                        } else {
                            v2 = (((v2 * v2) / v2) - float64(1.5))
                            if (32767 > 138) {
                                v51 := GenS2{F0: -1, F1: 2}
                                _ = v51
                            }
                        }
                    case "\n\t":
                        if false {
                            if false {
                                fmt.Println(v2, v2)
                                fmt.Println(fi35, fi35)
                            }
                        }
                    }
                    fmt.Println(v2, fi35)
                }
            default:
                var v52 bool = ((255 < 255) && true)
                _ = v52
            }
        default:
            for fi53 := 0; fi53 < 14; fi53++ {
                v2 = v2
                _ = fi53
            }
        }
        v2 += 1e10
        var v54 uint = 678
        _ = fi35
        _ = v54
    }
    v2 += (float64(0.1) * ((v2 * 0.0) - (0.0 - (v2 - (v2 * float64((1.5 - 0.1)))))))
    for fi55 := 0; fi55 < 2; fi55++ {
        if (!true) {
            fmt.Println(fi55)
        } else {
            v2 += v2
        }
        v2 += v2
        v2 = v2
        _ = fi55
    }
    if true {
        for fi56 := 0; fi56 < 3; fi56++ {
            var v57 uint8 = 0
            _ = fi56
            _ = v57
        }
        v2 = -2.25
        var v58 uint = 42
        _ = v58
    } else {
        var v59 float32 = float32(-2.25)
        fmt.Println(v2, v2)
        var v60 int8 = 1
        _ = v59
        _ = v60
    }
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
