package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() int8 {
    var v7 int = 127
    v7 = (v7 >> ((255) & 63))
    _ = v7
    return -79
}

func main() {
    v1 := []int{-1, -850, 4294967295, -2147483648}
    v2 := make(map[uint32]string)
    var v3 float32 = (float32(float64(((1.5 + 1.5) - 1e10))) * 1e10)
    var v4 int = -470
    var v5 float64 = 1.5
    var v6 int8 = -31
    fmt.Println(v5, v5, v4)
    var v8 bool = (false && false)
    for fi9 := 0; fi9 < 8; fi9++ {
        if (!(42 <= 127)) {
            fmt.Println(v8, v4, fi9)
            var v10 int32 = 1
            _ = len(v1)
            _ = v10
        }
        v8 = v8
        for fi11 := 0; fi11 < 2; fi11++ {
            v5 -= (v5 * ((v5 * ((3.14159 * (v5 - (v5 - 0.0))) - (float64(v3) - v5))) - v5))
            fmt.Println(v8, v4, fi9)
            _ = fi11
        }
        _ = fi9
    }
    v6 = -1
    v6 = (v6 ^ v6)
    var v12 float32 = (v3 / v3)
    switch "a" {
    case "hello":
        fmt.Println(v6, v6, v3)
    case "":
        var v13 uint64 = 4294967295
        _ = v13
    case "\n\t":
        v14 := GenS1{F0: ("foo bar" + "foo bar")}
        _ = v14
    }
    var v15 uint16 = 42
    for fi16 := 0; fi16 < 14; fi16++ {
        v4 += v4
        var v17 uint = 42
        var v18 uint16 = (v15 * v15)
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint32 = 415
    var v20 uint16 = v15
    for fi21 := 0; fi21 < 9; fi21++ {
        if v8 {
            var v22 float32 = ((-2.25 / 1e10) * (v12 / v3))
            for fi23 := 0; fi23 < 11; fi23++ {
                v5 = 0.1
                _ = fi23
            }
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v22
        }
        v1 = append(v1, -847)
        _ = fi21
    }
    var v24 string = ""
    fmt.Println(v8, v15)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
