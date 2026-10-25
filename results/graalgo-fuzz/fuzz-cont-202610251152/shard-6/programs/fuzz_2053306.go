package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int16; F2 int }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int16, p1 int64) uint16 {
    v7 := []int{-2147483648, -2147483648, -949, -1}
    if true {
        switch (("foo bar" + "") + "foo bar") {
        case "\n\t":
            var v8 uint8 = 42
            _ = v8
        default:
            fmt.Println(p1)
        }
        var v9 int8 = 106
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "hello"
    var v3 uint16 = 0
    v4 := []uint32{}
    var v5 uint16 = v3
    var v6 float32 = (1.5 + ((float32((1e10 - (((float64(float32(1e10)) - (0.1 + (1e10 / 0.1))) / 1.5) - (1.5 - 0.0)))) / 0.1) - (float32((3.14159 * 1e10)) - 0.1)))
    v10 := GenS1{F0: false, F1: 87, F2: uint32(v3)}
    if true {
        v5 = 65535
        var v11 bool = false
        _ = v11
    }
    v4 = append(v4, 2147483648)
    v12 := GenS1{F0: (!(!(false && true))), F1: 1, F2: 973}
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 int32 = 556
            v16 := GenS1{F0: ((3.14159 + 1e10) < 3.14159), F1: 42, F2: 459}
            v17 := GenS1{F0: (1e10 < float64(1.5)), F1: 831, F2: 2147483647}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    switch 42 {
    case -9223372036854775808:
        v2 = (v2 + v2)
    case 42:
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 uint = 0
            v5 = 659
            if (!false) {
                var v20 uint64 = 2
                v2 = v2
                _ = v20
            }
            _ = fi18
            _ = v19
        }
    case 4294967295:
        var v21 uint16 = v5
        _ = v21
    default:
        var v22 int64 = 703
        _ = v22
    }
    v23 := GenS1{F0: true, F1: 1, F2: 2}
    switch v2 {
    case "a":
        switch v2 {
        case "a":
            v3 = 32767
        case "hello":
            v24 := GenS1{F0: true, F1: 4294967295, F2: 4294967295}
            _ = v24
        }
    case "foo bar":
        if true {
            v25 := GenS2{F0: -2147483648}
            _ = v25
        }
    case "hello":
        var v26 int64 = 9223372036854775807
        _ = v26
    default:
        fmt.Println(v6, v5, v6)
    }
    if (!(1.5 > 1.5)) {
        var v27 string = ((v2 + v2) + "foo bar")
        _ = v27
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        if true {
            var v29 int16 = -530
            v3 += v3
            _ = v29
        } else {
            var v30 float32 = (v6 * (1.5 / 1.0))
            v6 = v6
            v2 = v2
            _ = v30
        }
        switch fi28 {
        case 107:
            var v31 uint8 = 255
            _ = v31
        default:
            var v32 float32 = (1e10 / (v6 - v6))
            _ = v32
        }
        for fi33 := 0; fi33 < 7; fi33++ {
            v34 := GenS1{F0: ((-557 > (^(fi28))) || false), F1: 18446744073709551615, F2: 42}
            v6 -= v6
            if (1 > 42) {
                switch "" {
                case "a":
                    v35 := GenS3{F0: 65535, F1: 706, F2: (^(fi28))}
                    _ = v35
                case "\n\t":
                    v5 = (v5 >> ((18446744073709551615) & 63))
                    v36 := GenS3{F0: 2147483648, F1: 42, F2: fi33}
                    _ = v36
                case "foo bar":
                    var v37 uint64 = 2
                    _ = v37
                default:
                    var v38 float32 = v6
                    _ = v38
                }
            } else {
                _ = len(v1)
                var v39 bool = true
                _ = v39
            }
            var v40 uint16 = (v3 - (v3 - v3))
            _ = fi33
            _ = v34
            _ = v40
        }
        _ = fi28
    }
    for fi41 := 0; fi41 < 10; fi41++ {
        for fi42 := 0; fi42 < 7; fi42++ {
            switch v2 {
            case "":
                _ = len(v4)
            case "\n\t":
                v5 = (v3 ^ v5)
            case "foo bar":
                v6 = -2.25
            default:
                v3 = (v5 * (v5 - (^(v3))))
                v43 := GenS4{F0: true, F1: 1}
                _ = v43
            }
            _ = fi42
        }
        v44 := GenS5{F0: 4294967295}
        _ = fi41
        _ = v44
    }
    v2 = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
