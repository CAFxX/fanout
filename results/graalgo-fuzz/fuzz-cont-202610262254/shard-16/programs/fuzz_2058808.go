package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 bool, p1 int8, p2 int8) uint32 {
    p0 = (!true)
    v6 := GenS1{F0: 95, F1: 600}
    if ((2 < 55) || (p0 || false)) {
        fmt.Println(p0)
    } else {
        fmt.Println(p1, p2, p0)
        switch 1 {
        case 1:
            fmt.Println(p1, p1)
        default:
            var v7 int = 32768
            _ = v7
        }
        if (506 <= 2) {
            p0 = (0 != 128)
            p0 = p0
            v8 := []int{9223372036854775807}
            _ = v8
        }
    }
    fmt.Println(p2)
    v9 := []int{-1, 415, 32768, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return 65535
}
func genFn1(p0 uint32, p1 int16, p2 string) uint32 {
    var v10 float64 = float64(3.14159)
    var v11 int64 = -702
    for fi12 := 0; fi12 < 8; fi12++ {
        switch fi12 {
        case 2147483647:
            v10 -= float64(1.5)
        case 32767:
            v10 += (float64(1e10) / float64(1.5))
        default:
            for fi13 := 0; fi13 < 7; fi13++ {
                v14 := []int{fi13}
                switch fi13 {
                case -2147483648:
                    var v15 uint8 = 1
                    _ = v15
                default:
                    if (v11 < (v11 ^ 9223372036854775807)) {
                        var v16 uint32 = p0
                        v17 := GenS1{F0: (p0 + v16), F1: 1}
                        var v18 uint = 65535
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                }
                var v19 bool = true
                _ = fi13
                _ = v14
                _ = v19
            }
        }
        _ = fi12
    }
    fmt.Println(p2)
    switch 1 {
    case 42:
        v20 := []int{4294967295, -1, -2147483648, 948}
        _ = v20
    default:
        v21 := []int{0, -1}
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 4294967295
}

func main() {
    var v1 float32 = 1e10
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 736
    v5 := []string{}
    for fi22 := 0; fi22 < 1; fi22++ {
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 uint8 = 139
            fmt.Println(v24)
            fmt.Println(v4)
            _ = fi23
            _ = v24
        }
        fmt.Println(fi22, v4, fi22)
        v25 := GenS2{F0: 42}
        _ = fi22
        _ = v25
    }
    v3[125] = 111
    v26 := GenS2{F0: 1000}
    v2 -= 1
    var v27 bool = (!(42 > 42))
    for fi28 := 0; fi28 < 5; fi28++ {
        v4 += v4
        if true {
            var v29 int = 127
            if v27 {
                var v30 float32 = (v1 - 3.14159)
                for fi31 := 0; fi31 < 4; fi31++ {
                    switch v29 {
                    case 127:
                        var v32 uint8 = 1
                        _ = v32
                    }
                    _ = fi31
                }
                fmt.Println(v4, fi28, v1)
                _ = v30
            } else {
                v2 = (^(v2))
            }
            delete(v3, 2147483647)
            _ = v29
        }
        _ = fi28
    }
    if ("a" < "") {
        v1 += float32(1e10)
    }
    if (!v27) {
        if v27 {
            fmt.Println(v2)
            var v33 uint8 = 255
            fmt.Println(v27, v2)
            _ = v33
        }
        var v34 int32 = 0
        _ = v34
    }
    fmt.Println(v2, v1, v27)
    var v35 int8 = 127
    switch 315 {
    case 32767:
        var v36 int64 = -703
        _ = v36
    case 127:
        v35 -= (v35 >> ((18446744073709551615) & 63))
    case -9223372036854775808:
        v37 := v3[906]
        _ = v37
    default:
        _ = len(v3)
    }
    var v38 uint64 = (v2 ^ v2)
    v1 -= v1
    var v39 bool = v27
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
