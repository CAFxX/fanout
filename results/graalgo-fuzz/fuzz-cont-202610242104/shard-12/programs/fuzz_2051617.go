package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int16 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    if false {
        v4 := GenS1{F0: -2147483648, F1: -9223372036854775808, F2: 32767}
        _ = v4
    }
    var v5 uint = 18446744073709551615
    var v6 string = ("\n\t" + "a")
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint16 = 456
        v9 := GenS1{F0: 1, F1: fi7, F2: int16(v5)}
        var v10 uint64 = 526
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    switch ("foo bar" + v6) {
    case "":
        v11 := GenS2{F0: -9223372036854775808}
        _ = v11
    case "hello":
        v12 := []int{-66, 42, 127}
        _ = v12
    default:
        if ((^(v5)) <= (v5 % ((v5) | 1))) {
            var v13 int16 = 1
            v6 = (((v6 + v6) + (v6 + "")) + (v6 + (v6 + "foo bar")))
            _ = v13
        } else {
            var v14 int = 1
            var v15 bool = true
            v16 := v1[346]
            _ = v14
            _ = v15
            _ = v16
        }
    }
    v6 = "\n\t"
    v5 -= 1
    v5 -= v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
