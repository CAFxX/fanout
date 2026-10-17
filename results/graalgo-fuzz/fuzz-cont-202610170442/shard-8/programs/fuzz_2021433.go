package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
func genFn0(p0 int16, p1 int) int32 {
    switch p1 {
    case -1:
        v7 := GenS1{F0: 368, F1: 18446744073709551615}
        _ = v7
    case 128:
        switch "hello" {
        case "foo bar":
            p0 -= p0
        case "a":
            var v8 int64 = 0
            _ = v8
        }
    default:
        if false {
            fmt.Println(p1, p1)
            fmt.Println(p0, p1)
        }
    }
    var v9 int = (p1 * -2147483648)
    var v10 int64 = 1
    p1 -= v9
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 float32 = -2.25
    var v2 float64 = 0.0
    var v3 int64 = 237
    v4 := []string{"a"}
    var v5 uint32 = 1
    var v6 string = (("hello" + ("hello" + "foo bar")) + "foo bar")
    v1 = (v1 / 0.1)
    if (true || false) {
        v5 = v5
        var v11 float32 = v1
        _ = v11
    }
    var v12 uint8 = 132
    var v13 uint32 = 1
    v4 = append(v4, (v6 + ""))
    v14 := GenS1{F0: (v3 & v3), F1: 255}
    var v15 int = -270
    v4 = append(v4, ("foo bar" + ""))
    v4 = append(v4, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
