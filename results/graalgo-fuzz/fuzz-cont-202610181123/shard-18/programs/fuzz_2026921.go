package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
func genFn0(p0 int) uint32 {
    fmt.Println(p0)
    var v7 uint32 = 35
    switch "\n\t" {
    case "\n\t":
        var v8 int = p0
        _ = v8
    case "a":
        fmt.Println(v7, p0)
    case "foo bar":
        v9 := []int{p0, p0, 42}
        _ = v9
    }
    _ = p0
    _ = v7
    return (v7 ^ 230)
}

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]string)
    v3 := []string{"foo bar"}
    v4 := []string{"foo bar", "\n\t"}
    var v5 string = ("" + ("hello" + "hello"))
    var v6 uint32 = 2147483647
    if (0 >= 42) {
        var v10 uint64 = 18446744073709551615
        _ = v10
    } else {
        v11 := GenS1{F0: v5}
        v2[v6] = "a"
        v4 = append(v4, v5)
        _ = v11
    }
    v1 += v1
    var v12 int8 = (-(v1))
    fmt.Println(v1, v5)
    var v13 uint64 = 18446744073709551615
    v4 = append(v4, v5)
    v13 = v13
    v14 := v2[0]
    v15 := GenS1{F0: v5}
    var v16 int8 = v12
    v17 := GenS1{F0: ("" + (((v14 + v14) + v14) + (v14 + "\n\t")))}
    v18 := GenS1{F0: ((v5 + "") + v14)}
    switch ((v5 + v5) + "") {
    case "":
        v5 = (v14 + ("hello" + "hello"))
    case "\n\t":
        fmt.Println(v13, v14)
    case "foo bar":
        var v19 uint16 = 32767
        _ = v19
    default:
        v5 = v5
    }
    v20 := GenS2{F0: -442}
    switch 955 {
    case 209:
        v21 := GenS2{F0: -233}
        _ = v21
    case -595:
        var v22 uint16 = 1
        _ = v22
    }
    var v23 uint8 = 127
    var v24 uint = uint(v13)
    v13 += v13
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
