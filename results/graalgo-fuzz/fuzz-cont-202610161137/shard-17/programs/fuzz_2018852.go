package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint16; F2 string }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int8 = 100
    v2 := []uint32{42, 42}
    var v3 int = -934
    var v4 float32 = -2.25
    switch "hello" {
    case "hello":
        switch (v3 >> ((787) & 63)) {
        case 1:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 4294967295:
            switch 4294967295 {
            case 2147483647:
                v5 := GenS1{F0: 0, F1: true, F2: 1}
                _ = v5
            case 204:
                if false {
                    fmt.Println(v1, v3)
                } else {
                    var v6 float32 = float32(((1e10 - float64((0.1 * (3.14159 / v4)))) / 0.1))
                    if true {
                        var v7 bool = false
                        _ = v7
                    } else {
                        var v8 float64 = 1.5
                        _ = v8
                    }
                    v3 -= v3
                    _ = v6
                }
            }
        case 42:
            v1 += -126
        }
    case "foo bar":
        var v9 int16 = -32768
        _ = v9
    }
    v3 += 2147483647
    v3 = v3
    v10 := GenS2{F0: -469, F1: 1, F2: ""}
    v3 = 2147483647
    v11 := GenS1{F0: 4294967295, F1: (9223372036854775807 < 42), F2: 49}
    v4 = v4
    v2 = append(v2, 488)
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    _ = len(v2)
    if (!true) {
        v2 = append(v2, 831)
        if (-18 == 42) {
            v12 := GenS3{F0: 237}
            fmt.Println(v4, v1)
            _ = v12
        } else {
            var v13 float32 = 1.5
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := GenS1{F0: 42, F1: false, F2: 42}
                _ = len(v2)
                fmt.Println(v3, fi14, v4)
                _ = fi14
                _ = v15
            }
            if false {
                var v16 int = v3
                _ = v16
            } else {
                var v17 int64 = 491
                if len(v2) > 0 {
                    v2[0] = 986
                }
                _ = v17
            }
            _ = v13
        }
    } else {
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        fmt.Println(v1, v4)
        switch "foo bar" {
        case "a":
            v1 += (-(v1))
        default:
            if len(v2) > 0 {
                v2[0] = 0
            }
        }
        var v19 float32 = float32(((-2.25 / (float64(v4) + 0.1)) - 1e10))
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
