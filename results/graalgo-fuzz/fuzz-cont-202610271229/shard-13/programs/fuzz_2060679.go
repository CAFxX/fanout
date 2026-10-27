package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0(p0 string) int32 {
    if ((1.5 * 0.1) <= (0.1 * 3.14159)) {
        var v7 uint16 = 84
        var v8 uint16 = (v7 << ((385) & 63))
        v9 := GenS1{F0: p0, F1: true, F2: 42}
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 uint16 = 75
        var v11 int64 = 0
        _ = v10
        _ = v11
    }
    var v12 int = 9223372036854775807
    fmt.Println(v12, v12)
    _ = p0
    _ = v12
    return 0
}
func genFn1(p0 string, p1 int32) int16 {
    var v13 int16 = -32768
    if (false && (437 >= 4294967295)) {
        v14 := GenS1{F0: "a", F1: (true || false), F2: 0}
        if (!true) {
            var v15 uint16 = 65535
            _ = v15
        }
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := []int{(fi16 >> ((42) & 63)), fi16}
            _ = fi16
            _ = v17
        }
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v13
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    var v4 float32 = ((3.14159 * (0.1 / -2.25)) + 3.14159)
    v5 := []string{("\n\t" + "\n\t"), (("hello" + "\n\t") + ("a" + (("" + "\n\t") + ("" + "a"))))}
    v6 := make(map[int]bool)
    if (100 > -38) {
        delete(v3, (((("hello" + "foo bar") + ("a" + "\n\t")) + "foo bar") + "hello"))
    } else {
        switch "hello" {
        case "foo bar":
            v4 = float32(0.1)
        case "hello":
            switch (("a" + "") + "\n\t") {
            case "foo bar":
                v4 = v4
            case "hello":
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 uint64 = 0
                    v4 = float32(1e10)
                    _ = fi18
                    _ = v19
                }
            default:
                var v20 float64 = (((((-2.25 + -2.25) - 3.14159) / 1.5) + (0.0 * -2.25)) - (0.1 + 1e10))
                _ = v20
            }
        }
        if true {
            v4 = (v4 - ((float32(3.14159) - 1e10) + 1.5))
            if (518 > 2) {
                v21 := GenS2{F0: 2, F1: 2}
                _ = v21
            }
            v4 += (((v4 / float32((-2.25 * 3.14159))) * v4) - float32(float64(0.0)))
        }
        var v22 bool = (true && false)
        _ = v22
    }
    var v23 uint = 0
    switch "\n\t" {
    case "a":
        var v24 int16 = -46
        _ = v24
    }
    var v25 int32 = 0
    v26 := GenS3{F0: 42, F1: 100}
    fmt.Println(v23, v4, v4)
    var v27 int16 = -243
    if (((true || false) || false) && (517 < 255)) {
        v23 = (v23 + 1)
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
        v28 := GenS3{F0: 1, F1: 0}
        _ = v28
    } else {
        fmt.Println(v23, v25, v23)
        v4 -= (1.5 - 1.5)
        for fi29 := 0; fi29 < 17; fi29++ {
            v5 = append(v5, "\n\t")
            _ = fi29
        }
    }
    v5 = append(v5, "")
    v30 := GenS3{F0: 953, F1: -19}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
