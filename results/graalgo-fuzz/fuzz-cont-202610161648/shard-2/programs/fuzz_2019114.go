package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int8 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint, p1 int) int64 {
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 int16 = 930
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 int8 = -1
            v8 := GenS1{F0: 42, F1: 128, F2: v7}
            var v9 int32 = 42
            fmt.Println(v5, fi6, fi6)
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        switch (^(p1)) {
        case 364:
            p1 -= p1
        }
        _ = fi4
        _ = v5
    }
    v10 := []int{p1, p1, p1, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    var v1 int64 = -1
    var v2 float32 = (1e10 + -2.25)
    var v3 float32 = (0.1 - v2)
    switch 42 {
    case 32768:
        v11 := []int{-2147483648, 9223372036854775807, 935, 9223372036854775807}
        _ = v11
    case 1:
        var v12 bool = (0 < 0)
        _ = v12
    }
    switch 42 {
    case -664:
        var v13 int64 = genFn0(871, int(v1))
        _ = v13
    case 42:
        for fi14 := 0; fi14 < 19; fi14++ {
            if (580 == -833) {
                v1 = -1
            }
            fmt.Println(v3, v2, v3)
            var v15 int = 42
            _ = fi14
            _ = v15
        }
    }
    switch 485 {
    case -497:
        var v16 string = ("foo bar" + "foo bar")
        _ = v16
    case 237:
        fmt.Println(v1, v1)
    default:
        for fi17 := 0; fi17 < 8; fi17++ {
            fmt.Println(fi17, v1)
            switch "hello" {
            case "foo bar":
                v18 := []int{(^(fi17)), fi17}
                _ = v18
            case "":
                v3 -= ((v3 + v2) * float32(0.1))
            default:
                var v19 int32 = -2147483648
                _ = v19
            }
            v2 += v3
            _ = fi17
        }
    }
    v2 = 3.14159
    var v20 int = -9223372036854775808
    switch v20 {
    case 1:
        fmt.Println(v2)
    case 775:
        fmt.Println(v2, v3)
    case -940:
        var v21 int32 = -2147483648
        _ = v21
    default:
        v20 = 0
    }
    fmt.Println(v1)
    v22 := GenS2{F0: false}
    v20 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
}
