package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint16; F1 bool; F2 string }
func genFn0() int16 {
    if (940 != 0) {
        if true {
            v6 := GenS1{F0: 42, F1: -128}
            var v7 float64 = ((((3.14159 + 1.5) * (-2.25 + float64(0.1))) - ((-2.25 - 1e10) * 0.0)) * 0.0)
            _ = v6
            _ = v7
        }
    }
    var v8 uint32 = 65535
    v9 := []int{0, -124, 2147483647}
    if (((!((18446744073709551615 < 0) && true)) || true) && (1 <= 127)) {
        switch 9223372036854775807 {
        case 2147483647:
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        default:
            if (!(!true)) {
                v10 := GenS1{F0: 800, F1: 0}
                _ = v10
            } else {
                var v11 string = (("foo bar" + ((("hello" + "\n\t") + ("\n\t" + "\n\t")) + (("\n\t" + ("" + "\n\t")) + ""))) + "\n\t")
                _ = v11
            }
        }
    } else {
        v8 += 65535
        v9 = append(v9, 307)
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 float32 = ((3.14159 - ((1.5 / 1e10) / (0.0 - (3.14159 / (0.0 + (float32(3.14159) - ((3.14159 + float32((0.1 / (1.5 / 1e10)))) * ((1.5 * -2.25) + 0.1)))))))) / 1e10)
            v8 = (^(v8))
            switch ("\n\t" + ("foo bar" + (("hello" + (("hello" + "a") + "\n\t")) + "foo bar"))) {
            case "":
                if ((false && true) || (1 >= 128)) {
                    v14 := GenS1{F0: -9223372036854775808, F1: 100}
                    v15 := GenS1{F0: -595, F1: 28}
                    _ = v14
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 2; fi16++ {
                        var v17 float64 = -2.25
                        switch "foo bar" {
                        case "\n\t":
                            var v18 uint64 = 1
                            _ = v18
                        case "hello":
                            if len(v9) > 0 {
                                v9[0] = fi12
                            }
                        default:
                            var v19 float32 = ((((((v13 - 0.0) * float32(0.1)) + (v13 - float32((3.14159 + v17)))) + (0.0 - -2.25)) - v13) - (0.0 / (v13 / 1.5)))
                            _ = v19
                        }
                        _ = fi16
                        _ = v17
                    }
                }
            case "a":
                for fi20 := 0; fi20 < 18; fi20++ {
                    fmt.Println(v13)
                    _ = fi20
                }
            case "foo bar":
                var v21 uint16 = 65535
                _ = v21
            }
            _ = fi12
            _ = v13
        }
    }
    v8 -= v8
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 0.0
    var v3 int = 905
    v4 := []int{(v3 ^ -9223372036854775808), v3, v3, v3}
    var v5 int32 = -158
    v22 := GenS2{F0: 1, F1: 117}
    var v23 int64 = 42
    switch (v3 >> ((1) & 63)) {
    case 1:
        if ("foo bar" == "hello") {
            v3 += v3
            var v24 int = v3
            if false {
                v25 := GenS1{F0: v23, F1: 100}
                _ = v25
            } else {
                var v26 uint32 = 65535
                var v27 int32 = v5
                var v28 uint8 = 86
                _ = v26
                _ = v27
                _ = v28
            }
            _ = v24
        }
    case 0:
        fmt.Println(v3, v23, v3)
    case 402:
        for fi29 := 0; fi29 < 11; fi29++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch (-(fi29)) {
            case 128:
                var v30 int16 = 42
                _ = v30
            case 322:
                switch "hello" {
                case "hello":
                    v23 += v23
                }
            case 42:
                var v31 uint16 = 2
                _ = v31
            default:
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    var v32 uint16 = 230
                    _ = v32
                case "":
                    v33 := GenS1{F0: 0, F1: 42}
                    _ = v33
                }
            }
            for fi34 := 0; fi34 < 6; fi34++ {
                delete(v1, ("a" + "\n\t"))
                if ((true || false) && true) {
                    fmt.Println(v5, fi34, fi34)
                    fmt.Println(v3, v3)
                    _ = len(v1)
                }
                _ = fi34
            }
            _ = fi29
        }
    default:
        var v35 uint8 = 42
        _ = v35
    }
    switch "a" {
    case "a":
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v36 int8 = 42
            fmt.Println(v36)
            _ = v36
        }
    case "hello":
        if true {
            v37 := GenS1{F0: (v23 | (v23 << ((65535) & 63))), F1: 76}
            v38 := GenS1{F0: (v23 / (((v23 & 0)) | 1)), F1: int8(v23)}
            _ = v37
            _ = v38
        }
    case "\n\t":
        if false {
            v4 = append(v4, v3)
        }
    default:
        v39 := GenS2{F0: 128, F1: 65535}
        _ = v39
    }
    _ = len(v4)
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        var v40 uint64 = 825
        fmt.Println(v3, v40)
        _ = v40
    }
    v2 -= (v2 * (((v2 / ((((float32(3.14159) + v2) / v2) / (v2 * (((0.1 - v2) - ((v2 + v2) - 0.1)) - float32((float64(((v2 / (v2 - v2)) * 0.0)) * (float64(v2) + 1.5)))))) + v2)) - float32((float64((v2 + float32(0.1))) + (-2.25 + ((1.5 / 3.14159) + (3.14159 / 3.14159)))))) / v2))
    v41 := GenS2{F0: 255, F1: 874}
    switch ("hello" + "\n\t") {
    case "\n\t":
        fmt.Println(v5, v23)
    default:
        fmt.Println(v2)
    }
    fmt.Println(v23, v3, v5)
    v2 -= v2
    v23 = (v23 / ((v23) | 1))
    v2 -= (float32(1e10) + 1e10)
    if (("foo bar" >= (("\n\t" + "") + "hello")) && true) {
        if (0.0 == 0.1) {
            var v42 int64 = v23
            var v43 int32 = v5
            _ = v42
            _ = v43
        } else {
            fmt.Println(v2, v23)
        }
        v2 = v2
        fmt.Println(v23, v2, v5)
    } else {
        v44 := GenS2{F0: 0, F1: 65535}
        v45 := GenS3{F0: 42, F1: (!true), F2: (("" + ("" + ("" + ""))) + "")}
        _ = v44
        _ = v45
    }
    fmt.Println(v23, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
