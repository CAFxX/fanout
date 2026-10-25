package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 string }
type GenS2 struct { F0 bool; F1 int; F2 int64 }
func genFn0() int32 {
    var v4 float32 = 0.1
    if false {
        v5 := GenS1{F0: 626, F1: 2147483647, F2: ""}
        var v6 float64 = -2.25
        _ = v5
        _ = v6
    }
    var v7 string = "hello"
    _ = v4
    _ = v7
    return -2147483648
}
func genFn1() int64 {
    if true {
        var v8 int = -2147483648
        fmt.Println(v8)
        _ = v8
    }
    var v9 int16 = 1000
    var v10 uint16 = 255
    if true {
        v10 += 0
    }
    _ = v10
    _ = v9
    return 201
}

func main() {
    var v1 int64 = 0
    var v2 string = ""
    var v3 uint16 = 0
    var v11 float64 = (1e10 * (1e10 * ((((-2.25 - float64(1e10)) - 0.0) * (((3.14159 / 1e10) * 0.0) + 0.1)) - (0.0 / 1.0))))
    var v12 float64 = v11
    fmt.Println(v12, v11)
    v13 := []int{}
    if false {
        var v14 uint = 4294967295
        _ = v14
    } else {
        fmt.Println(v12)
        _ = len(v13)
    }
    if true {
        var v15 uint = 18446744073709551615
        var v16 float64 = -2.25
        switch -9223372036854775808 {
        case 2147483647:
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        case 4294967295:
            var v17 int = 1
            _ = v17
        case 32768:
            fmt.Println(v15, v3, v12)
        }
        _ = v15
        _ = v16
    }
    v3 -= 163
    switch -579 {
    case 0:
        if ((false && false) && true) {
            if len(v13) > 0 {
                v13[0] = -662
            }
            v11 = 0.1
            if true {
                var v18 uint64 = 217
                if len(v13) > 0 {
                    v13[0] = 175
                }
                _ = v18
            }
        } else {
            v11 += v11
        }
    case 4294967295:
        v19 := GenS2{F0: (1 <= 0), F1: 0, F2: (v1 << ((18446744073709551615) & 63))}
        _ = v19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        fmt.Println(fi20, fi20, v12)
        fmt.Println(v3, v1)
        _ = fi20
    }
    fmt.Println(v2, v1)
    v3 = 999
    v13 = append(v13, 32768)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        v12 -= v11
        _ = fi21
    }
    var v22 uint64 = 2
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v23 := GenS1{F0: 1, F1: 271, F2: "hello"}
    var v24 uint8 = 139
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
