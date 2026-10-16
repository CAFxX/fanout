package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
func genFn0(p0 int8, p1 uint16) string {
    if ((p0 != p0) && ((p0 ^ p0) == p0)) {
        switch 150 {
        case 32767:
            var v7 float64 = -2.25
            _ = v7
        case -9223372036854775808:
            p1 = (p1 - (p1 << ((1) & 63)))
        case 0:
            fmt.Println(p0)
        }
        v8 := GenS1{F0: 1, F1: 18446744073709551615}
        _ = v8
    } else {
        v9 := []int{}
        p0 = p0
        _ = len(v9)
        _ = v9
    }
    switch "a" {
    case "hello":
        var v10 float64 = ((float64((-2.25 + 3.14159)) + (3.14159 + 1e10)) / 1.0)
        _ = v10
    case "foo bar":
        v11 := []int{-825, 32767, 42}
        _ = v11
    default:
        switch 543 {
        case 914:
            v12 := []int{127, 336}
            _ = v12
        case 4294967295:
            var v13 int16 = 32767
            _ = v13
        case -981:
            v14 := GenS1{F0: 1, F1: 0}
            _ = v14
        default:
            var v15 int8 = p0
            _ = v15
        }
    }
    p1 = 255
    var v16 int8 = (p0 & (p0 | p0))
    _ = p0
    _ = p1
    _ = v16
    return "\n\t"
}

func main() {
    var v1 int32 = 2147483647
    var v2 float64 = 1e10
    var v3 uint = 1
    v4 := make(map[string]bool)
    v5 := make(map[string]int)
    v6 := []int{254, 127, 0, 357}
    fmt.Println(v1, v1)
    var v17 int = 2147483647
    var v18 bool = (false || false)
    var v19 uint = (v3 % ((749) | 1))
    v19 += v3
    for fi20 := 0; fi20 < 12; fi20++ {
        switch "" {
        case "\n\t":
            v3 = (v3 ^ v3)
        case "hello":
            var v21 uint8 = 127
            _ = v21
        }
        v22 := GenS2{F0: fi20, F1: 506}
        if true {
            var v23 uint = v19
            fmt.Println(v23, v3)
            _ = v23
        } else {
            fmt.Println(v2)
        }
        _ = fi20
        _ = v22
    }
    v24 := GenS1{F0: 946, F1: 356}
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 int = 644
        _ = fi25
        _ = v26
    }
    v2 = 3.14159
    var v27 int64 = 9223372036854775807
    if (240 >= 42) {
        switch (v17 | 4294967295) {
        case -450:
            v1 -= int32(v27)
        }
        v28 := GenS1{F0: 4294967295, F1: 1}
        _ = v28
    } else {
        var v29 uint64 = 1
        for fi30 := 0; fi30 < 12; fi30++ {
            v1 += (v1 % ((2147483647) | 1))
            v31 := GenS3{F0: v3}
            _ = fi30
            _ = v31
        }
        _ = v29
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        switch 9223372036854775807 {
        case -982:
            if v18 {
                v27 += -380
                fmt.Println(v18)
                var v33 uint16 = 32767
                _ = v33
            } else {
                v17 = 4294967295
            }
        case 1:
            v34 := GenS4{F0: (v17 - fi32)}
            _ = v34
        case 2147483647:
            if len(v6) > 0 {
                v6[0] = (-(v17))
            }
        default:
            fmt.Println(v2)
        }
        v35 := GenS3{F0: v3}
        fmt.Println(fi32)
        _ = fi32
        _ = v35
    }
    fmt.Println(v18, v3)
    v3 -= 2
    fmt.Println(v1)
    for fi36 := 0; fi36 < 14; fi36++ {
        for fi37 := 0; fi37 < 2; fi37++ {
            if (!(9 <= 100)) {
                var v38 int32 = (v1 | v1)
                _ = v38
            } else {
                var v39 string = "a"
                var v40 int64 = (v27 % (((v27 / ((v27) | 1))) | 1))
                _ = v39
                _ = v40
            }
            var v41 uint16 = 1
            var v42 uint = (v3 / (((^(v3))) | 1))
            fmt.Println(v1, v17, v42)
            switch (v17 / ((fi36) | 1)) {
            case 32767:
                v19 -= v42
            default:
                v5[("foo bar" + "hello")] = (v17 << ((uint(v41)) & 63))
            }
            var v43 string = "foo bar"
            _ = fi37
            _ = v41
            _ = v42
            _ = v43
        }
        _ = fi36
    }
    v27 = (v27 * 0)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
