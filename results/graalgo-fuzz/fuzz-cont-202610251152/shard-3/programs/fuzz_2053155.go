package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint = 987
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    var v4 uint8 = 2
    var v5 float32 = (float32(0.1) / 1.0)
    v6 := make(map[int]bool)
    v7 := GenS1{F0: (v4 << (((v1 ^ 255)) & 63))}
    v3 = append(v3, 4294967295)
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 bool = true
        _ = fi8
        _ = v9
    }
    if (v1 != 0) {
        var v10 float32 = 1.5
        var v11 int64 = 611
        var v12 int16 = 42
        _ = v10
        _ = v11
        _ = v12
    } else {
        v4 = (^(v4))
        switch "foo bar" {
        case "foo bar":
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := v6[(-(fi13))]
                fmt.Println(fi13, v5)
                _ = fi13
                _ = v14
            }
        case "hello":
            if false {
                fmt.Println(v5, v4)
            } else {
                var v15 uint64 = 690
                if (!(false && ((v15 % ((2) | 1)) != v15))) {
                    v16 := GenS1{F0: 254}
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v1 -= v1
                        var v18 uint = (v1 & (v1 / ((v1) | 1)))
                        _ = fi17
                        _ = v18
                    }
                    v19 := GenS1{F0: 42}
                    _ = v16
                    _ = v19
                }
                _ = v15
            }
        case "\n\t":
            var v20 int64 = 1
            _ = v20
        }
        var v21 float64 = float64(-2.25)
        _ = v21
    }
    var v22 bool = true
    v5 += -2.25
    v4 -= 3
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 int16 = 42
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        var v26 float32 = 3.14159
        var v27 uint8 = v4
        v5 = 3.14159
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint8 = (v4 << ((18446744073709551615) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
