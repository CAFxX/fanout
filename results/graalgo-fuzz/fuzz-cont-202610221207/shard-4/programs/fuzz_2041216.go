package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint8; F1 int8; F2 string }

func main() {
    v1 := make(map[string]string)
    var v2 string = (("foo bar" + "a") + "")
    var v3 float64 = 1e10
    v4 := make(map[uint32]int)
    var v5 float32 = (0.1 + (1e10 / float32(((v3 + (v3 / (v3 + 0.1))) - float64(3.14159)))))
    for fi6 := 0; fi6 < 20; fi6++ {
        v2 = v2
        v3 = (v3 - v3)
        var v7 float32 = v5
        _ = fi6
        _ = v7
    }
    v5 -= v5
    fmt.Println(v2, v2)
    fmt.Println(v5)
    v8 := GenS1{F0: 0, F1: 697}
    var v9 int64 = 463
    var v10 uint16 = 992
    v11 := GenS2{F0: 191, F1: -104, F2: ((v2 + v2) + "\n\t")}
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
