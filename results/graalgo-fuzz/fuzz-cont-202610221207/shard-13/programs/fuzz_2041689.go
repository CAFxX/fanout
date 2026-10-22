package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 0
    var v3 int16 = v2
    var v4 int16 = (v2 * v3)
    var v5 bool = true
    var v6 string = ("foo bar" + "foo bar")
    switch 758 {
    case 1:
        v1["a"] = (9223372036854775807 < -213)
    case 128:
        var v7 uint = 4294967295
        _ = v7
    case -1:
        var v8 uint32 = 838
        _ = v8
    }
    v2 -= 32767
    for fi9 := 0; fi9 < 9; fi9++ {
        if ((!(((0.0 + 0.0) * ((1.5 * float32(3.14159)) / 1.0)) == (0.0 - float32(1.5)))) || v5) {
            v10 := GenS1{F0: -1}
            for fi11 := 0; fi11 < 19; fi11++ {
                v3 = 850
                v4 = v4
                fmt.Println(fi9, fi11)
                _ = fi11
            }
            _ = v10
        }
        _ = fi9
    }
    var v12 float32 = 1e10
    v13 := v1[(v6 + v6)]
    var v14 int64 = 0
    var v15 int64 = (v14 | -505)
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := GenS1{F0: 100}
        _ = fi16
        _ = v17
    }
    v18 := []int{-9223372036854775808, -992, 0}
    var v19 uint16 = 255
    var v20 int = -9223372036854775808
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
