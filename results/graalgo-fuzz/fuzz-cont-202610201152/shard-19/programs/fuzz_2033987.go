package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 18446744073709551615
    var v3 uint16 = 65535
    var v4 int = 2147483647
    var v5 float64 = (0.1 + 0.1)
    v6 := []uint32{1}
    fmt.Println(v2)
    if len(v6) > 0 {
        v6[0] = 0
    }
    if true {
        if ((-1 >= -1) || false) {
            v7 := GenS1{F0: (-128 > 127), F1: -2147483648}
            _ = v7
        }
    } else {
        if true {
            if true {
                var v8 int8 = 1
                _ = v8
            } else {
                delete(v1, 0)
            }
            v9 := GenS1{F0: true, F1: 2147483647}
            v10 := GenS1{F0: true, F1: 1}
            _ = v10
            _ = v9
        }
        if len(v6) > 0 {
            v6[0] = 42
        }
        if false {
            v4 = (v4 | v4)
        } else {
            v6 = append(v6, 2147483647)
        }
    }
    var v11 uint16 = 42
    switch v4 {
    case -2147483648:
        var v12 int8 = 1
        _ = v12
    case -156:
        var v13 float64 = (v5 - v5)
        _ = v13
    case 128:
        for fi14 := 0; fi14 < 1; fi14++ {
            if false {
                v11 += (^(v11))
                if (32767 <= 42) {
                    v15 := GenS2{F0: 0, F1: v3, F2: 1}
                    _ = v15
                } else {
                    v16 := GenS1{F0: false, F1: 42}
                    var v17 uint = 991
                    _ = v16
                    _ = v17
                }
                for fi18 := 0; fi18 < 3; fi18++ {
                    v1[947] = "foo bar"
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    fmt.Println(v2)
                    v5 = v5
                    v3 = (^(v3))
                    _ = fi18
                }
            } else {
                var v19 int64 = -784
                v2 -= 18446744073709551615
                _ = v19
            }
            _ = fi14
        }
    }
    if false {
        var v20 int64 = 42
        _ = v20
    }
    _ = len(v6)
    v21 := GenS2{F0: 541, F1: v11, F2: -603}
    v5 -= v5
    var v22 uint16 = 1
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 int = (v4 >> ((42) & 63))
        for fi25 := 0; fi25 < 9; fi25++ {
            v26 := GenS1{F0: (4294967295 != 42), F1: 1}
            v5 = ((v5 / v5) + 3.14159)
            v27 := GenS3{F0: -1, F1: (^(v3))}
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        v29 := GenS4{F0: -24}
        _ = fi28
        _ = v29
    }
    v6 = append(v6, 905)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
