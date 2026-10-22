package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int8 }
func genFn0(p0 string) int {
    v4 := GenS1{F0: 32768, F1: 624, F2: -128}
    v5 := GenS1{F0: -2147483648, F1: 1, F2: 97}
    _ = p0
    _ = v4
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 int16 = -562
    var v2 uint32 = 2
    var v3 int16 = (v1 << ((255) & 63))
    v6 := []int{127, genFn0(""), 4294967295, 42}
    var v7 string = "hello"
    v2 = 596
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v8 string = ((v7 + (v7 + ("a" + (("" + (v7 + ((v7 + ("hello" + v7)) + "foo bar"))) + (v7 + "\n\t"))))) + v7)
    v6 = append(v6, 4294967295)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = len(v6)
    for fi9 := 0; fi9 < 17; fi9++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi9
    }
    v7 = v7
    v10 := GenS1{F0: 1, F1: 0, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
