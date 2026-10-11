package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int64 }

func main() {
    var v1 int32 = -2147483648
    var v2 uint64 = 2
    v3 := make(map[string]uint32)
    var v4 uint8 = 118
    v5 := []uint32{4294967295, 2147483648, 0, 374}
    fmt.Println(v2, v1, v4)
    if (true || false) {
        var v6 uint16 = 457
        v7 := GenS1{F0: 0, F1: (v6 / (((v6 / ((v6) | 1))) | 1)), F2: 271}
        _ = v6
        _ = v7
    } else {
        fmt.Println(v2, v1, v2)
        v8 := v3[""]
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        v4 -= (v4 ^ (v4 + (^(v4))))
        v2 += (^(v2))
        _ = fi9
    }
    v4 = v4
    if (!((!true) && (true && (78 >= 42)))) {
        var v10 int8 = 72
        for fi11 := 0; fi11 < 15; fi11++ {
            v4 = (v4 % ((190) | 1))
            _ = fi11
        }
        switch 42 {
        case 2147483647:
            delete(v3, ("foo bar" + "hello"))
        case -1:
            switch -220 {
            case 2147483647:
                var v12 int64 = -493
                _ = v12
            }
        default:
            v13 := GenS1{F0: 65535, F1: uint16(v4), F2: 42}
            _ = v13
        }
        _ = v10
    }
    var v14 bool = true
    for fi15 := 0; fi15 < 5; fi15++ {
        if len(v5) > 0 {
            v5[0] = 2147483648
        }
        _ = fi15
    }
    var v16 float32 = 0.1
    if (0 <= 93) {
        var v17 string = ("a" + (("a" + (("a" + "") + ("a" + ""))) + (("hello" + "foo bar") + "foo bar")))
        _ = v17
    } else {
        var v18 int32 = (v1 & (v1 + v1))
        v19 := v3[""]
        _ = v18
        _ = v19
    }
    _ = len(v5)
    v20 := GenS1{F0: 535, F1: 696, F2: 0}
    fmt.Println(v1, v2)
    v21 := GenS1{F0: 970, F1: 32767, F2: 0}
    _ = len(v3)
    v1 -= -94
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
