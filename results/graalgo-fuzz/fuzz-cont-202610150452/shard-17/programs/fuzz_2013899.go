package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int16 }
func genFn0() int64 {
    var v6 float32 = 3.14159
    v6 += (v6 * 0.0)
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 int16 = -803
    v2 := make(map[uint32]string)
    var v3 uint64 = 18446744073709551615
    var v4 int16 = (^(v1))
    var v5 int64 = 9223372036854775807
    v1 -= 42
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 582, F1: 0, F2: 65535}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            v4 -= v4
            _ = fi10
        }
        _ = fi9
    }
    v11 := GenS2{F0: (v3 % ((v3) | 1)), F1: v1}
    var v12 uint = 662
    v13 := []int{}
    var v14 int32 = -972
    v15 := GenS2{F0: 558, F1: 32767}
    if (!true) {
        v3 = v3
        var v16 uint64 = v3
        v14 = 2147483647
        _ = v16
    }
    switch ("foo bar" + "hello") {
    case "a":
        var v17 uint = 18446744073709551615
        _ = v17
    case "":
        var v18 int32 = (v14 % (((v14 % (((v14 | (v14 / ((v14) | 1)))) | 1))) | 1))
        _ = v18
    case "hello":
        v19 := GenS2{F0: 0, F1: -165}
        _ = v19
    }
    v13 = append(v13, 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
