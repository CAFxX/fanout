package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
func genFn0(p0 uint16) uint {
    var v4 bool = (400 < 127)
    switch "\n\t" {
    case "\n\t":
        switch 32767 {
        case 1:
            var v5 bool = v4
            _ = v5
        case 4294967295:
            var v6 string = "foo bar"
            _ = v6
        case 9223372036854775807:
            v7 := []int{661, -18, -1}
            _ = v7
        default:
            v8 := []int{42}
            _ = v8
        }
    }
    p0 = p0
    if (true || true) {
        fmt.Println(v4, p0)
        v9 := GenS1{F0: -9223372036854775808, F1: 2147483647}
        _ = v9
    }
    var v10 int32 = 803
    _ = p0
    _ = v10
    _ = v4
    return 355
}

func main() {
    var v1 uint32 = 619
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    v11 := []int{}
    var v12 int = 0
    var v13 float64 = float64(0.0)
    var v14 bool = ((65 <= -128) && false)
    for fi15 := 0; fi15 < 2; fi15++ {
        fmt.Println(v14)
        var v16 int = (v12 * v12)
        _ = fi15
        _ = v16
    }
    v14 = v14
    fmt.Println(v14)
    var v17 int16 = 42
    if (((v13 + float64((float32(((0.0 + 0.0) - v13)) / 1.5))) * (((0.1 * (float64((1.5 - 0.1)) - 0.1)) / 0.1) - ((-2.25 + (v13 - ((v13 / ((3.14159 * v13) + (v13 + ((0.1 + (0.0 - (3.14159 / float64(float32(v13))))) + v13)))) / v13))) / v13))) != v13) {
        _ = len(v11)
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 string = (("" + "") + (("" + ("a" + "")) + ("foo bar" + ((("a" + "hello") + "hello") + ""))))
            v13 = 0.0
            var v20 float64 = 3.14159
            _ = fi18
            _ = v19
            _ = v20
        }
        v14 = true
    }
    var v21 float32 = -2.25
    v22 := GenS1{F0: 0, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
