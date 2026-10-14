package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int16; F2 int }
type GenS4 struct { F0 uint16; F1 string; F2 uint16 }
func genFn0(p0 int64, p1 int8) string {
    var v7 uint8 = 127
    v7 = 1
    p0 -= (p0 >> ((255) & 63))
    v8 := GenS1{F0: v7, F1: 42}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return ""
}
func genFn1() int8 {
    if (42 > 4294967295) {
        var v9 int32 = -2147483648
        if (!true) {
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int64 = 42
                switch "a" {
                case "\n\t":
                    fmt.Println(v9, fi10, v9)
                default:
                    fmt.Println(fi10, v9, fi10)
                }
                _ = fi10
                _ = v11
            }
            fmt.Println(v9)
        }
        var v12 uint64 = 0
        _ = v12
        _ = v9
    } else {
        v13 := GenS1{F0: 42, F1: 32767}
        if false {
            v14 := GenS1{F0: 255, F1: 1}
            switch (("foo bar" + ("a" + ("" + ("" + "hello")))) + "a") {
            case "a":
                var v15 float32 = (float32((((-2.25 * -2.25) * 0.1) / -2.25)) * 1.5)
                _ = v15
            }
            var v16 uint8 = 223
            _ = v14
            _ = v16
        } else {
            var v17 uint8 = 0
            _ = v17
        }
        _ = v13
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 int64 = -1
        var v20 int32 = 42
        var v21 int16 = 672
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    switch "foo bar" {
    case "hello":
        var v22 uint64 = 2
        _ = v22
    }
    var v23 float32 = (((((1e10 / float32(-2.25)) * ((1e10 / float32(float64(-2.25))) * ((float32(0.0) + (1e10 / -2.25)) * (1.5 + (1.5 * ((1.5 / float32(3.14159)) * 3.14159)))))) + (1.5 * -2.25)) + 0.0) - 1.5)
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 float64 = 3.14159
        v26 := []int{(fi24 * 994), 42, fi24}
        fmt.Println(v23, v23)
        _ = fi24
        _ = v25
        _ = v26
    }
    _ = v23
    return 100
}

func main() {
    var v1 int64 = 1
    var v2 int32 = 132
    var v3 uint64 = 42
    var v4 uint = 691
    var v5 string = "hello"
    v6 := make(map[uint32]uint32)
    v1 += -776
    for fi27 := 0; fi27 < 7; fi27++ {
        v1 = (v1 >> (((v4 << (((v4 & 314)) & 63))) & 63))
        var v28 float32 = 0.1
        fmt.Println(fi27, v28, fi27)
        _ = fi27
        _ = v28
    }
    if true {
        v29 := []int{-313, 32767}
        _ = v29
    } else {
        v3 = (^(v3))
        switch 9223372036854775807 {
        case 32768:
            var v30 int32 = 2147483647
            _ = v30
        case 128:
            if false {
                v2 = v2
            }
        }
    }
    var v31 uint8 = 90
    v3 = (v3 % (((^(v3))) | 1))
    fmt.Println(v4)
    var v32 int32 = (v2 + v2)
    var v33 string = v5
    v4 = (v4 + (v4 | (v4 | 255)))
    v5 = v33
    v31 -= (v31 << ((v4) & 63))
    if ((1e10 != 1e10) && true) {
        fmt.Println(v33, v32, v32)
        var v34 bool = (!false)
        _ = v34
    } else {
        v1 = v1
        _ = len(v6)
    }
    switch -362 {
    case 128:
        for fi35 := 0; fi35 < 15; fi35++ {
            var v36 uint16 = 1
            v37 := GenS2{F0: v4}
            for fi38 := 0; fi38 < 3; fi38++ {
                fmt.Println(v2, v3, v2)
                switch fi38 {
                case -1:
                    v5 = ""
                case 42:
                    switch (fi38 % ((int(v36)) | 1)) {
                    case 32767:
                        v1 -= v1
                    case -278:
                        if (v36 > v36) {
                            if (false && true) {
                                fmt.Println(v31, fi38, v36)
                                v39 := GenS1{F0: 2, F1: v36}
                                var v40 uint64 = v3
                                _ = v39
                                _ = v40
                            } else {
                                delete(v6, 701)
                                var v41 int8 = 127
                                fmt.Println(v36, v4, v41)
                                if true {
                                    if (false && (!true)) {
                                        fmt.Println(v5, v36)
                                    }
                                    switch "a" {
                                    case "hello":
                                        v6[968] = 2147483648
                                    case "":
                                        v42 := GenS1{F0: v31, F1: (^(v36))}
                                        _ = v42
                                    case "foo bar":
                                        v43 := []int{fi38, -9223372036854775808, fi35}
                                        _ = v43
                                    default:
                                        v5 = ("\n\t" + "foo bar")
                                    }
                                }
                                delete(v6, 42)
                                v32 = int32(v3)
                                _ = v41
                            }
                        }
                    case -9223372036854775808:
                        v33 = "hello"
                    }
                default:
                    fmt.Println(v36)
                }
                v6[850] = uint32(v3)
                _ = fi38
            }
            _ = fi35
            _ = v36
            _ = v37
        }
    case 9223372036854775807:
        for fi44 := 0; fi44 < 10; fi44++ {
            var v45 string = ""
            v46 := GenS3{F0: (v4 << (((v4 ^ (v4 & v4))) & 63)), F1: 502, F2: fi44}
            _ = fi44
            _ = v45
            _ = v46
        }
    case 0:
        var v47 int32 = 870
        _ = v47
    default:
        fmt.Println(v32, v2)
    }
    _ = len(v6)
    v2 -= (v2 * int32(v4))
    if (730 != 825) {
        if (-1 != 0) {
            v48 := GenS1{F0: 2, F1: 1}
            _ = v48
        } else {
            for fi49 := 0; fi49 < 5; fi49++ {
                var v50 int = 665
                _ = fi49
                _ = v50
            }
            var v51 uint = (v4 / (((^(v4))) | 1))
            _ = v51
        }
        switch -734 {
        case 42:
            var v52 float32 = float32(0.1)
            _ = v52
        case -2147483648:
            fmt.Println(v31)
        case -709:
            fmt.Println(v32)
        }
    }
    v53 := GenS4{F0: 255, F1: v33, F2: uint16(v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}
