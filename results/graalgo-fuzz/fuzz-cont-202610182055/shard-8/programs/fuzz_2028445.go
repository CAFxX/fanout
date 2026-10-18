package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint, p1 int8, p2 bool) bool {
    v7 := []int{}
    var v8 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (255 > 866)
}
func genFn1(p0 uint16, p1 int) int {
    v9 := []int{32767}
    p0 = p0
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int = (-(p1))
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS1{F0: -1, F1: p0}
            _ = len(v9)
            fmt.Println(fi10)
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = v9
    return (^(p1))
}

func main() {
    var v1 int32 = 42
    var v2 uint = 18446744073709551615
    var v3 float64 = (0.0 / 1.0)
    v4 := make(map[string]uint32)
    var v5 uint8 = 0
    var v6 int = 128
    switch v6 {
    case -9223372036854775808:
        v2 -= v2
    case 32767:
        var v14 float64 = float64((3.14159 - 1e10))
        _ = v14
    case -135:
        fmt.Println(v3, v6)
    default:
        switch v6 {
        case 4294967295:
            if ((0.0 / float32(v3)) >= (0.0 - float32(((v3 / 3.14159) / (v3 + float64(-2.25)))))) {
                var v15 uint8 = (v5 % ((v5) | 1))
                _ = v15
            } else {
                switch genFn1(690, (v6 ^ v6)) {
                case 2147483647:
                    fmt.Println(v5, v3, v3)
                case 4294967295:
                    v16 := []int{-2147483648}
                    _ = v16
                case 127:
                    var v17 uint32 = 575
                    _ = v17
                default:
                    var v18 uint = (v2 ^ (^(v2)))
                    _ = v18
                }
                if (false && (!false)) {
                    v19 := GenS1{F0: -1, F1: 1}
                    _ = v19
                } else {
                    v20 := GenS1{F0: 0, F1: 0}
                    v21 := GenS1{F0: 1, F1: 255}
                    _ = v20
                    _ = v21
                }
            }
        }
    }
    var v22 uint8 = v5
    for fi23 := 0; fi23 < 1; fi23++ {
        fmt.Println(v5, v3)
        _ = fi23
    }
    if false {
        v1 = v1
        if (255 != v5) {
            _ = len(v4)
        }
    }
    if false {
        if (v1 != (v1 << ((v2) & 63))) {
            v22 = (v22 - v22)
            switch "a" {
            case "\n\t":
                v24 := []int{int(v22), (v6 + v6), v6}
                _ = v24
            case "foo bar":
                v25 := GenS1{F0: -1, F1: 65535}
                _ = v25
            case "":
                var v26 int32 = 153
                _ = v26
            }
        } else {
            v3 = v3
            fmt.Println(v2, v6, v1)
        }
    } else {
        var v27 int16 = 1000
        v22 = (v22 >> ((813) & 63))
        _ = v27
    }
    v28 := GenS2{F0: 1, F1: 782, F2: 0}
    var v29 float64 = v3
    for fi30 := 0; fi30 < 15; fi30++ {
        if ((v22 != v22) && (!false)) {
            var v31 int64 = 42
            var v32 int32 = (v1 ^ v1)
            var v33 uint32 = 42
            _ = v31
            _ = v32
            _ = v33
        } else {
            switch fi30 {
            case 127:
                var v34 float64 = 0.0
                _ = v34
            case 1:
                switch 32768 {
                case 42:
                    v35 := []int{fi30, genFn1(42, (v6 + (^(fi30)))), int(v2), fi30}
                    _ = v35
                default:
                    var v36 bool = (true || (v1 >= (v1 & (v1 - v1))))
                    _ = v36
                }
            }
            var v37 int16 = 42
            if (!true) {
                v4["\n\t"] = 301
            } else {
                var v38 uint64 = 2
                var v39 int = fi30
                v40 := GenS3{F0: 152}
                _ = v38
                _ = v39
                _ = v40
            }
            _ = v37
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
