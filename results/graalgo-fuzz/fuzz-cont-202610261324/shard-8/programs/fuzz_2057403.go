package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 string; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint16, p2 int) string {
    var v6 int8 = -119
    v7 := GenS1{F0: 4294967295, F1: 903}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return ""
}
func genFn1(p0 uint16, p1 string, p2 string) uint {
    v8 := []int{-9223372036854775808, 127, -1}
    if false {
        v9 := GenS1{F0: 1, F1: 2147483647}
        if (42 >= 127) {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            var v10 int = -1
            _ = v10
        } else {
            var v11 float32 = 0.1
            _ = v11
        }
        var v12 float32 = ((3.14159 * 1.5) - 0.1)
        _ = v12
        _ = v9
    } else {
        v13 := GenS1{F0: 562, F1: 674}
        _ = v13
    }
    var v14 bool = (91 < 670)
    var v15 float64 = float64(1.5)
    v16 := GenS2{F0: 18446744073709551615, F1: (!(((!(!v14)) && v14) && false))}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    _ = v8
    return 783
}

func main() {
    var v1 string = ""
    var v2 uint32 = 2
    v3 := make(map[string]int)
    var v4 int64 = -428
    var v5 int64 = (v4 >> ((0) & 63))
    var v17 int16 = -32768
    v1 = v1
    var v18 uint16 = 65535
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS2{F0: 113, F1: false}
        v21 := GenS3{F0: v1, F1: v2, F2: 414}
        _ = fi19
        _ = v20
        _ = v21
    }
    if true {
        v18 = v18
    }
    switch "\n\t" {
    case "hello":
        v22 := GenS2{F0: 0, F1: true}
        _ = v22
    case "foo bar":
        if false {
            delete(v3, (v1 + v1))
            var v23 string = v1
            v24 := v3[("" + "foo bar")]
            _ = v23
            _ = v24
        } else {
            v18 -= v18
            fmt.Println(v5, v5, v17)
        }
    case "":
        v5 = 859
    }
    fmt.Println(v4)
    var v25 uint16 = v18
    v4 += v5
    v26 := v3[(v1 + "a")]
    for fi27 := 0; fi27 < 8; fi27++ {
        fmt.Println(v18, v2, v1)
        _ = fi27
    }
    v28 := GenS4{F0: uint32(v18)}
    if (true || (float32(3.14159) != -2.25)) {
        if ((^(v2)) >= v2) {
            var v29 int64 = v5
            for fi30 := 0; fi30 < 16; fi30++ {
                var v31 int16 = (v17 << ((582) & 63))
                v32 := v3[(v1 + ("a" + "\n\t"))]
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = v29
        }
        v1 = "hello"
        v26 = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
