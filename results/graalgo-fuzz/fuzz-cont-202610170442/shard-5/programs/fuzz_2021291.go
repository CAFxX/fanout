package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0(p0 bool) uint32 {
    switch "a" {
    case "a":
        for fi4 := 0; fi4 < 16; fi4++ {
            p0 = (42 != -1)
            _ = fi4
        }
    case "hello":
        var v5 uint16 = 1
        _ = v5
    default:
        var v6 uint64 = 18446744073709551615
        _ = v6
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        if p0 {
            v8 := GenS1{F0: 141}
            if true {
                v9 := GenS2{F0: 32768, F1: 2, F2: 127}
                _ = v9
            } else {
                switch (fi7 / ((fi7) | 1)) {
                case -9223372036854775808:
                    if p0 {
                        var v10 int64 = -1
                        v10 = v10
                        var v11 string = "a"
                        _ = v10
                        _ = v11
                    }
                case -656:
                    if p0 {
                        fmt.Println(fi7, fi7)
                        var v12 uint64 = 667
                        v12 -= v12
                        _ = v12
                    }
                default:
                    p0 = (!p0)
                }
                v13 := []int{fi7}
                switch (("" + (("\n\t" + "foo bar") + "foo bar")) + ("a" + "foo bar")) {
                case "hello":
                    v14 := GenS1{F0: 0}
                    _ = v14
                case "":
                    var v15 uint64 = 632
                    _ = v15
                default:
                    var v16 uint16 = 136
                    _ = v16
                }
                _ = v13
            }
            v17 := GenS1{F0: 126}
            _ = v17
            _ = v8
        } else {
            var v18 uint32 = 2147483648
            p0 = (!((!p0) && p0))
            _ = v18
        }
        _ = fi7
    }
    v19 := GenS3{F0: 42, F1: 213}
    for fi20 := 0; fi20 < 6; fi20++ {
        if ((-436 >= 32767) || p0) {
            var v21 uint32 = 2147483647
            p0 = false
            var v22 uint8 = 1
            _ = v21
            _ = v22
        }
        _ = fi20
    }
    switch "\n\t" {
    case "foo bar":
        var v23 bool = (p0 && p0)
        _ = v23
    default:
        v24 := GenS3{F0: 127, F1: 71}
        _ = v24
    }
    _ = p0
    _ = v19
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []uint32{(v1 - (v1 & 606))}
    var v3 int8 = -128
    fmt.Println(v3)
    fmt.Println(v1)
    v25 := GenS2{F0: 128, F1: 300, F2: 2}
    fmt.Println(v1, v1)
    fmt.Println(v1)
    v26 := GenS1{F0: 0}
    var v27 uint32 = (v1 / ((v1) | 1))
    v1 += v27
    v3 -= int8(v1)
    fmt.Println(v27)
    switch "" {
    case "":
        v27 = v1
    default:
        v27 = v1
    }
    switch "foo bar" {
    case "":
        var v28 int8 = (v3 >> ((563) & 63))
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
