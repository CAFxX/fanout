package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 bool; F1 uint16; F2 int }
func genFn0(p0 int32, p1 uint32, p2 uint16) int8 {
    p1 = 337
    p1 += 979
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint = 771
    var v2 int = -929
    v3 := []uint32{}
    var v4 string = "hello"
    var v5 bool = false
    v4 = (v4 + v4)
    fmt.Println(v4, v5, v4)
    v1 = v1
    switch v4 {
    case "foo bar":
        v2 += int(v1)
    case "hello":
        switch v2 {
        case 32767:
            var v6 int = 991
            _ = v6
        case 128:
            for fi7 := 0; fi7 < 8; fi7++ {
                for fi8 := 0; fi8 < 18; fi8++ {
                    v4 = "foo bar"
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    switch ("" + v4) {
                    case "\n\t":
                        fmt.Println(fi8)
                    }
                    v9 := GenS1{F0: 0, F1: true}
                    v10 := GenS2{F0: 42, F1: -9223372036854775808}
                    var v11 int32 = -333
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
                _ = fi7
            }
        }
    case "\n\t":
        var v12 int64 = -9223372036854775808
        _ = v12
    default:
        var v13 int8 = 1
        _ = v13
    }
    v3 = append(v3, 1)
    v2 = (^(v2))
    var v14 int64 = -1
    var v15 string = ""
    var v16 uint32 = 42
    var v17 uint = v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v16)
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 int64 = -9223372036854775808
        if v5 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch "\n\t" {
            case "foo bar":
                var v20 int64 = v14
                _ = v20
            case "\n\t":
                fmt.Println(v16, fi18, v19)
            }
            switch "hello" {
            case "foo bar":
                v19 = (v19 | 1)
            case "":
                v2 -= 2147483647
            case "a":
                v4 = ""
            default:
                v21 := GenS3{F0: v14, F1: 38}
                _ = v21
            }
        } else {
            if v5 {
                if (v5 && (!(float32(0.0) < (1.5 / ((0.1 + float32((1e10 * 1e10))) + float32(0.0)))))) {
                    var v22 uint32 = (v16 >> ((v1) & 63))
                    _ = v22
                } else {
                    var v23 float64 = -2.25
                    if len(v3) > 0 {
                        v3[0] = v16
                    }
                    _ = v23
                }
            }
            switch fi18 {
            case 127:
                v24 := GenS4{F0: v5, F1: 768, F2: -1}
                _ = v24
            case 128:
                v17 -= v1
            default:
                var v25 float64 = (0.0 - 1e10)
                _ = v25
            }
            v15 = v15
        }
        _ = fi18
        _ = v19
    }
    var v26 uint = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if ((0.0 + 1e10) > (-2.25 + (-2.25 / float32(((1e10 * 1.5) - (1e10 - -2.25)))))) {
        v4 = (v4 + v4)
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
