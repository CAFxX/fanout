package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32 }

func main() {
    var v1 int8 = -76
    var v2 uint16 = 1
    v3 := make(map[uint32]string)
    var v4 string = (("" + "foo bar") + "foo bar")
    var v5 float64 = 0.0
    switch 32768 {
    case 128:
        var v6 int = -1
        _ = v6
    case -247:
        var v7 int8 = 42
        _ = v7
    case 610:
        for fi8 := 0; fi8 < 3; fi8++ {
            if ((32767 >= 42) && (458 <= 928)) {
                var v9 uint32 = 65535
                v10 := GenS1{F0: (v9 | v9), F1: 42}
                if (!((2 >= 42) || (-1 <= -2147483648))) {
                    delete(v3, v9)
                }
                _ = v10
                _ = v9
            }
            var v11 float64 = 1.5
            v12 := GenS2{F0: fi8, F1: 67}
            _ = fi8
            _ = v11
            _ = v12
        }
    default:
        v5 = (float64(1e10) - ((1.5 + v5) - v5))
    }
    v13 := GenS1{F0: 608, F1: 2}
    if false {
        switch 4294967295 {
        case 42:
            switch 32767 {
            case -247:
                fmt.Println(v1)
            case -2147483648:
                var v14 int32 = 42
                _ = v14
            case 128:
                v4 = v4
            default:
                switch 947 {
                case 1:
                    for fi15 := 0; fi15 < 3; fi15++ {
                        v16 := GenS2{F0: fi15, F1: 566}
                        if (fi15 == -2147483648) {
                            var v17 uint = 90
                            for fi18 := 0; fi18 < 13; fi18++ {
                                delete(v3, 65535)
                                var v19 int64 = -999
                                v20 := []int{(fi18 - 1)}
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            _ = v17
                        } else {
                            if (2 >= 401) {
                                var v21 int64 = -9223372036854775808
                                if (-1 < -414) {
                                    var v22 int = fi15
                                    v4 = (v4 + "\n\t")
                                    _ = v22
                                }
                                _ = v21
                            } else {
                                for fi23 := 0; fi23 < 7; fi23++ {
                                    v1 = v1
                                    switch fi23 {
                                    case 42:
                                        if (0 == -243) {
                                            if true {
                                                var v24 uint32 = 65535
                                                v24 += v24
                                                if true {
                                                    fmt.Println(v5)
                                                    if false {
                                                        v25 := []int{fi15, fi23, (fi15 | fi15)}
                                                        var v26 int64 = -133
                                                        v1 -= v1
                                                        _ = v25
                                                        _ = v26
                                                    }
                                                } else {
                                                    if true {
                                                        fmt.Println(v1)
                                                        v27 := GenS1{F0: v24, F1: 307}
                                                        var v28 float64 = 0.1
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                }
                                                _ = v24
                                            }
                                            if true {
                                                fmt.Println(fi15, v4, fi15)
                                                fmt.Println(fi23, v1, v5)
                                                delete(v3, 790)
                                                fmt.Println(v5, v1)
                                            }
                                            switch 1 {
                                            case -783:
                                                var v29 uint32 = 4294967295
                                                v4 = v4
                                                v30 := GenS1{F0: 921, F1: 0}
                                                _ = v29
                                                _ = v30
                                            default:
                                                switch 1 {
                                                case 4294967295:
                                                    var v31 int64 = -9223372036854775808
                                                    _ = v31
                                                case 0:
                                                    if false {
                                                        var v32 int64 = 0
                                                        v33 := GenS3{F0: (v32 + (v32 / (((-(v32))) | 1))), F1: v32}
                                                        v32 += v32
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                case -255:
                                                    v34 := GenS4{F0: 2147483648}
                                                    _ = v34
                                                }
                                            }
                                        }
                                    case -2147483648:
                                        v4 = v4
                                    default:
                                        delete(v3, 65535)
                                    }
                                    _ = fi23
                                }
                                v4 = "hello"
                                for fi35 := 0; fi35 < 8; fi35++ {
                                    v2 += v2
                                    _ = fi35
                                }
                            }
                            var v36 float64 = v5
                            var v37 bool = true
                            _ = v36
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 6; fi38++ {
                            if (256 >= 0) {
                                v39 := GenS5{F0: -1, F1: v1}
                                fmt.Println(fi15, v2)
                                v40 := []int{32768, fi15, (fi38 / ((fi38) | 1))}
                                _ = v39
                                _ = v40
                            } else {
                                switch (fi15 - fi38) {
                                case 32768:
                                    if true {
                                        v5 -= (-2.25 - v5)
                                    }
                                }
                                v41 := GenS3{F0: int64(v2), F1: 9223372036854775807}
                                _ = v41
                            }
                            _ = fi38
                        }
                        _ = fi15
                        _ = v16
                    }
                }
            }
        default:
            var v42 int32 = 1
            _ = v42
        }
        v43 := GenS5{F0: -995, F1: (-(v1))}
        _ = v43
    }
    v44 := GenS1{F0: 0, F1: 2}
    if true {
        var v45 int32 = -659
        fmt.Println(v4, v2, v4)
        fmt.Println(v2, v5, v4)
        _ = v45
    }
    var v46 uint8 = uint8(v2)
    var v47 int = 32767
    var v48 bool = false
    for fi49 := 0; fi49 < 18; fi49++ {
        for fi50 := 0; fi50 < 14; fi50++ {
            v2 += 1
            if true {
                v3[1] = v4
            }
            _ = fi50
        }
        _ = fi49
    }
    var v51 float32 = -2.25
    if v48 {
        fmt.Println(v4, v51, v4)
        for fi52 := 0; fi52 < 20; fi52++ {
            for fi53 := 0; fi53 < 13; fi53++ {
                v48 = v48
                v1 = 100
                _ = fi53
            }
            _ = fi52
        }
        switch (v47 >> ((786) & 63)) {
        case 32767:
            for fi54 := 0; fi54 < 6; fi54++ {
                var v55 int64 = 1
                v56 := []int{(v47 + -629)}
                v4 = v4
                _ = fi54
                _ = v55
                _ = v56
            }
        case 1:
            fmt.Println(v2, v4, v51)
        }
    }
    var v57 int = -2147483648
    v4 = v4
    var v58 bool = false
    switch 1 {
    case 9223372036854775807:
        var v59 uint64 = 42
        _ = v59
    case -1:
        v60 := GenS6{F0: ("\n\t" + v4)}
        _ = v60
    default:
        switch "foo bar" {
        case "\n\t":
            for fi61 := 0; fi61 < 18; fi61++ {
                var v62 int = -437
                for fi63 := 0; fi63 < 6; fi63++ {
                    v3[2147483647] = "foo bar"
                    _ = fi63
                }
                var v64 uint32 = 2147483648
                _ = fi61
                _ = v62
                _ = v64
            }
        case "hello":
            var v65 uint8 = v46
            _ = v65
        }
    }
    for fi66 := 0; fi66 < 12; fi66++ {
        v67 := GenS7{F0: 0}
        var v68 uint16 = v2
        _ = len(v3)
        _ = fi66
        _ = v67
        _ = v68
    }
    switch (("\n\t" + v4) + "hello") {
    case "a":
        v69 := v3[65535]
        _ = v69
    default:
        switch 32768 {
        case 127:
            fmt.Println(v51)
        }
    }
    var v70 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v51
    _ = v57
    _ = v58
    _ = v70
}
