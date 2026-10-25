package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint }
type GenS2 struct { F0 int8; F1 uint16 }

func main() {
    var v1 uint16 = 944
    v2 := []string{("hello" + "hello")}
    v3 := make(map[int]string)
    var v4 bool = true
    var v5 string = "foo bar"
    if (v4 && (!(!(false && false)))) {
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 int64 = -1
            v7 += v7
            var v8 string = v5
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    if true {
        fmt.Println(v5, v4, v5)
        v9 := GenS1{F0: -9223372036854775808, F1: v1, F2: 4294967295}
        if (113 <= -128) {
            for fi10 := 0; fi10 < 20; fi10++ {
                if (false || (!(576 < 272))) {
                    switch "" {
                    case "\n\t":
                        var v11 uint8 = 2
                        _ = v11
                    }
                }
                _ = fi10
            }
        }
        _ = v9
    } else {
        var v12 float32 = -2.25
        v13 := GenS2{F0: -78, F1: (v1 | 872)}
        var v14 string = ((v5 + "") + ("hello" + "\n\t"))
        _ = v12
        _ = v13
        _ = v14
    }
    fmt.Println(v1)
    v5 = ("\n\t" + ("a" + v5))
    var v15 float32 = 0.1
    var v16 uint32 = 2147483647
    for fi17 := 0; fi17 < 16; fi17++ {
        delete(v3, -2147483648)
        fmt.Println(v4, v1, v1)
        _ = fi17
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v18 int64 = -451
    v18 += v18
    v5 = (((v5 + "") + v5) + v5)
    var v19 uint8 = 31
    if v4 {
        v20 := GenS2{F0: 127, F1: v1}
        var v21 float32 = -2.25
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := v3[fi22]
            v16 = (v16 + (v16 >> ((uint(v1)) & 63)))
            _ = fi22
            _ = v23
        }
        var v24 string = "hello"
        v25 := v3[32768]
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
