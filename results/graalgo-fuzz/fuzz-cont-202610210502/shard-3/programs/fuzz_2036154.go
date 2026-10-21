package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []uint32{}
    var v2 float32 = 1e10
    var v3 float32 = (1.5 - (((v2 * ((v2 * 0.0) / (-2.25 - (v2 + v2)))) + ((v2 * -2.25) + v2)) / v2))
    var v4 int32 = -2147483648
    var v5 uint32 = 4294967295
    fmt.Println(v4)
    var v6 uint = 146
    if false {
        var v7 float32 = ((((((v2 + 1.5) * ((1.5 / -2.25) - (float32(-2.25) * 1.5))) * v3) + 0.0) + v3) * (((v3 - (0.0 / v2)) * v2) + v3))
        _ = v7
    }
    v8 := GenS1{F0: 1}
    if len(v1) > 0 {
        v1[0] = v5
    }
    if false {
        if ((1 <= 42) && (42 > 0)) {
            switch -2147483648 {
            case 738:
                var v9 uint = (v6 * v6)
                _ = v9
            case 42:
                var v10 string = ("foo bar" + "hello")
                _ = v10
            default:
                if (3.14159 >= 1.5) {
                    v1 = append(v1, v5)
                    v6 = (v6 % ((v6) | 1))
                    fmt.Println(v4, v5)
                }
            }
            v2 = v2
        }
        v2 = (v3 + v2)
        fmt.Println(v2)
    } else {
        v2 += v3
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    if ("" >= "foo bar") {
        var v11 int16 = -385
        _ = v11
    } else {
        var v12 int = 32767
        v2 -= v2
        _ = v12
    }
    v2 = (0.1 * 0.0)
    v6 -= v6
    fmt.Println(v3, v2, v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        v13 := GenS2{F0: -333}
        _ = v13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
