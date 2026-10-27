package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint8) int64 {
    if false {
        p0 -= (p0 & 56)
        for fi4 := 0; fi4 < 18; fi4++ {
            fmt.Println(fi4)
            _ = fi4
        }
    } else {
        v5 := []int{0}
        p0 += 175
        _ = len(v5)
        _ = v5
    }
    p0 = p0
    var v6 uint32 = 642
    _ = p0
    _ = v6
    return 404
}

func main() {
    v1 := []int{0}
    v2 := []uint32{2147483648}
    v3 := make(map[string]string)
    v7 := GenS1{F0: 936, F1: (42 < 18446744073709551615)}
    var v8 uint16 = 65535
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 bool = true
        _ = fi9
        _ = v10
    }
    var v11 uint64 = 285
    var v12 uint16 = (v8 ^ (v8 + (v8 % ((v8) | 1))))
    var v13 uint64 = (v11 >> ((2) & 63))
    if (genFn0(227) > -1) {
        fmt.Println(v12, v12, v12)
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := v3["foo bar"]
            var v16 uint32 = 0
            v17 := GenS1{F0: 1, F1: true}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        fmt.Println(v12)
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        _ = len(v1)
        v11 = 18446744073709551615
        _ = fi18
    }
    switch 4294967295 {
    case 9223372036854775807:
        if false {
            v8 += 32767
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case 326:
        var v19 int8 = 2
        _ = v19
    case 720:
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 int = 1
            for fi22 := 0; fi22 < 11; fi22++ {
                var v23 int = fi20
                v1 = append(v1, 249)
                _ = len(v3)
                _ = fi22
                _ = v23
            }
            if len(v1) > 0 {
                v1[0] = -806
            }
            _ = fi20
            _ = v21
        }
    default:
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 bool = false
            _ = fi24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        for fi27 := 0; fi27 < 1; fi27++ {
            if (false || ((0.0 - 1.5) == float32(3.14159))) {
                fmt.Println(v12, v12)
            } else {
                var v28 string = "a"
                var v29 int8 = 47
                v8 += v8
                _ = v28
                _ = v29
            }
            _ = fi27
        }
        _ = fi26
    }
    var v30 uint32 = 143
    var v31 uint32 = v30
    switch 9223372036854775807 {
    case 32767:
        if (-2.25 <= 1e10) {
            if (-9223372036854775808 > 42) {
                switch -1 {
                case 0:
                    v12 -= v12
                default:
                    v31 = (^(v31))
                }
                var v32 int = 4294967295
                _ = v32
            }
            if false {
                var v33 float64 = float64(((3.14159 * -2.25) - (0.1 / 1e10)))
                _ = v33
            }
            fmt.Println(v12, v30)
        } else {
            switch "" {
            case "hello":
                for fi34 := 0; fi34 < 2; fi34++ {
                    var v35 float64 = 1.5
                    v30 += 2147483648
                    _ = fi34
                    _ = v35
                }
            case "foo bar":
                fmt.Println(v8, v12)
            default:
                fmt.Println(v30, v31)
            }
            v36 := GenS1{F0: 1, F1: false}
            for fi37 := 0; fi37 < 16; fi37++ {
                var v38 bool = (128 >= 1)
                _ = fi37
                _ = v38
            }
            _ = v36
        }
    case 32768:
        switch "hello" {
        case "\n\t":
            var v39 int32 = int32(v31)
            _ = v39
        case "a":
            var v40 float32 = -2.25
            _ = v40
        }
    case 1:
        var v41 int32 = -987
        _ = v41
    }
    if (-2147483648 >= -917) {
        v42 := GenS2{F0: ""}
        fmt.Println(v13, v13, v13)
        var v43 uint16 = v8
        _ = v42
        _ = v43
    }
    v44 := GenS3{F0: 0}
    for fi45 := 0; fi45 < 2; fi45++ {
        for fi46 := 0; fi46 < 4; fi46++ {
            v47 := GenS4{F0: uint32(v11)}
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v44
    _ = v7
    _ = v8
}
