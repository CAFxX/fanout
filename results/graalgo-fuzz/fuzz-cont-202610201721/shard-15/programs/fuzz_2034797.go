package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 bool; F1 uint64; F2 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"", "hello"}
    v3 := make(map[uint32]string)
    var v4 int16 = -303
    var v5 uint64 = 4294967295
    v6 := GenS1{F0: (("hello" + "a") + ("\n\t" + "\n\t")), F1: v4}
    if true {
        fmt.Println(v4)
        v7 := GenS1{F0: ("hello" + ("" + ("" + ""))), F1: -1}
        var v8 string = ""
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: "", F1: v4}
    for fi10 := 0; fi10 < 8; fi10++ {
        v4 = (-(v4))
        v11 := GenS1{F0: "foo bar", F1: v4}
        var v12 int = fi10
        _ = fi10
        _ = v11
        _ = v12
    }
    if (true && true) {
        var v13 float64 = ((1e10 * float64(0.0)) - (1.5 * (3.14159 * 3.14159)))
        v13 = v13
        _ = v13
    } else {
        v5 -= 1
        var v14 float32 = -2.25
        _ = v14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint32 = 907
        v4 = v4
        _ = fi15
        _ = v16
    }
    fmt.Println(v5)
    v1[2147483647] = 261
    var v17 int8 = 127
    switch -2147483648 {
    case 9223372036854775807:
        var v18 int64 = 453
        _ = v18
    case 764:
        delete(v1, 2147483647)
    case -930:
        v19 := GenS2{F0: (true || false), F1: v5, F2: (!(1e10 <= (1e10 * (0.0 - 1.5))))}
        _ = v19
    default:
        v20 := GenS1{F0: "hello", F1: 0}
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        v5 += (v5 | (^(v5)))
        _ = fi21
    }
    delete(v3, 2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
