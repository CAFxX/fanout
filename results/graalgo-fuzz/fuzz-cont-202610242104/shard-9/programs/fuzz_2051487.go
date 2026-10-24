package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 float32 = 3.14159
    var v3 float64 = (-2.25 / float64(1e10))
    v4 := make(map[int]string)
    var v5 float64 = -2.25
    var v6 int32 = -1
    for fi7 := 0; fi7 < 17; fi7++ {
        if (0 <= 1) {
            v5 -= v3
            v8 := []int{(fi7 << ((287) & 63)), fi7, 0}
            _ = v8
        }
        v9 := GenS1{F0: 28}
        v10 := GenS1{F0: 27}
        _ = fi7
        _ = v10
        _ = v9
    }
    var v11 bool = (46 > 2147483648)
    v2 -= v2
    v3 = v3
    var v12 uint = 65535
    var v13 float32 = (3.14159 * (v2 - (0.0 / ((v2 - float32(v5)) * v2))))
    delete(v1, "a")
    v14 := GenS1{F0: int16(v6)}
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 float64 = v3
        _ = fi15
        _ = v16
    }
    var v17 uint64 = 595
    if true {
        var v18 int16 = 0
        v12 += (v12 & (^(v12)))
        v19 := []int{480}
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        fmt.Println(v2, v6)
        _ = fi20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 uint32 = 1
        v23 := GenS1{F0: -1}
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: 1000}
    switch -2147483648 {
    case 32768:
        v25 := v4[-249]
        _ = v25
    case 9223372036854775807:
        _ = len(v4)
    default:
        delete(v4, 127)
    }
    var v26 string = "foo bar"
    v26 = v26
    v5 -= v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
