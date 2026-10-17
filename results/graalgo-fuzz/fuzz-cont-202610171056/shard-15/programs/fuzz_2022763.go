package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int8 }
func genFn0(p0 uint8) uint {
    if true {
        fmt.Println(p0)
        if true {
            for fi4 := 0; fi4 < 9; fi4++ {
                if true {
                    v5 := GenS1{F0: 1, F1: "a", F2: 1}
                    _ = v5
                }
                p0 = 9
                v6 := GenS1{F0: 65535, F1: (("" + "foo bar") + ((("" + "") + "") + (((("" + "hello") + ("" + "hello")) + "") + "foo bar"))), F2: 120}
                _ = fi4
                _ = v6
            }
            p0 = (p0 + p0)
        } else {
            var v7 float64 = float64(1.5)
            _ = v7
        }
    }
    p0 -= 243
    var v8 int8 = 127
    _ = p0
    _ = v8
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = (-2.25 + 1.5)
    v3 := make(map[int]string)
    var v9 uint16 = 127
    v3[9223372036854775807] = "\n\t"
    v2 = 0.1
    v9 -= v9
    v10 := []int{}
    var v11 bool = (!true)
    fmt.Println(v11)
    v12 := GenS1{F0: v9, F1: "", F2: 127}
    v9 = v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
