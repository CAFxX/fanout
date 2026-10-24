package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
func genFn0(p0 int32, p1 bool, p2 uint16) int64 {
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 string = ("a" + "foo bar")
        _ = fi6
        _ = v7
    }
    p0 -= (p0 / ((1) | 1))
    _ = p0
    _ = p1
    _ = p2
    return 892
}

func main() {
    v1 := []string{"\n\t", "hello"}
    var v2 bool = (!false)
    var v3 uint64 = 4294967295
    var v4 string = "\n\t"
    var v5 float64 = 0.0
    var v8 int8 = 127
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (-2147483648 == 0) {
        var v9 int = -2147483648
        var v10 float64 = (-2.25 - -2.25)
        _ = v10
        _ = v9
    }
    if true {
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 uint32 = uint32(v3)
            for fi13 := 0; fi13 < 20; fi13++ {
                v8 = (v8 / (((v8 % (((v8 % (((v8 ^ (v8 & v8))) | 1))) | 1))) | 1))
                _ = fi13
            }
            var v14 int64 = genFn0(0, v2, 65535)
            _ = fi11
            _ = v12
            _ = v14
        }
        if true {
            fmt.Println(v2)
            v15 := GenS1{F0: 1, F1: 281}
            _ = v15
        }
        var v16 int = 127
        _ = v16
    } else {
        switch int(v3) {
        case 974:
            v17 := GenS1{F0: 0, F1: 9223372036854775807}
            _ = v17
        }
        v5 -= (-2.25 + v5)
    }
    v2 = ((v3 / (((v3 ^ (v3 << ((0) & 63)))) | 1)) != v3)
    v2 = v2
    var v18 uint64 = (v3 % ((v3) | 1))
    if (390 < v18) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if len(v1) > 0 {
        v1[0] = "a"
    }
    v19 := GenS1{F0: 2, F1: 122}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
