package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 bool; F2 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 int16 = 936
    var v4 uint = 18446744073709551615
    var v5 uint16 = 1
    v1["a"] = 1
    if (float64(3.14159) < -2.25) {
        if true {
            fmt.Println(v3, v5)
            v4 -= (^(v4))
            v6 := GenS1{F0: -128, F1: (2147483648 > 298)}
            _ = v6
        }
        if (("foo bar" + ("foo bar" + ("a" + ("foo bar" + (("foo bar" + "hello") + "hello"))))) > "a") {
            fmt.Println(v4, v4, v3)
            var v7 int16 = v3
            v8 := v1["hello"]
            _ = v7
            _ = v8
        } else {
            v5 += (v5 * 42)
        }
    } else {
        var v9 uint = (^(v4))
        fmt.Println(v5)
        _ = v9
    }
    if (-66 > 984) {
        var v10 float32 = (1e10 / 1e10)
        for fi11 := 0; fi11 < 17; fi11++ {
            if (!false) {
                fmt.Println(v3, v4)
                var v12 bool = (1 < int(v5))
                for fi13 := 0; fi13 < 19; fi13++ {
                    if v12 {
                        v14 := GenS2{F0: 1, F1: 0, F2: 151}
                        fmt.Println(v4, v12, v4)
                        _ = v14
                    } else {
                        switch -157 {
                        case 32767:
                            v15 := GenS3{F0: v5, F1: v4}
                            _ = v15
                        case -2147483648:
                            if false {
                                if (!(false && v12)) {
                                    v3 = v3
                                    v3 = 1
                                }
                            }
                        case 127:
                            v4 += (v4 * v4)
                        }
                        switch fi11 {
                        case 2147483647:
                            switch (fi13 >> ((628) & 63)) {
                            case 9223372036854775807:
                                _ = len(v2)
                            case 2147483647:
                                var v16 uint16 = (v5 / ((255) | 1))
                                _ = v16
                            }
                        case 0:
                            if v12 {
                                v10 = -2.25
                                fmt.Println(v5, v12, v4)
                                var v17 int = 128
                                v5 = v5
                                _ = v17
                            }
                        case -9223372036854775808:
                            v12 = (!(127 == 2))
                        default:
                            v18 := v2[(fi13 + fi13)]
                            _ = v18
                        }
                    }
                    fmt.Println(fi11, v10)
                    switch "\n\t" {
                    case "a":
                        v3 = (v3 / ((1000) | 1))
                    case "foo bar":
                        delete(v2, fi13)
                    case "\n\t":
                        fmt.Println(fi13, v10)
                    default:
                        fmt.Println(fi13, fi13)
                    }
                    _ = fi13
                }
                _ = v12
            } else {
                if (100 >= -1) {
                    if (!true) {
                        fmt.Println(v3)
                        switch int(v3) {
                        case 0:
                            v2[fi11] = fi11
                        default:
                            v19 := GenS4{F0: fi11}
                            _ = v19
                        }
                        v20 := GenS5{F0: -2147483648, F1: (100 <= 100), F2: "foo bar"}
                        _ = v20
                    }
                    v21 := GenS6{F0: 42}
                    _ = v21
                } else {
                    var v22 float32 = float32((0.1 * (-2.25 / 0.1)))
                    fmt.Println(v5, fi11, fi11)
                    fmt.Println(v3, fi11, v3)
                    _ = v22
                }
                fmt.Println(v5, v5)
            }
            v23 := []int{0, (fi11 / ((fi11) | 1)), 1}
            v24 := GenS7{F0: 127}
            _ = fi11
            _ = v23
            _ = v24
        }
        _ = v10
    }
    v3 -= v3
    for fi25 := 0; fi25 < 4; fi25++ {
        var v26 int16 = v3
        _ = fi25
        _ = v26
    }
    switch -977 {
    case 0:
        switch 127 {
        case 0:
            if (0 < 1) {
                var v27 bool = (!true)
                v28 := GenS1{F0: 127, F1: false}
                _ = v27
                _ = v28
            } else {
                if true {
                    v29 := []int{0, 2147483647}
                    fmt.Println(v5, v5, v5)
                    fmt.Println(v3, v5)
                    _ = v29
                }
            }
        case -435:
            var v30 uint32 = 691
            _ = v30
        default:
            if false {
                v4 += v4
            }
        }
    case 128:
        fmt.Println(v5, v3, v3)
    case -471:
        var v31 float32 = -2.25
        _ = v31
    }
    var v32 uint16 = (v5 | 2)
    v5 -= 581
    var v33 uint32 = 126
    if false {
        switch int(v33) {
        case -2147483648:
            v34 := []int{9223372036854775807, 823}
            _ = v34
        case 561:
            var v35 string = "foo bar"
            _ = v35
        case 127:
            switch 413 {
            case -144:
                if true {
                    var v36 int = 540
                    var v37 uint8 = 42
                    _ = v36
                    _ = v37
                }
            case 148:
                if true {
                    var v38 bool = (true && (false && (true || true)))
                    _ = v38
                }
            case 32768:
                v39 := []int{-384}
                _ = v39
            default:
                var v40 string = "hello"
                _ = v40
            }
        }
        var v41 int32 = -205
        if true {
            v3 -= (v3 + v3)
            v42 := GenS7{F0: 127}
            for fi43 := 0; fi43 < 19; fi43++ {
                switch "a" {
                case "foo bar":
                    v44 := v1["a"]
                    _ = v44
                case "":
                    var v45 bool = (1.5 > ((0.0 - ((1e10 + float64(0.1)) / 1.0)) * 3.14159))
                    _ = v45
                case "a":
                    v46 := GenS7{F0: 127}
                    _ = v46
                default:
                    for fi47 := 0; fi47 < 12; fi47++ {
                        var v48 float32 = 3.14159
                        _ = fi47
                        _ = v48
                    }
                }
                if true {
                    var v49 int16 = v3
                    _ = v49
                } else {
                    var v50 bool = true
                    v41 -= (^(v41))
                    _ = v50
                }
                _ = fi43
            }
            _ = v42
        } else {
            for fi51 := 0; fi51 < 11; fi51++ {
                v52 := GenS6{F0: 0}
                v41 -= v41
                for fi53 := 0; fi53 < 1; fi53++ {
                    _ = len(v2)
                    switch "a" {
                    case "foo bar":
                        v54 := []int{fi53}
                        _ = v54
                    default:
                        v55 := GenS2{F0: 0, F1: 2, F2: 65535}
                        var v56 bool = false
                        _ = v55
                        _ = v56
                    }
                    _ = fi53
                }
                _ = fi51
                _ = v52
            }
            var v57 int8 = -1
            _ = v57
        }
        _ = v41
    }
    var v58 uint16 = v5
    var v59 float32 = 0.1
    v60 := GenS8{F0: v4}
    var v61 uint = (^(v4))
    for fi62 := 0; fi62 < 17; fi62++ {
        if false {
            for fi63 := 0; fi63 < 8; fi63++ {
                _ = len(v2)
                var v64 uint8 = 128
                if ((v4 + v4) <= (v61 << ((uint(v33)) & 63))) {
                    switch (fi63 & fi62) {
                    case 847:
                        v65 := GenS6{F0: 359}
                        _ = v65
                    }
                    v66 := GenS2{F0: v3, F1: 4294967295, F2: (v33 + v33)}
                    fmt.Println(v3)
                    v64 += (v64 / ((255) | 1))
                    _ = v66
                } else {
                    var v67 uint64 = 1
                    _ = v67
                }
                _ = fi63
                _ = v64
            }
            v33 += 1
        } else {
            v3 -= (v3 & 42)
        }
        _ = fi62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
