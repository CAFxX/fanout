package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint; F1 int16; F2 uint8 }
func genFn0() int64 {
    var v6 int32 = 42
    v6 -= v6
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int16 = 1000
        var v9 bool = true
        var v10 uint8 = 255
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := GenS1{F0: 2147483648, F1: -1}
    _ = v11
    _ = v6
    return 313
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint8 = 168
    v3 := make(map[uint32]string)
    v4 := []uint32{2147483647, 891, 2147483647}
    var v5 bool = (255 <= 2)
    v5 = true
    v12 := GenS1{F0: 2, F1: genFn0()}
    var v13 int32 = -626
    for fi14 := 0; fi14 < 13; fi14++ {
        fmt.Println(fi14, v5, v2)
        var v15 uint8 = v2
        var v16 int = 0
        _ = fi14
        _ = v15
        _ = v16
    }
    if (253 != (^(v13))) {
        v13 = (-(v13))
    }
    v1 = v1
    v1 = v1
    fmt.Println(v2, v1)
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := GenS1{F0: 699, F1: 1}
        var v19 uint16 = 779
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = len(v4)
    v1 -= 0
    switch ("\n\t" + ("foo bar" + ("a" + ""))) {
    case "\n\t":
        fmt.Println(v2, v13, v2)
    case "foo bar":
        v13 += v13
    case "hello":
        if (v13 >= v13) {
            var v20 bool = v5
            var v21 bool = false
            var v22 int32 = 1
            _ = v20
            _ = v21
            _ = v22
        }
    }
    if (v13 == -2147483648) {
        if ((v13 / ((-525) | 1)) == -825) {
            var v23 uint = 677
            v24 := GenS2{F0: v23, F1: 1000, F2: (v2 % ((v2) | 1))}
            _ = v23
            _ = v24
        }
    }
    v1 = v1
    fmt.Println(v5, v5, v5)
    var v25 int8 = 42
    fmt.Println(v2, v13, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
