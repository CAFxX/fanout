package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint8, p1 uint16) bool {
    switch ("" + "a") {
    case "hello":
        switch 42 {
        case 42:
            var v7 float32 = 3.14159
            _ = v7
        default:
            p0 -= 81
        }
    case "a":
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 uint16 = p1
            v10 := GenS1{F0: (!true), F1: 522}
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        fmt.Println(p1)
        var v12 bool = false
        _ = fi11
        _ = v12
    }
    p1 = 42
    if (!true) {
        p0 += 127
    }
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 int32 = 1
    var v2 uint16 = 107
    var v3 int32 = v1
    v4 := make(map[string]uint32)
    var v5 uint8 = 127
    var v6 uint64 = 4294967295
    for fi13 := 0; fi13 < 20; fi13++ {
        delete(v4, "hello")
        _ = fi13
    }
    if ((true && (-1 == 0)) && false) {
        switch ("foo bar" + "") {
        case "foo bar":
            var v14 int = 128
            _ = v14
        case "hello":
            if false {
                v2 = v2
            }
        case "":
            for fi15 := 0; fi15 < 16; fi15++ {
                v16 := v4["\n\t"]
                v1 += v1
                _ = fi15
                _ = v16
            }
        default:
            var v17 int8 = 127
            _ = v17
        }
    } else {
        var v18 float32 = 0.0
        _ = v18
    }
    v19 := v4["\n\t"]
    var v20 uint64 = (v6 % ((v6) | 1))
    if false {
        v4["hello"] = v19
        var v21 int8 = -128
        _ = v21
    }
    switch 1 {
    case 291:
        var v22 float64 = (float64(-2.25) - 0.0)
        _ = v22
    }
    switch -618 {
    case 9223372036854775807:
        v20 += v20
    case -1:
        v2 = 32767
    }
    v23 := []int{}
    if false {
        _ = len(v23)
        switch -215 {
        case -2147483648:
            v20 = v20
        }
        v24 := v4["hello"]
        _ = v24
    }
    var v25 int64 = -52
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
