package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
func genFn0(p0 int, p1 int64, p2 bool) int64 {
    for fi7 := 0; fi7 < 17; fi7++ {
        p0 = p0
        _ = fi7
    }
    switch "" {
    case "hello":
        var v8 uint16 = 1
        _ = v8
    }
    var v9 bool = p2
    p1 += -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (p1 ^ -488)
}
func genFn1() bool {
    if true {
        var v10 uint8 = 127
        v11 := []int{755}
        var v12 bool = (!false)
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 float32 = float32(-2.25)
    _ = v13
    return false
}

func main() {
    var v1 uint16 = 775
    var v2 float64 = (3.14159 - 0.0)
    v3 := make(map[int]int)
    v4 := []string{"a", ""}
    var v5 uint8 = 174
    var v6 uint32 = 65535
    var v14 uint8 = (v5 | 81)
    for fi15 := 0; fi15 < 7; fi15++ {
        v2 += (1.5 + 0.0)
        v16 := GenS1{F0: -2147483648}
        var v17 int8 = 100
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 42}
    var v19 int = 9223372036854775807
    for fi20 := 0; fi20 < 19; fi20++ {
        delete(v3, v19)
        v21 := v3[(v19 << ((1) & 63))]
        var v22 int8 = 69
        _ = fi20
        _ = v21
        _ = v22
    }
    v2 = v2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v23 int32 = 2147483647
    if true {
        v24 := GenS2{F0: int(v1)}
        var v25 uint = 42
        var v26 int = (v19 << ((1) & 63))
        _ = v24
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: v23}
    if true {
        _ = len(v3)
        _ = len(v3)
        var v28 bool = (!(325 != -235))
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
