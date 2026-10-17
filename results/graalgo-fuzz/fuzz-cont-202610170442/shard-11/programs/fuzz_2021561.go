package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int64 }
func genFn0() uint64 {
    v4 := []int{2147483647, 128, 66}
    var v5 uint32 = 2147483647
    var v6 int64 = 9223372036854775807
    switch "hello" {
    case "":
        for fi7 := 0; fi7 < 15; fi7++ {
            switch ("a" + "\n\t") {
            case "hello":
                for fi8 := 0; fi8 < 15; fi8++ {
                    var v9 bool = (!false)
                    if (false || (0.1 == -2.25)) {
                        var v10 int64 = v6
                        _ = v10
                    } else {
                        switch "" {
                        case "":
                            if v9 {
                                switch 575 {
                                case 0:
                                    v6 += (v6 * -185)
                                case 311:
                                    fmt.Println(fi7)
                                case 1:
                                    var v11 uint16 = 0
                                    _ = v11
                                }
                            }
                        case "hello":
                            var v12 uint8 = 1
                            _ = v12
                        }
                    }
                    _ = fi8
                    _ = v9
                }
            case "\n\t":
                for fi13 := 0; fi13 < 1; fi13++ {
                    if (90 > -120) {
                        fmt.Println(fi7)
                        v6 -= 1
                    }
                    _ = fi13
                }
            case "":
                if false {
                    v5 -= (v5 & 2)
                } else {
                    if false {
                        var v14 bool = (23 == 1)
                        v4 = append(v4, fi7)
                        fmt.Println(v6)
                        _ = v14
                    }
                    v5 = 2147483647
                }
            default:
                v6 -= v6
            }
            _ = fi7
        }
    default:
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 uint32 = (v5 + v5)
            _ = fi15
            _ = v16
        }
    }
    _ = v4
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -166
    v3 := make(map[int]int)
    v2 = v2
    v17 := GenS1{F0: 1, F1: 42, F2: 1}
    v18 := GenS1{F0: -1, F1: 210, F2: 431}
    var v19 int8 = 0
    delete(v1, (("" + ("" + "hello")) + "foo bar"))
    var v20 uint16 = 65535
    switch 32768 {
    case 0:
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 int16 = 309
            _ = fi21
            _ = v22
        }
    case 42:
        v20 -= 350
    }
    v19 -= (v19 >> ((65535) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
