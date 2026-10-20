package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 float32 = -2.25
    v2 := []string{"foo bar", "\n\t", ("a" + (("foo bar" + "a") + "\n\t")), (("a" + ("hello" + "a")) + "\n\t")}
    v3 := []int{0}
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 3; fi6++ {
            if false {
                var v7 float32 = 1.5
                fmt.Println(fi6, fi6, fi6)
                _ = v7
            } else {
                var v8 uint64 = 954
                v8 -= 0
                delete(v4, 0)
                _ = v8
            }
            v9 := GenS1{F0: 18446744073709551615, F1: 32767}
            v1 = 1e10
            _ = fi6
            _ = v9
        }
        v10 := GenS1{F0: 42, F1: -32768}
        _ = fi5
        _ = v10
    }
    fmt.Println(v1)
    var v11 uint8 = 1
    v1 += v1
    switch -734 {
    case -282:
        v1 = (v1 + 0.1)
    }
    fmt.Println(v11, v11)
    switch 127 {
    case 127:
        v12 := GenS1{F0: 42, F1: 32767}
        _ = v12
    case 9223372036854775807:
        v3 = append(v3, 4294967295)
    default:
        v11 = (v11 & v11)
    }
    var v13 int = 833
    if (112 >= 1) {
        v14 := GenS1{F0: 292, F1: 885}
        for fi15 := 0; fi15 < 5; fi15++ {
            v4[42] = true
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 uint32 = 42
                _ = fi16
                _ = v17
            }
            if (!((v11 + v11) <= 255)) {
                var v18 int64 = 705
                var v19 uint16 = 691
                switch "hello" {
                case "":
                    v20 := GenS1{F0: 59, F1: -1}
                    _ = v20
                }
                _ = v18
                _ = v19
            } else {
                _ = len(v4)
            }
            _ = fi15
        }
        v13 -= -2147483648
        _ = v14
    } else {
        var v21 int16 = 1
        for fi22 := 0; fi22 < 1; fi22++ {
            switch (v13 - (v13 + (v13 ^ v13))) {
            case 42:
                if (!true) {
                    v23 := v4[0]
                    var v24 int64 = -9223372036854775808
                    if (!(1e10 == 0.0)) {
                        if v23 {
                            fmt.Println(v21, v11, v24)
                            if ((v11 - (v11 * v11)) > (v11 << ((2) & 63))) {
                                for fi25 := 0; fi25 < 13; fi25++ {
                                    if v23 {
                                        if (v21 != (^(v21))) {
                                            v26 := v4[4294967295]
                                            fmt.Println(v11, v24)
                                            if false {
                                                fmt.Println(v21, fi25, v23)
                                                v27 := GenS1{F0: 953, F1: (v21 | v21)}
                                                _ = v27
                                            } else {
                                                fmt.Println(fi25, v1)
                                                v21 -= v21
                                            }
                                            _ = v26
                                        }
                                        switch "\n\t" {
                                        case "a":
                                            if (("hello" + ("a" + ((("" + (("\n\t" + (("hello" + ("hello" + ("\n\t" + ("a" + ((("a" + ("\n\t" + "hello")) + ("hello" + "hello")) + "a"))))) + "\n\t")) + "foo bar")) + "\n\t") + "\n\t"))) >= "a") {
                                                v11 -= v11
                                            }
                                        case "\n\t":
                                            switch "a" {
                                            case "hello":
                                                v4[2147483647] = (!true)
                                            case "":
                                                v11 = (v11 >> ((255) & 63))
                                            case "a":
                                                var v28 int16 = -615
                                                _ = v28
                                            }
                                            var v29 float64 = 1.5
                                            v30 := GenS1{F0: 1, F1: v21}
                                            _ = v29
                                            _ = v30
                                        case "hello":
                                            var v31 int8 = -128
                                            _ = v31
                                        }
                                    }
                                    _ = fi25
                                }
                                v4[343] = true
                                fmt.Println(v1, v21, v21)
                            }
                        }
                        v32 := GenS2{F0: 1}
                        _ = v32
                    } else {
                        var v33 uint16 = 2
                        if len(v2) > 0 {
                            v2[0] = (("foo bar" + "\n\t") + ("" + "foo bar"))
                        }
                        _ = v33
                    }
                    _ = v23
                    _ = v24
                } else {
                    if (!("" >= "foo bar")) {
                        var v34 float32 = v1
                        _ = v34
                    }
                    switch (fi22 / ((fi22) | 1)) {
                    case 4294967295:
                        var v35 int = v13
                        _ = v35
                    case 32768:
                        if ((fi22 + fi22) <= -1) {
                            switch (-(fi22)) {
                            case 9223372036854775807:
                                if (false && (!(65535 <= 151))) {
                                    switch ("a" + "") {
                                    case "":
                                        for fi36 := 0; fi36 < 14; fi36++ {
                                            v37 := GenS2{F0: 998}
                                            var v38 int = 2147483647
                                            if false {
                                                switch ("a" + ("foo bar" + "hello")) {
                                                case "hello":
                                                    switch "foo bar" {
                                                    case "a":
                                                        var v39 float32 = v1
                                                        _ = v39
                                                    }
                                                default:
                                                    v40 := GenS3{F0: v38, F1: 1}
                                                    _ = v40
                                                }
                                                var v41 uint8 = v11
                                                switch ("\n\t" + "\n\t") {
                                                case "a":
                                                    v13 = -2147483648
                                                case "":
                                                    if (1.5 != float32(float64((0.1 / (0.1 + (((v1 + v1) * float32((-2.25 + 0.0))) - 0.0)))))) {
                                                        var v42 int64 = -678
                                                        _ = v42
                                                    }
                                                }
                                                _ = v41
                                            }
                                            var v43 float32 = (((v1 / 1e10) + v1) + ((v1 / float32((float64((v1 + v1)) / (0.1 - -2.25)))) - (v1 - float32(0.0))))
                                            var v44 int = (v13 >> ((255) & 63))
                                            fmt.Println(fi22, v13, v43)
                                            _ = fi36
                                            _ = v37
                                            _ = v38
                                            _ = v43
                                            _ = v44
                                        }
                                    }
                                    v4[156] = false
                                } else {
                                    for fi45 := 0; fi45 < 9; fi45++ {
                                        var v46 float32 = v1
                                        _ = fi45
                                        _ = v46
                                    }
                                    _ = len(v4)
                                }
                            case 42:
                                v11 += v11
                            }
                        }
                    case -1:
                        v47 := GenS4{F0: v11}
                        _ = v47
                    }
                    var v48 int16 = (v21 ^ (v21 % (((v21 % ((v21) | 1))) | 1)))
                    _ = v48
                }
            }
            var v49 bool = (v21 == v21)
            for fi50 := 0; fi50 < 20; fi50++ {
                if v49 {
                    fmt.Println(v49, v1, v13)
                    v3 = append(v3, (v13 % (((fi50 << ((305) & 63))) | 1)))
                    var v51 uint64 = 387
                    _ = v51
                }
                fmt.Println(v13, v49)
                _ = fi50
            }
            _ = fi22
            _ = v49
        }
        v1 = float32(0.0)
        _ = v21
    }
    switch 42 {
    case 128:
        for fi52 := 0; fi52 < 4; fi52++ {
            v53 := v4[0]
            fmt.Println(fi52, v53)
            v54 := GenS1{F0: 18446744073709551615, F1: -1}
            _ = fi52
            _ = v53
            _ = v54
        }
    case 798:
        for fi55 := 0; fi55 < 15; fi55++ {
            for fi56 := 0; fi56 < 20; fi56++ {
                switch fi55 {
                case 798:
                    if true {
                        var v57 uint = 861
                        fmt.Println(v1, fi55)
                        fmt.Println(v13, v13)
                        _ = v57
                    }
                default:
                    var v58 int = (fi55 - (fi56 % ((9223372036854775807) | 1)))
                    _ = v58
                }
                if false {
                    v59 := GenS4{F0: (v11 + (^(v11)))}
                    var v60 uint32 = 582
                    if (!((!(!true)) && false)) {
                        var v61 float64 = float64((((((1.5 + -2.25) + v1) / (float32(1e10) / (v1 - 1.5))) * v1) / v1))
                        _ = v61
                    }
                    v13 = fi55
                    var v62 uint = 937
                    var v63 uint16 = uint16(v11)
                    _ = v59
                    _ = v60
                    _ = v62
                    _ = v63
                } else {
                    var v64 bool = false
                    fmt.Println(v1, v64)
                    v64 = v64
                    var v65 int = -9223372036854775808
                    _ = v64
                    _ = v65
                }
                v1 = -2.25
                _ = fi56
            }
            _ = fi55
        }
    case 300:
        if false {
            switch (v13 ^ (v13 >> ((1) & 63))) {
            case 224:
                if true {
                    v4[2147483647] = true
                }
            default:
                switch "\n\t" {
                case "hello":
                    v4[1] = true
                case "":
                    var v66 bool = (!(1 > 0))
                    _ = v66
                default:
                    var v67 int = v13
                    _ = v67
                }
            }
        } else {
            var v68 uint16 = 1
            _ = v68
        }
    }
    for fi69 := 0; fi69 < 14; fi69++ {
        _ = len(v2)
        if len(v3) > 0 {
            v3[0] = (fi69 & -121)
        }
        if true {
            v11 = 255
        } else {
            _ = len(v3)
            var v70 int = fi69
            _ = v70
        }
        _ = fi69
    }
    v71 := GenS2{F0: 133}
    switch "a" {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v71
}
