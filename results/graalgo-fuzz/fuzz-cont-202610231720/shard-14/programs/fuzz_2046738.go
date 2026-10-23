package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS4 struct { F0 string }
func genFn0(p0 int, p1 uint, p2 string) int32 {
    p2 = (p2 + "")
    p2 = p2
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}
func genFn1(p0 uint32, p1 int16) uint16 {
    v5 := GenS1{F0: (p1 | p1), F1: 65535, F2: 847}
    p0 = (^(p0))
    var v6 int = -9223372036854775808
    v7 := GenS1{F0: (p1 % ((p1) | 1)), F1: 0, F2: 42}
    fmt.Println(v6, v6, p1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 uint32 = 590
    var v8 int32 = 213
    fmt.Println(v4, v1, v1)
    var v9 int64 = 42
    var v10 uint64 = 18446744073709551615
    var v11 string = v1
    switch (v11 + "\n\t") {
    case "":
        v12 := []int{9223372036854775807, 1}
        _ = v12
    case "\n\t":
        v13 := []int{int(v10), 32768}
        _ = v13
    }
    var v14 int64 = (^(v9))
    if true {
        if true {
            var v15 float64 = -2.25
            _ = v15
        } else {
            var v16 uint = 65535
            if true {
                v17 := GenS2{F0: v4, F1: -567}
                var v18 uint8 = 0
                var v19 uint16 = uint16(v18)
                _ = v17
                _ = v18
                _ = v19
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                var v21 uint = (v16 + v16)
                _ = fi20
                _ = v21
            }
            _ = v16
        }
        v1 = ((("" + v1) + (v11 + "foo bar")) + ("\n\t" + "\n\t"))
    }
    var v22 int8 = -106
    var v23 uint64 = v10
    if (-9223372036854775808 == (v14 >> ((1) & 63))) {
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 float32 = 0.1
            _ = fi24
            _ = v25
        }
    } else {
        switch -973 {
        case 4294967295:
            v23 -= 42
        case 321:
            if false {
                switch -2147483648 {
                case 0:
                    switch (v1 + v11) {
                    case "\n\t":
                        v10 = v10
                    }
                default:
                    v26 := []int{-1}
                    _ = v26
                }
                v27 := GenS1{F0: -1, F1: 0, F2: (v8 + v8)}
                v14 += v14
                _ = v27
            } else {
                v28 := GenS3{F0: -1, F1: 255, F2: v23}
                _ = v28
            }
        case 942:
            for fi29 := 0; fi29 < 15; fi29++ {
                v30 := GenS4{F0: v1}
                var v31 int16 = -32768
                fmt.Println(v4)
                _ = fi29
                _ = v30
                _ = v31
            }
        }
    }
    v32 := v2[-791]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v8
    _ = v9
}
