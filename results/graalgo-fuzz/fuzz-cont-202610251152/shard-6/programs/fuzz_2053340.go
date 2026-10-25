package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint64; F1 string; F2 string }

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (!(!false))
    var v3 int64 = -1
    var v4 uint8 = 0
    v5 := make(map[string]bool)
    var v6 uint32 = 2
    v2 = (4294967295 <= 2)
    var v7 int64 = (v3 >> ((4294967295) & 63))
    if (1 == 0) {
        fmt.Println(v4, v6, v7)
        if ((v6 / ((v6) | 1)) == (v6 * (v6 | (v6 * (^(v6)))))) {
            v6 = v6
            fmt.Println(v6, v3, v7)
        } else {
            for fi8 := 0; fi8 < 19; fi8++ {
                for fi9 := 0; fi9 < 13; fi9++ {
                    v5[""] = ((float64((((0.1 / 1.5) + 1.5) * 1e10)) - -2.25) > 3.14159)
                    _ = fi9
                }
                _ = fi8
            }
            var v10 uint64 = 42
            v5[("\n\t" + "\n\t")] = (v10 == v10)
            _ = v10
        }
    }
    v11 := GenS1{F0: (v4 ^ 1), F1: 42}
    switch -973 {
    case 0:
        switch -1 {
        case 2147483647:
            var v12 uint32 = 1
            _ = v12
        case 0:
            v13 := []int{}
            _ = v13
        case 499:
            v14 := GenS1{F0: (^(v4)), F1: 65535}
            _ = v14
        }
    case -793:
        var v15 float64 = 3.14159
        _ = v15
    case -2147483648:
        v16 := []int{4294967295, -217, 2147483647}
        _ = v16
    default:
        var v17 int16 = 32767
        _ = v17
    }
    var v18 int = 9223372036854775807
    var v19 int = 127
    for fi20 := 0; fi20 < 11; fi20++ {
        if (false && (v6 <= 0)) {
            delete(v1, "a")
            for fi21 := 0; fi21 < 5; fi21++ {
                if (-2.25 == 0.1) {
                    fmt.Println(v6, v6, fi21)
                    v2 = v2
                    var v22 int16 = int16(v19)
                    _ = v22
                } else {
                    fmt.Println(v18)
                    var v23 uint64 = 216
                    var v24 int16 = int16(v6)
                    _ = v23
                    _ = v24
                }
                _ = fi21
            }
        }
        _ = fi20
    }
    var v25 float32 = 3.14159
    switch v18 {
    case 0:
        var v26 uint32 = (v6 - v6)
        _ = v26
    case 42:
        var v27 int64 = v3
        _ = v27
    case 4294967295:
        if ((v4 + (v4 << ((65535) & 63))) >= v4) {
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 bool = (v2 || (v4 > (v4 << ((0) & 63))))
                var v30 uint32 = (v6 * v6)
                _ = fi28
                _ = v29
                _ = v30
            }
            if ((v6 <= (^(v6))) && false) {
                v4 += v4
            } else {
                v19 = (v18 & (v19 + v19))
                fmt.Println(v6, v7, v18)
            }
        } else {
            if (18446744073709551615 <= 1) {
                v3 = int64(v6)
                for fi31 := 0; fi31 < 8; fi31++ {
                    v32 := v5[((("\n\t" + "a") + "\n\t") + "foo bar")]
                    _ = fi31
                    _ = v32
                }
                if (!(!(32767 < 2147483647))) {
                    v4 = v4
                    if (!true) {
                        var v33 uint8 = v4
                        v34 := GenS2{F0: 146, F1: v6}
                        switch "\n\t" {
                        case "hello":
                            for fi35 := 0; fi35 < 5; fi35++ {
                                v36 := []int{v18, v18}
                                fmt.Println(fi35, v18, v7)
                                _ = fi35
                                _ = v36
                            }
                        default:
                            v18 = int(v4)
                        }
                        _ = v33
                        _ = v34
                    }
                } else {
                    var v37 uint32 = v6
                    fmt.Println(v7, v18)
                    _ = v37
                }
            }
        }
    }
    var v38 uint8 = 1
    fmt.Println(v38, v19)
    fmt.Println(v18, v18)
    v3 -= v3
    v5["\n\t"] = (2147483647 >= 656)
    fmt.Println(v25, v38)
    if ((v38 | (v4 ^ v38)) > v4) {
        var v39 float32 = float32(1.5)
        var v40 int = v18
        var v41 uint16 = 1
        _ = v39
        _ = v40
        _ = v41
    } else {
        var v42 uint16 = 255
        var v43 uint32 = v6
        switch "" {
        case "foo bar":
            fmt.Println(v18, v43)
        }
        _ = v42
        _ = v43
    }
    if v2 {
        switch "\n\t" {
        case "\n\t":
            v6 = 2147483647
        case "hello":
            fmt.Println(v4, v7, v4)
        case "foo bar":
            fmt.Println(v3, v2)
        default:
            fmt.Println(v25, v18)
        }
        v44 := GenS3{F0: 407, F1: "foo bar", F2: "foo bar"}
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
