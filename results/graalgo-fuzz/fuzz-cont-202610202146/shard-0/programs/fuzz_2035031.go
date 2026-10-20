package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }

func main() {
    var v1 int = 275
    var v2 int64 = -9223372036854775808
    var v3 int32 = 2147483647
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 uint = 255
        _ = fi4
        _ = v5
    }
    v2 = (v2 >> ((0) & 63))
    v6 := GenS1{F0: int8(v2), F1: 65535, F2: "foo bar"}
    v2 += 690
    if (v3 != v3) {
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := GenS1{F0: -1, F1: 0, F2: "a"}
            _ = fi7
            _ = v8
        }
    }
    v1 -= 562
    v2 -= v2
    fmt.Println(v2, v3)
    v3 += v3
    if ((1.5 != (3.14159 / 1.5)) && (4294967295 < 0)) {
        var v9 uint16 = 900
        var v10 int8 = 71
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS1{F0: -128, F1: v9, F2: ((("hello" + "hello") + (("" + ("" + "hello")) + "a")) + "\n\t")}
            var v13 int32 = (v3 % ((-1) | 1))
            if (42 != 1) {
                v10 = (v10 >> ((65535) & 63))
                v13 = (v3 % ((v13) | 1))
                v2 += -9223372036854775808
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
        _ = v9
    }
    fmt.Println(v2, v1)
    v1 -= v1
    v1 = (v1 - v1)
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
}
