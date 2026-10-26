package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint) uint64 {
    v7 := GenS1{F0: 914, F1: 42}
    var v8 uint = p0
    p0 = 255
    if (!true) {
        var v9 int32 = 221
        var v10 uint64 = 18446744073709551615
        v11 := []int{4294967295}
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 uint32 = 65535
    _ = p0
    _ = v12
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 bool, p1 uint8, p2 uint32) int16 {
    v13 := GenS1{F0: 255, F1: (p2 >> ((0) & 63))}
    v14 := []int{4294967295, 257, 128, 1000}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return -522
}

func main() {
    var v1 uint16 = 32767
    v2 := []int{128, -1, -2147483648}
    var v3 uint16 = 42
    var v4 int = 42
    var v5 float32 = 1.5
    v6 := make(map[string]int)
    fmt.Println(v1, v3)
    var v15 int = 127
    var v16 uint8 = 0
    v17 := GenS1{F0: v3, F1: 65535}
    var v18 int8 = -45
    switch (("hello" + ("\n\t" + "a")) + "a") {
    case "":
        v18 = 0
    case "a":
        if false {
            if (1 >= 0) {
                var v19 int32 = 2147483647
                if true {
                    v6[("\n\t" + "foo bar")] = v15
                    v15 = (v15 ^ (v4 | (v15 & (^(v15)))))
                }
                v20 := GenS1{F0: v1, F1: 991}
                _ = v19
                _ = v20
            }
        }
    }
    var v21 uint = 2
    v1 = (v1 << ((v21) & 63))
    fmt.Println(v18)
    v2 = append(v2, v4)
    var v22 int = v4
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
