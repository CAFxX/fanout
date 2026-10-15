package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int }
type GenS2 struct { F0 uint8; F1 int64; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 uint16; F1 string; F2 bool }
type GenS8 struct { F0 uint16; F1 uint; F2 uint64 }

func main() {
    var v1 uint64 = 42
    var v2 int8 = -1
    v3 := make(map[int]int)
    var v4 int64 = -1
    var v5 uint16 = 1
    var v6 uint64 = v1
    var v7 uint32 = 42
    var v8 uint64 = 387
    switch 897 {
    case 785:
        if (!(true && (true || (127 >= 108)))) {
            v7 += 52
            var v9 float32 = -2.25
            _ = v9
        }
    case -699:
        switch ("" + "hello") {
        case "\n\t":
            v6 += v1
        case "foo bar":
            v10 := GenS1{F0: -219, F1: 42, F2: 127}
            _ = v10
        case "hello":
            v11 := GenS2{F0: 47, F1: (v4 ^ (v4 % ((9223372036854775807) | 1))), F2: true}
            _ = v11
        default:
            v7 = 213
        }
    case 0:
        fmt.Println(v4, v5, v7)
    default:
        var v12 uint8 = uint8(v7)
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v8 = 0
        _ = fi13
    }
    var v14 int8 = int8(v7)
    for fi15 := 0; fi15 < 20; fi15++ {
        if true {
            fmt.Println(fi15)
        } else {
            for fi16 := 0; fi16 < 9; fi16++ {
                fmt.Println(v5, fi16, v8)
                v5 -= 2
                fmt.Println(v6, fi15)
                _ = fi16
            }
        }
        _ = fi15
    }
    v3[2147483647] = 9223372036854775807
    var v17 int32 = int32(v8)
    if ((0.1 > 1e10) && (float32(1.5) <= float32((0.0 + (1e10 - 1.5))))) {
        switch "" {
        case "":
            if (!true) {
                v17 = (v17 ^ v17)
                if false {
                    var v18 int8 = v2
                    if false {
                        fmt.Println(v18)
                    }
                    _ = v18
                } else {
                    var v19 int16 = -999
                    switch 42 {
                    case -2147483648:
                        var v20 float32 = 1e10
                        _ = v20
                    case 287:
                        for fi21 := 0; fi21 < 7; fi21++ {
                            v17 = (v17 / ((v17) | 1))
                            if false {
                                for fi22 := 0; fi22 < 10; fi22++ {
                                    var v23 float64 = 0.1
                                    var v24 int8 = (-(v14))
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                                if false {
                                    fmt.Println(v5)
                                } else {
                                    var v25 float32 = float32((0.0 + float64(0.1)))
                                    v26 := []int{(fi21 | 0)}
                                    v27 := GenS3{F0: (^(fi21))}
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                            }
                            _ = fi21
                        }
                    }
                    _ = v19
                }
                if true {
                    fmt.Println(v2, v5)
                    v17 = v17
                } else {
                    v2 += 100
                }
            }
        }
    }
    v28 := []int{274, 4294967295, 42}
    if false {
        for fi29 := 0; fi29 < 10; fi29++ {
            if ((3.14159 + 3.14159) >= 1.5) {
                var v30 uint16 = (v5 % ((1) | 1))
                v31 := v3[(fi29 << ((1) & 63))]
                _ = v30
                _ = v31
            } else {
                var v32 uint8 = 255
                fmt.Println(v2, fi29)
                _ = v32
            }
            _ = len(v28)
            v33 := GenS4{F0: 0, F1: 927}
            _ = fi29
            _ = v33
        }
        var v34 int = 1
        _ = v34
    }
    if ((9223372036854775807 != -414) || false) {
        var v35 float64 = (-2.25 / -2.25)
        for fi36 := 0; fi36 < 5; fi36++ {
            for fi37 := 0; fi37 < 1; fi37++ {
                if (!true) {
                    var v38 int = -9223372036854775808
                    _ = len(v3)
                    _ = v38
                } else {
                    v3[1] = 882
                    switch "hello" {
                    case "\n\t":
                        switch (fi37 * -476) {
                        case -2147483648:
                            v39 := GenS3{F0: 9223372036854775807}
                            _ = v39
                        case 32768:
                            var v40 float64 = v35
                            _ = v40
                        }
                    case "hello":
                        var v41 uint32 = (^(v7))
                        _ = v41
                    case "foo bar":
                        v42 := GenS2{F0: 0, F1: -9223372036854775808, F2: ((v14 >= v14) || true)}
                        _ = v42
                    default:
                        var v43 string = ""
                        _ = v43
                    }
                    var v44 bool = ("a" < "")
                    _ = v44
                }
                if (int(v7) != (^(fi37))) {
                    var v45 int32 = (v17 % ((v17) | 1))
                    var v46 uint8 = 255
                    _ = v45
                    _ = v46
                } else {
                    v47 := GenS5{F0: 717, F1: "hello"}
                    var v48 float64 = -2.25
                    _ = v47
                    _ = v48
                }
                _ = fi37
            }
            switch "\n\t" {
            case "hello":
                v14 += 127
            case "foo bar":
                var v49 int64 = (-(v4))
                _ = v49
            }
            fmt.Println(fi36, v17, v5)
            _ = fi36
        }
        switch ("\n\t" + "foo bar") {
        case "a":
            if len(v28) >= 2 {
                v28 = v28[1:len(v28)]
            }
        case "foo bar":
            if (4294967295 >= v8) {
                delete(v3, -9223372036854775808)
            } else {
                if true {
                    if (0.0 >= ((1.5 / ((v35 * ((float64(((0.0 + 0.0) / -2.25)) - (1e10 * (((v35 / float64(((3.14159 * 3.14159) / (((1.5 / 1.0) - ((0.0 * (0.0 * ((((-2.25 / float32(v35)) * -2.25) - 3.14159) / 1e10))) / 0.1)) - (3.14159 + 0.1))))) + (v35 / (float64(1e10) - ((float64((0.1 * (((1.5 * ((((1.5 * (3.14159 / (1e10 / 3.14159))) / 1.0) + 0.1) + (((1.5 / 1.5) * (0.0 * (1e10 + (1.5 - 0.0)))) + (((-2.25 / (1.5 / 1e10)) * float32((v35 - 0.0))) - 3.14159)))) / 1.0) - 0.1))) - v35) - float64(-2.25))))) + v35))) / (v35 - v35))) - v35)) * v35)) {
                        switch "foo bar" {
                        case "hello":
                            v50 := GenS6{F0: (v17 | -409), F1: (^(v5))}
                            _ = v50
                        case "foo bar":
                            var v51 int = 187
                            _ = v51
                        case "\n\t":
                            v2 += 0
                        default:
                            switch "foo bar" {
                            case "foo bar":
                                v6 = (v6 ^ 78)
                            case "\n\t":
                                switch "a" {
                                case "a":
                                    switch -426 {
                                    case 128:
                                        if false {
                                            fmt.Println(v35, v1, v17)
                                            var v52 bool = false
                                            _ = v52
                                        }
                                    case -2147483648:
                                        if false {
                                            fmt.Println(v14, v5, v1)
                                        } else {
                                            for fi53 := 0; fi53 < 17; fi53++ {
                                                if len(v28) >= 2 {
                                                    v28 = v28[1:len(v28)]
                                                }
                                                var v54 uint16 = v5
                                                _ = fi53
                                                _ = v54
                                            }
                                            v1 = 57
                                            var v55 uint32 = 4294967295
                                            _ = v55
                                        }
                                    case -9223372036854775808:
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                    default:
                                        var v56 uint32 = v7
                                        _ = v56
                                    }
                                case "":
                                    var v57 float64 = float64(0.0)
                                    _ = v57
                                }
                            case "hello":
                                for fi58 := 0; fi58 < 15; fi58++ {
                                    if false {
                                        v5 += v5
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                        v59 := GenS7{F0: v5, F1: ("\n\t" + "a"), F2: (!false)}
                                        _ = v59
                                    }
                                    _ = fi58
                                }
                            default:
                                var v60 int16 = -1
                                _ = v60
                            }
                        }
                        if (false || (32767 < 0)) {
                            v17 = (v17 - 0)
                            if true {
                                if (!true) {
                                    v61 := GenS8{F0: (v5 << ((4294967295) & 63)), F1: 42, F2: (v8 / ((500) | 1))}
                                    var v62 bool = (!true)
                                    _ = v61
                                    _ = v62
                                }
                            } else {
                                v7 = 224
                            }
                        }
                    }
                    if (v17 < v17) {
                        for fi63 := 0; fi63 < 4; fi63++ {
                            for fi64 := 0; fi64 < 17; fi64++ {
                                if (!("a" >= (("hello" + "foo bar") + "foo bar"))) {
                                    var v65 int64 = v4
                                    _ = v65
                                } else {
                                    v66 := GenS2{F0: 2, F1: (v4 ^ (v4 - int64(v7))), F2: false}
                                    _ = len(v28)
                                    _ = v66
                                }
                                v8 += v8
                                var v67 float32 = (float32(((float64((0.1 * float32(v35))) * (v35 / ((3.14159 * v35) / float64((3.14159 * float32(v35)))))) + ((v35 * v35) - 1e10))) + 1.5)
                                _ = fi64
                                _ = v67
                            }
                            _ = fi63
                        }
                        v4 -= 615
                    } else {
                        var v68 bool = (255 >= 128)
                        _ = v68
                    }
                }
            }
        }
        _ = v35
    } else {
        delete(v3, -2147483648)
        fmt.Println(v4, v17, v8)
    }
    for fi69 := 0; fi69 < 5; fi69++ {
        v14 = v2
        _ = fi69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
