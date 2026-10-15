package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint = (^(v1))
    var v3 int8 = -128
    var v4 int16 = 1000
    var v5 int64 = -867
    v4 -= v4
    var v6 float32 = 1.5
    v7 := []int{0, 0, int(v1), 0}
    v8 := GenS1{F0: 42}
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS1{F0: 0}
        v5 = 276
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        switch "a" {
        case "foo bar":
            if len(v7) > 0 {
                v7[0] = (fi11 ^ (fi11 << ((v2) & 63)))
            }
        case "a":
            v2 -= v2
        default:
            var v12 int64 = -897
            _ = v12
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint16 = 331
        _ = len(v7)
        for fi15 := 0; fi15 < 4; fi15++ {
            if false {
                v14 += v14
                switch "a" {
                case "\n\t":
                    fmt.Println(fi15, v6)
                case "":
                    switch "hello" {
                    case "foo bar":
                        v5 -= 1
                    case "hello":
                        if false {
                            var v16 uint8 = 39
                            if (v4 <= int16(v3)) {
                                var v17 uint32 = 2
                                _ = v17
                            } else {
                                v2 = v1
                                switch 0 {
                                case -731:
                                    v6 = ((0.1 / ((1.5 / v6) + 0.0)) - (((float32((0.0 / -2.25)) + (v6 + 0.1)) * 1.5) / v6))
                                case 9223372036854775807:
                                    v18 := GenS1{F0: 42}
                                    _ = v18
                                default:
                                    var v19 uint64 = 42
                                    _ = v19
                                }
                            }
                            _ = v16
                        }
                    case "\n\t":
                        switch (^(fi15)) {
                        case 4294967295:
                            v1 += v2
                        case 42:
                            fmt.Println(v3, fi15)
                        default:
                            var v20 uint8 = 219
                            _ = v20
                        }
                    default:
                        if true {
                            if (v2 < 100) {
                                switch (-(fi13)) {
                                case -2147483648:
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v21 string = "hello"
                                        switch fi15 {
                                        case 32768:
                                            v22 := GenS2{F0: (v4 == (v4 % ((1000) | 1))), F1: 1}
                                            _ = v22
                                        case -9223372036854775808:
                                            var v23 uint32 = 2147483648
                                            _ = v23
                                        }
                                        var v24 int32 = 2147483647
                                        var v25 int32 = -1
                                        var v26 int32 = 0
                                        _ = v21
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    case "hello":
                                        v7 = append(v7, fi15)
                                    case "a":
                                        var v27 int16 = 42
                                        _ = v27
                                    }
                                case 0:
                                    var v28 bool = false
                                    _ = v28
                                }
                                v29 := GenS1{F0: 42}
                                v6 = 1.5
                                _ = v29
                            }
                            switch fi15 {
                            case 32768:
                                var v30 int16 = (-(v4))
                                _ = v30
                            case 0:
                                var v31 float32 = (((v6 - v6) - v6) - (v6 / 3.14159))
                                _ = v31
                            case 1:
                                var v32 uint16 = (v14 + 2)
                                _ = v32
                            default:
                                switch 128 {
                                case -9223372036854775808:
                                    if true {
                                        var v33 float32 = (v6 - v6)
                                        fmt.Println(v14, fi15, v2)
                                        var v34 string = "hello"
                                        switch v34 {
                                        case "\n\t":
                                            switch 127 {
                                            case 32767:
                                                v2 += (^(v2))
                                            case -211:
                                                if ("hello" >= v34) {
                                                    if true {
                                                        fmt.Println(fi15, v3, v14)
                                                    }
                                                    var v35 uint16 = (v14 / ((32767) | 1))
                                                    _ = v35
                                                }
                                            default:
                                                v36 := GenS1{F0: 18446744073709551615}
                                                _ = v36
                                            }
                                        default:
                                            v37 := GenS3{F0: false, F1: 42}
                                            _ = v37
                                        }
                                        _ = v33
                                        _ = v34
                                    }
                                default:
                                    var v38 int32 = 2147483647
                                    _ = v38
                                }
                            }
                        }
                    }
                case "a":
                    v1 -= 368
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v39 int32 = 2147483647
                _ = v39
            } else {
                var v40 uint32 = 917
                v2 = v1
                var v41 bool = ((v1 << (((v2 & v2)) & 63)) <= v1)
                _ = v40
                _ = v41
            }
            v5 += v5
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    v7 = append(v7, 9223372036854775807)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    switch (("" + ("\n\t" + "\n\t")) + "foo bar") {
    case "foo bar":
        var v42 float64 = 3.14159
        _ = v42
    }
    if ((0.0 != (0.1 * 1e10)) || (120 < 46)) {
        v3 += v3
        if true {
            if true {
                var v43 int64 = v5
                v44 := GenS1{F0: 870}
                _ = v43
                _ = v44
            }
            v45 := GenS4{F0: 0}
            _ = v45
        } else {
            var v46 int64 = v5
            fmt.Println(v6)
            v47 := GenS5{F0: 0}
            _ = v46
            _ = v47
        }
        if false {
            fmt.Println(v1)
        } else {
            for fi48 := 0; fi48 < 10; fi48++ {
                switch -984 {
                case 127:
                    if (true || (!(0.0 > (0.1 / v6)))) {
                        switch ("a" + (((((("a" + "foo bar") + ("" + "\n\t")) + "foo bar") + (("foo bar" + "") + "\n\t")) + "hello") + ("\n\t" + "\n\t"))) {
                        case "a":
                            for fi49 := 0; fi49 < 6; fi49++ {
                                var v50 uint8 = 92
                                v50 += v50
                                _ = fi49
                                _ = v50
                            }
                        default:
                            v51 := GenS5{F0: 2147483647}
                            _ = v51
                        }
                    }
                case -1:
                    for fi52 := 0; fi52 < 17; fi52++ {
                        var v53 uint64 = 0
                        switch 320 {
                        case -1:
                            v1 = (v2 & v1)
                        }
                        var v54 int8 = (v3 & v3)
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                case 2147483647:
                    var v55 uint64 = 287
                    _ = v55
                }
                _ = fi48
            }
            for fi56 := 0; fi56 < 13; fi56++ {
                var v57 bool = true
                _ = fi56
                _ = v57
            }
        }
    } else {
        if false {
            v58 := GenS6{F0: 42}
            _ = v58
        } else {
            var v59 int32 = -1
            fmt.Println(v3, v3)
            _ = v59
        }
        v4 = (v4 / ((-1) | 1))
        var v60 float64 = (0.1 - float64((v6 * float32(-2.25))))
        _ = v60
    }
    v61 := GenS1{F0: 2}
    for fi62 := 0; fi62 < 6; fi62++ {
        var v63 uint64 = 2
        _ = fi62
        _ = v63
    }
    var v64 uint8 = 143
    v6 -= (v6 - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v64
    _ = v7
    _ = v8
}
