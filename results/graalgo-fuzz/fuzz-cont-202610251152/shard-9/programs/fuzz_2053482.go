package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 bool }
func genFn0(p0 int8, p1 int) bool {
    var v6 float32 = (1e10 - (-2.25 * 0.1))
    v7 := GenS1{F0: ("a" + "hello"), F1: 42, F2: true}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return false
}

func main() {
    var v1 int = 32767
    v2 := []int{(v1 << ((65535) & 63))}
    var v3 uint32 = 2
    var v4 float32 = 1e10
    var v5 uint32 = 65535
    var v8 bool = (((2147483647 != -215) || true) && false)
    var v9 string = "hello"
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v10 := GenS1{F0: v9, F1: 401, F2: v8}
    var v11 float64 = -2.25
    for fi12 := 0; fi12 < 4; fi12++ {
        if v8 {
            if (-9223372036854775808 >= 1) {
                var v13 float64 = v11
                _ = v13
            } else {
                var v14 uint64 = 562
                var v15 float64 = 0.1
                if (348 < 255) {
                    var v16 bool = (42 != -980)
                    _ = v16
                } else {
                    var v17 float32 = float32(0.0)
                    _ = v17
                }
                _ = v14
                _ = v15
            }
            fmt.Println(v1, v8, v5)
        } else {
            for fi18 := 0; fi18 < 13; fi18++ {
                var v19 int = (^(fi12))
                _ = fi18
                _ = v19
            }
        }
        fmt.Println(v11, v8)
        _ = fi12
    }
    var v20 bool = v8
    v8 = (2 <= 1)
    fmt.Println(v3)
    fmt.Println(v1, v9, v4)
    var v21 uint32 = (v5 + 0)
    v3 -= (v21 << ((18446744073709551615) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
