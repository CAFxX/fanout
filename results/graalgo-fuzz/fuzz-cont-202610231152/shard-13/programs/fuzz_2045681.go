package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint64) uint16 {
    if true {
        var v4 uint64 = p1
        p1 += p1
        for fi5 := 0; fi5 < 9; fi5++ {
            v6 := []int{55, (fi5 / ((4294967295) | 1)), fi5, 127}
            _ = fi5
            _ = v6
        }
        _ = v4
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 32767
}
func genFn1(p0 uint, p1 uint64) uint16 {
    p0 -= p0
    fmt.Println(p1, p1)
    v7 := GenS1{F0: (42 < 65535)}
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 float64 = (-2.25 / ((3.14159 + (0.1 / (3.14159 * float64(float32(0.1))))) * 0.1))
    var v2 float64 = (float64((float32(1e10) / 1.0)) + (((v1 * 1e10) * (0.1 - 0.1)) + v1))
    v3 := []uint32{1, 1, 2}
    v1 += (float64(-2.25) - (((float64(0.1) + (v1 + float64((0.1 * 1.5)))) * 3.14159) / -2.25))
    for fi8 := 0; fi8 < 16; fi8++ {
        if (((fi8 >> ((42) & 63)) != fi8) && true) {
            if true {
                fmt.Println(fi8, v2, v1)
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 int64 = -436
                    v1 = v2
                    _ = fi9
                    _ = v10
                }
            }
            v1 -= float64(3.14159)
            v11 := GenS2{F0: -32768}
            _ = v11
        } else {
            var v12 uint32 = 628
            _ = v12
        }
        v13 := GenS3{F0: 42, F1: ("" + (("a" + "") + "hello"))}
        _ = fi8
        _ = v13
    }
    v14 := GenS3{F0: 958, F1: "foo bar"}
    var v15 int = 1
    if ((("a" + "a") + "a") < "hello") {
        v1 = 1.5
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi16 := 0; fi16 < 19; fi16++ {
            fmt.Println(v1, v15, v2)
            v17 := GenS4{F0: -1}
            switch ("" + "\n\t") {
            case "a":
                switch "" {
                case "foo bar":
                    switch "foo bar" {
                    case "":
                        fmt.Println(v2)
                    }
                case "hello":
                    fmt.Println(v1, v1)
                default:
                    if (28 > 0) {
                        v18 := GenS5{F0: 1}
                        switch (("foo bar" + "hello") + (("hello" + "a") + (("" + "\n\t") + "a"))) {
                        case "hello":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "a":
                            v19 := GenS1{F0: true}
                            _ = v19
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = v18
                    }
                }
            case "foo bar":
                fmt.Println(v15, v15)
            }
            _ = fi16
            _ = v17
        }
    } else {
        v1 -= v2
        var v20 int32 = 42
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        if (true || (1 > 2)) {
            v22 := GenS4{F0: -37}
            v2 = (-2.25 - v1)
            _ = v22
        } else {
            v2 = 0.0
            fmt.Println(v1)
        }
        var v23 int8 = 127
        v23 -= 16
        _ = fi21
        _ = v23
    }
    v15 -= 157
    v24 := GenS6{F0: 1}
    v25 := GenS1{F0: ((127 < -1) || false)}
    var v26 int16 = -252
    var v27 uint16 = 1
    for fi28 := 0; fi28 < 6; fi28++ {
        fmt.Println(v15, v26)
        _ = fi28
    }
    v3 = append(v3, 804)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
