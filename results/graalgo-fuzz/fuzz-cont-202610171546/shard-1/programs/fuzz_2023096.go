package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS3 struct { F0 int64; F1 string; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 int64; F2 uint32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 bool; F1 int; F2 uint16 }

func main() {
    v1 := []uint32{143}
    v2 := make(map[int]int)
    v3 := []int{127, 62, -9223372036854775808, 318}
    switch "hello" {
    case "hello":
        v4 := GenS1{F0: false}
        _ = v4
    case "a":
        var v5 uint8 = 1
        _ = v5
    case "":
        for fi6 := 0; fi6 < 9; fi6++ {
            if len(v1) > 0 {
                v1[0] = 1
            }
            var v7 float32 = (1e10 * 1.5)
            _ = fi6
            _ = v7
        }
    }
    v8 := v2[-1]
    switch (v8 >> ((4294967295) & 63)) {
    case 128:
        var v9 int32 = -579
        _ = v9
    case -190:
        var v10 uint64 = 18446744073709551615
        _ = v10
    default:
        v8 -= (v8 * (v8 + v8))
    }
    _ = len(v3)
    if true {
        for fi11 := 0; fi11 < 3; fi11++ {
            fmt.Println(v8, v8)
            if false {
                _ = len(v3)
                _ = len(v1)
                switch "" {
                case "":
                    v8 -= fi11
                case "\n\t":
                    if (206 <= 118) {
                        switch "foo bar" {
                        case "a":
                            for fi12 := 0; fi12 < 18; fi12++ {
                                fmt.Println(fi11, v8, fi12)
                                v1 = append(v1, 2147483647)
                                var v13 string = "foo bar"
                                if len(v3) > 0 {
                                    v3[0] = (v8 | (fi11 >> ((1) & 63)))
                                }
                                _ = fi12
                                _ = v13
                            }
                        case "":
                            if false {
                                var v14 uint16 = 827
                                v14 = (v14 ^ 0)
                                delete(v2, (fi11 / (((fi11 ^ (fi11 ^ v8))) | 1)))
                                _ = v14
                            } else {
                                fmt.Println(v8, v8)
                                for fi15 := 0; fi15 < 11; fi15++ {
                                    v16 := GenS2{F0: 42, F1: 42, F2: -1}
                                    fmt.Println(fi11)
                                    _ = fi15
                                    _ = v16
                                }
                            }
                        case "foo bar":
                            var v17 uint16 = 0
                            _ = v17
                        default:
                            v8 = v8
                        }
                        if false {
                            v8 = (v8 - -752)
                            var v18 float64 = 3.14159
                            _ = v18
                        } else {
                            if (704 != 18446744073709551615) {
                                var v19 float32 = ((-2.25 / 0.1) / (3.14159 + 3.14159))
                                if len(v1) > 0 {
                                    v1[0] = 65535
                                }
                                var v20 float64 = 1.5
                                _ = v19
                                _ = v20
                            }
                            if (false && true) {
                                v8 = 9223372036854775807
                                v21 := GenS3{F0: -9223372036854775808, F1: "hello", F2: 128}
                                _ = len(v3)
                                _ = v21
                            }
                        }
                    }
                }
            }
            if (4294967295 != 18446744073709551615) {
                v8 = v8
            }
            _ = fi11
        }
        if (((!true) && (42 <= 1)) || true) {
            v22 := GenS3{F0: 42, F1: "foo bar", F2: 0}
            for fi23 := 0; fi23 < 9; fi23++ {
                v24 := GenS4{F0: 0}
                _ = fi23
                _ = v24
            }
            _ = v22
        } else {
            v8 = -326
        }
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        v8 -= 4294967295
        v26 := v2[fi25]
        _ = fi25
        _ = v26
    }
    if true {
        var v27 float32 = 0.0
        v27 += (float32(((0.1 * -2.25) / 1.5)) / 0.1)
        if (true || false) {
            v1 = append(v1, 2)
            v8 = (-(v8))
            v28 := GenS5{F0: 224, F1: 255, F2: 42}
            _ = v28
        } else {
            switch "a" {
            case "a":
                v27 -= 0.0
            case "hello":
                var v29 int16 = 0
                _ = v29
            case "foo bar":
                fmt.Println(v8, v8)
            default:
                if ((!(127 >= 1)) || (1 <= 1)) {
                    v1 = append(v1, 56)
                } else {
                    v30 := GenS3{F0: 9223372036854775807, F1: "hello", F2: 2}
                    _ = v30
                }
            }
            v27 = v27
            if false {
                for fi31 := 0; fi31 < 19; fi31++ {
                    var v32 uint16 = 0
                    v8 -= 4294967295
                    _ = fi31
                    _ = v32
                }
                v33 := GenS5{F0: 632, F1: 18446744073709551615, F2: 1}
                _ = v33
            }
        }
        _ = v27
    } else {
        fmt.Println(v8)
        if false {
            v34 := GenS5{F0: 18446744073709551615, F1: 255, F2: 65535}
            _ = v34
        } else {
            v35 := GenS2{F0: 2147483647, F1: 0, F2: 580}
            for fi36 := 0; fi36 < 18; fi36++ {
                var v37 uint = 18446744073709551615
                v8 = fi36
                _ = fi36
                _ = v37
            }
            _ = v35
        }
        for fi38 := 0; fi38 < 19; fi38++ {
            fmt.Println(v8, fi38)
            fmt.Println(fi38, v8)
            _ = fi38
        }
    }
    switch 128 {
    case -1:
        switch (("foo bar" + "\n\t") + "\n\t") {
        case "foo bar":
            fmt.Println(v8)
        case "a":
            for fi39 := 0; fi39 < 16; fi39++ {
                var v40 string = ""
                _ = fi39
                _ = v40
            }
        case "":
            var v41 int = -9223372036854775808
            _ = v41
        default:
            v8 = (v8 * (v8 ^ v8))
        }
    case 42:
        var v42 float32 = 3.14159
        _ = v42
    case 327:
        for fi43 := 0; fi43 < 17; fi43++ {
            v44 := GenS6{F0: -114}
            switch 4294967295 {
            case 140:
                v45 := GenS4{F0: -2147483648}
                _ = v45
            case -654:
                _ = len(v2)
            case 4294967295:
                v8 = v8
            default:
                for fi46 := 0; fi46 < 4; fi46++ {
                    if (!false) {
                        v8 = 604
                    } else {
                        v47 := GenS6{F0: 127}
                        v8 -= fi43
                        switch (fi46 + fi46) {
                        case -752:
                            v2[fi46] = fi43
                        }
                        _ = v47
                    }
                    fmt.Println(fi46, v8, fi46)
                    _ = fi46
                }
            }
            _ = fi43
            _ = v44
        }
    default:
        fmt.Println(v8)
    }
    var v48 uint = 450
    var v49 string = ("foo bar" + ((((("foo bar" + "\n\t") + "hello") + (((("hello" + "foo bar") + "") + "a") + "a")) + "") + "hello"))
    if (false || (!false)) {
        var v50 uint16 = 50
        v2[v8] = v8
        v51 := GenS4{F0: 42}
        _ = v50
        _ = v51
    } else {
        var v52 bool = (42 >= 1)
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch ("foo bar" + v49) {
            case "a":
                var v53 uint64 = 940
                _ = v53
            case "\n\t":
                var v54 uint16 = 615
                _ = v54
            case "hello":
                v48 = v48
            }
        } else {
            var v55 float64 = (((float64(1e10) / (-2.25 - 3.14159)) - 0.0) + 0.0)
            v8 += v8
            _ = v55
        }
        _ = v52
    }
    for fi56 := 0; fi56 < 14; fi56++ {
        v8 -= -2147483648
        for fi57 := 0; fi57 < 19; fi57++ {
            v48 = v48
            v58 := v2[-47]
            var v59 uint32 = 4294967295
            _ = fi57
            _ = v58
            _ = v59
        }
        for fi60 := 0; fi60 < 6; fi60++ {
            if false {
                switch v8 {
                case 42:
                    var v61 uint8 = 90
                    _ = v61
                case 9223372036854775807:
                    switch fi60 {
                    case 0:
                        v2[813] = (fi56 * 0)
                    case 591:
                        var v62 float64 = 0.0
                        _ = v62
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                case 32767:
                    if ((true || (v49 < v49)) && false) {
                        v63 := GenS6{F0: 42}
                        v64 := GenS7{F0: fi56, F1: 0, F2: 2}
                        _ = v63
                        _ = v64
                    } else {
                        v65 := GenS6{F0: 0}
                        v66 := GenS2{F0: 42, F1: 42, F2: -1}
                        _ = v65
                        _ = v66
                    }
                default:
                    delete(v2, (v8 - v8))
                }
                if true {
                    switch "hello" {
                    case "":
                        if (int16(v48) <= -853) {
                            v67 := GenS7{F0: v8, F1: 42, F2: 358}
                            _ = v67
                        } else {
                            v8 = v8
                        }
                    default:
                        v68 := GenS2{F0: 2, F1: 18446744073709551615, F2: -119}
                        if (v48 <= (v48 + (v48 ^ v48))) {
                            v49 = ((v49 + "a") + "foo bar")
                            fmt.Println(fi60, v48, fi56)
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                        }
                        if (true && true) {
                            switch fi60 {
                            case 2147483647:
                                if (false || (0 <= 42)) {
                                    v48 = (v48 + (^(v48)))
                                } else {
                                    v8 = fi60
                                }
                            default:
                                v69 := GenS8{F0: 1}
                                _ = v69
                            }
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                            v49 = (((v49 + v49) + ((v49 + (v49 + "foo bar")) + v49)) + "hello")
                        }
                        switch (v8 & fi60) {
                        case 411:
                            v49 = v49
                        default:
                            v48 = v48
                            fmt.Println(fi56, fi56)
                        }
                        _ = v68
                    }
                    fmt.Println(fi60, v8)
                } else {
                    fmt.Println(fi56, fi56, v8)
                }
                v70 := GenS9{F0: true, F1: 1, F2: 1}
                var v71 string = "hello"
                _ = v70
                _ = v71
            }
            _ = fi60
        }
        _ = fi56
    }
    v72 := GenS9{F0: ((((false || (1000 == 1000)) && (v48 >= v48)) || (!(74 >= 127))) || false), F1: v8, F2: 1}
    v49 = (v49 + v49)
    if (float64((float32(1e10) / (3.14159 - -2.25))) <= 1.5) {
        if true {
            _ = len(v2)
        }
    } else {
        v1 = append(v1, 2147483648)
    }
    fmt.Println(v48, v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v72
    _ = v8
}
