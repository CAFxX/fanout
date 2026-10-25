package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 string }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 string }

func main() {
    var v1 int32 = -771
    var v2 int8 = 84
    var v3 int8 = v2
    v4 := make(map[uint32]bool)
    if false {
        var v5 uint16 = 42
        v6 := GenS1{F0: 1, F1: -128, F2: "foo bar"}
        v5 = v5
        _ = v5
        _ = v6
    } else {
        var v7 float64 = float64(0.1)
        _ = v7
    }
    v8 := []int{}
    for fi9 := 0; fi9 < 13; fi9++ {
        if true {
            fmt.Println(v3)
            v10 := GenS1{F0: -1, F1: 42, F2: "foo bar"}
            delete(v4, 187)
            _ = v10
        }
        if (1 > 36) {
            var v11 string = "a"
            _ = v11
        } else {
            switch -2147483648 {
            case 777:
                var v12 int16 = -32768
                _ = v12
            }
        }
        _ = fi9
    }
    if len(v8) > 0 {
        v8[0] = 32767
    }
    switch (("foo bar" + "hello") + ("a" + "hello")) {
    case "a":
        v2 = v3
    case "hello":
        if (!(42 >= 498)) {
            var v13 int = 32767
            v1 = v1
            v13 -= (^(v13))
            _ = v13
        } else {
            var v14 uint64 = 4294967295
            v15 := GenS2{F0: (v14 & (v14 << ((uint(v14)) & 63))), F1: uint32(v14)}
            _ = v14
            _ = v15
        }
    case "foo bar":
        v16 := GenS3{F0: ""}
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        v2 += 1
        v1 = (-(v1))
        _ = fi17
    }
    v2 -= 100
    var v18 float64 = (-2.25 + (3.14159 * 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
