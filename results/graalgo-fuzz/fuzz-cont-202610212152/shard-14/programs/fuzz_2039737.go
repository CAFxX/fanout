package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint; F1 bool; F2 int8 }
func genFn0() uint64 {
    var v4 int32 = 222
    if false {
        for fi5 := 0; fi5 < 5; fi5++ {
            v6 := GenS1{F0: (fi5 >> ((0) & 63)), F1: 18446744073709551615}
            fmt.Println(v4, v4)
            v4 = v4
            _ = fi5
            _ = v6
        }
        v4 += (v4 | v4)
    }
    v7 := GenS1{F0: 344, F1: 313}
    var v8 int8 = 100
    v9 := GenS1{F0: 721, F1: 806}
    _ = v4
    _ = v7
    _ = v8
    _ = v9
    return 932
}
func genFn1(p0 bool, p1 int64) int8 {
    switch 128 {
    case 127:
        var v10 uint = 65535
        _ = v10
    case 9223372036854775807:
        var v11 string = ("" + ("\n\t" + "hello"))
        _ = v11
    case -644:
        var v12 int64 = (p1 / ((9223372036854775807) | 1))
        _ = v12
    }
    var v13 int8 = 42
    p1 += (p1 ^ 1)
    v14 := GenS1{F0: 0, F1: 4294967295}
    v15 := GenS1{F0: 1, F1: 1}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return (v13 ^ v13)
}

func main() {
    var v1 float64 = 0.0
    var v2 string = ((("hello" + ("hello" + (("a" + "a") + ("\n\t" + "foo bar")))) + "a") + "")
    var v3 uint32 = 2147483647
    var v16 uint32 = v3
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 int = 2147483647
        var v19 uint64 = 1
        fmt.Println(v16)
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 uint = 104
    v21 := GenS1{F0: -2147483648, F1: v20}
    v2 = v2
    v2 = ("hello" + v2)
    switch (((v2 + (v2 + v2)) + v2) + v2) {
    case "\n\t":
        for fi22 := 0; fi22 < 3; fi22++ {
            v16 = v3
            _ = fi22
        }
    }
    v23 := GenS2{F0: v20, F1: true, F2: -1}
    v3 += (v16 << ((v20) & 63))
    v1 = (v1 - v1)
    var v24 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
}
