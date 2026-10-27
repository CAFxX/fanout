package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
func genFn0(p0 int, p1 uint) bool {
    var v7 uint8 = 158
    var v8 int64 = -1
    switch p0 {
    case 32767:
        if false {
            v8 = v8
            v8 += (v8 >> ((p1) & 63))
        }
    case 128:
        var v9 string = "a"
        _ = v9
    default:
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 int8 = 0
                v7 -= v7
                v13 := []int{-1, 1, (fi10 << ((p1) & 63)), (p0 + (p0 + fi11))}
                fmt.Println(p1)
                fmt.Println(v12, p0, v8)
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 uint16 = 2
            _ = fi10
            _ = v14
        }
    }
    fmt.Println(v7, v8, v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return true
}

func main() {
    var v1 bool = true
    v2 := make(map[string]uint32)
    v3 := []string{"", ((("\n\t" + "\n\t") + "foo bar") + ("hello" + "hello")), ((("hello" + "") + "\n\t") + "")}
    var v4 int64 = -1
    var v5 float32 = -2.25
    v6 := make(map[int]uint32)
    var v15 int8 = 85
    v16 := v2["\n\t"]
    if v1 {
        var v17 uint = 806
        var v18 bool = true
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 255, F1: 18446744073709551615}
    var v20 bool = (!v1)
    v21 := GenS1{F0: 1, F1: 4294967295}
    if false {
        switch 9223372036854775807 {
        case 9223372036854775807:
            v16 = (v16 ^ (v16 * 2147483648))
        case -9223372036854775808:
            if len(v3) > 0 {
                v3[0] = ("foo bar" + "a")
            }
        }
        var v22 bool = (819 == 2)
        var v23 uint64 = 2
        _ = v22
        _ = v23
    } else {
        if ((((0.0 - (v5 * v5)) - v5) + ((v5 * v5) / (v5 * ((0.0 / ((v5 * v5) * v5)) * v5)))) == v5) {
            v20 = v1
        } else {
            v5 -= v5
            v24 := GenS1{F0: 127, F1: 1}
            _ = v24
        }
    }
    var v25 string = ((("foo bar" + ("a" + ("hello" + ("foo bar" + "foo bar")))) + "") + ("\n\t" + "\n\t"))
    var v26 uint32 = v16
    var v27 int = 0
    v28 := GenS1{F0: 127, F1: 18446744073709551615}
    var v29 uint = 42
    v5 += ((v5 / v5) / ((((v5 + v5) - (v5 * 0.0)) / 1e10) / (v5 - v5)))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
