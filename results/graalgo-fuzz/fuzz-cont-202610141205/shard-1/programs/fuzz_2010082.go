package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16) int8 {
    switch "foo bar" {
    case "a":
        fmt.Println(p0)
    case "foo bar":
        var v5 int8 = 42
        _ = v5
    case "hello":
        var v6 float32 = 1e10
        _ = v6
    default:
        v7 := []int{-706}
        _ = v7
    }
    p0 = (p0 & 65535)
    p0 -= 782
    if (!true) {
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := []int{fi8, (fi8 & fi8), (fi8 + (fi8 >> ((1) & 63)))}
            _ = fi8
            _ = v9
        }
        p0 = (p0 / ((p0) | 1))
    }
    _ = p0
    return int8(p0)
}

func main() {
    var v1 uint16 = 42
    v2 := []int{127, 4294967295, 727}
    var v3 uint16 = (v1 >> ((42) & 63))
    v4 := []int{-2147483648}
    if true {
        v10 := GenS1{F0: (0.1 > (0.1 + 1.5))}
        if (1 == 1) {
            v2 = append(v2, -2147483648)
            v1 -= 0
        }
        _ = v10
    } else {
        v3 -= (v1 >> ((0) & 63))
    }
    var v11 uint64 = 184
    var v12 uint = 4294967295
    v3 -= (v3 % (((v3 | v3)) | 1))
    v3 -= 102
    var v13 uint16 = 42
    if len(v2) > 0 {
        v2[0] = 282
    }
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    _ = len(v4)
    for fi14 := 0; fi14 < 12; fi14++ {
        if (!(1 < 1)) {
            fmt.Println(v1, v1, v13)
            v15 := GenS1{F0: false}
            _ = v15
        } else {
            v16 := GenS1{F0: (v11 == v11)}
            if false {
                v11 = v11
                for fi17 := 0; fi17 < 11; fi17++ {
                    v13 -= v3
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v3, v1)
                    _ = fi17
                }
                _ = len(v2)
            }
            var v18 int = fi14
            _ = v16
            _ = v18
        }
        switch "" {
        case "foo bar":
            if (true || (42 <= -801)) {
                v4 = append(v4, fi14)
            } else {
                for fi19 := 0; fi19 < 20; fi19++ {
                    var v20 int8 = 42
                    var v21 int32 = int32(v20)
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                switch 42 {
                case 32767:
                    v22 := GenS2{F0: 861}
                    _ = v22
                default:
                    _ = len(v4)
                }
                v23 := GenS2{F0: int16(fi14)}
                _ = v23
            }
        case "a":
            v13 = uint16(v12)
        default:
            v1 = (v3 << ((v12) & 63))
        }
        _ = fi14
    }
    var v24 uint32 = 42
    var v25 float32 = 3.14159
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    fmt.Println(v12, v13)
    var v26 string = "hello"
    v3 -= v3
    var v27 int16 = -32768
    var v28 uint = v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
