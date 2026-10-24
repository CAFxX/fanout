package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
func genFn0() int16 {
    switch "foo bar" {
    case "a":
        var v5 uint = 42
        _ = v5
    case "hello":
        var v6 int16 = -1
        _ = v6
    default:
        v7 := []int{-775, 128, 78}
        _ = v7
    }
    v8 := []int{}
    _ = v8
    return -1
}

func main() {
    v1 := []int{-600, 1, 347, -1}
    var v2 int8 = 1
    var v3 string = ("" + "foo bar")
    v4 := make(map[uint32]string)
    if (-278 > -915) {
        for fi9 := 0; fi9 < 19; fi9++ {
            v3 = (v3 + v3)
            var v10 string = "hello"
            _ = fi9
            _ = v10
        }
        v3 = (v3 + v3)
        v11 := GenS1{F0: -737, F1: ""}
        _ = v11
    } else {
        if false {
            switch 609 {
            case 9223372036854775807:
                if false {
                    if true {
                        v12 := GenS1{F0: -2147483648, F1: (v3 + v3)}
                        var v13 float64 = (0.0 / float64(((((0.1 / -2.25) + float32(float64(((0.0 * float32(-2.25)) / 1.0)))) + float32(float64((1.5 / -2.25)))) * (0.0 + 1e10))))
                        v14 := GenS1{F0: 561, F1: v3}
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    var v15 float64 = (3.14159 + float64(float32((3.14159 + 3.14159))))
                    _ = v15
                } else {
                    v2 = v2
                    fmt.Println(v2, v3)
                }
            case -2147483648:
                v3 = "foo bar"
            }
            v2 -= -1
        }
    }
    switch 0 {
    case 2147483647:
        for fi16 := 0; fi16 < 6; fi16++ {
            v4[512] = v3
            var v17 int64 = 0
            _ = fi16
            _ = v17
        }
    case 1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -9223372036854775808:
        fmt.Println(v3, v3)
    default:
        var v18 string = v3
        _ = v18
    }
    v3 = v3
    v4[303] = (("foo bar" + ((v3 + ((v3 + "a") + "hello")) + ("\n\t" + (v3 + v3)))) + "hello")
    _ = len(v1)
    v3 = (v3 + ((v3 + "foo bar") + v3))
    var v19 int32 = -1
    if false {
        if ((int32(v2) <= v19) && (!(-539 >= 369))) {
            if len(v1) > 0 {
                v1[0] = 856
            }
        }
    } else {
        var v20 int16 = genFn0()
        fmt.Println(v20)
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
