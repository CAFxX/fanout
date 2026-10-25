package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int, p1 string, p2 int64) uint64 {
    switch ("a" + ((p1 + p1) + ((p1 + ("" + (("a" + p1) + p1))) + "\n\t"))) {
    case "hello":
        p1 = "a"
    }
    fmt.Println(p1, p1, p0)
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(fi7)
        if (-1 == 42) {
            p2 -= p2
            p0 = fi7
        } else {
            fmt.Println(p0)
            for fi8 := 0; fi8 < 9; fi8++ {
                p0 -= (p0 * 32768)
                if false {
                    p2 = -1
                }
                var v9 float64 = 0.0
                var v10 uint8 = 1
                _ = fi8
                _ = v10
                _ = v9
            }
            v11 := GenS1{F0: 2}
            _ = v11
        }
        p1 = "foo bar"
        _ = fi7
    }
    p0 = (p0 ^ int(p2))
    _ = p0
    _ = p1
    _ = p2
    return 479
}
func genFn1(p0 int8, p1 int32, p2 int32) bool {
    v12 := []int{-2147483648, -1, 128}
    if (!false) {
        var v13 bool = (0.0 < 1.5)
        _ = v13
    } else {
        v14 := GenS2{F0: (p0 % ((-1) | 1)), F1: -984}
        var v15 float32 = ((1e10 * 3.14159) + 3.14159)
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return (4294967295 != 2)
}

func main() {
    var v1 uint16 = 65535
    v2 := []int{128}
    var v3 uint8 = 42
    var v4 bool = ((true && true) || (869 > 0))
    var v5 uint64 = 4294967295
    var v6 string = "foo bar"
    fmt.Println(v3)
    v1 += 255
    var v16 string = v6
    v17 := GenS2{F0: 100, F1: 1000}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch v16 {
    case "hello":
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 int16 = 376
            _ = fi18
            _ = v19
        }
    }
    v6 = "foo bar"
    v5 -= v5
    var v20 uint8 = (v3 & v3)
    v4 = (!v4)
    var v21 uint = 0
    for fi22 := 0; fi22 < 6; fi22++ {
        fmt.Println(v21)
        v23 := GenS3{F0: 1}
        _ = fi22
        _ = v23
    }
    v24 := GenS3{F0: 2147483647}
    if v4 {
        var v25 int32 = 42
        fmt.Println(v3)
        var v26 float64 = 3.14159
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
