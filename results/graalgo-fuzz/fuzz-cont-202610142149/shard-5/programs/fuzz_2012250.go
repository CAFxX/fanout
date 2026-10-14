package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint32 }
func genFn0(p0 string, p1 uint, p2 string) int8 {
    var v4 int32 = -566
    for fi5 := 0; fi5 < 19; fi5++ {
        p2 = "foo bar"
        v6 := []int{fi5, -312, fi5}
        var v7 string = p0
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return -1
}

func main() {
    v1 := []int{-933, 32768, 32767}
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    v9 := GenS1{F0: "hello", F1: 100, F2: 347}
    if (-1 <= -1) {
        var v10 int64 = 0
        v10 = (^(v10))
        _ = len(v1)
        _ = v10
    }
    var v11 int = 9223372036854775807
    v12 := GenS1{F0: "a", F1: -1, F2: 1}
    fmt.Println(v11)
    switch "foo bar" {
    case "hello":
        fmt.Println(v11)
    default:
        v11 = 32768
    }
    switch ("\n\t" + "") {
    case "\n\t":
        _ = len(v1)
    case "foo bar":
        v13 := GenS1{F0: "a", F1: 17, F2: 2147483647}
        _ = v13
    case "a":
        fmt.Println(v11)
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v11)
    default:
        var v14 int64 = 9223372036854775807
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
