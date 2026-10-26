package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]bool)
    v3 := []string{"", "hello", ""}
    var v4 int64 = -516
    v5 := make(map[string]string)
    var v6 int8 = 100
    var v7 int8 = v6
    for fi8 := 0; fi8 < 18; fi8++ {
        if len(v3) > 0 {
            v3[0] = ("foo bar" + "hello")
        }
        _ = fi8
    }
    var v9 uint8 = 255
    _ = len(v3)
    if false {
        _ = len(v3)
        v6 = (v6 * -128)
    } else {
        v6 -= 127
        v9 = (v9 ^ v9)
    }
    v9 = (v9 / (((^(v9))) | 1))
    v7 = (v6 * (v6 | v7))
    if false {
        if true {
            var v10 int8 = v6
            var v11 int32 = v1
            if true {
                switch 42 {
                case 649:
                    if true {
                        var v12 uint32 = 1
                        var v13 bool = false
                        var v14 uint8 = 1
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 3; fi15++ {
                            var v16 uint = 196
                            v17 := GenS1{F0: 65535, F1: (v4 * int64(fi15))}
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    }
                default:
                    v18 := v2[22]
                    _ = v18
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v10
            _ = v11
        }
    } else {
        var v19 int = 730
        _ = v19
    }
    switch -9223372036854775808 {
    case 32768:
        v20 := GenS1{F0: 255, F1: v4}
        _ = v20
    case 0:
        v4 = v4
    default:
        var v21 uint64 = 2
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
