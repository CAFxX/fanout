package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 bool = ((false || ("\n\t" == "")) || (-45 < -125))
    var v3 int64 = 761
    v4 := make(map[string]bool)
    var v5 int16 = -1
    v6 := []string{("foo bar" + "foo bar"), ("hello" + "\n\t")}
    fmt.Println(v5, v5, v2)
    for fi7 := 0; fi7 < 2; fi7++ {
        if len(v6) > 0 {
            v6[0] = ""
        }
        v5 -= v5
        switch "hello" {
        case "a":
            var v8 uint32 = 4294967295
            _ = v8
        case "hello":
            fmt.Println(v5, v5)
        }
        _ = fi7
    }
    var v9 float64 = 0.0
    fmt.Println(v2)
    switch "a" {
    case "a":
        v3 -= (v3 >> ((4294967295) & 63))
    default:
        switch "foo bar" {
        case "hello":
            var v10 int64 = -9223372036854775808
            _ = v10
        case "foo bar":
            fmt.Println(v2, v9)
        }
    }
    fmt.Println(v9, v3)
    if false {
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 int8 = -92
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 8; fi14++ {
                v15 := GenS1{F0: 1}
                _ = fi14
                _ = v15
            }
            _ = fi11
        }
        v16 := GenS2{F0: 112, F1: 1, F2: 42}
        _ = v16
    }
    var v17 int16 = (v5 / ((v5) | 1))
    v17 += -54
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int8 = 81
        var v20 bool = (42 == 32767)
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
