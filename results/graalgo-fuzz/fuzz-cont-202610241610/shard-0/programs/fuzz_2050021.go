package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 string; F2 string }
func genFn0(p0 bool, p1 int64, p2 uint64) uint16 {
    if false {
        if (!((!(1 == 255)) || (809 >= 65535))) {
            v6 := GenS1{F0: 1}
            fmt.Println(p0)
            _ = v6
        }
        fmt.Println(p0, p2)
    } else {
        if p0 {
            var v7 int = 32768
            for fi8 := 0; fi8 < 17; fi8++ {
                if (false || false) {
                    p0 = (((p0 || false) || p0) && (true && (-702 >= int32(fi8))))
                    v9 := GenS2{F0: 18446744073709551615, F1: 95}
                    v10 := GenS3{F0: fi8, F1: 196}
                    _ = v10
                    _ = v9
                }
                var v11 int64 = p1
                var v12 int16 = -1
                _ = fi8
                _ = v11
                _ = v12
            }
            if false {
                var v13 uint = 2
                var v14 uint8 = 0
                _ = v13
                _ = v14
            }
            _ = v7
        }
        var v15 int = 1
        _ = v15
    }
    v16 := GenS3{F0: 4294967295, F1: 127}
    if (-2.25 < (1.5 + 0.0)) {
        switch "hello" {
        case "hello":
            for fi17 := 0; fi17 < 12; fi17++ {
                p0 = ((p2 | (p2 * p2)) == p2)
                _ = fi17
            }
        default:
            if (!(0.0 < 1.5)) {
                var v18 uint32 = 2
                if p0 {
                    v19 := GenS1{F0: 128}
                    if true {
                        switch "hello" {
                        case "foo bar":
                            if (42 != 42) {
                                v20 := []int{32768, 376, -672, 42}
                                var v21 int32 = 576
                                _ = v20
                                _ = v21
                            } else {
                                v22 := GenS1{F0: 32768}
                                _ = v22
                            }
                        case "hello":
                            for fi23 := 0; fi23 < 14; fi23++ {
                                p1 += p1
                                var v24 uint8 = 53
                                _ = fi23
                                _ = v24
                            }
                        }
                        p1 = p1
                    } else {
                        v25 := GenS2{F0: 4294967295, F1: 2}
                        var v26 string = ""
                        v27 := GenS3{F0: -1, F1: 255}
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = v19
                }
                var v28 float32 = (1.5 / 1.0)
                _ = v18
                _ = v28
            }
        }
        var v29 int = 4294967295
        v30 := GenS3{F0: v29, F1: 167}
        _ = v29
        _ = v30
    } else {
        v31 := GenS3{F0: 127, F1: 1}
        var v32 int = 42
        v33 := GenS1{F0: v32}
        _ = v31
        _ = v32
        _ = v33
    }
    var v34 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v34
    return 232
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    v4 := make(map[string]uint32)
    var v5 uint16 = 487
    var v35 int64 = -9223372036854775808
    for fi36 := 0; fi36 < 1; fi36++ {
        for fi37 := 0; fi37 < 15; fi37++ {
            v1 += v1
            _ = fi37
        }
        _ = fi36
    }
    v38 := GenS2{F0: 18446744073709551615, F1: 41}
    v39 := GenS4{F0: 2147483647}
    v40 := GenS5{F0: (v35 > v35), F1: (((("a" + "") + ("hello" + "a")) + "\n\t") + ("foo bar" + "foo bar")), F2: ("foo bar" + "a")}
    var v41 string = ""
    v1 = 227
    fmt.Println(v5)
    fmt.Println(v5)
    for fi42 := 0; fi42 < 17; fi42++ {
        v41 = "hello"
        var v43 float64 = (0.0 / ((1.5 * 3.14159) / 1.0))
        var v44 int = fi42
        _ = fi42
        _ = v43
        _ = v44
    }
    v45 := []int{2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v5
}
