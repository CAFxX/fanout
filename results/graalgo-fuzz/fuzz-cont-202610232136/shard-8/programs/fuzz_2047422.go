package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int64 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int32; F1 int32 }
func genFn0() string {
    if true {
        var v7 int16 = -83
        switch "\n\t" {
        case "hello":
            v7 -= (^(v7))
        case "\n\t":
            var v8 int = 42
            _ = v8
        case "a":
            if ((1.5 / float64(1.5)) != 3.14159) {
                var v9 int64 = 1
                if false {
                    fmt.Println(v7, v9)
                    v7 += v7
                    if true {
                        if false {
                            v9 = v9
                            for fi10 := 0; fi10 < 14; fi10++ {
                                v9 = (v9 << ((255) & 63))
                                _ = fi10
                            }
                            v7 = v7
                        }
                        v11 := GenS1{F0: v9, F1: -98, F2: (v9 - (v9 & (v9 >> ((181) & 63))))}
                        for fi12 := 0; fi12 < 16; fi12++ {
                            var v13 string = ("hello" + ("hello" + "a"))
                            v14 := []int{(fi12 * (fi12 / ((fi12) | 1)))}
                            fmt.Println(v13, v9, v9)
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        _ = v11
                    } else {
                        fmt.Println(v7)
                    }
                }
                _ = v9
            }
        }
        _ = v7
    } else {
        var v15 int64 = -656
        switch 432 {
        case -605:
            v15 = v15
        case 42:
            if (32767 <= 0) {
                v16 := []int{-1, 127}
                var v17 int32 = 42
                _ = v16
                _ = v17
            } else {
                fmt.Println(v15)
                v18 := []int{269, 4294967295, 793}
                switch -2147483648 {
                case 127:
                    v18 = append(v18, 32767)
                case 128:
                    v19 := GenS1{F0: (v15 | (v15 % ((v15) | 1))), F1: 1, F2: (v15 >> ((4294967295) & 63))}
                    _ = v19
                default:
                    v15 = v15
                }
                _ = v18
            }
        case 1:
            var v20 uint64 = 0
            _ = v20
        }
        var v21 int32 = 58
        _ = v15
        _ = v21
    }
    var v22 float32 = (-2.25 + 1e10)
    if (true && true) {
        v22 += v22
    }
    _ = v22
    return ""
}

func main() {
    var v1 int16 = 42
    v2 := []uint32{65535, 208, 65535, 42}
    v3 := []int{2147483647}
    var v4 uint64 = 4294967295
    v5 := []int{42, 674}
    v6 := make(map[int]uint32)
    var v23 string = "a"
    switch v23 {
    case "foo bar":
        fmt.Println(v23)
    case "\n\t":
        var v24 int = -9223372036854775808
        _ = v24
    case "a":
        v25 := GenS1{F0: 374, F1: 100, F2: 0}
        _ = v25
    default:
        var v26 uint8 = 1
        _ = v26
    }
    if len(v3) > 0 {
        v3[0] = 127
    }
    var v27 uint32 = 967
    var v28 float32 = 1e10
    var v29 uint64 = v4
    if true {
        v4 += v29
    }
    _ = len(v6)
    fmt.Println(v28, v23, v4)
    v28 += ((v28 * v28) * v28)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v28, v1)
    switch -1 {
    case -419:
        v1 -= (v1 << ((65535) & 63))
    }
    var v30 uint8 = 0
    for fi31 := 0; fi31 < 8; fi31++ {
        if true {
            if (-259 == v1) {
                v23 = v23
                for fi32 := 0; fi32 < 8; fi32++ {
                    v27 -= (v27 >> ((65535) & 63))
                    v27 = v27
                    fmt.Println(v23, v4, v23)
                    _ = fi32
                }
            }
            for fi33 := 0; fi33 < 10; fi33++ {
                v1 -= v1
                v34 := GenS2{F0: 127, F1: 788}
                var v35 int = 32767
                _ = fi33
                _ = v34
                _ = v35
            }
            if len(v5) > 0 {
                v5[0] = (fi31 / (((fi31 * fi31)) | 1))
            }
        }
        v6[(^(fi31))] = v27
        _ = fi31
    }
    v36 := GenS3{F0: 456, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
