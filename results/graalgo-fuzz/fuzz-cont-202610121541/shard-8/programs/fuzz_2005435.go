package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 bool; F1 bool; F2 int8 }

func main() {
    var v1 int = 4294967295
    v2 := make(map[int]int)
    var v3 uint16 = 255
    var v4 uint16 = 854
    var v5 string = "foo bar"
    var v6 bool = (65535 == 2147483647)
    switch "foo bar" {
    case "\n\t":
        var v7 int = (v1 | v1)
        _ = v7
    default:
        var v8 int16 = 1000
        _ = v8
    }
    var v9 bool = v6
    v4 = v3
    v3 += (v4 + v3)
    for fi10 := 0; fi10 < 1; fi10++ {
        switch (((("a" + (v5 + v5)) + v5) + "") + "\n\t") {
        case "\n\t":
            for fi11 := 0; fi11 < 14; fi11++ {
                fmt.Println(v3, v4)
                v12 := GenS1{F0: 0, F1: v5}
                var v13 float64 = float64(3.14159)
                _ = fi11
                _ = v12
                _ = v13
            }
        case "foo bar":
            fmt.Println(v3, v1, v5)
        default:
            if (2 < 204) {
                fmt.Println(v9)
                v14 := GenS2{F0: false, F1: (!v9), F2: 127}
                _ = v14
            } else {
                var v15 int16 = 1000
                var v16 uint16 = 65535
                _ = v15
                _ = v16
            }
        }
        _ = fi10
    }
    fmt.Println(v1)
    v1 += (v1 ^ 32768)
    v17 := GenS1{F0: 255, F1: (v5 + (("foo bar" + v5) + ("foo bar" + ((("a" + v5) + "\n\t") + "hello"))))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
