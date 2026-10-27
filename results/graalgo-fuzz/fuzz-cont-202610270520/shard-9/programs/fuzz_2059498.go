package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
func genFn0() bool {
    var v4 float32 = (-2.25 / 3.14159)
    v4 = (0.0 / 1.0)
    switch "\n\t" {
    case "":
        switch "foo bar" {
        case "a":
            var v5 float64 = -2.25
            _ = v5
        case "\n\t":
            v4 = float32(((1.5 / 3.14159) - (1.5 / 1.5)))
        case "":
            v4 = v4
        default:
            if false {
                v6 := GenS1{F0: -953, F1: 4294967295}
                _ = v6
            }
        }
    case "a":
        var v7 int8 = -53
        _ = v7
    }
    _ = v4
    return ((42 <= 592) || (!(275 <= 667)))
}

func main() {
    var v1 uint16 = 2
    var v2 string = "foo bar"
    v3 := make(map[int]uint32)
    v2 = ""
    var v8 int32 = 2147483647
    var v9 uint32 = 1
    v10 := GenS1{F0: (v8 + 150), F1: 377}
    _ = len(v3)
    v2 = "foo bar"
    v11 := []int{54, 32767, 0}
    var v12 float32 = 1.5
    switch (("" + "\n\t") + "a") {
    case "hello":
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 float32 = v12
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            var v15 int8 = 0
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    var v16 float32 = v12
    v3[-562] = 2147483648
    v16 -= float32(1.5)
    v9 = uint32(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
