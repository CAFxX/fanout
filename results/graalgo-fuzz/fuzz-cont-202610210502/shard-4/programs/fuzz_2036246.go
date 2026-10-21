package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int16 }
type GenS2 struct { F0 int16; F1 uint; F2 int8 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int8; F1 int64 }
func genFn0() uint64 {
    switch "" {
    case "foo bar":
        var v5 bool = (408 > 758)
        _ = v5
    case "":
        v6 := []int{32768, -2147483648, 38}
        _ = v6
    case "a":
        var v7 uint8 = 127
        _ = v7
    default:
        if false {
            var v8 uint32 = 4294967295
            var v9 int64 = 0
            _ = v8
            _ = v9
        }
    }
    switch -1 {
    case -708:
        var v10 int8 = 42
        _ = v10
    case 279:
        var v11 int64 = 9223372036854775807
        _ = v11
    default:
        var v12 int64 = 42
        _ = v12
    }
    var v13 string = ""
    _ = v13
    return 18446744073709551615
}

func main() {
    var v1 int = 128
    var v2 int8 = -58
    v3 := []string{"foo bar", ("hello" + ((("" + ("a" + "")) + "") + "\n\t"))}
    var v4 uint16 = 255
    v4 -= v4
    if (127 > v2) {
        switch "foo bar" {
        case "\n\t":
            v2 += (^(v2))
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "":
            fmt.Println(v2, v1, v2)
        }
    }
    if false {
        v14 := GenS1{F0: 0, F1: 65535, F2: -32768}
        _ = v14
    } else {
        switch (-(v1)) {
        case 521:
            v4 = v4
        case 32767:
            v1 -= 9223372036854775807
        default:
            var v15 float64 = (((0.1 - (float64(3.14159) * float64(float32(float64(3.14159))))) / (1e10 * (float64(1.5) - 0.0))) * ((1e10 - 1e10) + 1.5))
            _ = v15
        }
        v4 += (v4 + (^(v4)))
    }
    v16 := GenS2{F0: 62, F1: 42, F2: v2}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch v1 {
    case 128:
        var v17 float32 = ((0.0 - 0.1) - (float32(3.14159) * 1.5))
        _ = v17
    case -175:
        v3 = append(v3, ("a" + ""))
    case 127:
        switch (v1 | v1) {
        case 914:
            v2 = (v2 * v2)
        case 128:
            var v18 int = (v1 - v1)
            _ = v18
        }
    default:
        v19 := GenS3{F0: 647, F1: (v2 + -1)}
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS4{F0: v2, F1: 867}
        _ = fi20
        _ = v21
    }
    switch "a" {
    case "a":
        var v22 uint16 = (v4 | (v4 & 65535))
        _ = v22
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = ("foo bar" + ("" + "hello"))
        }
    case "\n\t":
        fmt.Println(v4)
    }
    if true {
        v2 -= -1
        v3 = append(v3, (("" + "") + "a"))
        var v23 int16 = 76
        _ = v23
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        fmt.Println(v4, v2)
        v2 = (-(v2))
        fmt.Println(fi24, fi24, v4)
        _ = fi24
    }
    v4 += (^(v4))
    if false {
        v25 := GenS3{F0: 1, F1: 100}
        _ = v25
    } else {
        var v26 int64 = 376
        v27 := GenS1{F0: 1, F1: 1, F2: 32767}
        _ = v26
        _ = v27
    }
    fmt.Println(v1, v1, v2)
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
