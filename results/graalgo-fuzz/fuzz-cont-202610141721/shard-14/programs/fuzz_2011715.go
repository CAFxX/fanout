package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 float64 = (((1.5 - (0.1 - 3.14159)) + (1e10 + 0.0)) - float64(1e10))
    v2 := []int{1, -2147483648}
    var v3 string = "foo bar"
    v4 := make(map[int]string)
    var v5 float64 = 0.0
    var v6 uint32 = 2
    if (1e10 == ((1.5 + (-2.25 + 1.5)) * 1.5)) {
        var v7 bool = false
        for fi8 := 0; fi8 < 4; fi8++ {
            v1 += (1.5 - v5)
            var v9 int16 = 42
            _ = fi8
            _ = v9
        }
        v1 -= 0.0
        _ = v7
    }
    var v10 uint8 = 2
    v5 = v1
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: -665}
        _ = fi11
        _ = v12
    }
    fmt.Println(v3)
    v3 = (v3 + v3)
    var v13 uint32 = 325
    var v14 int16 = 42
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 uint16 = 65535
        _ = fi15
        _ = v16
    }
    var v17 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
