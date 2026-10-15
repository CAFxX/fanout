package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]bool)
    var v3 int = 682
    v4 := []string{"", ("\n\t" + "foo bar"), "hello", ("foo bar" + "hello")}
    var v5 int64 = 94
    for fi6 := 0; fi6 < 11; fi6++ {
        v4 = append(v4, "a")
        _ = fi6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := v2[846]
        v9 := GenS1{F0: -128, F1: 4294967295, F2: fi7}
        var v10 uint64 = 0
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS1{F0: 0, F1: 0, F2: (v3 - (fi11 << ((42) & 63)))}
        fmt.Println(v5, v1, v3)
        v1 = float32(3.14159)
        _ = fi11
        _ = v12
    }
    fmt.Println(v1, v5)
    var v13 string = "foo bar"
    var v14 string = v13
    var v15 int32 = 813
    v16 := v2[0]
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 float32 = (1e10 / -2.25)
        var v19 int = fi17
        _ = fi17
        _ = v18
        _ = v19
    }
    v2[673] = (42 <= 42)
    var v20 int = 0
    var v21 int = v20
    v3 = v21
    v3 = v3
    switch (v13 + v14) {
    case "a":
        var v22 uint32 = 2147483647
        _ = v22
    case "foo bar":
        v5 = v5
    case "\n\t":
        for fi23 := 0; fi23 < 10; fi23++ {
            for fi24 := 0; fi24 < 12; fi24++ {
                var v25 uint8 = 42
                var v26 uint16 = 2
                switch (v3 & (fi23 / (((fi23 / (((v20 % (((fi23 << ((2) & 63))) | 1))) | 1))) | 1))) {
                case -9223372036854775808:
                    v26 -= (v26 + v26)
                case -829:
                    v2[42] = ((!false) || (!v16))
                case 42:
                    v26 -= uint16(v25)
                default:
                    var v27 float32 = float32(0.0)
                    _ = v27
                }
                _ = fi24
                _ = v25
                _ = v26
            }
            for fi28 := 0; fi28 < 4; fi28++ {
                fmt.Println(v21, v1, v20)
                var v29 int = fi28
                var v30 uint32 = 123
                _ = fi28
                _ = v29
                _ = v30
            }
            v13 = ""
            _ = fi23
        }
    }
    v16 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
