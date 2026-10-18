package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint8; F2 string }
type GenS3 struct { F0 int64; F1 uint32 }

func main() {
    var v1 int8 = -1
    var v2 int32 = -2147483648
    var v3 int32 = (v2 / ((0) | 1))
    v4 := make(map[int]string)
    v5 := []int{42}
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(v3)
        v4[(fi6 & fi6)] = ("foo bar" + "\n\t")
        var v7 uint32 = 705
        _ = fi6
        _ = v7
    }
    v2 += (v2 >> ((4294967295) & 63))
    v1 += 127
    if (!true) {
        switch "" {
        case "":
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 uint8 = 42
                v1 = (v1 >> ((588) & 63))
                v2 += v2
                _ = fi8
                _ = v9
            }
        case "a":
            switch ("\n\t" + "") {
            case "hello":
                v10 := GenS1{F0: ((429 == 241) || true)}
                _ = v10
            default:
                v3 = v3
            }
        default:
            var v11 int8 = v1
            _ = v11
        }
        v1 -= v1
    }
    v3 = (v2 - (v2 >> ((80) & 63)))
    if (780 > 0) {
        if true {
            if true {
                if true {
                    var v12 int = 32767
                    _ = v12
                }
            } else {
                var v13 int = -209
                for fi14 := 0; fi14 < 4; fi14++ {
                    var v15 uint64 = 380
                    _ = fi14
                    _ = v15
                }
                v3 += (v2 & -2147483648)
                _ = v13
            }
            fmt.Println(v1, v3, v3)
        }
    } else {
        v16 := GenS2{F0: 18446744073709551615, F1: 0, F2: "\n\t"}
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 int8 = (v1 ^ 1)
        for fi19 := 0; fi19 < 15; fi19++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    switch 32768 {
    case -2147483648:
        v1 -= (-(v1))
    default:
        var v20 uint32 = 274
        _ = v20
    }
    v21 := GenS1{F0: false}
    if (-2147483648 < 4294967295) {
        if ((160 > (v3 - v2)) && true) {
            v4[978] = "a"
        } else {
            var v22 float32 = 3.14159
            for fi23 := 0; fi23 < 2; fi23++ {
                var v24 int = int(v2)
                var v25 float32 = v22
                _ = fi23
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 5; fi26++ {
                fmt.Println(v3, v3)
                _ = fi26
            }
            _ = v22
        }
        var v27 string = "foo bar"
        v27 = ((v27 + v27) + (v27 + v27))
        _ = v27
    }
    v28 := v4[-1]
    v5 = append(v5, 83)
    switch 32768 {
    case 513:
        v4[-539] = v28
    }
    if (506 > (^(v2))) {
        v2 = v3
        var v29 float32 = 0.1
        v30 := GenS3{F0: 999, F1: 2147483647}
        _ = v29
        _ = v30
    }
    v1 = (v1 ^ -4)
    switch 42 {
    case 128:
        var v31 int16 = -1
        _ = v31
    default:
        _ = len(v5)
    }
    v5 = append(v5, int(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
