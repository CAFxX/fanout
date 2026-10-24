package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 string }

func main() {
    var v1 bool = true
    v2 := []int{-9223372036854775808, 9223372036854775807, 2147483647, 9223372036854775807}
    v3 := make(map[string]string)
    var v4 int = 128
    v5 := []string{(("" + "hello") + "foo bar"), "", "hello"}
    var v6 uint32 = 689
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v7 uint = 368
    v7 += v7
    fmt.Println(v6, v1, v4)
    switch v4 {
    case 1:
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 int16 = 48
            switch 2147483647 {
            case 1:
                for fi10 := 0; fi10 < 15; fi10++ {
                    v11 := GenS1{F0: -128, F1: 0}
                    _ = fi10
                    _ = v11
                }
            case 127:
                v12 := GenS2{F0: ""}
                _ = v12
            case 994:
                var v13 bool = (!v1)
                _ = v13
            default:
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
            }
            _ = fi8
            _ = v9
        }
    case -676:
        var v14 int32 = -382
        _ = v14
    case 0:
        var v15 uint8 = 126
        _ = v15
    default:
        fmt.Println(v7, v4)
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint32 = 2147483648
        if len(v5) > 0 {
            v5[0] = ("hello" + "foo bar")
        }
        if ((-663 <= 1) && v1) {
            v18 := GenS2{F0: ""}
            _ = v18
        }
        _ = fi16
        _ = v17
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        if v1 {
            var v20 float32 = 1e10
            _ = v20
        }
        v3["hello"] = "hello"
        _ = fi19
    }
    var v21 string = ("foo bar" + "")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
