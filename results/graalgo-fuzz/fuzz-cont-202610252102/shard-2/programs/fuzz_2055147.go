package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int64; F2 bool }
type GenS3 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int; F1 uint8; F2 int64 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS8 struct { F0 uint64 }
func genFn0() string {
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: 1}
            switch fi8 {
            case 32767:
                v10 := []int{}
                _ = v10
            case -2147483648:
                var v11 uint64 = 381
                _ = v11
            }
            var v12 float32 = 1.5
            _ = fi8
            _ = v12
            _ = v9
        }
        v13 := GenS2{F0: -155, F1: -9223372036854775808, F2: true}
        var v14 uint8 = 186
        _ = fi7
        _ = v13
        _ = v14
    }
    v15 := []int{32767}
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 uint64 = 0
        v18 := GenS2{F0: 1, F1: 42, F2: true}
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 string = ("foo bar" + "hello")
    if (3.14159 < float32((((-2.25 + float64(0.1)) / (((3.14159 + ((3.14159 * -2.25) / (3.14159 * float64(float32(-2.25))))) + 0.1) * 1e10)) - 1e10))) {
        var v20 uint8 = 1
        if true {
            var v21 int8 = 0
            if len(v15) > 0 {
                v15[0] = 32767
            }
            v20 -= v20
            _ = v21
        }
        _ = v20
    }
    _ = v15
    _ = v19
    return v19
}
func genFn1(p0 int, p1 uint16, p2 string) int64 {
    v22 := []int{42, 1, (p0 >> ((408) & 63))}
    switch "a" {
    case "\n\t":
        p1 = 32767
    case "":
        for fi23 := 0; fi23 < 8; fi23++ {
            if true {
                if true {
                    p2 = "hello"
                }
                p0 = p0
                v22 = append(v22, fi23)
            }
            if ((1e10 <= 0.1) || false) {
                v24 := GenS1{F0: 4294967295}
                if false {
                    v25 := GenS3{F0: (847 < 672), F1: 74, F2: 839}
                    _ = v25
                } else {
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    switch p2 {
                    case "\n\t":
                        p2 = ""
                    }
                    switch (p0 << ((309) & 63)) {
                    case 42:
                        p2 = "foo bar"
                    case 27:
                        if (!((float64(0.0) - 1e10) >= 1e10)) {
                            var v26 float32 = (0.1 - (1e10 + (float32((1.5 / ((-2.25 * float64(0.0)) + (-2.25 / 1.5)))) - 3.14159)))
                            var v27 uint8 = 2
                            for fi28 := 0; fi28 < 10; fi28++ {
                                var v29 uint64 = 414
                                _ = fi28
                                _ = v29
                            }
                            _ = v26
                            _ = v27
                        }
                    default:
                        var v30 uint64 = 2
                        _ = v30
                    }
                }
                _ = v24
            } else {
                fmt.Println(p2)
                var v31 int8 = -124
                var v32 uint64 = 545
                _ = v31
                _ = v32
            }
            var v33 int8 = 0
            _ = fi23
            _ = v33
        }
    case "a":
        fmt.Println(p2, p1, p0)
    }
    p0 -= p0
    v34 := GenS1{F0: 413}
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v34
    return 762
}

