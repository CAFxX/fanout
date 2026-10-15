package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int32 }

func main() {
    v1 := []uint32{42, 2}
    var v2 uint16 = 65535
    var v3 int16 = 42
    fmt.Println(v2)
    v2 -= 42
    v2 += (^(v2))
    var v4 string = ("foo bar" + "")
    for fi5 := 0; fi5 < 7; fi5++ {
        switch -1 {
        case -850:
            if true {
                v6 := GenS1{F0: -40, F1: (v4 + v4), F2: 0}
                fmt.Println(v2)
                var v7 int16 = v3
                _ = v6
                _ = v7
            } else {
                if len(v1) > 0 {
                    v1[0] = 487
                }
                v4 = v4
            }
        default:
            for fi8 := 0; fi8 < 9; fi8++ {
                fmt.Println(v2, v3, v4)
                var v9 uint = 4294967295
                v3 = v3
                fmt.Println(v4, v2, fi8)
                var v10 float32 = float32(1.5)
                var v11 uint16 = (^(v2))
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        }
        var v12 uint8 = 127
        _ = fi5
        _ = v12
    }
    var v13 int16 = (v3 >> ((82) & 63))
    switch 32768 {
    case -483:
        v2 -= v2
    }
    if (!(32768 != -2147483648)) {
        for fi14 := 0; fi14 < 17; fi14++ {
            if true {
                v15 := GenS1{F0: 78, F1: (v4 + ("\n\t" + (v4 + v4))), F2: 0}
                if len(v1) > 0 {
                    v1[0] = 90
                }
                v13 += (v3 << ((255) & 63))
                _ = v15
            }
            _ = fi14
        }
        var v16 uint16 = 1
        _ = v16
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v17 float32 = 0.1
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 string = v4
        var v20 int = 128
        var v21 int64 = -1
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
