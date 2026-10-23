package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int16 }
func genFn0(p0 int16) uint {
    v5 := []int{4294967295}
    var v6 int32 = 42
    var v7 int16 = p0
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: (fi8 & (fi8 ^ (fi8 >> ((499) & 63)))), F1: -1, F2: (p0 ^ (v7 % (((p0 | v7)) | 1)))}
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 bool = (32767 < 442)
            var v12 uint32 = 0
            v5 = append(v5, fi8)
            if len(v5) > 0 {
                v5[0] = (fi10 + (fi10 >> ((581) & 63)))
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 65535
    var v3 string = ("" + ("a" + "hello"))
    var v4 uint64 = 18446744073709551615
    v2 = v2
    var v13 uint8 = 255
    v2 -= 56
    v2 += v2
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 int16 = -32768
        _ = fi14
        _ = v15
    }
    fmt.Println(v13)
    switch "\n\t" {
    case "a":
        for fi16 := 0; fi16 < 12; fi16++ {
            for fi17 := 0; fi17 < 10; fi17++ {
                v13 = 20
                v13 = (v13 / ((v13) | 1))
                _ = len(v1)
                _ = fi17
            }
            for fi18 := 0; fi18 < 15; fi18++ {
                v3 = ((("foo bar" + v3) + "a") + v3)
                _ = fi18
            }
            _ = fi16
        }
    case "":
        if false {
            var v19 uint8 = (v13 + (v13 << ((genFn0(1)) & 63)))
            _ = v19
        }
    }
    fmt.Println(v3)
    for fi20 := 0; fi20 < 5; fi20++ {
        switch (fi20 ^ fi20) {
        case -9223372036854775808:
            var v21 uint32 = v2
            _ = v21
        default:
            v4 = 42
        }
        for fi22 := 0; fi22 < 20; fi22++ {
            switch (fi22 << ((genFn0(554)) & 63)) {
            case -32:
                v23 := v1[fi20]
                _ = v23
            default:
                v24 := v1[(fi22 << ((genFn0(-994)) & 63))]
                _ = v24
            }
            _ = len(v1)
            var v25 uint = 255
            _ = fi22
            _ = v25
        }
        v2 = v2
        _ = fi20
    }
    v13 += v13
    fmt.Println(v3)
    v3 = v3
    v26 := GenS1{F0: 1, F1: -9223372036854775808, F2: -32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
