package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 bool; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int = -9223372036854775808
    var v2 uint64 = 4294967295
    var v3 int32 = -972
    var v4 int16 = 1
    var v5 int64 = 42
    var v6 string = "hello"
    if true {
        v7 := GenS1{F0: -9223372036854775808, F1: (v5 - v5)}
        fmt.Println(v4)
        _ = v7
    }
    switch v6 {
    case "":
        v8 := []int{v1, (-(v1)), 2147483647, v1}
        _ = v8
    }
    fmt.Println(v4, v5, v1)
    for fi9 := 0; fi9 < 1; fi9++ {
        if (false || (42 == 65535)) {
            v10 := GenS2{F0: false, F1: v5, F2: v2}
            _ = v10
        } else {
            v5 += (v5 / ((v5) | 1))
        }
        v2 = (^(v2))
        _ = fi9
    }
    v1 += (v1 ^ v1)
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS1{F0: (fi11 - 32768), F1: v5}
        var v13 uint16 = 279
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 bool = true
    v14 = ((3.14159 < 3.14159) || v14)
    v3 = (v3 >> ((777) & 63))
    v3 -= -1
    v15 := []int{(v1 * -9223372036854775808)}
    if v14 {
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := GenS1{F0: (fi16 * (v1 * v1)), F1: -9223372036854775808}
            v18 := GenS3{F0: 0}
            _ = fi16
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 10; fi19++ {
            switch (fi19 / ((fi19) | 1)) {
            case 128:
                v20 := GenS2{F0: v14, F1: (^(v5)), F2: 2}
                _ = v20
            case -703:
                var v21 uint64 = (v2 + v2)
                _ = v21
            case 32768:
                v2 -= 18446744073709551615
            default:
                v14 = (3.14159 >= 0.0)
            }
            if len(v15) > 0 {
                v15[0] = 32767
            }
            _ = fi19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
