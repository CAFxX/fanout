package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 bool; F1 int16; F2 string }
type GenS3 struct { F0 int8; F1 string }
func genFn0(p0 int16) uint8 {
    var v5 uint = 2
    v6 := []int{-591}
    var v7 float32 = (float32(1e10) * (0.1 + (0.1 / 1.5)))
    _ = len(v6)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 60
}

func main() {
    v1 := make(map[string]int)
    var v2 string = (("\n\t" + ((("foo bar" + "\n\t") + "a") + "a")) + "hello")
    var v3 int64 = 1
    v4 := make(map[string]int)
    var v8 int64 = (v3 + -758)
    var v9 uint32 = 0
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 float32 = 1e10
        _ = fi10
        _ = v11
    }
    v8 = v3
    switch v2 {
    case "a":
        var v12 float32 = (3.14159 + float32(((-2.25 + 0.0) * -2.25)))
        _ = v12
    default:
        if false {
            v13 := GenS1{F0: 128, F1: (!(v2 == v2))}
            _ = v13
        } else {
            fmt.Println(v2, v8, v3)
        }
    }
    v8 -= -891
    fmt.Println(v9, v8, v2)
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS2{F0: true, F1: 961, F2: (v2 + ("" + v2))}
        v16 := []int{-9223372036854775808, 127, 9223372036854775807, fi14}
        for fi17 := 0; fi17 < 14; fi17++ {
            v2 = "foo bar"
            v18 := GenS3{F0: 42, F1: v2}
            var v19 int64 = (v3 >> ((0) & 63))
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    v3 = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
