package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint16; F2 uint }
func genFn0(p0 int16, p1 int32) bool {
    var v6 uint32 = 330
    if (169 == 255) {
        var v7 uint = 65535
        _ = v7
    }
    v6 += 2
    _ = p0
    _ = p1
    _ = v6
    return true
}
func genFn1(p0 uint, p1 uint16) int8 {
    v8 := GenS1{F0: ((((("" + ("\n\t" + "hello")) + (("foo bar" + "hello") + "")) + (("foo bar" + "") + ("a" + ("foo bar" + ("hello" + "hello"))))) + "\n\t") + "foo bar"), F1: 44, F2: 42}
    v9 := GenS1{F0: "", F1: 227, F2: 375}
    v10 := GenS2{F0: 1}
    v11 := GenS3{F0: "hello", F1: 32767}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 uint8 = 128
    v4 := []uint32{2147483648, 42}
    var v5 int64 = 745
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 6; fi13++ {
            fmt.Println(fi12, v3, fi12)
            _ = fi13
        }
        for fi14 := 0; fi14 < 1; fi14++ {
            fmt.Println(v3, fi14, fi14)
            _ = len(v2)
            switch "foo bar" {
            case "":
                v5 -= v5
                delete(v2, "")
                v5 += v5
                v15 := GenS4{F0: -43}
                _ = v15
            case "hello":
                v16 := GenS2{F0: 42}
                fmt.Println(v5, fi12)
                _ = v16
            case "foo bar":
                var v17 int16 = -1
                _ = v17
            }
            _ = fi14
        }
        delete(v1, 127)
        _ = fi12
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        if ((-(v5)) > v5) {
            v2["a"] = 4294967295
        } else {
            v3 -= v3
        }
        v19 := GenS4{F0: 0}
        if (int8(v5) == -15) {
            _ = len(v4)
            v3 = 2
            v4 = append(v4, 42)
        } else {
            for fi20 := 0; fi20 < 17; fi20++ {
                v21 := GenS1{F0: "hello", F1: v3, F2: 18446744073709551615}
                _ = fi20
                _ = v21
            }
            v3 = (v3 * (v3 / (((v3 + (v3 / ((v3) | 1)))) | 1)))
            var v22 uint16 = 42
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    v2["hello"] = 2147483647
    v3 -= v3
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 uint16 = 42
        _ = fi23
        _ = v24
    }
    switch -560 {
    case 1:
        fmt.Println(v5, v3)
    case -2147483648:
        v5 = (v5 >> ((280) & 63))
    case 928:
        v5 = v5
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            if len(v4) > 0 {
                v4[0] = 0
            }
            v27 := GenS5{F0: ("\n\t" + ("\n\t" + "\n\t")), F1: 65535, F2: 18446744073709551615}
            var v28 uint16 = 65535
            _ = fi26
            _ = v27
            _ = v28
        }
        v3 = v3
        switch (fi25 << ((375) & 63)) {
        case 128:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi25
    }
    v3 = v3
    v4 = append(v4, 4294967295)
    _ = len(v4)
    var v29 float32 = (3.14159 * 1.5)
    var v30 int = -772
    var v31 uint32 = 2
    v32 := GenS4{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
