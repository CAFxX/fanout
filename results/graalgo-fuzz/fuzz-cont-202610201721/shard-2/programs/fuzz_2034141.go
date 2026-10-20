package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0() string {
    var v5 string = ("foo bar" + "hello")
    for fi6 := 0; fi6 < 7; fi6++ {
        switch fi6 {
        case -9223372036854775808:
            v5 = v5
        case 0:
            var v7 string = ("a" + v5)
            _ = v7
        default:
            v8 := GenS1{F0: 0}
            _ = v8
        }
        _ = fi6
    }
    _ = v5
    return v5
}
func genFn1(p0 int16, p1 int32, p2 int) int32 {
    var v9 uint = 18446744073709551615
    var v10 bool = false
    v11 := []int{(p2 | (p2 >> ((65535) & 63))), 1, (p2 ^ (p2 & 509))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return (^(p1))
}

func main() {
    var v1 int16 = 987
    var v2 string = "hello"
    v3 := []string{"hello"}
    var v4 uint64 = 2
    v3 = append(v3, "")
    var v12 uint8 = 42
    if (487 < 2147483648) {
        v13 := GenS1{F0: -1}
        v14 := GenS1{F0: -128}
        _ = v13
        _ = v14
    } else {
        v12 = (v12 << ((1) & 63))
        var v15 uint64 = v4
        v15 += v15
        _ = v15
    }
    var v16 string = (v2 + v2)
    if len(v3) > 0 {
        v3[0] = "a"
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi17
    }
    _ = len(v3)
    fmt.Println(v1)
    v1 += (^(v1))
    v18 := GenS2{F0: -74, F1: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
