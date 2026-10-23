package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 int64; F1 uint32; F2 uint32 }

func main() {
    v1 := []int{-2147483648, 4294967295, 128}
    v2 := make(map[int]int)
    var v3 bool = (!(552 <= 2))
    v4 := make(map[int]string)
    var v5 int = -2147483648
    v6 := make(map[string]bool)
    var v7 int32 = 2147483647
    switch "a" {
    case "hello":
        fmt.Println(v3, v3)
    case "\n\t":
        v3 = ((!true) && (2 > 152))
    case "a":
        switch v5 {
        case 32767:
            v5 += (v5 + v5)
        default:
            var v8 uint16 = 65535
            _ = v8
        }
    }
    var v9 uint = 18446744073709551615
    if v3 {
        if (!v3) {
            v10 := GenS1{F0: (v9 - (^(v9))), F1: 42, F2: (-(v7))}
            _ = v10
        } else {
            v11 := GenS1{F0: v9, F1: -2147483648, F2: v7}
            if (1.5 == 3.14159) {
                v12 := GenS2{F0: 0}
                var v13 uint32 = 2147483647
                _ = v12
                _ = v13
            } else {
                v9 += v9
                if (!true) {
                    switch v5 {
                    case -782:
                        v14 := GenS1{F0: v9, F1: (v7 >> ((42) & 63)), F2: (v7 & (v7 << ((v9) & 63)))}
                        _ = v14
                    case 32767:
                        v9 += v9
                    case -82:
                        var v15 int64 = 9223372036854775807
                        _ = v15
                    default:
                        switch v5 {
                        case -2147483648:
                            v16 := GenS2{F0: 2}
                            _ = v16
                        default:
                            var v17 int = 42
                            _ = v17
                        }
                    }
                } else {
                    if v3 {
                        v18 := GenS2{F0: 833}
                        v19 := GenS3{F0: v5}
                        v20 := GenS4{F0: (v9 << ((v9) & 63))}
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 10; fi21++ {
                        if (4294967295 == 65535) {
                            if v3 {
                                var v22 uint64 = 42
                                fmt.Println(fi21, v9, v9)
                                var v23 bool = false
                                _ = v22
                                _ = v23
                            }
                            v7 = 1
                        }
                        v24 := GenS5{F0: 2, F1: (18446744073709551615 > v9), F2: 1}
                        _ = fi21
                        _ = v24
                    }
                }
            }
            _ = v11
        }
        if ((0.1 / 1.0) >= 1.5) {
            for fi25 := 0; fi25 < 16; fi25++ {
                var v26 float32 = ((-2.25 - -2.25) / (float32(float64(1e10)) / (0.1 * 3.14159)))
                _ = fi25
                _ = v26
            }
        } else {
            v1 = append(v1, v5)
        }
        var v27 int16 = 0
        _ = v27
    } else {
        var v28 int8 = -128
        fmt.Println(v7, v7)
        _ = v28
    }
    var v29 int16 = -307
    v29 = v29
    var v30 int16 = v29
    switch "foo bar" {
    case "foo bar":
        switch ("a" + ("hello" + "a")) {
        case "foo bar":
            var v31 int = v5
            _ = v31
        case "":
            var v32 uint = v9
            _ = v32
        case "hello":
            for fi33 := 0; fi33 < 6; fi33++ {
                if v3 {
                    v34 := GenS6{F0: 42, F1: (v7 >> ((v9) & 63))}
                    for fi35 := 0; fi35 < 18; fi35++ {
                        if (956 != 1) {
                            var v36 uint32 = 2
                            switch "" {
                            case "a":
                                v37 := GenS6{F0: 0, F1: (v7 * 42)}
                                if (77 >= -68) {
                                    var v38 uint = v9
                                    fmt.Println(v38, v3)
                                    _ = v38
                                }
                                _ = v37
                            case "":
                                v39 := GenS3{F0: fi35}
                                _ = v39
                            case "hello":
                                fmt.Println(v9)
                            }
                            var v40 uint64 = 18446744073709551615
                            _ = v36
                            _ = v40
                        } else {
                            v29 = (v29 + 887)
                            v3 = v3
                            if v3 {
                                var v41 uint64 = 42
                                _ = v41
                            }
                        }
                        switch -2147483648 {
                        case 42:
                            v42 := GenS3{F0: fi35}
                            _ = v42
                        case -2147483648:
                            v43 := GenS7{F0: int64(fi33), F1: 776, F2: 4294967295}
                            _ = v43
                        case 128:
                            v3 = (42 == v5)
                        }
                        v7 -= (v7 & 2147483647)
                        _ = fi35
                    }
                    if v3 {
                        for fi44 := 0; fi44 < 7; fi44++ {
                            v29 = v29
                            if v3 {
                                v5 -= (fi33 ^ int(v29))
                            }
                            _ = fi44
                        }
                        v29 = -466
                        var v45 string = "a"
                        _ = v45
                    }
                    _ = v34
                }
                _ = fi33
            }
        }
    }
    switch v5 {
    case 32767:
        switch "\n\t" {
        case "":
            var v46 int = (v5 & (v5 * (v5 * v5)))
            _ = v46
        default:
            fmt.Println(v7)
        }
    case 128:
        var v47 int = (v5 - v5)
        _ = v47
    case -2147483648:
        v48 := v6["foo bar"]
        _ = v48
    }
    v3 = (!(70 != (v5 >> ((v9) & 63))))
    switch v5 {
    case 9223372036854775807:
        switch "a" {
        case "hello":
            fmt.Println(v9, v5, v9)
        }
    case -2147483648:
        v7 += v7
    case 32768:
        if v3 {
            v5 = (v5 + (v5 ^ (v5 << (((v9 ^ (v9 ^ (^(v9))))) & 63))))
        }
    default:
        for fi49 := 0; fi49 < 18; fi49++ {
            v4[127] = (("hello" + "\n\t") + ("foo bar" + "a"))
            _ = fi49
        }
    }
    v50 := GenS3{F0: (v5 / (((v5 & (v5 ^ (v5 & -9223372036854775808)))) | 1))}
    v3 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v7
    _ = v9
}
