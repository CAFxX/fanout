package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 string }
func genFn0() int8 {
    var v4 uint = 474
    switch -414 {
    case 32768:
        v4 -= 47
    default:
        v5 := []int{2147483647}
        _ = v5
    }
    _ = v4
    return 100
}

func main() {
    var v1 int64 = -934
    var v2 int8 = -2
    v3 := make(map[int]bool)
    for fi6 := 0; fi6 < 3; fi6++ {
        v2 += int8(v1)
        var v7 uint32 = 1
        v8 := GenS1{F0: 529, F1: 2147483647}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS2{F0: -2147483648}
        switch fi9 {
        case -657:
            var v11 string = ((("" + ("" + (("\n\t" + "") + ""))) + "foo bar") + "")
            _ = v11
        default:
            if true {
                v2 += v2
                v1 -= int64(v2)
                v12 := GenS3{F0: v2}
                _ = v12
            }
        }
        _ = fi9
        _ = v10
    }
    v13 := []int{-2147483648, -9223372036854775808}
    v2 = (v2 % ((v2) | 1))
    fmt.Println(v2, v1)
    if len(v13) > 0 {
        v13[0] = 2147483647
    }
    v14 := GenS2{F0: 1}
    switch -2147483648 {
    case 4294967295:
        v15 := GenS4{F0: 42, F1: (("a" + (("foo bar" + "a") + "")) + "\n\t")}
        _ = v15
    }
    v2 -= -4
    var v16 uint16 = 965
    v1 -= (v1 ^ (v1 % (((v1 % (((v1 % ((v1) | 1))) | 1))) | 1)))
    if len(v13) > 0 {
        v13[0] = 127
    }
    v17 := GenS1{F0: 0, F1: 744}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
