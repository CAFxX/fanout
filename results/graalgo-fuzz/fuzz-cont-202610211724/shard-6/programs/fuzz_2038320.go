package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint64; F2 uint32 }
type GenS3 struct { F0 bool }
func genFn0(p0 int64, p1 int32) int8 {
    var v5 string = ""
    v6 := GenS1{F0: (p1 >> ((804) & 63))}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint8 = 1
    var v2 int64 = -434
    var v3 bool = true
    v4 := []int{959, 128, 1, 718}
    var v7 uint16 = 255
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v8 int64 = (v2 ^ 42)
    var v9 uint = 4294967295
    v10 := GenS1{F0: 0}
    if v3 {
        v1 = v1
        var v11 uint = v9
        _ = v11
    } else {
        if true {
            v2 = v2
            v8 = v8
        } else {
            if (!(v7 <= v7)) {
                fmt.Println(v2, v3)
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
            }
            var v12 int16 = -1
            _ = v12
        }
    }
    switch 147 {
    case 32767:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 0:
        fmt.Println(v7, v1)
    case 42:
        v8 -= (v2 % ((v2) | 1))
    }
    var v13 int32 = -2147483648
    var v14 int16 = 42
    if len(v4) > 0 {
        v4[0] = -374
    }
    _ = len(v4)
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 string = ((("" + "\n\t") + (("hello" + "\n\t") + ("a" + "a"))) + (("foo bar" + "\n\t") + "hello"))
        _ = len(v4)
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: "", F1: uint64(v7), F2: 42}
    for fi18 := 0; fi18 < 2; fi18++ {
        switch ("a" + ("" + "a")) {
        case "a":
            var v19 int16 = 1000
            _ = v19
        case "foo bar":
            fmt.Println(v7, v7, v1)
        default:
            v20 := GenS2{F0: "hello", F1: 0, F2: 42}
            _ = v20
        }
        v21 := GenS1{F0: v13}
        _ = fi18
        _ = v21
    }
    if v3 {
        v22 := GenS3{F0: v3}
        var v23 float64 = 0.0
        _ = v22
        _ = v23
    }
    if (true && (19 > 0)) {
        var v24 float64 = 3.14159
        _ = v24
    } else {
        fmt.Println(v7, v1)
        var v25 uint = v9
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
