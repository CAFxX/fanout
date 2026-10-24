package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 string; F2 uint }
func genFn0() int8 {
    v5 := []int{}
    v6 := GenS1{F0: -1}
    var v7 float64 = (3.14159 * (3.14159 + (((-2.25 - 1e10) / 1.5) * (0.1 / 3.14159))))
    v8 := GenS2{F0: true, F1: "a", F2: 418}
    v7 -= v7
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 395
    v3 := make(map[int]int)
    var v4 int8 = 1
    v4 -= (v4 | (v4 / (((v4 << ((255) & 63))) | 1)))
    if (255 == 2) {
        if ((v2 ^ (v2 | 42)) > (v2 >> ((0) & 63))) {
            switch 491 {
            case 1:
                for fi9 := 0; fi9 < 10; fi9++ {
                    v4 = -27
                    var v10 bool = (false || (true && (1 <= 95)))
                    _ = fi9
                    _ = v10
                }
            case 2147483647:
                v11 := []int{int(v2)}
                _ = v11
            default:
                fmt.Println(v4)
            }
            if false {
                fmt.Println(v4)
                if true {
                    v4 = (v4 & v4)
                } else {
                    var v12 uint = 42
                    _ = v12
                }
            }
            var v13 int = 127
            _ = v13
        } else {
            v1[v2] = 65535
        }
    }
    switch (("foo bar" + (("a" + "foo bar") + ("\n\t" + "hello"))) + "hello") {
    case "":
        var v14 float64 = (3.14159 / ((-2.25 / 1.0) * float64(0.1)))
        _ = v14
    case "foo bar":
        var v15 int32 = 42
        _ = v15
    }
    var v16 int32 = 126
    v2 -= v2
    if false {
        v2 = 2147483647
        var v17 int32 = -1
        _ = v17
    } else {
        v18 := GenS1{F0: genFn0()}
        _ = v18
    }
    var v19 int64 = 9223372036854775807
    v20 := []int{127, 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
