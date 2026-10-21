package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int32) int32 {
    var v4 uint16 = 0
    v4 -= v4
    v5 := []int{32768, 0, -905}
    _ = p0
    _ = v4
    _ = v5
    return (p0 % ((p0) | 1))
}
func genFn1() int32 {
    v6 := GenS1{F0: 32767}
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint = 255
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint16 = 2
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 1
    var v12 uint64 = 774
    _ = v11
    _ = v12
    _ = v6
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 3.14159
    var v3 float32 = float32(0.1)
    _ = len(v1)
    fmt.Println(v3)
    v13 := []int{-1, -212, -396}
    v2 -= 0.0
    v2 -= v3
    v14 := GenS1{F0: 708}
    v2 = (v2 * float32(((1.5 / (3.14159 / 3.14159)) * float64(v2))))
    var v15 float32 = 0.0
    var v16 uint16 = 172
    var v17 int32 = 2147483647
    var v18 uint32 = 596
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 string = ("" + "")
        _ = fi19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
