package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int }
type GenS2 struct { F0 string; F1 int32 }
func genFn0() int16 {
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := []int{fi4, (fi4 / ((fi4) | 1))}
        v5 = append(v5, 127)
        var v6 uint = 4294967295
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: 2147483648, F2: 32768}
    v8 := []int{785, 752}
    _ = v7
    _ = v8
    return 231
}
func genFn1() int64 {
    if ("foo bar" >= "") {
        v9 := []int{9223372036854775807, 240}
        _ = v9
    } else {
        var v10 int16 = -32768
        _ = v10
    }
    var v11 int64 = 1
    _ = v11
    return v11
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{557, 65535, 4294967295}
    var v3 string = "hello"
    if false {
        if true {
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 int8 = 93
                var v14 int8 = (^(v13))
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        fmt.Println(v3)
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    }
    if false {
        v15 := GenS2{F0: (("foo bar" + v3) + v3), F1: -2147483648}
        _ = v15
    } else {
        v3 = v3
    }
    var v16 uint8 = 42
    for fi17 := 0; fi17 < 4; fi17++ {
        if len(v2) > 0 {
            v2[0] = 0
        }
        fmt.Println(fi17, fi17)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi17
    }
    v18 := GenS2{F0: "hello", F1: 443}
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    var v19 int16 = 32767
    var v20 string = ("" + v3)
    _ = len(v2)
    for fi21 := 0; fi21 < 10; fi21++ {
        if (1 != 2147483647) {
            v1["foo bar"] = 2147483647
            var v22 int64 = -925
            _ = v22
        }
        var v23 uint32 = 118
        _ = fi21
        _ = v23
    }
    if (true || false) {
        fmt.Println(v20, v19)
    } else {
        var v24 uint32 = 230
        _ = v24
    }
    fmt.Println(v3)
    switch 2147483647 {
    case 32768:
        v20 = "foo bar"
    }
    var v25 float64 = float64((0.0 - 1.5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
}
