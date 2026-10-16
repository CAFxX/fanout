package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int8 }

func main() {
    var v1 string = ("a" + ("hello" + "a"))
    var v2 uint16 = 480
    var v3 uint8 = 0
    fmt.Println(v1, v2, v2)
    fmt.Println(v3)
    v1 = "a"
    if true {
        if true {
            v2 = (v2 / ((v2) | 1))
        }
        for fi4 := 0; fi4 < 4; fi4++ {
            var v5 float32 = ((1.5 + ((float32(0.1) * 1.5) + ((1.5 / (0.1 - ((1e10 * 1.5) * 0.0))) + float32((3.14159 * ((3.14159 - 0.1) * 0.1)))))) / 1.0)
            _ = fi4
            _ = v5
        }
    }
    var v6 string = v1
    for fi7 := 0; fi7 < 3; fi7++ {
        if (!true) {
            var v8 uint64 = 0
            if (true && true) {
                v6 = (v1 + v1)
                v3 = 127
                v1 = v1
            }
            v8 = (^(v8))
            _ = v8
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        if (-718 == 42) {
            fmt.Println(fi9, v2)
        }
        _ = fi9
    }
    v10 := GenS1{F0: 42}
    if false {
        v11 := []int{32768, 128, -626, 128}
        v3 -= (v3 * v3)
        _ = v11
    } else {
        v12 := GenS2{F0: 755, F1: 2, F2: 100}
        _ = v12
    }
    v2 += (v2 - v2)
    var v13 int8 = 1
    v14 := []int{}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    var v15 uint16 = (v2 ^ (v2 - v2))
    if (-9223372036854775808 != 42) {
        v15 += (v15 / ((v2) | 1))
        var v16 uint64 = 18446744073709551615
        _ = v16
    }
    v3 -= 42
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
}
