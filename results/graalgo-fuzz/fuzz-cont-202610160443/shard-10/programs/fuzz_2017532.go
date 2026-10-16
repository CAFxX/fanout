package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int }
func genFn0(p0 int64) int8 {
    switch -9223372036854775808 {
    case 4294967295:
        for fi5 := 0; fi5 < 20; fi5++ {
            if true {
                for fi6 := 0; fi6 < 4; fi6++ {
                    var v7 bool = true
                    _ = fi6
                    _ = v7
                }
            } else {
                p0 = 513
                var v8 uint32 = 2147483647
                for fi9 := 0; fi9 < 15; fi9++ {
                    var v10 bool = false
                    fmt.Println(v8, v8)
                    var v11 int16 = -186
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = v8
            }
            _ = fi5
        }
    case 2147483647:
        fmt.Println(p0)
    }
    if (false || false) {
        v12 := GenS1{F0: (true && false)}
        p0 += p0
        _ = v12
    } else {
        if false {
            for fi13 := 0; fi13 < 15; fi13++ {
                p0 -= p0
                _ = fi13
            }
            for fi14 := 0; fi14 < 18; fi14++ {
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 int8 = 127
                    _ = fi15
                    _ = v16
                }
                v17 := GenS2{F0: 13}
                var v18 uint8 = 255
                _ = fi14
                _ = v17
                _ = v18
            }
        }
        if true {
            v19 := GenS2{F0: 1000}
            fmt.Println(p0)
            v20 := GenS3{F0: 100, F1: 128}
            _ = v19
            _ = v20
        }
        var v21 uint32 = 2
        _ = v21
    }
    v22 := []int{4294967295, 676, 32767, -1}
    p0 = 9223372036854775807
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    _ = p0
    _ = v22
    return 45
}
func genFn1() string {
    var v23 int8 = -1
    fmt.Println(v23)
    _ = v23
    return ""
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int8 = -1
    var v3 int64 = v1
    v4 := []uint32{65535, 589}
    var v24 int32 = -173
    var v25 int32 = 2147483647
    var v26 uint8 = 1
    v3 = 9223372036854775807
    fmt.Println(v1, v25, v1)
    var v27 int = 9223372036854775807
    v28 := GenS3{F0: (^(v2)), F1: 2147483647}
    var v29 float64 = (0.0 * 0.1)
    if (-2.25 <= v29) {
        fmt.Println(v26, v24, v26)
        var v30 int = v27
        _ = v30
    }
    var v31 uint64 = 1
    v32 := GenS2{F0: 640}
    switch v27 {
    case -9223372036854775808:
        fmt.Println(v3)
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        if len(v4) > 0 {
            v4[0] = 2147483648
        }
        var v34 uint16 = 0
        v26 += v26
        _ = fi33
        _ = v34
    }
    v31 = (v31 ^ 2)
    var v35 uint16 = 0
    v24 += v24
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v4
}
