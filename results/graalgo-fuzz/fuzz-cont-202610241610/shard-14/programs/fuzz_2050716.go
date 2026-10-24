package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 uint8) uint64 {
    if (true && false) {
        if (!false) {
            for fi6 := 0; fi6 < 12; fi6++ {
                v7 := []int{fi6, -1, fi6, (fi6 % ((fi6) | 1))}
                fmt.Println(fi6)
                fmt.Println(p0)
                _ = fi6
                _ = v7
            }
        } else {
            switch "a" {
            case "a":
                for fi8 := 0; fi8 < 9; fi8++ {
                    v9 := []int{(fi8 >> ((255) & 63))}
                    v10 := GenS1{F0: -107}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            default:
                for fi11 := 0; fi11 < 6; fi11++ {
                    fmt.Println(fi11, fi11)
                    _ = fi11
                }
            }
        }
        v12 := GenS2{F0: 1}
        p0 = 2
        _ = v12
    }
    var v13 int64 = 1
    v13 += -1
    if false {
        v14 := []int{}
        _ = v14
    } else {
        if (42 <= -698) {
            v15 := GenS2{F0: 1}
            fmt.Println(p0, p0)
            var v16 uint = uint(p0)
            _ = v15
            _ = v16
        }
    }
    _ = p0
    _ = v13
    return 461
}

func main() {
    var v1 string = ("" + "\n\t")
    var v2 float32 = (1e10 * float32(0.0))
    v3 := []int{128}
    var v4 int64 = 1
    v5 := make(map[uint32]uint32)
    v2 -= 3.14159
    v17 := GenS1{F0: -1}
    v1 = v1
    var v18 uint8 = 154
    var v19 float32 = v2
    switch 32767 {
    case -1:
        _ = len(v5)
    case 9223372036854775807:
        switch -1 {
        case -9223372036854775808:
            v2 = v2
        case 1:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -2147483648:
            var v20 int32 = 713
            _ = v20
        default:
            v21 := GenS2{F0: 0}
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 8; fi22++ {
            v1 = "hello"
            for fi23 := 0; fi23 < 10; fi23++ {
                switch ("a" + (v1 + (v1 + v1))) {
                case "\n\t":
                    v24 := v5[65535]
                    _ = v24
                default:
                    if (-32768 > 42) {
                        var v25 int64 = (v4 >> ((931) & 63))
                        v26 := GenS3{F0: 282, F1: -32768}
                        _ = v25
                        _ = v26
                    }
                }
                fmt.Println(fi22)
                if len(v3) > 0 {
                    v3[0] = fi22
                }
                _ = fi23
            }
            _ = fi22
        }
    }
    switch "\n\t" {
    case "hello":
        for fi27 := 0; fi27 < 13; fi27++ {
            delete(v5, 4294967295)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi27
        }
    }
    v4 = 0
    for fi28 := 0; fi28 < 1; fi28++ {
        var v29 int32 = 94
        v19 += v19
        var v30 int32 = v29
        _ = fi28
        _ = v29
        _ = v30
    }
    v18 += (v18 | (v18 >> ((179) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
