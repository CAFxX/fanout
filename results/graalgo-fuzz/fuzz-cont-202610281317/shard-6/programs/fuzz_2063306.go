package main

import (
    "fmt"
)

func genFn0(p0 int16) uint8 {
    v4 := []int{-2147483648, -2147483648}
    var v5 uint16 = 1
    _ = p0
    _ = v4
    _ = v5
    return 13
}
func genFn1(p0 int32) bool {
    p0 -= -671
    var v6 int16 = 32767
    _ = p0
    _ = v6
    return false
}

func main() {
    var v1 float32 = ((1.5 - (1.5 * 0.1)) - 1.5)
    var v2 uint32 = 42
    var v3 int16 = 32767
    for fi7 := 0; fi7 < 3; fi7++ {
        v3 = (v3 - (v3 >> ((65535) & 63)))
        _ = fi7
    }
    v2 += 184
    var v8 int = 783
    v1 = (v1 - v1)
    fmt.Println(v1, v8, v8)
    v1 = ((((1e10 * 3.14159) - (3.14159 * 1e10)) + 0.0) + -2.25)
    if false {
        switch -949 {
        case 2147483647:
            var v9 uint8 = 224
            _ = v9
        case -908:
            v2 = (v2 & (v2 << ((2) & 63)))
        }
    } else {
        var v10 uint16 = 65535
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 int32 = -2147483648
            var v13 int8 = 0
            var v14 string = "foo bar"
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        v15 := []int{}
        _ = v10
        _ = v15
    }
    var v16 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
