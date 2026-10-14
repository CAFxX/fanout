package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int }
type GenS2 struct { F0 uint8; F1 int16; F2 uint64 }
func genFn0(p0 int64, p1 int16) uint16 {
    fmt.Println(p1)
    p1 += (p1 / ((p1) | 1))
    var v4 bool = (792 <= 255)
    _ = p0
    _ = p1
    _ = v4
    return 2
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    if (genFn0(1, 0) != 1) {
        v1 = (v1 | (v1 + v1))
    }
    switch 9223372036854775807 {
    case -1:
        if ((-786 > -1) && (!(!true))) {
            v1 = 0
        }
    case 4294967295:
        if true {
            var v5 int8 = 60
            var v6 uint8 = 2
            _ = v5
            _ = v6
        } else {
            var v7 uint64 = (v1 << ((333) & 63))
            _ = v7
        }
    case -2147483648:
        v8 := GenS1{F0: 128, F1: 589, F2: -2147483648}
        _ = v8
    }
    v9 := GenS2{F0: 2, F1: 32767, F2: (v1 - v1)}
    for fi10 := 0; fi10 < 4; fi10++ {
        if (!(false && false)) {
            v11 := []int{9223372036854775807, fi10, (fi10 + (^(fi10)))}
            var v12 int32 = 42
            _ = v11
            _ = v12
        }
        _ = fi10
    }
    v2[127] = "a"
    if true {
        fmt.Println(v1)
        v1 += (v1 ^ 0)
        v1 += (^(v1))
    }
    v1 = 18446744073709551615
    v13 := GenS1{F0: 461, F1: 127, F2: 9223372036854775807}
    fmt.Println(v1)
    v1 = (v1 - v1)
    v1 = v1
    var v14 string = ""
    v14 = ("a" + (v14 + (v14 + (v14 + (("\n\t" + "") + (("\n\t" + v14) + "\n\t"))))))
    v15 := []int{128, -2147483648, -733}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
