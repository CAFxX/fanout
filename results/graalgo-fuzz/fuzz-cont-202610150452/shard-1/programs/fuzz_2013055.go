package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int; F2 uint32 }
func genFn0(p0 uint16) string {
    var v5 uint32 = 416
    if (!((!(0 < -2147483648)) || (!(127 != 1)))) {
        var v6 int8 = 32
        _ = v6
    } else {
        var v7 uint64 = 762
        _ = v7
    }
    _ = p0
    _ = v5
    return ("" + "foo bar")
}
func genFn1() uint32 {
    v8 := GenS1{F0: 123}
    var v9 int8 = 17
    _ = v8
    _ = v9
    return 77
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    v3 := []string{"\n\t", "a", "foo bar", (("" + "a") + "foo bar")}
    var v4 int8 = 42
    v4 -= 42
    var v10 string = "a"
    fmt.Println(v10, v1, v4)
    v4 -= v4
    fmt.Println(v1, v10, v10)
    v10 = v10
    v11 := GenS2{F0: v4, F1: -2147483648, F2: 2}
    v12 := GenS1{F0: v4}
    var v13 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
