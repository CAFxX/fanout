package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint64; F2 int8 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int64 }

func main() {
    var v1 int8 = 42
    v2 := make(map[int]int)
    var v3 uint8 = 247
    fmt.Println(v1)
    v4 := GenS1{F0: 342}
    v5 := GenS2{F0: 0}
    for fi6 := 0; fi6 < 10; fi6++ {
        v3 += (v3 >> ((1) & 63))
        v1 = v1
        v7 := GenS3{F0: (fi6 + 127), F1: 18446744073709551615, F2: -72}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := []int{fi8, (fi8 - fi8), fi8}
        var v10 uint64 = 18446744073709551615
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int = 1
    fmt.Println(v1, v3)
    var v12 bool = (1e10 < 1e10)
    switch (v11 << ((0) & 63)) {
    case 1:
        delete(v2, (^(v11)))
    case -1:
        switch (^(v11)) {
        case 32767:
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := GenS4{F0: 9223372036854775807, F1: v1}
                v15 := []int{569}
                var v16 string = "\n\t"
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        case -1:
            var v17 uint64 = 1
            _ = v17
        case -9223372036854775808:
            var v18 int16 = -1
            _ = v18
        default:
            var v19 int = 9223372036854775807
            _ = v19
        }
    case -2147483648:
        if (((("foo bar" + ("foo bar" + "\n\t")) + "hello") + ("\n\t" + ((("hello" + "a") + "\n\t") + ("" + ("foo bar" + (("hello" + "hello") + "\n\t")))))) != ("a" + ("hello" + "\n\t"))) {
            v3 += v3
        } else {
            v11 = 32768
            var v20 uint16 = 42
            if v12 {
                switch "foo bar" {
                case "foo bar":
                    switch (v11 * (v11 ^ v11)) {
                    case -893:
                        fmt.Println(v11)
                    case -766:
                        var v21 uint32 = 278
                        _ = v21
                    default:
                        var v22 float64 = 1e10
                        _ = v22
                    }
                default:
                    fmt.Println(v12, v11)
                }
                fmt.Println(v12, v20)
            }
            _ = v20
        }
    }
    var v23 float64 = 1e10
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := []int{fi24}
        v11 = (fi24 & 2147483647)
        switch fi24 {
        case 127:
            v26 := GenS5{F0: 4294967295}
            _ = v26
        case 4294967295:
            for fi27 := 0; fi27 < 8; fi27++ {
                v28 := GenS6{F0: "\n\t", F1: 32767}
                v29 := GenS7{F0: 555}
                v11 = fi24
                _ = fi27
                _ = v28
                _ = v29
            }
        case -2147483648:
            var v30 bool = (!v12)
            _ = v30
        }
        _ = fi24
        _ = v25
    }
    v23 -= 1e10
    var v31 float32 = 0.1
    switch (v11 / ((v11) | 1)) {
    case 2147483647:
        var v32 uint = 16
        _ = v32
    case -206:
        v33 := []int{1, (v11 ^ 32767), v11}
        _ = v33
    default:
        v31 = v31
    }
    v34 := []int{v11, (v11 / ((v11) | 1)), (v11 & -1)}
    v12 = (!v12)
    v23 -= 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v34:", len(v34), v34)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
