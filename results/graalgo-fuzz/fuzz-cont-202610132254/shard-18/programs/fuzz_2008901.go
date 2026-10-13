package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int }
func genFn0() int {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint64 = 2
        if (true && true) {
            v8 += 2
            v8 = v8
            var v9 int64 = -590
            _ = v9
        } else {
            v10 := GenS1{F0: -9223372036854775808, F1: 4294967295, F2: (fi7 | (fi7 / (((fi7 / ((fi7) | 1))) | 1)))}
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    var v11 int = 32768
    var v12 float32 = 0.1
    var v13 int16 = 1000
    _ = v11
    _ = v12
    _ = v13
    return 32767
}

func main() {
    v1 := []string{"a", "", ""}
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    var v4 uint16 = 32767
    var v5 uint32 = 42
    var v6 uint32 = v5
    switch ("\n\t" + ("a" + "\n\t")) {
    case "foo bar":
        for fi14 := 0; fi14 < 3; fi14++ {
            v5 += (v5 / ((v5) | 1))
            var v15 uint8 = 255
            _ = fi14
            _ = v15
        }
    case "\n\t":
        var v16 uint16 = 1
        _ = v16
    case "a":
        for fi17 := 0; fi17 < 11; fi17++ {
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 uint8 = 127
                v6 = (v6 + v5)
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    default:
        var v20 int8 = 1
        _ = v20
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 uint8 = uint8(v4)
        var v23 float64 = 3.14159
        for fi24 := 0; fi24 < 18; fi24++ {
            switch "hello" {
            case "\n\t":
                v6 -= (^(v5))
            default:
                var v25 int32 = -2147483648
                _ = v25
            }
            _ = fi24
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    if true {
        if false {
            for fi26 := 0; fi26 < 12; fi26++ {
                switch "foo bar" {
                case "foo bar":
                    v27 := v3[(v5 - v5)]
                    _ = v27
                case "\n\t":
                    var v28 uint64 = 486
                    _ = v28
                }
                v29 := GenS1{F0: 42, F1: (fi26 % ((genFn0()) | 1)), F2: genFn0()}
                _ = fi26
                _ = v29
            }
        }
    }
    v5 += 65535
    for fi30 := 0; fi30 < 4; fi30++ {
        fmt.Println(v6, v5)
        _ = fi30
    }
    _ = len(v2)
    delete(v2, 128)
    delete(v2, genFn0())
    var v31 bool = true
    v31 = false
    var v32 bool = ((v4 ^ v4) >= (v4 << ((2) & 63)))
    fmt.Println(v6)
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
