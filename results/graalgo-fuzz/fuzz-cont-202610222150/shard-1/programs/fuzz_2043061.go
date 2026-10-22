package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
func genFn0(p0 int32, p1 uint8, p2 int8) string {
    var v5 int8 = p2
    switch "a" {
    case "\n\t":
        v6 := []int{}
        _ = v6
    default:
        p2 -= (v5 * (p2 >> ((1) & 63)))
    }
    v7 := []int{802, 18, 128, int(v5)}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return ("\n\t" + "hello")
}
func genFn1() int64 {
    v8 := []int{9223372036854775807, 1}
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int = 2147483647
        var v11 float32 = (1.5 * float32(float64((1.5 * (-2.25 / -2.25)))))
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v8
    return -125
}

func main() {
    var v1 int64 = 1
    v2 := make(map[int]bool)
    var v3 int64 = v1
    var v4 float32 = 1e10
    for fi12 := 0; fi12 < 19; fi12++ {
        v1 = 9223372036854775807
        delete(v2, fi12)
        _ = fi12
    }
    var v13 bool = false
    fmt.Println(v1, v13, v4)
    fmt.Println(v13, v1)
    var v14 uint8 = 195
    var v15 string = ("\n\t" + "foo bar")
    var v16 uint32 = 92
    var v17 int = 1
    switch v15 {
    case "\n\t":
        var v18 uint32 = 250
        _ = v18
    }
    fmt.Println(v16, v16, v16)
    var v19 int = v17
    v14 = v14
    var v20 int8 = 100
    switch (v15 + "a") {
    case "hello":
        var v21 float64 = ((0.1 * -2.25) + 1.5)
        _ = v21
    case "foo bar":
        v19 = (v17 | 0)
    case "a":
        for fi22 := 0; fi22 < 6; fi22++ {
            v23 := GenS1{F0: (255 <= 65535), F1: "a"}
            v24 := GenS1{F0: v13, F1: "foo bar"}
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    var v25 float32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
}
