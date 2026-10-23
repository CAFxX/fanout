package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
func genFn0(p0 int64, p1 bool, p2 string) int32 {
    var v6 bool = (255 <= 255)
    v7 := GenS1{F0: 1, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 int16 = 42
    var v2 int8 = 59
    var v3 float64 = ((1e10 - -2.25) * 1.5)
    var v4 int16 = v1
    var v5 int8 = (v2 << ((364) & 63))
    var v8 float64 = ((((0.0 / ((float64((0.1 - 3.14159)) + v3) + (3.14159 * 3.14159))) - (v3 - 3.14159)) + v3) / float64((float32(v3) + 1.5)))
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := []int{(-(fi9)), fi9, (-(fi9)), (fi9 + fi9)}
        _ = fi9
        _ = v10
    }
    var v11 bool = (((float64((1e10 / 1.0)) + 1e10) + v3) < v3)
    var v12 uint = 129
    var v13 int16 = 32767
    var v14 int16 = 0
    v15 := []int{-2147483648, 4294967295, 32767, 32768}
    var v16 int = 42
    for fi17 := 0; fi17 < 3; fi17++ {
        v16 = v16
        if v11 {
            fmt.Println(v13)
        }
        _ = fi17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        fmt.Println(v4, v12)
        var v19 string = "a"
        _ = fi18
        _ = v19
    }
    v2 = -1
    var v20 uint16 = 0
    var v21 uint16 = (v20 >> ((1) & 63))
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 uint16 = (^(v20))
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
