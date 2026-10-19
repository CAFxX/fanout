package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int32; F2 uint8 }
func genFn0() uint {
    if false {
        var v6 int16 = -466
        _ = v6
    } else {
        var v7 float64 = -2.25
        switch 4294967295 {
        case 42:
            v8 := GenS1{F0: 255}
            _ = v8
        case -218:
            fmt.Println(v7)
        }
        _ = v7
    }
    var v9 uint16 = 255
    var v10 int = 1
    var v11 uint32 = 2147483648
    _ = v10
    _ = v11
    _ = v9
    return 0
}
func genFn1(p0 string, p1 uint8) int16 {
    fmt.Println(p1)
    fmt.Println(p1)
    var v12 float32 = ((float32(3.14159) + ((float32(((float64((1.5 - (1.5 - -2.25))) * float64(1e10)) * 1e10)) / -2.25) + 0.1)) - 0.0)
    var v13 int = -1
    v12 -= v12
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 32767
}

func main() {
    var v1 bool = false
    v2 := []string{(("a" + "hello") + ("hello" + ("foo bar" + "\n\t"))), "\n\t", ""}
    var v3 uint16 = 42
    v4 := make(map[string]int)
    var v5 int16 = -706
    fmt.Println(v3, v3)
    switch 4294967295 {
    case 127:
        v5 = v5
    default:
        fmt.Println(v5, v3)
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 int16 = v5
        v16 := GenS2{F0: ("" + ""), F1: -2147483648, F2: 30}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := GenS2{F0: ((("\n\t" + "hello") + "hello") + ""), F1: 2147483647, F2: uint8(v3)}
        _ = fi17
        _ = v18
    }
    var v19 int8 = -128
    if (v1 || (0 == 4294967295)) {
        fmt.Println(v3)
        delete(v4, "hello")
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = 0
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 uint8 = 133
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
