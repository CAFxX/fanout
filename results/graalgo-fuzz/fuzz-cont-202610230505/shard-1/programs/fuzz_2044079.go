package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int }
func genFn0(p0 uint32, p1 uint16, p2 uint32) int8 {
    v5 := []int{-54, 32767}
    p2 -= 149
    p1 = (p1 | p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 1
    var v3 int = 1
    var v4 int64 = -1
    delete(v1, (v3 >> ((227) & 63)))
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 string = ("foo bar" + "\n\t")
        var v8 uint32 = 120
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float64 = (1.5 - float64((1.5 + 3.14159)))
    v10 := GenS1{F0: -2147483648, F1: ("foo bar" + "hello"), F2: (v3 >> ((0) & 63))}
    var v11 int8 = 69
    v12 := GenS1{F0: 1, F1: "a", F2: v3}
    switch (v3 * 128) {
    case 32767:
        switch 1 {
        case 32768:
            v3 += 332
        }
    case 4294967295:
        v3 -= 267
    case -945:
        v13 := GenS1{F0: 0, F1: "", F2: (v3 * int(v11))}
        _ = v13
    default:
        v2 = (v2 >> ((642) & 63))
    }
    v14 := GenS1{F0: -2147483648, F1: "foo bar", F2: (v3 & (v3 << ((629) & 63)))}
    var v15 int32 = -2147483648
    fmt.Println(v9, v4)
    fmt.Println(v4)
    v16 := GenS1{F0: v15, F1: "\n\t", F2: v3}
    v1[v3] = (v3 | v3)
    v11 = -70
    fmt.Println(v3)
    var v17 bool = (true && ((("a" + "foo bar") + "") < ""))
    var v18 uint32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
