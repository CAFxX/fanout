package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint16 = 1
    v2 := make(map[int]int)
    var v3 float32 = 1e10
    v4 := make(map[uint32]string)
    var v5 bool = (0 < 255)
    v6 := make(map[string]string)
    v7 := []int{}
    delete(v2, -1)
    v3 = (v3 * v3)
    v8 := GenS1{F0: 427}
    v9 := GenS1{F0: 0}
    v1 += v1
    switch 127 {
    case -9223372036854775808:
        v3 = 0.1
    case 0:
        var v10 string = ("\n\t" + "\n\t")
        _ = v10
    default:
        v11 := v6["foo bar"]
        _ = v11
    }
    v3 += 3.14159
    var v12 int64 = -1
    var v13 uint8 = 127
    v5 = v5
    switch ("foo bar" + "hello") {
    case "":
        if true {
            for fi14 := 0; fi14 < 1; fi14++ {
                switch "" {
                case "":
                    delete(v6, "hello")
                case "\n\t":
                    for fi15 := 0; fi15 < 17; fi15++ {
                        if ((v3 / v3) < 1.5) {
                            var v16 int16 = -1
                            fmt.Println(fi15)
                            v1 = 42
                            _ = v16
                        } else {
                            v13 = v13
                        }
                        if v5 {
                            fmt.Println(fi14, v3, fi14)
                            var v17 uint64 = 325
                            _ = v17
                        }
                        var v18 int16 = 1
                        if len(v7) > 0 {
                            v7[0] = (fi14 * fi14)
                        }
                        _ = fi15
                        _ = v18
                    }
                default:
                    for fi19 := 0; fi19 < 14; fi19++ {
                        var v20 int8 = 42
                        _ = fi19
                        _ = v20
                    }
                }
                v21 := GenS1{F0: 631}
                _ = fi14
                _ = v21
            }
            v7 = append(v7, 32768)
        }
    case "foo bar":
        v22 := GenS1{F0: 682}
        _ = v22
    case "a":
        if (1 != -2147483648) {
            for fi23 := 0; fi23 < 8; fi23++ {
                var v24 bool = ((v13 % ((v13) | 1)) == (v13 >> ((439) & 63)))
                v13 += (v13 + 127)
                for fi25 := 0; fi25 < 10; fi25++ {
                    var v26 int32 = -2147483648
                    if v5 {
                        v27 := v4[2147483647]
                        v28 := GenS1{F0: 2}
                        _ = v27
                        _ = v28
                    } else {
                        switch "" {
                        case "hello":
                            var v29 uint64 = 18446744073709551615
                            _ = v29
                        case "":
                            v5 = v24
                        case "a":
                            if false {
                                v3 += (0.0 + v3)
                            }
                        }
                    }
                    _ = fi25
                    _ = v26
                }
                _ = fi23
                _ = v24
            }
            v5 = ((-2.25 + -2.25) > (float64(v3) - (-2.25 / (((-2.25 + float64((v3 - (v3 / (0.0 - v3))))) / 1.0) + 0.0))))
        }
    default:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 int16 = -240
        _ = fi30
        _ = v31
    }
    v7 = append(v7, 42)
    for fi32 := 0; fi32 < 10; fi32++ {
        v5 = (!(v5 || (!(-2147483648 > 1))))
        _ = fi32
    }
    v33 := v2[128]
    if v5 {
        var v34 float64 = 3.14159
        v35 := GenS2{F0: 906, F1: (2147483648 == 42)}
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
