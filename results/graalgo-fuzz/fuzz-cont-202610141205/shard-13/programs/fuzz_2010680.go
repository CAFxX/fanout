package main

import (
    "fmt"
)

func genFn0(p0 int8, p1 uint32) uint16 {
    p1 += 2147483647
    var v7 float32 = (((float32(0.1) / (1.5 + 1e10)) * 3.14159) - (((3.14159 + 3.14159) * (1.5 - (1.5 / (((-2.25 - (1.5 * -2.25)) - (1e10 + 1.5)) / float32(1e10))))) + float32(1e10)))
    var v8 uint8 = 98
    v7 -= 1.5
    var v9 uint64 = 184
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 831
}
func genFn1(p0 uint, p1 int16) uint64 {
    var v10 int = int(p1)
    v10 -= v10
    v11 := []int{v10, (v10 >> ((255) & 63)), -2147483648}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    var v3 float32 = 1.5
    var v4 string = ""
    v5 := []uint32{0, 2147483648, 2147483648, 920}
    v6 := make(map[string]uint32)
    v4 = (v4 + ("hello" + ("a" + "hello")))
    _ = len(v6)
    fmt.Println(v4)
    if true {
        v3 += (0.1 * (v3 - ((v3 + ((v3 / ((-2.25 + float32((1.5 * ((float64(v3) + 0.0) + 1e10)))) * v3)) - 0.1)) + (1.5 * v3))))
    } else {
        fmt.Println(v4)
    }
    var v12 uint16 = 551
    v13 := v2[("foo bar" + "")]
    var v14 int16 = 42
    var v15 uint32 = 994
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
