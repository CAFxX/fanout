package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    var v1 uint8 = 42
    var v2 uint16 = 137
    v3 := []int{}
    v4 := make(map[uint32]int)
    fmt.Println(v1, v1)
    if true {
        for fi5 := 0; fi5 < 13; fi5++ {
            v1 -= (v1 & (^(v1)))
            if false {
                var v6 uint64 = 1
                v7 := GenS1{F0: 36}
                v6 += 42
                _ = v6
                _ = v7
            } else {
                for fi8 := 0; fi8 < 16; fi8++ {
                    var v9 uint64 = 1
                    var v10 uint = 42
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                v2 += v2
            }
            _ = fi5
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi11
        }
        v1 += 255
    }
    v1 += v1
    v1 += (v1 << ((18446744073709551615) & 63))
    v1 -= (^(v1))
    v2 += v2
    v2 = (v2 | 65535)
    v2 -= 621
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            _ = len(v3)
            v1 -= (v1 ^ 42)
            _ = fi13
        }
        if (-2.25 >= 1e10) {
            var v14 int8 = -128
            var v15 float32 = (float32(0.0) - 1.5)
            _ = v14
            _ = v15
        } else {
            var v16 float64 = 0.0
            _ = v16
        }
        _ = fi12
    }
    v17 := GenS2{F0: "a", F1: 0}
    v2 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
