package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int }

func main() {
    var v1 int8 = 0
    var v2 bool = (!(0 > 255))
    v3 := make(map[uint32]string)
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    var v6 float32 = 3.14159
    for fi7 := 0; fi7 < 11; fi7++ {
        delete(v5, 0)
        var v8 int8 = (v1 % (((v1 % ((71) | 1))) | 1))
        _ = fi7
        _ = v8
    }
    v6 = -2.25
    var v9 uint = 255
    var v10 int8 = v1
    if v2 {
        var v11 int64 = 0
        _ = v11
    }
    v1 += (v1 | v10)
    var v12 uint16 = 0
    fmt.Println(v2, v10)
    switch -9223372036854775808 {
    case -215:
        var v13 string = "a"
        _ = v13
    case 138:
        v14 := GenS1{F0: 0}
        _ = v14
    }
    v2 = v2
    v9 = 4294967295
    fmt.Println(v2, v1, v9)
    if (v2 || ((v2 && false) && (766 == 798))) {
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 uint16 = v12
            var v17 uint8 = 1
            _ = fi15
            _ = v16
            _ = v17
        }
        if v2 {
            for fi18 := 0; fi18 < 12; fi18++ {
                fmt.Println(v2)
                fmt.Println(v6, v2)
                var v19 int16 = 42
                _ = fi18
                _ = v19
            }
            fmt.Println(v2, v12, v9)
        } else {
            if (((float32(1.5) + 1.5) + v6) <= v6) {
                var v20 uint16 = (v12 << ((v9) & 63))
                _ = v20
            }
            v21 := GenS2{F0: (^(v12)), F1: 32767}
            _ = v21
        }
        v12 += 65535
    } else {
        v10 = 1
        switch 2147483647 {
        case 1:
            var v22 int32 = -2147483648
            _ = v22
        case 127:
            delete(v3, 680)
        case 771:
            for fi23 := 0; fi23 < 8; fi23++ {
                v10 += v1
                var v24 uint8 = 2
                var v25 int64 = 9223372036854775807
                _ = fi23
                _ = v24
                _ = v25
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
