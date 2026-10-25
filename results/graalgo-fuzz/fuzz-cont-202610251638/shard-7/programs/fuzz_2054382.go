package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int32 }
func genFn0() string {
    v6 := GenS1{F0: -1}
    var v7 int64 = 42
    v8 := GenS1{F0: (v7 + (v7 >> ((0) & 63)))}
    switch "\n\t" {
    case "\n\t":
        v9 := []int{-9223372036854775808, 573, -38, 2147483647}
        _ = v9
    }
    _ = v6
    _ = v7
    _ = v8
    return ("hello" + ("foo bar" + ("" + (("hello" + "") + (("\n\t" + (("hello" + "") + "")) + "")))))
}

func main() {
    var v1 bool = (((0.1 * -2.25) + -2.25) == ((((((0.0 - (1e10 * 1e10)) * float64(1.5)) - 1e10) - (float64(0.1) - 1.5)) / float64(0.1)) * ((-2.25 * ((-2.25 * (1.5 + (1.5 + -2.25))) + 0.1)) - float64(1e10))))
    var v2 uint64 = 286
    v3 := []uint32{275, 1}
    var v4 int64 = 534
    var v5 int8 = 0
    if v1 {
        if (1000 > -567) {
            var v10 bool = ((v5 & (v5 | -72)) > (v5 >> ((0) & 63)))
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        v1 = (127 != 2)
        if (!v1) {
            _ = len(v3)
            switch "foo bar" {
            case "hello":
                fmt.Println(v5, v5)
            case "a":
                var v12 int32 = 2147483647
                _ = v12
            default:
                v13 := GenS1{F0: v4}
                _ = v13
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v2 = 18446744073709551615
        _ = fi11
    }
    v14 := GenS2{F0: (("hello" + "hello") + "\n\t"), F1: -317}
    v4 -= -695
    v2 += v2
    var v15 uint = 884
    var v16 int16 = 32767
    var v17 uint32 = 0
    var v18 int32 = 0
    var v19 float64 = float64(0.0)
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 uint = 0
        switch ("a" + "hello") {
        case "hello":
            for fi22 := 0; fi22 < 17; fi22++ {
                v4 += v4
                v17 += v17
                _ = fi22
            }
        default:
            v23 := GenS2{F0: ("\n\t" + "a"), F1: v18}
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    var v24 int = int(v2)
    fmt.Println(v19, v1)
    var v25 int64 = int64(v2)
    v4 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
