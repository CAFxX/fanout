package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int, p1 uint8) uint64 {
    var v6 string = ("\n\t" + ("foo bar" + (("\n\t" + ("hello" + "")) + "\n\t")))
    p1 = 127
    switch (p0 / (((^(p0))) | 1)) {
    case -26:
        v6 = v6
    case 128:
        var v7 int64 = -826
        _ = v7
    default:
        var v8 string = "hello"
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 672
}

func main() {
    v1 := []string{}
    var v2 int16 = 1000
    var v3 int = 42
    var v4 float32 = 0.0
    var v5 uint16 = 520
    fmt.Println(v5, v5, v5)
    switch ("a" + ("hello" + "foo bar")) {
    case "hello":
        if false {
            v5 = 1
            fmt.Println(v5, v2)
        }
    case "\n\t":
        v1 = append(v1, "hello")
    }
    v9 := GenS1{F0: 0, F1: v5}
    var v10 float32 = (float32(float64(v4)) / (v4 - v4))
    if true {
        switch 149 {
        case -1:
            fmt.Println(v4, v4, v3)
        default:
            var v11 uint64 = 42
            _ = v11
        }
    }
    switch "\n\t" {
    case "":
        var v12 uint8 = 166
        _ = v12
    default:
        if false {
            var v13 float64 = 1.5
            switch ("hello" + "foo bar") {
            case "foo bar":
                var v14 uint16 = (v5 & v5)
                _ = v14
            }
            _ = v13
        }
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        fmt.Println(v3, fi15)
        _ = fi15
    }
    v16 := GenS1{F0: 1, F1: (^(v5))}
    v5 += 704
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
