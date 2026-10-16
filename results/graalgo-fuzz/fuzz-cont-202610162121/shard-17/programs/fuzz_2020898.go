package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint }
type GenS2 struct { F0 int32; F1 uint64; F2 int16 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 uint8; F2 int }

func main() {
    v1 := []string{"foo bar"}
    var v2 uint64 = 831
    var v3 string = (("a" + "\n\t") + (("a" + "a") + "foo bar"))
    var v4 int16 = 32767
    var v5 uint8 = 127
    switch "hello" {
    case "foo bar":
        v2 += v2
    default:
        var v6 uint = 65535
        _ = v6
    }
    v7 := GenS1{F0: 127, F1: 42, F2: uint(v2)}
    var v8 string = (v3 + v3)
    v4 -= (-(v4))
    if (1.5 != (((1.5 * 0.0) + 0.1) + (3.14159 / float32(0.1)))) {
        var v9 int8 = 42
        if (v3 >= v3) {
            v2 = (v2 - (v2 | v2))
        } else {
            v10 := GenS2{F0: -3, F1: 622, F2: 0}
            var v11 int64 = 9223372036854775807
            if ((2147483647 == 42) || ((v5 == (v5 / ((128) | 1))) && false)) {
                v9 += (v9 >> ((42) & 63))
            } else {
                fmt.Println(v2, v2, v3)
                var v12 bool = true
                _ = v12
            }
            _ = v10
            _ = v11
        }
        _ = v9
    } else {
        if true {
            if (-1 != 127) {
                v5 -= v5
                _ = len(v1)
                for fi13 := 0; fi13 < 7; fi13++ {
                    if false {
                        var v14 uint32 = 4294967295
                        v1 = append(v1, v8)
                        _ = v14
                    }
                    _ = fi13
                }
            } else {
                v4 -= 42
            }
        } else {
            v2 = v2
        }
        v3 = v3
        var v15 uint64 = (^(v2))
        _ = v15
    }
    v8 = (v3 + "a")
    if false {
        var v16 uint = 874
        v2 += v2
        v17 := GenS2{F0: 0, F1: (v2 ^ (v2 * v2)), F2: v4}
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 11; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                v4 += (v4 * (-(v4)))
                var v20 string = "foo bar"
                v20 = "hello"
                _ = fi19
                _ = v20
            }
            var v21 uint64 = v2
            if false {
                switch ("a" + v3) {
                case "hello":
                    switch (fi18 * 127) {
                    case -403:
                        v4 += (v4 << ((255) & 63))
                    case 2147483647:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        for fi22 := 0; fi22 < 14; fi22++ {
                            switch -30 {
                            case 42:
                                v23 := GenS3{F0: false, F1: (1 <= 1)}
                                switch v8 {
                                case "foo bar":
                                    _ = len(v1)
                                default:
                                    v3 = "foo bar"
                                }
                                var v24 int64 = -149
                                _ = v23
                                _ = v24
                            case 32767:
                                v2 += (v2 - 663)
                            case 128:
                                var v25 bool = true
                                _ = v25
                            default:
                                v5 = (^(v5))
                            }
                            switch "hello" {
                            case "foo bar":
                                var v26 float32 = (1.5 * 0.0)
                                _ = v26
                            case "\n\t":
                                if true {
                                    v21 = v2
                                    var v27 float64 = (1.5 * 1.5)
                                    _ = v27
                                }
                            }
                            v28 := GenS4{F0: -1}
                            _ = fi22
                            _ = v28
                        }
                    }
                default:
                    for fi29 := 0; fi29 < 11; fi29++ {
                        v5 = v5
                        _ = fi29
                    }
                }
                var v30 uint = 255
                _ = v30
            } else {
                if ((("" + "a") <= v8) || ((false && ("a" == v8)) && (((!false) || true) && (("\n\t" + ((("a" + ((("hello" + "a") + (v8 + v8)) + ((v3 + "\n\t") + ""))) + v3) + ((v3 + "hello") + "\n\t"))) > ((("" + (v8 + v8)) + "\n\t") + ""))))) {
                    switch "a" {
                    case "":
                        switch "\n\t" {
                        case "":
                            switch (v3 + "foo bar") {
                            case "hello":
                                v31 := GenS2{F0: -2147483648, F1: 596, F2: (v4 * (v4 * v4))}
                                _ = v31
                            case "":
                                v2 = (v2 + 2)
                            default:
                                if len(v1) > 0 {
                                    v1[0] = v8
                                }
                            }
                        case "\n\t":
                            v32 := GenS1{F0: 127, F1: (v4 << ((18446744073709551615) & 63)), F2: 743}
                            _ = v32
                        }
                    }
                    for fi33 := 0; fi33 < 14; fi33++ {
                        v3 = v3
                        var v34 uint = 2
                        v1 = append(v1, (v3 + v8))
                        _ = fi33
                        _ = v34
                    }
                }
            }
            _ = fi18
            _ = v21
        }
        fmt.Println(v4)
        if true {
            fmt.Println(v5, v4)
            v8 = ""
            _ = len(v1)
        } else {
            var v35 float32 = 1e10
            for fi36 := 0; fi36 < 13; fi36++ {
                v37 := GenS4{F0: 2147483647}
                _ = fi36
                _ = v37
            }
            _ = v35
        }
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        var v39 int = fi38
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        for fi40 := 0; fi40 < 11; fi40++ {
            v41 := GenS4{F0: 2147483647}
            if ("foo bar" <= v8) {
                v42 := GenS5{F0: 4294967295, F1: (v5 & (v5 % (((v5 / ((v5) | 1))) | 1))), F2: 127}
                _ = v42
            }
            v2 = 4294967295
            _ = fi40
            _ = v41
        }
        _ = fi38
        _ = v39
    }
    for fi43 := 0; fi43 < 17; fi43++ {
        var v44 bool = false
        var v45 uint32 = 626
        _ = fi43
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 7; fi46++ {
        fmt.Println(fi46, v5)
        _ = fi46
    }
    v1 = append(v1, ("\n\t" + v3))
    if false {
        fmt.Println(v5, v5, v5)
        v1 = append(v1, v3)
        v47 := GenS2{F0: 2147483647, F1: 567, F2: v4}
        _ = v47
    }
    var v48 int8 = 19
    for fi49 := 0; fi49 < 3; fi49++ {
        if (true || false) {
            var v50 uint64 = (^(v2))
            for fi51 := 0; fi51 < 9; fi51++ {
                var v52 string = v3
                if true {
                    v50 = (v2 / (((v2 >> ((418) & 63))) | 1))
                    switch (v3 + v3) {
                    case "foo bar":
                        var v53 float64 = ((float64(((1.5 + (-2.25 / 1.0)) + ((((((float32((0.1 * (0.0 + float64(0.1)))) + 0.0) + (((float32(((1e10 * float64((((float32(0.0) - ((0.1 / 1e10) * 1e10)) + 0.1) + -2.25))) * 1.5)) * -2.25) + ((((3.14159 - 3.14159) + 1.5) - 0.0) / 1.5)) + 1e10)) - 0.1) + (3.14159 - 0.1)) * (((1.5 - 1e10) / (-2.25 + -2.25)) / 1.5)) + 0.0))) * 0.1) + (3.14159 - 0.0))
                        v54 := GenS1{F0: 127, F1: v4, F2: 2}
                        _ = v53
                        _ = v54
                    }
                    v55 := GenS3{F0: true, F1: (false || true)}
                    v1 = append(v1, ("" + (v52 + v52)))
                    _ = v55
                } else {
                    v56 := GenS2{F0: 1, F1: v50, F2: 225}
                    var v57 uint16 = 255
                    if (true || true) {
                        v58 := GenS3{F0: (!(v4 == 827)), F1: true}
                        v57 = v57
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        _ = v58
                    } else {
                        var v59 int64 = 848
                        var v60 uint16 = (v57 - 1)
                        var v61 int16 = v4
                        _ = v59
                        _ = v60
                        _ = v61
                    }
                    _ = v56
                    _ = v57
                }
                var v62 string = ""
                _ = fi51
                _ = v52
                _ = v62
            }
            var v63 int8 = 1
            _ = v50
            _ = v63
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v64 uint16 = 65535
            v5 -= v5
            _ = v64
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (int64(v48) != 0) {
            switch ("\n\t" + v3) {
            case "hello":
                switch (fi49 % ((fi49) | 1)) {
                case 32768:
                    var v65 uint16 = 32767
                    _ = v65
                case 32767:
                    if len(v1) > 0 {
                        v1[0] = v3
                    }
                case 127:
                    switch fi49 {
                    case -9223372036854775808:
                        if false {
                            v2 += 18446744073709551615
                            v66 := GenS4{F0: 703}
                            _ = v66
                        } else {
                            switch v3 {
                            case "hello":
                                v1 = append(v1, "\n\t")
                            case "":
                                fmt.Println(v48, v48)
                            case "\n\t":
                                for fi67 := 0; fi67 < 18; fi67++ {
                                    v48 = (v48 + v48)
                                    var v68 int8 = 100
                                    v2 = 4294967295
                                    _ = fi67
                                    _ = v68
                                }
                            }
                        }
                    }
                default:
                    switch (fi49 % ((9223372036854775807) | 1)) {
                    case 1:
                        v48 = v48
                    case 0:
                        v69 := GenS1{F0: v48, F1: (v4 ^ (v4 | v4)), F2: 1}
                        _ = v69
                    case 32767:
                        switch (^(fi49)) {
                        case 2147483647:
                            fmt.Println(v48, v2, v2)
                        default:
                            switch v8 {
                            case "":
                                fmt.Println(v8, v4, v48)
                            case "foo bar":
                                fmt.Println(v3, v2)
                            case "a":
                                var v70 uint = 4294967295
                                _ = v70
                            }
                        }
                    default:
                        v1 = append(v1, "hello")
                    }
                }
            case "foo bar":
                v2 += 18446744073709551615
            default:
                v4 += v4
            }
        }
        _ = fi49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v7
    _ = v8
}
