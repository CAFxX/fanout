package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint64, p1 string) int64 {
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: (-1 > 1)}
        if false {
            switch p1 {
            case "":
                var v8 int8 = 42
                _ = v8
            case "foo bar":
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 string = "\n\t"
                    if false {
                        v10 = ""
                        if ((!true) || true) {
                            var v11 uint16 = 821
                            _ = v11
                        } else {
                            switch ("hello" + "") {
                            case "\n\t":
                                fmt.Println(v10, v10, fi9)
                            case "a":
                                var v12 float32 = 0.1
                                _ = v12
                            case "foo bar":
                                v13 := GenS2{F0: 2}
                                _ = v13
                            default:
                                var v14 int8 = 100
                                _ = v14
                            }
                        }
                        switch (-(fi9)) {
                        case 122:
                            v15 := []int{(fi9 & fi6), (fi6 * 127)}
                            _ = v15
                        case 32767:
                            v16 := GenS1{F0: false}
                            _ = v16
                        }
                        p1 = p1
                    }
                    _ = fi9
                    _ = v10
                }
            case "hello":
                p1 = p1
            }
            var v17 int8 = 100
            _ = v17
        }
        _ = fi6
        _ = v7
    }
    v18 := GenS1{F0: false}
    _ = p0
    _ = p1
    _ = v18
    return 1
}
func genFn1(p0 uint8, p1 uint64) int8 {
    v19 := GenS3{F0: "a"}
    v20 := []int{32767}
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 1
}

