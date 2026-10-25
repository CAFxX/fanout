package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 string, p1 uint8, p2 string) int8 {
    var v5 int32 = 1
    v5 = v5
    switch 32767 {
    case 32768:
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 uint8 = p1
            if (!(0 == 2147483647)) {
                v8 := GenS1{F0: -1, F1: 2}
                _ = v8
            }
            _ = fi6
            _ = v7
        }
    case 32767:
        var v9 uint16 = 2
        _ = v9
    default:
        v10 := []int{32768, 32768, 2147483647}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    v1 := []int{932, -1, -169, 1}
    var v2 float32 = 0.0
    v3 := []string{("hello" + "foo bar"), "a"}
    var v4 uint32 = 2147483647
    for fi11 := 0; fi11 < 20; fi11++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v4 = v4
        _ = fi11
    }
    if (true && (0 != -383)) {
        fmt.Println(v4, v4)
    }
    v2 = v2
    var v12 string = ("foo bar" + ("" + "foo bar"))
    switch -526 {
    case 1:
        var v13 bool = (v12 >= v12)
        _ = v13
    case 0:
        var v14 uint64 = 4294967295
        _ = v14
    case -620:
        v15 := GenS2{F0: 610}
        _ = v15
    }
    var v16 int8 = -1
    if ("" > v12) {
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            var v17 int16 = -751
            if true {
                v2 = (-2.25 + ((0.0 + float32(1e10)) + float32(0.1)))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v17
        }
        _ = len(v1)
    } else {
        var v18 bool = (-389 >= int32(v16))
        v18 = false
        v16 -= (v16 + (-(v16)))
        _ = v18
    }
    switch v12 {
    case "":
        fmt.Println(v2)
    }
    v16 -= v16
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    fmt.Println(v16, v4)
    if true {
        var v19 int = 32767
        if false {
            var v20 uint8 = 1
            var v21 float64 = 1e10
            fmt.Println(v4)
            _ = v20
            _ = v21
        }
        _ = v19
    }
    v2 -= v2
    v1 = append(v1, 9223372036854775807)
    var v22 int = 128
    if (v12 != (v12 + v12)) {
        var v23 uint = 1
        _ = v23
    }
    var v24 int = (v22 * v22)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
}
