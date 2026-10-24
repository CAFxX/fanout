package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint64; F1 uint; F2 int8 }
type GenS3 struct { F0 string; F1 uint32; F2 uint32 }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int32 = 42
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(v8)
            var v10 uint32 = 461
            _ = fi9
            _ = v10
        }
        fmt.Println(v8, fi7)
        _ = fi7
        _ = v8
    }
    var v11 int = -764
    _ = v11
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    var v3 int = -2147483648
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 42
    v6 := make(map[uint32]bool)
    v12 := GenS1{F0: 0, F1: 576}
    v5 = v5
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            v5 += 2147483647
            switch ("\n\t" + ("\n\t" + "hello")) {
            case "hello":
                switch (fi13 >> ((2) & 63)) {
                case 42:
                    v15 := GenS1{F0: 9223372036854775807, F1: -9223372036854775808}
                    _ = v15
                default:
                    v16 := GenS1{F0: -1, F1: 1}
                    _ = v16
                }
            }
            if (988 < 32767) {
                var v17 uint = 955
                _ = v17
            }
            _ = fi14
        }
        if true {
            var v18 int8 = 42
            v5 -= v5
            var v19 int = (^(fi13))
            _ = v18
            _ = v19
        } else {
            fmt.Println(v5, fi13)
        }
        _ = fi13
    }
    v5 = (v5 / (((v5 + v5)) | 1))
    for fi20 := 0; fi20 < 10; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            if (!false) {
                if (true && true) {
                    var v22 uint32 = (v5 ^ (v5 >> ((42) & 63)))
                    v23 := GenS1{F0: 42, F1: fi20}
                    var v24 int64 = -215
                    _ = v22
                    _ = v23
                    _ = v24
                }
                var v25 uint8 = 216
                if (("" >= "foo bar") || (fi21 <= v3)) {
                    v25 += v25
                    _ = len(v2)
                }
                _ = v25
            }
            fmt.Println(fi20)
            fmt.Println(v5)
            _ = fi21
        }
        for fi26 := 0; fi26 < 18; fi26++ {
            switch "foo bar" {
            case "foo bar":
                var v27 int16 = 42
                _ = v27
            case "\n\t":
                v5 = (v5 * 231)
            }
            var v28 uint16 = 924
            _ = fi26
            _ = v28
        }
        _ = fi20
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 float32 = (-2.25 / 3.14159)
        switch 4294967295 {
        case 95:
            var v31 int = (fi29 | fi29)
            _ = v31
        case 0:
            var v32 float32 = 1e10
            _ = v32
        default:
            v30 = v30
        }
        _ = fi29
        _ = v30
    }
    fmt.Println(v3)
    for fi33 := 0; fi33 < 16; fi33++ {
        if (2 <= 197) {
            switch "hello" {
            case "a":
                for fi34 := 0; fi34 < 8; fi34++ {
                    v35 := GenS2{F0: 1, F1: 365, F2: 127}
                    if false {
                        v3 -= 42
                        v5 -= (v5 & genFn0())
                        v36 := []int{fi34, fi34, v3}
                        var v37 int16 = -32768
                        if false {
                            var v38 uint8 = 217
                            var v39 float64 = ((-2.25 * 0.0) / 0.1)
                            var v40 float32 = 0.0
                            if len(v36) > 0 {
                                v36[0] = 2147483647
                            }
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        _ = v36
                        _ = v37
                    }
                    _ = fi34
                    _ = v35
                }
            case "":
                v41 := []int{(v3 ^ v3), (v3 << ((18446744073709551615) & 63)), (v3 + v3)}
                _ = v41
            case "foo bar":
                switch "a" {
                case "foo bar":
                    var v42 bool = (((!((v5 << ((779) & 63)) >= v5)) && (2 < 0)) || (-1 <= 9223372036854775807))
                    _ = v42
                case "a":
                    v5 = 2147483647
                case "\n\t":
                    v43 := GenS3{F0: "", F1: (^(v5)), F2: genFn0()}
                    _ = v43
                }
            }
            v44 := []int{128, (fi33 >> ((1) & 63))}
            v5 -= (v5 >> ((42) & 63))
            _ = v44
        } else {
            var v45 int64 = 0
            _ = v45
        }
        _ = fi33
    }
    fmt.Println(v3, v3)
    v5 = genFn0()
    v3 += v3
    var v46 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
