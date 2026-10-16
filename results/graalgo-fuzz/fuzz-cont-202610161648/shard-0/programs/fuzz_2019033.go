package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
func genFn0() string {
    var v7 int64 = -9223372036854775808
    v8 := []int{}
    _ = v7
    _ = v8
    return "a"
}

func main() {
    var v1 uint32 = 956
    v2 := make(map[uint32]bool)
    v3 := make(map[string]uint32)
    var v4 uint = 513
    var v5 string = ("a" + "\n\t")
    var v6 bool = true
    for fi9 := 0; fi9 < 5; fi9++ {
        if (v6 || ((false && (756 <= 871)) && false)) {
            if (-1 <= 42) {
                for fi10 := 0; fi10 < 14; fi10++ {
                    v11 := GenS1{F0: 784, F1: ((v5 + v5) + v5)}
                    v6 = (v5 == ((("foo bar" + "hello") + v5) + v5))
                    _ = fi10
                    _ = v11
                }
                _ = len(v3)
            }
            v5 = v5
            var v12 int64 = 0
            _ = v12
        }
        v13 := GenS1{F0: 971, F1: "a"}
        _ = fi9
        _ = v13
    }
    if ((v6 || false) || (722 > 255)) {
        v14 := GenS1{F0: 4294967295, F1: ("" + "hello")}
        fmt.Println(v6, v6)
        v1 -= (^(v1))
        _ = v14
    }
    v5 = ""
    var v15 uint = uint(v1)
    v16 := []int{42}
    fmt.Println(v4, v5)
    var v17 uint64 = 0
    _ = len(v16)
    for fi18 := 0; fi18 < 5; fi18++ {
        v17 -= v17
        var v19 int8 = -3
        if v6 {
            var v20 int64 = -9223372036854775808
            v6 = v6
            if (v4 > v15) {
                v21 := GenS1{F0: 1, F1: v5}
                var v22 int64 = (v20 | (v20 << ((255) & 63)))
                _ = v21
                _ = v22
            }
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    var v23 float32 = 3.14159
    var v24 int32 = 357
    var v25 uint64 = (v17 * v17)
    if v6 {
        var v26 uint32 = 2
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 int64 = -926
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    switch 788 {
    case 1:
        if len(v16) > 0 {
            v16[0] = 32768
        }
    case -846:
        var v29 uint = 42
        _ = v29
    case -381:
        for fi30 := 0; fi30 < 3; fi30++ {
            v24 -= (v24 + (v24 << ((v4) & 63)))
            v31 := GenS1{F0: (v17 - 0), F1: v5}
            if len(v16) > 0 {
                v16[0] = (fi30 / (((fi30 - fi30)) | 1))
            }
            _ = fi30
            _ = v31
        }
    default:
        var v32 int8 = -128
        _ = v32
    }
    if len(v16) > 0 {
        v16[0] = 839
    }
    var v33 uint32 = (v1 << ((v15) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
