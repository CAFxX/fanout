package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    var v3 uint32 = 1
    var v4 uint32 = (v3 + 2147483647)
    fmt.Println(v4, v3)
    var v5 uint32 = 65535
    v2[-328] = (!(false || ((4294967295 != 65535) && ("a" <= ("hello" + "")))))
    if (((true && false) && true) || false) {
        if (false || (382 != 766)) {
            switch "a" {
            case "foo bar":
                var v6 int32 = 2147483647
                _ = v6
            case "hello":
                for fi7 := 0; fi7 < 7; fi7++ {
                    var v8 uint64 = 2
                    if (4294967295 != (^(v8))) {
                        var v9 int64 = 1
                        _ = len(v1)
                        if (-977 >= -2147483648) {
                            for fi10 := 0; fi10 < 10; fi10++ {
                                v11 := []int{(^(fi10)), 13, fi10, 4294967295}
                                _ = fi10
                                _ = v11
                            }
                            fmt.Println(fi7)
                        }
                        _ = v9
                    }
                    _ = fi7
                    _ = v8
                }
            case "\n\t":
                for fi12 := 0; fi12 < 11; fi12++ {
                    var v13 int64 = -9223372036854775808
                    switch ("\n\t" + "") {
                    case "a":
                        for fi14 := 0; fi14 < 17; fi14++ {
                            v13 = (^(v13))
                            if ((65535 > 757) && false) {
                                v15 := []int{(fi14 | 127)}
                                _ = v15
                            } else {
                                _ = len(v1)
                                if (-2147483648 > -585) {
                                    if true {
                                        v16 := []int{}
                                        _ = v16
                                    } else {
                                        switch (-(fi14)) {
                                        case -1:
                                            var v17 string = ""
                                            var v18 int = fi14
                                            switch (fi12 << ((255) & 63)) {
                                            case 2147483647:
                                                if false {
                                                    if false {
                                                        if false {
                                                            var v19 uint = 18446744073709551615
                                                            _ = v19
                                                        }
                                                    }
                                                    var v20 bool = (390 > 4294967295)
                                                    fmt.Println(v5, v5, v17)
                                                    _ = v20
                                                } else {
                                                    switch fi12 {
                                                    case 42:
                                                        var v21 uint32 = v4
                                                        _ = v21
                                                    case -9223372036854775808:
                                                        v22 := []int{}
                                                        _ = v22
                                                    default:
                                                        var v23 float32 = (0.1 * float32(float64(0.1)))
                                                        _ = v23
                                                    }
                                                    v24 := GenS1{F0: (fi14 / ((420) | 1))}
                                                    v25 := []int{}
                                                    _ = v24
                                                    _ = v25
                                                }
                                            default:
                                                v13 = v13
                                            }
                                            _ = v17
                                            _ = v18
                                        case 42:
                                            fmt.Println(fi12)
                                            var v26 float32 = -2.25
                                            _ = v26
                                        case 9223372036854775807:
                                            if false {
                                                v27 := []int{fi14, 1, (^(fi12))}
                                                var v28 uint16 = 255
                                                var v29 uint64 = 693
                                                v28 -= (v28 / ((v28) | 1))
                                                v5 = 4294967295
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                            } else {
                                                v30 := GenS2{F0: -14}
                                                _ = v30
                                            }
                                        default:
                                            var v31 int8 = -128
                                            v32 := []int{(fi12 & 400)}
                                            fmt.Println(v5)
                                            switch ("" + "\n\t") {
                                            case "a":
                                                switch (^(fi14)) {
                                                case 205:
                                                    var v33 int64 = (v13 / ((v13) | 1))
                                                    _ = v33
                                                default:
                                                    switch -259 {
                                                    case 2147483647:
                                                        if ((fi12 & 9223372036854775807) > (fi12 % ((fi12) | 1))) {
                                                            if (true || (-2.25 > 3.14159)) {
                                                                v3 = 2147483647
                                                                var v34 uint32 = (^(v5))
                                                                v4 = v5
                                                                _ = v34
                                                            }
                                                            var v35 uint = 0
                                                            _ = v35
                                                        } else {
                                                            fmt.Println(v31, v3, v5)
                                                        }
                                                    case -9223372036854775808:
                                                        v3 -= 65535
                                                    }
                                                }
                                            default:
                                                var v36 uint16 = 730
                                                _ = v36
                                            }
                                            _ = v31
                                            _ = v32
                                        }
                                        v37 := []int{-25}
                                        var v38 int64 = v13
                                        _ = v37
                                        _ = v38
                                    }
                                    if false {
                                        switch 32767 {
                                        case 1:
                                            v39 := GenS3{F0: 4294967295, F1: 0}
                                            _ = v39
                                        default:
                                            fmt.Println(v13, v3)
                                        }
                                    }
                                    v5 = (v5 << ((0) & 63))
                                }
                            }
                            _ = fi14
                        }
                    case "foo bar":
                        if false {
                            if (false && false) {
                                v40 := GenS4{F0: 1, F1: 4294967295}
                                _ = v40
                            } else {
                                var v41 int = (fi12 >> ((2) & 63))
                                switch fi12 {
                                case 2147483647:
                                    fmt.Println(v4, fi12)
                                case 520:
                                    for fi42 := 0; fi42 < 2; fi42++ {
                                        v13 += (v13 ^ (v13 % ((1) | 1)))
                                        _ = fi42
                                    }
                                default:
                                    if (((v5 + (v4 * v4)) == 65535) && (1.5 < (1e10 / 1.5))) {
                                        v43 := GenS4{F0: 255, F1: 837}
                                        v5 = (^(v3))
                                        var v44 string = "\n\t"
                                        _ = v43
                                        _ = v44
                                    } else {
                                        var v45 float32 = (float32(((1.5 * (((3.14159 + float64(1.5)) - 1.5) - ((3.14159 * 0.0) - -2.25))) - ((0.1 * (0.1 * 0.0)) - 0.0))) / 1.5)
                                        if (!true) {
                                            v46 := GenS1{F0: v41}
                                            v47 := []int{v41, fi12, v41, 32767}
                                            _ = len(v2)
                                            _ = v46
                                            _ = v47
                                        } else {
                                            v48 := GenS2{F0: 100}
                                            for fi49 := 0; fi49 < 5; fi49++ {
                                                v3 += (v5 | (v3 * (v3 << ((2) & 63))))
                                                _ = fi49
                                            }
                                            var v50 float32 = v45
                                            _ = v48
                                            _ = v50
                                        }
                                        v51 := GenS4{F0: 1, F1: 160}
                                        _ = v45
                                        _ = v51
                                    }
                                }
                                var v52 float64 = 0.1
                                _ = v41
                                _ = v52
                            }
                        }
                    default:
                        for fi53 := 0; fi53 < 11; fi53++ {
                            v3 += (v3 / ((v3) | 1))
                            switch "foo bar" {
                            case "a":
                                v54 := GenS4{F0: 1, F1: 687}
                                _ = v54
                            }
                            v5 += v5
                            _ = fi53
                        }
                    }
                    v55 := GenS5{F0: 2}
                    _ = fi12
                    _ = v13
                    _ = v55
                }
            default:
                if false {
                    for fi56 := 0; fi56 < 14; fi56++ {
                        v4 -= (v5 ^ v3)
                        v4 = 2147483648
                        _ = fi56
                    }
                    switch (("\n\t" + "\n\t") + "\n\t") {
                    case "\n\t":
                        for fi57 := 0; fi57 < 8; fi57++ {
                            switch (-(fi57)) {
                            case 330:
                                for fi58 := 0; fi58 < 16; fi58++ {
                                    v59 := v2[(fi58 ^ fi58)]
                                    var v60 uint32 = (v3 ^ 1)
                                    switch (("" + ("" + "foo bar")) + ("a" + "a")) {
                                    case "hello":
                                        v3 -= (v4 >> ((85) & 63))
                                    }
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                            }
                            _ = fi57
                        }
                    }
                    v61 := []int{-932, 2147483647, 4294967295}
                    _ = v61
                } else {
                    v62 := GenS6{F0: "hello"}
                    _ = v62
                }
            }
        }
        var v63 uint = 65535
        if true {
            v5 = 2
            v3 -= (^(v5))
            for fi64 := 0; fi64 < 4; fi64++ {
                v65 := []int{int(v4), fi64, 0, fi64}
                var v66 int8 = 1
                _ = fi64
                _ = v65
                _ = v66
            }
        } else {
            if false {
                switch "foo bar" {
                case "":
                    switch 540 {
                    case 947:
                        v4 -= v3
                    case 294:
                        var v67 uint64 = 992
                        _ = v67
                    }
                case "a":
                    for fi68 := 0; fi68 < 20; fi68++ {
                        switch ("" + "foo bar") {
                        case "\n\t":
                            v69 := []int{-1}
                            _ = v69
                        case "a":
                            v70 := []int{(-(fi68)), fi68}
                            _ = v70
                        case "foo bar":
                            v63 -= 255
                        }
                        _ = fi68
                    }
                }
                v63 = v63
            } else {
                fmt.Println(v3, v3)
                v5 += (^(v4))
            }
        }
        _ = v63
    }
    switch "foo bar" {
    case "foo bar":
        if false {
            v4 = (v4 / ((v3) | 1))
        } else {
            var v71 bool = true
            for fi72 := 0; fi72 < 11; fi72++ {
                v73 := v1["hello"]
                v74 := GenS7{F0: 63}
                if (-128 == 115) {
                    fmt.Println(v71, v5, v71)
                    fmt.Println(v3, v3, v5)
                    fmt.Println(v5, v4)
                } else {
                    var v75 uint64 = 42
                    _ = v75
                }
                _ = fi72
                _ = v73
                _ = v74
            }
            _ = v71
        }
    case "\n\t":
        var v76 int = -164
        _ = v76
    }
    v77 := []int{421, -9223372036854775808}
    v3 += v4
    switch 679 {
    case 4294967295:
        switch ("foo bar" + ("" + "foo bar")) {
        case "hello":
            var v78 uint16 = 342
            _ = v78
        case "foo bar":
            if false {
                var v79 uint8 = 127
                _ = v79
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v77:", len(v77), v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
}
