package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint; F2 int32 }

func main() {
    var v1 int = -44
    var v2 uint32 = 4294967295
    var v3 int = (v1 + v1)
    var v4 int8 = 0
    v5 := make(map[int]string)
    var v6 int8 = (v4 & 42)
    if (false || false) {
        v7 := []int{v1, v3}
        _ = v7
    }
    v4 = (v6 % ((42) | 1))
    if true {
        v8 := GenS1{F0: v3, F1: 93, F2: -343}
        _ = v8
    } else {
        var v9 uint16 = 397
        if (42 < 200) {
            if (false || ((float32(((1e10 * 0.0) + ((0.0 + 0.0) + 0.1))) + 1.5) >= -2.25)) {
                switch (v3 >> ((255) & 63)) {
                case 378:
                    fmt.Println(v9, v6, v6)
                default:
                    v6 = v6
                }
                var v10 uint32 = 348
                _ = v10
            }
            var v11 float64 = -2.25
            _ = v11
        }
        _ = v9
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(v6)
        fmt.Println(v3)
        switch fi12 {
        case -833:
            fmt.Println(v4, fi12)
        default:
            var v13 int16 = 201
            _ = v13
        }
        _ = fi12
    }
    v2 += (v2 & v2)
    if true {
        v4 -= v4
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 uint64 = 676
            v16 := GenS2{F0: (false && (-2147483648 >= 42))}
            _ = fi14
            _ = v15
            _ = v16
        }
        v6 -= (v6 | (v6 / (((v6 >> ((443) & 63))) | 1)))
    } else {
        v17 := GenS2{F0: true}
        fmt.Println(v1)
        var v18 uint16 = 65535
        _ = v17
        _ = v18
    }
    var v19 bool = (false && true)
    if v19 {
        if (v1 != 197) {
            var v20 bool = v19
            var v21 int16 = 32767
            _ = v20
            _ = v21
        } else {
            fmt.Println(v1)
        }
        v22 := GenS1{F0: 2147483647, F1: 2, F2: -2147483648}
        if (0.1 < (1e10 * (0.0 * 1.5))) {
            for fi23 := 0; fi23 < 10; fi23++ {
                v24 := []int{fi23, (v1 | (^(v3)))}
                _ = fi23
                _ = v24
            }
            v5[v3] = (("" + ("a" + "")) + ("a" + "foo bar"))
            if (!v19) {
                var v25 int64 = 42
                _ = v25
            } else {
                for fi26 := 0; fi26 < 7; fi26++ {
                    var v27 uint8 = 128
                    switch "a" {
                    case "\n\t":
                        v5[v1] = (((("a" + "hello") + (("" + "a") + (("a" + ("a" + "\n\t")) + "foo bar"))) + "a") + ("foo bar" + ""))
                    case "foo bar":
                        v28 := GenS3{F0: v3, F1: 135, F2: 42}
                        _ = v28
                    }
                    _ = fi26
                    _ = v27
                }
            }
        }
        _ = v22
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
