package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int16 }
func genFn0() int {
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int16 = -888
        v7 -= v7
        _ = fi6
        _ = v7
    }
    var v8 uint = 4294967295
    _ = v8
    return -2147483648
}
func genFn1() bool {
    var v9 string = "hello"
    v9 = v9
    v10 := []int{367, genFn0(), -311, genFn0()}
    _ = v10
    _ = v9
    return true
}

func main() {
    var v1 uint = 2
    v2 := make(map[int]int)
    var v3 uint32 = 920
    v4 := []uint32{147, v3, v3, 192}
    var v5 int16 = 1
    fmt.Println(v3, v5)
    switch "a" {
    case "hello":
        v11 := GenS1{F0: "\n\t", F1: 100, F2: v5}
        _ = v11
    case "foo bar":
        v12 := GenS1{F0: "a", F1: 1, F2: (v5 * (v5 - (v5 ^ (-(v5)))))}
        _ = v12
    default:
        v3 = v3
    }
    v5 -= (^(v5))
    fmt.Println(v1, v3)
    v4 = append(v4, (v3 ^ 784))
    v3 = v3
    fmt.Println(v3, v5, v5)
    var v13 float32 = 1e10
    if true {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 int = genFn0()
            _ = len(v2)
            var v16 float32 = (((1.5 / v13) / v13) / ((((0.0 - 1e10) + 3.14159) * v13) / (((1e10 / v13) * 3.14159) + v13)))
            _ = fi14
            _ = v15
            _ = v16
        }
        v1 += 4294967295
    }
    fmt.Println(v1, v5, v3)
    if false {
        var v17 string = ((("a" + "") + "") + ("" + "\n\t"))
        _ = v17
    }
    switch "a" {
    case "":
        v18 := GenS1{F0: "a", F1: 100, F2: v5}
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
