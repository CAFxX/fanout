package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int32 }
type GenS2 struct { F0 uint16; F1 int8 }

func main() {
    var v1 string = ("hello" + ("hello" + ("a" + "")))
    v2 := make(map[int]uint32)
    var v3 uint16 = 255
    v4 := make(map[uint32]int)
    var v5 int = -42
    fmt.Println(v1)
    v5 -= v5
    v5 += 262
    v6 := []int{(v5 | v5)}
    v7 := v4[2147483647]
    v7 += 42
    v5 = 127
    if false {
        if true {
            switch v1 {
            case "hello":
                switch "a" {
                case "\n\t":
                    fmt.Println(v5)
                case "":
                    if true {
                        _ = len(v4)
                        var v8 uint16 = v3
                        delete(v2, (v5 + v7))
                        _ = v8
                    } else {
                        v1 = v1
                        if (!(!true)) {
                            if (-1 == 433) {
                                v9 := GenS1{F0: 65535, F1: v7, F2: -244}
                                var v10 uint = 2
                                _ = v10
                                _ = v9
                            }
                            if (false || false) {
                                var v11 float32 = (1e10 + 1.5)
                                if ((v5 << ((273) & 63)) != (-(v5))) {
                                    switch v7 {
                                    case 2147483647:
                                        switch (v1 + v1) {
                                        case "a":
                                            fmt.Println(v7, v1, v5)
                                        }
                                    case 1:
                                        _ = len(v6)
                                    case 32768:
                                        v5 = (v7 * (v5 & v7))
                                    default:
                                        switch (v1 + "hello") {
                                        case "foo bar":
                                            v12 := GenS1{F0: 58, F1: v7, F2: -966}
                                            _ = v12
                                        case "\n\t":
                                            v11 = float32(0.1)
                                        default:
                                            v3 -= v3
                                        }
                                    }
                                    switch v7 {
                                    case -1:
                                        v13 := GenS2{F0: 2, F1: 127}
                                        _ = v13
                                    default:
                                        var v14 bool = (100 == 100)
                                        _ = v14
                                    }
                                }
                                _ = v11
                            }
                        }
                        for fi15 := 0; fi15 < 15; fi15++ {
                            fmt.Println(fi15, fi15, v3)
                            _ = fi15
                        }
                    }
                }
            default:
                v16 := v4[4294967295]
                _ = v16
            }
        } else {
            var v17 float64 = 1.5
            _ = v17
        }
        v5 = v5
        v1 = v1
    }
    var v18 float64 = 0.0
    v19 := GenS2{F0: v3, F1: 127}
    fmt.Println(v5, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
