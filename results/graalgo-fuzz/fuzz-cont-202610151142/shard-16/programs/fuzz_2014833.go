package main

import (
    "fmt"
)


func main() {
    var v1 string = ((((("a" + ("foo bar" + "hello")) + "\n\t") + "foo bar") + "foo bar") + "")
    var v2 int64 = 9223372036854775807
    var v3 string = "foo bar"
    var v4 float64 = (1.5 + ((((1.5 + 3.14159) / 1.5) / -2.25) * ((1.5 * 0.0) / 1.0)))
    var v5 float32 = (1e10 + (float32(v4) / 1e10))
    var v6 string = v1
    var v7 float64 = 0.0
    if true {
        var v8 int16 = -599
        _ = v8
    } else {
        var v9 uint8 = 84
        switch v3 {
        case "hello":
            fmt.Println(v1, v1, v9)
        }
        _ = v9
    }
    var v10 uint8 = 2
    var v11 int = 4294967295
    if ((false && true) && true) {
        v12 := []int{}
        _ = v12
    } else {
        v4 = float64(1.5)
        var v13 float32 = v5
        _ = v13
    }
    v3 = "a"
    v14 := []int{561, v11}
    if len(v14) > 0 {
        v14[0] = 2147483647
    }
    var v15 uint64 = 600
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 uint32 = 42
        if true {
            var v18 int = (fi16 << ((2) & 63))
            _ = v18
        }
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
