package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint; F2 uint8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int16; F1 bool; F2 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }
func genFn0() uint32 {
    v6 := []int{}
    if true {
        v7 := GenS1{F0: 32767, F1: 2147483647, F2: 993}
        for fi8 := 0; fi8 < 16; fi8++ {
            if true {
                for fi9 := 0; fi9 < 13; fi9++ {
                    switch ("foo bar" + ("foo bar" + "\n\t")) {
                    case "foo bar":
                        _ = len(v6)
                    case "":
                        var v10 int = fi9
                        _ = v10
                    }
                    var v11 int = -1
                    v11 += 898
                    _ = fi9
                    _ = v11
                }
            }
            _ = fi8
        }
        _ = v7
    } else {
        var v12 int8 = 127
        for fi13 := 0; fi13 < 18; fi13++ {
            switch 0 {
            case 0:
                v12 -= 42
            }
            _ = fi13
        }
        var v14 int16 = -884
        _ = v12
        _ = v14
    }
    _ = v6
    return 42
}
func genFn1(p0 int8) uint8 {
    var v15 uint = 255
    var v16 uint = (v15 % (((^(v15))) | 1))
    var v17 int32 = 42
    fmt.Println(p0, v17, v17)
    var v18 uint32 = 149
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 1
    v3 := make(map[uint32]string)
    var v4 uint8 = 126
    v5 := make(map[string]bool)
    if (true || true) {
        switch ("" + "") {
        case "hello":
            fmt.Println(v4, v2)
        case "a":
            if false {
                var v19 int16 = -549
                switch 42 {
                case 689:
                    var v20 bool = (!true)
                    _ = v20
                }
                var v21 uint8 = genFn1(100)
                _ = v19
                _ = v21
            } else {
                var v22 bool = (-2.25 == 3.14159)
                _ = v22
            }
        }
    }
    var v23 bool = true
    var v24 uint16 = (v2 | (v2 ^ 0))
    switch -44 {
    case 32768:
        for fi25 := 0; fi25 < 11; fi25++ {
            switch fi25 {
            case 778:
                v2 = (v2 ^ v2)
            default:
                var v26 int = fi25
                _ = v26
            }
            v24 = (v2 - (^(v24)))
            _ = fi25
        }
    case 128:
        v27 := []int{0, 128, -9223372036854775808, 9223372036854775807}
        _ = v27
    }
    switch 4294967295 {
    case 2147483647:
        v28 := GenS1{F0: 32767, F1: 844, F2: 557}
        _ = v28
    case -9223372036854775808:
        v29 := GenS1{F0: 127, F1: 472, F2: 28}
        _ = v29
    default:
        fmt.Println(v4, v4)
    }
    v30 := GenS2{F0: (v23 && true), F1: 18446744073709551615, F2: v4}
    var v31 uint16 = (v2 % ((2) | 1))
    if true {
        for fi32 := 0; fi32 < 10; fi32++ {
            v5["\n\t"] = v23
            _ = fi32
        }
        fmt.Println(v2, v2)
    }
    var v33 int8 = 42
    for fi34 := 0; fi34 < 13; fi34++ {
        v35 := GenS3{F0: v33, F1: "foo bar"}
        for fi36 := 0; fi36 < 15; fi36++ {
            v37 := GenS3{F0: 100, F1: ""}
            _ = fi36
            _ = v37
        }
        _ = fi34
        _ = v35
    }
    switch "a" {
    case "a":
        if (18446744073709551615 == 4294967295) {
            v38 := []int{-1, int(v33)}
            if len(v38) > 0 {
                v38[0] = 4294967295
            }
            for fi39 := 0; fi39 < 9; fi39++ {
                switch fi39 {
                case 4294967295:
                    v31 = (^(v2))
                }
                _ = fi39
            }
            _ = v38
        } else {
            v40 := []int{21, 4294967295, 127, 0}
            if len(v40) > 0 {
                v40[0] = 0
            }
            v41 := GenS4{F0: -32768, F1: 9223372036854775807}
            _ = v40
            _ = v41
        }
    case "":
        if false {
            if ("\n\t" > ("foo bar" + "")) {
                for fi42 := 0; fi42 < 1; fi42++ {
                    if v23 {
                        v2 = v24
                        if (!(0 <= 1)) {
                            v24 = (v24 << ((65535) & 63))
                            v43 := GenS2{F0: false, F1: 9, F2: (v4 | 2)}
                            _ = v43
                        } else {
                            if true {
                                v44 := []int{fi42, fi42, (fi42 - (^(fi42)))}
                                var v45 int32 = 42
                                for fi46 := 0; fi46 < 2; fi46++ {
                                    if len(v44) > 0 {
                                        v44[0] = (fi46 | (fi42 * fi42))
                                    }
                                    var v47 uint16 = v24
                                    _ = fi46
                                    _ = v47
                                }
                                _ = v44
                                _ = v45
                            }
                            switch "a" {
                            case "foo bar":
                                if (v24 >= v2) {
                                    delete(v1, 9223372036854775807)
                                    v48 := GenS5{F0: 0, F1: v23, F2: 1}
                                    _ = v48
                                }
                            }
                            v4 -= v4
                        }
                        fmt.Println(v4, v31, v2)
                    }
                    if (1.5 == (float64(0.0) / -2.25)) {
                        fmt.Println(v23, v33)
                    } else {
                        for fi49 := 0; fi49 < 8; fi49++ {
                            if (!true) {
                                v31 = v24
                                v23 = v23
                                v2 = v31
                                v24 += v24
                                v4 -= (v4 >> ((0) & 63))
                            } else {
                                v50 := GenS6{F0: genFn1((-(v33)))}
                                v2 = v31
                                _ = v50
                            }
                            fmt.Println(v33, v23, v33)
                            var v51 uint = 167
                            if (!(!v23)) {
                                fmt.Println(fi42)
                            }
                            _ = fi49
                            _ = v51
                        }
                    }
                    _ = fi42
                }
                fmt.Println(v24)
            } else {
                switch ("foo bar" + (("\n\t" + "\n\t") + "")) {
                case "\n\t":
                    v52 := GenS1{F0: 962, F1: 379, F2: 1}
                    _ = v52
                case "hello":
                    switch "foo bar" {
                    case "":
                        switch (("" + "a") + "hello") {
                        case "":
                            v24 = v31
                        case "\n\t":
                            var v53 string = "hello"
                            _ = v53
                        case "a":
                            fmt.Println(v4, v33)
                        default:
                            for fi54 := 0; fi54 < 14; fi54++ {
                                switch (fi54 | (fi54 / ((fi54) | 1))) {
                                case 586:
                                    var v55 int8 = (v33 >> ((4294967295) & 63))
                                    _ = v55
                                }
                                var v56 int = fi54
                                fmt.Println(v24, v31)
                                _ = fi54
                                _ = v56
                            }
                        }
                    case "a":
                        v4 = 146
                    default:
                        var v57 bool = v23
                        _ = v57
                    }
                }
                switch 127 {
                case -918:
                    switch 128 {
                    case 32767:
                        var v58 uint = 2
                        _ = v58
                    case -665:
                        v59 := GenS7{F0: 42}
                        _ = v59
                    default:
                        for fi60 := 0; fi60 < 11; fi60++ {
                            v61 := []int{19}
                            _ = fi60
                            _ = v61
                        }
                    }
                case 2147483647:
                    for fi62 := 0; fi62 < 11; fi62++ {
                        v63 := []int{(fi62 & (fi62 << ((851) & 63))), fi62}
                        var v64 int16 = 1000
                        for fi65 := 0; fi65 < 10; fi65++ {
                            var v66 int = fi62
                            _ = fi65
                            _ = v66
                        }
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                case -9223372036854775808:
                    switch -1 {
                    case 9223372036854775807:
                        var v67 float64 = (1.5 / 1.5)
                        _ = v67
                    default:
                        v68 := GenS1{F0: 4294967295, F1: 269, F2: 2147483648}
                        _ = v68
                    }
                default:
                    var v69 int = -2147483648
                    _ = v69
                }
                for fi70 := 0; fi70 < 19; fi70++ {
                    v33 = 1
                    _ = fi70
                }
            }
            for fi71 := 0; fi71 < 15; fi71++ {
                v72 := GenS7{F0: 32767}
                v24 = v31
                _ = fi71
                _ = v72
            }
        }
    case "\n\t":
        var v73 uint8 = (^(v4))
        _ = v73
    default:
        var v74 uint32 = 994
        _ = v74
    }
    var v75 int = -9223372036854775808
    v2 -= 42
    var v76 float64 = 3.14159
    for fi77 := 0; fi77 < 10; fi77++ {
        for fi78 := 0; fi78 < 7; fi78++ {
            v31 = (v24 << ((99) & 63))
            v4 -= (v4 - (^(v4)))
            _ = fi78
        }
        var v79 string = ("\n\t" + "a")
        _ = fi77
        _ = v79
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
    _ = v5
    _ = v75
    _ = v76
}
