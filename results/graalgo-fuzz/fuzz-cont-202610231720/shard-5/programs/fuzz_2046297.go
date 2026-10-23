package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 bool; F1 uint8; F2 int }
func genFn0(p0 string, p1 int32) uint32 {
    var v6 int32 = (p1 | (p1 | p1))
    var v7 uint = 641
    switch 42 {
    case 843:
        var v8 int = 128
        _ = v8
    default:
        for fi9 := 0; fi9 < 1; fi9++ {
            v7 -= 836
            _ = fi9
        }
    }
    v10 := GenS1{F0: 455}
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 int16 = 32767
    v2 := []uint32{}
    v3 := []uint32{2147483648, 0, 518}
    var v4 int32 = 0
    v5 := []uint32{548, 4294967295}
    if (!(true || false)) {
        switch 42 {
        case -2147483648:
            for fi11 := 0; fi11 < 10; fi11++ {
                v12 := GenS2{F0: genFn0((("" + (((("\n\t" + "hello") + "foo bar") + (("a" + "foo bar") + "a")) + "\n\t")) + ("a" + "a")), v4)}
                v13 := GenS2{F0: 297}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        var v14 float64 = -2.25
        v5 = append(v5, 0)
        _ = v14
    } else {
        v15 := GenS3{F0: 0, F1: 9223372036854775807}
        _ = v15
    }
    v16 := GenS1{F0: 42}
    switch 685 {
    case 32768:
        v17 := GenS4{F0: (v4 ^ (v4 - (v4 << ((42) & 63)))), F1: (v1 % ((v1) | 1))}
        _ = v17
    case 42:
        var v18 int64 = 9223372036854775807
        _ = v18
    case 1:
        v4 -= (-(v4))
    }
    v1 = (^(v1))
    for fi19 := 0; fi19 < 8; fi19++ {
        v4 += (v4 >> ((2) & 63))
        switch 9223372036854775807 {
        case -612:
            v1 += 187
        case 32767:
            if (fi19 >= fi19) {
                var v20 uint = 4294967295
                var v21 int64 = 9223372036854775807
                _ = v20
                _ = v21
            } else {
                v1 = v1
            }
        case 0:
            for fi22 := 0; fi22 < 11; fi22++ {
                _ = len(v3)
                _ = len(v2)
                _ = fi22
            }
        default:
            if len(v2) > 0 {
                v2[0] = 659
            }
        }
        _ = fi19
    }
    v23 := GenS5{F0: false, F1: 255, F2: 127}
    v3 = append(v3, 2147483647)
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 int32 = (v4 >> ((18446744073709551615) & 63))
        _ = fi24
        _ = v25
    }
    v1 = v1
    v1 = (v1 << ((255) & 63))
    var v26 float64 = ((0.0 * 1.5) / 1.0)
    switch "a" {
    case "a":
        fmt.Println(v1, v4, v26)
    default:
        v26 -= v26
    }
    switch -1 {
    case 1:
        var v27 float32 = (3.14159 / 1e10)
        _ = v27
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        for fi29 := 0; fi29 < 15; fi29++ {
            if true {
                var v30 string = ""
                var v31 float32 = float32(0.1)
                v31 = ((((v31 / (v31 - 0.1)) - (1.5 / v31)) * (v31 - 1.5)) * 1e10)
                if len(v2) > 0 {
                    v2[0] = 0
                }
                if true {
                    if false {
                        _ = len(v3)
                    }
                    var v32 uint16 = 1
                    _ = v32
                }
                _ = v30
                _ = v31
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi29
        }
        if false {
            var v33 bool = (true || (!false))
            v4 = v4
            _ = v33
        } else {
            fmt.Println(v26, v26, v1)
            var v34 float64 = v26
            _ = v34
        }
        if len(v5) > 0 {
            v5[0] = 2147483647
        }
        _ = fi28
    }
    v35 := GenS4{F0: 29, F1: (v1 % ((v1) | 1))}
    for fi36 := 0; fi36 < 15; fi36++ {
        v1 -= -32768
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
