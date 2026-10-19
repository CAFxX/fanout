package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
func genFn0() int8 {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 int16 = -1
            var v8 int16 = 1
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 int8 = 95
        _ = fi5
        _ = v9
    }
    var v10 int32 = -965
    v11 := []int{0, 32767, 1, -108}
    switch "" {
    case "foo bar":
        var v12 int32 = (v10 ^ (v10 * v10))
        _ = v12
    case "a":
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    case "hello":
        v11 = append(v11, -122)
    default:
        v10 += 865
    }
    _ = v10
    _ = v11
    return 0
}

func main() {
    var v1 uint32 = 0
    var v2 int64 = 0
    var v3 int = 365
    v4 := []int{v3, (^(v3))}
    var v13 string = "hello"
    v14 := GenS1{F0: -32768, F1: 8}
    fmt.Println(v13)
    var v15 bool = false
    for fi16 := 0; fi16 < 14; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            v13 = (((v13 + v13) + (("hello" + (v13 + v13)) + v13)) + "hello")
            var v18 int8 = genFn0()
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            fmt.Println(v3, v15)
            _ = fi19
        }
        _ = fi16
    }
    v15 = v15
    if len(v4) > 0 {
        v4[0] = (v3 + (v3 << ((4294967295) & 63)))
    }
    v13 = ("a" + ("foo bar" + (v13 + "hello")))
    var v20 uint32 = v1
    v21 := GenS1{F0: 303, F1: 807}
    var v22 float64 = ((1.5 - float64(1.5)) / 1.0)
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 uint64 = 0
        fmt.Println(v24, fi23)
        _ = fi23
        _ = v24
    }
    v1 = v1
    if v15 {
        v20 = (v20 << ((uint(v1)) & 63))
        var v25 uint8 = 65
        for fi26 := 0; fi26 < 12; fi26++ {
            v20 += 816
            _ = fi26
        }
        _ = v25
    } else {
        var v27 uint32 = v1
        v13 = ""
        v3 = 153
        _ = v27
    }
    if v15 {
        var v28 uint8 = 128
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
