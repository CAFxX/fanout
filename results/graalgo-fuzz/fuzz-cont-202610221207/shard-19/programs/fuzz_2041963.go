package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    var v3 float32 = -2.25
    var v4 float32 = -2.25
    var v5 string = (("" + (("foo bar" + "") + "\n\t")) + (("" + "\n\t") + ("foo bar" + "\n\t")))
    var v6 int64 = -9223372036854775808
    _ = len(v1)
    var v7 uint16 = 886
    if false {
        v8 := GenS1{F0: 0, F1: 9223372036854775807}
        v9 := []int{-1, -2147483648}
        for fi10 := 0; fi10 < 5; fi10++ {
            v7 = (^(v7))
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi10
        }
        _ = v8
        _ = v9
    } else {
        if (!true) {
            v3 += 0.0
        }
        if (523 != (v6 >> ((4294967295) & 63))) {
            var v11 uint64 = 18446744073709551615
            var v12 int8 = 127
            v13 := GenS1{F0: 18446744073709551615, F1: (v6 >> ((65535) & 63))}
            _ = v11
            _ = v12
            _ = v13
        }
        switch -57 {
        case -823:
            var v14 bool = false
            _ = v14
        case 915:
            _ = len(v2)
        }
    }
    v15 := []int{744, -2147483648}
    fmt.Println(v6, v6, v5)
    var v16 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
