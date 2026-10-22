package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
func genFn0(p0 int, p1 int16) uint8 {
    for fi4 := 0; fi4 < 4; fi4++ {
        p1 -= p1
        for fi5 := 0; fi5 < 17; fi5++ {
            p0 -= 9223372036854775807
            _ = fi5
        }
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int64 = -349
            _ = fi6
            _ = v7
        }
        _ = fi4
    }
    var v8 uint16 = 255
    var v9 int32 = 331
    p1 = (p1 & 1)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 94
}

func main() {
    var v1 int = 32767
    var v2 uint32 = 4294967295
    v3 := make(map[uint32]uint32)
    switch v1 {
    case 361:
        var v10 int64 = 0
        _ = v10
    case 32768:
        var v11 uint8 = 128
        _ = v11
    case -763:
        v2 -= v2
    default:
        var v12 int64 = -9223372036854775808
        _ = v12
    }
    var v13 uint8 = genFn0((-(v1)), -648)
    v1 = 809
    v14 := []int{127}
    v15 := GenS1{F0: 2, F1: 1, F2: v1}
    for fi16 := 0; fi16 < 3; fi16++ {
        fmt.Println(v2, v2, v1)
        switch fi16 {
        case 663:
            v2 += 2
        case 127:
            if false {
                v1 -= fi16
                fmt.Println(v2)
                for fi17 := 0; fi17 < 12; fi17++ {
                    if (255 > 65535) {
                        v1 += v1
                    } else {
                        v13 = v13
                        v13 -= v13
                    }
                    v1 = (fi16 << ((42) & 63))
                    var v18 int16 = -310
                    _ = fi17
                    _ = v18
                }
            } else {
                for fi19 := 0; fi19 < 14; fi19++ {
                    if true {
                        var v20 int32 = -2147483648
                        v21 := GenS2{F0: "\n\t"}
                        _ = v20
                        _ = v21
                    }
                    _ = fi19
                }
                var v22 float32 = 0.1
                var v23 uint64 = 4294967295
                _ = v22
                _ = v23
            }
        case 128:
            v14 = append(v14, (-(fi16)))
        }
        v24 := GenS3{F0: 18446744073709551615}
        _ = fi16
        _ = v24
    }
    var v25 int8 = 102
    if false {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        var v26 int32 = 1
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
}
