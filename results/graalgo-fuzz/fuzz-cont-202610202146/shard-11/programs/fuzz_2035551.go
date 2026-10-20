package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
func genFn0(p0 string, p1 int64, p2 int) uint8 {
    v5 := []int{}
    var v6 int32 = -1
    var v7 int32 = (v6 % ((v6) | 1))
    var v8 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 uint16, p1 uint16, p2 bool) int {
    var v9 float32 = 0.1
    switch -465 {
    case 32767:
        if p2 {
            var v10 float64 = 3.14159
            fmt.Println(p1)
            var v11 float64 = float64((v9 + (float32(v10) + v9)))
            _ = v10
            _ = v11
        } else {
            if ((!false) && p2) {
                v9 += float32(1.5)
            } else {
                p2 = true
                var v12 string = ("a" + "")
                _ = v12
            }
        }
    default:
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 int64 = 0
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -2147483648
}

func main() {
    v1 := []uint32{2, 42}
    var v2 int64 = -9223372036854775808
    var v3 uint8 = 0
    var v4 string = "\n\t"
    v4 = v4
    v4 = ((((v4 + ("\n\t" + (((v4 + "") + "\n\t") + v4))) + v4) + "\n\t") + (("a" + "foo bar") + v4))
    for fi15 := 0; fi15 < 16; fi15++ {
        v3 = (^(v3))
        _ = fi15
    }
    fmt.Println(v4)
    var v16 uint64 = 1
    if (65535 >= 65535) {
        v17 := GenS1{F0: "hello", F1: -2147483648, F2: 1}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            fmt.Println(v4, v2)
            v2 -= 42
            _ = fi18
        }
    }
    v1 = append(v1, 2147483648)
    if (!(genFn0("hello", v2, genFn1(39, 255, false)) > v3)) {
        v1 = append(v1, 8)
    } else {
        v19 := GenS1{F0: v4, F1: 2147483647, F2: 2147483648}
        _ = v19
    }
    fmt.Println(v3, v3)
    var v20 uint8 = 42
    var v21 bool = true
    var v22 uint = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
