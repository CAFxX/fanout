package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 bool }
type GenS2 struct { F0 bool; F1 uint16 }
func genFn0(p0 int32, p1 uint8, p2 int16) string {
    fmt.Println(p2, p0, p1)
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int16 = p2
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}
func genFn1() bool {
    var v8 float64 = -2.25
    v8 -= -2.25
    fmt.Println(v8)
    _ = v8
    return false
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    var v3 string = (((((("hello" + ("hello" + "\n\t")) + "\n\t") + "\n\t") + "") + ("a" + "a")) + "a")
    var v4 float64 = 3.14159
    var v5 string = ((v3 + v3) + ("foo bar" + "\n\t"))
    fmt.Println(v4)
    v9 := GenS1{F0: 581, F1: 2, F2: true}
    var v10 uint64 = 4294967295
    var v11 int = 127
    v1[(v11 | -928)] = (v3 + v3)
    v12 := GenS1{F0: v11, F1: 2147483647, F2: true}
    v13 := []int{v11, 32768, (v11 & (v11 + (v11 | v11)))}
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 uint64 = 4294967295
        var v16 uint64 = 42
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int = (v11 >> ((1) & 63))
    var v18 string = v3
    var v19 uint64 = (^(v10))
    var v20 uint64 = v19
    v21 := GenS2{F0: false, F1: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
