package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 int64, p1 int8) int {
    v4 := GenS1{F0: 32768, F1: 910}
    if false {
        switch -611 {
        case 1:
            switch -1 {
            case 128:
                p1 -= (-(p1))
            case -469:
                if (0 <= 42) {
                    fmt.Println(p1)
                    v5 := GenS1{F0: -2147483648, F1: 2147483647}
                    p1 = p1
                    _ = v5
                } else {
                    p1 = p1
                    p0 += -1000
                }
            }
        case -90:
            p0 += p0
        case 392:
            fmt.Println(p1, p0)
        default:
            v6 := GenS1{F0: 0, F1: 43}
            _ = v6
        }
        var v7 uint8 = 127
        for fi8 := 0; fi8 < 14; fi8++ {
            v7 -= v7
            _ = fi8
        }
        _ = v7
    }
    var v9 uint32 = 2
    v9 += 4294967295
    for fi10 := 0; fi10 < 2; fi10++ {
        v11 := []int{fi10, -925}
        for fi12 := 0; fi12 < 11; fi12++ {
            if ((!false) && true) {
                if true {
                    var v13 int8 = (p1 & (-(p1)))
                    _ = v13
                } else {
                    var v14 float64 = (0.0 - (1.5 * (((1.5 / 1.0) / (float64(((((1.5 - -2.25) - 3.14159) / 1.0) - (0.0 - 0.1))) / (3.14159 - 1e10))) + 1.5)))
                    _ = v14
                }
            }
            switch "a" {
            case "foo bar":
                v9 = v9
            case "\n\t":
                fmt.Println(fi12)
            }
            switch ("\n\t" + (("a" + ("\n\t" + "")) + ("hello" + "\n\t"))) {
            case "a":
                fmt.Println(fi10, v9, v9)
            case "hello":
                p0 += -1
            case "\n\t":
                p0 = -1
            }
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v9
    return 42
}

func main() {
    v1 := []int{-1}
    var v2 int32 = 0
    var v3 uint64 = 18446744073709551615
    if (4294967295 >= 51) {
        fmt.Println(v3)
    } else {
        v15 := GenS1{F0: -109, F1: v2}
        for fi16 := 0; fi16 < 4; fi16++ {
            v17 := GenS1{F0: (^(fi16)), F1: -2147483648}
            v2 -= (v2 & (v2 + (v2 | v2)))
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    var v18 int32 = (v2 * v2)
    v19 := GenS1{F0: genFn0(42, -128), F1: -1}
    v20 := GenS1{F0: 32768, F1: v18}
    for fi21 := 0; fi21 < 4; fi21++ {
        fmt.Println(v2)
        var v22 int64 = 526
        _ = fi21
        _ = v22
    }
    switch ("" + "a") {
    case "\n\t":
        v23 := GenS1{F0: 768, F1: 2147483647}
        _ = v23
    }
    var v24 uint64 = (v3 / (((^(v3))) | 1))
    v18 = (v18 >> ((42) & 63))
    switch 127 {
    case 822:
        v25 := GenS1{F0: 127, F1: 2147483647}
        _ = v25
    case 0:
        if (v2 <= v18) {
            var v26 int64 = 0
            v18 = v2
            _ = v26
        }
    case 32768:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v27 := GenS2{F0: 1, F1: -741}
        _ = v27
    }
    var v28 float32 = 1.5
    v18 = (v18 | (v2 >> ((2) & 63)))
    v28 -= (0.0 - v28)
    v29 := GenS2{F0: 1, F1: -9223372036854775808}
    switch -297 {
    case 4294967295:
        var v30 uint16 = 133
        _ = v30
    case 32767:
        if len(v1) > 0 {
            v1[0] = 0
        }
    case 128:
        v31 := GenS1{F0: -307, F1: -146}
        _ = v31
    }
    var v32 uint64 = (^(v3))
    v2 -= v2
    v33 := GenS1{F0: 4294967295, F1: (v2 >> ((0) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
}
