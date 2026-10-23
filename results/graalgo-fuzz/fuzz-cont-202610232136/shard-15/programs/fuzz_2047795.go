package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 uint8 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]string)
    var v3 float64 = (-2.25 / 1e10)
    delete(v2, -2147483648)
    if (!true) {
        var v4 uint16 = 42
        _ = v4
    }
    var v5 uint16 = 2
    switch "\n\t" {
    case "hello":
        fmt.Println(v1, v5, v1)
    case "a":
        fmt.Println(v1, v5, v3)
    }
    v6 := []int{-1}
    switch 32767 {
    case -657:
        v1 += v1
    default:
        if (true || true) {
            if false {
                v7 := GenS1{F0: 2147483647}
                _ = v7
            } else {
                if ("a" >= "foo bar") {
                    if (false && true) {
                        var v8 float32 = 1.5
                        if false {
                            for fi9 := 0; fi9 < 10; fi9++ {
                                fmt.Println(v1, v5, v5)
                                v10 := GenS1{F0: fi9}
                                fmt.Println(v5, v1, v8)
                                _ = fi9
                                _ = v10
                            }
                            fmt.Println(v5, v8)
                            if false {
                                v1 = (v1 ^ 0)
                                switch "" {
                                case "hello":
                                    var v11 float64 = ((v3 / 0.1) / ((float64(float32((float64((v8 * v8)) + v3))) - float64(v8)) - v3))
                                    _ = v11
                                }
                            }
                        } else {
                            for fi12 := 0; fi12 < 20; fi12++ {
                                v13 := GenS1{F0: (fi12 / ((fi12) | 1))}
                                _ = fi12
                                _ = v13
                            }
                        }
                        v14 := v2[42]
                        _ = v14
                        _ = v8
                    }
                    if true {
                        for fi15 := 0; fi15 < 3; fi15++ {
                            v16 := v2[(fi15 / ((-9223372036854775808) | 1))]
                            _ = fi15
                            _ = v16
                        }
                        var v17 uint32 = 2
                        _ = v17
                    }
                } else {
                    switch "\n\t" {
                    case "hello":
                        if true {
                            for fi18 := 0; fi18 < 12; fi18++ {
                                for fi19 := 0; fi19 < 16; fi19++ {
                                    v20 := GenS2{F0: 1}
                                    _ = fi19
                                    _ = v20
                                }
                                switch fi18 {
                                case 34:
                                    fmt.Println(v3, fi18, v5)
                                default:
                                    v21 := GenS3{F0: 2, F1: 0}
                                    _ = v21
                                }
                                _ = fi18
                            }
                            v1 = (v1 & -1)
                        }
                    case "":
                        v22 := GenS4{F0: (("hello" + ("\n\t" + "\n\t")) + ("\n\t" + "hello")), F1: 42}
                        _ = v22
                    }
                }
                v5 -= (v5 / (((^(v5))) | 1))
                if (161 > 255) {
                    if false {
                        var v23 string = "a"
                        _ = v23
                    } else {
                        delete(v2, -2147483648)
                    }
                    var v24 int8 = 0
                    _ = v24
                } else {
                    fmt.Println(v1, v5, v5)
                }
            }
            switch "foo bar" {
            case "hello":
                if (42 > 2147483647) {
                    var v25 uint = 4294967295
                    _ = v25
                }
            }
            switch "\n\t" {
            case "":
                v26 := GenS4{F0: "", F1: 127}
                _ = v26
            case "hello":
                for fi27 := 0; fi27 < 7; fi27++ {
                    if false {
                        v5 += 32767
                    } else {
                        fmt.Println(v1)
                    }
                    _ = len(v6)
                    v5 += v5
                    _ = fi27
                }
            }
        }
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        if (true || true) {
            _ = len(v6)
            var v29 uint32 = 1
            delete(v2, fi28)
            _ = v29
        } else {
            var v30 int = 128
            var v31 float64 = v3
            _ = v30
            _ = v31
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v1 -= v1
        _ = fi28
    }
    if (3.14159 >= v3) {
        for fi32 := 0; fi32 < 6; fi32++ {
            fmt.Println(v3, v5, v1)
            _ = fi32
        }
    } else {
        if false {
            switch 1 {
            case 9223372036854775807:
                v1 = v1
            case 128:
                var v33 int64 = 9223372036854775807
                _ = v33
            default:
                _ = len(v6)
            }
            v3 += (v3 / (0.1 * v3))
            if false {
                v3 = float64(3.14159)
                var v34 float64 = ((v3 / 1.5) / (v3 + float64(0.1)))
                _ = v34
            } else {
                var v35 bool = false
                _ = v35
            }
        }
        v1 -= -2147483648
        fmt.Println(v1, v1, v3)
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        switch (^(fi36)) {
        case 32768:
            for fi37 := 0; fi37 < 20; fi37++ {
                if true {
                    if false {
                        v3 = (0.1 + (1.5 * 1.5))
                    }
                    var v38 float32 = float32(((v3 - (1e10 / v3)) * v3))
                    var v39 string = "\n\t"
                    _ = len(v6)
                    v1 = (v1 % ((v1) | 1))
                    var v40 string = v39
                    var v41 uint8 = 2
                    _ = v38
                    _ = v39
                    _ = v40
                    _ = v41
                }
                var v42 int16 = 1
                v42 -= (v42 >> ((65535) & 63))
                _ = fi37
                _ = v42
            }
        case -1:
            if (277 == 1) {
                switch (-(fi36)) {
                case 103:
                    var v43 int = fi36
                    _ = v43
                case 4294967295:
                    var v44 uint32 = 736
                    _ = v44
                case -9223372036854775808:
                    v5 = (v5 / (((^(v5))) | 1))
                }
            } else {
                switch "" {
                case "hello":
                    v45 := GenS1{F0: (fi36 - 0)}
                    _ = v45
                default:
                    v3 = (v3 / ((((1.5 / (float64(1.5) / v3)) + float64(float32(float64(0.1)))) / v3) + v3))
                }
            }
        }
        var v46 string = "foo bar"
        _ = fi36
        _ = v46
    }
    if true {
        v47 := GenS5{F0: (v3 != float64(((-2.25 * ((0.1 - 0.1) * float32((((1.5 - v3) / 1.0) * -2.25)))) + (3.14159 + 1e10)))), F1: 9223372036854775807}
        var v48 uint32 = 953
        _ = v47
        _ = v48
    } else {
        v1 = -418
    }
    switch 1 {
    case -352:
        v2[-115] = "hello"
    case -2147483648:
        switch ((("" + "a") + "") + ("foo bar" + "\n\t")) {
        case "":
            if (!((-58 > 127) || false)) {
                var v49 float32 = (-2.25 / 0.1)
                _ = v49
            }
        case "foo bar":
            v50 := GenS3{F0: 128, F1: 1}
            _ = v50
        default:
            v1 = v1
        }
    default:
        var v51 uint8 = 127
        _ = v51
    }
    switch 0 {
    case 127:
        v52 := GenS6{F0: 1, F1: 60}
        _ = v52
    case -179:
        delete(v2, 829)
    case 1:
        for fi53 := 0; fi53 < 12; fi53++ {
            v54 := GenS2{F0: 1}
            _ = fi53
            _ = v54
        }
    }
    var v55 int16 = -592
    v1 = 2147483647
    fmt.Println(v3, v55, v5)
    if false {
        v56 := v2[128]
        delete(v2, 802)
        v56 = ""
        _ = v56
    } else {
        switch 127 {
        case 994:
            if ("foo bar" <= "\n\t") {
                if false {
                    _ = len(v2)
                    v3 = (-2.25 - (((1e10 - 0.0) / 1e10) - v3))
                } else {
                    switch "a" {
                    case "hello":
                        for fi57 := 0; fi57 < 17; fi57++ {
                            var v58 int8 = 127
                            delete(v2, fi57)
                            _ = fi57
                            _ = v58
                        }
                    case "\n\t":
                        var v59 int16 = v55
                        _ = v59
                    case "":
                        v60 := v2[42]
                        _ = v60
                    }
                    var v61 uint16 = v5
                    v62 := GenS7{F0: 32767, F1: -1000}
                    _ = v61
                    _ = v62
                }
                if len(v6) > 0 {
                    v6[0] = -145
                }
                if true {
                    switch (("foo bar" + "") + "hello") {
                    case "":
                        v55 -= 815
                    }
                    fmt.Println(v55, v3)
                } else {
                    if false {
                        for fi63 := 0; fi63 < 8; fi63++ {
                            var v64 uint8 = 0
                            var v65 float64 = 1e10
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                        var v66 int = 9223372036854775807
                        _ = v66
                    }
                }
            } else {
                for fi67 := 0; fi67 < 16; fi67++ {
                    v68 := GenS7{F0: 42, F1: fi67}
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    switch fi67 {
                    case -1:
                        v69 := GenS1{F0: -2147483648}
                        _ = v69
                    default:
                        v55 = v55
                    }
                    _ = fi67
                    _ = v68
                }
                if true {
                    var v70 float64 = float64(0.0)
                    if ((false && (!true)) && false) {
                        fmt.Println(v70, v1, v55)
                        _ = len(v6)
                    } else {
                        _ = len(v2)
                        v6 = append(v6, -1)
                        v71 := GenS3{F0: 255, F1: 0}
                        _ = v71
                    }
                    _ = v70
                } else {
                    v72 := GenS8{F0: 42}
                    v5 = 255
                    _ = v72
                }
            }
        }
        var v73 int = 2147483647
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v55
    _ = v6
}
