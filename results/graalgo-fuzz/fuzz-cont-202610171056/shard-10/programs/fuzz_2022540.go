package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int8; F2 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 float64 = 1e10
    var v2 uint8 = 128
    v3 := []string{"foo bar", "hello", "a"}
    v4 := make(map[string]bool)
    var v5 bool = (float64(3.14159) < v1)
    var v6 string = ("\n\t" + ("a" + ((("a" + "\n\t") + "") + "hello")))
    switch 1 {
    case 580:
        v5 = v5
    case 9223372036854775807:
        v6 = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int64 = 42
        _ = len(v3)
        _ = fi7
        _ = v8
    }
    v2 = (v2 + (v2 - (v2 << ((42) & 63))))
    if v5 {
        switch 0 {
        case 127:
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 int8 = 100
                switch fi9 {
                case 9223372036854775807:
                    if len(v3) > 0 {
                        v3[0] = (v6 + (("\n\t" + "a") + "foo bar"))
                    }
                case 128:
                    var v11 bool = (902 >= 42)
                    _ = v11
                case 321:
                    for fi12 := 0; fi12 < 4; fi12++ {
                        var v13 float32 = (-2.25 / 3.14159)
                        v5 = (42 > 65535)
                        _ = fi12
                        _ = v13
                    }
                }
                for fi14 := 0; fi14 < 19; fi14++ {
                    fmt.Println(v6)
                    switch "\n\t" {
                    case "foo bar":
                        v5 = (false && true)
                        if (v5 && (42 <= 42)) {
                            var v15 int8 = v10
                            var v16 bool = v5
                            _ = v15
                            _ = v16
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        if false {
                            v17 := GenS1{F0: (v2 / ((43) | 1)), F1: -32768, F2: fi14}
                            var v18 int = -2147483648
                            if ((^(v2)) < (^(v2))) {
                                _ = len(v3)
                            }
                            _ = v17
                            _ = v18
                        }
                    }
                    _ = fi14
                }
                _ = fi9
                _ = v10
            }
        case -75:
            if v5 {
                fmt.Println(v1, v1)
                v1 += (((v1 * (0.1 / 1.0)) / v1) + v1)
            } else {
                v19 := GenS2{F0: -2147483648}
                fmt.Println(v5, v6, v1)
                _ = v19
            }
        }
    } else {
        switch (v6 + (v6 + v6)) {
        case "hello":
            for fi20 := 0; fi20 < 17; fi20++ {
                var v21 int16 = -1
                _ = fi20
                _ = v21
            }
        case "":
            v22 := GenS2{F0: 9223372036854775807}
            _ = v22
        case "\n\t":
            v5 = ((1.5 < v1) && (331 != 255))
        }
        v5 = (!v5)
    }
    if v5 {
        v5 = (!v5)
    } else {
        fmt.Println(v1, v6, v2)
        var v23 uint64 = 2
        v6 = v6
        _ = v23
    }
    v24 := GenS3{F0: 42, F1: 0, F2: 18446744073709551615}
    if v5 {
        switch 42 {
        case 42:
            for fi25 := 0; fi25 < 19; fi25++ {
                var v26 uint16 = 0
                if v5 {
                    var v27 int64 = -1
                    v3 = append(v3, v6)
                    for fi28 := 0; fi28 < 2; fi28++ {
                        v26 -= 1
                        _ = fi28
                    }
                    _ = v27
                }
                _ = fi25
                _ = v26
            }
        case 4294967295:
            v3 = append(v3, ((v6 + "\n\t") + ("hello" + (v6 + v6))))
        default:
            for fi29 := 0; fi29 < 11; fi29++ {
                v2 = (v2 & v2)
                v1 = v1
                _ = fi29
            }
        }
        v2 = (v2 ^ v2)
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 uint = 18446744073709551615
            v32 := v4[("hello" + ((((v6 + ((v6 + ((((((("\n\t" + "\n\t") + "foo bar") + v6) + (v6 + (v6 + "a"))) + ("" + (v6 + "a"))) + v6) + ("hello" + "hello"))) + ("hello" + ("a" + v6)))) + v6) + ((("hello" + v6) + "") + "\n\t")) + v6))]
            var v33 float64 = 3.14159
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
    } else {
        if ((("" <= (v6 + v6)) || (!v5)) && (65535 >= 42)) {
            v5 = v5
        }
        if v5 {
            delete(v4, v6)
        } else {
            switch v6 {
            case "a":
                switch -235 {
                case 32768:
                    v2 += v2
                case 127:
                    v6 = v6
                case -1:
                    var v34 float64 = v1
                    _ = v34
                }
            case "hello":
                var v35 string = ((v6 + ((("hello" + v6) + "\n\t") + v6)) + ("" + v6))
                _ = v35
            case "":
                v5 = (((v6 + "\n\t") + v6) <= v6)
            }
        }
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        switch (fi36 >> ((835) & 63)) {
        case 2147483647:
            for fi37 := 0; fi37 < 15; fi37++ {
                v6 = v6
                _ = fi37
            }
        case 32767:
            switch (-(fi36)) {
            case -2147483648:
                for fi38 := 0; fi38 < 5; fi38++ {
                    v1 = (float64((-2.25 / (1e10 - -2.25))) * (((1e10 - float64((-2.25 * -2.25))) * v1) / v1))
                    v6 = ("" + v6)
                    _ = fi38
                }
            case 9223372036854775807:
                var v39 int32 = 0
                _ = v39
            case 127:
                for fi40 := 0; fi40 < 7; fi40++ {
                    v5 = ((v6 + "foo bar") != (("hello" + (v6 + v6)) + v6))
                    if (v5 && (v5 || ((((v1 / ((v1 * -2.25) - (v1 - v1))) / v1) + ((v1 + -2.25) / 1.0)) >= 0.1))) {
                        var v41 int8 = 100
                        if ((fi40 << ((4294967295) & 63)) <= fi36) {
                            var v42 float64 = float64(1.5)
                            var v43 uint64 = 4294967295
                            if v5 {
                                var v44 uint8 = v2
                                var v45 int8 = (v41 ^ v41)
                                v46 := v4["a"]
                                _ = v44
                                _ = v45
                                _ = v46
                            } else {
                                var v47 int = (fi36 ^ fi36)
                                v48 := GenS4{F0: 0}
                                switch 9223372036854775807 {
                                case 1:
                                    delete(v4, "hello")
                                    v2 += v2
                                case -2147483648:
                                    v3 = append(v3, v6)
                                default:
                                    var v49 int64 = 1
                                    _ = v49
                                }
                                _ = v47
                                _ = v48
                            }
                            _ = v42
                            _ = v43
                        } else {
                            v5 = (!v5)
                            v1 = v1
                            var v50 int32 = 2147483647
                            _ = v50
                        }
                        _ = v41
                    }
                    fmt.Println(fi40)
                    var v51 int32 = 0
                    _ = fi40
                    _ = v51
                }
            }
        case 1:
            v52 := GenS4{F0: 730}
            _ = v52
        }
        v53 := GenS5{F0: fi36}
        _ = fi36
        _ = v53
    }
    var v54 float64 = (-2.25 * float64(1e10))
    if v5 {
        switch 680 {
        case 128:
            if ((!(2147483647 >= 2)) && v5) {
                switch 128 {
                case 9223372036854775807:
                    v6 = v6
                case 345:
                    v54 = 3.14159
                }
                v5 = (2147483647 > 42)
            } else {
                switch 32768 {
                case -1:
                    var v55 uint32 = 65535
                    _ = v55
                case 42:
                    v54 -= v1
                case 127:
                    for fi56 := 0; fi56 < 19; fi56++ {
                        v54 -= v1
                        if true {
                            var v57 int8 = 32
                            v58 := GenS6{F0: v57, F1: 42}
                            _ = v57
                            _ = v58
                        } else {
                            delete(v4, v6)
                            v59 := GenS1{F0: v2, F1: 885, F2: (fi56 + fi56)}
                            if false {
                                v60 := GenS7{F0: v2, F1: (v2 + v2)}
                                v2 += (v2 >> ((2) & 63))
                                _ = v60
                            }
                            _ = v59
                        }
                        if v5 {
                            delete(v4, "\n\t")
                            v1 += v1
                        } else {
                            var v61 int8 = -105
                            v62 := GenS2{F0: fi56}
                            v54 -= (1.5 / v1)
                            _ = v61
                            _ = v62
                        }
                        _ = fi56
                    }
                }
            }
        case 32768:
            delete(v4, ("a" + "foo bar"))
        case -232:
            if len(v3) > 0 {
                v3[0] = v6
            }
        default:
            var v63 int8 = -80
            _ = v63
        }
    } else {
        v2 = (v2 + 181)
        fmt.Println(v54)
        for fi64 := 0; fi64 < 16; fi64++ {
            v3 = append(v3, "foo bar")
            if (!(1.5 < -2.25)) {
                v54 = (0.0 + (v54 / ((1.5 + float64(float32(1e10))) - (v1 / v54))))
                for fi65 := 0; fi65 < 9; fi65++ {
                    v1 = v54
                    v66 := GenS3{F0: 0, F1: -31, F2: 42}
                    _ = fi65
                    _ = v66
                }
            }
            _ = fi64
        }
    }
    for fi67 := 0; fi67 < 8; fi67++ {
        var v68 float64 = (0.1 / (((0.1 / ((v1 / 3.14159) + v1)) - v54) / 1.0))
        _ = fi67
        _ = v68
    }
    if true {
        v5 = true
        if (false && (true && (v5 || true))) {
            v54 += ((((v54 * v1) / ((1.5 * (v54 - v1)) / 1.0)) + 0.0) * 3.14159)
            if v5 {
                var v69 uint64 = 18446744073709551615
                var v70 float64 = 1e10
                _ = v69
                _ = v70
            }
        }
        v71 := GenS4{F0: -915}
        _ = v71
    } else {
        v54 = 0.0
        if (!(((v2 + v2) <= (^(v2))) && false)) {
            _ = len(v3)
        } else {
            if (-1 >= 775) {
                var v72 float64 = (0.0 + v1)
                var v73 bool = ((!v5) || v5)
                _ = v72
                _ = v73
            } else {
                for fi74 := 0; fi74 < 11; fi74++ {
                    for fi75 := 0; fi75 < 18; fi75++ {
                        v76 := GenS3{F0: 1000, F1: -78, F2: 163}
                        v1 = ((float64(float32(3.14159)) * -2.25) / 3.14159)
                        v6 = "a"
                        switch 913 {
                        case 32768:
                            switch fi74 {
                            case -1:
                                var v77 int8 = -1
                                _ = v77
                            case 786:
                                delete(v4, v6)
                            }
                        case -110:
                            v6 = (v6 + (v6 + v6))
                        }
                        _ = fi75
                        _ = v76
                    }
                    _ = fi74
                }
                fmt.Println(v5)
                switch 9223372036854775807 {
                case -2147483648:
                    fmt.Println(v5, v5)
                default:
                    v1 += 0.0
                }
            }
            for fi78 := 0; fi78 < 17; fi78++ {
                var v79 int = (fi78 & (-(fi78)))
                _ = fi78
                _ = v79
            }
        }
    }
    if (620 <= 2147483647) {
        var v80 bool = (v6 != v6)
        var v81 uint32 = 2147483648
        v2 += 2
        _ = v80
        _ = v81
    }
    var v82 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v82
}
