package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 bool; F2 uint }
func genFn0(p0 uint32, p1 uint16) int64 {
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := []int{fi5, -588, (-(fi5))}
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 uint32 = (p0 >> ((184) & 63))
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(p0, p0, fi9)
        _ = fi9
    }
    p0 -= (p0 - 445)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 bool, p1 int) uint64 {
    p1 = p1
    v10 := GenS1{F0: p1}
    v11 := []int{p1, p1}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 295
}

func main() {
    var v1 uint16 = 425
    var v2 int16 = 1
    v3 := make(map[string]bool)
    var v4 int8 = -1
    delete(v3, "foo bar")
    switch "a" {
    case "hello":
        var v12 bool = (929 != 42)
        _ = v12
    case "":
        switch ("a" + "foo bar") {
        case "":
            var v13 float64 = ((float64(3.14159) - (3.14159 - 3.14159)) * (3.14159 * float64(3.14159)))
            _ = v13
        case "a":
            v14 := GenS2{F0: 18446744073709551615}
            _ = v14
        case "foo bar":
            if (-1 > v2) {
                var v15 float32 = (((0.1 * (-2.25 / 1e10)) * (1e10 / 3.14159)) / -2.25)
                v2 += -1
                v15 = v15
                _ = v15
            }
        default:
            for fi16 := 0; fi16 < 7; fi16++ {
                v1 += 32767
                v2 -= v2
                v17 := GenS2{F0: 4294967295}
                _ = fi16
                _ = v17
            }
        }
    default:
        for fi18 := 0; fi18 < 15; fi18++ {
            if false {
                var v19 uint8 = 0
                _ = v19
            }
            _ = fi18
        }
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := GenS3{F0: uint32(v1), F1: -1}
        v4 = (v4 / ((42) | 1))
        _ = fi20
        _ = v21
    }
    if (((v2 * 409) == (v2 << ((835) & 63))) && false) {
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 int = int(v4)
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 14; fi24++ {
            switch ("a" + "hello") {
            case "foo bar":
                for fi25 := 0; fi25 < 7; fi25++ {
                    var v26 int32 = -245
                    _ = fi25
                    _ = v26
                }
            default:
                if false {
                    for fi27 := 0; fi27 < 3; fi27++ {
                        v3["foo bar"] = false
                        v28 := GenS3{F0: 65535, F1: v4}
                        _ = fi27
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 17; fi29++ {
                        v30 := GenS2{F0: 763}
                        v1 = v1
                        v2 = (v2 | v2)
                        fmt.Println(v4)
                        _ = fi29
                        _ = v30
                    }
                    v31 := GenS3{F0: 2, F1: 60}
                    _ = v31
                } else {
                    for fi32 := 0; fi32 < 19; fi32++ {
                        var v33 uint64 = 4294967295
                        _ = fi32
                        _ = v33
                    }
                    v34 := GenS4{F0: fi24}
                    _ = v34
                }
            }
            _ = fi24
        }
        v1 = v1
    }
    switch 32767 {
    case 9223372036854775807:
        v1 = v1
    case 308:
        v35 := GenS1{F0: 32768}
        _ = v35
    case -2147483648:
        v36 := []int{}
        _ = v36
    }
    var v37 int16 = -297
    if (((^(v2)) > v2) && false) {
        var v38 uint8 = 255
        _ = v38
    }
    if true {
        v1 += 869
        v4 += v4
        v39 := GenS5{F0: -1, F1: false, F2: 200}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
