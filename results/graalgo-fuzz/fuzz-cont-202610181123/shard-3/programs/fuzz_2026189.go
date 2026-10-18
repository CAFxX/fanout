package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint32 }
func genFn0() uint8 {
    var v4 uint8 = 127
    var v5 int64 = 42
    v6 := GenS1{F0: 0, F1: 975}
    _ = v4
    _ = v5
    _ = v6
    return v4
}
func genFn1() uint64 {
    var v7 int = 0
    for fi8 := 0; fi8 < 15; fi8++ {
        v7 = fi8
        _ = fi8
    }
    var v9 int32 = -1
    _ = v7
    _ = v9
    return 42
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]int)
    var v3 int = -735
    if true {
        var v10 int = (v3 ^ v3)
        v11 := GenS2{F0: ""}
        _ = v10
        _ = v11
    }
    var v12 int8 = -1
    v13 := GenS1{F0: -9223372036854775808, F1: 0}
    var v14 uint8 = v1
    v12 = v12
    v3 += v3
    if (-1 != v12) {
        var v15 uint32 = 2147483648
        _ = v15
    }
    v16 := []int{(v3 + (v3 | v3)), -9223372036854775808}
    var v17 uint = 1
    for fi18 := 0; fi18 < 1; fi18++ {
        switch "a" {
        case "\n\t":
            switch (((("\n\t" + "") + "hello") + ("foo bar" + "")) + "foo bar") {
            case "a":
                var v19 uint = v17
                _ = v19
            case "foo bar":
                if (146 < 18446744073709551615) {
                    v17 = (v17 | (^(v17)))
                    switch "\n\t" {
                    case "":
                        if (9223372036854775807 == 42) {
                            v1 -= 0
                        }
                    case "a":
                        if ((v1 + (v14 / (((^(v1))) | 1))) > 42) {
                            v3 += (v3 << ((v17) & 63))
                            var v20 int8 = v12
                            var v21 uint32 = uint32(v14)
                            _ = v20
                            _ = v21
                        }
                    }
                    delete(v2, "hello")
                } else {
                    v22 := GenS2{F0: ("" + "foo bar")}
                    v17 = 634
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v24 := GenS2{F0: ("" + "hello")}
                        v25 := v2["\n\t"]
                        var v26 float32 = (float32(1.5) + (1.5 * 0.0))
                        var v27 bool = false
                        var v28 uint32 = 560
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    _ = v22
                }
            case "hello":
                v14 = (v14 % (((v1 % ((203) | 1))) | 1))
            default:
                fmt.Println(v1, v12)
            }
        }
        var v29 int16 = -32768
        _ = fi18
        _ = v29
    }
    if (!false) {
        if false {
            var v30 bool = ((1 == uint32(v17)) && (false && false))
            _ = v30
        } else {
            switch "a" {
            case "foo bar":
                if true {
                    if len(v16) > 0 {
                        v16[0] = (v3 * 2147483647)
                    }
                    if (true && true) {
                        switch v3 {
                        case -2147483648:
                            var v31 uint8 = 127
                            _ = v31
                        case 0:
                            if (-863 <= 1) {
                                var v32 bool = (false || ((v3 >= v3) && (true && false)))
                                _ = v32
                            } else {
                                var v33 string = ""
                                for fi34 := 0; fi34 < 10; fi34++ {
                                    v14 = 127
                                    if false {
                                        switch v33 {
                                        case "a":
                                            if false {
                                                var v35 uint16 = 343
                                                for fi36 := 0; fi36 < 15; fi36++ {
                                                    if false {
                                                        v14 = 26
                                                        fmt.Println(v12)
                                                        var v37 string = (v33 + "\n\t")
                                                        switch v33 {
                                                        case "foo bar":
                                                            if false {
                                                                v12 -= (v12 | v12)
                                                                if false {
                                                                    var v38 int = 128
                                                                    v39 := GenS3{F0: true, F1: 2}
                                                                    var v40 float64 = -2.25
                                                                    _ = v38
                                                                    _ = v39
                                                                    _ = v40
                                                                }
                                                            }
                                                        case "\n\t":
                                                            v17 -= (v17 / ((v17) | 1))
                                                        case "a":
                                                            if (2147483648 != 2) {
                                                                v35 = (^(v35))
                                                                v12 = (v12 << ((v17) & 63))
                                                            }
                                                        }
                                                        var v41 int8 = (^(v12))
                                                        _ = v37
                                                        _ = v41
                                                    }
                                                    v42 := v2[(v33 + v33)]
                                                    v43 := GenS3{F0: (128 <= 0), F1: 1}
                                                    _ = fi36
                                                    _ = v42
                                                    _ = v43
                                                }
                                                _ = v35
                                            }
                                        case "hello":
                                            v33 = "\n\t"
                                        case "\n\t":
                                            v1 += genFn0()
                                        }
                                    } else {
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                        v16 = append(v16, (v3 / ((0) | 1)))
                                        for fi44 := 0; fi44 < 16; fi44++ {
                                            v45 := v2["\n\t"]
                                            v46 := GenS1{F0: -1, F1: 1}
                                            fmt.Println(v17, v33, v17)
                                            _ = fi44
                                            _ = v45
                                            _ = v46
                                        }
                                    }
                                    _ = fi34
                                }
                                switch -306 {
                                case 174:
                                    for fi47 := 0; fi47 < 7; fi47++ {
                                        var v48 string = ((v33 + v33) + "")
                                        v16 = append(v16, fi47)
                                        v16 = append(v16, (-(v3)))
                                        _ = fi47
                                        _ = v48
                                    }
                                }
                                _ = v33
                            }
                        }
                    }
                } else {
                    _ = len(v16)
                    if false {
                        var v49 string = "foo bar"
                        _ = v49
                    }
                    var v50 uint16 = 317
                    _ = v50
                }
            }
        }
        if (!(!true)) {
            fmt.Println(v17, v3, v1)
            if (1 == -501) {
                switch v3 {
                case 42:
                    var v51 uint = (v17 | v17)
                    _ = v51
                case -2147483648:
                    v17 = 18446744073709551615
                }
            }
        } else {
            v12 += (v12 - v12)
            for fi52 := 0; fi52 < 12; fi52++ {
                v12 = v12
                if true {
                    fmt.Println(v12, v14, v1)
                    for fi53 := 0; fi53 < 5; fi53++ {
                        switch v3 {
                        case 128:
                            switch ("" + "\n\t") {
                            case "\n\t":
                                var v54 uint8 = 42
                                _ = v54
                            default:
                                v1 = v1
                            }
                        }
                        var v55 string = ("" + ("" + "foo bar"))
                        var v56 int64 = 42
                        _ = fi53
                        _ = v55
                        _ = v56
                    }
                    var v57 uint16 = 42
                    _ = v57
                }
                switch 32768 {
                case 9223372036854775807:
                    v17 -= v17
                default:
                    var v58 int32 = -1
                    _ = v58
                }
                _ = fi52
            }
            switch (-(v3)) {
            case 1:
                fmt.Println(v14, v12, v3)
            default:
                for fi59 := 0; fi59 < 16; fi59++ {
                    v14 -= v1
                    for fi60 := 0; fi60 < 19; fi60++ {
                        fmt.Println(fi59, fi59, v3)
                        _ = fi60
                    }
                    switch (v3 ^ 9223372036854775807) {
                    case 96:
                        if len(v16) > 0 {
                            v16[0] = 42
                        }
                    case 128:
                        if (false || ((42 == 0) || (false || true))) {
                            var v61 uint = v17
                            _ = v61
                        }
                    default:
                        fmt.Println(v1, v3)
                    }
                    _ = fi59
                }
            }
        }
    }
    var v62 int8 = (v12 - v12)
    var v63 uint64 = 14
    v14 = v1
    if (1 <= 285) {
        fmt.Println(v3, v63)
        switch v3 {
        case 9223372036854775807:
            v64 := v2["foo bar"]
            _ = v64
        default:
            var v65 float64 = 3.14159
            _ = v65
        }
        for fi66 := 0; fi66 < 15; fi66++ {
            if false {
                var v67 uint8 = (v1 * (v14 | (^(v14))))
                _ = v67
            }
            v17 = 42
            _ = fi66
        }
    } else {
        v17 -= (v17 * v17)
        v68 := GenS4{F0: false, F1: 306}
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        _ = v68
    }
    for fi69 := 0; fi69 < 9; fi69++ {
        switch ("" + "hello") {
        case "foo bar":
            switch (fi69 / ((fi69) | 1)) {
            case 128:
                var v70 uint = 1
                _ = v70
            default:
                v1 -= (v14 << ((v17) & 63))
            }
        case "":
            fmt.Println(v17, v14)
        default:
            delete(v2, "a")
        }
        v2[("hello" + "foo bar")] = 127
        var v71 float64 = 0.0
        _ = fi69
        _ = v71
    }
    var v72 uint8 = (v1 % ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v62
    _ = v63
    _ = v72
}
