package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint64) uint8 {
    switch 128 {
    case 32767:
        switch ("" + "") {
        case "\n\t":
            p1 -= (p1 * (p1 * (p1 * (p1 & 740))))
        default:
            var v6 int64 = 287
            _ = v6
        }
    case -9223372036854775808:
        for fi7 := 0; fi7 < 6; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                fmt.Println(p0, fi7, p0)
                var v9 string = ("a" + "a")
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    case 4294967295:
        p1 -= 1
    }
    if (42 < 1) {
        var v10 uint = 65535
        v11 := GenS1{F0: 1, F1: -228}
        _ = v10
        _ = v11
    }
    p0 += (p0 >> ((42) & 63))
    p1 += (p1 - (^(p1)))
    _ = p0
    _ = p1
    return 109
}

func main() {
    v1 := []int{1}
    var v2 int32 = 728
    v3 := make(map[string]bool)
    var v4 uint16 = 130
    var v5 uint64 = 341
    if true {
        v12 := GenS2{F0: 1}
        v4 = 42
        _ = v12
    }
    switch "" {
    case "a":
        if true {
            v13 := v3["\n\t"]
            _ = v13
        }
    default:
        fmt.Println(v4)
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        v5 += v5
        var v15 float64 = (-2.25 * 3.14159)
        v16 := GenS2{F0: 49}
        _ = fi14
        _ = v15
        _ = v16
    }
    if true {
        var v17 uint = 2
        _ = v17
    }
    if ((true && (true && false)) && (42 >= 118)) {
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 int32 = (v2 >> ((600) & 63))
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 uint8 = 127
                var v22 uint8 = v21
                v22 += (v21 + genFn0(800, v5))
                _ = fi20
                _ = v21
                _ = v22
            }
            if ((0.0 / 1e10) != 3.14159) {
                var v23 int16 = -920
                _ = v23
            } else {
                v2 -= v19
            }
            _ = fi18
            _ = v19
        }
    } else {
        _ = len(v1)
        fmt.Println(v2)
        var v24 int16 = 106
        _ = v24
    }
    fmt.Println(v5, v4)
    for fi25 := 0; fi25 < 3; fi25++ {
        delete(v3, "a")
        if true {
            for fi26 := 0; fi26 < 6; fi26++ {
                v27 := GenS1{F0: v2, F1: 42}
                if false {
                    v28 := GenS1{F0: v2, F1: -32768}
                    var v29 uint = 1
                    fmt.Println(v2, v2)
                    _ = v28
                    _ = v29
                }
                var v30 int8 = 42
                v5 -= v5
                _ = fi26
                _ = v27
                _ = v30
            }
            switch int(v5) {
            case 9223372036854775807:
                var v31 int64 = 0
                _ = v31
            }
        }
        var v32 uint8 = 127
        _ = fi25
        _ = v32
    }
    _ = len(v3)
    fmt.Println(v4, v5, v5)
    v4 -= v4
    delete(v3, ("a" + "\n\t"))
    var v33 float32 = 0.0
    v1 = append(v1, 343)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
