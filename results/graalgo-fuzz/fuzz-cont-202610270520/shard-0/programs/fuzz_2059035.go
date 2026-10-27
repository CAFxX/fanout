package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int64; F1 int8; F2 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 97
    var v3 string = ""
    v4 := make(map[string]bool)
    fmt.Println(v3, v2)
    v5 := []int{32767, 42, 2147483647}
    var v6 int8 = -18
    switch v3 {
    case "a":
        v5 = append(v5, 228)
    }
    var v7 bool = true
    v6 += v6
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: 18446744073709551615, F1: 255, F2: 0}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi8
        _ = v9
    }
    if (!v7) {
        for fi10 := 0; fi10 < 3; fi10++ {
            switch v3 {
            case "hello":
                var v11 uint64 = 1
                _ = v11
            case "a":
                switch v3 {
                case "a":
                    v3 = "a"
                case "\n\t":
                    v12 := GenS2{F0: (v6 - 111), F1: fi10}
                    _ = v12
                }
            }
            _ = len(v5)
            _ = fi10
        }
        _ = len(v5)
        v13 := GenS3{F0: int16(v2)}
        _ = v13
    }
    v2 -= v2
    var v14 string = (("" + (v3 + "foo bar")) + v3)
    if ((v7 && (!((-2.25 < 3.14159) && ((v2 + 2147483648) >= (v2 * 0))))) && v7) {
        v6 -= v6
        switch 0 {
        case 32768:
            _ = len(v4)
        case -2147483648:
            switch "a" {
            case "foo bar":
                fmt.Println(v2, v2)
            default:
                switch v14 {
                case "hello":
                    var v15 int = 360
                    _ = v15
                case "\n\t":
                    if v7 {
                        v16 := GenS2{F0: v6, F1: 9223372036854775807}
                        var v17 int16 = 288
                        _ = v16
                        _ = v17
                    }
                default:
                    for fi18 := 0; fi18 < 17; fi18++ {
                        fmt.Println(v7, v3, v14)
                        _ = fi18
                    }
                }
            }
        }
    }
    var v19 int8 = -128
    if v7 {
        v19 += 1
        var v20 uint64 = 92
        if (-39 == -1) {
            v21 := GenS4{F0: v20, F1: 942}
            var v22 float64 = 0.1
            _ = v21
            _ = v22
        }
        _ = v20
    } else {
        v3 = ((("" + "foo bar") + ((v3 + v14) + (v3 + v3))) + (v14 + v3))
    }
    if (729 <= -32768) {
        var v23 int = int(v2)
        _ = v23
    } else {
        v24 := GenS5{F0: -9223372036854775808, F1: v6, F2: 4294967295}
        _ = v24
    }
    var v25 bool = v7
    v3 = "foo bar"
    var v26 float64 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
