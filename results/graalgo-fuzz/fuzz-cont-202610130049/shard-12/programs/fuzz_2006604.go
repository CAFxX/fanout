package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int8; F1 int8 }

func main() {
    var v1 uint32 = 1
    v2 := []int{751, 127}
    var v3 uint = 42
    var v4 bool = false
    v2 = append(v2, -9223372036854775808)
    for fi5 := 0; fi5 < 8; fi5++ {
        if false {
            v3 += 4294967295
            _ = len(v2)
            if v4 {
                var v6 int = (fi5 + -787)
                v7 := GenS1{F0: 412}
                v4 = (v4 && ((v6 | v6) < -305))
                _ = v6
                _ = v7
            }
        } else {
            v2 = append(v2, (fi5 << ((v3) & 63)))
            if (fi5 > fi5) {
                fmt.Println(v1)
                v1 -= 2147483648
            }
        }
        v1 = v1
        fmt.Println(v1)
        _ = fi5
    }
    var v8 int32 = 621
    v3 = v3
    v4 = (-9223372036854775808 > 9223372036854775807)
    var v9 uint16 = 1
    v10 := GenS2{F0: v1, F1: 0, F2: 307}
    if len(v2) > 0 {
        v2[0] = -831
    }
    var v11 float32 = (3.14159 * ((((1e10 * -2.25) / 1.5) - ((0.1 - 0.0) / 1e10)) + 1.5))
    var v12 uint = (v3 ^ 4294967295)
    if v4 {
        var v13 float32 = float32(-2.25)
        fmt.Println(v13)
        _ = v13
    }
    v14 := GenS3{F0: (v9 < (v9 / ((v9) | 1))), F1: uint32(v9)}
    if (!(!(!v4))) {
        if v4 {
            for fi15 := 0; fi15 < 2; fi15++ {
                switch "a" {
                case "":
                    var v16 int64 = -9223372036854775808
                    _ = v16
                case "foo bar":
                    v4 = true
                }
                var v17 float64 = float64(0.0)
                if len(v2) > 0 {
                    v2[0] = (^(fi15))
                }
                _ = fi15
                _ = v17
            }
            v18 := GenS2{F0: v1, F1: 4294967295, F2: 754}
            _ = v18
        } else {
            var v19 uint8 = 248
            v3 += v3
            _ = v19
        }
        v20 := GenS1{F0: uint64(v1)}
        switch ("\n\t" + "foo bar") {
        case "a":
            var v21 uint64 = 4294967295
            _ = v21
        case "foo bar":
            for fi22 := 0; fi22 < 16; fi22++ {
                fmt.Println(v4, v9)
                _ = fi22
            }
        }
        _ = v20
    } else {
        switch "a" {
        case "hello":
            for fi23 := 0; fi23 < 10; fi23++ {
                fmt.Println(v3, v8, fi23)
                v8 = v8
                for fi24 := 0; fi24 < 15; fi24++ {
                    v3 = v12
                    _ = fi24
                }
                _ = fi23
            }
        }
        switch -209 {
        case 32768:
            var v25 int32 = 230
            _ = v25
        case 9223372036854775807:
            if true {
                var v26 int16 = 977
                v9 += (v9 & 2)
                var v27 uint8 = 42
                _ = v26
                _ = v27
            }
        case 344:
            v28 := GenS3{F0: false, F1: v1}
            _ = v28
        default:
            var v29 string = ("" + "foo bar")
            _ = v29
        }
        v30 := GenS4{F0: 72, F1: -66}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
