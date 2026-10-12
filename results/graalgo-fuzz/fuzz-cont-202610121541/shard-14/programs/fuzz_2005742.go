package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int8 }

func main() {
    var v1 uint = 654
    var v2 string = ("\n\t" + ("foo bar" + ("a" + "\n\t")))
    var v3 string = "a"
    for fi4 := 0; fi4 < 13; fi4++ {
        v5 := GenS1{F0: (v2 + (v3 + ((v2 + "hello") + (v2 + "hello"))))}
        var v6 string = v3
        fmt.Println(v6, fi4)
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        v3 = "a"
        _ = fi7
    }
    var v8 int = 42
    if false {
        v9 := []int{}
        if false {
            if true {
                if false {
                    var v10 uint8 = 0
                    var v11 int = v8
                    for fi12 := 0; fi12 < 6; fi12++ {
                        if true {
                            if true {
                                fmt.Println(v2, v11)
                            }
                            switch v11 {
                            case 714:
                                _ = len(v9)
                            case 2147483647:
                                for fi13 := 0; fi13 < 17; fi13++ {
                                    fmt.Println(fi12)
                                    _ = fi13
                                }
                            }
                            switch (v3 + v2) {
                            case "a":
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            case "foo bar":
                                var v14 int32 = 42
                                _ = v14
                            case "hello":
                                _ = len(v9)
                            default:
                                if (18446744073709551615 < 1) {
                                    fmt.Println(v8, fi12, v1)
                                    switch "" {
                                    case "foo bar":
                                        var v15 bool = (fi12 >= v11)
                                        _ = v15
                                    case "\n\t":
                                        if (-128 != 122) {
                                            _ = len(v9)
                                            v8 += v8
                                        } else {
                                            var v16 int8 = 100
                                            v17 := GenS1{F0: (v3 + v2)}
                                            _ = v16
                                            _ = v17
                                        }
                                    default:
                                        switch (v8 ^ (fi12 * v8)) {
                                        case -919:
                                            v10 += 1
                                        case 9223372036854775807:
                                            var v18 int64 = 1
                                            _ = v18
                                        case -9223372036854775808:
                                            var v19 uint32 = 207
                                            _ = v19
                                        default:
                                            v20 := GenS2{F0: 32767}
                                            _ = v20
                                        }
                                    }
                                }
                            }
                        } else {
                            _ = len(v9)
                            if ((!true) && true) {
                                switch (v2 + (v2 + "\n\t")) {
                                case "hello":
                                    switch (v11 % (((-(fi12))) | 1)) {
                                    case 546:
                                        v10 = v10
                                    case 32767:
                                        if len(v9) > 0 {
                                            v9[0] = fi12
                                        }
                                    case 9223372036854775807:
                                        switch "" {
                                        case "hello":
                                            var v21 uint64 = 42
                                            _ = v21
                                        case "a":
                                            var v22 uint8 = (^(v10))
                                            _ = v22
                                        default:
                                            v1 -= (v1 / ((18446744073709551615) | 1))
                                        }
                                    default:
                                        switch ("foo bar" + ("" + v2)) {
                                        case "":
                                            fmt.Println(v10, fi12, v11)
                                        case "a":
                                            if (32767 > 0) {
                                                v2 = v2
                                                fmt.Println(v10)
                                                for fi23 := 0; fi23 < 2; fi23++ {
                                                    _ = len(v9)
                                                    _ = fi23
                                                }
                                            } else {
                                                if (!((1e10 - (3.14159 * ((1.5 - float64(-2.25)) - -2.25))) > float64(1e10))) {
                                                    if false {
                                                        switch v3 {
                                                        case "\n\t":
                                                            v2 = ("\n\t" + v3)
                                                        default:
                                                            _ = len(v9)
                                                        }
                                                    } else {
                                                        v8 = 4294967295
                                                        v11 -= (v8 - 128)
                                                        for fi24 := 0; fi24 < 18; fi24++ {
                                                            var v25 uint16 = 255
                                                            v26 := GenS3{F0: (v1 != 18446744073709551615)}
                                                            _ = fi24
                                                            _ = v25
                                                            _ = v26
                                                        }
                                                    }
                                                    v8 = (fi12 ^ 127)
                                                }
                                            }
                                        default:
                                            v9 = append(v9, (v8 - (fi12 >> (((v1 - (^(v1)))) & 63))))
                                        }
                                    }
                                default:
                                    var v27 int8 = -1
                                    _ = v27
                                }
                                var v28 int64 = -1
                                if false {
                                    v29 := GenS4{F0: 850, F1: 331}
                                    for fi30 := 0; fi30 < 10; fi30++ {
                                        var v31 int64 = v28
                                        v28 += v28
                                        if (!true) {
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                            _ = len(v9)
                                        }
                                        var v32 int64 = 1
                                        fmt.Println(v28, v11, v3)
                                        _ = fi30
                                        _ = v31
                                        _ = v32
                                    }
                                    var v33 uint32 = 70
                                    _ = v29
                                    _ = v33
                                }
                                _ = v28
                            }
                            if true {
                                v3 = (v3 + v2)
                                v34 := GenS4{F0: 42, F1: 65535}
                                _ = v34
                            }
                        }
                        _ = fi12
                    }
                    _ = v10
                    _ = v11
                }
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            } else {
                for fi35 := 0; fi35 < 5; fi35++ {
                    fmt.Println(fi35, v2, v8)
                    _ = fi35
                }
            }
            if false {
                if false {
                    _ = len(v9)
                    switch 862 {
                    case -1:
                        v36 := GenS1{F0: "hello"}
                        _ = v36
                    case 114:
                        var v37 int64 = 9223372036854775807
                        _ = v37
                    default:
                        var v38 int64 = 0
                        _ = v38
                    }
                } else {
                    v3 = (v3 + v2)
                    var v39 int64 = 0
                    _ = v39
                }
                var v40 uint = 506
                _ = v40
            }
            switch v3 {
            case "a":
                fmt.Println(v3, v2, v2)
            case "":
                fmt.Println(v8, v3, v1)
            }
        } else {
            v41 := GenS4{F0: 32767, F1: 255}
            _ = v41
        }
        switch 42 {
        case -9223372036854775808:
            if (0 >= 829) {
                if ((0.0 + 1e10) != 0.0) {
                    if false {
                        switch (v3 + v2) {
                        case "foo bar":
                            var v42 uint16 = 250
                            _ = v42
                        case "a":
                            for fi43 := 0; fi43 < 8; fi43++ {
                                v3 = "a"
                                var v44 int = v8
                                _ = fi43
                                _ = v44
                            }
                        }
                    } else {
                        if false {
                            var v45 int16 = 0
                            _ = v45
                        }
                        if ((!(true && (!(42 == 2)))) && false) {
                            switch "hello" {
                            case "hello":
                                _ = len(v9)
                            case "a":
                                fmt.Println(v1, v1, v1)
                            }
                            switch v8 {
                            case 42:
                                switch v3 {
                                case "foo bar":
                                    switch 32768 {
                                    case 128:
                                        v2 = v2
                                    case 4294967295:
                                        var v46 int32 = -2147483648
                                        _ = v46
                                    case -1:
                                        v3 = v2
                                    }
                                case "a":
                                    var v47 float64 = 1.5
                                    _ = v47
                                case "":
                                    var v48 int16 = -857
                                    _ = v48
                                }
                            default:
                                if false {
                                    var v49 int = (v8 << ((65535) & 63))
                                    var v50 uint = (v1 * v1)
                                    _ = v49
                                    _ = v50
                                }
                            }
                            var v51 int16 = -252
                            _ = v51
                        }
                    }
                    for fi52 := 0; fi52 < 2; fi52++ {
                        v2 = (v3 + "hello")
                        v1 += 171
                        _ = fi52
                    }
                    switch "hello" {
                    case "\n\t":
                        v53 := GenS1{F0: "\n\t"}
                        _ = v53
                    }
                }
            } else {
                if (4294967295 < 2) {
                    var v54 bool = (18446744073709551615 != 2)
                    _ = v54
                }
                if ((2 <= 0) && (!(!(626 >= 372)))) {
                    var v55 uint16 = 0
                    _ = v55
                }
                if ((!true) && true) {
                    var v56 uint8 = 222
                    v1 += 908
                    _ = v56
                }
            }
        case 127:
            var v57 int8 = 1
            _ = v57
        case 2147483647:
            v1 -= (v1 % ((v1) | 1))
        default:
            for fi58 := 0; fi58 < 15; fi58++ {
                if len(v9) > 0 {
                    v9[0] = v8
                }
                for fi59 := 0; fi59 < 19; fi59++ {
                    v60 := GenS5{F0: 64, F1: 1}
                    var v61 uint16 = 209
                    var v62 uint32 = 630
                    var v63 string = v2
                    if true {
                        v62 = 2147483647
                        switch -2147483648 {
                        case 128:
                            var v64 int8 = int8(v62)
                            _ = v64
                        default:
                            var v65 float64 = 0.0
                            _ = v65
                        }
                    }
                    if len(v9) > 0 {
                        v9[0] = (fi58 % ((v8) | 1))
                    }
                    v63 = ""
                    if ((fi58 * (-(fi59))) <= fi58) {
                        v1 += (^(v1))
                    } else {
                        var v66 string = ("" + v63)
                        _ = v66
                    }
                    _ = fi59
                    _ = v60
                    _ = v61
                    _ = v62
                    _ = v63
                }
                v9 = append(v9, fi58)
                _ = fi58
            }
        }
        _ = v9
    }
    switch -748 {
    case 809:
        var v67 uint8 = 128
        _ = v67
    default:
        var v68 uint32 = 135
        _ = v68
    }
    if false {
        switch "a" {
        case "foo bar":
            if (18446744073709551615 < 943) {
                v8 -= (v8 % ((v8) | 1))
                var v69 int8 = -128
                _ = v69
            }
        }
        v70 := []int{(v8 | v8), v8, v8}
        _ = v70
    }
    switch "hello" {
    case "":
        v71 := []int{v8, v8}
        _ = v71
    }
    v8 = 879
    switch (v8 / ((v8) | 1)) {
    case 32767:
        for fi72 := 0; fi72 < 10; fi72++ {
            v73 := []int{-1, 9223372036854775807, v8, 42}
            _ = fi72
            _ = v73
        }
    }
    v74 := GenS3{F0: ((2147483647 >= 65535) || false)}
    if (1 < 146) {
        v1 -= 734
        v3 = v2
        var v75 uint = (v1 + (v1 / ((v1) | 1)))
        _ = v75
    } else {
        switch v8 {
        case -1:
            switch v8 {
            case 2147483647:
                var v76 float64 = float64(-2.25)
                _ = v76
            }
        case 1:
            fmt.Println(v1)
        case 313:
            v1 -= v1
        }
        var v77 string = v3
        _ = v77
    }
    v3 = v3
    for fi78 := 0; fi78 < 3; fi78++ {
        var v79 uint8 = 183
        fmt.Println(v8)
        _ = fi78
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v74
    _ = v8
}
