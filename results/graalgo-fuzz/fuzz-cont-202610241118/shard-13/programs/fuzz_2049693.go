package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int8; F1 int8; F2 uint16 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint8; F1 bool; F2 bool }
func genFn0(p0 string, p1 int32, p2 int64) bool {
    v4 := GenS1{F0: -9223372036854775808}
    v5 := GenS1{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return true
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 552
    v3 := make(map[int]string)
    v6 := GenS2{F0: 841, F1: 43, F2: 128}
    v2 -= v2
    v2 += v2
    v2 -= (-(v2))
    v7 := GenS1{F0: -693}
    v2 = v2
    switch "\n\t" {
    case "a":
        for fi8 := 0; fi8 < 5; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                if true {
                    var v10 uint8 = 2
                    _ = v10
                }
                _ = fi9
            }
            _ = fi8
        }
    default:
        if true {
            switch ("a" + "foo bar") {
            case "\n\t":
                var v11 bool = true
                _ = v11
            }
            var v12 int16 = -338
            fmt.Println(v12, v2)
            _ = v12
        } else {
            v13 := []int{1, 128, -1, 32767}
            _ = len(v1)
            v2 -= -893
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            fmt.Println(v2, v2, v2)
            v16 := GenS3{F0: 6, F1: 93, F2: 940}
            _ = fi15
            _ = v16
        }
        delete(v3, (fi14 >> ((504) & 63)))
        if (!false) {
            v2 = 0
        }
        _ = fi14
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v2)
    case "hello":
        v17 := GenS4{F0: "\n\t", F1: 42}
        _ = v17
    case "":
        switch ("" + "") {
        case "hello":
            _ = len(v3)
        case "a":
            v18 := []int{2147483647, -288, 32767, 4294967295}
            _ = v18
        default:
            v19 := []int{2147483647}
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        if (1 < 2) {
            v21 := v1[0]
            _ = v21
        } else {
            v22 := GenS5{F0: 100, F1: 42, F2: 32767}
            _ = v22
        }
        var v23 string = ("a" + "\n\t")
        if true {
            if true {
                var v24 bool = (!((328 == 474) || false))
                v2 = (v2 % ((-967) | 1))
                v24 = (3.14159 > 1e10)
                _ = v24
            } else {
                v23 = v23
            }
        }
        _ = fi20
        _ = v23
    }
    v3[-934] = ("hello" + ("foo bar" + (("a" + (("hello" + (("a" + "\n\t") + "hello")) + "")) + "hello")))
    v25 := GenS2{F0: 0, F1: 85, F2: 243}
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 bool = true
        for fi28 := 0; fi28 < 8; fi28++ {
            v29 := GenS3{F0: 235, F1: 4294967295, F2: 255}
            if (2 < 18446744073709551615) {
                v2 += -2147483648
                v2 = 646
                v2 += (v2 & 0)
            } else {
                v30 := GenS6{F0: 18446744073709551615, F1: 2}
                fmt.Println(fi26, fi28)
                v2 += v2
                v27 = v27
                _ = v30
            }
            var v31 uint = 18446744073709551615
            _ = fi28
            _ = v29
            _ = v31
        }
        _ = fi26
        _ = v27
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        for fi33 := 0; fi33 < 7; fi33++ {
            v2 = (v2 ^ v2)
            var v34 float32 = (3.14159 * 1e10)
            v35 := v1[31]
            var v36 bool = (!true)
            if (3.14159 != ((1e10 + 0.0) / 1.5)) {
                if (0 == 255) {
                    var v37 uint16 = 0
                    _ = v37
                }
                v38 := GenS7{F0: 42, F1: false, F2: (!(577 == 27))}
                _ = v38
            }
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
        }
        _ = fi32
    }
    switch -815 {
    case 32768:
        var v39 int64 = -48
        _ = v39
    case 33:
        v2 += v2
    default:
        var v40 uint8 = 255
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v6
    _ = v7
}
