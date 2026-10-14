package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -194
    var v3 int16 = int16(v2)
    v4 := make(map[string]string)
    v5 := []uint32{}
    switch 4294967295 {
    case -2147483648:
        var v6 uint64 = 966
        _ = v6
    case 128:
        var v7 uint = 0
        _ = v7
    }
    v2 = (v2 << ((1) & 63))
    switch -459 {
    case 511:
        v8 := v1[9223372036854775807]
        _ = v8
    case 128:
        v9 := GenS1{F0: (199 > 0), F1: 926}
        _ = v9
    case 2147483647:
        v5 = append(v5, 42)
    }
    var v10 int32 = 2
    var v11 uint8 = 1
    switch (("\n\t" + "\n\t") + ("foo bar" + (("" + "") + (("" + "foo bar") + "hello")))) {
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "\n\t":
        switch 42 {
        case 516:
            v12 := GenS1{F0: false, F1: 65535}
            _ = v12
        case 9223372036854775807:
            var v13 bool = false
            _ = v13
        }
    case "a":
        fmt.Println(v2, v11, v3)
    default:
        switch (("" + "hello") + (("\n\t" + ("foo bar" + "a")) + "hello")) {
        case "\n\t":
            if true {
                switch "" {
                case "foo bar":
                    _ = len(v1)
                default:
                    _ = len(v5)
                }
                v10 -= int32(v2)
                fmt.Println(v10, v10)
            }
        case "a":
            v3 -= (v3 << ((18446744073709551615) & 63))
        }
    }
    var v14 int16 = (^(v3))
    v10 = (v10 + (v10 ^ 1))
    fmt.Println(v11)
    for fi15 := 0; fi15 < 1; fi15++ {
        if false {
            if len(v5) > 0 {
                v5[0] = 1
            }
            v10 = v10
        }
        _ = fi15
    }
    fmt.Println(v3)
    if true {
        var v16 int8 = int8(v10)
        _ = v16
    }
    fmt.Println(v11, v10)
    var v17 int64 = v2
    fmt.Println(v10, v2, v3)
    var v18 float64 = 3.14159
    v14 -= 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
