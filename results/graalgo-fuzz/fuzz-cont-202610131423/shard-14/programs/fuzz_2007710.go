package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint64 }

func main() {
    v1 := []int{-554, 32767, 200}
    var v2 uint16 = 65535
    var v3 int = 439
    v4 := []int{-4, 9223372036854775807}
    var v5 int32 = -2147483648
    v6 := make(map[int]string)
    if len(v4) > 0 {
        v4[0] = (-(v3))
    }
    fmt.Println(v3, v3, v2)
    v7 := GenS1{F0: true}
    delete(v6, (-(v3)))
    for fi8 := 0; fi8 < 12; fi8++ {
        if false {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            var v9 float64 = 0.1
            switch "\n\t" {
            case "":
                v2 += 1
            case "foo bar":
                v2 -= v2
            }
            _ = v9
        }
        _ = fi8
    }
    v2 -= (v2 / ((v2) | 1))
    fmt.Println(v5, v3, v5)
    switch (("" + (((("foo bar" + "foo bar") + "foo bar") + (("" + "a") + ("" + "a"))) + "\n\t")) + "") {
    case "a":
        switch v3 {
        case -402:
            v6[-297] = ""
        case 4294967295:
            var v10 int8 = -1
            _ = v10
        case 32767:
            v11 := GenS1{F0: (!(true && (v2 == (v2 % ((0) | 1)))))}
            _ = v11
        default:
            if false {
                v3 += -470
                var v12 int32 = (v5 | v5)
                _ = v12
            }
        }
    case "\n\t":
        var v13 float64 = 3.14159
        _ = v13
    case "":
        if (!(356 != 2147483647)) {
            v5 += v5
        } else {
            v14 := v6[(v3 >> ((4294967295) & 63))]
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v15 int = v3
            _ = v14
            _ = v15
        }
    default:
        if false {
            var v16 float64 = 0.0
            fmt.Println(v3)
            _ = v16
        } else {
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := GenS1{F0: true}
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 5; fi19++ {
                v2 = v2
                var v20 uint32 = 4294967295
                if true {
                    if true {
                        v21 := GenS1{F0: (((1.5 - ((float32(-2.25) + (0.0 + -2.25)) / (3.14159 / (3.14159 * 1e10)))) - 0.1) < 0.0)}
                        _ = v21
                    }
                } else {
                    var v22 float64 = 3.14159
                    v23 := GenS2{F0: 32767, F1: 133}
                    _ = v22
                    _ = v23
                }
                _ = fi19
                _ = v20
            }
        }
    }
    v3 += (v3 % ((-2147483648) | 1))
    for fi24 := 0; fi24 < 14; fi24++ {
        delete(v6, 870)
        _ = fi24
    }
    if true {
        var v25 int32 = int32(v2)
        var v26 uint64 = 42
        fmt.Println(v5)
        _ = v25
        _ = v26
    } else {
        _ = len(v6)
        fmt.Println(v3)
        if (9223372036854775807 > -1) {
            v3 = (v3 | -150)
            if (!(!(float64(1e10) != (-2.25 - 1e10)))) {
                _ = len(v4)
                var v27 int8 = 1
                _ = v27
            }
        } else {
            v5 = (^(v5))
            switch (v3 >> ((18446744073709551615) & 63)) {
            case -745:
                for fi28 := 0; fi28 < 17; fi28++ {
                    _ = len(v4)
                    if (!(-79 <= 42)) {
                        v5 += v5
                    }
                    var v29 uint16 = 0
                    _ = fi28
                    _ = v29
                }
            case -1:
                var v30 int64 = 42
                _ = v30
            }
        }
    }
    v3 = (v3 >> ((64) & 63))
    fmt.Println(v2, v5)
    for fi31 := 0; fi31 < 3; fi31++ {
        v6[fi31] = ""
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
