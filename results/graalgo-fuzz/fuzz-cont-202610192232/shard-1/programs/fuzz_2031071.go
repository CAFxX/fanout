package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 65535
    var v3 uint16 = 495
    v4 := make(map[string]string)
    v5 := make(map[int]uint32)
    var v6 string = "a"
    var v7 float32 = 1e10
    v8 := GenS1{F0: 2147483647, F1: (!(false && (-9223372036854775808 > 0)))}
    v9 := []int{4294967295, int(v3), -248}
    var v10 int8 = 61
    for fi11 := 0; fi11 < 4; fi11++ {
        if true {
            v12 := v4[v6]
            if ((!false) && true) {
                var v13 float64 = ((0.1 * (1.5 / float64(float32(-2.25)))) * float64(1.5))
                v14 := GenS2{F0: 42}
                _ = v13
                _ = v14
            }
            v7 = 1.5
            _ = v12
        }
        _ = fi11
    }
    var v15 int8 = (-(v10))
    var v16 bool = ((!true) && (650 >= 0))
    var v17 string = v6
    v18 := GenS2{F0: v15}
    fmt.Println(v6, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
