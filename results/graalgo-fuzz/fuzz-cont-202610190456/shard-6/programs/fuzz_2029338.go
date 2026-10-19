package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint64) int16 {
    var v4 int8 = 42
    if false {
        fmt.Println(v4)
        var v5 uint64 = 510
        _ = v5
    } else {
        var v6 int32 = 1
        var v7 uint16 = 0
        fmt.Println(v6, v4, p0)
        _ = v6
        _ = v7
    }
    var v8 int32 = -609
    _ = p0
    _ = v4
    _ = v8
    return 1000
}

func main() {
    var v1 float32 = 0.0
    var v2 uint64 = 2
    v3 := make(map[int]string)
    v1 -= (-2.25 / ((float32(-2.25) * (1.5 / v1)) * (v1 / v1)))
    if false {
        v1 = v1
        switch 639 {
        case 204:
            if (!(32767 >= 630)) {
                v9 := GenS1{F0: 1, F1: 255}
                var v10 uint16 = 0
                _ = v10
                _ = v9
            } else {
                v3[556] = "\n\t"
            }
        }
    }
    var v11 int32 = -2147483648
    v3[0] = ("hello" + "hello")
    var v12 int8 = -128
    v13 := GenS1{F0: -319, F1: 14}
    v2 -= v2
    var v14 float64 = 0.1
    var v15 string = (("foo bar" + "hello") + ("\n\t" + (("foo bar" + "") + "")))
    var v16 string = ((v15 + v15) + "hello")
    var v17 uint32 = 287
    if (float64(0.1) < (v14 - v14)) {
        if false {
            v16 = "\n\t"
        } else {
            v18 := GenS2{F0: (!false), F1: v2}
            _ = v18
        }
        v19 := GenS3{F0: 255}
        _ = v19
    } else {
        var v20 float32 = v1
        fmt.Println(v2, v15, v11)
        _ = v20
    }
    v21 := GenS3{F0: 2}
    v22 := v3[32768]
    v15 = (v15 + "hello")
    if true {
        _ = len(v3)
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 uint64 = uint64(v17)
            _ = fi23
            _ = v24
        }
    } else {
        var v25 uint8 = 255
        v26 := []int{}
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
