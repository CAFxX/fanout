package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
func genFn0(p0 bool) uint8 {
    var v4 string = ("hello" + "a")
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 int8 = 1
            v8 := GenS1{F0: 65535, F1: 0}
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    if p0 {
        v9 := GenS1{F0: 2, F1: -958}
        _ = v9
    } else {
        fmt.Println(v4)
        v10 := []int{}
        _ = v10
    }
    p0 = (!true)
    switch v4 {
    case "":
        v11 := []int{9223372036854775807, 4294967295}
        _ = v11
    case "foo bar":
        var v12 uint32 = 65535
        _ = v12
    case "\n\t":
        v13 := []int{0, -217, -395, 811}
        _ = v13
    }
    _ = p0
    _ = v4
    return 128
}

func main() {
    v1 := []int{128, -2147483648, -9223372036854775808}
    var v2 float64 = ((1e10 + 3.14159) - 1e10)
    var v3 uint = 42
    v14 := GenS1{F0: 65535, F1: 1}
    if (736 != 42) {
        fmt.Println(v2)
        if (true || (!false)) {
            fmt.Println(v2)
        } else {
            v1 = append(v1, 297)
            v3 = v3
            switch -9223372036854775808 {
            case 922:
                if false {
                    v3 = (v3 & v3)
                }
            case 128:
                for fi15 := 0; fi15 < 3; fi15++ {
                    _ = len(v1)
                    _ = fi15
                }
            case 1:
                v16 := GenS1{F0: 255, F1: 42}
                _ = v16
            }
        }
        v17 := GenS1{F0: 1, F1: -1}
        _ = v17
    }
    fmt.Println(v2, v2)
    v18 := GenS1{F0: 65535, F1: -1}
    var v19 int64 = 242
    v19 += 949
    v20 := GenS1{F0: 548, F1: 937}
    v2 += v2
    var v21 float64 = v2
    v3 = (v3 >> (((v3 - 98)) & 63))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v22 := GenS1{F0: 42, F1: 0}
    var v23 int32 = -960
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
