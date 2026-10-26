package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint8; F1 int; F2 uint8 }
type GenS3 struct { F0 int; F1 uint16; F2 bool }
type GenS4 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS5 struct { F0 string; F1 uint32; F2 string }
func genFn0(p0 int8, p1 uint16) uint32 {
    p1 = 0
    v7 := GenS1{F0: 1, F1: 128}
    var v8 int = 2147483647
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint, p1 uint, p2 int16) uint8 {
    var v9 uint64 = 4294967295
    switch -2147483648 {
    case 527:
        if true {
            p1 = 209
        } else {
            switch "hello" {
            case "foo bar":
                v9 += (v9 / (((v9 ^ (v9 >> (((p0 / ((386) | 1))) & 63)))) | 1))
            case "\n\t":
                fmt.Println(p2, v9, v9)
            }
            var v10 int64 = -844
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := GenS1{F0: (v10 & -9223372036854775808), F1: fi11}
                switch ("a" + "\n\t") {
                case "foo bar":
                    if (!true) {
                        v13 := []int{fi11, 127, 220, (fi11 >> (((p0 * p0)) & 63))}
                        v10 += (v10 * v10)
                        _ = v13
                    } else {
                        var v14 uint = p0
                        v15 := GenS2{F0: 1, F1: fi11, F2: 128}
                        _ = v14
                        _ = v15
                    }
                case "hello":
                    var v16 int32 = 0
                    _ = v16
                default:
                    p1 = (p1 >> (((^(p0))) & 63))
                }
                _ = fi11
                _ = v12
            }
            _ = v10
        }
    case -646:
        if true {
            p2 -= p2
            var v17 int64 = -158
            _ = v17
        } else {
            var v18 int32 = 1
            fmt.Println(p2)
            _ = v18
        }
    case 842:
        switch 990 {
        case -2147483648:
            var v19 int8 = -9
            _ = v19
        case -520:
            v20 := []int{}
            _ = v20
        default:
            var v21 int8 = 42
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 127
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (576 <= 561)
    v3 := make(map[int]uint32)
    var v4 int32 = 42
    v5 := make(map[string]string)
    v6 := make(map[int]bool)
    for fi22 := 0; fi22 < 12; fi22++ {
        v23 := GenS2{F0: 117, F1: (fi22 ^ 2147483647), F2: 1}
        var v24 uint8 = 1
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 20; fi26++ {
            v27 := []int{}
            if (((float32(1e10) * (3.14159 * 1e10)) / -2.25) <= 1.5) {
                fmt.Println(fi26, v2)
                v2 = (0.0 > 3.14159)
            }
            _ = fi26
            _ = v27
        }
        if false {
            v28 := GenS3{F0: -9223372036854775808, F1: 1, F2: v2}
            v29 := GenS3{F0: 772, F1: 0, F2: (!(!v2))}
            _ = v28
            _ = v29
        } else {
            for fi30 := 0; fi30 < 6; fi30++ {
                v1[genFn0(50, 907)] = fi30
                fmt.Println(v2, v4, fi25)
                _ = fi30
            }
            if (134 <= 4294967295) {
                var v31 int16 = -32768
                for fi32 := 0; fi32 < 17; fi32++ {
                    var v33 string = "\n\t"
                    _ = fi32
                    _ = v33
                }
                _ = v31
            }
            fmt.Println(v2)
        }
        _ = fi25
    }
    switch ("foo bar" + "hello") {
    case "foo bar":
        v4 += v4
    case "\n\t":
        var v34 uint64 = 18446744073709551615
        _ = v34
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        switch ("hello" + "") {
        case "\n\t":
            var v36 uint32 = 4294967295
            _ = v36
        }
        _ = fi35
    }
    v5["foo bar"] = "a"
    if v2 {
        switch "" {
        case "":
            for fi37 := 0; fi37 < 16; fi37++ {
                v4 += v4
                _ = fi37
            }
        case "\n\t":
            var v38 string = ("" + "")
            _ = v38
        }
    }
    v39 := GenS4{F0: 2, F1: -9223372036854775808, F2: 42}
    if v2 {
        if (144 < 197) {
            v40 := GenS4{F0: 128, F1: 0, F2: 2}
            _ = v40
        }
        if v2 {
            if (1 < 255) {
                if v2 {
                    v2 = (171 != 255)
                    var v41 uint16 = 32767
                    _ = v41
                } else {
                    v2 = (9223372036854775807 <= 646)
                    var v42 uint16 = 740
                    var v43 float32 = 1.5
                    _ = v42
                    _ = v43
                }
                var v44 uint64 = 4294967295
                v2 = (true || true)
                _ = v44
            } else {
                v45 := GenS2{F0: 128, F1: -862, F2: 255}
                var v46 float64 = float64((3.14159 - (((float32((1e10 - (3.14159 / ((3.14159 - 1.5) * 1.5)))) * (1.5 / 3.14159)) - (0.1 - ((float32(1e10) - (float32(-2.25) + 0.0)) / -2.25))) - 1e10)))
                fmt.Println(v2, v4)
                _ = v45
                _ = v46
            }
            var v47 float64 = ((1.5 - -2.25) / 1.5)
            v2 = ("" >= ("foo bar" + "a"))
            _ = v47
        } else {
            v48 := GenS5{F0: ("\n\t" + "a"), F1: 2, F2: "\n\t"}
            var v49 int32 = v4
            _ = v48
            _ = v49
        }
    }
    v4 = -1
    v2 = true
    var v50 uint = 1
    for fi51 := 0; fi51 < 5; fi51++ {
        v52 := GenS2{F0: 42, F1: (-(fi51)), F2: 0}
        _ = fi51
        _ = v52
    }
    var v53 uint32 = 1
    v54 := GenS1{F0: 9223372036854775807, F1: 2147483647}
    var v55 int = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v50
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
