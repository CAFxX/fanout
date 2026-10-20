package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint64 }
func genFn0(p0 uint8, p1 uint16) int64 {
    var v4 int = 32767
    fmt.Println(v4, p0)
    v5 := GenS1{F0: 818}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 9223372036854775807
}
func genFn1(p0 uint64, p1 int) int64 {
    switch p1 {
    case 32768:
        var v6 int32 = 0
        _ = v6
    case 9223372036854775807:
        p1 = p1
    case 128:
        var v7 int8 = -35
        _ = v7
    default:
        var v8 float32 = 0.1
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := GenS1{F0: 42}
        for fi11 := 0; fi11 < 4; fi11++ {
            if ((((float64(((-2.25 / -2.25) + (1e10 * 1e10))) / 3.14159) / 0.1) - 3.14159) == float64(0.1)) {
                var v12 int8 = 42
                _ = v12
            } else {
                p1 = 32767
                p1 -= fi11
            }
            v13 := GenS1{F0: 42}
            _ = fi11
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    if true {
        v14 := GenS1{F0: 0}
        _ = v14
    } else {
        v15 := GenS2{F0: 42, F1: (p1 / ((p1) | 1))}
        _ = v15
    }
    if (1 == 0) {
        switch (p1 << ((831) & 63)) {
        case -759:
            v16 := GenS1{F0: 42}
            _ = v16
        }
    } else {
        p1 += (p1 | (p1 >> ((255) & 63)))
        var v17 bool = true
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := []int{(^(p1))}
            p1 = p1
            var v20 uint16 = 0
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = v17
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    var v3 uint64 = 18446744073709551615
    v21 := GenS2{F0: 127, F1: 881}
    v22 := GenS3{F0: 1}
    switch (("" + ("a" + ((("foo bar" + ("foo bar" + "\n\t")) + "a") + "foo bar"))) + ("\n\t" + "hello")) {
    case "":
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 uint32 = 2
            _ = fi23
            _ = v24
        }
    case "a":
        for fi25 := 0; fi25 < 7; fi25++ {
            v3 -= 0
            _ = fi25
        }
    case "\n\t":
        fmt.Println(v3)
    default:
        v26 := GenS4{F0: 2, F1: true, F2: v3}
        _ = v26
    }
    v3 = (v3 & (v3 * 324))
    fmt.Println(v3)
    fmt.Println(v3)
    _ = len(v1)
    for fi27 := 0; fi27 < 5; fi27++ {
        switch "\n\t" {
        case "hello":
            var v28 int16 = -256
            _ = v28
        default:
            v3 = 42
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            fmt.Println(fi27)
            _ = fi29
        }
        var v30 bool = (!(0 != 32767))
        _ = fi27
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
