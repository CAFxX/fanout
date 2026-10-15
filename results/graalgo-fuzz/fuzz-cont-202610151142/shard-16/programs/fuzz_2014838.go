package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int8; F1 string; F2 int16 }
type GenS3 struct { F0 int32; F1 uint16; F2 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64; F1 int8 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 6; fi4++ {
        v5 := []int{fi4, fi4, fi4, -287}
        var v6 float32 = -2.25
        v7 := GenS1{F0: 672, F1: 0}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 bool = (252 != 42)
    _ = v8
    return 127
}

func main() {
    var v1 uint32 = 803
    v2 := []string{("hello" + ("" + "")), "", ("a" + (("foo bar" + ("\n\t" + "foo bar")) + ("\n\t" + ("" + (("foo bar" + "\n\t") + ""))))), ""}
    var v3 bool = (539 > 2147483647)
    switch 32768 {
    case 2147483647:
        _ = len(v2)
    default:
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(v3, v3)
            v3 = v3
            _ = fi9
        }
    }
    var v10 float32 = (3.14159 + (-2.25 * 1.5))
    var v11 int32 = 0
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS2{F0: 94, F1: "a", F2: -1}
        v14 := GenS1{F0: 255, F1: 93}
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS1{F0: 2, F1: 2}
            _ = fi15
            _ = v16
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v17 uint32 = (v1 & (v1 / ((v1) | 1)))
    if len(v2) > 0 {
        v2[0] = ("" + "hello")
    }
    switch "\n\t" {
    case "":
        fmt.Println(v17, v11, v11)
    case "hello":
        v18 := GenS2{F0: -1, F1: "hello", F2: 550}
        _ = v18
    default:
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    }
    v19 := GenS1{F0: 938, F1: 41}
    if v3 {
        var v20 int = 0
        var v21 uint32 = (v1 / ((2147483647) | 1))
        switch -9223372036854775808 {
        case -2147483648:
            var v22 uint8 = 1
            _ = v22
        case 274:
            var v23 float64 = 1e10
            _ = v23
        case 4294967295:
            if true {
                _ = len(v2)
                v2 = append(v2, ("hello" + "a"))
                v17 = (v17 - 2147483647)
            } else {
                v24 := GenS3{F0: (-(v11)), F1: 32767, F2: 1}
                _ = v24
            }
        default:
            v20 -= (^(v20))
        }
        _ = v20
        _ = v21
    } else {
        v25 := GenS4{F0: (false || false)}
        v26 := GenS5{F0: 65535}
        _ = v25
        _ = v26
    }
    var v27 uint32 = (^(v1))
    v28 := GenS3{F0: (v11 >> ((234) & 63)), F1: 65535, F2: 42}
    v29 := GenS6{F0: 42, F1: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
