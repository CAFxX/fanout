package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int16 }
type GenS2 struct { F0 uint16; F1 string }
func genFn0(p0 int) string {
    var v5 int32 = 1
    v6 := GenS1{F0: (p0 / ((351) | 1)), F1: 1, F2: 42}
    _ = p0
    _ = v5
    _ = v6
    return ("\n\t" + "")
}

func main() {
    var v1 int32 = -494
    var v2 float32 = 1e10
    v3 := make(map[uint32]bool)
    var v4 int16 = 778
    var v7 uint32 = 553
    for fi8 := 0; fi8 < 10; fi8++ {
        v3[v7] = false
        var v9 float64 = ((float64(v2) / ((0.1 / 1e10) * 0.1)) - ((-2.25 / ((0.1 + float64(((-2.25 / (v2 - ((v2 - v2) / ((1e10 + v2) - v2)))) / v2))) + 0.1)) * 1.5))
        _ = fi8
        _ = v9
    }
    var v10 bool = (true || true)
    fmt.Println(v1)
    if (0.0 >= (v2 / v2)) {
        if (118 >= 133) {
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 float64 = 1.5
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 20; fi13++ {
                switch "a" {
                case "foo bar":
                    if (0 <= 887) {
                        var v14 string = "foo bar"
                        v14 = "a"
                        v2 = float32((0.1 / 1e10))
                        _ = v14
                    }
                case "hello":
                    v4 = v4
                default:
                    var v15 uint32 = (^(v7))
                    _ = v15
                }
                for fi16 := 0; fi16 < 11; fi16++ {
                    fmt.Println(v2, v10)
                    _ = fi16
                }
                _ = fi13
            }
            var v17 int16 = (v4 % (((v4 * v4)) | 1))
            _ = v17
        } else {
            var v18 uint8 = 128
            var v19 int = -784
            _ = v18
            _ = v19
        }
        _ = len(v3)
    }
    switch 1 {
    case 4294967295:
        var v20 uint8 = 76
        _ = v20
    case -9223372036854775808:
        var v21 int8 = 42
        _ = v21
    }
    switch -2147483648 {
    case 216:
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := GenS1{F0: fi22, F1: 1, F2: v4}
            _ = fi22
            _ = v23
        }
    case 42:
        var v24 int = -2147483648
        _ = v24
    case -840:
        v25 := GenS1{F0: -1, F1: 2, F2: (v4 & -1)}
        _ = v25
    }
    if true {
        for fi26 := 0; fi26 < 5; fi26++ {
            v27 := GenS1{F0: fi26, F1: 65535, F2: v4}
            switch "" {
            case "\n\t":
                if ((32767 > (v4 / ((-43) | 1))) && (((996 != 65535) && v10) || (2 != 4294967295))) {
                    var v28 uint16 = 32767
                    var v29 int64 = -9223372036854775808
                    _ = v28
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 8; fi30++ {
                        v1 -= 42
                        fmt.Println(v4)
                        _ = fi30
                    }
                }
            case "hello":
                v31 := GenS2{F0: 1, F1: "\n\t"}
                _ = v31
            case "foo bar":
                var v32 uint = 404
                _ = v32
            }
            _ = fi26
            _ = v27
        }
        var v33 uint = 65535
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
