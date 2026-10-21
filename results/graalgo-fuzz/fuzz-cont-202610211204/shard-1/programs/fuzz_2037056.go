package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 bool, p1 uint, p2 int32) int8 {
    if p0 {
        fmt.Println(p1, p2)
        var v6 uint64 = 629
        _ = v6
    } else {
        v7 := []int{-136, 127}
        fmt.Println(p2, p0, p1)
        _ = v7
    }
    var v8 int32 = (p2 % (((p2 * p2)) | 1))
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(p1, p1)
        _ = fi9
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 6
}

func main() {
    var v1 int32 = 202
    v2 := []int{113, -2147483648, 829}
    v3 := make(map[int]int)
    v4 := []uint32{2147483648, 1, 2147483648, 42}
    v5 := []int{32767}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v10 uint8 = 42
    if true {
        v1 = (v1 >> ((uint(v10)) & 63))
        var v11 uint32 = 1
        _ = v11
    } else {
        v10 -= v10
        if len(v4) > 0 {
            v4[0] = 2
        }
        v12 := GenS1{F0: 0}
        _ = v12
    }
    var v13 string = "a"
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS1{F0: 463}
        _ = fi14
        _ = v15
    }
    var v16 int8 = genFn0((true && (!(!false))), 267, v1)
    v1 = (v1 % ((-2147483648) | 1))
    v17 := GenS2{F0: 1, F1: 72}
    var v18 bool = true
    if v18 {
        var v19 uint = 128
        switch v13 {
        case "foo bar":
            if v18 {
                switch (v13 + "\n\t") {
                case "hello":
                    if v18 {
                        var v20 uint32 = 65535
                        var v21 float32 = ((1e10 + -2.25) + 0.0)
                        delete(v3, -356)
                        _ = v20
                        _ = v21
                    }
                case "foo bar":
                    v16 -= (v16 >> ((475) & 63))
                }
                var v22 uint32 = 4294967295
                var v23 float32 = float32(((0.0 + 0.1) / float64(float32(((1.5 / 3.14159) + float64(3.14159))))))
                _ = v22
                _ = v23
            } else {
                var v24 uint8 = v10
                fmt.Println(v16, v1, v24)
                v3[42] = 4294967295
                _ = v24
            }
        case "hello":
            if false {
                for fi25 := 0; fi25 < 2; fi25++ {
                    var v26 uint = 18446744073709551615
                    var v27 uint32 = 362
                    fmt.Println(v1, v26, v19)
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            } else {
                var v28 float64 = 1e10
                v1 = v1
                var v29 uint16 = 925
                _ = v28
                _ = v29
            }
        default:
            for fi30 := 0; fi30 < 10; fi30++ {
                var v31 int8 = (v16 * 0)
                var v32 uint16 = uint16(v10)
                var v33 uint32 = uint32(v19)
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
        }
        var v34 uint32 = 906
        _ = v19
        _ = v34
    }
    fmt.Println(v16, v16)
    var v35 int8 = (v16 + genFn0(v18, 159, v1))
    if false {
        switch v13 {
        case "a":
            _ = len(v4)
        case "foo bar":
            fmt.Println(v1, v10, v10)
        case "":
            v5 = append(v5, -690)
        }
    } else {
        for fi36 := 0; fi36 < 2; fi36++ {
            var v37 int64 = 1
            v38 := GenS1{F0: -32768}
            _ = fi36
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 10; fi39++ {
            v40 := GenS1{F0: -1}
            _ = fi39
            _ = v40
        }
        v41 := GenS3{F0: 903}
        _ = v41
    }
    v10 = (^(v10))
    switch v13 {
    case "foo bar":
        var v42 float32 = float32(1.5)
        _ = v42
    case "a":
        v13 = (v13 + "\n\t")
    case "hello":
        var v43 int8 = 100
        _ = v43
    }
    var v44 float64 = 1.5
    var v45 int32 = (v1 * (v1 - v1))
    for fi46 := 0; fi46 < 3; fi46++ {
        fmt.Println(v35)
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}
