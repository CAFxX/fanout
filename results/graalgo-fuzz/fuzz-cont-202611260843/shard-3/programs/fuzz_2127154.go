package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }
type GenS2 struct { F0 int8; F1 int16; F2 int64 }
type GenS3 struct { F0 int; F1 int; F2 bool }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int32, p1 string) int32 {
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 int8 = 0
        _ = fi5
        _ = v6
    }
    var v7 float64 = float64(3.14159)
    if false {
        p1 = p1
        if true {
            switch ((((p1 + (("\n\t" + p1) + "foo bar")) + "a") + "") + "a") {
            case "hello":
                for fi8 := 0; fi8 < 6; fi8++ {
                    var v9 string = p1
                    v10 := []int{599, fi8, fi8, 4294967295}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            default:
                fmt.Println(p0)
            }
            v11 := GenS1{F0: 50, F1: ((((((p1 + p1) + "") + "") + p1) + p1) + p1), F2: 42}
            v12 := []int{-499}
            _ = v11
            _ = v12
        }
        var v13 uint32 = 4294967295
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v7
    return (p0 ^ 792)
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]bool)
    v3 := []uint32{}
    var v4 int16 = -1
    for fi14 := 0; fi14 < 17; fi14++ {
        switch (fi14 >> ((42) & 63)) {
        case 32767:
            var v15 uint32 = 976
            _ = v15
        case 128:
            switch "" {
            case "\n\t":
                if ((fi14 >> ((65535) & 63)) != -1) {
                    var v16 uint32 = 0
                    _ = v16
                }
            case "a":
                var v17 float64 = (1e10 - (((1e10 * 0.0) - 1.5) / -2.25))
                _ = v17
            case "foo bar":
                var v18 uint = 2
                _ = v18
            }
        case -607:
            fmt.Println(fi14, v4)
        }
        v4 += (v4 | v4)
        var v19 uint16 = 42
        _ = fi14
        _ = v19
    }
    v20 := GenS2{F0: 0, F1: v4, F2: -814}
    v21 := GenS1{F0: 2, F1: "", F2: 18446744073709551615}
    v22 := GenS3{F0: -755, F1: 159, F2: (0 <= 719)}
    v2[42] = true
    v4 = (-(v4))
    var v23 int64 = -9223372036854775808
    v23 = (^(v23))
    v4 = (v4 + (v4 << ((432) & 63)))
    v23 = -9223372036854775808
    v24 := v2[890]
    fmt.Println(v4, v4)
    if true {
        switch "hello" {
        case "foo bar":
            _ = len(v2)
        case "hello":
            v25 := GenS4{F0: 93, F1: v23}
            _ = v25
        }
    } else {
        if len(v1) > 0 {
            v1[0] = ("foo bar" + "\n\t")
        }
        v3 = append(v3, 29)
    }
    if (!true) {
        v4 -= v4
    } else {
        fmt.Println(v23, v23)
        switch ("" + ("\n\t" + "")) {
        case "foo bar":
            fmt.Println(v24, v4)
        case "a":
            for fi26 := 0; fi26 < 9; fi26++ {
                switch fi26 {
                case 42:
                    for fi27 := 0; fi27 < 6; fi27++ {
                        var v28 float64 = 3.14159
                        var v29 string = "a"
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                default:
                    var v30 int64 = (v23 - 1)
                    _ = v30
                }
                for fi31 := 0; fi31 < 9; fi31++ {
                    _ = len(v2)
                    if v24 {
                        v32 := GenS5{F0: (^(v23))}
                        _ = v32
                    }
                    v24 = false
                    _ = fi31
                }
                _ = fi26
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
