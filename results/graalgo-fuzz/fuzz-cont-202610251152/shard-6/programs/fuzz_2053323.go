package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 int32 = 42
    var v2 uint64 = 1
    v3 := make(map[int]bool)
    var v4 float32 = (0.1 + -2.25)
    var v5 uint32 = 0
    v6 := []uint32{v5, v5, (v5 << ((1) & 63))}
    if (-9223372036854775808 >= 9223372036854775807) {
        switch 127 {
        case 42:
            v7 := GenS1{F0: -1}
            _ = v7
        default:
            var v8 int64 = -473
            _ = v8
        }
    } else {
        switch 2147483647 {
        case -2147483648:
            var v9 int64 = -1
            _ = v9
        case 0:
            if (!(false && false)) {
                for fi10 := 0; fi10 < 16; fi10++ {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        var v12 int32 = -1
                        _ = fi11
                        _ = v12
                    }
                    v13 := GenS1{F0: 0}
                    _ = len(v6)
                    _ = fi10
                    _ = v13
                }
            } else {
                for fi14 := 0; fi14 < 7; fi14++ {
                    switch "foo bar" {
                    case "a":
                        v4 = v4
                    case "foo bar":
                        var v15 uint32 = v5
                        _ = v15
                    default:
                        v5 -= 931
                    }
                    for fi16 := 0; fi16 < 16; fi16++ {
                        v5 += v5
                        var v17 int32 = v1
                        if ((!true) || true) {
                            switch "hello" {
                            case "\n\t":
                                switch (fi16 | fi14) {
                                case 32768:
                                    var v18 uint = 2
                                    _ = v18
                                case -493:
                                    if true {
                                        var v19 float32 = -2.25
                                        var v20 int = -9223372036854775808
                                        fmt.Println(v5, fi16, v4)
                                        _ = v19
                                        _ = v20
                                    }
                                case 327:
                                    switch "\n\t" {
                                    case "":
                                        var v21 int8 = -128
                                        _ = v21
                                    }
                                }
                            case "foo bar":
                                if false {
                                    switch fi14 {
                                    case 9223372036854775807:
                                        var v22 uint16 = 570
                                        _ = v22
                                    case 565:
                                        v17 += 2147483647
                                    case 0:
                                        v4 = v4
                                    }
                                } else {
                                    var v23 int64 = 42
                                    _ = v23
                                }
                            case "a":
                                v24 := GenS2{F0: fi14}
                                _ = v24
                            }
                            v5 = 1000
                            var v25 uint = 418
                            if (22 <= 851) {
                                var v26 uint16 = 847
                                _ = v26
                            }
                            var v27 bool = (int64(v2) != 9223372036854775807)
                            var v28 uint8 = 231
                            _ = v25
                            _ = v27
                            _ = v28
                        }
                        _ = fi16
                        _ = v17
                    }
                    _ = fi14
                }
                v29 := GenS2{F0: 1}
                _ = v29
            }
        case -1:
            switch 32767 {
            case 497:
                if (v1 != v1) {
                    var v30 int8 = 127
                    v31 := GenS3{F0: (-(v30))}
                    if true {
                        v2 = 18446744073709551615
                        for fi32 := 0; fi32 < 12; fi32++ {
                            fmt.Println(v4, v30)
                            _ = fi32
                        }
                    } else {
                        v4 -= 3.14159
                    }
                    _ = v30
                    _ = v31
                } else {
                    v33 := GenS1{F0: 42}
                    fmt.Println(v5, v1)
                    var v34 bool = true
                    _ = v33
                    _ = v34
                }
            }
        }
        v35 := GenS3{F0: -23}
        var v36 uint8 = 127
        _ = v35
        _ = v36
    }
    switch "hello" {
    case "foo bar":
        var v37 uint = 2
        _ = v37
    case "hello":
        switch 42 {
        case 4294967295:
            var v38 uint32 = 69
            _ = v38
        case -9223372036854775808:
            v5 = 2147483647
        case 266:
            fmt.Println(v2, v4)
        }
    case "a":
        for fi39 := 0; fi39 < 1; fi39++ {
            v4 += v4
            _ = fi39
        }
    default:
        if len(v6) > 0 {
            v6[0] = (^(v5))
        }
    }
    fmt.Println(v2, v2, v4)
    v40 := v3[548]
    if ((((v4 - 0.1) - (3.14159 + 1.5)) < v4) && v40) {
        for fi41 := 0; fi41 < 17; fi41++ {
            v42 := GenS2{F0: fi41}
            var v43 uint16 = 0
            switch fi41 {
            case 32768:
                for fi44 := 0; fi44 < 10; fi44++ {
                    if true {
                        fmt.Println(v1, v2, fi41)
                    } else {
                        var v45 float64 = (0.0 * 0.0)
                        v2 -= v2
                        _ = v45
                    }
                    v46 := GenS2{F0: 32767}
                    var v47 int32 = v1
                    _ = fi44
                    _ = v46
                    _ = v47
                }
            default:
                if (!((!(v1 <= (v1 << ((18446744073709551615) & 63)))) && false)) {
                    v4 += (float32(1e10) * float32((3.14159 + float64(((v4 / v4) + v4)))))
                    v48 := GenS4{F0: 240}
                    if (!v40) {
                        v1 += (v1 / ((v1) | 1))
                        v49 := GenS5{F0: 65535, F1: 306, F2: v2}
                        var v50 int8 = int8(v1)
                        _ = v49
                        _ = v50
                    }
                    _ = v48
                } else {
                    v51 := v3[105]
                    _ = v51
                }
            }
            _ = fi41
            _ = v42
            _ = v43
        }
        v2 = 2
    } else {
        v2 = 18446744073709551615
        fmt.Println(v1, v1, v5)
        v1 += (v1 | (^(v1)))
    }
    v52 := GenS6{F0: 255}
    if (180 <= 32767) {
        for fi53 := 0; fi53 < 13; fi53++ {
            switch ("" + "foo bar") {
            case "":
                v2 += (^(v2))
            case "foo bar":
                var v54 uint16 = 107
                _ = v54
            }
            _ = fi53
        }
        fmt.Println(v5, v40)
        v55 := GenS1{F0: 42}
        _ = v55
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi56 := 0; fi56 < 9; fi56++ {
        fmt.Println(v40, v4, v5)
        _ = fi56
    }
    if ("" < "") {
        v57 := GenS2{F0: 32767}
        _ = v57
    } else {
        v58 := GenS6{F0: 127}
        _ = v58
    }
    v59 := GenS7{F0: 844, F1: (^(v2))}
    fmt.Println(v1, v4, v1)
    var v60 uint32 = v5
    v61 := v3[-929]
    v5 -= (^(v60))
    var v62 uint16 = 2
    v40 = ((!(0 <= 42)) && v40)
    var v63 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v52
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
