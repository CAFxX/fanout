package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint8 }
func genFn0(p0 uint, p1 uint32) uint64 {
    var v4 int16 = 0
    for fi5 := 0; fi5 < 11; fi5++ {
        if false {
            if true {
                p0 = (p0 * (p0 ^ (p0 << ((p0) & 63))))
            }
        } else {
            var v6 int64 = -1
            _ = v6
        }
        _ = fi5
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 float64 = float64(0.0)
        var v9 uint = p0
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = -2.25
    v3 := make(map[uint32]bool)
    v2 = 0.1
    var v10 uint16 = 30
    var v11 float32 = (float32((-2.25 * (v2 - 3.14159))) * 1.5)
    if false {
        v12 := GenS1{F0: 199, F1: "foo bar", F2: 0}
        var v13 uint8 = 153
        _ = v12
        _ = v13
    } else {
        switch ("\n\t" + "a") {
        case "hello":
            v11 += (v11 * 1.5)
        case "":
            if true {
                v10 += (v10 / ((v10) | 1))
            }
        }
    }
    v14 := v3[563]
    fmt.Println(v2, v2)
    switch 1 {
    case -9223372036854775808:
        v10 += (v10 / ((v10) | 1))
    case -2147483648:
        v15 := GenS1{F0: 0, F1: ((("hello" + ("foo bar" + "\n\t")) + "hello") + ("foo bar" + ("" + "a"))), F2: 42}
        _ = v15
    case 0:
        var v16 float64 = 3.14159
        _ = v16
    default:
        var v17 uint32 = 2147483647
        _ = v17
    }
    var v18 int = -418
    var v19 uint8 = 100
    for fi20 := 0; fi20 < 14; fi20++ {
        v14 = (!false)
        switch v18 {
        case -9223372036854775808:
            var v21 uint16 = v10
            _ = v21
        case -458:
            var v22 int = v18
            _ = v22
        default:
            var v23 float64 = (v2 / ((-2.25 / v2) + (v2 * 0.1)))
            _ = v23
        }
        var v24 uint32 = 669
        _ = fi20
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
