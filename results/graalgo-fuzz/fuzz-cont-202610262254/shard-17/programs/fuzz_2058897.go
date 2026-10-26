package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
func genFn0(p0 uint64, p1 uint8, p2 uint16) uint16 {
    if true {
        var v7 uint64 = p0
        switch ("hello" + "\n\t") {
        case "hello":
            v8 := []int{}
            _ = v8
        case "a":
            v7 = (v7 | 0)
        case "foo bar":
            p2 = (p2 ^ (p2 << ((uint(p0)) & 63)))
        default:
            v9 := GenS1{F0: 1}
            _ = v9
        }
        _ = v7
    }
    fmt.Println(p0)
    var v10 uint = 849
    v11 := GenS1{F0: 720}
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS2{F0: p0}
        v14 := []int{fi12, fi12, fi12, fi12}
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return (p2 & (p2 << (((v10 >> ((0) & 63))) & 63)))
}

func main() {
    var v1 uint16 = 0
    var v2 float32 = 1e10
    var v3 uint32 = 2
    v4 := make(map[int]string)
    v5 := []string{(("" + ("" + (("hello" + "foo bar") + ("foo bar" + (((("" + "\n\t") + "a") + "") + "hello"))))) + "a"), ("" + ("hello" + "hello"))}
    v6 := make(map[int]int)
    fmt.Println(v2, v3)
    var v15 int64 = 0
    v16 := GenS3{F0: 32768}
    v2 -= 1e10
    v17 := GenS4{F0: (2 > 42)}
    if false {
        _ = len(v4)
        v18 := GenS2{F0: 42}
        _ = v18
    } else {
        v15 -= (v15 ^ 42)
        var v19 int64 = (v15 % ((int64(v3)) | 1))
        _ = v19
    }
    v15 = (-(v15))
    var v20 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
