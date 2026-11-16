package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint16 }

func main() {
    var v1 int = -361
    var v2 float32 = 1e10
    var v3 float64 = 3.14159
    v4 := make(map[string]uint32)
    var v5 int64 = -9223372036854775808
    var v6 int64 = v5
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := []int{}
        v3 = (((v3 * v3) + 3.14159) * v3)
        _ = fi7
        _ = v8
    }
    fmt.Println(v1, v6)
    var v9 string = ""
    for fi10 := 0; fi10 < 16; fi10++ {
        if true {
            v5 = v6
        } else {
            fmt.Println(v3, v6)
            fmt.Println(v9, v3, v5)
            fmt.Println(v9)
        }
        v2 = v2
        var v11 bool = false
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        if false {
            var v13 int32 = -461
            _ = v13
        }
        _ = fi12
    }
    if false {
        var v14 int16 = 32767
        v9 = "hello"
        var v15 float32 = (((v2 + v2) * float32(3.14159)) - float32((float64(v2) + v3)))
        _ = v14
        _ = v15
    }
    var v16 int64 = (-(v5))
    v17 := GenS1{F0: (!((!(!(919 > 0))) && false)), F1: 1, F2: 123}
    v2 = (v2 - (0.0 * v2))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
