package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
func genFn0(p0 bool) uint16 {
    v5 := []int{32768}
    switch 4294967295 {
    case 0:
        for fi6 := 0; fi6 < 20; fi6++ {
            v5 = append(v5, (fi6 & fi6))
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 int64 = 524
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
    default:
        switch -378 {
        case 128:
            p0 = (!(1e10 <= float32(1.5)))
        }
    }
    var v9 uint64 = 0
    v10 := GenS1{F0: 92, F1: 255}
    v9 = (v9 % ((v9) | 1))
    _ = p0
    _ = v10
    _ = v5
    _ = v9
    return 1
}

func main() {
    var v1 float32 = 0.0
    v2 := []int{127, -2147483648}
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    var v11 int64 = 1
    if ((2 != 2) && false) {
        v3[(("\n\t" + "\n\t") + "\n\t")] = (((!((2 != 2) || (false || (65535 == 4294967295)))) && (true || (-1 >= 0))) || (!false))
    } else {
        fmt.Println(v1)
    }
    var v12 uint32 = 65535
    v13 := GenS1{F0: v12, F1: 0}
    if true {
        switch ("\n\t" + "hello") {
        case "foo bar":
            if ((false || true) || false) {
                _ = len(v3)
                var v14 int8 = -128
                _ = v14
            }
        case "a":
            var v15 uint8 = 61
            _ = v15
        default:
            v2 = append(v2, 1)
        }
        if true {
            if false {
                v1 += 0.1
                v11 = (^(v11))
            }
            fmt.Println(v12, v11, v11)
            v16 := GenS1{F0: 2, F1: 136}
            _ = v16
        } else {
            v17 := GenS1{F0: 2147483648, F1: 65535}
            var v18 string = ""
            switch v18 {
            case "hello":
                v2 = append(v2, 4294967295)
            case "\n\t":
                delete(v3, ("\n\t" + v18))
            case "foo bar":
                v11 -= 1
            default:
                delete(v3, "a")
            }
            _ = v17
            _ = v18
        }
        fmt.Println(v11)
    }
    var v19 int64 = (v11 / (((v11 << ((692) & 63))) | 1))
    if (true || false) {
        _ = len(v3)
        v12 += (v12 * 0)
        var v20 float32 = v1
        _ = v20
    }
    v2 = append(v2, -352)
    if ((!false) || false) {
        for fi21 := 0; fi21 < 14; fi21++ {
            v1 -= (v1 - -2.25)
            v22 := v4[v12]
            v22 += 65535
            _ = fi21
            _ = v22
        }
        var v23 float32 = 0.0
        fmt.Println(v19, v1, v1)
        _ = v23
    }
    var v24 int64 = (^(v11))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
