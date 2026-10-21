package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int }
type GenS2 struct { F0 bool }
func genFn0(p0 uint16, p1 bool) uint8 {
    p0 = p0
    switch "hello" {
    case "a":
        var v4 int8 = 0
        _ = v4
    case "":
        fmt.Println(p0, p1)
    case "hello":
        fmt.Println(p1)
    }
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := []int{(fi5 + 128)}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 2147483647, F1: 1, F2: -1}
    v8 := []int{-158, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 uint16, p1 bool) int64 {
    v9 := []int{}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = p1
    _ = v9
    return 693
}

func main() {
    var v1 int16 = -68
    v2 := make(map[int]int)
    var v3 int32 = -921
    var v10 int64 = genFn1(2, true)
    v11 := v2[-100]
    var v12 float64 = (1e10 - ((-2.25 / (float64(((3.14159 * ((1e10 * 1.5) + ((3.14159 - ((0.1 - (3.14159 - (0.0 * 0.0))) / 0.1)) / 1e10))) / float32(1.5))) + float64(((0.0 / 1e10) - 0.1)))) - 0.1))
    v10 += 42
    var v13 int8 = -20
    fmt.Println(v1, v3)
    fmt.Println(v10, v1)
    if false {
        var v14 int64 = (v10 | int64(v13))
        _ = v14
    }
    v15 := v2[v11]
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int32 = 1
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 int = fi16
            v20 := []int{(v11 - 32767), 4294967295, fi18}
            v21 := GenS2{F0: false}
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    v1 = 1
    v22 := GenS2{F0: (16 >= (^(v13)))}
    switch v11 {
    case 885:
        v12 += (v12 - v12)
    default:
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 float32 = float32((((-2.25 * v12) * 0.0) - v12))
            _ = fi23
            _ = v24
        }
    }
    var v25 int16 = -32768
    v25 = (v1 & int16(v15))
    for fi26 := 0; fi26 < 20; fi26++ {
        v3 = (v3 & v3)
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v22
    _ = v25
    _ = v3
}
