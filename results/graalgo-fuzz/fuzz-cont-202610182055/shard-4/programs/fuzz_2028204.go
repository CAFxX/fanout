package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 147
    v3 := make(map[string]bool)
    var v4 bool = false
    delete(v3, "")
    switch 4294967295 {
    case 2147483647:
        if v4 {
            v2 = v2
        } else {
            fmt.Println(v4, v4)
        }
    case 127:
        v5 := []int{9223372036854775807, 1, 32767, int(v2)}
        _ = v5
    }
    if true {
        v3["\n\t"] = (32767 == 427)
        var v6 uint8 = 127
        _ = v6
    } else {
        var v7 bool = (0 > 2)
        _ = v7
    }
    var v8 float64 = 1.5
    v9 := []int{128, -548, 32767}
    var v10 int16 = 32767
    if len(v9) > 0 {
        v9[0] = 0
    }
    v11 := GenS1{F0: 0}
    var v12 bool = (!true)
    var v13 int = -521
    switch "\n\t" {
    case "\n\t":
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int32 = 0
            _ = fi14
            _ = v15
        }
    default:
        switch "" {
        case "foo bar":
            v13 = v13
        }
    }
    fmt.Println(v12, v8, v10)
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 uint32 = 858
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
