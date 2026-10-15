package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 int16; F1 int32 }
func genFn0(p0 int16, p1 uint64) int32 {
    var v4 uint16 = 0
    v5 := GenS1{F0: (true || false), F1: 213}
    v6 := GenS2{F0: 736, F1: -9223372036854775808}
    var v7 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 int32 = -256
    var v2 int64 = 9223372036854775807
    var v3 int32 = (v1 / (((v1 + (^(v1)))) | 1))
    switch (("foo bar" + "hello") + ("\n\t" + "foo bar")) {
    case "hello":
        if true {
            switch "" {
            case "foo bar":
                v8 := []int{-9223372036854775808}
                _ = v8
            case "a":
                if true {
                    var v9 float32 = (1e10 / 3.14159)
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 2; fi10++ {
                        for fi11 := 0; fi11 < 19; fi11++ {
                            v12 := []int{}
                            v13 := GenS3{F0: (^(v3))}
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        _ = fi10
                    }
                }
            case "\n\t":
                if (!true) {
                    switch -905 {
                    case -1:
                        v3 += v3
                    default:
                        v14 := GenS3{F0: v1}
                        _ = v14
                    }
                    v1 += v3
                    v3 = (v1 * (-(v3)))
                } else {
                    switch 4294967295 {
                    case -9223372036854775808:
                        if ((127 != 100) && (1.5 >= 0.1)) {
                            v15 := GenS1{F0: true, F1: 57}
                            _ = v15
                        } else {
                            if ("hello" > "hello") {
                                v2 = -9223372036854775808
                                v16 := []int{}
                                v1 = (v3 | (v3 ^ v1))
                                _ = v16
                            } else {
                                fmt.Println(v1, v2)
                                v17 := GenS4{F0: (v2 / ((v2) | 1))}
                                _ = v17
                            }
                        }
                    case -954:
                        if (!false) {
                            v18 := GenS5{F0: -1, F1: 4294967295}
                            switch 4294967295 {
                            case 42:
                                switch 2147483647 {
                                case -711:
                                    v2 = (^(v2))
                                default:
                                    v19 := GenS6{F0: (!true), F1: 42}
                                    _ = v19
                                }
                            case -9223372036854775808:
                                if (!true) {
                                    if false {
                                        v20 := []int{}
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                        _ = v20
                                    }
                                }
                            }
                            _ = v18
                        }
                    default:
                        v21 := GenS5{F0: -128, F1: 404}
                        _ = v21
                    }
                    v1 = -722
                }
            default:
                var v22 uint16 = 255
                _ = v22
            }
        }
    case "":
        switch "hello" {
        case "foo bar":
            v3 += v1
        case "a":
            var v23 float32 = 1.5
            _ = v23
        case "\n\t":
            v24 := GenS3{F0: -2147483648}
            _ = v24
        }
    case "a":
        var v25 string = "foo bar"
        _ = v25
    default:
        v3 -= (v1 | (v3 << ((984) & 63)))
    }
    switch "\n\t" {
    case "a":
        var v26 float32 = 1e10
        _ = v26
    case "hello":
        v27 := GenS3{F0: (v1 ^ v3)}
        _ = v27
    case "foo bar":
        switch "foo bar" {
        case "hello":
            if (254 == 255) {
                var v28 float64 = -2.25
                if (!false) {
                    if ((0.1 - (0.0 * 0.0)) < (float32((float64(-2.25) + v28)) + 3.14159)) {
                        fmt.Println(v1)
                    }
                } else {
                    v29 := []int{}
                    if (!true) {
                        for fi30 := 0; fi30 < 4; fi30++ {
                            for fi31 := 0; fi31 < 12; fi31++ {
                                v32 := GenS4{F0: (v2 & 0)}
                                _ = fi31
                                _ = v32
                            }
                            _ = fi30
                        }
                        for fi33 := 0; fi33 < 9; fi33++ {
                            var v34 int64 = v2
                            _ = fi33
                            _ = v34
                        }
                    } else {
                        v3 -= v3
                        for fi35 := 0; fi35 < 18; fi35++ {
                            v1 -= (-(v3))
                            v36 := GenS7{F0: 932, F1: -1}
                            _ = fi35
                            _ = v36
                        }
                        v37 := GenS3{F0: genFn0(474, 18446744073709551615)}
                        _ = v37
                    }
                    v38 := GenS2{F0: 1, F1: -1}
                    _ = v29
                    _ = v38
                }
                _ = v28
            }
        case "a":
            v1 -= (v3 | genFn0(-32768, 2))
        default:
            var v39 int32 = (^(v3))
            _ = v39
        }
    default:
        if true {
            v1 = 42
        } else {
            var v40 int32 = (v1 >> ((247) & 63))
            _ = v40
        }
    }
    var v41 float32 = (3.14159 / (-2.25 - 0.0))
    switch "hello" {
    case "foo bar":
        v42 := []int{836, 32768}
        _ = v42
    default:
        v43 := []int{0, -1, 173, -1}
        _ = v43
    }
    if (1 == 2) {
        var v44 uint8 = 128
        _ = v44
    } else {
        switch 32767 {
        case -148:
            v1 -= -1
        case 750:
            v45 := GenS8{F0: 42, F1: (v3 & v3)}
            _ = v45
        case 32768:
            v1 = (v3 - v1)
        }
        var v46 bool = false
        _ = v46
    }
    var v47 bool = (!(v1 < (^(v3))))
    v1 = (v1 | -2147483648)
    for fi48 := 0; fi48 < 16; fi48++ {
        var v49 uint8 = 0
        _ = fi48
        _ = v49
    }
    if (!(2 >= 249)) {
        switch -708 {
        case 49:
            fmt.Println(v47, v47, v2)
        case 32767:
            v47 = true
        case -1:
            var v50 int16 = 0
            _ = v50
        }
        v2 -= 9223372036854775807
        var v51 uint16 = 589
        _ = v51
    }
    if false {
        v1 = 458
    } else {
        v52 := []int{-33, -770, 32767, 42}
        _ = v52
    }
    for fi53 := 0; fi53 < 17; fi53++ {
        var v54 uint16 = 493
        switch "hello" {
        case "":
            var v55 uint32 = 1
            _ = v55
        case "\n\t":
            v3 += -1
        case "a":
            v56 := []int{(fi53 & 127), fi53, (-(fi53)), (-(fi53))}
            _ = v56
        }
        _ = fi53
        _ = v54
    }
    v41 = float32(3.14159)
    fmt.Println(v47)
    switch -1 {
    case 1:
        v57 := []int{127, 355, 128, 128}
        _ = v57
    case 9223372036854775807:
        for fi58 := 0; fi58 < 15; fi58++ {
            var v59 uint32 = 4
            var v60 uint32 = v59
            _ = fi58
            _ = v59
            _ = v60
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v47
}
