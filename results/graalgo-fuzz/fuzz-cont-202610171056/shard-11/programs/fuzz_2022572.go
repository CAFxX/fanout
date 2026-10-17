package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int16 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint16) string {
    fmt.Println(p0)
    var v7 uint64 = 970
    var v8 uint8 = 69
    v9 := []int{-2147483648, 0}
    for fi10 := 0; fi10 < 12; fi10++ {
        p0 += 1
        _ = fi10
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 18
    var v3 int64 = -883
    var v4 uint16 = 42
    v5 := []string{"a", ("foo bar" + "\n\t")}
    var v6 float32 = (3.14159 / (float32(0.0) + (((3.14159 + (0.1 - 1.5)) + -2.25) * 3.14159)))
    v3 -= 42
    v4 = (v4 | (v4 % ((290) | 1)))
    switch "" {
    case "hello":
        v5 = append(v5, "hello")
    }
    if true {
        v11 := GenS1{F0: 1, F1: 1, F2: 1000}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint = 0
            _ = fi12
            _ = v13
        }
        v5 = append(v5, "hello")
    }
    var v14 uint32 = 2147483647
    v6 += v6
    v5 = append(v5, "foo bar")
    var v15 uint64 = 18446744073709551615
    _ = len(v1)
    v16 := GenS1{F0: 255, F1: 1000, F2: -228}
    v17 := v1[42]
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            fmt.Println(v17, v6, v15)
            fmt.Println(v17, v15, fi19)
            _ = fi19
        }
        var v20 int64 = v3
        v20 = v3
        _ = fi18
        _ = v20
    }
    v21 := GenS2{F0: v3, F1: v14}
    fmt.Println(v6, v6)
    var v22 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
