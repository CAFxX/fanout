package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint32; F1 string }
func genFn0(p0 uint16, p1 int8, p2 uint32) int16 {
    v6 := []int{174, 2147483647}
    var v7 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return int16(p1)
}
func genFn1(p0 int64, p1 uint) int64 {
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 int8 = -87
        v9 = v9
        var v10 uint64 = 477
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{-315, -2147483648, 9223372036854775807, 32768}
    _ = p0
    _ = p1
    _ = v11
    return (p0 * (p0 + p0))
}

func main() {
    var v1 uint64 = 988
    v2 := make(map[uint32]bool)
    var v3 uint16 = 65535
    var v4 uint16 = v3
    var v5 int8 = 127
    if (!true) {
        v12 := GenS1{F0: v1, F1: 43, F2: 332}
        _ = v12
    }
    var v13 float64 = (3.14159 + 3.14159)
    if true {
        var v14 bool = true
        _ = v14
    }
    fmt.Println(v13, v5)
    if ((v4 < (v3 + v3)) || false) {
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 string = "\n\t"
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 uint = 540
                _ = fi17
                _ = v18
            }
            _ = fi15
            _ = v16
        }
        _ = len(v2)
        v19 := []int{390, -211, -673}
        _ = v19
    } else {
        fmt.Println(v3, v13, v13)
    }
    if true {
        if false {
            var v20 bool = ((v3 >> ((255) & 63)) == 0)
            fmt.Println(v4, v1)
            switch "foo bar" {
            case "":
                var v21 uint32 = 1
                _ = v21
            case "a":
                fmt.Println(v4, v1, v4)
            case "\n\t":
                var v22 uint8 = 127
                _ = v22
            }
            _ = v20
        }
        var v23 int16 = -849
        _ = v23
    } else {
        fmt.Println(v4, v4)
        var v24 uint32 = 2147483647
        v25 := []int{}
        _ = v24
        _ = v25
    }
    v26 := GenS1{F0: v1, F1: (v1 - v1), F2: v1}
    fmt.Println(v5, v3, v5)
    delete(v2, 2)
    var v27 float32 = 3.14159
    fmt.Println(v27, v27)
    if false {
        v28 := []int{-453, 2147483647, 128}
        for fi29 := 0; fi29 < 9; fi29++ {
            fmt.Println(v5, v13)
            fmt.Println(v1, v13)
            _ = fi29
        }
        _ = len(v28)
        _ = v28
    }
    if (0 > genFn0(492, (v5 << ((2) & 63)), 266)) {
        fmt.Println(v4)
        switch 916 {
        case 128:
            if (!(297 > 574)) {
                fmt.Println(v3)
                var v30 float32 = float32((((-2.25 + float64(((1e10 / -2.25) + (3.14159 + float32(((v13 / (v13 + -2.25)) - 3.14159)))))) / (0.0 + v13)) * (v13 / -2.25)))
                _ = v30
            } else {
                var v31 uint64 = v1
                for fi32 := 0; fi32 < 4; fi32++ {
                    v33 := GenS2{F0: 2147483648, F1: "a"}
                    _ = fi32
                    _ = v33
                }
                var v34 float64 = 1e10
                _ = v31
                _ = v34
            }
        case -880:
            var v35 int = 0
            _ = v35
        case -9223372036854775808:
            var v36 uint64 = 94
            _ = v36
        }
    }
    v37 := GenS1{F0: (v1 - (v1 >> ((211) & 63))), F1: uint64(v3), F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
