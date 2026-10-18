package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int) uint16 {
    var v6 int64 = 0
    var v7 int16 = -32768
    _ = p0
    _ = v6
    _ = v7
    return 286
}
func genFn1(p0 uint8, p1 uint32) int32 {
    if true {
        v8 := GenS1{F0: 65535}
        var v9 string = "hello"
        _ = v8
        _ = v9
    }
    if false {
        var v10 uint16 = 1
        _ = v10
    } else {
        var v11 float64 = float64(0.0)
        _ = v11
    }
    if false {
        v12 := []int{1, 2147483647, 755, int(p1)}
        _ = v12
    }
    v13 := []int{32768, -2147483648, 127}
    _ = p0
    _ = p1
    _ = v13
    return 42
}

func main() {
    var v1 bool = true
    var v2 int64 = 1
    var v3 uint16 = 423
    var v4 float32 = -2.25
    var v5 int16 = int16(v2)
    v14 := []int{2147483647}
    var v15 int16 = v5
    v3 = 32767
    var v16 int8 = 100
    v16 -= -1
    v3 = 1
    fmt.Println(v3, v15, v16)
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            if (1 < 128) {
                if len(v14) > 0 {
                    v14[0] = -510
                }
                v1 = (840 <= v3)
                if len(v14) > 0 {
                    v14[0] = 42
                }
            }
            _ = fi18
        }
        v2 += (v2 - (v2 ^ -1))
        _ = fi17
    }
    var v19 uint32 = 162
    v2 += (^(v2))
    var v20 uint = 2
    if len(v14) > 0 {
        v14[0] = 0
    }
    fmt.Println(v19, v1)
    var v21 uint16 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
