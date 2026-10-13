package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 string) int16 {
    fmt.Println(p0)
    var v7 float32 = (3.14159 + (3.14159 - -2.25))
    var v8 uint16 = 65535
    v9 := []int{1, 582, 471, -1}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int) uint16 {
    var v10 float32 = float32(float64(0.0))
    if false {
        v11 := []int{808, (-(p0))}
        _ = v11
    }
    var v12 int = p0
    fmt.Println(v10, p0)
    v10 += v10
    _ = p0
    _ = v10
    _ = v12
    return 32767
}

func main() {
    var v1 bool = (190 != 99)
    v2 := make(map[string]uint32)
    var v3 string = (((((("\n\t" + "") + "a") + ("hello" + "\n\t")) + ("" + "hello")) + "hello") + "hello")
    var v4 int32 = 42
    v5 := make(map[string]int)
    var v6 int64 = 1
    var v13 int8 = 127
    var v14 float64 = 0.0
    if v1 {
        if false {
            fmt.Println(v14)
            v3 = "foo bar"
            v15 := GenS1{F0: v1}
            _ = v15
        }
    } else {
        v16 := GenS1{F0: v1}
        v13 = (^(v13))
        _ = v16
    }
    var v17 uint32 = 65535
    var v18 int64 = (^(v6))
    var v19 float32 = (0.1 + -2.25)
    if (int8(v17) <= v13) {
        var v20 uint32 = (v17 % ((v17) | 1))
        _ = v20
    }
    switch v3 {
    case "foo bar":
        v21 := []int{108}
        _ = v21
    case "hello":
        v3 = v3
    default:
        var v22 int = -1
        _ = v22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        if (127 > 43) {
            v24 := []int{fi23}
            _ = v24
        }
        var v25 uint32 = (^(v17))
        _ = fi23
        _ = v25
    }
    v3 = "hello"
    if (v19 > float32(v14)) {
        fmt.Println(v3, v19)
        v26 := []int{}
        _ = v26
    }
    var v27 string = v3
    _ = len(v2)
    v28 := GenS1{F0: (v1 || (!(4294967295 > 42)))}
    v17 = 353
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
