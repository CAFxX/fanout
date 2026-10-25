package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int }
func genFn0(p0 int8, p1 uint32) int {
    switch 32767 {
    case 2147483647:
        p1 = (p1 - 305)
    case 253:
        fmt.Println(p0, p0)
    }
    fmt.Println(p0, p1)
    var v4 uint8 = 0
    var v5 uint32 = p1
    var v6 bool = (!(32768 < -2147483648))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 32767
}
func genFn1() bool {
    var v7 uint = 4294967295
    v7 -= v7
    v7 = 211
    _ = v7
    return (((true && (148 < 2147483647)) || true) || (4294967295 >= 65535))
}

func main() {
    var v1 uint = 139
    var v2 float32 = 0.0
    v3 := []int{2147483647, 127}
    for fi8 := 0; fi8 < 4; fi8++ {
        v1 = v1
        _ = fi8
    }
    v9 := GenS1{F0: (-597 <= 42), F1: ("hello" + "foo bar"), F2: 1}
    var v10 float32 = v2
    v1 += (^(v1))
    v10 = ((-2.25 / v10) / 1e10)
    fmt.Println(v1, v1)
    fmt.Println(v10, v1, v1)
    v11 := GenS1{F0: (true && false), F1: "\n\t", F2: 969}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
