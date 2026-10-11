package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int64 }
type GenS2 struct { F0 int32 }
func genFn0(p0 bool, p1 uint8, p2 uint8) int32 {
    fmt.Println(p1, p0)
    fmt.Println(p1)
    var v5 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    var v1 uint = 780
    v2 := make(map[string]string)
    v3 := []string{"foo bar", "\n\t", ("foo bar" + ("" + "")), "\n\t"}
    v4 := []int{1, -9223372036854775808, 105}
    v6 := v2[("" + ("\n\t" + "\n\t"))]
    v7 := GenS1{F0: 4294967295, F1: 2147483648, F2: 9223372036854775807}
    v1 -= (v1 / ((v1) | 1))
    var v8 uint = 1
    switch v6 {
    case "":
        fmt.Println(v6, v6)
    }
    v8 += (v1 >> ((65535) & 63))
    var v9 string = "hello"
    for fi10 := 0; fi10 < 7; fi10++ {
        v2[v9] = "a"
        switch v6 {
        case "foo bar":
            delete(v2, v6)
        case "hello":
            v11 := GenS2{F0: -1}
            _ = v11
        default:
            var v12 uint8 = 128
            _ = v12
        }
        _ = fi10
    }
    var v13 int = -833
    v14 := GenS2{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
