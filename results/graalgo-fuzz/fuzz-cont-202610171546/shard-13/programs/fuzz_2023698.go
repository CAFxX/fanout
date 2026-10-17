package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 bool; F1 int64; F2 uint64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 0
    var v3 int32 = -651
    v4 := make(map[string]string)
    v5 := []int{127, 2147483647, 87, 1}
    delete(v1, 4294967295)
    fmt.Println(v3, v3)
    v6 := GenS1{F0: 0, F1: 157}
    v2 -= 65535
    switch 32767 {
    case 730:
        switch -2147483648 {
        case 32767:
            v2 -= 65535
        case 127:
            v2 += v2
        case 32768:
            fmt.Println(v2, v2)
        default:
            switch 42 {
            case 212:
                if true {
                    switch ("\n\t" + "a") {
                    case "hello":
                        if (false || (!(490 >= 32))) {
                            fmt.Println(v3)
                            for fi7 := 0; fi7 < 20; fi7++ {
                                switch "" {
                                case "":
                                    v3 = v3
                                }
                                _ = fi7
                            }
                            if (42 >= -128) {
                                v8 := GenS1{F0: 4294967295, F1: 255}
                                _ = v8
                            }
                        }
                    default:
                        var v9 int16 = -1
                        _ = v9
                    }
                    v2 += (v2 / ((v2) | 1))
                    var v10 int64 = 0
                    _ = v10
                }
            case 1:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case -928:
                var v11 uint = 2
                _ = v11
            default:
                delete(v1, -2147483648)
            }
        }
    case 2147483647:
        delete(v1, 2147483647)
    default:
        var v12 uint32 = 4294967295
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        if ((-2.25 != 1e10) || (!false)) {
            v3 = v3
            v14 := GenS2{F0: (false || (!true)), F1: 567, F2: 226}
            _ = v14
        }
        v2 -= v2
        _ = fi13
    }
    v3 -= (v3 & 2147483647)
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 20; fi16++ {
            _ = len(v5)
            _ = fi16
        }
        var v17 uint64 = 776
        _ = fi15
        _ = v17
    }
    v3 = v3
    if len(v5) > 0 {
        v5[0] = 871
    }
    var v18 uint64 = 943
    fmt.Println(v2, v2)
    var v19 float64 = float64(0.1)
    v3 -= 686
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := GenS2{F0: false, F1: 7, F2: v18}
        v19 += (v19 - v19)
        v18 -= (v18 & (^(v18)))
        _ = fi20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
