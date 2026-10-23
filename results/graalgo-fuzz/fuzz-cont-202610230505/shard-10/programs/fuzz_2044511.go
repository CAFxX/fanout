package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint64; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint16; F1 uint16 }
type GenS11 struct { F0 bool; F1 int64; F2 uint16 }
type GenS12 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 0
    var v3 int32 = 0
    var v4 string = (("a" + "") + "hello")
    switch ("a" + "\n\t") {
    case "hello":
        _ = len(v1)
    }
    if ((42 == 4294967295) || true) {
        var v5 int = -91
        fmt.Println(v2, v3)
        _ = v5
    } else {
        var v6 bool = false
        var v7 uint32 = 612
        for fi8 := 0; fi8 < 1; fi8++ {
            _ = len(v1)
            var v9 string = (v4 + v4)
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    var v10 uint16 = 32767
    var v11 string = (v4 + "hello")
    for fi12 := 0; fi12 < 3; fi12++ {
        fmt.Println(v3, v3)
        var v13 float64 = (3.14159 - 3.14159)
        switch fi12 {
        case 4294967295:
            v14 := GenS1{F0: 42}
            _ = v14
        case 32767:
            v15 := GenS2{F0: v11}
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v16 := GenS1{F0: 255}
    v2 -= (v2 << ((504) & 63))
    var v17 uint32 = 1
    v18 := v1[2147483648]
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 float64 = (float64(1.5) * ((1e10 / float64(((((1.5 - float32((0.1 / (3.14159 * 0.1)))) * -2.25) * -2.25) + (3.14159 * -2.25)))) + (-2.25 - 1e10)))
            _ = fi20
            _ = v21
        }
        v22 := GenS3{F0: 75, F1: 0, F2: v4}
        _ = fi19
        _ = v22
    }
    fmt.Println(v4, v4, v3)
    v23 := []int{}
    v24 := GenS2{F0: v11}
    var v25 int16 = -831
    switch ("\n\t" + "hello") {
    case "foo bar":
        if (!false) {
            switch ("hello" + "") {
            case "hello":
                _ = len(v23)
            default:
                var v26 uint = 0
                _ = v26
            }
            _ = len(v23)
            if false {
                switch v2 {
                case 2147483647:
                    for fi27 := 0; fi27 < 11; fi27++ {
                        if false {
                            var v28 uint = 4294967295
                            v29 := GenS4{F0: "foo bar"}
                            for fi30 := 0; fi30 < 11; fi30++ {
                                v3 -= v3
                                _ = fi30
                            }
                            _ = v28
                            _ = v29
                        } else {
                            if len(v23) > 0 {
                                v23[0] = -330
                            }
                            switch -1 {
                            case -9223372036854775808:
                                v17 += v17
                            case 9223372036854775807:
                                switch v2 {
                                case -2147483648:
                                    v31 := GenS1{F0: 133}
                                    _ = v31
                                case 42:
                                    var v32 int64 = -126
                                    _ = v32
                                }
                            case 1:
                                for fi33 := 0; fi33 < 2; fi33++ {
                                    var v34 float32 = 3.14159
                                    switch v4 {
                                    case "\n\t":
                                        if (fi27 >= 32767) {
                                            var v35 int16 = v25
                                            switch ((v11 + (v4 + "a")) + v4) {
                                            case "":
                                                v10 = v10
                                            case "hello":
                                                var v36 uint16 = 171
                                                _ = v36
                                            default:
                                                var v37 uint32 = v17
                                                _ = v37
                                            }
                                            switch (v2 - (fi33 | fi27)) {
                                            case 127:
                                                v38 := GenS2{F0: v18}
                                                _ = v38
                                            case 9223372036854775807:
                                                if (v17 > v17) {
                                                    if (42 == v25) {
                                                        fmt.Println(v10)
                                                    } else {
                                                        v11 = v18
                                                        v39 := v1[v17]
                                                        if len(v23) >= 2 {
                                                            v23 = v23[1:len(v23)]
                                                        }
                                                        switch (fi33 - v2) {
                                                        case -1:
                                                            if true {
                                                                if (uint32(v10) > v17) {
                                                                    switch v11 {
                                                                    case "a":
                                                                        if len(v23) >= 2 {
                                                                            v23 = v23[1:len(v23)]
                                                                        }
                                                                    case "\n\t":
                                                                        if (true && (true && (fi33 < (v2 * v2)))) {
                                                                            v40 := GenS5{F0: v3}
                                                                            _ = v40
                                                                        }
                                                                    case "foo bar":
                                                                        v41 := GenS6{F0: 0, F1: 83}
                                                                        _ = v41
                                                                    default:
                                                                        v42 := GenS6{F0: 809, F1: 239}
                                                                        _ = v42
                                                                    }
                                                                }
                                                                var v43 uint32 = v17
                                                                _ = v43
                                                            }
                                                        case -2147483648:
                                                            if (492 != 125) {
                                                                var v44 uint16 = (v10 | v10)
                                                                v45 := GenS1{F0: 2}
                                                                var v46 bool = true
                                                                var v47 uint64 = 138
                                                                v39 = ""
                                                                _ = v44
                                                                _ = v45
                                                                _ = v46
                                                                _ = v47
                                                            } else {
                                                                var v48 float32 = 0.1
                                                                _ = v48
                                                            }
                                                        default:
                                                            fmt.Println(v2, v35, v39)
                                                        }
                                                        if true {
                                                            v17 = (v17 - 42)
                                                            var v49 uint32 = v17
                                                            _ = v49
                                                        }
                                                        _ = v39
                                                    }
                                                    var v50 uint8 = uint8(v10)
                                                    if ((fi27 << ((0) & 63)) >= (fi27 % ((fi33) | 1))) {
                                                        v10 = v10
                                                        var v51 uint = 357
                                                        v10 -= uint16(v17)
                                                        fmt.Println(v3)
                                                        delete(v1, (v17 * v17))
                                                        _ = v51
                                                    } else {
                                                        v35 -= (v35 | int16(v3))
                                                    }
                                                    _ = v50
                                                }
                                            }
                                            _ = v35
                                        }
                                    case "hello":
                                        switch "foo bar" {
                                        case "a":
                                            fmt.Println(v2)
                                        case "hello":
                                            var v52 int = fi27
                                            _ = v52
                                        case "\n\t":
                                            v11 = (("\n\t" + ("\n\t" + ((v4 + v18) + v18))) + ("a" + "a"))
                                            if false {
                                                if len(v23) > 0 {
                                                    v23[0] = -2147483648
                                                }
                                            } else {
                                                _ = len(v1)
                                            }
                                        default:
                                            switch (("hello" + (v18 + "")) + (v4 + "\n\t")) {
                                            case "foo bar":
                                                var v53 uint8 = 42
                                                _ = v53
                                            }
                                        }
                                    case "a":
                                        var v54 string = v18
                                        _ = v54
                                    default:
                                        v1[v17] = ("\n\t" + (("a" + "hello") + v4))
                                    }
                                    _ = fi33
                                    _ = v34
                                }
                            default:
                                var v55 int32 = 1
                                _ = v55
                            }
                            v56 := GenS7{F0: "hello"}
                            _ = v56
                        }
                        _ = fi27
                    }
                case -2147483648:
                    v4 = v4
                default:
                    var v57 float32 = 3.14159
                    _ = v57
                }
                if (true || (true && (753 == 0))) {
                    for fi58 := 0; fi58 < 6; fi58++ {
                        _ = len(v1)
                        _ = fi58
                    }
                    if true {
                        for fi59 := 0; fi59 < 2; fi59++ {
                            var v60 uint8 = 127
                            if false {
                                v1[v17] = ""
                                if true {
                                    v11 = ""
                                }
                                v61 := GenS6{F0: 73, F1: (v60 & (v60 / (((v60 % ((2) | 1))) | 1)))}
                                _ = v61
                            }
                            v18 = "hello"
                            _ = fi59
                            _ = v60
                        }
                        for fi62 := 0; fi62 < 19; fi62++ {
                            v63 := v1[(v17 - 42)]
                            _ = fi62
                            _ = v63
                        }
                        for fi64 := 0; fi64 < 11; fi64++ {
                            switch -1 {
                            case 4294967295:
                                if true {
                                    for fi65 := 0; fi65 < 16; fi65++ {
                                        v11 = v4
                                        fmt.Println(v10, v2, v17)
                                        _ = fi65
                                    }
                                    v66 := GenS6{F0: uint(v10), F1: 1}
                                    var v67 bool = true
                                    _ = v66
                                    _ = v67
                                }
                            case -1:
                                var v68 int = 127
                                _ = v68
                            }
                            v69 := GenS8{F0: 1}
                            _ = fi64
                            _ = v69
                        }
                    } else {
                        fmt.Println(v11, v2, v2)
                        fmt.Println(v10, v2)
                    }
                    if true {
                        v18 = ""
                    }
                }
            } else {
                for fi70 := 0; fi70 < 18; fi70++ {
                    var v71 int32 = v3
                    _ = fi70
                    _ = v71
                }
                v10 = (v10 << ((865) & 63))
                if (false || (true && false)) {
                    var v72 int16 = (v25 ^ (v25 >> ((884) & 63)))
                    v4 = v11
                    var v73 int64 = 371
                    _ = v72
                    _ = v73
                }
            }
        } else {
            v74 := GenS9{F0: ("" + v11)}
            v3 = -1
            _ = v74
        }
    case "":
        if true {
            v18 = ("a" + ("\n\t" + v18))
            v10 += 42
            switch v4 {
            case "\n\t":
                switch -618 {
                case 42:
                    var v75 bool = true
                    _ = v75
                }
            case "a":
                fmt.Println(v18, v4)
            }
        }
    case "a":
        if (!false) {
            for fi76 := 0; fi76 < 20; fi76++ {
                switch "" {
                case "a":
                    v4 = v18
                case "\n\t":
                    v11 = v4
                case "":
                    v77 := GenS10{F0: 42, F1: v10}
                    _ = v77
                default:
                    v1[2147483648] = v18
                }
                if len(v23) > 0 {
                    v23[0] = (v2 & (^(fi76)))
                }
                _ = fi76
            }
            v4 = "a"
            v78 := GenS11{F0: true, F1: int64(v10), F2: 42}
            _ = v78
        }
    }
    for fi79 := 0; fi79 < 6; fi79++ {
        for fi80 := 0; fi80 < 14; fi80++ {
            v81 := GenS12{F0: (!true)}
            fmt.Println(v17, v2)
            _ = fi80
            _ = v81
        }
        for fi82 := 0; fi82 < 12; fi82++ {
            fmt.Println(fi82, fi82, v18)
            _ = fi82
        }
        for fi83 := 0; fi83 < 16; fi83++ {
            if false {
                var v84 int32 = v3
                _ = v84
            }
            _ = len(v23)
            v3 = 1
            _ = fi83
        }
        _ = fi79
    }
    var v85 int8 = 42
    _ = len(v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v85
}
