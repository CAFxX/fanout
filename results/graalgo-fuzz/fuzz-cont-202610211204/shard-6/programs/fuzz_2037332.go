package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int32, p1 uint32, p2 string) int {
    p1 = 766
    var v6 uint32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 32768
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 0.0
    var v3 float64 = 1.5
    var v4 int = 127
    var v5 uint8 = 74
    var v7 bool = (!((1 <= 0) && true))
    var v8 int32 = 2147483647
    v9 := GenS1{F0: 643}
    var v10 int32 = (v8 & 1)
    fmt.Println(v3)
    v11 := v1["\n\t"]
    v8 = (v8 ^ -2147483648)
    var v12 string = "foo bar"
    switch genFn0(-2147483648, (v11 * 214), "hello") {
    case 42:
        v10 = v10
    case -1:
        v5 = v5
    case 353:
        v12 = (v12 + "foo bar")
    }
    fmt.Println(v3, v7, v10)
    switch v4 {
    case 1:
        switch v12 {
        case "hello":
            v3 = v3
        case "\n\t":
            var v13 string = (("\n\t" + ("a" + v12)) + v12)
            _ = v13
        }
    case 24:
        v1[((((v12 + (v12 + "foo bar")) + (v12 + "a")) + ("foo bar" + v12)) + v12)] = 675
    default:
        for fi14 := 0; fi14 < 11; fi14++ {
            v2 -= (0.1 / ((0.0 + v2) * v2))
            _ = fi14
        }
    }
    var v15 int32 = (v10 - (v10 - 42))
    fmt.Println(v2)
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 uint = 2
        _ = fi16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
