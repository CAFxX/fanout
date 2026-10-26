package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint32) int {
    v7 := GenS1{F0: 4294967295, F1: 0}
    v8 := []int{-329, 0}
    _ = p0
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int16) uint64 {
    var v9 int8 = -12
    if ((p0 == p0) || ((!(!false)) || (1 > -1))) {
        fmt.Println(v9, v9)
    }
    _ = p0
    _ = v9
    return 0
}

func main() {
    var v1 int16 = 446
    var v2 bool = false
    v3 := []string{"\n\t", ((("" + ("foo bar" + "foo bar")) + ("a" + ("hello" + ("foo bar" + "foo bar")))) + "")}
    v4 := []int{135, 623, 32768, 4294967295}
    v5 := make(map[string]uint32)
    var v6 uint8 = 129
    if v2 {
        _ = len(v5)
        fmt.Println(v1, v2, v2)
        switch 0 {
        case 42:
            var v10 int32 = 1
            _ = v10
        }
    } else {
        var v11 float32 = 0.0
        var v12 uint = 4294967295
        _ = len(v4)
        _ = v11
        _ = v12
    }
    v2 = (!v2)
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS2{F0: 0}
        _ = fi13
        _ = v14
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        var v15 uint = 55
        _ = v15
    } else {
        v6 -= (^(v6))
    }
    v16 := GenS2{F0: 2}
    var v17 float32 = ((3.14159 * float32(-2.25)) * (3.14159 + float32(0.1)))
    v18 := GenS3{F0: -93}
    var v19 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
