package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 int; F2 int64 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint8; F1 string; F2 bool }
func genFn0() uint8 {
    var v6 uint16 = 42
    v7 := []int{740, -749}
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 string = ("\n\t" + "\n\t")
    var v2 uint32 = 877
    var v3 float32 = 1e10
    var v4 int8 = 1
    var v5 uint64 = 1
    var v8 int64 = 245
    fmt.Println(v5)
    switch v1 {
    case "\n\t":
        var v9 uint8 = 42
        _ = v9
    case "a":
        v10 := GenS1{F0: 1, F1: uint32(v5), F2: "\n\t"}
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := []int{(fi11 % (((^(fi11))) | 1))}
        v13 := GenS1{F0: 551, F1: v2, F2: v1}
        _ = fi11
        _ = v12
        _ = v13
    }
    if (v3 == float32(((0.1 + float64(v3)) + ((0.0 * -2.25) + 0.0)))) {
        v14 := GenS2{F0: (v5 % ((v5) | 1))}
        v5 -= 2
        _ = v14
    }
    v15 := []int{-2147483648}
    if false {
        var v16 int = -2147483648
        v17 := GenS3{F0: 695, F1: v16, F2: -1}
        if false {
            v18 := GenS2{F0: v5}
            v19 := GenS3{F0: 335, F1: (-(v16)), F2: (v8 >> ((1) & 63))}
            _ = v18
            _ = v19
        } else {
            fmt.Println(v8, v2)
        }
        _ = v16
        _ = v17
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        switch v1 {
        case "a":
            switch 185 {
            case 1:
                switch fi20 {
                case 4294967295:
                    for fi21 := 0; fi21 < 14; fi21++ {
                        switch fi21 {
                        case -501:
                            v3 = v3
                        default:
                            v22 := GenS4{F0: v4, F1: 1000}
                            _ = v22
                        }
                        _ = len(v15)
                        _ = fi21
                    }
                }
            case 128:
                if false {
                    v4 = 1
                    if false {
                        for fi23 := 0; fi23 < 20; fi23++ {
                            v5 += (v5 << ((485) & 63))
                            if ((!false) && true) {
                                v24 := GenS5{F0: v1}
                                _ = v24
                            } else {
                                v8 -= 9223372036854775807
                                if (0 > -255) {
                                    var v25 float32 = (((float32(0.0) + 0.1) + v3) / float32((1.5 + 1.5)))
                                    var v26 uint16 = 108
                                    var v27 int32 = 880
                                    v28 := GenS3{F0: (v26 ^ v26), F1: fi20, F2: (v8 % ((v8) | 1))}
                                    switch 32767 {
                                    case 1:
                                        fmt.Println(v1)
                                    case 2147483647:
                                        var v29 uint8 = 16
                                        v30 := GenS1{F0: 2, F1: (v2 + v2), F2: v1}
                                        _ = v29
                                        _ = v30
                                    default:
                                        v3 = 0.1
                                    }
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                                v4 += v4
                            }
                            v5 -= v5
                            _ = fi23
                        }
                        for fi31 := 0; fi31 < 11; fi31++ {
                            if (false && ((!true) && (true || true))) {
                                switch (fi31 ^ (-(fi31))) {
                                case -312:
                                    v8 = (v8 * v8)
                                case 9223372036854775807:
                                    var v32 uint8 = 0
                                    _ = v32
                                case 4294967295:
                                    var v33 int = (fi31 % ((-2147483648) | 1))
                                    _ = v33
                                default:
                                    if (!((v2 >> ((uint(v5)) & 63)) != (v2 >> ((610) & 63)))) {
                                        if true {
                                            fmt.Println(fi31, fi20, v8)
                                            if false {
                                                v8 += 432
                                                v34 := GenS6{F0: (((v1 + (v1 + "\n\t")) + v1) + "")}
                                                _ = len(v15)
                                                v35 := GenS4{F0: (v4 - (^(v4))), F1: 32767}
                                                _ = v34
                                                _ = v35
                                            } else {
                                                var v36 uint = 42
                                                v1 = v1
                                                switch "foo bar" {
                                                case "":
                                                    v5 += v5
                                                default:
                                                    var v37 uint32 = 209
                                                    _ = v37
                                                }
                                                _ = v36
                                            }
                                            var v38 uint64 = v5
                                            _ = v38
                                        } else {
                                            v39 := GenS7{F0: -734, F1: genFn0()}
                                            v40 := GenS7{F0: 362, F1: 0}
                                            if len(v15) > 0 {
                                                v15[0] = (^(fi20))
                                            }
                                            _ = v39
                                            _ = v40
                                        }
                                        v15 = append(v15, 32767)
                                        v1 = (v1 + "hello")
                                        fmt.Println(v4, v8)
                                        var v41 int64 = v8
                                        _ = v41
                                    } else {
                                        v42 := GenS3{F0: 0, F1: -746, F2: (v8 ^ 518)}
                                        var v43 uint64 = (v5 ^ (v5 + (v5 & (^(v5)))))
                                        _ = v42
                                        _ = v43
                                    }
                                }
                                fmt.Println(fi20, v8)
                                switch fi31 {
                                case 2147483647:
                                    switch (fi31 + (fi31 << ((991) & 63))) {
                                    case 247:
                                        v44 := GenS8{F0: 42, F1: 4294967295}
                                        _ = v44
                                    case -2147483648:
                                        v45 := GenS9{F0: (false && ("hello" >= (v1 + "")))}
                                        _ = v45
                                    case 974:
                                        v4 += (v4 >> ((37) & 63))
                                    default:
                                        v1 = ("a" + v1)
                                    }
                                case 128:
                                    v46 := GenS10{F0: 65535}
                                    _ = v46
                                default:
                                    var v47 uint64 = v5
                                    _ = v47
                                }
                            }
                            var v48 float64 = -2.25
                            _ = fi31
                            _ = v48
                        }
                        var v49 int8 = (-(v4))
                        _ = v49
                    }
                } else {
                    if false {
                        fmt.Println(v1, v1, v5)
                    } else {
                        var v50 uint64 = v5
                        _ = v50
                    }
                    v8 += v8
                    for fi51 := 0; fi51 < 11; fi51++ {
                        v1 = ""
                        if len(v15) > 0 {
                            v15[0] = 1
                        }
                        _ = fi51
                    }
                }
            }
        case "\n\t":
            v15 = append(v15, (fi20 | -9223372036854775808))
        default:
            fmt.Println(v1, v2)
        }
        var v52 uint = 255
        v5 -= 42
        _ = fi20
        _ = v52
    }
    v53 := GenS11{F0: 255, F1: (v1 + v1), F2: (!false)}
    fmt.Println(v4, v4, v5)
    var v54 string = (v1 + "\n\t")
    for fi55 := 0; fi55 < 14; fi55++ {
        v56 := GenS6{F0: ((v54 + v1) + ((((((v54 + "\n\t") + (v1 + v54)) + v54) + v54) + (v1 + v1)) + (v54 + (v1 + ((v1 + ((v54 + "foo bar") + "\n\t")) + "\n\t")))))}
        v1 = "\n\t"
        v5 -= v5
        _ = fi55
        _ = v56
    }
    var v57 int32 = -1
    v8 = (v8 >> ((65535) & 63))
    v8 += v8
    switch ((("hello" + ((("a" + v54) + "") + v1)) + v1) + ((v54 + (v54 + "a")) + (v54 + ((v1 + "\n\t") + v1)))) {
    case "foo bar":
        var v58 uint32 = (v2 + v2)
        _ = v58
    case "\n\t":
        if (42 >= 2) {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            v3 += v3
            var v59 int32 = 1
            _ = v59
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v57
    _ = v8
}
