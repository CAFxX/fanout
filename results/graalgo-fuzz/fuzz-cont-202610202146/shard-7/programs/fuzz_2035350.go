package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint64) int16 {
    var v7 int16 = -684
    fmt.Println(p0, p0)
    p0 -= p0
    _ = p0
    _ = v7
    return (v7 - v7)
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]bool)
    var v3 int8 = 127
    v4 := make(map[uint32]string)
    v5 := []int{32767, 241, -594, 1}
    v6 := make(map[int]uint32)
    switch -2147483648 {
    case -249:
        if v1 {
            if true {
                var v8 float32 = (1.5 / (-2.25 + (1.5 - (3.14159 + (0.0 / 1e10)))))
                _ = v8
            }
        } else {
            var v9 int8 = -37
            _ = v9
        }
    case 0:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v10 uint = 2
    var v11 bool = false
    _ = len(v5)
    v11 = v11
    var v12 int8 = (v3 << ((v10) & 63))
    v13 := GenS1{F0: 255}
    var v14 int64 = 0
    _ = len(v4)
    v6[-2147483648] = 0
    if len(v5) > 0 {
        v5[0] = 1
    }
    fmt.Println(v1)
    v15 := GenS1{F0: 255}
    if v1 {
        v11 = (v14 <= -167)
        v5 = append(v5, -2147483648)
        if (!v11) {
            var v16 int16 = 42
            v11 = ((0.0 - 0.0) < ((0.0 / -2.25) + 3.14159))
            _ = v16
        } else {
            var v17 float64 = ((0.0 - 1e10) + 0.0)
            _ = v17
        }
    }
    var v18 string = ""
    v1 = (1e10 != 1.5)
    v19 := GenS1{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
