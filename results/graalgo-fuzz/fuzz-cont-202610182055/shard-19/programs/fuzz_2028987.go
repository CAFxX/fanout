package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint32, p1 int, p2 int8) int16 {
    if ("a" != "hello") {
        p1 = 4294967295
        var v7 uint = 42
        _ = v7
    }
    p1 = 32768
    fmt.Println(p2)
    var v8 int32 = 0
    var v9 int8 = -3
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -584
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (true && false)
    var v3 int8 = -6
    var v4 int32 = 2147483647
    v5 := make(map[uint32]bool)
    var v6 string = "\n\t"
    v6 = "foo bar"
    delete(v5, 0)
    if v2 {
        var v10 int = 239
        if v2 {
            v11 := GenS1{F0: 1, F1: 127, F2: 42}
            var v12 uint16 = 255
            _ = v11
            _ = v12
        } else {
            v13 := v1[v10]
            _ = v13
        }
        _ = v10
    }
    fmt.Println(v3)
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 int8 = (^(v3))
        var v16 float32 = -2.25
        _ = fi14
        _ = v15
        _ = v16
    }
    v4 += (v4 + (v4 ^ v4))
    v4 += v4
    v2 = v2
    fmt.Println(v2)
    v17 := GenS1{F0: genFn0(944, 42, v3), F1: 1, F2: 0}
    if (!(2 < 65)) {
        v18 := v1[32768]
        _ = v18
    } else {
        if false {
            if v2 {
                var v19 float32 = ((((-2.25 / 1.0) - 1.5) - 0.1) / (0.0 + float32(0.1)))
                v20 := GenS1{F0: genFn0(4294967295, 14, 111), F1: v3, F2: 2}
                _ = v19
                _ = v20
            }
            switch v6 {
            case "a":
                switch v6 {
                case "":
                    var v21 int32 = v4
                    _ = v21
                default:
                    var v22 uint = 42
                    _ = v22
                }
            case "hello":
                v23 := []int{}
                _ = v23
            }
        } else {
            var v24 float32 = ((0.1 / 0.1) * (3.14159 + 1.5))
            v2 = (!v2)
            _ = v24
        }
        var v25 uint64 = 18446744073709551615
        _ = v25
    }
    fmt.Println(v6, v6)
    v26 := GenS2{F0: (v3 + (v3 % (((v3 + (-(v3)))) | 1))), F1: 65535}
    var v27 uint16 = 42
    switch (v6 + v6) {
    case "hello":
        _ = len(v5)
    case "foo bar":
        v2 = (v3 != (^(v3)))
    default:
        var v28 int = -2147483648
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
