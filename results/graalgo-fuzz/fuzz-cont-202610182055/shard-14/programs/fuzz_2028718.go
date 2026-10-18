package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() int16 {
    var v6 int32 = 0
    if (-1 == 292) {
        v6 = (v6 - v6)
        var v7 float32 = 1e10
        for fi8 := 0; fi8 < 10; fi8++ {
            v7 -= v7
            _ = fi8
        }
        _ = v7
    }
    _ = v6
    return -658
}

func main() {
    var v1 uint16 = 1
    var v2 uint8 = 127
    v3 := make(map[string]uint32)
    var v4 float32 = 0.0
    var v5 float64 = 0.1
    var v9 uint8 = (v2 & 255)
    v9 = (v9 - v2)
    v10 := []int{42, 610, -9223372036854775808}
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 int32 = 864
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: ("" + "foo bar")}
    v14 := v3[("" + "\n\t")]
    var v15 int64 = 0
    var v16 int = 137
    fmt.Println(v4, v1, v9)
    if (!true) {
        fmt.Println(v14, v14, v5)
        fmt.Println(v1)
        var v17 string = "hello"
        _ = v17
    }
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
