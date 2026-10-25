package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }

func main() {
    var v1 string = ((("hello" + ("a" + "hello")) + "hello") + (((("" + ("\n\t" + "a")) + "\n\t") + (("hello" + ("a" + "a")) + ("\n\t" + "hello"))) + "\n\t"))
    v2 := make(map[string]bool)
    var v3 float64 = float64((((1.5 - 0.0) - 3.14159) - 0.1))
    v4 := []int{480, 32768, -792, 9223372036854775807}
    v1 = ""
    if false {
        if true {
            fmt.Println(v3, v3)
            v3 = float64(3.14159)
            fmt.Println(v3, v3)
        }
        var v5 int8 = 0
        _ = v5
    }
    fmt.Println(v1, v1)
    v1 = ("a" + (("a" + "a") + "a"))
    fmt.Println(v3, v3)
    if (!false) {
        for fi6 := 0; fi6 < 7; fi6++ {
            fmt.Println(fi6)
            v7 := GenS1{F0: -128, F1: 1, F2: (v1 + v1)}
            if len(v4) > 0 {
                v4[0] = (fi6 ^ (fi6 - fi6))
            }
            _ = fi6
            _ = v7
        }
        var v8 int = -34
        v3 = float64(0.1)
        _ = v8
    } else {
        fmt.Println(v3, v1)
    }
    v3 = (1.5 + v3)
    var v9 int8 = 42
    for fi10 := 0; fi10 < 14; fi10++ {
        if true {
            v1 = v1
            fmt.Println(fi10, fi10, fi10)
        } else {
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 uint16 = 380
                _ = fi11
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v14 := GenS2{F0: 42, F1: (v9 + (v9 - (^(v9)))), F2: 9223372036854775807}
            switch fi13 {
            case 0:
                fmt.Println(fi10, v1, v1)
            case -9223372036854775808:
                var v15 uint = 4294967295
                _ = v15
            default:
                v16 := GenS1{F0: v9, F1: 90, F2: v1}
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 int = int(v9)
            _ = fi17
            _ = v18
        }
        _ = fi10
    }
    if (0 != 0) {
        v3 += v3
        var v19 string = v1
        _ = v19
    } else {
        v20 := GenS2{F0: 2147483648, F1: 42, F2: -665}
        v21 := GenS1{F0: -73, F1: 255, F2: v1}
        switch 9223372036854775807 {
        case 128:
            v9 = (^(v9))
        case 4294967295:
            for fi22 := 0; fi22 < 7; fi22++ {
                v9 += (v9 + v9)
                v23 := GenS1{F0: v9, F1: 647, F2: v1}
                _ = fi22
                _ = v23
            }
        }
        _ = v20
        _ = v21
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v24 uint = 0
    var v25 int = -9223372036854775808
    v25 = (v25 ^ v25)
    fmt.Println(v3, v1)
    v26 := GenS2{F0: 492, F1: int8(v25), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v9
}
