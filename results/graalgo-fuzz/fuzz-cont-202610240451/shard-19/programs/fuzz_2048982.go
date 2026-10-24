package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint = 384
    v2 := []string{"", "hello"}
    var v3 float32 = (-2.25 * 0.1)
    var v4 int16 = -31
    v3 = -2.25
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    v1 = v1
    var v5 float32 = float32((1e10 + -2.25))
    var v6 int8 = 86
    for fi7 := 0; fi7 < 2; fi7++ {
        switch "" {
        case "\n\t":
            var v8 int64 = 42
            _ = v8
        case "a":
            fmt.Println(v6, v5, v6)
        case "hello":
            for fi9 := 0; fi9 < 7; fi9++ {
                switch ("" + (((("\n\t" + "") + ("hello" + "a")) + "") + "a")) {
                case "":
                    v5 = (v5 - v3)
                default:
                    if true {
                        var v10 int16 = (-(v4))
                        _ = v10
                    }
                }
                if true {
                    _ = len(v2)
                    fmt.Println(v3)
                    v5 = v3
                }
                _ = fi9
            }
        }
        fmt.Println(v5, v6, fi7)
        _ = fi7
    }
    fmt.Println(v6, v5)
    v11 := GenS1{F0: 2147483648}
    fmt.Println(v4, v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
