package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 uint16; F2 uint32 }
type GenS7 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint32, p1 int8) int {
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 int16 = 32767
        fmt.Println(fi4, v5, p1)
        var v6 int8 = p1
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint32 = (^(p0))
    var v8 uint64 = 771
    v9 := GenS1{F0: 42, F1: (p1 * (p1 / ((p1) | 1)))}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return int(p1)
}

func main() {
    v1 := make(map[string]string)
    var v2 string = "\n\t"
    var v3 uint = 515
    switch "hello" {
    case "foo bar":
        if false {
            switch v2 {
            case "foo bar":
                v10 := GenS1{F0: 42, F1: -1}
                _ = v10
            case "hello":
                v3 -= (^(v3))
            }
            var v11 int32 = 0
            for fi12 := 0; fi12 < 20; fi12++ {
                if (false && (true && (42 >= -1))) {
                    v11 = -868
                }
                fmt.Println(fi12, v11, v11)
                if (-2.25 > 0.0) {
                    fmt.Println(v3, v2)
                }
                _ = fi12
            }
            _ = v11
        } else {
            var v13 uint8 = 0
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 uint16 = 65535
                v2 = v2
                switch (fi14 - fi14) {
                case 295:
                    if ((-2.25 / 1e10) <= 0.1) {
                        switch "" {
                        case "a":
                            if (v3 == 1) {
                                v2 = v2
                                if (!(9223372036854775807 <= -10)) {
                                    for fi16 := 0; fi16 < 9; fi16++ {
                                        switch (((v2 + (v2 + ("a" + ""))) + v2) + "") {
                                        case "hello":
                                            delete(v1, v2)
                                        case "foo bar":
                                            switch (fi14 & int(v13)) {
                                            case 1:
                                                var v17 uint8 = (v13 - v13)
                                                _ = v17
                                            default:
                                                var v18 uint64 = 42
                                                _ = v18
                                            }
                                        case "\n\t":
                                            if false {
                                                fmt.Println(v3, v13, fi14)
                                            }
                                        }
                                        v19 := v1[(v2 + v2)]
                                        _ = fi16
                                        _ = v19
                                    }
                                    var v20 float64 = 3.14159
                                    v21 := []int{32768, fi14, fi14}
                                    _ = v20
                                    _ = v21
                                } else {
                                    var v22 string = (("" + (v2 + "foo bar")) + v2)
                                    _ = v22
                                }
                                var v23 string = "foo bar"
                                _ = v23
                            } else {
                                v24 := GenS2{F0: 458, F1: 1}
                                _ = v24
                            }
                        case "\n\t":
                            fmt.Println(v2)
                        case "foo bar":
                            v15 -= 42
                        }
                    } else {
                        v1[""] = v2
                    }
                }
                _ = fi14
                _ = v15
            }
            var v25 uint16 = 32767
            _ = v13
            _ = v25
        }
    case "":
        v2 = v2
    case "hello":
        switch -1 {
        case 127:
            v3 += (v3 >> ((1) & 63))
        default:
            var v26 uint64 = 0
            _ = v26
        }
    }
    switch (v2 + ((v2 + v2) + "hello")) {
    case "hello":
        if (!true) {
            if false {
                for fi27 := 0; fi27 < 10; fi27++ {
                    v2 = (("hello" + v2) + v2)
                    fmt.Println(v3, v2, v3)
                    _ = fi27
                }
                v28 := GenS3{F0: 0, F1: 2147483648}
                _ = v28
            }
        }
    }
    var v29 bool = (true && (!false))
    switch "\n\t" {
    case "a":
        var v30 int = -1
        _ = v30
    case "\n\t":
        switch v2 {
        case "\n\t":
            for fi31 := 0; fi31 < 14; fi31++ {
                switch (^(fi31)) {
                case -66:
                    v32 := []int{(fi31 << (((v3 + v3)) & 63)), genFn0(2, 42)}
                    _ = v32
                case 4294967295:
                    _ = len(v1)
                case -9223372036854775808:
                    v2 = v2
                }
                _ = fi31
            }
        case "":
            v2 = "a"
        case "a":
            switch genFn0(2147483648, 67) {
            case 2147483647:
                switch ((((v2 + v2) + v2) + "hello") + "foo bar") {
                case "foo bar":
                    for fi33 := 0; fi33 < 10; fi33++ {
                        v34 := GenS3{F0: 42, F1: 303}
                        if (76 != 255) {
                            v35 := GenS3{F0: 42, F1: 0}
                            _ = v35
                        }
                        _ = fi33
                        _ = v34
                    }
                case "hello":
                    for fi36 := 0; fi36 < 11; fi36++ {
                        for fi37 := 0; fi37 < 16; fi37++ {
                            if v29 {
                                var v38 bool = (!((443 > 65535) && v29))
                                fmt.Println(v29)
                                var v39 uint = 255
                                _ = v38
                                _ = v39
                            }
                            if (1 >= -1) {
                                if true {
                                    v40 := GenS4{F0: 42}
                                    fmt.Println(fi36, v29)
                                    if v29 {
                                        var v41 uint = v3
                                        fmt.Println(v2, v2, fi36)
                                        v2 = ""
                                        var v42 int16 = 1
                                        _ = v41
                                        _ = v42
                                    } else {
                                        var v43 uint16 = uint16(v3)
                                        if v29 {
                                            if ((0 < 127) || (v43 == (v43 >> (((v3 / ((v3) | 1))) & 63)))) {
                                                v44 := GenS1{F0: 0, F1: -1}
                                                var v45 string = "hello"
                                                var v46 uint8 = 1
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                            } else {
                                                v47 := []int{genFn0(1, 127)}
                                                if (!("a" > ("" + v2))) {
                                                    switch -942 {
                                                    case 32767:
                                                        if true {
                                                            v1[v2] = (v2 + v2)
                                                            var v48 uint64 = 1
                                                            switch (fi37 | (fi37 ^ genFn0(0, 100))) {
                                                            case 962:
                                                                if true {
                                                                    var v49 int = genFn0(636, 42)
                                                                    fmt.Println(v43, fi36, v3)
                                                                    _ = v49
                                                                }
                                                            default:
                                                                v2 = "foo bar"
                                                            }
                                                            _ = v48
                                                        }
                                                    case -698:
                                                        fmt.Println(v29, v2)
                                                    default:
                                                        var v50 int64 = 0
                                                        _ = len(v1)
                                                        _ = v50
                                                    }
                                                    if (v29 && v29) {
                                                        v43 = v43
                                                        _ = len(v1)
                                                        v3 -= v3
                                                    } else {
                                                        if (0.0 < (1.5 / -2.25)) {
                                                            v51 := v1["\n\t"]
                                                            _ = v51
                                                        } else {
                                                            var v52 int64 = -1
                                                            _ = v52
                                                        }
                                                    }
                                                    var v53 string = "foo bar"
                                                    _ = v53
                                                }
                                                fmt.Println(fi36, v2)
                                                _ = v47
                                            }
                                            var v54 float64 = (((-2.25 + -2.25) / 1.0) / 1.5)
                                            v54 += ((v54 + (((((v54 - 3.14159) / v54) * 3.14159) + v54) / (((v54 * 0.0) + v54) - (v54 * (0.0 / v54))))) * v54)
                                            _ = v54
                                        } else {
                                            fmt.Println(v43, fi37, v2)
                                            v55 := []int{128, fi37, (fi37 + fi36), (fi36 >> (((v3 | 4294967295)) & 63))}
                                            v56 := v1[("" + "")]
                                            _ = v55
                                            _ = v56
                                        }
                                        _ = v43
                                    }
                                    fmt.Println(v2)
                                    _ = v40
                                } else {
                                    switch fi37 {
                                    case 906:
                                        switch v2 {
                                        case "\n\t":
                                            switch fi37 {
                                            case 784:
                                                delete(v1, (v2 + v2))
                                            default:
                                                fmt.Println(fi36)
                                            }
                                        case "hello":
                                            v2 = (v2 + "hello")
                                        }
                                    case 1:
                                        fmt.Println(v2, v3, v2)
                                    case -1:
                                        var v57 uint = 255
                                        _ = v57
                                    default:
                                        v2 = v2
                                    }
                                }
                            }
                            v58 := GenS5{F0: 204}
                            _ = fi37
                            _ = v58
                        }
                        fmt.Println(v29, v29, v29)
                        for fi59 := 0; fi59 < 8; fi59++ {
                            v29 = (v29 && v29)
                            fmt.Println(fi59)
                            _ = fi59
                        }
                        _ = fi36
                    }
                case "":
                    delete(v1, ("" + ""))
                default:
                    v29 = true
                }
            case 180:
                for fi60 := 0; fi60 < 14; fi60++ {
                    var v61 string = ("a" + v2)
                    _ = fi60
                    _ = v61
                }
            }
        }
    case "":
        var v62 float64 = 3.14159
        _ = v62
    default:
        switch v2 {
        case "":
            fmt.Println(v3, v2, v2)
        }
    }
    v63 := GenS4{F0: 957}
    switch -2147483648 {
    case -2147483648:
        v29 = v29
    case 42:
        switch 130 {
        case -708:
            v3 = v3
        default:
            var v64 uint64 = 18446744073709551615
            _ = v64
        }
    case 32767:
        v3 = (^(v3))
    }
    v65 := []int{32768, 128}
    if len(v65) >= 2 {
        v65 = v65[1:len(v65)]
    }
    switch v2 {
    case "a":
        switch 0 {
        case 32768:
            for fi66 := 0; fi66 < 18; fi66++ {
                v67 := GenS6{F0: (fi66 << ((867) & 63)), F1: 1, F2: 65535}
                for fi68 := 0; fi68 < 12; fi68++ {
                    _ = len(v1)
                    fmt.Println(v3, v3)
                    _ = fi68
                }
                fmt.Println(v29, v3)
                _ = fi66
                _ = v67
            }
        case -516:
            delete(v1, "foo bar")
        }
    case "":
        var v69 uint16 = 831
        _ = v69
    case "\n\t":
        v70 := GenS4{F0: 32767}
        _ = v70
    }
    v71 := GenS7{F0: 1, F1: 1}
    fmt.Println(v29, v3, v2)
    for fi72 := 0; fi72 < 18; fi72++ {
        v2 = "a"
        v29 = v29
        _ = fi72
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v65:", len(v65), v65)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v63
    _ = v65
    _ = v71
}
