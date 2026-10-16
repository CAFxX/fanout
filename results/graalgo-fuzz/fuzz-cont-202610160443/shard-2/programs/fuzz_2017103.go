package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
func genFn0(p0 uint16, p1 uint64) bool {
    p1 += 0
    p1 -= p1
    var v6 bool = ((255 >= 256) || false)
    _ = p0
    _ = p1
    _ = v6
    return (127 > -2147483648)
}

func main() {
    var v1 uint32 = 2147483647
    var v2 bool = false
    v3 := []uint32{v1, v1}
    var v4 float32 = 0.0
    v5 := []uint32{(v1 * (v1 * (v1 ^ v1))), (^(v1))}
    v7 := GenS1{F0: 42, F1: 0}
    v2 = (((!((-1 == 2147483647) && (42 >= 2))) || v2) || v2)
    fmt.Println(v4, v4)
    v8 := GenS1{F0: 1, F1: -25}
    v4 = v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = v4
    switch 42 {
    case -978:
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 bool = v2
            switch "a" {
            case "\n\t":
                var v11 int64 = 46
                _ = v11
            }
            var v12 string = "hello"
            _ = fi9
            _ = v10
            _ = v12
        }
    }
    var v13 uint64 = 954
    v14 := GenS1{F0: -647, F1: 0}
    var v15 int16 = 817
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
