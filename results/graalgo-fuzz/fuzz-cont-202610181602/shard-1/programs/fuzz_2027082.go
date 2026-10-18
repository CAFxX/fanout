package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 int16 = -1
    var v3 string = "hello"
    v4 := []int{42, 127, 2147483647}
    if (1 < -128) {
        _ = len(v1)
        v2 += 1000
        if false {
            v2 += (v2 >> ((612) & 63))
            v5 := GenS1{F0: -269, F1: v2}
            if false {
                fmt.Println(v2)
                v3 = v3
                v6 := GenS1{F0: -1, F1: v2}
                _ = v6
            } else {
                for fi7 := 0; fi7 < 19; fi7++ {
                    var v8 float32 = 0.1
                    _ = fi7
                    _ = v8
                }
                v3 = (v3 + v3)
            }
            _ = v5
        }
    } else {
        fmt.Println(v2)
    }
    var v9 uint32 = 1
    v2 -= v2
    delete(v1, v3)
    delete(v1, "a")
    delete(v1, ("a" + "hello"))
    v2 += (v2 - (v2 * (^(v2))))
    v3 = v3
    var v10 bool = false
    if v10 {
        v2 -= (^(v2))
        if (1 != v2) {
            fmt.Println(v3)
        }
        v9 += v9
    }
    v9 -= v9
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    var v11 int8 = -93
    var v12 uint16 = 32767
    if v10 {
        v12 += uint16(v9)
    }
    var v13 int32 = 2147483647
    fmt.Println(v2, v11, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
