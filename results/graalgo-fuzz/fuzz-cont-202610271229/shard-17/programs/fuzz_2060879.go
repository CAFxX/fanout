package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int16; F2 int8 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint; F1 int32; F2 uint8 }
type GenS9 struct { F0 int32; F1 uint64 }

func main() {
    var v1 int8 = -128
    var v2 uint8 = 2
    var v3 uint8 = 64
    v4 := []int{601, 9223372036854775807, 127}
    v5 := []string{"hello", "", "hello", "a"}
    v6 := make(map[uint32]int)
    switch "hello" {
    case "\n\t":
        v7 := GenS1{F0: v3}
        _ = v7
    case "hello":
        var v8 float64 = ((((1e10 + 1e10) - -2.25) * 1.5) - 1e10)
        _ = v8
    case "":
        for fi9 := 0; fi9 < 4; fi9++ {
            if false {
                var v10 int16 = -32768
                if true {
                    fmt.Println(v2)
                }
                _ = v10
            } else {
                switch fi9 {
                case 42:
                    if ((v2 / ((v2) | 1)) > 239) {
                        if ((("hello" + "hello") + ("foo bar" + ("a" + "\n\t"))) >= "hello") {
                            var v11 int64 = -307
                            _ = v11
                        } else {
                            v2 = v2
                        }
                    } else {
                        fmt.Println(v3, v3, v3)
                        v1 = v1
                    }
                default:
                    for fi12 := 0; fi12 < 3; fi12++ {
                        if (18446744073709551615 != 1) {
                            var v13 int8 = v1
                            fmt.Println(fi9)
                            _ = v13
                        } else {
                            var v14 int = -313
                            delete(v6, 1)
                            v14 += 1
                            _ = v14
                        }
                        v1 += 100
                        var v15 float64 = 3.14159
                        fmt.Println(v1)
                        _ = fi12
                        _ = v15
                    }
                }
                fmt.Println(v2)
                fmt.Println(fi9, fi9)
            }
            var v16 float32 = 0.0
            for fi17 := 0; fi17 < 4; fi17++ {
                if (0 >= 286) {
                    var v18 uint64 = 42
                    _ = v18
                }
                v16 -= 3.14159
                switch "foo bar" {
                case "":
                    v1 = v1
                case "foo bar":
                    v16 -= v16
                case "\n\t":
                    v19 := GenS1{F0: v3}
                    _ = v19
                }
                v2 += v3
                _ = fi17
            }
            _ = fi9
            _ = v16
        }
    }
    switch "\n\t" {
    case "hello":
        switch 9223372036854775807 {
        case 32767:
            v1 = int8(v3)
        case -245:
            var v20 float32 = (float32(1e10) / 1.0)
            _ = v20
        case -12:
            switch "\n\t" {
            case "\n\t":
                v4 = append(v4, -2147483648)
            case "a":
                var v21 float64 = (0.1 + 3.14159)
                _ = v21
            }
        default:
            for fi22 := 0; fi22 < 13; fi22++ {
                for fi23 := 0; fi23 < 17; fi23++ {
                    switch (fi23 * 32767) {
                    case 128:
                        if (!false) {
                            fmt.Println(fi23)
                        } else {
                            if (2 <= 274) {
                                var v24 float64 = ((((float64(0.0) / 0.1) * (1.5 * 0.0)) - 3.14159) + (0.1 + 3.14159))
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if true {
                                    fmt.Println(v1, v2)
                                }
                                _ = v24
                            }
                            var v25 int16 = 0
                            _ = v25
                        }
                    case -991:
                        v26 := GenS2{F0: 127}
                        _ = v26
                    case -904:
                        if true {
                            fmt.Println(v1)
                            v27 := GenS1{F0: (^(v3))}
                            var v28 int = (fi22 + 32768)
                            _ = v27
                            _ = v28
                        }
                    }
                    v29 := v6[0]
                    var v30 int8 = (v1 / (((v1 * (^(v1)))) | 1))
                    v31 := GenS3{F0: v29, F1: -933, F2: (v30 + (-(v30)))}
                    _ = fi23
                    _ = v29
                    _ = v30
                    _ = v31
                }
                for fi32 := 0; fi32 < 19; fi32++ {
                    delete(v6, 0)
                    _ = fi32
                }
                _ = fi22
            }
        }
    case "":
        fmt.Println(v2, v2)
    default:
        if true {
            if false {
                var v33 float64 = (0.0 - 1.5)
                _ = v33
            } else {
                _ = len(v5)
                var v34 string = "a"
                var v35 int16 = 42
                _ = v34
                _ = v35
            }
            v3 = 255
            if ((42 == -1) || true) {
                v3 -= (v3 | v3)
                v36 := GenS4{F0: 2, F1: 9223372036854775807}
                _ = v36
            } else {
                var v37 int8 = v1
                v38 := GenS5{F0: 166, F1: 4294967295}
                fmt.Println(v3, v37)
                _ = v37
                _ = v38
            }
        }
    }
    var v39 int = -1
    if true {
        v40 := GenS6{F0: 2}
        _ = len(v6)
        _ = v40
    }
    if (!(false || false)) {
        var v41 int32 = -912
        _ = v41
    } else {
        fmt.Println(v2, v39, v1)
        v1 += v1
        v42 := GenS7{F0: v39}
        _ = v42
    }
    if (false && (882 < 482)) {
        if true {
            v43 := GenS8{F0: uint(v3), F1: 42, F2: 128}
            for fi44 := 0; fi44 < 4; fi44++ {
                v3 = 2
                v2 += (^(v2))
                switch 128 {
                case 127:
                    if (((v1 * v1) >= (v1 | 100)) && (9223372036854775807 <= 9223372036854775807)) {
                        v45 := GenS1{F0: v2}
                        _ = v45
                    }
                case 1:
                    v1 -= -123
                }
                _ = fi44
            }
            _ = v43
        }
    }
    switch (-(v39)) {
    case 2147483647:
        for fi46 := 0; fi46 < 7; fi46++ {
            var v47 uint = 18446744073709551615
            _ = fi46
            _ = v47
        }
    case -1:
        switch "\n\t" {
        case "":
            var v48 int = -110
            _ = v48
        case "\n\t":
            v49 := GenS4{F0: 0, F1: 9223372036854775807}
            _ = v49
        case "foo bar":
            var v50 int16 = -1
            _ = v50
        }
    default:
        var v51 uint16 = 2
        _ = v51
    }
    for fi52 := 0; fi52 < 17; fi52++ {
        var v53 uint16 = 887
        v1 = v1
        _ = fi52
        _ = v53
    }
    if false {
        v54 := GenS9{F0: 42, F1: 1}
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
