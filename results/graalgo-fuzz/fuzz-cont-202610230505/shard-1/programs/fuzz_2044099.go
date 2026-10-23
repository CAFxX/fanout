package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint32 }
func genFn0(p0 int32, p1 string, p2 int16) uint16 {
    fmt.Println(p0)
    v7 := []int{int(p2), -9223372036854775808}
    switch "hello" {
    case "":
        v8 := GenS1{F0: 2147483647, F1: false, F2: 1}
        _ = v8
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}
func genFn1(p0 uint8, p1 uint8, p2 uint8) int8 {
    var v9 bool = false
    v10 := GenS1{F0: 175, F1: (2147483648 != 2147483648), F2: 118}
    v11 := GenS2{F0: 2, F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 100
}

func main() {
    var v1 int16 = -32768
    var v2 uint16 = 32767
    v3 := make(map[int]uint32)
    var v4 int64 = 42
    var v5 int = 9223372036854775807
    var v6 int32 = 1
    v12 := GenS3{F0: 466, F1: 100, F2: 42}
    if true {
        v13 := []int{v5, v5}
        if ((v1 / (((v1 + (v1 * (^(v1))))) | 1)) == (v1 + (-(v1)))) {
            v13 = append(v13, 4294967295)
            fmt.Println(v5)
        }
        if false {
            fmt.Println(v6, v4, v2)
        }
        _ = v13
    }
    var v14 string = ("foo bar" + "")
    v6 -= 1
    v4 += -1
    var v15 uint16 = v2
    var v16 uint8 = 2
    v17 := GenS2{F0: v16, F1: (^(v16))}
    var v18 uint16 = genFn0((v6 ^ (-(v6))), "foo bar", (v1 ^ (v1 | (v1 % ((-472) | 1)))))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
