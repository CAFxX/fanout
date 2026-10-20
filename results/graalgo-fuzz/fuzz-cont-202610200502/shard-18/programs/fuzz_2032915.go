package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 string; F1 int16; F2 uint }
func genFn0(p0 string) uint8 {
    fmt.Println(p0)
    v4 := GenS1{F0: -2147483648, F1: 0, F2: 42}
    v5 := []int{2147483647, 42, -1}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, 127)
    _ = p0
    _ = v4
    _ = v5
    return 30
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 1
    var v3 int64 = 781
    for fi6 := 0; fi6 < 13; fi6++ {
        if false {
            v7 := GenS1{F0: (fi6 ^ fi6), F1: 4294967295, F2: 255}
            switch ("foo bar" + "foo bar") {
            case "hello":
                v2 += (v2 ^ 125)
            case "\n\t":
                v3 += v3
            case "a":
                v8 := GenS1{F0: 168, F1: 1, F2: 127}
                _ = v8
            }
            v9 := []int{127, (fi6 | (fi6 ^ (fi6 % ((-2147483648) | 1)))), 9223372036854775807}
            _ = v7
            _ = v9
        } else {
            var v10 int32 = 999
            v2 = (v2 >> ((0) & 63))
            _ = v10
        }
        _ = fi6
    }
    switch "\n\t" {
    case "hello":
        var v11 int64 = v3
        _ = v11
    case "\n\t":
        var v12 float32 = 0.0
        _ = v12
    default:
        v13 := GenS1{F0: 1, F1: 41, F2: 42}
        _ = v13
    }
    var v14 int32 = 2147483647
    fmt.Println(v2)
    var v15 uint = 42
    var v16 int32 = (v14 | int32(v3))
    if ((18446744073709551615 != 4294967295) && true) {
        v15 = v15
        v15 += (v15 * (v15 - 667))
        if false {
            for fi17 := 0; fi17 < 18; fi17++ {
                if false {
                    v18 := GenS2{F0: (v15 * (^(v15))), F1: (^(v15))}
                    var v19 float64 = (-2.25 * (1.5 - 0.0))
                    _ = v18
                    _ = v19
                }
                v20 := []int{-862, (fi17 / ((fi17) | 1))}
                _ = fi17
                _ = v20
            }
            switch "a" {
            case "":
                var v21 int32 = v16
                _ = v21
            case "a":
                fmt.Println(v3)
            case "\n\t":
                for fi22 := 0; fi22 < 1; fi22++ {
                    switch "a" {
                    case "":
                        v2 -= (v2 >> (((^(v15))) & 63))
                    }
                    var v23 uint = 255
                    v16 -= v16
                    _ = fi22
                    _ = v23
                }
            }
        }
    }
    var v24 float64 = 0.1
    v25 := GenS2{F0: v15, F1: v15}
    v1[1] = 127
    if (558 == 4294967295) {
        v26 := GenS3{F0: "hello", F1: 1000, F2: 18446744073709551615}
        v15 = v15
        _ = v26
    } else {
        var v27 uint8 = 149
        _ = len(v1)
        _ = v27
    }
    var v28 int = -497
    var v29 int32 = v14
    switch (("hello" + ("foo bar" + "hello")) + ("a" + "foo bar")) {
    case "a":
        var v30 float64 = (v24 / -2.25)
        _ = v30
    }
    fmt.Println(v24, v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
}
