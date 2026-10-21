package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint16 }
func genFn0() uint32 {
    var v5 uint8 = 63
    fmt.Println(v5)
    var v6 int = 268
    switch "\n\t" {
    case "\n\t":
        for fi7 := 0; fi7 < 3; fi7++ {
            fmt.Println(v6, v6)
            v8 := []int{v6}
            _ = fi7
            _ = v8
        }
    case "foo bar":
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 13; fi10++ {
                v6 += fi9
                _ = fi10
            }
            _ = fi9
        }
    case "":
        for fi11 := 0; fi11 < 15; fi11++ {
            v5 = (v5 / (((v5 | v5)) | 1))
            v6 -= fi11
            _ = fi11
        }
    default:
        var v12 uint64 = 4294967295
        _ = v12
    }
    _ = v5
    _ = v6
    return 2147483648
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = true
    var v3 float32 = (((1.5 * (0.0 * 3.14159)) / 3.14159) - -2.25)
    v4 := []string{("a" + ("a" + "foo bar")), ""}
    var v13 uint32 = genFn0()
    v14 := GenS1{F0: 838, F1: 128}
    var v15 uint = 845
    var v16 float64 = -2.25
    delete(v1, -2147483648)
    if true {
        if v2 {
            v17 := GenS2{F0: genFn0()}
            v18 := GenS3{F0: 100, F1: genFn0(), F2: 42}
            for fi19 := 0; fi19 < 8; fi19++ {
                v16 -= ((-2.25 - v16) / v16)
                _ = fi19
            }
            _ = v17
            _ = v18
        }
    }
    var v20 uint16 = 2
    var v21 uint64 = 2
    switch 1 {
    case 42:
        var v22 uint = 18446744073709551615
        _ = v22
    case -981:
        v21 = (^(v21))
    case 32768:
        v16 -= v16
    default:
        var v23 int = 42
        _ = v23
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        fmt.Println(v16, v21, v21)
        if false {
            if (((v13 << ((0) & 63)) > v13) || false) {
                v25 := GenS1{F0: 4294967295, F1: 4294967295}
                v26 := GenS3{F0: -30, F1: 248, F2: v20}
                _ = v25
                _ = v26
            } else {
                var v27 int16 = 42
                v28 := GenS1{F0: (v21 ^ v21), F1: 9223372036854775807}
                _ = v27
                _ = v28
            }
        }
        _ = fi24
    }
    v2 = (!(!false))
    var v29 int64 = 42
    v21 = (v21 ^ 2)
    var v30 uint32 = 1
    v31 := GenS2{F0: 4294967295}
    fmt.Println(v20, v16, v2)
    if false {
        v32 := GenS2{F0: (v13 - uint32(v15))}
        v21 -= v21
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
