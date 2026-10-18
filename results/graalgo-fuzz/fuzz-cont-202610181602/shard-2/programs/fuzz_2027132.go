package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int16 }
func genFn0(p0 uint) int16 {
    var v5 int32 = 42
    v6 := GenS1{F0: 0, F1: (true && false), F2: -32768}
    v7 := []int{127, 32767}
    p0 = (p0 + p0)
    if (1000 > 942) {
        p0 -= p0
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint16, p1 string, p2 string) string {
    v8 := []int{-2147483648, -1, 0, -1}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v9 int32 = 0
    v9 = 198
    fmt.Println(p2, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 65535
    var v3 float32 = (0.0 + (3.14159 * (0.1 + 3.14159)))
    v4 := []string{("foo bar" + "foo bar"), "foo bar", ("" + "hello"), ("\n\t" + ("foo bar" + "\n\t"))}
    v3 = float32(1e10)
    var v10 uint32 = 2
    var v11 bool = (false || (531 >= 1))
    delete(v1, 42)
    var v12 float64 = 3.14159
    var v13 uint32 = (v10 >> ((42) & 63))
    var v14 float64 = 1.5
    var v15 float64 = (v14 / ((v12 + v14) - (0.1 * (-2.25 - 0.0))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
