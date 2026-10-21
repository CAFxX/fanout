package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int64; F1 int64; F2 int8 }
func genFn0(p0 uint32, p1 string) uint8 {
    v5 := GenS1{F0: 42}
    var v6 uint = 0
    fmt.Println(p0, p1, p0)
    var v7 uint64 = 2
    var v8 uint8 = 42
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 int8 = 0
    var v2 int8 = v1
    var v3 float32 = (3.14159 / 1e10)
    var v4 int = 127
    var v9 int64 = 1
    switch "\n\t" {
    case "a":
        var v10 uint16 = 32767
        _ = v10
    case "\n\t":
        var v11 uint = 4294967295
        _ = v11
    }
    var v12 uint64 = 42
    v13 := []int{v4, 342, -805, 2147483647}
    fmt.Println(v9)
    if true {
        for fi14 := 0; fi14 < 9; fi14++ {
            _ = len(v13)
            _ = fi14
        }
    }
    fmt.Println(v9)
    switch (v4 << ((514) & 63)) {
    case -36:
        if (0.1 == ((1.5 * (3.14159 / (1e10 * (float64(float32((1e10 / 1.5))) - -2.25)))) * -2.25)) {
            if true {
                switch "a" {
                case "":
                    if ((!((^(v12)) >= (^(v12)))) || (779 > 567)) {
                        v12 += v12
                    } else {
                        var v15 uint8 = 2
                        fmt.Println(v12)
                        _ = v15
                    }
                }
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 bool = (1.5 <= (v3 + v3))
                    _ = fi16
                    _ = v17
                }
            } else {
                if len(v13) > 0 {
                    v13[0] = (-(v4))
                }
                var v18 float64 = 1e10
                for fi19 := 0; fi19 < 20; fi19++ {
                    switch ("hello" + "hello") {
                    case "":
                        _ = len(v13)
                    case "hello":
                        if false {
                            for fi20 := 0; fi20 < 2; fi20++ {
                                if true {
                                    v21 := GenS1{F0: 42}
                                    v2 += int8(v9)
                                    var v22 bool = (true || (0 < v2))
                                    _ = v21
                                    _ = v22
                                } else {
                                    v13 = append(v13, (v4 << ((2) & 63)))
                                    switch 42 {
                                    case 1:
                                        var v23 int32 = -1
                                        _ = v23
                                    case 2147483647:
                                        var v24 uint8 = 255
                                        _ = v24
                                    }
                                    switch "\n\t" {
                                    case "a":
                                        if len(v13) > 0 {
                                            v13[0] = 205
                                        }
                                    case "foo bar":
                                        var v25 int8 = v1
                                        _ = v25
                                    case "":
                                        if len(v13) > 0 {
                                            v13[0] = fi19
                                        }
                                    default:
                                        if (true && false) {
                                            var v26 uint64 = v12
                                            if false {
                                                switch "\n\t" {
                                                case "":
                                                    _ = len(v13)
                                                case "foo bar":
                                                    v18 = v18
                                                default:
                                                    v27 := GenS1{F0: -1000}
                                                    _ = v27
                                                }
                                            } else {
                                                if len(v13) > 0 {
                                                    v13[0] = (fi19 & fi20)
                                                }
                                                v12 += v26
                                            }
                                            var v28 string = "a"
                                            _ = v26
                                            _ = v28
                                        }
                                    }
                                }
                                _ = fi20
                            }
                            for fi29 := 0; fi29 < 2; fi29++ {
                                v30 := GenS2{F0: v4, F1: 117}
                                v3 = ((v3 + float32(1.5)) / float32((float64((v3 / (v3 - (v3 / 1.0)))) / v18)))
                                _ = fi29
                                _ = v30
                            }
                            if (-525 >= 0) {
                                var v31 int16 = -32768
                                var v32 uint16 = 0
                                _ = v31
                                _ = v32
                            }
                        }
                    case "\n\t":
                        v12 -= 0
                    }
                    _ = len(v13)
                    _ = fi19
                }
                _ = v18
            }
            v2 = (v2 / ((v2) | 1))
            v9 = v9
        } else {
            for fi33 := 0; fi33 < 8; fi33++ {
                if (!true) {
                    if (!false) {
                        var v34 string = "a"
                        _ = v34
                    }
                    if true {
                        v35 := GenS1{F0: -776}
                        v12 += v12
                        _ = v35
                    }
                    if (941 >= int16(v2)) {
                        v2 = v2
                        v2 -= v2
                        switch "" {
                        case "foo bar":
                            for fi36 := 0; fi36 < 20; fi36++ {
                                switch fi36 {
                                case -1:
                                    v37 := GenS1{F0: 317}
                                    _ = v37
                                default:
                                    v38 := GenS2{F0: int(v9), F1: (v2 + v1)}
                                    _ = v38
                                }
                                v39 := GenS3{F0: 1, F1: v9, F2: v2}
                                _ = fi36
                                _ = v39
                            }
                        case "\n\t":
                            v40 := GenS1{F0: 42}
                            _ = v40
                        }
                    } else {
                        var v41 float64 = (((0.1 + float64(v3)) / (0.0 - ((1e10 * (0.1 + 1e10)) - -2.25))) / 3.14159)
                        _ = v41
                    }
                } else {
                    _ = len(v13)
                    if true {
                        for fi42 := 0; fi42 < 1; fi42++ {
                            if (true || (true || true)) {
                                var v43 uint8 = 2
                                v44 := GenS2{F0: 596, F1: v1}
                                if len(v13) > 0 {
                                    v13[0] = int(v43)
                                }
                                switch "" {
                                case "":
                                    if (((float64(3.14159) - 1e10) - float64(0.1)) != 1.5) {
                                        v13 = append(v13, (v4 | (v4 & fi42)))
                                        if (!true) {
                                            var v45 float64 = 0.0
                                            fmt.Println(v1, v43, v12)
                                            v4 += 128
                                            _ = v45
                                        } else {
                                            if true {
                                                fmt.Println(v4, v9, v3)
                                            }
                                            switch 32767 {
                                            case -9223372036854775808:
                                                var v46 uint8 = (^(v43))
                                                _ = v46
                                            case -2147483648:
                                                var v47 int8 = v2
                                                v48 := GenS1{F0: 524}
                                                var v49 float32 = (((((0.1 - (v3 * (3.14159 - 1.5))) + (v3 / (v3 / 1e10))) + (v3 / 0.1)) / v3) - v3)
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                            }
                                        }
                                    }
                                case "hello":
                                    v43 = v43
                                }
                                _ = v43
                                _ = v44
                            } else {
                                v12 += (v12 * 4294967295)
                                if (!(0.0 > v3)) {
                                    v12 = 1
                                }
                            }
                            _ = fi42
                        }
                        if false {
                            var v50 int = v4
                            _ = v50
                        } else {
                            v4 -= 0
                            var v51 int8 = v1
                            _ = v51
                        }
                    }
                }
                v1 = (v1 * int8(v12))
                fmt.Println(v3, v12, v3)
                _ = fi33
            }
        }
    default:
        fmt.Println(v12, v9)
    }
    for fi52 := 0; fi52 < 19; fi52++ {
        var v53 uint16 = 633
        _ = fi52
        _ = v53
    }
    var v54 float32 = v3
    if (genFn0(42, "\n\t") >= 57) {
        var v55 int32 = -1
        for fi56 := 0; fi56 < 5; fi56++ {
            v4 -= 319
            for fi57 := 0; fi57 < 19; fi57++ {
                fmt.Println(v4, v12)
                _ = fi57
            }
            fmt.Println(fi56, v2, v12)
            _ = fi56
        }
        _ = v55
    }
    _ = len(v13)
    v58 := GenS2{F0: v4, F1: (v2 % (((v2 + 127)) | 1))}
    if false {
        var v59 int64 = (v9 - 694)
        if true {
            if true {
                if (-1 != 0) {
                    v1 = -128
                    var v60 uint64 = v12
                    _ = v60
                } else {
                    fmt.Println(v2, v54, v54)
                }
            }
        } else {
            var v61 uint64 = (v12 + v12)
            _ = v61
        }
        _ = v59
    }
    switch v4 {
    case 1:
        var v62 uint8 = 127
        _ = v62
    case 0:
        var v63 int = 0
        _ = v63
    case -1:
        var v64 int64 = v9
        _ = v64
    default:
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v58
    _ = v9
}
