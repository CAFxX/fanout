package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }
type GenS2 struct { F0 int; F1 bool; F2 int }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 uint8; F1 uint; F2 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "hello"
    var v3 float64 = 1e10
    var v4 string = (v2 + v2)
    var v5 float32 = ((0.0 + 1.5) - ((3.14159 - 1e10) * (1e10 / 3.14159)))
    var v6 int16 = 42
    if true {
        var v7 int32 = 2147483647
        _ = v7
    } else {
        var v8 uint32 = 662
        v4 = ("a" + v4)
        v9 := v1[600]
        _ = v8
        _ = v9
    }
    var v10 float32 = (v5 + v5)
    if ((2147483648 <= 2) || (2147483647 != 42)) {
        for fi11 := 0; fi11 < 4; fi11++ {
            if false {
                v6 -= 267
            } else {
                for fi12 := 0; fi12 < 6; fi12++ {
                    v13 := GenS1{F0: (fi12 - (fi11 & fi11)), F1: 4294967295, F2: v4}
                    _ = fi12
                    _ = v13
                }
                v1[1] = (94 != 1)
                if (0 != 1) {
                    v14 := GenS2{F0: (fi11 % ((62) | 1)), F1: (true || true), F2: fi11}
                    _ = v14
                } else {
                    v2 = v2
                    switch (((v2 + "\n\t") + v4) + v4) {
                    case "":
                        v15 := []int{-806}
                        _ = v15
                    case "foo bar":
                        var v16 uint32 = 2147483647
                        _ = v16
                    case "\n\t":
                        if true {
                            var v17 float32 = (v10 / v10)
                            _ = v17
                        } else {
                            v4 = ((v4 + ((("foo bar" + v4) + ("foo bar" + "foo bar")) + "a")) + (v4 + "foo bar"))
                            var v18 int64 = 232
                            _ = v18
                        }
                    default:
                        for fi19 := 0; fi19 < 18; fi19++ {
                            var v20 uint = 88
                            _ = fi19
                            _ = v20
                        }
                    }
                }
            }
            _ = fi11
        }
        var v21 int16 = v6
        for fi22 := 0; fi22 < 18; fi22++ {
            v23 := GenS2{F0: 4294967295, F1: false, F2: fi22}
            for fi24 := 0; fi24 < 3; fi24++ {
                switch (("\n\t" + "hello") + "\n\t") {
                case "\n\t":
                    fmt.Println(v21)
                    _ = len(v1)
                    if false {
                        fmt.Println(v4, v5)
                        v25 := GenS3{F0: 2, F1: 50}
                        _ = v25
                    } else {
                        switch -157 {
                        case 32767:
                            v26 := GenS3{F0: 18446744073709551615, F1: 2}
                            _ = v26
                        default:
                            if (!true) {
                                var v27 int = (fi22 + fi22)
                                _ = v27
                            } else {
                                var v28 int32 = -1
                                var v29 uint16 = 562
                                v29 += (v29 >> ((2) & 63))
                                _ = v28
                                _ = v29
                            }
                        }
                        v30 := GenS1{F0: (fi22 ^ (fi22 % (((-(fi24))) | 1))), F1: 42, F2: v4}
                        v31 := v1[1]
                        v6 = -1
                        switch fi22 {
                        case -9223372036854775808:
                            v32 := v1[2]
                            switch "foo bar" {
                            case "a":
                                fmt.Println(v4, v32)
                            }
                            _ = v32
                        }
                        _ = v30
                        _ = v31
                    }
                case "hello":
                    v21 += -1
                case "a":
                    if (true || (-128 == 114)) {
                        v33 := GenS1{F0: (fi24 | 805), F1: 1, F2: v2}
                        _ = v33
                    }
                }
                _ = len(v1)
                _ = fi24
            }
            for fi34 := 0; fi34 < 14; fi34++ {
                v6 += 1
                _ = fi34
            }
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        var v35 int64 = 9223372036854775807
        _ = v35
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        v37 := GenS1{F0: fi36, F1: 606, F2: v2}
        v4 = v4
        v10 = 1e10
        _ = fi36
        _ = v37
    }
    fmt.Println(v5, v2, v10)
    switch ("\n\t" + (((v4 + v4) + "") + (("hello" + (v2 + "hello")) + "hello"))) {
    case "hello":
        v6 = (^(v6))
    case "foo bar":
        var v38 float64 = (float64(v5) + ((((((float64(v5) - (float64((((v5 + v5) * v10) * v5)) * 0.0)) + 0.0) / float64((v10 / 1.0))) * 3.14159) * 3.14159) - v3))
        _ = v38
    }
    switch (v2 + ((v4 + "a") + ("foo bar" + (v2 + "a")))) {
    case "foo bar":
        v1[2147483647] = false
    case "\n\t":
        switch int(v6) {
        case -1:
            if (857 != 0) {
                delete(v1, 2147483647)
                v39 := GenS4{F0: 854, F1: 526}
                _ = v39
            } else {
                for fi40 := 0; fi40 < 19; fi40++ {
                    if (100 > 0) {
                        for fi41 := 0; fi41 < 9; fi41++ {
                            var v42 uint16 = 42
                            switch v2 {
                            case "hello":
                                v3 = 1e10
                            default:
                                v3 += (v3 - v3)
                            }
                            _ = fi41
                            _ = v42
                        }
                        delete(v1, 0)
                        v2 = v2
                    }
                    fmt.Println(v10, v3, v5)
                    _ = fi40
                }
                var v43 uint16 = 32767
                _ = v43
            }
        default:
            for fi44 := 0; fi44 < 12; fi44++ {
                var v45 int32 = 42
                _ = fi44
                _ = v45
            }
        }
    case "hello":
        v46 := []int{32768}
        _ = v46
    default:
        v4 = "hello"
    }
    v47 := GenS5{F0: 24, F1: 236, F2: 1}
    if (-1 == -9223372036854775808) {
        fmt.Println(v5)
        for fi48 := 0; fi48 < 1; fi48++ {
            var v49 int8 = -128
            fmt.Println(fi48, v6, v5)
            _ = fi48
            _ = v49
        }
        var v50 string = (v4 + "hello")
        _ = v50
    }
    v2 = (v4 + (("foo bar" + "foo bar") + ""))
    var v51 int16 = v6
    var v52 float64 = v3
    if true {
        for fi53 := 0; fi53 < 15; fi53++ {
            v6 += (v51 << ((2) & 63))
            for fi54 := 0; fi54 < 10; fi54++ {
                var v55 float64 = (v3 + (v3 * 1.5))
                var v56 uint32 = 1
                _ = len(v1)
                v57 := GenS1{F0: 0, F1: 857, F2: ("\n\t" + "foo bar")}
                var v58 float32 = v5
                _ = fi54
                _ = v55
                _ = v56
                _ = v57
                _ = v58
            }
            _ = fi53
        }
        v59 := []int{-9223372036854775808, -9223372036854775808}
        v51 -= v6
        _ = v59
    } else {
        for fi60 := 0; fi60 < 6; fi60++ {
            v61 := []int{(-(fi60)), 4294967295, (fi60 % (((fi60 * fi60)) | 1)), 0}
            _ = fi60
            _ = v61
        }
        var v62 uint16 = 160
        v10 += float32(v3)
        _ = v62
    }
    v51 -= v51
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
