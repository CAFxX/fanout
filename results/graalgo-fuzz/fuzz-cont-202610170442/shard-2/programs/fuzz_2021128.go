package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int; F1 int; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 float64 = -2.25
    var v2 bool = true
    v3 := []int{}
    v4 := []string{(("\n\t" + "a") + ""), "", ("" + "foo bar"), ""}
    v5 := []string{"foo bar", "\n\t", ("hello" + ("" + "a"))}
    if (v2 || (9223372036854775807 > 1)) {
        var v6 uint32 = 2
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                v2 = (!v2)
                _ = fi8
            }
            _ = fi7
        }
        _ = v6
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 -= -2.25
    fmt.Println(v1, v1)
    v2 = (2147483647 >= -2147483648)
    var v9 float64 = v1
    for fi10 := 0; fi10 < 20; fi10++ {
        if (!(2 != 65535)) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            if true {
                var v11 int8 = 1
                _ = v11
            }
        }
        var v12 int16 = 138
        if (v2 && (255 < 13)) {
            v13 := GenS1{F0: 912}
            v2 = (2147483647 >= 557)
            if ("\n\t" != ("" + "foo bar")) {
                for fi14 := 0; fi14 < 19; fi14++ {
                    if len(v5) > 0 {
                        v5[0] = (((((("foo bar" + "hello") + "") + ("foo bar" + "")) + "") + ("\n\t" + ("hello" + ""))) + (("foo bar" + "a") + "foo bar"))
                    }
                    var v15 int = fi14
                    switch ("\n\t" + ("" + "")) {
                    case "hello":
                        var v16 int = 1
                        _ = v16
                    case "a":
                        var v17 int = 128
                        switch 4294967295 {
                        case -974:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        v18 := GenS1{F0: 2147483647}
                        _ = v17
                        _ = v18
                    case "foo bar":
                        v2 = (255 == 0)
                    default:
                        v19 := GenS2{F0: 1, F1: 1000}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v19
                    }
                    _ = fi14
                    _ = v15
                }
            } else {
                switch -841 {
                case 729:
                    fmt.Println(fi10, v12, v1)
                case 1:
                    fmt.Println(v12)
                }
            }
            _ = v13
        } else {
            v2 = ((true || (!v2)) && (v12 != v12))
        }
        _ = fi10
        _ = v12
    }
    v20 := GenS1{F0: 65535}
    v3 = append(v3, -717)
    v21 := GenS1{F0: 42}
    var v22 string = ("foo bar" + "\n\t")
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := GenS3{F0: v2, F1: ((v22 + v22) + v22)}
        if (1 < fi23) {
            switch fi23 {
            case 32767:
                v5 = append(v5, ("hello" + v22))
            default:
                v9 += 1.5
            }
            if (18446744073709551615 < 0) {
                v4 = append(v4, "hello")
            } else {
                v2 = (!true)
                switch fi23 {
                case -9223372036854775808:
                    for fi25 := 0; fi25 < 12; fi25++ {
                        var v26 float64 = (0.0 / 1e10)
                        v27 := GenS4{F0: fi25, F1: (fi25 | (fi25 % (((fi25 << ((255) & 63))) | 1))), F2: 2}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                case 4294967295:
                    if v2 {
                        for fi28 := 0; fi28 < 4; fi28++ {
                            switch 42 {
                            case 32767:
                                var v29 float32 = (-2.25 / -2.25)
                                _ = v29
                            }
                            switch "foo bar" {
                            case "hello":
                                var v30 uint64 = 42
                                _ = v30
                            case "":
                                _ = len(v5)
                                v3 = append(v3, (-(fi23)))
                                var v31 int8 = 116
                                _ = v31
                            default:
                                v32 := GenS4{F0: 144, F1: 128, F2: 100}
                                _ = v32
                            }
                            if (42 < 2) {
                                v22 = "a"
                                v3 = append(v3, -1)
                                fmt.Println(fi23, v22, v9)
                                v9 = v9
                                fmt.Println(v22, fi23, v22)
                            }
                            if v2 {
                                var v33 bool = (!false)
                                v34 := GenS1{F0: 856}
                                v9 += v9
                                _ = v33
                                _ = v34
                            } else {
                                v22 = ("a" + v22)
                                if v2 {
                                    _ = len(v3)
                                }
                                var v35 int64 = -1
                                _ = v35
                            }
                            _ = fi28
                        }
                    } else {
                        switch fi23 {
                        case 32767:
                            var v36 float32 = -2.25
                            _ = v36
                        case 0:
                            for fi37 := 0; fi37 < 19; fi37++ {
                                var v38 uint16 = 65535
                                _ = fi37
                                _ = v38
                            }
                        case 604:
                            v2 = v2
                        default:
                            v39 := GenS2{F0: 42, F1: -32768}
                            _ = v39
                        }
                        v22 = v22
                        v1 += ((v9 * 0.1) + v9)
                    }
                }
                var v40 uint8 = 2
                _ = v40
            }
            v41 := GenS4{F0: fi23, F1: fi23, F2: 455}
            _ = v41
        }
        for fi42 := 0; fi42 < 12; fi42++ {
            v9 += v1
            var v43 uint16 = 533
            var v44 uint16 = v43
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi23
        _ = v24
    }
    if v2 {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v2, v1)
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = 193
            }
        case "hello":
            var v45 float64 = 0.1
            _ = v45
        default:
            v22 = v22
        }
        if (334 >= 0) {
            v9 -= (v9 / v1)
            for fi46 := 0; fi46 < 19; fi46++ {
                for fi47 := 0; fi47 < 9; fi47++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi47
                }
                _ = fi46
            }
        } else {
            v22 = (((v22 + v22) + (v22 + v22)) + "a")
            for fi48 := 0; fi48 < 5; fi48++ {
                fmt.Println(fi48)
                v49 := GenS1{F0: 0}
                v50 := GenS5{F0: 2, F1: 612}
                _ = fi48
                _ = v49
                _ = v50
            }
            switch v22 {
            case "foo bar":
                var v51 int64 = -9223372036854775808
                _ = v51
            default:
                for fi52 := 0; fi52 < 5; fi52++ {
                    v1 = ((-2.25 + v9) / 0.1)
                    var v53 int8 = 127
                    for fi54 := 0; fi54 < 15; fi54++ {
                        var v55 string = ("\n\t" + "\n\t")
                        var v56 float32 = 3.14159
                        v56 += (1e10 - v56)
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                    _ = fi52
                    _ = v53
                }
            }
        }
    }
    var v57 uint8 = 1
    fmt.Println(v2)
    for fi58 := 0; fi58 < 18; fi58++ {
        var v59 int64 = -740
        _ = fi58
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v9
}
