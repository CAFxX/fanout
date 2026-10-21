package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }

func main() {
    var v1 int8 = 57
    var v2 string = ("foo bar" + "hello")
    v3 := []int{-2147483648, 4294967295, -786}
    v4 := []uint32{}
    fmt.Println(v2, v1)
    v1 += (v1 << ((255) & 63))
    v5 := GenS1{F0: -2147483648, F1: 662}
    var v6 int16 = 0
    v2 = "a"
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: 1, F1: 1}
        switch (fi7 * (fi7 / (((fi7 << ((42) & 63))) | 1))) {
        case -1:
            var v9 int64 = -1
            _ = v9
        case 405:
            v4 = append(v4, 65535)
        }
        _ = fi7
        _ = v8
    }
    v10 := GenS1{F0: int32(v6), F1: 255}
    if false {
        if (false && false) {
            var v11 uint16 = 560
            var v12 float32 = 0.0
            _ = v11
            _ = v12
        }
        var v13 int8 = v1
        if false {
            fmt.Println(v1)
        } else {
            for fi14 := 0; fi14 < 2; fi14++ {
                v1 += (v1 ^ (v13 >> ((65535) & 63)))
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 uint32 = 273
                    var v17 uint16 = 769
                    _ = len(v4)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                switch v2 {
                case "hello":
                    v1 += int8(fi14)
                }
                _ = fi14
            }
            if (-23 <= 127) {
                v13 -= (v13 ^ int8(v6))
                for fi18 := 0; fi18 < 9; fi18++ {
                    v19 := GenS1{F0: -807, F1: 42}
                    _ = fi18
                    _ = v19
                }
            } else {
                var v20 float32 = (0.0 * 0.1)
                var v21 uint16 = 65535
                _ = v20
                _ = v21
            }
        }
        _ = v13
    }
    if true {
        fmt.Println(v1, v6, v1)
        var v22 int16 = 42
        switch 878 {
        case 536:
            v23 := GenS1{F0: 2147483647, F1: 1}
            _ = v23
        case 32767:
            var v24 bool = (!false)
            _ = v24
        }
        _ = v22
    }
    var v25 uint32 = 290
    v1 = v1
    var v26 int = -170
    var v27 int16 = int16(v26)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
