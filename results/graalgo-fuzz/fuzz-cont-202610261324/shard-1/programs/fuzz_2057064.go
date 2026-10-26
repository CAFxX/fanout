package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []uint32{4294967295, 0, 65535}
    var v2 uint8 = 156
    var v3 float32 = 0.0
    var v4 uint16 = 992
    if false {
        v2 -= 26
        v5 := GenS1{F0: 1}
        v2 = (v2 ^ (v2 + v2))
        _ = v5
    } else {
        v6 := GenS1{F0: 65535}
        _ = v6
    }
    switch 32767 {
    case 128:
        fmt.Println(v2)
    }
    fmt.Println(v2, v3)
    if false {
        fmt.Println(v2, v4, v3)
        var v7 int8 = 0
        var v8 int32 = -2147483648
        _ = v7
        _ = v8
    } else {
        v3 -= v3
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 = ((3.14159 + v3) - (v3 / 1.5))
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        _ = len(v1)
        _ = fi9
    }
    v3 = 1e10
    v3 += (v3 - 3.14159)
    v4 -= v4
    v10 := GenS1{F0: 2}
    v3 = 3.14159
    fmt.Println(v2, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
