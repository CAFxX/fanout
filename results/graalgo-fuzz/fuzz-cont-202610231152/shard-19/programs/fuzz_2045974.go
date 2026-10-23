package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
func genFn0() int64 {
    var v7 uint = 473
    v8 := []int{9, -209}
    v8 = append(v8, -617)
    _ = v7
    _ = v8
    return 9223372036854775807
}
func genFn1(p0 int32, p1 uint64) uint32 {
    var v9 uint16 = 2
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint64 = 2
        var v12 uint16 = 32767
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v9
    return 468
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]bool)
    var v3 float64 = 0.1
    var v4 int = -419
    var v5 uint32 = 4294967295
    var v6 int = v4
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := []int{(fi13 << ((v1) & 63)), 4294967295, int(v5)}
        v1 -= 18446744073709551615
        _ = fi13
        _ = v14
    }
    if true {
        var v15 string = "hello"
        v2[((("foo bar" + v15) + ((v15 + "foo bar") + "")) + ("foo bar" + ((v15 + v15) + (v15 + ("foo bar" + v15)))))] = (v3 <= 1.5)
        fmt.Println(v4, v4)
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        switch (v4 & (fi16 >> (((v1 >> ((583) & 63))) & 63))) {
        case -2147483648:
            for fi17 := 0; fi17 < 19; fi17++ {
                fmt.Println(fi17, fi16, v3)
                fmt.Println(v3)
                var v18 int8 = -125
                _ = fi17
                _ = v18
            }
        case 32768:
            var v19 uint16 = 32767
            _ = v19
        case 32767:
            v1 = uint(v5)
        default:
            v3 += (v3 - v3)
        }
        _ = fi16
    }
    var v20 string = (("foo bar" + (("hello" + "\n\t") + "")) + (("a" + "foo bar") + ("foo bar" + "a")))
    fmt.Println(v3, v6, v3)
    v21 := GenS1{F0: v5, F1: 42}
    v22 := []int{}
    for fi23 := 0; fi23 < 15; fi23++ {
        v3 += v3
        v2[v20] = false
        _ = fi23
    }
    _ = len(v2)
    v6 = (v4 % ((43) | 1))
    v24 := GenS1{F0: genFn1(1, 281), F1: 1}
    v25 := GenS1{F0: (v5 << ((v1) & 63)), F1: 2147483647}
    if len(v22) > 0 {
        v22[0] = 2147483647
    }
    var v26 uint64 = 0
    v27 := GenS1{F0: v5, F1: 42}
    v26 -= (v26 + 18446744073709551615)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
