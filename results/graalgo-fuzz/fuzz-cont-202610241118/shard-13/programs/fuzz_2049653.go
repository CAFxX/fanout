package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 string; F1 int }
func genFn0(p0 bool) int16 {
    var v4 int = -2147483648
    for fi5 := 0; fi5 < 7; fi5++ {
        p0 = p0
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := GenS1{F0: 216, F1: "hello", F2: 0}
            switch "foo bar" {
            case "a":
                switch "hello" {
                case "a":
                    var v8 int64 = -283
                    _ = v8
                case "":
                    p0 = (float32(3.14159) < float32(3.14159))
                case "foo bar":
                    if (42 > 1) {
                        if p0 {
                            v9 := GenS1{F0: 807, F1: "hello", F2: 127}
                            switch (-(fi5)) {
                            case 4294967295:
                                v4 -= fi6
                                switch "" {
                                case "hello":
                                    var v10 uint64 = 4294967295
                                    _ = v10
                                case "foo bar":
                                    p0 = (65535 == 1)
                                default:
                                    if p0 {
                                        var v11 uint8 = 177
                                        _ = v11
                                    }
                                }
                                p0 = p0
                            case -92:
                                v4 -= v4
                            case 42:
                                var v12 uint64 = 1
                                _ = v12
                            }
                            var v13 uint8 = 127
                            _ = v13
                            _ = v9
                        }
                    }
                default:
                    var v14 uint = 420
                    _ = v14
                }
            case "hello":
                fmt.Println(p0, v4)
                var v15 bool = p0
                if (1e10 <= ((float64(0.0) - (float64(3.14159) * 1e10)) + 0.1)) {
                    switch (fi5 ^ 127) {
                    case -27:
                        v16 := GenS1{F0: 0, F1: "hello", F2: 0}
                        _ = v16
                    case 480:
                        if ((!false) && (fi6 != (fi6 ^ (v4 / ((42) | 1))))) {
                            v17 := []int{}
                            v18 := GenS1{F0: 655, F1: "", F2: 255}
                            _ = v17
                            _ = v18
                        }
                    }
                }
                var v19 uint8 = 128
                _ = v15
                _ = v19
            }
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        p0 = (p0 && (!((((("\n\t" + (("hello" + "foo bar") + "foo bar")) + "hello") + "") + "foo bar") != "foo bar")))
        _ = fi20
    }
    v4 += v4
    v21 := GenS1{F0: 65535, F1: "foo bar", F2: 0}
    _ = p0
    _ = v21
    _ = v4
    return 783
}
func genFn1(p0 uint64) int64 {
    var v22 int = 0
    v23 := []int{}
    var v24 uint8 = 127
    p0 = (p0 + 1)
    if (-2.25 < ((float64(1.5) - 0.1) / 1e10)) {
        v24 -= 2
        v25 := GenS2{F0: 42, F1: 512}
        _ = v25
    }
    _ = p0
    _ = v22
    _ = v23
    _ = v24
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{2147483648, 426, 65535}
    var v3 bool = (!(!(1 == 414)))
    if false {
        delete(v1, ("" + ("" + "foo bar")))
        var v26 int16 = -32768
        _ = v26
    }
    if true {
        v3 = v3
    }
    v1["a"] = v3
    v27 := GenS3{F0: "foo bar", F1: 2147483647}
    var v28 bool = (0.0 < 1.5)
    var v29 float32 = 1.5
    var v30 float64 = float64((v29 * v29))
    v31 := GenS2{F0: 1000, F1: 1000}
    switch "foo bar" {
    case "hello":
        var v32 uint16 = 255
        _ = v32
    }
    v29 -= 1e10
    for fi33 := 0; fi33 < 8; fi33++ {
        v28 = ((90 < 128) && (291 < -795))
        var v34 uint64 = 794
        _ = fi33
        _ = v34
    }
    var v35 float32 = v29
    for fi36 := 0; fi36 < 7; fi36++ {
        var v37 uint16 = 255
        fmt.Println(fi36, v29)
        switch (fi36 ^ fi36) {
        case 42:
            if true {
                if v3 {
                    v29 = (1e10 / v29)
                    var v38 uint64 = 4294967295
                    _ = v38
                }
                if ((1 <= 1) && true) {
                    for fi39 := 0; fi39 < 2; fi39++ {
                        var v40 int32 = -2147483648
                        var v41 uint = 1
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    v3 = v3
                    v28 = false
                }
                v35 = v35
            } else {
                v28 = true
            }
        case 9223372036854775807:
            for fi42 := 0; fi42 < 17; fi42++ {
                fmt.Println(v28, v35)
                _ = fi42
            }
        default:
            for fi43 := 0; fi43 < 11; fi43++ {
                var v44 uint16 = v37
                v2 = append(v2, 2147483647)
                _ = fi43
                _ = v44
            }
        }
        _ = fi36
        _ = v37
    }
    if (!false) {
        v45 := GenS3{F0: "foo bar", F1: 1}
        _ = v45
    } else {
        fmt.Println(v28, v28, v29)
        var v46 int8 = 42
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
}
