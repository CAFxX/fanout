package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 string; F1 int32; F2 int64 }
type GenS4 struct { F0 uint; F1 uint; F2 int32 }
func genFn0(p0 string, p1 uint32, p2 uint8) uint64 {
    p1 = (p1 & 2147483647)
    for fi4 := 0; fi4 < 2; fi4++ {
        for fi5 := 0; fi5 < 9; fi5++ {
            v6 := []int{(-(fi5)), (fi5 + (fi5 & (fi4 | (-(fi5))))), (-(fi5)), 1}
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: 255, F1: 42, F2: 42}
        _ = fi4
        _ = v7
    }
    p0 = (p0 + ((p0 + p0) + p0))
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 uint8) uint8 {
    for fi8 := 0; fi8 < 12; fi8++ {
        if (false && (-2.25 != (1e10 - (0.1 - (float64(0.0) / 1.5))))) {
            p0 = 1
        }
        fmt.Println(p0, fi8)
        p0 = p0
        _ = fi8
    }
    var v9 float64 = 1.5
    switch -150 {
    case 9223372036854775807:
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := []int{int(p0)}
            _ = fi10
            _ = v11
        }
    case 1:
        v12 := []int{261, 32767}
        _ = v12
    case 42:
        v9 = float64(0.0)
    default:
        var v13 bool = (358 != 354)
        _ = v13
    }
    p0 = 42
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 string = "a"
            _ = fi15
            _ = v16
        }
        v17 := GenS1{F0: 255, F1: 42, F2: 323}
        _ = fi14
        _ = v17
    }
    _ = p0
    _ = v9
    return (p0 / (((p0 + (p0 >> ((4294967295) & 63)))) | 1))
}

func main() {
    v1 := []uint32{65535, 2147483648, 0, 1}
    var v2 int32 = 570
    v3 := make(map[string]int)
    v18 := GenS1{F0: 255, F1: 1, F2: 229}
    var v19 uint32 = 1
    var v20 string = ("\n\t" + "\n\t")
    v2 -= 0
    switch "" {
    case "foo bar":
        var v21 uint16 = 975
        _ = v21
    case "hello":
        v22 := GenS2{F0: 127, F1: 32767}
        _ = v22
    case "a":
        fmt.Println(v19)
    }
    fmt.Println(v2, v19)
    var v23 int16 = 0
    v24 := GenS2{F0: 42, F1: 0}
    var v25 int32 = int32(v19)
    fmt.Println(v19, v20, v23)
    switch v20 {
    case "a":
        fmt.Println(v25)
    case "foo bar":
        switch (v20 + "hello") {
        case "foo bar":
            var v26 int16 = 32767
            _ = v26
        case "a":
            for fi27 := 0; fi27 < 10; fi27++ {
                switch "\n\t" {
                case "a":
                    for fi28 := 0; fi28 < 12; fi28++ {
                        var v29 int64 = -1
                        var v30 uint16 = 451
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                }
                _ = fi27
            }
        case "\n\t":
            v31 := GenS3{F0: v20, F1: (v25 << ((4294967295) & 63)), F2: -1}
            _ = v31
        }
    case "\n\t":
        for fi32 := 0; fi32 < 12; fi32++ {
            v1 = append(v1, v19)
            var v33 uint8 = 68
            var v34 uint8 = genFn1(v33)
            _ = fi32
            _ = v33
            _ = v34
        }
    default:
        fmt.Println(v20, v20, v2)
    }
    v35 := GenS4{F0: 255, F1: 2, F2: v2}
    v36 := GenS4{F0: 18446744073709551615, F1: 597, F2: (-(v25))}
    fmt.Println(v2)
    fmt.Println(v23)
    var v37 uint16 = 0
    v3[""] = -2147483648
    var v38 int32 = -742
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
