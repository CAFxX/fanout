package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
func genFn0() uint8 {
    v4 := GenS1{F0: ((!(false || true)) || true)}
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := []int{(fi5 - fi5), 1}
        _ = fi5
        _ = v6
    }
    var v7 uint = 176
    v8 := GenS1{F0: true}
    switch "foo bar" {
    case "foo bar":
        switch "\n\t" {
        case "hello":
            v7 -= (v7 % ((v7) | 1))
        case "\n\t":
            fmt.Println(v7)
        }
    case "a":
        if (false && false) {
            var v9 bool = true
            _ = v9
        }
    case "\n\t":
        var v10 uint32 = 0
        _ = v10
    }
    _ = v4
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{("foo bar" + "foo bar")}
    var v3 float32 = 0.0
    var v11 float32 = -2.25
    fmt.Println(v3)
    for fi12 := 0; fi12 < 2; fi12++ {
        v1 = v1
        v11 = 0.1
        v13 := GenS1{F0: true}
        _ = fi12
        _ = v13
    }
    v2 = append(v2, "foo bar")
    var v14 string = "hello"
    fmt.Println(v11)
    v15 := GenS2{F0: 1}
    v11 = 0.0
    fmt.Println(v1, v14, v11)
    var v16 int8 = 0
    var v17 int8 = -128
    switch v14 {
    case "\n\t":
        fmt.Println(v14, v11, v17)
    default:
        fmt.Println(v17)
    }
    var v18 float64 = (0.1 + 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
