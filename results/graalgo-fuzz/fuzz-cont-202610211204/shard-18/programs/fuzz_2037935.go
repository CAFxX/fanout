package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
func genFn0() uint32 {
    if (-1 > 193) {
        var v4 string = ("" + "foo bar")
        for fi5 := 0; fi5 < 7; fi5++ {
            v4 = v4
            _ = fi5
        }
        var v6 float32 = 1.5
        _ = v4
        _ = v6
    }
    var v7 float64 = (3.14159 * (3.14159 / 1.0))
    var v8 uint = 0
    fmt.Println(v8)
    _ = v7
    _ = v8
    return 2147483648
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint32 = 964
    var v3 bool = false
    v2 += (v2 << ((1) & 63))
    v9 := []int{128, 25}
    _ = len(v9)
    switch "\n\t" {
    case "foo bar":
        for fi10 := 0; fi10 < 1; fi10++ {
            v2 = genFn0()
            _ = fi10
        }
    case "a":
        if v3 {
            v2 += 65535
            if (v3 || v3) {
                var v11 uint32 = (v2 >> ((0) & 63))
                _ = v11
            } else {
                fmt.Println(v1)
                var v12 bool = true
                _ = v12
            }
        } else {
            var v13 int8 = 42
            v2 += (v2 - genFn0())
            _ = v13
        }
    case "\n\t":
        if (v3 || (false || v3)) {
            _ = len(v9)
        }
    }
    var v14 uint8 = 255
    if (32767 < -16) {
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 int64 = 9223372036854775807
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 bool = v3
                v9 = append(v9, (fi15 & (fi15 * -1)))
                if (!(1 != -1)) {
                    v19 := GenS1{F0: 42, F1: 18446744073709551615}
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
            var v20 uint32 = (^(v2))
            _ = fi15
            _ = v16
            _ = v20
        }
        v21 := GenS1{F0: -1, F1: 42}
        _ = v21
    }
    if (!(!(v3 || (1 < int16(v1))))) {
        var v22 uint32 = v2
        var v23 uint = uint(v1)
        _ = v22
        _ = v23
    }
    switch -2147483648 {
    case 1:
        fmt.Println(v14, v14)
    case 0:
        fmt.Println(v1, v1, v14)
    case -68:
        if (-796 != -67) {
            v2 = genFn0()
        } else {
            fmt.Println(v2)
            fmt.Println(v2, v2, v1)
            v14 += v14
        }
    }
    var v24 int8 = int8(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v9
}
