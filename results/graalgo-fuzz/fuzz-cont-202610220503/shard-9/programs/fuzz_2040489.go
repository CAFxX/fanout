package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint }
func genFn0(p0 uint, p1 int8) bool {
    v6 := []int{}
    fmt.Println(p1)
    for fi7 := 0; fi7 < 7; fi7++ {
        p0 = p0
        var v8 uint64 = 0
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS1{F0: 216}
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v11 := GenS2{F0: 749, F1: p0}
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return false
}

func main() {
    var v1 int32 = 42
    var v2 uint64 = 335
    v3 := make(map[string]string)
    var v4 float64 = 1e10
    var v5 int8 = -1
    fmt.Println(v1, v1)
    var v12 uint32 = 1
    var v13 float32 = (0.1 - (((0.0 - 0.0) + 1.5) * -2.25))
    for fi14 := 0; fi14 < 18; fi14++ {
        switch fi14 {
        case 42:
            delete(v3, (("\n\t" + ("a" + "foo bar")) + "\n\t"))
        case 32767:
            var v15 bool = (v12 > v12)
            _ = v15
        }
        var v16 float64 = v4
        var v17 uint8 = 127
        _ = fi14
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: (v1 + v1), F1: 42}
    v4 -= 1e10
    v1 -= v1
    v19 := []int{}
    if true {
        v20 := GenS1{F0: 2}
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 int8 = (v5 >> ((642) & 63))
            _ = len(v3)
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    v23 := GenS1{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
