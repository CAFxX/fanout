package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 uint32, p1 uint16) int32 {
    v7 := []int{}
    var v8 bool = (-934 <= 29)
    v9 := GenS1{F0: p1, F1: 242}
    var v10 uint8 = 255
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1() int8 {
    v11 := []int{2147483647}
    _ = len(v11)
    var v12 bool = (!true)
    _ = len(v11)
    var v13 string = ("\n\t" + "hello")
    _ = v11
    _ = v12
    _ = v13
    return 100
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    v4 := []string{"", "foo bar", ""}
    var v5 string = "foo bar"
    v6 := make(map[int]uint32)
    if false {
        fmt.Println(v1)
    }
    fmt.Println(v5)
    v1 = (v1 ^ v1)
    v5 = ((v5 + ("a" + "")) + "\n\t")
    if (255 != 76) {
        if len(v4) > 0 {
            v4[0] = ""
        }
        v14 := GenS2{F0: 32767, F1: 42}
        _ = v14
    } else {
        _ = len(v6)
        v1 -= 1
    }
    v1 += v1
    v1 = genFn0(4294967295, 1)
    v5 = "hello"
    v15 := GenS1{F0: 227, F1: 2147483648}
    var v16 int16 = 620
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
