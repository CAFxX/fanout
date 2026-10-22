package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
func genFn0(p0 uint16, p1 int, p2 uint32) bool {
    var v6 float64 = 0.1
    v7 := GenS1{F0: ((!true) || (((0.1 + -2.25) == 1e10) || (((^(p0)) >= 0) && (243 < 0)))), F1: -128, F2: 42}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (0 < 760)
}

func main() {
    var v1 uint32 = 934
    v2 := []int{451}
    v3 := make(map[string]string)
    var v4 int64 = 42
    var v5 float32 = 1.5
    switch 1 {
    case -635:
        v5 -= float32(3.14159)
    case 0:
        var v8 float32 = ((float32(0.1) + v5) + float32((-2.25 - (0.1 / -2.25))))
        _ = v8
    case 906:
        if true {
            var v9 int32 = 0
            _ = v9
        }
    }
    var v10 float32 = (1e10 / v5)
    for fi11 := 0; fi11 < 1; fi11++ {
        _ = len(v3)
        v4 = (v4 - int64(fi11))
        _ = fi11
    }
    var v12 uint16 = 379
    v5 -= -2.25
    var v13 uint = 255
    var v14 float64 = 0.1
    var v15 int32 = 42
    var v16 uint8 = 1
    v17 := GenS1{F0: false, F1: 42, F2: (v12 | (v12 << (((v13 % ((65535) | 1))) & 63)))}
    for fi18 := 0; fi18 < 12; fi18++ {
        _ = len(v2)
        v13 = v13
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
