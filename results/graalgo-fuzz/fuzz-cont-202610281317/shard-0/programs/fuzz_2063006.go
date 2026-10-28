package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int8; F2 int64 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint; F1 string; F2 int }
type GenS6 struct { F0 uint }
func genFn0(p0 uint32, p1 int) string {
    var v4 int = (p1 | 127)
    if true {
        fmt.Println(p0, p1)
        for fi5 := 0; fi5 < 11; fi5++ {
            fmt.Println(fi5, p0)
            for fi6 := 0; fi6 < 18; fi6++ {
                var v7 float32 = (0.0 * (3.14159 * 1e10))
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
        p1 = v4
    } else {
        var v8 string = ("foo bar" + "hello")
        v9 := []int{32768}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int16 = 0
        var v12 uint16 = 65535
        if true {
            p1 = 0
        } else {
            switch (^(fi10)) {
            case -345:
                v13 := []int{(v4 ^ (fi10 >> ((18446744073709551615) & 63)))}
                _ = v13
            case 128:
                if (1 >= 28) {
                    for fi14 := 0; fi14 < 11; fi14++ {
                        if (42 != 140) {
                            switch "hello" {
                            case "a":
                                fmt.Println(v12, v12, v11)
                            case "foo bar":
                                var v15 uint = 2
                                _ = v15
                            }
                            fmt.Println(p0)
                            v4 = 128
                        } else {
                            var v16 uint32 = (p0 - 2)
                            var v17 uint64 = 0
                            switch fi10 {
                            case 32767:
                                var v18 string = "\n\t"
                                _ = v18
                            default:
                                switch "a" {
                                case "":
                                    if true {
                                        fmt.Println(v16, p0)
                                        switch p1 {
                                        case 9223372036854775807:
                                            if (!true) {
                                                switch fi14 {
                                                case 198:
                                                    if false {
                                                        if (p0 > v16) {
                                                            var v19 uint = 1
                                                            if true {
                                                                v19 = (^(v19))
                                                                var v20 string = (("a" + ("\n\t" + ((((("hello" + ("hello" + "")) + ("foo bar" + (("" + ("hello" + "")) + ("" + ("\n\t" + ("foo bar" + "foo bar")))))) + "\n\t") + "a") + "\n\t"))) + "a")
                                                                var v21 float32 = 0.0
                                                                _ = v20
                                                                _ = v21
                                                            }
                                                            _ = v19
                                                        } else {
                                                            v22 := []int{}
                                                            v22 = append(v22, fi10)
                                                            var v23 uint32 = v16
                                                            _ = v22
                                                            _ = v23
                                                        }
                                                    }
                                                case -1:
                                                    p1 -= (-(fi10))
                                                case 0:
                                                    fmt.Println(v12, v17, p0)
                                                }
                                                if false {
                                                    if true {
                                                        fmt.Println(p1)
                                                    }
                                                }
                                                v24 := GenS1{F0: (v4 ^ fi14)}
                                                _ = v24
                                            }
                                        default:
                                            var v25 int16 = -370
                                            _ = v25
                                        }
                                        v26 := GenS2{F0: v11}
                                        _ = v26
                                    }
                                case "a":
                                    switch v4 {
                                    case -9223372036854775808:
                                        if true {
                                            p0 = (v16 >> ((65535) & 63))
                                            p1 = fi10
                                        } else {
                                            var v27 uint16 = v12
                                            _ = v27
                                        }
                                    case 127:
                                        v17 += uint64(v16)
                                    default:
                                        var v28 uint16 = 42
                                        _ = v28
                                    }
                                default:
                                    var v29 uint16 = v12
                                    _ = v29
                                }
                            }
                            _ = v16
                            _ = v17
                        }
                        var v30 int32 = 42
                        v31 := GenS3{F0: (p0 ^ p0), F1: 100, F2: 0}
                        _ = fi14
                        _ = v30
                        _ = v31
                    }
                } else {
                    if (127 == -113) {
                        fmt.Println(p0)
                    }
                    if (!(4294967295 == 18446744073709551615)) {
                        v32 := GenS4{F0: "foo bar", F1: 0}
                        v33 := GenS1{F0: v4}
                        _ = v32
                        _ = v33
                    }
                    fmt.Println(v12, p1)
                }
            case 0:
                v11 = (v11 | -32768)
            }
            v34 := GenS5{F0: 739, F1: (("" + (("\n\t" + ("\n\t" + ("" + "foo bar"))) + ("hello" + "hello"))) + "a"), F2: (p1 << ((964) & 63))}
            v35 := []int{}
            _ = v34
            _ = v35
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    v36 := []int{-512, int(p0)}
    _ = p0
    _ = p1
    _ = v36
    _ = v4
    return ""
}

func main() {
    var v1 float64 = (0.1 - 0.1)
    var v2 uint16 = 830
    v3 := make(map[string]string)
    v1 = float64(((0.0 - -2.25) + ((1.5 * -2.25) - ((0.0 - float32((v1 - v1))) * 1.5))))
    for fi37 := 0; fi37 < 4; fi37++ {
        v2 += v2
        v38 := GenS5{F0: 255, F1: "hello", F2: fi37}
        for fi39 := 0; fi39 < 5; fi39++ {
            v1 += 3.14159
            var v40 uint16 = (v2 / ((v2) | 1))
            _ = fi39
            _ = v40
        }
        _ = fi37
        _ = v38
    }
    switch "a" {
    case "foo bar":
        v1 = 0.0
    case "\n\t":
        var v41 float32 = -2.25
        _ = v41
    case "hello":
        v1 -= v1
    default:
        for fi42 := 0; fi42 < 18; fi42++ {
            v43 := GenS5{F0: 18446744073709551615, F1: ("foo bar" + "foo bar"), F2: (-(fi42))}
            var v44 uint8 = uint8(v2)
            _ = fi42
            _ = v43
            _ = v44
        }
    }
    v45 := GenS2{F0: 0}
    for fi46 := 0; fi46 < 12; fi46++ {
        fmt.Println(fi46, v2, fi46)
        _ = fi46
    }
    if true {
        if false {
            var v47 int32 = -222
            var v48 string = ("" + "hello")
            for fi49 := 0; fi49 < 13; fi49++ {
                v2 = (v2 / (((v2 >> ((42) & 63))) | 1))
                _ = fi49
            }
            _ = v47
            _ = v48
        }
    }
    switch "\n\t" {
    case "hello":
        if false {
            for fi50 := 0; fi50 < 20; fi50++ {
                switch "\n\t" {
                case "\n\t":
                    v51 := GenS6{F0: 10}
                    _ = v51
                }
                var v52 string = "hello"
                _ = fi50
                _ = v52
            }
            var v53 int16 = 789
            v54 := []int{553}
            _ = v53
            _ = v54
        } else {
            fmt.Println(v2, v1)
        }
    case "a":
        v2 -= v2
    }
    for fi55 := 0; fi55 < 9; fi55++ {
        switch ("hello" + "") {
        case "":
            v56 := GenS4{F0: "", F1: 65535}
            _ = v56
        case "\n\t":
            fmt.Println(fi55, v2, v2)
        }
        _ = fi55
    }
    if true {
        if false {
            v1 -= -2.25
            fmt.Println(v2, v1)
            fmt.Println(v1)
        }
        v1 += v1
    }
    v57 := GenS3{F0: 406, F1: -1, F2: 30}
    var v58 bool = false
    if (!v58) {
        switch -9223372036854775808 {
        case -2147483648:
            v58 = (-9223372036854775808 == 42)
        case -1:
            for fi59 := 0; fi59 < 1; fi59++ {
                var v60 uint16 = (v2 >> ((716) & 63))
                _ = fi59
                _ = v60
            }
        default:
            v2 = (v2 % (((v2 ^ (v2 * (v2 + v2)))) | 1))
        }
        var v61 float32 = -2.25
        v1 = float64(((v61 * v61) / (((0.1 + (((v61 + v61) / v61) * ((1e10 * float32(float64((1.5 * float32(v1))))) / 1.5))) - 0.1) + v61)))
        _ = v61
    }
    var v62 bool = true
    v3[("foo bar" + "foo bar")] = ""
    fmt.Println(v2, v2)
    switch 317 {
    case 1:
        _ = len(v3)
    case 617:
        var v63 uint8 = 191
        _ = v63
    case 128:
        var v64 float32 = 1.5
        _ = v64
    default:
        for fi65 := 0; fi65 < 17; fi65++ {
            v66 := GenS1{F0: fi65}
            var v67 bool = (65535 > 4294967295)
            v3[("a" + "hello")] = "\n\t"
            _ = fi65
            _ = v66
            _ = v67
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v58:", v58)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v57
    _ = v58
    _ = v62
}
