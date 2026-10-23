package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 int32 = -1
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    v4 := []int{127}
    var v5 bool = (1 >= 9223372036854775807)
    var v6 int16 = -32768
    v7 := v2[""]
    if (1 < 352) {
        v6 = v6
    }
    var v8 int = -1
    if (v5 && v5) {
        v9 := v3[727]
        _ = v9
    } else {
        fmt.Println(v6, v7, v1)
        var v10 int = v8
        _ = v10
    }
    _ = len(v4)
    v11 := GenS1{F0: v8}
    v8 += (^(v8))
    var v12 float64 = 0.1
    v3[58] = "a"
    v1 = v1
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS1{F0: (fi13 % (((fi13 / ((42) | 1))) | 1))}
        if (((int(v6) > fi13) || v5) && true) {
            v5 = ((v6 - v6) != int16(fi13))
            for fi15 := 0; fi15 < 14; fi15++ {
                if (true || (18446744073709551615 == 329)) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if false {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v16 uint = 1
                        v5 = ((!true) && v5)
                        _ = v16
                    }
                }
                if ((v6 ^ (v6 ^ (v6 % ((42) | 1)))) != 1000) {
                    var v17 int = (fi13 * (-(v8)))
                    v7 = v7
                    if len(v4) > 0 {
                        v4[0] = v8
                    }
                    _ = v17
                }
                var v18 int = v8
                _ = fi15
                _ = v18
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                v5 = v5
                _ = fi19
            }
        } else {
            v4 = append(v4, v8)
            v2[v7] = "hello"
            switch (fi13 | v8) {
            case 32767:
                if (722 < 590) {
                    var v20 float64 = (v12 * (0.0 * ((v12 * 1.5) - 0.0)))
                    fmt.Println(v7, v8)
                    _ = v20
                }
            default:
                var v21 uint64 = 747
                _ = v21
            }
        }
        _ = fi13
        _ = v14
    }
    if (42 > 24) {
        if v5 {
            v12 += 0.1
            var v22 int64 = -985
            _ = v22
        }
    } else {
        var v23 bool = v5
        v24 := GenS1{F0: (v8 + 9223372036854775807)}
        _ = v23
        _ = v24
    }
    if ((false || ((1e10 / 1.0) <= v12)) || (4294967295 < 65535)) {
        v1 += v1
    } else {
        for fi25 := 0; fi25 < 3; fi25++ {
            fmt.Println(v5, v8)
            for fi26 := 0; fi26 < 20; fi26++ {
                v2[v7] = ""
                var v27 uint8 = 58
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
        var v28 int = (v8 >> ((1) & 63))
        v12 += v12
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
