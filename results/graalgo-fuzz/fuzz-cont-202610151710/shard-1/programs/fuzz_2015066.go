package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS2 struct { F0 int32; F1 string }
func genFn0(p0 int32, p1 int8, p2 uint16) bool {
    var v6 uint = 18446744073709551615
    fmt.Println(v6, p2, v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ((!false) && true)
}
func genFn1(p0 int16) uint32 {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int16 = 295
        var v9 uint64 = 42
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 = -505
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(p0, fi10)
        p0 = -320
        var v11 uint64 = 18446744073709551615
        _ = fi10
        _ = v11
    }
    switch ("hello" + (("" + "\n\t") + "foo bar")) {
    case "a":
        var v12 int16 = 42
        _ = v12
    case "":
        p0 = p0
    case "hello":
        if (1e10 != (((3.14159 - 0.1) / -2.25) + 1e10)) {
            v13 := []int{9223372036854775807}
            var v14 uint64 = 2
            var v15 int = -587
            _ = v13
            _ = v14
            _ = v15
        }
    default:
        p0 = p0
    }
    _ = p0
    return 653
}

func main() {
    var v1 int64 = 525
    var v2 uint32 = 803
    v3 := make(map[uint32]string)
    var v4 int8 = 42
    var v5 uint = 255
    v1 += (v1 % (((v1 | 0)) | 1))
    var v16 uint = 460
    var v17 uint32 = v2
    if false {
        if false {
            var v18 uint16 = 257
            v19 := []int{127}
            _ = v18
            _ = v19
        } else {
            fmt.Println(v4)
            switch "\n\t" {
            case "\n\t":
                switch 42 {
                case 9223372036854775807:
                    for fi20 := 0; fi20 < 14; fi20++ {
                        v21 := GenS1{F0: 161, F1: true, F2: 207}
                        v22 := GenS2{F0: 434, F1: "\n\t"}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
            default:
                var v23 float32 = 0.0
                _ = v23
            }
        }
    } else {
        var v24 uint = v16
        _ = v24
    }
    var v25 int32 = -2147483648
    v2 -= (v2 | (^(v17)))
    v26 := []int{int(v17), -2147483648, 142}
    var v27 int16 = -501
    switch 32768 {
    case 2147483647:
        var v28 int16 = (^(v27))
        _ = v28
    }
    v17 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
