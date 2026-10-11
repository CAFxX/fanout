package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int16; F1 string; F2 bool }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int8; F1 int64; F2 bool }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 42
    v3 := make(map[uint32]string)
    var v4 int = 2147483647
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int64 = 1
        _ = fi5
        _ = v6
    }
    v4 = -1
    v4 = (v4 ^ -1)
    var v7 string = "hello"
    if ((!(false || ((!true) || true))) && true) {
        var v8 int64 = 1
        _ = v8
    } else {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 int8 = 127
            var v11 int64 = 437
            _ = fi9
            _ = v10
            _ = v11
        }
        v7 = (((("\n\t" + v7) + "hello") + v7) + "hello")
        switch "" {
        case "hello":
            var v12 uint16 = 42
            _ = v12
        case "foo bar":
            v7 = v7
        default:
            switch "" {
            case "foo bar":
                v13 := []int{(^(v4)), v4, 0}
                _ = v13
            case "\n\t":
                for fi14 := 0; fi14 < 20; fi14++ {
                    fmt.Println(v2, v4, v7)
                    switch ("a" + (((v7 + ("a" + ("foo bar" + "\n\t"))) + v7) + "\n\t")) {
                    case "a":
                        v4 = (v4 ^ fi14)
                    case "\n\t":
                        v2 = v2
                    default:
                        var v15 int64 = 732
                        _ = v15
                    }
                    _ = fi14
                }
            case "":
                v16 := GenS1{F0: (v4 - v4), F1: 2147483647}
                _ = v16
            }
        }
    }
    v4 -= v4
    switch v7 {
    case "a":
        v2 = 0
    }
    v17 := GenS2{F0: 677, F1: 0}
    for fi18 := 0; fi18 < 7; fi18++ {
        if false {
            var v19 string = v7
            v7 = v19
            v20 := GenS3{F0: -32768, F1: "foo bar", F2: false}
            _ = v19
            _ = v20
        }
        switch v7 {
        case "\n\t":
            if (90 != 1) {
                var v21 uint64 = 139
                for fi22 := 0; fi22 < 16; fi22++ {
                    _ = len(v1)
                    v21 -= 100
                    v23 := v3[0]
                    _ = fi22
                    _ = v23
                }
                if true {
                    if ((!true) && false) {
                        v24 := v1[42]
                        _ = v24
                    }
                    switch v4 {
                    case 0:
                        fmt.Println(v7)
                    case 2147483647:
                        v7 = ("\n\t" + (("" + "hello") + v7))
                    case 127:
                        switch (fi18 ^ int(v2)) {
                        case 1:
                            switch 127 {
                            case 128:
                                if (!true) {
                                    var v25 int32 = v2
                                    _ = v25
                                }
                            }
                        case -131:
                            v26 := GenS4{F0: 1, F1: 4294967295}
                            _ = v26
                        default:
                            if (false || true) {
                                var v27 int32 = v2
                                v28 := []int{fi18, (-(fi18)), 32768, 2147483647}
                                if (true && ((false && false) || (42 >= 42))) {
                                    var v29 uint16 = 938
                                    for fi30 := 0; fi30 < 8; fi30++ {
                                        var v31 uint16 = v29
                                        _ = fi30
                                        _ = v31
                                    }
                                    _ = v29
                                }
                                _ = v27
                                _ = v28
                            } else {
                                v32 := v3[4294967295]
                                v21 = (v21 - (v21 ^ v21))
                                _ = v32
                            }
                        }
                    default:
                        v2 -= -2147483648
                    }
                    var v33 int8 = int8(v4)
                    _ = v33
                } else {
                    switch v7 {
                    case "\n\t":
                        v34 := GenS5{F0: 0, F1: 9223372036854775807, F2: false}
                        _ = v34
                    case "a":
                        var v35 int16 = 42
                        _ = v35
                    case "hello":
                        delete(v3, 2147483648)
                    default:
                        var v36 uint8 = 127
                        _ = v36
                    }
                    if false {
                        v37 := GenS6{F0: ("\n\t" + ((("a" + ("" + v7)) + (("foo bar" + "") + v7)) + (v7 + (v7 + "hello")))), F1: 255}
                        _ = v37
                    } else {
                        switch (v7 + "a") {
                        case "\n\t":
                            switch "\n\t" {
                            case "a":
                                switch (v7 + (v7 + ((v7 + (("a" + "hello") + ("foo bar" + "a"))) + ""))) {
                                case "hello":
                                    for fi38 := 0; fi38 < 6; fi38++ {
                                        switch (^(v4)) {
                                        case 32767:
                                            v2 = (-(v2))
                                        case -2147483648:
                                            v7 = ("" + (v7 + ("\n\t" + v7)))
                                        }
                                        switch (v4 << ((0) & 63)) {
                                        case 4294967295:
                                            if true {
                                                if (!((!((!(4294967295 != 1)) || (42 > 65535))) || (1e10 == -2.25))) {
                                                    switch "a" {
                                                    case "foo bar":
                                                        v2 = v2
                                                    case "a":
                                                        var v39 uint64 = v21
                                                        _ = v39
                                                    case "hello":
                                                        var v40 int = (v4 << ((849) & 63))
                                                        _ = v40
                                                    default:
                                                        fmt.Println(fi18, fi18)
                                                    }
                                                }
                                            } else {
                                                v21 -= v21
                                                var v41 uint = 1
                                                v4 = (-(fi18))
                                                _ = v41
                                            }
                                        case 42:
                                            v42 := GenS3{F0: 1000, F1: "", F2: true}
                                            _ = v42
                                        case 9223372036854775807:
                                            var v43 int32 = (v2 / (((-(v2))) | 1))
                                            _ = v43
                                        default:
                                            v44 := []int{}
                                            _ = v44
                                        }
                                        switch ("" + "a") {
                                        case "foo bar":
                                            fmt.Println(v2, v7, fi18)
                                        case "a":
                                            var v45 uint16 = 2
                                            _ = v45
                                        }
                                        _ = fi38
                                    }
                                case "a":
                                    var v46 uint = 361
                                    _ = v46
                                case "foo bar":
                                    v7 = ""
                                }
                            case "hello":
                                var v47 int8 = 127
                                _ = v47
                            case "":
                                var v48 int64 = -44
                                _ = v48
                            default:
                                if (!true) {
                                    if ("\n\t" <= "foo bar") {
                                        if (42 >= 1) {
                                            switch (v7 + (v7 + (("foo bar" + ("\n\t" + v7)) + v7))) {
                                            case "hello":
                                                v21 -= v21
                                            case "foo bar":
                                                for fi49 := 0; fi49 < 5; fi49++ {
                                                    switch fi18 {
                                                    case 0:
                                                        v4 -= (fi49 ^ (v4 >> ((829) & 63)))
                                                    default:
                                                        delete(v3, uint32(v21))
                                                        v21 += 449
                                                        if false {
                                                            var v50 int32 = (v2 & (v2 * (v2 >> ((18446744073709551615) & 63))))
                                                            v51 := GenS2{F0: 1000, F1: 99}
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    }
                                                    if ((!true) || (88 != 218)) {
                                                        switch (("foo bar" + "\n\t") + v7) {
                                                        case "foo bar":
                                                            var v52 uint = 1
                                                            _ = v52
                                                        case "hello":
                                                            v7 = v7
                                                        case "a":
                                                            v53 := GenS3{F0: 32767, F1: (v7 + "a"), F2: false}
                                                            _ = v53
                                                        default:
                                                            var v54 string = (("foo bar" + v7) + v7)
                                                            _ = v54
                                                        }
                                                        v7 = ("hello" + "hello")
                                                        if (float32(1.5) <= 0.0) {
                                                            v21 -= 0
                                                            switch v4 {
                                                            case 2147483647:
                                                                var v55 bool = false
                                                                _ = v55
                                                            default:
                                                                v56 := GenS7{F0: 4294967295}
                                                                _ = v56
                                                            }
                                                            var v57 bool = false
                                                            var v58 uint = 626
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                        var v59 int8 = 27
                                                        _ = v59
                                                    } else {
                                                        v4 += (fi18 | fi49)
                                                    }
                                                    if true {
                                                        v7 = "\n\t"
                                                    }
                                                    _ = fi49
                                                }
                                            case "":
                                                v60 := GenS3{F0: 1000, F1: v7, F2: true}
                                                _ = v60
                                            default:
                                                v61 := GenS2{F0: 947, F1: 128}
                                                _ = v61
                                            }
                                        }
                                    }
                                }
                            }
                        case "a":
                            v4 = (v4 ^ fi18)
                        }
                        for fi62 := 0; fi62 < 10; fi62++ {
                            switch (fi62 % ((v4) | 1)) {
                            case 42:
                                var v63 float64 = 0.0
                                _ = v63
                            }
                            v64 := []int{}
                            var v65 int16 = 901
                            if len(v64) > 0 {
                                v64[0] = fi18
                            }
                            _ = fi62
                            _ = v64
                            _ = v65
                        }
                        switch -295 {
                        case 4294967295:
                            v66 := GenS4{F0: (v21 % (((v21 + (v21 + v21))) | 1)), F1: 4294967295}
                            _ = v66
                        default:
                            var v67 string = (v7 + ("a" + v7))
                            _ = v67
                        }
                    }
                }
                _ = v21
            }
        default:
            fmt.Println(v2, v7)
        }
        _ = fi18
    }
    v68 := []int{(v4 >> ((290) & 63))}
    fmt.Println(v2)
    for fi69 := 0; fi69 < 18; fi69++ {
        v70 := GenS2{F0: 32767, F1: 2}
        _ = fi69
        _ = v70
    }
    for fi71 := 0; fi71 < 2; fi71++ {
        if len(v68) >= 2 {
            v68 = v68[1:len(v68)]
        }
        _ = fi71
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v68:", len(v68), v68)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v68
    _ = v7
}
