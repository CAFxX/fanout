package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 string; F1 int8 }
func genFn0() uint {
    var v6 uint16 = 673
    var v7 int64 = 1
    v7 += v7
    v8 := GenS1{F0: 0, F1: 100}
    switch "" {
    case "foo bar":
        var v9 uint = 269
        _ = v9
    case "":
        if (32768 == -1) {
            v6 += 32767
        } else {
            switch ("hello" + "\n\t") {
            case "a":
                var v10 bool = true
                _ = v10
            case "hello":
                v7 -= v7
            case "":
                v11 := []int{-2147483648, 42}
                _ = v11
            default:
                v12 := GenS2{F0: -9223372036854775808}
                _ = v12
            }
            v7 -= (v7 >> ((0) & 63))
        }
    case "\n\t":
        v6 = (v6 ^ (^(v6)))
    default:
        var v13 int32 = -1
        _ = v13
    }
    _ = v6
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 uint16) int {
    fmt.Println(p0)
    var v14 float64 = (-2.25 / float64(float32(1.5)))
    var v15 float64 = float64(1.5)
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int8 = 1
        v14 += v15
        _ = fi16
        _ = v17
    }
    fmt.Println(v15, v15)
    _ = p0
    _ = v14
    _ = v15
    return 585
}

func main() {
    var v1 uint64 = 841
    v2 := make(map[int]bool)
    v3 := []string{"a", "hello"}
    var v4 int32 = 1
    var v5 uint32 = 0
    v18 := v2[genFn1(2)]
    var v19 int8 = 22
    v20 := GenS3{F0: (^(v19)), F1: 2}
    if true {
        v18 = v18
    }
    v21 := GenS2{F0: int(v1)}
    if v18 {
        v5 -= v5
        v19 -= -69
        if v18 {
            var v22 int64 = 9223372036854775807
            switch int(v22) {
            case -9223372036854775808:
                var v23 uint64 = v1
                _ = v23
            }
            _ = v22
        }
    } else {
        v24 := GenS3{F0: (v19 ^ 1), F1: 919}
        v25 := GenS4{F0: (("hello" + ("a" + "foo bar")) + "\n\t"), F1: (v19 + int8(v5))}
        var v26 int32 = v4
        _ = v24
        _ = v25
        _ = v26
    }
    v19 -= (v19 & (v19 ^ v19))
    var v27 uint = 42
    var v28 int32 = (v4 ^ -2147483648)
    fmt.Println(v1, v18)
    var v29 int16 = 648
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
