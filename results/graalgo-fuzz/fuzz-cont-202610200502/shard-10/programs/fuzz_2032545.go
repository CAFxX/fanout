package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 0
    v3 := make(map[string]string)
    v4 := []string{"\n\t", "\n\t", ((("foo bar" + "hello") + ("" + "\n\t")) + "\n\t")}
    v5 := make(map[string]string)
    fmt.Println(v2)
    v3["\n\t"] = ("" + ("a" + "a"))
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 float32 = (float32(1.5) + 3.14159)
        v2 = 1
        _ = fi6
        _ = v7
    }
    var v8 uint64 = 298
    if false {
        var v9 string = "hello"
        _ = len(v4)
        v10 := GenS1{F0: 18446744073709551615, F1: -1, F2: 14}
        _ = v10
        _ = v9
    } else {
        var v11 uint64 = 4294967295
        var v12 uint = 268
        var v13 string = ("" + "a")
        _ = v11
        _ = v12
        _ = v13
    }
    v2 -= (v2 % ((v2) | 1))
    fmt.Println(v8, v8)
    if (!false) {
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
        v8 += v8
        v2 += (v2 / (((v2 % ((v2) | 1))) | 1))
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        switch "a" {
        case "":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            v15 := v1[-925]
            _ = v15
        }
        _ = fi14
    }
    switch ("hello" + (("a" + "") + "a")) {
    case "\n\t":
        var v16 int16 = 1000
        _ = v16
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v17 := GenS1{F0: 254, F1: 42, F2: 1}
    _ = len(v4)
    v2 += v2
    if true {
        v2 += 0
        switch 2147483647 {
        case -864:
            var v18 uint = 0
            _ = v18
        case -1:
            v2 = v2
        case 32767:
            v2 = v2
        }
    }
    var v19 uint32 = 1
    v20 := GenS1{F0: 65535, F1: 42, F2: 119}
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
