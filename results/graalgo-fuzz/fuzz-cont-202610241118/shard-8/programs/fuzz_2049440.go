package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int8; F2 int32 }
type GenS4 struct { F0 uint64; F1 uint; F2 int }
type GenS5 struct { F0 uint32; F1 int; F2 uint }
type GenS6 struct { F0 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"a", ("foo bar" + "hello")}
    var v3 uint = 511
    v4 := make(map[int]string)
    switch "\n\t" {
    case "a":
        if ("foo bar" != "hello") {
            v3 += 42
        }
    }
    v3 = (^(v3))
    v3 += v3
    v3 = v3
    if (981 > (v3 >> ((v3) & 63))) {
        if (128 >= -555) {
            v5 := GenS1{F0: -878, F1: (^(v3))}
            v2 = append(v2, "a")
            _ = v5
        }
    }
    _ = len(v2)
    v6 := GenS1{F0: 42, F1: (v3 & v3)}
    var v7 bool = false
    switch 855 {
    case 127:
        if v7 {
            switch "foo bar" {
            case "":
                var v8 float64 = float64(((-2.25 + -2.25) + (((-2.25 - 0.0) / (1.5 + 1.5)) - 1e10)))
                _ = v8
            }
            v9 := GenS1{F0: -1, F1: (^(v3))}
            _ = len(v2)
            _ = v9
        } else {
            v10 := GenS2{F0: 189}
            v3 += 2
            v11 := GenS3{F0: v7, F1: 42, F2: -960}
            _ = v10
            _ = v11
        }
    default:
        v12 := GenS3{F0: v7, F1: 0, F2: 2147483647}
        _ = v12
    }
    v7 = (float32((float64(0.1) - 0.1)) != (0.0 / -2.25))
    var v13 uint16 = 0
    fmt.Println(v13, v3, v13)
    switch -683 {
    case -2147483648:
        v14 := GenS4{F0: 992, F1: v3, F2: 128}
        _ = v14
    }
    if true {
        var v15 uint64 = 746
        switch 2147483647 {
        case 175:
            fmt.Println(v7, v15)
        case 42:
            v16 := GenS5{F0: 210, F1: 32768, F2: v3}
            _ = v16
        case -1:
            var v17 uint16 = (^(v13))
            _ = v17
        default:
            v3 += 2
        }
        _ = v15
    } else {
        v3 += v3
        if v7 {
            if false {
                for fi18 := 0; fi18 < 2; fi18++ {
                    v13 = v13
                    switch (^(fi18)) {
                    case -2147483648:
                        var v19 int8 = 78
                        _ = v19
                    case 0:
                        fmt.Println(v3, fi18)
                    case -115:
                        v20 := GenS6{F0: fi18}
                        _ = v20
                    }
                    var v21 string = ("hello" + "\n\t")
                    _ = fi18
                    _ = v21
                }
            }
            v3 += (v3 << ((627) & 63))
        } else {
            switch 32768 {
            case -1:
                v3 = (v3 + (v3 >> (((v3 / ((467) | 1))) & 63)))
            case 536:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                v22 := GenS6{F0: 127}
                _ = v22
            }
        }
        if v7 {
            v23 := GenS2{F0: 1}
            var v24 uint64 = 4294967295
            if false {
                var v25 int32 = 42
                var v26 uint64 = v24
                _ = v25
                _ = v26
            }
            _ = v23
            _ = v24
        }
    }
    switch "a" {
    case "a":
        switch -629 {
        case -2147483648:
            for fi27 := 0; fi27 < 12; fi27++ {
                _ = len(v4)
                if (2147483647 == 441) {
                    switch "\n\t" {
                    case "hello":
                        _ = len(v2)
                    }
                } else {
                    v13 = 42
                    if (!((v3 >= (v3 - (v3 & (^(v3))))) || false)) {
                        for fi28 := 0; fi28 < 9; fi28++ {
                            var v29 float64 = 1e10
                            _ = fi28
                            _ = v29
                        }
                        var v30 float32 = 3.14159
                        if ((940 <= -864) && (42 > 42)) {
                            v7 = v7
                        }
                        _ = v30
                    } else {
                        v13 -= (v13 | (v13 << ((v3) & 63)))
                        var v31 string = ("" + ("\n\t" + "hello"))
                        var v32 float64 = float64(1e10)
                        _ = v31
                        _ = v32
                    }
                    v33 := GenS6{F0: fi27}
                    _ = v33
                }
                v7 = (!((((1.5 * ((float64(1e10) * (0.1 - (-2.25 + 1.5))) * -2.25)) + 0.1) / (1.5 / 1e10)) == float64((1.5 * -2.25))))
                _ = fi27
            }
        case 128:
            delete(v1, 42)
        case -1:
            var v34 bool = (!(-2.25 == (0.1 + 0.1)))
            _ = v34
        default:
            for fi35 := 0; fi35 < 13; fi35++ {
                v36 := v4[fi35]
                if (-752 != -999) {
                    if (!v7) {
                        if v7 {
                            var v37 float64 = (3.14159 + 3.14159)
                            fmt.Println(v13)
                            var v38 uint32 = 253
                            _ = v37
                            _ = v38
                        }
                        if (!v7) {
                            var v39 int64 = -9223372036854775808
                            switch (v36 + (((((v36 + v36) + "foo bar") + "a") + v36) + ("foo bar" + (v36 + v36)))) {
                            case "foo bar":
                                if true {
                                    if (-100 <= -92) {
                                        v3 += (v3 | v3)
                                    }
                                    if (!(v7 || (0 != 18446744073709551615))) {
                                        var v40 uint8 = uint8(v13)
                                        _ = v40
                                    }
                                    var v41 int32 = -139
                                    _ = v41
                                }
                            default:
                                fmt.Println(v13, v36, fi35)
                            }
                            _ = v39
                        }
                    }
                } else {
                    var v42 int32 = 2147483647
                    var v43 bool = v7
                    var v44 uint64 = 1
                    _ = v42
                    _ = v43
                    _ = v44
                }
                _ = fi35
                _ = v36
            }
        }
    case "\n\t":
        v7 = (-1 != 0)
    case "":
        if ((-2147483648 >= 0) && ((-1 <= 1) && (127 < 128))) {
            if v7 {
                v7 = v7
                var v45 float64 = 0.1
                _ = v45
            } else {
                _ = len(v2)
            }
            v46 := v1[2147483647]
            _ = v46
        } else {
            switch ("hello" + "hello") {
            case "a":
                var v47 int = -1
                _ = v47
            case "\n\t":
                fmt.Println(v7, v3, v7)
            case "hello":
                v7 = v7
            }
        }
    }
    fmt.Println(v3, v13, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
