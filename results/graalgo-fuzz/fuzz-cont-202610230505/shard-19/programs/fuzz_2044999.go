package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int8; F1 uint64; F2 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    var v4 int32 = 468
    var v5 int16 = 1000
    var v6 uint64 = 42
    switch "a" {
    case "hello":
        v7 := v3[4294967295]
        _ = v7
    }
    switch 128 {
    case 42:
        for fi8 := 0; fi8 < 4; fi8++ {
            if (42 == 102) {
                var v9 int32 = 1
                v9 -= (v4 >> ((0) & 63))
                _ = v9
            }
            _ = fi8
        }
    case 217:
        v3[4294967295] = 65535
    case 9223372036854775807:
        switch "\n\t" {
        case "a":
            v4 -= 2147483647
        case "\n\t":
            var v10 int64 = -438
            _ = v10
        default:
            v3[2147483647] = 411
        }
    }
    if true {
        if true {
            switch "" {
            case "hello":
                v4 += v4
            case "":
                v6 = v6
            case "foo bar":
                v4 -= (v4 % (((-(v4))) | 1))
            }
        }
        v5 = v5
    }
    v4 -= (v4 * 370)
    v4 = (v4 << ((4294967295) & 63))
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int16 = v5
        _ = fi11
        _ = v12
    }
    var v13 uint16 = 924
    switch 9223372036854775807 {
    case -2147483648:
        switch "" {
        case "foo bar":
            v14 := GenS1{F0: 0, F1: 0}
            _ = v14
        }
    default:
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 int8 = 0
            fmt.Println(v16)
            v5 += -1
            _ = fi15
            _ = v16
        }
    }
    var v17 float64 = 1e10
    v18 := GenS1{F0: 8, F1: 255}
    if (v5 > v5) {
        if (0.1 <= (0.1 + 1.5)) {
            fmt.Println(v13, v4, v5)
        } else {
            v19 := []int{-2147483648, -86}
            for fi20 := 0; fi20 < 8; fi20++ {
                v6 = v6
                _ = fi20
            }
            var v21 uint8 = 127
            _ = v19
            _ = v21
        }
        switch "foo bar" {
        case "a":
            switch "\n\t" {
            case "\n\t":
                _ = len(v1)
            default:
                if true {
                    var v22 string = "hello"
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v6 = v6
                        v4 = v4
                        fmt.Println(v13)
                        _ = fi23
                    }
                    _ = v22
                }
            }
        case "\n\t":
            for fi24 := 0; fi24 < 12; fi24++ {
                v25 := GenS1{F0: -1, F1: 18446744073709551615}
                v26 := []int{}
                for fi27 := 0; fi27 < 20; fi27++ {
                    var v28 int = 908
                    v26 = append(v26, 1)
                    v28 += (-(v28))
                    _ = fi27
                    _ = v28
                }
                _ = fi24
                _ = v25
                _ = v26
            }
        }
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        switch "" {
        case "\n\t":
            for fi29 := 0; fi29 < 5; fi29++ {
                v30 := GenS2{F0: -1, F1: (^(v6)), F2: 1}
                for fi31 := 0; fi31 < 2; fi31++ {
                    v17 += (v17 / v17)
                    v5 = v5
                    v32 := []int{fi31, -848, fi29, fi29}
                    _ = fi31
                    _ = v32
                }
                _ = fi29
                _ = v30
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
