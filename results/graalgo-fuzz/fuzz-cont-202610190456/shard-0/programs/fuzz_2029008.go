package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint32 }
type GenS2 struct { F0 string; F1 int32; F2 int }

func main() {
    var v1 string = "a"
    v2 := []int{9223372036854775807, 1, 2147483647}
    var v3 int64 = -135
    if ((255 > 0) || true) {
        v3 -= (v3 * v3)
        var v4 uint16 = 98
        v5 := GenS1{F0: -662, F1: 4294967295, F2: 2147483647}
        _ = v4
        _ = v5
    }
    v6 := GenS1{F0: -994, F1: 595, F2: 65535}
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v7 := GenS2{F0: "hello", F1: 330, F2: 32768}
    var v8 string = "\n\t"
    _ = len(v2)
    var v9 uint = 65535
    v8 = v1
    switch ((v1 + v8) + v8) {
    case "":
        var v10 bool = false
        _ = v10
    }
    v2 = append(v2, 2147483647)
    v8 = v8
    v11 := GenS2{F0: ("\n\t" + v8), F1: -833, F2: 127}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
