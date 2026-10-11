package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint16, p2 string) int8 {
    v5 := []int{int(p1), int(p1), 2147483647, 2147483647}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}
func genFn1(p0 uint16, p1 int8) uint64 {
    switch "foo bar" {
    case "hello":
        fmt.Println(p0)
    case "":
        fmt.Println(p0)
    case "\n\t":
        for fi6 := 0; fi6 < 3; fi6++ {
            p1 = p1
            switch fi6 {
            case 32767:
                p1 = (p1 >> ((1) & 63))
            }
            p1 = (p1 * (p1 ^ genFn0(655, (p0 << ((295) & 63)), (("" + (("foo bar" + "\n\t") + "hello")) + (("a" + (("" + "hello") + ("foo bar" + ""))) + (("\n\t" + "") + "hello"))))))
            _ = fi6
        }
    }
    p0 -= (^(p0))
    v7 := []int{-688, 397, 32768}
    var v8 uint64 = 4294967295
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]uint32)
    v3 := []int{-971}
    var v4 bool = (4294967295 < 2)
    v4 = false
    for fi9 := 0; fi9 < 19; fi9++ {
        v4 = (((888 >= 32767) && ((((3.14159 * 3.14159) / 3.14159) - (3.14159 + ((float32(-2.25) * float32((float64(float32(1e10)) / 0.1))) / (1.5 + 1.5)))) == 1e10)) || (v4 || v4))
        _ = fi9
    }
    fmt.Println(v4)
    _ = len(v1)
    _ = len(v2)
    v4 = true
    v3 = append(v3, 127)
    switch -175 {
    case 882:
        v10 := GenS1{F0: 1000}
        _ = v10
    case -976:
        v3 = append(v3, -821)
    case 1:
        var v11 int8 = -128
        _ = v11
    }
    v4 = ("a" == ("foo bar" + (("" + "a") + "")))
    v4 = (!(2 >= 1))
    var v12 uint16 = 255
    v12 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
