package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := []uint32{4294967295, 2147483647}
    var v2 int32 = -124
    var v3 int64 = 9223372036854775807
    var v4 uint16 = 742
    var v5 int = 0
    var v6 bool = ("\n\t" != (("" + "\n\t") + ""))
    for fi7 := 0; fi7 < 18; fi7++ {
        v6 = true
        v6 = true
        switch "\n\t" {
        case "\n\t":
            var v8 string = ""
            _ = v8
        case "a":
            fmt.Println(v3, v3)
        }
        _ = fi7
    }
    v2 -= v2
    v3 = v3
    switch "a" {
    case "\n\t":
        var v9 int = v5
        _ = v9
    case "foo bar":
        var v10 uint16 = (v4 << ((769) & 63))
        _ = v10
    }
    v5 = 1
    switch (v5 >> ((255) & 63)) {
    case -9223372036854775808:
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
    case 127:
        var v11 int8 = 39
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        fmt.Println(v2)
        if (482 == 42) {
            v4 = v4
        }
        _ = fi12
    }
    switch 32768 {
    case 32768:
        v5 = (v5 + 32768)
    default:
        v13 := GenS1{F0: 0}
        _ = v13
    }
    if true {
        var v14 bool = ((126 > 2) && ("" < "foo bar"))
        _ = v14
    } else {
        v15 := GenS1{F0: v4}
        var v16 float64 = 0.1
        var v17 float64 = ((v16 - (v16 + (0.0 + 1e10))) * 3.14159)
        _ = v15
        _ = v16
        _ = v17
    }
    v6 = (v2 <= v2)
    var v18 float32 = (0.0 / (1.5 - 0.1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
