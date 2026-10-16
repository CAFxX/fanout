package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 string }
type GenS2 struct { F0 uint }
func genFn0(p0 uint) uint32 {
    p0 -= p0
    var v7 uint32 = 4294967295
    fmt.Println(v7)
    v7 -= (v7 ^ (^(v7)))
    v7 = v7
    _ = p0
    _ = v7
    return v7
}
func genFn1() uint {
    switch -458 {
    case -9223372036854775808:
        var v8 bool = false
        _ = v8
    default:
        var v9 uint = 18446744073709551615
        _ = v9
    }
    if true {
        var v10 uint32 = 0
        _ = v10
    }
    var v11 float64 = ((float64(-2.25) * ((1.5 + 0.1) - 1.5)) - float64(1e10))
    var v12 int32 = 42
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 uint64 = 706
    var v2 int64 = -9223372036854775808
    var v3 float32 = 1.5
    var v4 bool = true
    var v5 int8 = 127
    var v6 uint32 = 2147483648
    v13 := GenS1{F0: -1, F1: 42, F2: "hello"}
    v14 := []int{}
    if len(v14) > 0 {
        v14[0] = 572
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    var v15 uint8 = 54
    v3 = ((v3 + ((v3 / ((v3 - float32(0.1)) * v3)) * v3)) / float32(0.1))
    fmt.Println(v4, v2, v2)
    fmt.Println(v4, v1, v2)
    switch "foo bar" {
    case "":
        if len(v14) > 0 {
            v14[0] = 341
        }
    case "\n\t":
        switch "" {
        case "foo bar":
            var v16 uint32 = v6
            _ = v16
        }
    case "a":
        _ = len(v14)
    }
    if v4 {
        for fi17 := 0; fi17 < 18; fi17++ {
            v4 = v4
            v18 := GenS1{F0: (v2 ^ (-(v2))), F1: (v2 - v2), F2: "hello"}
            _ = fi17
            _ = v18
        }
    }
    switch "a" {
    case "foo bar":
        if len(v14) > 0 {
            v14[0] = 2147483647
        }
    case "\n\t":
        var v19 uint8 = v15
        _ = v19
    }
    var v20 int = 1
    for fi21 := 0; fi21 < 6; fi21++ {
        if v4 {
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        }
        fmt.Println(v5, v5)
        _ = fi21
    }
    if ((-2.25 + 0.1) == ((((v3 + v3) * -2.25) + (v3 - v3)) - v3)) {
        if (v3 == 3.14159) {
            v20 -= int(v15)
        } else {
            if len(v14) > 0 {
                v14[0] = v20
            }
            v22 := GenS2{F0: 685}
            _ = v22
        }
        var v23 uint32 = 764
        var v24 string = ("" + "hello")
        _ = v23
        _ = v24
    } else {
        if len(v14) > 0 {
            v14[0] = 31
        }
    }
    v25 := GenS2{F0: 42}
    var v26 float64 = (1e10 - (0.1 * 0.1))
    var v27 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
