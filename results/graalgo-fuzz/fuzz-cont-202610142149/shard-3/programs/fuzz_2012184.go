package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int64; F2 uint8 }
func genFn0(p0 uint64) uint {
    v4 := GenS1{F0: 812, F1: 255, F2: 42}
    v5 := GenS2{F0: 0}
    p0 = 1
    v6 := GenS3{F0: 1, F1: 1, F2: 2}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 862
}

func main() {
    var v1 uint32 = 65535
    var v2 uint32 = (v1 << ((400) & 63))
    var v3 int64 = 42
    v7 := []int{int(v2)}
    v8 := GenS1{F0: 289, F1: 88, F2: v3}
    v3 = (v3 & v3)
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int16 = 42
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        if (2 >= 2) {
            v3 = (v3 ^ (v3 + v3))
            v2 -= (^(v2))
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        } else {
            v3 -= (v3 % ((v3) | 1))
            var v12 uint32 = 2147483647
            _ = v12
        }
        if (0 <= -933) {
            if (0 <= 0) {
                if (!true) {
                    v13 := GenS3{F0: 2147483647, F1: 0, F2: 127}
                    _ = v13
                }
            }
            if (!((0.1 != (0.0 - 3.14159)) && true)) {
                var v14 int = 96
                _ = v14
            }
        }
        var v15 float32 = 0.0
        _ = fi11
        _ = v15
    }
    var v16 uint16 = 338
    v7 = append(v7, -139)
    var v17 int8 = 127
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := GenS1{F0: v16, F1: 0, F2: v3}
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
