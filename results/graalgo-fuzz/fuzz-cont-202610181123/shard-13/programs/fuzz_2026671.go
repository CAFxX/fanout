package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 string; F1 int; F2 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 0.1
    var v3 uint32 = 839
    v4 := make(map[string]bool)
    var v5 uint64 = 0
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 string = ("\n\t" + "a")
            _ = fi7
            _ = v8
        }
        var v9 float64 = 0.0
        _ = fi6
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            v4[("hello" + "a")] = (42 <= -629)
            fmt.Println(v3)
            _ = fi11
        }
        if true {
            v3 += (v3 / (((v3 + (v3 << ((uint(v3)) & 63)))) | 1))
            v12 := []int{fi10, (fi10 - fi10), -9223372036854775808, fi10}
            var v13 uint = 18446744073709551615
            _ = v12
            _ = v13
        }
        v4["foo bar"] = true
        _ = fi10
    }
    var v14 int32 = 2147483647
    if (-2.25 < (0.1 - (v2 + (v2 + v2)))) {
        switch 42 {
        case -621:
            v15 := GenS1{F0: 9223372036854775807, F1: 679}
            _ = v15
        default:
            v16 := GenS2{F0: "", F1: -715, F2: 2}
            _ = v16
        }
        v2 = v2
    } else {
        fmt.Println(v3)
        v17 := []int{-718, -861}
        _ = v17
    }
    var v18 int = -539
    v5 = (v5 ^ (v5 & v5))
    switch (v18 - v18) {
    case 2147483647:
        if true {
            fmt.Println(v14, v2, v2)
            v19 := v4[("" + "a")]
            fmt.Println(v5, v19)
            _ = v19
        } else {
            v20 := []int{(v18 * -510)}
            _ = v20
        }
    case 32767:
        v5 = (v5 + 42)
    }
    v21 := []int{-1}
    v5 += 606
    if false {
        fmt.Println(v14)
        v21 = append(v21, (v18 | (v18 | (v18 + v18))))
    }
    var v22 int32 = (v14 + v14)
    switch (("\n\t" + "hello") + "a") {
    case "a":
        var v23 uint8 = 127
        _ = v23
    case "hello":
        var v24 string = "hello"
        _ = v24
    case "":
        var v25 float32 = (float32(1.5) / v2)
        _ = v25
    default:
        fmt.Println(v14)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
