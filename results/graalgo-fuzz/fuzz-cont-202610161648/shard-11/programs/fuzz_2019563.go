package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int16 = 32767
        fmt.Println(v8, fi7)
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 608}
    _ = v9
    return 127
}

func main() {
    v1 := []uint32{0, 2147483647, 2147483648, 42}
    var v2 uint = 1
    v3 := make(map[string]bool)
    var v4 uint32 = 128
    v5 := []uint32{uint32(v2), (v4 % ((v4) | 1))}
    var v6 uint16 = 65535
    var v10 int32 = 2147483647
    var v11 uint32 = 1
    v6 = 707
    switch ("a" + "hello") {
    case "foo bar":
        v12 := GenS1{F0: 18446744073709551615}
        _ = v12
    default:
        var v13 bool = (float32((float64((1.5 - float32(3.14159))) - 3.14159)) >= float32(1.5))
        _ = v13
    }
    var v14 uint64 = 0
    v6 -= v6
    fmt.Println(v4, v10, v11)
    fmt.Println(v2, v6)
    v15 := GenS1{F0: (v14 - (^(v14)))}
    for fi16 := 0; fi16 < 14; fi16++ {
        v14 = 440
        var v17 int = (fi16 - -1)
        _ = fi16
        _ = v17
    }
    v14 -= v14
    v6 += 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
