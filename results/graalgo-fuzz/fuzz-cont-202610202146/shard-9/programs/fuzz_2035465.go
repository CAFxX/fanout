package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int; F2 string }
func genFn0(p0 int16) int64 {
    fmt.Println(p0)
    v5 := GenS1{F0: ""}
    v6 := GenS1{F0: (("foo bar" + ("" + "")) + "a")}
    fmt.Println(p0)
    _ = p0
    _ = v5
    _ = v6
    return 878
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = (2147483648 <= 7)
    v3 := []int{}
    v4 := []uint32{65535, 0, 65535}
    _ = len(v1)
    if v2 {
        v2 = (-9223372036854775808 == 344)
    }
    v2 = v2
    v7 := GenS2{F0: 42, F1: -228, F2: "\n\t"}
    v8 := GenS1{F0: "\n\t"}
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int8 = -1
        _ = len(v1)
        _ = fi9
        _ = v10
    }
    fmt.Println(v2)
    var v11 uint16 = 255
    switch "hello" {
    case "\n\t":
        v12 := GenS2{F0: -128, F1: 127, F2: "\n\t"}
        _ = v12
    case "a":
        if v2 {
            v13 := GenS1{F0: "foo bar"}
            _ = v13
        } else {
            v3 = append(v3, -1)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
