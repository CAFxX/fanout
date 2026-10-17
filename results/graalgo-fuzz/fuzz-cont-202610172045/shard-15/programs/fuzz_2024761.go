package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 int64) uint32 {
    var v7 string = "a"
    if true {
        var v8 float64 = 0.0
        if true {
            v9 := []int{}
            if (false || true) {
                for fi10 := 0; fi10 < 19; fi10++ {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    v8 -= v8
                    _ = fi10
                }
                if len(v9) > 0 {
                    v9[0] = 791
                }
                if (!(v7 >= "foo bar")) {
                    _ = len(v9)
                    var v11 int16 = -258
                    var v12 int32 = -2147483648
                    _ = v11
                    _ = v12
                } else {
                    if ((false || (false && true)) && true) {
                        v13 := GenS1{F0: v7, F1: 42}
                        p0 = (p0 | p0)
                        _ = v13
                    } else {
                        v7 = "hello"
                    }
                    v8 -= float64((-2.25 + 1e10))
                    p0 = (p0 * -1)
                }
            } else {
                var v14 int8 = 42
                for fi15 := 0; fi15 < 13; fi15++ {
                    switch fi15 {
                    case 610:
                        if false {
                            switch fi15 {
                            case -511:
                                if (255 == 1) {
                                    fmt.Println(p0, p0, v7)
                                }
                            case -9223372036854775808:
                                var v16 float32 = 1.5
                                _ = v16
                            case 959:
                                var v17 uint32 = 319
                                _ = v17
                            default:
                                v14 -= v14
                            }
                            v18 := GenS2{F0: v14, F1: 400, F2: 0}
                            _ = v18
                        }
                    default:
                        for fi19 := 0; fi19 < 2; fi19++ {
                            v9 = append(v9, (fi15 | (fi15 << ((65535) & 63))))
                            if true {
                                v20 := GenS2{F0: (v14 / ((v14) | 1)), F1: 0, F2: 4294967295}
                                fmt.Println(v8, v8, fi19)
                                _ = v20
                            } else {
                                v9 = append(v9, 0)
                                _ = len(v9)
                                v9 = append(v9, fi15)
                                var v21 int64 = (p0 + p0)
                                _ = v21
                            }
                            switch v7 {
                            case "hello":
                                var v22 float32 = 1e10
                                _ = v22
                            case "foo bar":
                                switch (v7 + "a") {
                                case "\n\t":
                                    v7 = "hello"
                                }
                            case "":
                                if true {
                                    v8 = v8
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    var v23 uint32 = 42
                                    _ = v23
                                }
                            default:
                                var v24 int = fi19
                                _ = v24
                            }
                            _ = fi19
                        }
                    }
                    _ = fi15
                }
                var v25 int32 = 42
                _ = v14
                _ = v25
            }
            v26 := GenS1{F0: (v7 + v7), F1: 42}
            _ = v26
            _ = v9
        } else {
            fmt.Println(v8)
            for fi27 := 0; fi27 < 19; fi27++ {
                p0 -= p0
                v28 := []int{(fi27 ^ (fi27 >> ((2) & 63))), (fi27 / ((fi27) | 1)), int(p0), fi27}
                _ = fi27
                _ = v28
            }
        }
        _ = v8
    }
    _ = p0
    _ = v7
    return 2147483647
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{-9223372036854775808, 42, 32767, 852}
    var v3 int16 = 1000
    var v4 int16 = 34
    v5 := []string{"\n\t", "hello", (("foo bar" + "a") + (("foo bar" + (("foo bar" + "") + ("" + ("hello" + "hello")))) + (("\n\t" + "foo bar") + ("foo bar" + "hello"))))}
    var v6 string = ""
    var v29 uint8 = 71
    var v30 bool = true
    for fi31 := 0; fi31 < 19; fi31++ {
        var v32 float64 = float64(((1e10 + (0.0 - (((1e10 + ((0.1 * 3.14159) + (3.14159 + 0.1))) - -2.25) - ((0.1 + 1.5) * 3.14159)))) - 3.14159))
        if (!v30) {
            switch v6 {
            case "\n\t":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        } else {
            v3 = (v4 % ((v3) | 1))
            var v33 uint = 4294967295
            var v34 int32 = -8
            _ = v33
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    v30 = false
    v29 -= 0
    switch 1 {
    case 0:
        var v35 uint8 = 0
        _ = v35
    case 127:
        var v36 uint = 255
        _ = v36
    }
    var v37 uint32 = genFn0(341)
    v38 := GenS1{F0: "hello", F1: (v37 + (v37 % ((v37) | 1)))}
    if (genFn0(9223372036854775807) == v37) {
        v39 := v1[1]
        _ = v39
    } else {
        var v40 float32 = 0.1
        var v41 uint64 = 4294967295
        _ = v40
        _ = v41
    }
    v5 = append(v5, (v6 + (v6 + v6)))
    switch "hello" {
    case "foo bar":
        if false {
            switch v6 {
            case "a":
                var v42 string = v6
                _ = v42
            default:
                var v43 uint16 = 255
                _ = v43
            }
            var v44 uint64 = 42
            _ = v44
        } else {
            _ = len(v1)
        }
    }
    v6 = "hello"
    switch 735 {
    case -1:
        for fi45 := 0; fi45 < 4; fi45++ {
            if (!(!(v30 && false))) {
                if len(v2) > 0 {
                    v2[0] = fi45
                }
                if (fi45 != 4294967295) {
                    _ = len(v2)
                    if v30 {
                        if true {
                            for fi46 := 0; fi46 < 11; fi46++ {
                                v1[42] = v30
                                v37 = (v37 % ((2) | 1))
                                delete(v1, fi45)
                                _ = fi46
                            }
                        }
                        switch (-(fi45)) {
                        case 28:
                            v47 := GenS2{F0: 42, F1: 1, F2: 274}
                            _ = v47
                        case -539:
                            switch fi45 {
                            case -958:
                                if v30 {
                                    fmt.Println(v30)
                                    v37 += 4294967295
                                }
                            case -238:
                                switch v6 {
                                case "hello":
                                    if false {
                                        v3 -= v4
                                        var v48 uint32 = v37
                                        var v49 uint16 = 1
                                        _ = v48
                                        _ = v49
                                    }
                                default:
                                    v50 := GenS3{F0: fi45, F1: "\n\t"}
                                    _ = v50
                                }
                            case -2147483648:
                                switch "a" {
                                case "":
                                    fmt.Println(v29, v4, v6)
                                case "hello":
                                    var v51 string = ""
                                    _ = v51
                                case "\n\t":
                                    switch (((v6 + v6) + ((v6 + v6) + v6)) + (v6 + ((v6 + (v6 + (v6 + ""))) + "hello"))) {
                                    case "a":
                                        v3 -= -32768
                                    case "":
                                        if len(v2) > 0 {
                                            v2[0] = -1
                                        }
                                    case "\n\t":
                                        var v52 uint64 = 151
                                        _ = v52
                                    default:
                                        fmt.Println(v30, fi45, v4)
                                    }
                                default:
                                    fmt.Println(v3, v37, v3)
                                }
                            }
                        default:
                            switch v6 {
                            case "":
                                v29 += 2
                            case "hello":
                                switch (fi45 ^ (fi45 ^ fi45)) {
                                case 128:
                                    var v53 bool = (208 > uint16(v29))
                                    _ = v53
                                case 9223372036854775807:
                                    var v54 string = v6
                                    _ = v54
                                }
                            case "a":
                                switch (fi45 % ((fi45) | 1)) {
                                case 695:
                                    switch (-(fi45)) {
                                    case 2147483647:
                                        v6 = v6
                                    }
                                case -9223372036854775808:
                                    var v55 int64 = -9223372036854775808
                                    _ = v55
                                case 2147483647:
                                    v37 = v37
                                default:
                                    v56 := GenS2{F0: -9, F1: 167, F2: 4294967295}
                                    _ = v56
                                }
                            }
                        }
                        for fi57 := 0; fi57 < 4; fi57++ {
                            v58 := GenS4{F0: ((v3 | int16(v37)) != v4), F1: (v29 | v29)}
                            _ = fi57
                            _ = v58
                        }
                    }
                    var v59 int8 = 77
                    _ = v59
                } else {
                    v29 += v29
                    var v60 int64 = -1
                    _ = v60
                }
                v29 = (v29 * 128)
            }
            delete(v1, (fi45 << ((2) & 63)))
            if v30 {
                fmt.Println(v37, v4)
            } else {
                for fi61 := 0; fi61 < 18; fi61++ {
                    if true {
                        if true {
                            v29 -= 0
                            var v62 float32 = 0.1
                            v29 -= (v29 << ((0) & 63))
                            _ = v62
                        } else {
                            v29 += 128
                            var v63 uint = 255
                            _ = v63
                        }
                    } else {
                        if (1 != 9223372036854775807) {
                            _ = len(v1)
                            if v30 {
                                switch "hello" {
                                case "\n\t":
                                    v29 += 1
                                case "":
                                    v64 := GenS1{F0: v6, F1: (v37 + 2147483648)}
                                    _ = v64
                                case "foo bar":
                                    v1[(^(fi61))] = v30
                                }
                                _ = len(v5)
                                if v30 {
                                    v1[(-(fi45))] = v30
                                }
                                v37 -= (v37 / ((4294967295) | 1))
                            }
                            var v65 int16 = (v4 >> ((514) & 63))
                            _ = v65
                        }
                        var v66 uint = 255
                        v67 := GenS2{F0: 127, F1: 32767, F2: uint(v29)}
                        _ = v66
                        _ = v67
                    }
                    _ = fi61
                }
            }
            _ = fi45
        }
    case 42:
        fmt.Println(v6, v4)
    case 32768:
        var v68 uint = 256
        _ = v68
    default:
        var v69 int = -2147483648
        _ = v69
    }
    v37 = (^(v37))
    v4 = (^(v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
