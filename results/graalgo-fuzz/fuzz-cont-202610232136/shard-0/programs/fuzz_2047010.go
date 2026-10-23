package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int16; F1 int32; F2 int16 }
type GenS3 struct { F0 int64; F1 int; F2 uint16 }
type GenS4 struct { F0 uint16; F1 int; F2 int64 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint64) string {
    p0 = p0
    if false {
        var v5 int = 42
        var v6 string = "a"
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(p0, p0)
            fmt.Println(p0, p0)
            _ = fi7
        }
        p0 -= (p0 + 18446744073709551615)
    }
    fmt.Println(p0)
    switch "a" {
    case "a":
        switch "hello" {
        case "foo bar":
            var v8 int32 = 42
            _ = v8
        case "a":
            v9 := GenS1{F0: 930, F1: 116}
            _ = v9
        default:
            fmt.Println(p0)
        }
    case "foo bar":
        v10 := []int{-558, 2147483647}
        _ = v10
    default:
        v11 := GenS2{F0: 32767, F1: 42, F2: 0}
        _ = v11
    }
    switch "a" {
    case "hello":
        v12 := GenS3{F0: -156, F1: 127, F2: 2}
        _ = v12
    }
    _ = p0
    return "a"
}
func genFn1(p0 string, p1 uint) int64 {
    p1 -= p1
    v13 := []int{606, 32}
    var v14 int8 = 0
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 1
}

func main() {
    var v1 int32 = 42
    v2 := make(map[string]int)
    v3 := []string{"a", ""}
    var v4 uint8 = 2
    var v15 uint32 = 65535
    v4 = (^(v4))
    if false {
        v2["hello"] = -9223372036854775808
        var v16 float64 = 1e10
        v4 = (v4 & (v4 ^ v4))
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 16; fi17++ {
            switch 0 {
            case 127:
                switch -692 {
                case -861:
                    v1 = (v1 % ((-563) | 1))
                case -282:
                    var v18 uint = uint(v15)
                    _ = v18
                case 42:
                    if true {
                        v19 := GenS3{F0: 525, F1: -2147483648, F2: 42}
                        var v20 uint32 = v15
                        _ = v19
                        _ = v20
                    }
                }
            case 32767:
                v21 := GenS4{F0: 1, F1: fi17, F2: -1}
                _ = v21
            case 9223372036854775807:
                v4 -= 127
            }
            fmt.Println(fi17)
            v2[""] = (fi17 & 1)
            _ = fi17
        }
        var v22 uint64 = 2
        _ = v22
    }
    var v23 int8 = -128
    v24 := GenS5{F0: 2}
    v15 = (v15 - (v15 << ((255) & 63)))
    var v25 float64 = (1e10 * 1e10)
    v2["hello"] = -2147483648
    var v26 uint8 = v4
    v1 = int32(v26)
    var v27 uint = 877
    var v28 int = 32767
    var v29 float64 = (v25 / v25)
    if false {
        v30 := GenS1{F0: int64(v23), F1: 32767}
        var v31 int32 = -1
        for fi32 := 0; fi32 < 12; fi32++ {
            v1 -= 150
            _ = fi32
        }
        _ = v30
        _ = v31
    } else {
        var v33 bool = false
        switch (v28 >> ((1) & 63)) {
        case 127:
            var v34 uint64 = 0
            _ = v34
        case 128:
            fmt.Println(v26, v15, v23)
        case 330:
            if (4294967295 > 1) {
                if ((!v33) && false) {
                    v33 = (false || v33)
                } else {
                    v3 = append(v3, ("hello" + ("foo bar" + "hello")))
                    v29 = 1.5
                }
                v27 = v27
            }
        }
        _ = v33
    }
    v35 := GenS2{F0: 27, F1: v1, F2: 752}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v4
}
