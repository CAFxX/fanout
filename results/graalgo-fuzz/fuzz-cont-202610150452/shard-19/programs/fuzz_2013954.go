package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int64; F2 int64 }
func genFn0(p0 uint) int {
    var v4 uint64 = 1
    fmt.Println(v4, p0)
    v5 := GenS1{F0: 32767, F1: 2147483647, F2: (^(v4))}
    v6 := GenS1{F0: 0, F1: -172, F2: (v4 % ((v4) | 1))}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 uint8 = 2
    v2 := []int{128, 0, 128, 2147483647}
    v3 := make(map[uint32]string)
    v1 += (^(v1))
    var v7 int64 = 0
    var v8 uint = 2
    switch genFn0((v8 << ((4294967295) & 63))) {
    case 180:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 4294967295:
        _ = len(v3)
    }
    if ("a" >= "hello") {
        for fi9 := 0; fi9 < 18; fi9++ {
            for fi10 := 0; fi10 < 3; fi10++ {
                v1 = v1
                v1 = (v1 ^ 85)
                v8 -= v8
                v11 := GenS1{F0: 32767, F1: 0, F2: uint64(v1)}
                switch ("foo bar" + ("foo bar" + "a")) {
                case "hello":
                    v12 := v3[666]
                    _ = v12
                case "a":
                    if (1 >= 42) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        switch (^(fi10)) {
                        case 32768:
                            v1 += v1
                        case -890:
                            fmt.Println(fi9)
                        }
                    } else {
                        v13 := GenS1{F0: 1, F1: -2147483648, F2: 2}
                        _ = v13
                    }
                default:
                    switch "\n\t" {
                    case "hello":
                        v7 = (v7 / ((9223372036854775807) | 1))
                    case "\n\t":
                        switch fi10 {
                        case 4294967295:
                            var v14 uint = (v8 % (((^(v8))) | 1))
                            v2 = append(v2, fi10)
                            if (!(true || false)) {
                                var v15 int32 = -2147483648
                                var v16 float64 = 1e10
                                _ = v15
                                _ = v16
                            } else {
                                v14 = 615
                                v17 := GenS1{F0: 869, F1: -348, F2: 664}
                                _ = v17
                            }
                            _ = v14
                        case -9223372036854775808:
                            _ = len(v2)
                        case 427:
                            delete(v3, 2)
                        }
                    case "foo bar":
                        v18 := GenS2{F0: 2, F1: (v7 / ((v7) | 1)), F2: (v7 & (v7 + v7))}
                        _ = v18
                    }
                }
                var v19 uint32 = 2
                var v20 float64 = (-2.25 - 3.14159)
                v20 = v20
                _ = fi10
                _ = v11
                _ = v19
                _ = v20
            }
            _ = fi9
        }
        v1 = v1
        v21 := GenS2{F0: 4294967295, F1: v7, F2: v7}
        _ = v21
    }
    fmt.Println(v1, v8)
    var v22 int32 = 534
    v23 := GenS2{F0: 65535, F1: (v7 & (-(v7))), F2: v7}
    var v24 uint16 = 690
    fmt.Println(v8, v22, v8)
    for fi25 := 0; fi25 < 16; fi25++ {
        if (uint32(v24) <= 65535) {
            v7 += v7
        }
        fmt.Println(v7)
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v7
    _ = v8
}
