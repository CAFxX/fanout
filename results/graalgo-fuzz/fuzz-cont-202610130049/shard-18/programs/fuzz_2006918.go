package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int64 }
type GenS2 struct { F0 int64; F1 string; F2 uint64 }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 uint64 = 0
    v5 := GenS1{F0: 4294967295, F1: 516, F2: 201}
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int8 = 0
        _ = fi6
        _ = v7
    }
    fmt.Println(v4, v1)
    switch 199 {
    case 4294967295:
        var v8 string = "a"
        _ = v8
    case 2147483647:
        for fi9 := 0; fi9 < 9; fi9++ {
            if (false || (0 == 0)) {
                var v10 int = fi9
                v10 -= v10
                v11 := GenS1{F0: 4294967295, F1: 801, F2: 1}
                _ = v10
                _ = v11
            }
            fmt.Println(v4, v1)
            v12 := GenS2{F0: 42, F1: "", F2: v4}
            _ = fi9
            _ = v12
        }
    case 42:
        switch ((("" + ("\n\t" + "foo bar")) + ("\n\t" + "")) + "foo bar") {
        case "a":
            v13 := GenS2{F0: 165, F1: "hello", F2: (v4 - v4)}
            _ = v13
        case "hello":
            fmt.Println(v1, v1)
        case "":
            v4 -= (^(v4))
        default:
            var v14 string = ((((("" + "\n\t") + "a") + "") + ("\n\t" + "hello")) + "hello")
            _ = v14
        }
    }
    fmt.Println(v4)
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 uint32 = 2147483647
        var v17 uint16 = 2
        v1 += (^(v1))
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        var v18 uint64 = 380
        _ = v18
    } else {
        if ((0 <= 0) || (-32768 == -229)) {
            delete(v2, 32767)
            v19 := []int{908, 2147483647, 1}
            _ = v19
        }
        var v20 float32 = (1.5 + 3.14159)
        _ = v20
    }
    v4 = v4
    v4 += v4
    var v21 float32 = -2.25
    v22 := []int{}
    v23 := GenS2{F0: -250, F1: ("foo bar" + "a"), F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
