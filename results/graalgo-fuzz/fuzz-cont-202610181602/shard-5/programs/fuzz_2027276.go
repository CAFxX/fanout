package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 string }
func genFn0(p0 uint64) uint32 {
    v4 := []int{745, 9223372036854775807}
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint = 0
        _ = fi5
        _ = v6
    }
    p0 -= (p0 + (p0 & p0))
    _ = p0
    _ = v4
    return 42
}
func genFn1() string {
    var v7 int8 = -1
    var v8 uint32 = 42
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    var v3 uint16 = 42
    v3 -= v3
    v9 := GenS1{F0: -621, F1: 27}
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS2{F0: ((("a" + "\n\t") + ("\n\t" + "hello")) + "foo bar")}
        _ = fi10
        _ = v11
    }
    v3 = v3
    v3 -= v3
    var v12 int32 = 42
    fmt.Println(v3, v3)
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
