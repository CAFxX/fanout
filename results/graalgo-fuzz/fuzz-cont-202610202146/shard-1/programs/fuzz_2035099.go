package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint32 }
type GenS2 struct { F0 int16; F1 int32; F2 int8 }
func genFn0() int32 {
    var v4 bool = (1.5 != 1e10)
    v4 = v4
    v4 = (128 == 1)
    _ = v4
    return -100
}
func genFn1(p0 uint) bool {
    v5 := GenS1{F0: p0, F1: -1, F2: 0}
    var v6 int8 = 100
    switch "hello" {
    case "a":
        v6 = v6
    case "":
        v6 = v6
    case "hello":
        v7 := GenS2{F0: 32767, F1: -2147483648, F2: (-(v6))}
        _ = v7
    }
    _ = p0
    _ = v5
    _ = v6
    return false
}

func main() {
    var v1 int = -665
    var v2 uint32 = 2
    var v3 uint64 = 572
    if (-66 != -67) {
        var v8 bool = (1e10 <= 1.5)
        _ = v8
    }
    v9 := []int{0, -1}
    v10 := GenS2{F0: 1000, F1: 732, F2: 0}
    var v11 int = (v1 / ((-486) | 1))
    if true {
        v12 := GenS2{F0: int16(v11), F1: genFn0(), F2: 1}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v12
    }
    switch v11 {
    case 9223372036854775807:
        v2 = 957
    case 2147483647:
        v11 = (v1 % (((v1 + v11)) | 1))
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            if false {
                v11 = fi14
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi14
        }
        _ = fi13
    }
    fmt.Println(v2, v11, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
