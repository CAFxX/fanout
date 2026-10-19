package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint64 }
func genFn0() string {
    v6 := GenS1{F0: 2147483647, F1: 58, F2: 18446744073709551615}
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint8 = 2
        fmt.Println(v8, v8)
        _ = fi7
        _ = v8
    }
    if (1 >= 4294967295) {
        var v9 int8 = 100
        switch "foo bar" {
        case "foo bar":
            if true {
                var v10 int64 = -9223372036854775808
                _ = v10
            } else {
                switch -2147483648 {
                case 9223372036854775807:
                    v11 := []int{-842, -1, -9223372036854775808, 127}
                    _ = v11
                default:
                    v12 := []int{}
                    _ = v12
                }
            }
        default:
            switch 9223372036854775807 {
            case 552:
                v9 -= v9
            default:
                fmt.Println(v9)
            }
        }
        _ = v9
    } else {
        var v13 string = "hello"
        v13 = v13
        _ = v13
    }
    var v14 bool = (205 > 128)
    v15 := GenS1{F0: 739, F1: 42, F2: 403}
    _ = v14
    _ = v15
    _ = v6
    return "hello"
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = -543
    var v3 int32 = -961
    v4 := []string{"a", "\n\t", (("\n\t" + "foo bar") + "\n\t"), ("" + "\n\t")}
    var v5 int32 = 42
    var v16 int32 = (^(v3))
    switch "\n\t" {
    case "\n\t":
        var v17 int64 = -1
        _ = v17
    case "a":
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 uint32 = 4294967295
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 int = fi20
                var v22 uint = 133
                _ = fi20
                _ = v21
                _ = v22
            }
            fmt.Println(fi18, v3, v19)
            _ = fi18
            _ = v19
        }
    }
    var v23 uint16 = 255
    for fi24 := 0; fi24 < 17; fi24++ {
        for fi25 := 0; fi25 < 8; fi25++ {
            v23 = 65535
            _ = fi25
        }
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 float32 = 0.1
            _ = fi26
            _ = v27
        }
        _ = fi24
    }
    v1["\n\t"] = 9223372036854775807
    fmt.Println(v5, v16, v23)
    v2 += 0
    v28 := GenS1{F0: 42, F1: 243, F2: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
