package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int16 }
func genFn0(p0 string, p1 uint8) uint {
    v6 := []int{9223372036854775807, -55, 1}
    var v7 float32 = (3.14159 / 3.14159)
    if (-1 <= 127) {
        v8 := GenS1{F0: 630, F1: p0}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 151
}
func genFn1(p0 uint8, p1 int, p2 uint32) int64 {
    var v9 string = ""
    fmt.Println(v9, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 42
}

func main() {
    var v1 int32 = -500
    v2 := make(map[int]string)
    var v3 int = 42
    v4 := make(map[uint32]bool)
    var v5 float32 = float32(1.5)
    v10 := GenS1{F0: 590, F1: "a"}
    v3 = v3
    fmt.Println(v3)
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint8 = 127
        v1 += v1
        var v13 uint64 = 4294967295
        _ = fi11
        _ = v12
        _ = v13
    }
    if (!(612 > 42)) {
        if true {
            if (2 > 2147483647) {
                for fi14 := 0; fi14 < 8; fi14++ {
                    fmt.Println(v5)
                    _ = fi14
                }
                for fi15 := 0; fi15 < 14; fi15++ {
                    v5 += 3.14159
                    var v16 uint32 = 2147483647
                    _ = fi15
                    _ = v16
                }
            } else {
                if (false || ((!true) && false)) {
                    v17 := GenS2{F0: 42}
                    v18 := []int{v3, 50, v3}
                    for fi19 := 0; fi19 < 18; fi19++ {
                        v5 += (v5 * 1.5)
                        _ = fi19
                    }
                    _ = v17
                    _ = v18
                } else {
                    if (1e10 != v5) {
                        var v20 bool = (-1 == 9223372036854775807)
                        _ = v20
                    }
                    v21 := v2[443]
                    _ = v21
                }
            }
            if (!false) {
                var v22 uint16 = 255
                fmt.Println(v1)
                _ = v22
            }
            v23 := v2[v3]
            _ = v23
        } else {
            fmt.Println(v3, v1)
        }
        v3 -= v3
        v1 = (v1 * v1)
    } else {
        v5 += v5
        fmt.Println(v3, v3)
        v1 += -1
    }
    v24 := []int{(-(v3)), v3, 1, v3}
    for fi25 := 0; fi25 < 14; fi25++ {
        if true {
            var v26 float32 = v5
            var v27 string = "hello"
            _ = v26
            _ = v27
        }
        fmt.Println(v5)
        _ = fi25
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        v29 := GenS1{F0: 255, F1: ""}
        v5 = -2.25
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
        _ = fi30
    }
    fmt.Println(v3, v3, v1)
    v3 += v3
    v31 := GenS3{F0: v1, F1: -1}
    _ = len(v2)
    if false {
        if (true || (!(4294967295 <= 683))) {
            var v32 int64 = 1
            v4[42] = false
            _ = v32
        } else {
            v3 -= v3
            fmt.Println(v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
