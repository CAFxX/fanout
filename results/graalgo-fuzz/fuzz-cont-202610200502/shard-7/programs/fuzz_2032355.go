package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint64 }
type GenS2 struct { F0 string; F1 uint16; F2 string }
type GenS3 struct { F0 bool; F1 int64; F2 uint32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{1, -233, 9223372036854775807, 32767}
    v3 := make(map[string]bool)
    v4 := []uint32{2, 1}
    var v5 int16 = -870
    v5 += (v5 ^ (v5 >> ((44) & 63)))
    var v6 bool = (!false)
    var v7 float64 = 0.1
    if (32767 != 2147483647) {
        var v8 uint = 255
        _ = v8
    }
    switch 0 {
    case -676:
        switch 127 {
        case 127:
            switch "foo bar" {
            case "\n\t":
                if (1.5 > 0.0) {
                    var v9 int64 = -636
                    fmt.Println(v5)
                    var v10 uint16 = 897
                    _ = v10
                    _ = v9
                } else {
                    var v11 uint16 = 0
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v13 := GenS1{F0: v6, F1: (("\n\t" + ("\n\t" + "")) + "\n\t"), F2: 4294967295}
                        var v14 int = fi12
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = v11
                }
            case "hello":
                v15 := GenS2{F0: "\n\t", F1: 1, F2: "foo bar"}
                _ = v15
            }
        case -2147483648:
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 uint16 = 65535
                for fi18 := 0; fi18 < 15; fi18++ {
                    if (v7 != v7) {
                        var v19 string = ""
                        delete(v1, 65535)
                        switch v19 {
                        case "":
                            var v20 int64 = -9223372036854775808
                            _ = v20
                        default:
                            v5 = v5
                        }
                        _ = v19
                    }
                    v3[("" + ("hello" + (("foo bar" + ("\n\t" + (("a" + "a") + "a"))) + ("a" + "a"))))] = (!v6)
                    v21 := v3["hello"]
                    _ = fi18
                    _ = v21
                }
                _ = fi16
                _ = v17
            }
        case -887:
            var v22 int8 = -119
            _ = v22
        default:
            v6 = (v6 || (0 == 0))
        }
    }
    if ((18446744073709551615 == 255) || v6) {
        for fi23 := 0; fi23 < 18; fi23++ {
            v3["hello"] = v6
            _ = fi23
        }
        switch 563 {
        case -703:
            var v24 float32 = -2.25
            _ = v24
        }
    }
    var v25 uint8 = 128
    v25 += v25
    v6 = v6
    var v26 int32 = -16
    switch "hello" {
    case "":
        for fi27 := 0; fi27 < 11; fi27++ {
            var v28 int64 = -1
            _ = fi27
            _ = v28
        }
    case "hello":
        if false {
            for fi29 := 0; fi29 < 20; fi29++ {
                var v30 uint16 = 1
                _ = fi29
                _ = v30
            }
            fmt.Println(v25, v5, v7)
            v6 = ((-(v26)) != (v26 / ((v26) | 1)))
        }
    case "a":
        var v31 float64 = ((v7 / ((v7 / v7) / -2.25)) / (v7 * v7))
        _ = v31
    }
    var v32 int64 = 1
    switch ("hello" + ("" + "")) {
    case "a":
        if false {
            v25 = 42
            var v33 string = "hello"
            _ = v33
        } else {
            var v34 string = "foo bar"
            v35 := GenS3{F0: ((-2.25 / 3.14159) >= ((0.1 + 0.0) - 0.1)), F1: v32, F2: 615}
            _ = v34
            _ = v35
        }
    case "foo bar":
        v32 -= (v32 | (v32 / (((v32 % (((v32 << ((2) & 63))) | 1))) | 1)))
    case "\n\t":
        var v36 uint64 = 515
        _ = v36
    }
    if true {
        v26 += (-(v26))
    } else {
        if (65535 != 1) {
            v2 = append(v2, -9223372036854775808)
            v6 = v6
            if len(v4) > 0 {
                v4[0] = 65535
            }
        } else {
            v26 = v26
            fmt.Println(v7, v32, v25)
        }
        for fi37 := 0; fi37 < 12; fi37++ {
            fmt.Println(v32, v5, v6)
            var v38 uint32 = 0
            _ = fi37
            _ = v38
        }
        v39 := GenS1{F0: ((2 == 29) && false), F1: ("\n\t" + ""), F2: 1}
        _ = v39
    }
    if v6 {
        if v6 {
            var v40 int64 = v32
            var v41 uint8 = v25
            _ = v40
            _ = v41
        }
        fmt.Println(v32, v32)
    } else {
        v42 := GenS3{F0: true, F1: v32, F2: 319}
        _ = v42
    }
    v5 = (^(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
