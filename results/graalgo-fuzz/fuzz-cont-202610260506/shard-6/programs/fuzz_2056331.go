package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint8 }
func genFn0(p0 uint64) uint16 {
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := GenS1{F0: 1, F1: 0, F2: 0}
        p0 = p0
        _ = fi4
        _ = v5
    }
    v6 := []int{227}
    p0 = p0
    _ = p0
    _ = v6
    return 32
}

func main() {
    v1 := []string{"", ("a" + "\n\t"), ("\n\t" + "a")}
    var v2 float32 = 0.0
    var v3 int32 = 0
    v7 := GenS1{F0: -559, F1: 42, F2: 1}
    v2 = (v2 + ((v2 * v2) * (0.1 / float32(3.14159))))
    v2 = v2
    var v8 float32 = 0.0
    if (true && true) {
        var v9 uint8 = 128
        _ = v9
    } else {
        if false {
            v2 -= v8
            fmt.Println(v3, v8, v8)
        } else {
            if false {
                v3 = (v3 << ((868) & 63))
                if false {
                    v2 -= float32(1e10)
                    for fi10 := 0; fi10 < 8; fi10++ {
                        v3 -= (v3 * v3)
                        _ = fi10
                    }
                }
                v8 = 1.5
            }
            v3 += 2147483647
            var v11 int64 = -1
            _ = v11
        }
    }
    v12 := GenS1{F0: -621, F1: 0, F2: 2}
    v8 += (v2 / v2)
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
