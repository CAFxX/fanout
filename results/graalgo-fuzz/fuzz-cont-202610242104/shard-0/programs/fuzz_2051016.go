package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int64 = 42
    v2 := []uint32{2147483648, 811, 2147483647, 358}
    var v3 uint32 = 42
    var v4 float32 = 0.1
    switch ("a" + (("foo bar" + "\n\t") + "\n\t")) {
    case "\n\t":
        if (127 >= 32767) {
            var v5 int = 128
            if true {
                v3 = (^(v3))
            }
            _ = v5
        } else {
            for fi6 := 0; fi6 < 16; fi6++ {
                v1 = (v1 % ((v1) | 1))
                fmt.Println(v1, v3)
                _ = fi6
            }
        }
    case "a":
        v7 := GenS1{F0: 255, F1: 100}
        _ = v7
    default:
        var v8 int64 = v1
        _ = v8
    }
    v4 += (1e10 * v4)
    var v9 int = -9223372036854775808
    _ = len(v2)
    v9 = v9
    v2 = append(v2, v3)
    v9 += (v9 * -13)
    v4 += -2.25
    fmt.Println(v9, v1, v1)
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 string = ("hello" + "")
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 20; fi13++ {
            fmt.Println(v3, v9, fi10)
            var v14 int64 = v1
            v4 += v4
            _ = fi13
            _ = v14
        }
        var v15 string = "hello"
        _ = fi10
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 float64 = 1.5
        var v18 string = "\n\t"
        v18 = ("" + (v18 + v18))
        _ = fi16
        _ = v17
        _ = v18
    }
    if false {
        var v19 uint32 = (v3 + 63)
        _ = v19
    } else {
        v3 -= v3
    }
    switch "hello" {
    case "hello":
        v20 := GenS1{F0: 2, F1: 127}
        _ = v20
    case "foo bar":
        switch (-(v9)) {
        case 128:
            v1 = v1
        }
    case "a":
        v3 += v3
    default:
        v21 := GenS1{F0: 2, F1: -128}
        _ = v21
    }
    v22 := GenS2{F0: 42}
    fmt.Println(v1, v4, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
