package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 string }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int; F1 uint64 }
func genFn0(p0 uint8) string {
    v6 := []int{0, 32767, 4294967295}
    switch ("\n\t" + "") {
    case "a":
        var v7 int64 = -1
        _ = v7
    case "hello":
        var v8 int32 = -840
        _ = v8
    default:
        p0 = (p0 | p0)
    }
    _ = p0
    _ = v6
    return ("\n\t" + ("\n\t" + "\n\t"))
}
func genFn1() int64 {
    v9 := GenS1{F0: true, F1: false, F2: ("foo bar" + "")}
    var v10 bool = false
    if v10 {
        var v11 uint = 65535
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS2{F0: 759, F1: 660}
            v11 += (v11 / ((0) | 1))
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 int32 = 0
            v10 = (65535 >= 2147483647)
            _ = fi14
            _ = v15
        }
        _ = v11
    } else {
        fmt.Println(v10)
        var v16 string = ""
        v17 := []int{4294967295, 396, 266, -1}
        _ = v16
        _ = v17
    }
    v18 := []int{867, 9223372036854775807}
    _ = v10
    _ = v18
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{4294967295}
    var v3 uint = 180
    var v4 uint32 = 1
    var v5 int32 = 2147483647
    for fi19 := 0; fi19 < 2; fi19++ {
        fmt.Println(v5, v4, fi19)
        v20 := GenS3{F0: v4, F1: 235}
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: 32767, F1: (v4 / (((v4 + 1)) | 1))}
    switch "a" {
    case "hello":
        v1 = v1
    case "":
        v3 += uint(v4)
    }
    v5 -= v5
    if false {
        var v22 bool = ((false || (false || (299 != -2147483648))) && (1 != 255))
        if false {
            if len(v2) > 0 {
                v2[0] = 0
            }
        } else {
            switch v1 {
            case "a":
                var v23 uint64 = 2
                _ = v23
            case "\n\t":
                v5 = 1
            case "":
                if ((127 < 2) || (998 == 121)) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    var v24 string = v1
                    _ = v24
                }
            default:
                v25 := GenS4{F0: 972, F1: 0}
                _ = v25
            }
        }
        _ = v22
    } else {
        var v26 int16 = 1
        v3 -= (v3 << ((v3) & 63))
        _ = v26
    }
    v3 = v3
    var v27 float32 = 1.5
    if false {
        v1 = v1
    } else {
        switch v1 {
        case "hello":
            v1 = v1
        case "foo bar":
            if true {
                switch 42 {
                case 4294967295:
                    fmt.Println(v5, v5)
                case 127:
                    v27 = ((0.0 * v27) / v27)
                }
                fmt.Println(v3, v1, v1)
                v28 := GenS4{F0: -398, F1: 733}
                _ = v28
            }
        default:
            if len(v2) > 0 {
                v2[0] = (v4 ^ v4)
            }
        }
        var v29 int32 = (v5 * (v5 * 42))
        _ = v29
    }
    _ = len(v2)
    for fi30 := 0; fi30 < 14; fi30++ {
        v3 -= v3
        v31 := GenS2{F0: 1000, F1: (v4 / ((128) | 1))}
        _ = fi30
        _ = v31
    }
    v27 -= float32(float64((v27 * v27)))
    fmt.Println(v1, v27)
    v27 -= 3.14159
    var v32 int = -1
    v27 += (v27 / v27)
    v33 := GenS2{F0: -415, F1: (v4 | (v4 ^ v4))}
    fmt.Println(v5, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
