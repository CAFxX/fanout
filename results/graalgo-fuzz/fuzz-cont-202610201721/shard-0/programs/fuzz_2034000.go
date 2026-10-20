package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int16 }
type GenS2 struct { F0 int; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int64 }

func main() {
    var v1 int16 = 0
    v2 := []int{127, 0, 668}
    v3 := []string{}
    var v4 uint8 = 29
    var v5 string = (("" + "") + "hello")
    var v6 uint16 = 65535
    var v7 int8 = -128
    if len(v3) > 0 {
        v3[0] = ((v5 + v5) + (v5 + (("" + v5) + v5)))
    }
    switch v5 {
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "hello":
        for fi8 := 0; fi8 < 4; fi8++ {
            _ = len(v2)
            fmt.Println(v7)
            _ = fi8
        }
    default:
        var v9 uint16 = v6
        _ = v9
    }
    _ = len(v3)
    v10 := GenS1{F0: -1, F1: 255, F2: (v1 << ((2) & 63))}
    for fi11 := 0; fi11 < 8; fi11++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi11
    }
    fmt.Println(v4)
    if false {
        switch -436 {
        case 127:
            for fi12 := 0; fi12 < 7; fi12++ {
                v1 -= v1
                _ = fi12
            }
        }
    }
    fmt.Println(v1)
    _ = len(v2)
    switch (v5 + (v5 + v5)) {
    case "\n\t":
        v13 := GenS1{F0: 786, F1: 255, F2: (v1 | 32767)}
        _ = v13
    case "hello":
        v6 += v6
    case "a":
        if false {
            var v14 int32 = -286
            _ = v14
        } else {
            var v15 int16 = (^(v1))
            _ = v15
        }
    default:
        if (!(true || ((v6 ^ (^(v6))) <= (v6 & 0)))) {
            var v16 string = "foo bar"
            fmt.Println(v16, v4)
            var v17 float64 = (0.1 * (((((-2.25 + ((float64(1e10) - 3.14159) + -2.25)) + float64(0.1)) - 0.1) * 0.0) + 1e10))
            _ = v16
            _ = v17
        } else {
            var v18 uint = 0
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 int64 = int64(v18)
                _ = fi19
                _ = v20
            }
            v21 := GenS2{F0: 127, F1: 600, F2: (v4 >> (((v18 | v18)) & 63))}
            _ = v18
            _ = v21
        }
    }
    v22 := GenS1{F0: 810, F1: 65535, F2: int16(v4)}
    v23 := GenS3{F0: 2147483648, F1: -9223372036854775808}
    if len(v2) > 0 {
        v2[0] = 127
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        v25 := GenS3{F0: 42, F1: 1}
        v26 := GenS1{F0: -687, F1: 2, F2: 32767}
        fmt.Println(v4)
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        var v28 uint32 = 65535
        v3 = append(v3, ("" + (v5 + v5)))
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
