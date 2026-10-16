package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []int{32767, 128, -277}
    var v2 uint16 = 0
    var v3 float64 = 1.5
    var v4 float32 = 0.0
    var v5 int64 = -1
    var v6 int = 1
    switch ("" + "hello") {
    case "hello":
        v6 = (v6 % ((127) | 1))
    case "\n\t":
        var v7 uint8 = 14
        _ = v7
    default:
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 string = ""
            _ = fi8
            _ = v9
        }
    }
    switch (("\n\t" + "a") + "foo bar") {
    case "\n\t":
        v6 -= 9223372036854775807
    }
    var v10 float64 = (0.0 * float64(-2.25))
    v10 += (1e10 / ((0.1 / (v3 + v10)) - ((v10 / (v10 + float64(v4))) - v10)))
    fmt.Println(v2)
    for fi11 := 0; fi11 < 3; fi11++ {
        v5 = 9223372036854775807
        v12 := GenS1{F0: 141}
        switch "\n\t" {
        case "a":
            v3 = (v3 - v3)
        }
        _ = fi11
        _ = v12
    }
    v1 = append(v1, v6)
    var v13 bool = false
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if v13 {
        var v14 uint16 = 255
        _ = v14
    }
    v1 = append(v1, (v6 >> ((120) & 63)))
    v1 = append(v1, (^(v6)))
    var v15 int64 = (v5 / ((v5) | 1))
    v5 = v5
    var v16 uint32 = 454
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
