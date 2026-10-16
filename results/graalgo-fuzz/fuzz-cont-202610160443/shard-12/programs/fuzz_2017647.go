package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int64; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int8 }

func main() {
    v1 := []string{"", "hello", ("\n\t" + "hello"), "\n\t"}
    v2 := make(map[int]uint32)
    v3 := []string{}
    var v4 uint32 = 42
    var v5 uint32 = 1
    v6 := make(map[string]string)
    var v7 float64 = 0.0
    v8 := v2[9223372036854775807]
    var v9 uint32 = 0
    if (4294967295 >= -486) {
        if ((!(228 != 600)) && (!false)) {
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS1{F0: 42, F1: -128}
                fmt.Println(fi10, v8)
                _ = fi10
                _ = v11
            }
            v12 := GenS2{F0: 9223372036854775807, F1: 42, F2: int32(v9)}
            _ = v12
        }
        v13 := v6[("" + "a")]
        fmt.Println(v7, v5)
        _ = v13
    } else {
        v1 = append(v1, "foo bar")
    }
    var v14 int = 4294967295
    v15 := GenS2{F0: 9223372036854775807, F1: 42, F2: 2147483647}
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v16 := GenS2{F0: -176, F1: 250, F2: 2147483647}
    if true {
        v17 := GenS3{F0: 25, F1: (v5 * 2), F2: 1}
        _ = len(v6)
        var v18 uint = 885
        _ = v17
        _ = v18
    } else {
        v7 -= float64((0.0 - float32(((v7 + v7) * v7))))
        fmt.Println(v7)
        if (833 < 42) {
            var v19 int32 = 828
            _ = v19
        }
    }
    v20 := GenS3{F0: 68, F1: v9, F2: -81}
    v5 = 308
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch (-(v14)) {
    case -120:
        v4 = (v5 << ((255) & 63))
    case 32767:
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 int = (v14 << ((42) & 63))
            var v23 int8 = 0
            fmt.Println(v8, v5)
            _ = fi21
            _ = v22
            _ = v23
        }
    case 1:
        v9 -= v8
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
