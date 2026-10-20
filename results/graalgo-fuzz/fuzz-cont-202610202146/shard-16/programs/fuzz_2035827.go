package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int32; F1 int64; F2 bool }
type GenS3 struct { F0 uint; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint }
type GenS5 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := []int{42, 410, 2147483647}
    var v2 uint32 = 696
    var v3 int8 = 65
    for fi4 := 0; fi4 < 5; fi4++ {
        v3 += (^(v3))
        _ = fi4
    }
    v5 := GenS1{F0: 0, F1: 42}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v6 int8 = 42
    if (!(v2 != 874)) {
        switch ("foo bar" + ("foo bar" + "foo bar")) {
        case "\n\t":
            v7 := GenS2{F0: 1, F1: 42, F2: true}
            _ = v7
        case "":
            if false {
                v3 += (v3 & (v3 << ((1) & 63)))
            }
        }
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int64 = -771
            _ = fi8
            _ = v9
        }
        switch -2147483648 {
        case 9223372036854775807:
            _ = len(v1)
        case 0:
            if (32767 <= 1000) {
                _ = len(v1)
            }
        }
    }
    v10 := GenS2{F0: 1, F1: -1, F2: true}
    fmt.Println(v6)
    v11 := GenS3{F0: 832, F1: -76, F2: 42}
    if (true && (false || true)) {
        v12 := GenS3{F0: 18446744073709551615, F1: (v6 + v3), F2: 0}
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 18; fi14++ {
                v6 -= (v3 | (v6 << ((651) & 63)))
                _ = len(v1)
                _ = fi14
            }
            _ = fi13
        }
        _ = v12
    } else {
        v15 := GenS4{F0: 65535, F1: -2147483648, F2: 4294967295}
        v16 := GenS5{F0: 42, F1: v2}
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 int64 = -9223372036854775808
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 bool = false
        _ = fi20
        _ = v21
    }
    v6 = 127
    for fi22 := 0; fi22 < 13; fi22++ {
        if (true && false) {
            v23 := GenS5{F0: 4294967295, F1: v2}
            v24 := GenS5{F0: 1, F1: v2}
            var v25 float64 = ((0.1 * 3.14159) * float64((0.1 * (0.0 * 1e10))))
            _ = v23
            _ = v24
            _ = v25
        } else {
            v3 = (v3 << ((1) & 63))
        }
        _ = fi22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
