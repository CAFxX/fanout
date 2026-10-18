package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint8, p1 uint32) int16 {
    switch "a" {
    case "":
        v5 := []int{1, 103, 1, 128}
        _ = v5
    case "a":
        var v6 int8 = -1
        _ = v6
    default:
        p1 += (p1 | (p1 >> ((255) & 63)))
    }
    if (true || (-444 != -1)) {
        var v7 string = "hello"
        v8 := GenS1{F0: 18446744073709551615, F1: uint(p1)}
        _ = v7
        _ = v8
    } else {
        if true {
            if false {
                p0 -= 253
                v9 := GenS2{F0: 42, F1: 0}
                _ = v9
            }
            var v10 int32 = 1
            _ = v10
        }
    }
    _ = p0
    _ = p1
    return -141
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[int]uint32)
    var v3 float64 = (3.14159 / 1.5)
    var v4 uint64 = 187
    v11 := []int{}
    v4 = 18446744073709551615
    delete(v2, int(v4))
    switch "foo bar" {
    case "hello":
        v4 -= (v4 * (v4 - (^(v4))))
    case "\n\t":
        delete(v2, (v1 / ((v1) | 1)))
    case "":
        _ = len(v11)
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint64 = v4
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v4 -= v4
    v3 += v3
    v2[-751] = 773
    v3 -= 0.1
    var v15 int32 = 2147483647
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := v2[fi17]
            _ = fi17
            _ = v18
        }
        _ = len(v2)
        _ = fi16
    }
    v1 += v1
    for fi19 := 0; fi19 < 13; fi19++ {
        fmt.Println(v15, v1)
        v2[(v1 | (fi19 / (((v1 | (fi19 & (v1 - v1)))) | 1)))] = 0
        v4 += 764
        _ = fi19
    }
    if true {
        var v20 int8 = 0
        v20 = v20
        _ = v20
    }
    var v21 float64 = (0.0 * 1e10)
    v22 := v2[v1]
    var v23 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
