package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 int; F2 int32 }
type GenS6 struct { F0 int16; F1 uint64; F2 int16 }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 int64; F1 int8 }
type GenS9 struct { F0 string; F1 int16 }
type GenS10 struct { F0 int64 }
func genFn0() int16 {
    var v7 float64 = 3.14159
    for fi8 := 0; fi8 < 2; fi8++ {
        if true {
            switch fi8 {
            case 788:
                v9 := GenS1{F0: 48}
                _ = v9
            }
            v7 += (((v7 + (((v7 / 0.1) * (v7 / v7)) - v7)) - -2.25) * v7)
            fmt.Println(fi8, v7)
        } else {
            fmt.Println(fi8)
            v10 := GenS1{F0: 12}
            _ = v10
        }
        fmt.Println(v7)
        _ = fi8
    }
    if (true && (4294967295 > 2147483647)) {
        if true {
            v7 = v7
            v7 = v7
            var v11 int8 = -14
            _ = v11
        } else {
            var v12 uint32 = 2
            var v13 int64 = -654
            _ = v12
            _ = v13
        }
    }
    var v14 uint64 = 1
    v7 += (v7 - (0.0 - (v7 + (v7 / -2.25))))
    _ = v14
    _ = v7
    return 32767
}

func main() {
    var v1 bool = true
    v2 := []string{}
    var v3 uint32 = 955
    var v4 float32 = (1e10 * 0.1)
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]uint32)
    v1 = v1
    var v15 uint64 = 42
    var v16 bool = (!(v1 && (0.0 <= 0.0)))
    v3 = (v3 | 1)
    _ = len(v2)
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := GenS1{F0: 0}
        _ = fi17
        _ = v18
    }
    if v1 {
        for fi19 := 0; fi19 < 16; fi19++ {
            if v16 {
                v20 := GenS2{F0: true, F1: 561, F2: 536}
                _ = v20
            } else {
                switch fi19 {
                case -9223372036854775808:
                    var v21 string = "hello"
                    _ = v21
                default:
                    if false {
                        for fi22 := 0; fi22 < 13; fi22++ {
                            if ((489 <= fi22) && (false && (!false))) {
                                fmt.Println(fi22)
                            } else {
                                if ((v15 - v15) == 596) {
                                    if true {
                                        switch ((("a" + "") + "a") + "\n\t") {
                                        case "\n\t":
                                            switch "foo bar" {
                                            case "":
                                                v15 = v15
                                            case "foo bar":
                                                v23 := GenS1{F0: 173}
                                                _ = v23
                                            case "\n\t":
                                                switch fi22 {
                                                case -1:
                                                    var v24 uint = 533
                                                    _ = v24
                                                }
                                            }
                                        case "foo bar":
                                            switch ((("" + "a") + "hello") + ((((("\n\t" + ("hello" + "")) + ((("foo bar" + "hello") + "foo bar") + ("foo bar" + "hello"))) + "foo bar") + "a") + "hello")) {
                                            case "hello":
                                                if true {
                                                    if ((!v1) || ((v15 % ((uint64(v3)) | 1)) == (v15 << ((255) & 63)))) {
                                                        if ((-2.25 >= 1.5) || (false || (127 > 128))) {
                                                            v25 := GenS3{F0: 0, F1: -2147483648}
                                                            if (1 == -456) {
                                                                var v26 float32 = -2.25
                                                                fmt.Println(v4)
                                                                var v27 int16 = 1
                                                                _ = v26
                                                                _ = v27
                                                            }
                                                            if (65535 == 42) {
                                                                _ = len(v2)
                                                            } else {
                                                                switch "" {
                                                                case "\n\t":
                                                                    v28 := GenS4{F0: (^(v15))}
                                                                    _ = v28
                                                                case "hello":
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v29 := GenS1{F0: 135}
                                                                        _ = v29
                                                                    default:
                                                                        v30 := GenS5{F0: 9223372036854775807, F1: 683, F2: 42}
                                                                        _ = v30
                                                                    }
                                                                }
                                                            }
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v3 -= (^(v3))
                                                            }
                                                            _ = v25
                                                        } else {
                                                            if (int64(fi19) == 0) {
                                                                var v31 bool = v16
                                                                v31 = v1
                                                                fmt.Println(fi22, v16)
                                                                _ = v31
                                                            } else {
                                                                if len(v2) > 0 {
                                                                    v2[0] = "foo bar"
                                                                }
                                                            }
                                                        }
                                                        var v32 int64 = -9223372036854775808
                                                        _ = v32
                                                    }
                                                }
                                            case "foo bar":
                                                switch (^(fi19)) {
                                                case 42:
                                                    fmt.Println(fi19, v16)
                                                case 9223372036854775807:
                                                    var v33 int = (fi19 * fi19)
                                                    _ = v33
                                                default:
                                                    switch ("foo bar" + ("\n\t" + "")) {
                                                    case "\n\t":
                                                        switch "\n\t" {
                                                        case "hello":
                                                            var v34 int = 2147483647
                                                            _ = v34
                                                        }
                                                        v35 := GenS6{F0: 0, F1: 0, F2: 1}
                                                        v15 += (v15 << ((4294967295) & 63))
                                                        _ = v35
                                                    case "hello":
                                                        v36 := GenS7{F0: v3, F1: 2147483647}
                                                        _ = v36
                                                    default:
                                                        _ = len(v6)
                                                    }
                                                }
                                                v15 = 387
                                                v4 = v4
                                            case "":
                                                if (false || v1) {
                                                    var v37 int16 = -629
                                                    v38 := GenS8{F0: 42, F1: 42}
                                                    _ = v37
                                                    _ = v38
                                                }
                                            }
                                        }
                                        v15 = v15
                                        var v39 int64 = -1
                                        _ = v39
                                    }
                                } else {
                                    v15 = v15
                                }
                                v3 = (v3 >> ((2) & 63))
                                fmt.Println(v1)
                            }
                            _ = fi22
                        }
                    }
                }
                for fi40 := 0; fi40 < 17; fi40++ {
                    v41 := GenS9{F0: (("foo bar" + "hello") + "foo bar"), F1: 0}
                    _ = fi40
                    _ = v41
                }
            }
            _ = fi19
        }
        for fi42 := 0; fi42 < 19; fi42++ {
            for fi43 := 0; fi43 < 3; fi43++ {
                var v44 uint64 = v15
                fmt.Println(v15, fi43)
                switch ("\n\t" + ("a" + "")) {
                case "":
                    v45 := GenS4{F0: v15}
                    _ = v45
                }
                _ = fi43
                _ = v44
            }
            v6[(v3 / (((v3 / ((v3) | 1))) | 1))] = (v3 ^ 42)
            var v46 float32 = 1e10
            _ = fi42
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 11; fi47++ {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        switch "\n\t" {
        case "foo bar":
            for fi48 := 0; fi48 < 4; fi48++ {
                var v49 int16 = -1
                v50 := GenS1{F0: 120}
                _ = fi48
                _ = v49
                _ = v50
            }
        case "a":
            v15 = 2
        case "hello":
            for fi51 := 0; fi51 < 19; fi51++ {
                if v1 {
                    var v52 uint64 = 18446744073709551615
                    v53 := v6[(v3 * (v3 / ((v3) | 1)))]
                    _ = v52
                    _ = v53
                } else {
                    var v54 bool = v16
                    v2 = append(v2, ("" + "foo bar"))
                    if (genFn0() < 32767) {
                        if v54 {
                            if (!(v15 > 2)) {
                                v4 -= v4
                                var v55 int8 = 100
                                v55 = 127
                                _ = v55
                            } else {
                                fmt.Println(v1)
                            }
                            if (-220 >= 584) {
                                v54 = v1
                                _ = len(v2)
                            } else {
                                v56 := GenS5{F0: -1000, F1: fi47, F2: 1}
                                _ = len(v2)
                                var v57 int16 = 1
                                var v58 uint32 = v3
                                var v59 uint8 = 128
                                v15 += v15
                                var v60 uint = 372
                                _ = v56
                                _ = v57
                                _ = v58
                                _ = v59
                                _ = v60
                            }
                            fmt.Println(fi51)
                        } else {
                            var v61 int16 = 529
                            var v62 uint32 = v3
                            v6[4294967295] = v62
                            _ = v61
                            _ = v62
                        }
                        fmt.Println(fi51)
                    }
                    _ = v54
                }
                fmt.Println(v3, v15, v3)
                v63 := GenS6{F0: 0, F1: 18446744073709551615, F2: genFn0()}
                _ = fi51
                _ = v63
            }
        default:
            for fi64 := 0; fi64 < 16; fi64++ {
                v1 = v16
                _ = fi64
            }
        }
        for fi65 := 0; fi65 < 15; fi65++ {
            if v16 {
                v15 = (v15 + v15)
                var v66 int64 = -866
                v15 += (v15 % (((v15 * v15)) | 1))
                v67 := GenS10{F0: -1}
                if true {
                    v68 := GenS8{F0: -9223372036854775808, F1: 42}
                    _ = v68
                }
                v69 := v5["a"]
                var v70 uint8 = 9
                _ = v66
                _ = v67
                _ = v69
                _ = v70
            }
            switch (fi65 & fi65) {
            case 306:
                v4 = ((v4 / v4) * -2.25)
            case 971:
                var v71 int = (^(fi47))
                _ = v71
            default:
                v3 -= 42
            }
            v1 = ("" > (("hello" + "\n\t") + (((((("foo bar" + "a") + "a") + "a") + "a") + "a") + ("a" + "a"))))
            if ((!(1 > 42)) && true) {
                v72 := GenS5{F0: -1, F1: fi47, F2: -1}
                _ = v72
            }
            var v73 float32 = -2.25
            _ = fi65
            _ = v73
        }
        _ = fi47
    }
    fmt.Println(v3)
    var v74 uint = 643
    var v75 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v75
}
