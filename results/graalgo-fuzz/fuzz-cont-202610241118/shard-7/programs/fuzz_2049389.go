package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint }

func main() {
    var v1 float32 = (0.0 + (0.1 - 0.1))
    var v2 uint8 = 201
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 float32 = (v1 * v1)
    var v6 uint64 = 870
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint8 = (v2 * v2)
        v9 := v3[(^(fi7))]
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(v6, v1)
        var v11 uint = 1
        _ = fi10
        _ = v11
    }
    v6 = 0
    if (!true) {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 int32 = 0
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            for fi15 := 0; fi15 < 19; fi15++ {
                v1 = 1.5
                _ = fi15
            }
            v1 = (float32(-2.25) - (v5 * ((v1 * v5) * v1)))
            _ = fi14
        }
        var v16 float64 = 1e10
        _ = v16
    } else {
        v17 := GenS1{F0: v6, F1: ((true && ("a" < "hello")) || (v6 == (^(v6)))), F2: (v2 >> ((987) & 63))}
        v5 = v5
        for fi18 := 0; fi18 < 20; fi18++ {
            fmt.Println(v6, v1)
            _ = fi18
        }
        _ = v17
    }
    v6 = (v6 ^ v6)
    v19 := GenS2{F0: 622, F1: 42, F2: 42}
    var v20 int = 127
    if (45 == (v6 ^ 2)) {
        v21 := []int{(^(v20)), (v20 + v20)}
        var v22 string = "a"
        _ = v21
        _ = v22
    } else {
        var v23 bool = ("foo bar" > ("a" + ""))
        fmt.Println(v2, v23)
        var v24 int16 = -32768
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
