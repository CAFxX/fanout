package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int }

func main() {
    var v1 float64 = (-2.25 * 1.5)
    v2 := make(map[uint32]bool)
    var v3 uint32 = 4294967295
    v4 := []int{32767}
    v5 := GenS1{F0: v3, F1: (((((("foo bar" + ("a" + "a")) + "") + "foo bar") + "\n\t") + (("a" + ("a" + "\n\t")) + "foo bar")) + ("a" + "foo bar")), F2: 903}
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint16 = 1
        v2[v3] = (1 < 0)
        _ = fi6
        _ = v7
    }
    var v8 uint16 = 634
    v8 += v8
    v9 := v2[(v3 + v3)]
    if v9 {
        _ = len(v4)
        v1 -= v1
    } else {
        var v10 uint = 1
        _ = v10
    }
    fmt.Println(v9, v8, v1)
    var v11 uint64 = 2
    if v9 {
        var v12 uint = 42
        for fi13 := 0; fi13 < 7; fi13++ {
            if v9 {
                var v14 int32 = 2147483647
                var v15 uint64 = (v11 / (((^(v11))) | 1))
                _ = v14
                _ = v15
            }
            _ = fi13
        }
        _ = v12
    } else {
        v1 += (1e10 + float64((((-2.25 + 0.1) / 0.1) + float32(v1))))
        v16 := GenS1{F0: v3, F1: "foo bar", F2: 32767}
        _ = v16
    }
    v11 = 505
    v3 -= v3
    v3 = v3
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 uint8 = 1
            if ((0 <= 0) && v9) {
                if (!(!(!true))) {
                    var v20 uint16 = (v8 - v8)
                    _ = v20
                } else {
                    v21 := GenS1{F0: (v3 - v3), F1: ("a" + ""), F2: fi17}
                    var v22 uint32 = 393
                    var v23 uint = 1
                    var v24 string = (("\n\t" + "hello") + "foo bar")
                    v9 = ((v3 >> (((v23 >> (((^(v23))) & 63))) & 63)) <= 42)
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                var v25 uint64 = (v11 | 0)
                v11 = (^(v25))
                _ = v25
            }
            _ = fi18
            _ = v19
        }
        delete(v2, v3)
        _ = fi17
    }
    v3 = v3
    v1 = 0.1
    var v26 uint = 595
    v9 = (32767 >= -32768)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