func main() {
    var v1 int64 = -1
    v2 := []string{"", "a", "foo bar", "a"}
    v3 := make(map[int]bool)
    var v4 string = (("foo bar" + ("\n\t" + "hello")) + (("\n\t" + "a") + (("" + "\n\t") + ("foo bar" + "a"))))
    var v5 string = ""
    v6 := make(map[uint32]uint32)
    v35 := GenS2{F0: -1, F1: genFn1(750, 63, (v5 + ("hello" + "foo bar"))), F2: ((!true) && (false && (0.0 >= 0.1)))}
    if false {
        var v36 float64 = -2.25
        fmt.Println(v5, v4)
        _ = v36
    } else {
        _ = len(v2)
        var v37 uint32 = 70
        var v38 uint64 = 4294967295
        _ = v37
        _ = v38
    }
    v39 := GenS4{F0: 128, F1: 92, F2: (v1 ^ -1)}
    v5 = "foo bar"
    switch -564 {
    case 47:
        if true {
            v1 = (v1 >> ((298) & 63))
        } else {
            for fi40 := 0; fi40 < 4; fi40++ {
                v4 = v4
                _ = fi40
            }
            if true {
                var v41 bool = (1 < 0)
                v4 = v4
                _ = v41
            } else {
                v5 = (v4 + v5)
                var v42 uint = 18446744073709551615
                var v43 uint32 = 674
                _ = v42
                _ = v43
            }
            _ = len(v2)
        }
    }
    if len(v2) > 0 {
        v2[0] = v5
    }
    if false {
        v4 = "foo bar"
        v44 := GenS5{F0: -27, F1: 32767}
        switch v4 {
        case "\n\t":
            v4 = ((v5 + (("foo bar" + "\n\t") + "")) + "foo bar")
        case "foo bar":
            for fi45 := 0; fi45 < 16; fi45++ {
                v46 := GenS2{F0: 706, F1: -1, F2: (!false)}
                if len(v2) > 0 {
                    v2[0] = (v5 + (v4 + "hello"))
                }
                _ = fi45
                _ = v46
            }
        case "":
            if false {
                var v47 int = 4294967295
                if (!(v1 > genFn1(4294967295, 371, "\n\t"))) {
                    if true {
                        var v48 uint16 = 1
                        _ = len(v2)
                        v47 += v47
                        _ = v48
                    }
                    var v49 float64 = 0.0
                    fmt.Println(v1)
                    _ = v49
                } else {
                    var v50 bool = false
                    _ = v50
                }
                _ = v47
            }
        }
        _ = v44
    } else {
        for fi51 := 0; fi51 < 16; fi51++ {
            fmt.Println(v5, v1, v5)
            if (false || true) {
                switch (v4 + v5) {
                case "foo bar":
                    v2 = append(v2, "\n\t")
                default:
                    for fi52 := 0; fi52 < 4; fi52++ {
                        var v53 uint8 = 0
                        v5 = "\n\t"
                        if (1 >= 0) {
                            if true {
                                _ = len(v2)
                                v54 := GenS6{F0: v4, F1: fi51}
                                switch v4 {
                                case "foo bar":
                                    var v55 uint16 = 2
                                    _ = v55
                                case "\n\t":
                                    var v56 uint16 = 2
                                    _ = v56
                                }
                                _ = v54
                            } else {
                                v5 = ("foo bar" + v5)
                            }
                            v4 = (v5 + "foo bar")
                            var v57 uint32 = 65535
                            v58 := GenS3{F0: (!false), F1: 2, F2: v57}
                            _ = v57
                            _ = v58
                        } else {
                            v59 := v6[540]
                            var v60 int16 = 0
                            if true {
                                v4 = (("a" + (v5 + "hello")) + ("hello" + v5))
                                fmt.Println(v59, v1, v53)
                            } else {
                                if false {
                                    var v61 int32 = -522
                                    _ = v61
                                }
                                if true {
                                    v4 = v4
                                }
                            }
                            v60 = v60
                            _ = v59
                            _ = v60
                        }
                        _ = fi52
                        _ = v53
                    }
                }
                v62 := v6[42]
                _ = v62
            } else {
                switch (v5 + v4) {
                case "":
                    if true {
                        v63 := GenS7{F0: 176, F1: 152, F2: 42}
                        v4 = "foo bar"
                        switch 32768 {
                        case -273:
                            var v64 float32 = 3.14159
                            _ = v64
                        default:
                            for fi65 := 0; fi65 < 17; fi65++ {
                                var v66 string = (v5 + "\n\t")
                                var v67 uint8 = 18
                                _ = len(v2)
                                var v68 int32 = -2147483648
                                _ = fi65
                                _ = v66
                                _ = v67
                                _ = v68
                            }
                        }
                        _ = v63
                    }
                case "foo bar":
                    v1 = -9223372036854775808
                default:
                    var v69 string = v4
                    _ = v69
                }
            }
            v70 := GenS5{F0: -97, F1: (fi51 * 127)}
            _ = fi51
            _ = v70
        }
    }
    for fi71 := 0; fi71 < 19; fi71++ {
        v72 := GenS8{F0: 1}
        _ = len(v2)
        fmt.Println(fi71, v4, fi71)
        _ = fi71
        _ = v72
    }
    var v73 uint64 = 1000
    var v74 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
}
