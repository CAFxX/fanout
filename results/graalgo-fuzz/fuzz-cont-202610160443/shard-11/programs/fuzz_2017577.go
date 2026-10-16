package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int64 }

func main() {
    var v1 uint32 = 42
    v2 := []uint32{v1}
    var v3 uint16 = 42
    v4 := GenS1{F0: -1}
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 float32 = (1.5 + 1e10)
        if false {
            var v7 uint8 = 127
            _ = v7
        }
        if false {
            for fi8 := 0; fi8 < 15; fi8++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v1 -= uint32(v3)
                _ = fi8
            }
            fmt.Println(v3, fi5, v6)
        } else {
            v9 := GenS1{F0: 1}
            var v10 uint16 = (^(v3))
            _ = v10
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    var v11 uint32 = v1
    fmt.Println(v1, v11, v3)
    if false {
        v1 = 65535
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 bool = (false || (true || false))
                _ = fi13
                _ = v14
            }
            v2 = append(v2, v11)
            _ = fi12
        }
    } else {
        for fi15 := 0; fi15 < 8; fi15++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v1, v11)
            _ = fi15
        }
        var v16 uint32 = (^(v1))
        _ = v16
    }
    fmt.Println(v3, v1, v1)
    var v17 int16 = 32767
    if len(v2) > 0 {
        v2[0] = v11
    }
    switch 127 {
    case 806:
        v2 = append(v2, (v11 * v11))
    case 32767:
        switch -9223372036854775808 {
        case -2147483648:
            v11 -= v11
        case 32767:
            v18 := GenS2{F0: -2147483648, F1: 9223372036854775807}
            _ = v18
        case 990:
            var v19 string = "foo bar"
            _ = v19
        }
    case 838:
        v20 := GenS1{F0: v17}
        _ = v20
    }
    v2 = append(v2, 434)
    if true {
        if false {
            var v21 uint16 = v3
            _ = v21
        }
        var v22 uint8 = 0
        _ = v22
    }
    if (-94 == 42) {
        fmt.Println(v3, v3)
        v11 = 42
    } else {
        for fi23 := 0; fi23 < 15; fi23++ {
            fmt.Println(v11, v3, v3)
            _ = fi23
        }
        var v24 int8 = 100
        v1 = (v1 ^ 900)
        _ = v24
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        v1 = (v11 * 2147483648)
        fmt.Println(fi25, v1, v1)
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
