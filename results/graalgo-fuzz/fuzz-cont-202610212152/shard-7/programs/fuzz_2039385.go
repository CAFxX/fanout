package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
func genFn0() string {
    v5 := GenS1{F0: 328}
    if (!(false && true)) {
        var v6 int16 = 1
        switch (("foo bar" + "") + "a") {
        case "":
            if (18446744073709551615 >= 42) {
                if true {
                    if true {
                        for fi7 := 0; fi7 < 3; fi7++ {
                            v8 := []int{(fi7 % ((fi7) | 1)), fi7}
                            fmt.Println(v6, v6)
                            _ = fi7
                            _ = v8
                        }
                        for fi9 := 0; fi9 < 1; fi9++ {
                            var v10 int32 = 427
                            v11 := []int{(fi9 * -1)}
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                    } else {
                        switch "" {
                        case "a":
                            var v12 float32 = ((float32(1.5) - ((1e10 * (0.1 * (0.1 - float32(((float64(((-2.25 - (1.5 / (float32(((1.5 - -2.25) * (((0.1 / 1.0) * 0.1) - float64((1.5 - 0.0))))) / 1.0))) + 0.1)) * 0.0) / 1.0))))) / 1.5)) * -2.25)
                            _ = v12
                        }
                        var v13 bool = true
                        _ = v13
                    }
                } else {
                    v6 = -32768
                }
                switch -1 {
                case -1:
                    v6 -= (^(v6))
                case 4294967295:
                    var v14 int16 = (v6 & v6)
                    _ = v14
                }
            } else {
                v6 -= v6
            }
        case "\n\t":
            v15 := GenS2{F0: 127}
            _ = v15
        default:
            v6 = (v6 - (^(v6)))
        }
        fmt.Println(v6)
        _ = v6
    } else {
        var v16 float64 = 1.5
        var v17 uint32 = 4294967295
        v18 := GenS3{F0: 42, F1: 255}
        _ = v16
        _ = v17
        _ = v18
    }
    switch (("a" + "\n\t") + "") {
    case "hello":
        var v19 int = 128
        _ = v19
    case "foo bar":
        if (4294967295 >= 1) {
            for fi20 := 0; fi20 < 4; fi20++ {
                if true {
                    var v21 float32 = 1e10
                    var v22 int16 = 285
                    switch fi20 {
                    case 127:
                        v21 += 1.5
                    default:
                        var v23 int32 = -660
                        _ = v23
                    }
                    _ = v21
                    _ = v22
                } else {
                    v24 := []int{fi20, (fi20 + -563), 127}
                    fmt.Println(fi20)
                    _ = v24
                }
                var v25 float64 = ((((float64((0.0 + float32(-2.25))) + 0.1) * float64(0.1)) + float64(0.0)) * float64((0.1 * 3.14159)))
                v25 = (((v25 * 0.1) / 0.1) + ((float64(1.5) * v25) / 1.5))
                _ = fi20
                _ = v25
            }
            v26 := []int{-127, 127, 9223372036854775807}
            if (469 == 2) {
                var v27 int64 = -1
                var v28 bool = false
                _ = v27
                _ = v28
            } else {
                var v29 bool = true
                _ = v29
            }
            _ = v26
        } else {
            var v30 int32 = -2147483648
            switch -775 {
            case 42:
                var v31 uint64 = 1
                _ = v31
            case -1:
                v32 := []int{}
                _ = v32
            case -390:
                switch "\n\t" {
                case "a":
                    v33 := []int{42, 128}
                    _ = v33
                case "\n\t":
                    var v34 uint32 = 571
                    _ = v34
                case "foo bar":
                    v30 -= v30
                default:
                    switch "foo bar" {
                    case "foo bar":
                        var v35 bool = (!(127 <= 1))
                        _ = v35
                    case "a":
                        var v36 int64 = 1
                        _ = v36
                    }
                }
            default:
                var v37 float64 = -2.25
                _ = v37
            }
            _ = v30
        }
    default:
        for fi38 := 0; fi38 < 11; fi38++ {
            fmt.Println(fi38)
            var v39 int16 = -32768
            v40 := GenS2{F0: 255}
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    v41 := []int{-2147483648, 9223372036854775807, 32768}
    var v42 uint8 = 0
    _ = v41
    _ = v42
    _ = v5
    return ("hello" + (("hello" + (("foo bar" + "foo bar") + ("foo bar" + "\n\t"))) + ("" + "hello")))
}

func main() {
    var v1 bool = true
    var v2 uint = 1
    v3 := make(map[int]int)
    var v4 int = 9223372036854775807
    v43 := GenS1{F0: 42}
    v2 += (v2 >> ((v2) & 63))
    var v44 int32 = -2147483648
    switch "\n\t" {
    case "a":
        var v45 float32 = 0.0
        _ = v45
    default:
        v4 += v4
    }
    fmt.Println(v1, v4)
    fmt.Println(v44)
    v2 = v2
    v46 := GenS1{F0: 65535}
    if v1 {
        switch v4 {
        case -463:
            v47 := GenS4{F0: 867}
            _ = v47
        case 4294967295:
            v48 := GenS5{F0: true}
            _ = v48
        }
        v44 = (v44 << (((v2 ^ v2)) & 63))
        v49 := []int{127, v4, -543, (v4 | (v4 & (v4 + int(v2))))}
        _ = v49
    } else {
        fmt.Println(v44, v4)
        for fi50 := 0; fi50 < 20; fi50++ {
            switch "a" {
            case "":
                v4 = v4
            case "hello":
                v51 := []int{-795, (v4 | v4), fi50, 32768}
                _ = v51
            case "foo bar":
                if v1 {
                    fmt.Println(v1, v2)
                    var v52 bool = v1
                    _ = v52
                } else {
                    var v53 uint8 = 40
                    var v54 int = fi50
                    for fi55 := 0; fi55 < 18; fi55++ {
                        v1 = ((false && v1) && (((v1 && true) || ((v2 ^ v2) > 2)) || (v1 && true)))
                        _ = fi55
                    }
                    _ = v53
                    _ = v54
                }
            }
            v56 := []int{2147483647}
            fmt.Println(v1)
            _ = fi50
            _ = v56
        }
    }
    var v57 bool = ((v2 == v2) || v1)
    for fi58 := 0; fi58 < 11; fi58++ {
        for fi59 := 0; fi59 < 17; fi59++ {
            fmt.Println(fi58)
            v1 = false
            var v60 int64 = 760
            _ = fi59
            _ = v60
        }
        v61 := []int{(v4 & fi58), fi58, (-(fi58)), (fi58 * fi58)}
        if v1 {
            v2 = 833
            v4 += v4
            if (42 < 2) {
                if true {
                    var v62 bool = v1
                    switch ("\n\t" + "\n\t") {
                    case "a":
                        v63 := GenS2{F0: 1}
                        _ = v63
                    }
                    _ = v62
                } else {
                    for fi64 := 0; fi64 < 2; fi64++ {
                        v2 -= v2
                        switch "foo bar" {
                        case "":
                            fmt.Println(v57)
                        case "hello":
                            v57 = v57
                        case "foo bar":
                            var v65 uint64 = 0
                            _ = v65
                        }
                        _ = fi64
                    }
                    v44 += (v44 * v44)
                }
            }
        } else {
            var v66 uint32 = 4294967295
            var v67 int = (v4 - (v4 * (fi58 ^ -20)))
            _ = v66
            _ = v67
        }
        _ = fi58
        _ = v61
    }
    if (true || (2 < 1)) {
        var v68 int = (v4 % ((-2147483648) | 1))
        _ = v68
    }
    v57 = (0.1 == -2.25)
    for fi69 := 0; fi69 < 12; fi69++ {
        v4 = (-(v4))
        _ = fi69
    }
    if v57 {
        for fi70 := 0; fi70 < 3; fi70++ {
            v1 = ((!v1) && v1)
            var v71 uint64 = 42
            _ = fi70
            _ = v71
        }
        v72 := GenS2{F0: 48}
        v57 = false
        _ = v72
    }
    v1 = true
    var v73 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v57:", v57)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v46
    _ = v57
    _ = v73
}
