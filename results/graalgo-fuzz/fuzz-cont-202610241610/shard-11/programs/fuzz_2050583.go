package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint8 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int8; F1 int32; F2 uint }

func main() {
    var v1 bool = (42 > 18446744073709551615)
    v2 := []uint32{65535}
    v3 := make(map[int]bool)
    var v4 uint = 0
    if (v1 || (42 >= -1)) {
        v5 := GenS1{F0: 0, F1: 42, F2: 0}
        _ = v5
    }
    var v6 uint32 = 1
    if len(v2) > 0 {
        v2[0] = 65535
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (926 <= 1) {
        v7 := GenS2{F0: v4, F1: -133}
        fmt.Println(v6, v6, v4)
        for fi8 := 0; fi8 < 5; fi8++ {
            for fi9 := 0; fi9 < 7; fi9++ {
                switch "hello" {
                case "":
                    var v10 uint32 = v6
                    _ = v10
                default:
                    fmt.Println(fi9, v1, v1)
                }
                v1 = false
                _ = fi9
            }
            _ = fi8
        }
        _ = v7
    } else {
        var v11 bool = v1
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 uint64 = 0
            v14 := GenS1{F0: 127, F1: fi12, F2: 2}
            v1 = v11
            _ = fi12
            _ = v13
            _ = v14
        }
        v6 += (v6 % ((v6) | 1))
        _ = v11
    }
    v4 += (v4 >> (((v4 ^ (v4 & (v4 & 546)))) & 63))
    if (!v1) {
        v4 = (v4 / ((v4) | 1))
        fmt.Println(v6, v1)
    } else {
        v4 = (v4 * 351)
    }
    v15 := GenS1{F0: 255, F1: -2147483648, F2: 107}
    for fi16 := 0; fi16 < 19; fi16++ {
        fmt.Println(v6, fi16)
        for fi17 := 0; fi17 < 19; fi17++ {
            fmt.Println(fi16)
            var v18 int64 = -9223372036854775808
            v1 = (!(!v1))
            _ = fi17
            _ = v18
        }
        v19 := GenS3{F0: 0, F1: -1, F2: 1}
        _ = fi16
        _ = v19
    }
    var v20 int64 = 1
    v21 := GenS1{F0: 216, F1: 42, F2: 241}
    v2 = append(v2, 65535)
    if ((!(!v1)) && false) {
        if false {
            switch -1 {
            case 1:
                v6 = (v6 & (v6 - (v6 - v6)))
            }
            for fi22 := 0; fi22 < 17; fi22++ {
                v20 = -9223372036854775808
                _ = fi22
            }
        } else {
            v4 = (v4 / (((v4 % ((v4) | 1))) | 1))
        }
        if len(v2) > 0 {
            v2[0] = v6
        }
        if (-239 > -2147483648) {
            var v23 int64 = v20
            fmt.Println(v20)
            var v24 int64 = v20
            _ = v23
            _ = v24
        }
    } else {
        var v25 uint = (^(v4))
        switch 2147483647 {
        case -9223372036854775808:
            var v26 float32 = 1.5
            _ = v26
        }
        for fi27 := 0; fi27 < 14; fi27++ {
            v28 := GenS2{F0: v25, F1: 1}
            _ = fi27
            _ = v28
        }
        _ = v25
    }
    if v1 {
        var v29 uint = 4294967295
        if v1 {
            v1 = false
        }
        var v30 int8 = -128
        _ = v29
        _ = v30
    } else {
        var v31 bool = v1
        for fi32 := 0; fi32 < 2; fi32++ {
            fmt.Println(v20, v1)
            _ = fi32
        }
        for fi33 := 0; fi33 < 19; fi33++ {
            var v34 int8 = 42
            v1 = (v31 && v31)
            _ = fi33
            _ = v34
        }
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v6
}
