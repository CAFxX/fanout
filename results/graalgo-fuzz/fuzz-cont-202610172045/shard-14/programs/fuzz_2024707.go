package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int16; F2 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64 }
func genFn0(p0 uint, p1 bool) uint16 {
    var v6 uint32 = 529
    v6 -= 488
    var v7 uint32 = (v6 ^ 4294967295)
    p1 = (2 == 36)
    var v8 int32 = 42
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint8 = 158
    var v2 int = 770
    v3 := []string{((("foo bar" + ("" + "")) + "hello") + (((("\n\t" + "") + (("hello" + "\n\t") + "")) + "") + "foo bar")), "hello", ("hello" + ("" + "hello")), "\n\t"}
    v4 := make(map[string]bool)
    var v5 int64 = -1
    if (true || (90 >= 1)) {
        for fi9 := 0; fi9 < 17; fi9++ {
            if true {
                v10 := GenS1{F0: 560, F1: 127, F2: 155}
                v1 = 73
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi11 := 0; fi11 < 6; fi11++ {
                        v3 = append(v3, ("a" + ("\n\t" + "hello")))
                        _ = fi11
                    }
                    for fi12 := 0; fi12 < 3; fi12++ {
                        if (0 > -226) {
                            fmt.Println(fi9, fi12)
                            var v13 float64 = (1.5 + ((3.14159 * 1e10) / 1.5))
                            _ = v13
                        } else {
                            v2 = (fi9 + (fi9 / (((fi9 % ((761) | 1))) | 1)))
                            if true {
                                v4[""] = (!(!((float32((0.1 - -2.25)) + 1e10) == float32(1e10))))
                                fmt.Println(v1)
                                v2 = (v2 >> ((18446744073709551615) & 63))
                                if (false || (0.1 < 0.0)) {
                                    var v14 string = "a"
                                    var v15 uint8 = 165
                                    _ = v14
                                    _ = v15
                                } else {
                                    fmt.Println(fi9)
                                    v5 += -9223372036854775808
                                    v5 = v5
                                }
                                if false {
                                    if (2147483647 <= 2147483648) {
                                        v5 = v5
                                        var v16 uint64 = 4294967295
                                        _ = len(v3)
                                        v2 -= fi9
                                        _ = v16
                                    } else {
                                        switch ("\n\t" + "hello") {
                                        case "foo bar":
                                            _ = len(v4)
                                        case "\n\t":
                                            switch ("hello" + "\n\t") {
                                            case "\n\t":
                                                if true {
                                                    v2 = fi12
                                                    var v17 int64 = 9223372036854775807
                                                    _ = v17
                                                } else {
                                                    switch ((("hello" + "\n\t") + "") + "") {
                                                    case "\n\t":
                                                        var v18 uint32 = 4294967295
                                                        _ = v18
                                                    case "foo bar":
                                                        v19 := GenS1{F0: 810, F1: -65, F2: 42}
                                                        _ = v19
                                                    case "a":
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v3 = append(v3, "\n\t")
                                                        case "hello":
                                                            if true {
                                                                var v20 float32 = 1e10
                                                                if true {
                                                                    v2 -= (-(fi12))
                                                                    delete(v4, (("a" + "\n\t") + ("\n\t" + "hello")))
                                                                }
                                                                fmt.Println(v2)
                                                                _ = v20
                                                            }
                                                        case "\n\t":
                                                            v21 := GenS2{F0: 2}
                                                            _ = v21
                                                        }
                                                    default:
                                                        var v22 int64 = (v5 + 9223372036854775807)
                                                        _ = v22
                                                    }
                                                    v23 := GenS1{F0: 0, F1: 65, F2: 2}
                                                    _ = v23
                                                }
                                            default:
                                                v24 := GenS3{F0: 47, F1: 32767, F2: 2147483647}
                                                _ = v24
                                            }
                                        case "hello":
                                            if false {
                                                var v25 uint64 = 4294967295
                                                _ = v25
                                            }
                                        }
                                        v26 := GenS4{F0: 0}
                                        _ = v26
                                    }
                                    var v27 int8 = 9
                                    _ = v27
                                }
                                v28 := GenS3{F0: 65535, F1: -292, F2: 2147483647}
                                delete(v4, "")
                                _ = v28
                            } else {
                                var v29 int32 = -1
                                var v30 uint8 = (v1 ^ v1)
                                v29 -= -875
                                fmt.Println(v2)
                                v31 := GenS1{F0: 0, F1: -37, F2: 0}
                                var v32 float64 = 0.0
                                switch ("a" + "") {
                                case "foo bar":
                                    var v33 uint32 = 0
                                    _ = v33
                                }
                                v1 = (^(v1))
                                _ = v29
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                        }
                        fmt.Println(v2, fi9, fi9)
                        var v34 float32 = -2.25
                        _ = fi12
                        _ = v34
                    }
                }
                _ = v10
            }
            if true {
                v2 = v2
                v5 = (v5 & (-(v5)))
                v35 := GenS5{F0: v1, F1: 42}
                _ = v35
            }
            var v36 float64 = 1.5
            _ = fi9
            _ = v36
        }
        v4["foo bar"] = (false || (!(42 <= 42)))
    } else {
        var v37 int = v2
        _ = v37
    }
    delete(v4, "")
    v5 = (v5 / (((-(v5))) | 1))
    if false {
        fmt.Println(v2, v2, v5)
        fmt.Println(v1, v2, v2)
        for fi38 := 0; fi38 < 13; fi38++ {
            v5 -= v5
            v39 := GenS2{F0: 65535}
            var v40 uint32 = 1
            _ = fi38
            _ = v39
            _ = v40
        }
    } else {
        switch "a" {
        case "\n\t":
            v1 = (v1 | (v1 >> ((4294967295) & 63)))
        case "":
            var v41 uint16 = 42
            _ = v41
        case "hello":
            var v42 uint8 = (v1 - (v1 << ((255) & 63)))
            _ = v42
        }
        fmt.Println(v1, v2)
    }
    if (v5 == -9223372036854775808) {
        v5 = v5
    } else {
        v4["\n\t"] = false
    }
    v5 += v5
    switch v2 {
    case 4294967295:
        for fi43 := 0; fi43 < 19; fi43++ {
            for fi44 := 0; fi44 < 15; fi44++ {
                if false {
                    if (false || (1.5 == 0.1)) {
                        switch "" {
                        case "\n\t":
                            v45 := GenS1{F0: 2147483647, F1: -128, F2: 0}
                            _ = v45
                        }
                    }
                    fmt.Println(fi44, fi43, v5)
                }
                _ = fi44
            }
            if true {
                delete(v4, ("" + "\n\t"))
                v46 := GenS4{F0: 32767}
                var v47 uint16 = 608
                _ = v46
                _ = v47
            } else {
                switch "a" {
                case "":
                    v5 = v5
                case "hello":
                    switch "" {
                    case "hello":
                        switch "hello" {
                        case "foo bar":
                            if false {
                                var v48 int16 = 1
                                _ = v48
                            }
                        case "hello":
                            _ = len(v4)
                        }
                    case "a":
                        if false {
                            switch ("\n\t" + "") {
                            case "foo bar":
                                for fi49 := 0; fi49 < 5; fi49++ {
                                    var v50 uint = 326
                                    _ = fi49
                                    _ = v50
                                }
                            default:
                                v51 := GenS5{F0: v1, F1: 937}
                                _ = v51
                            }
                        } else {
                            fmt.Println(fi43, v1)
                        }
                    case "foo bar":
                        v52 := GenS1{F0: 2147483648, F1: -75, F2: 2}
                        _ = v52
                    }
                case "foo bar":
                    v53 := GenS6{F0: 1}
                    _ = v53
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v2 = 1
            }
            v1 = 2
            _ = fi43
        }
    case 0:
        switch "foo bar" {
        case "":
            v54 := GenS2{F0: 0}
            _ = v54
        }
    }
    for fi55 := 0; fi55 < 17; fi55++ {
        for fi56 := 0; fi56 < 16; fi56++ {
            var v57 int = 128
            _ = fi56
            _ = v57
        }
        if (!(1.5 > float32(1.5))) {
            switch "foo bar" {
            case "\n\t":
                delete(v4, "")
            case "hello":
                var v58 uint64 = 0
                _ = v58
            default:
                for fi59 := 0; fi59 < 15; fi59++ {
                    v5 += (-(v5))
                    v2 -= (v2 ^ fi55)
                    v3 = append(v3, ("hello" + ("a" + "a")))
                    _ = fi59
                }
            }
            v2 = fi55
        }
        if false {
            v2 += fi55
            v60 := GenS7{F0: v5}
            v2 -= (-(v2))
            _ = v60
        }
        _ = fi55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
