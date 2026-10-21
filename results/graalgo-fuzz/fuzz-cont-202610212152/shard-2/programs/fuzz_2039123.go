package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 int; F1 uint8; F2 int32 }

func main() {
    var v1 int16 = 42
    var v2 float64 = 0.0
    var v3 int32 = 989
    v4 := GenS1{F0: -2147483648, F1: int64(v3)}
    for fi5 := 0; fi5 < 13; fi5++ {
        switch (fi5 >> ((625) & 63)) {
        case 128:
            v6 := GenS1{F0: (fi5 << ((42) & 63)), F1: -9223372036854775808}
            _ = v6
        case -2147483648:
            v7 := GenS1{F0: (fi5 * -125), F1: -964}
            _ = v7
        case 1:
            fmt.Println(v3)
        default:
            v8 := GenS1{F0: (fi5 & (fi5 | (^(fi5)))), F1: 1}
            _ = v8
        }
        v9 := GenS1{F0: (fi5 ^ fi5), F1: -725}
        _ = fi5
        _ = v9
    }
    switch 106 {
    case 42:
        v1 -= (-(v1))
    case 32767:
        v1 -= v1
    case 644:
        var v10 uint32 = 0
        _ = v10
    }
    var v11 uint16 = 255
    fmt.Println(v2)
    switch ("" + "") {
    case "":
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := []int{}
            switch "\n\t" {
            case "a":
                v14 := GenS2{F0: 42, F1: 26, F2: 576}
                _ = v14
            case "":
                switch -106 {
                case -2147483648:
                    v11 += (v11 + v11)
                case 127:
                    _ = len(v13)
                case 32768:
                    if len(v13) > 0 {
                        v13[0] = (fi12 << ((26) & 63))
                    }
                default:
                    for fi15 := 0; fi15 < 4; fi15++ {
                        v16 := GenS3{F0: 42, F1: false}
                        _ = fi15
                        _ = v16
                    }
                }
            default:
                v11 += v11
            }
            _ = fi12
            _ = v13
        }
    case "hello":
        switch ("a" + "foo bar") {
        case "\n\t":
            v1 -= -32768
        case "hello":
            v17 := GenS2{F0: 76, F1: -128, F2: 447}
            _ = v17
        }
    case "foo bar":
        switch 4294967295 {
        case 127:
            var v18 uint64 = 4294967295
            _ = v18
        case -848:
            v19 := GenS4{F0: 18446744073709551615, F1: v3}
            _ = v19
        case 32768:
            v20 := GenS5{F0: 127, F1: 128, F2: v3}
            _ = v20
        }
    default:
        v21 := []int{2147483647}
        _ = v21
    }
    switch "foo bar" {
    case "\n\t":
        var v22 int64 = 42
        _ = v22
    case "":
        fmt.Println(v11, v11, v3)
    }
    fmt.Println(v1)
    var v23 bool = true
    switch "\n\t" {
    case "foo bar":
        v24 := []int{}
        _ = v24
    default:
        if (389 >= 2) {
            v25 := []int{42}
            v26 := GenS5{F0: -886, F1: 42, F2: -820}
            var v27 int = 2147483647
            _ = v25
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 10; fi28++ {
                for fi29 := 0; fi29 < 10; fi29++ {
                    var v30 int8 = 100
                    _ = fi29
                    _ = v30
                }
                _ = fi28
            }
            for fi31 := 0; fi31 < 10; fi31++ {
                if (fi31 == 0) {
                    var v32 bool = (v23 || v23)
                    var v33 bool = v32
                    _ = v32
                    _ = v33
                } else {
                    var v34 uint32 = 4294967295
                    _ = v34
                }
                _ = fi31
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
