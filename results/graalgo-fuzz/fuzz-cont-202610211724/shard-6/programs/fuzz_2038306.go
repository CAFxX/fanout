package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
type GenS2 struct { F0 uint16; F1 string }

func main() {
    var v1 float64 = 1e10
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]uint32)
    if true {
        v1 += v1
        v5 := GenS1{F0: 226, F1: 42, F2: false}
        _ = v5
    } else {
        v6 := []int{262, -9223372036854775808, 32767, -9223372036854775808}
        _ = v6
    }
    fmt.Println(v3, v3)
    v7 := []int{-828, 32767, -239}
    var v8 float32 = 1e10
    v3 += v3
    var v9 bool = true
    v10 := GenS1{F0: 255, F1: -24, F2: (false || ((v3 ^ v3) >= v3))}
    fmt.Println(v8, v8, v1)
    v9 = v9
    v8 -= float32(v1)
    if v9 {
        v11 := v2[v3]
        for fi12 := 0; fi12 < 8; fi12++ {
            for fi13 := 0; fi13 < 16; fi13++ {
                fmt.Println(fi13)
                v14 := GenS1{F0: 4294967295, F1: 42, F2: false}
                var v15 int16 = 915
                _ = fi13
                _ = v14
                _ = v15
            }
            if (!v9) {
                fmt.Println(v1, v1, v11)
                for fi16 := 0; fi16 < 13; fi16++ {
                    v3 += v3
                    v9 = false
                    v17 := GenS2{F0: 894, F1: (("hello" + "foo bar") + "\n\t")}
                    var v18 float64 = (v1 + 0.1)
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            } else {
                fmt.Println(v1, v8)
                _ = len(v4)
            }
            _ = fi12
        }
        fmt.Println(v8, v1, v11)
        _ = v11
    }
    v19 := GenS2{F0: 0, F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
