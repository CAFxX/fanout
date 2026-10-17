package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }

func main() {
    v1 := []uint32{}
    var v2 bool = true
    var v3 uint = 56
    var v4 int8 = -119
    var v5 int = 2147483647
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if v2 {
        if (!(1 >= 1)) {
            v5 = v5
        }
    }
    var v6 string = ("foo bar" + (("\n\t" + ("\n\t" + "foo bar")) + ((("foo bar" + ("a" + "a")) + "foo bar") + ("" + ""))))
    var v7 int8 = (v4 >> (((^(v3))) & 63))
    v8 := GenS1{F0: 2147483647}
    for fi9 := 0; fi9 < 16; fi9++ {
        v7 -= (v7 ^ (v4 << ((v3) & 63)))
        v5 += fi9
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        _ = fi9
    }
    var v10 int = v5
    if (1e10 >= 1.5) {
        v6 = "hello"
        v11 := GenS1{F0: 65535}
        var v12 uint32 = 488
        _ = v11
        _ = v12
    } else {
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
        v5 -= (v5 - (v10 % ((-659) | 1)))
        v10 += v10
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS1{F0: 2}
        v15 := GenS2{F0: v3}
        if len(v1) > 0 {
            v1[0] = 479
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 int64 = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
