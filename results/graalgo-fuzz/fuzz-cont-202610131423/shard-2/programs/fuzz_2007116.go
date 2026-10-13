package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int; F1 int }
func genFn0() uint64 {
    var v4 float64 = float64((1e10 - -2.25))
    if false {
        v5 := GenS1{F0: 585, F1: 2}
        _ = v5
    }
    var v6 string = "\n\t"
    _ = v4
    _ = v6
    return 1
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[uint32]string)
    v3 := []uint32{2147483647, 335, v1, v1}
    for fi7 := 0; fi7 < 10; fi7++ {
        v3 = append(v3, v1)
        v1 = v1
        var v8 float64 = -2.25
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: 1}
    delete(v2, v1)
    v1 -= (v1 * (v1 & (v1 - (v1 + v1))))
    var v10 uint32 = 2
    v1 += (v1 >> ((255) & 63))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "a" {
    case "\n\t":
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    case "hello":
        var v11 int64 = -422
        _ = v11
    case "":
        switch 97 {
        case 792:
            v1 = v10
        case 127:
            var v12 uint32 = 0
            _ = v12
        case -126:
            if (false || false) {
                for fi13 := 0; fi13 < 14; fi13++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi13
                }
            } else {
                var v14 float64 = 0.1
                if false {
                    fmt.Println(v1)
                    v15 := GenS2{F0: 4294967295, F1: 953}
                    if false {
                        if true {
                            var v16 uint32 = (v1 - v10)
                            fmt.Println(v14, v14)
                            var v17 uint32 = (v10 / ((v16) | 1))
                            _ = v16
                            _ = v17
                        } else {
                            if true {
                                v10 = 713
                                for fi18 := 0; fi18 < 14; fi18++ {
                                    for fi19 := 0; fi19 < 20; fi19++ {
                                        v3 = append(v3, (v1 | v1))
                                        v20 := v2[(^(v1))]
                                        _ = fi19
                                        _ = v20
                                    }
                                    var v21 uint16 = 42
                                    _ = fi18
                                    _ = v21
                                }
                            }
                        }
                    }
                    _ = v15
                }
                _ = v14
            }
        }
    default:
        v10 = (v1 << ((64) & 63))
    }
    if (false && (false && ((!(42 >= 795)) || false))) {
        v10 -= v10
        v1 = v1
    }
    var v22 int = -636
    v1 = v10
    var v23 uint8 = 223
    fmt.Println(v22)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v24 uint64 = 926
    v23 = (v23 << ((916) & 63))
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        v3 = append(v3, v1)
    case "\n\t":
        v22 = v22
    }
    var v25 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v9
}
