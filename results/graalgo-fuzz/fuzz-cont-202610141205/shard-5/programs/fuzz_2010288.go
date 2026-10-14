package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
func genFn0(p0 int32) int64 {
    v7 := []int{-714}
    if ((437 == 82) || false) {
        var v8 string = "hello"
        _ = v8
    }
    _ = p0
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    var v4 uint = 65535
    v5 := []string{}
    v6 := make(map[int]bool)
    var v9 string = ("\n\t" + ("\n\t" + (("" + ("a" + "")) + "foo bar")))
    v9 = v9
    v9 = (v9 + v9)
    switch -817 {
    case 212:
        v4 -= (v4 + (v4 / ((v4) | 1)))
    case -217:
        if false {
            for fi10 := 0; fi10 < 6; fi10++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi10
            }
            v11 := GenS1{F0: 127, F1: true}
            _ = v11
        }
    }
    v12 := v1[v9]
    var v13 float64 = 1e10
    if true {
        v9 = (v9 + v9)
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        v9 = "a"
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
