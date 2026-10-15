package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS2 struct { F0 bool; F1 int64; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint8; F2 string }
type GenS5 struct { F0 uint64; F1 int8; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint }
func genFn0(p0 uint8, p1 int32, p2 uint32) uint {
    var v6 bool = false
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{-1, fi7, (fi7 - (fi7 / ((fi7) | 1)))}
        var v9 uint = 65535
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 295
}

func main() {
    var v1 bool = true
    var v2 uint32 = 272
    v3 := []string{"", "a"}
    var v4 int64 = 602
    var v5 uint8 = 247
    if v1 {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 string = ((("a" + "") + "") + (("hello" + "hello") + ("\n\t" + "\n\t")))
            _ = fi10
            _ = v11
        }
        if (0.1 == 1.5) {
            _ = len(v3)
            var v12 float64 = -2.25
            _ = v12
        }
        switch 128 {
        case 4294967295:
            switch "" {
            case "\n\t":
                fmt.Println(v1, v5, v4)
            case "foo bar":
                var v13 uint64 = 2
                _ = v13
            case "hello":
                if false {
                    v14 := GenS1{F0: v4, F1: (v5 >> ((genFn0(v5, -746, 2)) & 63)), F2: 2}
                    var v15 uint64 = 0
                    _ = v14
                    _ = v15
                } else {
                    v16 := GenS1{F0: (v4 - (v4 | v4)), F1: (v5 ^ v5), F2: v5}
                    fmt.Println(v2)
                    switch 51 {
                    case -408:
                        v4 -= (v4 - v4)
                    default:
                        v5 += v5
                    }
                    _ = v16
                }
            default:
                if false {
                    fmt.Println(v1)
                    v4 = v4
                    if (-698 <= -750) {
                        fmt.Println(v4)
                        var v17 float64 = 3.14159
                        for fi18 := 0; fi18 < 11; fi18++ {
                            var v19 string = ("hello" + ("" + "foo bar"))
                            v2 += 946
                            _ = fi18
                            _ = v19
                        }
                        _ = v17
                    } else {
                        v20 := GenS1{F0: (v4 % ((0) | 1)), F1: (v5 ^ (v5 * (v5 * v5))), F2: (^(v5))}
                        var v21 int64 = (v4 >> ((470) & 63))
                        fmt.Println(v21)
                        _ = v20
                        _ = v21
                    }
                }
            }
        case -823:
            v22 := GenS2{F0: (0 > 255), F1: 1, F2: 461}
            _ = v22
        case 2147483647:
            if ((v4 | v4) <= 9223372036854775807) {
                v23 := GenS1{F0: v4, F1: (^(v5)), F2: (v5 | (v5 ^ v5))}
                _ = v23
            } else {
                switch -1 {
                case -9223372036854775808:
                    if v1 {
                        for fi24 := 0; fi24 < 12; fi24++ {
                            if v1 {
                                var v25 float32 = (float32(float64(float32((float64(0.1) - (0.0 * (1.5 / (1.5 * ((0.0 / ((1.5 - float64(3.14159)) - 0.0)) + 1.5)))))))) / 1e10)
                                _ = v25
                            } else {
                                var v26 uint64 = uint64(v2)
                                _ = v26
                            }
                            v27 := GenS3{F0: v4}
                            _ = fi24
                            _ = v27
                        }
                        var v28 float64 = 3.14159
                        _ = v28
                    }
                case -906:
                    fmt.Println(v5, v1)
                case 0:
                    var v29 bool = (v1 && (697 <= 4294967295))
                    _ = v29
                default:
                    _ = len(v3)
                }
                var v30 string = "foo bar"
                switch v30 {
                case "":
                    v2 = (v2 >> ((genFn0((v5 | v5), 2147483647, 453)) & 63))
                case "a":
                    var v31 uint32 = 65535
                    _ = v31
                }
                _ = v30
            }
        default:
            if false {
                if (true || (false || (128 <= (v5 << ((2) & 63))))) {
                    v2 = 2147483648
                    fmt.Println(v1, v1, v5)
                    for fi32 := 0; fi32 < 11; fi32++ {
                        fmt.Println(v4, v4, v4)
                        _ = fi32
                    }
                }
                var v33 int32 = -452
                fmt.Println(v33, v1)
                _ = v33
            }
        }
    } else {
        var v34 float64 = 1e10
        switch -2147483648 {
        case -918:
            v35 := GenS4{F0: v4, F1: 255, F2: "hello"}
            _ = v35
        case -996:
            v36 := GenS5{F0: 2, F1: -128, F2: 1000}
            _ = v36
        default:
            v5 += v5
        }
        for fi37 := 0; fi37 < 19; fi37++ {
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            _ = len(v3)
            for fi38 := 0; fi38 < 19; fi38++ {
                v39 := GenS2{F0: (247 >= 18446744073709551615), F1: (v4 / ((v4) | 1)), F2: 238}
                switch (fi37 ^ 22) {
                case 2147483647:
                    var v40 int64 = v4
                    _ = v40
                case -943:
                    _ = len(v3)
                    v4 += v4
                    if v1 {
                        v34 -= (-2.25 - 1e10)
                    }
                case 128:
                    fmt.Println(v5)
                }
                _ = fi38
                _ = v39
            }
            _ = fi37
        }
        _ = v34
    }
    if v1 {
        v41 := GenS6{F0: 778}
        switch ((("a" + ("hello" + "")) + "\n\t") + "") {
        case "a":
            switch "a" {
            case "\n\t":
                v42 := GenS3{F0: v4}
                _ = v42
            default:
                v43 := GenS7{F0: v2}
                _ = v43
            }
        case "foo bar":
            v1 = v1
        case "hello":
            v5 += v5
        default:
            var v44 int32 = -1
            _ = v44
        }
        _ = v41
    }
    if v1 {
        v5 += v5
    } else {
        fmt.Println(v5, v4, v1)
        var v45 uint8 = 2
        _ = v45
    }
    var v46 string = ""
    switch 403 {
    case -1:
        var v47 int16 = 976
        _ = v47
    }
    v48 := GenS8{F0: (-(v4))}
    var v49 int8 = -128
    var v50 uint = 365
    switch -2147483648 {
    case 4294967295:
        v3 = append(v3, "a")
    case 32767:
        fmt.Println(v5, v5, v50)
    default:
        var v51 int16 = 1000
        _ = v51
    }
    switch 128 {
    case 4294967295:
        v52 := GenS8{F0: v4}
        _ = v52
    case -73:
        v4 += (v4 * v4)
    case -507:
        for fi53 := 0; fi53 < 14; fi53++ {
            switch fi53 {
            case 9223372036854775807:
                v5 -= 127
            case -871:
                v54 := GenS2{F0: (uint32(v5) > 1), F1: v4, F2: v50}
                _ = v54
            default:
                fmt.Println(fi53, fi53, v49)
            }
            if v1 {
                for fi55 := 0; fi55 < 3; fi55++ {
                    var v56 int64 = (v4 % (((v4 / ((v4) | 1))) | 1))
                    v5 -= 243
                    v2 += (v2 / ((v2) | 1))
                    _ = fi55
                    _ = v56
                }
                var v57 string = ""
                v58 := GenS9{F0: (v50 % ((genFn0(v5, 42, 1)) | 1))}
                _ = v57
                _ = v58
            }
            fmt.Println(v49, v4)
            _ = fi53
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
