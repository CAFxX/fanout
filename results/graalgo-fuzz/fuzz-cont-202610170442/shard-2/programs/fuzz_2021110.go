package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int32, p1 uint) int16 {
    for fi4 := 0; fi4 < 13; fi4++ {
        v5 := GenS1{F0: 832, F1: 1}
        _ = fi4
        _ = v5
    }
    p1 = p1
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float32 = (1e10 - 3.14159)
        var v8 uint64 = 760
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    return -32768
}
func genFn1() int8 {
    var v9 uint32 = 2147483647
    v9 = 42
    _ = v9
    return 127
}

func main() {
    var v1 int8 = 127
    var v2 float32 = 1.5
    v3 := make(map[uint32]bool)
    v1 -= 85
    v10 := v3[65535]
    fmt.Println(v2, v1, v10)
    v10 = v10
    for fi11 := 0; fi11 < 18; fi11++ {
        switch (^(fi11)) {
        case -1:
            var v12 int64 = -1
            _ = v12
        case 4294967295:
            v2 -= v2
        case 127:
            v2 = -2.25
        }
        v13 := GenS1{F0: 42, F1: 32767}
        _ = fi11
        _ = v13
    }
    v1 = v1
    if true {
        v14 := GenS1{F0: 42, F1: 255}
        delete(v3, 104)
        var v15 int = -599
        _ = v14
        _ = v15
    }
    var v16 uint8 = 229
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 uint32 = 4294967295
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        v16 -= 1
        fmt.Println(v1, v2)
        var v20 int8 = (^(v1))
        _ = fi19
        _ = v20
    }
    fmt.Println(v16, v1, v2)
    var v21 int8 = (v1 % (((v1 >> ((2) & 63))) | 1))
    var v22 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
