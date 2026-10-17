package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int64; F1 uint16 }

func main() {
    v1 := []uint32{2147483648}
    var v2 int64 = 0
    v3 := make(map[uint32]int)
    var v4 float32 = -2.25
    var v5 uint = 0
    var v6 string = "hello"
    var v7 uint16 = 32767
    if (-128 < -1) {
        for fi8 := 0; fi8 < 8; fi8++ {
            v5 = 0
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 float32 = (float32(float64(v4)) * v4)
                v4 += (((v10 / (v4 * v10)) - 1.5) * v4)
                if (fi8 <= fi9) {
                    v11 := GenS1{F0: (v5 << ((v5) & 63)), F1: 392}
                    _ = v11
                } else {
                    v10 += (v4 * v10)
                }
                if (!((v5 / (((v5 - v5)) | 1)) < v5)) {
                    v12 := GenS1{F0: (v5 ^ v5), F1: 432}
                    _ = v12
                } else {
                    if false {
                        switch v6 {
                        case "hello":
                            var v13 uint16 = (^(v7))
                            _ = v13
                        case "foo bar":
                            fmt.Println(v5)
                            var v14 uint16 = v7
                            _ = v14
                        case "a":
                            var v15 uint8 = 154
                            _ = v15
                        }
                        v2 += v2
                    } else {
                        v3[4294967295] = (fi8 % ((int(v2)) | 1))
                        v4 += 0.0
                        v16 := GenS2{F0: 42, F1: (v7 * v7)}
                        _ = v16
                    }
                    switch (fi8 - fi8) {
                    case -79:
                        var v17 int = (fi9 / ((4294967295) | 1))
                        _ = v17
                    case -1:
                        var v18 uint64 = 2
                        _ = v18
                    }
                    if true {
                        v10 = ((v4 / v10) - v10)
                    }
                }
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        v4 = v4
    }
    if true {
        switch "a" {
        case "foo bar":
            v5 -= 1
        }
    } else {
        for fi19 := 0; fi19 < 3; fi19++ {
            v7 -= v7
            for fi20 := 0; fi20 < 11; fi20++ {
                switch "" {
                case "hello":
                    v7 = (^(v7))
                case "\n\t":
                    v4 = -2.25
                }
                v1 = append(v1, 274)
                var v21 int8 = -1
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        if true {
            v22 := GenS1{F0: v5, F1: 1000}
            _ = v22
        }
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
    }
    v7 += 42
    if true {
        if ((v7 % ((263) | 1)) >= (^(v7))) {
            fmt.Println(v5, v7)
        }
    }
    fmt.Println(v7, v2)
    var v23 int32 = -2147483648
    var v24 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
