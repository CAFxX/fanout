package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
func genFn0(p0 int8, p1 bool, p2 uint64) int8 {
    var v5 uint64 = 0
    if (0 <= 4294967295) {
        var v6 int = 127
        _ = v6
    } else {
        v7 := []int{127, -624, -1, 32767}
        _ = v7
    }
    var v8 float64 = -2.25
    p0 += -79
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return (^(p0))
}
func genFn1(p0 uint64, p1 uint32, p2 uint) int8 {
    var v9 uint64 = (^(p0))
    v9 -= p0
    v10 := GenS1{F0: v9, F1: (127 > 42)}
    if ((^(p2)) > p2) {
        fmt.Println(v9, p0, p1)
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    if true {
        var v11 int = -1
        _ = v11
    }
    var v12 uint16 = 1
    var v13 string = "foo bar"
    _ = len(v2)
    if false {
        var v14 int16 = 1
        v15 := v3[2147483647]
        v16 := GenS1{F0: 1, F1: false}
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 uint16 = (v12 << ((18446744073709551615) & 63))
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := GenS1{F0: 42, F1: (true && false)}
        _ = fi18
        _ = v19
    }
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
