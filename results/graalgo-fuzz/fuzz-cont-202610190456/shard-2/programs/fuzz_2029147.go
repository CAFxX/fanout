package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint; F1 bool; F2 uint32 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 int64; F1 uint64; F2 int64 }
type GenS8 struct { F0 int16; F1 uint64; F2 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    v3 := []string{"hello", "", "foo bar"}
    var v4 int64 = 9223372036854775807
    var v5 int32 = -383
    var v6 int64 = v4
    if false {
        v6 = (v4 % ((1) | 1))
        if (-32768 < 152) {
            var v7 float64 = -2.25
            if true {
                v4 = (^(v4))
                v6 += v6
                delete(v1, 108)
            }
            var v8 uint = 290
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: 4294967295}
            _ = v9
        }
        var v10 uint32 = 2147483647
        _ = v10
    } else {
        if false {
            var v11 uint16 = 42
            _ = v11
        } else {
            var v12 float32 = 3.14159
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 string = ""
                var v15 int8 = 127
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 int32 = 1
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: 0}
        _ = v18
    }
    if false {
        v19 := GenS1{F0: 2147483647}
        _ = v19
    } else {
        delete(v1, 1)
        v2[32767] = "a"
    }
    v20 := GenS1{F0: 2}
    var v21 float64 = -2.25
    if true {
        v3 = append(v3, ("foo bar" + "hello"))
        switch -497 {
        case 127:
            v22 := GenS1{F0: 2147483647}
            _ = v22
        case 9223372036854775807:
            v3 = append(v3, "a")
        case 32768:
            for fi23 := 0; fi23 < 7; fi23++ {
                for fi24 := 0; fi24 < 18; fi24++ {
                    var v25 uint64 = 1
                    _ = fi24
                    _ = v25
                }
                _ = len(v3)
                _ = fi23
            }
        default:
            switch ("a" + "") {
            case "foo bar":
                var v26 int = -400
                _ = v26
            }
        }
        v5 -= -1
    } else {
        v27 := GenS1{F0: 42}
        if (0.0 >= float32(v21)) {
            v4 = 81
            var v28 uint32 = 2147483647
            _ = v28
        }
        v4 = v6
        _ = v27
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        switch ("" + "hello") {
        case "a":
            switch fi29 {
            case -522:
                v30 := GenS1{F0: 5}
                _ = v30
            default:
                if ((true || (false || ((4294967295 <= 114) && false))) && ((false || true) && true)) {
                    var v31 int = fi29
                    var v32 uint8 = 127
                    _ = v31
                    _ = v32
                }
            }
        default:
            v33 := GenS2{F0: 0, F1: 62}
            _ = v33
        }
        if ((false || false) && (((false || (true || true)) || false) || (2 <= 931))) {
            for fi34 := 0; fi34 < 2; fi34++ {
                switch "hello" {
                case "a":
                    v5 += (v5 ^ -1)
                    v5 -= 1
                    v6 = v4
                case "\n\t":
                    v35 := GenS3{F0: 255, F1: (float64(0.0) < ((v21 - (((v21 + (v21 * 0.1)) * v21) - -2.25)) + v21)), F2: 2}
                    _ = v35
                default:
                    fmt.Println(v5)
                }
                _ = fi34
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            delete(v1, 0)
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            if false {
                var v37 uint8 = 1
                delete(v2, 0)
                switch "" {
                case "hello":
                    if ((!((v6 % ((-938) | 1)) >= -409)) && true) {
                        v38 := GenS2{F0: 42, F1: (v37 - 255)}
                        var v39 int = fi36
                        _ = v38
                        _ = v39
                    }
                }
                _ = v37
            } else {
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
            }
            _ = fi36
        }
        _ = fi29
    }
    switch 1 {
    case -514:
        if (true && false) {
            for fi40 := 0; fi40 < 2; fi40++ {
                switch "\n\t" {
                case "foo bar":
                    for fi41 := 0; fi41 < 4; fi41++ {
                        var v42 float64 = (v21 / v21)
                        if true {
                            fmt.Println(v42, v6, v21)
                            var v43 bool = false
                            _ = v43
                        } else {
                            var v44 bool = false
                            v45 := GenS4{F0: 1, F1: 430}
                            var v46 int32 = (v5 >> ((202) & 63))
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        var v47 uint32 = 519
                        if false {
                            switch 2147483647 {
                            case 2147483647:
                                v48 := GenS5{F0: 255}
                                _ = v48
                            case 9223372036854775807:
                                fmt.Println(v42, v42)
                            }
                        } else {
                            var v49 uint32 = v47
                            v6 -= (v4 | (v4 / ((v4) | 1)))
                            _ = v49
                        }
                        _ = fi41
                        _ = v42
                        _ = v47
                    }
                case "hello":
                    var v50 float64 = v21
                    _ = v50
                default:
                    var v51 string = ("a" + "\n\t")
                    _ = v51
                }
                v52 := GenS4{F0: 42, F1: fi40}
                var v53 int = (fi40 >> ((4294967295) & 63))
                _ = fi40
                _ = v52
                _ = v53
            }
        }
    case 1:
        fmt.Println(v4, v6)
    case 127:
        if ((2 <= 2) || false) {
            _ = len(v3)
            _ = len(v2)
            for fi54 := 0; fi54 < 11; fi54++ {
                v55 := GenS5{F0: 65535}
                var v56 float32 = float32(v21)
                var v57 int32 = v5
                _ = fi54
                _ = v55
                _ = v56
                _ = v57
            }
        }
    }
    if false {
        if ((-2.25 > (v21 / v21)) && ((3.14159 / -2.25) < 3.14159)) {
            var v58 uint16 = 57
            v58 -= v58
            _ = v58
        } else {
            v59 := GenS5{F0: 42}
            _ = v59
        }
    }
    if (18446744073709551615 < 65535) {
        delete(v2, -434)
    }
    v60 := GenS6{F0: true, F1: 32767}
    _ = len(v3)
    fmt.Println(v6, v21, v4)
    v4 = v6
    delete(v1, 2147483647)
    if true {
        if true {
            v5 = v5
        } else {
            var v61 float32 = ((((3.14159 * (1.5 / 1.0)) + 3.14159) / ((((0.0 * -2.25) - ((((((1e10 * 1e10) - (-2.25 + (1e10 * 3.14159))) - (0.0 + (-2.25 / (-2.25 - (-2.25 / 1e10))))) - (3.14159 - 3.14159)) - (3.14159 * float32(3.14159))) / (0.1 / float32(v21)))) * 0.1) / 1.0)) * (0.1 / 1e10))
            var v62 uint16 = 255
            v63 := GenS5{F0: (v62 >> ((4294967295) & 63))}
            _ = v61
            _ = v62
            _ = v63
        }
        switch "" {
        case "a":
            v5 = 1
        case "hello":
            for fi64 := 0; fi64 < 17; fi64++ {
                var v65 int8 = 100
                _ = fi64
                _ = v65
            }
        case "":
            var v66 uint = 18446744073709551615
            _ = v66
        default:
            if (2 != 773) {
                v67 := GenS7{F0: -485, F1: 581, F2: 0}
                v5 += v5
                _ = v67
            }
        }
        var v68 int = 127
        _ = v68
    } else {
        v69 := GenS8{F0: 0, F1: 341, F2: 537}
        _ = v69
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
