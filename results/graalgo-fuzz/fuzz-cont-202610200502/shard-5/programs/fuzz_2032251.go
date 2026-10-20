package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }

func main() {
    v1 := []uint32{2147483648, 0, 1}
    v2 := []string{"hello", "", "a"}
    v3 := make(map[uint32]uint32)
    if true {
        switch "" {
        case "hello":
            var v4 uint16 = 1
            _ = v4
        case "":
            var v5 uint8 = 147
            _ = v5
        case "foo bar":
            v6 := v3[2147483648]
            _ = v6
        }
        if (false && true) {
            var v7 bool = false
            switch "hello" {
            case "hello":
                var v8 uint64 = 2
                _ = v8
            case "":
                v7 = false
            }
            _ = v7
        } else {
            var v9 uint32 = 728
            _ = v9
        }
        var v10 int32 = 215
        _ = v10
    } else {
        var v11 int32 = -1
        var v12 uint = 1
        _ = v11
        _ = v12
    }
    switch 42 {
    case -303:
        var v13 uint = 4294967295
        _ = v13
    case 4294967295:
        delete(v3, 2147483648)
    case 127:
        if ((542 <= 255) && false) {
            var v14 int = 127
            if true {
                var v15 int = 1
                v15 = 91
                _ = v15
            } else {
                var v16 float64 = float64(float32(0.1))
                v14 -= v14
                _ = v16
            }
            _ = v14
        } else {
            v17 := v3[358]
            v17 = v17
            _ = v17
        }
    default:
        switch 42 {
        case 968:
            switch 489 {
            case 747:
                switch 32767 {
                case 526:
                    var v18 float64 = float64((0.1 - -2.25))
                    _ = v18
                case 4294967295:
                    v19 := GenS1{F0: 2, F1: 0}
                    _ = v19
                }
            case -864:
                var v20 int64 = 42
                _ = v20
            case 127:
                switch -593 {
                case 32768:
                    var v21 int = 127
                    _ = v21
                case 32767:
                    var v22 int = 275
                    _ = v22
                case 35:
                    var v23 uint64 = 0
                    _ = v23
                }
            default:
                var v24 int64 = 1
                _ = v24
            }
        }
    }
    var v25 uint64 = 4294967295
    v25 = (^(v25))
    v25 += 2
    fmt.Println(v25)
    if (42 != 0) {
        v26 := v3[2147483647]
        v3[v26] = (v26 % (((v26 >> ((65535) & 63))) | 1))
        _ = v26
    } else {
        v1 = append(v1, 0)
        v2 = append(v2, "foo bar")
    }
    var v27 int = 9223372036854775807
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        if true {
            var v28 int64 = 0
            _ = v28
        }
    case "\n\t":
        v25 -= 294
    case "":
        var v29 int8 = 1
        _ = v29
    }
    v30 := GenS1{F0: 18446744073709551615, F1: 2}
    var v31 uint16 = 1
    var v32 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
