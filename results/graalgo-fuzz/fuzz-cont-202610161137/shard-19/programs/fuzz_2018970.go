package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint, p1 uint) uint32 {
    v5 := GenS1{F0: 1}
    switch "\n\t" {
    case "a":
        p0 += 4294967295
    case "foo bar":
        p1 = 1
    default:
        p0 += (^(p1))
    }
    _ = p0
    _ = p1
    _ = v5
    return 2
}

func main() {
    var v1 float64 = (0.0 / 1.0)
    v2 := make(map[int]uint32)
    var v3 int = 1
    var v4 string = ("a" + "\n\t")
    v1 = 1.5
    var v6 int16 = -726
    var v7 int16 = 1
    v8 := GenS2{F0: 65535, F1: 474}
    v9 := GenS2{F0: 204, F1: 2}
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int16 = (-(v6))
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v2[(v3 / (((-(v3))) | 1))] = 4294967295
    if true {
        switch v3 {
        case -39:
            v3 += v3
        }
        v13 := []int{127, 541, (v3 << ((65535) & 63))}
        _ = v13
    } else {
        var v14 uint8 = 200
        for fi15 := 0; fi15 < 20; fi15++ {
            if (v4 > "foo bar") {
                switch 33 {
                case -412:
                    switch "" {
                    case "foo bar":
                        if (89 < 100) {
                            for fi16 := 0; fi16 < 9; fi16++ {
                                v7 -= (^(v6))
                                var v17 uint8 = (v14 * v14)
                                var v18 int = int(v14)
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            var v19 bool = false
                            _ = v19
                        } else {
                            var v20 int32 = 2147483647
                            v7 = v7
                            if true {
                                var v21 int = -315
                                for fi22 := 0; fi22 < 13; fi22++ {
                                    var v23 string = v4
                                    fmt.Println(v6)
                                    if false {
                                        v14 -= (^(v14))
                                    } else {
                                        var v24 uint32 = 2
                                        _ = v24
                                    }
                                    _ = fi22
                                    _ = v23
                                }
                                v20 -= (v20 / ((v20) | 1))
                                _ = v21
                            } else {
                                var v25 uint64 = 420
                                _ = v25
                            }
                            _ = v20
                        }
                    }
                default:
                    var v26 int = v3
                    _ = v26
                }
                v27 := v2[fi15]
                for fi28 := 0; fi28 < 2; fi28++ {
                    v14 = v14
                    var v29 int8 = 100
                    var v30 uint32 = v27
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = v27
            }
            switch v4 {
            case "\n\t":
                switch (fi15 | (v3 * (v3 | (v3 * (fi15 - fi15))))) {
                case 128:
                    fmt.Println(v1, fi15)
                case 9223372036854775807:
                    fmt.Println(v4, v1, v14)
                }
            case "hello":
                v14 += (v14 * v14)
            }
            v31 := GenS2{F0: 887, F1: 441}
            _ = fi15
            _ = v31
        }
        var v32 uint16 = 692
        _ = v14
        _ = v32
    }
    fmt.Println(v1)
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
