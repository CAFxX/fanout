package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 bool = (38 >= 207)
    var v2 uint16 = 255
    v3 := []int{898, 2147483647, 127, 4294967295}
    v4 := make(map[string]int)
    v5 := []string{"foo bar", ("hello" + "a"), ("" + "\n\t")}
    v6 := make(map[int]int)
    fmt.Println(v1, v1)
    for fi7 := 0; fi7 < 12; fi7++ {
        if len(v5) > 0 {
            v5[0] = "a"
        }
        _ = fi7
    }
    switch "\n\t" {
    case "":
        var v8 int16 = -708
        _ = v8
    }
    if false {
        fmt.Println(v2, v2)
        var v9 uint = 4294967295
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v9
    }
    var v10 int64 = 1
    v11 := GenS1{F0: 128}
    var v12 float64 = 1e10
    fmt.Println(v1, v12, v1)
    if (65535 >= 255) {
        var v13 float32 = (1e10 + 1e10)
        var v14 uint32 = 0
        _ = v13
        _ = v14
    }
    v5 = append(v5, "a")
    for fi15 := 0; fi15 < 17; fi15++ {
        v1 = (fi15 != (fi15 << ((255) & 63)))
        _ = fi15
    }
    if v1 {
        v10 -= (v10 * (v10 - v10))
        for fi16 := 0; fi16 < 17; fi16++ {
            for fi17 := 0; fi17 < 14; fi17++ {
                switch "" {
                case "foo bar":
                    v5 = append(v5, "\n\t")
                case "hello":
                    var v18 int16 = 42
                    _ = v18
                }
                var v19 uint8 = 241
                _ = fi17
                _ = v19
            }
            _ = fi16
        }
    }
    switch "" {
    case "\n\t":
        v10 += (v10 / ((v10) | 1))
    case "hello":
        fmt.Println(v10, v12, v2)
    default:
        for fi20 := 0; fi20 < 13; fi20++ {
            fmt.Println(v1)
            _ = fi20
        }
    }
    v10 += (-(v10))
    v1 = v1
    v21 := GenS1{F0: 255}
    var v22 uint16 = (v2 >> ((18446744073709551615) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
