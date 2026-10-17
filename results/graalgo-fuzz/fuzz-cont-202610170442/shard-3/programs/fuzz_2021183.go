package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int32; F1 string; F2 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint) uint64 {
    v6 := []int{961}
    switch 127 {
    case 32768:
        v6 = append(v6, 9223372036854775807)
    case -1:
        var v7 int64 = 896
        _ = v7
    case 365:
        fmt.Println(p0)
    }
    if (-2.25 <= 1.5) {
        switch "foo bar" {
        case "foo bar":
            var v8 int64 = 42
            _ = v8
        case "a":
            var v9 int8 = 8
            _ = v9
        default:
            if len(v6) > 0 {
                v6[0] = 0
            }
        }
    }
    p0 = 65535
    _ = p0
    _ = v6
    return 42
}

func main() {
    var v1 uint16 = 400
    var v2 float32 = -2.25
    v3 := make(map[int]bool)
    var v4 uint64 = 496
    var v5 int16 = 102
    v2 -= (float32(1.5) * v2)
    v5 += v5
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 int64 = 9223372036854775807
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        delete(v3, (fi12 | fi12))
        _ = fi12
    }
    _ = len(v3)
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := v3[(fi13 << ((201) & 63))]
        if true {
            if true {
                v15 := GenS1{F0: 0, F1: ((1e10 + v2) < 1.5)}
                _ = v15
            } else {
                v5 = v5
            }
            v16 := GenS1{F0: 100, F1: ((fi13 ^ fi13) == fi13)}
            _ = v16
        } else {
            _ = len(v3)
        }
        _ = fi13
        _ = v14
    }
    switch -453 {
    case 4294967295:
        var v17 float32 = float32(((1e10 * ((-2.25 + 1e10) - -2.25)) * (0.1 * 0.1)))
        _ = v17
    case 0:
        for fi18 := 0; fi18 < 3; fi18++ {
            v2 = v2
            if (v5 <= -32768) {
                var v19 uint64 = (v4 >> ((937) & 63))
                fmt.Println(v4)
                var v20 uint32 = uint32(v1)
                _ = v19
                _ = v20
            }
            _ = fi18
        }
    }
    v21 := GenS2{F0: 536, F1: "foo bar", F2: 0}
    for fi22 := 0; fi22 < 7; fi22++ {
        v23 := []int{1, (fi22 - 4294967295), 584, 4294967295}
        var v24 int8 = 100
        _ = fi22
        _ = v23
        _ = v24
    }
    if false {
        if (true && (uint64(v1) == (v4 / ((v4) | 1)))) {
            switch -4 {
            case 892:
                v1 = v1
            case 4294967295:
                fmt.Println(v2, v5)
            case -1:
                v4 += 2
            default:
                switch "\n\t" {
                case "\n\t":
                    for fi25 := 0; fi25 < 16; fi25++ {
                        delete(v3, fi25)
                        var v26 uint64 = uint64(v1)
                        _ = fi25
                        _ = v26
                    }
                case "hello":
                    v27 := GenS2{F0: -1, F1: ((("a" + (((("foo bar" + (("" + "a") + (("\n\t" + ("" + "\n\t")) + "hello"))) + "hello") + "a") + "hello")) + ("hello" + "")) + "hello"), F2: -956}
                    _ = v27
                case "a":
                    var v28 float64 = 3.14159
                    _ = v28
                }
            }
            if false {
                var v29 uint16 = (v1 + (v1 % (((v1 / ((65535) | 1))) | 1)))
                for fi30 := 0; fi30 < 3; fi30++ {
                    if false {
                        v5 -= (v5 + 438)
                        if (42 >= 220) {
                            v31 := GenS3{F0: uint64(v1)}
                            _ = v31
                        }
                        var v32 string = ""
                        _ = v32
                    } else {
                        delete(v3, (-(fi30)))
                        v33 := GenS4{F0: true, F1: uint16(v4)}
                        _ = v33
                    }
                    switch (fi30 % ((9223372036854775807) | 1)) {
                    case 32767:
                        v34 := []int{fi30, (-(fi30)), -9223372036854775808, (fi30 % (((-(fi30))) | 1))}
                        _ = v34
                    case 42:
                        for fi35 := 0; fi35 < 8; fi35++ {
                            if (1 <= 2) {
                                var v36 float64 = -2.25
                                _ = v36
                            }
                            var v37 uint = 65535
                            v4 -= 42
                            _ = fi35
                            _ = v37
                        }
                    default:
                        v5 = v5
                    }
                    v38 := GenS1{F0: 42, F1: true}
                    _ = fi30
                    _ = v38
                }
                _ = v29
            }
        }
    }
    if false {
        v39 := GenS3{F0: 18446744073709551615}
        _ = v39
    }
    fmt.Println(v1)
    switch "hello" {
    case "foo bar":
        var v40 float64 = 3.14159
        _ = v40
    case "\n\t":
        v41 := []int{42, 32768, -376, int(v4)}
        _ = v41
    }
    var v42 uint64 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
