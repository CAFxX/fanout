package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint }
type GenS2 struct { F0 uint; F1 uint8; F2 string }
func genFn0() int8 {
    v7 := []int{-778, 4294967295, 2147483647, 4294967295}
    var v8 uint8 = 1
    v8 += (v8 % (((v8 | v8)) | 1))
    fmt.Println(v8)
    _ = v7
    _ = v8
    return 98
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{65535, 0, 65535}
    var v3 float32 = float32(0.1)
    var v4 int8 = 100
    var v5 bool = false
    var v6 int = 526
    var v9 int64 = -1
    v4 = v4
    fmt.Println(v5)
    v9 -= (-(v9))
    var v10 int64 = v9
    var v11 bool = ((genFn0() <= (v4 & v4)) && (!((v4 + v4) == (v4 ^ v4))))
    switch "hello" {
    case "a":
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 2; fi13++ {
                v5 = (0 > 2)
                _ = fi13
            }
            var v14 float64 = (-2.25 - 0.0)
            v15 := GenS1{F0: genFn0(), F1: (v10 % ((v9) | 1)), F2: 650}
            _ = fi12
            _ = v14
            _ = v15
        }
    case "":
        if v5 {
            if v11 {
                v16 := GenS1{F0: v4, F1: v10, F2: 324}
                _ = v16
            }
        } else {
            var v17 uint8 = 127
            _ = v17
        }
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v18 := GenS2{F0: 65535, F1: 174, F2: (("\n\t" + ("foo bar" + "a")) + ("a" + "a"))}
    if len(v2) > 0 {
        v2[0] = 2
    }
    v19 := GenS2{F0: 732, F1: 1, F2: "hello"}
    var v20 uint32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
