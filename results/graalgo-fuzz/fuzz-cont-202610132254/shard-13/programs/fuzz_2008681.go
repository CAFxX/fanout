package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint; F1 int32; F2 int64 }
func genFn0(p0 string) int8 {
    var v7 int64 = -911
    v7 = (v7 >> ((2) & 63))
    _ = p0
    _ = v7
    return -1
}
func genFn1(p0 int8, p1 bool) string {
    var v8 int16 = 32767
    for fi9 := 0; fi9 < 4; fi9++ {
        switch (fi9 + 0) {
        case 1:
            v10 := GenS1{F0: p0, F1: 1}
            _ = v10
        }
        _ = fi9
    }
    var v11 string = "hello"
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return (("hello" + v11) + v11)
}

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[int]bool)
    var v3 bool = false
    var v4 int8 = -104
    var v5 float32 = 3.14159
    var v6 uint32 = 5
    var v12 uint64 = 4
    fmt.Println(v12, v6, v5)
    var v13 float64 = 0.0
    var v14 uint = 42
    switch 0 {
    case 9223372036854775807:
        v3 = true
    }
    v6 -= 2147483648
    switch 0 {
    case 0:
        v15 := GenS1{F0: (-(v4)), F1: -9223372036854775808}
        _ = v15
    }
    v16 := GenS2{F0: (v14 % ((v14) | 1)), F1: 2147483647, F2: -1}
    var v17 uint8 = 42
    v18 := GenS2{F0: 544, F1: 0, F2: int64(v6)}
    if false {
        v13 = (1.5 - float64(float32(-2.25)))
    } else {
        v19 := GenS2{F0: (v14 - (v14 ^ (^(v14)))), F1: 1, F2: 1}
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
