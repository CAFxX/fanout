package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint16; F2 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 int; F1 uint; F2 int8 }
type GenS9 struct { F0 string }
func genFn0(p0 int8) int32 {
    v7 := []int{}
    if (true || (-1 != -859)) {
        if ((p0 + p0) > 1) {
            if (42 == 42) {
                for fi8 := 0; fi8 < 9; fi8++ {
                    v9 := GenS1{F0: -2147483648, F1: -32768}
                    var v10 float64 = (0.1 * float64(0.1))
                    if (!(true || false)) {
                        switch fi8 {
                        case 32767:
                            switch (fi8 - -1) {
                            case -1:
                                var v11 int64 = 1
                                _ = v11
                            case 127:
                                for fi12 := 0; fi12 < 4; fi12++ {
                                    fmt.Println(v10)
                                    _ = fi12
                                }
                            }
                        default:
                            if true {
                                if len(v7) > 0 {
                                    v7[0] = fi8
                                }
                                var v13 uint8 = 1
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                _ = v13
                            } else {
                                fmt.Println(v10, p0)
                                if true {
                                    switch fi8 {
                                    case 850:
                                        var v14 int = 32768
                                        _ = v14
                                    case 738:
                                        var v15 float32 = -2.25
                                        _ = v15
                                    case 0:
                                        switch "" {
                                        case "hello":
                                            v16 := GenS2{F0: 2, F1: 42}
                                            _ = v16
                                        case "\n\t":
                                            for fi17 := 0; fi17 < 6; fi17++ {
                                                if (255 == 145) {
                                                    p0 = p0
                                                    var v18 string = ""
                                                    v19 := GenS3{F0: 9}
                                                    _ = v18
                                                    _ = v19
                                                }
                                                _ = fi17
                                            }
                                        case "a":
                                            var v20 int64 = 42
                                            _ = v20
                                        }
                                    default:
                                        var v21 int8 = (p0 - (p0 / (((p0 & p0)) | 1)))
                                        _ = v21
                                    }
                                }
                            }
                        }
                        var v22 string = "\n\t"
                        _ = v22
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                for fi23 := 0; fi23 < 7; fi23++ {
                    v24 := GenS4{F0: 210, F1: 0}
                    _ = fi23
                    _ = v24
                }
            } else {
                if false {
                    p0 += -83
                    switch "" {
                    case "a":
                        if false {
                            var v25 uint = 1
                            _ = v25
                        } else {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            p0 = -128
                            fmt.Println(p0)
                        }
                    default:
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    var v26 bool = (false || (false && (!(false && (-2.25 <= float32(0.1))))))
                    _ = v26
                } else {
                    var v27 string = "a"
                    v7 = append(v7, int(p0))
                    switch (v27 + ("foo bar" + "foo bar")) {
                    case "hello":
                        for fi28 := 0; fi28 < 9; fi28++ {
                            v29 := GenS5{F0: 24, F1: 948, F2: 42}
                            v30 := GenS5{F0: 18446744073709551615, F1: 2, F2: -2147483648}
                            for fi31 := 0; fi31 < 5; fi31++ {
                                switch v27 {
                                case "foo bar":
                                    if (!(0 <= 9223372036854775807)) {
                                        if (!(42 > 98)) {
                                            v27 = v27
                                            v32 := GenS4{F0: -2147483648, F1: 127}
                                            var v33 uint32 = 1
                                            fmt.Println(fi28, fi31)
                                            _ = v32
                                            _ = v33
                                        } else {
                                            var v34 string = v27
                                            _ = v34
                                        }
                                        var v35 uint64 = 18446744073709551615
                                        p0 = p0
                                        var v36 bool = (!false)
                                        _ = v35
                                        _ = v36
                                    }
                                case "hello":
                                    var v37 uint32 = 42
                                    _ = v37
                                default:
                                    var v38 uint16 = 928
                                    _ = v38
                                }
                                v39 := GenS6{F0: -846}
                                fmt.Println(v27)
                                _ = fi31
                                _ = v39
                            }
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    case "\n\t":
                        var v40 int8 = (-(p0))
                        _ = v40
                    case "":
                        var v41 uint = 420
                        _ = v41
                    default:
                        var v42 int32 = 2147483647
                        _ = v42
                    }
                    _ = v27
                }
            }
            v43 := GenS3{F0: -21}
            var v44 bool = true
            _ = v43
            _ = v44
        } else {
            switch 32768 {
            case 2147483647:
                switch "hello" {
                case "foo bar":
                    var v45 uint32 = 65535
                    _ = v45
                case "\n\t":
                    v46 := GenS7{F0: 65535, F1: 2}
                    _ = v46
                default:
                    var v47 bool = true
                    _ = v47
                }
            }
        }
    }
    var v48 uint32 = 868
    switch -2147483648 {
    case -2147483648:
        for fi49 := 0; fi49 < 11; fi49++ {
            var v50 float32 = 0.1
            if true {
                if true {
                    v48 += (v48 * (^(v48)))
                    p0 = p0
                    var v51 int8 = -128
                    _ = v51
                }
            } else {
                switch ("\n\t" + ("" + "")) {
                case "\n\t":
                    switch fi49 {
                    case -204:
                        if ("\n\t" != "foo bar") {
                            fmt.Println(v50, p0)
                        }
                    case 4294967295:
                        if len(v7) > 0 {
                            v7[0] = fi49
                        }
                    case 127:
                        switch 0 {
                        case 0:
                            _ = len(v7)
                        }
                    }
                case "hello":
                    for fi52 := 0; fi52 < 2; fi52++ {
                        fmt.Println(fi52)
                        if (false || (true || false)) {
                            if (0 <= 127) {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                v7 = append(v7, -752)
                                switch 32768 {
                                case 42:
                                    var v53 uint16 = 32767
                                    _ = v53
                                case -9223372036854775808:
                                    p0 = 14
                                case -1:
                                    v54 := GenS8{F0: (fi49 << ((4294967295) & 63)), F1: 101, F2: p0}
                                    _ = v54
                                }
                            }
                            fmt.Println(v50, fi49)
                        }
                        var v55 uint = 42
                        if (-32768 == 42) {
                            var v56 int = fi52
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = v56
                        } else {
                            var v57 float64 = -2.25
                            v58 := GenS2{F0: v48, F1: 127}
                            _ = v57
                            _ = v58
                        }
                        _ = fi52
                        _ = v55
                    }
                case "a":
                    fmt.Println(p0)
                default:
                    for fi59 := 0; fi59 < 15; fi59++ {
                        p0 -= -128
                        _ = fi59
                    }
                }
            }
            _ = fi49
            _ = v50
        }
    case -845:
        v7 = append(v7, -1)
    case 32768:
        v48 = v48
    default:
        var v60 int16 = -801
        _ = v60
    }
    _ = p0
    _ = v48
    _ = v7
    return 85
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]uint32)
    var v3 string = "\n\t"
    v4 := []string{("" + v3), "\n\t", ((v3 + (v3 + (("a" + v3) + "\n\t"))) + "\n\t")}
    v5 := make(map[int]string)
    var v6 uint8 = 52
    _ = len(v5)
    var v61 int8 = 33
    if (!((v6 >> ((0) & 63)) <= v6)) {
        v3 = ("\n\t" + "\n\t")
        fmt.Println(v3, v6, v61)
        _ = len(v5)
    }
    v1 = (v1 ^ 0)
    v6 = 0
    if ((42 != -446) || ((0.0 / ((1e10 * (0.1 * 0.0)) - (0.0 + 0.1))) >= 0.0)) {
        v61 = (v61 - -115)
        v6 = 127
    }
    var v62 uint16 = 684
    var v63 int32 = -2147483648
    if (true || true) {
        v62 -= 2
    } else {
        if (v1 > (v1 % (((^(v1))) | 1))) {
            if (true || (v63 == (-(v63)))) {
                var v64 int16 = -211
                for fi65 := 0; fi65 < 4; fi65++ {
                    var v66 int32 = (v63 / ((42) | 1))
                    _ = fi65
                    _ = v66
                }
                fmt.Println(v1, v64)
                _ = v64
            }
            v63 -= (v63 << ((65535) & 63))
        }
        if true {
            switch 127 {
            case 2147483647:
                fmt.Println(v1, v1)
            case 0:
                v67 := GenS9{F0: "hello"}
                _ = v67
            }
            v3 = "foo bar"
        } else {
            v68 := GenS9{F0: v3}
            var v69 int64 = 336
            _ = v68
            _ = v69
        }
    }
    v1 = (v1 * v1)
    delete(v2, (^(v1)))
    v63 = v63
    v1 += (v1 ^ v1)
    v61 = v61
    v3 = v3
    if false {
        fmt.Println(v62, v61)
    }
    var v70 float32 = -2.25
    var v71 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v63
    _ = v70
    _ = v71
}
