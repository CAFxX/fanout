package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint8 = 1
    var v2 int8 = 110
    var v3 int = -9223372036854775808
    var v4 uint = 450
    v5 := []string{("\n\t" + "")}
    var v6 int32 = 42
    if (uint16(v1) < 2) {
        fmt.Println(v4)
        v2 += v2
        var v7 int32 = (v6 & (v6 ^ v6))
        _ = v7
    } else {
        fmt.Println(v6, v1)
        v8 := GenS1{F0: 1, F1: 1}
        _ = v8
    }
    switch "\n\t" {
    case "":
        switch (("hello" + "foo bar") + "\n\t") {
        case "":
            v9 := GenS2{F0: 32768}
            _ = v9
        default:
            var v10 bool = false
            _ = v10
        }
    default:
        switch v3 {
        case 32768:
            if len(v5) > 0 {
                v5[0] = "hello"
            }
        }
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            v5 = append(v5, "hello")
            v3 -= (-(v3))
            _ = fi12
        }
        _ = fi11
    }
    var v13 int32 = (v6 + (-(v6)))
    switch "" {
    case "":
        v4 -= (v4 + (v4 | 65535))
    }
    switch "\n\t" {
    case "foo bar":
        if (!true) {
            v14 := GenS1{F0: 4294967295, F1: 42}
            _ = v14
        } else {
            var v15 string = "foo bar"
            var v16 uint16 = 1
            for fi17 := 0; fi17 < 14; fi17++ {
                v5 = append(v5, v15)
                var v18 bool = false
                fmt.Println(v2)
                _ = fi17
                _ = v18
            }
            _ = v15
            _ = v16
        }
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "a":
        v19 := GenS1{F0: 2147483648, F1: -1}
        _ = v19
    }
    v5 = append(v5, "\n\t")
    v2 -= (v2 / (((-(v2))) | 1))
    v5 = append(v5, ("hello" + "\n\t"))
    v3 = -890
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v2 += (^(v2))
    v2 = (v2 * -77)
    v6 = 1
    var v20 float64 = (1.5 * 1.5)
    fmt.Println(v6, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
