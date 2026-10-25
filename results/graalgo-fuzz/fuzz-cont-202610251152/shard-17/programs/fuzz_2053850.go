package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int16; F1 int32; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint8; F1 int8; F2 int8 }
func genFn0(p0 int16, p1 uint16) int16 {
    var v4 int64 = 9223372036854775807
    var v5 uint16 = 226
    p0 = -32768
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p0
}
func genFn1() bool {
    var v6 uint64 = 2
    var v7 uint32 = 1
    v8 := []int{-755}
    _ = v6
    _ = v7
    _ = v8
    return (127 < 100)
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = (1.5 - 3.14159)
    var v3 uint8 = 42
    var v9 uint32 = 837
    v9 = (v9 ^ v9)
    delete(v1, "\n\t")
    if (v9 < 65) {
        switch 713 {
        case -742:
            v10 := GenS1{F0: -1, F1: "\n\t"}
            _ = v10
        case -157:
            var v11 int64 = 980
            _ = v11
        case 2147483647:
            for fi12 := 0; fi12 < 16; fi12++ {
                if (100 <= -1) {
                    v13 := GenS2{F0: 0, F1: 679, F2: v3}
                    _ = v13
                }
                if (!(false || (0 < int64(v9)))) {
                    if false {
                        if (!true) {
                            v14 := []int{(fi12 % ((fi12) | 1)), fi12, -9223372036854775808, fi12}
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 17; fi15++ {
                                var v16 uint16 = uint16(v9)
                                _ = fi15
                                _ = v16
                            }
                        }
                        for fi17 := 0; fi17 < 2; fi17++ {
                            v1[""] = (uint16(v9) > 2)
                            var v18 uint16 = 399
                            if ((!((0 >= -2147483648) && (uint8(v18) > v3))) || false) {
                                v18 += 854
                            }
                            _ = fi17
                            _ = v18
                        }
                    } else {
                        v19 := GenS3{F0: 107, F1: 4294967295, F2: v9}
                        var v20 int16 = 1
                        v21 := []int{(fi12 | (fi12 + fi12)), fi12, (fi12 & fi12), -591}
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    v9 = (v9 >> ((1) & 63))
                }
                _ = fi12
            }
        default:
            if ((-2147483648 > 895) && true) {
                v22 := GenS3{F0: -1, F1: 522, F2: 2147483648}
                for fi23 := 0; fi23 < 13; fi23++ {
                    v24 := []int{(-(fi23)), (fi23 << ((1) & 63)), fi23, (fi23 + 2147483647)}
                    _ = fi23
                    _ = v24
                }
                _ = v22
            }
        }
    } else {
        v25 := GenS2{F0: 769, F1: 1, F2: v3}
        v26 := []int{2147483647}
        var v27 int = int(v9)
        _ = v25
        _ = v26
        _ = v27
    }
    v9 = v9
    v3 -= (v3 | v3)
    if ((-1 == 0) && (uint32(v3) <= (^(v9)))) {
        switch 4294967295 {
        case 32767:
            v2 += float64(((0.1 * 1e10) * 3.14159))
        case -304:
            v28 := GenS4{F0: v3}
            _ = v28
        default:
            for fi29 := 0; fi29 < 20; fi29++ {
                v3 = v3
                _ = fi29
            }
        }
        var v30 float32 = (1e10 * 3.14159)
        for fi31 := 0; fi31 < 1; fi31++ {
            if true {
                v9 = (v9 >> ((42) & 63))
                var v32 float32 = ((((v30 + float32(v2)) / 0.1) / v30) * 0.1)
                if true {
                    if (42 >= int32(v9)) {
                        v2 += (((v2 + (v2 * (v2 + v2))) + 1.5) / v2)
                        if true {
                            delete(v1, ((((("a" + "\n\t") + "foo bar") + "") + ("foo bar" + "a")) + "foo bar"))
                            switch "" {
                            case "foo bar":
                                switch "\n\t" {
                                case "\n\t":
                                    var v33 float32 = 1.5
                                    _ = v33
                                default:
                                    for fi34 := 0; fi34 < 10; fi34++ {
                                        v35 := GenS1{F0: -1, F1: "hello"}
                                        if ("\n\t" <= "a") {
                                            if false {
                                                var v36 uint16 = 32767
                                                _ = v36
                                            }
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                v37 := GenS5{F0: (-(fi34)), F1: 2147483647}
                                                _ = v37
                                            case "hello":
                                                if false {
                                                    v32 = v30
                                                }
                                            }
                                            switch "\n\t" {
                                            case "":
                                                var v38 int = 32767
                                                _ = v38
                                            case "foo bar":
                                                v3 = 2
                                            default:
                                                _ = len(v1)
                                            }
                                            v39 := GenS3{F0: -78, F1: 715, F2: v9}
                                            if false {
                                                var v40 uint = 42
                                                var v41 uint = (v40 << (((v40 << ((v40) & 63))) & 63))
                                                _ = v40
                                                _ = v41
                                            } else {
                                                fmt.Println(v30)
                                                v2 += 0.1
                                                if false {
                                                    var v42 int16 = -409
                                                    if ((false && false) || true) {
                                                        if true {
                                                            v43 := GenS4{F0: 84}
                                                            switch "hello" {
                                                            case "a":
                                                                var v44 int16 = v42
                                                                _ = v44
                                                            case "":
                                                                if (((v9 & (v9 & (v9 << ((4294967295) & 63)))) > v9) || ((fi34 << ((4294967295) & 63)) == (fi31 | fi31))) {
                                                                    var v45 int16 = 0
                                                                    _ = v45
                                                                }
                                                            }
                                                            var v46 uint32 = (^(v9))
                                                            _ = v43
                                                            _ = v46
                                                        } else {
                                                            var v47 uint = 226
                                                            var v48 uint64 = 2
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                    }
                                                    var v49 bool = false
                                                    v9 = 2147483648
                                                    _ = v42
                                                    _ = v49
                                                }
                                            }
                                            _ = v39
                                        }
                                        var v50 int32 = 2147483647
                                        _ = len(v1)
                                        v9 += v9
                                        _ = fi34
                                        _ = v35
                                        _ = v50
                                    }
                                }
                            case "hello":
                                var v51 uint = 976
                                _ = v51
                            default:
                                v52 := GenS2{F0: 0, F1: -2147483648, F2: (v3 ^ uint8(v9))}
                                _ = v52
                            }
                        }
                    } else {
                        var v53 float32 = (v32 * v30)
                        fmt.Println(v53, v30)
                        if (true || false) {
                            for fi54 := 0; fi54 < 2; fi54++ {
                                if true {
                                    v30 = v32
                                }
                                if (false || true) {
                                    var v55 string = ""
                                    _ = v55
                                } else {
                                    fmt.Println(v3, v3, v53)
                                }
                                _ = fi54
                            }
                        } else {
                            v56 := GenS6{F0: (v3 | 128), F1: -128, F2: 25}
                            v30 = 1e10
                            _ = v56
                        }
                        _ = v53
                    }
                } else {
                    if (v32 == -2.25) {
                        v57 := v1[""]
                        _ = v57
                    }
                }
                _ = v32
            }
            var v58 string = "hello"
            _ = fi31
            _ = v58
        }
        _ = v30
    }
    _ = len(v1)
    v1["hello"] = (1.5 != 0.0)
    var v59 int = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v59:", v59)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v59
    _ = v9
}
