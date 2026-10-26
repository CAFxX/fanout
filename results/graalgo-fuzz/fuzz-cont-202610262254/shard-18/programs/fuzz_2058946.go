package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    var v1 uint8 = 42
    v2 := make(map[uint32]string)
    var v3 float32 = -2.25
    var v4 int64 = -521
    fmt.Println(v4)
    var v5 int32 = 1
    v6 := GenS1{F0: 4294967295, F1: 32767}
    switch "hello" {
    case "foo bar":
        if ((v1 == 7) || (1 >= 42)) {
            var v7 int8 = 25
            switch 2147483647 {
            case 4294967295:
                switch "" {
                case "\n\t":
                    fmt.Println(v1, v4, v3)
                case "":
                    _ = len(v2)
                case "a":
                    var v8 string = ""
                    _ = v8
                default:
                    v7 -= v7
                }
            case -187:
                if ((1.5 + 3.14159) > 1e10) {
                    var v9 string = ("\n\t" + ("foo bar" + "\n\t"))
                    _ = v9
                }
            default:
                v10 := []int{-2147483648, -9223372036854775808}
                _ = v10
            }
            _ = v7
        } else {
            var v11 uint = 18446744073709551615
            v1 = (v1 ^ v1)
            if true {
                fmt.Println(v1)
                for fi12 := 0; fi12 < 17; fi12++ {
                    _ = len(v2)
                    v2[30] = "\n\t"
                    v3 = v3
                    _ = fi12
                }
            } else {
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 uint16 = 0
                    switch fi13 {
                    case 2147483647:
                        for fi15 := 0; fi15 < 4; fi15++ {
                            v5 -= 2147483647
                            v1 += (v1 % ((v1) | 1))
                            _ = fi15
                        }
                    case -95:
                        switch 2147483647 {
                        case 435:
                            _ = len(v2)
                        case -9223372036854775808:
                            var v16 float64 = float64(v3)
                            _ = v16
                        }
                    case -2147483648:
                        var v17 uint8 = 255
                        _ = v17
                    default:
                        v3 = v3
                    }
                    _ = fi13
                    _ = v14
                }
                _ = len(v2)
            }
            _ = v11
        }
    }
    fmt.Println(v1, v4)
    if ((!(0 == 100)) && false) {
        v18 := GenS1{F0: 11, F1: -650}
        _ = v18
    }
    v5 = 42
    if false {
        var v19 uint64 = 4294967295
        v4 -= (^(v4))
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
