package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
func genFn0(p0 int16, p1 int32) uint {
    var v5 string = "\n\t"
    switch ((v5 + v5) + "a") {
    case "foo bar":
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := []int{fi6, fi6}
            _ = fi6
            _ = v7
        }
    case "a":
        v8 := []int{128, 343, 4294967295}
        _ = v8
    case "hello":
        var v9 bool = (42 == 18446744073709551615)
        _ = v9
    default:
        var v10 uint16 = 255
        _ = v10
    }
    v11 := GenS1{F0: -60, F1: -51}
    var v12 int64 = 717
    var v13 int16 = (p0 & p0)
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v5
    return 42
}

func main() {
    var v1 float32 = (((3.14159 * -2.25) / float32(-2.25)) + 3.14159)
    v2 := []string{""}
    var v3 uint16 = 32767
    var v4 float64 = 0.0
    v1 += 3.14159
    fmt.Println(v3)
    fmt.Println(v1)
    var v14 int16 = 415
    var v15 float64 = ((v4 / v4) - 3.14159)
    switch 9223372036854775807 {
    case -2147483648:
        var v16 float32 = (v1 / v1)
        _ = v16
    }
    if false {
        if true {
            var v17 int32 = 1
            _ = v17
        }
    } else {
        v3 += (v3 + (v3 * (v3 & 1)))
        fmt.Println(v3, v1, v1)
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
