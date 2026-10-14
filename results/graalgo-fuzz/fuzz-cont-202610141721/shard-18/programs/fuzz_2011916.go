package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
func genFn0() int {
    var v4 int8 = 100
    v4 = v4
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 int32 = 42
    v2 := []uint32{2147483647, 121}
    v3 := make(map[string]int)
    switch "\n\t" {
    case "foo bar":
        v1 = 0
    }
    var v5 uint = 42
    v1 += (v1 + v1)
    fmt.Println(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 -= v1
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: -926}
            _ = fi7
            _ = v8
        }
        var v9 int = fi6
        v10 := GenS1{F0: v9}
        _ = fi6
        _ = v10
        _ = v9
    }
    if false {
        v1 = v1
        v1 += (-(v1))
        var v11 uint = (v5 >> ((v5) & 63))
        _ = v11
    }
    var v12 float32 = -2.25
    v13 := GenS2{F0: true}
    for fi14 := 0; fi14 < 7; fi14++ {
        delete(v3, "foo bar")
        _ = fi14
    }
    fmt.Println(v1, v5)
    v12 -= 0.0
    v1 = (v1 % (((-(v1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v5
}
