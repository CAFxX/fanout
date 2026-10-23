package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint32 }
func genFn0(p0 int32) int16 {
    var v7 int64 = -9223372036854775808
    if ((0.0 < ((0.0 / 3.14159) / 0.1)) && (false && (false && true))) {
        fmt.Println(p0)
    } else {
        v7 = -9223372036854775808
        if ((!(!(!(!(1.5 >= 0.1))))) || (2 < 2147483648)) {
            p0 += p0
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 bool = true
                _ = fi8
                _ = v9
            }
        } else {
            switch "" {
            case "\n\t":
                v10 := []int{1, 128}
                _ = v10
            case "hello":
                for fi11 := 0; fi11 < 1; fi11++ {
                    v7 = 1
                    _ = fi11
                }
            }
            v12 := GenS1{F0: -9223372036854775808, F1: "\n\t", F2: 0}
            var v13 uint = 649
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = v7
    return -32768
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    var v4 uint8 = 128
    var v5 uint32 = 916
    var v6 uint64 = 741
    var v14 int8 = 121
    v5 -= (v5 | v5)
    var v15 uint16 = 32767
    v14 = v14
    v16 := []int{}
    var v17 bool = (32767 >= 32767)
    _ = len(v1)
    var v18 int8 = v14
    v6 = v6
    if ((!(!v17)) || v17) {
        var v19 float32 = 1e10
        v20 := GenS1{F0: 373, F1: ("" + "foo bar"), F2: v5}
        fmt.Println(v5, v18)
        _ = v19
        _ = v20
    } else {
        switch "\n\t" {
        case "\n\t":
            var v21 string = "hello"
            _ = v21
        default:
            var v22 int = -2147483648
            _ = v22
        }
        v23 := GenS1{F0: -879, F1: "\n\t", F2: 2147483647}
        if v17 {
            fmt.Println(v6)
        }
        _ = v23
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 uint16 = (v15 / ((32767) | 1))
        _ = fi24
        _ = v25
    }
    v15 = (v15 + 1)
    var v26 float64 = 0.1
    v27 := v1[(^(v5))]
    var v28 uint64 = (v6 | v6)
    switch "a" {
    case "a":
        var v29 uint = 361
        _ = v29
    case "foo bar":
        switch 1 {
        case 258:
            v18 = -1
        case 42:
            fmt.Println(v6, v4)
        case 2147483647:
            v30 := GenS1{F0: -795, F1: "", F2: (v5 ^ (^(v27)))}
            _ = v30
        }
    }
    v16 = append(v16, -1)
    for fi31 := 0; fi31 < 5; fi31++ {
        switch fi31 {
        case 237:
            for fi32 := 0; fi32 < 4; fi32++ {
                var v33 uint16 = uint16(v4)
                var v34 uint16 = (v15 & (v15 % ((v15) | 1)))
                _ = fi32
                _ = v33
                _ = v34
            }
        case 32768:
            v28 = v6
        case 0:
            var v35 uint32 = (v27 ^ v5)
            _ = v35
        default:
            if (!false) {
                fmt.Println(v27, v4)
            } else {
                v36 := GenS1{F0: fi31, F1: "", F2: v5}
                fmt.Println(v14)
                v18 -= (v18 * (v18 << ((65535) & 63)))
                _ = v36
            }
        }
        var v37 int64 = 9223372036854775807
        _ = fi31
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
