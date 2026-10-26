package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    var v1 int16 = 1000
    var v2 uint8 = 255
    var v3 float64 = 1.5
    var v4 int = 32768
    v4 += v4
    if (1 < 1) {
        v3 = v3
        if (-1 <= 42) {
            var v5 string = "foo bar"
            for fi6 := 0; fi6 < 11; fi6++ {
                fmt.Println(fi6)
                _ = fi6
            }
            _ = v5
        } else {
            fmt.Println(v1)
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 float32 = float32((v3 * 0.1))
                fmt.Println(v4, v4, v1)
                _ = fi7
                _ = v8
            }
            var v9 uint = 2
            _ = v9
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 int64 = 518
                _ = fi11
                _ = v12
            }
            v1 = v1
            _ = fi10
        }
    } else {
        v13 := []int{(-(v4)), (v4 & -2147483648), (-(v4))}
        v1 += v1
        _ = v13
    }
    v1 = -1
    var v14 int32 = 1
    if false {
        v15 := []int{(^(v4)), 1, v4}
        _ = v15
    }
    v16 := []int{127, v4, v4, v4}
    v17 := GenS1{F0: (^(v2)), F1: 506}
    v3 -= -2.25
    if false {
        fmt.Println(v14)
        fmt.Println(v14)
        v3 = ((v3 + v3) * v3)
    } else {
        var v18 uint32 = 602
        var v19 int16 = v1
        v4 = v4
        _ = v18
        _ = v19
    }
    var v20 int16 = (v1 & -944)
    var v21 bool = ((true && true) || true)
    switch "a" {
    case "\n\t":
        fmt.Println(v21)
    case "hello":
        v2 += v2
    }
    v14 -= v14
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    var v22 float64 = v3
    if len(v16) > 0 {
        v16[0] = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
