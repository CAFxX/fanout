package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 string, p1 int64, p2 int8) uint8 {
    v4 := GenS1{F0: 42, F1: 100, F2: 1}
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 8
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 1e10
    var v3 int64 = 9223372036854775807
    switch ("" + "\n\t") {
    case "hello":
        v1["a"] = "hello"
    case "\n\t":
        fmt.Println(v2, v2)
    }
    var v5 uint = 255
    delete(v1, "hello")
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 128, F1: -79, F2: 210}
        var v8 bool = (false && (((765 >= 1) && (0 == 870)) && false))
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: -9223372036854775808, F1: -128, F2: (^(v5))}
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS2{F0: -128}
        fmt.Println(v5, v5)
        var v12 int8 = 0
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        v1["hello"] = ("foo bar" + "")
        var v13 string = "\n\t"
        _ = v13
    } else {
        v2 = ((float64(float32((1.5 * (((float64(((0.1 * (-2.25 - 1e10)) * 1e10)) * (0.1 * (v2 + (v2 - ((((v2 - float64(((-2.25 - (((3.14159 / 1.0) / (float32(0.1) / 1.5)) + 0.0)) * (-2.25 / (0.0 - (0.1 * (float32(float64(-2.25)) / 0.1))))))) * (1.5 - v2)) / v2) + v2))))) * ((v2 - ((v2 / 1e10) * v2)) + v2)) + v2)))) * v2) * (v2 / (3.14159 - (3.14159 / (v2 / v2)))))
    }
    v5 = v5
    var v14 uint32 = 42
    var v15 int16 = 0
    var v16 int64 = v3
    v1[("foo bar" + "hello")] = "\n\t"
    v17 := []int{-1, 42, -1, -9223372036854775808}
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 int64 = v3
        fmt.Println(v3, v19, fi18)
        v16 += 664
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v21 := GenS3{F0: 9}
        for fi22 := 0; fi22 < 17; fi22++ {
            v5 -= v5
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
