package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint }
type GenS2 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    v1 := []uint32{824, 1, 248}
    var v2 int16 = -997
    var v3 uint8 = 174
    var v4 uint32 = 2147483647
    fmt.Println(v3)
    var v5 int = 32767
    for fi6 := 0; fi6 < 16; fi6++ {
        if false {
            switch (fi6 / (((^(fi6))) | 1)) {
            case 32768:
                _ = len(v1)
            }
            fmt.Println(v4)
            var v7 uint64 = 0
            _ = v7
        }
        fmt.Println(v5, v3, v2)
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int16 = v2
            _ = fi8
            _ = v9
        }
        _ = fi6
    }
    fmt.Println(v5, v3, v4)
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            v1 = append(v1, 2147483647)
            v5 -= v5
            v12 := GenS1{F0: (-6 > 42), F1: v2, F2: 255}
            v1 = append(v1, v4)
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v1 = append(v1, 65535)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch v5 {
    case -9223372036854775808:
        if true {
            v13 := GenS2{F0: (v3 & (^(v3))), F1: 65535, F2: v2}
            fmt.Println(v5)
            v5 = 32768
            _ = v13
        } else {
            for fi14 := 0; fi14 < 16; fi14++ {
                _ = len(v1)
                _ = fi14
            }
            v5 = 128
        }
    case 127:
        v15 := GenS3{F0: true, F1: 0}
        _ = v15
    default:
        fmt.Println(v4, v3)
    }
    v2 = (v2 - (v2 | (v2 - (v2 << ((42) & 63)))))
    if len(v1) > 0 {
        v1[0] = 2
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
