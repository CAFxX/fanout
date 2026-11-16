package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 string }

func main() {
    var v1 int8 = -107
    v2 := make(map[uint32]string)
    v3 := []int{-9223372036854775808}
    if false {
        if true {
            var v4 bool = (18446744073709551615 <= 255)
            for fi5 := 0; fi5 < 4; fi5++ {
                if (100 == 2147483648) {
                    fmt.Println(fi5, fi5, v4)
                } else {
                    switch (fi5 - (fi5 / ((fi5) | 1))) {
                    case -9223372036854775808:
                        var v6 int32 = 1
                        _ = v6
                    }
                    v1 += int8(fi5)
                }
                _ = fi5
            }
            _ = v4
        } else {
            v1 = -97
            var v7 uint16 = 2
            v7 = (v7 * (v7 >> ((1) & 63)))
            _ = v7
        }
        v1 = (v1 | v1)
        var v8 int8 = 100
        _ = v8
    } else {
        fmt.Println(v1)
        for fi9 := 0; fi9 < 19; fi9++ {
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 int16 = -1
                _ = fi10
                _ = v11
            }
            v12 := GenS1{F0: v1, F1: 123, F2: "foo bar"}
            _ = fi9
            _ = v12
        }
        v1 = v1
    }
    v1 -= v1
    v1 += -3
    var v13 float32 = 0.1
    switch "hello" {
    case "\n\t":
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 int32 = -2147483648
            v15 += v15
            _ = fi14
            _ = v15
        }
    case "foo bar":
        v1 = v1
    case "a":
        v13 += v13
    }
    var v16 uint64 = 42
    var v17 float32 = (float32(-2.25) * (v13 / (v13 + v13)))
    fmt.Println(v16, v16, v16)
    var v18 uint8 = 127
    v13 = v13
    v2[238] = "foo bar"
    var v19 uint16 = 160
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
