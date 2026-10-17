package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 string }
func genFn0() string {
    var v6 int = 1
    v6 += v6
    var v7 int64 = 42
    var v8 uint = 774
    _ = v6
    _ = v7
    _ = v8
    return (("\n\t" + (("" + (("a" + "\n\t") + (("" + "foo bar") + ""))) + ("\n\t" + "hello"))) + ("foo bar" + "hello"))
}
func genFn1(p0 uint64) int64 {
    var v9 uint8 = 154
    var v10 uint = 42
    if ((32767 == 127) || true) {
        p0 -= p0
    } else {
        v11 := GenS1{F0: -450, F1: 1}
        _ = v11
    }
    var v12 int8 = 100
    _ = p0
    _ = v10
    _ = v12
    _ = v9
    return 1
}

func main() {
    var v1 uint32 = 1
    var v2 int16 = 1
    var v3 uint64 = 1
    var v4 int8 = 37
    var v5 int32 = int32(v2)
    v13 := []int{}
    if true {
        v14 := GenS2{F0: v3, F1: 18446744073709551615, F2: (v2 >> ((255) & 63))}
        fmt.Println(v2, v5)
        _ = v14
    }
    fmt.Println(v1, v4)
    v15 := GenS2{F0: uint64(v1), F1: 1, F2: -1}
    for fi16 := 0; fi16 < 8; fi16++ {
        fmt.Println(v3, v5)
        v1 -= 1
        var v17 int = fi16
        _ = fi16
        _ = v17
    }
    v18 := GenS3{F0: (!((true && (v5 != (^(v5)))) && ((127 < 10) || (1.5 == 0.1))))}
    v19 := GenS4{F0: 100, F1: "hello"}
    v13 = append(v13, -9223372036854775808)
    fmt.Println(v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
