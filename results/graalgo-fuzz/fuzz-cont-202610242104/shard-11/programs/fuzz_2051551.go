package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int }
type GenS2 struct { F0 uint64; F1 int32; F2 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 int; F2 uint }
type GenS5 struct { F0 uint16; F1 uint32; F2 int64 }
func genFn0(p0 uint8, p1 int32) string {
    var v5 uint8 = (p0 << ((18446744073709551615) & 63))
    v5 += v5
    _ = p0
    _ = p1
    _ = v5
    return ("" + "\n\t")
}
func genFn1(p0 uint16, p1 uint, p2 uint16) int16 {
    var v6 uint16 = 500
    if (-9223372036854775808 < -9223372036854775808) {
        v7 := GenS1{F0: (^(v6)), F1: -1, F2: 32768}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 6; fi8++ {
            if true {
                v9 := GenS1{F0: p2, F1: (^(fi8)), F2: fi8}
                _ = v9
            } else {
                var v10 uint32 = 938
                v11 := GenS2{F0: 445, F1: 2147483647, F2: 2147483647}
                fmt.Println(p2, v10)
                _ = v10
                _ = v11
            }
            if false {
                switch (fi8 - -256) {
                case -280:
                    var v12 string = "foo bar"
                    _ = v12
                case 32767:
                    p0 += (p0 * 255)
                }
            }
            var v13 uint64 = 18446744073709551615
            _ = fi8
            _ = v13
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            fmt.Println(p2, p1)
            var v15 int16 = -599
            v6 = 32767
            _ = fi14
            _ = v15
        }
        fmt.Println(p1, p2, p1)
    }
    v16 := []int{1}
    fmt.Println(p1)
    for fi17 := 0; fi17 < 13; fi17++ {
        p0 = (p0 % ((1) | 1))
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(fi18, fi18)
            fmt.Println(v6)
            _ = fi18
        }
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v6
    return -32768
}

func main() {
    var v1 int16 = 32767
    var v2 int8 = 42
    v3 := make(map[uint32]string)
    v4 := []uint32{671}
    v1 += (v1 << ((1) & 63))
    switch 4294967295 {
    case -1:
        _ = len(v3)
    default:
        switch 32767 {
        case -538:
            for fi19 := 0; fi19 < 3; fi19++ {
                switch "foo bar" {
                case "foo bar":
                    var v20 int64 = 136
                    _ = v20
                }
                _ = fi19
            }
        case -2147483648:
            _ = len(v4)
        case 127:
            v21 := GenS3{F0: 1}
            _ = v21
        default:
            v1 += genFn1(395, 0, 65535)
        }
    }
    var v22 uint32 = 2
    fmt.Println(v22)
    if len(v4) > 0 {
        v4[0] = 0
    }
    var v23 int64 = -1
    _ = len(v4)
    v2 -= (v2 + (v2 / ((v2) | 1)))
    v23 = (v23 * (-(v23)))
    v2 += (v2 & (v2 % (((^(v2))) | 1)))
    v24 := GenS2{F0: 637, F1: 1, F2: 1}
    v25 := GenS3{F0: (^(v22))}
    for fi26 := 0; fi26 < 20; fi26++ {
        if true {
            var v27 uint = 0
            _ = v27
        }
        v28 := GenS4{F0: "\n\t", F1: (fi26 << ((728) & 63)), F2: 1}
        delete(v3, v22)
        _ = fi26
        _ = v28
    }
    v29 := GenS5{F0: 992, F1: 793, F2: v23}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
}
