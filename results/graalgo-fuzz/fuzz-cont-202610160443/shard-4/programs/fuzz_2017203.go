package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS4 struct { F0 bool; F1 int32; F2 string }
func genFn0() uint {
    if (true || true) {
        var v7 int8 = 127
        _ = v7
    }
    if ((1 > 125) && true) {
        var v8 int8 = 42
        var v9 uint64 = 44
        _ = v8
        _ = v9
    }
    switch "a" {
    case "hello":
        v10 := []int{325, 42, 42}
        _ = v10
    default:
        var v11 float32 = (0.1 - ((1e10 * ((0.1 + 0.1) * 1.5)) - (1e10 - 1.5)))
        _ = v11
    }
    return 4294967295
}

func main() {
    var v1 uint8 = 127
    v2 := []uint32{348, 590, 879, 148}
    v3 := []uint32{2147483648, 660}
    var v4 uint8 = v1
    v5 := []string{"a"}
    var v6 int8 = -71
    for fi12 := 0; fi12 < 5; fi12++ {
        if true {
            v6 = v6
            for fi13 := 0; fi13 < 14; fi13++ {
                v6 = v6
                var v14 int64 = 42
                _ = fi13
                _ = v14
            }
            var v15 bool = ((1.5 / 1e10) < 0.1)
            _ = v15
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            fmt.Println(v1, fi12, fi16)
            _ = fi16
        }
        _ = fi12
    }
    var v17 uint8 = (^(v1))
    _ = len(v5)
    v18 := GenS1{F0: -1, F1: 2147483647, F2: 127}
    fmt.Println(v6, v1)
    for fi19 := 0; fi19 < 5; fi19++ {
        v6 -= (v6 >> ((65535) & 63))
        if false {
            fmt.Println(v6, v1)
        }
        _ = fi19
    }
    if true {
        if (("hello" >= (("\n\t" + ("foo bar" + "a")) + "foo bar")) && (!false)) {
            v20 := GenS2{F0: -207, F1: 643}
            var v21 int16 = -227
            fmt.Println(v1, v1, v4)
            _ = v20
            _ = v21
        } else {
            var v22 uint = 1
            _ = v22
        }
    }
    var v23 uint16 = 1
    for fi24 := 0; fi24 < 2; fi24++ {
        fmt.Println(v1, v17, fi24)
        switch "\n\t" {
        case "\n\t":
            for fi25 := 0; fi25 < 3; fi25++ {
                fmt.Println(v6)
                _ = fi25
            }
        default:
            var v26 string = ("\n\t" + "")
            _ = v26
        }
        if len(v3) > 0 {
            v3[0] = 396
        }
        _ = fi24
    }
    if true {
        if false {
            var v27 uint8 = (v4 << ((65535) & 63))
            v27 = 255
            _ = v27
        }
        v17 -= (v4 % ((128) | 1))
        v6 += (^(v6))
    } else {
        fmt.Println(v17, v23, v23)
        v28 := GenS3{F0: (v4 << ((0) & 63)), F1: 221, F2: -835}
        v29 := GenS4{F0: (1 > 2147483647), F1: 2147483647, F2: ("foo bar" + "\n\t")}
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
