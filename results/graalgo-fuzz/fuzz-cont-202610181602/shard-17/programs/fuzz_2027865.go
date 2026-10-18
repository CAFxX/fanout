package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint16 }
func genFn0() bool {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 float64 = float64(1.5)
        var v9 string = ""
        for fi10 := 0; fi10 < 6; fi10++ {
            v9 = "foo bar"
            _ = fi10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v11 string = ("\n\t" + "foo bar")
    v11 = (v11 + v11)
    var v12 float32 = 0.1
    _ = v11
    _ = v12
    return (true && ((!((32767 >= 1) || true)) && false))
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 207
    v3 := []uint32{2147483648, 342, 65535, 65535}
    v4 := []uint32{}
    var v5 uint64 = 1
    var v6 float32 = (3.14159 / 1.0)
    v2 += v2
    for fi13 := 0; fi13 < 9; fi13++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v6 += 0.1
        _ = fi13
    }
    v14 := v1["foo bar"]
    v2 = v2
    v2 += (v2 >> ((1) & 63))
    var v15 uint64 = 725
    for fi16 := 0; fi16 < 8; fi16++ {
        if false {
            v15 = 891
            var v17 uint32 = 42
            _ = v17
        }
        v18 := GenS1{F0: 18446744073709551615}
        fmt.Println(fi16)
        _ = fi16
        _ = v18
    }
    if (false || true) {
        var v19 int32 = 2147483647
        switch (v14 / ((128) | 1)) {
        case 4294967295:
            var v20 float64 = 0.0
            _ = v20
        case 32768:
            if (0.1 < (0.1 / 1.5)) {
                for fi21 := 0; fi21 < 18; fi21++ {
                    v6 -= float32((0.1 + 0.1))
                    v4 = append(v4, 1)
                    _ = fi21
                }
                for fi22 := 0; fi22 < 4; fi22++ {
                    for fi23 := 0; fi23 < 16; fi23++ {
                        v3 = append(v3, 1)
                        var v24 uint64 = 1
                        var v25 int64 = 9223372036854775807
                        if ((!true) || ((true && true) || (fi22 != v14))) {
                            v25 += v25
                            switch int(v19) {
                            case 128:
                                v26 := GenS1{F0: 2}
                                if (0 > v25) {
                                    if true {
                                        v6 = v6
                                        v14 = v14
                                    } else {
                                        if (!(!true)) {
                                            if (1 >= 549) {
                                                fmt.Println(v2, fi22, v14)
                                            } else {
                                                switch "" {
                                                case "hello":
                                                    switch ("a" + "hello") {
                                                    case "":
                                                        if (float32(-2.25) < (3.14159 / -2.25)) {
                                                            v4 = append(v4, 906)
                                                            if true {
                                                                var v27 int16 = -32768
                                                                fmt.Println(v19)
                                                                var v28 int16 = 221
                                                                var v29 uint32 = 2
                                                                _ = v27
                                                                _ = v28
                                                                _ = v29
                                                            }
                                                            v6 = float32(1.5)
                                                            fmt.Println(fi23)
                                                            fmt.Println(fi22, v5)
                                                        } else {
                                                            v15 += v24
                                                            fmt.Println(v25, fi23, v24)
                                                            var v30 float64 = ((float64(v6) * -2.25) * 1.5)
                                                            _ = v30
                                                        }
                                                    case "foo bar":
                                                        v6 = v6
                                                    default:
                                                        v15 -= (v15 ^ v15)
                                                    }
                                                }
                                                v31 := GenS2{F0: -32768, F1: "a"}
                                                var v32 int = -536
                                                _ = v31
                                                _ = v32
                                            }
                                            fmt.Println(v19, v15, v15)
                                        } else {
                                            v33 := GenS1{F0: v24}
                                            v15 = (^(v15))
                                            var v34 int8 = 42
                                            _ = v33
                                            _ = v34
                                        }
                                    }
                                    var v35 uint32 = 42
                                    _ = v35
                                } else {
                                    if (!false) {
                                        var v36 string = "hello"
                                        _ = v36
                                    }
                                    if true {
                                        if (3.14159 != v6) {
                                            var v37 float32 = 0.0
                                            _ = v37
                                        }
                                        v15 = (^(v5))
                                        if ((v2 == (v2 % ((v2) | 1))) || (85 >= -128)) {
                                            var v38 int = fi23
                                            v24 -= 0
                                            _ = v38
                                        } else {
                                            v25 += 0
                                            v39 := GenS1{F0: (^(v15))}
                                            switch v14 {
                                            case 42:
                                                v24 += (v24 % (((v15 * (^(v5)))) | 1))
                                            }
                                            switch "" {
                                            case "hello":
                                                if true {
                                                    _ = len(v3)
                                                    v2 += (v2 << ((408) & 63))
                                                } else {
                                                    var v40 uint64 = (v24 + v24)
                                                    _ = v40
                                                }
                                            default:
                                                var v41 uint16 = 2
                                                _ = v41
                                            }
                                            _ = v39
                                        }
                                    } else {
                                        fmt.Println(v25, v5)
                                        switch (fi22 - (-(v14))) {
                                        case 729:
                                            fmt.Println(v6, v15, v24)
                                        default:
                                            switch "hello" {
                                            case "":
                                                v42 := GenS2{F0: 42, F1: (("a" + "") + "hello")}
                                                _ = v42
                                            case "hello":
                                                switch 0 {
                                                case -1:
                                                    var v43 uint8 = 255
                                                    _ = v43
                                                case 42:
                                                    var v44 float64 = 1.5
                                                    _ = v44
                                                case 144:
                                                    var v45 int16 = 1
                                                    _ = v45
                                                }
                                            case "foo bar":
                                                var v46 int = -26
                                                _ = v46
                                            default:
                                                var v47 uint64 = 1
                                                v15 = v24
                                                _ = v47
                                            }
                                        }
                                    }
                                }
                                _ = v26
                            default:
                                v25 = v25
                            }
                            fmt.Println(v2)
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = fi22
                }
            }
        }
        for fi48 := 0; fi48 < 6; fi48++ {
            if ((((("hello" + ("a" + "")) + "a") + (("" + "") + "")) < "hello") || ((("\n\t" + "a") + (((("a" + "") + "") + ("" + "a")) + "\n\t")) > "a")) {
                v49 := GenS1{F0: (v15 >> ((644) & 63))}
                var v50 int64 = int64(v5)
                var v51 float64 = 1e10
                _ = v49
                _ = v50
                _ = v51
            }
            var v52 bool = true
            _ = fi48
            _ = v52
        }
        _ = v19
    }
    v3 = append(v3, 1)
    for fi53 := 0; fi53 < 5; fi53++ {
        v54 := v1["a"]
        v55 := v1["a"]
        var v56 bool = (!true)
        _ = fi53
        _ = v54
        _ = v55
        _ = v56
    }
    for fi57 := 0; fi57 < 7; fi57++ {
        if false {
            for fi58 := 0; fi58 < 11; fi58++ {
                v59 := GenS3{F0: 32767, F1: 0, F2: 291}
                v2 -= uint8(v15)
                switch ("" + "hello") {
                case "hello":
                    v15 -= 936
                case "a":
                    var v60 int32 = 312
                    _ = v60
                case "\n\t":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi58
                _ = v59
            }
            v61 := GenS1{F0: (v15 + v15)}
            if false {
                for fi62 := 0; fi62 < 10; fi62++ {
                    var v63 int32 = 10
                    _ = fi62
                    _ = v63
                }
                v2 -= v2
            } else {
                var v64 bool = true
                var v65 int32 = 42
                _ = v64
                _ = v65
            }
            _ = v61
        }
        _ = fi57
    }
    switch (-(v14)) {
    case 9223372036854775807:
        v66 := GenS3{F0: 864, F1: 1, F2: v5}
        _ = v66
    case 641:
        var v67 string = ("hello" + (("" + ("" + "foo bar")) + "\n\t"))
        _ = v67
    case 32767:
        var v68 uint = 42
        _ = v68
    }
    switch "a" {
    case "hello":
        var v69 uint32 = 621
        _ = v69
    default:
        v70 := GenS4{F0: 1, F1: 18446744073709551615, F2: 392}
        _ = v70
    }
    for fi71 := 0; fi71 < 18; fi71++ {
        v72 := v1["foo bar"]
        fmt.Println(fi71)
        _ = fi71
        _ = v72
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
