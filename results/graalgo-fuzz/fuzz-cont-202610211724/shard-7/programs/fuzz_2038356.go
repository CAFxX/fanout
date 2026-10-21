package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint32 }
type GenS2 struct { F0 int16; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0(p0 uint16, p1 uint8) int {
    var v7 int = 128
    if (42 > 65535) {
        if false {
            switch v7 {
            case 127:
                if true {
                    switch (v7 + v7) {
                    case 128:
                        fmt.Println(p1, v7, p0)
                    case 2147483647:
                        switch "\n\t" {
                        case "foo bar":
                            v8 := []int{(^(v7)), v7, 4294967295, v7}
                            _ = v8
                        default:
                            p1 = p1
                        }
                    }
                }
            default:
                v9 := GenS1{F0: 1, F1: 42, F2: 65535}
                _ = v9
            }
            v10 := GenS1{F0: 1, F1: int16(v7), F2: 567}
            _ = v10
        } else {
            p1 -= 255
        }
        v11 := []int{(v7 << ((0) & 63))}
        _ = v11
    }
    var v12 uint16 = p0
    switch (-(v7)) {
    case -307:
        v7 += 0
    case -2147483648:
        for fi13 := 0; fi13 < 19; fi13++ {
            v12 += p0
            _ = fi13
        }
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    return 4294967295
}
func genFn1(p0 string, p1 int64) int32 {
    v14 := GenS1{F0: -29, F1: -32768, F2: 4294967295}
    fmt.Println(p0, p1)
    v15 := []int{}
    if (0 > 445) {
        for fi16 := 0; fi16 < 17; fi16++ {
            p0 = "hello"
            _ = fi16
        }
    } else {
        p0 = "hello"
    }
    v17 := GenS2{F0: 0, F1: -422, F2: 236}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v17
    return 42
}

func main() {
    var v1 int32 = -1
    var v2 int32 = v1
    var v3 float32 = -2.25
    v4 := make(map[uint32]bool)
    v5 := make(map[int]uint32)
    var v6 float64 = ((-2.25 - 1.5) * 3.14159)
    v18 := []int{genFn0(1, 2), -148, -2147483648, -583}
    if false {
        var v19 float32 = v3
        _ = v19
    }
    v2 = v2
    v6 -= -2.25
    var v20 string = ("\n\t" + "")
    if true {
        var v21 int32 = v1
        var v22 int = 128
        var v23 uint16 = 65535
        _ = v21
        _ = v22
        _ = v23
    }
    fmt.Println(v1)
    if true {
        var v24 float64 = 1e10
        var v25 int16 = 42
        if true {
            for fi26 := 0; fi26 < 1; fi26++ {
                v27 := v4[2]
                _ = fi26
                _ = v27
            }
            switch 127 {
            case 4294967295:
                var v28 uint16 = 42
                _ = v28
            case 128:
                v24 -= 0.0
            case -321:
                switch "foo bar" {
                case "\n\t":
                    var v29 int16 = v25
                    _ = v29
                case "a":
                    v3 += 1.5
                }
            default:
                fmt.Println(v6)
            }
        } else {
            v30 := GenS3{F0: 32767, F1: false}
            v24 -= (3.14159 + (0.1 + v6))
            _ = v30
        }
        _ = v24
        _ = v25
    }
    v3 = ((v3 * v3) - (v3 - (0.1 / (v3 / v3))))
    v31 := GenS2{F0: -32768, F1: 0, F2: 1}
    v32 := GenS3{F0: 42, F1: (!true)}
    switch 32767 {
    case 0:
        if true {
            v33 := GenS3{F0: 2, F1: ((3.14159 - 0.1) >= (v6 + 0.1))}
            _ = v33
        } else {
            var v34 int = 42
            _ = v34
        }
    case 2147483647:
        var v35 int64 = -9223372036854775808
        _ = v35
    default:
        v20 = v20
    }
    v36 := GenS2{F0: 1000, F1: 32767, F2: 255}
    v37 := GenS3{F0: 421, F1: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
