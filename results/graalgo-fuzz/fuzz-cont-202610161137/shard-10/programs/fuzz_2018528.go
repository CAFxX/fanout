package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 365
    v3 := make(map[int]string)
    var v4 int8 = 127
    var v5 string = ""
    var v6 float64 = ((((float64(1.5) * (3.14159 * ((float64(0.0) / 0.1) - 1.5))) - 1e10) / 0.1) * 0.1)
    fmt.Println(v4, v2, v4)
    switch 675 {
    case -1:
        delete(v1, -988)
    }
    fmt.Println(v2)
    v4 += int8(v2)
    fmt.Println(v2, v2, v6)
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint8 = 0
        var v9 uint32 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int64 = -171
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 string = v5
        _ = fi11
        _ = v12
    }
    v6 -= v6
    var v13 uint32 = 4294967295
    v2 = v2
    v2 += v2
    var v14 int32 = 1
    v15 := GenS1{F0: -784, F1: ("" + ((v5 + v5) + v5)), F2: 42}
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
