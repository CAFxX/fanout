package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16 }
func genFn0() int64 {
    switch "foo bar" {
    case "foo bar":
        var v4 uint = 0
        _ = v4
    case "\n\t":
        for fi5 := 0; fi5 < 5; fi5++ {
            for fi6 := 0; fi6 < 17; fi6++ {
                fmt.Println(fi6)
                var v7 uint8 = 129
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
    default:
        if ((24 < 0) || false) {
            v8 := []int{2147483647}
            switch 941 {
            case -9223372036854775808:
                var v9 uint64 = 42
                _ = v9
            }
            _ = v8
        }
    }
    switch 1 {
    case 244:
        v10 := []int{-9223372036854775808}
        _ = v10
    case 920:
        var v11 uint64 = 4294967295
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        switch 967 {
        case 239:
            var v13 bool = (35 == 0)
            _ = v13
        case 567:
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 uint16 = 32767
                var v16 string = ((("\n\t" + "a") + "foo bar") + ("\n\t" + "a"))
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        fmt.Println(fi12)
        if false {
            if false {
                v17 := GenS1{F0: 2, F1: 255, F2: 0}
                var v18 int16 = 42
                _ = v17
                _ = v18
            }
            v19 := GenS1{F0: 1, F1: 0, F2: 1}
            var v20 uint = 42
            _ = v19
            _ = v20
        }
        _ = fi12
    }
    var v21 uint = 0
    switch "\n\t" {
    case "\n\t":
        if (true || false) {
            fmt.Println(v21)
            v21 -= v21
        }
    case "":
        switch 2147483647 {
        case -118:
            v22 := GenS2{F0: 32767, F1: 2147483647}
            _ = v22
        case 128:
            if true {
                var v23 int = -2147483648
                fmt.Println(v23, v21)
                var v24 int16 = -946
                _ = v23
                _ = v24
            }
        default:
            switch 32768 {
            case -2147483648:
                v21 -= (v21 + v21)
            case 42:
                v21 = v21
            case 1:
                v25 := GenS2{F0: 811, F1: 2147483647}
                _ = v25
            default:
                v26 := GenS3{F0: 0}
                _ = v26
            }
        }
    case "hello":
        if (1 != 2) {
            switch "\n\t" {
            case "hello":
                if true {
                    v27 := GenS1{F0: 1, F1: 2, F2: 2}
                    v28 := []int{-416, -1, -798, -801}
                    for fi29 := 0; fi29 < 10; fi29++ {
                        var v30 int16 = -91
                        _ = len(v28)
                        var v31 int8 = -85
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = v27
                    _ = v28
                } else {
                    v32 := []int{128, 442, -1, -2147483648}
                    for fi33 := 0; fi33 < 18; fi33++ {
                        v21 = v21
                        for fi34 := 0; fi34 < 9; fi34++ {
                            v32 = append(v32, fi33)
                            fmt.Println(fi33, fi34)
                            v21 += v21
                            v21 = (v21 + v21)
                            _ = fi34
                        }
                        v35 := GenS4{F0: 100}
                        _ = fi33
                        _ = v35
                    }
                    fmt.Println(v21)
                    _ = v32
                }
            case "foo bar":
                v36 := GenS1{F0: 161, F1: 127, F2: 153}
                _ = v36
            }
            fmt.Println(v21)
            v21 -= v21
        } else {
            var v37 uint32 = 543
            for fi38 := 0; fi38 < 8; fi38++ {
                v37 += v37
                v39 := GenS2{F0: 42, F1: -994}
                for fi40 := 0; fi40 < 10; fi40++ {
                    var v41 int = 42
                    var v42 bool = true
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                _ = fi38
                _ = v39
            }
            _ = v37
        }
    default:
        v43 := GenS2{F0: 32767, F1: 42}
        _ = v43
    }
    _ = v21
    return -607
}
func genFn1() string {
    v44 := GenS1{F0: 42, F1: 118, F2: 1}
    v45 := []int{2147483647, 9223372036854775807, 0}
    var v46 uint = 386
    if false {
        var v47 float64 = -2.25
        switch 128 {
        case -1:
            var v48 int16 = -1
            _ = v48
        case -233:
            var v49 int64 = 570
            _ = v49
        }
        for fi50 := 0; fi50 < 13; fi50++ {
            var v51 string = "a"
            var v52 uint16 = 2
            _ = fi50
            _ = v51
            _ = v52
        }
        _ = v47
    }
    _ = v44
    _ = v45
    _ = v46
    return ""
}

func main() {
    var v1 float32 = float32(-2.25)
    var v2 uint64 = 2
    v3 := make(map[uint32]int)
    v1 += v1
    var v53 int = 32768
    var v54 float32 = ((v1 / v1) * v1)
    var v55 int64 = -242
    switch "hello" {
    case "":
        v56 := GenS4{F0: 0}
        _ = v56
    case "hello":
        switch (v53 % ((4294967295) | 1)) {
        case 42:
            if true {
                var v57 float64 = 1e10
                var v58 int8 = -59
                delete(v3, 2147483648)
                _ = v57
                _ = v58
            }
        case -9223372036854775808:
            v59 := []int{(-(v53))}
            _ = v59
        }
    default:
        var v60 int16 = -111
        _ = v60
    }
    v2 += (v2 << ((2) & 63))
    switch "\n\t" {
    case "a":
        for fi61 := 0; fi61 < 18; fi61++ {
            v62 := []int{(-(fi61)), fi61, v53, fi61}
            v1 -= v1
            if false {
                v63 := GenS4{F0: -128}
                switch "a" {
                case "\n\t":
                    if (1 <= 9223372036854775807) {
                        var v64 uint32 = uint32(v2)
                        v55 += v55
                        _ = v64
                    } else {
                        fmt.Println(v53)
                        switch int(v55) {
                        case -444:
                            if (!(-85 < -2147483648)) {
                                v55 -= v55
                                v62 = append(v62, (fi61 + v53))
                                fmt.Println(fi61, fi61)
                            }
                        case -1:
                            var v65 int = (fi61 << ((930) & 63))
                            _ = v65
                        case 9223372036854775807:
                            v66 := GenS4{F0: 127}
                            _ = v66
                        default:
                            switch v53 {
                            case -2147483648:
                                var v67 float32 = v54
                                _ = v67
                            case 9223372036854775807:
                                var v68 int = fi61
                                _ = v68
                            }
                        }
                        for fi69 := 0; fi69 < 3; fi69++ {
                            v54 = v1
                            v70 := GenS1{F0: 255, F1: 0, F2: 0}
                            if false {
                                v54 = (v1 - (0.1 - (0.0 - (v54 + float32((-2.25 - ((3.14159 * 1.5) / 3.14159)))))))
                            }
                            _ = fi69
                            _ = v70
                        }
                    }
                case "":
                    var v71 uint16 = 32767
                    _ = v71
                case "a":
                    var v72 bool = ((true && (!(668 == 0))) || true)
                    _ = v72
                }
                _ = v63
            }
            _ = fi61
            _ = v62
        }
    case "\n\t":
        v2 -= 40
    default:
        v73 := []int{32768, v53, v53}
        _ = v73
    }
    v74 := []int{v53, (^(v53)), v53}
    if (!false) {
        var v75 uint8 = 203
        _ = v75
    }
    fmt.Println(v54, v1, v54)
    fmt.Println(v2, v2, v1)
    fmt.Println(v2, v53, v54)
    v76 := GenS5{F0: 42}
    switch v53 {
    case 781:
        var v77 float64 = (((((((1.5 - float64(3.14159)) * (1e10 + (float64(0.0) + -2.25))) - (3.14159 + ((1e10 - 0.0) * 3.14159))) / 0.1) * 3.14159) + (0.1 / 1.0)) - 1e10)
        _ = v77
    default:
        v53 = (v53 % ((v53) | 1))
    }
    switch "foo bar" {
    case "":
        v78 := GenS4{F0: 0}
        _ = v78
    default:
        if false {
            var v79 int8 = 127
            _ = v79
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v74:", len(v74), v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v53
    _ = v54
    _ = v55
    _ = v74
    _ = v76
}
