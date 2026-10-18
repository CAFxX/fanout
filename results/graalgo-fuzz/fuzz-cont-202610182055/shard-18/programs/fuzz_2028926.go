package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint16 }
func genFn0() int32 {
    var v7 uint16 = 1
    if (true && true) {
        fmt.Println(v7)
        var v8 float32 = (float32((((0.0 + (-2.25 / -2.25)) + 0.0) / -2.25)) * ((-2.25 - 0.0) - 1e10))
        _ = v8
    }
    v9 := []int{2147483647}
    v10 := GenS1{F0: 1, F1: 1, F2: (v7 | (v7 * 32767))}
    v11 := GenS1{F0: 42, F1: v7, F2: (v7 >> ((0) & 63))}
    _ = v10
    _ = v11
    _ = v7
    _ = v9
    return -1
}

func main() {
    var v1 bool = (true && (true || ((float32(0.1) - (-2.25 - 1e10)) > 0.1)))
    v2 := make(map[uint32]bool)
    var v3 int = 128
    var v4 uint32 = 273
    v5 := make(map[string]bool)
    var v6 float32 = ((0.1 * (1e10 * (1.5 * 1e10))) + float32((-2.25 / (float64((3.14159 - 0.0)) / 1e10))))
    v5["hello"] = (v3 == 0)
    fmt.Println(v3, v6, v4)
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 int8 = 97
        _ = fi12
        _ = v13
    }
    v1 = false
    var v14 float32 = (1e10 - (v6 * (-2.25 - (float32(1e10) - float32(-2.25)))))
    v3 = -1
    var v15 uint64 = 50
    var v16 float32 = (float32(0.0) - 1.5)
    if ((!(42 >= -2147483648)) || v1) {
        for fi17 := 0; fi17 < 8; fi17++ {
            if (!true) {
                fmt.Println(v15, v1, v1)
            } else {
                for fi18 := 0; fi18 < 12; fi18++ {
                    var v19 uint64 = (v15 ^ (v15 * 42))
                    if v1 {
                        v20 := GenS1{F0: 224, F1: 957, F2: 0}
                        fmt.Println(fi18, v14, v15)
                        v19 += uint64(v4)
                        _ = v20
                    }
                    _ = fi18
                    _ = v19
                }
                v16 = v14
            }
            v14 = v14
            _ = fi17
        }
        var v21 string = ("a" + "")
        _ = v21
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
