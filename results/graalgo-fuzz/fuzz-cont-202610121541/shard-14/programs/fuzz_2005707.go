package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
func genFn0(p0 string, p1 uint, p2 string) int {
    v7 := []int{42, -2147483648}
    p0 = ("a" + p2)
    p2 = p2
    fmt.Println(p0, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 127
}
func genFn1(p0 uint, p1 uint64, p2 uint16) uint64 {
    var v8 bool = true
    v9 := GenS1{F0: 912, F1: -128}
    p1 -= (p1 ^ p1)
    if true {
        v10 := []int{136, 4294967295, -9223372036854775808, 790}
        var v11 int32 = -2147483648
        v12 := GenS1{F0: 67, F1: -36}
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    var v4 int8 = -1
    var v5 int8 = (v4 + (v4 / ((-1) | 1)))
    v6 := make(map[int]bool)
    var v13 uint16 = 2
    if true {
        if false {
            switch 9223372036854775807 {
            case 32767:
                v14 := []int{687}
                _ = v14
            case -9223372036854775808:
                fmt.Println(v13, v13)
            default:
                switch ("hello" + (("\n\t" + "foo bar") + "a")) {
                case "foo bar":
                    var v15 uint = 0
                    _ = v15
                default:
                    fmt.Println(v5, v5, v4)
                }
            }
            var v16 int32 = 2147483647
            _ = v16
        } else {
            switch ("" + "a") {
            case "foo bar":
                v17 := GenS1{F0: 0, F1: v5}
                _ = v17
            case "\n\t":
                var v18 uint64 = genFn1(42, 440, (v13 ^ v13))
                _ = v18
            case "hello":
                var v19 uint16 = 375
                _ = v19
            }
            var v20 uint32 = 0
            _ = v20
        }
    } else {
        v21 := GenS1{F0: 238, F1: -128}
        _ = v21
    }
    var v22 uint64 = 0
    for fi23 := 0; fi23 < 8; fi23++ {
        v24 := GenS1{F0: 42, F1: (v4 | v4)}
        v5 = v5
        for fi25 := 0; fi25 < 11; fi25++ {
            var v26 bool = true
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    v27 := GenS1{F0: 42, F1: -95}
    v4 = v5
    if true {
        switch "" {
        case "hello":
            switch ("foo bar" + "") {
            case "a":
                v5 -= v4
            default:
                if true {
                    var v28 int16 = 0
                    var v29 float64 = -2.25
                    _ = v28
                    _ = v29
                }
            }
        case "foo bar":
            v5 -= v5
        case "a":
            var v30 bool = ((-600 > 42) && false)
            _ = v30
        }
        var v31 uint8 = 2
        _ = v31
    }
    if ((0.1 >= ((-2.25 / (0.1 / 0.1)) - (((1e10 - float64((0.1 - 1e10))) / (float64((0.1 / (float32(0.1) + 3.14159))) / float64(((1e10 * 1.5) / 1.0)))) * 0.0))) && false) {
        v4 -= -1
        v32 := []int{4294967295, 32768}
        _ = v32
    } else {
        v4 = (-(v5))
    }
    var v33 uint16 = (v13 / (((v13 - v13)) | 1))
    var v34 uint64 = genFn1(42, v22, 2)
    v35 := v6[128]
    v36 := GenS1{F0: 2147483647, F1: (v5 ^ (v4 * -1))}
    if (v35 && false) {
        var v37 int64 = 9223372036854775807
        var v38 int32 = -655
        v35 = true
        _ = v37
        _ = v38
    } else {
        v39 := []int{9223372036854775807, 2147483647}
        var v40 float64 = 3.14159
        _ = v39
        _ = v40
    }
    v41 := GenS1{F0: -1, F1: v4}
    for fi42 := 0; fi42 < 10; fi42++ {
        var v43 float32 = 1e10
        _ = fi42
        _ = v43
    }
    if (1 <= -1) {
        switch -9223372036854775808 {
        case 128:
            v34 = v34
        }
        v5 -= (v5 | 100)
    } else {
        v35 = false
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        var v45 bool = v35
        v46 := v6[(fi44 ^ (fi44 & fi44))]
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v47
    _ = v5
    _ = v6
}
