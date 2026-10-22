package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint64; F1 int8; F2 int16 }
func genFn0(p0 int8) uint {
    var v7 string = "foo bar"
    v8 := GenS1{F0: 32767, F1: -9223372036854775808}
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint16 = 693
    v2 := make(map[int]int)
    v3 := []string{"a"}
    var v4 uint16 = 712
    var v5 int32 = 1
    v6 := make(map[uint32]string)
    v4 = v1
    v5 = 415
    v1 += v1
    var v9 bool = false
    var v10 int8 = 0
    if ((v9 || (!(-1 < -995))) && true) {
        fmt.Println(v9)
        v11 := GenS1{F0: 1, F1: int(v1)}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 4; fi12++ {
            for fi13 := 0; fi13 < 6; fi13++ {
                v14 := GenS1{F0: v4, F1: (fi12 + 32767)}
                v5 += v5
                _ = fi13
                _ = v14
            }
            if false {
                fmt.Println(v1, v1)
                var v15 uint8 = 69
                v9 = v9
                _ = v15
            }
            if v9 {
                v3 = append(v3, "\n\t")
            }
            _ = fi12
        }
        v16 := GenS2{F0: 954, F1: (v10 | (v10 / (((-(v10))) | 1))), F2: 32767}
        if v9 {
            var v17 int8 = (v10 / (((v10 >> ((genFn0((v10 * v10))) & 63))) | 1))
            _ = v17
        }
        _ = v16
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        fmt.Println(v5, v9, v4)
        if v9 {
            delete(v6, uint32(v1))
            _ = len(v2)
        } else {
            var v19 uint8 = 128
            _ = v19
        }
        _ = fi18
    }
    v9 = (-998 <= 127)
    v1 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
