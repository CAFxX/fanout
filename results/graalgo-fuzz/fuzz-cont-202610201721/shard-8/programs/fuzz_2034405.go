package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint, p2 int32) bool {
    var v5 uint8 = 1
    v6 := GenS1{F0: (p0 << (((p1 | (p1 / ((4294967295) | 1)))) & 63)), F1: 2, F2: "\n\t"}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (!true)
}

func main() {
    v1 := []uint32{490, 1, 2147483648}
    v2 := []uint32{886}
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 0
    fmt.Println(v4)
    v4 += (v4 ^ (v4 + 187))
    v4 += 155
    var v7 uint16 = 2
    v8 := GenS2{F0: 704}
    if false {
        v4 = 91
        var v9 uint16 = v7
        _ = v9
    } else {
        var v10 int8 = 50
        var v11 uint64 = 170
        var v12 uint = 255
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 32767}
    if ((!true) || (!(v4 >= 42))) {
        for fi14 := 0; fi14 < 12; fi14++ {
            if true {
                switch "\n\t" {
                case "foo bar":
                    switch "" {
                    case "":
                        var v15 bool = (!true)
                        _ = v15
                    case "\n\t":
                        for fi16 := 0; fi16 < 1; fi16++ {
                            _ = len(v3)
                            var v17 int32 = 1
                            _ = fi16
                            _ = v17
                        }
                    default:
                        for fi18 := 0; fi18 < 13; fi18++ {
                            var v19 uint16 = (v7 / ((v7) | 1))
                            _ = fi18
                            _ = v19
                        }
                    }
                case "a":
                    var v20 uint32 = 863
                    _ = v20
                }
                var v21 uint = 18446744073709551615
                v7 += v7
                _ = v21
            } else {
                var v22 int8 = 120
                fmt.Println(v22, fi14, v22)
                _ = v22
            }
            _ = fi14
        }
        var v23 int64 = 0
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v23
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 int64 = 1
        _ = fi24
        _ = v25
    }
    v26 := GenS1{F0: 44, F1: 1, F2: ("foo bar" + "foo bar")}
    fmt.Println(v7)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
