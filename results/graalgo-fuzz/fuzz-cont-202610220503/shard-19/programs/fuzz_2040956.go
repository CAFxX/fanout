package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]uint32)
    var v4 uint64 = 2
    v5 := []uint32{65535, 291}
    var v6 int16 = -32768
    switch 4294967295 {
    case 1:
        fmt.Println(v6)
    default:
        v6 += (-(v6))
    }
    v6 = (v6 | v6)
    v1[381] = -497
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := v3["\n\t"]
        _ = fi7
        _ = v8
    }
    var v9 int64 = 9223372036854775807
    var v10 bool = true
    switch -1 {
    case 4294967295:
        var v11 uint64 = v4
        _ = v11
    case -881:
        if v10 {
            switch 4294967295 {
            case 0:
                for fi12 := 0; fi12 < 19; fi12++ {
                    fmt.Println(v9)
                    _ = fi12
                }
            case 937:
                switch "a" {
                case "foo bar":
                    for fi13 := 0; fi13 < 10; fi13++ {
                        _ = len(v1)
                        _ = fi13
                    }
                case "\n\t":
                    var v14 uint8 = 0
                    _ = v14
                }
            default:
                switch "hello" {
                case "\n\t":
                    switch -9223372036854775808 {
                    case 9223372036854775807:
                        if false {
                            var v15 uint = 2
                            _ = v15
                        }
                    case -409:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                default:
                    v6 += (^(v6))
                }
            }
        } else {
            var v16 bool = v10
            _ = v16
        }
    default:
        if (v10 && true) {
            if v10 {
                v2[0] = (127 < 935)
                v1[4294967295] = 110
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 uint8 = 1
                    v9 -= -9223372036854775808
                    _ = fi17
                    _ = v18
                }
            } else {
                var v19 int = 1
                _ = v19
            }
            v20 := GenS1{F0: 87, F1: 0, F2: 1}
            _ = v20
        } else {
            delete(v3, "hello")
            switch "" {
            case "a":
                v4 = (v4 * v4)
            case "\n\t":
                v4 += (v4 << ((255) & 63))
            case "hello":
                v21 := GenS1{F0: -1, F1: 42, F2: -128}
                _ = v21
            }
            if ((1.5 * (1.5 * 0.0)) == 0.0) {
                for fi22 := 0; fi22 < 18; fi22++ {
                    fmt.Println(v10, v9, fi22)
                    v23 := GenS1{F0: -1, F1: 686, F2: 0}
                    _ = fi22
                    _ = v23
                }
                var v24 uint8 = 166
                _ = v24
            }
        }
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        fmt.Println(v9)
        _ = fi25
    }
    switch 32768 {
    case 448:
        v9 += 0
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
