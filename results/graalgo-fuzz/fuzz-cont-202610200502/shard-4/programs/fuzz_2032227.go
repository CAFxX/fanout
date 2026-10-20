package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]bool)
    var v3 int = 276
    v4 := GenS1{F0: "\n\t"}
    v1 = (v1 >> ((18446744073709551615) & 63))
    v1 -= v1
    v1 -= int32(v3)
    if false {
        for fi5 := 0; fi5 < 7; fi5++ {
            if true {
                if true {
                    v6 := GenS1{F0: "\n\t"}
                    var v7 int32 = (v1 / ((v1) | 1))
                    var v8 bool = (((v7 - (-(v7))) >= v7) && false)
                    _ = v6
                    _ = v7
                    _ = v8
                } else {
                    if (true || ((true && true) && (false && false))) {
                        v9 := []int{42, fi5, 0}
                        _ = v9
                    }
                }
                v10 := GenS2{F0: 694}
                _ = v10
            } else {
                var v11 uint32 = 628
                v11 = v11
                v3 = (fi5 | fi5)
                _ = v11
            }
            v12 := GenS1{F0: "\n\t"}
            _ = fi5
            _ = v12
        }
        delete(v2, (v3 >> ((1) & 63)))
    }
    v3 = (v3 ^ (v3 * 2147483647))
    v13 := GenS3{F0: 9223372036854775807, F1: 9223372036854775807}
    var v14 float64 = float64(1e10)
    v15 := GenS4{F0: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
