package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint }
func genFn0(p0 uint16, p1 uint32) int64 {
    var v7 bool = true
    for fi8 := 0; fi8 < 7; fi8++ {
        if true {
            var v9 int64 = 725
            p0 -= p0
            v10 := GenS1{F0: 18446744073709551615, F1: 1, F2: 2}
            _ = v10
            _ = v9
        }
        var v11 uint16 = p0
        _ = fi8
        _ = v11
    }
    fmt.Println(v7, p1, v7)
    p0 += 334
    _ = p0
    _ = p1
    _ = v7
    return 793
}
func genFn1(p0 int8, p1 uint) uint64 {
    if (3.14159 >= 0.1) {
        fmt.Println(p0, p1)
        switch (("foo bar" + (("\n\t" + "hello") + ((("foo bar" + "\n\t") + ("foo bar" + "hello")) + ("" + "foo bar")))) + "hello") {
        case "foo bar":
            var v12 uint64 = 4294967295
            _ = v12
        default:
            switch 42 {
            case 128:
                v13 := []int{0}
                _ = v13
            case 2147483647:
                p0 = (p0 / ((p0) | 1))
            default:
                fmt.Println(p0)
            }
        }
    } else {
        v14 := []int{0, -2147483648, 127, 4294967295}
        v14 = append(v14, 32767)
        _ = v14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            switch "hello" {
            case "foo bar":
                var v17 int = fi16
                _ = v17
            case "hello":
                var v18 int32 = -337
                _ = v18
            case "a":
                var v19 int = (fi16 << ((p1) & 63))
                _ = v19
            }
            p0 = (p0 + (p0 - (p0 + -128)))
            _ = fi16
        }
        _ = fi15
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := GenS1{F0: 1, F1: 1, F2: 0}
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 uint8 = 127
            v25 := GenS1{F0: (p1 ^ p1), F1: -432, F2: 91}
            _ = fi23
            _ = v24
            _ = v25
        }
        for fi26 := 0; fi26 < 2; fi26++ {
            fmt.Println(p0, p1)
            _ = fi26
        }
        _ = fi20
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint32 = 186
    v2 := []int{815, 523}
    var v3 bool = (0.0 >= 1.5)
    var v4 bool = false
    var v5 uint8 = 42
    v6 := []int{32767, -9223372036854775808}
    v4 = v4
    v3 = (1 > 255)
    var v27 float64 = 0.1
    v1 = v1
    v5 = v5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = v1
    for fi28 := 0; fi28 < 11; fi28++ {
        v27 = v27
        _ = fi28
    }
    switch -1 {
    case 4294967295:
        var v29 int64 = 1
        _ = v29
    case 32767:
        v27 -= v27
    }
    fmt.Println(v1, v3, v3)
    if (284 <= 2) {
        if (v1 > 619) {
            for fi30 := 0; fi30 < 14; fi30++ {
                v31 := GenS1{F0: 90, F1: -2147483648, F2: 4294967295}
                if true {
                    fmt.Println(v5, v27, v1)
                    var v32 uint64 = 2
                    v4 = v4
                    _ = v32
                } else {
                    v5 = v5
                    fmt.Println(v4)
                }
                _ = fi30
                _ = v31
            }
            v27 -= 1e10
            for fi33 := 0; fi33 < 2; fi33++ {
                if (1.5 >= -2.25) {
                    v34 := GenS1{F0: 564, F1: 118, F2: 42}
                    _ = v34
                } else {
                    var v35 int16 = 0
                    _ = v35
                }
                _ = fi33
            }
        }
        switch 127 {
        case -9223372036854775808:
            v36 := GenS2{F0: 2, F1: 0, F2: 354}
            _ = v36
        case 4294967295:
            v37 := GenS2{F0: 745, F1: 0, F2: 65535}
            _ = v37
        default:
            v38 := GenS2{F0: 512, F1: -2147483648, F2: 42}
            _ = v38
        }
    } else {
        v4 = true
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
