package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 string }
type GenS2 struct { F0 int64 }
func genFn0(p0 int, p1 uint16) uint {
    if true {
        var v6 uint64 = 0
        if false {
            v7 := GenS1{F0: 65535, F1: "\n\t", F2: ("a" + "hello")}
            v8 := GenS1{F0: 2, F1: "\n\t", F2: ("\n\t" + "a")}
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 uint8 = 74
                var v11 float32 = ((-2.25 - (-2.25 * ((0.0 * (0.1 * (-2.25 - 0.0))) + (1e10 - -2.25)))) / 0.1)
                v12 := []int{}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            _ = v7
            _ = v8
        }
        _ = v6
    }
    p0 -= p0
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "\n\t"
    v3 := make(map[string]uint32)
    var v4 int16 = 757
    v5 := []int{-9223372036854775808, -9223372036854775808, 0, 128}
    v2 = (("" + v2) + "hello")
    var v13 uint16 = 32767
    v2 = ((v2 + "\n\t") + ("hello" + v2))
    v13 -= (^(v13))
    _ = len(v1)
    var v14 uint8 = 0
    v15 := GenS2{F0: 0}
    var v16 bool = false
    v2 = v2
    _ = len(v3)
    _ = len(v3)
    v13 -= (v13 * 65535)
    v2 = (v2 + "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
