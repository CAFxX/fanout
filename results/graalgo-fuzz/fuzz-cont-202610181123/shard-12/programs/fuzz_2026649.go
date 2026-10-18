package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int8; F1 int; F2 uint }
type GenS4 struct { F0 uint64; F1 uint; F2 int32 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint16, p1 uint8, p2 int32) string {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int8 = -128
        _ = fi7
        _ = v8
    }
    p0 += 2
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}

func main() {
    v1 := []string{("foo bar" + "a"), "\n\t", "a"}
    var v2 string = "foo bar"
    var v3 uint16 = 969
    v4 := make(map[int]int)
    v5 := make(map[int]uint32)
    v6 := []string{(v2 + ("foo bar" + ("hello" + v2))), "a", (v2 + (v2 + (("\n\t" + ((v2 + ("a" + v2)) + (v2 + "foo bar"))) + ("hello" + v2))))}
    if (!true) {
        _ = len(v1)
        v3 = v3
        if ((float32((((0.0 / -2.25) / ((((((((0.0 * (-2.25 / 3.14159)) - (1.5 * 1.5)) - 0.0) + ((1.5 / 1.0) + (1.5 + -2.25))) + (3.14159 - 1e10)) + (float64(-2.25) - 0.1)) + (float64(1e10) * (1.5 / (1.5 * ((0.1 / ((float64((0.1 * float32((0.0 / 1.0)))) - 1e10) / 1.0)) / 1.0))))) / 1.0)) + float64((1.5 + (-2.25 - (((1e10 - 0.1) / (1e10 / (0.0 + (3.14159 / (float32(-2.25) / 1.0))))) - 0.0)))))) * (0.0 * (3.14159 + -2.25))) <= (3.14159 - 3.14159)) {
            v3 -= (v3 - 1)
        }
    } else {
        var v9 uint64 = 324
        delete(v5, -600)
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        switch -9223372036854775808 {
        case 127:
            var v11 float32 = 0.0
            _ = v11
        case 880:
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 int = (fi10 >> ((28) & 63))
                _ = fi12
                _ = v13
            }
        case 32767:
            v14 := GenS1{F0: 637}
            _ = v14
        default:
            v3 = 0
        }
        _ = len(v1)
        _ = fi10
    }
    switch 127 {
    case -9223372036854775808:
        v2 = (v2 + v2)
    }
    var v15 int64 = -9223372036854775808
    if (!(false || false)) {
        v16 := GenS2{F0: 255, F1: 679}
        v4[int(v15)] = 1
        _ = v16
    } else {
        v15 += (v15 - (v15 & (v15 << ((447) & 63))))
        var v17 uint16 = 2
        var v18 float64 = 0.0
        _ = v17
        _ = v18
    }
    switch v2 {
    case "hello":
        v19 := GenS1{F0: -100}
        _ = v19
    default:
        v4[32767] = -9223372036854775808
    }
    if false {
        var v20 int64 = (v15 * (v15 - (v15 ^ (v15 >> ((2) & 63)))))
        v21 := GenS1{F0: 128}
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 string = v2
                fmt.Println(v2)
                if false {
                    if (v3 <= (v3 & 2)) {
                        switch (^(fi23)) {
                        case 4294967295:
                            fmt.Println(v15, v24, v2)
                        case 2147483647:
                            fmt.Println(v24)
                        case 1:
                            var v25 bool = (!true)
                            _ = v25
                        }
                        v6 = append(v6, ((v24 + "") + v2))
                        v15 = (v15 % ((v15) | 1))
                    }
                } else {
                    v26 := GenS3{F0: 0, F1: -2147483648, F2: 255}
                    _ = v26
                }
                _ = fi23
                _ = v24
            }
            var v27 float64 = (1.5 / -2.25)
            delete(v4, fi22)
            _ = fi22
            _ = v27
        }
        _ = v20
        _ = v21
    } else {
        var v28 float64 = 1e10
        _ = v28
    }
    fmt.Println(v3, v3, v2)
    v29 := GenS4{F0: 2, F1: 255, F2: -1}
    v30 := GenS5{F0: 85}
    for fi31 := 0; fi31 < 3; fi31++ {
        v3 = v3
        var v32 uint64 = 42
        if true {
            switch ((v2 + (v2 + v2)) + "foo bar") {
            case "hello":
                _ = len(v5)
            case "\n\t":
                v5[fi31] = 2147483648
            case "a":
                var v33 int8 = 42
                _ = v33
            default:
                v34 := GenS2{F0: v3, F1: 65535}
                _ = v34
            }
        } else {
            var v35 float32 = (0.1 * (1.5 - ((((1.5 - -2.25) + 0.1) * 0.0) * (0.1 / float32(0.1)))))
            delete(v4, fi31)
            var v36 uint32 = 2147483648
            _ = v35
            _ = v36
        }
        _ = fi31
        _ = v32
    }
    v3 -= v3
    v2 = (v2 + "")
    fmt.Println(v2)
    var v37 int16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
