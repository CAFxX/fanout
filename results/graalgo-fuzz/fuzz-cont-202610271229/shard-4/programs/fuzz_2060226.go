package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 uint32; F2 int }
type GenS4 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint) int8 {
    v7 := []int{413, 32768, -1, 42}
    v8 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 int32, p1 int8, p2 uint8) int32 {
    v9 := []int{-2147483648, 1}
    v10 := GenS2{F0: ("a" + "foo bar")}
    v11 := GenS1{F0: 444}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return p0
}

func main() {
    var v1 float32 = (float32(-2.25) / -2.25)
    v2 := make(map[int]uint32)
    var v3 int8 = 127
    v4 := make(map[uint32]uint32)
    var v5 int64 = 9223372036854775807
    v6 := make(map[int]bool)
    var v12 string = ((("a" + (("" + "hello") + "\n\t")) + "foo bar") + "foo bar")
    var v13 int32 = genFn1(-963, (v3 / (((v3 | (v3 << ((642) & 63)))) | 1)), 127)
    delete(v4, 4294967295)
    var v14 uint64 = 42
    v15 := GenS1{F0: 861}
    switch ("\n\t" + v12) {
    case "hello":
        var v16 int16 = -32768
        _ = v16
    default:
        var v17 uint16 = 226
        _ = v17
    }
    v3 = v3
    v18 := GenS2{F0: (v12 + ("" + ("hello" + "")))}
    fmt.Println(v5, v5, v14)
    if true {
        var v19 uint8 = 110
        v4[42] = 2147483647
        v20 := GenS3{F0: 127, F1: 4294967295, F2: 127}
        _ = v19
        _ = v20
    } else {
        v21 := GenS4{F0: 2147483648, F1: (-(v3))}
        fmt.Println(v5, v12, v12)
        v6[32768] = true
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
