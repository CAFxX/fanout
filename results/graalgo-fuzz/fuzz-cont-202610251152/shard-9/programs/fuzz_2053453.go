package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int; F1 int8; F2 string }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 string }
func genFn0(p0 int) string {
    fmt.Println(p0)
    if false {
        var v6 string = ""
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        switch 127 {
        case 127:
            v8 := GenS1{F0: 619}
            _ = v8
        case 2147483647:
            var v9 float32 = 3.14159
            _ = v9
        case 9223372036854775807:
            p0 = (p0 ^ fi7)
        default:
            for fi10 := 0; fi10 < 3; fi10++ {
                fmt.Println(fi10, fi7, p0)
                _ = fi10
            }
        }
        _ = fi7
    }
    v11 := []int{p0}
    _ = p0
    _ = v11
    return ("foo bar" + "a")
}
func genFn1(p0 uint32, p1 uint32, p2 uint16) uint32 {
    p2 = p2
    v12 := []int{32767}
    v13 := GenS2{F0: -9223372036854775808, F1: 327}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return (p0 * 4294967295)
}

func main() {
    var v1 int = 42
    v2 := make(map[string]uint32)
    var v3 int16 = 1000
    v4 := make(map[string]int)
    v5 := []int{v1, v1, (v1 * v1)}
    fmt.Println(v1, v3)
    if false {
        v14 := GenS1{F0: 18446744073709551615}
        if false {
            v1 -= (^(v1))
            for fi15 := 0; fi15 < 17; fi15++ {
                if (-733 >= 1) {
                    v4["a"] = (v1 >> ((1) & 63))
                    var v16 uint32 = 307
                    _ = v16
                } else {
                    v1 = (-(v1))
                    v1 -= (fi15 * fi15)
                    if false {
                        v1 = v1
                        switch "\n\t" {
                        case "hello":
                            switch 0 {
                            case 1:
                                if (!(((float32(1e10) - 3.14159) + float32((-2.25 + ((0.1 / 3.14159) * -2.25)))) <= (3.14159 / 1.5))) {
                                    for fi17 := 0; fi17 < 6; fi17++ {
                                        fmt.Println(fi15)
                                        v18 := GenS3{F0: -128, F1: 127}
                                        _ = fi17
                                        _ = v18
                                    }
                                }
                            case 128:
                                v19 := GenS2{F0: (^(fi15)), F1: 0}
                                _ = v19
                            case 32768:
                                var v20 int = fi15
                                _ = v20
                            default:
                                v1 = v1
                            }
                        }
                        fmt.Println(v1)
                    } else {
                        v21 := GenS3{F0: -43, F1: 76}
                        if len(v5) > 0 {
                            v5[0] = (v1 * (fi15 - 32767))
                        }
                        var v22 int8 = -1
                        _ = v21
                        _ = v22
                    }
                }
                if false {
                    fmt.Println(v1, v1, fi15)
                    v3 += v3
                    _ = len(v5)
                } else {
                    switch 127 {
                    case 127:
                        v23 := GenS1{F0: 479}
                        _ = v23
                    }
                    if true {
                        switch 128 {
                        case -551:
                            switch ("hello" + ("\n\t" + "\n\t")) {
                            case "":
                                for fi24 := 0; fi24 < 17; fi24++ {
                                    var v25 string = ""
                                    _ = fi24
                                    _ = v25
                                }
                            default:
                                v2[""] = 65535
                            }
                        }
                        var v26 bool = (("foo bar" > "") || true)
                        _ = v26
                    } else {
                        v27 := GenS4{F0: -9223372036854775808, F1: -128, F2: "\n\t"}
                        _ = v27
                    }
                    switch (fi15 / ((-943) | 1)) {
                    case 1:
                        v28 := GenS5{F0: 1, F1: 255}
                        _ = v28
                    }
                }
                if true {
                    if ("a" < "a") {
                        for fi29 := 0; fi29 < 3; fi29++ {
                            var v30 bool = true
                            fmt.Println(fi15, fi15, fi15)
                            _ = fi29
                            _ = v30
                        }
                    }
                    switch (v1 * v1) {
                    case 32768:
                        if (1 <= 1000) {
                            v31 := GenS4{F0: fi15, F1: 37, F2: ""}
                            _ = v31
                        } else {
                            switch 462 {
                            case 42:
                                for fi32 := 0; fi32 < 19; fi32++ {
                                    v4["\n\t"] = fi32
                                    if (255 >= 0) {
                                        v1 -= 9223372036854775807
                                        v3 = (v3 & 716)
                                        v33 := GenS2{F0: v1, F1: 1}
                                        _ = v33
                                    }
                                    var v34 uint64 = 18446744073709551615
                                    _ = fi32
                                    _ = v34
                                }
                            default:
                                if (true || (42 < 899)) {
                                    for fi35 := 0; fi35 < 19; fi35++ {
                                        var v36 int8 = -1
                                        fmt.Println(fi35, v36)
                                        if false {
                                            _ = len(v2)
                                            v3 -= -1
                                            var v37 uint16 = 65535
                                            v38 := GenS3{F0: (v36 << ((599) & 63)), F1: 255}
                                            _ = v37
                                            _ = v38
                                        }
                                        if (-9223372036854775808 <= 1) {
                                            switch fi35 {
                                            case 32767:
                                                v1 += v1
                                            case 9223372036854775807:
                                                var v39 int16 = (v3 >> ((383) & 63))
                                                _ = v39
                                            case -9223372036854775808:
                                                v40 := GenS2{F0: -1, F1: 959}
                                                _ = v40
                                            default:
                                                v3 = (v3 + -32768)
                                            }
                                            var v41 float64 = 0.0
                                            v3 -= v3
                                            v3 = v3
                                            var v42 int64 = -1
                                            _ = v41
                                            _ = v42
                                        }
                                        switch v1 {
                                        case 128:
                                            v3 -= v3
                                        case 2147483647:
                                            fmt.Println(v36, v3)
                                        }
                                        _ = fi35
                                        _ = v36
                                    }
                                } else {
                                    var v43 uint64 = 0
                                    for fi44 := 0; fi44 < 17; fi44++ {
                                        switch v1 {
                                        case 32768:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case 127:
                                            v3 = (v3 >> ((255) & 63))
                                        default:
                                            fmt.Println(fi15)
                                        }
                                        switch ("a" + "hello") {
                                        case "foo bar":
                                            v45 := GenS6{F0: 891}
                                            _ = v45
                                        default:
                                            var v46 int = v1
                                            _ = v46
                                        }
                                        _ = fi44
                                    }
                                    v43 = (v43 & 2)
                                    _ = v43
                                }
                            }
                            var v47 bool = ((-60 != 0) && (-2.25 > ((-2.25 / 0.1) / 1.0)))
                            _ = v47
                        }
                    case 1:
                        for fi48 := 0; fi48 < 1; fi48++ {
                            switch ("hello" + "") {
                            case "foo bar":
                                fmt.Println(fi15, v3)
                            case "\n\t":
                                var v49 bool = false
                                _ = v49
                            case "":
                                var v50 int16 = v3
                                _ = v50
                            default:
                                fmt.Println(fi48, v3)
                            }
                            v1 += (fi48 - (fi48 * v1))
                            if (1 > (v3 + v3)) {
                                v3 += -721
                            }
                            _ = fi48
                        }
                    }
                }
                _ = fi15
            }
        }
        _ = v14
    } else {
        v51 := GenS6{F0: 445}
        v52 := v4["hello"]
        _ = v51
        _ = v52
    }
    var v53 int16 = 409
    var v54 uint32 = 0
    var v55 int8 = 0
    v55 += (v55 | v55)
    for fi56 := 0; fi56 < 10; fi56++ {
        var v57 int16 = v3
        var v58 uint32 = (v54 % (((v54 % ((v54) | 1))) | 1))
        _ = fi56
        _ = v57
        _ = v58
    }
    v59 := GenS7{F0: 68, F1: 9223372036854775807}
    if (97 >= 4294967295) {
        _ = len(v2)
    }
    for fi60 := 0; fi60 < 20; fi60++ {
        if (!false) {
            switch "hello" {
            case "a":
                fmt.Println(v54)
            case "hello":
                v1 = fi60
            }
            if (int(v55) < v1) {
                fmt.Println(v53, v55)
                v3 = (v3 ^ 1)
                var v61 float64 = 0.0
                _ = v61
            }
            for fi62 := 0; fi62 < 6; fi62++ {
                v1 = 4294967295
                _ = fi62
            }
        }
        v63 := GenS8{F0: ""}
        var v64 int8 = (v55 / (((v55 * v55)) | 1))
        _ = fi60
        _ = v63
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v59
}
