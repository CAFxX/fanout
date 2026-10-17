package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int8; F1 int64; F2 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint }
func genFn0() uint8 {
    v7 := GenS1{F0: 42, F1: 2147483648}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 float32 = 3.14159
        fmt.Println(v9, v9)
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := []int{-2147483648, (fi8 / (((fi10 | fi8)) | 1)), fi8}
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    _ = v7
    return 59
}

func main() {
    var v1 bool = ((0 >= 857) && (!(0 <= 0)))
    var v2 int32 = 467
    var v3 int8 = -128
    v4 := []uint32{65535, 334, 32}
    var v5 uint = 4294967295
    v6 := []string{("a" + "\n\t")}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "hello" {
    case "a":
        switch "a" {
        case "a":
            for fi12 := 0; fi12 < 2; fi12++ {
                v3 -= (^(v3))
                _ = fi12
            }
        case "foo bar":
            v13 := GenS1{F0: (-(v3)), F1: 370}
            _ = v13
        default:
            switch -9223372036854775808 {
            case 32768:
                var v14 int64 = 42
                _ = v14
            case 0:
                switch 1 {
                case 0:
                    fmt.Println(v5, v5)
                default:
                    if len(v4) > 0 {
                        v4[0] = 631
                    }
                }
            case -527:
                if (-1 <= 0) {
                    var v15 uint32 = 711
                    var v16 uint16 = 32767
                    _ = v15
                    _ = v16
                } else {
                    if v1 {
                        v17 := GenS1{F0: 1, F1: 0}
                        var v18 int = -9223372036854775808
                        if (true || v1) {
                            fmt.Println(v3, v3, v5)
                            fmt.Println(v5)
                            if false {
                                fmt.Println(v18)
                                switch v18 {
                                case -720:
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v19 uint8 = 0
                                        _ = v19
                                    default:
                                        if len(v4) > 0 {
                                            v4[0] = 65535
                                        }
                                    }
                                case 546:
                                    var v20 int = (v18 % (((v18 >> ((v5) & 63))) | 1))
                                    _ = v20
                                case 0:
                                    switch ("hello" + ("\n\t" + "\n\t")) {
                                    case "\n\t":
                                        fmt.Println(v3)
                                    case "a":
                                        v5 -= 4294967295
                                    default:
                                        _ = len(v4)
                                    }
                                default:
                                    var v21 uint8 = 128
                                    _ = v21
                                }
                            } else {
                                v2 = (-(v2))
                            }
                        } else {
                            for fi22 := 0; fi22 < 17; fi22++ {
                                v5 = 0
                                _ = len(v4)
                                _ = fi22
                            }
                        }
                        _ = v17
                        _ = v18
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        v5 = (^(v5))
                        v23 := GenS1{F0: -1, F1: 975}
                        _ = v23
                    }
                }
            default:
                var v24 string = "\n\t"
                _ = v24
            }
        }
    default:
        if (!true) {
            var v25 uint32 = 4294967295
            _ = v25
        } else {
            v1 = (!true)
            v1 = true
            var v26 int = 42
            _ = v26
        }
    }
    var v27 int8 = v3
    if ((!(false && (42 <= 2))) && (!(true || (42 != 2)))) {
        if (42 >= 927) {
            if ((((3.14159 - float64((0.0 / (float32(1.5) - 0.1)))) - 1e10) / (0.1 + -2.25)) < 3.14159) {
                var v28 int = 9223372036854775807
                fmt.Println(v5, v27, v3)
                fmt.Println(v27, v27)
                _ = v28
            } else {
                v29 := GenS2{F0: -128, F1: 510, F2: 42}
                v30 := GenS1{F0: (v27 + v27), F1: 0}
                v31 := GenS2{F0: v3, F1: 635, F2: 212}
                _ = v29
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 1; fi32++ {
                if ((810 >= int64(v3)) || (9223372036854775807 > -904)) {
                    var v33 uint32 = 2147483648
                    fmt.Println(v5)
                    var v34 int = int(v2)
                    _ = v33
                    _ = v34
                }
                fmt.Println(v2)
                _ = fi32
            }
            v27 -= (v3 + v3)
        } else {
            switch (("a" + "a") + "hello") {
            case "foo bar":
                v1 = v1
            }
            var v35 uint = 18446744073709551615
            v35 -= v35
            _ = v35
        }
        v36 := GenS1{F0: int8(v5), F1: 348}
        var v37 uint64 = 4294967295
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 3; fi38++ {
        if v1 {
            v39 := GenS2{F0: v27, F1: 1, F2: 42}
            v40 := GenS3{F0: 529}
            _ = v39
            _ = v40
        } else {
            fmt.Println(v27, v1)
            fmt.Println(fi38)
        }
        v3 += v27
        for fi41 := 0; fi41 < 17; fi41++ {
            v4 = append(v4, 42)
            _ = fi41
        }
        _ = fi38
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        if true {
            v4 = append(v4, 804)
            v5 += 4294967295
        } else {
            v3 = v27
            fmt.Println(v27, v3, v5)
            switch "" {
            case "a":
                var v43 bool = (0 < v27)
                _ = v43
            case "\n\t":
                if ((609 == 2) && (0 >= 127)) {
                    var v44 uint32 = 42
                    _ = v44
                }
            case "hello":
                v2 -= 0
            default:
                switch (fi42 - fi42) {
                case 1:
                    var v45 int16 = 42
                    _ = v45
                }
            }
        }
        var v46 int = (^(fi42))
        switch ("a" + ((("hello" + "foo bar") + "a") + "")) {
        case "":
            v47 := GenS4{F0: 65535, F1: 42}
            _ = v47
        case "foo bar":
            if false {
                v48 := GenS1{F0: v27, F1: 42}
                _ = v48
            } else {
                v49 := GenS3{F0: 4294967295}
                var v50 int16 = 0
                _ = v49
                _ = v50
            }
        }
        _ = fi42
        _ = v46
    }
    var v51 uint16 = 400
    var v52 bool = (1 < 2147483648)
    if true {
        if false {
            fmt.Println(v3)
            if v1 {
                v51 = v51
                v53 := GenS5{F0: 2, F1: v5}
                _ = v53
            }
            fmt.Println(v2, v1, v27)
        } else {
            v54 := GenS1{F0: (-(v27)), F1: 1}
            for fi55 := 0; fi55 < 4; fi55++ {
                fmt.Println(v51)
                var v56 int16 = 42
                _ = fi55
                _ = v56
            }
            v3 += 91
            _ = v54
        }
        if v52 {
            var v57 int8 = 127
            _ = v57
        }
        var v58 int64 = 68
        _ = v58
    }
    switch "foo bar" {
    case "hello":
        v2 = v2
    case "\n\t":
        v5 = (v5 ^ (v5 & (v5 * uint(v51))))
    }
    var v59 int64 = -9223372036854775808
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v59
    _ = v6
}
