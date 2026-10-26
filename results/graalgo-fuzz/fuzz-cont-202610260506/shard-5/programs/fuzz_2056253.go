package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0(p0 bool, p1 bool) int {
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 0}
        _ = fi5
        _ = v6
    }
    switch "hello" {
    case "foo bar":
        v7 := GenS1{F0: 127}
        _ = v7
    case "":
        p0 = (((p0 && p1) && true) || p0)
    case "hello":
        var v8 uint32 = 2
        _ = v8
    default:
        for fi9 := 0; fi9 < 3; fi9++ {
            fmt.Println(fi9)
            var v10 string = ("" + ((((("\n\t" + ((("" + "\n\t") + "a") + "a")) + "hello") + ("a" + ("" + ("" + ((("hello" + (("a" + "foo bar") + (("a" + "") + "a"))) + "hello") + "a"))))) + ("" + "hello")) + "\n\t"))
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint64 = 838
    var v2 uint16 = 65535
    var v3 uint16 = 1
    v4 := make(map[uint32]bool)
    var v11 int16 = -32768
    fmt.Println(v3, v11)
    v2 = (v3 % (((v2 ^ (v3 % (((v3 >> ((0) & 63))) | 1)))) | 1))
    v1 -= 42
    v3 += v3
    var v12 int32 = 0
    v13 := GenS2{F0: 9223372036854775807, F1: -800}
    var v14 int64 = -9223372036854775808
    v1 = v1
    fmt.Println(v2, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
