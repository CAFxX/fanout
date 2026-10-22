package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int64 }
func genFn0(p0 uint64, p1 int32, p2 string) int64 {
    var v5 uint16 = 32767
    var v6 int8 = 100
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 string = ("hello" + "foo bar")
    var v2 uint8 = 255
    var v3 uint = 0
    v4 := make(map[int]int)
    var v8 uint16 = 255
    var v9 string = "foo bar"
    if true {
        v3 += uint(v8)
    }
    v2 = v2
    v1 = (v9 + v1)
    if true {
        if true {
            fmt.Println(v1, v8)
            v1 = "a"
        }
    }
    v10 := GenS1{F0: (v8 >> (((v3 / (((^(v3))) | 1))) & 63)), F1: v2, F2: int64(v2)}
    var v11 float64 = (1.5 / (float64(float32(float64(1.5))) / 1.5))
    fmt.Println(v9, v11)
    fmt.Println(v1, v1)
    v9 = (((((("" + ((v9 + "a") + "hello")) + v1) + ("hello" + v9)) + (v1 + (((v1 + v9) + "\n\t") + (v1 + ((v1 + (v1 + (v9 + ("" + "a")))) + v9))))) + ((v9 + "foo bar") + v9)) + v1)
    v9 = v1
    v11 = 3.14159
    fmt.Println(v11, v1, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
