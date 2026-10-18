package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 string; F1 string; F2 int }
func genFn0(p0 int32) string {
    p0 -= -200
    p0 -= (p0 >> ((0) & 63))
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := GenS1{F0: -1, F1: true}
        fmt.Println(fi5, fi5)
        _ = fi5
        _ = v6
    }
    _ = p0
    return ""
}
func genFn1(p0 int8) int8 {
    p0 = p0
    var v7 string = ""
    switch 127 {
    case -287:
        fmt.Println(p0)
    }
    var v8 uint8 = 2
    _ = p0
    _ = v7
    _ = v8
    return -44
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float64 = -2.25
    var v3 uint = 42
    var v4 uint = v3
    fmt.Println(v3)
    switch "foo bar" {
    case "":
        var v9 string = "a"
        _ = v9
    default:
        v3 -= (v3 >> ((v4) & 63))
    }
    var v10 uint32 = 4294967295
    v1 = 0
    switch ((((("foo bar" + ("\n\t" + ("" + ("\n\t" + ("hello" + ("" + "")))))) + "a") + "a") + ("\n\t" + "\n\t")) + ("\n\t" + ("a" + "a"))) {
    case "a":
        var v11 int32 = -2147483648
        _ = v11
    case "hello":
        var v12 uint16 = 42
        _ = v12
    default:
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := []int{fi13, fi13, (fi13 & 202)}
            var v15 uint32 = 1
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    fmt.Println(v3, v1, v2)
    switch "foo bar" {
    case "\n\t":
        v16 := []int{2147483647, -9223372036854775808}
        _ = v16
    default:
        v17 := []int{-207, 9223372036854775807, 670, 128}
        _ = v17
    }
    var v18 uint64 = 160
    v2 += (v2 - (v2 - v2))
    v19 := GenS2{F0: "hello", F1: "a", F2: 127}
    v3 -= 387
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
