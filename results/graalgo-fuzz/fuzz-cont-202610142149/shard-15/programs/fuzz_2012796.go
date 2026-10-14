package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint16, p1 int32, p2 int) uint16 {
    v4 := GenS1{F0: p0}
    v5 := []int{p2}
    var v6 bool = true
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{}
    var v3 uint64 = 18446744073709551615
    if (42 == 1000) {
        if len(v2) > 0 {
            v2[0] = 2
        }
        var v7 uint32 = 2
        if (true || (!(!false))) {
            switch -121 {
            case -917:
                v7 -= (^(v7))
            case -1:
                v7 = v7
            case 4294967295:
                if len(v2) > 0 {
                    v2[0] = v7
                }
            default:
                if (1 >= 42) {
                    var v8 int64 = 190
                    v8 = 42
                    for fi9 := 0; fi9 < 7; fi9++ {
                        v1[(v7 * (v7 ^ (v7 / (((v7 | 2147483648)) | 1))))] = false
                        switch (fi9 >> ((812) & 63)) {
                        case 4294967295:
                            v10 := GenS2{F0: v8, F1: v7}
                            _ = v10
                        case 9223372036854775807:
                            v3 = v3
                        default:
                            if false {
                                switch "foo bar" {
                                case "foo bar":
                                    for fi11 := 0; fi11 < 20; fi11++ {
                                        var v12 int32 = 1
                                        _ = fi11
                                        _ = v12
                                    }
                                }
                                v7 = (v7 >> ((684) & 63))
                                var v13 uint64 = 999
                                _ = v13
                            }
                        }
                        _ = fi9
                    }
                    _ = v8
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = v7
    } else {
        v14 := v1[4294967295]
        if len(v2) > 0 {
            v2[0] = 596
        }
        _ = v14
    }
    var v15 int64 = -1
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 bool = (!((!(false && (2147483647 >= 752))) || (!(1 == 194))))
        _ = fi16
        _ = v17
    }
    v3 += v3
    switch 32767 {
    case 32767:
        fmt.Println(v3, v15)
    case 842:
        var v18 uint = 255
        _ = v18
    }
    v15 = (v15 & v15)
    v19 := v1[uint32(v3)]
    fmt.Println(v19, v3, v15)
    v19 = v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
}
