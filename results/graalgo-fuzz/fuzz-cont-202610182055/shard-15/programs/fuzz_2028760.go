package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 bool; F1 bool; F2 string }
type GenS6 struct { F0 uint16; F1 int8; F2 uint }
type GenS7 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS8 struct { F0 int64; F1 uint32 }
type GenS9 struct { F0 uint8; F1 uint; F2 bool }
type GenS10 struct { F0 uint8; F1 uint8; F2 string }

func main() {
    var v1 int8 = -128
    v2 := make(map[int]bool)
    var v3 float32 = 3.14159
    var v4 uint64 = 534
    for fi5 := 0; fi5 < 11; fi5++ {
        switch "hello" {
        case "":
            var v6 string = "hello"
            _ = v6
        case "\n\t":
            var v7 int32 = 1
            _ = v7
        }
        v4 = (v4 >> ((705) & 63))
        _ = fi5
    }
    if (true && false) {
        fmt.Println(v4, v3)
        if (790 >= 1) {
            switch "a" {
            case "foo bar":
                v4 = 563
            case "\n\t":
                switch 1 {
                case -788:
                    switch -9223372036854775808 {
                    case 4294967295:
                        v3 -= v3
                    case 127:
                        v8 := GenS1{F0: false}
                        _ = v8
                    default:
                        _ = len(v2)
                    }
                case 32768:
                    if false {
                        v9 := GenS1{F0: true}
                        if true {
                            fmt.Println(v3)
                            v10 := GenS2{F0: 749}
                            for fi11 := 0; fi11 < 15; fi11++ {
                                v12 := GenS3{F0: "hello"}
                                _ = fi11
                                _ = v12
                            }
                            _ = v10
                        }
                        _ = v9
                    }
                default:
                    if (1e10 < (1e10 + ((float64(v3) / -2.25) * 0.1))) {
                        v13 := []int{-438, 0, -23, 546}
                        _ = v13
                    }
                }
            case "":
                v14 := []int{-1, 0, 9223372036854775807, 32768}
                _ = v14
            }
            v1 = (v1 & v1)
        }
        var v15 int64 = 42
        _ = v15
    } else {
        var v16 int16 = 1000
        _ = v16
    }
    v17 := []int{127, 127, -159}
    v18 := GenS2{F0: 1}
    fmt.Println(v3, v4, v4)
    for fi19 := 0; fi19 < 10; fi19++ {
        v3 -= 3.14159
        if false {
            if (false || false) {
                fmt.Println(v3, fi19, v3)
            } else {
                for fi20 := 0; fi20 < 13; fi20++ {
                    delete(v2, -1)
                    v21 := GenS4{F0: -796, F1: ""}
                    _ = fi20
                    _ = v21
                }
                _ = len(v2)
                v17 = append(v17, fi19)
            }
            v4 += (v4 >> ((65535) & 63))
        } else {
            var v22 uint64 = (v4 & 18446744073709551615)
            var v23 float32 = 0.1
            if true {
                fmt.Println(fi19, fi19, fi19)
                var v24 uint8 = 1
                _ = v24
            }
            _ = v22
            _ = v23
        }
        fmt.Println(v1)
        _ = fi19
    }
    switch 32768 {
    case -9223372036854775808:
        v25 := GenS1{F0: false}
        _ = v25
    case 9223372036854775807:
        v26 := v2[int(v1)]
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        if (!(("a" > "hello") && true)) {
            v28 := GenS5{F0: false, F1: ("a" > ""), F2: ("" + ("hello" + "a"))}
            for fi29 := 0; fi29 < 3; fi29++ {
                v30 := GenS6{F0: 42, F1: 1, F2: 2}
                switch "hello" {
                case "hello":
                    switch "a" {
                    case "":
                        var v31 string = "hello"
                        _ = v31
                    case "\n\t":
                        v32 := GenS7{F0: 19, F1: 0, F2: (v4 % ((v4) | 1))}
                        _ = v32
                    case "a":
                        v4 = (^(v4))
                    }
                default:
                    fmt.Println(v1)
                }
                if true {
                    switch "a" {
                    case "hello":
                        if false {
                            fmt.Println(fi27, v1, v1)
                            var v33 int32 = -2147483648
                            _ = v33
                        } else {
                            var v34 uint64 = v4
                            v4 += (v34 + (v4 - (^(v34))))
                            v1 += v1
                            _ = v34
                        }
                        fmt.Println(v3, fi27, fi29)
                        v2[(fi29 % (((fi29 - (^(fi27)))) | 1))] = (!false)
                    case "a":
                        fmt.Println(fi29)
                    default:
                        var v35 string = (("\n\t" + ("hello" + "")) + ("\n\t" + ""))
                        _ = v35
                    }
                    switch fi29 {
                    case 32768:
                        if false {
                            if false {
                                delete(v2, (fi29 | (fi29 / ((fi29) | 1))))
                                switch (fi27 & (fi27 + (fi29 - (fi29 / ((534) | 1))))) {
                                case 9223372036854775807:
                                    var v36 float64 = 1.5
                                    _ = v36
                                }
                                v3 -= float32(-2.25)
                                if (v3 < v3) {
                                    if ((true || ((v4 * 42) >= (v4 % (((^(v4))) | 1)))) && (!true)) {
                                        var v37 int32 = 2147483647
                                        v37 = (v37 ^ v37)
                                        v38 := GenS6{F0: 42, F1: (-(v1)), F2: 65535}
                                        _ = v37
                                        _ = v38
                                    } else {
                                        if true {
                                            var v39 uint16 = 32767
                                            var v40 uint32 = 1
                                            v2[fi27] = false
                                            if (!true) {
                                                v40 += v40
                                                var v41 int8 = v1
                                                if true {
                                                    v42 := GenS8{F0: -1, F1: (v40 % ((v40) | 1))}
                                                    _ = v42
                                                } else {
                                                    v1 = (v41 ^ -128)
                                                }
                                                _ = v41
                                            } else {
                                                v1 -= v1
                                            }
                                            v40 = (v40 ^ (v40 ^ (v40 | 299)))
                                            _ = v39
                                            _ = v40
                                        } else {
                                            switch ("\n\t" + "\n\t") {
                                            case "a":
                                                _ = len(v17)
                                            case "foo bar":
                                                v3 += v3
                                            case "":
                                                v43 := GenS9{F0: 215, F1: 65535, F2: false}
                                                _ = v43
                                            }
                                        }
                                    }
                                    var v44 float32 = v3
                                    var v45 uint16 = 1
                                    _ = v44
                                    _ = v45
                                } else {
                                    v3 -= 1.5
                                    delete(v2, (fi29 >> ((790) & 63)))
                                    switch "a" {
                                    case "":
                                        v46 := GenS8{F0: 914, F1: 2147483648}
                                        _ = v46
                                    }
                                }
                                v47 := GenS6{F0: 32767, F1: (v1 / ((100) | 1)), F2: 1}
                                _ = v47
                            }
                            v3 -= v3
                            v3 = 3.14159
                            if (false && (-497 > -1)) {
                                v2[fi27] = (uint16(v4) > 0)
                            } else {
                                fmt.Println(fi29, fi27)
                            }
                            if true {
                                var v48 int8 = (v1 + (v1 >> ((65535) & 63)))
                                var v49 uint64 = 0
                                _ = v48
                                _ = v49
                            }
                            v4 = v4
                            var v50 uint = 719
                            _ = v50
                        }
                    case 624:
                        switch "a" {
                        case "foo bar":
                            v1 -= (^(v1))
                        default:
                            if true {
                                var v51 float64 = (((-2.25 + 0.1) + 1e10) - (0.1 / ((1.5 + (-2.25 * (float64(1.5) / 1.5))) + (1.5 * 0.0))))
                                _ = v51
                            } else {
                                if ((true || false) && ((fi29 | fi29) != (^(fi27)))) {
                                    v52 := v2[fi27]
                                    v53 := GenS5{F0: v52, F1: ("a" < "\n\t"), F2: ("foo bar" + "")}
                                    _ = v52
                                    _ = v53
                                } else {
                                    fmt.Println(v3)
                                    if len(v17) > 0 {
                                        v17[0] = fi29
                                    }
                                }
                                delete(v2, (^(fi27)))
                            }
                        }
                    default:
                        fmt.Println(fi29)
                    }
                    v54 := GenS10{F0: 2, F1: 42, F2: "foo bar"}
                    if true {
                        v4 += 42
                        v4 += (v4 >> ((65535) & 63))
                        switch (("hello" + "\n\t") + "\n\t") {
                        case "a":
                            var v55 float64 = 0.1
                            _ = v55
                        case "hello":
                            v4 = v4
                        case "foo bar":
                            v3 = (float32(3.14159) / (((((1e10 * v3) / 1.5) - v3) - (0.0 - (v3 - float32(3.14159)))) / 1.5))
                        default:
                            if false {
                                v56 := GenS10{F0: 128, F1: 219, F2: ""}
                                _ = v56
                            }
                        }
                    }
                    _ = v54
                } else {
                    var v57 uint16 = 2
                    v3 += ((v3 - float32(-2.25)) + v3)
                    _ = v57
                }
                _ = fi29
                _ = v30
            }
            var v58 int8 = (-(v1))
            _ = v28
            _ = v58
        }
        v59 := GenS8{F0: 0, F1: 390}
        fmt.Println(v3)
        _ = fi27
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
