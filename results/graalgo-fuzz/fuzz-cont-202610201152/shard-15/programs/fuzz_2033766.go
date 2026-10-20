package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint64 }
func genFn0(p0 int16, p1 uint64, p2 string) uint64 {
    v7 := GenS1{F0: p2, F1: 1}
    v8 := []int{-1, -468, 0}
    _ = len(v8)
    var v9 int8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p1
}

func main() {
    v1 := []string{}
    var v2 string = (("" + (((("foo bar" + "foo bar") + "\n\t") + ("hello" + ("hello" + ("\n\t" + "hello")))) + "a")) + ("" + "a"))
    var v3 int = 2147483647
    v4 := make(map[uint32]uint32)
    var v5 float32 = (1e10 * float32(1e10))
    var v6 uint16 = 32767
    v3 = 127
    v3 -= v3
    v2 = v2
    if true {
        if len(v1) > 0 {
            v1[0] = ""
        }
        v2 = v2
        switch "a" {
        case "hello":
            v4[913] = 2147483648
        }
    }
    v5 += v5
    v5 -= (1.5 - 1e10)
    var v10 uint8 = 83
    switch "a" {
    case "foo bar":
        v11 := GenS1{F0: "foo bar", F1: uint8(v6)}
        _ = v11
    case "":
        delete(v4, 2147483647)
    case "a":
        var v12 int = 2147483647
        _ = v12
    }
    if true {
        fmt.Println(v10, v2, v10)
    } else {
        var v13 string = ""
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 uint = 255
            delete(v4, 4294967295)
            fmt.Println(v3)
            _ = fi14
            _ = v15
        }
        v3 += (v3 * (-(v3)))
        _ = v13
    }
    var v16 float64 = 0.1
    _ = len(v1)
    v2 = v2
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := GenS2{F0: -9223372036854775808, F1: (v10 | (v10 + (v10 + (v10 >> ((286) & 63))))), F2: 18446744073709551615}
        _ = fi17
        _ = v18
    }
    delete(v4, 2147483647)
    var v19 uint = 433
    var v20 int16 = 513
    var v21 uint32 = 2147483648
    v6 -= v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
