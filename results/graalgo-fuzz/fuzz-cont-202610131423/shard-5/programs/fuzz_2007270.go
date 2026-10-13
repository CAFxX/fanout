package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint, p1 int) uint8 {
    p0 = 18446744073709551615
    if (p1 <= (p1 + 1)) {
        var v7 uint16 = 0
        p1 = -743
        _ = v7
    }
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v8
    return 42
}
func genFn1(p0 int) int8 {
    v9 := []int{p0, 9223372036854775807, (p0 ^ (p0 << ((65535) & 63))), p0}
    p0 += p0
    v10 := GenS1{F0: 2}
    _ = p0
    _ = v10
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 18446744073709551615
    var v3 int8 = -57
    var v4 int64 = 9223372036854775807
    var v5 int16 = 42
    var v6 float32 = 1e10
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 int16 = (v5 ^ v5)
        switch "foo bar" {
        case "a":
            var v13 uint64 = v2
            _ = v13
        default:
            if true {
                v2 -= v2
                v14 := GenS2{F0: (-(v12)), F1: 84}
                fmt.Println(v4, v3, fi11)
                _ = v14
            } else {
                v4 = (v4 + (v4 | (^(v4))))
                fmt.Println(v6, v12)
                v15 := GenS1{F0: 0}
                _ = v15
            }
        }
        var v16 int32 = 2147483647
        _ = fi11
        _ = v12
        _ = v16
    }
    v17 := GenS1{F0: 933}
    v5 = -1
    _ = len(v1)
    var v18 string = ""
    if (!((v3 % ((genFn1(79)) | 1)) >= (v3 % ((-59) | 1)))) {
        var v19 string = "a"
        _ = v19
    }
    delete(v1, -2147483648)
    fmt.Println(v6, v3)
    v1[-188] = v18
    var v20 uint64 = v2
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
