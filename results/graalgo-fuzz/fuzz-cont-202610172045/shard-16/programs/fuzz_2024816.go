package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint8; F1 int; F2 int }
type GenS5 struct { F0 int16; F1 int32; F2 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint; F1 uint32; F2 int8 }
type GenS10 struct { F0 uint32; F1 uint64; F2 int64 }
func genFn0(p0 int8, p1 uint8) uint8 {
    v7 := GenS1{F0: 57}
    var v8 uint = 718
    var v9 bool = ((p0 ^ 81) >= (p0 & p0))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 int = 501
    v2 := make(map[string]bool)
    v3 := []string{"foo bar", "", "\n\t"}
    var v4 float32 = (-2.25 * 1.5)
    v5 := []int{v1, -1, (v1 | (v1 | 1))}
    var v6 uint16 = 65535
    if true {
        for fi10 := 0; fi10 < 12; fi10++ {
            v4 += (v4 / (3.14159 - v4))
            v1 += fi10
            _ = fi10
        }
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := GenS2{F0: 0, F1: 1}
            v6 = v6
            for fi13 := 0; fi13 < 12; fi13++ {
                fmt.Println(fi11)
                v2["\n\t"] = true
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
        if true {
            for fi14 := 0; fi14 < 3; fi14++ {
                for fi15 := 0; fi15 < 3; fi15++ {
                    if ((-2.25 > 3.14159) || true) {
                        if (762 <= -840) {
                            if false {
                                _ = len(v2)
                                v1 = (fi15 ^ (v1 << ((916) & 63)))
                                v16 := GenS3{F0: v6, F1: 735}
                                _ = v16
                            }
                            v17 := GenS4{F0: 0, F1: 29, F2: (fi15 ^ (fi14 * v1))}
                            fmt.Println(fi15, fi15, v1)
                            _ = v17
                        }
                    } else {
                        v6 += (v6 << ((uint(v6)) & 63))
                        if (!((2 >= 255) && true)) {
                            switch "\n\t" {
                            case "foo bar":
                                v5 = append(v5, v1)
                            case "hello":
                                v18 := GenS1{F0: 498}
                                _ = v18
                            }
                            v19 := GenS4{F0: 127, F1: fi15, F2: (v1 >> ((uint(v6)) & 63))}
                            var v20 uint32 = 1
                            v6 = (v6 ^ (v6 / (((^(v6))) | 1)))
                            switch ((("\n\t" + ("hello" + "foo bar")) + "foo bar") + "foo bar") {
                            case "\n\t":
                                if (int64(v20) == 42) {
                                    var v21 int64 = -416
                                    fmt.Println(fi15, v1)
                                    var v22 uint = 62
                                    _ = v21
                                    _ = v22
                                } else {
                                    v23 := GenS2{F0: -793, F1: -1}
                                    fmt.Println(fi15, v1)
                                    _ = v23
                                }
                            case "":
                                var v24 uint = 42
                                _ = v24
                            default:
                                switch fi15 {
                                case 32767:
                                    v25 := v2["foo bar"]
                                    _ = v25
                                case 128:
                                    v26 := v2[("foo bar" + "\n\t")]
                                    _ = v26
                                }
                            }
                            _ = v19
                            _ = v20
                        }
                        if ((!false) || (3.14159 <= (0.0 + v4))) {
                            if true {
                                delete(v2, "")
                            } else {
                                if true {
                                    var v27 float64 = float64((v4 - v4))
                                    v28 := GenS5{F0: 690, F1: 2147483647, F2: "hello"}
                                    if (v1 < (-(fi15))) {
                                        if false {
                                            v29 := GenS6{F0: -883}
                                            switch ("hello" + (("" + "a") + "hello")) {
                                            case "hello":
                                                fmt.Println(fi14, v27)
                                                v2["\n\t"] = false
                                                if (!((!false) && (32767 == (fi14 << ((42) & 63))))) {
                                                    var v30 int64 = -9223372036854775808
                                                    _ = v30
                                                }
                                            case "a":
                                                switch (fi15 << ((0) & 63)) {
                                                case 128:
                                                    if true {
                                                        v4 = v4
                                                        if false {
                                                            fmt.Println(v6, fi15)
                                                        } else {
                                                            v31 := GenS3{F0: 32767, F1: -816}
                                                            v5 = append(v5, (v1 * (fi14 / (((v1 / ((v1) | 1))) | 1))))
                                                            v32 := v2[("a" + "a")]
                                                            var v33 int8 = 40
                                                            _ = v31
                                                            _ = v32
                                                            _ = v33
                                                        }
                                                    }
                                                case 32768:
                                                    fmt.Println(v27, v27, v27)
                                                }
                                            case "\n\t":
                                                v4 -= float32(((v27 * ((float64((((v4 + (v4 * (0.0 * -2.25))) * float32(0.1)) / v4)) * v27) - (v27 * v27))) / v27))
                                            default:
                                                var v34 uint8 = 128
                                                _ = v34
                                            }
                                            _ = v29
                                        } else {
                                            v27 -= ((-2.25 - (1.5 + (v27 - 1.5))) / (v27 - float64(1.5)))
                                        }
                                        v3 = append(v3, "a")
                                        delete(v2, "\n\t")
                                    } else {
                                        v1 += (-(v1))
                                    }
                                    var v35 float32 = v4
                                    _ = v27
                                    _ = v28
                                    _ = v35
                                } else {
                                    if true {
                                        v2[(("foo bar" + "a") + (("a" + "hello") + "hello"))] = true
                                        v36 := GenS5{F0: -1, F1: 0, F2: ("hello" + ("hello" + ("hello" + "\n\t")))}
                                        _ = v36
                                    } else {
                                        var v37 uint64 = 2
                                        _ = v37
                                    }
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v38 float64 = 0.0
                                _ = v38
                            }
                        }
                        var v39 int64 = 1
                        _ = v39
                    }
                    fmt.Println(v6, v1)
                    _ = fi15
                }
                v6 = v6
                _ = fi14
            }
            v40 := GenS3{F0: (v6 * v6), F1: -909}
            var v41 int8 = 127
            _ = v40
            _ = v41
        }
    } else {
        var v42 float64 = 0.0
        var v43 int8 = -128
        _ = v42
        _ = v43
    }
    v44 := v2[("" + "\n\t")]
    v4 = (v4 / v4)
    _ = len(v5)
    for fi45 := 0; fi45 < 18; fi45++ {
        switch "hello" {
        case "hello":
            var v46 float32 = float32((0.0 + (((1.5 - 1e10) - (1e10 / (float64(v4) + 3.14159))) / (0.0 - 0.1))))
            _ = v46
        case "\n\t":
            var v47 uint8 = 1
            _ = v47
        case "":
            v48 := GenS2{F0: -1, F1: 127}
            _ = v48
        }
        _ = fi45
    }
    for fi49 := 0; fi49 < 8; fi49++ {
        v6 += v6
        _ = fi49
    }
    var v50 uint8 = 255
    for fi51 := 0; fi51 < 2; fi51++ {
        var v52 string = ("a" + "a")
        switch 0 {
        case -9223372036854775808:
            if true {
                for fi53 := 0; fi53 < 9; fi53++ {
                    v1 -= int(v50)
                    v6 = v6
                    var v54 uint = 129
                    _ = len(v2)
                    _ = fi53
                    _ = v54
                }
            } else {
                v1 += v1
                var v55 float32 = v4
                for fi56 := 0; fi56 < 19; fi56++ {
                    var v57 int = fi56
                    delete(v2, v52)
                    v58 := GenS7{F0: 267}
                    _ = fi56
                    _ = v57
                    _ = v58
                }
                _ = v55
            }
        case 4294967295:
            var v59 int8 = 8
            _ = v59
        }
        _ = fi51
        _ = v52
    }
    switch "a" {
    case "\n\t":
        var v60 int64 = 1
        _ = v60
    case "a":
        switch "\n\t" {
        case "foo bar":
            var v61 int64 = 1
            _ = v61
        }
    }
    v62 := GenS1{F0: 0}
    switch "\n\t" {
    case "\n\t":
        var v63 uint32 = 258
        _ = v63
    default:
        v44 = true
    }
    var v64 int64 = 1
    v65 := GenS8{F0: 635}
    v66 := GenS5{F0: 288, F1: int32(v50), F2: "hello"}
    switch v1 {
    case -2147483648:
        for fi67 := 0; fi67 < 18; fi67++ {
            switch (fi67 >> ((uint(v50)) & 63)) {
            case 0:
                switch ("hello" + "a") {
                case "hello":
                    switch ("a" + "hello") {
                    case "foo bar":
                        fmt.Println(v44, v6)
                    case "a":
                        v68 := GenS9{F0: 255, F1: 582, F2: 42}
                        _ = v68
                    default:
                        for fi69 := 0; fi69 < 3; fi69++ {
                            v1 -= 0
                            switch "a" {
                            case "hello":
                                if ((-1 == -128) && true) {
                                    fmt.Println(fi69, fi69)
                                }
                            case "":
                                var v70 int8 = 0
                                if (4294967295 < 1) {
                                    var v71 int64 = (v64 >> ((uint(v6)) & 63))
                                    _ = v71
                                }
                                var v72 int8 = v70
                                _ = v70
                                _ = v72
                            case "foo bar":
                                var v73 int16 = 0
                                _ = v73
                            }
                            _ = fi69
                        }
                    }
                case "":
                    v2["hello"] = v44
                }
            case 42:
                v50 = v50
            }
            if len(v5) > 0 {
                v5[0] = 861
            }
            if v44 {
                var v74 uint = 596
                _ = v74
            }
            _ = fi67
        }
    case 32767:
        v75 := GenS10{F0: 65535, F1: 16, F2: (v64 >> ((431) & 63))}
        _ = v75
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v50
    _ = v6
    _ = v62
    _ = v64
    _ = v65
    _ = v66
}
