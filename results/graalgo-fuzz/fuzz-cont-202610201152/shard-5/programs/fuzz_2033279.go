package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
func genFn0() uint64 {
    var v6 int8 = -1
    var v7 float64 = 1e10
    var v8 uint32 = 2147483647
    _ = v6
    _ = v7
    _ = v8
    return 478
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 2
    var v3 uint16 = 1
    var v4 int32 = 2147483647
    var v5 uint32 = 895
    v9 := []int{128}
    var v10 uint = 255
    fmt.Println(v2)
    var v11 int32 = v4
    v10 = (v10 >> (((v10 + v2)) & 63))
    fmt.Println(v5, v5)
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := v1[v5]
        _ = fi12
        _ = v13
    }
    fmt.Println(v4, v5)
    v9 = append(v9, 1)
    v2 += 775
    var v14 int64 = -9223372036854775808
    if (0.1 == 3.14159) {
        v2 -= 18446744073709551615
        for fi15 := 0; fi15 < 8; fi15++ {
            delete(v1, (^(v5)))
            _ = fi15
        }
    } else {
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 uint8 = 42
            _ = fi16
            _ = v17
        }
        fmt.Println(v14, v3)
        if (1000 < 1000) {
            v18 := GenS1{F0: (!(-128 < 78)), F1: 1000}
            var v19 int = 32767
            _ = v18
            _ = v19
        }
    }
    v20 := GenS1{F0: ("hello" <= "hello"), F1: -975}
    v21 := GenS2{F0: true}
    if true {
        switch 1 {
        case 2147483647:
            delete(v1, 65535)
        default:
            v2 = (^(v10))
        }
    } else {
        fmt.Println(v11)
        var v22 uint64 = 4294967295
        _ = v22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        switch "hello" {
        case "hello":
            var v24 int64 = 1
            _ = v24
        case "":
            var v25 float64 = ((1.5 / 1.0) * float64(((1e10 / (0.1 / 1e10)) - ((1e10 + (0.0 - (1e10 - (0.1 + 1e10)))) + (1.5 / 0.1)))))
            _ = v25
        default:
            var v26 uint = (v2 | 65535)
            _ = v26
        }
        var v27 bool = ((v3 & (v3 << ((1) & 63))) == (v3 * v3))
        _ = fi23
        _ = v27
    }
    if false {
        var v28 uint = v10
        v29 := GenS1{F0: ((42 <= 42) || false), F1: -1}
        if (((v4 >> ((42) & 63)) >= (-(v4))) || true) {
            if (true || true) {
                _ = len(v9)
            } else {
                v30 := GenS2{F0: (!((!true) && true))}
                if true {
                    fmt.Println(v3, v10, v5)
                }
                _ = v30
            }
            var v31 float64 = ((3.14159 / (0.1 / float64(-2.25))) + -2.25)
            _ = v31
        } else {
            v32 := GenS2{F0: false}
            _ = v32
        }
        _ = v28
        _ = v29
    }
    switch "foo bar" {
    case "foo bar":
        var v33 uint32 = (^(v5))
        _ = v33
    case "\n\t":
        fmt.Println(v10, v5)
    case "a":
        if true {
            var v34 int16 = 741
            switch "a" {
            case "":
                v35 := GenS3{F0: 1}
                _ = v35
            case "a":
                fmt.Println(v5, v3)
            default:
                v2 = v2
            }
            _ = v34
        } else {
            var v36 uint16 = (v3 & 32767)
            if ((v3 * v36) > 0) {
                var v37 int8 = -90
                _ = len(v9)
                _ = v37
            }
            var v38 int16 = 0
            _ = v36
            _ = v38
        }
    default:
        v14 -= (v14 >> (((v10 / ((v10) | 1))) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
