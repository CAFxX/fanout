package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64, p1 uint64, p2 int8) uint32 {
    fmt.Println(p0, p0)
    p1 = (^(p1))
    v6 := GenS1{F0: (p2 - (-(p2)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483648
}
func genFn1(p0 int, p1 bool) uint {
    fmt.Println(p0, p1)
    p1 = p1
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int64 = 0
        fmt.Println(p0)
        v9 := GenS1{F0: -106}
        _ = fi7
        _ = v8
        _ = v9
    }
    switch ("\n\t" + "") {
    case "foo bar":
        p0 = p0
    default:
        switch (p0 + (p0 >> ((4294967295) & 63))) {
        case -2147483648:
            v10 := GenS1{F0: -1}
            _ = v10
        case 32768:
            p0 -= (p0 * (p0 >> ((0) & 63)))
        case 74:
            switch "foo bar" {
            case "foo bar":
                var v11 float32 = float32((0.0 - 1.5))
                _ = v11
            default:
                p1 = p1
            }
        default:
            if ((309 >= 293) || (1 == 2)) {
                p1 = (449 == 0)
                p1 = (((((0.0 / (1e10 * 0.1)) - float32((0.0 + 1e10))) * 1e10) * 3.14159) != (0.1 * float32(1e10)))
            }
        }
    }
    var v12 int16 = 1
    _ = p0
    _ = p1
    _ = v12
    return 4294967295
}

func main() {
    var v1 string = ("foo bar" + "\n\t")
    v2 := []int{}
    var v3 int32 = 1
    var v4 int = 127
    v5 := make(map[string]uint32)
    v1 = v1
    var v13 int8 = 100
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 int64 = -738
        var v16 uint8 = 102
        var v17 bool = (v16 == 210)
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    switch ("a" + v1) {
    case "\n\t":
        _ = len(v2)
    default:
        if (((127 != 255) || (32767 <= 2)) && (9223372036854775807 >= 575)) {
            fmt.Println(v3, v13, v4)
            for fi18 := 0; fi18 < 5; fi18++ {
                v3 += (v3 - 2147483647)
                _ = fi18
            }
            v3 = -398
        }
    }
    if ((((true && true) && false) || (!(891 >= (v4 & (v4 - v4))))) || false) {
        switch v4 {
        case -1:
            if (v13 != -128) {
                fmt.Println(v4, v13, v13)
                v2 = append(v2, v4)
            }
        case 648:
            var v19 uint32 = 2147483647
            _ = v19
        }
        fmt.Println(v13)
        var v20 uint16 = 255
        _ = v20
    } else {
        v21 := GenS2{F0: (v4 << ((65535) & 63))}
        v3 += v3
        _ = len(v5)
        _ = v21
    }
    fmt.Println(v4)
    for fi22 := 0; fi22 < 12; fi22++ {
        v3 += v3
        v13 -= v13
        var v23 int = v4
        _ = fi22
        _ = v23
    }
    if (93 < 0) {
        if (1 >= 0) {
            var v24 uint32 = 605
            fmt.Println(v24)
            for fi25 := 0; fi25 < 17; fi25++ {
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 float64 = 1.5
                    var v28 int16 = 0
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                v24 = (v24 << ((536) & 63))
                if true {
                    var v29 int16 = -32768
                    v30 := GenS3{F0: 1}
                    v31 := GenS2{F0: v4}
                    _ = v29
                    _ = v30
                    _ = v31
                } else {
                    v13 = v13
                }
                _ = fi25
            }
            _ = v24
        }
        v1 = (v1 + v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
