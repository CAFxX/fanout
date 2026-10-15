package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 bool; F1 uint64; F2 bool }
func genFn0() int16 {
    v5 := []int{}
    switch -493 {
    case 1:
        var v6 int8 = 74
        _ = v6
    default:
        v5 = append(v5, 32767)
    }
    _ = v5
    return 1
}
func genFn1(p0 int, p1 int64) uint64 {
    v7 := []int{9223372036854775807}
    var v8 int8 = 127
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint64 = 197
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 0, F1: 197}
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"a"}
    v3 := make(map[int]uint32)
    v4 := []string{"foo bar", "a", "a"}
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint = 4294967295
        var v14 int16 = 1
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        if (true || (1 == 1000)) {
            switch "a" {
            case "\n\t":
                var v16 float64 = 3.14159
                _ = v16
            case "hello":
                v17 := GenS1{F0: 32767, F1: 2147483648}
                _ = v17
            case "foo bar":
                if true {
                    var v18 float32 = 3.14159
                    v19 := GenS1{F0: -634, F1: 2147483647}
                    var v20 float32 = 1e10
                    _ = v18
                    _ = v19
                    _ = v20
                }
            }
        }
        var v21 uint32 = 507
        _ = fi15
        _ = v21
    }
    v22 := GenS2{F0: 573, F1: 1}
    var v23 int = 32768
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := GenS3{F0: (0.1 == float32(3.14159)), F1: 42, F2: false}
        _ = fi24
        _ = v25
    }
    var v26 float64 = 1.5
    switch (("\n\t" + "a") + "\n\t") {
    case "\n\t":
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 float64 = (v26 / (v26 - v26))
            var v29 int32 = 2147483647
            _ = fi27
            _ = v28
            _ = v29
        }
    case "foo bar":
        v30 := GenS1{F0: 831, F1: 1}
        _ = v30
    case "a":
        v23 = v23
    }
    var v31 uint = 18446744073709551615
    _ = len(v1)
    if (v26 != ((((((v26 / v26) * float64(3.14159)) + (((((1e10 - float64(((0.0 * -2.25) - 3.14159))) + (0.1 + -2.25)) + v26) - ((v26 + 1.5) / (v26 * v26))) - 1e10)) + (v26 * (0.1 / v26))) / (((v26 - v26) * v26) - (v26 - (v26 - float64(1e10))))) - 1e10)) {
        for fi32 := 0; fi32 < 13; fi32++ {
            v33 := GenS2{F0: 4294967295, F1: -752}
            var v34 float32 = (0.1 * ((0.0 + 1.5) + (3.14159 + -2.25)))
            _ = fi32
            _ = v33
            _ = v34
        }
        var v35 uint32 = 2
        _ = v35
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        switch "hello" {
        case "\n\t":
            var v37 bool = true
            _ = v37
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v31 = (^(v31))
        var v38 uint64 = 1
        _ = fi36
        _ = v38
    }
    switch (v23 * 613) {
    case 0:
        fmt.Println(v23, v31, v23)
    default:
        fmt.Println(v26, v26)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v31
    _ = v4
}
