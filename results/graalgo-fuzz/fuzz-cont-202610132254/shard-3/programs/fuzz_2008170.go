package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
func genFn0(p0 uint64, p1 uint8, p2 uint16) int64 {
    var v7 uint32 = 257
    v8 := GenS1{F0: 42, F1: -946}
    p0 = 42
    var v9 uint8 = p1
    if (-994 <= 814) {
        if (0.0 > float32(1e10)) {
            v10 := []int{-9223372036854775808}
            if ((!false) && (!((p0 * uint64(v7)) < p0))) {
                v10 = append(v10, 0)
                var v11 int16 = 1
                _ = len(v10)
                _ = v11
            }
            _ = v10
        }
        var v12 uint8 = 24
        if false {
            v13 := []int{-785, 32768, -61, -2147483648}
            var v14 string = "hello"
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = v13
            _ = v14
        } else {
            if false {
                fmt.Println(p0, p0)
                var v15 string = "hello"
                v16 := []int{1, 128, -2147483648}
                _ = v15
                _ = v16
            }
            var v17 int = -1
            _ = v17
        }
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 9223372036854775807
}
func genFn1() uint8 {
    if (!(0 <= 2147483648)) {
        if (true || (2 == 676)) {
            var v18 bool = (0 == 65535)
            v18 = v18
            v18 = v18
            _ = v18
        }
    } else {
        var v19 uint = 18446744073709551615
        _ = v19
    }
    if true {
        if true {
            for fi20 := 0; fi20 < 7; fi20++ {
                var v21 string = "foo bar"
                for fi22 := 0; fi22 < 16; fi22++ {
                    v21 = v21
                    _ = fi22
                }
                _ = fi20
                _ = v21
            }
        }
        v23 := []int{127, 4294967295, 32768, 0}
        var v24 uint8 = 128
        _ = v23
        _ = v24
    }
    var v25 float32 = 0.1
    for fi26 := 0; fi26 < 15; fi26++ {
        var v27 int16 = 42
        fmt.Println(v25)
        v28 := []int{(-(fi26)), fi26}
        _ = fi26
        _ = v27
        _ = v28
    }
    _ = v25
    return 128
}

func main() {
    var v1 uint8 = 42
    var v2 int16 = -1
    var v3 string = ""
    var v4 float32 = (3.14159 - 0.1)
    v5 := []int{-2147483648, 127, 0}
    v6 := make(map[uint32]int)
    v4 = 3.14159
    var v29 int16 = (^(v2))
    if len(v5) > 0 {
        v5[0] = 32767
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        var v31 int = -1
        for fi32 := 0; fi32 < 4; fi32++ {
            _ = len(v5)
            var v33 uint16 = 42
            _ = fi32
            _ = v33
        }
        _ = fi30
        _ = v31
    }
    fmt.Println(v4, v4)
    if (v1 != (v1 / (((v1 + (v1 * (v1 - v1)))) | 1))) {
        v34 := GenS1{F0: (v29 - (v2 % ((v2) | 1))), F1: int32(v1)}
        for fi35 := 0; fi35 < 18; fi35++ {
            if false {
                v4 -= (1.5 / 1.0)
            }
            for fi36 := 0; fi36 < 6; fi36++ {
                v2 -= v29
                var v37 float32 = (v4 - v4)
                _ = fi36
                _ = v37
            }
            fmt.Println(v4)
            _ = fi35
        }
        fmt.Println(v2, v1)
        _ = v34
    } else {
        if (2147483647 == 42) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v38 := GenS1{F0: 1000, F1: 1}
            _ = v38
        }
    }
    switch (v3 + (v3 + (v3 + v3))) {
    case "hello":
        var v39 int32 = -2147483648
        _ = v39
    default:
        var v40 int32 = 307
        _ = v40
    }
    if (!false) {
        var v41 uint = 0
        var v42 int16 = v29
        _ = v41
        _ = v42
    }
    delete(v6, 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
