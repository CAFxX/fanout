package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int8 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint }
type GenS3 struct { F0 int16 }
func genFn0(p0 bool, p1 uint8, p2 uint16) int {
    v5 := []int{}
    var v6 int64 = -9223372036854775808
    if p0 {
        p1 = p1
    } else {
        var v7 uint32 = 65535
        for fi8 := 0; fi8 < 17; fi8++ {
            v5 = append(v5, -2147483648)
            v5 = append(v5, 9223372036854775807)
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := GenS1{F0: v6, F1: fi9, F2: -100}
                if false {
                    var v11 uint8 = (p1 >> ((18446744073709551615) & 63))
                    var v12 int64 = (v6 >> ((443) & 63))
                    _ = len(v5)
                    _ = v11
                    _ = v12
                } else {
                    fmt.Println(p0)
                }
                fmt.Println(p1)
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        var v13 float32 = 1e10
        _ = v13
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 484
}
func genFn1(p0 int, p1 int32) int8 {
    v14 := []int{}
    if len(v14) > 0 {
        v14[0] = (p0 | genFn0(((false && true) && (!false)), 127, 2))
    }
    p1 -= 0
    _ = p0
    _ = p1
    _ = v14
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 1
    var v3 int = -1
    v4 := []uint32{4294967295}
    v4 = append(v4, 65535)
    if false {
        v2 = (^(v2))
        var v15 uint8 = v2
        var v16 uint64 = 2
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            v3 -= 32767
            _ = fi17
        }
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(v2, v2)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v19 bool = false
            _ = fi18
            _ = v19
        }
        v20 := GenS1{F0: 0, F1: (v3 / ((genFn0((true && (-2.25 <= 0.0)), v2, 236)) | 1)), F2: 42}
        _ = v20
    }
    switch "" {
    case "a":
        v2 = v2
    }
    v3 += 128
    var v21 string = ("a" + "hello")
    v22 := GenS1{F0: 42, F1: genFn0((true || (100 != genFn1(v3, 1))), (v2 - (v2 & (v2 | (v2 + v2)))), 65535), F2: 127}
    _ = len(v1)
    if ((1e10 * 0.1) != 1e10) {
        var v23 int = 127
        fmt.Println(v23, v21)
        _ = v23
    } else {
        v3 = v3
        v4 = append(v4, 13)
    }
    var v24 int32 = -2147483648
    for fi25 := 0; fi25 < 9; fi25++ {
        switch (v3 % ((fi25) | 1)) {
        case 2147483647:
            v4 = append(v4, 65535)
        case -824:
            v21 = "a"
        case 128:
            v26 := GenS2{F0: v2, F1: 0, F2: 65535}
            _ = v26
        default:
            if true {
                switch v21 {
                case "\n\t":
                    var v27 int32 = -414
                    _ = v27
                }
                v1[2147483647] = 1
                fmt.Println(v2, v24, v2)
            }
        }
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 string = v21
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi28
            _ = v29
        }
        v2 -= (^(v2))
        _ = fi25
    }
    switch v3 {
    case 267:
        var v30 int16 = 32767
        _ = v30
    case 1:
        for fi31 := 0; fi31 < 18; fi31++ {
            if (471 >= 212) {
                if ((0.0 * 0.1) >= 0.0) {
                    var v32 uint8 = v2
                    var v33 int8 = 42
                    v21 = v21
                    _ = v32
                    _ = v33
                } else {
                    v21 = v21
                    v21 = (v21 + v21)
                }
            }
            var v34 uint = 65535
            v35 := GenS3{F0: 963}
            _ = fi31
            _ = v34
            _ = v35
        }
    }
    v36 := GenS2{F0: v2, F1: -3, F2: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v36
    _ = v4
}
