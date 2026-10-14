package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }

func main() {
    v1 := []string{"\n\t", "hello", ""}
    var v2 uint64 = 159
    var v3 uint32 = 134
    var v4 uint8 = 255
    var v5 float32 = 3.14159
    var v6 int64 = 147
    fmt.Println(v6, v2)
    for fi7 := 0; fi7 < 13; fi7++ {
        v5 = -2.25
        _ = fi7
    }
    fmt.Println(v5, v6)
    var v8 int = 32768
    fmt.Println(v2, v5, v6)
    var v9 int16 = 1000
    v10 := GenS1{F0: (v6 | v6), F1: (v9 / ((v9) | 1))}
    for fi11 := 0; fi11 < 5; fi11++ {
        v6 = v6
        _ = fi11
    }
    v2 += 1
    switch ("hello" + "a") {
    case "hello":
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    case "foo bar":
        switch (v8 >> ((42) & 63)) {
        case -9223372036854775808:
            v4 -= (v4 % ((v4) | 1))
        default:
            if true {
                v9 += v9
            }
        }
    case "\n\t":
        for fi12 := 0; fi12 < 2; fi12++ {
            v1 = append(v1, "\n\t")
            _ = fi12
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
