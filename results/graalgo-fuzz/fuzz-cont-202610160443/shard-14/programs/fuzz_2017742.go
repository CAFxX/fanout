package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint16; F2 int16 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (-2.25 + -2.25)
    v3 := make(map[uint32]int)
    v4 := GenS1{F0: 2147483648}
    var v5 uint8 = 128
    v5 = v5
    if true {
        if false {
            v5 += (v5 / ((255) | 1))
            v2 = 0.1
        } else {
            if false {
                v5 -= (v5 * 87)
                if true {
                    var v6 string = "\n\t"
                    var v7 string = v6
                    _ = v6
                    _ = v7
                }
            } else {
                v8 := []int{}
                v9 := v3[0]
                v8 = append(v8, v9)
                _ = v8
                _ = v9
            }
        }
        v10 := GenS2{F0: -1, F1: 0, F2: 186}
        v5 += (v5 | 1)
        _ = v10
    }
    v1[2147483647] = false
    fmt.Println(v5, v5)
    fmt.Println(v2)
    if (false && ((!(32768 != 127)) || true)) {
        v11 := GenS1{F0: 42}
        _ = v11
    } else {
        v12 := GenS3{F0: -1}
        _ = v12
    }
    var v13 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
