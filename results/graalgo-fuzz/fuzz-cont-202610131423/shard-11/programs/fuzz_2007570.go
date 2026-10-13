package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 string }
func genFn0() string {
    var v4 int = 128
    var v5 int32 = -103
    v4 -= v4
    if false {
        switch v4 {
        case -1:
            var v6 uint64 = 4294967295
            _ = v6
        case 128:
            var v7 uint = 513
            _ = v7
        }
        v8 := []int{-2147483648, 127}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = v8
    }
    _ = v4
    _ = v5
    return "a"
}

func main() {
    var v1 string = "foo bar"
    v2 := []int{127, 32768, 42}
    var v3 uint = 65535
    v3 = v3
    var v9 uint16 = 2
    if (4294967295 != 18446744073709551615) {
        switch 4294967295 {
        case -2147483648:
            var v10 uint = (v3 ^ (v3 * (^(v3))))
            _ = v10
        case 32768:
            for fi11 := 0; fi11 < 6; fi11++ {
                fmt.Println(fi11, fi11, v9)
                _ = fi11
            }
        }
        fmt.Println(v9)
    }
    v12 := GenS1{F0: 0, F1: 127, F2: v1}
    v1 = (v1 + v1)
    fmt.Println(v3)
    fmt.Println(v3, v1, v9)
    switch ((v1 + v1) + (("\n\t" + ((v1 + ((v1 + v1) + v1)) + ((("foo bar" + v1) + "") + (v1 + v1)))) + ("a" + "hello"))) {
    case "hello":
        var v13 uint8 = 36
        _ = v13
    case "a":
        for fi14 := 0; fi14 < 19; fi14++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v15 int16 = -950
            _ = fi14
            _ = v15
        }
    default:
        v3 -= (v3 >> (((v3 & (^(v3)))) & 63))
    }
    switch "" {
    case "foo bar":
        var v16 int64 = 0
        _ = v16
    case "a":
        fmt.Println(v3, v3)
    }
    var v17 int = 2147483647
    v1 = v1
    _ = len(v2)
    if (-512 < 1) {
        v18 := GenS1{F0: 0, F1: 0, F2: "\n\t"}
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
