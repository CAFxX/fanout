package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS3 struct { F0 bool; F1 int; F2 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -533
    var v3 string = "a"
    var v4 float64 = (1.5 + -2.25)
    v5 := make(map[string]string)
    var v6 uint8 = 42
    var v7 float64 = -2.25
    var v8 uint8 = 194
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int16 = -1
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 int = fi9
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    if false {
        var v13 int = (v2 + -9223372036854775808)
        v14 := GenS1{F0: 827}
        _ = v13
        _ = v14
    }
    v4 = float64(1.5)
    v15 := GenS1{F0: 0}
    v16 := GenS2{F0: (v6 + v8), F1: 2, F2: 0}
    v17 := GenS1{F0: 1}
    if false {
        v3 = v3
    } else {
        fmt.Println(v6, v8, v6)
        var v18 string = (v3 + v3)
        _ = v18
    }
    v19 := GenS3{F0: false, F1: (v2 * (v2 ^ (-(v2)))), F2: (v6 % ((39) | 1))}
    if false {
        v3 = ((v3 + (v3 + v3)) + (v3 + ((("hello" + "") + "") + ("a" + "hello"))))
    }
    switch (-(v2)) {
    case -984:
        v7 -= v7
    case -2147483648:
        v6 += (^(v6))
    case 9223372036854775807:
        v4 -= v7
    }
    var v20 string = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
