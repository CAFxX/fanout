package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint8 }
func genFn0() int {
    v4 := []int{42}
    var v5 uint64 = 617
    if (0 > -507) {
        fmt.Println(v5)
    } else {
        var v6 uint16 = 0
        _ = v6
    }
    v7 := GenS1{F0: -2147483648}
    _ = v4
    _ = v5
    _ = v7
    return 1
}
func genFn1() int64 {
    var v8 uint16 = 2
    fmt.Println(v8)
    _ = v8
    return 953
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    v3 := []uint32{697}
    var v9 uint32 = 3
    v10 := GenS1{F0: -107}
    var v11 float32 = 1e10
    switch "hello" {
    case "a":
        var v12 string = ("\n\t" + "")
        _ = v12
    default:
        var v13 float32 = v11
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS2{F0: fi14, F1: 988}
        var v16 float64 = 1e10
        v9 = (v9 ^ v9)
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint8 = 255
    v18 := GenS3{F0: (v17 / ((v17) | 1))}
    _ = len(v2)
    var v19 float32 = 0.0
    var v20 string = ("foo bar" + (("hello" + "hello") + "foo bar"))
    fmt.Println(v11, v20)
    for fi21 := 0; fi21 < 8; fi21++ {
        if len(v3) > 0 {
            v3[0] = v9
        }
        _ = fi21
    }
    v20 = v20
    var v22 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v9
}
