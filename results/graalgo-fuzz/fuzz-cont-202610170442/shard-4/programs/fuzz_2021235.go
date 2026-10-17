package main

import (
    "fmt"
)

func genFn0(p0 int64, p1 string, p2 uint) uint8 {
    var v4 uint32 = 2147483647
    p1 = ("hello" + p1)
    fmt.Println(p2)
    var v5 uint32 = 419
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 255
}

func main() {
    var v1 int32 = -125
    var v2 float32 = float32((3.14159 * (0.1 - 0.0)))
    v3 := []string{""}
    switch -825 {
    case -573:
        if (-1 < -119) {
            var v6 uint64 = 42
            var v7 int8 = 100
            var v8 int16 = 42
            _ = v6
            _ = v7
            _ = v8
        } else {
            var v9 int = -2147483648
            _ = v9
        }
    }
    var v10 string = "a"
    for fi11 := 0; fi11 < 10; fi11++ {
        v1 = (v1 >> ((487) & 63))
        _ = fi11
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        v2 = (v2 * 1e10)
        v2 = (((v2 + v2) * (-2.25 - 0.0)) * 1e10)
        var v13 bool = (127 > 151)
        _ = fi12
        _ = v13
    }
    v1 += (v1 / ((v1) | 1))
    fmt.Println(v10, v1)
    for fi14 := 0; fi14 < 14; fi14++ {
        v1 += v1
        _ = fi14
    }
    v2 += v2
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
