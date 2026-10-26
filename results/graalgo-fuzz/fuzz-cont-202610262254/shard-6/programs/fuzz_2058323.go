package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int16 }
func genFn0(p0 int32) bool {
    v6 := []int{42, 2147483647}
    switch "a" {
    case "a":
        v6 = append(v6, 127)
    case "foo bar":
        switch "hello" {
        case "\n\t":
            p0 = p0
        case "foo bar":
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 uint = 18446744073709551615
                _ = fi7
                _ = v8
            }
        }
    case "\n\t":
        p0 += (p0 / (((p0 >> ((4294967295) & 63))) | 1))
    }
    _ = p0
    _ = v6
    return (0 > 0)
}
func genFn1(p0 int16, p1 uint8, p2 bool) uint16 {
    v9 := []int{127, 32768, 128}
    var v10 int = 224
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 65535
}

func main() {
    var v1 int32 = -38
    var v2 int64 = 954
    v3 := []int{-653}
    var v4 int16 = 32767
    v5 := make(map[int]uint32)
    v11 := GenS1{F0: 87}
    v2 += v2
    if false {
        var v12 int16 = (^(v4))
        var v13 int16 = -897
        _ = v12
        _ = v13
    } else {
        if true {
            var v14 uint8 = 0
            fmt.Println(v4, v2, v4)
            _ = v14
        }
        if true {
            var v15 int32 = v1
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            fmt.Println(v2)
            _ = v15
        } else {
            var v16 uint64 = 18446744073709551615
            var v17 string = "a"
            if (!true) {
                var v18 uint16 = 2
                _ = v18
            } else {
                v16 -= 18446744073709551615
            }
            _ = v16
            _ = v17
        }
    }
    v2 = (v2 % (((v2 | v2)) | 1))
    if len(v3) > 0 {
        v3[0] = 0
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := GenS1{F0: 127}
        switch "hello" {
        case "a":
            if true {
                v4 += (v4 - (v4 << ((18446744073709551615) & 63)))
            }
        default:
            for fi21 := 0; fi21 < 5; fi21++ {
                v1 += v1
                v22 := GenS1{F0: 137}
                v23 := GenS1{F0: 255}
                fmt.Println(v2, fi19, fi21)
                _ = len(v5)
                _ = fi21
                _ = v22
                _ = v23
            }
        }
        _ = fi19
        _ = v20
    }
    v2 -= v2
    v24 := GenS2{F0: 0}
    if ((!(255 <= 481)) || (4294967295 >= 2)) {
        var v25 bool = (4294967295 != 1)
        _ = v25
    } else {
        if (-471 != -1) {
            v26 := GenS2{F0: 42}
            if false {
                var v27 int8 = 86
                v27 = 1
                _ = v27
            } else {
                var v28 int = 127
                delete(v5, (v28 & 127))
                var v29 uint64 = 623
                _ = v28
                _ = v29
            }
            _ = v26
        } else {
            _ = len(v3)
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 uint16 = 655
                v32 := GenS3{F0: 65535, F1: (v31 + (^(v31))), F2: (v4 << ((65535) & 63))}
                _ = fi30
                _ = v31
                _ = v32
            }
        }
    }
    v33 := GenS4{F0: 2, F1: 18, F2: 32767}
    for fi34 := 0; fi34 < 2; fi34++ {
        fmt.Println(v4, fi34)
        v2 = v2
        _ = fi34
    }
    fmt.Println(v1, v4)
    fmt.Println(v4, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
