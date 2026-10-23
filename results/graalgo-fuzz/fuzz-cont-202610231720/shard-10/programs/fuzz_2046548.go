package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int64, p1 uint16, p2 uint8) uint {
    p1 = p1
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1() string {
    switch (((("a" + "foo bar") + "a") + "foo bar") + "a") {
    case "\n\t":
        if true {
            v5 := []int{-101}
            _ = v5
        } else {
            var v6 uint = 542
            _ = v6
        }
    case "hello":
        v7 := []int{-9223372036854775808, 128}
        _ = v7
    case "a":
        var v8 uint8 = 128
        _ = v8
    default:
        v9 := []int{-532, 0}
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint64 = 4294967295
        v12 := []int{0, (fi10 >> ((255) & 63)), 1, (-(fi10))}
        _ = fi10
        _ = v11
        _ = v12
    }
    if (!(!false)) {
        var v13 bool = ((true || (!(((65535 >= 0) && (!((42 < 32768) || (-773 <= -32768)))) || true))) && true)
        if (132 != 924) {
            for fi14 := 0; fi14 < 18; fi14++ {
                for fi15 := 0; fi15 < 12; fi15++ {
                    v16 := GenS1{F0: 65535, F1: (4294967295 == 65535), F2: 65535}
                    _ = fi15
                    _ = v16
                }
                v13 = ((42 == 1) || false)
                _ = fi14
            }
            v13 = (false || (v13 || true))
        }
        _ = v13
    } else {
        var v17 uint64 = 36
        v18 := []int{32768}
        _ = v17
        _ = v18
    }
    return "foo bar"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = true
    v3 := []string{("a" + ((((("hello" + "foo bar") + "") + ("" + ("foo bar" + "\n\t"))) + (("" + (((("\n\t" + "foo bar") + "hello") + "foo bar") + "foo bar")) + (("\n\t" + "foo bar") + ""))) + "a"))}
    v4 := make(map[int]string)
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 float32 = 1e10
        _ = fi19
        _ = v20
    }
    var v21 float64 = 1.5
    switch ("foo bar" + "") {
    case "\n\t":
        switch "" {
        case "a":
            switch 32768 {
            case -506:
                var v22 int64 = 9223372036854775807
                _ = v22
            case 128:
                switch 127 {
                case 32767:
                    var v23 int64 = 0
                    _ = v23
                default:
                    v24 := GenS1{F0: 24, F1: v2, F2: 4294967295}
                    _ = v24
                }
            case 42:
                var v25 float64 = (v21 / float64(-2.25))
                _ = v25
            default:
                v21 = v21
            }
        }
    case "a":
        if false {
            if true {
                var v26 int32 = 42
                v27 := GenS1{F0: genFn0(1, 822, 255), F1: (!v2), F2: 2}
                _ = v26
                _ = v27
            }
        }
    case "foo bar":
        v21 = v21
    }
    if (211 != 128) {
        if ((919 >= -1) && (!(818 > 0))) {
            v21 += (v21 + (v21 * (1e10 * v21)))
            switch "" {
            case "a":
                if (49 >= 1) {
                    if v2 {
                        var v28 float64 = 0.0
                        _ = v28
                    } else {
                        var v29 int16 = 452
                        var v30 int16 = 220
                        _ = v29
                        _ = v30
                    }
                    if (-2.25 <= v21) {
                        v2 = (0.1 >= -2.25)
                        for fi31 := 0; fi31 < 20; fi31++ {
                            switch "foo bar" {
                            case "\n\t":
                                v21 -= v21
                            }
                            var v32 uint64 = 1
                            for fi33 := 0; fi33 < 7; fi33++ {
                                v34 := v4[fi31]
                                v21 -= v21
                                _ = fi33
                                _ = v34
                            }
                            _ = fi31
                            _ = v32
                        }
                        switch "\n\t" {
                        case "\n\t":
                            v35 := GenS1{F0: 2, F1: true, F2: 520}
                            _ = v35
                        case "":
                            for fi36 := 0; fi36 < 16; fi36++ {
                                for fi37 := 0; fi37 < 16; fi37++ {
                                    v21 += v21
                                    _ = fi37
                                }
                                switch fi36 {
                                case 32768:
                                    _ = len(v3)
                                default:
                                    if v2 {
                                        if v2 {
                                            v38 := GenS2{F0: -628, F1: 42}
                                            _ = v38
                                        } else {
                                            for fi39 := 0; fi39 < 19; fi39++ {
                                                v40 := GenS3{F0: 706}
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v2 = (true || true)
                                                _ = fi39
                                                _ = v40
                                            }
                                            switch fi36 {
                                            case 89:
                                                v41 := GenS3{F0: 2}
                                                _ = v41
                                            case 4294967295:
                                                fmt.Println(v2, fi36)
                                            case 9223372036854775807:
                                                v21 = v21
                                            }
                                        }
                                    }
                                }
                                fmt.Println(v21)
                                _ = fi36
                            }
                        case "a":
                            fmt.Println(v21, v2)
                        default:
                            if len(v3) > 0 {
                                v3[0] = (("foo bar" + "a") + "a")
                            }
                        }
                    } else {
                        for fi42 := 0; fi42 < 4; fi42++ {
                            v2 = (true || true)
                            if v2 {
                                v21 = v21
                            }
                            var v43 int = 32767
                            _ = fi42
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 14; fi44++ {
                            v21 -= v21
                            var v45 uint32 = 1
                            _ = fi44
                            _ = v45
                        }
                    }
                } else {
                    for fi46 := 0; fi46 < 9; fi46++ {
                        var v47 int16 = -32768
                        v48 := GenS4{F0: 42, F1: 4294967295}
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    if ("" < "\n\t") {
                        v2 = false
                        delete(v1, ("hello" + ""))
                    }
                }
            case "":
                fmt.Println(v21, v21)
            }
        } else {
            v21 -= 0.1
        }
    }
    fmt.Println(v21)
    for fi49 := 0; fi49 < 9; fi49++ {
        for fi50 := 0; fi50 < 15; fi50++ {
            switch (fi50 + 127) {
            case 4294967295:
                v21 += (float64(0.0) - -2.25)
            case 737:
                fmt.Println(v21, fi49)
            case -530:
                fmt.Println(v21, v21, fi49)
            default:
                var v51 uint32 = 0
                _ = v51
            }
            _ = fi50
        }
        v52 := GenS5{F0: 255, F1: 2}
        switch (fi49 + (fi49 / (((fi49 >> ((genFn0(-1, 255, 255)) & 63))) | 1))) {
        case 905:
            var v53 uint64 = 0
            _ = v53
        default:
            v21 += (1.5 / v21)
        }
        _ = fi49
        _ = v52
    }
    v2 = v2
    v54 := GenS4{F0: -9223372036854775808, F1: genFn0(0, 212, 128)}
    for fi55 := 0; fi55 < 1; fi55++ {
        v56 := GenS5{F0: 0, F1: 4294967295}
        for fi57 := 0; fi57 < 11; fi57++ {
            if v2 {
                v2 = (v2 && v2)
            }
            v21 -= (v21 + (((v21 - v21) / v21) - v21))
            _ = fi57
        }
        var v58 int16 = 1000
        _ = fi55
        _ = v56
        _ = v58
    }
    v21 = v21
    var v59 uint8 = 206
    for fi60 := 0; fi60 < 2; fi60++ {
        fmt.Println(v59)
        _ = fi60
    }
    var v61 int8 = 127
    if ((v2 && (!(1e10 == (((((v21 - (3.14159 * 0.1)) - 1e10) + ((((3.14159 * (1.5 / 1.0)) - float64(1e10)) + v21) * 0.0)) / v21) - 3.14159)))) && v2) {
        v61 += (v61 | (v61 ^ 42))
        for fi62 := 0; fi62 < 16; fi62++ {
            var v63 string = (((("\n\t" + ("a" + "\n\t")) + "\n\t") + "") + "foo bar")
            _ = fi62
            _ = v63
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v54
    _ = v59
    _ = v61
}
