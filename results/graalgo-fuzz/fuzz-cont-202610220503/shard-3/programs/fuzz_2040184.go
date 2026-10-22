package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
func genFn0() int8 {
    var v5 int8 = 1
    switch -1 {
    case 700:
        switch "foo bar" {
        case "foo bar":
            v6 := GenS1{F0: v5, F1: 42}
            _ = v6
        }
    default:
        v7 := []int{1}
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: 865}
    if ((3.14159 / ((3.14159 - (3.14159 - 1.5)) + float32(-2.25))) > ((3.14159 * 1.5) * 0.1)) {
        v5 -= (v5 % ((v5) | 1))
    }
    _ = v5
    _ = v8
    return v5
}
func genFn1(p0 int, p1 uint) int16 {
    if (!(true || ((^(p0)) >= p0))) {
        p0 += (p0 | p0)
        fmt.Println(p1, p1)
    }
    v9 := []int{p0, p0, (p0 * 1), (p0 >> (((^(p1))) & 63))}
    var v10 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 1000
}

func main() {
    var v1 uint16 = 32767
    var v2 uint = 998
    var v3 uint = (^(v2))
    var v4 int64 = 238
    switch 420 {
    case 42:
        var v11 uint8 = 150
        _ = v11
    case 179:
        v12 := []int{127}
        _ = v12
    case 1:
        fmt.Println(v4)
    }
    v2 -= (^(v3))
    v13 := []int{32768, 911, -9223372036854775808}
    var v14 bool = false
    v15 := GenS1{F0: genFn0(), F1: uint64(v1)}
    switch ("foo bar" + (("" + "foo bar") + (("hello" + ("hello" + "a")) + "hello"))) {
    case "":
        switch -2147483648 {
        case 32768:
            var v16 uint8 = 241
            _ = v16
        case 128:
            if len(v13) > 0 {
                v13[0] = 308
            }
        case -107:
            var v17 int32 = 111
            _ = v17
        }
    default:
        var v18 int = -9223372036854775808
        _ = v18
    }
    if (-32768 >= -104) {
        if len(v13) > 0 {
            v13[0] = 0
        }
    }
    var v19 int = 19
    if false {
        v20 := GenS1{F0: 100, F1: 575}
        _ = v20
    } else {
        var v21 uint = (v3 << ((v3) & 63))
        var v22 uint8 = 77
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        v2 += v3
        var v24 uint32 = 7
        fmt.Println(v2, v3, v14)
        _ = fi23
        _ = v24
    }
    var v25 int64 = (v4 & v4)
    var v26 int32 = -443
    for fi27 := 0; fi27 < 13; fi27++ {
        var v28 float32 = (1.5 / 3.14159)
        var v29 uint16 = v1
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 uint = v2
    v14 = v14
    var v31 string = (("" + "hello") + "a")
    for fi32 := 0; fi32 < 16; fi32++ {
        fmt.Println(v19, v25, v26)
        _ = fi32
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        fmt.Println(v25, v26, v31)
        switch (fi33 >> ((986) & 63)) {
        case 4294967295:
            v34 := GenS1{F0: -1, F1: 413}
            _ = v34
        default:
            v35 := GenS1{F0: -128, F1: 928}
            _ = v35
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
