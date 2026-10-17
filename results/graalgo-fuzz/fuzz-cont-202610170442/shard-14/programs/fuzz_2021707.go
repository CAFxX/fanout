package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int64; F1 string }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 uint = 2
    var v4 float32 = float32(float64(-2.25))
    var v5 int8 = 100
    var v6 uint32 = 4294967295
    v3 = (v3 << ((v3) & 63))
    if (("" + ((("hello" + "\n\t") + "hello") + "a")) <= "a") {
        if true {
            if (!true) {
                v7 := v1[(^(v6))]
                v6 += v6
                _ = v7
            }
        }
    }
    v3 -= v3
    v8 := GenS1{F0: false, F1: 0}
    if (v5 < (v5 | v5)) {
        v4 = -2.25
    } else {
        if true {
            v9 := GenS2{F0: 0, F1: ""}
            _ = v9
        }
    }
    v10 := GenS1{F0: (!(32767 >= -592)), F1: -2147483648}
    v11 := v1[(^(v6))]
    v5 = v5
    v4 = v4
    var v12 float32 = v4
    v13 := []int{-956, 32768, -880, -620}
    switch 1 {
    case 2147483647:
        fmt.Println(v6)
    }
    v14 := v2[128]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
