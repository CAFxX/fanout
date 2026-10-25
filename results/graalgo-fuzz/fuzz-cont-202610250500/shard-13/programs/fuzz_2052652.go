package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint64, p1 uint64, p2 string) int32 {
    if true {
        var v7 int8 = 1
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int8 = v7
            switch p2 {
            case "":
                if ((fi8 & (fi8 + (fi8 % (((fi8 + 569)) | 1)))) < int(v7)) {
                    p2 = p2
                    if true {
                        var v10 bool = ((^(v7)) == -128)
                        _ = v10
                    }
                }
            case "\n\t":
                v7 = (^(v9))
            case "foo bar":
                var v11 int = (fi8 << ((453) & 63))
                _ = v11
            default:
                v9 -= (^(v7))
            }
            _ = fi8
            _ = v9
        }
        v12 := []int{}
        _ = v12
        _ = v7
    }
    var v13 uint32 = 42
    v13 = (v13 + 744)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 2147483647
}

func main() {
    v1 := []uint32{2147483647, 78}
    v2 := []string{"foo bar"}
    var v3 int = 1
    var v4 uint8 = 0
    var v5 uint64 = 1
    v6 := make(map[int]uint32)
    switch "a" {
    case "hello":
        var v14 uint16 = 0
        _ = v14
    case "":
        switch "hello" {
        case "":
            v5 -= (v5 << ((2) & 63))
        case "a":
            v15 := GenS1{F0: genFn0((^(v5)), (v5 >> ((758) & 63)), (("hello" + "a") + ("" + "a")))}
            _ = v15
        default:
            var v16 uint16 = uint16(v5)
            _ = v16
        }
    }
    switch "hello" {
    case "\n\t":
        v4 -= 42
    case "foo bar":
        v3 -= v3
    }
    if false {
        var v17 uint = 224
        for fi18 := 0; fi18 < 11; fi18++ {
            if (false && (!false)) {
                v5 = v5
            }
            v19 := GenS1{F0: -1}
            if ((v4 & v4) > 231) {
                v20 := GenS2{F0: 1, F1: (v17 >> (((v17 * (v17 - v17))) & 63)), F2: (v4 % ((2) | 1))}
                v4 += v4
                _ = v20
            }
            _ = fi18
            _ = v19
        }
        var v21 uint16 = 0
        _ = v17
        _ = v21
    }
    if false {
        v4 = (^(v4))
        v22 := GenS1{F0: -2147483648}
        _ = v22
    } else {
        v5 += (v5 + v5)
        var v23 int32 = 0
        var v24 uint64 = v5
        _ = v23
        _ = v24
    }
    if (255 >= 65535) {
        var v25 uint32 = 940
        _ = v25
    }
    var v26 int16 = 0
    for fi27 := 0; fi27 < 14; fi27++ {
        if len(v1) > 0 {
            v1[0] = 2
        }
        if true {
            var v28 int = -2147483648
            _ = v28
        } else {
            v29 := GenS2{F0: 0, F1: 255, F2: v4}
            _ = len(v2)
            for fi30 := 0; fi30 < 19; fi30++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi30
            }
            _ = v29
        }
        v31 := v6[(v3 & 4294967295)]
        _ = fi27
        _ = v31
    }
    var v32 uint = 557
    fmt.Println(v26, v5)
    fmt.Println(v3, v5, v5)
    if true {
        var v33 uint64 = (^(v5))
        v34 := GenS3{F0: 2, F1: 55}
        _ = v33
        _ = v34
    } else {
        fmt.Println(v4, v26)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
