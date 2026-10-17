package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint64, p1 int16) string {
    var v4 float32 = 1.5
    var v5 string = ("foo bar" + "")
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return v5
}
func genFn1(p0 int, p1 string, p2 string) uint {
    var v6 uint64 = 0
    v7 := []int{(p0 % (((p0 % ((p0) | 1))) | 1)), 128, p0}
    switch p1 {
    case "hello":
        var v8 uint32 = 641
        _ = v8
    case "\n\t":
        var v9 float64 = 1e10
        _ = v9
    default:
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 uint = 503
            fmt.Println(fi10, p0)
            v7 = append(v7, fi10)
            _ = fi10
            _ = v11
        }
    }
    if len(v7) > 0 {
        v7[0] = (p0 | (p0 | 9223372036854775807))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int8 = 1
    var v2 string = ((("a" + ("foo bar" + "\n\t")) + "\n\t") + "foo bar")
    var v3 int64 = -9223372036854775808
    v1 = (v1 / (((v1 % ((1) | 1))) | 1))
    var v12 bool = false
    var v13 int32 = 2147483647
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := GenS1{F0: v2}
        v1 += v1
        _ = fi14
        _ = v15
    }
    var v16 float32 = 0.0
    for fi17 := 0; fi17 < 1; fi17++ {
        switch ("hello" + "\n\t") {
        case "a":
            v18 := []int{(fi17 >> ((2) & 63)), fi17, (fi17 & fi17), (fi17 ^ (^(fi17)))}
            _ = v18
        default:
            fmt.Println(v3, v3)
        }
        var v19 int64 = (v3 + 9223372036854775807)
        for fi20 := 0; fi20 < 16; fi20++ {
            v12 = v12
            _ = fi20
        }
        _ = fi17
        _ = v19
    }
    v1 += 42
    if false {
        v21 := GenS1{F0: ((v2 + (v2 + v2)) + "a")}
        _ = v21
    }
    var v22 uint64 = 42
    v3 += (v3 ^ (v3 - v3))
    var v23 uint16 = 42
    var v24 int = 4294967295
    var v25 uint8 = 127
    v26 := []int{(v24 | 128), v24, (v24 + v24)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
