package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint64 = 42
    var v2 bool = (2147483647 > -758)
    var v3 int32 = -887
    var v4 uint8 = 2
    if v2 {
        fmt.Println(v4, v3, v1)
        v5 := GenS1{F0: 776}
        if (v3 <= int32(v1)) {
            fmt.Println(v2)
            v6 := []int{498, 1}
            var v7 int = -2147483648
            _ = v6
            _ = v7
        }
        _ = v5
    }
    fmt.Println(v3)
    v8 := []int{739, 4294967295, -2147483648}
    v4 = v4
    if (2 == 42) {
        if len(v8) > 0 {
            v8[0] = 4294967295
        }
        var v9 int8 = -39
        fmt.Println(v9, v9)
        _ = v9
    } else {
        v3 += v3
        v4 = v4
    }
    fmt.Println(v1)
    var v10 int16 = 487
    for fi11 := 0; fi11 < 19; fi11++ {
        if (v3 <= (v3 << ((65535) & 63))) {
            var v12 int16 = (v10 ^ (v10 << ((2) & 63)))
            _ = v12
        }
        var v13 uint16 = 42
        var v14 int16 = v10
        _ = fi11
        _ = v13
        _ = v14
    }
    var v15 bool = (208 <= 0)
    for fi16 := 0; fi16 < 4; fi16++ {
        v10 -= (^(v10))
        _ = fi16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        switch ("" + ("\n\t" + ("" + ("hello" + "a")))) {
        case "foo bar":
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 int8 = 127
                _ = fi18
                _ = v19
            }
        case "hello":
            v8 = append(v8, 127)
        case "a":
            fmt.Println(v15)
        default:
            v1 = 0
        }
        var v20 uint64 = v1
        v10 = -218
        _ = fi17
        _ = v20
    }
    if len(v8) > 0 {
        v8[0] = -9223372036854775808
    }
    var v21 int8 = -27
    v21 += v21
    var v22 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v8
}
