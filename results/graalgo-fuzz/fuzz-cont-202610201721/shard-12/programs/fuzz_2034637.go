package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int32 }
func genFn0(p0 int16, p1 uint64, p2 uint32) int8 {
    var v7 uint = uint(p2)
    if false {
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(p0, v7, v7)
            for fi9 := 0; fi9 < 11; fi9++ {
                v10 := GenS1{F0: 1, F1: 0, F2: 9223372036854775807}
                if false {
                    v7 = 1
                }
                _ = fi9
                _ = v10
            }
            v11 := GenS2{F0: (p0 % ((p0) | 1))}
            _ = fi8
            _ = v11
        }
    }
    v7 += v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -128
}
func genFn1(p0 uint16) int {
    p0 = (p0 - (p0 ^ 32767))
    var v12 int16 = 0
    var v13 float64 = 3.14159
    _ = p0
    _ = v12
    _ = v13
    return -588
}

func main() {
    v1 := []uint32{1, 1, 1, 0}
    var v2 int64 = 0
    var v3 uint = 953
    v4 := []uint32{65535, 4294967295, 42, 2}
    var v5 bool = false
    v6 := make(map[uint32]int)
    var v14 uint8 = 1
    if (v5 && (v2 > v2)) {
        switch -1 {
        case 128:
            for fi15 := 0; fi15 < 11; fi15++ {
                if (!((v5 && v5) && false)) {
                    if (1 != 405) {
                        if v5 {
                            var v16 bool = (0 < 1)
                            for fi17 := 0; fi17 < 18; fi17++ {
                                v18 := GenS1{F0: 2, F1: 115, F2: v2}
                                _ = fi17
                                _ = v18
                            }
                            _ = v16
                        } else {
                            switch "\n\t" {
                            case "":
                                switch fi15 {
                                case -690:
                                    fmt.Println(v3, v5, v14)
                                }
                            default:
                                v19 := GenS1{F0: 32767, F1: 1, F2: v2}
                                _ = v19
                            }
                            fmt.Println(fi15)
                        }
                    } else {
                        fmt.Println(v14)
                    }
                }
                if ((!(!(!(0.0 == 1.5)))) && (!v5)) {
                    v3 -= v3
                    if v5 {
                        var v20 int8 = 42
                        if ((!(0 != -1)) && true) {
                            delete(v6, 117)
                        } else {
                            var v21 bool = (v5 || (545 > 331))
                            _ = v21
                        }
                        _ = v20
                    }
                    v14 = v14
                }
                _ = fi15
            }
        }
        for fi22 := 0; fi22 < 19; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 float64 = 0.1
                v24 -= float64(0.0)
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    }
    fmt.Println(v14, v3)
    v25 := GenS3{F0: ("foo bar" + "hello"), F1: 42}
    var v26 uint32 = 0
    if v5 {
        switch 563 {
        case 1:
            switch ("\n\t" + "a") {
            case "a":
                if v5 {
                    for fi27 := 0; fi27 < 9; fi27++ {
                        var v28 uint16 = 1
                        _ = fi27
                        _ = v28
                    }
                    if false {
                        v29 := GenS3{F0: "\n\t", F1: -1}
                        _ = v29
                    } else {
                        _ = len(v6)
                        switch 32768 {
                        case 128:
                            v30 := GenS1{F0: 2, F1: (v26 << ((v3) & 63)), F2: 979}
                            _ = v30
                        }
                    }
                    v31 := v6[v26]
                    _ = v31
                }
            case "\n\t":
                v32 := GenS3{F0: (("" + "\n\t") + ("foo bar" + "hello")), F1: -550}
                _ = v32
            case "":
                v33 := v6[(v26 >> (((^(v3))) & 63))]
                _ = v33
            }
        case -407:
            v34 := v6[v26]
            _ = v34
        }
    } else {
        for fi35 := 0; fi35 < 18; fi35++ {
            var v36 int64 = (v2 ^ (-(v2)))
            _ = fi35
            _ = v36
        }
        v14 = 1
    }
    var v37 int64 = v2
    fmt.Println(v14, v3, v2)
    if v5 {
        v38 := v6[(v26 >> ((v3) & 63))]
        _ = v38
    }
    var v39 int64 = v2
    fmt.Println(v37, v5, v39)
    var v40 float64 = ((float64(0.1) - -2.25) / (1e10 + (1e10 * 0.1)))
    for fi41 := 0; fi41 < 14; fi41++ {
        v42 := GenS1{F0: 2, F1: (v26 >> (((v3 - v3)) & 63)), F2: 0}
        _ = fi41
        _ = v42
    }
    v6[v26] = -537
    v2 += v39
    var v43 float32 = 0.0
    v44 := GenS3{F0: "foo bar", F1: -2147483648}
    fmt.Println(v43, v14, v43)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
