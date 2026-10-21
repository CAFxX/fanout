package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint8; F1 string; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int8; F2 int8 }
type GenS5 struct { F0 uint64; F1 int; F2 uint8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8 }
func genFn0(p0 int) uint64 {
    p0 = (p0 | p0)
    var v6 string = (((("" + "foo bar") + "foo bar") + "a") + "foo bar")
    switch (p0 ^ p0) {
    case 678:
        for fi7 := 0; fi7 < 4; fi7++ {
            p0 -= (p0 << ((2) & 63))
            _ = fi7
        }
    case 42:
        for fi8 := 0; fi8 < 19; fi8++ {
            v6 = v6
            fmt.Println(fi8, fi8, p0)
            v9 := GenS1{F0: 18446744073709551615, F1: 65535, F2: 666}
            _ = fi8
            _ = v9
        }
    case 310:
        v10 := GenS2{F0: 255, F1: ("foo bar" + v6), F2: 0}
        _ = v10
    }
    _ = p0
    _ = v6
    return 1
}
func genFn1(p0 string) int32 {
    v11 := GenS3{F0: 32767}
    fmt.Println(p0)
    v12 := GenS4{F0: 1, F1: -83, F2: -42}
    p0 = p0
    if (false || ((float64((-2.25 * (-2.25 / float32((((-2.25 + 3.14159) + (float64(3.14159) / (0.1 + 1e10))) + (0.0 + float64(0.0))))))) > 1.5) && true)) {
        v13 := GenS5{F0: 4294967295, F1: 127, F2: 189}
        fmt.Println(p0)
        _ = v13
    } else {
        v14 := GenS3{F0: 255}
        _ = v14
    }
    _ = p0
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 int32 = 0
    var v2 int16 = -1
    var v3 uint32 = 1
    v4 := []uint32{}
    v5 := []string{"hello", (("hello" + "foo bar") + "a"), "hello", "foo bar"}
    v1 -= v1
    fmt.Println(v3, v3, v3)
    var v15 uint32 = (v3 >> ((4294967295) & 63))
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 18; fi17++ {
            v18 := GenS4{F0: (v2 + 0), F1: 42, F2: 127}
            var v19 bool = (fi17 == (fi16 << ((2) & 63)))
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    var v20 int16 = (v2 / ((v2) | 1))
    v21 := GenS6{F0: 737}
    if true {
        var v22 uint64 = genFn0(535)
        v2 = v2
        _ = v22
    }
    var v23 uint64 = 136
    v5 = append(v5, "")
    if (v20 < (v20 << ((4294967295) & 63))) {
        v4 = append(v4, v15)
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 uint32 = v15
        var v26 int8 = 127
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 float32 = float32((-2.25 / -2.25))
            var v29 float64 = 0.1
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        v3 -= 65535
        _ = fi30
    }
    v31 := GenS2{F0: uint8(v23), F1: "\n\t", F2: (v1 >> ((255) & 63))}
    fmt.Println(v1, v15, v2)
    switch "hello" {
    case "foo bar":
        fmt.Println(v3, v2)
    case "a":
        var v32 int32 = 1
        _ = v32
    case "":
        for fi33 := 0; fi33 < 5; fi33++ {
            v34 := GenS7{F0: 0}
            var v35 int = fi33
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    var v36 int8 = -128
    var v37 int8 = int8(v23)
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
