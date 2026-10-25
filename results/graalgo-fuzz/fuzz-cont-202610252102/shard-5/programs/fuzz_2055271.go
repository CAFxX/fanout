package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint64; F1 bool; F2 int32 }
func genFn0(p0 int64, p1 uint, p2 int) int8 {
    p2 = -460
    if true {
        v5 := []int{26, p2}
        _ = v5
    }
    v6 := GenS1{F0: 26, F1: -2147483648, F2: (((("foo bar" + "foo bar") + "a") + "") + "a")}
    var v7 uint8 = 0
    var v8 int = (p2 | 9223372036854775807)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    var v4 int64 = 1
    v1 = ((v1 + "foo bar") + "hello")
    v9 := GenS1{F0: 127, F1: 0, F2: v1}
    var v10 float64 = float64((((0.0 / 3.14159) / (1e10 - -2.25)) - (3.14159 * 0.0)))
    v3[42] = 127
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 uint16 = 422
        v13 := []int{-431, 53, 32767}
        v14 := v2[25]
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    v10 += 1e10
    v15 := []int{-200, 283, -2147483648, -136}
    fmt.Println(v4, v1)
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    var v16 float64 = 0.0
    v17 := v2[326]
    var v18 int64 = -9223372036854775808
    v1 = v1
    switch "\n\t" {
    case "a":
        v19 := GenS2{F0: -502}
        _ = v19
    case "\n\t":
        if false {
            v10 -= -2.25
        } else {
            var v20 uint64 = 1
            v1 = ((("a" + ("foo bar" + v1)) + "") + v1)
            var v21 uint8 = 255
            _ = v20
            _ = v21
        }
    default:
        v18 -= v4
    }
    var v22 string = "\n\t"
    var v23 bool = (-2.25 == -2.25)
    for fi24 := 0; fi24 < 18; fi24++ {
        switch (^(v17)) {
        case 42:
            if (v23 || (42 == 65535)) {
                var v25 uint64 = 2
                if (!(!v23)) {
                    fmt.Println(v18)
                    switch v22 {
                    case "foo bar":
                        v22 = v1
                    default:
                        fmt.Println(v1, v23, v18)
                    }
                } else {
                    v22 = v22
                    v26 := GenS3{F0: 1, F1: v25, F2: 246}
                    _ = v26
                }
                _ = v25
            } else {
                v27 := GenS4{F0: 42, F1: v23, F2: 1}
                _ = v27
            }
        default:
            if true {
                for fi28 := 0; fi28 < 11; fi28++ {
                    v1 = v22
                    switch fi24 {
                    case 4294967295:
                        var v29 string = v1
                        _ = v29
                    }
                    _ = fi28
                }
            } else {
                if v23 {
                    var v30 uint = 255
                    v22 = "foo bar"
                    switch (fi24 % ((fi24) | 1)) {
                    case 127:
                        var v31 bool = (127 != (fi24 - 1))
                        _ = v31
                    }
                    _ = v30
                }
                var v32 int32 = 42
                v33 := GenS3{F0: 4294967295, F1: 42, F2: 42}
                _ = v32
                _ = v33
            }
        }
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
