package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint64 }
func genFn0(p0 uint64) int8 {
    for fi4 := 0; fi4 < 6; fi4++ {
        p0 -= 52
        v5 := GenS1{F0: 0, F1: -2147483648, F2: 458}
        var v6 uint = 71
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        switch (fi7 << ((2) & 63)) {
        case 127:
            p0 += 42
        default:
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := GenS1{F0: 255, F1: -895, F2: 858}
                var v10 uint16 = uint16(p0)
                v11 := []int{(fi8 << ((65535) & 63))}
                p0 += p0
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        }
        var v12 uint32 = 42
        v12 += 2147483647
        _ = fi7
        _ = v12
    }
    v13 := []int{}
    for fi14 := 0; fi14 < 18; fi14++ {
        fmt.Println(fi14, fi14)
        v15 := GenS1{F0: 255, F1: 2147483647, F2: 27}
        switch ("foo bar" + "a") {
        case "a":
            for fi16 := 0; fi16 < 16; fi16++ {
                fmt.Println(p0, fi14)
                p0 += 42
                var v17 uint16 = 2
                _ = fi16
                _ = v17
            }
        case "\n\t":
            var v18 float32 = ((1e10 * (((float32(3.14159) * 1.5) / float32((-2.25 - 0.0))) / 0.1)) - (3.14159 / 1.5))
            _ = v18
        case "hello":
            v19 := GenS1{F0: 2, F1: -1, F2: 1}
            _ = v19
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(p0)
    _ = p0
    _ = v13
    return 42
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    switch 42 {
    case 32767:
        v1 = -2.25
    default:
        var v20 string = "a"
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        fmt.Println(v1, v1)
        fmt.Println(fi21)
        var v22 int16 = 0
        _ = fi21
        _ = v22
    }
    if ((-32768 == 42) && (0 > 619)) {
        v1 -= (1.5 - v1)
        v23 := GenS1{F0: 295, F1: 42, F2: 255}
        if true {
            v1 -= ((((v1 / (float64((((1e10 / 1.0) / 1.0) / (1.5 * 0.1))) + v1)) * float64(1.5)) - 1e10) / 1.0)
        }
        _ = v23
    } else {
        switch "\n\t" {
        case "a":
            v1 -= float64(1.5)
        default:
            v24 := GenS1{F0: 126, F1: -1, F2: 253}
            _ = v24
        }
        var v25 uint = 42
        _ = v25
    }
    switch "hello" {
    case "hello":
        if false {
            if (!((0.0 + float32(float64((-2.25 - float32(((v1 - 0.0) * v1)))))) == 3.14159)) {
                for fi26 := 0; fi26 < 5; fi26++ {
                    var v27 int32 = 384
                    _ = fi26
                    _ = v27
                }
            } else {
                if (((((0.1 / v1) / (v1 + (v1 + (float64(((float32((-2.25 * float64(0.0))) / 1.0) - 0.1)) + float64(3.14159))))) / v1) / 0.1) >= float64((3.14159 * 1e10))) {
                    fmt.Println(v1)
                    delete(v3, "")
                } else {
                    v1 -= v1
                    var v28 uint8 = 124
                    v3[("hello" + ((((("hello" + "hello") + "hello") + ("" + "a")) + "a") + "foo bar"))] = 65535
                    _ = v28
                }
                var v29 int = 127
                var v30 bool = false
                _ = v29
                _ = v30
            }
        }
    default:
        switch -353 {
        case 0:
            var v31 uint = 293
            _ = v31
        default:
            v32 := []int{}
            _ = v32
        }
    }
    var v33 bool = false
    for fi34 := 0; fi34 < 19; fi34++ {
        v35 := v3["\n\t"]
        if false {
            fmt.Println(v33, v35)
            for fi36 := 0; fi36 < 7; fi36++ {
                delete(v2, "hello")
                v37 := GenS1{F0: 65535, F1: -1, F2: 139}
                v1 += v1
                _ = fi36
                _ = v37
            }
        }
        _ = fi34
        _ = v35
    }
    var v38 int8 = 37
    v39 := []int{9223372036854775807, 32768, 42}
    v40 := GenS2{F0: (v38 * 100)}
    v41 := GenS3{F0: 994, F1: (v38 - -1), F2: 42}
    v38 -= -1
    v38 = genFn0(2)
    delete(v2, ("\n\t" + ""))
    if (((32767 >= 255) && v33) || (!true)) {
        fmt.Println(v38, v33, v33)
        var v42 int64 = 9223372036854775807
        delete(v2, "hello")
        _ = v42
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v39
    _ = v40
    _ = v41
}
