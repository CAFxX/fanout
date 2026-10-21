package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int }
type GenS3 struct { F0 bool; F1 uint8 }
func genFn0(p0 int, p1 string, p2 uint64) int8 {
    var v7 int32 = 42
    var v8 int32 = (v7 % ((-831) | 1))
    v9 := []int{}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 100
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 821
    var v3 string = (("hello" + "hello") + "hello")
    var v4 uint = 42
    v5 := make(map[int]string)
    var v6 float32 = 0.0
    v10 := []int{-1, 2147483647}
    var v11 int16 = (v2 >> ((v4) & 63))
    v6 += (-2.25 / (float32(0.0) / (float32((-2.25 / 1.0)) / float32((float64((v6 - -2.25)) * 3.14159)))))
    for fi12 := 0; fi12 < 16; fi12++ {
        if (-2147483648 < 333) {
            v13 := GenS1{F0: 2147483647}
            v14 := GenS1{F0: 1}
            _ = v13
            _ = v14
        }
        if (!false) {
            v15 := GenS2{F0: 0, F1: 2147483647, F2: 32767}
            _ = v15
        }
        _ = fi12
    }
    v16 := GenS3{F0: (42 != 9223372036854775807), F1: 127}
    v2 -= 1000
    fmt.Println(v6, v2, v4)
    fmt.Println(v2, v6, v3)
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 int8 = genFn0(4294967295, v3, 18446744073709551615)
        _ = fi17
        _ = v18
    }
    var v19 int = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
