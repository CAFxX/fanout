package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int }
func genFn0(p0 uint64) int16 {
    var v6 uint8 = 1
    v7 := GenS1{F0: 255}
    var v8 int8 = -26
    v9 := []int{-2147483648, 0, 32768}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -32768
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 100
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    var v5 bool = (1 == -1)
    var v10 bool = v5
    v11 := GenS2{F0: 0, F1: 42}
    fmt.Println(v5, v2)
    fmt.Println(v5, v5)
    var v12 uint = 65535
    if v5 {
        fmt.Println(v5, v12)
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 int = (^(fi13))
            var v15 int16 = -1
            v12 = v12
            _ = fi13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            for fi17 := 0; fi17 < 13; fi17++ {
                v18 := GenS2{F0: v12, F1: 128}
                v5 = (!v5)
                _ = fi17
                _ = v18
            }
            v10 = (false && v5)
            v19 := []int{fi16}
            _ = fi16
            _ = v19
        }
    }
    fmt.Println(v10, v5, v12)
    delete(v4, 966)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
