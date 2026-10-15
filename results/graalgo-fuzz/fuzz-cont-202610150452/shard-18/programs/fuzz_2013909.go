package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
func genFn0(p0 uint16) uint8 {
    fmt.Println(p0)
    var v4 uint64 = 2
    fmt.Println(p0, v4)
    var v5 float32 = 0.1
    _ = p0
    _ = v4
    _ = v5
    return 16
}
func genFn1(p0 string, p1 uint32, p2 int) uint16 {
    v6 := []int{(p2 >> ((314) & 63)), 128, p2}
    p2 = (p2 ^ (p2 ^ p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 255
}

func main() {
    var v1 string = "hello"
    var v2 float64 = 0.1
    v3 := make(map[int]string)
    v7 := GenS1{F0: v1, F1: 42}
    for fi8 := 0; fi8 < 20; fi8++ {
        v1 = "\n\t"
        _ = fi8
    }
    switch v1 {
    case "a":
        fmt.Println(v2, v1)
    case "hello":
        if ((!(!true)) && ((!(false && true)) || (-2.25 != v2))) {
            fmt.Println(v2)
            v1 = v1
        } else {
            v9 := GenS1{F0: (v1 + ""), F1: 824}
            v10 := []int{9223372036854775807, 128}
            _ = v10
            _ = v9
        }
    }
    v11 := []int{128}
    var v12 uint8 = 0
    v1 = v1
    if (genFn1(v1, 2147483647, 32767) < 831) {
        for fi13 := 0; fi13 < 10; fi13++ {
            _ = len(v11)
            _ = fi13
        }
        if false {
            v1 = v1
        }
        v14 := GenS1{F0: v1, F1: 42}
        _ = v14
    } else {
        if true {
            v2 -= ((v2 * (v2 * 3.14159)) - v2)
            _ = len(v3)
        } else {
            if (1 < 2) {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                var v15 int64 = 42
                _ = v15
            }
        }
        v12 -= (v12 * v12)
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
}
