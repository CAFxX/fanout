package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int }

func main() {
    var v1 int = 1
    var v2 uint16 = 42
    v3 := make(map[string]int)
    var v4 uint8 = 239
    v5 := make(map[uint32]int)
    var v6 uint32 = 2
    v4 += 128
    v5[v6] = (v1 << ((830) & 63))
    if (!false) {
        var v7 bool = true
        v8 := []int{v1}
        _ = v7
        _ = v8
    }
    v4 -= v4
    v2 += (v2 & v2)
    fmt.Println(v2)
    v9 := []int{v1, v1, (v1 % ((42) | 1)), (v1 << ((uint(v6)) & 63))}
    var v10 uint64 = 0
    switch v1 {
    case 0:
        v9 = append(v9, (-(v1)))
    case 106:
        for fi11 := 0; fi11 < 15; fi11++ {
            if (65535 >= 4294967295) {
                _ = len(v3)
            }
            for fi12 := 0; fi12 < 6; fi12++ {
                v1 = 127
                if ((v6 % (((^(v6))) | 1)) < 42) {
                    v13 := GenS1{F0: (v6 >> ((uint(v10)) & 63)), F1: "a", F2: (fi12 - int(v10))}
                    _ = len(v9)
                    _ = v13
                } else {
                    var v14 uint16 = v2
                    v6 += v6
                    v5[v6] = fi12
                    _ = v14
                }
                _ = fi12
            }
            _ = fi11
        }
    case 88:
        var v15 int64 = 42
        _ = v15
    default:
        v5[v6] = (^(v1))
    }
    v6 += (v6 * (v6 & v6))
    switch (v1 - (v1 | (v1 * (v1 & (v1 & -342))))) {
    case 32767:
        _ = len(v9)
    case 32768:
        if (v1 > 32768) {
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 float32 = 1.5
                if (!(!true)) {
                    switch v1 {
                    case 0:
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    case -47:
                        v10 = v10
                    default:
                        v18 := GenS1{F0: v6, F1: "hello", F2: v1}
                        _ = v18
                    }
                    var v19 uint = 4294967295
                    _ = v19
                }
                delete(v5, (v6 ^ v6))
                _ = fi16
                _ = v17
            }
        }
    }
    var v20 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
