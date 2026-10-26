package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint64; F2 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 uint }

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]string)
    var v3 float32 = 3.14159
    v4 := make(map[uint32]string)
    var v5 uint16 = 536
    switch (("" + "foo bar") + "hello") {
    case "foo bar":
        var v6 int16 = 32767
        _ = v6
    case "":
        switch ("\n\t" + ("hello" + ("foo bar" + "\n\t"))) {
        case "":
            if (((0.1 - 1.5) + 1.5) < 3.14159) {
                v7 := v4[4294967295]
                fmt.Println(v1)
                _ = v7
            } else {
                v1 = 1
                v8 := v2[65535]
                _ = v8
            }
        case "foo bar":
            var v9 float64 = 0.1
            _ = v9
        case "hello":
            fmt.Println(v1, v3, v5)
        }
    default:
        if (!(v5 <= v5)) {
            var v10 float32 = v3
            _ = v10
        }
    }
    v11 := GenS1{F0: -1, F1: (0 != -124)}
    v12 := GenS1{F0: 42, F1: (42 != -1)}
    switch "a" {
    case "hello":
        if ((!(v1 >= -128)) && true) {
            v1 += (v1 + (v1 + v1))
        } else {
            fmt.Println(v1, v5, v5)
        }
    case "\n\t":
        v1 = v1
    case "":
        v13 := []int{-2147483648, -1}
        _ = v13
    default:
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := GenS2{F0: 0}
            _ = fi14
            _ = v15
        }
    }
    var v16 uint64 = 42
    v16 += 0
    switch "\n\t" {
    case "hello":
        if false {
            if false {
                var v17 uint8 = 128
                var v18 float64 = (1e10 - 0.1)
                _ = v17
                _ = v18
            } else {
                var v19 bool = (-387 < -806)
                for fi20 := 0; fi20 < 7; fi20++ {
                    switch fi20 {
                    case 128:
                        v19 = (!(v19 || (((v5 >> ((18446744073709551615) & 63)) != v5) && (v1 == (-(v1))))))
                    case -2147483648:
                        fmt.Println(v1, v1)
                    case 9223372036854775807:
                        var v21 uint8 = 42
                        _ = v21
                    }
                    var v22 uint = 436
                    var v23 uint32 = 65535
                    _ = fi20
                    _ = v22
                    _ = v23
                }
                _ = v19
            }
        } else {
            switch "hello" {
            case "":
                for fi24 := 0; fi24 < 18; fi24++ {
                    v25 := []int{fi24}
                    delete(v4, 314)
                    v25 = append(v25, 32767)
                    _ = fi24
                    _ = v25
                }
            case "foo bar":
                if (false || true) {
                    v4[2147483647] = "a"
                    v26 := GenS1{F0: -507, F1: (100 >= v1)}
                    _ = v26
                } else {
                    var v27 uint16 = (v5 * 255)
                    v28 := []int{9223372036854775807, 1, 4294967295}
                    v5 = v5
                    _ = v27
                    _ = v28
                }
            }
            if (v16 > 2) {
                if (v3 < 3.14159) {
                    fmt.Println(v3)
                    fmt.Println(v1)
                }
                v3 += v3
            }
        }
    case "a":
        for fi29 := 0; fi29 < 9; fi29++ {
            if (true || true) {
                var v30 int64 = 1
                for fi31 := 0; fi31 < 12; fi31++ {
                    if true {
                        if (v1 > v1) {
                            var v32 float64 = 0.0
                            v33 := GenS3{F0: -19, F1: 4294967295, F2: (-2147483648 <= -2147483648)}
                            v34 := GenS3{F0: (v1 ^ (v1 - 42)), F1: (v16 & 18446744073709551615), F2: (558 == 1)}
                            var v35 float32 = (0.1 - float32(v32))
                            _ = v32
                            _ = v33
                            _ = v34
                            _ = v35
                        } else {
                            fmt.Println(v16, v16, v5)
                            if false {
                                v16 = v16
                                var v36 uint = 42
                                switch fi29 {
                                case 32768:
                                    fmt.Println(v1)
                                default:
                                    v37 := []int{}
                                    _ = v37
                                }
                                _ = v36
                            } else {
                                fmt.Println(v16)
                                v38 := v2[0]
                                _ = v38
                            }
                        }
                        v39 := GenS3{F0: (-(v1)), F1: (v16 >> ((42) & 63)), F2: (((1.5 - v3) >= 0.1) && false)}
                        v5 += v5
                        _ = v39
                    }
                    if ("foo bar" <= "a") {
                        var v40 uint64 = (v16 - 0)
                        fmt.Println(v3, fi31)
                        _ = v40
                    }
                    if false {
                        if (!false) {
                            fmt.Println(v1, v5)
                        }
                    }
                    _ = fi31
                }
                _ = v30
            } else {
                switch (fi29 - fi29) {
                case 601:
                    v41 := []int{}
                    _ = v41
                }
                v1 -= (v1 ^ v1)
            }
            var v42 string = ""
            var v43 float64 = -2.25
            _ = fi29
            _ = v42
            _ = v43
        }
    default:
        v44 := GenS4{F0: -659}
        _ = v44
    }
    var v45 uint16 = v5
    v46 := []int{9223372036854775807, -446, -2147483648}
    var v47 int8 = 0
    if len(v46) >= 2 {
        v46 = v46[1:len(v46)]
    }
    if (!true) {
        var v48 float32 = (v3 / v3)
        if (true || (1e10 <= ((float32(-2.25) * v3) / 3.14159))) {
            var v49 uint8 = 42
            v50 := GenS2{F0: -292}
            _ = v49
            _ = v50
        }
        _ = len(v46)
        _ = v48
    }
    if false {
        switch "foo bar" {
        case "hello":
            v51 := GenS2{F0: -1}
            _ = v51
        case "foo bar":
            v2[4294967295] = "foo bar"
        case "a":
            v4[2147483648] = (("foo bar" + ("foo bar" + "a")) + ("foo bar" + ("hello" + "hello")))
        default:
            for fi52 := 0; fi52 < 1; fi52++ {
                switch "hello" {
                case "":
                    var v53 float32 = float32((0.0 - 1.5))
                    _ = v53
                case "\n\t":
                    v5 = 0
                case "foo bar":
                    for fi54 := 0; fi54 < 15; fi54++ {
                        if (!true) {
                            var v55 uint16 = 2
                            _ = v55
                        }
                        fmt.Println(v45, v5, fi52)
                        var v56 uint64 = (v16 % ((v16) | 1))
                        delete(v2, 1)
                        _ = fi54
                        _ = v56
                    }
                }
                v57 := GenS1{F0: 381, F1: false}
                _ = fi52
                _ = v57
            }
        }
        v45 = (v5 ^ (v45 << ((2) & 63)))
        v2[1] = "a"
    } else {
        var v58 int32 = 0
        _ = v58
    }
    v47 -= (^(v1))
    delete(v4, 0)
    delete(v4, 2147483647)
    v59 := GenS5{F0: v16, F1: 340}
    var v60 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v59
    _ = v60
}
