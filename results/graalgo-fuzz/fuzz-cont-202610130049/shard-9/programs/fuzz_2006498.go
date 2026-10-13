package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 bool }

func main() {
    v1 := []int{-2147483648}
    v2 := []int{9223372036854775807, 1, 32768, 0}
    v3 := make(map[int]int)
    var v4 uint = 255
    var v5 int64 = 9223372036854775807
    var v6 float64 = 3.14159
    fmt.Println(v6, v6)
    v7 := GenS1{F0: 2147483648, F1: (v5 | -727), F2: 42}
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            v6 = (v6 / v6)
            _ = fi9
        }
        fmt.Println(v6)
        _ = fi8
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(v6)
        _ = fi10
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        v4 = v4
        if true {
            v4 = 140
        }
        var v12 int32 = -864
        _ = fi11
        _ = v12
    }
    if (!((false && false) || (255 != 255))) {
        v3[573] = 128
    }
    _ = len(v1)
    fmt.Println(v6)
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint64 = 205
        for fi15 := 0; fi15 < 5; fi15++ {
            v5 -= v5
            var v16 uint16 = 32767
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    v17 := GenS2{F0: -1, F1: 255}
    var v18 int64 = v5
    if (1 == 996) {
        v2 = append(v2, -889)
    } else {
        var v19 uint32 = 144
        v20 := GenS1{F0: (v19 & v19), F1: v5, F2: -128}
        if (true && false) {
            if false {
                v6 = 3.14159
                if false {
                    for fi21 := 0; fi21 < 6; fi21++ {
                        var v22 uint16 = 54
                        _ = fi21
                        _ = v22
                    }
                    var v23 string = "foo bar"
                    _ = v23
                }
                v19 = (v19 & (^(v19)))
            } else {
                for fi24 := 0; fi24 < 19; fi24++ {
                    var v25 string = "a"
                    _ = fi24
                    _ = v25
                }
                v26 := GenS3{F0: true}
                _ = v26
            }
            v27 := GenS3{F0: true}
            if true {
                var v28 int = 4294967295
                var v29 int32 = 256
                _ = v28
                _ = v29
            }
            _ = v27
        }
        _ = v19
        _ = v20
    }
    var v30 uint8 = 1
    if false {
        v30 -= (v30 / ((0) | 1))
    } else {
        v1 = append(v1, 42)
        if (((v18 % (((v18 * v5)) | 1)) > int64(v30)) || true) {
            fmt.Println(v30)
        }
    }
    if true {
        var v31 uint16 = 255
        v31 = uint16(v4)
        _ = v31
    } else {
        switch 9223372036854775807 {
        case 4294967295:
            var v32 float64 = float64(1.5)
            _ = v32
        case 127:
            v30 -= v30
        case 0:
            if (-32768 != 0) {
                var v33 int32 = 0
                var v34 bool = false
                _ = v33
                _ = v34
            } else {
                fmt.Println(v5, v30)
                var v35 int16 = 521
                _ = v35
            }
        }
        v6 += (v6 * (v6 / -2.25))
        fmt.Println(v5, v5, v5)
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        _ = len(v2)
        _ = fi36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
