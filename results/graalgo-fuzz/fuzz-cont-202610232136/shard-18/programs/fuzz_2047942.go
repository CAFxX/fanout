package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 string; F1 int16 }

func main() {
    v1 := []string{}
    var v2 uint = 255
    var v3 uint64 = 418
    var v4 string = ("\n\t" + "foo bar")
    var v5 int8 = 127
    fmt.Println(v3)
    var v6 uint64 = v3
    for fi7 := 0; fi7 < 14; fi7++ {
        v3 -= 42
        if false {
            var v8 uint64 = v3
            _ = v8
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(v6, fi7, v2)
            v2 = v2
            v3 = uint64(v2)
            v6 = v6
            if (!(fi7 > fi7)) {
                var v10 int = 469
                _ = v10
            }
            _ = fi9
        }
        _ = fi7
    }
    var v11 float64 = 1.5
    v12 := GenS1{F0: v4, F1: 114}
    v11 = (v11 - ((v11 + 1e10) - v11))
    var v13 int64 = 1
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 uint8 = 2
        v11 += ((v11 - float64(0.1)) - 1.5)
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: v4, F1: -938}
    v6 += uint64(v2)
    var v17 uint = v2
    if false {
        var v18 uint64 = v3
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
