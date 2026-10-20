package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int, p1 int8, p2 uint16) int64 {
    var v5 int32 = -1
    var v6 uint32 = 904
    if (!(42 < 128)) {
        if (!(3.14159 < 0.0)) {
            if true {
                fmt.Println(v6)
                if true {
                    var v7 int64 = -1
                    for fi8 := 0; fi8 < 5; fi8++ {
                        if false {
                            var v9 float64 = 1e10
                            _ = v9
                        } else {
                            v10 := []int{-776, (fi8 + fi8)}
                            _ = v10
                        }
                        for fi11 := 0; fi11 < 1; fi11++ {
                            v6 += v6
                            fmt.Println(fi8)
                            _ = fi11
                        }
                        _ = fi8
                    }
                    switch (p0 >> ((489) & 63)) {
                    case 198:
                        p2 = p2
                    case 32768:
                        v12 := []int{p0, 2147483647}
                        _ = v12
                    }
                    _ = v7
                } else {
                    switch p0 {
                    case -2147483648:
                        v13 := GenS1{F0: p2, F1: v6, F2: (-(v5))}
                        _ = v13
                    case 128:
                        fmt.Println(v6)
                    }
                }
                var v14 uint16 = 506
                _ = v14
            }
            switch "" {
            case "\n\t":
                v15 := GenS1{F0: p2, F1: v6, F2: (v5 ^ v5)}
                _ = v15
            case "foo bar":
                v16 := GenS2{F0: 518, F1: "foo bar"}
                _ = v16
            case "a":
                fmt.Println(p2, p1, p1)
            default:
                v6 += 940
            }
        }
    }
    var v17 int8 = p1
    switch "foo bar" {
    case "foo bar":
        var v18 uint8 = 151
        _ = v18
    case "hello":
        v19 := GenS2{F0: 65535, F1: "a"}
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v5
    _ = v6
    return -9223372036854775808
}
func genFn1(p0 int16) uint8 {
    p0 += p0
    v20 := []int{32767, 128, 913, 9223372036854775807}
    var v21 int16 = 0
    if (-2.25 != 0.0) {
        var v22 string = "hello"
        v21 -= v21
        var v23 int = 128
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        v25 := GenS3{F0: 42, F1: true}
        var v26 float32 = float32(((0.1 * ((0.1 + (1.5 / (3.14159 * (1.5 + 3.14159)))) + 0.0)) - -2.25))
        _ = fi24
        _ = v25
        _ = v26
    }
    _ = p0
    _ = v20
    _ = v21
    return 42
}

func main() {
    v1 := []string{(("\n\t" + "foo bar") + "foo bar")}
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 42
    var v4 uint8 = v3
    v3 = (v4 ^ 42)
    delete(v2, 2147483647)
    v27 := v2[449]
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    v28 := GenS4{F0: 2147483647}
    v27 = 2147483648
    v29 := GenS3{F0: 18446744073709551615, F1: (32768 < 32768)}
    switch ("a" + "hello") {
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "\n\t":
        v30 := GenS5{F0: 4294967295}
        _ = v30
    case "foo bar":
        v1 = append(v1, ("" + ""))
    default:
        var v31 float64 = 1e10
        _ = v31
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        for fi33 := 0; fi33 < 16; fi33++ {
            var v34 int8 = 121
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            _ = len(v2)
            _ = fi33
            _ = v34
        }
        var v35 uint64 = 2
        for fi36 := 0; fi36 < 11; fi36++ {
            var v37 int = 128
            v37 += fi36
            _ = fi36
            _ = v37
        }
        _ = fi32
        _ = v35
    }
    var v38 float64 = 0.1
    v3 -= (^(v3))
    fmt.Println(v38, v27, v3)
    switch "" {
    case "":
        for fi39 := 0; fi39 < 8; fi39++ {
            var v40 float32 = 0.1
            _ = fi39
            _ = v40
        }
    case "hello":
        for fi41 := 0; fi41 < 3; fi41++ {
            for fi42 := 0; fi42 < 15; fi42++ {
                var v43 uint32 = (v27 >> ((uint(v3)) & 63))
                v3 += genFn1(1)
                _ = fi42
                _ = v43
            }
            switch ("\n\t" + (("\n\t" + "") + ("\n\t" + "foo bar"))) {
            case "hello":
                var v44 float64 = (v38 * v38)
                _ = v44
            case "":
                var v45 int64 = -207
                _ = v45
            default:
                switch fi41 {
                case 0:
                    for fi46 := 0; fi46 < 1; fi46++ {
                        if (!true) {
                            var v47 int32 = -1
                            var v48 int16 = -1
                            _ = v47
                            _ = v48
                        }
                        var v49 int32 = 677
                        _ = fi46
                        _ = v49
                    }
                case 1:
                    var v50 uint8 = genFn1(1000)
                    _ = v50
                case 9223372036854775807:
                    var v51 float32 = -2.25
                    _ = v51
                default:
                    switch (fi41 & 1) {
                    case -501:
                        var v52 uint = 736
                        _ = v52
                    case -1:
                        if false {
                            v53 := GenS6{F0: 1}
                            _ = v53
                        }
                    case 9223372036854775807:
                        var v54 uint64 = 4294967295
                        _ = v54
                    default:
                        _ = len(v2)
                    }
                }
            }
            _ = fi41
        }
    case "foo bar":
        v3 -= genFn1(-32768)
    }
    var v55 int16 = 414
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v4
    _ = v55
}
