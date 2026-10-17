package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
func genFn0(p0 int8, p1 int) uint {
    switch "a" {
    case "":
        var v4 int8 = p0
        _ = v4
    }
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 1, F1: 0}
        var v7 int64 = -9223372036854775808
        _ = fi5
        _ = v6
        _ = v7
    }
    if false {
        var v8 int16 = -1
        var v9 uint16 = 2
        for fi10 := 0; fi10 < 16; fi10++ {
            for fi11 := 0; fi11 < 2; fi11++ {
                p0 += (^(p0))
                fmt.Println(v9)
                v12 := []int{}
                v13 := GenS1{F0: 555, F1: 1}
                p0 += (p0 % ((0) | 1))
                fmt.Println(p0, v9, p0)
                _ = fi11
                _ = v12
                _ = v13
            }
            v14 := GenS1{F0: v9, F1: 2147483647}
            _ = fi10
            _ = v14
        }
        _ = v8
        _ = v9
    } else {
        p1 += (p1 % (((p1 | (p1 << ((584) & 63)))) | 1))
        p0 = p0
        var v15 uint64 = 42
        _ = v15
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]bool)
    var v3 int16 = -697
    for fi16 := 0; fi16 < 8; fi16++ {
        v3 = v3
        switch fi16 {
        case 1:
            fmt.Println(fi16, fi16, v1)
        default:
            delete(v2, ("hello" + (("\n\t" + "\n\t") + "hello")))
        }
        v3 = (v3 ^ v3)
        _ = fi16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 int16 = v3
        _ = fi17
        _ = v18
    }
    v3 += v3
    var v19 uint16 = 1
    switch -775 {
    case -2147483648:
        v20 := GenS1{F0: (v19 >> ((2) & 63)), F1: 42}
        _ = v20
    case 32768:
        fmt.Println(v1)
    case 9223372036854775807:
        fmt.Println(v19, v1)
    default:
        v3 += -1
    }
    v1 += (^(v1))
    if (false || (127 >= -61)) {
        var v21 uint32 = 65535
        _ = v21
    }
    switch 32768 {
    case 0:
        var v22 uint16 = v19
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
