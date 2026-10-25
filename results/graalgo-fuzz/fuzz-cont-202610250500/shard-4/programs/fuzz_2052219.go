package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint64) string {
    var v4 float64 = 0.0
    var v5 uint16 = 203
    v4 += float64(-2.25)
    switch 42 {
    case 128:
        v6 := GenS1{F0: 42, F1: -116, F2: 1}
        _ = v6
    default:
        v5 -= (^(v5))
    }
    _ = p0
    _ = v4
    _ = v5
    return ("" + ("foo bar" + ("hello" + "hello")))
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 255
    v3 := []uint32{405, 0, 4294967295, 89}
    if false {
        if true {
            if (0 != -10) {
                if len(v3) > 0 {
                    v3[0] = 65535
                }
            } else {
                switch ("foo bar" + "hello") {
                case "":
                    for fi7 := 0; fi7 < 18; fi7++ {
                        if (42 > 65535) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        var v8 int64 = int64(fi7)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi7
                        _ = v8
                    }
                case "\n\t":
                    if len(v3) > 0 {
                        v3[0] = 65535
                    }
                case "foo bar":
                    var v9 float64 = (float64(0.1) - 0.1)
                    _ = v9
                }
                var v10 uint8 = 42
                _ = v10
            }
            v1 = append(v1, 2147483647)
            v11 := GenS2{F0: 100, F1: 929}
            _ = v11
        } else {
            v2 = (v2 >> ((180) & 63))
            v2 = v2
            v2 = 2
        }
    } else {
        var v12 int16 = 885
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi13
    }
    var v14 int64 = -785
    v14 = 0
    v15 := GenS2{F0: -13, F1: 18446744073709551615}
    fmt.Println(v14, v14)
    v16 := GenS1{F0: (^(v2)), F1: 127, F2: 1}
    var v17 uint8 = 42
    v18 := GenS1{F0: (v2 << ((255) & 63)), F1: 100, F2: 18446744073709551615}
    v14 = (v14 << ((0) & 63))
    var v19 uint64 = 4294967295
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v20 int32 = -917
    if true {
        var v21 int64 = (v14 << ((255) & 63))
        _ = v21
    }
    v19 = v19
    v17 -= 1
    v19 = v19
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
