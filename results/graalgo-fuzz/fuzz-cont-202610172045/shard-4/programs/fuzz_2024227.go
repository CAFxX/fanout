package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }

func main() {
    var v1 int64 = -712
    var v2 uint8 = 127
    var v3 int = 9223372036854775807
    v4 := make(map[int]int)
    switch (v3 & v3) {
    case 2147483647:
        v4[-9223372036854775808] = 0
    case -1:
        for fi5 := 0; fi5 < 6; fi5++ {
            for fi6 := 0; fi6 < 9; fi6++ {
                v2 = v2
                fmt.Println(v1, v2, v2)
                fmt.Println(v3, v1)
                _ = fi6
            }
            _ = fi5
        }
    }
    if ("a" <= ("" + "hello")) {
        var v7 uint32 = 968
        var v8 int64 = (v1 ^ 42)
        _ = v7
        _ = v8
    }
    switch (v3 * 128) {
    case -680:
        v3 = (v3 + v3)
    }
    v9 := GenS1{F0: 42, F1: v3}
    v2 += 2
    v10 := GenS1{F0: 42, F1: -896}
    for fi11 := 0; fi11 < 15; fi11++ {
        fmt.Println(v2)
        var v12 float64 = float64(-2.25)
        _ = fi11
        _ = v12
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
