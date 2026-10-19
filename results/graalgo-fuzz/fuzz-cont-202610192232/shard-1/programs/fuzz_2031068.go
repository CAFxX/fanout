package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int64) uint {
    v5 := []int{127, 4294967295}
    if false {
        fmt.Println(p0)
    }
    v6 := GenS1{F0: 588, F1: 0, F2: 32767}
    _ = p0
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 string = ((("" + "foo bar") + "foo bar") + "foo bar")
    v2 := make(map[uint32]uint32)
    var v3 float32 = ((-2.25 / 1.0) * (float32(3.14159) / (0.1 / ((0.1 - 1e10) + 0.0))))
    var v4 float32 = v3
    v7 := GenS1{F0: 733, F1: 108, F2: 1}
    v3 = (v3 / 1.5)
    v1 = (v1 + (("hello" + v1) + v1))
    if (42 >= -933) {
        fmt.Println(v4, v3, v3)
        var v8 int64 = 9223372036854775807
        for fi9 := 0; fi9 < 2; fi9++ {
            v4 -= ((float32(1e10) - 0.0) / ((float32((3.14159 * ((3.14159 + 3.14159) * ((1e10 + -2.25) - 1e10)))) / v4) * (v3 - (1.5 * v4))))
            _ = fi9
        }
        _ = v8
    }
    delete(v2, 2)
    v3 -= ((v4 / 1.0) + float32(0.1))
    var v10 bool = true
    switch (((("a" + v1) + (v1 + "hello")) + v1) + v1) {
    case "foo bar":
        if (!(758 <= 4294967295)) {
            if (0.1 < 1.5) {
                for fi11 := 0; fi11 < 16; fi11++ {
                    var v12 int = 9223372036854775807
                    v4 = 0.0
                    var v13 uint64 = 658
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            } else {
                if false {
                    if v10 {
                        v14 := GenS2{F0: 80}
                        _ = v14
                    }
                    delete(v2, 357)
                    v1 = (v1 + v1)
                }
            }
            fmt.Println(v4, v4)
        }
    case "":
        fmt.Println(v4, v4)
    default:
        var v15 bool = (!(-66 <= -2147483648))
        _ = v15
    }
    var v16 uint16 = 255
    switch 32767 {
    case 515:
        var v17 float32 = (v4 * (0.1 * (v4 / v3)))
        _ = v17
    case -876:
        delete(v2, 1)
    }
    if ((-2.25 * (float64(1e10) + 0.1)) > float64(-2.25)) {
        var v18 uint8 = 42
        _ = v18
    }
    v1 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
