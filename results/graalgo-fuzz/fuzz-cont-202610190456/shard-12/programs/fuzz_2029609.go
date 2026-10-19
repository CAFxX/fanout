package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint8 }
func genFn0() string {
    var v6 uint16 = 32767
    v7 := []int{-27, -106, 979}
    var v8 int32 = -354
    v9 := GenS1{F0: ((1 < 65535) && true), F1: 65535, F2: 995}
    var v10 bool = ((602 > 2) || (1000 > 32767))
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int8 = 125
    var v3 bool = true
    v4 := []uint32{42, 707, 65535, 841}
    var v5 uint = 65535
    var v11 float64 = 0.0
    var v12 uint64 = v1
    var v13 uint32 = 0
    v13 += v13
    v14 := GenS2{F0: (^(v5)), F1: v13, F2: -1}
    v15 := GenS3{F0: v13, F1: 927}
    v16 := GenS4{F0: uint8(v13)}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v17 uint32 = 1
    v12 += (v1 ^ (v12 << (((^(v5))) & 63)))
    fmt.Println(v17, v13)
    if true {
        if v3 {
            var v18 uint8 = 0
            for fi19 := 0; fi19 < 3; fi19++ {
                switch fi19 {
                case 9223372036854775807:
                    v18 += (v18 | v18)
                }
                _ = fi19
            }
            _ = v18
        } else {
            v11 += v11
            v13 += (^(v13))
            var v20 uint64 = (v1 >> ((723) & 63))
            _ = v20
        }
        v21 := GenS2{F0: v5, F1: v17, F2: 2}
        var v22 bool = true
        _ = v21
        _ = v22
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v23 := GenS2{F0: (v5 / ((4294967295) | 1)), F1: v13, F2: -1}
    var v24 uint32 = v17
    var v25 float32 = (((1e10 + 0.1) - (0.0 * 0.1)) + (3.14159 + (0.0 / 1.0)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
