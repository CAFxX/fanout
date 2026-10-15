package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint; F1 string; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint; F2 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 string }
func genFn0() int64 {
    var v6 int16 = 1
    if (127 == -128) {
        v7 := []int{32768, 0, -1, 32767}
        _ = v7
    }
    var v8 int16 = v6
    v6 += v6
    _ = v6
    _ = v8
    return 1
}

func main() {
    var v1 float64 = (0.0 + float64(1e10))
    var v2 float32 = (((1e10 + 0.0) - 3.14159) * (((0.0 * 3.14159) * -2.25) + 0.1))
    var v3 uint8 = 0
    v4 := make(map[uint32]bool)
    var v5 int8 = 127
    switch -1 {
    case -1:
        switch 127 {
        case 967:
            var v9 float32 = ((((1.5 - v2) * (1.5 + v2)) - (v2 + ((0.1 / ((-2.25 / 1e10) + (v2 / (v2 / v2)))) * (v2 - float32(v1))))) * v2)
            _ = v9
        case 60:
            v10 := v4[0]
            _ = v10
        case 42:
            for fi11 := 0; fi11 < 17; fi11++ {
                v2 += v2
                _ = fi11
            }
        default:
            v12 := GenS1{F0: v3, F1: "hello"}
            _ = v12
        }
    case 42:
        if true {
            v13 := GenS1{F0: 0, F1: "\n\t"}
            _ = v13
        }
    case -276:
        v14 := GenS1{F0: 49, F1: "\n\t"}
        _ = v14
    }
    switch ("foo bar" + "") {
    case "\n\t":
        v3 -= (v3 % (((v3 + v3)) | 1))
    case "hello":
        v1 = 1.5
    case "a":
        delete(v4, 0)
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 uint8 = v3
        var v17 uint = 255
        v17 += (v17 | (v17 | v17))
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        var v18 int64 = 0
        _ = v18
    } else {
        v5 += (v5 & v5)
    }
    var v19 bool = ((187 <= 432) && false)
    var v20 uint32 = 71
    v19 = true
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 uint8 = v3
        if (!(!v19)) {
            switch int(v22) {
            case 32767:
                for fi23 := 0; fi23 < 11; fi23++ {
                    delete(v4, v20)
                    var v24 uint8 = (^(v22))
                    _ = fi23
                    _ = v24
                }
            case -999:
                var v25 uint8 = (v3 ^ v3)
                _ = v25
            case 4294967295:
                switch (fi21 & 42) {
                case -2147483648:
                    v26 := GenS2{F0: 42, F1: v20}
                    _ = v26
                case 227:
                    var v27 int = (^(fi21))
                    _ = v27
                }
            default:
                if v19 {
                    if (v19 && v19) {
                        switch fi21 {
                        case 128:
                            v20 = 1
                        case 47:
                            fmt.Println(v19, v5, v22)
                        default:
                            switch ((("" + "hello") + (((("foo bar" + "hello") + "a") + "\n\t") + "a")) + ("hello" + "foo bar")) {
                            case "":
                                var v28 uint8 = v22
                                _ = v28
                            }
                        }
                    }
                    var v29 int16 = 42
                    v30 := GenS1{F0: v3, F1: ""}
                    _ = v29
                    _ = v30
                } else {
                    switch (fi21 + (fi21 / ((-1) | 1))) {
                    case 9223372036854775807:
                        for fi31 := 0; fi31 < 10; fi31++ {
                            switch ("hello" + "hello") {
                            case "a":
                                v22 = (v22 & (v22 - 115))
                            default:
                                var v32 uint8 = v3
                                _ = v32
                            }
                            _ = fi31
                        }
                    case 1:
                        if (1 < 42) {
                            var v33 uint64 = 1
                            _ = v33
                        }
                    }
                    var v34 uint32 = (v20 % ((42) | 1))
                    if (2147483647 >= 0) {
                        v35 := []int{84, fi21, fi21, fi21}
                        if (!false) {
                            v22 += (v3 * (v3 ^ (^(v22))))
                            for fi36 := 0; fi36 < 16; fi36++ {
                                var v37 int32 = -993
                                v22 += (v3 ^ (v22 + (v22 >> ((0) & 63))))
                                if true {
                                    v38 := GenS3{F0: uint(v20), F1: (("\n\t" + ("a" + "a")) + "a"), F2: 1}
                                    v39 := GenS1{F0: (^(v22)), F1: ("foo bar" + ("foo bar" + ("\n\t" + ("hello" + ""))))}
                                    var v40 int32 = int32(v34)
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                }
                                if v19 {
                                    if true {
                                        fmt.Println(v37, v2, fi36)
                                        v41 := GenS1{F0: (v22 / (((v3 >> ((0) & 63))) | 1)), F1: "a"}
                                        v37 = -453
                                        _ = v41
                                    } else {
                                        var v42 string = (("hello" + ("hello" + "")) + (("" + ("\n\t" + "")) + ("hello" + ("" + "hello"))))
                                        if len(v35) > 0 {
                                            v35[0] = (fi21 ^ fi21)
                                        }
                                        fmt.Println(v2, v42)
                                        _ = v42
                                    }
                                    var v43 uint8 = v3
                                    switch (fi36 >> ((65535) & 63)) {
                                    case -667:
                                        v1 -= ((0.1 + -2.25) - v1)
                                    default:
                                        v44 := GenS2{F0: 65535, F1: (v20 + (v20 & v20))}
                                        _ = v44
                                    }
                                    v45 := GenS4{F0: (^(v34)), F1: 16, F2: ("foo bar" + "hello")}
                                    var v46 int8 = 1
                                    if ((!v19) && v19) {
                                        var v47 bool = (-891 > fi21)
                                        _ = v47
                                    }
                                    _ = v43
                                    _ = v45
                                    _ = v46
                                }
                                _ = fi36
                                _ = v37
                            }
                            v5 -= (v5 >> ((255) & 63))
                        }
                        _ = v35
                    }
                    _ = v34
                }
            }
            fmt.Println(v3)
            if (0 == 1) {
                for fi48 := 0; fi48 < 16; fi48++ {
                    v1 = ((float64(1.5) / (0.1 + v1)) + v1)
                    _ = fi48
                }
                v49 := []int{(fi21 - fi21)}
                for fi50 := 0; fi50 < 18; fi50++ {
                    if (1 > 0) {
                        _ = len(v49)
                        if len(v49) >= 2 {
                            v49 = v49[1:len(v49)]
                        }
                    }
                    var v51 int = 42
                    var v52 int = fi50
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                _ = v49
            }
        } else {
            switch "" {
            case "a":
                if (v20 != (v20 / ((1) | 1))) {
                    if (!(fi21 > 0)) {
                        v1 = float64((0.0 / v2))
                        v5 += -128
                    }
                    v20 = 1
                } else {
                    switch ("" + ("hello" + "\n\t")) {
                    case "hello":
                        var v53 uint8 = uint8(v20)
                        _ = v53
                    case "a":
                        for fi54 := 0; fi54 < 19; fi54++ {
                            switch (fi21 / (((fi54 / (((fi54 % ((int(v5)) | 1))) | 1))) | 1)) {
                            case 1:
                                v3 += (v22 % ((v22) | 1))
                            case 128:
                                var v55 uint16 = 903
                                _ = v55
                            default:
                                v1 += ((1.5 - v1) - 3.14159)
                            }
                            v20 = (v20 + 42)
                            v2 += v2
                            _ = fi54
                        }
                    default:
                        for fi56 := 0; fi56 < 17; fi56++ {
                            v4[(v20 % (((v20 * v20)) | 1))] = ((((1.5 / v2) / float32(((float64(3.14159) - v1) / (v1 / (v1 + (0.1 * 1.5)))))) != 0.1) && false)
                            fmt.Println(v19, v19, v19)
                            fmt.Println(v22, v19)
                            _ = fi56
                        }
                    }
                }
            case "":
                switch -9223372036854775808 {
                case 387:
                    fmt.Println(v19, v22, v3)
                case 42:
                    switch fi21 {
                    case -9223372036854775808:
                        if v19 {
                            var v57 uint16 = 65535
                            v4[(^(v20))] = false
                            delete(v4, (v20 << ((uint(v22)) & 63)))
                            _ = v57
                        }
                    case -146:
                        var v58 int16 = 180
                        _ = v58
                    default:
                        var v59 uint8 = (^(v3))
                        _ = v59
                    }
                case -1:
                    v60 := GenS2{F0: 822, F1: (v20 + (^(v20)))}
                    _ = v60
                }
            default:
                v61 := GenS5{F0: 0}
                _ = v61
            }
            switch fi21 {
            case 9223372036854775807:
                v5 = 1
            default:
                for fi62 := 0; fi62 < 11; fi62++ {
                    switch (fi21 & (fi62 - fi62)) {
                    case -2147483648:
                        var v63 float64 = 1e10
                        var v64 uint16 = 1
                        v65 := GenS6{F0: (v5 - (v5 * (v5 & v5))), F1: ("hello" + "foo bar")}
                        _ = v63
                        _ = v64
                        _ = v65
                    }
                    if v19 {
                        var v66 float32 = v2
                        var v67 uint = 167
                        _ = v66
                        _ = v67
                    } else {
                        var v68 uint64 = 42
                        v22 -= (v22 >> ((444) & 63))
                        _ = v68
                    }
                    _ = fi62
                }
            }
        }
        _ = fi21
        _ = v22
    }
    fmt.Println(v1, v5)
    switch -1 {
    case 32768:
        var v69 float32 = float32(v1)
        _ = v69
    default:
        v2 += v2
    }
    delete(v4, (v20 % (((v20 | v20)) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
