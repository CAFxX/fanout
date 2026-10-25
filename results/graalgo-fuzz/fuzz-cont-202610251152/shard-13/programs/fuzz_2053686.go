package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 bool }
func genFn0(p0 int) int64 {
    fmt.Println(p0)
    var v6 bool = (32767 == 530)
    for fi7 := 0; fi7 < 14; fi7++ {
        p0 = fi7
        var v8 uint16 = 1
        _ = fi7
        _ = v8
    }
    p0 -= (p0 - 9223372036854775807)
    _ = p0
    _ = v6
    return -461
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 124
    var v3 int8 = 78
    var v4 float64 = (1e10 / 3.14159)
    v5 := make(map[string]int)
    var v9 int8 = 54
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 bool = false
        v12 := []int{(fi10 >> ((0) & 63)), (^(fi10))}
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        v2 += v2
        var v13 int16 = 0
        _ = v13
    }
    fmt.Println(v4, v4, v4)
    v2 -= v2
    fmt.Println(v4, v2)
    v14 := v5["\n\t"]
    var v15 string = "\n\t"
    var v16 int8 = (-(v3))
    delete(v1, 4294967295)
    v17 := GenS1{F0: (-(v14)), F1: true, F2: (!false)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
