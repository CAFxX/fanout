package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int32 = -2147483648
    var v2 uint32 = 1
    v3 := make(map[uint32]string)
    var v4 float64 = 1e10
    var v5 int16 = -465
    v2 -= v2
    v6 := []int{0}
    var v7 float64 = ((v4 + (float64(float32(v4)) * float64(1.5))) + 0.0)
    v1 = v1
    var v8 int = 128
    delete(v3, (^(v2)))
    var v9 uint64 = 4294967295
    fmt.Println(v7)
    var v10 int32 = (v1 ^ v1)
    fmt.Println(v7)
    for fi11 := 0; fi11 < 5; fi11++ {
        if false {
            v7 = ((-2.25 / (v4 * (v4 / 1.5))) / (0.0 + v4))
        } else {
            fmt.Println(v7, fi11, v9)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v12 uint64 = (v9 / ((4294967295) | 1))
            _ = v12
        }
        _ = fi11
    }
    if false {
        v4 -= (v4 * 0.1)
        var v13 int32 = (v1 / ((v10) | 1))
        if true {
            v14 := v3[(v2 * v2)]
            v15 := GenS1{F0: -128}
            delete(v3, 2)
            _ = v14
            _ = v15
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            delete(v3, v2)
        }
        _ = v13
    } else {
        if len(v6) > 0 {
            v6[0] = v8
        }
        v16 := GenS1{F0: -1}
        var v17 string = ("\n\t" + "hello")
        _ = v16
        _ = v17
    }
    var v18 int8 = 100
    var v19 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
