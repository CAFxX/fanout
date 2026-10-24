package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32; F1 uint; F2 int }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int32; F1 string; F2 bool }
type GenS13 struct { F0 int8; F1 uint16 }
func genFn0() string {
    var v7 uint32 = 65535
    if true {
        if true {
            v7 -= v7
        }
        fmt.Println(v7)
        v7 = v7
    } else {
        var v8 int8 = 127
        v8 = (^(v8))
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 string = "a"
                v7 += (^(v7))
                v8 = v8
                _ = fi10
                _ = v11
            }
            v8 += v8
            _ = fi9
        }
        _ = v8
    }
    switch 32767 {
    case 2147483647:
        for fi12 := 0; fi12 < 17; fi12++ {
            switch 32767 {
            case 1:
                v13 := GenS1{F0: 255, F1: 42}
                _ = v13
            default:
                var v14 int32 = -242
                _ = v14
            }
            _ = fi12
        }
    }
    if true {
        for fi15 := 0; fi15 < 2; fi15++ {
            fmt.Println(v7, fi15)
            switch "foo bar" {
            case "hello":
                if (2147483648 == v7) {
                    var v16 uint = 4294967295
                    v16 = 255
                    var v17 float32 = (-2.25 - float32(-2.25))
                    _ = v16
                    _ = v17
                } else {
                    var v18 float64 = (float64(0.1) - -2.25)
                    _ = v18
                }
            case "foo bar":
                v7 = (v7 / (((v7 * 2147483648)) | 1))
            case "":
                v19 := GenS2{F0: 18446744073709551615, F1: 255, F2: 1}
                _ = v19
            default:
                v20 := GenS1{F0: 2, F1: 1}
                _ = v20
            }
            _ = fi15
        }
        v21 := GenS3{F0: 2147483647}
        _ = v21
    } else {
        if false {
            switch "" {
            case "":
                fmt.Println(v7)
            case "\n\t":
                v22 := []int{42}
                _ = v22
            case "a":
                v23 := []int{-9223372036854775808, 9223372036854775807, -2147483648}
                _ = v23
            }
            v24 := GenS4{F0: -67, F1: 4294967295}
            _ = v24
        }
        var v25 uint32 = v7
        _ = v25
    }
    v26 := []int{1, 4294967295}
    _ = v26
    _ = v7
    return "hello"
}
func genFn1() int {
    for fi27 := 0; fi27 < 14; fi27++ {
        var v28 int64 = 42
        var v29 uint16 = 794
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        var v31 int16 = 1
        v32 := GenS3{F0: fi30}
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 uint = 2
    var v34 uint = (v33 * v33)
    if false {
        var v35 int32 = 42
        _ = v35
    }
    _ = v33
    _ = v34
    return -851
}

func main() {
    var v1 int = 32768
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    var v4 bool = (!false)
    var v5 uint64 = 2
    v6 := make(map[string]bool)
    var v36 int16 = 879
    if v4 {
        var v37 string = "foo bar"
        _ = v37
    } else {
        v38 := GenS5{F0: 2}
        if true {
            if v4 {
                v39 := GenS6{F0: 128}
                v5 -= v5
                _ = v39
            } else {
                fmt.Println(v1, v5)
                switch "" {
                case "foo bar":
                    var v40 float64 = 3.14159
                    _ = v40
                case "a":
                    v1 = (v1 << ((442) & 63))
                default:
                    if ((("\n\t" + "\n\t") <= ("" + "foo bar")) && (!(1e10 > 3.14159))) {
                        var v41 int64 = 52
                        _ = v41
                    }
                }
            }
            if (v4 || true) {
                v42 := []int{int(v5), 898, (^(v1))}
                _ = v42
            } else {
                var v43 uint = 42
                var v44 float32 = ((((-2.25 / (1e10 * 1e10)) / (0.0 + (-2.25 * 3.14159))) + (1e10 - (0.1 - -2.25))) * 3.14159)
                _ = v43
                _ = v44
            }
            fmt.Println(v5, v5, v1)
        }
        v4 = v4
        _ = v38
    }
    var v45 uint32 = 2
    v46 := GenS7{F0: v45, F1: 4294967295, F2: (v1 % ((v1) | 1))}
    var v47 int8 = -128
    for fi48 := 0; fi48 < 2; fi48++ {
        var v49 string = "hello"
        switch genFn1() {
        case 32768:
            for fi50 := 0; fi50 < 8; fi50++ {
                var v51 uint64 = v5
                var v52 uint = 255
                _ = fi50
                _ = v51
                _ = v52
            }
        default:
            for fi53 := 0; fi53 < 8; fi53++ {
                switch (v49 + "foo bar") {
                case "a":
                    fmt.Println(v4)
                case "hello":
                    var v54 int16 = (v36 & v36)
                    _ = v54
                case "\n\t":
                    fmt.Println(v47)
                }
                v49 = v49
                switch ((("\n\t" + "") + (v49 + (v49 + "hello"))) + v49) {
                case "\n\t":
                    var v55 uint64 = v5
                    _ = v55
                }
                _ = fi53
            }
        }
        fmt.Println(v1, v4)
        _ = fi48
        _ = v49
    }
    switch 42 {
    case -154:
        for fi56 := 0; fi56 < 13; fi56++ {
            fmt.Println(v47, v45)
            if (1 != 1) {
                var v57 int32 = 231
                v5 = v5
                if (0.0 < 0.0) {
                    v4 = true
                    var v58 int8 = v47
                    _ = v58
                }
                _ = v57
            }
            v59 := []int{}
            _ = fi56
            _ = v59
        }
    case 0:
        fmt.Println(v1, v47, v45)
    default:
        for fi60 := 0; fi60 < 16; fi60++ {
            var v61 uint = 535
            var v62 int32 = 42
            for fi63 := 0; fi63 < 6; fi63++ {
                fmt.Println(v61, v5)
                switch "" {
                case "foo bar":
                    var v64 float32 = 1e10
                    _ = v64
                default:
                    v65 := GenS5{F0: v45}
                    _ = v65
                }
                if (-2.25 == 0.0) {
                    var v66 bool = v4
                    fmt.Println(v5)
                    v67 := GenS8{F0: -9223372036854775808, F1: 1}
                    _ = v66
                    _ = v67
                } else {
                    v47 = v47
                    switch genFn1() {
                    case 128:
                        v68 := GenS9{F0: (fi63 - int(v61))}
                        _ = v68
                    case 607:
                        if true {
                            var v69 float64 = (1.5 / 1e10)
                            v1 = v1
                            var v70 uint = 451
                            _ = v69
                            _ = v70
                        }
                    case -1:
                        v71 := GenS10{F0: 2}
                        _ = v71
                    default:
                        v72 := GenS11{F0: int32(v5)}
                        _ = v72
                    }
                    switch "a" {
                    case "foo bar":
                        switch int(v62) {
                        case 653:
                            if v4 {
                                fmt.Println(v1, v1)
                            } else {
                                var v73 int16 = v36
                                v45 = v45
                                var v74 uint8 = 2
                                var v75 int64 = 351
                                _ = v73
                                _ = v74
                                _ = v75
                            }
                        }
                    case "\n\t":
                        var v76 float32 = 0.1
                        _ = v76
                    case "hello":
                        fmt.Println(fi60, fi60)
                    default:
                        v1 -= 32767
                    }
                }
                _ = fi63
            }
            _ = fi60
            _ = v61
            _ = v62
        }
    }
    v36 = v36
    var v77 uint8 = 0
    v36 = v36
    if true {
        var v78 uint16 = 42
        _ = v78
    } else {
        switch (v1 / (((v1 << ((255) & 63))) | 1)) {
        case 42:
            var v79 int8 = 87
            _ = v79
        case 32768:
            if false {
                switch (v1 | v1) {
                case 434:
                    _ = len(v6)
                default:
                    fmt.Println(v5, v77, v36)
                }
            }
        default:
            switch ("foo bar" + (("\n\t" + ("" + "foo bar")) + "\n\t")) {
            case "foo bar":
                v80 := GenS5{F0: (v45 % ((uint32(v77)) | 1))}
                _ = v80
            case "a":
                for fi81 := 0; fi81 < 7; fi81++ {
                    v5 += v5
                    if (3.14159 != float32(float64(-2.25))) {
                        fmt.Println(v45)
                        v82 := GenS12{F0: 747, F1: "\n\t", F2: (65535 < 18446744073709551615)}
                        var v83 float64 = (((float64((1.5 - float32(float64(1.5)))) + (float64(1.5) / 1e10)) - 3.14159) / 3.14159)
                        _ = v82
                        _ = v83
                    }
                    _ = fi81
                }
            default:
                v84 := GenS8{F0: 1, F1: 1}
                _ = v84
            }
        }
    }
    var v85 int = 127
    v86 := GenS13{F0: (v47 | -128), F1: 278}
    v45 += v45
    if (42 >= 2147483647) {
        if v4 {
            fmt.Println(v1)
            v6["\n\t"] = (42 >= 938)
            v45 = (v45 + v45)
        } else {
            v85 += v85
        }
        var v87 int = v1
        _ = v87
    }
    fmt.Println(v36)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v77:", v77)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v77
    _ = v85
    _ = v86
}
