package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint }
type GenS2 struct { F0 uint16; F1 int32 }
func genFn0(p0 int16, p1 uint, p2 uint) uint32 {
    if (1 >= -128) {
        if (167 == 42) {
            fmt.Println(p2, p1, p2)
            p1 += p2
        }
        p0 = p0
        v4 := []int{32768, -2147483648}
        _ = v4
    }
    switch (("a" + "") + "\n\t") {
    case "":
        fmt.Println(p2, p0, p1)
    case "\n\t":
        var v5 float32 = float32(-2.25)
        _ = v5
    case "hello":
        for fi6 := 0; fi6 < 10; fi6++ {
            v7 := []int{}
            _ = fi6
            _ = v7
        }
    default:
        if (p1 > (p2 / (((p1 << ((119) & 63))) | 1))) {
            if false {
                for fi8 := 0; fi8 < 14; fi8++ {
                    if (true || ((((!true) || (0 != -128)) || false) || true)) {
                        var v9 int16 = p0
                        var v10 uint8 = 132
                        var v11 int64 = 1
                        _ = v10
                        _ = v11
                        _ = v9
                    }
                    for fi12 := 0; fi12 < 10; fi12++ {
                        p1 += (p1 - (p1 + 100))
                        fmt.Println(p1)
                        if false {
                            var v13 int32 = 0
                            v14 := GenS1{F0: -33, F1: 0, F2: p2}
                            _ = v13
                            _ = v14
                        }
                        _ = fi12
                    }
                    _ = fi8
                }
                for fi15 := 0; fi15 < 4; fi15++ {
                    v16 := GenS1{F0: 92, F1: 0, F2: p1}
                    _ = fi15
                    _ = v16
                }
            }
        }
    }
    if true {
        p0 -= -812
        v17 := []int{-900, 4294967295}
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}
func genFn1(p0 uint16, p1 string, p2 int32) int {
    v18 := []int{32767, 2147483647, 658, 32768}
    var v19 bool = (-1 < (p2 << ((518) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return -394
}

func main() {
    var v1 int16 = 50
    var v2 bool = (-1 != 0)
    var v3 string = (("" + "a") + ("hello" + "foo bar"))
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 int = fi20
        _ = fi20
        _ = v21
    }
    if (1 != 4294967295) {
        fmt.Println(v1, v2)
        switch 32768 {
        case 128:
            var v22 int = 32767
            _ = v22
        case -341:
            v23 := []int{}
            _ = v23
        case 4294967295:
            v1 = v1
        }
        switch 32768 {
        case 9223372036854775807:
            v1 = v1
        case -430:
            var v24 float64 = (((1.5 * 0.1) * 3.14159) * ((0.1 - 1.5) - (1.5 + (1e10 + 1e10))))
            _ = v24
        case 128:
            v3 = v3
        }
    }
    v1 = v1
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 int32 = -2147483648
        var v27 float32 = 0.0
        var v28 int16 = (^(v1))
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    v29 := []int{}
    switch 0 {
    case 127:
        for fi30 := 0; fi30 < 14; fi30++ {
            v1 -= (v1 | (v1 / ((0) | 1)))
            v3 = (((v3 + "foo bar") + (("a" + "foo bar") + v3)) + v3)
            _ = fi30
        }
    case 42:
        var v31 int32 = 1
        _ = v31
    case -9223372036854775808:
        var v32 int32 = 2147483647
        _ = v32
    }
    fmt.Println(v2)
    for fi33 := 0; fi33 < 7; fi33++ {
        var v34 float64 = 0.0
        _ = fi33
        _ = v34
    }
    switch 42 {
    case 32768:
        v35 := GenS1{F0: -107, F1: 768, F2: 65535}
        _ = v35
    case 9223372036854775807:
        v3 = (("foo bar" + (((v3 + (v3 + "foo bar")) + ("a" + v3)) + v3)) + "hello")
    }
    var v36 uint32 = 2
    for fi37 := 0; fi37 < 6; fi37++ {
        if v2 {
            switch genFn1(287, ("\n\t" + (v3 + v3)), 159) {
            case 2147483647:
                v38 := GenS1{F0: -74, F1: 876, F2: 188}
                _ = v38
            case -1:
                var v39 float32 = 3.14159
                _ = v39
            }
            fmt.Println(v1)
            v40 := GenS2{F0: 42, F1: 42}
            _ = v40
        }
        _ = fi37
    }
    v36 = v36
    var v41 uint32 = (v36 - 2147483648)
    for fi42 := 0; fi42 < 14; fi42++ {
        if false {
            v29 = append(v29, (fi42 & 0))
        }
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v36
    _ = v41
}
