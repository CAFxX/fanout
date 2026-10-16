package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []uint32{42}
    var v2 bool = false
    v3 := make(map[int]string)
    v4 := []uint32{}
    var v5 int16 = -32768
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: (fi6 / (((fi6 | fi6)) | 1)), F1: (fi6 >> ((65535) & 63)), F2: 2}
        fmt.Println(v2, v2, v2)
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 42
    var v9 int16 = v5
    var v10 float64 = (1.5 - -2.25)
    v8 = (v8 - (v8 << ((2) & 63)))
    switch "" {
    case "hello":
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 int = fi11
            v12 = v12
            var v13 int16 = 1000
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    v9 += v9
    fmt.Println(v2, v2)
    var v14 uint16 = 255
    var v15 int = 1
    var v16 uint = 42
    for fi17 := 0; fi17 < 14; fi17++ {
        v18 := v3[-287]
        if v2 {
            var v19 string = "\n\t"
            if ((1e10 - (0.0 + 0.0)) <= -2.25) {
                v20 := GenS2{F0: int32(v15)}
                v18 = "a"
                var v21 int32 = -2147483648
                _ = v20
                _ = v21
            }
            v5 += v5
            _ = v19
        }
        v10 = ((((1.5 / v10) / 1.0) + v10) * (v10 + (v10 / ((v10 / v10) - v10))))
        _ = fi17
        _ = v18
    }
    var v22 float32 = (float32(((v10 / (v10 - (v10 - v10))) + 0.0)) / 3.14159)
    if v2 {
        v9 -= v5
        var v23 int64 = 9223372036854775807
        fmt.Println(v14)
        _ = v23
    } else {
        v5 = 369
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        v3[(fi24 % (((v15 >> ((v16) & 63))) | 1))] = "foo bar"
        for fi25 := 0; fi25 < 17; fi25++ {
            v8 += v8
            _ = fi25
        }
        if v2 {
            fmt.Println(v2)
            for fi26 := 0; fi26 < 13; fi26++ {
                v2 = (127 < 1)
                var v27 int = (v15 % ((v15) | 1))
                if (((v15 & (fi26 >> ((749) & 63))) < (v15 | 32768)) || false) {
                    var v28 int8 = 1
                    v3[(-(fi24))] = "a"
                    _ = v28
                } else {
                    if (!v2) {
                        v10 -= float64(float32(1e10))
                        _ = len(v3)
                        var v29 uint64 = 16
                        _ = v29
                    }
                }
                _ = fi26
                _ = v27
            }
        }
        _ = fi24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
