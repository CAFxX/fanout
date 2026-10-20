package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 bool }
type GenS2 struct { F0 int64 }
func genFn0() int32 {
    var v5 int8 = -128
    fmt.Println(v5)
    v5 = 42
    _ = v5
    return -172
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 127
    v3 := []string{}
    v4 := make(map[int]int)
    var v6 float32 = ((float32((0.0 + (0.0 - ((1.5 * (0.1 + (0.0 + -2.25))) + 1e10)))) / (-2.25 + 1.5)) - 3.14159)
    v2 = v2
    if true {
        var v7 uint16 = 32767
        _ = v7
    } else {
        v8 := GenS1{F0: -128, F1: false, F2: (0.0 > v6)}
        var v9 uint8 = v2
        _ = v8
        _ = v9
    }
    var v10 uint = 1
    for fi11 := 0; fi11 < 8; fi11++ {
        switch (fi11 + fi11) {
        case 42:
            v12 := GenS2{F0: 0}
            _ = v12
        }
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := GenS2{F0: -813}
            fmt.Println(fi13, v10, v2)
            _ = fi13
            _ = v14
        }
        _ = fi11
    }
    v6 = v6
    v6 += 0.0
    if (-1 > -9223372036854775808) {
        var v15 int64 = 9223372036854775807
        var v16 bool = true
        v6 -= v6
        _ = v15
        _ = v16
    } else {
        switch ("foo bar" + "a") {
        case "foo bar":
            delete(v1, 2147483647)
        }
    }
    var v17 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
