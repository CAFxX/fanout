package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int }
type GenS2 struct { F0 uint8; F1 bool; F2 uint8 }
func genFn0() uint64 {
    var v7 int8 = -1
    v7 -= 0
    _ = v7
    return 1
}
func genFn1(p0 int16) uint16 {
    v8 := []int{0}
    var v9 uint = 4294967295
    fmt.Println(p0, v9)
    _ = p0
    _ = v8
    _ = v9
    return 15
}

func main() {
    var v1 uint16 = 222
    var v2 int = 32768
    var v3 uint64 = 18446744073709551615
    var v4 int16 = 1000
    var v5 uint32 = 0
    var v6 bool = false
    if (false && v6) {
        var v10 bool = ((0 <= 1) || v6)
        _ = v10
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        v1 += genFn1(-499)
        _ = fi11
    }
    v12 := []int{}
    v13 := GenS1{F0: genFn0(), F1: v2, F2: (-(v2))}
    switch (^(v2)) {
    case -394:
        v3 = 4294967295
    case 1:
        for fi14 := 0; fi14 < 3; fi14++ {
            v1 = v1
            _ = fi14
        }
    case 42:
        v12 = append(v12, (v2 >> ((65535) & 63)))
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 int32 = int32(v1)
        _ = fi15
        _ = v16
    }
    fmt.Println(v2, v4)
    v5 = v5
    if v6 {
        var v17 bool = v6
        var v18 uint16 = (v1 / (((v1 << ((255) & 63))) | 1))
        switch 32768 {
        case 32767:
            var v19 int8 = -128
            _ = v19
        case -488:
            fmt.Println(v1, v1, v17)
        default:
            if (!false) {
                if ((v4 % (((v4 << ((352) & 63))) | 1)) == v4) {
                    var v20 int16 = v4
                    fmt.Println(v5)
                    _ = v20
                }
                v21 := GenS1{F0: (v3 & 736), F1: -2147483648, F2: (-(v2))}
                switch "" {
                case "\n\t":
                    var v22 uint64 = (v3 & v3)
                    _ = v22
                }
                _ = v21
            }
        }
        _ = v17
        _ = v18
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        _ = len(v12)
        var v24 int16 = (^(v4))
        v5 = v5
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        v4 = 1
        _ = fi25
    }
    switch ("" + ("hello" + "hello")) {
    case "foo bar":
        fmt.Println(v6)
    case "":
        var v26 int32 = int32(v4)
        _ = v26
    case "hello":
        v27 := GenS2{F0: 223, F1: (!(219 != v3)), F2: 0}
        _ = v27
    }
    var v28 uint8 = 127
    if (-1 == 162) {
        v5 = v5
    } else {
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 uint32 = 0
            var v31 int = int(v3)
            _ = fi29
            _ = v30
            _ = v31
        }
        v32 := GenS2{F0: 128, F1: (!(float32(3.14159) == -2.25)), F2: 2}
        _ = v32
    }
    var v33 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
