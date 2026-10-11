package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
func genFn0() string {
    var v6 uint8 = 255
    var v7 string = "hello"
    _ = v6
    _ = v7
    return (((v7 + (((v7 + v7) + "hello") + (v7 + (("\n\t" + "foo bar") + v7)))) + "a") + (v7 + v7))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 18446744073709551615
    var v3 uint32 = 2147483647
    v4 := []int{-85, 128}
    var v5 int16 = -1
    _ = len(v1)
    v8 := v1[(v3 >> ((42) & 63))]
    v9 := GenS1{F0: "hello", F1: 0}
    fmt.Println(v3)
    fmt.Println(v5)
    for fi10 := 0; fi10 < 5; fi10++ {
        fmt.Println(v5)
        _ = fi10
    }
    v3 -= (^(v3))
    v2 -= 31
    v11 := GenS1{F0: (("\n\t" + (v8 + v8)) + ("\n\t" + v8)), F1: 2}
    var v12 string = (v8 + v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
