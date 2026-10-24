package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }

func main() {
    var v1 int8 = 100
    v2 := make(map[string]uint32)
    var v3 int8 = v1
    var v4 int16 = 1
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 bool = ((v3 != v3) || true)
        switch int(v4) {
        case -2147483648:
            v6 = (1 < 2)
        default:
            var v7 string = "foo bar"
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int32 = 42
        _ = fi8
        _ = v9
    }
    var v10 int8 = (v1 & (^(v1)))
    fmt.Println(v4, v4, v10)
    fmt.Println(v10, v1)
    v11 := GenS1{F0: 1, F1: (-1 != 1)}
    var v12 uint = 245
    fmt.Println(v3)
    v12 = v12
    if (-1 != 42) {
        var v13 bool = false
        switch -795 {
        case -2147483648:
            fmt.Println(v4, v10)
        case 645:
            switch "foo bar" {
            case "foo bar":
                delete(v2, ("\n\t" + "foo bar"))
            case "hello":
                v14 := GenS1{F0: 497, F1: v13}
                _ = v14
            default:
                var v15 int8 = 80
                _ = v15
            }
        case 442:
            var v16 uint8 = 124
            _ = v16
        default:
            if (42 > -39) {
                v3 += (v1 - (v10 & (v1 - v1)))
            } else {
                if true {
                    if v13 {
                        v1 = v1
                    } else {
                        var v17 int64 = 9223372036854775807
                        for fi18 := 0; fi18 < 20; fi18++ {
                            var v19 string = (("hello" + "") + "foo bar")
                            var v20 float64 = ((-2.25 / 0.1) * (1.5 - 0.1))
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        _ = v17
                    }
                }
                fmt.Println(v3)
            }
        }
        v12 += (v12 * v12)
        _ = v13
    } else {
        v4 += (v4 << ((v12) & 63))
        var v21 string = "\n\t"
        _ = v21
    }
    v10 = int8(v4)
    v3 = (v3 / ((v10) | 1))
    v12 += 223
    v22 := []int{32768, 42, -2147483648, -9223372036854775808}
    var v23 int = 4294967295
    if len(v22) > 0 {
        v22[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
