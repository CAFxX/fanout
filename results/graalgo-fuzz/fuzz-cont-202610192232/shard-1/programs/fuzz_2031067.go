package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 4294967295
    v3 := make(map[string]int)
    var v4 int8 = 111
    if (42 >= 2147483647) {
        for fi5 := 0; fi5 < 1; fi5++ {
            v6 := []int{}
            _ = fi5
            _ = v6
        }
        fmt.Println(v4, v2)
    }
    v7 := GenS1{F0: false, F1: ""}
    if true {
        v8 := []int{}
        var v9 uint = 1
        fmt.Println(v2, v4, v4)
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 19; fi10++ {
            v4 = v4
            var v11 uint64 = 847
            _ = fi10
            _ = v11
        }
        var v12 int16 = 1000
        _ = v12
    }
    v2 = v2
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 int16 = -524
            _ = fi14
            _ = v15
        }
        var v16 int = (fi13 << ((uint(v2)) & 63))
        _ = fi13
        _ = v16
    }
    switch 2147483647 {
    case 128:
        fmt.Println(v4, v4)
    }
    var v17 bool = (1 != 1)
    var v18 bool = v17
    v2 = (v2 ^ (v2 & 247))
    switch 32767 {
    case 1:
        var v19 uint32 = 4294967295
        _ = v19
    }
    v20 := GenS1{F0: v18, F1: "foo bar"}
    var v21 int64 = 9223372036854775807
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
