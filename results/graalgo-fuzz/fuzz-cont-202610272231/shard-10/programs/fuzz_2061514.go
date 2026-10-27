package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 uint }

func main() {
    var v1 uint64 = 1
    var v2 float32 = (-2.25 * 0.1)
    var v3 uint8 = 2
    var v4 uint16 = 65535
    var v5 int64 = -620
    v6 := GenS1{F0: v1}
    for fi7 := 0; fi7 < 14; fi7++ {
        switch (fi7 >> ((255) & 63)) {
        case 32768:
            switch "\n\t" {
            case "foo bar":
                v8 := GenS1{F0: v1}
                _ = v8
            default:
                for fi9 := 0; fi9 < 2; fi9++ {
                    v10 := []int{}
                    _ = fi9
                    _ = v10
                }
            }
        case -783:
            v11 := []int{}
            _ = v11
        case -111:
            var v12 uint64 = (v1 | (v1 / ((0) | 1)))
            _ = v12
        }
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS1{F0: (v1 - 42)}
                _ = fi13
                _ = v14
            }
        case "\n\t":
            var v15 float64 = 0.0
            _ = v15
        default:
            var v16 bool = ((v1 % (((v1 + v1)) | 1)) != (v1 + 42))
            _ = v16
        }
        var v17 int = (fi7 + fi7)
        _ = fi7
        _ = v17
    }
    fmt.Println(v1, v1, v1)
    switch 303 {
    case -152:
        for fi18 := 0; fi18 < 13; fi18++ {
            if (!false) {
                var v19 int64 = (v5 << ((65535) & 63))
                var v20 int32 = -266
                var v21 string = "foo bar"
                _ = v19
                _ = v20
                _ = v21
            }
            var v22 uint = 534
            _ = fi18
            _ = v22
        }
    }
    var v23 float32 = (3.14159 - ((0.1 - ((0.1 - -2.25) / v2)) + v2))
    var v24 uint64 = v1
    v25 := GenS1{F0: v24}
    switch 2147483647 {
    case -223:
        for fi26 := 0; fi26 < 16; fi26++ {
            for fi27 := 0; fi27 < 7; fi27++ {
                if false {
                    var v28 int16 = 1
                    switch ("hello" + ("\n\t" + "")) {
                    case "\n\t":
                        var v29 string = ("hello" + "")
                        _ = v29
                    case "hello":
                        v5 += int64(v28)
                    case "":
                        v3 = (v3 << ((1) & 63))
                    }
                    _ = v28
                } else {
                    switch "a" {
                    case "a":
                        switch -9223372036854775808 {
                        case 380:
                            v1 = (v24 << ((1) & 63))
                        default:
                            v30 := []int{(fi26 << ((4294967295) & 63)), (fi26 & fi26), (fi26 >> ((460) & 63)), (fi27 % (((fi27 - (^(fi27)))) | 1))}
                            _ = v30
                        }
                    }
                    if (-470 == 0) {
                        if ((-110 < -1) && (fi26 != (fi26 | -704))) {
                            switch (fi26 * fi27) {
                            case 170:
                                switch ((("a" + "\n\t") + "a") + "\n\t") {
                                case "hello":
                                    switch (("a" + ("foo bar" + "\n\t")) + "a") {
                                    case "hello":
                                        var v31 uint16 = v4
                                        _ = v31
                                    default:
                                        v32 := []int{(fi27 | fi27), fi26, 454}
                                        _ = v32
                                    }
                                }
                            case 127:
                                fmt.Println(v1)
                                v33 := GenS2{F0: ((("foo bar" + ("a" + "a")) + "hello") + "hello"), F1: 807}
                                if false {
                                    if true {
                                        if (v5 <= (v5 ^ v5)) {
                                            v34 := []int{42, (fi27 % ((0) | 1)), (fi27 | -9223372036854775808)}
                                            _ = v34
                                        } else {
                                            var v35 float64 = (0.1 / (1.5 - (float64((v23 + (v2 * 1e10))) - 1.5)))
                                            var v36 uint8 = 97
                                            v3 = (^(v3))
                                            _ = v35
                                            _ = v36
                                        }
                                        v37 := GenS3{F0: 127, F1: 2}
                                        switch ("foo bar" + "foo bar") {
                                        case "hello":
                                            v5 = (v5 >> ((103) & 63))
                                        default:
                                            fmt.Println(v24, v3, v1)
                                        }
                                        _ = v37
                                    }
                                    v24 += v1
                                }
                                _ = v33
                            default:
                                v5 = int64(v3)
                                v1 = v24
                            }
                            v24 = 1
                            var v38 int32 = -1
                            _ = v38
                        }
                        v5 += (v5 - v5)
                        fmt.Println(v5)
                    } else {
                        if false {
                            if true {
                                fmt.Println(v1)
                                v3 = (^(v3))
                            } else {
                                v39 := GenS4{F0: 1}
                                if ((v5 | int64(v24)) <= (v5 % ((-1) | 1))) {
                                    if true {
                                        if (v5 > v5) {
                                            var v40 uint = 1
                                            _ = v40
                                        } else {
                                            fmt.Println(v3)
                                        }
                                        var v41 bool = false
                                        _ = v41
                                    }
                                    v1 = 0
                                    var v42 int64 = -905
                                    _ = v42
                                }
                                switch 0 {
                                case 32768:
                                    v43 := GenS4{F0: 1}
                                    _ = v43
                                case 535:
                                    if (0 == v5) {
                                        fmt.Println(v1)
                                        v44 := []int{}
                                        fmt.Println(v1, v23)
                                        if len(v44) >= 2 {
                                            v44 = v44[1:len(v44)]
                                        }
                                        if (!(false || (0.0 < 1e10))) {
                                            var v45 bool = true
                                            var v46 float32 = v23
                                            _ = v45
                                            _ = v46
                                        }
                                        _ = v44
                                    }
                                }
                                var v47 bool = true
                                if ((^(v4)) > v4) {
                                    var v48 int32 = 36
                                    _ = v48
                                } else {
                                    var v49 float32 = (float32(1.5) - ((v2 - 1e10) - ((0.1 + 1.5) + -2.25)))
                                    var v50 uint32 = uint32(v4)
                                    v51 := []int{149, (fi27 * (fi27 >> ((2) & 63))), 9223372036854775807}
                                    switch (^(fi27)) {
                                    case 822:
                                        fmt.Println(fi26)
                                    case 18:
                                        v52 := GenS1{F0: 2}
                                        _ = v52
                                    case 328:
                                        var v53 int16 = 1
                                        _ = v53
                                    }
                                    var v54 string = "a"
                                    if v47 {
                                        var v55 int8 = 43
                                        if len(v51) > 0 {
                                            v51[0] = (fi27 & -1)
                                        }
                                        v51 = append(v51, (fi26 >> ((uint(v50)) & 63)))
                                        if (!((v55 < 1) || (((^(v24)) != v1) || (!(!v47))))) {
                                            var v56 uint8 = (^(v3))
                                            v50 = 2
                                            var v57 int64 = -1
                                            _ = v56
                                            _ = v57
                                        } else {
                                            v5 -= (^(v5))
                                            if (!(!(v47 || (!(-1 == -2147483648))))) {
                                                v58 := GenS1{F0: 18446744073709551615}
                                                fmt.Println(v54, v54, fi27)
                                                if len(v51) > 0 {
                                                    v51[0] = 128
                                                }
                                                _ = v58
                                            }
                                            v47 = v47
                                        }
                                        v50 = (v50 * v50)
                                        _ = v55
                                    }
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                    _ = v54
                                }
                                _ = v39
                                _ = v47
                            }
                            var v59 uint64 = v24
                            _ = v59
                        }
                        if (!true) {
                            if true {
                                v60 := []int{}
                                _ = v60
                            }
                            var v61 uint64 = (v24 * (v24 | uint64(v3)))
                            v23 += -2.25
                            _ = v61
                        }
                    }
                    if (!false) {
                        v62 := GenS5{F0: 1000, F1: 42}
                        _ = v62
                    }
                    v63 := GenS4{F0: 124}
                    v64 := []int{128, -912, 42, fi26}
                    _ = v63
                    _ = v64
                }
                _ = fi27
            }
            var v65 bool = (false || true)
            _ = fi26
            _ = v65
        }
    case 2147483647:
        for fi66 := 0; fi66 < 1; fi66++ {
            v67 := []int{609, (fi66 | fi66), (fi66 - int(v4)), fi66}
            v24 += v1
            _ = fi66
            _ = v67
        }
    default:
        v68 := []int{}
        _ = v68
    }
    v69 := GenS2{F0: "", F1: 1}
    v70 := []int{-931}
    for fi71 := 0; fi71 < 15; fi71++ {
        if len(v70) >= 2 {
            v70 = v70[1:len(v70)]
        }
        for fi72 := 0; fi72 < 17; fi72++ {
            v73 := GenS1{F0: v24}
            v24 = 0
            _ = fi72
            _ = v73
        }
        _ = fi71
    }
    for fi74 := 0; fi74 < 17; fi74++ {
        var v75 uint8 = (v3 % ((v3) | 1))
        v24 = (^(v1))
        v2 = (1e10 - 0.1)
        _ = fi74
        _ = v75
    }
    switch -218 {
    case -2147483648:
        var v76 uint64 = 2
        _ = v76
    }
    v4 = (v4 / (((v4 | v4)) | 1))
    v77 := GenS4{F0: 46}
    var v78 string = "a"
    v79 := GenS2{F0: v78, F1: 1}
    var v80 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v70:", len(v70), v70)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
    _ = v77
    _ = v78
    _ = v79
    _ = v80
}
