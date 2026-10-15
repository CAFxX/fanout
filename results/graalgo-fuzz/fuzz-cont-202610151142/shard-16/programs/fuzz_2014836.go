package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int8 }
func genFn0(p0 string, p1 int) uint32 {
    p0 = (("a" + p0) + p0)
    v7 := []int{}
    _ = len(v7)
    v8 := GenS1{F0: (128 != 2), F1: ("\n\t" + (p0 + p0)), F2: 127}
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int32) int {
    p0 -= 510
    v9 := GenS1{F0: (!true), F1: (("hello" + "hello") + "hello"), F2: 100}
    fmt.Println(p0)
    _ = p0
    _ = v9
    return -590
}

func main() {
    var v1 uint8 = 201
    v2 := []string{}
    var v3 int8 = -67
    var v4 uint32 = 65535
    var v5 bool = (((v4 ^ v4) < (^(v4))) && ("foo bar" != "hello"))
    v6 := make(map[int]uint32)
    var v10 int32 = -1
    var v11 uint8 = v1
    var v12 uint8 = v11
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            if false {
                var v15 int = 9223372036854775807
                _ = v15
            } else {
                v3 -= (^(v3))
            }
            _ = fi14
        }
        for fi16 := 0; fi16 < 5; fi16++ {
            if v5 {
                v17 := GenS1{F0: v5, F1: "\n\t", F2: (v3 / (((v3 - (v3 & (v3 * v3)))) | 1))}
                _ = v17
            } else {
                v5 = v5
                _ = len(v6)
            }
            fmt.Println(v4, v11)
            _ = fi16
        }
        _ = fi13
    }
    var v18 string = "a"
    delete(v6, -401)
    _ = len(v6)
    fmt.Println(v12, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
