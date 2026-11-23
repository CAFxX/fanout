package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 int, p1 int64, p2 uint8) string {
    if false {
        p0 = 875
    } else {
        v5 := []int{p0, p0}
        for fi6 := 0; fi6 < 2; fi6++ {
            p1 += (-(p1))
            _ = fi6
        }
        _ = v5
    }
    var v7 float64 = 1e10
    v8 := GenS1{F0: 0}
    if true {
        for fi9 := 0; fi9 < 7; fi9++ {
            if false {
                var v10 int32 = 590
                _ = v10
            } else {
                v11 := []int{fi9, p0}
                _ = v11
            }
            p2 += p2
            p0 -= p0
            _ = fi9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ("a" + ((("foo bar" + (("\n\t" + "hello") + "a")) + "a") + "a"))
}

func main() {
    var v1 uint = 802
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 int16 = -217
    v12 := GenS1{F0: 65535}
    if (42 <= -988) {
        v4 = (v4 & (^(v4)))
        fmt.Println(v4, v4)
    } else {
        var v13 int = 2147483647
        v14 := GenS1{F0: 1}
        _ = v13
        _ = v14
    }
    if (0 >= 41) {
        v15 := []int{-2147483648}
        switch 434 {
        case 42:
            v16 := GenS2{F0: 240, F1: 683}
            _ = v16
        case -448:
            var v17 uint = v1
            _ = v17
        }
        if false {
            var v18 int16 = v4
            _ = v18
        }
        _ = v15
    }
    v4 += v4
    var v19 int64 = 229
    v20 := GenS2{F0: 127, F1: -2147483648}
    if false {
        if false {
            v1 += (v1 * v1)
        }
        switch 1 {
        case 128:
            v21 := []int{127, 42, -9223372036854775808}
            _ = v21
        }
    }
    var v22 uint64 = 0
    var v23 float64 = 0.1
    fmt.Println(v23, v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
