package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 bool; F2 int16 }
func genFn0() string {
    if (-2.25 == 3.14159) {
        var v5 uint8 = 2
        fmt.Println(v5)
        _ = v5
    }
    v6 := []int{}
    v7 := GenS1{F0: ("hello" + "hello"), F1: 0, F2: -1}
    _ = v6
    _ = v7
    return (("" + "hello") + "foo bar")
}
func genFn1() int8 {
    if true {
        var v8 int8 = 0
        var v9 uint16 = 42
        _ = v8
        _ = v9
    }
    switch -626 {
    case -9223372036854775808:
        var v10 uint16 = 255
        _ = v10
    case -2147483648:
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 uint16 = 0
            v12 -= (v12 & (v12 >> ((65535) & 63)))
            _ = fi11
            _ = v12
        }
    default:
        var v13 int16 = -1
        _ = v13
    }
    return -128
}

func main() {
    var v1 float32 = (1e10 / 3.14159)
    var v2 uint16 = 1
    var v3 int64 = -478
    var v4 string = ""
    var v14 int = 0
    fmt.Println(v3, v1)
    var v15 int8 = 9
    var v16 int32 = 42
    for fi17 := 0; fi17 < 18; fi17++ {
        v3 += v3
        var v18 float32 = 1.5
        _ = fi17
        _ = v18
    }
    var v19 uint64 = 2
    for fi20 := 0; fi20 < 15; fi20++ {
        v21 := GenS2{F0: "\n\t"}
        _ = fi20
        _ = v21
    }
    var v22 uint32 = 42
    v19 += (v19 ^ 89)
    switch (v14 << ((2) & 63)) {
    case -2147483648:
        v22 = v22
    case -966:
        if true {
            fmt.Println(v16, v16)
            v15 += 42
        }
    }
    fmt.Println(v22, v14, v22)
    v23 := GenS3{F0: 0, F1: ((v14 * v14) >= 32768), F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
