package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint }

func main() {
    var v1 int16 = -563
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    if false {
        switch 42 {
        case 42:
            v4 := GenS1{F0: 2, F1: 65535}
            _ = v4
        case 2147483647:
            v5 := GenS1{F0: 64, F1: 486}
            _ = v5
        default:
            var v6 int8 = -128
            _ = v6
        }
        if false {
            fmt.Println(v1)
        } else {
            v1 += v1
            var v7 int64 = 474
            _ = v7
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := GenS1{F0: 2, F1: 409}
                var v11 int8 = 42
                _ = fi9
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 int32 = -310
                _ = fi12
                _ = v13
            }
            var v14 uint32 = 2147483648
            _ = fi8
            _ = v14
        }
    } else {
        var v15 uint = 1
        v16 := GenS2{F0: 269, F1: -399, F2: 1}
        if (255 <= 0) {
            v15 -= v15
            var v17 int = -2147483648
            v17 += v17
            _ = v17
        }
        _ = v15
        _ = v16
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 string = "\n\t"
        for fi20 := 0; fi20 < 15; fi20++ {
            v1 -= 1000
            var v21 float64 = (1.5 - ((float64(-2.25) * 1.5) * (-2.25 * 0.0)))
            var v22 uint8 = 2
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v1)
    for fi23 := 0; fi23 < 2; fi23++ {
        if (!true) {
            for fi24 := 0; fi24 < 12; fi24++ {
                switch "\n\t" {
                case "foo bar":
                    v1 = (v1 ^ 1000)
                }
                _ = fi24
            }
        }
        v25 := []int{(fi23 >> ((18446744073709551615) & 63)), (fi23 / (((fi23 - fi23)) | 1)), (fi23 / ((1) | 1))}
        _ = fi23
        _ = v25
    }
    var v26 int64 = 0
    var v27 uint16 = 1
    v1 += (v1 - (v1 >> ((0) & 63)))
    v28 := []int{0}
    var v29 bool = false
    for fi30 := 0; fi30 < 12; fi30++ {
        v29 = (-2147483648 <= -2147483648)
        if (true || (!false)) {
            v31 := GenS3{F0: 1}
            v3["\n\t"] = 151
            v32 := GenS1{F0: 65535, F1: 1}
            _ = v31
            _ = v32
        } else {
            for fi33 := 0; fi33 < 6; fi33++ {
                v1 += 0
                var v34 uint16 = v27
                switch (-(fi30)) {
                case 496:
                    v35 := GenS2{F0: 447, F1: -627, F2: 4294967295}
                    _ = v35
                case -2147483648:
                    v26 -= (-(v26))
                case -9223372036854775808:
                    v1 += v1
                default:
                    switch "foo bar" {
                    case "\n\t":
                        if (uint8(v27) != 255) {
                            v27 = v27
                            fmt.Println(fi30, v29)
                        }
                        var v36 int8 = 24
                        _ = v36
                    case "a":
                        var v37 int16 = -32768
                        _ = v37
                    }
                }
                _ = fi33
                _ = v34
            }
            v38 := GenS1{F0: 255, F1: (v27 + (v27 | v27))}
            if (((3.14159 - (0.1 + (0.0 - 3.14159))) * 3.14159) >= 1.5) {
                var v39 bool = (!(v29 || true))
                if (!v39) {
                    delete(v2, "a")
                    v40 := GenS4{F0: int32(v27), F1: 18446744073709551615}
                    v29 = v29
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 9; fi41++ {
                        var v42 uint32 = 1
                        v43 := GenS2{F0: 1, F1: 119, F2: 1}
                        v26 += 633
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    v26 = (v26 ^ v26)
                    var v44 uint16 = (^(v27))
                    _ = v44
                }
                switch "foo bar" {
                case "foo bar":
                    v45 := GenS3{F0: 2147483648}
                    _ = v45
                case "hello":
                    v46 := GenS2{F0: 412, F1: -2147483648, F2: 2}
                    _ = v46
                case "":
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                }
                _ = v39
            }
            _ = v38
        }
        _ = fi30
    }
    var v47 int64 = 1
    if false {
        switch 42 {
        case -176:
            v29 = v29
        case -431:
            var v48 bool = v29
            _ = v48
        }
    } else {
        v27 = (v27 >> ((42) & 63))
        fmt.Println(v26, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v47
}
