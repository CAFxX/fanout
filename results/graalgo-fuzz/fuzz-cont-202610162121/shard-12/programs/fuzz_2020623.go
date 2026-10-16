package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int32 }
func genFn0(p0 uint8, p1 int, p2 uint32) int64 {
    p1 = 1
    p0 = p0
    var v4 int = 2147483647
    switch ("foo bar" + "\n\t") {
    case "a":
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := []int{(v4 % (((v4 & p1)) | 1)), p1}
            _ = fi5
            _ = v6
        }
    case "\n\t":
        v7 := GenS1{F0: -1, F1: 4294967295, F2: -1}
        _ = v7
    case "hello":
        switch ("a" + "\n\t") {
        case "\n\t":
            if false {
                v8 := GenS1{F0: p1, F1: 2147483647, F2: 0}
                _ = v8
            } else {
                p1 += v4
            }
        case "hello":
            var v9 int = -745
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 1e10
    v3 := make(map[int]int)
    var v10 float64 = 0.0
    var v11 uint8 = 128
    if (65535 > 2) {
        v12 := GenS1{F0: 42, F1: uint32(v11), F2: 1}
        v2 -= -2.25
        _ = v12
    }
    var v13 bool = ((v2 != (float32(v10) + (v2 - (0.1 + v2)))) && false)
    if v13 {
        fmt.Println(v11, v2, v13)
        delete(v1, -1)
        _ = len(v1)
    }
    v2 += v2
    var v14 uint16 = 1
    var v15 int = 128
    v11 -= 127
    _ = len(v1)
    v2 -= float32(v10)
    v16 := GenS1{F0: (^(v15)), F1: 928, F2: -1}
    _ = len(v3)
    switch ("a" + ("a" + ("hello" + (("hello" + "foo bar") + "hello")))) {
    case "\n\t":
        if v13 {
            v17 := GenS1{F0: v15, F1: 155, F2: -285}
            _ = v17
        }
    }
    v2 += (v2 / -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
