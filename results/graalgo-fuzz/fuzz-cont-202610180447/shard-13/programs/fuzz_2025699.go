package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint16 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int8 = 1
    v3 := make(map[int]string)
    if false {
        var v4 uint32 = 42
        _ = v4
    } else {
        v1 = 784
        var v5 int16 = -32768
        _ = v5
    }
    fmt.Println(v1)
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: v2, F1: 323, F2: 1}
            v2 -= v2
            fmt.Println(v1, v1, v2)
            var v9 uint64 = 0
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    delete(v3, 128)
    switch -9223372036854775808 {
    case 127:
        v10 := GenS1{F0: (v2 << ((4294967295) & 63)), F1: 2, F2: 1}
        _ = v10
    }
    var v11 float64 = (((3.14159 + 1.5) + 0.0) * 1.5)
    v2 = v2
    v12 := GenS1{F0: 127, F1: 230, F2: 1}
    var v13 int32 = 1
    var v14 uint64 = 42
    v11 = (v11 - (float64(1e10) + -2.25))
    v15 := GenS1{F0: (v2 % ((-1) | 1)), F1: 691, F2: 1}
    var v16 float64 = (((0.0 - 1e10) - 1.5) + v11)
    delete(v3, 0)
    v3[32767] = ("foo bar" + ((("" + "hello") + ("a" + "foo bar")) + ""))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
