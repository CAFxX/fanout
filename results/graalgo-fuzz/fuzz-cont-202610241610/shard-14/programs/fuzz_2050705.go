package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 bool; F1 int; F2 string }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 string; F1 string; F2 int32 }
type GenS7 struct { F0 int16; F1 uint32; F2 uint32 }
func genFn0() uint64 {
    v6 := GenS1{F0: 2, F1: 255}
    v7 := GenS1{F0: 1, F1: 65535}
    var v8 uint8 = 255
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := []uint32{4294967295}
    var v2 int8 = 100
    var v3 float32 = 0.1
    v4 := []int{}
    v5 := make(map[string]string)
    v9 := GenS1{F0: 115, F1: 2}
    if (100 > (v2 % ((-56) | 1))) {
        for fi10 := 0; fi10 < 13; fi10++ {
            fmt.Println(v2, v3, v2)
            _ = fi10
        }
        v11 := GenS1{F0: 230, F1: 1}
        if true {
            v12 := GenS2{F0: 0, F1: 0}
            var v13 int32 = 0
            _ = v12
            _ = v13
        }
        _ = v11
    }
    var v14 float32 = v3
    for fi15 := 0; fi15 < 4; fi15++ {
        v14 -= v14
        _ = fi15
    }
    _ = len(v1)
    if (185 > 108) {
        var v16 float64 = (1e10 + (float64((0.1 * v3)) + 0.0))
        switch 610 {
        case -2147483648:
            var v17 int8 = 127
            _ = v17
        case 542:
            v16 = (0.0 - 0.1)
        default:
            fmt.Println(v14, v2, v3)
        }
        _ = v16
    }
    v2 = 42
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 int64 = 42
        var v20 int = -598
        _ = fi18
        _ = v19
        _ = v20
    }
    v3 -= (((v14 / v3) + v3) * ((v3 - v3) + v14))
    if false {
        v2 = 20
        if (65535 >= 0) {
            v21 := GenS1{F0: 32, F1: 1}
            v22 := GenS1{F0: 159, F1: 2}
            _ = v21
            _ = v22
        }
        var v23 bool = (42 <= 9223372036854775807)
        _ = v23
    }
    if false {
        var v24 string = "foo bar"
        switch 128 {
        case -235:
            v3 = ((0.1 + v3) / v3)
        }
        switch (((v24 + (v24 + v24)) + v24) + v24) {
        case "foo bar":
            fmt.Println(v24)
        }
        _ = v24
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        if true {
            v3 = v14
        }
        if true {
            for fi27 := 0; fi27 < 9; fi27++ {
                v2 -= 93
                switch fi26 {
                case 127:
                    v14 += ((1.5 + ((v14 / 1e10) + v3)) / (v14 * -2.25))
                case 32768:
                    fmt.Println(v3, v2, v2)
                default:
                    var v28 float64 = 0.0
                    var v29 float32 = (1.5 - (v14 / v3))
                    _ = v28
                    _ = v29
                }
                _ = fi27
            }
            v2 -= v2
        }
        v30 := GenS3{F0: 255, F1: 536}
        _ = fi26
        _ = v30
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if ((1e10 + 1.5) > (3.14159 * ((0.0 / 1.0) / 3.14159))) {
        if false {
            v31 := GenS3{F0: 2, F1: 65535}
            _ = v31
        }
        for fi32 := 0; fi32 < 13; fi32++ {
            if len(v1) > 0 {
                v1[0] = 42
            }
            switch -509 {
            case -39:
                var v33 int64 = 9223372036854775807
                _ = v33
            default:
                switch fi32 {
                case -1:
                    if (!(2147483647 >= 65535)) {
                        fmt.Println(v14, v2, v14)
                    } else {
                        v14 -= v14
                    }
                case 1:
                    switch "hello" {
                    case "hello":
                        switch "hello" {
                        case "hello":
                            for fi34 := 0; fi34 < 20; fi34++ {
                                fmt.Println(fi34, v3)
                                v35 := GenS4{F0: ((!true) && (v14 != float32(1e10))), F1: fi32, F2: "hello"}
                                fmt.Println(v14, v3, v14)
                                if (661 <= 700) {
                                    var v36 uint32 = 4294967295
                                    _ = v36
                                }
                                _ = fi34
                                _ = v35
                            }
                        default:
                            for fi37 := 0; fi37 < 7; fi37++ {
                                var v38 int16 = 688
                                v39 := v5[("" + "\n\t")]
                                v39 = (v39 + "a")
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                        }
                    case "a":
                        var v40 bool = ((false && false) || ((255 == 1) || false))
                        _ = v40
                    }
                case -390:
                    v3 -= -2.25
                default:
                    for fi41 := 0; fi41 < 9; fi41++ {
                        v42 := GenS2{F0: 1, F1: 0}
                        switch fi41 {
                        case 0:
                            var v43 uint16 = 65535
                            _ = v43
                        case 127:
                            switch "foo bar" {
                            case "":
                                fmt.Println(v2)
                            case "foo bar":
                                switch ((("foo bar" + "\n\t") + "") + ("a" + "\n\t")) {
                                case "hello":
                                    fmt.Println(fi41, fi32)
                                case "foo bar":
                                    v44 := GenS1{F0: 2, F1: 65535}
                                    _ = v44
                                case "":
                                    switch (fi41 & fi32) {
                                    case 4294967295:
                                        v14 -= float32(1.5)
                                    case -159:
                                        v14 = 1e10
                                    default:
                                        _ = len(v4)
                                    }
                                default:
                                    if ((615 >= 1) && (4294967295 != 2147483648)) {
                                        switch fi32 {
                                        case -695:
                                            switch "hello" {
                                            case "hello":
                                                if (0 != 32767) {
                                                    fmt.Println(fi41, v2)
                                                } else {
                                                    v45 := v5["hello"]
                                                    var v46 int8 = (v2 % ((v2) | 1))
                                                    v14 += -2.25
                                                    _ = v45
                                                    _ = v46
                                                }
                                            case "a":
                                                if ((((1.5 * 0.1) * 0.1) == ((float64((0.0 + (-2.25 / v14))) + 1e10) / 3.14159)) || (1 <= 42)) {
                                                    var v47 uint8 = 119
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v14 += 3.14159
                                                    case "hello":
                                                        switch fi32 {
                                                        case 32767:
                                                            v48 := GenS1{F0: 2, F1: 32767}
                                                            _ = v48
                                                        case 2147483647:
                                                            v47 -= v47
                                                        }
                                                    case "":
                                                        v2 = v2
                                                    }
                                                    fmt.Println(v2, fi32)
                                                    _ = v47
                                                } else {
                                                    delete(v5, "\n\t")
                                                    if true {
                                                        var v49 int64 = int64(v2)
                                                        v49 = (v49 ^ -383)
                                                        _ = v49
                                                    } else {
                                                        fmt.Println(v14, fi41)
                                                    }
                                                    var v50 float32 = v14
                                                    _ = v50
                                                }
                                                var v51 int = (fi41 / (((^(fi41))) | 1))
                                                _ = v51
                                            }
                                        default:
                                            fmt.Println(v14, v14)
                                        }
                                    }
                                }
                                _ = len(v5)
                                delete(v5, ("hello" + "hello"))
                                if false {
                                    var v52 uint8 = 127
                                    v52 -= (v52 * (v52 >> ((232) & 63)))
                                    fmt.Println(v14, v52)
                                    _ = v52
                                }
                                var v53 int16 = 581
                                var v54 uint64 = 42
                                _ = v53
                                _ = v54
                            }
                        case 32768:
                            if true {
                                delete(v5, "\n\t")
                                v55 := GenS5{F0: 824, F1: 65535}
                                _ = v55
                            }
                            var v56 int8 = 30
                            if (56 > 2) {
                                if (9223372036854775807 <= 42) {
                                    if true {
                                        var v57 float64 = (0.1 / (-2.25 + 1e10))
                                        switch ("hello" + "a") {
                                        case "\n\t":
                                            fmt.Println(v56, fi32, fi32)
                                        default:
                                            var v58 uint32 = 1
                                            _ = v58
                                        }
                                        v59 := v5["a"]
                                        _ = v57
                                        _ = v59
                                    }
                                    if (true || (!(593 > -927))) {
                                        v60 := GenS1{F0: 0, F1: 1}
                                        _ = v60
                                    } else {
                                        switch int(v56) {
                                        case 605:
                                            _ = len(v4)
                                        }
                                        switch (-(fi41)) {
                                        case 128:
                                            if true {
                                                delete(v5, "a")
                                                var v61 string = ("\n\t" + "hello")
                                                _ = v61
                                            }
                                        default:
                                            switch (fi32 * -936) {
                                            case -477:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                            _ = v56
                        default:
                            fmt.Println(v2, v2)
                        }
                        v62 := GenS6{F0: "", F1: "\n\t", F2: 42}
                        _ = fi41
                        _ = v42
                        _ = v62
                    }
                }
            }
            for fi63 := 0; fi63 < 11; fi63++ {
                fmt.Println(v2)
                _ = fi63
            }
            _ = fi32
        }
        switch -1 {
        case 32768:
            if true {
                var v64 uint8 = 42
                _ = v64
            } else {
                _ = len(v5)
                v14 += v3
            }
        case 127:
            fmt.Println(v2)
        case -431:
            v65 := GenS7{F0: -455, F1: 479, F2: 0}
            _ = v65
        default:
            var v66 int64 = 1
            _ = v66
        }
    }
    var v67 float32 = 1e10
    var v68 float64 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
    _ = v68
    _ = v9
}
