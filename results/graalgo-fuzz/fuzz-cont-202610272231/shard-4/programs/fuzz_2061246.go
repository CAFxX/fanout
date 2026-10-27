package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int }
type GenS4 struct { F0 int32 }

func main() {
    v1 := []uint32{383}
    var v2 uint = 4294967295
    v3 := make(map[string]string)
    var v4 int8 = -1
    for fi5 := 0; fi5 < 9; fi5++ {
        if (-2147483648 <= fi5) {
            var v6 uint64 = 0
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            fmt.Println(v4, v2)
            _ = v6
        } else {
            if ((v4 - (^(v4))) >= (v4 % (((v4 ^ -1)) | 1))) {
                if false {
                    v2 -= v2
                }
            }
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 int8 = 0
                v4 = v4
                v9 := GenS1{F0: (v2 * (v2 * v2)), F1: 2147483647}
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        _ = fi5
    }
    v10 := GenS2{F0: 426, F1: 42}
    v2 += v2
    if (!false) {
        var v11 bool = ((((("\n\t" + "hello") + "") + (("hello" + "foo bar") + "hello")) + ("" + "foo bar")) >= "foo bar")
        v4 = v4
        v12 := GenS3{F0: 74, F1: 1, F2: 668}
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 uint64 = 1
        v15 := GenS2{F0: 255, F1: -271}
        fmt.Println(v14, fi13)
        _ = fi13
        _ = v14
        _ = v15
    }
    v4 -= (v4 ^ (^(v4)))
    v4 = -128
    v1 = append(v1, 625)
    v16 := GenS4{F0: -1}
    fmt.Println(v4, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
