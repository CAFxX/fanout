package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8) int {
    var v7 int32 = -2147483648
    v7 += v7
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: -32768, F1: 414}
        v10 := []int{741}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{}
    v12 := GenS2{F0: true}
    _ = p0
    _ = v11
    _ = v12
    _ = v7
    return -2147483648
}

func main() {
    v1 := []int{-1, 524, 0, 620}
    v2 := make(map[string]int)
    var v3 int64 = 116
    var v4 int64 = v3
    var v5 bool = (1 <= 100)
    var v6 int16 = -1
    v3 -= 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v13 string = "a"
    if (100 < 1) {
        fmt.Println(v6, v6, v13)
        v3 += v4
    }
    var v14 bool = true
    var v15 uint16 = 2
    var v16 uint64 = 0
    switch ("\n\t" + (("foo bar" + (v13 + ((v13 + "hello") + v13))) + "\n\t")) {
    case "\n\t":
        v3 -= (v3 + 1)
    default:
        v4 += (v3 + v4)
    }
    v5 = false
    var v17 int64 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
