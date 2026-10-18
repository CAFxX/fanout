package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 bool; F2 bool }
type GenS3 struct { F0 uint16; F1 string; F2 int16 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 string; F1 uint16; F2 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 0
    var v3 int64 = 0
    var v4 string = "a"
    v5 := make(map[uint32]uint32)
    var v6 float64 = 1e10
    var v7 uint32 = 0
    v8 := GenS1{F0: 9223372036854775807}
    var v9 float32 = 3.14159
    v2 = v2
    fmt.Println(v6, v6)
    var v10 int32 = 2147483647
    if false {
        if ((!(!false)) || false) {
            fmt.Println(v3)
        }
        v11 := GenS1{F0: 42}
        v12 := []int{42}
        _ = v11
        _ = v12
    }
    fmt.Println(v10, v10, v6)
    v6 -= v6
    v13 := v5[65535]
    if true {
        if true {
            fmt.Println(v4, v10, v6)
            v14 := []int{-9223372036854775808}
            _ = v14
        } else {
            var v15 int = -987
            v13 += (v13 & 4294967295)
            _ = v15
        }
        v16 := GenS2{F0: v4, F1: (false && true), F2: false}
        var v17 bool = true
        _ = v16
        _ = v17
    } else {
        v18 := GenS1{F0: 9223372036854775807}
        var v19 uint64 = v2
        switch -2147483648 {
        case 4294967295:
            switch (v4 + "a") {
            case "hello":
                var v20 string = "\n\t"
                _ = v20
            case "\n\t":
                v21 := GenS3{F0: 28, F1: "foo bar", F2: 32767}
                _ = v21
            case "a":
                var v22 int8 = -128
                _ = v22
            default:
                for fi23 := 0; fi23 < 20; fi23++ {
                    v24 := GenS2{F0: v4, F1: true, F2: false}
                    switch "a" {
                    case "\n\t":
                        v19 = v2
                    case "a":
                        switch fi23 {
                        case -2147483648:
                            _ = len(v1)
                        case 4294967295:
                            var v25 uint64 = 18446744073709551615
                            _ = v25
                        case 32768:
                            fmt.Println(v2, v13, v19)
                        default:
                            v26 := GenS4{F0: (v3 << ((0) & 63)), F1: (v7 ^ 65535)}
                            _ = v26
                        }
                    default:
                        switch "foo bar" {
                        case "foo bar":
                            var v27 int = int(v7)
                            _ = v27
                        }
                    }
                    for fi28 := 0; fi28 < 9; fi28++ {
                        v29 := GenS5{F0: v4, F1: 32767, F2: v10}
                        _ = fi28
                        _ = v29
                    }
                    _ = fi23
                    _ = v24
                }
            }
        default:
            switch 4294967295 {
            case 631:
                v30 := GenS1{F0: -2147483648}
                _ = v30
            case 9223372036854775807:
                if (-1 > 42) {
                    if false {
                        var v31 int32 = (v10 / ((2147483647) | 1))
                        _ = v31
                    } else {
                        var v32 int32 = (v10 + 2147483647)
                        _ = v32
                    }
                }
            default:
                fmt.Println(v2, v9, v4)
            }
        }
        _ = v18
        _ = v19
    }
    v33 := GenS4{F0: (v3 / ((v3) | 1)), F1: 2147483647}
    if (!((v6 - (float64((v9 - 3.14159)) - v6)) < float64(float32(v6)))) {
        v10 -= v10
    }
    var v34 uint64 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
