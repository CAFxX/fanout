package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int }

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]bool)
    var v3 int8 = -1
    var v4 int16 = 32767
    v5 := make(map[uint32]string)
    var v6 uint16 = 399
    v7 := GenS1{F0: (1 != 354), F1: -1, F2: v6}
    v3 += v3
    if (((18446744073709551615 > 4294967295) && false) && true) {
        v8 := GenS2{F0: 65535}
        _ = v8
    }
    if false {
        for fi9 := 0; fi9 < 16; fi9++ {
            fmt.Println(fi9, fi9)
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 uint = 2
                v1 = 42
                v12 := v2[2147483648]
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
        for fi13 := 0; fi13 < 6; fi13++ {
            fmt.Println(v1, v4)
            _ = fi13
        }
    } else {
        switch 42 {
        case 9223372036854775807:
            if (false || false) {
                var v14 int64 = 405
                _ = v14
            } else {
                v15 := GenS3{F0: 552, F1: v6}
                var v16 int16 = -32768
                var v17 int = 883
                _ = v15
                _ = v16
                _ = v17
            }
        }
        v4 = -1
        fmt.Println(v1, v4, v4)
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        v3 = int8(fi18)
        if false {
            v4 = (v4 | (v4 - (v4 >> ((4294967295) & 63))))
            var v19 uint8 = 1
            _ = v19
        } else {
            if (fi18 != 2147483647) {
                if false {
                    v20 := GenS4{F0: 132}
                    _ = v20
                }
            } else {
                fmt.Println(v3, v3)
            }
            v21 := GenS4{F0: (^(fi18))}
            _ = v21
        }
        _ = fi18
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        fmt.Println(v4, v6)
        _ = fi22
    }
    var v23 int8 = (v3 & v3)
    var v24 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
