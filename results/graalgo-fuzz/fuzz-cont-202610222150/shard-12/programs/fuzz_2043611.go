package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int64; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8; F1 string; F2 uint }
type GenS4 struct { F0 uint8; F1 int64; F2 int8 }
func genFn0(p0 uint64, p1 int64, p2 int16) string {
    p2 += (p2 % ((p2) | 1))
    v7 := []int{4294967295, 2147483647, int(p2), 9223372036854775807}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    switch 617 {
    case 1:
        var v8 uint16 = 1
        _ = v8
    case 32767:
        v9 := GenS1{F0: -1, F1: 127}
        _ = v9
    case -17:
        var v10 float64 = -2.25
        _ = v10
    default:
        var v11 float32 = float32(0.0)
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ((("" + "hello") + "foo bar") + "a")
}

func main() {
    var v1 float32 = 1e10
    var v2 int = 127
    v3 := make(map[string]string)
    var v4 int32 = -2147483648
    v5 := []int{42, v2, v2}
    v6 := make(map[int]bool)
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        for fi12 := 0; fi12 < 6; fi12++ {
            fmt.Println(v2, v1, v1)
            _ = fi12
        }
    case "hello":
        v13 := GenS2{F0: 42, F1: 42, F2: 42}
        _ = v13
    case "a":
        v2 += -2147483648
    }
    switch v2 {
    case -9223372036854775808:
        v2 -= (^(v2))
    case -1:
        v14 := GenS3{F0: 1, F1: "a", F2: 255}
        _ = v14
    default:
        switch "\n\t" {
        case "a":
            v15 := GenS1{F0: 985, F1: (v2 >> ((2) & 63))}
            _ = v15
        }
    }
    var v16 uint64 = 1
    if len(v5) > 0 {
        v5[0] = (v2 * 1)
    }
    v4 = (v4 - v4)
    switch ("\n\t" + "") {
    case "\n\t":
        if (106 != 1) {
            v16 = (v16 ^ (^(v16)))
            delete(v6, -972)
            fmt.Println(v4, v4)
        }
    }
    v17 := GenS3{F0: -20, F1: "foo bar", F2: 255}
    var v18 int16 = 1000
    if false {
        fmt.Println(v16)
        fmt.Println(v18)
        if false {
            var v19 uint = 255
            var v20 string = ""
            _ = v19
            _ = v20
        } else {
            if len(v5) > 0 {
                v5[0] = (v2 | (v2 % ((v2) | 1)))
            }
            fmt.Println(v4, v16, v4)
            if false {
                var v21 int16 = (-(v18))
                var v22 int16 = v18
                var v23 uint32 = 0
                _ = v21
                _ = v22
                _ = v23
            } else {
                for fi24 := 0; fi24 < 12; fi24++ {
                    v25 := GenS4{F0: 0, F1: -1, F2: -49}
                    v1 = (v1 + v1)
                    _ = fi24
                    _ = v25
                }
                v26 := GenS3{F0: 1, F1: ("\n\t" + "\n\t"), F2: 0}
                var v27 int8 = 127
                _ = v26
                _ = v27
            }
        }
    }
    v4 = (v4 & v4)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v28 bool = (false || (505 > 2147483647))
    fmt.Println(v28, v28)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
