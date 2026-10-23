package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 float64 = 1.5
    var v2 int = 0
    v3 := []string{"", ("\n\t" + ("foo bar" + "\n\t")), "foo bar", ""}
    var v4 uint = 826
    var v5 float32 = 1.5
    var v6 int8 = 100
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(fi7, v2, v4)
        if false {
            v5 = ((v5 - -2.25) + (((v5 * v5) - v5) * ((((v5 / v5) / 1.5) / (v5 * ((v5 * (v5 / (v5 + v5))) * float32(-2.25)))) * float32(-2.25))))
        }
        var v8 int8 = v6
        _ = fi7
        _ = v8
    }
    v4 = v4
    fmt.Println(v2)
    fmt.Println(v5, v6)
    if false {
        var v9 float32 = float32(1.5)
        var v10 uint = v4
        v4 = (v10 % ((v4) | 1))
        _ = v10
        _ = v9
    }
    v5 = -2.25
    v5 = (1.5 / (1.5 * -2.25))
    fmt.Println(v1, v6, v2)
    v11 := GenS1{F0: 2147483647}
    if ((529 != 2147483647) || true) {
        fmt.Println(v5)
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            switch ((("\n\t" + "hello") + "\n\t") + "") {
            case "\n\t":
                _ = len(v3)
            default:
                var v13 uint8 = 72
                _ = v13
            }
            _ = fi12
        }
        var v14 bool = ((111 != 0) || false)
        v15 := GenS1{F0: 266}
        _ = v14
        _ = v15
    }
    var v16 uint = (v4 >> (((v4 & v4)) & 63))
    v17 := GenS1{F0: 1}
    v6 = (^(v6))
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
