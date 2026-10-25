package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 uint8; F2 bool }
type GenS4 struct { F0 string; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16; F1 string; F2 string }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 uint64; F1 uint8 }
type GenS9 struct { F0 uint32; F1 uint8; F2 bool }
type GenS10 struct { F0 uint32; F1 int; F2 int16 }
type GenS11 struct { F0 uint16 }

func main() {
    var v1 bool = (127 == 135)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 int16 = 1000
    var v5 float32 = 0.1
    var v6 uint64 = 0
    v1 = v1
    v1 = v1
    v7 := []int{}
    v8 := GenS1{F0: (v6 ^ (v6 % ((v6) | 1))), F1: -12}
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint16 = 32767
        delete(v2, 2147483647)
        for fi11 := 0; fi11 < 4; fi11++ {
            if (42 <= -2147483648) {
                var v12 int64 = 1
                v6 = v6
                v13 := GenS1{F0: (v6 | 42), F1: (^(fi11))}
                _ = v12
                _ = v13
            }
            if v1 {
                if true {
                    v14 := GenS2{F0: -79}
                    var v15 int64 = -9223372036854775808
                    if false {
                        if v1 {
                            if (((((3.14159 + float64(v5)) * 0.0) / (1e10 / float64(v5))) + (0.1 * -2.25)) <= (((1.5 * float64((v5 + v5))) + (-2.25 - ((0.0 / 1e10) + 0.0))) + 0.0)) {
                                fmt.Println(v1)
                                v16 := GenS1{F0: 2, F1: (fi9 * -1)}
                                _ = v16
                            }
                            var v17 int8 = 125
                            var v18 uint = 68
                            _ = v17
                            _ = v18
                        } else {
                            if v1 {
                                v19 := GenS3{F0: 9223372036854775807, F1: 2, F2: (42 > (^(v15)))}
                                v15 += v15
                                _ = len(v2)
                                switch fi11 {
                                case 32768:
                                    fmt.Println(v15, v1, v1)
                                    v20 := GenS2{F0: (v15 % ((v15) | 1))}
                                    _ = v20
                                case 127:
                                    var v21 uint8 = 255
                                    _ = v21
                                case 2147483647:
                                    v22 := GenS4{F0: "hello", F1: -1, F2: v6}
                                    v23 := GenS5{F0: v6}
                                    _ = v22
                                    _ = v23
                                }
                                fmt.Println(v5, v6)
                                _ = v19
                            }
                            var v24 int = (fi11 | 9223372036854775807)
                            v7 = append(v7, 967)
                            v10 = v10
                            v15 -= (v15 * -216)
                            _ = v24
                        }
                        var v25 uint16 = (v10 ^ 255)
                        fmt.Println(fi9)
                        _ = v25
                    } else {
                        if (!true) {
                            if false {
                                fmt.Println(v6)
                                var v26 uint64 = (v6 * v6)
                                _ = v26
                            } else {
                                var v27 uint = 2
                                _ = v27
                            }
                            if false {
                                var v28 int64 = v15
                                fmt.Println(v4)
                                var v29 int8 = 10
                                _ = v28
                                _ = v29
                            }
                        } else {
                            var v30 float64 = 1.5
                            _ = v30
                        }
                        v3[(^(fi9))] = 2
                    }
                    if v1 {
                        if v1 {
                            var v31 int8 = 8
                            var v32 uint64 = v6
                            _ = v31
                            _ = v32
                        } else {
                            v33 := GenS6{F0: v4, F1: "a", F2: "hello"}
                            if ((!v1) && v1) {
                                if v1 {
                                    fmt.Println(v1)
                                }
                                v10 = v10
                                delete(v3, (fi11 | 9223372036854775807))
                                var v34 float32 = v5
                                fmt.Println(v34)
                                delete(v2, 878)
                                v35 := GenS7{F0: 1, F1: (25 == (v15 * v15))}
                                _ = v34
                                _ = v35
                            }
                            v4 -= v4
                            _ = v33
                        }
                    } else {
                        if (v15 <= v15) {
                            delete(v2, 2)
                        }
                        if (!((!true) && v1)) {
                            var v36 int32 = -1
                            _ = len(v7)
                            _ = v36
                        }
                    }
                    _ = v14
                    _ = v15
                }
                var v37 uint16 = v10
                _ = v37
            }
            v38 := GenS7{F0: 2147483648, F1: (-620 >= 1)}
            _ = fi11
            _ = v38
        }
        _ = fi9
        _ = v10
    }
    var v39 float32 = (1e10 * v5)
    v40 := GenS3{F0: int(v6), F1: 255, F2: (255 <= 0)}
    if (1 >= 1) {
        _ = len(v7)
    } else {
        switch ("" + "hello") {
        case "a":
            switch -202 {
            case 2147483647:
                if (-568 != 244) {
                    fmt.Println(v6, v5)
                    var v41 float64 = (float64(1e10) + 1.5)
                    _ = v41
                }
            case 1:
                v42 := GenS8{F0: v6, F1: 2}
                _ = v42
            default:
                delete(v3, 1)
            }
        case "foo bar":
            switch 75 {
            case -1:
                var v43 string = "a"
                _ = v43
            case 32768:
                v1 = (!false)
            case 32767:
                var v44 bool = (true || (!(42 > 303)))
                _ = v44
            default:
                for fi45 := 0; fi45 < 20; fi45++ {
                    if true {
                        if (!("a" < "hello")) {
                            v5 += 3.14159
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            if false {
                                fmt.Println(v6, v6)
                                if v1 {
                                    for fi46 := 0; fi46 < 13; fi46++ {
                                        var v47 int16 = (-(v4))
                                        v2[2] = 497
                                        _ = fi46
                                        _ = v47
                                    }
                                    v5 = -2.25
                                    if false {
                                        var v48 float64 = (-2.25 / 3.14159)
                                        var v49 int64 = 595
                                        _ = v48
                                        _ = v49
                                    }
                                }
                            }
                        } else {
                            v50 := GenS3{F0: 4294967295, F1: 201, F2: (((0 != 18446744073709551615) || ((390 != 1) || v1)) && false)}
                            v6 = v6
                            _ = v50
                        }
                        if v1 {
                            var v51 bool = (!(2 < 2))
                            for fi52 := 0; fi52 < 9; fi52++ {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                switch (fi52 >> ((2) & 63)) {
                                case 774:
                                    switch (fi45 << ((724) & 63)) {
                                    case 615:
                                        v53 := GenS9{F0: 2147483647, F1: 103, F2: true}
                                        _ = v53
                                    case -9223372036854775808:
                                        if len(v7) > 0 {
                                            v7[0] = fi45
                                        }
                                    case 32767:
                                        switch "\n\t" {
                                        case "\n\t":
                                            switch (("foo bar" + "") + ("hello" + "hello")) {
                                            case "a":
                                                if ("" == ("" + "hello")) {
                                                    var v54 string = "foo bar"
                                                    v4 -= 32767
                                                    var v55 uint = 0
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        if ((v6 % ((v6) | 1)) > 4294967295) {
                                                            var v56 float32 = 1e10
                                                            var v57 string = "\n\t"
                                                            var v58 int16 = (v4 * 1000)
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                    }
                                                }
                                            case "hello":
                                                fmt.Println(fi52, v5)
                                                v59 := v3[(fi45 - (fi52 ^ (fi45 ^ (fi52 % (((fi52 >> ((255) & 63))) | 1)))))]
                                                _ = v59
                                            case "foo bar":
                                                var v60 int32 = 42
                                                _ = v60
                                            }
                                        }
                                    default:
                                        v1 = ((0.1 / ((1.5 * 1e10) - (v39 - -2.25))) == v5)
                                        if true {
                                            switch fi52 {
                                            case 32768:
                                                var v61 uint64 = (v6 | 556)
                                                _ = v61
                                            case 128:
                                                v39 -= 0.1
                                            case 2147483647:
                                                switch "\n\t" {
                                                case "hello":
                                                    var v62 uint64 = (v6 >> ((65535) & 63))
                                                    _ = v62
                                                case "foo bar":
                                                    var v63 int32 = -2147483648
                                                    _ = v63
                                                default:
                                                    var v64 int32 = 42
                                                    switch 280 {
                                                    case 95:
                                                        var v65 bool = (v5 != v39)
                                                        _ = v65
                                                    case -988:
                                                        v51 = (-1 != (v64 * v64))
                                                    }
                                                    var v66 float64 = (1.5 / 1e10)
                                                    _ = v64
                                                    _ = v66
                                                }
                                            }
                                        } else {
                                            var v67 uint = 42
                                            v68 := GenS10{F0: 1, F1: fi52, F2: 895}
                                            fmt.Println(v5)
                                            _ = v67
                                            _ = v68
                                        }
                                    }
                                }
                                _ = fi52
                            }
                            _ = v51
                        }
                    }
                    for fi69 := 0; fi69 < 4; fi69++ {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        var v70 uint32 = 4294967295
                        v71 := GenS7{F0: uint32(v6), F1: (-128 >= 43)}
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                    _ = fi45
                }
            }
        case "\n\t":
            for fi72 := 0; fi72 < 6; fi72++ {
                var v73 int64 = 295
                v74 := GenS11{F0: 902}
                _ = fi72
                _ = v73
                _ = v74
            }
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    var v75 int16 = (v4 | v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v75
    _ = v8
}
