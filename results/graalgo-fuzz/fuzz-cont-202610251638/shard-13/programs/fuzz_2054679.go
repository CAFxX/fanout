package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint8; F2 string }
type GenS3 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 uint16; F2 int32 }

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]uint32)
    v3 := []string{("a" + "a"), "", "hello", ("a" + (("" + "\n\t") + ("" + "")))}
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]bool)
    var v6 uint16 = 1
    v7 := GenS1{F0: -9223372036854775808}
    v4[2147483647] = 42
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int64 = 0
        switch 0 {
        case 484:
            v10 := GenS2{F0: 42, F1: 128, F2: "foo bar"}
            _ = v10
        case 0:
            var v11 float32 = -2.25
            _ = v11
        }
        v12 := GenS3{F0: 42, F1: 0, F2: 42}
        _ = fi8
        _ = v12
        _ = v9
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            fmt.Println(v1)
            fmt.Println(fi13, v1)
            _ = fi14
        }
        _ = fi13
    }
    if false {
        var v15 float64 = 3.14159
        switch "" {
        case "\n\t":
            if (((v15 / 1.0) != v15) || ((((float64(1e10) - 0.1) / v15) / v15) == (v15 * 3.14159))) {
                for fi16 := 0; fi16 < 3; fi16++ {
                    v15 = (float64((float32(v15) * (3.14159 - 0.1))) + -2.25)
                    fmt.Println(fi16, v1, v6)
                    _ = fi16
                }
            }
        case "a":
            var v17 uint = 98
            _ = v17
        }
        _ = v15
    }
    if false {
        fmt.Println(v6)
    } else {
        v18 := v2[4294967295]
        v19 := GenS4{F0: 620, F1: "hello"}
        fmt.Println(v18)
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        if len(v3) > 0 {
            v3[0] = "a"
        }
        v21 := GenS2{F0: -128, F1: (v1 >> ((414) & 63)), F2: "hello"}
        _ = fi20
        _ = v21
    }
    if true {
        for fi22 := 0; fi22 < 4; fi22++ {
            _ = len(v3)
            v1 = 62
            _ = fi22
        }
    } else {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    }
    switch 920 {
    case 869:
        v1 -= v1
    }
    var v23 float64 = -2.25
    for fi24 := 0; fi24 < 20; fi24++ {
        if false {
            var v25 int64 = 9223372036854775807
            for fi26 := 0; fi26 < 10; fi26++ {
                v27 := v2[258]
                fmt.Println(v25, v1)
                _ = fi26
                _ = v27
            }
            _ = v25
        } else {
            v28 := GenS5{F0: 9223372036854775807}
            _ = v28
        }
        _ = fi24
    }
    v29 := GenS6{F0: "", F1: v6, F2: -531}
    v30 := GenS6{F0: (("foo bar" + (((("" + "a") + "") + "\n\t") + "foo bar")) + "\n\t"), F1: (v6 + 812), F2: -950}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!true) {
        var v31 uint16 = v6
        _ = v31
    }
    fmt.Println(v1, v6, v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
