package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int32 }
func genFn0(p0 uint, p1 bool, p2 int8) int16 {
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint8 = 255
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}

func main() {
    v1 := []int{32768}
    var v2 bool = false
    var v3 bool = true
    v4 := []string{("foo bar" + ("foo bar" + "hello")), "a"}
    v8 := GenS1{F0: v3, F1: "hello", F2: 1}
    var v9 int8 = -76
    v9 -= (v9 / ((v9) | 1))
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    if (33 >= 0) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            v9 = 9
            v9 = v9
            v3 = (((0.1 * 1e10) != (0.1 / 1.5)) || (-9223372036854775808 > 42))
            _ = fi10
        }
        fmt.Println(v2, v3, v9)
    }
    fmt.Println(v2, v2)
    var v11 float32 = (3.14159 / 1e10)
    v9 = 1
    fmt.Println(v2, v3)
    switch "hello" {
    case "":
        fmt.Println(v9)
    }
    fmt.Println(v2, v11)
    v3 = false
    v12 := GenS1{F0: (0 != 42), F1: "a", F2: 2147483647}
    v2 = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
