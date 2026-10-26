package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint }

func main() {
    v1 := []string{"foo bar", (("\n\t" + "a") + "a"), "a", ("" + "")}
    v2 := []int{}
    v3 := make(map[string]int)
    var v4 bool = false
    var v5 float64 = 1e10
    v6 := make(map[string]uint32)
    switch "foo bar" {
    case "hello":
        if v4 {
            switch "a" {
            case "a":
                var v7 string = "\n\t"
                _ = v7
            default:
                switch "hello" {
                case "a":
                    v4 = ((!v4) || false)
                }
            }
        }
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        v4 = (((-9223372036854775808 < -1) && v4) && true)
        v9 := GenS1{F0: 72}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 0, F1: 4294967295}
    var v11 uint16 = 937
    v5 -= -2.25
    v11 -= (^(v11))
    v1 = append(v1, "hello")
    if v4 {
        v4 = (32767 >= 42)
        v12 := GenS1{F0: (v11 << ((65535) & 63))}
        _ = v12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
