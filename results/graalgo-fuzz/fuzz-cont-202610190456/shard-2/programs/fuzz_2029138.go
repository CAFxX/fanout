package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int32; F2 int8 }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int8; F1 uint64; F2 string }

func main() {
    var v1 uint8 = 0
    var v2 uint32 = 2147483648
    var v3 int8 = -1
    v4 := []string{}
    var v5 uint32 = v2
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: 4294967295, F1: 255}
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: 18446744073709551615, F1: 0}
            var v10 int = (fi8 * (fi8 ^ fi6))
            v11 := GenS2{F0: (v1 ^ v1), F1: uint32(v1)}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v12 float32 = 1.5
    switch 370 {
    case -59:
        fmt.Println(v1, v2, v12)
    case 1:
        switch 502 {
        case -9223372036854775808:
            var v13 int64 = -27
            _ = v13
        default:
            v14 := GenS1{F0: 862, F1: 430}
            _ = v14
        }
    case -496:
        var v15 int64 = -245
        _ = v15
    }
    var v16 float64 = 1.5
    v17 := GenS3{F0: (v1 % ((128) | 1)), F1: -375, F2: (v3 << ((42) & 63))}
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 int8 = v3
        var v20 uint16 = 1
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 int = (^(fi18))
            _ = fi21
            _ = v22
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    var v23 string = "hello"
    switch ("hello" + "\n\t") {
    case "\n\t":
        var v24 string = ((("a" + ("" + v23)) + v23) + v23)
        _ = v24
    case "foo bar":
        if (true || (1e10 >= 0.1)) {
            var v25 uint = 42
            var v26 uint8 = v1
            v4 = append(v4, "hello")
            _ = v25
            _ = v26
        } else {
            switch v23 {
            case "a":
                var v27 int16 = 1000
                _ = v27
            case "hello":
                var v28 int32 = 42
                _ = v28
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            switch ("foo bar" + v23) {
            case "a":
                v3 += v3
            case "\n\t":
                var v29 float64 = float64(float32(-2.25))
                _ = v29
            case "hello":
                var v30 uint8 = v1
                _ = v30
            }
            if len(v4) > 0 {
                v4[0] = v23
            }
        }
    case "a":
        if (1 != -9223372036854775808) {
            for fi31 := 0; fi31 < 7; fi31++ {
                v5 = 1
                v32 := GenS4{F0: ("\n\t" + ((v23 + "hello") + v23)), F1: fi31}
                if true {
                    v1 -= v1
                    var v33 uint32 = 2147483647
                    _ = v33
                } else {
                    v34 := GenS5{F0: 42, F1: 0}
                    v35 := GenS6{F0: (v3 / ((100) | 1)), F1: 18446744073709551615, F2: "\n\t"}
                    _ = v34
                    _ = v35
                }
                _ = fi31
                _ = v32
            }
        }
    }
    if false {
        var v36 int = 1
        v37 := GenS6{F0: v3, F1: 2, F2: "\n\t"}
        v4 = append(v4, "\n\t")
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        v16 = float64(v12)
        if false {
            var v39 int16 = -223
            _ = v39
        }
        v1 -= (v1 / (((v1 >> ((42) & 63))) | 1))
        _ = fi38
    }
    v5 = (v5 % (((v5 << ((42) & 63))) | 1))
    var v40 uint32 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
