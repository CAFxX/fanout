package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int8; F2 int }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (0.1 / -2.25)
    var v3 int8 = 127
    v4 := make(map[uint32]string)
    delete(v4, 665)
    v2 = v2
    var v5 uint8 = 42
    var v6 bool = true
    if v6 {
        if ((((v2 / v2) / (v2 + v2)) * 0.1) != float32((0.1 * 3.14159))) {
            v2 = (v2 + float32(-2.25))
            switch "" {
            case "a":
                v5 += 255
            case "hello":
                v4[2] = ("" + (("foo bar" + (("a" + "\n\t") + "foo bar")) + "a"))
            case "\n\t":
                fmt.Println(v6, v5)
            default:
                var v7 string = (("\n\t" + ("foo bar" + ("" + "foo bar"))) + "")
                _ = v7
            }
            fmt.Println(v6)
        }
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 int64 = 1
            _ = fi8
            _ = v9
        }
        v3 -= 20
    } else {
        if v6 {
            v6 = (("\n\t" + ("a" + "")) >= ((("a" + "") + "a") + ""))
            v10 := GenS1{F0: 275}
            v5 += (v5 / ((1) | 1))
            _ = v10
        }
    }
    v1[0] = false
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 bool = v6
        switch "\n\t" {
        case "\n\t":
            v13 := GenS2{F0: -1, F1: v3, F2: fi11}
            _ = v13
        }
        var v14 int64 = 648
        _ = fi11
        _ = v12
        _ = v14
    }
    var v15 uint32 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
