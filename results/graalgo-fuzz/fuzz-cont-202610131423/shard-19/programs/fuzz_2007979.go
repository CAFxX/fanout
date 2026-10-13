package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := []int{19, 4294967295}
    v2 := make(map[int]string)
    var v3 int = 4294967295
    var v4 uint16 = 32767
    v5 := []string{}
    var v6 int64 = 419
    var v7 uint64 = 4294967295
    var v8 uint64 = uint64(v4)
    var v9 int16 = 1000
    var v10 int8 = 1
    v11 := GenS1{F0: 42, F1: 4294967295}
    var v12 uint = 18446744073709551615
    var v13 string = ("" + "\n\t")
    if (((986 != 2147483647) || false) || ((v12 & 84) > (v12 - v12))) {
        var v14 uint16 = (^(v4))
        v5 = append(v5, v13)
        _ = v14
    }
    var v15 float64 = (1e10 - float64(3.14159))
    for fi16 := 0; fi16 < 10; fi16++ {
        if (!true) {
            var v17 uint8 = 238
            _ = v17
        }
        switch (("hello" + v13) + v13) {
        case "\n\t":
            _ = len(v1)
        case "":
            if ((^(v9)) >= v9) {
                var v18 float64 = v15
                _ = v18
            } else {
                if ((v12 | v12) > v12) {
                    v2[0] = v13
                } else {
                    var v19 int32 = 0
                    var v20 float64 = 0.0
                    if len(v5) > 0 {
                        v5[0] = (v13 + "hello")
                    }
                    _ = v19
                    _ = v20
                }
                var v21 int8 = (-(v10))
                _ = v21
            }
        case "foo bar":
            fmt.Println(v6, v7)
        }
        fmt.Println(v3, v12, v13)
        _ = fi16
    }
    v22 := GenS2{F0: v4, F1: v3, F2: v10}
    v23 := GenS3{F0: 65535, F1: 1}
    v9 += v9
    var v24 uint64 = v8
    v6 = (v6 ^ (v6 & (v6 << (((v12 % ((v12) | 1))) & 63))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
