package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int16 }
func genFn0() int {
    for fi6 := 0; fi6 < 8; fi6++ {
        if (!true) {
            fmt.Println(fi6)
            var v7 int64 = 9223372036854775807
            var v8 string = ""
            _ = v7
            _ = v8
        } else {
            switch ("\n\t" + "\n\t") {
            case "a":
                v9 := []int{-144, fi6}
                _ = v9
            default:
                if ((1.5 * 0.1) != (0.1 - (0.1 - 3.14159))) {
                    var v10 uint64 = 4294967295
                    _ = v10
                }
            }
            var v11 uint32 = 957
            _ = v11
        }
        for fi12 := 0; fi12 < 9; fi12++ {
            fmt.Println(fi12, fi12)
            var v13 uint = 138
            if false {
                if false {
                    var v14 uint16 = 32767
                    v13 += uint(v14)
                    _ = v14
                }
                v15 := []int{(fi12 & fi6), fi6, (fi12 / ((fi12) | 1))}
                _ = v15
            }
            _ = fi12
            _ = v13
        }
        _ = fi6
    }
    var v16 uint32 = 4294967295
    _ = v16
    return 2147483647
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 4294967295
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 534
    v17 := GenS1{F0: uint16(v2), F1: 1, F2: -1}
    v18 := []int{genFn0(), 128, 880, genFn0()}
    v3[("foo bar" + "a")] = "\n\t"
    var v19 int16 = 32767
    var v20 uint = 938
    fmt.Println(v19)
    v19 += (v19 & 0)
    _ = len(v1)
    var v21 int16 = -942
    v4[65535] = 0
    v5 = (v5 | (v5 * v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
