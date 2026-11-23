package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64; F1 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 1
    var v3 uint = 18446744073709551615
    if false {
        v4 := GenS1{F0: 32767, F1: (^(v2)), F2: 42}
        v2 = 127
        for fi5 := 0; fi5 < 17; fi5++ {
            switch ("" + "a") {
            case "":
                var v6 uint32 = 2147483647
                _ = v6
            }
            for fi7 := 0; fi7 < 10; fi7++ {
                v8 := GenS1{F0: 65535, F1: (v2 % ((v2) | 1)), F2: (fi5 + (^(fi7)))}
                _ = fi7
                _ = v8
            }
            _ = fi5
        }
        _ = v4
    }
    v9 := GenS1{F0: 32767, F1: v2, F2: -2147483648}
    v3 = (v3 >> ((v3) & 63))
    switch -1 {
    case 2147483647:
        if (1 != 915) {
            switch "" {
            case "":
                var v10 int = 0
                _ = v10
            case "hello":
                v2 -= v2
            case "a":
                for fi11 := 0; fi11 < 5; fi11++ {
                    if true {
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(v3, fi11)
                        case "hello":
                            if false {
                                var v12 bool = false
                                var v13 uint64 = 18446744073709551615
                                switch fi11 {
                                case 2147483647:
                                    var v14 string = ("foo bar" + "hello")
                                    _ = v14
                                case 127:
                                    if (v12 && (!(((fi11 * fi11) >= fi11) || (42 >= 4294967295)))) {
                                        var v15 float64 = 1e10
                                        _ = v15
                                    }
                                default:
                                    for fi16 := 0; fi16 < 10; fi16++ {
                                        var v17 int = fi11
                                        _ = fi16
                                        _ = v17
                                    }
                                }
                                _ = v12
                                _ = v13
                            }
                        default:
                            fmt.Println(v2)
                        }
                        var v18 int64 = 661
                        v2 = 0
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 6; fi19++ {
                            v2 += v2
                            fmt.Println(fi11, fi11, v2)
                            _ = fi19
                        }
                    }
                    for fi20 := 0; fi20 < 14; fi20++ {
                        delete(v1, 0)
                        _ = fi20
                    }
                    _ = fi11
                }
            }
            var v21 float32 = 0.0
            _ = v21
        }
    }
    var v22 float64 = 3.14159
    switch "a" {
    case "foo bar":
        var v23 string = "a"
        _ = v23
    case "hello":
        switch 32768 {
        case -2147483648:
            v3 = (v3 % ((42) | 1))
        }
    case "a":
        switch ("\n\t" + "hello") {
        case "":
            switch "a" {
            case "\n\t":
                if true {
                    for fi24 := 0; fi24 < 2; fi24++ {
                        v1[65535] = ((("foo bar" + (("foo bar" + "a") + ("" + ("a" + "hello")))) + "\n\t") + ((("hello" + "a") + "\n\t") + "foo bar"))
                        if false {
                            switch fi24 {
                            case -1:
                                v3 = v3
                            case 901:
                                for fi25 := 0; fi25 < 2; fi25++ {
                                    fmt.Println(v3, fi25)
                                    _ = fi25
                                }
                            }
                            switch fi24 {
                            case 0:
                                switch "" {
                                case "hello":
                                    v26 := GenS1{F0: 279, F1: (^(v2)), F2: -1}
                                    _ = v26
                                case "a":
                                    for fi27 := 0; fi27 < 4; fi27++ {
                                        switch (fi27 | 2147483647) {
                                        case -9223372036854775808:
                                            fmt.Println(fi24, v22, v22)
                                        case 2147483647:
                                            fmt.Println(fi24, fi27)
                                        }
                                        _ = fi27
                                    }
                                case "":
                                    if (!false) {
                                        v28 := GenS2{F0: 549, F1: false}
                                        for fi29 := 0; fi29 < 14; fi29++ {
                                            switch (("foo bar" + "a") + "") {
                                            case "foo bar":
                                                v1[599] = "a"
                                            case "\n\t":
                                                v30 := GenS3{F0: v2, F1: fi29, F2: 0}
                                                _ = v30
                                            }
                                            var v31 string = ("" + (("foo bar" + ("a" + "\n\t")) + ""))
                                            if (128 != fi29) {
                                                delete(v1, 4294967295)
                                                v31 = "hello"
                                                v22 = (((0.1 + 1.5) + (((((0.0 - float64((3.14159 / (3.14159 / float32(0.1))))) / -2.25) + (0.1 - (v22 + 1.5))) - (v22 + (float64(3.14159) + (float64(0.0) * v22)))) - float64(0.1))) + (1.5 * 3.14159))
                                                var v32 float64 = (1e10 - v22)
                                                _ = v32
                                            } else {
                                                v33 := []int{fi24, fi24, fi24, fi24}
                                                if len(v33) >= 2 {
                                                    v33 = v33[1:len(v33)]
                                                }
                                                _ = v33
                                            }
                                            if (0.1 != (((-2.25 / 3.14159) - 1e10) / (1.5 * (float32(0.0) - float32(((3.14159 * 0.0) - 1.5)))))) {
                                                if false {
                                                    var v34 string = v31
                                                    if false {
                                                        v31 = v31
                                                    }
                                                    v35 := v1[42]
                                                    _ = v34
                                                    _ = v35
                                                }
                                                var v36 float32 = float32(1.5)
                                                _ = v36
                                            }
                                            if true {
                                                fmt.Println(v31, v3)
                                                if true {
                                                    v1[1] = ("foo bar" + v31)
                                                    fmt.Println(fi29, v22, v3)
                                                    v31 = v31
                                                } else {
                                                    var v37 string = v31
                                                    _ = v37
                                                }
                                            } else {
                                                v38 := GenS4{F0: 0}
                                                _ = v38
                                            }
                                            _ = fi29
                                            _ = v31
                                        }
                                        v2 = v2
                                        _ = v28
                                    }
                                default:
                                    for fi39 := 0; fi39 < 1; fi39++ {
                                        fmt.Println(v2, v22)
                                        switch (fi24 >> ((v3) & 63)) {
                                        case 0:
                                            var v40 bool = true
                                            _ = v40
                                        case 127:
                                            v41 := GenS3{F0: v2, F1: fi24, F2: 100}
                                            _ = v41
                                        }
                                        v42 := GenS4{F0: 1}
                                        v43 := GenS5{F0: -539}
                                        _ = fi39
                                        _ = v42
                                        _ = v43
                                    }
                                }
                            case 9223372036854775807:
                                for fi44 := 0; fi44 < 16; fi44++ {
                                    var v45 int64 = 0
                                    v22 += v22
                                    _ = fi44
                                    _ = v45
                                }
                            default:
                                for fi46 := 0; fi46 < 15; fi46++ {
                                    var v47 string = (("a" + "\n\t") + "")
                                    var v48 uint8 = 127
                                    switch (fi24 + -168) {
                                    case 9223372036854775807:
                                        v49 := GenS6{F0: 1, F1: 0}
                                        _ = v49
                                    case 32767:
                                        if (!(2147483648 < 65535)) {
                                            v3 = 65535
                                            var v50 uint8 = v48
                                            _ = v50
                                        } else {
                                            v51 := GenS2{F0: -9223372036854775808, F1: false}
                                            v22 = v22
                                            v22 += (1.5 * v22)
                                            _ = v51
                                        }
                                    }
                                    _ = fi46
                                    _ = v47
                                    _ = v48
                                }
                            }
                        } else {
                            fmt.Println(v22, fi24, v22)
                        }
                        _ = fi24
                    }
                    v52 := GenS4{F0: 553}
                    _ = v52
                } else {
                    var v53 bool = ((!(-191 == -2147483648)) || (111 < -15))
                    _ = v53
                }
            case "a":
                for fi54 := 0; fi54 < 3; fi54++ {
                    switch fi54 {
                    case -9223372036854775808:
                        for fi55 := 0; fi55 < 6; fi55++ {
                            v22 = 1.5
                            _ = fi55
                        }
                    case -224:
                        v3 += v3
                    }
                    if false {
                        v3 += (^(v3))
                    } else {
                        if true {
                            if ((fi54 & fi54) <= 32767) {
                                v1[65535] = ("hello" + "foo bar")
                                var v56 float64 = 1e10
                                _ = v56
                            }
                            for fi57 := 0; fi57 < 17; fi57++ {
                                fmt.Println(v2)
                                _ = fi57
                            }
                            var v58 int64 = 1
                            _ = v58
                        }
                        for fi59 := 0; fi59 < 16; fi59++ {
                            v60 := GenS4{F0: -1}
                            _ = fi59
                            _ = v60
                        }
                        switch ("hello" + "") {
                        case "\n\t":
                            if false {
                                v61 := v1[2]
                                if false {
                                    var v62 float64 = ((v22 * v22) * v22)
                                    _ = v62
                                } else {
                                    switch (v61 + (v61 + (("a" + "foo bar") + ((v61 + (v61 + "")) + v61)))) {
                                    case "":
                                        fmt.Println(v3, v22, v3)
                                    case "foo bar":
                                        v63 := GenS1{F0: 0, F1: 0, F2: int(v2)}
                                        _ = v63
                                    case "\n\t":
                                        v3 = v3
                                    }
                                    v61 = v61
                                }
                                _ = v61
                            }
                        case "":
                            if ((324 < 255) || true) {
                                var v64 bool = (100 >= -1)
                                _ = v64
                            } else {
                                for fi65 := 0; fi65 < 18; fi65++ {
                                    var v66 uint = v3
                                    _ = fi65
                                    _ = v66
                                }
                                var v67 uint64 = 4294967295
                                fmt.Println(v2)
                                _ = v67
                            }
                        case "hello":
                            if true {
                                var v68 uint32 = 4294967295
                                _ = v68
                            }
                        }
                    }
                    _ = fi54
                }
            case "":
                var v69 uint = 18446744073709551615
                _ = v69
            }
        case "hello":
            var v70 uint = uint(v2)
            _ = v70
        default:
            var v71 string = ("a" + "a")
            _ = v71
        }
    default:
        v72 := []int{2147483647, 974}
        _ = v72
    }
    v3 -= uint(v2)
    fmt.Println(v3, v2, v3)
    v73 := GenS4{F0: -2147483648}
    v74 := []int{-119}
    for fi75 := 0; fi75 < 11; fi75++ {
        if ((!(("a" != "") && (!(175 != 380)))) && (!(2147483648 == 1))) {
            if true {
                delete(v1, 224)
                v74 = append(v74, 2147483647)
                if ((fi75 | fi75) >= fi75) {
                    var v76 bool = (v22 < (3.14159 + (v22 + v22)))
                    v77 := GenS5{F0: 1000}
                    _ = v76
                    _ = v77
                } else {
                    for fi78 := 0; fi78 < 15; fi78++ {
                        if len(v74) > 0 {
                            v74[0] = (fi75 / (((fi75 >> (((v3 / (((v3 & (v3 / (((v3 + (v3 >> ((v3) & 63)))) | 1)))) | 1))) & 63))) | 1))
                        }
                        fmt.Println(v22, v3)
                        switch "" {
                        case "hello":
                            fmt.Println(v3, fi78, v2)
                        }
                        v1[42] = ""
                        if len(v74) > 0 {
                            v74[0] = fi75
                        }
                        _ = fi78
                    }
                    for fi79 := 0; fi79 < 17; fi79++ {
                        v3 -= 2
                        _ = fi79
                    }
                }
            }
            switch ("" + "foo bar") {
            case "\n\t":
                v74 = append(v74, fi75)
            }
        }
        fmt.Println(fi75, v2, v2)
        var v80 float64 = ((v22 + float64(1e10)) * ((v22 / (v22 / v22)) * v22))
        _ = fi75
        _ = v80
    }
    fmt.Println(v2)
    _ = len(v74)
    v22 -= (1e10 * 1.5)
    fmt.Println(v3, v3, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v74:", len(v74), v74)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v73
    _ = v74
    _ = v9
}
