package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint8) bool {
    p0 = 133
    v6 := GenS1{F0: 128}
    switch 2147483647 {
    case -1:
        v7 := GenS2{F0: 0, F1: 1}
        _ = v7
    case 128:
        if (((1e10 + 1.5) - -2.25) >= 3.14159) {
            p0 = (p0 % ((p0) | 1))
        }
    case 0:
        var v8 bool = (((1 >= uint(p0)) && (!(!false))) && false)
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := []int{42, (fi9 ^ fi9), (fi9 * fi9)}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v6
    return (((4294967295 <= 32768) || (true || (!((true || ((-128 >= 90) && true)) || false)))) || true)
}

func main() {
    var v1 uint16 = 1
    var v2 int64 = -673
    var v3 uint16 = (v1 >> ((2) & 63))
    var v4 uint32 = 949
    var v5 bool = (!false)
    v11 := []int{}
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 int = 995
            v4 += v4
            _ = fi13
            _ = v14
        }
        var v15 uint8 = 1
        _ = fi12
        _ = v15
    }
    fmt.Println(v5, v3)
    v1 -= (v1 & (v3 | 2))
    v4 = 42
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 int32 = -2147483648
        fmt.Println(fi16)
        _ = fi16
        _ = v17
    }
    var v18 string = "hello"
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v19 float32 = 3.14159
    switch ("hello" + ((v18 + (((("a" + v18) + (v18 + v18)) + v18) + v18)) + "foo bar")) {
    case "foo bar":
        _ = len(v11)
    case "a":
        v11 = append(v11, 9223372036854775807)
    default:
        v20 := GenS2{F0: 4294967295, F1: (^(v1))}
        _ = v20
    }
    fmt.Println(v3, v1)
    fmt.Println(v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
