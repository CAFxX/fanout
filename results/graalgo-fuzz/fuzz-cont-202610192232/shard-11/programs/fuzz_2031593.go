package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint; F1 int8; F2 int8 }

func main() {
    var v1 uint8 = 42
    var v2 int32 = -583
    var v3 uint32 = 65535
    var v4 uint = 678
    var v5 uint8 = 45
    v6 := []int{127, 4294967295}
    if (v4 == v4) {
        var v7 float32 = 3.14159
        _ = v7
    } else {
        if len(v6) > 0 {
            v6[0] = 753
        }
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 string = "a"
            v4 = (v4 & (v4 - v4))
            var v11 string = ""
            _ = fi9
            _ = v10
            _ = v11
        }
        switch "foo bar" {
        case "hello":
            var v12 uint64 = 42
            _ = v12
        case "foo bar":
            fmt.Println(v1, v3, v5)
        case "\n\t":
            var v13 uint64 = 2
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            v3 = v3
            _ = fi14
        }
        _ = fi8
    }
    var v15 string = (("" + "a") + ("" + "hello"))
    v16 := GenS1{F0: (v4 * 2), F1: 0}
    v6 = append(v6, 655)
    var v17 uint32 = uint32(v5)
    v18 := GenS2{F0: v4, F1: 0, F2: 127}
    for fi19 := 0; fi19 < 10; fi19++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi19
    }
    switch "\n\t" {
    case "a":
        var v20 uint64 = 2
        _ = v20
    case "":
        var v21 uint8 = 128
        _ = v21
    default:
        fmt.Println(v4, v3)
    }
    v17 += 4294967295
    v2 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
