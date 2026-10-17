package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 int64; F2 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int; F1 int16 }

func main() {
    var v1 string = (("hello" + ((("hello" + ((("" + "\n\t") + ((("\n\t" + "foo bar") + "a") + "a")) + "a")) + "") + ("" + ("\n\t" + "a")))) + "hello")
    v2 := make(map[uint32]bool)
    var v3 int32 = -968
    v4 := make(map[string]int)
    v5 := []int{1, 32767, 127, 2147483647}
    var v6 uint = 18446744073709551615
    if (65535 <= 825) {
        var v7 int8 = 127
        _ = v7
    } else {
        for fi8 := 0; fi8 < 7; fi8++ {
            if (-1 < 42) {
                v9 := GenS1{F0: (((!((42 > int16(fi8)) && true)) || true) || true), F1: -128}
                _ = v9
            }
            _ = fi8
        }
        var v10 uint32 = 2
        _ = v10
    }
    var v11 string = "hello"
    if (4294967295 == 2147483647) {
        var v12 int8 = 100
        var v13 bool = false
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: (!("hello" >= "")), F1: 0}
    fmt.Println(v6, v6, v6)
    fmt.Println(v11, v1, v11)
    switch 32768 {
    case 464:
        if (-32768 >= -32768) {
            var v15 int64 = -9223372036854775808
            v15 += v15
            _ = v15
        }
    case 127:
        if false {
            v16 := GenS2{F0: -1}
            v1 = "\n\t"
            v3 -= v3
            _ = v16
        }
    case 2147483647:
        if false {
            switch 32768 {
            case -1:
                if false {
                    if ((true || ((255 >= 42) || (!(v3 != v3)))) && false) {
                        for fi17 := 0; fi17 < 4; fi17++ {
                            fmt.Println(v3, v1, v6)
                            if (v3 > int32(v6)) {
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    v11 = (v1 + ((((v11 + v11) + ((("hello" + (v11 + (v11 + ((("" + ("" + v1)) + v11) + v1)))) + v11) + "a")) + "\n\t") + (("a" + v11) + v11)))
                                    _ = fi18
                                }
                                v6 += (^(v6))
                                v3 = v3
                            } else {
                                if false {
                                    v11 = v1
                                    for fi19 := 0; fi19 < 6; fi19++ {
                                        switch ((((v1 + ("" + v11)) + (v1 + ((v11 + ("\n\t" + v11)) + ("hello" + v1)))) + "hello") + (("a" + v1) + v1)) {
                                        case "foo bar":
                                            switch 127 {
                                            case 9223372036854775807:
                                                switch (fi19 * 2147483647) {
                                                case 0:
                                                    var v20 uint64 = 42
                                                    _ = v20
                                                case 9223372036854775807:
                                                    v2[1] = true
                                                }
                                            default:
                                                v21 := GenS3{F0: -9223372036854775808, F1: 42, F2: 0}
                                                _ = v21
                                            }
                                        default:
                                            v4[v1] = (fi19 ^ fi19)
                                        }
                                        switch fi17 {
                                        case 127:
                                            if false {
                                                v22 := GenS3{F0: -9223372036854775808, F1: 1, F2: 356}
                                                if (!false) {
                                                    if false {
                                                        v23 := v4[v11]
                                                        _ = v23
                                                    } else {
                                                        fmt.Println(v3, v11)
                                                    }
                                                    if (127 != -1) {
                                                        v24 := GenS2{F0: -1}
                                                        v11 = "\n\t"
                                                        var v25 string = v1
                                                        _ = v24
                                                        _ = v25
                                                    }
                                                    if true {
                                                        v3 += (v3 + -2147483648)
                                                        var v26 string = v11
                                                        _ = v26
                                                    }
                                                }
                                                var v27 int64 = -1
                                                if true {
                                                    if (536 > 206) {
                                                        _ = len(v5)
                                                        _ = len(v5)
                                                    }
                                                    delete(v2, 4294967295)
                                                }
                                                var v28 uint64 = 4294967295
                                                _ = v22
                                                _ = v27
                                                _ = v28
                                            } else {
                                                var v29 float32 = 0.0
                                                if false {
                                                    v1 = (v1 + v11)
                                                    v30 := GenS1{F0: true, F1: -104}
                                                    _ = v30
                                                } else {
                                                    if (!true) {
                                                        var v31 int16 = -971
                                                        fmt.Println(v1)
                                                        _ = v31
                                                    } else {
                                                        fmt.Println(fi17, fi19, v29)
                                                        _ = len(v4)
                                                    }
                                                    if (0 == (fi17 << ((2) & 63))) {
                                                        var v32 uint64 = 92
                                                        fmt.Println(v29, v3, v29)
                                                        var v33 uint16 = 444
                                                        if (99 > 127) {
                                                            v34 := GenS1{F0: false, F1: 1}
                                                            v35 := GenS4{F0: 0}
                                                            switch (fi19 >> ((v6) & 63)) {
                                                            case 762:
                                                                v36 := GenS5{F0: 128, F1: 9223372036854775807, F2: -1}
                                                                _ = v36
                                                            case -1:
                                                                var v37 float64 = (0.1 - -2.25)
                                                                v38 := GenS6{F0: 0}
                                                                if true {
                                                                    switch -843 {
                                                                    case -912:
                                                                        delete(v2, 4294967295)
                                                                    default:
                                                                        var v39 bool = false
                                                                        _ = v39
                                                                    }
                                                                    v40 := v4[""]
                                                                    v6 += 42
                                                                    _ = v40
                                                                } else {
                                                                    v3 = v3
                                                                    switch int(v33) {
                                                                    case -852:
                                                                        v29 += v29
                                                                    case -294:
                                                                        var v41 uint32 = uint32(v32)
                                                                        _ = v41
                                                                    case 42:
                                                                        delete(v2, 1)
                                                                    }
                                                                }
                                                                var v42 uint = 2
                                                                _ = v37
                                                                _ = v38
                                                                _ = v42
                                                            case 99:
                                                                fmt.Println(v29, fi19)
                                                            default:
                                                                switch ("a" + v11) {
                                                                case "\n\t":
                                                                    var v43 int64 = 0
                                                                    _ = v43
                                                                case "":
                                                                    var v44 float32 = (v29 * v29)
                                                                    _ = v44
                                                                }
                                                            }
                                                            _ = v34
                                                            _ = v35
                                                        } else {
                                                            if ((0.1 * 3.14159) != (-2.25 * -2.25)) {
                                                                var v45 bool = true
                                                                _ = v45
                                                            } else {
                                                                var v46 int32 = v3
                                                                fmt.Println(v6, fi19, v1)
                                                                switch v11 {
                                                                case "hello":
                                                                    if len(v5) > 0 {
                                                                        v5[0] = (fi19 & (fi17 | int(v32)))
                                                                    }
                                                                default:
                                                                    var v47 uint8 = 2
                                                                    _ = v47
                                                                }
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                delete(v2, 1)
                                                                _ = v46
                                                            }
                                                        }
                                                        var v48 uint32 = 257
                                                        _ = v32
                                                        _ = v33
                                                        _ = v48
                                                    }
                                                }
                                                _ = v29
                                            }
                                        }
                                        if true {
                                            v1 = v1
                                            if (false && true) {
                                                var v49 int32 = (v3 ^ -1)
                                                v11 = ""
                                                v6 += (^(v6))
                                                _ = v49
                                            }
                                            v50 := GenS7{F0: 822}
                                            _ = v50
                                        }
                                        v51 := GenS2{F0: int64(v3)}
                                        _ = fi19
                                        _ = v51
                                    }
                                    v11 = ("a" + "")
                                }
                                for fi52 := 0; fi52 < 5; fi52++ {
                                    v6 += (v6 ^ (v6 << ((1) & 63)))
                                    var v53 float32 = (float32(0.0) + (0.0 * ((float32(3.14159) - (float32((0.0 / 0.1)) - 3.14159)) + 1e10)))
                                    _ = fi52
                                    _ = v53
                                }
                                switch (fi17 << (((^(v6))) & 63)) {
                                case 42:
                                    var v54 bool = true
                                    _ = v54
                                }
                            }
                            v6 = v6
                            _ = fi17
                        }
                    }
                }
            case -174:
                v11 = ("" + "\n\t")
            default:
                v55 := GenS5{F0: 42, F1: 413, F2: 0}
                _ = v55
            }
            if false {
                v56 := GenS5{F0: 127, F1: 0, F2: 0}
                var v57 float64 = 0.1
                _ = v56
                _ = v57
            }
        }
    default:
        fmt.Println(v11, v6, v3)
    }
    if true {
        v58 := GenS1{F0: (("a" + v1) >= (v1 + ((v1 + "\n\t") + "hello"))), F1: -128}
        switch int(v3) {
        case 0:
            var v59 float64 = (float64((-2.25 * (-2.25 / 1.5))) + (0.0 - (float64(0.1) - (1e10 + 3.14159))))
            _ = v59
        case -840:
            fmt.Println(v6, v6, v11)
        case -2147483648:
            v6 -= 628
        }
        _ = v58
    } else {
        fmt.Println(v6)
    }
    var v60 uint64 = 1
    for fi61 := 0; fi61 < 5; fi61++ {
        var v62 float32 = (float32(1.5) + 0.1)
        _ = fi61
        _ = v62
    }
    if (!(2 > 53)) {
        v6 += v6
    }
    for fi63 := 0; fi63 < 9; fi63++ {
        switch (fi63 | fi63) {
        case -832:
            v6 = (v6 ^ v6)
        case 127:
            v64 := GenS8{F0: fi63, F1: 42}
            _ = v64
        }
        _ = fi63
    }
    for fi65 := 0; fi65 < 1; fi65++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v66 := GenS3{F0: -915, F1: 1000, F2: 407}
        var v67 uint32 = 2147483648
        _ = fi65
        _ = v66
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