func main() {
    var v1 int32 = 42
    v2 := []string{}
    var v3 int8 = 42
    v4 := make(map[int]int)
    var v5 int8 = v3
    switch "\n\t" {
    case "hello":
        var v21 bool = true
        _ = v21
    case "a":
        switch -9223372036854775808 {
        case 0:
            switch 0 {
            case 958:
                var v22 float64 = 0.1
                _ = v22
            }
        case -744:
            if (!(717 == 473)) {
                switch "" {
                case "a":
                    var v23 int64 = -1
                    _ = v23
                case "foo bar":
                    var v24 uint32 = 1
                    _ = v24
                case "hello":
                    for fi25 := 0; fi25 < 11; fi25++ {
                        for fi26 := 0; fi26 < 18; fi26++ {
                            switch ("\n\t" + "hello") {
                            case "a":
                                if (-2.25 > 3.14159) {
                                    v1 += (v1 % ((0) | 1))
                                }
                            case "":
                                fmt.Println(fi26, v1)
                            }
                            v5 = (v5 + (^(v3)))
                            v27 := GenS4{F0: 65535}
                            var v28 int32 = (v1 * v1)
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        v5 = (v3 - (v5 >> ((255) & 63)))
                        if true {
                            if true {
                                for fi29 := 0; fi29 < 16; fi29++ {
                                    v5 += (v3 - v5)
                                    _ = fi29
                                }
                                for fi30 := 0; fi30 < 16; fi30++ {
                                    var v31 float32 = float32((((0.0 + (1e10 - -2.25)) / (3.14159 * 0.1)) + 1.5))
                                    v3 += genFn1(128, 1)
                                    _ = fi30
                                    _ = v31
                                }
                                _ = len(v4)
                            } else {
                                if false {
                                    fmt.Println(v3)
                                    for fi32 := 0; fi32 < 13; fi32++ {
                                        v33 := GenS4{F0: 42}
                                        var v34 bool = (float64(1.5) < ((1e10 / -2.25) + ((3.14159 - float64(((1.5 - ((0.1 - 1e10) / 0.1)) + float32(((1.5 + -2.25) / 1.5))))) - 3.14159)))
                                        fmt.Println(fi32)
                                        _ = fi32
                                        _ = v33
                                        _ = v34
                                    }
                                } else {
                                    v5 += -1
                                    v35 := GenS3{F0: "foo bar"}
                                    var v36 bool = (1e10 == 0.0)
                                    _ = v35
                                    _ = v36
                                }
                                var v37 uint64 = 1
                                for fi38 := 0; fi38 < 19; fi38++ {
                                    fmt.Println(v3, v1, v37)
                                    v39 := GenS5{F0: false, F1: -9223372036854775808}
                                    var v40 uint = 2
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                                _ = v37
                            }
                        }
                        _ = fi25
                    }
                }
            } else {
                v5 += v3
                switch ("foo bar" + "") {
                case "foo bar":
                    v5 = -128
                case "hello":
                    v3 = genFn1(127, 18446744073709551615)
                default:
                    if true {
                        fmt.Println(v3, v3, v5)
                        for fi41 := 0; fi41 < 15; fi41++ {
                            for fi42 := 0; fi42 < 4; fi42++ {
                                v4[fi41] = fi41
                                v5 += (v5 >> ((255) & 63))
                                var v43 bool = (!false)
                                v1 -= 155
                                _ = fi42
                                _ = v43
                            }
                            if (-1 == v1) {
                                switch fi41 {
                                case 0:
                                    var v44 bool = true
                                    _ = v44
                                case 128:
                                    for fi45 := 0; fi45 < 18; fi45++ {
                                        var v46 uint8 = 128
                                        fmt.Println(v3, fi41, fi41)
                                        _ = fi45
                                        _ = v46
                                    }
                                case 9223372036854775807:
                                    v1 = v1
                                }
                                if (false && true) {
                                    for fi47 := 0; fi47 < 20; fi47++ {
                                        v48 := GenS6{F0: (^(v1))}
                                        v1 += (v1 | 42)
                                        _ = fi47
                                        _ = v48
                                    }
                                    if true {
                                        for fi49 := 0; fi49 < 13; fi49++ {
                                            var v50 bool = true
                                            _ = fi49
                                            _ = v50
                                        }
                                    }
                                    _ = len(v4)
                                }
                            } else {
                                v3 = v3
                            }
                            _ = fi41
                        }
                    } else {
                        v3 -= (v5 & (v3 * v5))
                        var v51 float64 = (0.1 / ((0.0 / 3.14159) + (1e10 + ((-2.25 / -2.25) / 1.0))))
                        delete(v4, 32768)
                        _ = v51
                    }
                }
                fmt.Println(v5, v1)
            }
        case 128:
            switch "hello" {
            case "hello":
                v1 = v1
            }
        }
    case "\n\t":
        if (391 < 359) {
            fmt.Println(v1, v3, v1)
        }
    default:
        for fi52 := 0; fi52 < 1; fi52++ {
            var v53 uint8 = 1
            var v54 int32 = v1
            for fi55 := 0; fi55 < 10; fi55++ {
                v3 = (v3 ^ (v3 ^ genFn1((v53 << ((0) & 63)), 1)))
                _ = fi55
            }
            _ = fi52
            _ = v53
            _ = v54
        }
    }
    fmt.Println(v1, v3, v1)
    var v56 int = 32767
    var v57 bool = (608 < genFn0(42, ""))
    switch "foo bar" {
    case "foo bar":
        _ = len(v2)
    case "a":
        v56 -= (v56 + 127)
    case "\n\t":
        var v58 uint64 = 1
        _ = v58
    default:
        v59 := GenS4{F0: 1}
        _ = v59
    }
    var v60 float32 = 1e10
    if (!false) {
        _ = len(v2)
    } else {
        if v57 {
            fmt.Println(v56)
            var v61 int = (^(v56))
            var v62 int16 = -32768
            _ = v61
            _ = v62
        }
        v3 = (v5 / ((v5) | 1))
        for fi63 := 0; fi63 < 1; fi63++ {
            var v64 bool = (true && ((v5 % ((v3) | 1)) != genFn1(0, 4294967295)))
            var v65 int = (fi63 & (fi63 * fi63))
            for fi66 := 0; fi66 < 9; fi66++ {
                switch ("a" + ("foo bar" + "\n\t")) {
                case "\n\t":
                    var v67 uint32 = 1
                    _ = v67
                case "a":
                    var v68 int16 = 42
                    _ = v68
                case "foo bar":
                    v4[(fi66 << ((0) & 63))] = 715
                default:
                    if false {
                        v4[(^(fi66))] = fi63
                    }
                }
                v57 = v64
                fmt.Println(v57, v65)
                var v69 uint16 = 249
                _ = fi66
                _ = v69
            }
            _ = fi63
            _ = v64
            _ = v65
        }
    }
    var v70 int32 = (v1 / ((42) | 1))
    fmt.Println(v60, v60)
    v71 := GenS6{F0: v1}
    switch v56 {
    case 0:
        var v72 uint = 523
        _ = v72
    default:
        var v73 uint8 = 197
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v60:", v60)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v60
    _ = v70
    _ = v71
}
