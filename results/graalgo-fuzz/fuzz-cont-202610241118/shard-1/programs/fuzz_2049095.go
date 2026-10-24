package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int16 }
func genFn0(p0 int32, p1 uint64) bool {
    fmt.Println(p0, p0)
    v7 := []int{32768, -143, 976}
    v8 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (1.5 > 1e10)
}
func genFn1(p0 int8) string {
    p0 += 0
    if true {
        if (1 <= -1) {
            p0 -= (p0 * p0)
        } else {
            p0 += p0
            var v9 int = 2147483647
            v9 = v9
            _ = v9
        }
        var v10 int64 = -403
        _ = v10
    } else {
        p0 = p0
        switch "hello" {
        case "\n\t":
            var v11 float32 = (3.14159 / 1.5)
            _ = v11
        case "":
            switch "hello" {
            case "foo bar":
                var v12 uint32 = 849
                _ = v12
            case "hello":
                v13 := []int{879, 35, 234}
                _ = v13
            }
        case "hello":
            fmt.Println(p0)
        }
    }
    _ = p0
    return "\n\t"
}

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 190
    var v4 int = 128
    v5 := make(map[string]uint32)
    var v6 bool = true
    var v14 int = (v4 ^ (v4 << ((463) & 63)))
    v5["a"] = 42
    delete(v2, 1)
    switch v14 {
    case 127:
        v1 = (v1 + v1)
    }
    var v15 uint64 = 59
    for fi16 := 0; fi16 < 2; fi16++ {
        switch -318 {
        case -725:
            var v17 int8 = -128
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            v3 = v3
            v19 := GenS1{F0: 42}
            fmt.Println(fi16, v1)
            _ = fi18
            _ = v19
        }
        _ = fi16
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 float32 = (0.0 + -2.25)
        _ = fi20
        _ = v21
    }
    if v6 {
        var v22 int8 = 42
        var v23 uint16 = 2
        var v24 int32 = (v1 + v1)
        _ = v22
        _ = v23
        _ = v24
    } else {
        switch 4294967295 {
        case 32768:
            switch v14 {
            case 32768:
                var v25 uint8 = (v3 >> ((42) & 63))
                _ = v25
            case 9223372036854775807:
                v4 = v14
            default:
                v26 := GenS2{F0: 913, F1: 1}
                _ = v26
            }
        }
        fmt.Println(v6, v1)
        fmt.Println(v15, v15)
    }
    v3 += (v3 & v3)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
