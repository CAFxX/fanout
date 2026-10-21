package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int8 }
func genFn0(p0 int32) uint {
    if (32767 < 2) {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 float64 = ((1e10 - 3.14159) * (float64(0.0) * -2.25))
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 int32 = p0
                _ = fi9
                _ = v10
            }
            v8 -= (v8 - 0.1)
            _ = fi7
            _ = v8
        }
        var v11 int64 = 9223372036854775807
        _ = v11
    }
    var v12 float64 = 1.5
    var v13 uint8 = 42
    fmt.Println(v13)
    var v14 bool = false
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 4294967295
    var v3 bool = false
    var v4 int = 780
    var v5 uint64 = 809
    v6 := make(map[string]uint32)
    v6[("a" + (("foo bar" + ("" + ("" + ""))) + ""))] = (v2 ^ v2)
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint16 = 823
        if v3 {
            var v17 int64 = 9223372036854775807
            for fi18 := 0; fi18 < 11; fi18++ {
                if (v16 < (v16 * 2)) {
                    var v19 uint32 = 814
                    v20 := GenS1{F0: v17, F1: -126, F2: 42}
                    fmt.Println(v3, v2, v5)
                    _ = v19
                    _ = v20
                } else {
                    delete(v1, (v4 & (fi18 | fi15)))
                    var v21 string = ""
                    switch (v4 / (((v4 / ((-1) | 1))) | 1)) {
                    case 2147483647:
                        var v22 uint = genFn0(-2147483648)
                        _ = v22
                    default:
                        var v23 uint64 = (v5 ^ v5)
                        _ = v23
                    }
                    _ = v21
                }
                switch "" {
                case "hello":
                    if v3 {
                        v2 += 367
                        v24 := []int{4294967295, -1, int(v17)}
                        _ = v24
                    }
                    var v25 int16 = int16(v17)
                    _ = v25
                }
                _ = fi18
            }
            var v26 int8 = 1
            _ = v17
            _ = v26
        }
        _ = len(v6)
        _ = fi15
        _ = v16
    }
    fmt.Println(v2, v4, v5)
    v27 := GenS1{F0: 0, F1: 127, F2: 0}
    if false {
        v3 = true
        v28 := []int{}
        _ = v28
    }
    var v29 uint32 = (v2 >> ((genFn0(-877)) & 63))
    switch "hello" {
    case "hello":
        var v30 int32 = -873
        _ = v30
    case "foo bar":
        for fi31 := 0; fi31 < 6; fi31++ {
            v3 = ((false || false) && (v3 || v3))
            var v32 float32 = (float32(((3.14159 - 3.14159) - 0.0)) - 0.0)
            _ = fi31
            _ = v32
        }
    default:
        v3 = (!(-2147483648 <= -527))
    }
    v33 := GenS1{F0: -1, F1: 127, F2: -55}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
