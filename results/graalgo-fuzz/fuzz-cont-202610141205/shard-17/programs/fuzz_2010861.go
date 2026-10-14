package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint; F1 int32; F2 int }
type GenS3 struct { F0 int64; F1 string; F2 int64 }
func genFn0(p0 string, p1 int) int8 {
    switch p1 {
    case 128:
        v5 := GenS1{F0: "foo bar", F1: 1}
        _ = v5
    case -9223372036854775808:
        p0 = ((p0 + p0) + "a")
    case 2147483647:
        p1 -= p1
    }
    if true {
        var v6 int32 = 305
        if false {
            for fi7 := 0; fi7 < 19; fi7++ {
                var v8 uint32 = 2
                fmt.Println(p1)
                _ = fi7
                _ = v8
            }
        } else {
            v9 := []int{}
            if (1 > -1) {
                p1 = (p1 + (p1 / ((p1) | 1)))
                if (0 > 42) {
                    fmt.Println(v6)
                }
            }
            var v10 int8 = 26
            _ = v10
            _ = v9
        }
        _ = v6
    }
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int = -2147483648
    var v2 float32 = 1e10
    v3 := []int{(v1 | (v1 | (-(v1)))), (v1 / (((v1 + v1)) | 1)), v1}
    v4 := []int{-2147483648}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = 3.14159
    switch "" {
    case "hello":
        var v11 string = ("" + ("" + ("" + ("foo bar" + ""))))
        _ = v11
    case "a":
        fmt.Println(v1, v1)
    case "foo bar":
        fmt.Println(v2, v2)
    }
    fmt.Println(v2)
    v2 += float32(0.0)
    if len(v4) > 0 {
        v4[0] = 42
    }
    if true {
        for fi12 := 0; fi12 < 14; fi12++ {
            v2 += v2
            var v13 int16 = -32768
            fmt.Println(v2, v13, v2)
            _ = fi12
            _ = v13
        }
    } else {
        if false {
            v2 += v2
            fmt.Println(v2, v1)
        } else {
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 float64 = ((-2.25 - 3.14159) / (1.5 - float64(0.0)))
                v16 := GenS2{F0: 247, F1: -2147483648, F2: v1}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    }
    v17 := GenS2{F0: 150, F1: 2147483647, F2: v1}
    if true {
        v18 := GenS3{F0: 1, F1: ("" + "\n\t"), F2: 9223372036854775807}
        for fi19 := 0; fi19 < 14; fi19++ {
            v4 = append(v4, (^(v1)))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi19
        }
        _ = v18
    }
    _ = len(v3)
    for fi20 := 0; fi20 < 16; fi20++ {
        v2 = v2
        var v21 int8 = 1
        _ = fi20
        _ = v21
    }
    fmt.Println(v2, v1)
    v22 := GenS3{F0: -366, F1: ("a" + "hello"), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
