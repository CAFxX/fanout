package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint64; F1 int; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 2
    v3 := make(map[string]uint32)
    v4 := []uint32{17, 0}
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]bool)
    v2 = 42
    v2 += v2
    fmt.Println(v2)
    v2 -= v2
    var v7 int64 = 0
    var v8 bool = true
    v5[2] = ("hello" + "foo bar")
    for fi9 := 0; fi9 < 17; fi9++ {
        if (65535 >= 255) {
            v4 = append(v4, 65535)
            switch fi9 {
            case 127:
                v10 := GenS1{F0: v2, F1: 18446744073709551615}
                _ = v10
            case 128:
                v7 -= (v7 - (v7 << ((694) & 63)))
            default:
                if (496 == -2147483648) {
                    v11 := v5[4294967295]
                    if len(v4) > 0 {
                        v4[0] = 399
                    }
                    v2 = 979
                    _ = v11
                }
            }
            var v12 uint64 = v2
            _ = v12
        } else {
            switch ("\n\t" + "\n\t") {
            case "a":
                v2 -= (v2 & v2)
            case "":
                v2 = (v2 % ((4294967295) | 1))
            }
            for fi13 := 0; fi13 < 11; fi13++ {
                v2 -= (v2 * v2)
                v2 += v2
                var v14 bool = (!true)
                _ = fi13
                _ = v14
            }
            if ((!(false && (fi9 >= fi9))) || v8) {
                v15 := GenS1{F0: 2, F1: uint(v2)}
                _ = v15
            }
        }
        _ = fi9
    }
    fmt.Println(v8)
    switch "a" {
    case "\n\t":
        var v16 float32 = -2.25
        _ = v16
    case "foo bar":
        for fi17 := 0; fi17 < 7; fi17++ {
            delete(v6, 2147483648)
            _ = fi17
        }
    }
    v8 = true
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 int8 = 0
        _ = len(v4)
        _ = len(v5)
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        v7 = v7
        _ = fi20
    }
    if true {
        v21 := GenS2{F0: 1, F1: 127, F2: 706}
        var v22 string = "\n\t"
        var v23 float64 = (0.1 + float64((0.1 * -2.25)))
        _ = v21
        _ = v22
        _ = v23
    } else {
        v24 := GenS3{F0: 120, F1: 2147483647, F2: 32767}
        _ = len(v5)
        v25 := GenS3{F0: 1, F1: 294, F2: 666}
        _ = v24
        _ = v25
    }
    var v26 bool = false
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = v2
    if (1 > 42) {
        if true {
            if v26 {
                for fi27 := 0; fi27 < 2; fi27++ {
                    for fi28 := 0; fi28 < 12; fi28++ {
                        var v29 int32 = -756
                        _ = fi28
                        _ = v29
                    }
                    var v30 uint64 = v2
                    _ = fi27
                    _ = v30
                }
            }
            _ = len(v4)
            v2 += v2
        }
    } else {
        if v8 {
            v26 = (255 < 255)
            fmt.Println(v2, v8, v8)
            v2 += (v2 ^ v2)
        } else {
            if v8 {
                v31 := GenS3{F0: 42, F1: 2, F2: 42}
                v2 = (v2 + (^(v2)))
                if len(v4) > 0 {
                    v4[0] = 0
                }
                _ = v31
            }
            _ = len(v4)
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
