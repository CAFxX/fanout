package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int32, p1 int32, p2 string) uint64 {
    fmt.Println(p2)
    v6 := GenS1{F0: (^(p0)), F1: -128}
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 string = ("\n\t" + p2)
        p2 = v8
        _ = fi7
        _ = v8
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 int64 = 1
    var v2 string = ""
    v3 := []uint32{4294967295}
    var v4 uint16 = 32767
    var v5 int64 = (^(v1))
    if (114 >= 42) {
        v3 = append(v3, 65535)
    }
    v9 := GenS1{F0: 42, F1: -1}
    var v10 uint8 = 42
    var v11 uint16 = v4
    switch -527 {
    case 127:
        v12 := GenS2{F0: 0, F1: 385}
        _ = v12
    case 2147483647:
        v2 = v2
    case 32768:
        v2 = v2
    default:
        v2 = (v2 + "foo bar")
    }
    v4 -= (v4 * 42)
    var v13 float64 = 1e10
    if true {
        fmt.Println(v5, v4)
        v3 = append(v3, 2147483647)
        v11 = (v4 | (^(v11)))
    } else {
        if (2147483647 <= 0) {
            if false {
                var v14 uint16 = (v4 + (v11 ^ (v4 - (v4 / (((^(v11))) | 1)))))
                _ = len(v3)
                v11 = v11
                _ = v14
            } else {
                for fi15 := 0; fi15 < 17; fi15++ {
                    for fi16 := 0; fi16 < 11; fi16++ {
                        var v17 uint16 = (^(v4))
                        v18 := GenS1{F0: 2147483647, F1: 73}
                        var v19 uint8 = (v10 - v10)
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = fi15
                }
            }
            if true {
                v5 = (v5 - v1)
                for fi20 := 0; fi20 < 10; fi20++ {
                    v11 = 94
                    var v21 int32 = -2147483648
                    var v22 uint64 = 2
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
        } else {
            var v23 uint16 = v11
            v5 = (v1 >> ((741) & 63))
            var v24 uint64 = 42
            _ = v23
            _ = v24
        }
        var v25 uint64 = 922
        _ = v25
    }
    v5 = (v1 << ((42) & 63))
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 1; fi27++ {
            v11 += v4
            v28 := GenS3{F0: (^(v4))}
            _ = fi27
            _ = v28
        }
        if (true && ((-1 >= (v5 >> ((477) & 63))) && true)) {
            v11 = (v11 >> ((107) & 63))
            v2 = (v2 + v2)
        } else {
            v11 += v11
        }
        _ = fi26
    }
    v10 += (v10 << ((1) & 63))
    var v29 int16 = int16(v10)
    v4 -= (v11 % ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
