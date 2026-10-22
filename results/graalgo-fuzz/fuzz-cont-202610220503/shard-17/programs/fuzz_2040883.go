package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 int32 = 2147483647
    var v4 bool = true
    v4 = v4
    v5 := []int{-342, -2147483648, 128}
    var v6 int8 = -50
    for fi7 := 0; fi7 < 9; fi7++ {
        v3 -= (-(v3))
        _ = fi7
    }
    var v8 string = ("hello" + "foo bar")
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 = 0
    var v9 int8 = 42
    if ((!(!(false || true))) && v4) {
        fmt.Println(v8, v4)
        v10 := GenS1{F0: (2 > 0)}
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS1{F0: v4}
        fmt.Println(v3, v9)
        v5 = append(v5, -2147483648)
        _ = fi11
        _ = v12
    }
    if (42 > 42) {
        if false {
            var v13 bool = true
            _ = v13
        }
    } else {
        var v14 float64 = float64(((1.5 / (1.5 * float32(-2.25))) + -2.25))
        _ = v14
    }
    v15 := GenS1{F0: v4}
    v9 = v6
    fmt.Println(v9, v8, v6)
    switch -1 {
    case 962:
        var v16 uint32 = 2
        _ = v16
    default:
        var v17 uint32 = 0
        _ = v17
    }
    v18 := GenS1{F0: (2147483647 == v3)}
    v19 := GenS2{F0: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
