package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int; F1 uint16; F2 uint }
func genFn0() bool {
    if (42 >= 1) {
        var v4 int16 = 1
        v4 += v4
        _ = v4
    } else {
        var v5 uint64 = 18446744073709551615
        v6 := []int{32768, 880, 9223372036854775807, 1}
        for fi7 := 0; fi7 < 6; fi7++ {
            v5 += v5
            _ = fi7
        }
        _ = v5
        _ = v6
    }
    v8 := GenS1{F0: ((1e10 - 3.14159) > 0.0)}
    v9 := GenS1{F0: true}
    _ = v8
    _ = v9
    return true
}
func genFn1() uint16 {
    if (884 > 42) {
        if true {
            if true {
                switch -2147483648 {
                case 0:
                    var v10 uint16 = 835
                    _ = v10
                case -2147483648:
                    var v11 int8 = 42
                    _ = v11
                case 42:
                    switch ("hello" + "hello") {
                    case "":
                        var v12 uint = 4294967295
                        _ = v12
                    case "hello":
                        var v13 uint16 = 65535
                        _ = v13
                    case "foo bar":
                        v14 := []int{}
                        _ = v14
                    default:
                        v15 := GenS2{F0: 352, F1: 4294967295, F2: 732}
                        _ = v15
                    }
                }
            } else {
                var v16 float32 = 3.14159
                fmt.Println(v16)
                _ = v16
            }
            switch "" {
            case "":
                var v17 uint8 = 46
                _ = v17
            case "foo bar":
                v18 := []int{-397, 42, 32768}
                _ = v18
            default:
                switch 296 {
                case -797:
                    v19 := []int{-9223372036854775808, 127, 32767}
                    _ = v19
                case 0:
                    var v20 uint64 = 0
                    _ = v20
                case 32767:
                    var v21 uint8 = 128
                    _ = v21
                default:
                    for fi22 := 0; fi22 < 5; fi22++ {
                        var v23 uint64 = 789
                        v23 = v23
                        if true {
                            if false {
                                v23 += v23
                                v24 := []int{(fi22 + fi22), fi22}
                                v23 -= 2
                                _ = v24
                            }
                        }
                        _ = fi22
                        _ = v23
                    }
                }
            }
            v25 := []int{42, -512, 0, 42}
            _ = v25
        } else {
            v26 := GenS2{F0: 2, F1: 478, F2: 958}
            for fi27 := 0; fi27 < 2; fi27++ {
                var v28 int8 = 0
                _ = fi27
                _ = v28
            }
            _ = v26
        }
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 uint16 = 2
            _ = fi29
            _ = v30
        }
    }
    var v31 float64 = 1e10
    v31 -= v31
    var v32 uint = 313
    _ = v31
    _ = v32
    return 232
}

func main() {
    v1 := []string{"foo bar", "\n\t"}
    var v2 int64 = -9223372036854775808
    v3 := []uint32{2147483647, 2}
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    v2 = (v2 % (((v2 - (v2 & (v2 - 1)))) | 1))
    v2 += 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        if (!true) {
            var v34 bool = ((42 != 42) || (!false))
            _ = v34
        } else {
            for fi35 := 0; fi35 < 14; fi35++ {
                fmt.Println(fi33, v2)
                _ = fi35
            }
        }
        if false {
            var v36 int8 = 73
            _ = v36
        }
        _ = fi33
    }
    v2 = v2
    switch 250 {
    case 32768:
        if (!false) {
            switch -523 {
            case 127:
                for fi37 := 0; fi37 < 18; fi37++ {
                    var v38 string = (("\n\t" + "foo bar") + ("a" + "a"))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v38, v2)
                    _ = fi37
                    _ = v38
                }
            case -430:
                var v39 string = "hello"
                _ = v39
            }
        } else {
            v40 := GenS2{F0: genFn1(), F1: 562, F2: 1}
            _ = v40
        }
    }
    v41 := GenS3{F0: 2147483647, F1: 2147483647}
    v42 := GenS3{F0: 755, F1: -1}
    v2 -= v2
    var v43 int32 = -725
    var v44 bool = (((float64((float32(((float64(((-2.25 / (float32(1.5) / 3.14159)) - -2.25)) - 0.1) + 1.5)) / float32(3.14159))) - (((0.1 + 1.5) - ((0.1 / 1.5) - (0.1 / 3.14159))) + 1e10)) / 1e10) < (0.1 * (1e10 + (0.0 * 0.0))))
    switch ("hello" + "\n\t") {
    case "\n\t":
        for fi45 := 0; fi45 < 14; fi45++ {
            var v46 int8 = 1
            v2 = (v2 >> ((65535) & 63))
            var v47 int8 = 100
            _ = fi45
            _ = v46
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 5; fi48++ {
        v49 := GenS1{F0: v44}
        _ = fi48
        _ = v49
    }
    v50 := GenS4{F0: 42, F1: v43}
    if (((("a" + "hello") + ("a" + "a")) > "\n\t") && false) {
        for fi51 := 0; fi51 < 20; fi51++ {
            var v52 bool = (!((v2 - (v2 - 9223372036854775807)) == 1))
            _ = fi51
            _ = v52
        }
        if false {
            if v44 {
                switch -108 {
                case -217:
                    var v53 int64 = v2
                    _ = v53
                case 390:
                    var v54 string = ""
                    _ = v54
                }
                for fi55 := 0; fi55 < 20; fi55++ {
                    var v56 string = "a"
                    _ = fi55
                    _ = v56
                }
                var v57 string = "a"
                _ = v57
            } else {
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                var v58 int32 = v43
                _ = v58
            }
            var v59 uint = 1
            v60 := GenS5{F0: 18446744073709551615, F1: "hello"}
            _ = v59
            _ = v60
        }
        v61 := GenS6{F0: 42, F1: 1, F2: 0}
        _ = v61
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v50
}
