package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int64 }
type GenS2 struct { F0 string; F1 string; F2 int }
func genFn0(p0 bool, p1 uint32, p2 int16) uint64 {
    v4 := []int{32768, 705, 32767}
    if true {
        if p0 {
            if false {
                if (1e10 <= (-2.25 / -2.25)) {
                    for fi5 := 0; fi5 < 5; fi5++ {
                        fmt.Println(p1, p0)
                        var v6 float64 = 1.5
                        _ = fi5
                        _ = v6
                    }
                    for fi7 := 0; fi7 < 7; fi7++ {
                        fmt.Println(p0, p2, p0)
                        _ = fi7
                    }
                }
                for fi8 := 0; fi8 < 19; fi8++ {
                    fmt.Println(p2, fi8)
                    var v9 string = ""
                    var v10 float32 = (1e10 * float32(1.5))
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                v4 = append(v4, 9223372036854775807)
            }
        }
    } else {
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 uint32 = 2147483647
            for fi13 := 0; fi13 < 5; fi13++ {
                p0 = p0
                fmt.Println(v12)
                p2 = p2
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
        p2 -= 1
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS1{F0: (p0 && (255 != 0)), F1: p2, F2: 42}
            _ = fi14
            _ = v15
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if p0 {
        var v16 int = 32768
        _ = v16
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            v4 = append(v4, -2147483648)
            _ = fi17
        }
    }
    v18 := GenS1{F0: (!false), F1: 776, F2: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v4
    return 18446744073709551615
}
func genFn1() int32 {
    var v19 float32 = (0.0 / (-2.25 + float32(0.1)))
    if false {
        v20 := []int{-210, -914, 241, -382}
        var v21 uint8 = 127
        _ = v20
        _ = v21
    }
    _ = v19
    return -1
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[int]uint32)
    var v3 uint32 = 0
    v22 := GenS2{F0: "a", F1: "hello", F2: 32768}
    v23 := v2[v1]
    var v24 uint16 = 194
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 int8 = -128
        _ = fi25
        _ = v26
    }
    var v27 float32 = ((0.1 + (0.0 * 1.5)) - -2.25)
    if true {
        fmt.Println(v3, v27)
        fmt.Println(v24, v23, v1)
        v23 += v3
    } else {
        v2[(v1 / ((int(v3)) | 1))] = (v23 << ((uint(v3)) & 63))
        for fi28 := 0; fi28 < 14; fi28++ {
            for fi29 := 0; fi29 < 11; fi29++ {
                v1 = fi28
                _ = fi29
            }
            var v30 bool = true
            _ = fi28
            _ = v30
        }
    }
    v1 += (v1 | v1)
    if (!true) {
        if true {
            v27 -= 1.5
            v3 -= 0
            v27 = -2.25
        }
        var v31 uint32 = v3
        _ = v31
    }
    var v32 int32 = genFn1()
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v32
}
