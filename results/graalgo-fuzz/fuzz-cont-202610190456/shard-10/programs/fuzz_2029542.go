package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint32 }
func genFn0() uint16 {
    v4 := GenS1{F0: -12, F1: 1000}
    v5 := []int{1, 4294967295, 32768, -792}
    var v6 uint64 = 335
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1() uint32 {
    v7 := []int{32767, -52}
    var v8 int8 = 103
    v8 += v8
    fmt.Println(v8)
    v8 = 42
    _ = v7
    _ = v8
    return 900
}

func main() {
    var v1 uint64 = 373
    v2 := make(map[string]string)
    v3 := []uint32{0, 768}
    var v9 uint16 = 2
    var v10 uint64 = v1
    v11 := GenS2{F0: 2}
    v12 := GenS2{F0: 2147483648}
    var v13 uint64 = 18446744073709551615
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 string = "\n\t"
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 2147483647}
    v10 -= 2
    var v17 string = ("a" + ("a" + "a"))
    v3 = append(v3, genFn1())
    for fi18 := 0; fi18 < 9; fi18++ {
        if false {
            v13 = (^(v13))
            _ = len(v2)
            var v19 uint16 = (v9 >> ((18446744073709551615) & 63))
            _ = v19
        }
        var v20 float32 = 1.5
        v13 = (^(v1))
        _ = fi18
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
