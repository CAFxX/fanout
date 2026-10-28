package main

import (
    "fmt"
)


func main() {
    var v1 uint = 145
    var v2 uint32 = 42
    var v3 uint16 = 65535
    var v4 int16 = -35
    var v5 uint16 = v3
    if ((^(v4)) != (-(v4))) {
        v6 := []int{-9223372036854775808, 127, 500, 9223372036854775807}
        fmt.Println(v2, v3)
        for fi7 := 0; fi7 < 12; fi7++ {
            switch ("a" + "foo bar") {
            case "\n\t":
                var v8 int8 = -1
                _ = v8
            case "":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi7
        }
        _ = v6
    } else {
        v1 = (v1 - v1)
    }
    var v9 float32 = 0.0
    v10 := []int{2147483647, 2147483647, 0}
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint64 = 18446744073709551615
        v12 -= (^(v12))
        var v13 int = (fi11 ^ fi11)
        _ = fi11
        _ = v12
        _ = v13
    }
    v5 += (v3 >> ((2) & 63))
    v4 -= v4
    var v14 uint32 = v2
    fmt.Println(v14, v4)
    var v15 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
