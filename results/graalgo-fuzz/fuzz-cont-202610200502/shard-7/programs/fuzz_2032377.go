package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint8) int8 {
    v7 := GenS1{F0: ((true || (42 != -1)) && (false && true))}
    v8 := []int{-1}
    var v9 bool = true
    fmt.Println(p0, p0)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -52
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = -2.25
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    var v5 uint8 = 0
    var v6 float32 = (v2 * (((((v2 / v2) / v2) - v2) / (v2 * (v2 - v2))) - v2))
    var v10 float64 = -2.25
    v6 += (0.0 + v2)
    fmt.Println(v5, v10)
    v11 := GenS1{F0: false}
    var v12 int32 = 1
    var v13 float32 = 1e10
    var v14 int64 = 42
    v15 := GenS1{F0: false}
    var v16 string = ("\n\t" + "foo bar")
    var v17 string = v16
    var v18 int = -509
    switch v18 {
    case -873:
        v19 := GenS2{F0: v14}
        _ = v19
    case -9223372036854775808:
        v20 := v1[0]
        _ = v20
    default:
        fmt.Println(v2)
    }
    var v21 uint32 = 65535
    v14 = (^(v14))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
