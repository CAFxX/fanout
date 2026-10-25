package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 uint64; F2 uint8 }

func main() {
    var v1 float64 = 3.14159
    var v2 uint = 255
    v3 := []int{440, 2147483647, -685}
    var v4 uint32 = 2147483647
    var v5 bool = true
    v6 := make(map[string]string)
    switch "\n\t" {
    case "hello":
        v7 := GenS1{F0: "foo bar", F1: 0}
        _ = v7
    default:
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    }
    var v8 bool = (("foo bar" + "") <= "a")
    var v9 int16 = -100
    v9 = v9
    v10 := v6[("" + ((("foo bar" + "a") + "hello") + "hello"))]
    if ((!(65535 != 65535)) && true) {
        var v11 uint32 = 2147483648
        var v12 uint8 = 1
        _ = v11
        _ = v12
    } else {
        if (1 <= 0) {
            var v13 int8 = -1
            v9 += (v9 + (v9 >> (((v2 % ((v2) | 1))) & 63)))
            v14 := GenS2{F0: 4294967295, F1: 66}
            _ = v13
            _ = v14
        }
    }
    var v15 int64 = 1
    var v16 float32 = (-2.25 / 0.1)
    var v17 uint32 = (^(v4))
    v18 := GenS2{F0: v2, F1: 2}
    fmt.Println(v10, v15)
    var v19 uint8 = 46
    v20 := GenS3{F0: (v19 * v19), F1: (v19 & (v19 * (v19 / (((v19 & (v19 << ((v2) & 63)))) | 1))))}
    v21 := GenS4{F0: 775}
    if (true && v5) {
        if true {
            if v5 {
                if v5 {
                    fmt.Println(v10, v19, v4)
                    v22 := GenS4{F0: 672}
                    _ = v22
                }
                var v23 uint16 = 2
                var v24 int32 = -655
                _ = v23
                _ = v24
            }
            v17 += (v17 + (v4 >> ((v2) & 63)))
        }
        if v5 {
            delete(v6, v10)
        } else {
            fmt.Println(v19, v9)
            var v25 uint32 = v17
            _ = v25
        }
        var v26 uint32 = v4
        _ = v26
    } else {
        v9 = -908
        if (32767 >= 821) {
            fmt.Println(v15)
        } else {
            fmt.Println(v2, v8)
            v15 -= (v15 - (-(v15)))
            var v27 string = v10
            _ = v27
        }
    }
    if (v15 < (v15 + (v15 ^ -771))) {
        var v28 float32 = v16
        var v29 uint32 = v4
        _ = v28
        _ = v29
    } else {
        var v30 int8 = -99
        switch -2147483648 {
        case 9223372036854775807:
            v31 := GenS5{F0: -2147483648, F1: 0, F2: v19}
            _ = v31
        case -41:
            var v32 int8 = (v30 - v30)
            _ = v32
        case 644:
            v15 += (v15 - v15)
        default:
            v10 = (v10 + ((v10 + (v10 + v10)) + v10))
        }
        for fi33 := 0; fi33 < 11; fi33++ {
            v34 := v6[(v10 + v10)]
            var v35 uint64 = 4294967295
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = v30
    }
    switch "hello" {
    case "hello":
        if ((v5 && (0 > 65535)) && v5) {
            v19 = 1
        }
    case "foo bar":
        switch "foo bar" {
        case "a":
            switch v10 {
            case "hello":
                var v36 uint = v2
                _ = v36
            case "foo bar":
                v37 := GenS5{F0: -1, F1: 0, F2: (v19 + v19)}
                _ = v37
            case "":
                fmt.Println(v16, v8, v5)
            }
        default:
            v6[v10] = (v10 + ((v10 + "foo bar") + "hello"))
        }
    case "":
        fmt.Println(v17, v2, v10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
