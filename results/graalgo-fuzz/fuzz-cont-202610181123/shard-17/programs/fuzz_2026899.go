package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int8 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int8; F1 int64; F2 int8 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 uint8; F1 uint; F2 int }

func main() {
    var v1 int64 = -1
    v2 := make(map[int]uint32)
    v3 := []string{"", ("\n\t" + "hello"), ((("a" + "\n\t") + "") + "hello")}
    var v4 string = "\n\t"
    v5 := make(map[int]bool)
    var v6 uint32 = 4294967295
    v2[0] = (^(v6))
    var v7 bool = false
    v7 = true
    var v8 int16 = 32767
    v9 := GenS1{F0: 255, F1: 1, F2: int8(v6)}
    if ((v4 + v4) < v4) {
        if len(v3) > 0 {
            v3[0] = ("\n\t" + v4)
        }
        v10 := GenS2{F0: 0, F1: (0 > 65535)}
        switch "hello" {
        case "a":
            if ((v8 / ((-458) | 1)) >= v8) {
                v6 -= (v6 >> ((448) & 63))
                fmt.Println(v6, v1)
                switch "" {
                case "a":
                    for fi11 := 0; fi11 < 14; fi11++ {
                        v4 = v4
                        v6 = 1
                        if v7 {
                            switch v4 {
                            case "a":
                                v12 := GenS3{F0: v6, F1: 0}
                                _ = v12
                            }
                            if v7 {
                                switch fi11 {
                                case 4294967295:
                                    var v13 uint = 2
                                    _ = v13
                                case -9223372036854775808:
                                    if (!(true && true)) {
                                        var v14 uint64 = uint64(v6)
                                        _ = v14
                                    } else {
                                        for fi15 := 0; fi15 < 4; fi15++ {
                                            if v7 {
                                                if true {
                                                    switch v4 {
                                                    case "\n\t":
                                                        var v16 float32 = 0.0
                                                        _ = v16
                                                    case "hello":
                                                        _ = len(v5)
                                                    default:
                                                        var v17 float32 = 0.1
                                                        _ = v17
                                                    }
                                                }
                                                v1 -= 0
                                                fmt.Println(fi11, fi11, v8)
                                            }
                                            if (255 > 127) {
                                                v6 -= (v6 + 65535)
                                                v8 -= (-(v8))
                                                v18 := GenS3{F0: (v6 / ((v6) | 1)), F1: 4294967295}
                                                _ = v18
                                            } else {
                                                switch 2147483647 {
                                                case 317:
                                                    v19 := GenS4{F0: -95, F1: (v1 + (v1 >> ((42) & 63))), F2: int8(v1)}
                                                    fmt.Println(v6, v1, v6)
                                                    v20 := GenS3{F0: (v6 & (v6 & (v6 & v6))), F1: 773}
                                                    _ = v19
                                                    _ = v20
                                                case 4294967295:
                                                    var v21 uint16 = 255
                                                    _ = v21
                                                case -2147483648:
                                                    var v22 uint8 = 128
                                                    _ = v22
                                                default:
                                                    v23 := GenS3{F0: 2147483647, F1: 18446744073709551615}
                                                    _ = v23
                                                }
                                                v24 := v5[-9223372036854775808]
                                                _ = v24
                                            }
                                            if (((true || ((4294967295 >= 2) || v7)) && (0 <= 2)) && (419 < 2147483647)) {
                                                if (1e10 >= 1.5) {
                                                    fmt.Println(v1, v6, fi11)
                                                    fmt.Println(v6, fi15, v1)
                                                }
                                                v4 = v4
                                                v25 := GenS4{F0: 53, F1: -417, F2: 100}
                                                _ = v25
                                            } else {
                                                v1 = v1
                                                var v26 string = (("a" + (v4 + ("a" + "foo bar"))) + v4)
                                                switch fi15 {
                                                case 1:
                                                    var v27 uint32 = (v6 & 65535)
                                                    _ = v27
                                                case 9223372036854775807:
                                                    v28 := GenS5{F0: 827, F1: (v8 << ((4294967295) & 63))}
                                                    var v29 string = v4
                                                    var v30 int8 = -128
                                                    _ = v28
                                                    _ = v29
                                                    _ = v30
                                                default:
                                                    switch (^(fi15)) {
                                                    case -263:
                                                        v4 = "hello"
                                                    case 740:
                                                        if (-128 != 1) {
                                                            fmt.Println(v26, v7, fi11)
                                                            var v31 uint64 = 2
                                                            var v32 int64 = (^(v1))
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                    case 0:
                                                        var v33 int64 = 72
                                                        _ = v33
                                                    }
                                                }
                                                _ = v26
                                            }
                                            v6 -= (v6 - (v6 * v6))
                                            v34 := GenS3{F0: 883, F1: 18446744073709551615}
                                            _ = fi15
                                            _ = v34
                                        }
                                    }
                                }
                                fmt.Println(fi11, fi11, v7)
                            } else {
                                v35 := GenS5{F0: (^(fi11)), F1: (v8 | v8)}
                                for fi36 := 0; fi36 < 8; fi36++ {
                                    switch (v4 + "hello") {
                                    case "foo bar":
                                        v4 = ("hello" + v4)
                                    }
                                    v37 := GenS4{F0: 33, F1: (v1 | v1), F2: -128}
                                    v6 = (v6 / ((4294967295) | 1))
                                    _ = fi36
                                    _ = v37
                                }
                                v2[(fi11 >> ((42) & 63))] = v6
                                _ = v35
                            }
                        }
                        _ = fi11
                    }
                case "\n\t":
                    fmt.Println(v7, v1)
                default:
                    if (!(!(-1 <= int8(v8)))) {
                        var v38 float64 = (float64(1e10) / 1e10)
                        _ = len(v3)
                        _ = v38
                    } else {
                        v1 -= v1
                        v3 = append(v3, v4)
                    }
                }
            }
        }
        _ = v10
    }
    v39 := GenS6{F0: 42, F1: 578, F2: -9223372036854775808}
    _ = len(v2)
    v4 = ""
    if (true || ((127 <= 0) || (!v7))) {
        v8 = v8
        for fi40 := 0; fi40 < 12; fi40++ {
            if (1 > 1) {
                var v41 uint64 = 363
                var v42 uint64 = (v41 * (v41 - (v41 - v41)))
                var v43 int = 2147483647
                _ = v41
                _ = v42
                _ = v43
            }
            _ = fi40
        }
    } else {
        for fi44 := 0; fi44 < 14; fi44++ {
            for fi45 := 0; fi45 < 10; fi45++ {
                v46 := GenS1{F0: 182, F1: -128, F2: -70}
                var v47 uint64 = 860
                var v48 uint16 = 1
                _ = fi45
                _ = v46
                _ = v47
                _ = v48
            }
            var v49 int16 = (v8 >> ((778) & 63))
            _ = fi44
            _ = v49
        }
        for fi50 := 0; fi50 < 4; fi50++ {
            v3 = append(v3, v4)
            _ = fi50
        }
    }
    for fi51 := 0; fi51 < 10; fi51++ {
        fmt.Println(fi51, v7, v8)
        var v52 bool = (v1 < (v1 & (^(v1))))
        for fi53 := 0; fi53 < 10; fi53++ {
            v54 := GenS2{F0: 188, F1: true}
            v8 += v8
            _ = fi53
            _ = v54
        }
        _ = fi51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
