package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }

func main() {
    var v1 int64 = 659
    var v2 uint64 = 349
    v3 := []uint32{4294967295, 911}
    v4 := GenS1{F0: 127, F1: (-180 == 970)}
    switch -627 {
    case -58:
        var v5 int8 = -69
        _ = v5
    default:
        var v6 int8 = 127
        _ = v6
    }
    var v7 string = "foo bar"
    v3 = append(v3, 0)
    if (32767 < -1) {
        var v8 uint32 = 2
        _ = len(v3)
        var v9 int8 = -71
        _ = v8
        _ = v9
    }
    var v10 uint16 = 105
    var v11 string = v7
    fmt.Println(v2, v7)
    v12 := GenS1{F0: -47, F1: (true && (-2147483648 == 337))}
    var v13 uint64 = 286
    fmt.Println(v1, v11, v11)
    switch -974 {
    case -1:
        if false {
            fmt.Println(v11, v1)
            var v14 string = (v11 + (("hello" + (("hello" + v7) + ("a" + v7))) + v11))
            _ = v14
        }
    case 32767:
        if (-268 >= 1000) {
            v7 = "foo bar"
            v11 = v11
            v11 = "\n\t"
        }
    case -2147483648:
        var v15 string = ""
        _ = v15
    default:
        if (!false) {
            v7 = v7
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v16 float32 = (-2.25 * 1.5)
            _ = v16
        }
    }
    var v17 float64 = 1.5
    var v18 uint = 65535
    var v19 int64 = (v1 & 9223372036854775807)
    switch 9223372036854775807 {
    case -2147483648:
        v11 = "\n\t"
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 10; fi21++ {
            v7 = v7
            v2 -= (^(v2))
            if (false || (v11 >= "a")) {
                v11 = "\n\t"
                if len(v3) > 0 {
                    v3[0] = 1
                }
                v17 = (v17 + (float64(float32(-2.25)) * ((-2.25 + float64(0.1)) / v17)))
            }
            _ = fi21
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
