package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int; F1 uint8; F2 uint64 }
func genFn0(p0 uint, p1 uint16, p2 int64) uint64 {
    p0 = 418
    v5 := []int{2147483647, -150, -2147483648, -23}
    var v6 int8 = 32
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int64, p1 int8) uint32 {
    v7 := []int{9223372036854775807, 1, 128, -225}
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            fmt.Println(fi9, p1, p0)
            if (("a" + ("foo bar" + "\n\t")) != ("foo bar" + "a")) {
                var v10 uint64 = 18446744073709551615
                v10 += (^(v10))
                var v11 int = fi8
                var v12 bool = (0 >= 1)
                v11 = (fi8 ^ fi8)
                _ = v10
                _ = v11
                _ = v12
            }
            p1 += 127
            _ = fi9
        }
        _ = fi8
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v7
    return 2147483648
}

func main() {
    var v1 uint8 = 238
    var v2 int32 = 2147483647
    var v3 uint16 = 2
    var v4 int32 = v2
    v13 := []int{32767, 127}
    v2 = 1
    v3 += (v3 | (v3 >> ((966) & 63)))
    v1 = 255
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 int16 = 902
        v2 = (v4 + v2)
        fmt.Println(v1, v15)
        _ = fi14
        _ = v15
    }
    if (-1 <= -32768) {
        if (((v1 ^ v1) >= v1) || (false && (702 != int(v1)))) {
            var v16 int64 = 1
            _ = v16
        } else {
            v3 = 626
            if false {
                if false {
                    v17 := GenS1{F0: -1, F1: 2}
                    fmt.Println(v1)
                    _ = v17
                }
                v1 += v1
                var v18 uint = 2
                _ = v18
            }
            if (42 > 791) {
                v19 := GenS1{F0: 42, F1: (v3 - 1)}
                v20 := GenS2{F0: 39, F1: v1, F2: genFn0(580, v3, -1)}
                var v21 float32 = 1.5
                _ = v19
                _ = v20
                _ = v21
            }
        }
        _ = len(v13)
        v3 = v3
    } else {
        fmt.Println(v4)
        _ = len(v13)
        v3 -= v3
    }
    v3 -= 981
    if (("\n\t" + "") == "\n\t") {
        if (4294967295 > 4294967295) {
            v22 := GenS2{F0: 0, F1: 26, F2: 2}
            _ = v22
        }
        if (!true) {
            _ = len(v13)
            _ = len(v13)
            switch -9223372036854775808 {
            case 9223372036854775807:
                var v23 float64 = -2.25
                _ = v23
            }
        } else {
            var v24 uint = 18446744073709551615
            v24 += 41
            v25 := GenS2{F0: 42, F1: 127, F2: 4294967295}
            _ = v24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        v27 := GenS1{F0: 0, F1: v3}
        if len(v13) > 0 {
            v13[0] = fi26
        }
        fmt.Println(v1, fi26, v2)
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        v3 = (^(v3))
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
