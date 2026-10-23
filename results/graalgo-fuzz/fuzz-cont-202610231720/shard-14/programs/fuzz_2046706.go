package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int64; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int64) bool {
    v6 := GenS1{F0: 42}
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{(fi7 / ((fi7) | 1))}
        v9 := GenS1{F0: 4294967295}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 0}
    p1 = 9223372036854775807
    var v11 int16 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v6
    return (1 <= 6)
}
func genFn1(p0 string, p1 int8) uint32 {
    v12 := GenS2{F0: 42, F1: 706}
    fmt.Println(p0, p0)
    for fi13 := 0; fi13 < 14; fi13++ {
        switch (fi13 & (fi13 / (((fi13 % (((fi13 - (fi13 * fi13))) | 1))) | 1))) {
        case 318:
            var v14 string = ((p0 + ((((p0 + "hello") + p0) + p0) + "\n\t")) + "a")
            _ = v14
        case 765:
            fmt.Println(fi13, p0, p1)
        case 4294967295:
            var v15 uint32 = 1
            _ = v15
        default:
            var v16 uint16 = 255
            _ = v16
        }
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v12
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = -32768
    var v3 uint32 = 2147483648
    v4 := []uint32{(v3 / (((v3 / (((v3 - v3)) | 1))) | 1)), v3, (v3 % ((v3) | 1)), 2147483648}
    var v5 float32 = 0.1
    v3 = (v3 & 2)
    v17 := GenS3{F0: 0, F1: -1, F2: genFn1("foo bar", -13)}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v18 := GenS4{F0: 0, F1: 65535}
    if false {
        var v19 uint = 600
        _ = v19
    }
    var v20 float64 = float64(3.14159)
    delete(v1, v3)
    var v21 uint64 = 1
    var v22 uint32 = (^(v3))
    v23 := GenS5{F0: 255, F1: 65535}
    v2 += (^(v2))
    fmt.Println(v22)
    var v24 int64 = 9223372036854775807
    v4 = append(v4, (v3 | v3))
    switch "\n\t" {
    case "":
        v25 := GenS6{F0: ""}
        _ = v25
    case "foo bar":
        switch "" {
        case "hello":
            if len(v4) > 0 {
                v4[0] = genFn1("foo bar", 42)
            }
        case "\n\t":
            var v26 int64 = v24
            _ = v26
        case "":
            v27 := GenS7{F0: 4294967295}
            _ = v27
        }
    }
    v20 += v20
    var v28 int16 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
