package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 4294967295
    v3 := make(map[uint32]uint32)
    var v4 bool = true
    var v5 uint = 808
    var v6 uint64 = (v2 * (^(v2)))
    v5 += 255
    for fi7 := 0; fi7 < 8; fi7++ {
        if false {
            fmt.Println(v4, v4)
        }
        _ = fi7
    }
    var v8 int8 = -128
    fmt.Println(v4, v4, v5)
    var v9 int = 2147483647
    var v10 int8 = 1
    if (3.14159 <= (-2.25 / 1.5)) {
        var v11 uint32 = 1
        if ((("" + "foo bar") + ((("" + ((("hello" + "") + "hello") + "foo bar")) + ((("\n\t" + "") + "hello") + "hello")) + "foo bar")) <= (("" + "hello") + ("a" + ("a" + ("a" + ("foo bar" + "\n\t")))))) {
            var v12 int8 = int8(v9)
            switch ("hello" + "hello") {
            case "foo bar":
                for fi13 := 0; fi13 < 11; fi13++ {
                    var v14 uint64 = (v2 ^ v6)
                    _ = fi13
                    _ = v14
                }
            case "\n\t":
                v15 := []int{int(v5), 42, (v9 - (v9 | v9))}
                _ = v15
            default:
                v2 = (^(v6))
            }
            _ = v12
        } else {
            v16 := GenS1{F0: v4, F1: -2147483648, F2: 1000}
            _ = v16
        }
        _ = v11
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        if (0 != 654) {
            fmt.Println(v8, v6, v6)
            if v4 {
                switch ("foo bar" + ("hello" + "hello")) {
                case "\n\t":
                    v10 = (v8 << ((v5) & 63))
                case "hello":
                    v18 := GenS1{F0: v4, F1: 550, F2: 1}
                    _ = v18
                case "":
                    v19 := GenS1{F0: ((4294967295 < 2) || v4), F1: -1, F2: -1}
                    _ = v19
                }
                for fi20 := 0; fi20 < 12; fi20++ {
                    if v4 {
                        var v21 int32 = 0
                        var v22 uint64 = 18446744073709551615
                        v23 := GenS1{F0: v4, F1: v21, F2: 32767}
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        var v24 float64 = 1.5
                        _ = v24
                    }
                    var v25 int32 = 42
                    _ = fi20
                    _ = v25
                }
                v26 := GenS1{F0: (v8 != (v10 * (v10 + v10))), F1: -422, F2: 498}
                _ = v26
            } else {
                v8 = (^(v10))
                var v27 uint64 = 1
                v28 := GenS1{F0: true, F1: -62, F2: 1}
                _ = v27
                _ = v28
            }
        }
        fmt.Println(v2, v6, v10)
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
