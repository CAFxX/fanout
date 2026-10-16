package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int32 }
func genFn0(p0 int) int64 {
    p0 = -2147483648
    v4 := GenS1{F0: ("a" + ("hello" + "")), F1: 42, F2: 42}
    p0 -= p0
    var v5 string = ("" + "")
    var v6 int16 = -32768
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 int = 37
    v2 := []uint32{65535, 42}
    var v3 string = "foo bar"
    for fi7 := 0; fi7 < 16; fi7++ {
        v1 = (v1 | (v1 + (v1 - 2147483647)))
        _ = fi7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 bool = true
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 int64 = genFn0((fi8 >> ((18446744073709551615) & 63)))
            _ = fi10
            _ = v11
        }
        if (2147483647 >= 765) {
            var v12 uint16 = 1
            if len(v2) > 0 {
                v2[0] = 1
            }
            _ = v12
        } else {
            fmt.Println(v1, v1, v3)
        }
        _ = fi8
        _ = v9
    }
    var v13 uint = 42
    var v14 float32 = (1.5 / 1.5)
    var v15 int = (v1 >> (((v13 - 42)) & 63))
    v14 = 1e10
    v2 = append(v2, 2147483647)
    var v16 uint64 = 2
    fmt.Println(v1, v15)
    if false {
        _ = len(v2)
    }
    switch v15 {
    case 4294967295:
        var v17 int32 = -1
        _ = v17
    }
    var v18 bool = false
    _ = len(v2)
    v16 += (^(v16))
    var v19 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
