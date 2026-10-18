package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint }
type GenS2 struct { F0 bool; F1 uint32; F2 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 0.1
    v3 := make(map[uint32]uint32)
    var v4 string = ("hello" + "hello")
    var v5 uint32 = 65535
    v6 := []int{}
    if false {
        var v7 float64 = (v2 - (v2 + v2))
        switch 4294967295 {
        case 128:
            for fi8 := 0; fi8 < 12; fi8++ {
                v4 = "foo bar"
                v6 = append(v6, 723)
                v4 = (v4 + "a")
                _ = fi8
            }
        case -9223372036854775808:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 uint16 = 2
                v1[v4] = (v5 ^ 670)
                _ = fi9
                _ = v10
            }
        }
        _ = v7
    }
    switch v4 {
    case "a":
        var v11 uint32 = (v5 & v5)
        _ = v11
    case "":
        if false {
            v12 := GenS1{F0: -128, F1: (v5 << ((1) & 63)), F2: 287}
            _ = v12
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v2 -= (-2.25 * (v2 / v2))
        }
    case "\n\t":
        if (v4 < "hello") {
            switch 9223372036854775807 {
            case 32768:
                fmt.Println(v5)
            case 128:
                v1[("" + (("foo bar" + v4) + v4))] = 42
            case -2147483648:
                v4 = (v4 + v4)
            default:
                var v13 uint = 411
                _ = v13
            }
            if (1.5 <= 1.5) {
                v14 := GenS1{F0: 110, F1: v5, F2: 700}
                v15 := v3[v5]
                var v16 uint32 = 2147483647
                _ = v14
                _ = v15
                _ = v16
            }
            v5 -= 40
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        if true {
            if true {
                _ = len(v1)
                _ = len(v1)
            }
        } else {
            for fi18 := 0; fi18 < 5; fi18++ {
                _ = len(v1)
                v19 := GenS1{F0: 100, F1: 151, F2: 42}
                _ = fi18
                _ = v19
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        var v20 uint32 = v5
        _ = fi17
        _ = v20
    }
    var v21 uint8 = 19
    _ = len(v6)
    if (!(!((((v4 + ("\n\t" + v4)) + v4) + ("" + "")) > (("hello" + v4) + "")))) {
        if (0 != 2147483647) {
            var v22 string = v4
            for fi23 := 0; fi23 < 2; fi23++ {
                v24 := GenS2{F0: (2147483647 <= 0), F1: v5, F2: 42}
                fmt.Println(v2)
                v6 = append(v6, fi23)
                _ = fi23
                _ = v24
            }
            var v25 float64 = float64(float32(float64(((-2.25 + 3.14159) / -2.25))))
            _ = v22
            _ = v25
        } else {
            fmt.Println(v5)
            var v26 uint32 = v5
            _ = v26
        }
        var v27 uint64 = uint64(v21)
        _ = v27
    } else {
        var v28 uint32 = v5
        delete(v1, "\n\t")
        _ = v28
    }
    var v29 int64 = 499
    fmt.Println(v2, v21)
    v21 -= (v21 >> ((42) & 63))
    if (42 != -1) {
        var v30 int64 = v29
        v31 := GenS2{F0: false, F1: (v5 / ((v5) | 1)), F2: -70}
        _ = v30
        _ = v31
    } else {
        v2 += (v2 * float64(3.14159))
        v4 = (v4 + "foo bar")
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
