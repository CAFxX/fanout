package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS4 struct { F0 int }

func main() {
    v1 := []int{-9223372036854775808, 32767}
    var v2 string = "foo bar"
    v3 := []int{1, 128, 128, 128}
    v4 := make(map[string]bool)
    var v5 string = (("a" + v2) + v2)
    v6 := make(map[uint32]bool)
    if true {
        v7 := GenS1{F0: 0, F1: 1}
        _ = v7
    }
    if (0 != 127) {
        if (2147483647 > 1) {
            var v8 int64 = 0
            if (false && false) {
                v5 = "hello"
            } else {
                var v9 int16 = -32768
                var v10 uint32 = 2147483648
                _ = v10
                _ = v9
            }
            _ = v8
        }
        var v11 float32 = float32(float64((0.1 - (0.1 + float32((1.5 * (((0.1 * (1.5 / float64(1.5))) + 1e10) / (-2.25 + float64(float32(((0.1 + 0.1) / -2.25)))))))))))
        var v12 uint = 4294967295
        _ = v11
        _ = v12
    } else {
        v13 := GenS2{F0: 1000}
        var v14 string = (v5 + (v5 + v5))
        _ = v13
        _ = v14
    }
    if (1e10 != -2.25) {
        var v15 int64 = 1
        v4[v5] = (381 < 383)
        var v16 bool = true
        _ = v15
        _ = v16
    }
    _ = len(v6)
    _ = len(v1)
    switch (v5 + ((v2 + ((v5 + v2) + "hello")) + v2)) {
    case "":
        switch 983 {
        case 32768:
            v17 := GenS2{F0: 42}
            _ = v17
        case 2147483647:
            fmt.Println(v5, v2)
        default:
            for fi18 := 0; fi18 < 20; fi18++ {
                for fi19 := 0; fi19 < 19; fi19++ {
                    v2 = "foo bar"
                    v2 = (("hello" + v2) + (((v5 + v5) + v5) + v2))
                    v20 := GenS3{F0: 469, F1: 38, F2: 42}
                    var v21 int8 = -117
                    v22 := GenS4{F0: (fi18 * (fi19 + 2147483647))}
                    fmt.Println(fi18, fi18)
                    fmt.Println(fi19)
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                _ = len(v3)
                _ = fi18
            }
        }
    case "a":
        for fi23 := 0; fi23 < 4; fi23++ {
            _ = len(v4)
            _ = fi23
        }
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = -2147483648
        }
    }
    var v24 string = v5
    var v25 int8 = -128
    v24 = v24
    var v26 bool = (2 == 65535)
    if v26 {
        v25 -= v25
        if (true || (v25 <= v25)) {
            if v26 {
                for fi27 := 0; fi27 < 5; fi27++ {
                    if ((!((1 > 1) && v26)) || (42 != 484)) {
                        v28 := GenS1{F0: (-(fi27)), F1: 6}
                        _ = v28
                    }
                    var v29 uint64 = 18446744073709551615
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi27
                    _ = v29
                }
                v24 = v5
            } else {
                for fi30 := 0; fi30 < 16; fi30++ {
                    _ = len(v4)
                    var v31 bool = (1 == 1)
                    _ = fi30
                    _ = v31
                }
                v26 = false
            }
            v25 -= v25
        }
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        var v33 int = 9223372036854775807
        v33 = (v33 - v33)
        var v34 int8 = v25
        _ = fi32
        _ = v33
        _ = v34
    }
    delete(v6, 748)
    v35 := GenS4{F0: 42}
    _ = len(v4)
    v24 = v2
    fmt.Println(v24, v24)
    fmt.Println(v5, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
